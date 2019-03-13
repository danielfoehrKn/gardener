// +build !ignore_autogenerated

// Code generated by conversion-gen. DO NOT EDIT.

package v1alpha1

import (
	unsafe "unsafe"

	core "github.com/gardener/gardener/pkg/apis/core"
	v1 "k8s.io/api/core/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*Cloud)(nil), (*core.Cloud)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_Cloud_To_core_Cloud(a.(*Cloud), b.(*core.Cloud), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*core.Cloud)(nil), (*Cloud)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_core_Cloud_To_v1alpha1_Cloud(a.(*core.Cloud), b.(*Cloud), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ClusterInfo)(nil), (*core.ClusterInfo)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ClusterInfo_To_core_ClusterInfo(a.(*ClusterInfo), b.(*core.ClusterInfo), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*core.ClusterInfo)(nil), (*ClusterInfo)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_core_ClusterInfo_To_v1alpha1_ClusterInfo(a.(*core.ClusterInfo), b.(*ClusterInfo), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*Condition)(nil), (*core.Condition)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_Condition_To_core_Condition(a.(*Condition), b.(*core.Condition), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*core.Condition)(nil), (*Condition)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_core_Condition_To_v1alpha1_Condition(a.(*core.Condition), b.(*Condition), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ControllerDeployment)(nil), (*core.ControllerDeployment)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ControllerDeployment_To_core_ControllerDeployment(a.(*ControllerDeployment), b.(*core.ControllerDeployment), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*core.ControllerDeployment)(nil), (*ControllerDeployment)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_core_ControllerDeployment_To_v1alpha1_ControllerDeployment(a.(*core.ControllerDeployment), b.(*ControllerDeployment), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ControllerInstallation)(nil), (*core.ControllerInstallation)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ControllerInstallation_To_core_ControllerInstallation(a.(*ControllerInstallation), b.(*core.ControllerInstallation), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*core.ControllerInstallation)(nil), (*ControllerInstallation)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_core_ControllerInstallation_To_v1alpha1_ControllerInstallation(a.(*core.ControllerInstallation), b.(*ControllerInstallation), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ControllerInstallationList)(nil), (*core.ControllerInstallationList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ControllerInstallationList_To_core_ControllerInstallationList(a.(*ControllerInstallationList), b.(*core.ControllerInstallationList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*core.ControllerInstallationList)(nil), (*ControllerInstallationList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_core_ControllerInstallationList_To_v1alpha1_ControllerInstallationList(a.(*core.ControllerInstallationList), b.(*ControllerInstallationList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ControllerInstallationSpec)(nil), (*core.ControllerInstallationSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ControllerInstallationSpec_To_core_ControllerInstallationSpec(a.(*ControllerInstallationSpec), b.(*core.ControllerInstallationSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*core.ControllerInstallationSpec)(nil), (*ControllerInstallationSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_core_ControllerInstallationSpec_To_v1alpha1_ControllerInstallationSpec(a.(*core.ControllerInstallationSpec), b.(*ControllerInstallationSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ControllerInstallationStatus)(nil), (*core.ControllerInstallationStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ControllerInstallationStatus_To_core_ControllerInstallationStatus(a.(*ControllerInstallationStatus), b.(*core.ControllerInstallationStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*core.ControllerInstallationStatus)(nil), (*ControllerInstallationStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_core_ControllerInstallationStatus_To_v1alpha1_ControllerInstallationStatus(a.(*core.ControllerInstallationStatus), b.(*ControllerInstallationStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ControllerRegistration)(nil), (*core.ControllerRegistration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ControllerRegistration_To_core_ControllerRegistration(a.(*ControllerRegistration), b.(*core.ControllerRegistration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*core.ControllerRegistration)(nil), (*ControllerRegistration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_core_ControllerRegistration_To_v1alpha1_ControllerRegistration(a.(*core.ControllerRegistration), b.(*ControllerRegistration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ControllerRegistrationList)(nil), (*core.ControllerRegistrationList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ControllerRegistrationList_To_core_ControllerRegistrationList(a.(*ControllerRegistrationList), b.(*core.ControllerRegistrationList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*core.ControllerRegistrationList)(nil), (*ControllerRegistrationList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_core_ControllerRegistrationList_To_v1alpha1_ControllerRegistrationList(a.(*core.ControllerRegistrationList), b.(*ControllerRegistrationList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ControllerRegistrationSpec)(nil), (*core.ControllerRegistrationSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ControllerRegistrationSpec_To_core_ControllerRegistrationSpec(a.(*ControllerRegistrationSpec), b.(*core.ControllerRegistrationSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*core.ControllerRegistrationSpec)(nil), (*ControllerRegistrationSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_core_ControllerRegistrationSpec_To_v1alpha1_ControllerRegistrationSpec(a.(*core.ControllerRegistrationSpec), b.(*ControllerRegistrationSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ControllerResource)(nil), (*core.ControllerResource)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ControllerResource_To_core_ControllerResource(a.(*ControllerResource), b.(*core.ControllerResource), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*core.ControllerResource)(nil), (*ControllerResource)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_core_ControllerResource_To_v1alpha1_ControllerResource(a.(*core.ControllerResource), b.(*ControllerResource), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*Endpoint)(nil), (*core.Endpoint)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_Endpoint_To_core_Endpoint(a.(*Endpoint), b.(*core.Endpoint), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*core.Endpoint)(nil), (*Endpoint)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_core_Endpoint_To_v1alpha1_Endpoint(a.(*core.Endpoint), b.(*Endpoint), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*Kubernetes)(nil), (*core.Kubernetes)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_Kubernetes_To_core_Kubernetes(a.(*Kubernetes), b.(*core.Kubernetes), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*core.Kubernetes)(nil), (*Kubernetes)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_core_Kubernetes_To_v1alpha1_Kubernetes(a.(*core.Kubernetes), b.(*Kubernetes), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*Logging)(nil), (*core.Logging)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_Logging_To_core_Logging(a.(*Logging), b.(*core.Logging), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*core.Logging)(nil), (*Logging)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_core_Logging_To_v1alpha1_Logging(a.(*core.Logging), b.(*Logging), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*Monitoring)(nil), (*core.Monitoring)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_Monitoring_To_core_Monitoring(a.(*Monitoring), b.(*core.Monitoring), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*core.Monitoring)(nil), (*Monitoring)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_core_Monitoring_To_v1alpha1_Monitoring(a.(*core.Monitoring), b.(*Monitoring), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*Plant)(nil), (*core.Plant)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_Plant_To_core_Plant(a.(*Plant), b.(*core.Plant), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*core.Plant)(nil), (*Plant)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_core_Plant_To_v1alpha1_Plant(a.(*core.Plant), b.(*Plant), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*PlantList)(nil), (*core.PlantList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_PlantList_To_core_PlantList(a.(*PlantList), b.(*core.PlantList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*core.PlantList)(nil), (*PlantList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_core_PlantList_To_v1alpha1_PlantList(a.(*core.PlantList), b.(*PlantList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*PlantSpec)(nil), (*core.PlantSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_PlantSpec_To_core_PlantSpec(a.(*PlantSpec), b.(*core.PlantSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*core.PlantSpec)(nil), (*PlantSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_core_PlantSpec_To_v1alpha1_PlantSpec(a.(*core.PlantSpec), b.(*PlantSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*PlantStatus)(nil), (*core.PlantStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_PlantStatus_To_core_PlantStatus(a.(*PlantStatus), b.(*core.PlantStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*core.PlantStatus)(nil), (*PlantStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_core_PlantStatus_To_v1alpha1_PlantStatus(a.(*core.PlantStatus), b.(*PlantStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ProviderConfig)(nil), (*core.ProviderConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ProviderConfig_To_core_ProviderConfig(a.(*ProviderConfig), b.(*core.ProviderConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*core.ProviderConfig)(nil), (*ProviderConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_core_ProviderConfig_To_v1alpha1_ProviderConfig(a.(*core.ProviderConfig), b.(*ProviderConfig), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1alpha1_Cloud_To_core_Cloud(in *Cloud, out *core.Cloud, s conversion.Scope) error {
	out.Type = in.Type
	out.Region = in.Region
	return nil
}

// Convert_v1alpha1_Cloud_To_core_Cloud is an autogenerated conversion function.
func Convert_v1alpha1_Cloud_To_core_Cloud(in *Cloud, out *core.Cloud, s conversion.Scope) error {
	return autoConvert_v1alpha1_Cloud_To_core_Cloud(in, out, s)
}

func autoConvert_core_Cloud_To_v1alpha1_Cloud(in *core.Cloud, out *Cloud, s conversion.Scope) error {
	out.Type = in.Type
	out.Region = in.Region
	return nil
}

// Convert_core_Cloud_To_v1alpha1_Cloud is an autogenerated conversion function.
func Convert_core_Cloud_To_v1alpha1_Cloud(in *core.Cloud, out *Cloud, s conversion.Scope) error {
	return autoConvert_core_Cloud_To_v1alpha1_Cloud(in, out, s)
}

func autoConvert_v1alpha1_ClusterInfo_To_core_ClusterInfo(in *ClusterInfo, out *core.ClusterInfo, s conversion.Scope) error {
	if err := Convert_v1alpha1_Cloud_To_core_Cloud(&in.Cloud, &out.Cloud, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_Kubernetes_To_core_Kubernetes(&in.Kubernetes, &out.Kubernetes, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_ClusterInfo_To_core_ClusterInfo is an autogenerated conversion function.
func Convert_v1alpha1_ClusterInfo_To_core_ClusterInfo(in *ClusterInfo, out *core.ClusterInfo, s conversion.Scope) error {
	return autoConvert_v1alpha1_ClusterInfo_To_core_ClusterInfo(in, out, s)
}

func autoConvert_core_ClusterInfo_To_v1alpha1_ClusterInfo(in *core.ClusterInfo, out *ClusterInfo, s conversion.Scope) error {
	if err := Convert_core_Cloud_To_v1alpha1_Cloud(&in.Cloud, &out.Cloud, s); err != nil {
		return err
	}
	if err := Convert_core_Kubernetes_To_v1alpha1_Kubernetes(&in.Kubernetes, &out.Kubernetes, s); err != nil {
		return err
	}
	return nil
}

// Convert_core_ClusterInfo_To_v1alpha1_ClusterInfo is an autogenerated conversion function.
func Convert_core_ClusterInfo_To_v1alpha1_ClusterInfo(in *core.ClusterInfo, out *ClusterInfo, s conversion.Scope) error {
	return autoConvert_core_ClusterInfo_To_v1alpha1_ClusterInfo(in, out, s)
}

func autoConvert_v1alpha1_Condition_To_core_Condition(in *Condition, out *core.Condition, s conversion.Scope) error {
	out.Type = core.ConditionType(in.Type)
	out.Status = v1.ConditionStatus(in.Status)
	out.LastTransitionTime = in.LastTransitionTime
	out.LastUpdateTime = in.LastUpdateTime
	out.Reason = in.Reason
	out.Message = in.Message
	return nil
}

// Convert_v1alpha1_Condition_To_core_Condition is an autogenerated conversion function.
func Convert_v1alpha1_Condition_To_core_Condition(in *Condition, out *core.Condition, s conversion.Scope) error {
	return autoConvert_v1alpha1_Condition_To_core_Condition(in, out, s)
}

func autoConvert_core_Condition_To_v1alpha1_Condition(in *core.Condition, out *Condition, s conversion.Scope) error {
	out.Type = ConditionType(in.Type)
	out.Status = v1.ConditionStatus(in.Status)
	out.LastTransitionTime = in.LastTransitionTime
	out.LastUpdateTime = in.LastUpdateTime
	out.Reason = in.Reason
	out.Message = in.Message
	return nil
}

// Convert_core_Condition_To_v1alpha1_Condition is an autogenerated conversion function.
func Convert_core_Condition_To_v1alpha1_Condition(in *core.Condition, out *Condition, s conversion.Scope) error {
	return autoConvert_core_Condition_To_v1alpha1_Condition(in, out, s)
}

func autoConvert_v1alpha1_ControllerDeployment_To_core_ControllerDeployment(in *ControllerDeployment, out *core.ControllerDeployment, s conversion.Scope) error {
	out.Type = in.Type
	out.ProviderConfig = (*core.ProviderConfig)(unsafe.Pointer(in.ProviderConfig))
	return nil
}

// Convert_v1alpha1_ControllerDeployment_To_core_ControllerDeployment is an autogenerated conversion function.
func Convert_v1alpha1_ControllerDeployment_To_core_ControllerDeployment(in *ControllerDeployment, out *core.ControllerDeployment, s conversion.Scope) error {
	return autoConvert_v1alpha1_ControllerDeployment_To_core_ControllerDeployment(in, out, s)
}

func autoConvert_core_ControllerDeployment_To_v1alpha1_ControllerDeployment(in *core.ControllerDeployment, out *ControllerDeployment, s conversion.Scope) error {
	out.Type = in.Type
	out.ProviderConfig = (*ProviderConfig)(unsafe.Pointer(in.ProviderConfig))
	return nil
}

// Convert_core_ControllerDeployment_To_v1alpha1_ControllerDeployment is an autogenerated conversion function.
func Convert_core_ControllerDeployment_To_v1alpha1_ControllerDeployment(in *core.ControllerDeployment, out *ControllerDeployment, s conversion.Scope) error {
	return autoConvert_core_ControllerDeployment_To_v1alpha1_ControllerDeployment(in, out, s)
}

func autoConvert_v1alpha1_ControllerInstallation_To_core_ControllerInstallation(in *ControllerInstallation, out *core.ControllerInstallation, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1alpha1_ControllerInstallationSpec_To_core_ControllerInstallationSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_ControllerInstallationStatus_To_core_ControllerInstallationStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_ControllerInstallation_To_core_ControllerInstallation is an autogenerated conversion function.
func Convert_v1alpha1_ControllerInstallation_To_core_ControllerInstallation(in *ControllerInstallation, out *core.ControllerInstallation, s conversion.Scope) error {
	return autoConvert_v1alpha1_ControllerInstallation_To_core_ControllerInstallation(in, out, s)
}

func autoConvert_core_ControllerInstallation_To_v1alpha1_ControllerInstallation(in *core.ControllerInstallation, out *ControllerInstallation, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_core_ControllerInstallationSpec_To_v1alpha1_ControllerInstallationSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_core_ControllerInstallationStatus_To_v1alpha1_ControllerInstallationStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_core_ControllerInstallation_To_v1alpha1_ControllerInstallation is an autogenerated conversion function.
func Convert_core_ControllerInstallation_To_v1alpha1_ControllerInstallation(in *core.ControllerInstallation, out *ControllerInstallation, s conversion.Scope) error {
	return autoConvert_core_ControllerInstallation_To_v1alpha1_ControllerInstallation(in, out, s)
}

func autoConvert_v1alpha1_ControllerInstallationList_To_core_ControllerInstallationList(in *ControllerInstallationList, out *core.ControllerInstallationList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]core.ControllerInstallation)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1alpha1_ControllerInstallationList_To_core_ControllerInstallationList is an autogenerated conversion function.
func Convert_v1alpha1_ControllerInstallationList_To_core_ControllerInstallationList(in *ControllerInstallationList, out *core.ControllerInstallationList, s conversion.Scope) error {
	return autoConvert_v1alpha1_ControllerInstallationList_To_core_ControllerInstallationList(in, out, s)
}

func autoConvert_core_ControllerInstallationList_To_v1alpha1_ControllerInstallationList(in *core.ControllerInstallationList, out *ControllerInstallationList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]ControllerInstallation)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_core_ControllerInstallationList_To_v1alpha1_ControllerInstallationList is an autogenerated conversion function.
func Convert_core_ControllerInstallationList_To_v1alpha1_ControllerInstallationList(in *core.ControllerInstallationList, out *ControllerInstallationList, s conversion.Scope) error {
	return autoConvert_core_ControllerInstallationList_To_v1alpha1_ControllerInstallationList(in, out, s)
}

func autoConvert_v1alpha1_ControllerInstallationSpec_To_core_ControllerInstallationSpec(in *ControllerInstallationSpec, out *core.ControllerInstallationSpec, s conversion.Scope) error {
	out.RegistrationRef = in.RegistrationRef
	out.SeedRef = in.SeedRef
	return nil
}

// Convert_v1alpha1_ControllerInstallationSpec_To_core_ControllerInstallationSpec is an autogenerated conversion function.
func Convert_v1alpha1_ControllerInstallationSpec_To_core_ControllerInstallationSpec(in *ControllerInstallationSpec, out *core.ControllerInstallationSpec, s conversion.Scope) error {
	return autoConvert_v1alpha1_ControllerInstallationSpec_To_core_ControllerInstallationSpec(in, out, s)
}

func autoConvert_core_ControllerInstallationSpec_To_v1alpha1_ControllerInstallationSpec(in *core.ControllerInstallationSpec, out *ControllerInstallationSpec, s conversion.Scope) error {
	out.RegistrationRef = in.RegistrationRef
	out.SeedRef = in.SeedRef
	return nil
}

// Convert_core_ControllerInstallationSpec_To_v1alpha1_ControllerInstallationSpec is an autogenerated conversion function.
func Convert_core_ControllerInstallationSpec_To_v1alpha1_ControllerInstallationSpec(in *core.ControllerInstallationSpec, out *ControllerInstallationSpec, s conversion.Scope) error {
	return autoConvert_core_ControllerInstallationSpec_To_v1alpha1_ControllerInstallationSpec(in, out, s)
}

func autoConvert_v1alpha1_ControllerInstallationStatus_To_core_ControllerInstallationStatus(in *ControllerInstallationStatus, out *core.ControllerInstallationStatus, s conversion.Scope) error {
	out.Conditions = *(*[]core.Condition)(unsafe.Pointer(&in.Conditions))
	out.ProviderStatus = (*core.ProviderConfig)(unsafe.Pointer(in.ProviderStatus))
	return nil
}

// Convert_v1alpha1_ControllerInstallationStatus_To_core_ControllerInstallationStatus is an autogenerated conversion function.
func Convert_v1alpha1_ControllerInstallationStatus_To_core_ControllerInstallationStatus(in *ControllerInstallationStatus, out *core.ControllerInstallationStatus, s conversion.Scope) error {
	return autoConvert_v1alpha1_ControllerInstallationStatus_To_core_ControllerInstallationStatus(in, out, s)
}

func autoConvert_core_ControllerInstallationStatus_To_v1alpha1_ControllerInstallationStatus(in *core.ControllerInstallationStatus, out *ControllerInstallationStatus, s conversion.Scope) error {
	out.Conditions = *(*[]Condition)(unsafe.Pointer(&in.Conditions))
	out.ProviderStatus = (*ProviderConfig)(unsafe.Pointer(in.ProviderStatus))
	return nil
}

// Convert_core_ControllerInstallationStatus_To_v1alpha1_ControllerInstallationStatus is an autogenerated conversion function.
func Convert_core_ControllerInstallationStatus_To_v1alpha1_ControllerInstallationStatus(in *core.ControllerInstallationStatus, out *ControllerInstallationStatus, s conversion.Scope) error {
	return autoConvert_core_ControllerInstallationStatus_To_v1alpha1_ControllerInstallationStatus(in, out, s)
}

func autoConvert_v1alpha1_ControllerRegistration_To_core_ControllerRegistration(in *ControllerRegistration, out *core.ControllerRegistration, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1alpha1_ControllerRegistrationSpec_To_core_ControllerRegistrationSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_ControllerRegistration_To_core_ControllerRegistration is an autogenerated conversion function.
func Convert_v1alpha1_ControllerRegistration_To_core_ControllerRegistration(in *ControllerRegistration, out *core.ControllerRegistration, s conversion.Scope) error {
	return autoConvert_v1alpha1_ControllerRegistration_To_core_ControllerRegistration(in, out, s)
}

func autoConvert_core_ControllerRegistration_To_v1alpha1_ControllerRegistration(in *core.ControllerRegistration, out *ControllerRegistration, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_core_ControllerRegistrationSpec_To_v1alpha1_ControllerRegistrationSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

// Convert_core_ControllerRegistration_To_v1alpha1_ControllerRegistration is an autogenerated conversion function.
func Convert_core_ControllerRegistration_To_v1alpha1_ControllerRegistration(in *core.ControllerRegistration, out *ControllerRegistration, s conversion.Scope) error {
	return autoConvert_core_ControllerRegistration_To_v1alpha1_ControllerRegistration(in, out, s)
}

func autoConvert_v1alpha1_ControllerRegistrationList_To_core_ControllerRegistrationList(in *ControllerRegistrationList, out *core.ControllerRegistrationList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]core.ControllerRegistration)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1alpha1_ControllerRegistrationList_To_core_ControllerRegistrationList is an autogenerated conversion function.
func Convert_v1alpha1_ControllerRegistrationList_To_core_ControllerRegistrationList(in *ControllerRegistrationList, out *core.ControllerRegistrationList, s conversion.Scope) error {
	return autoConvert_v1alpha1_ControllerRegistrationList_To_core_ControllerRegistrationList(in, out, s)
}

func autoConvert_core_ControllerRegistrationList_To_v1alpha1_ControllerRegistrationList(in *core.ControllerRegistrationList, out *ControllerRegistrationList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]ControllerRegistration)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_core_ControllerRegistrationList_To_v1alpha1_ControllerRegistrationList is an autogenerated conversion function.
func Convert_core_ControllerRegistrationList_To_v1alpha1_ControllerRegistrationList(in *core.ControllerRegistrationList, out *ControllerRegistrationList, s conversion.Scope) error {
	return autoConvert_core_ControllerRegistrationList_To_v1alpha1_ControllerRegistrationList(in, out, s)
}

func autoConvert_v1alpha1_ControllerRegistrationSpec_To_core_ControllerRegistrationSpec(in *ControllerRegistrationSpec, out *core.ControllerRegistrationSpec, s conversion.Scope) error {
	out.Resources = *(*[]core.ControllerResource)(unsafe.Pointer(&in.Resources))
	out.Deployment = (*core.ControllerDeployment)(unsafe.Pointer(in.Deployment))
	return nil
}

// Convert_v1alpha1_ControllerRegistrationSpec_To_core_ControllerRegistrationSpec is an autogenerated conversion function.
func Convert_v1alpha1_ControllerRegistrationSpec_To_core_ControllerRegistrationSpec(in *ControllerRegistrationSpec, out *core.ControllerRegistrationSpec, s conversion.Scope) error {
	return autoConvert_v1alpha1_ControllerRegistrationSpec_To_core_ControllerRegistrationSpec(in, out, s)
}

func autoConvert_core_ControllerRegistrationSpec_To_v1alpha1_ControllerRegistrationSpec(in *core.ControllerRegistrationSpec, out *ControllerRegistrationSpec, s conversion.Scope) error {
	out.Resources = *(*[]ControllerResource)(unsafe.Pointer(&in.Resources))
	out.Deployment = (*ControllerDeployment)(unsafe.Pointer(in.Deployment))
	return nil
}

// Convert_core_ControllerRegistrationSpec_To_v1alpha1_ControllerRegistrationSpec is an autogenerated conversion function.
func Convert_core_ControllerRegistrationSpec_To_v1alpha1_ControllerRegistrationSpec(in *core.ControllerRegistrationSpec, out *ControllerRegistrationSpec, s conversion.Scope) error {
	return autoConvert_core_ControllerRegistrationSpec_To_v1alpha1_ControllerRegistrationSpec(in, out, s)
}

func autoConvert_v1alpha1_ControllerResource_To_core_ControllerResource(in *ControllerResource, out *core.ControllerResource, s conversion.Scope) error {
	out.Kind = in.Kind
	out.Type = in.Type
	return nil
}

// Convert_v1alpha1_ControllerResource_To_core_ControllerResource is an autogenerated conversion function.
func Convert_v1alpha1_ControllerResource_To_core_ControllerResource(in *ControllerResource, out *core.ControllerResource, s conversion.Scope) error {
	return autoConvert_v1alpha1_ControllerResource_To_core_ControllerResource(in, out, s)
}

func autoConvert_core_ControllerResource_To_v1alpha1_ControllerResource(in *core.ControllerResource, out *ControllerResource, s conversion.Scope) error {
	out.Kind = in.Kind
	out.Type = in.Type
	return nil
}

// Convert_core_ControllerResource_To_v1alpha1_ControllerResource is an autogenerated conversion function.
func Convert_core_ControllerResource_To_v1alpha1_ControllerResource(in *core.ControllerResource, out *ControllerResource, s conversion.Scope) error {
	return autoConvert_core_ControllerResource_To_v1alpha1_ControllerResource(in, out, s)
}

func autoConvert_v1alpha1_Endpoint_To_core_Endpoint(in *Endpoint, out *core.Endpoint, s conversion.Scope) error {
	out.Name = in.Name
	out.URL = in.URL
	return nil
}

// Convert_v1alpha1_Endpoint_To_core_Endpoint is an autogenerated conversion function.
func Convert_v1alpha1_Endpoint_To_core_Endpoint(in *Endpoint, out *core.Endpoint, s conversion.Scope) error {
	return autoConvert_v1alpha1_Endpoint_To_core_Endpoint(in, out, s)
}

func autoConvert_core_Endpoint_To_v1alpha1_Endpoint(in *core.Endpoint, out *Endpoint, s conversion.Scope) error {
	out.Name = in.Name
	out.URL = in.URL
	return nil
}

// Convert_core_Endpoint_To_v1alpha1_Endpoint is an autogenerated conversion function.
func Convert_core_Endpoint_To_v1alpha1_Endpoint(in *core.Endpoint, out *Endpoint, s conversion.Scope) error {
	return autoConvert_core_Endpoint_To_v1alpha1_Endpoint(in, out, s)
}

func autoConvert_v1alpha1_Kubernetes_To_core_Kubernetes(in *Kubernetes, out *core.Kubernetes, s conversion.Scope) error {
	out.Version = in.Version
	return nil
}

// Convert_v1alpha1_Kubernetes_To_core_Kubernetes is an autogenerated conversion function.
func Convert_v1alpha1_Kubernetes_To_core_Kubernetes(in *Kubernetes, out *core.Kubernetes, s conversion.Scope) error {
	return autoConvert_v1alpha1_Kubernetes_To_core_Kubernetes(in, out, s)
}

func autoConvert_core_Kubernetes_To_v1alpha1_Kubernetes(in *core.Kubernetes, out *Kubernetes, s conversion.Scope) error {
	out.Version = in.Version
	return nil
}

// Convert_core_Kubernetes_To_v1alpha1_Kubernetes is an autogenerated conversion function.
func Convert_core_Kubernetes_To_v1alpha1_Kubernetes(in *core.Kubernetes, out *Kubernetes, s conversion.Scope) error {
	return autoConvert_core_Kubernetes_To_v1alpha1_Kubernetes(in, out, s)
}

func autoConvert_v1alpha1_Logging_To_core_Logging(in *Logging, out *core.Logging, s conversion.Scope) error {
	out.Endpoints = *(*[]core.Endpoint)(unsafe.Pointer(&in.Endpoints))
	return nil
}

// Convert_v1alpha1_Logging_To_core_Logging is an autogenerated conversion function.
func Convert_v1alpha1_Logging_To_core_Logging(in *Logging, out *core.Logging, s conversion.Scope) error {
	return autoConvert_v1alpha1_Logging_To_core_Logging(in, out, s)
}

func autoConvert_core_Logging_To_v1alpha1_Logging(in *core.Logging, out *Logging, s conversion.Scope) error {
	out.Endpoints = *(*[]Endpoint)(unsafe.Pointer(&in.Endpoints))
	return nil
}

// Convert_core_Logging_To_v1alpha1_Logging is an autogenerated conversion function.
func Convert_core_Logging_To_v1alpha1_Logging(in *core.Logging, out *Logging, s conversion.Scope) error {
	return autoConvert_core_Logging_To_v1alpha1_Logging(in, out, s)
}

func autoConvert_v1alpha1_Monitoring_To_core_Monitoring(in *Monitoring, out *core.Monitoring, s conversion.Scope) error {
	out.Endpoints = *(*[]core.Endpoint)(unsafe.Pointer(&in.Endpoints))
	return nil
}

// Convert_v1alpha1_Monitoring_To_core_Monitoring is an autogenerated conversion function.
func Convert_v1alpha1_Monitoring_To_core_Monitoring(in *Monitoring, out *core.Monitoring, s conversion.Scope) error {
	return autoConvert_v1alpha1_Monitoring_To_core_Monitoring(in, out, s)
}

func autoConvert_core_Monitoring_To_v1alpha1_Monitoring(in *core.Monitoring, out *Monitoring, s conversion.Scope) error {
	out.Endpoints = *(*[]Endpoint)(unsafe.Pointer(&in.Endpoints))
	return nil
}

// Convert_core_Monitoring_To_v1alpha1_Monitoring is an autogenerated conversion function.
func Convert_core_Monitoring_To_v1alpha1_Monitoring(in *core.Monitoring, out *Monitoring, s conversion.Scope) error {
	return autoConvert_core_Monitoring_To_v1alpha1_Monitoring(in, out, s)
}

func autoConvert_v1alpha1_Plant_To_core_Plant(in *Plant, out *core.Plant, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1alpha1_PlantSpec_To_core_PlantSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_PlantStatus_To_core_PlantStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_Plant_To_core_Plant is an autogenerated conversion function.
func Convert_v1alpha1_Plant_To_core_Plant(in *Plant, out *core.Plant, s conversion.Scope) error {
	return autoConvert_v1alpha1_Plant_To_core_Plant(in, out, s)
}

func autoConvert_core_Plant_To_v1alpha1_Plant(in *core.Plant, out *Plant, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_core_PlantSpec_To_v1alpha1_PlantSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_core_PlantStatus_To_v1alpha1_PlantStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_core_Plant_To_v1alpha1_Plant is an autogenerated conversion function.
func Convert_core_Plant_To_v1alpha1_Plant(in *core.Plant, out *Plant, s conversion.Scope) error {
	return autoConvert_core_Plant_To_v1alpha1_Plant(in, out, s)
}

func autoConvert_v1alpha1_PlantList_To_core_PlantList(in *PlantList, out *core.PlantList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]core.Plant)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1alpha1_PlantList_To_core_PlantList is an autogenerated conversion function.
func Convert_v1alpha1_PlantList_To_core_PlantList(in *PlantList, out *core.PlantList, s conversion.Scope) error {
	return autoConvert_v1alpha1_PlantList_To_core_PlantList(in, out, s)
}

func autoConvert_core_PlantList_To_v1alpha1_PlantList(in *core.PlantList, out *PlantList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]Plant)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_core_PlantList_To_v1alpha1_PlantList is an autogenerated conversion function.
func Convert_core_PlantList_To_v1alpha1_PlantList(in *core.PlantList, out *PlantList, s conversion.Scope) error {
	return autoConvert_core_PlantList_To_v1alpha1_PlantList(in, out, s)
}

func autoConvert_v1alpha1_PlantSpec_To_core_PlantSpec(in *PlantSpec, out *core.PlantSpec, s conversion.Scope) error {
	out.SecretRef = in.SecretRef
	out.Monitoring = (*core.Monitoring)(unsafe.Pointer(in.Monitoring))
	out.Logging = (*core.Logging)(unsafe.Pointer(in.Logging))
	return nil
}

// Convert_v1alpha1_PlantSpec_To_core_PlantSpec is an autogenerated conversion function.
func Convert_v1alpha1_PlantSpec_To_core_PlantSpec(in *PlantSpec, out *core.PlantSpec, s conversion.Scope) error {
	return autoConvert_v1alpha1_PlantSpec_To_core_PlantSpec(in, out, s)
}

func autoConvert_core_PlantSpec_To_v1alpha1_PlantSpec(in *core.PlantSpec, out *PlantSpec, s conversion.Scope) error {
	out.SecretRef = in.SecretRef
	out.Monitoring = (*Monitoring)(unsafe.Pointer(in.Monitoring))
	out.Logging = (*Logging)(unsafe.Pointer(in.Logging))
	return nil
}

// Convert_core_PlantSpec_To_v1alpha1_PlantSpec is an autogenerated conversion function.
func Convert_core_PlantSpec_To_v1alpha1_PlantSpec(in *core.PlantSpec, out *PlantSpec, s conversion.Scope) error {
	return autoConvert_core_PlantSpec_To_v1alpha1_PlantSpec(in, out, s)
}

func autoConvert_v1alpha1_PlantStatus_To_core_PlantStatus(in *PlantStatus, out *core.PlantStatus, s conversion.Scope) error {
	out.Conditions = *(*[]core.Condition)(unsafe.Pointer(&in.Conditions))
	out.ObservedGeneration = (*int64)(unsafe.Pointer(in.ObservedGeneration))
	out.ClusterInfo = (*core.ClusterInfo)(unsafe.Pointer(in.ClusterInfo))
	return nil
}

// Convert_v1alpha1_PlantStatus_To_core_PlantStatus is an autogenerated conversion function.
func Convert_v1alpha1_PlantStatus_To_core_PlantStatus(in *PlantStatus, out *core.PlantStatus, s conversion.Scope) error {
	return autoConvert_v1alpha1_PlantStatus_To_core_PlantStatus(in, out, s)
}

func autoConvert_core_PlantStatus_To_v1alpha1_PlantStatus(in *core.PlantStatus, out *PlantStatus, s conversion.Scope) error {
	out.Conditions = *(*[]Condition)(unsafe.Pointer(&in.Conditions))
	out.ObservedGeneration = (*int64)(unsafe.Pointer(in.ObservedGeneration))
	out.ClusterInfo = (*ClusterInfo)(unsafe.Pointer(in.ClusterInfo))
	return nil
}

// Convert_core_PlantStatus_To_v1alpha1_PlantStatus is an autogenerated conversion function.
func Convert_core_PlantStatus_To_v1alpha1_PlantStatus(in *core.PlantStatus, out *PlantStatus, s conversion.Scope) error {
	return autoConvert_core_PlantStatus_To_v1alpha1_PlantStatus(in, out, s)
}

func autoConvert_v1alpha1_ProviderConfig_To_core_ProviderConfig(in *ProviderConfig, out *core.ProviderConfig, s conversion.Scope) error {
	out.RawExtension = in.RawExtension
	return nil
}

// Convert_v1alpha1_ProviderConfig_To_core_ProviderConfig is an autogenerated conversion function.
func Convert_v1alpha1_ProviderConfig_To_core_ProviderConfig(in *ProviderConfig, out *core.ProviderConfig, s conversion.Scope) error {
	return autoConvert_v1alpha1_ProviderConfig_To_core_ProviderConfig(in, out, s)
}

func autoConvert_core_ProviderConfig_To_v1alpha1_ProviderConfig(in *core.ProviderConfig, out *ProviderConfig, s conversion.Scope) error {
	out.RawExtension = in.RawExtension
	return nil
}

// Convert_core_ProviderConfig_To_v1alpha1_ProviderConfig is an autogenerated conversion function.
func Convert_core_ProviderConfig_To_v1alpha1_ProviderConfig(in *core.ProviderConfig, out *ProviderConfig, s conversion.Scope) error {
	return autoConvert_core_ProviderConfig_To_v1alpha1_ProviderConfig(in, out, s)
}
