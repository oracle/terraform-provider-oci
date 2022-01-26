// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"github.com/oracle/oci-go-sdk/v56/common"
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
)

var mappingWorkRequestSummaryOperationType = map[string]WorkRequestSummaryOperationTypeEnum{
	"ENABLE_DATA_SAFE_CONFIGURATION":            WorkRequestSummaryOperationTypeEnableDataSafeConfiguration,
	"CREATE_PRIVATE_ENDPOINT":                   WorkRequestSummaryOperationTypeCreatePrivateEndpoint,
	"UPDATE_PRIVATE_ENDPOINT":                   WorkRequestSummaryOperationTypeUpdatePrivateEndpoint,
	"DELETE_PRIVATE_ENDPOINT":                   WorkRequestSummaryOperationTypeDeletePrivateEndpoint,
	"CHANGE_PRIVATE_ENDPOINT_COMPARTMENT":       WorkRequestSummaryOperationTypeChangePrivateEndpointCompartment,
	"CREATE_ONPREM_CONNECTOR":                   WorkRequestSummaryOperationTypeCreateOnpremConnector,
	"UPDATE_ONPREM_CONNECTOR":                   WorkRequestSummaryOperationTypeUpdateOnpremConnector,
	"DELETE_ONPREM_CONNECTOR":                   WorkRequestSummaryOperationTypeDeleteOnpremConnector,
	"UPDATE_ONPREM_CONNECTOR_WALLET":            WorkRequestSummaryOperationTypeUpdateOnpremConnectorWallet,
	"CHANGE_ONPREM_CONNECTOR_COMPARTMENT":       WorkRequestSummaryOperationTypeChangeOnpremConnectorCompartment,
	"CREATE_TARGET_DATABASE":                    WorkRequestSummaryOperationTypeCreateTargetDatabase,
	"UPDATE_TARGET_DATABASE":                    WorkRequestSummaryOperationTypeUpdateTargetDatabase,
	"ACTIVATE_TARGET_DATABASE":                  WorkRequestSummaryOperationTypeActivateTargetDatabase,
	"DEACTIVATE_TARGET_DATABASE":                WorkRequestSummaryOperationTypeDeactivateTargetDatabase,
	"DELETE_TARGET_DATABASE":                    WorkRequestSummaryOperationTypeDeleteTargetDatabase,
	"CHANGE_TARGET_DATABASE_COMPARTMENT":        WorkRequestSummaryOperationTypeChangeTargetDatabaseCompartment,
	"CREATE_USER_ASSESSMENT":                    WorkRequestSummaryOperationTypeCreateUserAssessment,
	"ASSESS_USER_ASSESSMENT":                    WorkRequestSummaryOperationTypeAssessUserAssessment,
	"CREATE_SNAPSHOT_USER_ASSESSMENT":           WorkRequestSummaryOperationTypeCreateSnapshotUserAssessment,
	"CREATE_SCHEDULE_USER_ASSESSMENT":           WorkRequestSummaryOperationTypeCreateScheduleUserAssessment,
	"COMPARE_WITH_BASELINE_USER_ASSESSMENT":     WorkRequestSummaryOperationTypeCompareWithBaselineUserAssessment,
	"DELETE_USER_ASSESSMENT":                    WorkRequestSummaryOperationTypeDeleteUserAssessment,
	"UPDATE_USER_ASSESSMENT":                    WorkRequestSummaryOperationTypeUpdateUserAssessment,
	"CHANGE_USER_ASSESSMENT_COMPARTMENT":        WorkRequestSummaryOperationTypeChangeUserAssessmentCompartment,
	"SET_USER_ASSESSMENT_BASELINE":              WorkRequestSummaryOperationTypeSetUserAssessmentBaseline,
	"UNSET_USER_ASSESSMENT_BASELINE":            WorkRequestSummaryOperationTypeUnsetUserAssessmentBaseline,
	"GENERATE_USER_ASSESSMENT_REPORT":           WorkRequestSummaryOperationTypeGenerateUserAssessmentReport,
	"CREATE_SECURITY_ASSESSMENT":                WorkRequestSummaryOperationTypeCreateSecurityAssessment,
	"CREATE_SECURITY_ASSESSMENT_NOW":            WorkRequestSummaryOperationTypeCreateSecurityAssessmentNow,
	"ASSESS_SECURITY_ASSESSMENT":                WorkRequestSummaryOperationTypeAssessSecurityAssessment,
	"CREATE_SNAPSHOT_SECURITY_ASSESSMENT":       WorkRequestSummaryOperationTypeCreateSnapshotSecurityAssessment,
	"CREATE_SCHEDULE_SECURITY_ASSESSMENT":       WorkRequestSummaryOperationTypeCreateScheduleSecurityAssessment,
	"COMPARE_WITH_BASELINE_SECURITY_ASSESSMENT": WorkRequestSummaryOperationTypeCompareWithBaselineSecurityAssessment,
	"DELETE_SECURITY_ASSESSMENT":                WorkRequestSummaryOperationTypeDeleteSecurityAssessment,
	"UPDATE_SECURITY_ASSESSMENT":                WorkRequestSummaryOperationTypeUpdateSecurityAssessment,
	"CHANGE_SECURITY_ASSESSMENT_COMPARTMENT":    WorkRequestSummaryOperationTypeChangeSecurityAssessmentCompartment,
	"SET_SECURITY_ASSESSMENT_BASELINE":          WorkRequestSummaryOperationTypeSetSecurityAssessmentBaseline,
	"UNSET_SECURITY_ASSESSMENT_BASELINE":        WorkRequestSummaryOperationTypeUnsetSecurityAssessmentBaseline,
	"GENERATE_SECURITY_ASSESSMENT_REPORT":       WorkRequestSummaryOperationTypeGenerateSecurityAssessmentReport,
}

// GetWorkRequestSummaryOperationTypeEnumValues Enumerates the set of values for WorkRequestSummaryOperationTypeEnum
func GetWorkRequestSummaryOperationTypeEnumValues() []WorkRequestSummaryOperationTypeEnum {
	values := make([]WorkRequestSummaryOperationTypeEnum, 0)
	for _, v := range mappingWorkRequestSummaryOperationType {
		values = append(values, v)
	}
	return values
}

// WorkRequestSummaryStatusEnum Enum with underlying type: string
type WorkRequestSummaryStatusEnum string

// Set of constants representing the allowable values for WorkRequestSummaryStatusEnum
const (
	WorkRequestSummaryStatusAccepted   WorkRequestSummaryStatusEnum = "ACCEPTED"
	WorkRequestSummaryStatusInProgress WorkRequestSummaryStatusEnum = "IN_PROGRESS"
	WorkRequestSummaryStatusFailed     WorkRequestSummaryStatusEnum = "FAILED"
	WorkRequestSummaryStatusSucceeded  WorkRequestSummaryStatusEnum = "SUCCEEDED"
)

var mappingWorkRequestSummaryStatus = map[string]WorkRequestSummaryStatusEnum{
	"ACCEPTED":    WorkRequestSummaryStatusAccepted,
	"IN_PROGRESS": WorkRequestSummaryStatusInProgress,
	"FAILED":      WorkRequestSummaryStatusFailed,
	"SUCCEEDED":   WorkRequestSummaryStatusSucceeded,
}

// GetWorkRequestSummaryStatusEnumValues Enumerates the set of values for WorkRequestSummaryStatusEnum
func GetWorkRequestSummaryStatusEnumValues() []WorkRequestSummaryStatusEnum {
	values := make([]WorkRequestSummaryStatusEnum, 0)
	for _, v := range mappingWorkRequestSummaryStatus {
		values = append(values, v)
	}
	return values
}
