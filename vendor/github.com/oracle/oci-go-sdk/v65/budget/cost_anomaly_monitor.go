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

// CostAnomalyMonitor A CostAnomalyMonitor.
type CostAnomalyMonitor struct {

	// The OCID of the Cost Anomaly Monitor.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The name of the cost anomaly monitor. Avoid entering confidential information.
	Name *string `mandatory:"true" json:"name"`

	// The current state of the cost monitor.
	LifecycleState CostAnomalyMonitorLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	TargetResourceFilter *TargetResourceFilter `mandatory:"true" json:"targetResourceFilter"`

	// The time that the cost monitor was created.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time that the cost monitor was last updated.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// Type of cost monitor
	Type MonitorTypeEnum `mandatory:"false" json:"type,omitempty"`

	// The description of the budget.
	Description *string `mandatory:"false" json:"description"`

	// The current state details of the cost monitor.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	CostAlertSubscriptionMap *CostAlertSubscriptionMap `mandatory:"false" json:"costAlertSubscriptionMap"`

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

func (m CostAnomalyMonitor) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CostAnomalyMonitor) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCostAnomalyMonitorLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetCostAnomalyMonitorLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingMonitorTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetMonitorTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CostAnomalyMonitorLifecycleStateEnum Enum with underlying type: string
type CostAnomalyMonitorLifecycleStateEnum string

// Set of constants representing the allowable values for CostAnomalyMonitorLifecycleStateEnum
const (
	CostAnomalyMonitorLifecycleStateActive   CostAnomalyMonitorLifecycleStateEnum = "ACTIVE"
	CostAnomalyMonitorLifecycleStateInactive CostAnomalyMonitorLifecycleStateEnum = "INACTIVE"
	CostAnomalyMonitorLifecycleStateDeleted  CostAnomalyMonitorLifecycleStateEnum = "DELETED"
)

var mappingCostAnomalyMonitorLifecycleStateEnum = map[string]CostAnomalyMonitorLifecycleStateEnum{
	"ACTIVE":   CostAnomalyMonitorLifecycleStateActive,
	"INACTIVE": CostAnomalyMonitorLifecycleStateInactive,
	"DELETED":  CostAnomalyMonitorLifecycleStateDeleted,
}

var mappingCostAnomalyMonitorLifecycleStateEnumLowerCase = map[string]CostAnomalyMonitorLifecycleStateEnum{
	"active":   CostAnomalyMonitorLifecycleStateActive,
	"inactive": CostAnomalyMonitorLifecycleStateInactive,
	"deleted":  CostAnomalyMonitorLifecycleStateDeleted,
}

// GetCostAnomalyMonitorLifecycleStateEnumValues Enumerates the set of values for CostAnomalyMonitorLifecycleStateEnum
func GetCostAnomalyMonitorLifecycleStateEnumValues() []CostAnomalyMonitorLifecycleStateEnum {
	values := make([]CostAnomalyMonitorLifecycleStateEnum, 0)
	for _, v := range mappingCostAnomalyMonitorLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetCostAnomalyMonitorLifecycleStateEnumStringValues Enumerates the set of values in String for CostAnomalyMonitorLifecycleStateEnum
func GetCostAnomalyMonitorLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"INACTIVE",
		"DELETED",
	}
}

// GetMappingCostAnomalyMonitorLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCostAnomalyMonitorLifecycleStateEnum(val string) (CostAnomalyMonitorLifecycleStateEnum, bool) {
	enum, ok := mappingCostAnomalyMonitorLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
