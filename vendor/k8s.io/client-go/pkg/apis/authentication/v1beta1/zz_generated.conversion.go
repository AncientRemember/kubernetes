// +build !ignore_autogenerated

/*
Copyright 2016 The Kubernetes Authors.

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

// This file was autogenerated by conversion-gen. Do not edit it manually!

package v1beta1

import (
	api "k8s.io/client-go/pkg/api"
	authentication "k8s.io/client-go/pkg/apis/authentication"
	conversion "k8s.io/client-go/pkg/conversion"
	runtime "k8s.io/client-go/pkg/runtime"
)

func init() {
	SchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedConversionFuncs(
		Convert_v1beta1_TokenReview_To_authentication_TokenReview,
		Convert_authentication_TokenReview_To_v1beta1_TokenReview,
		Convert_v1beta1_TokenReviewSpec_To_authentication_TokenReviewSpec,
		Convert_authentication_TokenReviewSpec_To_v1beta1_TokenReviewSpec,
		Convert_v1beta1_TokenReviewStatus_To_authentication_TokenReviewStatus,
		Convert_authentication_TokenReviewStatus_To_v1beta1_TokenReviewStatus,
		Convert_v1beta1_UserInfo_To_authentication_UserInfo,
		Convert_authentication_UserInfo_To_v1beta1_UserInfo,
	)
}

func autoConvert_v1beta1_TokenReview_To_authentication_TokenReview(in *TokenReview, out *authentication.TokenReview, s conversion.Scope) error {
	if err := api.Convert_unversioned_TypeMeta_To_unversioned_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	// TODO: Inefficient conversion - can we improve it?
	if err := s.Convert(&in.ObjectMeta, &out.ObjectMeta, 0); err != nil {
		return err
	}
	if err := Convert_v1beta1_TokenReviewSpec_To_authentication_TokenReviewSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1beta1_TokenReviewStatus_To_authentication_TokenReviewStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

func Convert_v1beta1_TokenReview_To_authentication_TokenReview(in *TokenReview, out *authentication.TokenReview, s conversion.Scope) error {
	return autoConvert_v1beta1_TokenReview_To_authentication_TokenReview(in, out, s)
}

func autoConvert_authentication_TokenReview_To_v1beta1_TokenReview(in *authentication.TokenReview, out *TokenReview, s conversion.Scope) error {
	if err := api.Convert_unversioned_TypeMeta_To_unversioned_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	// TODO: Inefficient conversion - can we improve it?
	if err := s.Convert(&in.ObjectMeta, &out.ObjectMeta, 0); err != nil {
		return err
	}
	if err := Convert_authentication_TokenReviewSpec_To_v1beta1_TokenReviewSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_authentication_TokenReviewStatus_To_v1beta1_TokenReviewStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

func Convert_authentication_TokenReview_To_v1beta1_TokenReview(in *authentication.TokenReview, out *TokenReview, s conversion.Scope) error {
	return autoConvert_authentication_TokenReview_To_v1beta1_TokenReview(in, out, s)
}

func autoConvert_v1beta1_TokenReviewSpec_To_authentication_TokenReviewSpec(in *TokenReviewSpec, out *authentication.TokenReviewSpec, s conversion.Scope) error {
	out.Token = in.Token
	return nil
}

func Convert_v1beta1_TokenReviewSpec_To_authentication_TokenReviewSpec(in *TokenReviewSpec, out *authentication.TokenReviewSpec, s conversion.Scope) error {
	return autoConvert_v1beta1_TokenReviewSpec_To_authentication_TokenReviewSpec(in, out, s)
}

func autoConvert_authentication_TokenReviewSpec_To_v1beta1_TokenReviewSpec(in *authentication.TokenReviewSpec, out *TokenReviewSpec, s conversion.Scope) error {
	out.Token = in.Token
	return nil
}

func Convert_authentication_TokenReviewSpec_To_v1beta1_TokenReviewSpec(in *authentication.TokenReviewSpec, out *TokenReviewSpec, s conversion.Scope) error {
	return autoConvert_authentication_TokenReviewSpec_To_v1beta1_TokenReviewSpec(in, out, s)
}

func autoConvert_v1beta1_TokenReviewStatus_To_authentication_TokenReviewStatus(in *TokenReviewStatus, out *authentication.TokenReviewStatus, s conversion.Scope) error {
	out.Authenticated = in.Authenticated
	if err := Convert_v1beta1_UserInfo_To_authentication_UserInfo(&in.User, &out.User, s); err != nil {
		return err
	}
	out.Error = in.Error
	return nil
}

func Convert_v1beta1_TokenReviewStatus_To_authentication_TokenReviewStatus(in *TokenReviewStatus, out *authentication.TokenReviewStatus, s conversion.Scope) error {
	return autoConvert_v1beta1_TokenReviewStatus_To_authentication_TokenReviewStatus(in, out, s)
}

func autoConvert_authentication_TokenReviewStatus_To_v1beta1_TokenReviewStatus(in *authentication.TokenReviewStatus, out *TokenReviewStatus, s conversion.Scope) error {
	out.Authenticated = in.Authenticated
	if err := Convert_authentication_UserInfo_To_v1beta1_UserInfo(&in.User, &out.User, s); err != nil {
		return err
	}
	out.Error = in.Error
	return nil
}

func Convert_authentication_TokenReviewStatus_To_v1beta1_TokenReviewStatus(in *authentication.TokenReviewStatus, out *TokenReviewStatus, s conversion.Scope) error {
	return autoConvert_authentication_TokenReviewStatus_To_v1beta1_TokenReviewStatus(in, out, s)
}

func autoConvert_v1beta1_UserInfo_To_authentication_UserInfo(in *UserInfo, out *authentication.UserInfo, s conversion.Scope) error {
	out.Username = in.Username
	out.UID = in.UID
	out.Groups = in.Groups
	if in.Extra != nil {
		in, out := &in.Extra, &out.Extra
		*out = make(map[string]authentication.ExtraValue, len(*in))
		for key, val := range *in {
			newVal := new(authentication.ExtraValue)
			// TODO: Inefficient conversion - can we improve it?
			if err := s.Convert(&val, newVal, 0); err != nil {
				return err
			}
			(*out)[key] = *newVal
		}
	} else {
		out.Extra = nil
	}
	return nil
}

func Convert_v1beta1_UserInfo_To_authentication_UserInfo(in *UserInfo, out *authentication.UserInfo, s conversion.Scope) error {
	return autoConvert_v1beta1_UserInfo_To_authentication_UserInfo(in, out, s)
}

func autoConvert_authentication_UserInfo_To_v1beta1_UserInfo(in *authentication.UserInfo, out *UserInfo, s conversion.Scope) error {
	out.Username = in.Username
	out.UID = in.UID
	out.Groups = in.Groups
	if in.Extra != nil {
		in, out := &in.Extra, &out.Extra
		*out = make(map[string]ExtraValue, len(*in))
		for key, val := range *in {
			newVal := new(ExtraValue)
			// TODO: Inefficient conversion - can we improve it?
			if err := s.Convert(&val, newVal, 0); err != nil {
				return err
			}
			(*out)[key] = *newVal
		}
	} else {
		out.Extra = nil
	}
	return nil
}

func Convert_authentication_UserInfo_To_v1beta1_UserInfo(in *authentication.UserInfo, out *UserInfo, s conversion.Scope) error {
	return autoConvert_authentication_UserInfo_To_v1beta1_UserInfo(in, out, s)
}
