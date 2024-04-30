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

// GenerateHealthReportDetails Details to use when performing health check on a masking policy.
type GenerateHealthReportDetails struct {

	// The type of health check. The default behaviour is to perform all health checks.
	CheckType GenerateHealthReportDetailsCheckTypeEnum `mandatory:"false" json:"checkType,omitempty"`

	// The OCID of the target database to use for the masking policy
	// health check. The targetId associated with the masking policy
	// is used if this is not passed.
	TargetId *string `mandatory:"false" json:"targetId"`

	// The OCID of the compartment where the health report resource should be created.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// The tablespace that should be used to estimate space.
	// If no tablespace is provided, the DEFAULT tablespace is used.
	Tablespace *string `mandatory:"false" json:"tablespace"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m GenerateHealthReportDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m GenerateHealthReportDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingGenerateHealthReportDetailsCheckTypeEnum(string(m.CheckType)); !ok && m.CheckType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CheckType: %s. Supported values are: %s.", m.CheckType, strings.Join(GetGenerateHealthReportDetailsCheckTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// GenerateHealthReportDetailsCheckTypeEnum Enum with underlying type: string
type GenerateHealthReportDetailsCheckTypeEnum string

// Set of constants representing the allowable values for GenerateHealthReportDetailsCheckTypeEnum
const (
	GenerateHealthReportDetailsCheckTypeAll GenerateHealthReportDetailsCheckTypeEnum = "ALL"
)

var mappingGenerateHealthReportDetailsCheckTypeEnum = map[string]GenerateHealthReportDetailsCheckTypeEnum{
	"ALL": GenerateHealthReportDetailsCheckTypeAll,
}

var mappingGenerateHealthReportDetailsCheckTypeEnumLowerCase = map[string]GenerateHealthReportDetailsCheckTypeEnum{
	"all": GenerateHealthReportDetailsCheckTypeAll,
}

// GetGenerateHealthReportDetailsCheckTypeEnumValues Enumerates the set of values for GenerateHealthReportDetailsCheckTypeEnum
func GetGenerateHealthReportDetailsCheckTypeEnumValues() []GenerateHealthReportDetailsCheckTypeEnum {
	values := make([]GenerateHealthReportDetailsCheckTypeEnum, 0)
	for _, v := range mappingGenerateHealthReportDetailsCheckTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetGenerateHealthReportDetailsCheckTypeEnumStringValues Enumerates the set of values in String for GenerateHealthReportDetailsCheckTypeEnum
func GetGenerateHealthReportDetailsCheckTypeEnumStringValues() []string {
	return []string{
		"ALL",
	}
}

// GetMappingGenerateHealthReportDetailsCheckTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGenerateHealthReportDetailsCheckTypeEnum(val string) (GenerateHealthReportDetailsCheckTypeEnum, bool) {
	enum, ok := mappingGenerateHealthReportDetailsCheckTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
