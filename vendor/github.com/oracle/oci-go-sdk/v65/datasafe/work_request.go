// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// WorkRequest An asynchronous work request.
type WorkRequest struct {

	// The resources that are affected by the work request.
	OperationType WorkRequestOperationTypeEnum `mandatory:"true" json:"operationType"`

	// The current status of the work request.
	Status WorkRequestStatusEnum `mandatory:"true" json:"status"`

	// The OCID of the work request.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the compartment that contains the work request.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The resources that are affected by this work request.
	Resources []WorkRequestResource `mandatory:"true" json:"resources"`

	// Progress of the work request in percentage.
	PercentComplete *float32 `mandatory:"true" json:"percentComplete"`

	// The date and time the work request was accepted, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeAccepted *common.SDKTime `mandatory:"true" json:"timeAccepted"`

	// The date and time the work request transitioned from ACCEPTED to IN_PROGRESS, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeStarted *common.SDKTime `mandatory:"false" json:"timeStarted"`

	// The date and time the work request reached a terminal state, either FAILED or SUCCEEDED. Format is defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
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
	WorkRequestOperationTypeEnableDataSafeConfiguration               WorkRequestOperationTypeEnum = "ENABLE_DATA_SAFE_CONFIGURATION"
	WorkRequestOperationTypeCreatePrivateEndpoint                     WorkRequestOperationTypeEnum = "CREATE_PRIVATE_ENDPOINT"
	WorkRequestOperationTypeUpdatePrivateEndpoint                     WorkRequestOperationTypeEnum = "UPDATE_PRIVATE_ENDPOINT"
	WorkRequestOperationTypeDeletePrivateEndpoint                     WorkRequestOperationTypeEnum = "DELETE_PRIVATE_ENDPOINT"
	WorkRequestOperationTypeChangePrivateEndpointCompartment          WorkRequestOperationTypeEnum = "CHANGE_PRIVATE_ENDPOINT_COMPARTMENT"
	WorkRequestOperationTypeCreateOnpremConnector                     WorkRequestOperationTypeEnum = "CREATE_ONPREM_CONNECTOR"
	WorkRequestOperationTypeUpdateOnpremConnector                     WorkRequestOperationTypeEnum = "UPDATE_ONPREM_CONNECTOR"
	WorkRequestOperationTypeDeleteOnpremConnector                     WorkRequestOperationTypeEnum = "DELETE_ONPREM_CONNECTOR"
	WorkRequestOperationTypeUpdateOnpremConnectorWallet               WorkRequestOperationTypeEnum = "UPDATE_ONPREM_CONNECTOR_WALLET"
	WorkRequestOperationTypeChangeOnpremConnectorCompartment          WorkRequestOperationTypeEnum = "CHANGE_ONPREM_CONNECTOR_COMPARTMENT"
	WorkRequestOperationTypeCreateTargetDatabase                      WorkRequestOperationTypeEnum = "CREATE_TARGET_DATABASE"
	WorkRequestOperationTypeUpdateTargetDatabase                      WorkRequestOperationTypeEnum = "UPDATE_TARGET_DATABASE"
	WorkRequestOperationTypeActivateTargetDatabase                    WorkRequestOperationTypeEnum = "ACTIVATE_TARGET_DATABASE"
	WorkRequestOperationTypeDeactivateTargetDatabase                  WorkRequestOperationTypeEnum = "DEACTIVATE_TARGET_DATABASE"
	WorkRequestOperationTypeDeleteTargetDatabase                      WorkRequestOperationTypeEnum = "DELETE_TARGET_DATABASE"
	WorkRequestOperationTypeChangeTargetDatabaseCompartment           WorkRequestOperationTypeEnum = "CHANGE_TARGET_DATABASE_COMPARTMENT"
	WorkRequestOperationTypeCreatePeerTargetDatabase                  WorkRequestOperationTypeEnum = "CREATE_PEER_TARGET_DATABASE"
	WorkRequestOperationTypeUpdatePeerTargetDatabase                  WorkRequestOperationTypeEnum = "UPDATE_PEER_TARGET_DATABASE"
	WorkRequestOperationTypeDeletePeerTargetDatabase                  WorkRequestOperationTypeEnum = "DELETE_PEER_TARGET_DATABASE"
	WorkRequestOperationTypeRefreshTargetDatabase                     WorkRequestOperationTypeEnum = "REFRESH_TARGET_DATABASE"
	WorkRequestOperationTypeProvisionPolicy                           WorkRequestOperationTypeEnum = "PROVISION_POLICY"
	WorkRequestOperationTypeRetrievePolicy                            WorkRequestOperationTypeEnum = "RETRIEVE_POLICY"
	WorkRequestOperationTypeUpdatePolicy                              WorkRequestOperationTypeEnum = "UPDATE_POLICY"
	WorkRequestOperationTypeChangePolicyCompartment                   WorkRequestOperationTypeEnum = "CHANGE_POLICY_COMPARTMENT"
	WorkRequestOperationTypeCreateUserAssessment                      WorkRequestOperationTypeEnum = "CREATE_USER_ASSESSMENT"
	WorkRequestOperationTypeAssessUserAssessment                      WorkRequestOperationTypeEnum = "ASSESS_USER_ASSESSMENT"
	WorkRequestOperationTypeCreateSnapshotUserAssessment              WorkRequestOperationTypeEnum = "CREATE_SNAPSHOT_USER_ASSESSMENT"
	WorkRequestOperationTypeCreateScheduleUserAssessment              WorkRequestOperationTypeEnum = "CREATE_SCHEDULE_USER_ASSESSMENT"
	WorkRequestOperationTypeCompareWithBaselineUserAssessment         WorkRequestOperationTypeEnum = "COMPARE_WITH_BASELINE_USER_ASSESSMENT"
	WorkRequestOperationTypeDeleteUserAssessment                      WorkRequestOperationTypeEnum = "DELETE_USER_ASSESSMENT"
	WorkRequestOperationTypeUpdateUserAssessment                      WorkRequestOperationTypeEnum = "UPDATE_USER_ASSESSMENT"
	WorkRequestOperationTypeChangeUserAssessmentCompartment           WorkRequestOperationTypeEnum = "CHANGE_USER_ASSESSMENT_COMPARTMENT"
	WorkRequestOperationTypeSetUserAssessmentBaseline                 WorkRequestOperationTypeEnum = "SET_USER_ASSESSMENT_BASELINE"
	WorkRequestOperationTypeUnsetUserAssessmentBaseline               WorkRequestOperationTypeEnum = "UNSET_USER_ASSESSMENT_BASELINE"
	WorkRequestOperationTypeGenerateUserAssessmentReport              WorkRequestOperationTypeEnum = "GENERATE_USER_ASSESSMENT_REPORT"
	WorkRequestOperationTypeCreateSecurityAssessment                  WorkRequestOperationTypeEnum = "CREATE_SECURITY_ASSESSMENT"
	WorkRequestOperationTypeCreateSecurityAssessmentNow               WorkRequestOperationTypeEnum = "CREATE_SECURITY_ASSESSMENT_NOW"
	WorkRequestOperationTypeAssessSecurityAssessment                  WorkRequestOperationTypeEnum = "ASSESS_SECURITY_ASSESSMENT"
	WorkRequestOperationTypeCreateSnapshotSecurityAssessment          WorkRequestOperationTypeEnum = "CREATE_SNAPSHOT_SECURITY_ASSESSMENT"
	WorkRequestOperationTypeCreateScheduleSecurityAssessment          WorkRequestOperationTypeEnum = "CREATE_SCHEDULE_SECURITY_ASSESSMENT"
	WorkRequestOperationTypeCompareWithBaselineSecurityAssessment     WorkRequestOperationTypeEnum = "COMPARE_WITH_BASELINE_SECURITY_ASSESSMENT"
	WorkRequestOperationTypeDeleteSecurityAssessment                  WorkRequestOperationTypeEnum = "DELETE_SECURITY_ASSESSMENT"
	WorkRequestOperationTypeUpdateSecurityAssessment                  WorkRequestOperationTypeEnum = "UPDATE_SECURITY_ASSESSMENT"
	WorkRequestOperationTypeUpdateFindingRisk                         WorkRequestOperationTypeEnum = "UPDATE_FINDING_RISK"
	WorkRequestOperationTypeChangeSecurityAssessmentCompartment       WorkRequestOperationTypeEnum = "CHANGE_SECURITY_ASSESSMENT_COMPARTMENT"
	WorkRequestOperationTypeSetSecurityAssessmentBaseline             WorkRequestOperationTypeEnum = "SET_SECURITY_ASSESSMENT_BASELINE"
	WorkRequestOperationTypeUnsetSecurityAssessmentBaseline           WorkRequestOperationTypeEnum = "UNSET_SECURITY_ASSESSMENT_BASELINE"
	WorkRequestOperationTypeGenerateSecurityAssessmentReport          WorkRequestOperationTypeEnum = "GENERATE_SECURITY_ASSESSMENT_REPORT"
	WorkRequestOperationTypeCreateAuditProfile                        WorkRequestOperationTypeEnum = "CREATE_AUDIT_PROFILE"
	WorkRequestOperationTypeCalculateVolume                           WorkRequestOperationTypeEnum = "CALCULATE_VOLUME"
	WorkRequestOperationTypeCalculateCollectedVolume                  WorkRequestOperationTypeEnum = "CALCULATE_COLLECTED_VOLUME"
	WorkRequestOperationTypeCreateDbSecurityConfig                    WorkRequestOperationTypeEnum = "CREATE_DB_SECURITY_CONFIG"
	WorkRequestOperationTypeRefreshDbSecurityConfig                   WorkRequestOperationTypeEnum = "REFRESH_DB_SECURITY_CONFIG"
	WorkRequestOperationTypeUpdateDbSecurityConfig                    WorkRequestOperationTypeEnum = "UPDATE_DB_SECURITY_CONFIG"
	WorkRequestOperationTypeChangeDbSecurityConfigCompartment         WorkRequestOperationTypeEnum = "CHANGE_DB_SECURITY_CONFIG_COMPARTMENT"
	WorkRequestOperationTypeGenerateFirewallPolicy                    WorkRequestOperationTypeEnum = "GENERATE_FIREWALL_POLICY"
	WorkRequestOperationTypeUpdateFirewallPolicy                      WorkRequestOperationTypeEnum = "UPDATE_FIREWALL_POLICY"
	WorkRequestOperationTypeChangeFirewallPolicyCompartment           WorkRequestOperationTypeEnum = "CHANGE_FIREWALL_POLICY_COMPARTMENT"
	WorkRequestOperationTypeDeleteFirewallPolicy                      WorkRequestOperationTypeEnum = "DELETE_FIREWALL_POLICY"
	WorkRequestOperationTypeCreateSqlCollection                       WorkRequestOperationTypeEnum = "CREATE_SQL_COLLECTION"
	WorkRequestOperationTypeUpdateSqlCollection                       WorkRequestOperationTypeEnum = "UPDATE_SQL_COLLECTION"
	WorkRequestOperationTypeStartSqlCollection                        WorkRequestOperationTypeEnum = "START_SQL_COLLECTION"
	WorkRequestOperationTypeStopSqlCollection                         WorkRequestOperationTypeEnum = "STOP_SQL_COLLECTION"
	WorkRequestOperationTypeDeleteSqlCollection                       WorkRequestOperationTypeEnum = "DELETE_SQL_COLLECTION"
	WorkRequestOperationTypeChangeSqlCollectionCompartment            WorkRequestOperationTypeEnum = "CHANGE_SQL_COLLECTION_COMPARTMENT"
	WorkRequestOperationTypeRefreshSqlCollectionLogInsights           WorkRequestOperationTypeEnum = "REFRESH_SQL_COLLECTION_LOG_INSIGHTS"
	WorkRequestOperationTypePurgeSqlCollectionLogs                    WorkRequestOperationTypeEnum = "PURGE_SQL_COLLECTION_LOGS"
	WorkRequestOperationTypeRefreshViolations                         WorkRequestOperationTypeEnum = "REFRESH_VIOLATIONS"
	WorkRequestOperationTypeCreateArchival                            WorkRequestOperationTypeEnum = "CREATE_ARCHIVAL"
	WorkRequestOperationTypeUpdateSecurityPolicy                      WorkRequestOperationTypeEnum = "UPDATE_SECURITY_POLICY"
	WorkRequestOperationTypeChangeSecurityPolicyCompartment           WorkRequestOperationTypeEnum = "CHANGE_SECURITY_POLICY_COMPARTMENT"
	WorkRequestOperationTypeUpdateSecurityPolicyDeployment            WorkRequestOperationTypeEnum = "UPDATE_SECURITY_POLICY_DEPLOYMENT"
	WorkRequestOperationTypeChangeSecurityPolicyDeploymentCompartment WorkRequestOperationTypeEnum = "CHANGE_SECURITY_POLICY_DEPLOYMENT_COMPARTMENT"
	WorkRequestOperationTypeAuditTrail                                WorkRequestOperationTypeEnum = "AUDIT_TRAIL"
	WorkRequestOperationTypeDeleteAuditTrail                          WorkRequestOperationTypeEnum = "DELETE_AUDIT_TRAIL"
	WorkRequestOperationTypeDiscoverAuditTrails                       WorkRequestOperationTypeEnum = "DISCOVER_AUDIT_TRAILS"
	WorkRequestOperationTypeUpdateAuditTrail                          WorkRequestOperationTypeEnum = "UPDATE_AUDIT_TRAIL"
	WorkRequestOperationTypeUpdateAuditProfile                        WorkRequestOperationTypeEnum = "UPDATE_AUDIT_PROFILE"
	WorkRequestOperationTypeAuditChangeCompartment                    WorkRequestOperationTypeEnum = "AUDIT_CHANGE_COMPARTMENT"
	WorkRequestOperationTypeCreateReportDefinition                    WorkRequestOperationTypeEnum = "CREATE_REPORT_DEFINITION"
	WorkRequestOperationTypeUpdateReportDefinition                    WorkRequestOperationTypeEnum = "UPDATE_REPORT_DEFINITION"
	WorkRequestOperationTypeChangeReportDefinitionCompartment         WorkRequestOperationTypeEnum = "CHANGE_REPORT_DEFINITION_COMPARTMENT"
	WorkRequestOperationTypeDeleteReportDefinition                    WorkRequestOperationTypeEnum = "DELETE_REPORT_DEFINITION"
	WorkRequestOperationTypeGenerateReport                            WorkRequestOperationTypeEnum = "GENERATE_REPORT"
	WorkRequestOperationTypeChangeReportCompartment                   WorkRequestOperationTypeEnum = "CHANGE_REPORT_COMPARTMENT"
	WorkRequestOperationTypeDeleteArchiveRetrieval                    WorkRequestOperationTypeEnum = "DELETE_ARCHIVE_RETRIEVAL"
	WorkRequestOperationTypeCreateArchiveRetrieval                    WorkRequestOperationTypeEnum = "CREATE_ARCHIVE_RETRIEVAL"
	WorkRequestOperationTypeUpdateArchiveRetrieval                    WorkRequestOperationTypeEnum = "UPDATE_ARCHIVE_RETRIEVAL"
	WorkRequestOperationTypeChangeArchiveRetrievalCompartment         WorkRequestOperationTypeEnum = "CHANGE_ARCHIVE_RETRIEVAL_COMPARTMENT"
	WorkRequestOperationTypeUpdateAlert                               WorkRequestOperationTypeEnum = "UPDATE_ALERT"
	WorkRequestOperationTypeTargetAlertPolicyAssociation              WorkRequestOperationTypeEnum = "TARGET_ALERT_POLICY_ASSOCIATION"
	WorkRequestOperationTypeCreateSensitiveDataModel                  WorkRequestOperationTypeEnum = "CREATE_SENSITIVE_DATA_MODEL"
	WorkRequestOperationTypeUpdateSensitiveDataModel                  WorkRequestOperationTypeEnum = "UPDATE_SENSITIVE_DATA_MODEL"
	WorkRequestOperationTypeDeleteSensitiveDataModel                  WorkRequestOperationTypeEnum = "DELETE_SENSITIVE_DATA_MODEL"
	WorkRequestOperationTypeUploadSensitiveDataModel                  WorkRequestOperationTypeEnum = "UPLOAD_SENSITIVE_DATA_MODEL"
	WorkRequestOperationTypeGenerateSensitiveDataModelForDownload     WorkRequestOperationTypeEnum = "GENERATE_SENSITIVE_DATA_MODEL_FOR_DOWNLOAD"
	WorkRequestOperationTypeCreateSensitiveColumn                     WorkRequestOperationTypeEnum = "CREATE_SENSITIVE_COLUMN"
	WorkRequestOperationTypeUpdateSensitiveColumn                     WorkRequestOperationTypeEnum = "UPDATE_SENSITIVE_COLUMN"
	WorkRequestOperationTypePatchSensitiveColumns                     WorkRequestOperationTypeEnum = "PATCH_SENSITIVE_COLUMNS"
	WorkRequestOperationTypeCreateDiscoveryJob                        WorkRequestOperationTypeEnum = "CREATE_DISCOVERY_JOB"
	WorkRequestOperationTypeDeleteDiscoveryJob                        WorkRequestOperationTypeEnum = "DELETE_DISCOVERY_JOB"
	WorkRequestOperationTypePatchDiscoveryJobResult                   WorkRequestOperationTypeEnum = "PATCH_DISCOVERY_JOB_RESULT"
	WorkRequestOperationTypeApplyDiscoveryJobResult                   WorkRequestOperationTypeEnum = "APPLY_DISCOVERY_JOB_RESULT"
	WorkRequestOperationTypeGenerateDiscoveryReport                   WorkRequestOperationTypeEnum = "GENERATE_DISCOVERY_REPORT"
	WorkRequestOperationTypeCreateSensitiveType                       WorkRequestOperationTypeEnum = "CREATE_SENSITIVE_TYPE"
	WorkRequestOperationTypeUpdateSensitiveType                       WorkRequestOperationTypeEnum = "UPDATE_SENSITIVE_TYPE"
	WorkRequestOperationTypeCreateMaskingPolicy                       WorkRequestOperationTypeEnum = "CREATE_MASKING_POLICY"
	WorkRequestOperationTypeUpdateMaskingPolicy                       WorkRequestOperationTypeEnum = "UPDATE_MASKING_POLICY"
	WorkRequestOperationTypeDeleteMaskingPolicy                       WorkRequestOperationTypeEnum = "DELETE_MASKING_POLICY"
	WorkRequestOperationTypeUploadMaskingPolicy                       WorkRequestOperationTypeEnum = "UPLOAD_MASKING_POLICY"
	WorkRequestOperationTypeGenerateMaskingPolicyForDownload          WorkRequestOperationTypeEnum = "GENERATE_MASKING_POLICY_FOR_DOWNLOAD"
	WorkRequestOperationTypeCreateMaskingColumn                       WorkRequestOperationTypeEnum = "CREATE_MASKING_COLUMN"
	WorkRequestOperationTypeUpdateMaskingColumn                       WorkRequestOperationTypeEnum = "UPDATE_MASKING_COLUMN"
	WorkRequestOperationTypePatchMaskingColumns                       WorkRequestOperationTypeEnum = "PATCH_MASKING_COLUMNS"
	WorkRequestOperationTypeGenerateMaskingReport                     WorkRequestOperationTypeEnum = "GENERATE_MASKING_REPORT"
	WorkRequestOperationTypeCreateLibraryMaskingFormat                WorkRequestOperationTypeEnum = "CREATE_LIBRARY_MASKING_FORMAT"
	WorkRequestOperationTypeUpdateLibraryMaskingFormat                WorkRequestOperationTypeEnum = "UPDATE_LIBRARY_MASKING_FORMAT"
	WorkRequestOperationTypeAddColumnsFromSdm                         WorkRequestOperationTypeEnum = "ADD_COLUMNS_FROM_SDM"
	WorkRequestOperationTypeMaskingJob                                WorkRequestOperationTypeEnum = "MASKING_JOB"
	WorkRequestOperationTypeCreateDifference                          WorkRequestOperationTypeEnum = "CREATE_DIFFERENCE"
	WorkRequestOperationTypeDeleteDifference                          WorkRequestOperationTypeEnum = "DELETE_DIFFERENCE"
	WorkRequestOperationTypeUpdateDifference                          WorkRequestOperationTypeEnum = "UPDATE_DIFFERENCE"
	WorkRequestOperationTypePatchDifference                           WorkRequestOperationTypeEnum = "PATCH_DIFFERENCE"
	WorkRequestOperationTypeApplyDifference                           WorkRequestOperationTypeEnum = "APPLY_DIFFERENCE"
	WorkRequestOperationTypeMaskPolicyGenerateHealthReport            WorkRequestOperationTypeEnum = "MASK_POLICY_GENERATE_HEALTH_REPORT"
	WorkRequestOperationTypeMaskPolicyDeleteHealthReport              WorkRequestOperationTypeEnum = "MASK_POLICY_DELETE_HEALTH_REPORT"
	WorkRequestOperationTypeAbortMasking                              WorkRequestOperationTypeEnum = "ABORT_MASKING"
	WorkRequestOperationTypeCreateSecurityPolicyReport                WorkRequestOperationTypeEnum = "CREATE_SECURITY_POLICY_REPORT"
	WorkRequestOperationTypeRefreshSecurityPolicyCache                WorkRequestOperationTypeEnum = "REFRESH_SECURITY_POLICY_CACHE"
	WorkRequestOperationTypeDeleteSecurityPolicyCache                 WorkRequestOperationTypeEnum = "DELETE_SECURITY_POLICY_CACHE"
	WorkRequestOperationTypeCreateSchedule                            WorkRequestOperationTypeEnum = "CREATE_SCHEDULE"
	WorkRequestOperationTypeRemoveScheduleReport                      WorkRequestOperationTypeEnum = "REMOVE_SCHEDULE_REPORT"
	WorkRequestOperationTypeUpdateAllAlert                            WorkRequestOperationTypeEnum = "UPDATE_ALL_ALERT"
	WorkRequestOperationTypePatchTargetAlertPolicyAssociation         WorkRequestOperationTypeEnum = "PATCH_TARGET_ALERT_POLICY_ASSOCIATION"
)

var mappingWorkRequestOperationTypeEnum = map[string]WorkRequestOperationTypeEnum{
	"ENABLE_DATA_SAFE_CONFIGURATION":                WorkRequestOperationTypeEnableDataSafeConfiguration,
	"CREATE_PRIVATE_ENDPOINT":                       WorkRequestOperationTypeCreatePrivateEndpoint,
	"UPDATE_PRIVATE_ENDPOINT":                       WorkRequestOperationTypeUpdatePrivateEndpoint,
	"DELETE_PRIVATE_ENDPOINT":                       WorkRequestOperationTypeDeletePrivateEndpoint,
	"CHANGE_PRIVATE_ENDPOINT_COMPARTMENT":           WorkRequestOperationTypeChangePrivateEndpointCompartment,
	"CREATE_ONPREM_CONNECTOR":                       WorkRequestOperationTypeCreateOnpremConnector,
	"UPDATE_ONPREM_CONNECTOR":                       WorkRequestOperationTypeUpdateOnpremConnector,
	"DELETE_ONPREM_CONNECTOR":                       WorkRequestOperationTypeDeleteOnpremConnector,
	"UPDATE_ONPREM_CONNECTOR_WALLET":                WorkRequestOperationTypeUpdateOnpremConnectorWallet,
	"CHANGE_ONPREM_CONNECTOR_COMPARTMENT":           WorkRequestOperationTypeChangeOnpremConnectorCompartment,
	"CREATE_TARGET_DATABASE":                        WorkRequestOperationTypeCreateTargetDatabase,
	"UPDATE_TARGET_DATABASE":                        WorkRequestOperationTypeUpdateTargetDatabase,
	"ACTIVATE_TARGET_DATABASE":                      WorkRequestOperationTypeActivateTargetDatabase,
	"DEACTIVATE_TARGET_DATABASE":                    WorkRequestOperationTypeDeactivateTargetDatabase,
	"DELETE_TARGET_DATABASE":                        WorkRequestOperationTypeDeleteTargetDatabase,
	"CHANGE_TARGET_DATABASE_COMPARTMENT":            WorkRequestOperationTypeChangeTargetDatabaseCompartment,
	"CREATE_PEER_TARGET_DATABASE":                   WorkRequestOperationTypeCreatePeerTargetDatabase,
	"UPDATE_PEER_TARGET_DATABASE":                   WorkRequestOperationTypeUpdatePeerTargetDatabase,
	"DELETE_PEER_TARGET_DATABASE":                   WorkRequestOperationTypeDeletePeerTargetDatabase,
	"REFRESH_TARGET_DATABASE":                       WorkRequestOperationTypeRefreshTargetDatabase,
	"PROVISION_POLICY":                              WorkRequestOperationTypeProvisionPolicy,
	"RETRIEVE_POLICY":                               WorkRequestOperationTypeRetrievePolicy,
	"UPDATE_POLICY":                                 WorkRequestOperationTypeUpdatePolicy,
	"CHANGE_POLICY_COMPARTMENT":                     WorkRequestOperationTypeChangePolicyCompartment,
	"CREATE_USER_ASSESSMENT":                        WorkRequestOperationTypeCreateUserAssessment,
	"ASSESS_USER_ASSESSMENT":                        WorkRequestOperationTypeAssessUserAssessment,
	"CREATE_SNAPSHOT_USER_ASSESSMENT":               WorkRequestOperationTypeCreateSnapshotUserAssessment,
	"CREATE_SCHEDULE_USER_ASSESSMENT":               WorkRequestOperationTypeCreateScheduleUserAssessment,
	"COMPARE_WITH_BASELINE_USER_ASSESSMENT":         WorkRequestOperationTypeCompareWithBaselineUserAssessment,
	"DELETE_USER_ASSESSMENT":                        WorkRequestOperationTypeDeleteUserAssessment,
	"UPDATE_USER_ASSESSMENT":                        WorkRequestOperationTypeUpdateUserAssessment,
	"CHANGE_USER_ASSESSMENT_COMPARTMENT":            WorkRequestOperationTypeChangeUserAssessmentCompartment,
	"SET_USER_ASSESSMENT_BASELINE":                  WorkRequestOperationTypeSetUserAssessmentBaseline,
	"UNSET_USER_ASSESSMENT_BASELINE":                WorkRequestOperationTypeUnsetUserAssessmentBaseline,
	"GENERATE_USER_ASSESSMENT_REPORT":               WorkRequestOperationTypeGenerateUserAssessmentReport,
	"CREATE_SECURITY_ASSESSMENT":                    WorkRequestOperationTypeCreateSecurityAssessment,
	"CREATE_SECURITY_ASSESSMENT_NOW":                WorkRequestOperationTypeCreateSecurityAssessmentNow,
	"ASSESS_SECURITY_ASSESSMENT":                    WorkRequestOperationTypeAssessSecurityAssessment,
	"CREATE_SNAPSHOT_SECURITY_ASSESSMENT":           WorkRequestOperationTypeCreateSnapshotSecurityAssessment,
	"CREATE_SCHEDULE_SECURITY_ASSESSMENT":           WorkRequestOperationTypeCreateScheduleSecurityAssessment,
	"COMPARE_WITH_BASELINE_SECURITY_ASSESSMENT":     WorkRequestOperationTypeCompareWithBaselineSecurityAssessment,
	"DELETE_SECURITY_ASSESSMENT":                    WorkRequestOperationTypeDeleteSecurityAssessment,
	"UPDATE_SECURITY_ASSESSMENT":                    WorkRequestOperationTypeUpdateSecurityAssessment,
	"UPDATE_FINDING_RISK":                           WorkRequestOperationTypeUpdateFindingRisk,
	"CHANGE_SECURITY_ASSESSMENT_COMPARTMENT":        WorkRequestOperationTypeChangeSecurityAssessmentCompartment,
	"SET_SECURITY_ASSESSMENT_BASELINE":              WorkRequestOperationTypeSetSecurityAssessmentBaseline,
	"UNSET_SECURITY_ASSESSMENT_BASELINE":            WorkRequestOperationTypeUnsetSecurityAssessmentBaseline,
	"GENERATE_SECURITY_ASSESSMENT_REPORT":           WorkRequestOperationTypeGenerateSecurityAssessmentReport,
	"CREATE_AUDIT_PROFILE":                          WorkRequestOperationTypeCreateAuditProfile,
	"CALCULATE_VOLUME":                              WorkRequestOperationTypeCalculateVolume,
	"CALCULATE_COLLECTED_VOLUME":                    WorkRequestOperationTypeCalculateCollectedVolume,
	"CREATE_DB_SECURITY_CONFIG":                     WorkRequestOperationTypeCreateDbSecurityConfig,
	"REFRESH_DB_SECURITY_CONFIG":                    WorkRequestOperationTypeRefreshDbSecurityConfig,
	"UPDATE_DB_SECURITY_CONFIG":                     WorkRequestOperationTypeUpdateDbSecurityConfig,
	"CHANGE_DB_SECURITY_CONFIG_COMPARTMENT":         WorkRequestOperationTypeChangeDbSecurityConfigCompartment,
	"GENERATE_FIREWALL_POLICY":                      WorkRequestOperationTypeGenerateFirewallPolicy,
	"UPDATE_FIREWALL_POLICY":                        WorkRequestOperationTypeUpdateFirewallPolicy,
	"CHANGE_FIREWALL_POLICY_COMPARTMENT":            WorkRequestOperationTypeChangeFirewallPolicyCompartment,
	"DELETE_FIREWALL_POLICY":                        WorkRequestOperationTypeDeleteFirewallPolicy,
	"CREATE_SQL_COLLECTION":                         WorkRequestOperationTypeCreateSqlCollection,
	"UPDATE_SQL_COLLECTION":                         WorkRequestOperationTypeUpdateSqlCollection,
	"START_SQL_COLLECTION":                          WorkRequestOperationTypeStartSqlCollection,
	"STOP_SQL_COLLECTION":                           WorkRequestOperationTypeStopSqlCollection,
	"DELETE_SQL_COLLECTION":                         WorkRequestOperationTypeDeleteSqlCollection,
	"CHANGE_SQL_COLLECTION_COMPARTMENT":             WorkRequestOperationTypeChangeSqlCollectionCompartment,
	"REFRESH_SQL_COLLECTION_LOG_INSIGHTS":           WorkRequestOperationTypeRefreshSqlCollectionLogInsights,
	"PURGE_SQL_COLLECTION_LOGS":                     WorkRequestOperationTypePurgeSqlCollectionLogs,
	"REFRESH_VIOLATIONS":                            WorkRequestOperationTypeRefreshViolations,
	"CREATE_ARCHIVAL":                               WorkRequestOperationTypeCreateArchival,
	"UPDATE_SECURITY_POLICY":                        WorkRequestOperationTypeUpdateSecurityPolicy,
	"CHANGE_SECURITY_POLICY_COMPARTMENT":            WorkRequestOperationTypeChangeSecurityPolicyCompartment,
	"UPDATE_SECURITY_POLICY_DEPLOYMENT":             WorkRequestOperationTypeUpdateSecurityPolicyDeployment,
	"CHANGE_SECURITY_POLICY_DEPLOYMENT_COMPARTMENT": WorkRequestOperationTypeChangeSecurityPolicyDeploymentCompartment,
	"AUDIT_TRAIL":                                   WorkRequestOperationTypeAuditTrail,
	"DELETE_AUDIT_TRAIL":                            WorkRequestOperationTypeDeleteAuditTrail,
	"DISCOVER_AUDIT_TRAILS":                         WorkRequestOperationTypeDiscoverAuditTrails,
	"UPDATE_AUDIT_TRAIL":                            WorkRequestOperationTypeUpdateAuditTrail,
	"UPDATE_AUDIT_PROFILE":                          WorkRequestOperationTypeUpdateAuditProfile,
	"AUDIT_CHANGE_COMPARTMENT":                      WorkRequestOperationTypeAuditChangeCompartment,
	"CREATE_REPORT_DEFINITION":                      WorkRequestOperationTypeCreateReportDefinition,
	"UPDATE_REPORT_DEFINITION":                      WorkRequestOperationTypeUpdateReportDefinition,
	"CHANGE_REPORT_DEFINITION_COMPARTMENT":          WorkRequestOperationTypeChangeReportDefinitionCompartment,
	"DELETE_REPORT_DEFINITION":                      WorkRequestOperationTypeDeleteReportDefinition,
	"GENERATE_REPORT":                               WorkRequestOperationTypeGenerateReport,
	"CHANGE_REPORT_COMPARTMENT":                     WorkRequestOperationTypeChangeReportCompartment,
	"DELETE_ARCHIVE_RETRIEVAL":                      WorkRequestOperationTypeDeleteArchiveRetrieval,
	"CREATE_ARCHIVE_RETRIEVAL":                      WorkRequestOperationTypeCreateArchiveRetrieval,
	"UPDATE_ARCHIVE_RETRIEVAL":                      WorkRequestOperationTypeUpdateArchiveRetrieval,
	"CHANGE_ARCHIVE_RETRIEVAL_COMPARTMENT":          WorkRequestOperationTypeChangeArchiveRetrievalCompartment,
	"UPDATE_ALERT":                                  WorkRequestOperationTypeUpdateAlert,
	"TARGET_ALERT_POLICY_ASSOCIATION":               WorkRequestOperationTypeTargetAlertPolicyAssociation,
	"CREATE_SENSITIVE_DATA_MODEL":                   WorkRequestOperationTypeCreateSensitiveDataModel,
	"UPDATE_SENSITIVE_DATA_MODEL":                   WorkRequestOperationTypeUpdateSensitiveDataModel,
	"DELETE_SENSITIVE_DATA_MODEL":                   WorkRequestOperationTypeDeleteSensitiveDataModel,
	"UPLOAD_SENSITIVE_DATA_MODEL":                   WorkRequestOperationTypeUploadSensitiveDataModel,
	"GENERATE_SENSITIVE_DATA_MODEL_FOR_DOWNLOAD":    WorkRequestOperationTypeGenerateSensitiveDataModelForDownload,
	"CREATE_SENSITIVE_COLUMN":                       WorkRequestOperationTypeCreateSensitiveColumn,
	"UPDATE_SENSITIVE_COLUMN":                       WorkRequestOperationTypeUpdateSensitiveColumn,
	"PATCH_SENSITIVE_COLUMNS":                       WorkRequestOperationTypePatchSensitiveColumns,
	"CREATE_DISCOVERY_JOB":                          WorkRequestOperationTypeCreateDiscoveryJob,
	"DELETE_DISCOVERY_JOB":                          WorkRequestOperationTypeDeleteDiscoveryJob,
	"PATCH_DISCOVERY_JOB_RESULT":                    WorkRequestOperationTypePatchDiscoveryJobResult,
	"APPLY_DISCOVERY_JOB_RESULT":                    WorkRequestOperationTypeApplyDiscoveryJobResult,
	"GENERATE_DISCOVERY_REPORT":                     WorkRequestOperationTypeGenerateDiscoveryReport,
	"CREATE_SENSITIVE_TYPE":                         WorkRequestOperationTypeCreateSensitiveType,
	"UPDATE_SENSITIVE_TYPE":                         WorkRequestOperationTypeUpdateSensitiveType,
	"CREATE_MASKING_POLICY":                         WorkRequestOperationTypeCreateMaskingPolicy,
	"UPDATE_MASKING_POLICY":                         WorkRequestOperationTypeUpdateMaskingPolicy,
	"DELETE_MASKING_POLICY":                         WorkRequestOperationTypeDeleteMaskingPolicy,
	"UPLOAD_MASKING_POLICY":                         WorkRequestOperationTypeUploadMaskingPolicy,
	"GENERATE_MASKING_POLICY_FOR_DOWNLOAD":          WorkRequestOperationTypeGenerateMaskingPolicyForDownload,
	"CREATE_MASKING_COLUMN":                         WorkRequestOperationTypeCreateMaskingColumn,
	"UPDATE_MASKING_COLUMN":                         WorkRequestOperationTypeUpdateMaskingColumn,
	"PATCH_MASKING_COLUMNS":                         WorkRequestOperationTypePatchMaskingColumns,
	"GENERATE_MASKING_REPORT":                       WorkRequestOperationTypeGenerateMaskingReport,
	"CREATE_LIBRARY_MASKING_FORMAT":                 WorkRequestOperationTypeCreateLibraryMaskingFormat,
	"UPDATE_LIBRARY_MASKING_FORMAT":                 WorkRequestOperationTypeUpdateLibraryMaskingFormat,
	"ADD_COLUMNS_FROM_SDM":                          WorkRequestOperationTypeAddColumnsFromSdm,
	"MASKING_JOB":                                   WorkRequestOperationTypeMaskingJob,
	"CREATE_DIFFERENCE":                             WorkRequestOperationTypeCreateDifference,
	"DELETE_DIFFERENCE":                             WorkRequestOperationTypeDeleteDifference,
	"UPDATE_DIFFERENCE":                             WorkRequestOperationTypeUpdateDifference,
	"PATCH_DIFFERENCE":                              WorkRequestOperationTypePatchDifference,
	"APPLY_DIFFERENCE":                              WorkRequestOperationTypeApplyDifference,
	"MASK_POLICY_GENERATE_HEALTH_REPORT":            WorkRequestOperationTypeMaskPolicyGenerateHealthReport,
	"MASK_POLICY_DELETE_HEALTH_REPORT":              WorkRequestOperationTypeMaskPolicyDeleteHealthReport,
	"ABORT_MASKING":                                 WorkRequestOperationTypeAbortMasking,
	"CREATE_SECURITY_POLICY_REPORT":                 WorkRequestOperationTypeCreateSecurityPolicyReport,
	"REFRESH_SECURITY_POLICY_CACHE":                 WorkRequestOperationTypeRefreshSecurityPolicyCache,
	"DELETE_SECURITY_POLICY_CACHE":                  WorkRequestOperationTypeDeleteSecurityPolicyCache,
	"CREATE_SCHEDULE":                               WorkRequestOperationTypeCreateSchedule,
	"REMOVE_SCHEDULE_REPORT":                        WorkRequestOperationTypeRemoveScheduleReport,
	"UPDATE_ALL_ALERT":                              WorkRequestOperationTypeUpdateAllAlert,
	"PATCH_TARGET_ALERT_POLICY_ASSOCIATION":         WorkRequestOperationTypePatchTargetAlertPolicyAssociation,
}

var mappingWorkRequestOperationTypeEnumLowerCase = map[string]WorkRequestOperationTypeEnum{
	"enable_data_safe_configuration":                WorkRequestOperationTypeEnableDataSafeConfiguration,
	"create_private_endpoint":                       WorkRequestOperationTypeCreatePrivateEndpoint,
	"update_private_endpoint":                       WorkRequestOperationTypeUpdatePrivateEndpoint,
	"delete_private_endpoint":                       WorkRequestOperationTypeDeletePrivateEndpoint,
	"change_private_endpoint_compartment":           WorkRequestOperationTypeChangePrivateEndpointCompartment,
	"create_onprem_connector":                       WorkRequestOperationTypeCreateOnpremConnector,
	"update_onprem_connector":                       WorkRequestOperationTypeUpdateOnpremConnector,
	"delete_onprem_connector":                       WorkRequestOperationTypeDeleteOnpremConnector,
	"update_onprem_connector_wallet":                WorkRequestOperationTypeUpdateOnpremConnectorWallet,
	"change_onprem_connector_compartment":           WorkRequestOperationTypeChangeOnpremConnectorCompartment,
	"create_target_database":                        WorkRequestOperationTypeCreateTargetDatabase,
	"update_target_database":                        WorkRequestOperationTypeUpdateTargetDatabase,
	"activate_target_database":                      WorkRequestOperationTypeActivateTargetDatabase,
	"deactivate_target_database":                    WorkRequestOperationTypeDeactivateTargetDatabase,
	"delete_target_database":                        WorkRequestOperationTypeDeleteTargetDatabase,
	"change_target_database_compartment":            WorkRequestOperationTypeChangeTargetDatabaseCompartment,
	"create_peer_target_database":                   WorkRequestOperationTypeCreatePeerTargetDatabase,
	"update_peer_target_database":                   WorkRequestOperationTypeUpdatePeerTargetDatabase,
	"delete_peer_target_database":                   WorkRequestOperationTypeDeletePeerTargetDatabase,
	"refresh_target_database":                       WorkRequestOperationTypeRefreshTargetDatabase,
	"provision_policy":                              WorkRequestOperationTypeProvisionPolicy,
	"retrieve_policy":                               WorkRequestOperationTypeRetrievePolicy,
	"update_policy":                                 WorkRequestOperationTypeUpdatePolicy,
	"change_policy_compartment":                     WorkRequestOperationTypeChangePolicyCompartment,
	"create_user_assessment":                        WorkRequestOperationTypeCreateUserAssessment,
	"assess_user_assessment":                        WorkRequestOperationTypeAssessUserAssessment,
	"create_snapshot_user_assessment":               WorkRequestOperationTypeCreateSnapshotUserAssessment,
	"create_schedule_user_assessment":               WorkRequestOperationTypeCreateScheduleUserAssessment,
	"compare_with_baseline_user_assessment":         WorkRequestOperationTypeCompareWithBaselineUserAssessment,
	"delete_user_assessment":                        WorkRequestOperationTypeDeleteUserAssessment,
	"update_user_assessment":                        WorkRequestOperationTypeUpdateUserAssessment,
	"change_user_assessment_compartment":            WorkRequestOperationTypeChangeUserAssessmentCompartment,
	"set_user_assessment_baseline":                  WorkRequestOperationTypeSetUserAssessmentBaseline,
	"unset_user_assessment_baseline":                WorkRequestOperationTypeUnsetUserAssessmentBaseline,
	"generate_user_assessment_report":               WorkRequestOperationTypeGenerateUserAssessmentReport,
	"create_security_assessment":                    WorkRequestOperationTypeCreateSecurityAssessment,
	"create_security_assessment_now":                WorkRequestOperationTypeCreateSecurityAssessmentNow,
	"assess_security_assessment":                    WorkRequestOperationTypeAssessSecurityAssessment,
	"create_snapshot_security_assessment":           WorkRequestOperationTypeCreateSnapshotSecurityAssessment,
	"create_schedule_security_assessment":           WorkRequestOperationTypeCreateScheduleSecurityAssessment,
	"compare_with_baseline_security_assessment":     WorkRequestOperationTypeCompareWithBaselineSecurityAssessment,
	"delete_security_assessment":                    WorkRequestOperationTypeDeleteSecurityAssessment,
	"update_security_assessment":                    WorkRequestOperationTypeUpdateSecurityAssessment,
	"update_finding_risk":                           WorkRequestOperationTypeUpdateFindingRisk,
	"change_security_assessment_compartment":        WorkRequestOperationTypeChangeSecurityAssessmentCompartment,
	"set_security_assessment_baseline":              WorkRequestOperationTypeSetSecurityAssessmentBaseline,
	"unset_security_assessment_baseline":            WorkRequestOperationTypeUnsetSecurityAssessmentBaseline,
	"generate_security_assessment_report":           WorkRequestOperationTypeGenerateSecurityAssessmentReport,
	"create_audit_profile":                          WorkRequestOperationTypeCreateAuditProfile,
	"calculate_volume":                              WorkRequestOperationTypeCalculateVolume,
	"calculate_collected_volume":                    WorkRequestOperationTypeCalculateCollectedVolume,
	"create_db_security_config":                     WorkRequestOperationTypeCreateDbSecurityConfig,
	"refresh_db_security_config":                    WorkRequestOperationTypeRefreshDbSecurityConfig,
	"update_db_security_config":                     WorkRequestOperationTypeUpdateDbSecurityConfig,
	"change_db_security_config_compartment":         WorkRequestOperationTypeChangeDbSecurityConfigCompartment,
	"generate_firewall_policy":                      WorkRequestOperationTypeGenerateFirewallPolicy,
	"update_firewall_policy":                        WorkRequestOperationTypeUpdateFirewallPolicy,
	"change_firewall_policy_compartment":            WorkRequestOperationTypeChangeFirewallPolicyCompartment,
	"delete_firewall_policy":                        WorkRequestOperationTypeDeleteFirewallPolicy,
	"create_sql_collection":                         WorkRequestOperationTypeCreateSqlCollection,
	"update_sql_collection":                         WorkRequestOperationTypeUpdateSqlCollection,
	"start_sql_collection":                          WorkRequestOperationTypeStartSqlCollection,
	"stop_sql_collection":                           WorkRequestOperationTypeStopSqlCollection,
	"delete_sql_collection":                         WorkRequestOperationTypeDeleteSqlCollection,
	"change_sql_collection_compartment":             WorkRequestOperationTypeChangeSqlCollectionCompartment,
	"refresh_sql_collection_log_insights":           WorkRequestOperationTypeRefreshSqlCollectionLogInsights,
	"purge_sql_collection_logs":                     WorkRequestOperationTypePurgeSqlCollectionLogs,
	"refresh_violations":                            WorkRequestOperationTypeRefreshViolations,
	"create_archival":                               WorkRequestOperationTypeCreateArchival,
	"update_security_policy":                        WorkRequestOperationTypeUpdateSecurityPolicy,
	"change_security_policy_compartment":            WorkRequestOperationTypeChangeSecurityPolicyCompartment,
	"update_security_policy_deployment":             WorkRequestOperationTypeUpdateSecurityPolicyDeployment,
	"change_security_policy_deployment_compartment": WorkRequestOperationTypeChangeSecurityPolicyDeploymentCompartment,
	"audit_trail":                                   WorkRequestOperationTypeAuditTrail,
	"delete_audit_trail":                            WorkRequestOperationTypeDeleteAuditTrail,
	"discover_audit_trails":                         WorkRequestOperationTypeDiscoverAuditTrails,
	"update_audit_trail":                            WorkRequestOperationTypeUpdateAuditTrail,
	"update_audit_profile":                          WorkRequestOperationTypeUpdateAuditProfile,
	"audit_change_compartment":                      WorkRequestOperationTypeAuditChangeCompartment,
	"create_report_definition":                      WorkRequestOperationTypeCreateReportDefinition,
	"update_report_definition":                      WorkRequestOperationTypeUpdateReportDefinition,
	"change_report_definition_compartment":          WorkRequestOperationTypeChangeReportDefinitionCompartment,
	"delete_report_definition":                      WorkRequestOperationTypeDeleteReportDefinition,
	"generate_report":                               WorkRequestOperationTypeGenerateReport,
	"change_report_compartment":                     WorkRequestOperationTypeChangeReportCompartment,
	"delete_archive_retrieval":                      WorkRequestOperationTypeDeleteArchiveRetrieval,
	"create_archive_retrieval":                      WorkRequestOperationTypeCreateArchiveRetrieval,
	"update_archive_retrieval":                      WorkRequestOperationTypeUpdateArchiveRetrieval,
	"change_archive_retrieval_compartment":          WorkRequestOperationTypeChangeArchiveRetrievalCompartment,
	"update_alert":                                  WorkRequestOperationTypeUpdateAlert,
	"target_alert_policy_association":               WorkRequestOperationTypeTargetAlertPolicyAssociation,
	"create_sensitive_data_model":                   WorkRequestOperationTypeCreateSensitiveDataModel,
	"update_sensitive_data_model":                   WorkRequestOperationTypeUpdateSensitiveDataModel,
	"delete_sensitive_data_model":                   WorkRequestOperationTypeDeleteSensitiveDataModel,
	"upload_sensitive_data_model":                   WorkRequestOperationTypeUploadSensitiveDataModel,
	"generate_sensitive_data_model_for_download":    WorkRequestOperationTypeGenerateSensitiveDataModelForDownload,
	"create_sensitive_column":                       WorkRequestOperationTypeCreateSensitiveColumn,
	"update_sensitive_column":                       WorkRequestOperationTypeUpdateSensitiveColumn,
	"patch_sensitive_columns":                       WorkRequestOperationTypePatchSensitiveColumns,
	"create_discovery_job":                          WorkRequestOperationTypeCreateDiscoveryJob,
	"delete_discovery_job":                          WorkRequestOperationTypeDeleteDiscoveryJob,
	"patch_discovery_job_result":                    WorkRequestOperationTypePatchDiscoveryJobResult,
	"apply_discovery_job_result":                    WorkRequestOperationTypeApplyDiscoveryJobResult,
	"generate_discovery_report":                     WorkRequestOperationTypeGenerateDiscoveryReport,
	"create_sensitive_type":                         WorkRequestOperationTypeCreateSensitiveType,
	"update_sensitive_type":                         WorkRequestOperationTypeUpdateSensitiveType,
	"create_masking_policy":                         WorkRequestOperationTypeCreateMaskingPolicy,
	"update_masking_policy":                         WorkRequestOperationTypeUpdateMaskingPolicy,
	"delete_masking_policy":                         WorkRequestOperationTypeDeleteMaskingPolicy,
	"upload_masking_policy":                         WorkRequestOperationTypeUploadMaskingPolicy,
	"generate_masking_policy_for_download":          WorkRequestOperationTypeGenerateMaskingPolicyForDownload,
	"create_masking_column":                         WorkRequestOperationTypeCreateMaskingColumn,
	"update_masking_column":                         WorkRequestOperationTypeUpdateMaskingColumn,
	"patch_masking_columns":                         WorkRequestOperationTypePatchMaskingColumns,
	"generate_masking_report":                       WorkRequestOperationTypeGenerateMaskingReport,
	"create_library_masking_format":                 WorkRequestOperationTypeCreateLibraryMaskingFormat,
	"update_library_masking_format":                 WorkRequestOperationTypeUpdateLibraryMaskingFormat,
	"add_columns_from_sdm":                          WorkRequestOperationTypeAddColumnsFromSdm,
	"masking_job":                                   WorkRequestOperationTypeMaskingJob,
	"create_difference":                             WorkRequestOperationTypeCreateDifference,
	"delete_difference":                             WorkRequestOperationTypeDeleteDifference,
	"update_difference":                             WorkRequestOperationTypeUpdateDifference,
	"patch_difference":                              WorkRequestOperationTypePatchDifference,
	"apply_difference":                              WorkRequestOperationTypeApplyDifference,
	"mask_policy_generate_health_report":            WorkRequestOperationTypeMaskPolicyGenerateHealthReport,
	"mask_policy_delete_health_report":              WorkRequestOperationTypeMaskPolicyDeleteHealthReport,
	"abort_masking":                                 WorkRequestOperationTypeAbortMasking,
	"create_security_policy_report":                 WorkRequestOperationTypeCreateSecurityPolicyReport,
	"refresh_security_policy_cache":                 WorkRequestOperationTypeRefreshSecurityPolicyCache,
	"delete_security_policy_cache":                  WorkRequestOperationTypeDeleteSecurityPolicyCache,
	"create_schedule":                               WorkRequestOperationTypeCreateSchedule,
	"remove_schedule_report":                        WorkRequestOperationTypeRemoveScheduleReport,
	"update_all_alert":                              WorkRequestOperationTypeUpdateAllAlert,
	"patch_target_alert_policy_association":         WorkRequestOperationTypePatchTargetAlertPolicyAssociation,
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
		"CREATE_TARGET_DATABASE",
		"UPDATE_TARGET_DATABASE",
		"ACTIVATE_TARGET_DATABASE",
		"DEACTIVATE_TARGET_DATABASE",
		"DELETE_TARGET_DATABASE",
		"CHANGE_TARGET_DATABASE_COMPARTMENT",
		"CREATE_PEER_TARGET_DATABASE",
		"UPDATE_PEER_TARGET_DATABASE",
		"DELETE_PEER_TARGET_DATABASE",
		"REFRESH_TARGET_DATABASE",
		"PROVISION_POLICY",
		"RETRIEVE_POLICY",
		"UPDATE_POLICY",
		"CHANGE_POLICY_COMPARTMENT",
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
		"UPDATE_FINDING_RISK",
		"CHANGE_SECURITY_ASSESSMENT_COMPARTMENT",
		"SET_SECURITY_ASSESSMENT_BASELINE",
		"UNSET_SECURITY_ASSESSMENT_BASELINE",
		"GENERATE_SECURITY_ASSESSMENT_REPORT",
		"CREATE_AUDIT_PROFILE",
		"CALCULATE_VOLUME",
		"CALCULATE_COLLECTED_VOLUME",
		"CREATE_DB_SECURITY_CONFIG",
		"REFRESH_DB_SECURITY_CONFIG",
		"UPDATE_DB_SECURITY_CONFIG",
		"CHANGE_DB_SECURITY_CONFIG_COMPARTMENT",
		"GENERATE_FIREWALL_POLICY",
		"UPDATE_FIREWALL_POLICY",
		"CHANGE_FIREWALL_POLICY_COMPARTMENT",
		"DELETE_FIREWALL_POLICY",
		"CREATE_SQL_COLLECTION",
		"UPDATE_SQL_COLLECTION",
		"START_SQL_COLLECTION",
		"STOP_SQL_COLLECTION",
		"DELETE_SQL_COLLECTION",
		"CHANGE_SQL_COLLECTION_COMPARTMENT",
		"REFRESH_SQL_COLLECTION_LOG_INSIGHTS",
		"PURGE_SQL_COLLECTION_LOGS",
		"REFRESH_VIOLATIONS",
		"CREATE_ARCHIVAL",
		"UPDATE_SECURITY_POLICY",
		"CHANGE_SECURITY_POLICY_COMPARTMENT",
		"UPDATE_SECURITY_POLICY_DEPLOYMENT",
		"CHANGE_SECURITY_POLICY_DEPLOYMENT_COMPARTMENT",
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
		"CREATE_DIFFERENCE",
		"DELETE_DIFFERENCE",
		"UPDATE_DIFFERENCE",
		"PATCH_DIFFERENCE",
		"APPLY_DIFFERENCE",
		"MASK_POLICY_GENERATE_HEALTH_REPORT",
		"MASK_POLICY_DELETE_HEALTH_REPORT",
		"ABORT_MASKING",
		"CREATE_SECURITY_POLICY_REPORT",
		"REFRESH_SECURITY_POLICY_CACHE",
		"DELETE_SECURITY_POLICY_CACHE",
		"CREATE_SCHEDULE",
		"REMOVE_SCHEDULE_REPORT",
		"UPDATE_ALL_ALERT",
		"PATCH_TARGET_ALERT_POLICY_ASSOCIATION",
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
	WorkRequestStatusSuspending WorkRequestStatusEnum = "SUSPENDING"
	WorkRequestStatusSuspended  WorkRequestStatusEnum = "SUSPENDED"
)

var mappingWorkRequestStatusEnum = map[string]WorkRequestStatusEnum{
	"ACCEPTED":    WorkRequestStatusAccepted,
	"IN_PROGRESS": WorkRequestStatusInProgress,
	"FAILED":      WorkRequestStatusFailed,
	"SUCCEEDED":   WorkRequestStatusSucceeded,
	"CANCELING":   WorkRequestStatusCanceling,
	"CANCELED":    WorkRequestStatusCanceled,
	"SUSPENDING":  WorkRequestStatusSuspending,
	"SUSPENDED":   WorkRequestStatusSuspended,
}

var mappingWorkRequestStatusEnumLowerCase = map[string]WorkRequestStatusEnum{
	"accepted":    WorkRequestStatusAccepted,
	"in_progress": WorkRequestStatusInProgress,
	"failed":      WorkRequestStatusFailed,
	"succeeded":   WorkRequestStatusSucceeded,
	"canceling":   WorkRequestStatusCanceling,
	"canceled":    WorkRequestStatusCanceled,
	"suspending":  WorkRequestStatusSuspending,
	"suspended":   WorkRequestStatusSuspended,
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
		"SUSPENDING",
		"SUSPENDED",
	}
}

// GetMappingWorkRequestStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingWorkRequestStatusEnum(val string) (WorkRequestStatusEnum, bool) {
	enum, ok := mappingWorkRequestStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
