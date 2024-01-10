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

// WorkRequestSummary A description of the work request's status.
type WorkRequestSummary struct {

	// The identifier of the work request.
	Id *string `mandatory:"true" json:"id"`

	// The identifier of the compartment that contains the work request.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The identifier of the Digital Assistant instance to which this work request pertains.
	OdaInstanceId *string `mandatory:"true" json:"odaInstanceId"`

	// The identifier of the resource to which this work request pertains.
	ResourceId *string `mandatory:"true" json:"resourceId"`

	// The type of the operation that's associated with the work request.
	RequestAction WorkRequestSummaryRequestActionEnum `mandatory:"true" json:"requestAction"`

	// The status of current work request.
	Status WorkRequestSummaryStatusEnum `mandatory:"true" json:"status"`

	// The resources that this work request affects.
	Resources []WorkRequestResource `mandatory:"true" json:"resources"`
}

func (m WorkRequestSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m WorkRequestSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingWorkRequestSummaryRequestActionEnum(string(m.RequestAction)); !ok && m.RequestAction != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RequestAction: %s. Supported values are: %s.", m.RequestAction, strings.Join(GetWorkRequestSummaryRequestActionEnumStringValues(), ",")))
	}
	if _, ok := GetMappingWorkRequestSummaryStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetWorkRequestSummaryStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// WorkRequestSummaryRequestActionEnum Enum with underlying type: string
type WorkRequestSummaryRequestActionEnum string

// Set of constants representing the allowable values for WorkRequestSummaryRequestActionEnum
const (
	WorkRequestSummaryRequestActionCreateOdaInstance                   WorkRequestSummaryRequestActionEnum = "CREATE_ODA_INSTANCE"
	WorkRequestSummaryRequestActionUpgradeOdaInstance                  WorkRequestSummaryRequestActionEnum = "UPGRADE_ODA_INSTANCE"
	WorkRequestSummaryRequestActionDeleteOdaInstance                   WorkRequestSummaryRequestActionEnum = "DELETE_ODA_INSTANCE"
	WorkRequestSummaryRequestActionPurgeOdaInstance                    WorkRequestSummaryRequestActionEnum = "PURGE_ODA_INSTANCE"
	WorkRequestSummaryRequestActionRecoverOdaInstance                  WorkRequestSummaryRequestActionEnum = "RECOVER_ODA_INSTANCE"
	WorkRequestSummaryRequestActionStopOdaInstance                     WorkRequestSummaryRequestActionEnum = "STOP_ODA_INSTANCE"
	WorkRequestSummaryRequestActionStartOdaInstance                    WorkRequestSummaryRequestActionEnum = "START_ODA_INSTANCE"
	WorkRequestSummaryRequestActionChangeOdaInstanceCompartment        WorkRequestSummaryRequestActionEnum = "CHANGE_ODA_INSTANCE_COMPARTMENT"
	WorkRequestSummaryRequestActionChangeCustEncKey                    WorkRequestSummaryRequestActionEnum = "CHANGE_CUST_ENC_KEY"
	WorkRequestSummaryRequestActionDeactCustEncKey                     WorkRequestSummaryRequestActionEnum = "DEACT_CUST_ENC_KEY"
	WorkRequestSummaryRequestActionCreateAssociation                   WorkRequestSummaryRequestActionEnum = "CREATE_ASSOCIATION"
	WorkRequestSummaryRequestActionDeleteAssociation                   WorkRequestSummaryRequestActionEnum = "DELETE_ASSOCIATION"
	WorkRequestSummaryRequestActionUpdateEntitlementsForCacct          WorkRequestSummaryRequestActionEnum = "UPDATE_ENTITLEMENTS_FOR_CACCT"
	WorkRequestSummaryRequestActionLookupOdaInstancesForCacct          WorkRequestSummaryRequestActionEnum = "LOOKUP_ODA_INSTANCES_FOR_CACCT"
	WorkRequestSummaryRequestActionCreateOdaInstanceAttachment         WorkRequestSummaryRequestActionEnum = "CREATE_ODA_INSTANCE_ATTACHMENT"
	WorkRequestSummaryRequestActionUpdateOdaInstanceAttachment         WorkRequestSummaryRequestActionEnum = "UPDATE_ODA_INSTANCE_ATTACHMENT"
	WorkRequestSummaryRequestActionDeleteOdaInstanceAttachment         WorkRequestSummaryRequestActionEnum = "DELETE_ODA_INSTANCE_ATTACHMENT"
	WorkRequestSummaryRequestActionCreateImportedPackage               WorkRequestSummaryRequestActionEnum = "CREATE_IMPORTED_PACKAGE"
	WorkRequestSummaryRequestActionUpdateImportedPackage               WorkRequestSummaryRequestActionEnum = "UPDATE_IMPORTED_PACKAGE"
	WorkRequestSummaryRequestActionDeleteImportedPackage               WorkRequestSummaryRequestActionEnum = "DELETE_IMPORTED_PACKAGE"
	WorkRequestSummaryRequestActionImportBot                           WorkRequestSummaryRequestActionEnum = "IMPORT_BOT"
	WorkRequestSummaryRequestActionCreateSkill                         WorkRequestSummaryRequestActionEnum = "CREATE_SKILL"
	WorkRequestSummaryRequestActionCloneSkill                          WorkRequestSummaryRequestActionEnum = "CLONE_SKILL"
	WorkRequestSummaryRequestActionExtendSkill                         WorkRequestSummaryRequestActionEnum = "EXTEND_SKILL"
	WorkRequestSummaryRequestActionVersionSkill                        WorkRequestSummaryRequestActionEnum = "VERSION_SKILL"
	WorkRequestSummaryRequestActionExportSkill                         WorkRequestSummaryRequestActionEnum = "EXPORT_SKILL"
	WorkRequestSummaryRequestActionCreateDigitalAssistant              WorkRequestSummaryRequestActionEnum = "CREATE_DIGITAL_ASSISTANT"
	WorkRequestSummaryRequestActionCloneDigitalAssistant               WorkRequestSummaryRequestActionEnum = "CLONE_DIGITAL_ASSISTANT"
	WorkRequestSummaryRequestActionExtendDigitalAssistant              WorkRequestSummaryRequestActionEnum = "EXTEND_DIGITAL_ASSISTANT"
	WorkRequestSummaryRequestActionVersionDigitalAssistant             WorkRequestSummaryRequestActionEnum = "VERSION_DIGITAL_ASSISTANT"
	WorkRequestSummaryRequestActionExportDigitalAssistant              WorkRequestSummaryRequestActionEnum = "EXPORT_DIGITAL_ASSISTANT"
	WorkRequestSummaryRequestActionCreateOdaPrivateEndpoint            WorkRequestSummaryRequestActionEnum = "CREATE_ODA_PRIVATE_ENDPOINT"
	WorkRequestSummaryRequestActionDeleteOdaPrivateEndpoint            WorkRequestSummaryRequestActionEnum = "DELETE_ODA_PRIVATE_ENDPOINT"
	WorkRequestSummaryRequestActionUpdateOdaPrivateEndpoint            WorkRequestSummaryRequestActionEnum = "UPDATE_ODA_PRIVATE_ENDPOINT"
	WorkRequestSummaryRequestActionChangeOdaPrivateEndpointCompartment WorkRequestSummaryRequestActionEnum = "CHANGE_ODA_PRIVATE_ENDPOINT_COMPARTMENT"
	WorkRequestSummaryRequestActionCreateOdaPrivateEndpointScanProxy   WorkRequestSummaryRequestActionEnum = "CREATE_ODA_PRIVATE_ENDPOINT_SCAN_PROXY"
	WorkRequestSummaryRequestActionDeleteOdaPrivateEndpointScanProxy   WorkRequestSummaryRequestActionEnum = "DELETE_ODA_PRIVATE_ENDPOINT_SCAN_PROXY"
	WorkRequestSummaryRequestActionCreateOdaPrivateEndpointAttachment  WorkRequestSummaryRequestActionEnum = "CREATE_ODA_PRIVATE_ENDPOINT_ATTACHMENT"
	WorkRequestSummaryRequestActionDeleteOdaPrivateEndpointAttachment  WorkRequestSummaryRequestActionEnum = "DELETE_ODA_PRIVATE_ENDPOINT_ATTACHMENT"
)

var mappingWorkRequestSummaryRequestActionEnum = map[string]WorkRequestSummaryRequestActionEnum{
	"CREATE_ODA_INSTANCE":                     WorkRequestSummaryRequestActionCreateOdaInstance,
	"UPGRADE_ODA_INSTANCE":                    WorkRequestSummaryRequestActionUpgradeOdaInstance,
	"DELETE_ODA_INSTANCE":                     WorkRequestSummaryRequestActionDeleteOdaInstance,
	"PURGE_ODA_INSTANCE":                      WorkRequestSummaryRequestActionPurgeOdaInstance,
	"RECOVER_ODA_INSTANCE":                    WorkRequestSummaryRequestActionRecoverOdaInstance,
	"STOP_ODA_INSTANCE":                       WorkRequestSummaryRequestActionStopOdaInstance,
	"START_ODA_INSTANCE":                      WorkRequestSummaryRequestActionStartOdaInstance,
	"CHANGE_ODA_INSTANCE_COMPARTMENT":         WorkRequestSummaryRequestActionChangeOdaInstanceCompartment,
	"CHANGE_CUST_ENC_KEY":                     WorkRequestSummaryRequestActionChangeCustEncKey,
	"DEACT_CUST_ENC_KEY":                      WorkRequestSummaryRequestActionDeactCustEncKey,
	"CREATE_ASSOCIATION":                      WorkRequestSummaryRequestActionCreateAssociation,
	"DELETE_ASSOCIATION":                      WorkRequestSummaryRequestActionDeleteAssociation,
	"UPDATE_ENTITLEMENTS_FOR_CACCT":           WorkRequestSummaryRequestActionUpdateEntitlementsForCacct,
	"LOOKUP_ODA_INSTANCES_FOR_CACCT":          WorkRequestSummaryRequestActionLookupOdaInstancesForCacct,
	"CREATE_ODA_INSTANCE_ATTACHMENT":          WorkRequestSummaryRequestActionCreateOdaInstanceAttachment,
	"UPDATE_ODA_INSTANCE_ATTACHMENT":          WorkRequestSummaryRequestActionUpdateOdaInstanceAttachment,
	"DELETE_ODA_INSTANCE_ATTACHMENT":          WorkRequestSummaryRequestActionDeleteOdaInstanceAttachment,
	"CREATE_IMPORTED_PACKAGE":                 WorkRequestSummaryRequestActionCreateImportedPackage,
	"UPDATE_IMPORTED_PACKAGE":                 WorkRequestSummaryRequestActionUpdateImportedPackage,
	"DELETE_IMPORTED_PACKAGE":                 WorkRequestSummaryRequestActionDeleteImportedPackage,
	"IMPORT_BOT":                              WorkRequestSummaryRequestActionImportBot,
	"CREATE_SKILL":                            WorkRequestSummaryRequestActionCreateSkill,
	"CLONE_SKILL":                             WorkRequestSummaryRequestActionCloneSkill,
	"EXTEND_SKILL":                            WorkRequestSummaryRequestActionExtendSkill,
	"VERSION_SKILL":                           WorkRequestSummaryRequestActionVersionSkill,
	"EXPORT_SKILL":                            WorkRequestSummaryRequestActionExportSkill,
	"CREATE_DIGITAL_ASSISTANT":                WorkRequestSummaryRequestActionCreateDigitalAssistant,
	"CLONE_DIGITAL_ASSISTANT":                 WorkRequestSummaryRequestActionCloneDigitalAssistant,
	"EXTEND_DIGITAL_ASSISTANT":                WorkRequestSummaryRequestActionExtendDigitalAssistant,
	"VERSION_DIGITAL_ASSISTANT":               WorkRequestSummaryRequestActionVersionDigitalAssistant,
	"EXPORT_DIGITAL_ASSISTANT":                WorkRequestSummaryRequestActionExportDigitalAssistant,
	"CREATE_ODA_PRIVATE_ENDPOINT":             WorkRequestSummaryRequestActionCreateOdaPrivateEndpoint,
	"DELETE_ODA_PRIVATE_ENDPOINT":             WorkRequestSummaryRequestActionDeleteOdaPrivateEndpoint,
	"UPDATE_ODA_PRIVATE_ENDPOINT":             WorkRequestSummaryRequestActionUpdateOdaPrivateEndpoint,
	"CHANGE_ODA_PRIVATE_ENDPOINT_COMPARTMENT": WorkRequestSummaryRequestActionChangeOdaPrivateEndpointCompartment,
	"CREATE_ODA_PRIVATE_ENDPOINT_SCAN_PROXY":  WorkRequestSummaryRequestActionCreateOdaPrivateEndpointScanProxy,
	"DELETE_ODA_PRIVATE_ENDPOINT_SCAN_PROXY":  WorkRequestSummaryRequestActionDeleteOdaPrivateEndpointScanProxy,
	"CREATE_ODA_PRIVATE_ENDPOINT_ATTACHMENT":  WorkRequestSummaryRequestActionCreateOdaPrivateEndpointAttachment,
	"DELETE_ODA_PRIVATE_ENDPOINT_ATTACHMENT":  WorkRequestSummaryRequestActionDeleteOdaPrivateEndpointAttachment,
}

var mappingWorkRequestSummaryRequestActionEnumLowerCase = map[string]WorkRequestSummaryRequestActionEnum{
	"create_oda_instance":                     WorkRequestSummaryRequestActionCreateOdaInstance,
	"upgrade_oda_instance":                    WorkRequestSummaryRequestActionUpgradeOdaInstance,
	"delete_oda_instance":                     WorkRequestSummaryRequestActionDeleteOdaInstance,
	"purge_oda_instance":                      WorkRequestSummaryRequestActionPurgeOdaInstance,
	"recover_oda_instance":                    WorkRequestSummaryRequestActionRecoverOdaInstance,
	"stop_oda_instance":                       WorkRequestSummaryRequestActionStopOdaInstance,
	"start_oda_instance":                      WorkRequestSummaryRequestActionStartOdaInstance,
	"change_oda_instance_compartment":         WorkRequestSummaryRequestActionChangeOdaInstanceCompartment,
	"change_cust_enc_key":                     WorkRequestSummaryRequestActionChangeCustEncKey,
	"deact_cust_enc_key":                      WorkRequestSummaryRequestActionDeactCustEncKey,
	"create_association":                      WorkRequestSummaryRequestActionCreateAssociation,
	"delete_association":                      WorkRequestSummaryRequestActionDeleteAssociation,
	"update_entitlements_for_cacct":           WorkRequestSummaryRequestActionUpdateEntitlementsForCacct,
	"lookup_oda_instances_for_cacct":          WorkRequestSummaryRequestActionLookupOdaInstancesForCacct,
	"create_oda_instance_attachment":          WorkRequestSummaryRequestActionCreateOdaInstanceAttachment,
	"update_oda_instance_attachment":          WorkRequestSummaryRequestActionUpdateOdaInstanceAttachment,
	"delete_oda_instance_attachment":          WorkRequestSummaryRequestActionDeleteOdaInstanceAttachment,
	"create_imported_package":                 WorkRequestSummaryRequestActionCreateImportedPackage,
	"update_imported_package":                 WorkRequestSummaryRequestActionUpdateImportedPackage,
	"delete_imported_package":                 WorkRequestSummaryRequestActionDeleteImportedPackage,
	"import_bot":                              WorkRequestSummaryRequestActionImportBot,
	"create_skill":                            WorkRequestSummaryRequestActionCreateSkill,
	"clone_skill":                             WorkRequestSummaryRequestActionCloneSkill,
	"extend_skill":                            WorkRequestSummaryRequestActionExtendSkill,
	"version_skill":                           WorkRequestSummaryRequestActionVersionSkill,
	"export_skill":                            WorkRequestSummaryRequestActionExportSkill,
	"create_digital_assistant":                WorkRequestSummaryRequestActionCreateDigitalAssistant,
	"clone_digital_assistant":                 WorkRequestSummaryRequestActionCloneDigitalAssistant,
	"extend_digital_assistant":                WorkRequestSummaryRequestActionExtendDigitalAssistant,
	"version_digital_assistant":               WorkRequestSummaryRequestActionVersionDigitalAssistant,
	"export_digital_assistant":                WorkRequestSummaryRequestActionExportDigitalAssistant,
	"create_oda_private_endpoint":             WorkRequestSummaryRequestActionCreateOdaPrivateEndpoint,
	"delete_oda_private_endpoint":             WorkRequestSummaryRequestActionDeleteOdaPrivateEndpoint,
	"update_oda_private_endpoint":             WorkRequestSummaryRequestActionUpdateOdaPrivateEndpoint,
	"change_oda_private_endpoint_compartment": WorkRequestSummaryRequestActionChangeOdaPrivateEndpointCompartment,
	"create_oda_private_endpoint_scan_proxy":  WorkRequestSummaryRequestActionCreateOdaPrivateEndpointScanProxy,
	"delete_oda_private_endpoint_scan_proxy":  WorkRequestSummaryRequestActionDeleteOdaPrivateEndpointScanProxy,
	"create_oda_private_endpoint_attachment":  WorkRequestSummaryRequestActionCreateOdaPrivateEndpointAttachment,
	"delete_oda_private_endpoint_attachment":  WorkRequestSummaryRequestActionDeleteOdaPrivateEndpointAttachment,
}

// GetWorkRequestSummaryRequestActionEnumValues Enumerates the set of values for WorkRequestSummaryRequestActionEnum
func GetWorkRequestSummaryRequestActionEnumValues() []WorkRequestSummaryRequestActionEnum {
	values := make([]WorkRequestSummaryRequestActionEnum, 0)
	for _, v := range mappingWorkRequestSummaryRequestActionEnum {
		values = append(values, v)
	}
	return values
}

// GetWorkRequestSummaryRequestActionEnumStringValues Enumerates the set of values in String for WorkRequestSummaryRequestActionEnum
func GetWorkRequestSummaryRequestActionEnumStringValues() []string {
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

// GetMappingWorkRequestSummaryRequestActionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingWorkRequestSummaryRequestActionEnum(val string) (WorkRequestSummaryRequestActionEnum, bool) {
	enum, ok := mappingWorkRequestSummaryRequestActionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// WorkRequestSummaryStatusEnum Enum with underlying type: string
type WorkRequestSummaryStatusEnum string

// Set of constants representing the allowable values for WorkRequestSummaryStatusEnum
const (
	WorkRequestSummaryStatusAccepted   WorkRequestSummaryStatusEnum = "ACCEPTED"
	WorkRequestSummaryStatusInProgress WorkRequestSummaryStatusEnum = "IN_PROGRESS"
	WorkRequestSummaryStatusSucceeded  WorkRequestSummaryStatusEnum = "SUCCEEDED"
	WorkRequestSummaryStatusFailed     WorkRequestSummaryStatusEnum = "FAILED"
	WorkRequestSummaryStatusCanceling  WorkRequestSummaryStatusEnum = "CANCELING"
	WorkRequestSummaryStatusCanceled   WorkRequestSummaryStatusEnum = "CANCELED"
)

var mappingWorkRequestSummaryStatusEnum = map[string]WorkRequestSummaryStatusEnum{
	"ACCEPTED":    WorkRequestSummaryStatusAccepted,
	"IN_PROGRESS": WorkRequestSummaryStatusInProgress,
	"SUCCEEDED":   WorkRequestSummaryStatusSucceeded,
	"FAILED":      WorkRequestSummaryStatusFailed,
	"CANCELING":   WorkRequestSummaryStatusCanceling,
	"CANCELED":    WorkRequestSummaryStatusCanceled,
}

var mappingWorkRequestSummaryStatusEnumLowerCase = map[string]WorkRequestSummaryStatusEnum{
	"accepted":    WorkRequestSummaryStatusAccepted,
	"in_progress": WorkRequestSummaryStatusInProgress,
	"succeeded":   WorkRequestSummaryStatusSucceeded,
	"failed":      WorkRequestSummaryStatusFailed,
	"canceling":   WorkRequestSummaryStatusCanceling,
	"canceled":    WorkRequestSummaryStatusCanceled,
}

// GetWorkRequestSummaryStatusEnumValues Enumerates the set of values for WorkRequestSummaryStatusEnum
func GetWorkRequestSummaryStatusEnumValues() []WorkRequestSummaryStatusEnum {
	values := make([]WorkRequestSummaryStatusEnum, 0)
	for _, v := range mappingWorkRequestSummaryStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetWorkRequestSummaryStatusEnumStringValues Enumerates the set of values in String for WorkRequestSummaryStatusEnum
func GetWorkRequestSummaryStatusEnumStringValues() []string {
	return []string{
		"ACCEPTED",
		"IN_PROGRESS",
		"SUCCEEDED",
		"FAILED",
		"CANCELING",
		"CANCELED",
	}
}

// GetMappingWorkRequestSummaryStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingWorkRequestSummaryStatusEnum(val string) (WorkRequestSummaryStatusEnum, bool) {
	enum, ok := mappingWorkRequestSummaryStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
