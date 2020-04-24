package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"k8s.io/api/imagepolicy/v1alpha1"

	"github.com/weiqiang333/kube-admission-image/pkg/admission"
)

// ImagesAdmission images 决策入口
func ImagesAdmission(c *gin.Context) {
	var imageReview v1alpha1.ImageReview
	var review v1alpha1.ImageReview

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Printf("body read fail: %v", err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	err = json.Unmarshal(body, &imageReview)
	if err != nil {
		log.Printf("body json Unmarshal fail: %v", err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	admissionPolicy(c, imageReview, review)
}

func admissionPolicy(c *gin.Context, imageReview, review v1alpha1.ImageReview) {
	review.APIVersion = "imagepolicy.k8s.io/v1alpha1"
	review.Kind = "ImageReview"
	review.Status.Allowed = true

	// NamePolicy
	allow, reason, err := admission.NamePolicy(imageReview)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if !allow {
		review.Status.Allowed = allow
		review.Status.Reason = reason
		c.JSON(http.StatusOK, review)
		return
	}

	// UnauthorizedSourcePolicy
	allow, reason, err = admission.UnauthorizedSourcePolicy(imageReview)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if !allow {
		review.Status.Allowed = allow
		review.Status.Reason = reason
		c.JSON(http.StatusOK, review)
		return
	}

	c.JSON(http.StatusOK, review)
}
