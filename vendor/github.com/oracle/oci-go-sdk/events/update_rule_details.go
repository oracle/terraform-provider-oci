// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Events API
//
// API for the Events Service. Use this API to manage rules and actions that create automation
// in your tenancy. For more information, see Overview of Events (https://docs.cloud.oracle.com/iaas/Content/Events/Concepts/eventsoverview.htm).
//

package events

import (
	"github.com/oracle/oci-go-sdk/common"
)

// UpdateRuleDetails The rule attributes that you can update.
type UpdateRuleDetails struct {

	// A string that describes the rule. It does not have to be unique, and you can change it. Avoid entering
	// confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// A string that describes the details of the rule. It does not have to be unique, and you can change it. Avoid entering
	// confidential information.
	Description *string `mandatory:"false" json:"description"`

	// Whether or not this rule is currently enabled.
	// Example: `true`
	IsEnabled *bool `mandatory:"false" json:"isEnabled"`

	// A filter that specifies the event that will trigger actions associated with this rule. A few
	// important things to remember about filters:
	// * Fields not mentioned in the condition are ignored. You can create a valid filter that matches
	// all events with two curly brackets: `{}`
	//   For more examples, see
	// Matching Events with Filters (https://docs.cloud.oracle.com/iaas/Content/Events/Concepts/filterevents.htm).
	// * For a condition with fileds to match an event, the event must contain all the field names
	// listed in the condition. Field names must appear in the condition with the same nesting
	// structure used in the event.
	//   For a list of reference events, see
	// Services that Produce Events (https://docs.cloud.oracle.com/iaas/Content/Events/Reference/eventsproducers.htm).
	// * Rules apply to events in the compartment in which you create them and any child compartments.
	// This means that a condition specified by a rule only matches events emitted from resources in
	// the compartment or any of its child compartments.
	// * The condition is a string value in a JSON object, but numbers in conditions are converted
	// from strings to numbers before they are evaluated for matches. This means that 100, 100.0 or
	// 1.0e2 are all considered equal.
	// * Boolean values are converted to numbers and then evaluated. This means true and True are
	// considered equal, as are False and false.
	// * Wildcard matching is supported with the asterisk (*) character.
	//   For examples of wildcard matching, see
	// Matching Events with Filters (https://docs.cloud.oracle.com/iaas/Content/Events/Concepts/filterevents.htm)
	// Example: `\"eventType\": \"com.oraclecloud.databaseservice.autonomous.database.backup.end\"`
	Condition *string `mandatory:"false" json:"condition"`

	Actions *ActionDetailsList `mandatory:"false" json:"actions"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. Exists for cross-compatibility only.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m UpdateRuleDetails) String() string {
	return common.PointerString(m)
}
