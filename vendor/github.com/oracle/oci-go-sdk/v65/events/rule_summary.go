// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Events API
//
// API for the Events Service. Use this API to manage rules and actions that create automation
// in your tenancy. For more information, see Overview of Events (https://docs.oracle.com/iaas/Content/Events/Concepts/eventsoverview.htm).
//

package events

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// RuleSummary The summary details of rules for Events. For more information, see
// Managing Rules for Events (https://docs.oracle.com/iaas/Content/Events/Task/managingrules.htm).
type RuleSummary struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of this rule.
	Id *string `mandatory:"true" json:"id"`

	// A string that describes the rule. It does not have to be unique, and you can change it. Avoid entering
	// confidential information.
	// Example: `"This rule sends a notification upon completion of DbaaS backup."`
	DisplayName *string `mandatory:"true" json:"displayName"`

	LifecycleState RuleLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// A filter that specifies the event that will trigger actions associated with this rule. A few
	// important things to remember about filters:
	// * Fields not mentioned in the condition are ignored. You can create a valid filter that matches
	// all events with two curly brackets: `{}`
	//   For more examples, see
	// Matching Events with Filters (https://docs.oracle.com/iaas/Content/Events/Concepts/filterevents.htm).
	// * For a condition with fields to match an event, the event must contain all the field names
	// listed in the condition. Field names must appear in the condition with the same nesting
	// structure used in the event.
	//   For a list of reference events, see
	// Services that Produce Events (https://docs.oracle.com/iaas/Content/Events/Reference/eventsproducers.htm).
	// * Rules apply to events in the compartment in which you create them and any child compartments.
	// This means that a condition specified by a rule only matches events emitted from resources in
	// the compartment or any of its child compartments.
	// * Wildcard matching is supported with the asterisk (*) character.
	//   For examples of wildcard matching, see
	// Matching Events with Filters (https://docs.oracle.com/iaas/Content/Events/Concepts/filterevents.htm)
	// Example: `\"eventType\": \"com.oraclecloud.databaseservice.autonomous.database.backup.end\"`
	Condition *string `mandatory:"true" json:"condition"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment to which this rule belongs.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Whether or not this rule is currently enabled.
	// Example: `true`
	IsEnabled *bool `mandatory:"true" json:"isEnabled"`

	// The time this rule was created, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339)
	// timestamp format.
	// Example: `2018-09-12T22:47:12.613Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// A string that describes the details of the rule. It does not have to be unique, and you can change it. Avoid entering
	// confidential information.
	Description *string `mandatory:"false" json:"description"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. Exists for cross-compatibility only.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m RuleSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RuleSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingRuleLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetRuleLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
