package admission

import (
	"fmt"
	"log"
	"strings"

	"github.com/spf13/viper"
	"github.com/weiqiang333/kube-admission-image/pkg/method"

	"github.com/containers/image/docker/reference"
	"k8s.io/api/imagepolicy/v1alpha1"
)

// NamePolicy
// response: allow bool, reason string, error
func NamePolicy(imageReview v1alpha1.ImageReview) (bool, string, error) {
	allow := true
	nameRejectPolicy := viper.GetStringSlice("nameRejectPolicy")

	// 拒绝使用 latest version images
	// 并对不遵守 https://github.com/containers/image/pull/220 container 规范化进行警示, 不拒绝
	if method.FindsStringSlice(nameRejectPolicy, "latestTag") {
		if allow, reason, err := policyLatestTag(imageReview); !allow || err != nil {
			return allow, reason, err
		}
	}

	return allow, "", nil
}

// policyLatestTag
func policyLatestTag(imageReview v1alpha1.ImageReview) (bool, string, error) {
	allow := true
	for _, container := range imageReview.Spec.Containers {
		usingLatest, err := isUsingLatestTag(container.Image)
		if err != nil {
			log.Printf("Error while parsing image name (%v): %+v", container.Image, err)
			return allow, "", fmt.Errorf("Error while parsing image name (%v): %+v", container.Image, err)
		}
		if usingLatest {
			allow = false
			log.Printf("Images using latest tag are not allowed: %s in namespace %s", container.Image, imageReview.Spec.Namespace)
			return allow, "Images using latest tag are not allowed", nil
		}
	}
	return allow, "", nil
}

func isUsingLatestTag(image string) (bool, error) {
	named, err := reference.ParseNormalizedNamed(image)
	if err != nil {
		return false, err
	}

	return strings.HasSuffix(reference.TagNameOnly(named).String(), ":latest"), nil
}
