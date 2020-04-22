package admission

import (
	"fmt"
	"log"
	"strings"

	"github.com/containers/image/docker/reference"
	"k8s.io/api/imagepolicy/v1alpha1"
)

// NamePolicy
func NamePolicy(imageReview v1alpha1.ImageReview) (bool, error) {
	allow := true

	// 拒绝使用 latest version images
	// 并对不遵守 https://github.com/containers/image/pull/220 container 规范化进行警示, 不拒绝
	if allow, err := policyLatestTag(imageReview); !allow {
		return allow, err
	}

	return allow, nil
}

// policyLatestTag
func policyLatestTag(imageReview v1alpha1.ImageReview) (bool, error) {
	allow := true
	for _, container := range imageReview.Spec.Containers {
		usingLatest, err := isUsingLatestTag(container.Image)
		if err != nil {
			log.Printf("Error while parsing image name: %+v", err)
			return allow, fmt.Errorf("Error while parsing image name: %+v", err)
		}
		if usingLatest {
			allow = false
			return allow, fmt.Errorf("Images using latest tag are not allowed")
		}
	}
	return allow, nil
}

func isUsingLatestTag(image string) (bool, error) {
	named, err := reference.ParseNormalizedNamed(image)
	if err != nil {
		return false, err
	}

	return strings.HasSuffix(reference.TagNameOnly(named).String(), ":latest"), nil
}
