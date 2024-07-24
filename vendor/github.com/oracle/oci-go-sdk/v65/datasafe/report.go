// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// Report The description of the report.
type Report struct {

	// The OCID of the report.
	Id *string `mandatory:"true" json:"id"`

	// Name of the report.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID of the compartment containing the report.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Specifies the date and time the report was generated.
	TimeGenerated *common.SDKTime `mandatory:"true" json:"timeGenerated"`

	// The current state of the audit report.
	LifecycleState ReportLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The OCID of the report definition.
	ReportDefinitionId *string `mandatory:"false" json:"reportDefinitionId"`

	// Specifies a description of the report.
	Description *string `mandatory:"false" json:"description"`

	// Specifies the format of report to be .xls or .pdf or .json
	MimeType ReportMimeTypeEnum `mandatory:"false" json:"mimeType,omitempty"`

	// The type of the audit report.
	Type ReportTypeEnum `mandatory:"false" json:"type,omitempty"`

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

func (m Report) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Report) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingReportLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetReportLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingReportMimeTypeEnum(string(m.MimeType)); !ok && m.MimeType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for MimeType: %s. Supported values are: %s.", m.MimeType, strings.Join(GetReportMimeTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingReportTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetReportTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ReportMimeTypeEnum Enum with underlying type: string
type ReportMimeTypeEnum string

// Set of constants representing the allowable values for ReportMimeTypeEnum
const (
	ReportMimeTypePdf ReportMimeTypeEnum = "PDF"
	ReportMimeTypeXls ReportMimeTypeEnum = "XLS"
)

var mappingReportMimeTypeEnum = map[string]ReportMimeTypeEnum{
	"PDF": ReportMimeTypePdf,
	"XLS": ReportMimeTypeXls,
}

var mappingReportMimeTypeEnumLowerCase = map[string]ReportMimeTypeEnum{
	"pdf": ReportMimeTypePdf,
	"xls": ReportMimeTypeXls,
}

// GetReportMimeTypeEnumValues Enumerates the set of values for ReportMimeTypeEnum
func GetReportMimeTypeEnumValues() []ReportMimeTypeEnum {
	values := make([]ReportMimeTypeEnum, 0)
	for _, v := range mappingReportMimeTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetReportMimeTypeEnumStringValues Enumerates the set of values in String for ReportMimeTypeEnum
func GetReportMimeTypeEnumStringValues() []string {
	return []string{
		"PDF",
		"XLS",
	}
}

// GetMappingReportMimeTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingReportMimeTypeEnum(val string) (ReportMimeTypeEnum, bool) {
	enum, ok := mappingReportMimeTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
