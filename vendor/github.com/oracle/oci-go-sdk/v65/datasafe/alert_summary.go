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

// AlertSummary Summary of a Data Safe Alert.
type AlertSummary struct {

	// The OCID of the alert.
	Id *string `mandatory:"true" json:"id"`

	// The status of the alert.
	Status AlertStatusEnum `mandatory:"true" json:"status"`

	// The display name of the alert.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Severity level of the alert.
	Severity AlertSeverityEnum `mandatory:"true" json:"severity"`

	// The OCID of the compartment that contains the alert.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Creation date and time of the alert, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// Last date and time the alert was updated, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The current state of the alert.
	LifecycleState AlertLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The details of the alert.
	Description *string `mandatory:"false" json:"description"`

	// Creation date and time of the operation that triggered alert, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	OperationTime *common.SDKTime `mandatory:"false" json:"operationTime"`

	// The operation that triggered alert.
	Operation *string `mandatory:"false" json:"operation"`

	// The result of the operation (event) that triggered alert.
	OperationStatus AlertSummaryOperationStatusEnum `mandatory:"false" json:"operationStatus,omitempty"`

	// Array of OCIDs of the target database.
	TargetIds []string `mandatory:"false" json:"targetIds"`

	// Array of names of the target database.
	TargetNames []string `mandatory:"false" json:"targetNames"`

	// The OCID of the policy that triggered alert.
	PolicyId *string `mandatory:"false" json:"policyId"`

	// Type of the alert. Indicates the Data Safe feature triggering the alert.
	AlertType AlertTypeEnum `mandatory:"false" json:"alertType,omitempty"`

	// Map that contains maps of values.
	//  Example: `{"Operations": {"CostCenter": "42"}}`
	FeatureDetails map[string]map[string]interface{} `mandatory:"false" json:"featureDetails"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m AlertSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AlertSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAlertStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetAlertStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAlertSeverityEnum(string(m.Severity)); !ok && m.Severity != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Severity: %s. Supported values are: %s.", m.Severity, strings.Join(GetAlertSeverityEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAlertLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetAlertLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingAlertSummaryOperationStatusEnum(string(m.OperationStatus)); !ok && m.OperationStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OperationStatus: %s. Supported values are: %s.", m.OperationStatus, strings.Join(GetAlertSummaryOperationStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAlertTypeEnum(string(m.AlertType)); !ok && m.AlertType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AlertType: %s. Supported values are: %s.", m.AlertType, strings.Join(GetAlertTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AlertSummaryOperationStatusEnum Enum with underlying type: string
type AlertSummaryOperationStatusEnum string

// Set of constants representing the allowable values for AlertSummaryOperationStatusEnum
const (
	AlertSummaryOperationStatusSucceeded AlertSummaryOperationStatusEnum = "SUCCEEDED"
	AlertSummaryOperationStatusFailed    AlertSummaryOperationStatusEnum = "FAILED"
)

var mappingAlertSummaryOperationStatusEnum = map[string]AlertSummaryOperationStatusEnum{
	"SUCCEEDED": AlertSummaryOperationStatusSucceeded,
	"FAILED":    AlertSummaryOperationStatusFailed,
}

var mappingAlertSummaryOperationStatusEnumLowerCase = map[string]AlertSummaryOperationStatusEnum{
	"succeeded": AlertSummaryOperationStatusSucceeded,
	"failed":    AlertSummaryOperationStatusFailed,
}

// GetAlertSummaryOperationStatusEnumValues Enumerates the set of values for AlertSummaryOperationStatusEnum
func GetAlertSummaryOperationStatusEnumValues() []AlertSummaryOperationStatusEnum {
	values := make([]AlertSummaryOperationStatusEnum, 0)
	for _, v := range mappingAlertSummaryOperationStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetAlertSummaryOperationStatusEnumStringValues Enumerates the set of values in String for AlertSummaryOperationStatusEnum
func GetAlertSummaryOperationStatusEnumStringValues() []string {
	return []string{
		"SUCCEEDED",
		"FAILED",
	}
}

// GetMappingAlertSummaryOperationStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAlertSummaryOperationStatusEnum(val string) (AlertSummaryOperationStatusEnum, bool) {
	enum, ok := mappingAlertSummaryOperationStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
