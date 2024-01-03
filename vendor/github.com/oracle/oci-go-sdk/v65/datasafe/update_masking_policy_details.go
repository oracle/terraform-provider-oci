// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// UpdateMaskingPolicyDetails Details to update a masking policy.
type UpdateMaskingPolicyDetails struct {

	// The display name of the masking policy. The name does not have to be unique, and it's changeable.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The description of the masking policy.
	Description *string `mandatory:"false" json:"description"`

	// Indicates if the temporary tables created during a masking operation should be dropped after masking. It's enabled by default.
	// Set this attribute to false to preserve the temporary tables. Masking creates temporary tables that map the original sensitive
	// data values to mask values. By default, these temporary tables are dropped after masking. But, in some cases, you may want
	// to preserve this information to track how masking changed your data. Note that doing so compromises security. These tables
	// must be dropped before the database is available for unprivileged users.
	IsDropTempTablesEnabled *bool `mandatory:"false" json:"isDropTempTablesEnabled"`

	// Indicates if redo logging is enabled during a masking operation. It's disabled by default. Set this attribute to true to
	// enable redo logging. By default, masking disables redo logging and flashback logging to purge any original unmasked
	// data from logs. However, in certain circumstances when you only want to test masking, rollback changes, and retry masking,
	// you could enable logging and use a flashback database to retrieve the original unmasked data after it has been masked.
	IsRedoLoggingEnabled *bool `mandatory:"false" json:"isRedoLoggingEnabled"`

	// Indicates if statistics gathering is enabled. It's enabled by default. Set this attribute to false to disable statistics
	// gathering. The masking process gathers statistics on masked database tables after masking completes.
	IsRefreshStatsEnabled *bool `mandatory:"false" json:"isRefreshStatsEnabled"`

	// Specifies options to enable parallel execution when running data masking. Allowed values are 'NONE' (no parallelism),
	// 'DEFAULT' (the Oracle Database computes the optimum degree of parallelism) or an integer value to be used as the degree
	// of parallelism. Parallel execution helps effectively use multiple CPUs and improve masking performance. Refer to the
	// Oracle Database parallel execution framework when choosing an explicit degree of parallelism.
	ParallelDegree *string `mandatory:"false" json:"parallelDegree"`

	// Specifies how to recompile invalid objects post data masking. Allowed values are 'SERIAL' (recompile in serial),
	// 'PARALLEL' (recompile in parallel), 'NONE' (do not recompile). If it's set to PARALLEL, the value of parallelDegree
	// attribute is used. Use the built-in UTL_RECOMP package to recompile any remaining invalid objects after masking completes.
	Recompile MaskingPolicyRecompileEnum `mandatory:"false" json:"recompile,omitempty"`

	// A pre-masking script, which can contain SQL and PL/SQL statements. It's executed before
	// the core masking script generated using the masking policy. It's usually used to perform
	// any preparation or prerequisite work before masking data.
	PreMaskingScript *string `mandatory:"false" json:"preMaskingScript"`

	// A post-masking script, which can contain SQL and PL/SQL statements. It's executed after
	// the core masking script generated using the masking policy. It's usually used to perform
	// additional transformation or cleanup work after masking.
	PostMaskingScript *string `mandatory:"false" json:"postMaskingScript"`

	ColumnSource UpdateColumnSourceDetails `mandatory:"false" json:"columnSource"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m UpdateMaskingPolicyDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateMaskingPolicyDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingMaskingPolicyRecompileEnum(string(m.Recompile)); !ok && m.Recompile != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Recompile: %s. Supported values are: %s.", m.Recompile, strings.Join(GetMaskingPolicyRecompileEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *UpdateMaskingPolicyDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName             *string                           `json:"displayName"`
		Description             *string                           `json:"description"`
		IsDropTempTablesEnabled *bool                             `json:"isDropTempTablesEnabled"`
		IsRedoLoggingEnabled    *bool                             `json:"isRedoLoggingEnabled"`
		IsRefreshStatsEnabled   *bool                             `json:"isRefreshStatsEnabled"`
		ParallelDegree          *string                           `json:"parallelDegree"`
		Recompile               MaskingPolicyRecompileEnum        `json:"recompile"`
		PreMaskingScript        *string                           `json:"preMaskingScript"`
		PostMaskingScript       *string                           `json:"postMaskingScript"`
		ColumnSource            updatecolumnsourcedetails         `json:"columnSource"`
		FreeformTags            map[string]string                 `json:"freeformTags"`
		DefinedTags             map[string]map[string]interface{} `json:"definedTags"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	m.Description = model.Description

	m.IsDropTempTablesEnabled = model.IsDropTempTablesEnabled

	m.IsRedoLoggingEnabled = model.IsRedoLoggingEnabled

	m.IsRefreshStatsEnabled = model.IsRefreshStatsEnabled

	m.ParallelDegree = model.ParallelDegree

	m.Recompile = model.Recompile

	m.PreMaskingScript = model.PreMaskingScript

	m.PostMaskingScript = model.PostMaskingScript

	nn, e = model.ColumnSource.UnmarshalPolymorphicJSON(model.ColumnSource.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.ColumnSource = nn.(UpdateColumnSourceDetails)
	} else {
		m.ColumnSource = nil
	}

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	return
}
