// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// IngestTimeRuleSummary Summary of an ingest time rule.
type IngestTimeRuleSummary struct {

	// The log analytics entity OCID. This ID is a reference used by log analytics features and it represents
	// a resource that is provisioned and managed by the customer on their premises or on the cloud.
	Id *string `mandatory:"true" json:"id"`

	// Compartment Identifier OCID  (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The ingest time rule display name.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Description for this resource.
	Description *string `mandatory:"false" json:"description"`

	// The date and time the resource was created, in the format defined by RFC3339.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The date and time the resource was last updated, in the format defined by RFC3339.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The current state of the ingest time rule.
	LifecycleState ConfigLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// A flag indicating whether or not the ingest time rule is enabled.
	IsEnabled *bool `mandatory:"false" json:"isEnabled"`

	// The ingest time rule condition kind.
	ConditionKind IngestTimeRuleSummaryConditionKindEnum `mandatory:"false" json:"conditionKind,omitempty"`

	// The ingest time rule condition field name.
	FieldName *string `mandatory:"false" json:"fieldName"`

	// The ingest time rule condition field value.
	FieldValue *string `mandatory:"false" json:"fieldValue"`
}

func (m IngestTimeRuleSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m IngestTimeRuleSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingConfigLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetConfigLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingIngestTimeRuleSummaryConditionKindEnum(string(m.ConditionKind)); !ok && m.ConditionKind != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ConditionKind: %s. Supported values are: %s.", m.ConditionKind, strings.Join(GetIngestTimeRuleSummaryConditionKindEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// IngestTimeRuleSummaryConditionKindEnum Enum with underlying type: string
type IngestTimeRuleSummaryConditionKindEnum string

// Set of constants representing the allowable values for IngestTimeRuleSummaryConditionKindEnum
const (
	IngestTimeRuleSummaryConditionKindField IngestTimeRuleSummaryConditionKindEnum = "FIELD"
)

var mappingIngestTimeRuleSummaryConditionKindEnum = map[string]IngestTimeRuleSummaryConditionKindEnum{
	"FIELD": IngestTimeRuleSummaryConditionKindField,
}

var mappingIngestTimeRuleSummaryConditionKindEnumLowerCase = map[string]IngestTimeRuleSummaryConditionKindEnum{
	"field": IngestTimeRuleSummaryConditionKindField,
}

// GetIngestTimeRuleSummaryConditionKindEnumValues Enumerates the set of values for IngestTimeRuleSummaryConditionKindEnum
func GetIngestTimeRuleSummaryConditionKindEnumValues() []IngestTimeRuleSummaryConditionKindEnum {
	values := make([]IngestTimeRuleSummaryConditionKindEnum, 0)
	for _, v := range mappingIngestTimeRuleSummaryConditionKindEnum {
		values = append(values, v)
	}
	return values
}

// GetIngestTimeRuleSummaryConditionKindEnumStringValues Enumerates the set of values in String for IngestTimeRuleSummaryConditionKindEnum
func GetIngestTimeRuleSummaryConditionKindEnumStringValues() []string {
	return []string{
		"FIELD",
	}
}

// GetMappingIngestTimeRuleSummaryConditionKindEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingIngestTimeRuleSummaryConditionKindEnum(val string) (IngestTimeRuleSummaryConditionKindEnum, bool) {
	enum, ok := mappingIngestTimeRuleSummaryConditionKindEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
