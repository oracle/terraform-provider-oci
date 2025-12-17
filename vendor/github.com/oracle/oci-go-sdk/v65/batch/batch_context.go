// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Batch API
//
// Use the Batch Control Plane API to encapsulate and manage all aspects of computationally intensive jobs.
//

package batch

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// BatchContext Representation of a batch context and its configurations.
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized, talk to
// an administrator. If you're an administrator who needs to write policies to give users access, see
// Getting Started with Policies (https://docs.oracle.com/iaas/Content/Identity/policiesgs/get-started-with-policies.htm).
type BatchContext struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the batch context.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The date and time the batch context was created, in the format defined by RFC 3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The current state of the batch context.
	LifecycleState BatchContextLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"true" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"true" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"true" json:"systemTags"`

	// List of job priority configurations related to the batch context.
	JobPriorityConfigurations []JobPriorityConfiguration `mandatory:"true" json:"jobPriorityConfigurations"`

	Network *Network `mandatory:"true" json:"network"`

	// List of fleet configurations related to the batch context.
	Fleets []Fleet `mandatory:"true" json:"fleets"`

	// Mapping of concurrent/shared resources used in job tasks to their limits.
	Entitlements map[string]int `mandatory:"true" json:"entitlements"`

	// Summarized information about the batch context.
	Description *string `mandatory:"false" json:"description"`

	// The date and time the batch context was updated, in the format defined by RFC 3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// A message that describes the current state in more detail. For example,   can be used to provide actionable information for a resource in the Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	LoggingConfiguration LoggingConfiguration `mandatory:"false" json:"loggingConfiguration"`
}

func (m BatchContext) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m BatchContext) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingBatchContextLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetBatchContextLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *BatchContext) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Description               *string                           `json:"description"`
		TimeUpdated               *common.SDKTime                   `json:"timeUpdated"`
		LifecycleDetails          *string                           `json:"lifecycleDetails"`
		LoggingConfiguration      loggingconfiguration              `json:"loggingConfiguration"`
		Id                        *string                           `json:"id"`
		CompartmentId             *string                           `json:"compartmentId"`
		DisplayName               *string                           `json:"displayName"`
		TimeCreated               *common.SDKTime                   `json:"timeCreated"`
		LifecycleState            BatchContextLifecycleStateEnum    `json:"lifecycleState"`
		FreeformTags              map[string]string                 `json:"freeformTags"`
		DefinedTags               map[string]map[string]interface{} `json:"definedTags"`
		SystemTags                map[string]map[string]interface{} `json:"systemTags"`
		JobPriorityConfigurations []JobPriorityConfiguration        `json:"jobPriorityConfigurations"`
		Network                   *Network                          `json:"network"`
		Fleets                    []fleet                           `json:"fleets"`
		Entitlements              map[string]int                    `json:"entitlements"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Description = model.Description

	m.TimeUpdated = model.TimeUpdated

	m.LifecycleDetails = model.LifecycleDetails

	nn, e = model.LoggingConfiguration.UnmarshalPolymorphicJSON(model.LoggingConfiguration.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.LoggingConfiguration = nn.(LoggingConfiguration)
	} else {
		m.LoggingConfiguration = nil
	}

	m.Id = model.Id

	m.CompartmentId = model.CompartmentId

	m.DisplayName = model.DisplayName

	m.TimeCreated = model.TimeCreated

	m.LifecycleState = model.LifecycleState

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.SystemTags = model.SystemTags

	m.JobPriorityConfigurations = make([]JobPriorityConfiguration, len(model.JobPriorityConfigurations))
	copy(m.JobPriorityConfigurations, model.JobPriorityConfigurations)
	m.Network = model.Network

	m.Fleets = make([]Fleet, len(model.Fleets))
	for i, n := range model.Fleets {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.Fleets[i] = nn.(Fleet)
		} else {
			m.Fleets[i] = nil
		}
	}
	m.Entitlements = model.Entitlements

	return
}

// BatchContextLifecycleStateEnum Enum with underlying type: string
type BatchContextLifecycleStateEnum string

// Set of constants representing the allowable values for BatchContextLifecycleStateEnum
const (
	BatchContextLifecycleStateCreating       BatchContextLifecycleStateEnum = "CREATING"
	BatchContextLifecycleStateActive         BatchContextLifecycleStateEnum = "ACTIVE"
	BatchContextLifecycleStateInactive       BatchContextLifecycleStateEnum = "INACTIVE"
	BatchContextLifecycleStateUpdating       BatchContextLifecycleStateEnum = "UPDATING"
	BatchContextLifecycleStateNeedsAttention BatchContextLifecycleStateEnum = "NEEDS_ATTENTION"
	BatchContextLifecycleStateDeleting       BatchContextLifecycleStateEnum = "DELETING"
	BatchContextLifecycleStateDeleted        BatchContextLifecycleStateEnum = "DELETED"
	BatchContextLifecycleStateFailed         BatchContextLifecycleStateEnum = "FAILED"
)

var mappingBatchContextLifecycleStateEnum = map[string]BatchContextLifecycleStateEnum{
	"CREATING":        BatchContextLifecycleStateCreating,
	"ACTIVE":          BatchContextLifecycleStateActive,
	"INACTIVE":        BatchContextLifecycleStateInactive,
	"UPDATING":        BatchContextLifecycleStateUpdating,
	"NEEDS_ATTENTION": BatchContextLifecycleStateNeedsAttention,
	"DELETING":        BatchContextLifecycleStateDeleting,
	"DELETED":         BatchContextLifecycleStateDeleted,
	"FAILED":          BatchContextLifecycleStateFailed,
}

var mappingBatchContextLifecycleStateEnumLowerCase = map[string]BatchContextLifecycleStateEnum{
	"creating":        BatchContextLifecycleStateCreating,
	"active":          BatchContextLifecycleStateActive,
	"inactive":        BatchContextLifecycleStateInactive,
	"updating":        BatchContextLifecycleStateUpdating,
	"needs_attention": BatchContextLifecycleStateNeedsAttention,
	"deleting":        BatchContextLifecycleStateDeleting,
	"deleted":         BatchContextLifecycleStateDeleted,
	"failed":          BatchContextLifecycleStateFailed,
}

// GetBatchContextLifecycleStateEnumValues Enumerates the set of values for BatchContextLifecycleStateEnum
func GetBatchContextLifecycleStateEnumValues() []BatchContextLifecycleStateEnum {
	values := make([]BatchContextLifecycleStateEnum, 0)
	for _, v := range mappingBatchContextLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetBatchContextLifecycleStateEnumStringValues Enumerates the set of values in String for BatchContextLifecycleStateEnum
func GetBatchContextLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"INACTIVE",
		"UPDATING",
		"NEEDS_ATTENTION",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingBatchContextLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBatchContextLifecycleStateEnum(val string) (BatchContextLifecycleStateEnum, bool) {
	enum, ok := mappingBatchContextLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
