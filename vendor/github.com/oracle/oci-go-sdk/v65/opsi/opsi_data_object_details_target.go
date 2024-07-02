// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Ops Insights API
//
// Use the Ops Insights API to perform data extraction operations to obtain database
// resource utilization, performance statistics, and reference information. For more information,
// see About Oracle Cloud Infrastructure Ops Insights (https://docs.cloud.oracle.com/en-us/iaas/operations-insights/doc/operations-insights.html).
//

package opsi

import (
	"strings"
)

// OpsiDataObjectDetailsTargetEnum Enum with underlying type: string
type OpsiDataObjectDetailsTargetEnum string

// Set of constants representing the allowable values for OpsiDataObjectDetailsTargetEnum
const (
	OpsiDataObjectDetailsTargetIndividualOpsidataobject          OpsiDataObjectDetailsTargetEnum = "INDIVIDUAL_OPSIDATAOBJECT"
	OpsiDataObjectDetailsTargetOpsidataobjecttypeOpsidataobjects OpsiDataObjectDetailsTargetEnum = "OPSIDATAOBJECTTYPE_OPSIDATAOBJECTS"
)

var mappingOpsiDataObjectDetailsTargetEnum = map[string]OpsiDataObjectDetailsTargetEnum{
	"INDIVIDUAL_OPSIDATAOBJECT":          OpsiDataObjectDetailsTargetIndividualOpsidataobject,
	"OPSIDATAOBJECTTYPE_OPSIDATAOBJECTS": OpsiDataObjectDetailsTargetOpsidataobjecttypeOpsidataobjects,
}

var mappingOpsiDataObjectDetailsTargetEnumLowerCase = map[string]OpsiDataObjectDetailsTargetEnum{
	"individual_opsidataobject":          OpsiDataObjectDetailsTargetIndividualOpsidataobject,
	"opsidataobjecttype_opsidataobjects": OpsiDataObjectDetailsTargetOpsidataobjecttypeOpsidataobjects,
}

// GetOpsiDataObjectDetailsTargetEnumValues Enumerates the set of values for OpsiDataObjectDetailsTargetEnum
func GetOpsiDataObjectDetailsTargetEnumValues() []OpsiDataObjectDetailsTargetEnum {
	values := make([]OpsiDataObjectDetailsTargetEnum, 0)
	for _, v := range mappingOpsiDataObjectDetailsTargetEnum {
		values = append(values, v)
	}
	return values
}

// GetOpsiDataObjectDetailsTargetEnumStringValues Enumerates the set of values in String for OpsiDataObjectDetailsTargetEnum
func GetOpsiDataObjectDetailsTargetEnumStringValues() []string {
	return []string{
		"INDIVIDUAL_OPSIDATAOBJECT",
		"OPSIDATAOBJECTTYPE_OPSIDATAOBJECTS",
	}
}

// GetMappingOpsiDataObjectDetailsTargetEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOpsiDataObjectDetailsTargetEnum(val string) (OpsiDataObjectDetailsTargetEnum, bool) {
	enum, ok := mappingOpsiDataObjectDetailsTargetEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
