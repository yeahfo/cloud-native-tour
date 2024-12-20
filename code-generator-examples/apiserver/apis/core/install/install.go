package install

import (
	"github.com/yeahfo/cloud-native-tour/code-generator-examples/apiserver/apis/core"
	v1 "github.com/yeahfo/cloud-native-tour/code-generator-examples/apiserver/apis/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
)

func Install(scheme *runtime.Scheme) {
	utilruntime.Must(core.AddToScheme(scheme))
	utilruntime.Must(v1.AddToScheme(scheme))
	utilruntime.Must(scheme.SetVersionPriority(v1.SchemeGroupVersion))
}
