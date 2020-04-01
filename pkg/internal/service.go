package internal

import (
	"github.com/weaveworks/flagger/pkg/apis/flagger/v1beta1"
	"strings"
)

// the value should be seperated by "," , labels combination should distinguish canary and primary workloads.
// these labels used to select canary and primary as seperated workloads.
const ALICLOUD_CANARY_DISTINGUISH_LABELS = "alicloud.canary.distinguish.labels"

// the value should be seperated by "," , labels combination should be same for canary and primary workloads.
// this labels used to select canary and primary as a whole app.
const ALICLOUD_CANARY_GENERAL_LABELS = "alicloud.canary.general.labels"

// check canary distinguish labels existing, if existed, return true and labels slice, otherwise return false and nil slice.
func CanaryDistinguishLabelsExisted(canary *v1beta1.Canary) ([]string, bool) {
	labels, exist := canary.Annotations[ALICLOUD_CANARY_DISTINGUISH_LABELS]
	if !exist || labels == "" {
		return nil, false
	}
	labelsSlice := strings.Split(labels, ",")
	if len(labelsSlice) == 0 {
		return nil, false
	}
	return labelsSlice, true
}

func CanaryGeneralLabelsExisted(canary *v1beta1.Canary) ([]string, bool) {
	labels, exist := canary.Annotations[ALICLOUD_CANARY_GENERAL_LABELS]
	if !exist || labels == "" {
		return nil, false
	}
	labelsSlice := strings.Split(labels, ",")
	if len(labelsSlice) == 0 {
		return nil, false
	}
	return labelsSlice, true
}

