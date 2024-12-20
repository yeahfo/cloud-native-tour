package v1

import (
	helloyeahfogithubio "github.com/yeahfo/cloud-native-tour/api-aggregation-lib-multi-version/pkg/api/hello.yeahfo.github.io"
	hellov1 "github.com/yeahfo/cloud-native-tour/api/hello.yeahfo.github.io/v1"
	"k8s.io/apimachinery/pkg/conversion"
)

// Convert_helloyeahfogithubio_Foo_To_v1_Foo ...
//
//goland:noinspection GoSnakeCaseUsage
func Convert_helloyeahfogithubio_Foo_To_v1_Foo(in *helloyeahfogithubio.Foo, out *hellov1.Foo, scope conversion.Scope) error {
	// call func generated by conversion-gen
	// autoConvert_helloyeahfogithubio_Foo_To_v1_Foo will call Convert_helloyeahfogithubio_FooSpec_To_v1_FooSpec
	if err := autoConvert_helloyeahfogithubio_Foo_To_v1_Foo(in, out, scope); err != nil {
		return err
	}
	// do conversion here
	if out.Annotations == nil {
		out.Annotations = map[string]string{}
	}
	out.Annotations[AnnotationImage] = in.Spec.Image
	return nil
}

// Convert_helloyeahfogithubio_FooSpec_To_v1_FooSpec
//
//goland:noinspection GoSnakeCaseUsage
func Convert_helloyeahfogithubio_FooSpec_To_v1_FooSpec(in *helloyeahfogithubio.FooSpec, out *hellov1.FooSpec, _ conversion.Scope) error {
	out.Msg = in.Config.Msg
	out.Description = in.Config.Description
	return nil
}

// Convert_v1_Foo_To_helloyeahfogithubio_Foo
//
//goland:noinspection GoSnakeCaseUsage
func Convert_v1_Foo_To_helloyeahfogithubio_Foo(in *hellov1.Foo, out *helloyeahfogithubio.Foo, scope conversion.Scope) error {
	// call func generated by conversion-gen
	// autoConvert_v1_Foo_To_helloyeahfogithubio_Foo will call Convert_v1_FooSpec_To_helloyeahfogithubio_FooSpec
	if err := autoConvert_v1_Foo_To_helloyeahfogithubio_Foo(in, out, scope); err != nil {
		return err
	}
	// do conversion here
	out.Spec.Image = in.Annotations[AnnotationImage]
	return nil
}

// Convert_v1_FooSpec_To_helloyeahfogithubio_FooSpec
//
//goland:noinspection GoSnakeCaseUsage
func Convert_v1_FooSpec_To_helloyeahfogithubio_FooSpec(in *hellov1.FooSpec, out *helloyeahfogithubio.FooSpec, scope conversion.Scope) error {
	out.Config.Msg = in.Msg
	out.Config.Description = in.Description
	return nil
}
