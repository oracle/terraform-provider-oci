// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
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

	// Specifies the data and time the report definition was created.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The date and time the report definition was update.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Additional SCIM filters used to define the report.
	ScimFilter *string `mandatory:"false" json:"scimFilter"`

	// An array of column objects in the order (left to right) displayed in the report. A column object stores all information about a column, including the name displayed on the UI, corresponding field name in the data source, data type of the column, and column visibility (if the column is visible to the user).
	ColumnInfo []Column `mandatory:"false" json:"columnInfo"`

	// An array of columnFilter objects. A columnFilter object stores all information about a column filter including field name, an operator, one or more expressions, if the filter is enabled, or if the filter is hidden.
	ColumnFilters []ColumnFilter `mandatory:"false" json:"columnFilters"`

	// An array of column sorting objects. Each column sorting object stores the column name to be sorted and if the sorting is in ascending order; sorting is done by the first column in the array, then by the second column in the array, etc.
	ColumnSortings []ColumnSorting `mandatory:"false" json:"columnSortings"`

	// An array of report summary objects in the order (left to right)  displayed in the report.  A  report summary object stores all information about summary of report to be displayed, including the name displayed on UI, the display order, corresponding group by and count of values, summary visibility (if the summary is visible to user).
	Summary []Summary `mandatory:"false" json:"summary"`

	// The schedule to generate the report periodically in the specified format:
	// <version-string>;<version-specific-schedule>
	// Allowed version strings - "v1"
	// v1's version specific schedule -<ss> <mm> <hh> <day-of-week> <day-of-month>
	// Each of the above fields potentially introduce constraints. A workrequest is created only
	// when clock time satisfies all the constraints. Constraints introduced:
	// 1. seconds = <ss> (So, the allowed range for <ss> is [0, 59])
	// 2. minutes = <mm> (So, the allowed range for <mm> is [0, 59])
	// 3. hours = <hh> (So, the allowed range for <hh> is [0, 23])
	// 4. <day-of-week> can be either '*' (without quotes or a number between 1(Monday) and 7(Sunday))
	// No constraint introduced when it is '*'. When not, day of week must equal the given value
	// 5. <day-of-month> can be either '*' (without quotes or a number between 1 and 28)
	// No constraint introduced when it is '*'. When not, day of month must equal the given value
	Schedule *string `mandatory:"false" json:"schedule"`

	// Specifies the format of the report ( either XLS or PDF )
	ScheduledReportMimeType ReportDefinitionScheduledReportMimeTypeEnum `mandatory:"false" json:"scheduledReportMimeType,omitempty"`

	// Specifies the limit on the number of rows in the report.
	ScheduledReportRowLimit *int `mandatory:"false" json:"scheduledReportRowLimit"`

	// The name of the report to be scheduled.
	ScheduledReportName *string `mandatory:"false" json:"scheduledReportName"`

	// The OCID of the compartment in which the scheduled resource should be created.
	ScheduledReportCompartmentId *string `mandatory:"false" json:"scheduledReportCompartmentId"`

	// The time span for the records in the report to be scheduled.
	// <period-value><period>
	// Allowed period strings - "H","D","M","Y"
	// Each of the above fields potentially introduce constraints. A workRequest is created only
	// when period-value satisfies all the constraints. Constraints introduced:
	// 1. period = H (The allowed range for period-value is [1, 23])
	// 2. period = D (The allowed range for period-value is [1, 30])
	// 3. period = M (The allowed range for period-value is [1, 11])
	// 4. period = Y (The minimum period-value is 1)
	RecordTimeSpan *string `mandatory:"false" json:"recordTimeSpan"`

	// The list of the data protection regulations/standards used in the report that will help demonstrate compliance.
	ComplianceStandards []string `mandatory:"false" json:"complianceStandards"`

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
	if _, ok := GetMappingReportDefinitionScheduledReportMimeTypeEnum(string(m.ScheduledReportMimeType)); !ok && m.ScheduledReportMimeType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ScheduledReportMimeType: %s. Supported values are: %s.", m.ScheduledReportMimeType, strings.Join(GetReportDefinitionScheduledReportMimeTypeEnumStringValues(), ",")))
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
	ReportDefinitionCategoryCustomReports      ReportDefinitionCategoryEnum = "CUSTOM_REPORTS"
	ReportDefinitionCategorySummary            ReportDefinitionCategoryEnum = "SUMMARY"
	ReportDefinitionCategoryActivityAuditing   ReportDefinitionCategoryEnum = "ACTIVITY_AUDITING"
	ReportDefinitionCategoryFirewallViolations ReportDefinitionCategoryEnum = "FIREWALL_VIOLATIONS"
	ReportDefinitionCategoryAllowedSql         ReportDefinitionCategoryEnum = "ALLOWED_SQL"
)

var mappingReportDefinitionCategoryEnum = map[string]ReportDefinitionCategoryEnum{
	"CUSTOM_REPORTS":      ReportDefinitionCategoryCustomReports,
	"SUMMARY":             ReportDefinitionCategorySummary,
	"ACTIVITY_AUDITING":   ReportDefinitionCategoryActivityAuditing,
	"FIREWALL_VIOLATIONS": ReportDefinitionCategoryFirewallViolations,
	"ALLOWED_SQL":         ReportDefinitionCategoryAllowedSql,
}

var mappingReportDefinitionCategoryEnumLowerCase = map[string]ReportDefinitionCategoryEnum{
	"custom_reports":      ReportDefinitionCategoryCustomReports,
	"summary":             ReportDefinitionCategorySummary,
	"activity_auditing":   ReportDefinitionCategoryActivityAuditing,
	"firewall_violations": ReportDefinitionCategoryFirewallViolations,
	"allowed_sql":         ReportDefinitionCategoryAllowedSql,
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
		"FIREWALL_VIOLATIONS",
		"ALLOWED_SQL",
	}
}

// GetMappingReportDefinitionCategoryEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingReportDefinitionCategoryEnum(val string) (ReportDefinitionCategoryEnum, bool) {
	enum, ok := mappingReportDefinitionCategoryEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ReportDefinitionScheduledReportMimeTypeEnum Enum with underlying type: string
type ReportDefinitionScheduledReportMimeTypeEnum string

// Set of constants representing the allowable values for ReportDefinitionScheduledReportMimeTypeEnum
const (
	ReportDefinitionScheduledReportMimeTypePdf ReportDefinitionScheduledReportMimeTypeEnum = "PDF"
	ReportDefinitionScheduledReportMimeTypeXls ReportDefinitionScheduledReportMimeTypeEnum = "XLS"
)

var mappingReportDefinitionScheduledReportMimeTypeEnum = map[string]ReportDefinitionScheduledReportMimeTypeEnum{
	"PDF": ReportDefinitionScheduledReportMimeTypePdf,
	"XLS": ReportDefinitionScheduledReportMimeTypeXls,
}

var mappingReportDefinitionScheduledReportMimeTypeEnumLowerCase = map[string]ReportDefinitionScheduledReportMimeTypeEnum{
	"pdf": ReportDefinitionScheduledReportMimeTypePdf,
	"xls": ReportDefinitionScheduledReportMimeTypeXls,
}

// GetReportDefinitionScheduledReportMimeTypeEnumValues Enumerates the set of values for ReportDefinitionScheduledReportMimeTypeEnum
func GetReportDefinitionScheduledReportMimeTypeEnumValues() []ReportDefinitionScheduledReportMimeTypeEnum {
	values := make([]ReportDefinitionScheduledReportMimeTypeEnum, 0)
	for _, v := range mappingReportDefinitionScheduledReportMimeTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetReportDefinitionScheduledReportMimeTypeEnumStringValues Enumerates the set of values in String for ReportDefinitionScheduledReportMimeTypeEnum
func GetReportDefinitionScheduledReportMimeTypeEnumStringValues() []string {
	return []string{
		"PDF",
		"XLS",
	}
}

// GetMappingReportDefinitionScheduledReportMimeTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingReportDefinitionScheduledReportMimeTypeEnum(val string) (ReportDefinitionScheduledReportMimeTypeEnum, bool) {
	enum, ok := mappingReportDefinitionScheduledReportMimeTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
