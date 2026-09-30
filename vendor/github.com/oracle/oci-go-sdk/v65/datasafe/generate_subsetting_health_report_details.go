// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// GenerateSubsettingHealthReportDetails Details to use when performing health check on a subsetting policy.
type GenerateSubsettingHealthReportDetails struct {
	TargetCredentials *Credentials `mandatory:"true" json:"targetCredentials"`

	// The type of health check. The default behaviour is to perform all health checks. TABLESPACE_CHECK performs only the tablespace health check.
	CheckType GenerateSubsettingHealthReportDetailsCheckTypeEnum `mandatory:"false" json:"checkType,omitempty"`

	// The OCID of the target database to use for the subsetting policy
	// health check. The targetId associated with the subsetting policy
	// is used if this is not passed.
	TargetId *string `mandatory:"false" json:"targetId"`

	// The OCID of the compartment where the health report resource should be created.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// The tablespace that should be used to estimate space.
	// If no tablespace is provided, the DEFAULT tablespace is used.
	Tablespace *string `mandatory:"false" json:"tablespace"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m GenerateSubsettingHealthReportDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m GenerateSubsettingHealthReportDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingGenerateSubsettingHealthReportDetailsCheckTypeEnum(string(m.CheckType)); !ok && m.CheckType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CheckType: %s. Supported values are: %s.", m.CheckType, strings.Join(GetGenerateSubsettingHealthReportDetailsCheckTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// GenerateSubsettingHealthReportDetailsCheckTypeEnum Enum with underlying type: string
type GenerateSubsettingHealthReportDetailsCheckTypeEnum string

// Set of constants representing the allowable values for GenerateSubsettingHealthReportDetailsCheckTypeEnum
const (
	GenerateSubsettingHealthReportDetailsCheckTypeAll             GenerateSubsettingHealthReportDetailsCheckTypeEnum = "ALL"
	GenerateSubsettingHealthReportDetailsCheckTypeTablespaceCheck GenerateSubsettingHealthReportDetailsCheckTypeEnum = "TABLESPACE_CHECK"
)

var mappingGenerateSubsettingHealthReportDetailsCheckTypeEnum = map[string]GenerateSubsettingHealthReportDetailsCheckTypeEnum{
	"ALL":              GenerateSubsettingHealthReportDetailsCheckTypeAll,
	"TABLESPACE_CHECK": GenerateSubsettingHealthReportDetailsCheckTypeTablespaceCheck,
}

var mappingGenerateSubsettingHealthReportDetailsCheckTypeEnumLowerCase = map[string]GenerateSubsettingHealthReportDetailsCheckTypeEnum{
	"all":              GenerateSubsettingHealthReportDetailsCheckTypeAll,
	"tablespace_check": GenerateSubsettingHealthReportDetailsCheckTypeTablespaceCheck,
}

// GetGenerateSubsettingHealthReportDetailsCheckTypeEnumValues Enumerates the set of values for GenerateSubsettingHealthReportDetailsCheckTypeEnum
func GetGenerateSubsettingHealthReportDetailsCheckTypeEnumValues() []GenerateSubsettingHealthReportDetailsCheckTypeEnum {
	values := make([]GenerateSubsettingHealthReportDetailsCheckTypeEnum, 0)
	for _, v := range mappingGenerateSubsettingHealthReportDetailsCheckTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetGenerateSubsettingHealthReportDetailsCheckTypeEnumStringValues Enumerates the set of values in String for GenerateSubsettingHealthReportDetailsCheckTypeEnum
func GetGenerateSubsettingHealthReportDetailsCheckTypeEnumStringValues() []string {
	return []string{
		"ALL",
		"TABLESPACE_CHECK",
	}
}

// GetMappingGenerateSubsettingHealthReportDetailsCheckTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGenerateSubsettingHealthReportDetailsCheckTypeEnum(val string) (GenerateSubsettingHealthReportDetailsCheckTypeEnum, bool) {
	enum, ok := mappingGenerateSubsettingHealthReportDetailsCheckTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
