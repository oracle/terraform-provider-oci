// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Digital Assistant Service Instance API
//
// API to create and maintain Oracle Digital Assistant service instances.
//

package oda

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// WorkRequest The description of work request, including its status.
type WorkRequest struct {

	// The identifier of the work request.
	Id *string `mandatory:"true" json:"id"`

	// The identifier of the compartment that contains the work request.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The identifier of the Digital Assistant instance to which this work request pertains.
	OdaInstanceId *string `mandatory:"true" json:"odaInstanceId"`

	// The identifier of the resource to which this work request pertains.
	ResourceId *string `mandatory:"true" json:"resourceId"`

	// The type of the operation that's associated with the work request.
	RequestAction WorkRequestRequestActionEnum `mandatory:"true" json:"requestAction"`

	// The status of current work request.
	Status WorkRequestStatusEnum `mandatory:"true" json:"status"`

	// The resources that this work request affects.
	Resources []WorkRequestResource `mandatory:"true" json:"resources"`

	// Percentage of the request completed.
	PercentComplete *float32 `mandatory:"true" json:"percentComplete"`

	// The date and time that the request was created, as described in
	// RFC 3339 (https://tools.ietf.org/rfc/rfc3339), section 14.29.
	TimeAccepted *common.SDKTime `mandatory:"true" json:"timeAccepted"`

	// A short message that provides more detail about the current status.
	// For example, if a work request fails, then this may include information
	// about why it failed.
	StatusMessage *string `mandatory:"false" json:"statusMessage"`

	// The date and time that the request was started, as described in RFC 3339 (https://tools.ietf.org/rfc/rfc3339), CKQ
	// section 14.29.
	TimeStarted *common.SDKTime `mandatory:"false" json:"timeStarted"`

	// The date and time that the object finished, as described in RFC 3339 (https://tools.ietf.org/rfc/rfc3339). CKQ
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
	if _, ok := GetMappingWorkRequestRequestActionEnum(string(m.RequestAction)); !ok && m.RequestAction != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RequestAction: %s. Supported values are: %s.", m.RequestAction, strings.Join(GetWorkRequestRequestActionEnumStringValues(), ",")))
	}
	if _, ok := GetMappingWorkRequestStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetWorkRequestStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// WorkRequestRequestActionEnum Enum with underlying type: string
type WorkRequestRequestActionEnum string

// Set of constants representing the allowable values for WorkRequestRequestActionEnum
const (
	WorkRequestRequestActionCreateOdaInstance                   WorkRequestRequestActionEnum = "CREATE_ODA_INSTANCE"
	WorkRequestRequestActionUpgradeOdaInstance                  WorkRequestRequestActionEnum = "UPGRADE_ODA_INSTANCE"
	WorkRequestRequestActionDeleteOdaInstance                   WorkRequestRequestActionEnum = "DELETE_ODA_INSTANCE"
	WorkRequestRequestActionPurgeOdaInstance                    WorkRequestRequestActionEnum = "PURGE_ODA_INSTANCE"
	WorkRequestRequestActionRecoverOdaInstance                  WorkRequestRequestActionEnum = "RECOVER_ODA_INSTANCE"
	WorkRequestRequestActionStopOdaInstance                     WorkRequestRequestActionEnum = "STOP_ODA_INSTANCE"
	WorkRequestRequestActionStartOdaInstance                    WorkRequestRequestActionEnum = "START_ODA_INSTANCE"
	WorkRequestRequestActionChangeOdaInstanceCompartment        WorkRequestRequestActionEnum = "CHANGE_ODA_INSTANCE_COMPARTMENT"
	WorkRequestRequestActionChangeCustEncKey                    WorkRequestRequestActionEnum = "CHANGE_CUST_ENC_KEY"
	WorkRequestRequestActionDeactCustEncKey                     WorkRequestRequestActionEnum = "DEACT_CUST_ENC_KEY"
	WorkRequestRequestActionCreateAssociation                   WorkRequestRequestActionEnum = "CREATE_ASSOCIATION"
	WorkRequestRequestActionDeleteAssociation                   WorkRequestRequestActionEnum = "DELETE_ASSOCIATION"
	WorkRequestRequestActionCreatePcsInstance                   WorkRequestRequestActionEnum = "CREATE_PCS_INSTANCE"
	WorkRequestRequestActionUpdateEntitlementsForCacct          WorkRequestRequestActionEnum = "UPDATE_ENTITLEMENTS_FOR_CACCT"
	WorkRequestRequestActionLookupOdaInstancesForCacct          WorkRequestRequestActionEnum = "LOOKUP_ODA_INSTANCES_FOR_CACCT"
	WorkRequestRequestActionCreateOdaInstanceAttachment         WorkRequestRequestActionEnum = "CREATE_ODA_INSTANCE_ATTACHMENT"
	WorkRequestRequestActionUpdateOdaInstanceAttachment         WorkRequestRequestActionEnum = "UPDATE_ODA_INSTANCE_ATTACHMENT"
	WorkRequestRequestActionDeleteOdaInstanceAttachment         WorkRequestRequestActionEnum = "DELETE_ODA_INSTANCE_ATTACHMENT"
	WorkRequestRequestActionCreateImportedPackage               WorkRequestRequestActionEnum = "CREATE_IMPORTED_PACKAGE"
	WorkRequestRequestActionUpdateImportedPackage               WorkRequestRequestActionEnum = "UPDATE_IMPORTED_PACKAGE"
	WorkRequestRequestActionDeleteImportedPackage               WorkRequestRequestActionEnum = "DELETE_IMPORTED_PACKAGE"
	WorkRequestRequestActionImportBot                           WorkRequestRequestActionEnum = "IMPORT_BOT"
	WorkRequestRequestActionCreateSkill                         WorkRequestRequestActionEnum = "CREATE_SKILL"
	WorkRequestRequestActionCloneSkill                          WorkRequestRequestActionEnum = "CLONE_SKILL"
	WorkRequestRequestActionExtendSkill                         WorkRequestRequestActionEnum = "EXTEND_SKILL"
	WorkRequestRequestActionVersionSkill                        WorkRequestRequestActionEnum = "VERSION_SKILL"
	WorkRequestRequestActionExportSkill                         WorkRequestRequestActionEnum = "EXPORT_SKILL"
	WorkRequestRequestActionCreateDigitalAssistant              WorkRequestRequestActionEnum = "CREATE_DIGITAL_ASSISTANT"
	WorkRequestRequestActionCloneDigitalAssistant               WorkRequestRequestActionEnum = "CLONE_DIGITAL_ASSISTANT"
	WorkRequestRequestActionExtendDigitalAssistant              WorkRequestRequestActionEnum = "EXTEND_DIGITAL_ASSISTANT"
	WorkRequestRequestActionVersionDigitalAssistant             WorkRequestRequestActionEnum = "VERSION_DIGITAL_ASSISTANT"
	WorkRequestRequestActionExportDigitalAssistant              WorkRequestRequestActionEnum = "EXPORT_DIGITAL_ASSISTANT"
	WorkRequestRequestActionCreateOdaPrivateEndpoint            WorkRequestRequestActionEnum = "CREATE_ODA_PRIVATE_ENDPOINT"
	WorkRequestRequestActionDeleteOdaPrivateEndpoint            WorkRequestRequestActionEnum = "DELETE_ODA_PRIVATE_ENDPOINT"
	WorkRequestRequestActionUpdateOdaPrivateEndpoint            WorkRequestRequestActionEnum = "UPDATE_ODA_PRIVATE_ENDPOINT"
	WorkRequestRequestActionChangeOdaPrivateEndpointCompartment WorkRequestRequestActionEnum = "CHANGE_ODA_PRIVATE_ENDPOINT_COMPARTMENT"
	WorkRequestRequestActionCreateOdaPrivateEndpointScanProxy   WorkRequestRequestActionEnum = "CREATE_ODA_PRIVATE_ENDPOINT_SCAN_PROXY"
	WorkRequestRequestActionDeleteOdaPrivateEndpointScanProxy   WorkRequestRequestActionEnum = "DELETE_ODA_PRIVATE_ENDPOINT_SCAN_PROXY"
	WorkRequestRequestActionCreateOdaPrivateEndpointAttachment  WorkRequestRequestActionEnum = "CREATE_ODA_PRIVATE_ENDPOINT_ATTACHMENT"
	WorkRequestRequestActionDeleteOdaPrivateEndpointAttachment  WorkRequestRequestActionEnum = "DELETE_ODA_PRIVATE_ENDPOINT_ATTACHMENT"
)

var mappingWorkRequestRequestActionEnum = map[string]WorkRequestRequestActionEnum{
	"CREATE_ODA_INSTANCE":                     WorkRequestRequestActionCreateOdaInstance,
	"UPGRADE_ODA_INSTANCE":                    WorkRequestRequestActionUpgradeOdaInstance,
	"DELETE_ODA_INSTANCE":                     WorkRequestRequestActionDeleteOdaInstance,
	"PURGE_ODA_INSTANCE":                      WorkRequestRequestActionPurgeOdaInstance,
	"RECOVER_ODA_INSTANCE":                    WorkRequestRequestActionRecoverOdaInstance,
	"STOP_ODA_INSTANCE":                       WorkRequestRequestActionStopOdaInstance,
	"START_ODA_INSTANCE":                      WorkRequestRequestActionStartOdaInstance,
	"CHANGE_ODA_INSTANCE_COMPARTMENT":         WorkRequestRequestActionChangeOdaInstanceCompartment,
	"CHANGE_CUST_ENC_KEY":                     WorkRequestRequestActionChangeCustEncKey,
	"DEACT_CUST_ENC_KEY":                      WorkRequestRequestActionDeactCustEncKey,
	"CREATE_ASSOCIATION":                      WorkRequestRequestActionCreateAssociation,
	"DELETE_ASSOCIATION":                      WorkRequestRequestActionDeleteAssociation,
	"CREATE_PCS_INSTANCE":                     WorkRequestRequestActionCreatePcsInstance,
	"UPDATE_ENTITLEMENTS_FOR_CACCT":           WorkRequestRequestActionUpdateEntitlementsForCacct,
	"LOOKUP_ODA_INSTANCES_FOR_CACCT":          WorkRequestRequestActionLookupOdaInstancesForCacct,
	"CREATE_ODA_INSTANCE_ATTACHMENT":          WorkRequestRequestActionCreateOdaInstanceAttachment,
	"UPDATE_ODA_INSTANCE_ATTACHMENT":          WorkRequestRequestActionUpdateOdaInstanceAttachment,
	"DELETE_ODA_INSTANCE_ATTACHMENT":          WorkRequestRequestActionDeleteOdaInstanceAttachment,
	"CREATE_IMPORTED_PACKAGE":                 WorkRequestRequestActionCreateImportedPackage,
	"UPDATE_IMPORTED_PACKAGE":                 WorkRequestRequestActionUpdateImportedPackage,
	"DELETE_IMPORTED_PACKAGE":                 WorkRequestRequestActionDeleteImportedPackage,
	"IMPORT_BOT":                              WorkRequestRequestActionImportBot,
	"CREATE_SKILL":                            WorkRequestRequestActionCreateSkill,
	"CLONE_SKILL":                             WorkRequestRequestActionCloneSkill,
	"EXTEND_SKILL":                            WorkRequestRequestActionExtendSkill,
	"VERSION_SKILL":                           WorkRequestRequestActionVersionSkill,
	"EXPORT_SKILL":                            WorkRequestRequestActionExportSkill,
	"CREATE_DIGITAL_ASSISTANT":                WorkRequestRequestActionCreateDigitalAssistant,
	"CLONE_DIGITAL_ASSISTANT":                 WorkRequestRequestActionCloneDigitalAssistant,
	"EXTEND_DIGITAL_ASSISTANT":                WorkRequestRequestActionExtendDigitalAssistant,
	"VERSION_DIGITAL_ASSISTANT":               WorkRequestRequestActionVersionDigitalAssistant,
	"EXPORT_DIGITAL_ASSISTANT":                WorkRequestRequestActionExportDigitalAssistant,
	"CREATE_ODA_PRIVATE_ENDPOINT":             WorkRequestRequestActionCreateOdaPrivateEndpoint,
	"DELETE_ODA_PRIVATE_ENDPOINT":             WorkRequestRequestActionDeleteOdaPrivateEndpoint,
	"UPDATE_ODA_PRIVATE_ENDPOINT":             WorkRequestRequestActionUpdateOdaPrivateEndpoint,
	"CHANGE_ODA_PRIVATE_ENDPOINT_COMPARTMENT": WorkRequestRequestActionChangeOdaPrivateEndpointCompartment,
	"CREATE_ODA_PRIVATE_ENDPOINT_SCAN_PROXY":  WorkRequestRequestActionCreateOdaPrivateEndpointScanProxy,
	"DELETE_ODA_PRIVATE_ENDPOINT_SCAN_PROXY":  WorkRequestRequestActionDeleteOdaPrivateEndpointScanProxy,
	"CREATE_ODA_PRIVATE_ENDPOINT_ATTACHMENT":  WorkRequestRequestActionCreateOdaPrivateEndpointAttachment,
	"DELETE_ODA_PRIVATE_ENDPOINT_ATTACHMENT":  WorkRequestRequestActionDeleteOdaPrivateEndpointAttachment,
}

var mappingWorkRequestRequestActionEnumLowerCase = map[string]WorkRequestRequestActionEnum{
	"create_oda_instance":                     WorkRequestRequestActionCreateOdaInstance,
	"upgrade_oda_instance":                    WorkRequestRequestActionUpgradeOdaInstance,
	"delete_oda_instance":                     WorkRequestRequestActionDeleteOdaInstance,
	"purge_oda_instance":                      WorkRequestRequestActionPurgeOdaInstance,
	"recover_oda_instance":                    WorkRequestRequestActionRecoverOdaInstance,
	"stop_oda_instance":                       WorkRequestRequestActionStopOdaInstance,
	"start_oda_instance":                      WorkRequestRequestActionStartOdaInstance,
	"change_oda_instance_compartment":         WorkRequestRequestActionChangeOdaInstanceCompartment,
	"change_cust_enc_key":                     WorkRequestRequestActionChangeCustEncKey,
	"deact_cust_enc_key":                      WorkRequestRequestActionDeactCustEncKey,
	"create_association":                      WorkRequestRequestActionCreateAssociation,
	"delete_association":                      WorkRequestRequestActionDeleteAssociation,
	"create_pcs_instance":                     WorkRequestRequestActionCreatePcsInstance,
	"update_entitlements_for_cacct":           WorkRequestRequestActionUpdateEntitlementsForCacct,
	"lookup_oda_instances_for_cacct":          WorkRequestRequestActionLookupOdaInstancesForCacct,
	"create_oda_instance_attachment":          WorkRequestRequestActionCreateOdaInstanceAttachment,
	"update_oda_instance_attachment":          WorkRequestRequestActionUpdateOdaInstanceAttachment,
	"delete_oda_instance_attachment":          WorkRequestRequestActionDeleteOdaInstanceAttachment,
	"create_imported_package":                 WorkRequestRequestActionCreateImportedPackage,
	"update_imported_package":                 WorkRequestRequestActionUpdateImportedPackage,
	"delete_imported_package":                 WorkRequestRequestActionDeleteImportedPackage,
	"import_bot":                              WorkRequestRequestActionImportBot,
	"create_skill":                            WorkRequestRequestActionCreateSkill,
	"clone_skill":                             WorkRequestRequestActionCloneSkill,
	"extend_skill":                            WorkRequestRequestActionExtendSkill,
	"version_skill":                           WorkRequestRequestActionVersionSkill,
	"export_skill":                            WorkRequestRequestActionExportSkill,
	"create_digital_assistant":                WorkRequestRequestActionCreateDigitalAssistant,
	"clone_digital_assistant":                 WorkRequestRequestActionCloneDigitalAssistant,
	"extend_digital_assistant":                WorkRequestRequestActionExtendDigitalAssistant,
	"version_digital_assistant":               WorkRequestRequestActionVersionDigitalAssistant,
	"export_digital_assistant":                WorkRequestRequestActionExportDigitalAssistant,
	"create_oda_private_endpoint":             WorkRequestRequestActionCreateOdaPrivateEndpoint,
	"delete_oda_private_endpoint":             WorkRequestRequestActionDeleteOdaPrivateEndpoint,
	"update_oda_private_endpoint":             WorkRequestRequestActionUpdateOdaPrivateEndpoint,
	"change_oda_private_endpoint_compartment": WorkRequestRequestActionChangeOdaPrivateEndpointCompartment,
	"create_oda_private_endpoint_scan_proxy":  WorkRequestRequestActionCreateOdaPrivateEndpointScanProxy,
	"delete_oda_private_endpoint_scan_proxy":  WorkRequestRequestActionDeleteOdaPrivateEndpointScanProxy,
	"create_oda_private_endpoint_attachment":  WorkRequestRequestActionCreateOdaPrivateEndpointAttachment,
	"delete_oda_private_endpoint_attachment":  WorkRequestRequestActionDeleteOdaPrivateEndpointAttachment,
}

// GetWorkRequestRequestActionEnumValues Enumerates the set of values for WorkRequestRequestActionEnum
func GetWorkRequestRequestActionEnumValues() []WorkRequestRequestActionEnum {
	values := make([]WorkRequestRequestActionEnum, 0)
	for _, v := range mappingWorkRequestRequestActionEnum {
		values = append(values, v)
	}
	return values
}

// GetWorkRequestRequestActionEnumStringValues Enumerates the set of values in String for WorkRequestRequestActionEnum
func GetWorkRequestRequestActionEnumStringValues() []string {
	return []string{
		"CREATE_ODA_INSTANCE",
		"UPGRADE_ODA_INSTANCE",
		"DELETE_ODA_INSTANCE",
		"PURGE_ODA_INSTANCE",
		"RECOVER_ODA_INSTANCE",
		"STOP_ODA_INSTANCE",
		"START_ODA_INSTANCE",
		"CHANGE_ODA_INSTANCE_COMPARTMENT",
		"CHANGE_CUST_ENC_KEY",
		"DEACT_CUST_ENC_KEY",
		"CREATE_ASSOCIATION",
		"DELETE_ASSOCIATION",
		"CREATE_PCS_INSTANCE",
		"UPDATE_ENTITLEMENTS_FOR_CACCT",
		"LOOKUP_ODA_INSTANCES_FOR_CACCT",
		"CREATE_ODA_INSTANCE_ATTACHMENT",
		"UPDATE_ODA_INSTANCE_ATTACHMENT",
		"DELETE_ODA_INSTANCE_ATTACHMENT",
		"CREATE_IMPORTED_PACKAGE",
		"UPDATE_IMPORTED_PACKAGE",
		"DELETE_IMPORTED_PACKAGE",
		"IMPORT_BOT",
		"CREATE_SKILL",
		"CLONE_SKILL",
		"EXTEND_SKILL",
		"VERSION_SKILL",
		"EXPORT_SKILL",
		"CREATE_DIGITAL_ASSISTANT",
		"CLONE_DIGITAL_ASSISTANT",
		"EXTEND_DIGITAL_ASSISTANT",
		"VERSION_DIGITAL_ASSISTANT",
		"EXPORT_DIGITAL_ASSISTANT",
		"CREATE_ODA_PRIVATE_ENDPOINT",
		"DELETE_ODA_PRIVATE_ENDPOINT",
		"UPDATE_ODA_PRIVATE_ENDPOINT",
		"CHANGE_ODA_PRIVATE_ENDPOINT_COMPARTMENT",
		"CREATE_ODA_PRIVATE_ENDPOINT_SCAN_PROXY",
		"DELETE_ODA_PRIVATE_ENDPOINT_SCAN_PROXY",
		"CREATE_ODA_PRIVATE_ENDPOINT_ATTACHMENT",
		"DELETE_ODA_PRIVATE_ENDPOINT_ATTACHMENT",
	}
}

// GetMappingWorkRequestRequestActionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingWorkRequestRequestActionEnum(val string) (WorkRequestRequestActionEnum, bool) {
	enum, ok := mappingWorkRequestRequestActionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// WorkRequestStatusEnum Enum with underlying type: string
type WorkRequestStatusEnum string

// Set of constants representing the allowable values for WorkRequestStatusEnum
const (
	WorkRequestStatusAccepted   WorkRequestStatusEnum = "ACCEPTED"
	WorkRequestStatusInProgress WorkRequestStatusEnum = "IN_PROGRESS"
	WorkRequestStatusSucceeded  WorkRequestStatusEnum = "SUCCEEDED"
	WorkRequestStatusFailed     WorkRequestStatusEnum = "FAILED"
	WorkRequestStatusCanceling  WorkRequestStatusEnum = "CANCELING"
	WorkRequestStatusCanceled   WorkRequestStatusEnum = "CANCELED"
)

var mappingWorkRequestStatusEnum = map[string]WorkRequestStatusEnum{
	"ACCEPTED":    WorkRequestStatusAccepted,
	"IN_PROGRESS": WorkRequestStatusInProgress,
	"SUCCEEDED":   WorkRequestStatusSucceeded,
	"FAILED":      WorkRequestStatusFailed,
	"CANCELING":   WorkRequestStatusCanceling,
	"CANCELED":    WorkRequestStatusCanceled,
}

var mappingWorkRequestStatusEnumLowerCase = map[string]WorkRequestStatusEnum{
	"accepted":    WorkRequestStatusAccepted,
	"in_progress": WorkRequestStatusInProgress,
	"succeeded":   WorkRequestStatusSucceeded,
	"failed":      WorkRequestStatusFailed,
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
		"SUCCEEDED",
		"FAILED",
		"CANCELING",
		"CANCELED",
	}
}

// GetMappingWorkRequestStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingWorkRequestStatusEnum(val string) (WorkRequestStatusEnum, bool) {
	enum, ok := mappingWorkRequestStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
