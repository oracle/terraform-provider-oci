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

// CostAlertSubscription A CostAlertSubscription.
type CostAlertSubscription struct {

	// The OCID of the Cost Alert Subscription.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the compartment which hold the cost alert subscription resource.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The name of the cost alert subscription. Avoid entering confidential information.
	Name *string `mandatory:"true" json:"name"`

	// The current state of the cost alert subscription.
	LifecycleState CostAlertSubscriptionLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The notification channels string.
	Channels *string `mandatory:"true" json:"channels"`

	// The time that the cost alert subscription was created.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time that the cost alert subscription was updated.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// List of monitor identifiers
	CostAnomalyMonitors *interface{} `mandatory:"true" json:"costAnomalyMonitors"`

	// The description of the cost alert subscription.
	Description *string `mandatory:"false" json:"description"`

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

func (m CostAlertSubscription) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CostAlertSubscription) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCostAlertSubscriptionLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetCostAlertSubscriptionLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CostAlertSubscriptionLifecycleStateEnum Enum with underlying type: string
type CostAlertSubscriptionLifecycleStateEnum string

// Set of constants representing the allowable values for CostAlertSubscriptionLifecycleStateEnum
const (
	CostAlertSubscriptionLifecycleStateActive   CostAlertSubscriptionLifecycleStateEnum = "ACTIVE"
	CostAlertSubscriptionLifecycleStateInactive CostAlertSubscriptionLifecycleStateEnum = "INACTIVE"
)

var mappingCostAlertSubscriptionLifecycleStateEnum = map[string]CostAlertSubscriptionLifecycleStateEnum{
	"ACTIVE":   CostAlertSubscriptionLifecycleStateActive,
	"INACTIVE": CostAlertSubscriptionLifecycleStateInactive,
}

var mappingCostAlertSubscriptionLifecycleStateEnumLowerCase = map[string]CostAlertSubscriptionLifecycleStateEnum{
	"active":   CostAlertSubscriptionLifecycleStateActive,
	"inactive": CostAlertSubscriptionLifecycleStateInactive,
}

// GetCostAlertSubscriptionLifecycleStateEnumValues Enumerates the set of values for CostAlertSubscriptionLifecycleStateEnum
func GetCostAlertSubscriptionLifecycleStateEnumValues() []CostAlertSubscriptionLifecycleStateEnum {
	values := make([]CostAlertSubscriptionLifecycleStateEnum, 0)
	for _, v := range mappingCostAlertSubscriptionLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetCostAlertSubscriptionLifecycleStateEnumStringValues Enumerates the set of values in String for CostAlertSubscriptionLifecycleStateEnum
func GetCostAlertSubscriptionLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"INACTIVE",
	}
}

// GetMappingCostAlertSubscriptionLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCostAlertSubscriptionLifecycleStateEnum(val string) (CostAlertSubscriptionLifecycleStateEnum, bool) {
	enum, ok := mappingCostAlertSubscriptionLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
