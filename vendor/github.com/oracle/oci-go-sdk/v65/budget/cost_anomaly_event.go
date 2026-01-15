// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Budgets API
//
// Use the Budgets API to manage budgets and budget alerts. For more information, see Budgets Overview (https://docs.oracle.com/iaas/Content/Billing/Concepts/budgetsoverview.htm).
//

package budget

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CostAnomalyEvent A CostAnomalyEvent.
type CostAnomalyEvent struct {

	// The OCID of the Cost Anomaly Event.
	Id *string `mandatory:"true" json:"id"`

	// The name of the associated cost monitor.
	CostAnomalyName *string `mandatory:"true" json:"costAnomalyName"`

	// The OCID of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The current state of the cost anomaly event.
	LifecycleState CostAnomalyEventLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The OCID of the associated cost monitor.
	CostMonitorId *string `mandatory:"true" json:"costMonitorId"`

	// The name of the associated cost monitor.
	CostMonitorName *string `mandatory:"true" json:"costMonitorName"`

	TargetResourceFilter *TargetResourceFilter `mandatory:"true" json:"targetResourceFilter"`

	// The event date of the anomaly event.
	TimeAnomalyEventDate *common.SDKTime `mandatory:"true" json:"timeAnomalyEventDate"`

	// The created time of the cost anomaly event.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The updated time of the cost anomaly event.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// Type of cost monitor
	CostMonitorType MonitorTypeEnum `mandatory:"false" json:"costMonitorType,omitempty"`

	// The cost impact of the detected anomaly.
	CostImpact *float64 `mandatory:"false" json:"costImpact"`

	// The cost variance percentage of the detected anomaly.
	CostVariancePercentage *float64 `mandatory:"false" json:"costVariancePercentage"`

	RootCauseDetail *RootCauseDetail `mandatory:"false" json:"rootCauseDetail"`

	// The feedback response for the cost anomaly event.
	FeedbackResponse CostAnomalyEventFeedbackResponseEnum `mandatory:"false" json:"feedbackResponse,omitempty"`

	// Defined tags for this resource. Each key is predefined and scoped to a
	// namespace. For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// System tags can be viewed by users, but can only be created by the system.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m CostAnomalyEvent) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CostAnomalyEvent) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCostAnomalyEventLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetCostAnomalyEventLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingMonitorTypeEnum(string(m.CostMonitorType)); !ok && m.CostMonitorType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CostMonitorType: %s. Supported values are: %s.", m.CostMonitorType, strings.Join(GetMonitorTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCostAnomalyEventFeedbackResponseEnum(string(m.FeedbackResponse)); !ok && m.FeedbackResponse != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for FeedbackResponse: %s. Supported values are: %s.", m.FeedbackResponse, strings.Join(GetCostAnomalyEventFeedbackResponseEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CostAnomalyEventLifecycleStateEnum Enum with underlying type: string
type CostAnomalyEventLifecycleStateEnum string

// Set of constants representing the allowable values for CostAnomalyEventLifecycleStateEnum
const (
	CostAnomalyEventLifecycleStateActive   CostAnomalyEventLifecycleStateEnum = "ACTIVE"
	CostAnomalyEventLifecycleStateInactive CostAnomalyEventLifecycleStateEnum = "INACTIVE"
)

var mappingCostAnomalyEventLifecycleStateEnum = map[string]CostAnomalyEventLifecycleStateEnum{
	"ACTIVE":   CostAnomalyEventLifecycleStateActive,
	"INACTIVE": CostAnomalyEventLifecycleStateInactive,
}

var mappingCostAnomalyEventLifecycleStateEnumLowerCase = map[string]CostAnomalyEventLifecycleStateEnum{
	"active":   CostAnomalyEventLifecycleStateActive,
	"inactive": CostAnomalyEventLifecycleStateInactive,
}

// GetCostAnomalyEventLifecycleStateEnumValues Enumerates the set of values for CostAnomalyEventLifecycleStateEnum
func GetCostAnomalyEventLifecycleStateEnumValues() []CostAnomalyEventLifecycleStateEnum {
	values := make([]CostAnomalyEventLifecycleStateEnum, 0)
	for _, v := range mappingCostAnomalyEventLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetCostAnomalyEventLifecycleStateEnumStringValues Enumerates the set of values in String for CostAnomalyEventLifecycleStateEnum
func GetCostAnomalyEventLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"INACTIVE",
	}
}

// GetMappingCostAnomalyEventLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCostAnomalyEventLifecycleStateEnum(val string) (CostAnomalyEventLifecycleStateEnum, bool) {
	enum, ok := mappingCostAnomalyEventLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// CostAnomalyEventFeedbackResponseEnum Enum with underlying type: string
type CostAnomalyEventFeedbackResponseEnum string

// Set of constants representing the allowable values for CostAnomalyEventFeedbackResponseEnum
const (
	CostAnomalyEventFeedbackResponseAccurateAnomaly CostAnomalyEventFeedbackResponseEnum = "ACCURATE_ANOMALY"
	CostAnomalyEventFeedbackResponseExpectedAnomaly CostAnomalyEventFeedbackResponseEnum = "EXPECTED_ANOMALY"
)

var mappingCostAnomalyEventFeedbackResponseEnum = map[string]CostAnomalyEventFeedbackResponseEnum{
	"ACCURATE_ANOMALY": CostAnomalyEventFeedbackResponseAccurateAnomaly,
	"EXPECTED_ANOMALY": CostAnomalyEventFeedbackResponseExpectedAnomaly,
}

var mappingCostAnomalyEventFeedbackResponseEnumLowerCase = map[string]CostAnomalyEventFeedbackResponseEnum{
	"accurate_anomaly": CostAnomalyEventFeedbackResponseAccurateAnomaly,
	"expected_anomaly": CostAnomalyEventFeedbackResponseExpectedAnomaly,
}

// GetCostAnomalyEventFeedbackResponseEnumValues Enumerates the set of values for CostAnomalyEventFeedbackResponseEnum
func GetCostAnomalyEventFeedbackResponseEnumValues() []CostAnomalyEventFeedbackResponseEnum {
	values := make([]CostAnomalyEventFeedbackResponseEnum, 0)
	for _, v := range mappingCostAnomalyEventFeedbackResponseEnum {
		values = append(values, v)
	}
	return values
}

// GetCostAnomalyEventFeedbackResponseEnumStringValues Enumerates the set of values in String for CostAnomalyEventFeedbackResponseEnum
func GetCostAnomalyEventFeedbackResponseEnumStringValues() []string {
	return []string{
		"ACCURATE_ANOMALY",
		"EXPECTED_ANOMALY",
	}
}

// GetMappingCostAnomalyEventFeedbackResponseEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCostAnomalyEventFeedbackResponseEnum(val string) (CostAnomalyEventFeedbackResponseEnum, bool) {
	enum, ok := mappingCostAnomalyEventFeedbackResponseEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
