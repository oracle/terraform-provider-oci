// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fusion Applications Environment Management API
//
// Use the Fusion Applications Environment Management API to manage the environments where your Fusion Applications run. For more information, see the Fusion Applications Environment Management documentation (https://docs.cloud.oracle.com/iaas/Content/fusion-applications/home.htm).
//

package fusionapps

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// WorkRequest A description of workrequest status
type WorkRequest struct {

	// Possible operation types.
	OperationType WorkRequestOperationTypeEnum `mandatory:"true" json:"operationType"`

	// Possible operation status.
	Status WorkRequestStatusEnum `mandatory:"true" json:"status"`

	// The id of the work request.
	Id *string `mandatory:"true" json:"id"`

	// The ocid of the compartment that contains the work request. Work requests should be scoped to
	// the same compartment as the resource the work request affects. If the work request affects multiple resources,
	// and those resources are not in the same compartment, it is up to the service team to pick the primary
	// resource whose compartment should be used
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The resources affected by this work request.
	Resources []WorkRequestResource `mandatory:"true" json:"resources"`

	// Percentage of the request completed.
	PercentComplete *float32 `mandatory:"true" json:"percentComplete"`

	// The date and time the request was created, as described in
	// RFC 3339 (https://tools.ietf.org/rfc/rfc3339), section 14.29.
	TimeAccepted *common.SDKTime `mandatory:"true" json:"timeAccepted"`

	// The date and time the request was started, as described in RFC 3339 (https://tools.ietf.org/rfc/rfc3339),
	// section 14.29.
	TimeStarted *common.SDKTime `mandatory:"false" json:"timeStarted"`

	// The date and time the object was finished, as described in RFC 3339 (https://tools.ietf.org/rfc/rfc3339).
	TimeFinished *common.SDKTime `mandatory:"false" json:"timeFinished"`
}

func (m WorkRequest) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m WorkRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingWorkRequestOperationTypeEnum(string(m.OperationType)); !ok && m.OperationType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OperationType: %s. Supported values are: %s.", m.OperationType, strings.Join(GetWorkRequestOperationTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingWorkRequestStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetWorkRequestStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// WorkRequestOperationTypeEnum Enum with underlying type: string
type WorkRequestOperationTypeEnum string

// Set of constants representing the allowable values for WorkRequestOperationTypeEnum
const (
	WorkRequestOperationTypeCreateFusionEnvironment                  WorkRequestOperationTypeEnum = "CREATE_FUSION_ENVIRONMENT"
	WorkRequestOperationTypeUpdateFusionEnvironment                  WorkRequestOperationTypeEnum = "UPDATE_FUSION_ENVIRONMENT"
	WorkRequestOperationTypeResetFusionEnvironmentAdminPassword      WorkRequestOperationTypeEnum = "RESET_FUSION_ENVIRONMENT_ADMIN_PASSWORD"
	WorkRequestOperationTypeScaleFusionEnvironment                   WorkRequestOperationTypeEnum = "SCALE_FUSION_ENVIRONMENT"
	WorkRequestOperationTypeArchiveFusionEnvironment                 WorkRequestOperationTypeEnum = "ARCHIVE_FUSION_ENVIRONMENT"
	WorkRequestOperationTypeRestoreFusionEnvironment                 WorkRequestOperationTypeEnum = "RESTORE_FUSION_ENVIRONMENT"
	WorkRequestOperationTypeCreateServiceInstance                    WorkRequestOperationTypeEnum = "CREATE_SERVICE_INSTANCE"
	WorkRequestOperationTypeUpdateServiceInstance                    WorkRequestOperationTypeEnum = "UPDATE_SERVICE_INSTANCE"
	WorkRequestOperationTypeDetachServiceInstance                    WorkRequestOperationTypeEnum = "DETACH_SERVICE_INSTANCE"
	WorkRequestOperationTypeAddUser                                  WorkRequestOperationTypeEnum = "ADD_USER"
	WorkRequestOperationTypeRemoveUser                               WorkRequestOperationTypeEnum = "REMOVE_USER"
	WorkRequestOperationTypeDeleteFusionEnvironment                  WorkRequestOperationTypeEnum = "DELETE_FUSION_ENVIRONMENT"
	WorkRequestOperationTypeChangeFusionEnvironmentCompartment       WorkRequestOperationTypeEnum = "CHANGE_FUSION_ENVIRONMENT_COMPARTMENT"
	WorkRequestOperationTypeUpgradeFusionEnvironment                 WorkRequestOperationTypeEnum = "UPGRADE_FUSION_ENVIRONMENT"
	WorkRequestOperationTypeCreateFusionEnvironmentFamily            WorkRequestOperationTypeEnum = "CREATE_FUSION_ENVIRONMENT_FAMILY"
	WorkRequestOperationTypeDeleteFusionEnvironmentFamily            WorkRequestOperationTypeEnum = "DELETE_FUSION_ENVIRONMENT_FAMILY"
	WorkRequestOperationTypeUpdateFusionEnvironmentFamily            WorkRequestOperationTypeEnum = "UPDATE_FUSION_ENVIRONMENT_FAMILY"
	WorkRequestOperationTypeChangeFusionEnvironmentFamilyCompartment WorkRequestOperationTypeEnum = "CHANGE_FUSION_ENVIRONMENT_FAMILY_COMPARTMENT"
	WorkRequestOperationTypeRefreshFusionEnvironment                 WorkRequestOperationTypeEnum = "REFRESH_FUSION_ENVIRONMENT"
	WorkRequestOperationTypeExecuteColdPatch                         WorkRequestOperationTypeEnum = "EXECUTE_COLD_PATCH"
	WorkRequestOperationTypeDataMaskFusionEnvironment                WorkRequestOperationTypeEnum = "DATA_MASK_FUSION_ENVIRONMENT"
	WorkRequestOperationTypeInitiateExtract                          WorkRequestOperationTypeEnum = "INITIATE_EXTRACT"
	WorkRequestOperationTypeSubscriptionSuspend                      WorkRequestOperationTypeEnum = "SUBSCRIPTION_SUSPEND"
	WorkRequestOperationTypeSubscriptionExpire                       WorkRequestOperationTypeEnum = "SUBSCRIPTION_EXPIRE"
	WorkRequestOperationTypeSubscriptionUpdate                       WorkRequestOperationTypeEnum = "SUBSCRIPTION_UPDATE"
	WorkRequestOperationTypeSubscriptionResume                       WorkRequestOperationTypeEnum = "SUBSCRIPTION_RESUME"
	WorkRequestOperationTypeSubscriptionTerminate                    WorkRequestOperationTypeEnum = "SUBSCRIPTION_TERMINATE"
)

var mappingWorkRequestOperationTypeEnum = map[string]WorkRequestOperationTypeEnum{
	"CREATE_FUSION_ENVIRONMENT":                    WorkRequestOperationTypeCreateFusionEnvironment,
	"UPDATE_FUSION_ENVIRONMENT":                    WorkRequestOperationTypeUpdateFusionEnvironment,
	"RESET_FUSION_ENVIRONMENT_ADMIN_PASSWORD":      WorkRequestOperationTypeResetFusionEnvironmentAdminPassword,
	"SCALE_FUSION_ENVIRONMENT":                     WorkRequestOperationTypeScaleFusionEnvironment,
	"ARCHIVE_FUSION_ENVIRONMENT":                   WorkRequestOperationTypeArchiveFusionEnvironment,
	"RESTORE_FUSION_ENVIRONMENT":                   WorkRequestOperationTypeRestoreFusionEnvironment,
	"CREATE_SERVICE_INSTANCE":                      WorkRequestOperationTypeCreateServiceInstance,
	"UPDATE_SERVICE_INSTANCE":                      WorkRequestOperationTypeUpdateServiceInstance,
	"DETACH_SERVICE_INSTANCE":                      WorkRequestOperationTypeDetachServiceInstance,
	"ADD_USER":                                     WorkRequestOperationTypeAddUser,
	"REMOVE_USER":                                  WorkRequestOperationTypeRemoveUser,
	"DELETE_FUSION_ENVIRONMENT":                    WorkRequestOperationTypeDeleteFusionEnvironment,
	"CHANGE_FUSION_ENVIRONMENT_COMPARTMENT":        WorkRequestOperationTypeChangeFusionEnvironmentCompartment,
	"UPGRADE_FUSION_ENVIRONMENT":                   WorkRequestOperationTypeUpgradeFusionEnvironment,
	"CREATE_FUSION_ENVIRONMENT_FAMILY":             WorkRequestOperationTypeCreateFusionEnvironmentFamily,
	"DELETE_FUSION_ENVIRONMENT_FAMILY":             WorkRequestOperationTypeDeleteFusionEnvironmentFamily,
	"UPDATE_FUSION_ENVIRONMENT_FAMILY":             WorkRequestOperationTypeUpdateFusionEnvironmentFamily,
	"CHANGE_FUSION_ENVIRONMENT_FAMILY_COMPARTMENT": WorkRequestOperationTypeChangeFusionEnvironmentFamilyCompartment,
	"REFRESH_FUSION_ENVIRONMENT":                   WorkRequestOperationTypeRefreshFusionEnvironment,
	"EXECUTE_COLD_PATCH":                           WorkRequestOperationTypeExecuteColdPatch,
	"DATA_MASK_FUSION_ENVIRONMENT":                 WorkRequestOperationTypeDataMaskFusionEnvironment,
	"INITIATE_EXTRACT":                             WorkRequestOperationTypeInitiateExtract,
	"SUBSCRIPTION_SUSPEND":                         WorkRequestOperationTypeSubscriptionSuspend,
	"SUBSCRIPTION_EXPIRE":                          WorkRequestOperationTypeSubscriptionExpire,
	"SUBSCRIPTION_UPDATE":                          WorkRequestOperationTypeSubscriptionUpdate,
	"SUBSCRIPTION_RESUME":                          WorkRequestOperationTypeSubscriptionResume,
	"SUBSCRIPTION_TERMINATE":                       WorkRequestOperationTypeSubscriptionTerminate,
}

var mappingWorkRequestOperationTypeEnumLowerCase = map[string]WorkRequestOperationTypeEnum{
	"create_fusion_environment":                    WorkRequestOperationTypeCreateFusionEnvironment,
	"update_fusion_environment":                    WorkRequestOperationTypeUpdateFusionEnvironment,
	"reset_fusion_environment_admin_password":      WorkRequestOperationTypeResetFusionEnvironmentAdminPassword,
	"scale_fusion_environment":                     WorkRequestOperationTypeScaleFusionEnvironment,
	"archive_fusion_environment":                   WorkRequestOperationTypeArchiveFusionEnvironment,
	"restore_fusion_environment":                   WorkRequestOperationTypeRestoreFusionEnvironment,
	"create_service_instance":                      WorkRequestOperationTypeCreateServiceInstance,
	"update_service_instance":                      WorkRequestOperationTypeUpdateServiceInstance,
	"detach_service_instance":                      WorkRequestOperationTypeDetachServiceInstance,
	"add_user":                                     WorkRequestOperationTypeAddUser,
	"remove_user":                                  WorkRequestOperationTypeRemoveUser,
	"delete_fusion_environment":                    WorkRequestOperationTypeDeleteFusionEnvironment,
	"change_fusion_environment_compartment":        WorkRequestOperationTypeChangeFusionEnvironmentCompartment,
	"upgrade_fusion_environment":                   WorkRequestOperationTypeUpgradeFusionEnvironment,
	"create_fusion_environment_family":             WorkRequestOperationTypeCreateFusionEnvironmentFamily,
	"delete_fusion_environment_family":             WorkRequestOperationTypeDeleteFusionEnvironmentFamily,
	"update_fusion_environment_family":             WorkRequestOperationTypeUpdateFusionEnvironmentFamily,
	"change_fusion_environment_family_compartment": WorkRequestOperationTypeChangeFusionEnvironmentFamilyCompartment,
	"refresh_fusion_environment":                   WorkRequestOperationTypeRefreshFusionEnvironment,
	"execute_cold_patch":                           WorkRequestOperationTypeExecuteColdPatch,
	"data_mask_fusion_environment":                 WorkRequestOperationTypeDataMaskFusionEnvironment,
	"initiate_extract":                             WorkRequestOperationTypeInitiateExtract,
	"subscription_suspend":                         WorkRequestOperationTypeSubscriptionSuspend,
	"subscription_expire":                          WorkRequestOperationTypeSubscriptionExpire,
	"subscription_update":                          WorkRequestOperationTypeSubscriptionUpdate,
	"subscription_resume":                          WorkRequestOperationTypeSubscriptionResume,
	"subscription_terminate":                       WorkRequestOperationTypeSubscriptionTerminate,
}

// GetWorkRequestOperationTypeEnumValues Enumerates the set of values for WorkRequestOperationTypeEnum
func GetWorkRequestOperationTypeEnumValues() []WorkRequestOperationTypeEnum {
	values := make([]WorkRequestOperationTypeEnum, 0)
	for _, v := range mappingWorkRequestOperationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetWorkRequestOperationTypeEnumStringValues Enumerates the set of values in String for WorkRequestOperationTypeEnum
func GetWorkRequestOperationTypeEnumStringValues() []string {
	return []string{
		"CREATE_FUSION_ENVIRONMENT",
		"UPDATE_FUSION_ENVIRONMENT",
		"RESET_FUSION_ENVIRONMENT_ADMIN_PASSWORD",
		"SCALE_FUSION_ENVIRONMENT",
		"ARCHIVE_FUSION_ENVIRONMENT",
		"RESTORE_FUSION_ENVIRONMENT",
		"CREATE_SERVICE_INSTANCE",
		"UPDATE_SERVICE_INSTANCE",
		"DETACH_SERVICE_INSTANCE",
		"ADD_USER",
		"REMOVE_USER",
		"DELETE_FUSION_ENVIRONMENT",
		"CHANGE_FUSION_ENVIRONMENT_COMPARTMENT",
		"UPGRADE_FUSION_ENVIRONMENT",
		"CREATE_FUSION_ENVIRONMENT_FAMILY",
		"DELETE_FUSION_ENVIRONMENT_FAMILY",
		"UPDATE_FUSION_ENVIRONMENT_FAMILY",
		"CHANGE_FUSION_ENVIRONMENT_FAMILY_COMPARTMENT",
		"REFRESH_FUSION_ENVIRONMENT",
		"EXECUTE_COLD_PATCH",
		"DATA_MASK_FUSION_ENVIRONMENT",
		"INITIATE_EXTRACT",
		"SUBSCRIPTION_SUSPEND",
		"SUBSCRIPTION_EXPIRE",
		"SUBSCRIPTION_UPDATE",
		"SUBSCRIPTION_RESUME",
		"SUBSCRIPTION_TERMINATE",
	}
}

// GetMappingWorkRequestOperationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingWorkRequestOperationTypeEnum(val string) (WorkRequestOperationTypeEnum, bool) {
	enum, ok := mappingWorkRequestOperationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// WorkRequestStatusEnum Enum with underlying type: string
type WorkRequestStatusEnum string

// Set of constants representing the allowable values for WorkRequestStatusEnum
const (
	WorkRequestStatusAccepted   WorkRequestStatusEnum = "ACCEPTED"
	WorkRequestStatusInProgress WorkRequestStatusEnum = "IN_PROGRESS"
	WorkRequestStatusFailed     WorkRequestStatusEnum = "FAILED"
	WorkRequestStatusSucceeded  WorkRequestStatusEnum = "SUCCEEDED"
	WorkRequestStatusCanceling  WorkRequestStatusEnum = "CANCELING"
	WorkRequestStatusCanceled   WorkRequestStatusEnum = "CANCELED"
)

var mappingWorkRequestStatusEnum = map[string]WorkRequestStatusEnum{
	"ACCEPTED":    WorkRequestStatusAccepted,
	"IN_PROGRESS": WorkRequestStatusInProgress,
	"FAILED":      WorkRequestStatusFailed,
	"SUCCEEDED":   WorkRequestStatusSucceeded,
	"CANCELING":   WorkRequestStatusCanceling,
	"CANCELED":    WorkRequestStatusCanceled,
}

var mappingWorkRequestStatusEnumLowerCase = map[string]WorkRequestStatusEnum{
	"accepted":    WorkRequestStatusAccepted,
	"in_progress": WorkRequestStatusInProgress,
	"failed":      WorkRequestStatusFailed,
	"succeeded":   WorkRequestStatusSucceeded,
	"canceling":   WorkRequestStatusCanceling,
	"canceled":    WorkRequestStatusCanceled,
}

// GetWorkRequestStatusEnumValues Enumerates the set of values for WorkRequestStatusEnum
func GetWorkRequestStatusEnumValues() []WorkRequestStatusEnum {
	values := make([]WorkRequestStatusEnum, 0)
	for _, v := range mappingWorkRequestStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetWorkRequestStatusEnumStringValues Enumerates the set of values in String for WorkRequestStatusEnum
func GetWorkRequestStatusEnumStringValues() []string {
	return []string{
		"ACCEPTED",
		"IN_PROGRESS",
		"FAILED",
		"SUCCEEDED",
		"CANCELING",
		"CANCELED",
	}
}

// GetMappingWorkRequestStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingWorkRequestStatusEnum(val string) (WorkRequestStatusEnum, bool) {
	enum, ok := mappingWorkRequestStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
