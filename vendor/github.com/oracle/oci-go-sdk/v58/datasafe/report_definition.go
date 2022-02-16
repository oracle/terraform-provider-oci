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

// ReportDefinition Description of report definition.
type ReportDefinition struct {

	// Name of the report definition.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID of the report definition.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the compartment containing the report definition.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The current state of the report.
	LifecycleState ReportDefinitionLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The OCID of the parent report definition. In the case of seeded report definition, this is same as definition OCID.
	ParentId *string `mandatory:"false" json:"parentId"`

	// Specifies the name of the category that this report belongs to.
	Category ReportDefinitionCategoryEnum `mandatory:"false" json:"category,omitempty"`

	// A description of the report definition.
	Description *string `mandatory:"false" json:"description"`

	// Specifies the name of a resource that provides data for the report. For example alerts, events.
	DataSource ReportDefinitionDataSourceEnum `mandatory:"false" json:"dataSource,omitempty"`

	// Signifies whether the definition is seeded or user defined. Values can either be 'true' or 'false'.
	IsSeeded *bool `mandatory:"false" json:"isSeeded"`

	// Specifies how the report definitions are ordered in the display.
	DisplayOrder *int `mandatory:"false" json:"displayOrder"`

	// Specifies the time at which the report definition was created.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The date and time of the report definition update in Data Safe.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Additional scim filters used to specialize the report.
	ScimFilter *string `mandatory:"false" json:"scimFilter"`

	// An array of column objects in the order (left to right) displayed in the report. A column object stores all information about a column, including the name displayed on the UI, corresponding field name in the data source, data type of the column, and column visibility (if the column is visible to the user).
	ColumnInfo []Column `mandatory:"false" json:"columnInfo"`

	// An array of column filter objects. A column Filter object stores all information about a column filter including field name, an operator, one or more expressions, if the filter is enabled, or if the filter is hidden.
	ColumnFilters []ColumnFilter `mandatory:"false" json:"columnFilters"`

	// An array of column sorting objects. Each column sorting object stores the column name to be sorted and if the sorting is in ascending order; sorting is done by the first column in the array, then by the second column in the array, etc.
	ColumnSortings []ColumnSorting `mandatory:"false" json:"columnSortings"`

	// An array of report summary objects in the order (left to right)  displayed in the report.  A  report summary object stores all information about summary of report to be displayed, including the name displayed on UI, the display order, corresponding group by and count of values, summary visibility (if the summary is visible to user).
	Summary []Summary `mandatory:"false" json:"summary"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m ReportDefinition) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ReportDefinition) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingReportDefinitionLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetReportDefinitionLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingReportDefinitionCategoryEnum(string(m.Category)); !ok && m.Category != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Category: %s. Supported values are: %s.", m.Category, strings.Join(GetReportDefinitionCategoryEnumStringValues(), ",")))
	}
	if _, ok := GetMappingReportDefinitionDataSourceEnum(string(m.DataSource)); !ok && m.DataSource != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DataSource: %s. Supported values are: %s.", m.DataSource, strings.Join(GetReportDefinitionDataSourceEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ReportDefinitionCategoryEnum Enum with underlying type: string
type ReportDefinitionCategoryEnum string

// Set of constants representing the allowable values for ReportDefinitionCategoryEnum
const (
	ReportDefinitionCategoryCustomReports    ReportDefinitionCategoryEnum = "CUSTOM_REPORTS"
	ReportDefinitionCategorySummary          ReportDefinitionCategoryEnum = "SUMMARY"
	ReportDefinitionCategoryActivityAuditing ReportDefinitionCategoryEnum = "ACTIVITY_AUDITING"
)

var mappingReportDefinitionCategoryEnum = map[string]ReportDefinitionCategoryEnum{
	"CUSTOM_REPORTS":    ReportDefinitionCategoryCustomReports,
	"SUMMARY":           ReportDefinitionCategorySummary,
	"ACTIVITY_AUDITING": ReportDefinitionCategoryActivityAuditing,
}

// GetReportDefinitionCategoryEnumValues Enumerates the set of values for ReportDefinitionCategoryEnum
func GetReportDefinitionCategoryEnumValues() []ReportDefinitionCategoryEnum {
	values := make([]ReportDefinitionCategoryEnum, 0)
	for _, v := range mappingReportDefinitionCategoryEnum {
		values = append(values, v)
	}
	return values
}

// GetReportDefinitionCategoryEnumStringValues Enumerates the set of values in String for ReportDefinitionCategoryEnum
func GetReportDefinitionCategoryEnumStringValues() []string {
	return []string{
		"CUSTOM_REPORTS",
		"SUMMARY",
		"ACTIVITY_AUDITING",
	}
}

// GetMappingReportDefinitionCategoryEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingReportDefinitionCategoryEnum(val string) (ReportDefinitionCategoryEnum, bool) {
	mappingReportDefinitionCategoryEnumIgnoreCase := make(map[string]ReportDefinitionCategoryEnum)
	for k, v := range mappingReportDefinitionCategoryEnum {
		mappingReportDefinitionCategoryEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingReportDefinitionCategoryEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
