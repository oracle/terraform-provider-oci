// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// ReportDefinitionSummary Summary of report definition.
type ReportDefinitionSummary struct {

	// Name of the report definition.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID of the report definition.
	Id *string `mandatory:"true" json:"id"`

	// Specifies the time at which the report definition was created.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The OCID of the compartment containing the report definition.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The current state of the report
	LifecycleState ReportDefinitionLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Specifies the name of the category that this report belongs to.
	Category ReportDefinitionSummaryCategoryEnum `mandatory:"false" json:"category,omitempty"`

	// A description of the report definition.
	Description *string `mandatory:"false" json:"description"`

	// Signifies whether the definition is seeded or user defined. Values can either be 'true' or 'false'.
	IsSeeded *bool `mandatory:"false" json:"isSeeded"`

	// Specifies how the report definitions are ordered in the display.
	DisplayOrder *int `mandatory:"false" json:"displayOrder"`

	// The date and time of the report definition update in Data Safe.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Specifies the name of a resource that provides data for the report. For example alerts, events.
	DataSource ReportDefinitionDataSourceEnum `mandatory:"false" json:"dataSource,omitempty"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m ReportDefinitionSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ReportDefinitionSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingReportDefinitionLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetReportDefinitionLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingReportDefinitionSummaryCategoryEnum(string(m.Category)); !ok && m.Category != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Category: %s. Supported values are: %s.", m.Category, strings.Join(GetReportDefinitionSummaryCategoryEnumStringValues(), ",")))
	}
	if _, ok := GetMappingReportDefinitionDataSourceEnum(string(m.DataSource)); !ok && m.DataSource != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DataSource: %s. Supported values are: %s.", m.DataSource, strings.Join(GetReportDefinitionDataSourceEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ReportDefinitionSummaryCategoryEnum Enum with underlying type: string
type ReportDefinitionSummaryCategoryEnum string

// Set of constants representing the allowable values for ReportDefinitionSummaryCategoryEnum
const (
	ReportDefinitionSummaryCategoryCustomReports    ReportDefinitionSummaryCategoryEnum = "CUSTOM_REPORTS"
	ReportDefinitionSummaryCategorySummary          ReportDefinitionSummaryCategoryEnum = "SUMMARY"
	ReportDefinitionSummaryCategoryActivityAuditing ReportDefinitionSummaryCategoryEnum = "ACTIVITY_AUDITING"
)

var mappingReportDefinitionSummaryCategoryEnum = map[string]ReportDefinitionSummaryCategoryEnum{
	"CUSTOM_REPORTS":    ReportDefinitionSummaryCategoryCustomReports,
	"SUMMARY":           ReportDefinitionSummaryCategorySummary,
	"ACTIVITY_AUDITING": ReportDefinitionSummaryCategoryActivityAuditing,
}

// GetReportDefinitionSummaryCategoryEnumValues Enumerates the set of values for ReportDefinitionSummaryCategoryEnum
func GetReportDefinitionSummaryCategoryEnumValues() []ReportDefinitionSummaryCategoryEnum {
	values := make([]ReportDefinitionSummaryCategoryEnum, 0)
	for _, v := range mappingReportDefinitionSummaryCategoryEnum {
		values = append(values, v)
	}
	return values
}

// GetReportDefinitionSummaryCategoryEnumStringValues Enumerates the set of values in String for ReportDefinitionSummaryCategoryEnum
func GetReportDefinitionSummaryCategoryEnumStringValues() []string {
	return []string{
		"CUSTOM_REPORTS",
		"SUMMARY",
		"ACTIVITY_AUDITING",
	}
}

// GetMappingReportDefinitionSummaryCategoryEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingReportDefinitionSummaryCategoryEnum(val string) (ReportDefinitionSummaryCategoryEnum, bool) {
	mappingReportDefinitionSummaryCategoryEnumIgnoreCase := make(map[string]ReportDefinitionSummaryCategoryEnum)
	for k, v := range mappingReportDefinitionSummaryCategoryEnum {
		mappingReportDefinitionSummaryCategoryEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingReportDefinitionSummaryCategoryEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
