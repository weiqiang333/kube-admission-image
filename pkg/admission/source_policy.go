package admission

import (
	"fmt"

	"github.com/spf13/viper"
	"k8s.io/api/imagepolicy/v1alpha1"

	"github.com/weiqiang333/kube-admission-image/pkg/method"
)

// UnauthorizedSourcePolicy
// response: allow bool, reason string, error
func UnauthorizedSourcePolicy(imageReview v1alpha1.ImageReview) (bool, string, error) {
	allow := true

	// 拒绝/允许 指定来源
	if allow, reason, err := designatedSourcePolicy(imageReview); !allow || err != nil {
		return allow, reason, err
	}

	return allow, "", nil
}

func designatedSourcePolicy(imageReview v1alpha1.ImageReview) (bool, string, error) {
	allow := true
	sourceDefaultPolicy := viper.GetString("sourceDefaultPolicy")
	sourceAllowPolicy := viper.GetStringSlice("sourceAllowPolicy")
	sourceRejectPolicy := viper.GetStringSlice("sourceRejectPolicy")

	for _, container := range imageReview.Spec.Containers {
		domain, _ := method.SplitDockerDomain(container.Image)
		if method.FindsStringSlice(sourceRejectPolicy, domain) {
			allow = false
			reason := fmt.Sprintf("UnauthorizedSourcePolicy, %s violation designatedSource %v", domain, sourceRejectPolicy)
			return allow, reason, nil
		}
		if sourceDefaultPolicy == "reject" {
			if !method.FindsStringSlice(sourceAllowPolicy, domain) {
				allow = false
				reason := fmt.Sprintf("UnauthorizedSourcePolicy, %s does not exist designatedSource %v", domain, sourceAllowPolicy)
				return allow, reason, nil
			}
		}
	}

	return allow, "", nil
}
