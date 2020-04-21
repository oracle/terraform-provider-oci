// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
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
	// * For a condition with fields to match an event, the event must contain all the field names
	// listed in the condition. Field names must appear in the condition with the same nesting
	// structure used in the event.
	//   For a list of reference events, see
	// Services that Produce Events (https://docs.cloud.oracle.com/iaas/Content/Events/Reference/eventsproducers.htm).
	// * Rules apply to events in the compartment in which you create them and any child compartments.
	// This means that a condition specified by a rule only matches events emitted from resources in
	// the compartment or any of its child compartments.
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
