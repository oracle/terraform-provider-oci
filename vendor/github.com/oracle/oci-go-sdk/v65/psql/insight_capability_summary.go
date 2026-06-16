// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// PGSQL Control Plane API
//
// Use the OCI Database with PostgreSQL API to manage resources such as database systems, database nodes, backups, and configurations.
// For information, see the user guide documentation for the service (https://docs.oracle.com/iaas/Content/postgresql/home.htm).
//

package psql

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// InsightCapabilitySummary Describes supported insight types and their capabilities.
type InsightCapabilitySummary struct {

	// Echo of the requested insight type.
	InsightType InsightCapabilitySummaryInsightTypeEnum `mandatory:"true" json:"insightType"`

	// Supported insight data types for this insight type.
	DataTypeCapabilities []InsightDataTypeCapability `mandatory:"true" json:"dataTypeCapabilities"`

	// Human-readable description of the insight type.
	Description *string `mandatory:"false" json:"description"`
}

func (m InsightCapabilitySummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m InsightCapabilitySummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingInsightCapabilitySummaryInsightTypeEnum(string(m.InsightType)); !ok && m.InsightType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for InsightType: %s. Supported values are: %s.", m.InsightType, strings.Join(GetInsightCapabilitySummaryInsightTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// InsightCapabilitySummaryInsightTypeEnum Enum with underlying type: string
type InsightCapabilitySummaryInsightTypeEnum string

// Set of constants representing the allowable values for InsightCapabilitySummaryInsightTypeEnum
const (
	InsightCapabilitySummaryInsightTypeQueryInsight InsightCapabilitySummaryInsightTypeEnum = "QUERY_INSIGHT"
)

var mappingInsightCapabilitySummaryInsightTypeEnum = map[string]InsightCapabilitySummaryInsightTypeEnum{
	"QUERY_INSIGHT": InsightCapabilitySummaryInsightTypeQueryInsight,
}

var mappingInsightCapabilitySummaryInsightTypeEnumLowerCase = map[string]InsightCapabilitySummaryInsightTypeEnum{
	"query_insight": InsightCapabilitySummaryInsightTypeQueryInsight,
}

// GetInsightCapabilitySummaryInsightTypeEnumValues Enumerates the set of values for InsightCapabilitySummaryInsightTypeEnum
func GetInsightCapabilitySummaryInsightTypeEnumValues() []InsightCapabilitySummaryInsightTypeEnum {
	values := make([]InsightCapabilitySummaryInsightTypeEnum, 0)
	for _, v := range mappingInsightCapabilitySummaryInsightTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetInsightCapabilitySummaryInsightTypeEnumStringValues Enumerates the set of values in String for InsightCapabilitySummaryInsightTypeEnum
func GetInsightCapabilitySummaryInsightTypeEnumStringValues() []string {
	return []string{
		"QUERY_INSIGHT",
	}
}

// GetMappingInsightCapabilitySummaryInsightTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingInsightCapabilitySummaryInsightTypeEnum(val string) (InsightCapabilitySummaryInsightTypeEnum, bool) {
	enum, ok := mappingInsightCapabilitySummaryInsightTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
