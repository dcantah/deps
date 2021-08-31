package deps

import (
	"fmt"

	api "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	runtime "k8s.io/cri-api/pkg/apis/runtime/v1"

	"k8s.io/klog/v2"
)

type typetest struct {
	uid types.UID
}

func Deps() {
	klog.V(5).Infof("Test")
	fmt.Println(api.StreamTypeError)
	fmt.Println(typetest{})
	fmt.Println(&runtime.AuthConfig{})
}
