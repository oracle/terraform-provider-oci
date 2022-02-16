// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// WorkRequestSummary Summary of a work request.
type WorkRequestSummary struct {

	// The asynchronous operation tracked by this work request.
	OperationType WorkRequestSummaryOperationTypeEnum `mandatory:"true" json:"operationType"`

	// The current status of the work request.
	Status WorkRequestSummaryStatusEnum `mandatory:"true" json:"status"`

	// The OCID of the work request.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the compartment that contains the work request.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The resources that are affected by the work request.
	Resources []WorkRequestResource `mandatory:"true" json:"resources"`

	// Progress of the work request in percentage.
	PercentComplete *float32 `mandatory:"true" json:"percentComplete"`

	// The date and time the work request was accepted, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeAccepted *common.SDKTime `mandatory:"true" json:"timeAccepted"`

	// The date and time the work request transitioned from ACCEPTED to IN_PROGRESS, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeStarted *common.SDKTime `mandatory:"false" json:"timeStarted"`

	// The date and time the work request reached a terminal state, either FAILED or SUCCEEDED, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeFinished *common.SDKTime `mandatory:"false" json:"timeFinished"`
}

func (m WorkRequestSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m WorkRequestSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingWorkRequestSummaryOperationTypeEnum(string(m.OperationType)); !ok && m.OperationType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OperationType: %s. Supported values are: %s.", m.OperationType, strings.Join(GetWorkRequestSummaryOperationTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingWorkRequestSummaryStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetWorkRequestSummaryStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// WorkRequestSummaryOperationTypeEnum Enum with underlying type: string
type WorkRequestSummaryOperationTypeEnum string

// Set of constants representing the allowable values for WorkRequestSummaryOperationTypeEnum
const (
	WorkRequestSummaryOperationTypeEnableDataSafeConfiguration           WorkRequestSummaryOperationTypeEnum = "ENABLE_DATA_SAFE_CONFIGURATION"
	WorkRequestSummaryOperationTypeCreatePrivateEndpoint                 WorkRequestSummaryOperationTypeEnum = "CREATE_PRIVATE_ENDPOINT"
	WorkRequestSummaryOperationTypeUpdatePrivateEndpoint                 WorkRequestSummaryOperationTypeEnum = "UPDATE_PRIVATE_ENDPOINT"
	WorkRequestSummaryOperationTypeDeletePrivateEndpoint                 WorkRequestSummaryOperationTypeEnum = "DELETE_PRIVATE_ENDPOINT"
	WorkRequestSummaryOperationTypeChangePrivateEndpointCompartment      WorkRequestSummaryOperationTypeEnum = "CHANGE_PRIVATE_ENDPOINT_COMPARTMENT"
	WorkRequestSummaryOperationTypeCreateOnpremConnector                 WorkRequestSummaryOperationTypeEnum = "CREATE_ONPREM_CONNECTOR"
	WorkRequestSummaryOperationTypeUpdateOnpremConnector                 WorkRequestSummaryOperationTypeEnum = "UPDATE_ONPREM_CONNECTOR"
	WorkRequestSummaryOperationTypeDeleteOnpremConnector                 WorkRequestSummaryOperationTypeEnum = "DELETE_ONPREM_CONNECTOR"
	WorkRequestSummaryOperationTypeUpdateOnpremConnectorWallet           WorkRequestSummaryOperationTypeEnum = "UPDATE_ONPREM_CONNECTOR_WALLET"
	WorkRequestSummaryOperationTypeChangeOnpremConnectorCompartment      WorkRequestSummaryOperationTypeEnum = "CHANGE_ONPREM_CONNECTOR_COMPARTMENT"
	WorkRequestSummaryOperationTypeProvisionPolicy                       WorkRequestSummaryOperationTypeEnum = "PROVISION_POLICY"
	WorkRequestSummaryOperationTypeRetrievePolicy                        WorkRequestSummaryOperationTypeEnum = "RETRIEVE_POLICY"
	WorkRequestSummaryOperationTypeUpdatePolicy                          WorkRequestSummaryOperationTypeEnum = "UPDATE_POLICY"
	WorkRequestSummaryOperationTypeChangePolicyCompartment               WorkRequestSummaryOperationTypeEnum = "CHANGE_POLICY_COMPARTMENT"
	WorkRequestSummaryOperationTypeCreateTargetDatabase                  WorkRequestSummaryOperationTypeEnum = "CREATE_TARGET_DATABASE"
	WorkRequestSummaryOperationTypeUpdateTargetDatabase                  WorkRequestSummaryOperationTypeEnum = "UPDATE_TARGET_DATABASE"
	WorkRequestSummaryOperationTypeActivateTargetDatabase                WorkRequestSummaryOperationTypeEnum = "ACTIVATE_TARGET_DATABASE"
	WorkRequestSummaryOperationTypeDeactivateTargetDatabase              WorkRequestSummaryOperationTypeEnum = "DEACTIVATE_TARGET_DATABASE"
	WorkRequestSummaryOperationTypeDeleteTargetDatabase                  WorkRequestSummaryOperationTypeEnum = "DELETE_TARGET_DATABASE"
	WorkRequestSummaryOperationTypeChangeTargetDatabaseCompartment       WorkRequestSummaryOperationTypeEnum = "CHANGE_TARGET_DATABASE_COMPARTMENT"
	WorkRequestSummaryOperationTypeCreateUserAssessment                  WorkRequestSummaryOperationTypeEnum = "CREATE_USER_ASSESSMENT"
	WorkRequestSummaryOperationTypeAssessUserAssessment                  WorkRequestSummaryOperationTypeEnum = "ASSESS_USER_ASSESSMENT"
	WorkRequestSummaryOperationTypeCreateSnapshotUserAssessment          WorkRequestSummaryOperationTypeEnum = "CREATE_SNAPSHOT_USER_ASSESSMENT"
	WorkRequestSummaryOperationTypeCreateScheduleUserAssessment          WorkRequestSummaryOperationTypeEnum = "CREATE_SCHEDULE_USER_ASSESSMENT"
	WorkRequestSummaryOperationTypeCompareWithBaselineUserAssessment     WorkRequestSummaryOperationTypeEnum = "COMPARE_WITH_BASELINE_USER_ASSESSMENT"
	WorkRequestSummaryOperationTypeDeleteUserAssessment                  WorkRequestSummaryOperationTypeEnum = "DELETE_USER_ASSESSMENT"
	WorkRequestSummaryOperationTypeUpdateUserAssessment                  WorkRequestSummaryOperationTypeEnum = "UPDATE_USER_ASSESSMENT"
	WorkRequestSummaryOperationTypeChangeUserAssessmentCompartment       WorkRequestSummaryOperationTypeEnum = "CHANGE_USER_ASSESSMENT_COMPARTMENT"
	WorkRequestSummaryOperationTypeSetUserAssessmentBaseline             WorkRequestSummaryOperationTypeEnum = "SET_USER_ASSESSMENT_BASELINE"
	WorkRequestSummaryOperationTypeUnsetUserAssessmentBaseline           WorkRequestSummaryOperationTypeEnum = "UNSET_USER_ASSESSMENT_BASELINE"
	WorkRequestSummaryOperationTypeGenerateUserAssessmentReport          WorkRequestSummaryOperationTypeEnum = "GENERATE_USER_ASSESSMENT_REPORT"
	WorkRequestSummaryOperationTypeCreateSecurityAssessment              WorkRequestSummaryOperationTypeEnum = "CREATE_SECURITY_ASSESSMENT"
	WorkRequestSummaryOperationTypeCreateSecurityAssessmentNow           WorkRequestSummaryOperationTypeEnum = "CREATE_SECURITY_ASSESSMENT_NOW"
	WorkRequestSummaryOperationTypeAssessSecurityAssessment              WorkRequestSummaryOperationTypeEnum = "ASSESS_SECURITY_ASSESSMENT"
	WorkRequestSummaryOperationTypeCreateSnapshotSecurityAssessment      WorkRequestSummaryOperationTypeEnum = "CREATE_SNAPSHOT_SECURITY_ASSESSMENT"
	WorkRequestSummaryOperationTypeCreateScheduleSecurityAssessment      WorkRequestSummaryOperationTypeEnum = "CREATE_SCHEDULE_SECURITY_ASSESSMENT"
	WorkRequestSummaryOperationTypeCompareWithBaselineSecurityAssessment WorkRequestSummaryOperationTypeEnum = "COMPARE_WITH_BASELINE_SECURITY_ASSESSMENT"
	WorkRequestSummaryOperationTypeDeleteSecurityAssessment              WorkRequestSummaryOperationTypeEnum = "DELETE_SECURITY_ASSESSMENT"
	WorkRequestSummaryOperationTypeUpdateSecurityAssessment              WorkRequestSummaryOperationTypeEnum = "UPDATE_SECURITY_ASSESSMENT"
	WorkRequestSummaryOperationTypeChangeSecurityAssessmentCompartment   WorkRequestSummaryOperationTypeEnum = "CHANGE_SECURITY_ASSESSMENT_COMPARTMENT"
	WorkRequestSummaryOperationTypeSetSecurityAssessmentBaseline         WorkRequestSummaryOperationTypeEnum = "SET_SECURITY_ASSESSMENT_BASELINE"
	WorkRequestSummaryOperationTypeUnsetSecurityAssessmentBaseline       WorkRequestSummaryOperationTypeEnum = "UNSET_SECURITY_ASSESSMENT_BASELINE"
	WorkRequestSummaryOperationTypeGenerateSecurityAssessmentReport      WorkRequestSummaryOperationTypeEnum = "GENERATE_SECURITY_ASSESSMENT_REPORT"
	WorkRequestSummaryOperationTypeCalculateVolume                       WorkRequestSummaryOperationTypeEnum = "CALCULATE_VOLUME"
	WorkRequestSummaryOperationTypeCalculateCollectedVolume              WorkRequestSummaryOperationTypeEnum = "CALCULATE_COLLECTED_VOLUME"
	WorkRequestSummaryOperationTypeAuditTrail                            WorkRequestSummaryOperationTypeEnum = "AUDIT_TRAIL"
	WorkRequestSummaryOperationTypeDeleteAuditTrail                      WorkRequestSummaryOperationTypeEnum = "DELETE_AUDIT_TRAIL"
	WorkRequestSummaryOperationTypeDiscoverAuditTrails                   WorkRequestSummaryOperationTypeEnum = "DISCOVER_AUDIT_TRAILS"
	WorkRequestSummaryOperationTypeUpdateAuditTrail                      WorkRequestSummaryOperationTypeEnum = "UPDATE_AUDIT_TRAIL"
	WorkRequestSummaryOperationTypeUpdateAuditProfile                    WorkRequestSummaryOperationTypeEnum = "UPDATE_AUDIT_PROFILE"
	WorkRequestSummaryOperationTypeAuditChangeCompartment                WorkRequestSummaryOperationTypeEnum = "AUDIT_CHANGE_COMPARTMENT"
	WorkRequestSummaryOperationTypeCreateReportDefinition                WorkRequestSummaryOperationTypeEnum = "CREATE_REPORT_DEFINITION"
	WorkRequestSummaryOperationTypeUpdateReportDefinition                WorkRequestSummaryOperationTypeEnum = "UPDATE_REPORT_DEFINITION"
	WorkRequestSummaryOperationTypeChangeReportDefinitionCompartment     WorkRequestSummaryOperationTypeEnum = "CHANGE_REPORT_DEFINITION_COMPARTMENT"
	WorkRequestSummaryOperationTypeDeleteReportDefinition                WorkRequestSummaryOperationTypeEnum = "DELETE_REPORT_DEFINITION"
	WorkRequestSummaryOperationTypeGenerateReport                        WorkRequestSummaryOperationTypeEnum = "GENERATE_REPORT"
	WorkRequestSummaryOperationTypeChangeReportCompartment               WorkRequestSummaryOperationTypeEnum = "CHANGE_REPORT_COMPARTMENT"
	WorkRequestSummaryOperationTypeDeleteArchiveRetrieval                WorkRequestSummaryOperationTypeEnum = "DELETE_ARCHIVE_RETRIEVAL"
	WorkRequestSummaryOperationTypeCreateArchiveRetrieval                WorkRequestSummaryOperationTypeEnum = "CREATE_ARCHIVE_RETRIEVAL"
	WorkRequestSummaryOperationTypeUpdateArchiveRetrieval                WorkRequestSummaryOperationTypeEnum = "UPDATE_ARCHIVE_RETRIEVAL"
	WorkRequestSummaryOperationTypeChangeArchiveRetrievalCompartment     WorkRequestSummaryOperationTypeEnum = "CHANGE_ARCHIVE_RETRIEVAL_COMPARTMENT"
	WorkRequestSummaryOperationTypeUpdateAlert                           WorkRequestSummaryOperationTypeEnum = "UPDATE_ALERT"
	WorkRequestSummaryOperationTypeTargetAlertPolicyAssociation          WorkRequestSummaryOperationTypeEnum = "TARGET_ALERT_POLICY_ASSOCIATION"
	WorkRequestSummaryOperationTypeCreateSensitiveDataModel              WorkRequestSummaryOperationTypeEnum = "CREATE_SENSITIVE_DATA_MODEL"
	WorkRequestSummaryOperationTypeUpdateSensitiveDataModel              WorkRequestSummaryOperationTypeEnum = "UPDATE_SENSITIVE_DATA_MODEL"
	WorkRequestSummaryOperationTypeDeleteSensitiveDataModel              WorkRequestSummaryOperationTypeEnum = "DELETE_SENSITIVE_DATA_MODEL"
	WorkRequestSummaryOperationTypeUploadSensitiveDataModel              WorkRequestSummaryOperationTypeEnum = "UPLOAD_SENSITIVE_DATA_MODEL"
	WorkRequestSummaryOperationTypeGenerateSensitiveDataModelForDownload WorkRequestSummaryOperationTypeEnum = "GENERATE_SENSITIVE_DATA_MODEL_FOR_DOWNLOAD"
	WorkRequestSummaryOperationTypeCreateSensitiveColumn                 WorkRequestSummaryOperationTypeEnum = "CREATE_SENSITIVE_COLUMN"
	WorkRequestSummaryOperationTypeUpdateSensitiveColumn                 WorkRequestSummaryOperationTypeEnum = "UPDATE_SENSITIVE_COLUMN"
	WorkRequestSummaryOperationTypePatchSensitiveColumns                 WorkRequestSummaryOperationTypeEnum = "PATCH_SENSITIVE_COLUMNS"
	WorkRequestSummaryOperationTypeCreateDiscoveryJob                    WorkRequestSummaryOperationTypeEnum = "CREATE_DISCOVERY_JOB"
	WorkRequestSummaryOperationTypeDeleteDiscoveryJob                    WorkRequestSummaryOperationTypeEnum = "DELETE_DISCOVERY_JOB"
	WorkRequestSummaryOperationTypePatchDiscoveryJobResult               WorkRequestSummaryOperationTypeEnum = "PATCH_DISCOVERY_JOB_RESULT"
	WorkRequestSummaryOperationTypeApplyDiscoveryJobResult               WorkRequestSummaryOperationTypeEnum = "APPLY_DISCOVERY_JOB_RESULT"
	WorkRequestSummaryOperationTypeGenerateDiscoveryReport               WorkRequestSummaryOperationTypeEnum = "GENERATE_DISCOVERY_REPORT"
	WorkRequestSummaryOperationTypeCreateSensitiveType                   WorkRequestSummaryOperationTypeEnum = "CREATE_SENSITIVE_TYPE"
	WorkRequestSummaryOperationTypeUpdateSensitiveType                   WorkRequestSummaryOperationTypeEnum = "UPDATE_SENSITIVE_TYPE"
	WorkRequestSummaryOperationTypeCreateMaskingPolicy                   WorkRequestSummaryOperationTypeEnum = "CREATE_MASKING_POLICY"
	WorkRequestSummaryOperationTypeUpdateMaskingPolicy                   WorkRequestSummaryOperationTypeEnum = "UPDATE_MASKING_POLICY"
	WorkRequestSummaryOperationTypeDeleteMaskingPolicy                   WorkRequestSummaryOperationTypeEnum = "DELETE_MASKING_POLICY"
	WorkRequestSummaryOperationTypeUploadMaskingPolicy                   WorkRequestSummaryOperationTypeEnum = "UPLOAD_MASKING_POLICY"
	WorkRequestSummaryOperationTypeGenerateMaskingPolicyForDownload      WorkRequestSummaryOperationTypeEnum = "GENERATE_MASKING_POLICY_FOR_DOWNLOAD"
	WorkRequestSummaryOperationTypeCreateMaskingColumn                   WorkRequestSummaryOperationTypeEnum = "CREATE_MASKING_COLUMN"
	WorkRequestSummaryOperationTypeUpdateMaskingColumn                   WorkRequestSummaryOperationTypeEnum = "UPDATE_MASKING_COLUMN"
	WorkRequestSummaryOperationTypePatchMaskingColumns                   WorkRequestSummaryOperationTypeEnum = "PATCH_MASKING_COLUMNS"
	WorkRequestSummaryOperationTypeGenerateMaskingReport                 WorkRequestSummaryOperationTypeEnum = "GENERATE_MASKING_REPORT"
	WorkRequestSummaryOperationTypeCreateLibraryMaskingFormat            WorkRequestSummaryOperationTypeEnum = "CREATE_LIBRARY_MASKING_FORMAT"
	WorkRequestSummaryOperationTypeUpdateLibraryMaskingFormat            WorkRequestSummaryOperationTypeEnum = "UPDATE_LIBRARY_MASKING_FORMAT"
	WorkRequestSummaryOperationTypeAddColumnsFromSdm                     WorkRequestSummaryOperationTypeEnum = "ADD_COLUMNS_FROM_SDM"
	WorkRequestSummaryOperationTypeMaskingJob                            WorkRequestSummaryOperationTypeEnum = "MASKING_JOB"
)

var mappingWorkRequestSummaryOperationTypeEnum = map[string]WorkRequestSummaryOperationTypeEnum{
	"ENABLE_DATA_SAFE_CONFIGURATION":             WorkRequestSummaryOperationTypeEnableDataSafeConfiguration,
	"CREATE_PRIVATE_ENDPOINT":                    WorkRequestSummaryOperationTypeCreatePrivateEndpoint,
	"UPDATE_PRIVATE_ENDPOINT":                    WorkRequestSummaryOperationTypeUpdatePrivateEndpoint,
	"DELETE_PRIVATE_ENDPOINT":                    WorkRequestSummaryOperationTypeDeletePrivateEndpoint,
	"CHANGE_PRIVATE_ENDPOINT_COMPARTMENT":        WorkRequestSummaryOperationTypeChangePrivateEndpointCompartment,
	"CREATE_ONPREM_CONNECTOR":                    WorkRequestSummaryOperationTypeCreateOnpremConnector,
	"UPDATE_ONPREM_CONNECTOR":                    WorkRequestSummaryOperationTypeUpdateOnpremConnector,
	"DELETE_ONPREM_CONNECTOR":                    WorkRequestSummaryOperationTypeDeleteOnpremConnector,
	"UPDATE_ONPREM_CONNECTOR_WALLET":             WorkRequestSummaryOperationTypeUpdateOnpremConnectorWallet,
	"CHANGE_ONPREM_CONNECTOR_COMPARTMENT":        WorkRequestSummaryOperationTypeChangeOnpremConnectorCompartment,
	"PROVISION_POLICY":                           WorkRequestSummaryOperationTypeProvisionPolicy,
	"RETRIEVE_POLICY":                            WorkRequestSummaryOperationTypeRetrievePolicy,
	"UPDATE_POLICY":                              WorkRequestSummaryOperationTypeUpdatePolicy,
	"CHANGE_POLICY_COMPARTMENT":                  WorkRequestSummaryOperationTypeChangePolicyCompartment,
	"CREATE_TARGET_DATABASE":                     WorkRequestSummaryOperationTypeCreateTargetDatabase,
	"UPDATE_TARGET_DATABASE":                     WorkRequestSummaryOperationTypeUpdateTargetDatabase,
	"ACTIVATE_TARGET_DATABASE":                   WorkRequestSummaryOperationTypeActivateTargetDatabase,
	"DEACTIVATE_TARGET_DATABASE":                 WorkRequestSummaryOperationTypeDeactivateTargetDatabase,
	"DELETE_TARGET_DATABASE":                     WorkRequestSummaryOperationTypeDeleteTargetDatabase,
	"CHANGE_TARGET_DATABASE_COMPARTMENT":         WorkRequestSummaryOperationTypeChangeTargetDatabaseCompartment,
	"CREATE_USER_ASSESSMENT":                     WorkRequestSummaryOperationTypeCreateUserAssessment,
	"ASSESS_USER_ASSESSMENT":                     WorkRequestSummaryOperationTypeAssessUserAssessment,
	"CREATE_SNAPSHOT_USER_ASSESSMENT":            WorkRequestSummaryOperationTypeCreateSnapshotUserAssessment,
	"CREATE_SCHEDULE_USER_ASSESSMENT":            WorkRequestSummaryOperationTypeCreateScheduleUserAssessment,
	"COMPARE_WITH_BASELINE_USER_ASSESSMENT":      WorkRequestSummaryOperationTypeCompareWithBaselineUserAssessment,
	"DELETE_USER_ASSESSMENT":                     WorkRequestSummaryOperationTypeDeleteUserAssessment,
	"UPDATE_USER_ASSESSMENT":                     WorkRequestSummaryOperationTypeUpdateUserAssessment,
	"CHANGE_USER_ASSESSMENT_COMPARTMENT":         WorkRequestSummaryOperationTypeChangeUserAssessmentCompartment,
	"SET_USER_ASSESSMENT_BASELINE":               WorkRequestSummaryOperationTypeSetUserAssessmentBaseline,
	"UNSET_USER_ASSESSMENT_BASELINE":             WorkRequestSummaryOperationTypeUnsetUserAssessmentBaseline,
	"GENERATE_USER_ASSESSMENT_REPORT":            WorkRequestSummaryOperationTypeGenerateUserAssessmentReport,
	"CREATE_SECURITY_ASSESSMENT":                 WorkRequestSummaryOperationTypeCreateSecurityAssessment,
	"CREATE_SECURITY_ASSESSMENT_NOW":             WorkRequestSummaryOperationTypeCreateSecurityAssessmentNow,
	"ASSESS_SECURITY_ASSESSMENT":                 WorkRequestSummaryOperationTypeAssessSecurityAssessment,
	"CREATE_SNAPSHOT_SECURITY_ASSESSMENT":        WorkRequestSummaryOperationTypeCreateSnapshotSecurityAssessment,
	"CREATE_SCHEDULE_SECURITY_ASSESSMENT":        WorkRequestSummaryOperationTypeCreateScheduleSecurityAssessment,
	"COMPARE_WITH_BASELINE_SECURITY_ASSESSMENT":  WorkRequestSummaryOperationTypeCompareWithBaselineSecurityAssessment,
	"DELETE_SECURITY_ASSESSMENT":                 WorkRequestSummaryOperationTypeDeleteSecurityAssessment,
	"UPDATE_SECURITY_ASSESSMENT":                 WorkRequestSummaryOperationTypeUpdateSecurityAssessment,
	"CHANGE_SECURITY_ASSESSMENT_COMPARTMENT":     WorkRequestSummaryOperationTypeChangeSecurityAssessmentCompartment,
	"SET_SECURITY_ASSESSMENT_BASELINE":           WorkRequestSummaryOperationTypeSetSecurityAssessmentBaseline,
	"UNSET_SECURITY_ASSESSMENT_BASELINE":         WorkRequestSummaryOperationTypeUnsetSecurityAssessmentBaseline,
	"GENERATE_SECURITY_ASSESSMENT_REPORT":        WorkRequestSummaryOperationTypeGenerateSecurityAssessmentReport,
	"CALCULATE_VOLUME":                           WorkRequestSummaryOperationTypeCalculateVolume,
	"CALCULATE_COLLECTED_VOLUME":                 WorkRequestSummaryOperationTypeCalculateCollectedVolume,
	"AUDIT_TRAIL":                                WorkRequestSummaryOperationTypeAuditTrail,
	"DELETE_AUDIT_TRAIL":                         WorkRequestSummaryOperationTypeDeleteAuditTrail,
	"DISCOVER_AUDIT_TRAILS":                      WorkRequestSummaryOperationTypeDiscoverAuditTrails,
	"UPDATE_AUDIT_TRAIL":                         WorkRequestSummaryOperationTypeUpdateAuditTrail,
	"UPDATE_AUDIT_PROFILE":                       WorkRequestSummaryOperationTypeUpdateAuditProfile,
	"AUDIT_CHANGE_COMPARTMENT":                   WorkRequestSummaryOperationTypeAuditChangeCompartment,
	"CREATE_REPORT_DEFINITION":                   WorkRequestSummaryOperationTypeCreateReportDefinition,
	"UPDATE_REPORT_DEFINITION":                   WorkRequestSummaryOperationTypeUpdateReportDefinition,
	"CHANGE_REPORT_DEFINITION_COMPARTMENT":       WorkRequestSummaryOperationTypeChangeReportDefinitionCompartment,
	"DELETE_REPORT_DEFINITION":                   WorkRequestSummaryOperationTypeDeleteReportDefinition,
	"GENERATE_REPORT":                            WorkRequestSummaryOperationTypeGenerateReport,
	"CHANGE_REPORT_COMPARTMENT":                  WorkRequestSummaryOperationTypeChangeReportCompartment,
	"DELETE_ARCHIVE_RETRIEVAL":                   WorkRequestSummaryOperationTypeDeleteArchiveRetrieval,
	"CREATE_ARCHIVE_RETRIEVAL":                   WorkRequestSummaryOperationTypeCreateArchiveRetrieval,
	"UPDATE_ARCHIVE_RETRIEVAL":                   WorkRequestSummaryOperationTypeUpdateArchiveRetrieval,
	"CHANGE_ARCHIVE_RETRIEVAL_COMPARTMENT":       WorkRequestSummaryOperationTypeChangeArchiveRetrievalCompartment,
	"UPDATE_ALERT":                               WorkRequestSummaryOperationTypeUpdateAlert,
	"TARGET_ALERT_POLICY_ASSOCIATION":            WorkRequestSummaryOperationTypeTargetAlertPolicyAssociation,
	"CREATE_SENSITIVE_DATA_MODEL":                WorkRequestSummaryOperationTypeCreateSensitiveDataModel,
	"UPDATE_SENSITIVE_DATA_MODEL":                WorkRequestSummaryOperationTypeUpdateSensitiveDataModel,
	"DELETE_SENSITIVE_DATA_MODEL":                WorkRequestSummaryOperationTypeDeleteSensitiveDataModel,
	"UPLOAD_SENSITIVE_DATA_MODEL":                WorkRequestSummaryOperationTypeUploadSensitiveDataModel,
	"GENERATE_SENSITIVE_DATA_MODEL_FOR_DOWNLOAD": WorkRequestSummaryOperationTypeGenerateSensitiveDataModelForDownload,
	"CREATE_SENSITIVE_COLUMN":                    WorkRequestSummaryOperationTypeCreateSensitiveColumn,
	"UPDATE_SENSITIVE_COLUMN":                    WorkRequestSummaryOperationTypeUpdateSensitiveColumn,
	"PATCH_SENSITIVE_COLUMNS":                    WorkRequestSummaryOperationTypePatchSensitiveColumns,
	"CREATE_DISCOVERY_JOB":                       WorkRequestSummaryOperationTypeCreateDiscoveryJob,
	"DELETE_DISCOVERY_JOB":                       WorkRequestSummaryOperationTypeDeleteDiscoveryJob,
	"PATCH_DISCOVERY_JOB_RESULT":                 WorkRequestSummaryOperationTypePatchDiscoveryJobResult,
	"APPLY_DISCOVERY_JOB_RESULT":                 WorkRequestSummaryOperationTypeApplyDiscoveryJobResult,
	"GENERATE_DISCOVERY_REPORT":                  WorkRequestSummaryOperationTypeGenerateDiscoveryReport,
	"CREATE_SENSITIVE_TYPE":                      WorkRequestSummaryOperationTypeCreateSensitiveType,
	"UPDATE_SENSITIVE_TYPE":                      WorkRequestSummaryOperationTypeUpdateSensitiveType,
	"CREATE_MASKING_POLICY":                      WorkRequestSummaryOperationTypeCreateMaskingPolicy,
	"UPDATE_MASKING_POLICY":                      WorkRequestSummaryOperationTypeUpdateMaskingPolicy,
	"DELETE_MASKING_POLICY":                      WorkRequestSummaryOperationTypeDeleteMaskingPolicy,
	"UPLOAD_MASKING_POLICY":                      WorkRequestSummaryOperationTypeUploadMaskingPolicy,
	"GENERATE_MASKING_POLICY_FOR_DOWNLOAD":       WorkRequestSummaryOperationTypeGenerateMaskingPolicyForDownload,
	"CREATE_MASKING_COLUMN":                      WorkRequestSummaryOperationTypeCreateMaskingColumn,
	"UPDATE_MASKING_COLUMN":                      WorkRequestSummaryOperationTypeUpdateMaskingColumn,
	"PATCH_MASKING_COLUMNS":                      WorkRequestSummaryOperationTypePatchMaskingColumns,
	"GENERATE_MASKING_REPORT":                    WorkRequestSummaryOperationTypeGenerateMaskingReport,
	"CREATE_LIBRARY_MASKING_FORMAT":              WorkRequestSummaryOperationTypeCreateLibraryMaskingFormat,
	"UPDATE_LIBRARY_MASKING_FORMAT":              WorkRequestSummaryOperationTypeUpdateLibraryMaskingFormat,
	"ADD_COLUMNS_FROM_SDM":                       WorkRequestSummaryOperationTypeAddColumnsFromSdm,
	"MASKING_JOB":                                WorkRequestSummaryOperationTypeMaskingJob,
}

// GetWorkRequestSummaryOperationTypeEnumValues Enumerates the set of values for WorkRequestSummaryOperationTypeEnum
func GetWorkRequestSummaryOperationTypeEnumValues() []WorkRequestSummaryOperationTypeEnum {
	values := make([]WorkRequestSummaryOperationTypeEnum, 0)
	for _, v := range mappingWorkRequestSummaryOperationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetWorkRequestSummaryOperationTypeEnumStringValues Enumerates the set of values in String for WorkRequestSummaryOperationTypeEnum
func GetWorkRequestSummaryOperationTypeEnumStringValues() []string {
	return []string{
		"ENABLE_DATA_SAFE_CONFIGURATION",
		"CREATE_PRIVATE_ENDPOINT",
		"UPDATE_PRIVATE_ENDPOINT",
		"DELETE_PRIVATE_ENDPOINT",
		"CHANGE_PRIVATE_ENDPOINT_COMPARTMENT",
		"CREATE_ONPREM_CONNECTOR",
		"UPDATE_ONPREM_CONNECTOR",
		"DELETE_ONPREM_CONNECTOR",
		"UPDATE_ONPREM_CONNECTOR_WALLET",
		"CHANGE_ONPREM_CONNECTOR_COMPARTMENT",
		"PROVISION_POLICY",
		"RETRIEVE_POLICY",
		"UPDATE_POLICY",
		"CHANGE_POLICY_COMPARTMENT",
		"CREATE_TARGET_DATABASE",
		"UPDATE_TARGET_DATABASE",
		"ACTIVATE_TARGET_DATABASE",
		"DEACTIVATE_TARGET_DATABASE",
		"DELETE_TARGET_DATABASE",
		"CHANGE_TARGET_DATABASE_COMPARTMENT",
		"CREATE_USER_ASSESSMENT",
		"ASSESS_USER_ASSESSMENT",
		"CREATE_SNAPSHOT_USER_ASSESSMENT",
		"CREATE_SCHEDULE_USER_ASSESSMENT",
		"COMPARE_WITH_BASELINE_USER_ASSESSMENT",
		"DELETE_USER_ASSESSMENT",
		"UPDATE_USER_ASSESSMENT",
		"CHANGE_USER_ASSESSMENT_COMPARTMENT",
		"SET_USER_ASSESSMENT_BASELINE",
		"UNSET_USER_ASSESSMENT_BASELINE",
		"GENERATE_USER_ASSESSMENT_REPORT",
		"CREATE_SECURITY_ASSESSMENT",
		"CREATE_SECURITY_ASSESSMENT_NOW",
		"ASSESS_SECURITY_ASSESSMENT",
		"CREATE_SNAPSHOT_SECURITY_ASSESSMENT",
		"CREATE_SCHEDULE_SECURITY_ASSESSMENT",
		"COMPARE_WITH_BASELINE_SECURITY_ASSESSMENT",
		"DELETE_SECURITY_ASSESSMENT",
		"UPDATE_SECURITY_ASSESSMENT",
		"CHANGE_SECURITY_ASSESSMENT_COMPARTMENT",
		"SET_SECURITY_ASSESSMENT_BASELINE",
		"UNSET_SECURITY_ASSESSMENT_BASELINE",
		"GENERATE_SECURITY_ASSESSMENT_REPORT",
		"CALCULATE_VOLUME",
		"CALCULATE_COLLECTED_VOLUME",
		"AUDIT_TRAIL",
		"DELETE_AUDIT_TRAIL",
		"DISCOVER_AUDIT_TRAILS",
		"UPDATE_AUDIT_TRAIL",
		"UPDATE_AUDIT_PROFILE",
		"AUDIT_CHANGE_COMPARTMENT",
		"CREATE_REPORT_DEFINITION",
		"UPDATE_REPORT_DEFINITION",
		"CHANGE_REPORT_DEFINITION_COMPARTMENT",
		"DELETE_REPORT_DEFINITION",
		"GENERATE_REPORT",
		"CHANGE_REPORT_COMPARTMENT",
		"DELETE_ARCHIVE_RETRIEVAL",
		"CREATE_ARCHIVE_RETRIEVAL",
		"UPDATE_ARCHIVE_RETRIEVAL",
		"CHANGE_ARCHIVE_RETRIEVAL_COMPARTMENT",
		"UPDATE_ALERT",
		"TARGET_ALERT_POLICY_ASSOCIATION",
		"CREATE_SENSITIVE_DATA_MODEL",
		"UPDATE_SENSITIVE_DATA_MODEL",
		"DELETE_SENSITIVE_DATA_MODEL",
		"UPLOAD_SENSITIVE_DATA_MODEL",
		"GENERATE_SENSITIVE_DATA_MODEL_FOR_DOWNLOAD",
		"CREATE_SENSITIVE_COLUMN",
		"UPDATE_SENSITIVE_COLUMN",
		"PATCH_SENSITIVE_COLUMNS",
		"CREATE_DISCOVERY_JOB",
		"DELETE_DISCOVERY_JOB",
		"PATCH_DISCOVERY_JOB_RESULT",
		"APPLY_DISCOVERY_JOB_RESULT",
		"GENERATE_DISCOVERY_REPORT",
		"CREATE_SENSITIVE_TYPE",
		"UPDATE_SENSITIVE_TYPE",
		"CREATE_MASKING_POLICY",
		"UPDATE_MASKING_POLICY",
		"DELETE_MASKING_POLICY",
		"UPLOAD_MASKING_POLICY",
		"GENERATE_MASKING_POLICY_FOR_DOWNLOAD",
		"CREATE_MASKING_COLUMN",
		"UPDATE_MASKING_COLUMN",
		"PATCH_MASKING_COLUMNS",
		"GENERATE_MASKING_REPORT",
		"CREATE_LIBRARY_MASKING_FORMAT",
		"UPDATE_LIBRARY_MASKING_FORMAT",
		"ADD_COLUMNS_FROM_SDM",
		"MASKING_JOB",
	}
}

// GetMappingWorkRequestSummaryOperationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingWorkRequestSummaryOperationTypeEnum(val string) (WorkRequestSummaryOperationTypeEnum, bool) {
	mappingWorkRequestSummaryOperationTypeEnumIgnoreCase := make(map[string]WorkRequestSummaryOperationTypeEnum)
	for k, v := range mappingWorkRequestSummaryOperationTypeEnum {
		mappingWorkRequestSummaryOperationTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingWorkRequestSummaryOperationTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// WorkRequestSummaryStatusEnum Enum with underlying type: string
type WorkRequestSummaryStatusEnum string

// Set of constants representing the allowable values for WorkRequestSummaryStatusEnum
const (
	WorkRequestSummaryStatusAccepted   WorkRequestSummaryStatusEnum = "ACCEPTED"
	WorkRequestSummaryStatusInProgress WorkRequestSummaryStatusEnum = "IN_PROGRESS"
	WorkRequestSummaryStatusFailed     WorkRequestSummaryStatusEnum = "FAILED"
	WorkRequestSummaryStatusSucceeded  WorkRequestSummaryStatusEnum = "SUCCEEDED"
	WorkRequestSummaryStatusCanceling  WorkRequestSummaryStatusEnum = "CANCELING"
	WorkRequestSummaryStatusCanceled   WorkRequestSummaryStatusEnum = "CANCELED"
	WorkRequestSummaryStatusSuspending WorkRequestSummaryStatusEnum = "SUSPENDING"
	WorkRequestSummaryStatusSuspended  WorkRequestSummaryStatusEnum = "SUSPENDED"
)

var mappingWorkRequestSummaryStatusEnum = map[string]WorkRequestSummaryStatusEnum{
	"ACCEPTED":    WorkRequestSummaryStatusAccepted,
	"IN_PROGRESS": WorkRequestSummaryStatusInProgress,
	"FAILED":      WorkRequestSummaryStatusFailed,
	"SUCCEEDED":   WorkRequestSummaryStatusSucceeded,
	"CANCELING":   WorkRequestSummaryStatusCanceling,
	"CANCELED":    WorkRequestSummaryStatusCanceled,
	"SUSPENDING":  WorkRequestSummaryStatusSuspending,
	"SUSPENDED":   WorkRequestSummaryStatusSuspended,
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
		"FAILED",
		"SUCCEEDED",
		"CANCELING",
		"CANCELED",
		"SUSPENDING",
		"SUSPENDED",
	}
}

// GetMappingWorkRequestSummaryStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingWorkRequestSummaryStatusEnum(val string) (WorkRequestSummaryStatusEnum, bool) {
	mappingWorkRequestSummaryStatusEnumIgnoreCase := make(map[string]WorkRequestSummaryStatusEnum)
	for k, v := range mappingWorkRequestSummaryStatusEnum {
		mappingWorkRequestSummaryStatusEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingWorkRequestSummaryStatusEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
