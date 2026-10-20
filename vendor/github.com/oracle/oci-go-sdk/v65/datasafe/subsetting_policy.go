// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// SubsettingPolicy A subsetting policy defines the approach to subset data in a target database. It's basically
// a collection of rules applied on tables or schemas, to reduce rows in them. A subsetting policy can
// be used to subset multiple databases provided that they have the same schema design.
type SubsettingPolicy struct {

	// The OCID of the subsetting policy
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the compartment that contains the subsetting policy
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The date and time the subsetting policy was created, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339)
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The current state of the subsetting policy
	LifecycleState SubsettingPolicyLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time the subsetting policy was last updated, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339)
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// Indicates if redo logging is enabled during a subsetting operation. It's disabled by default. Set this attribute to true to
	// enable redo logging. By default, subsetting disables redo logging and flashback logging to purge any original
	// data from logs. However, in certain circumstances when you only want to test subsetting, rollback changes, and retry subsetting,
	// you could enable logging and use a flashback database to retrieve the original data after it has been subsetted.
	IsRedoLoggingEnabled *bool `mandatory:"true" json:"isRedoLoggingEnabled"`

	// Indicates if statistics gathering is enabled. It's enabled by default. Set this attribute to false to disable statistics
	// gathering. The subsetting process gathers statistics on database tables after subsetting completes
	IsRefreshStatsEnabled *bool `mandatory:"true" json:"isRefreshStatsEnabled"`

	// Specifies options to enable parallel execution when running data subsetting. Allowed values are 'NONE' (no parallelism),
	// 'DEFAULT' (the Oracle Database computes the optimum degree of parallelism) or an integer value to be used as the degree
	// of parallelism. Parallel execution helps effectively use multiple CPUs and improve subsetting performance. Refer to the
	// Oracle Database parallel execution framework when choosing an explicit degree of parallelism
	ParallelDegree *string `mandatory:"true" json:"parallelDegree"`

	// Specifies how to recompile invalid objects post data subsetting. Allowed values are 'SERIAL' (recompile in serial),
	// 'PARALLEL' (recompile in parallel), 'NONE' (do not recompile). If it's set to PARALLEL, the value of parallelDegree
	// attribute is used. Use the built-in UTL_RECOMP package to recompile any remaining invalid objects after subsetting completes
	Recompile SubsettingPolicyRecompileEnum `mandatory:"true" json:"recompile"`

	// The display name of the subsetting policy
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The description of the subsetting policy
	Description *string `mandatory:"false" json:"description"`

	// Strategy to be applied for tables which are not impacted by any of the subsetting rules
	UnrelatedTablesAction SubsettingPolicyUnrelatedTablesActionEnum `mandatory:"false" json:"unrelatedTablesAction,omitempty"`

	// A pre-subsetting script, which can contain SQL and PL/SQL statements. It's executed before
	// the core subsetting script generated using the subsetting policy. It's usually used to perform
	// any preparation or prerequisite work before subsetting data
	PreSubsettingScript *string `mandatory:"false" json:"preSubsettingScript"`

	// A post-subsetting script, which can contain SQL and PL/SQL statements. It's executed after
	// the core subsetting script generated using the subsetting policy. It's usually used to perform
	// additional transformation or cleanup work after subsetting.
	PostSubsettingScript *string `mandatory:"false" json:"postSubsettingScript"`

	SchemaSource SchemaSourceDetails `mandatory:"false" json:"schemaSource"`

	// The OCID of the masking policy associated with this subsetting policy
	MaskingPolicyId *string `mandatory:"false" json:"maskingPolicyId"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m SubsettingPolicy) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SubsettingPolicy) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSubsettingPolicyLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetSubsettingPolicyLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSubsettingPolicyRecompileEnum(string(m.Recompile)); !ok && m.Recompile != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Recompile: %s. Supported values are: %s.", m.Recompile, strings.Join(GetSubsettingPolicyRecompileEnumStringValues(), ",")))
	}

	if _, ok := GetMappingSubsettingPolicyUnrelatedTablesActionEnum(string(m.UnrelatedTablesAction)); !ok && m.UnrelatedTablesAction != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for UnrelatedTablesAction: %s. Supported values are: %s.", m.UnrelatedTablesAction, strings.Join(GetSubsettingPolicyUnrelatedTablesActionEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *SubsettingPolicy) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName           *string                                   `json:"displayName"`
		Description           *string                                   `json:"description"`
		UnrelatedTablesAction SubsettingPolicyUnrelatedTablesActionEnum `json:"unrelatedTablesAction"`
		PreSubsettingScript   *string                                   `json:"preSubsettingScript"`
		PostSubsettingScript  *string                                   `json:"postSubsettingScript"`
		SchemaSource          schemasourcedetails                       `json:"schemaSource"`
		MaskingPolicyId       *string                                   `json:"maskingPolicyId"`
		FreeformTags          map[string]string                         `json:"freeformTags"`
		DefinedTags           map[string]map[string]interface{}         `json:"definedTags"`
		Id                    *string                                   `json:"id"`
		CompartmentId         *string                                   `json:"compartmentId"`
		TimeCreated           *common.SDKTime                           `json:"timeCreated"`
		LifecycleState        SubsettingPolicyLifecycleStateEnum        `json:"lifecycleState"`
		TimeUpdated           *common.SDKTime                           `json:"timeUpdated"`
		IsRedoLoggingEnabled  *bool                                     `json:"isRedoLoggingEnabled"`
		IsRefreshStatsEnabled *bool                                     `json:"isRefreshStatsEnabled"`
		ParallelDegree        *string                                   `json:"parallelDegree"`
		Recompile             SubsettingPolicyRecompileEnum             `json:"recompile"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	m.Description = model.Description

	m.UnrelatedTablesAction = model.UnrelatedTablesAction

	m.PreSubsettingScript = model.PreSubsettingScript

	m.PostSubsettingScript = model.PostSubsettingScript

	nn, e = model.SchemaSource.UnmarshalPolymorphicJSON(model.SchemaSource.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.SchemaSource = nn.(SchemaSourceDetails)
	} else {
		m.SchemaSource = nil
	}

	m.MaskingPolicyId = model.MaskingPolicyId

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.Id = model.Id

	m.CompartmentId = model.CompartmentId

	m.TimeCreated = model.TimeCreated

	m.LifecycleState = model.LifecycleState

	m.TimeUpdated = model.TimeUpdated

	m.IsRedoLoggingEnabled = model.IsRedoLoggingEnabled

	m.IsRefreshStatsEnabled = model.IsRefreshStatsEnabled

	m.ParallelDegree = model.ParallelDegree

	m.Recompile = model.Recompile

	return
}

// SubsettingPolicyLifecycleStateEnum Enum with underlying type: string
type SubsettingPolicyLifecycleStateEnum string

// Set of constants representing the allowable values for SubsettingPolicyLifecycleStateEnum
const (
	SubsettingPolicyLifecycleStateCreating       SubsettingPolicyLifecycleStateEnum = "CREATING"
	SubsettingPolicyLifecycleStateActive         SubsettingPolicyLifecycleStateEnum = "ACTIVE"
	SubsettingPolicyLifecycleStateUpdating       SubsettingPolicyLifecycleStateEnum = "UPDATING"
	SubsettingPolicyLifecycleStateDeleting       SubsettingPolicyLifecycleStateEnum = "DELETING"
	SubsettingPolicyLifecycleStateDeleted        SubsettingPolicyLifecycleStateEnum = "DELETED"
	SubsettingPolicyLifecycleStateNeedsAttention SubsettingPolicyLifecycleStateEnum = "NEEDS_ATTENTION"
	SubsettingPolicyLifecycleStateFailed         SubsettingPolicyLifecycleStateEnum = "FAILED"
)

var mappingSubsettingPolicyLifecycleStateEnum = map[string]SubsettingPolicyLifecycleStateEnum{
	"CREATING":        SubsettingPolicyLifecycleStateCreating,
	"ACTIVE":          SubsettingPolicyLifecycleStateActive,
	"UPDATING":        SubsettingPolicyLifecycleStateUpdating,
	"DELETING":        SubsettingPolicyLifecycleStateDeleting,
	"DELETED":         SubsettingPolicyLifecycleStateDeleted,
	"NEEDS_ATTENTION": SubsettingPolicyLifecycleStateNeedsAttention,
	"FAILED":          SubsettingPolicyLifecycleStateFailed,
}

var mappingSubsettingPolicyLifecycleStateEnumLowerCase = map[string]SubsettingPolicyLifecycleStateEnum{
	"creating":        SubsettingPolicyLifecycleStateCreating,
	"active":          SubsettingPolicyLifecycleStateActive,
	"updating":        SubsettingPolicyLifecycleStateUpdating,
	"deleting":        SubsettingPolicyLifecycleStateDeleting,
	"deleted":         SubsettingPolicyLifecycleStateDeleted,
	"needs_attention": SubsettingPolicyLifecycleStateNeedsAttention,
	"failed":          SubsettingPolicyLifecycleStateFailed,
}

// GetSubsettingPolicyLifecycleStateEnumValues Enumerates the set of values for SubsettingPolicyLifecycleStateEnum
func GetSubsettingPolicyLifecycleStateEnumValues() []SubsettingPolicyLifecycleStateEnum {
	values := make([]SubsettingPolicyLifecycleStateEnum, 0)
	for _, v := range mappingSubsettingPolicyLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetSubsettingPolicyLifecycleStateEnumStringValues Enumerates the set of values in String for SubsettingPolicyLifecycleStateEnum
func GetSubsettingPolicyLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"UPDATING",
		"DELETING",
		"DELETED",
		"NEEDS_ATTENTION",
		"FAILED",
	}
}

// GetMappingSubsettingPolicyLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSubsettingPolicyLifecycleStateEnum(val string) (SubsettingPolicyLifecycleStateEnum, bool) {
	enum, ok := mappingSubsettingPolicyLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SubsettingPolicyRecompileEnum Enum with underlying type: string
type SubsettingPolicyRecompileEnum string

// Set of constants representing the allowable values for SubsettingPolicyRecompileEnum
const (
	SubsettingPolicyRecompileSerial   SubsettingPolicyRecompileEnum = "SERIAL"
	SubsettingPolicyRecompileParallel SubsettingPolicyRecompileEnum = "PARALLEL"
	SubsettingPolicyRecompileNone     SubsettingPolicyRecompileEnum = "NONE"
)

var mappingSubsettingPolicyRecompileEnum = map[string]SubsettingPolicyRecompileEnum{
	"SERIAL":   SubsettingPolicyRecompileSerial,
	"PARALLEL": SubsettingPolicyRecompileParallel,
	"NONE":     SubsettingPolicyRecompileNone,
}

var mappingSubsettingPolicyRecompileEnumLowerCase = map[string]SubsettingPolicyRecompileEnum{
	"serial":   SubsettingPolicyRecompileSerial,
	"parallel": SubsettingPolicyRecompileParallel,
	"none":     SubsettingPolicyRecompileNone,
}

// GetSubsettingPolicyRecompileEnumValues Enumerates the set of values for SubsettingPolicyRecompileEnum
func GetSubsettingPolicyRecompileEnumValues() []SubsettingPolicyRecompileEnum {
	values := make([]SubsettingPolicyRecompileEnum, 0)
	for _, v := range mappingSubsettingPolicyRecompileEnum {
		values = append(values, v)
	}
	return values
}

// GetSubsettingPolicyRecompileEnumStringValues Enumerates the set of values in String for SubsettingPolicyRecompileEnum
func GetSubsettingPolicyRecompileEnumStringValues() []string {
	return []string{
		"SERIAL",
		"PARALLEL",
		"NONE",
	}
}

// GetMappingSubsettingPolicyRecompileEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSubsettingPolicyRecompileEnum(val string) (SubsettingPolicyRecompileEnum, bool) {
	enum, ok := mappingSubsettingPolicyRecompileEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SubsettingPolicyUnrelatedTablesActionEnum Enum with underlying type: string
type SubsettingPolicyUnrelatedTablesActionEnum string

// Set of constants representing the allowable values for SubsettingPolicyUnrelatedTablesActionEnum
const (
	SubsettingPolicyUnrelatedTablesActionTruncate SubsettingPolicyUnrelatedTablesActionEnum = "TRUNCATE"
	SubsettingPolicyUnrelatedTablesActionKeep     SubsettingPolicyUnrelatedTablesActionEnum = "KEEP"
)

var mappingSubsettingPolicyUnrelatedTablesActionEnum = map[string]SubsettingPolicyUnrelatedTablesActionEnum{
	"TRUNCATE": SubsettingPolicyUnrelatedTablesActionTruncate,
	"KEEP":     SubsettingPolicyUnrelatedTablesActionKeep,
}

var mappingSubsettingPolicyUnrelatedTablesActionEnumLowerCase = map[string]SubsettingPolicyUnrelatedTablesActionEnum{
	"truncate": SubsettingPolicyUnrelatedTablesActionTruncate,
	"keep":     SubsettingPolicyUnrelatedTablesActionKeep,
}

// GetSubsettingPolicyUnrelatedTablesActionEnumValues Enumerates the set of values for SubsettingPolicyUnrelatedTablesActionEnum
func GetSubsettingPolicyUnrelatedTablesActionEnumValues() []SubsettingPolicyUnrelatedTablesActionEnum {
	values := make([]SubsettingPolicyUnrelatedTablesActionEnum, 0)
	for _, v := range mappingSubsettingPolicyUnrelatedTablesActionEnum {
		values = append(values, v)
	}
	return values
}

// GetSubsettingPolicyUnrelatedTablesActionEnumStringValues Enumerates the set of values in String for SubsettingPolicyUnrelatedTablesActionEnum
func GetSubsettingPolicyUnrelatedTablesActionEnumStringValues() []string {
	return []string{
		"TRUNCATE",
		"KEEP",
	}
}

// GetMappingSubsettingPolicyUnrelatedTablesActionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSubsettingPolicyUnrelatedTablesActionEnum(val string) (SubsettingPolicyUnrelatedTablesActionEnum, bool) {
	enum, ok := mappingSubsettingPolicyUnrelatedTablesActionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
