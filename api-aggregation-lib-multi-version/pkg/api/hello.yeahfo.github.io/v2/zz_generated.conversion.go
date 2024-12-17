//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
// Code generated by conversion-gen. DO NOT EDIT.

package v2

import (
	unsafe "unsafe"

	helloyeahfogithubio "github.com/yeahfo/cloud-native-tour/api-aggregation-lib-multi-version/pkg/api/hello.yeahfo.github.io"
	helloyeahfogithubiov2 "github.com/yeahfo/cloud-native-tour/api/hello.yeahfo.github.io/v2"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*helloyeahfogithubiov2.Foo)(nil), (*helloyeahfogithubio.Foo)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v2_Foo_To_helloyeahfogithubio_Foo(a.(*helloyeahfogithubiov2.Foo), b.(*helloyeahfogithubio.Foo), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*helloyeahfogithubio.Foo)(nil), (*helloyeahfogithubiov2.Foo)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_helloyeahfogithubio_Foo_To_v2_Foo(a.(*helloyeahfogithubio.Foo), b.(*helloyeahfogithubiov2.Foo), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*helloyeahfogithubiov2.FooCondition)(nil), (*helloyeahfogithubio.FooCondition)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v2_FooCondition_To_helloyeahfogithubio_FooCondition(a.(*helloyeahfogithubiov2.FooCondition), b.(*helloyeahfogithubio.FooCondition), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*helloyeahfogithubio.FooCondition)(nil), (*helloyeahfogithubiov2.FooCondition)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_helloyeahfogithubio_FooCondition_To_v2_FooCondition(a.(*helloyeahfogithubio.FooCondition), b.(*helloyeahfogithubiov2.FooCondition), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*helloyeahfogithubiov2.FooConfig)(nil), (*helloyeahfogithubio.FooConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v2_FooConfig_To_helloyeahfogithubio_FooConfig(a.(*helloyeahfogithubiov2.FooConfig), b.(*helloyeahfogithubio.FooConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*helloyeahfogithubio.FooConfig)(nil), (*helloyeahfogithubiov2.FooConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_helloyeahfogithubio_FooConfig_To_v2_FooConfig(a.(*helloyeahfogithubio.FooConfig), b.(*helloyeahfogithubiov2.FooConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*helloyeahfogithubiov2.FooList)(nil), (*helloyeahfogithubio.FooList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v2_FooList_To_helloyeahfogithubio_FooList(a.(*helloyeahfogithubiov2.FooList), b.(*helloyeahfogithubio.FooList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*helloyeahfogithubio.FooList)(nil), (*helloyeahfogithubiov2.FooList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_helloyeahfogithubio_FooList_To_v2_FooList(a.(*helloyeahfogithubio.FooList), b.(*helloyeahfogithubiov2.FooList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*helloyeahfogithubiov2.FooSpec)(nil), (*helloyeahfogithubio.FooSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v2_FooSpec_To_helloyeahfogithubio_FooSpec(a.(*helloyeahfogithubiov2.FooSpec), b.(*helloyeahfogithubio.FooSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*helloyeahfogithubio.FooSpec)(nil), (*helloyeahfogithubiov2.FooSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_helloyeahfogithubio_FooSpec_To_v2_FooSpec(a.(*helloyeahfogithubio.FooSpec), b.(*helloyeahfogithubiov2.FooSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*helloyeahfogithubiov2.FooStatus)(nil), (*helloyeahfogithubio.FooStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v2_FooStatus_To_helloyeahfogithubio_FooStatus(a.(*helloyeahfogithubiov2.FooStatus), b.(*helloyeahfogithubio.FooStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*helloyeahfogithubio.FooStatus)(nil), (*helloyeahfogithubiov2.FooStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_helloyeahfogithubio_FooStatus_To_v2_FooStatus(a.(*helloyeahfogithubio.FooStatus), b.(*helloyeahfogithubiov2.FooStatus), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v2_Foo_To_helloyeahfogithubio_Foo(in *helloyeahfogithubiov2.Foo, out *helloyeahfogithubio.Foo, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v2_FooSpec_To_helloyeahfogithubio_FooSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v2_FooStatus_To_helloyeahfogithubio_FooStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v2_Foo_To_helloyeahfogithubio_Foo is an autogenerated conversion function.
func Convert_v2_Foo_To_helloyeahfogithubio_Foo(in *helloyeahfogithubiov2.Foo, out *helloyeahfogithubio.Foo, s conversion.Scope) error {
	return autoConvert_v2_Foo_To_helloyeahfogithubio_Foo(in, out, s)
}

func autoConvert_helloyeahfogithubio_Foo_To_v2_Foo(in *helloyeahfogithubio.Foo, out *helloyeahfogithubiov2.Foo, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_helloyeahfogithubio_FooSpec_To_v2_FooSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_helloyeahfogithubio_FooStatus_To_v2_FooStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_helloyeahfogithubio_Foo_To_v2_Foo is an autogenerated conversion function.
func Convert_helloyeahfogithubio_Foo_To_v2_Foo(in *helloyeahfogithubio.Foo, out *helloyeahfogithubiov2.Foo, s conversion.Scope) error {
	return autoConvert_helloyeahfogithubio_Foo_To_v2_Foo(in, out, s)
}

func autoConvert_v2_FooCondition_To_helloyeahfogithubio_FooCondition(in *helloyeahfogithubiov2.FooCondition, out *helloyeahfogithubio.FooCondition, s conversion.Scope) error {
	out.Type = helloyeahfogithubio.FooConditionType(in.Type)
	out.Status = v1.ConditionStatus(in.Status)
	return nil
}

// Convert_v2_FooCondition_To_helloyeahfogithubio_FooCondition is an autogenerated conversion function.
func Convert_v2_FooCondition_To_helloyeahfogithubio_FooCondition(in *helloyeahfogithubiov2.FooCondition, out *helloyeahfogithubio.FooCondition, s conversion.Scope) error {
	return autoConvert_v2_FooCondition_To_helloyeahfogithubio_FooCondition(in, out, s)
}

func autoConvert_helloyeahfogithubio_FooCondition_To_v2_FooCondition(in *helloyeahfogithubio.FooCondition, out *helloyeahfogithubiov2.FooCondition, s conversion.Scope) error {
	out.Type = helloyeahfogithubiov2.FooConditionType(in.Type)
	out.Status = corev1.ConditionStatus(in.Status)
	return nil
}

// Convert_helloyeahfogithubio_FooCondition_To_v2_FooCondition is an autogenerated conversion function.
func Convert_helloyeahfogithubio_FooCondition_To_v2_FooCondition(in *helloyeahfogithubio.FooCondition, out *helloyeahfogithubiov2.FooCondition, s conversion.Scope) error {
	return autoConvert_helloyeahfogithubio_FooCondition_To_v2_FooCondition(in, out, s)
}

func autoConvert_v2_FooConfig_To_helloyeahfogithubio_FooConfig(in *helloyeahfogithubiov2.FooConfig, out *helloyeahfogithubio.FooConfig, s conversion.Scope) error {
	out.Msg = in.Msg
	out.Description = in.Description
	return nil
}

// Convert_v2_FooConfig_To_helloyeahfogithubio_FooConfig is an autogenerated conversion function.
func Convert_v2_FooConfig_To_helloyeahfogithubio_FooConfig(in *helloyeahfogithubiov2.FooConfig, out *helloyeahfogithubio.FooConfig, s conversion.Scope) error {
	return autoConvert_v2_FooConfig_To_helloyeahfogithubio_FooConfig(in, out, s)
}

func autoConvert_helloyeahfogithubio_FooConfig_To_v2_FooConfig(in *helloyeahfogithubio.FooConfig, out *helloyeahfogithubiov2.FooConfig, s conversion.Scope) error {
	out.Msg = in.Msg
	out.Description = in.Description
	return nil
}

// Convert_helloyeahfogithubio_FooConfig_To_v2_FooConfig is an autogenerated conversion function.
func Convert_helloyeahfogithubio_FooConfig_To_v2_FooConfig(in *helloyeahfogithubio.FooConfig, out *helloyeahfogithubiov2.FooConfig, s conversion.Scope) error {
	return autoConvert_helloyeahfogithubio_FooConfig_To_v2_FooConfig(in, out, s)
}

func autoConvert_v2_FooList_To_helloyeahfogithubio_FooList(in *helloyeahfogithubiov2.FooList, out *helloyeahfogithubio.FooList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]helloyeahfogithubio.Foo)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v2_FooList_To_helloyeahfogithubio_FooList is an autogenerated conversion function.
func Convert_v2_FooList_To_helloyeahfogithubio_FooList(in *helloyeahfogithubiov2.FooList, out *helloyeahfogithubio.FooList, s conversion.Scope) error {
	return autoConvert_v2_FooList_To_helloyeahfogithubio_FooList(in, out, s)
}

func autoConvert_helloyeahfogithubio_FooList_To_v2_FooList(in *helloyeahfogithubio.FooList, out *helloyeahfogithubiov2.FooList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]helloyeahfogithubiov2.Foo)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_helloyeahfogithubio_FooList_To_v2_FooList is an autogenerated conversion function.
func Convert_helloyeahfogithubio_FooList_To_v2_FooList(in *helloyeahfogithubio.FooList, out *helloyeahfogithubiov2.FooList, s conversion.Scope) error {
	return autoConvert_helloyeahfogithubio_FooList_To_v2_FooList(in, out, s)
}

func autoConvert_v2_FooSpec_To_helloyeahfogithubio_FooSpec(in *helloyeahfogithubiov2.FooSpec, out *helloyeahfogithubio.FooSpec, s conversion.Scope) error {
	out.Image = in.Image
	if err := Convert_v2_FooConfig_To_helloyeahfogithubio_FooConfig(&in.Config, &out.Config, s); err != nil {
		return err
	}
	return nil
}

// Convert_v2_FooSpec_To_helloyeahfogithubio_FooSpec is an autogenerated conversion function.
func Convert_v2_FooSpec_To_helloyeahfogithubio_FooSpec(in *helloyeahfogithubiov2.FooSpec, out *helloyeahfogithubio.FooSpec, s conversion.Scope) error {
	return autoConvert_v2_FooSpec_To_helloyeahfogithubio_FooSpec(in, out, s)
}

func autoConvert_helloyeahfogithubio_FooSpec_To_v2_FooSpec(in *helloyeahfogithubio.FooSpec, out *helloyeahfogithubiov2.FooSpec, s conversion.Scope) error {
	out.Image = in.Image
	if err := Convert_helloyeahfogithubio_FooConfig_To_v2_FooConfig(&in.Config, &out.Config, s); err != nil {
		return err
	}
	return nil
}

// Convert_helloyeahfogithubio_FooSpec_To_v2_FooSpec is an autogenerated conversion function.
func Convert_helloyeahfogithubio_FooSpec_To_v2_FooSpec(in *helloyeahfogithubio.FooSpec, out *helloyeahfogithubiov2.FooSpec, s conversion.Scope) error {
	return autoConvert_helloyeahfogithubio_FooSpec_To_v2_FooSpec(in, out, s)
}

func autoConvert_v2_FooStatus_To_helloyeahfogithubio_FooStatus(in *helloyeahfogithubiov2.FooStatus, out *helloyeahfogithubio.FooStatus, s conversion.Scope) error {
	out.Phase = helloyeahfogithubio.FooPhase(in.Phase)
	out.Conditions = *(*[]helloyeahfogithubio.FooCondition)(unsafe.Pointer(&in.Conditions))
	return nil
}

// Convert_v2_FooStatus_To_helloyeahfogithubio_FooStatus is an autogenerated conversion function.
func Convert_v2_FooStatus_To_helloyeahfogithubio_FooStatus(in *helloyeahfogithubiov2.FooStatus, out *helloyeahfogithubio.FooStatus, s conversion.Scope) error {
	return autoConvert_v2_FooStatus_To_helloyeahfogithubio_FooStatus(in, out, s)
}

func autoConvert_helloyeahfogithubio_FooStatus_To_v2_FooStatus(in *helloyeahfogithubio.FooStatus, out *helloyeahfogithubiov2.FooStatus, s conversion.Scope) error {
	out.Phase = helloyeahfogithubiov2.FooPhase(in.Phase)
	out.Conditions = *(*[]helloyeahfogithubiov2.FooCondition)(unsafe.Pointer(&in.Conditions))
	return nil
}

// Convert_helloyeahfogithubio_FooStatus_To_v2_FooStatus is an autogenerated conversion function.
func Convert_helloyeahfogithubio_FooStatus_To_v2_FooStatus(in *helloyeahfogithubio.FooStatus, out *helloyeahfogithubiov2.FooStatus, s conversion.Scope) error {
	return autoConvert_helloyeahfogithubio_FooStatus_To_v2_FooStatus(in, out, s)
}