// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// MaskingPolicy A masking policy defines the approach to mask data in a target database. It's basically
// a collection of columns to be masked, called masking columns, and the associated masking
// formats to be used to mask these columns. A masking policy can be used to mask multiple
// databases provided that they have the same schema design. For more information, see
// <a href="https://docs.oracle.com/en/cloud/paas/data-safe/udscs/masking-policies.html">Masking Policies </a>
// in the Oracle Data Safe documentation.
type MaskingPolicy struct {

	// The OCID of the masking policy.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the compartment that contains the masking policy.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The display name of the masking policy.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The date and time the masking policy was created, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The current state of the masking policy.
	LifecycleState MaskingLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time the masking policy was last updated, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339)
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// Indicates if the temporary tables created during a masking operation should be dropped after masking. It's enabled by default.
	// Set this attribute to false to preserve the temporary tables. Masking creates temporary tables that map the original sensitive
	// data values to mask values. By default, these temporary tables are dropped after masking. But, in some cases, you may want
	// to preserve this information to track how masking changed your data. Note that doing so compromises security. These tables
	// must be dropped before the database is available for unprivileged users.
	IsDropTempTablesEnabled *bool `mandatory:"true" json:"isDropTempTablesEnabled"`

	// Indicates if redo logging is enabled during a masking operation. It's disabled by default. Set this attribute to true to
	// enable redo logging. By default, masking disables redo logging and flashback logging to purge any original unmasked
	// data from logs. However, in certain circumstances when you only want to test masking, rollback changes, and retry masking,
	// you could enable logging and use a flashback database to retrieve the original unmasked data after it has been masked.
	IsRedoLoggingEnabled *bool `mandatory:"true" json:"isRedoLoggingEnabled"`

	// Indicates if statistics gathering is enabled. It's enabled by default. Set this attribute to false to disable statistics
	// gathering. The masking process gathers statistics on masked database tables after masking completes.
	IsRefreshStatsEnabled *bool `mandatory:"true" json:"isRefreshStatsEnabled"`

	// Specifies options to enable parallel execution when running data masking. Allowed values are 'NONE' (no parallelism),
	// 'DEFAULT' (the Oracle Database computes the optimum degree of parallelism) or an integer value to be used as the degree
	// of parallelism. Parallel execution helps effectively use multiple CPUs and improve masking performance. Refer to the
	// Oracle Database parallel execution framework when choosing an explicit degree of parallelism.
	ParallelDegree *string `mandatory:"true" json:"parallelDegree"`

	// Specifies how to recompile invalid objects post data masking. Allowed values are 'SERIAL' (recompile in serial),
	// 'PARALLEL' (recompile in parallel), 'NONE' (do not recompile). If it's set to PARALLEL, the value of parallelDegree
	// attribute is used. Use the built-in UTL_RECOMP package to recompile any remaining invalid objects after masking completes.
	Recompile MaskingPolicyRecompileEnum `mandatory:"true" json:"recompile"`

	// The description of the masking policy.
	Description *string `mandatory:"false" json:"description"`

	// A pre-masking script, which can contain SQL and PL/SQL statements. It's executed before
	// the core masking script generated using the masking policy. It's usually used to perform
	// any preparation or prerequisite work before masking data.
	PreMaskingScript *string `mandatory:"false" json:"preMaskingScript"`

	// A post-masking script, which can contain SQL and PL/SQL statements. It's executed after
	// the core masking script generated using the masking policy. It's usually used to perform
	// additional transformation or cleanup work after masking.
	PostMaskingScript *string `mandatory:"false" json:"postMaskingScript"`

	ColumnSource ColumnSourceDetails `mandatory:"false" json:"columnSource"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m MaskingPolicy) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MaskingPolicy) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingMaskingLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetMaskingLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingMaskingPolicyRecompileEnum(string(m.Recompile)); !ok && m.Recompile != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Recompile: %s. Supported values are: %s.", m.Recompile, strings.Join(GetMaskingPolicyRecompileEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *MaskingPolicy) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Description             *string                           `json:"description"`
		PreMaskingScript        *string                           `json:"preMaskingScript"`
		PostMaskingScript       *string                           `json:"postMaskingScript"`
		ColumnSource            columnsourcedetails               `json:"columnSource"`
		FreeformTags            map[string]string                 `json:"freeformTags"`
		DefinedTags             map[string]map[string]interface{} `json:"definedTags"`
		Id                      *string                           `json:"id"`
		CompartmentId           *string                           `json:"compartmentId"`
		DisplayName             *string                           `json:"displayName"`
		TimeCreated             *common.SDKTime                   `json:"timeCreated"`
		LifecycleState          MaskingLifecycleStateEnum         `json:"lifecycleState"`
		TimeUpdated             *common.SDKTime                   `json:"timeUpdated"`
		IsDropTempTablesEnabled *bool                             `json:"isDropTempTablesEnabled"`
		IsRedoLoggingEnabled    *bool                             `json:"isRedoLoggingEnabled"`
		IsRefreshStatsEnabled   *bool                             `json:"isRefreshStatsEnabled"`
		ParallelDegree          *string                           `json:"parallelDegree"`
		Recompile               MaskingPolicyRecompileEnum        `json:"recompile"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Description = model.Description

	m.PreMaskingScript = model.PreMaskingScript

	m.PostMaskingScript = model.PostMaskingScript

	nn, e = model.ColumnSource.UnmarshalPolymorphicJSON(model.ColumnSource.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.ColumnSource = nn.(ColumnSourceDetails)
	} else {
		m.ColumnSource = nil
	}

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.Id = model.Id

	m.CompartmentId = model.CompartmentId

	m.DisplayName = model.DisplayName

	m.TimeCreated = model.TimeCreated

	m.LifecycleState = model.LifecycleState

	m.TimeUpdated = model.TimeUpdated

	m.IsDropTempTablesEnabled = model.IsDropTempTablesEnabled

	m.IsRedoLoggingEnabled = model.IsRedoLoggingEnabled

	m.IsRefreshStatsEnabled = model.IsRefreshStatsEnabled

	m.ParallelDegree = model.ParallelDegree

	m.Recompile = model.Recompile

	return
}

// MaskingPolicyRecompileEnum Enum with underlying type: string
type MaskingPolicyRecompileEnum string

// Set of constants representing the allowable values for MaskingPolicyRecompileEnum
const (
	MaskingPolicyRecompileSerial   MaskingPolicyRecompileEnum = "SERIAL"
	MaskingPolicyRecompileParallel MaskingPolicyRecompileEnum = "PARALLEL"
	MaskingPolicyRecompileNone     MaskingPolicyRecompileEnum = "NONE"
)

var mappingMaskingPolicyRecompileEnum = map[string]MaskingPolicyRecompileEnum{
	"SERIAL":   MaskingPolicyRecompileSerial,
	"PARALLEL": MaskingPolicyRecompileParallel,
	"NONE":     MaskingPolicyRecompileNone,
}

var mappingMaskingPolicyRecompileEnumLowerCase = map[string]MaskingPolicyRecompileEnum{
	"serial":   MaskingPolicyRecompileSerial,
	"parallel": MaskingPolicyRecompileParallel,
	"none":     MaskingPolicyRecompileNone,
}

// GetMaskingPolicyRecompileEnumValues Enumerates the set of values for MaskingPolicyRecompileEnum
func GetMaskingPolicyRecompileEnumValues() []MaskingPolicyRecompileEnum {
	values := make([]MaskingPolicyRecompileEnum, 0)
	for _, v := range mappingMaskingPolicyRecompileEnum {
		values = append(values, v)
	}
	return values
}

// GetMaskingPolicyRecompileEnumStringValues Enumerates the set of values in String for MaskingPolicyRecompileEnum
func GetMaskingPolicyRecompileEnumStringValues() []string {
	return []string{
		"SERIAL",
		"PARALLEL",
		"NONE",
	}
}

// GetMappingMaskingPolicyRecompileEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMaskingPolicyRecompileEnum(val string) (MaskingPolicyRecompileEnum, bool) {
	enum, ok := mappingMaskingPolicyRecompileEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
