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

// CreateSubsettingPolicyDetails Details to create a new subsetting policy. Use either a sensitive data model or a target database as the source
// of schemas for subsetting.
// To use a sensitive data model as the source of schemas, set the schemaSource attribute to SENSITIVE_DATA_MODEL and
// provide the sensitiveDataModelId attribute. The schemas from the SDM will be used for subsetting.
// To use a target database as the source of schemas, set the schemaSource attribute to TARGET and provide the targetId
// attribute along with the schemasForSubsetting list.
// The schema list can also be provided or updated later using the UpdateSubsettingPolicy operation.
// After creating a subsetting policy, you can use operations to add or modify subsetting rules as needed
type CreateSubsettingPolicyDetails struct {

	// The OCID of the compartment where the subsetting policy should be created
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	SchemaSource CreateSchemaSourceDetails `mandatory:"true" json:"schemaSource"`

	// The display name of the subsetting policy. The name does not have to be unique, and it's changeable
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The description of the subsetting policy
	Description *string `mandatory:"false" json:"description"`

	// Indicates if redo logging is enabled during a subsetting operation. It's disabled by default. Set this attribute to true to
	// enable redo logging. By default, subsetting disables redo logging and flashback logging to purge any original
	// data from logs. However, in certain circumstances when you only want to test subsetting, rollback changes, and retry subsetting,
	// you could enable logging and use a flashback database to retrieve the original data after it has been subsetted
	IsRedoLoggingEnabled *bool `mandatory:"false" json:"isRedoLoggingEnabled"`

	// Indicates if statistics gathering is enabled. It's enabled by default. Set this attribute to false to disable statistics
	// gathering. The subsetting process gathers statistics on subsetted database tables after subsetting completes
	IsRefreshStatsEnabled *bool `mandatory:"false" json:"isRefreshStatsEnabled"`

	// Specifies options to enable parallel execution when running data subsetting. Allowed values are 'NONE' (no parallelism),
	// 'DEFAULT' (the Oracle Database computes the optimum degree of parallelism) or an integer value to be used as the degree
	// of parallelism. Parallel execution helps effectively use multiple CPUs and improve subsetting performance. Refer to the
	// Oracle Database parallel execution framework when choosing an explicit degree of parallelism
	ParallelDegree *string `mandatory:"false" json:"parallelDegree"`

	// Specifies how to recompile invalid objects post data subsetting. Allowed values are 'SERIAL' (recompile in serial),
	// 'PARALLEL' (recompile in parallel), 'NONE' (do not recompile). If it's set to PARALLEL, the value of parallelDegree
	// attribute is used. Use the built-in UTL_RECOMP package to recompile any remaining invalid objects after subsetting completes
	Recompile SubsettingPolicyRecompileEnum `mandatory:"false" json:"recompile,omitempty"`

	// Strategy to be applied for tables which are not impacted by any of the subsetting rules
	UnrelatedTablesAction SubsettingPolicyUnrelatedTablesActionEnum `mandatory:"false" json:"unrelatedTablesAction,omitempty"`

	// A pre-subsetting script, which can contain SQL and PL/SQL statements. It's executed before
	// the subsetting process. It's usually used to perform
	// any preparation or prerequisite work before subsetting data.
	PreSubsettingScript *string `mandatory:"false" json:"preSubsettingScript"`

	// A post-subsetting script, which can contain SQL and PL/SQL statements. It's executed after
	// the subsetting process. It's usually used to perform
	// additional transformation or cleanup work after subsetting data.
	PostSubsettingScript *string `mandatory:"false" json:"postSubsettingScript"`

	// The OCID of the masking policy to associate with this subsetting policy
	MaskingPolicyId *string `mandatory:"false" json:"maskingPolicyId"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateSubsettingPolicyDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateSubsettingPolicyDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

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
func (m *CreateSubsettingPolicyDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName           *string                                   `json:"displayName"`
		Description           *string                                   `json:"description"`
		IsRedoLoggingEnabled  *bool                                     `json:"isRedoLoggingEnabled"`
		IsRefreshStatsEnabled *bool                                     `json:"isRefreshStatsEnabled"`
		ParallelDegree        *string                                   `json:"parallelDegree"`
		Recompile             SubsettingPolicyRecompileEnum             `json:"recompile"`
		UnrelatedTablesAction SubsettingPolicyUnrelatedTablesActionEnum `json:"unrelatedTablesAction"`
		PreSubsettingScript   *string                                   `json:"preSubsettingScript"`
		PostSubsettingScript  *string                                   `json:"postSubsettingScript"`
		MaskingPolicyId       *string                                   `json:"maskingPolicyId"`
		FreeformTags          map[string]string                         `json:"freeformTags"`
		DefinedTags           map[string]map[string]interface{}         `json:"definedTags"`
		CompartmentId         *string                                   `json:"compartmentId"`
		SchemaSource          createschemasourcedetails                 `json:"schemaSource"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	m.Description = model.Description

	m.IsRedoLoggingEnabled = model.IsRedoLoggingEnabled

	m.IsRefreshStatsEnabled = model.IsRefreshStatsEnabled

	m.ParallelDegree = model.ParallelDegree

	m.Recompile = model.Recompile

	m.UnrelatedTablesAction = model.UnrelatedTablesAction

	m.PreSubsettingScript = model.PreSubsettingScript

	m.PostSubsettingScript = model.PostSubsettingScript

	m.MaskingPolicyId = model.MaskingPolicyId

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.CompartmentId = model.CompartmentId

	nn, e = model.SchemaSource.UnmarshalPolymorphicJSON(model.SchemaSource.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.SchemaSource = nn.(CreateSchemaSourceDetails)
	} else {
		m.SchemaSource = nil
	}

	return
}
