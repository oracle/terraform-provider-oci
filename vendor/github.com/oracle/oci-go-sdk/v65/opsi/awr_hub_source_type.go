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

// AwrHubSourceTypeEnum Enum with underlying type: string
type AwrHubSourceTypeEnum string

// Set of constants representing the allowable values for AwrHubSourceTypeEnum
const (
	AwrHubSourceTypeAdwS                 AwrHubSourceTypeEnum = "ADW_S"
	AwrHubSourceTypeAtpS                 AwrHubSourceTypeEnum = "ATP_S"
	AwrHubSourceTypeAdwD                 AwrHubSourceTypeEnum = "ADW_D"
	AwrHubSourceTypeAtpD                 AwrHubSourceTypeEnum = "ATP_D"
	AwrHubSourceTypeExternalPdb          AwrHubSourceTypeEnum = "EXTERNAL_PDB"
	AwrHubSourceTypeExternalNoncdb       AwrHubSourceTypeEnum = "EXTERNAL_NONCDB"
	AwrHubSourceTypeComanagedVmCdb       AwrHubSourceTypeEnum = "COMANAGED_VM_CDB"
	AwrHubSourceTypeComanagedVmPdb       AwrHubSourceTypeEnum = "COMANAGED_VM_PDB"
	AwrHubSourceTypeComanagedVmNoncdb    AwrHubSourceTypeEnum = "COMANAGED_VM_NONCDB"
	AwrHubSourceTypeComanagedBmCdb       AwrHubSourceTypeEnum = "COMANAGED_BM_CDB"
	AwrHubSourceTypeComanagedBmPdb       AwrHubSourceTypeEnum = "COMANAGED_BM_PDB"
	AwrHubSourceTypeComanagedBmNoncdb    AwrHubSourceTypeEnum = "COMANAGED_BM_NONCDB"
	AwrHubSourceTypeComanagedExacsCdb    AwrHubSourceTypeEnum = "COMANAGED_EXACS_CDB"
	AwrHubSourceTypeComanagedExacsPdb    AwrHubSourceTypeEnum = "COMANAGED_EXACS_PDB"
	AwrHubSourceTypeComanagedExacsNoncdb AwrHubSourceTypeEnum = "COMANAGED_EXACS_NONCDB"
	AwrHubSourceTypeUndefined            AwrHubSourceTypeEnum = "UNDEFINED"
)

var mappingAwrHubSourceTypeEnum = map[string]AwrHubSourceTypeEnum{
	"ADW_S":                  AwrHubSourceTypeAdwS,
	"ATP_S":                  AwrHubSourceTypeAtpS,
	"ADW_D":                  AwrHubSourceTypeAdwD,
	"ATP_D":                  AwrHubSourceTypeAtpD,
	"EXTERNAL_PDB":           AwrHubSourceTypeExternalPdb,
	"EXTERNAL_NONCDB":        AwrHubSourceTypeExternalNoncdb,
	"COMANAGED_VM_CDB":       AwrHubSourceTypeComanagedVmCdb,
	"COMANAGED_VM_PDB":       AwrHubSourceTypeComanagedVmPdb,
	"COMANAGED_VM_NONCDB":    AwrHubSourceTypeComanagedVmNoncdb,
	"COMANAGED_BM_CDB":       AwrHubSourceTypeComanagedBmCdb,
	"COMANAGED_BM_PDB":       AwrHubSourceTypeComanagedBmPdb,
	"COMANAGED_BM_NONCDB":    AwrHubSourceTypeComanagedBmNoncdb,
	"COMANAGED_EXACS_CDB":    AwrHubSourceTypeComanagedExacsCdb,
	"COMANAGED_EXACS_PDB":    AwrHubSourceTypeComanagedExacsPdb,
	"COMANAGED_EXACS_NONCDB": AwrHubSourceTypeComanagedExacsNoncdb,
	"UNDEFINED":              AwrHubSourceTypeUndefined,
}

var mappingAwrHubSourceTypeEnumLowerCase = map[string]AwrHubSourceTypeEnum{
	"adw_s":                  AwrHubSourceTypeAdwS,
	"atp_s":                  AwrHubSourceTypeAtpS,
	"adw_d":                  AwrHubSourceTypeAdwD,
	"atp_d":                  AwrHubSourceTypeAtpD,
	"external_pdb":           AwrHubSourceTypeExternalPdb,
	"external_noncdb":        AwrHubSourceTypeExternalNoncdb,
	"comanaged_vm_cdb":       AwrHubSourceTypeComanagedVmCdb,
	"comanaged_vm_pdb":       AwrHubSourceTypeComanagedVmPdb,
	"comanaged_vm_noncdb":    AwrHubSourceTypeComanagedVmNoncdb,
	"comanaged_bm_cdb":       AwrHubSourceTypeComanagedBmCdb,
	"comanaged_bm_pdb":       AwrHubSourceTypeComanagedBmPdb,
	"comanaged_bm_noncdb":    AwrHubSourceTypeComanagedBmNoncdb,
	"comanaged_exacs_cdb":    AwrHubSourceTypeComanagedExacsCdb,
	"comanaged_exacs_pdb":    AwrHubSourceTypeComanagedExacsPdb,
	"comanaged_exacs_noncdb": AwrHubSourceTypeComanagedExacsNoncdb,
	"undefined":              AwrHubSourceTypeUndefined,
}

// GetAwrHubSourceTypeEnumValues Enumerates the set of values for AwrHubSourceTypeEnum
func GetAwrHubSourceTypeEnumValues() []AwrHubSourceTypeEnum {
	values := make([]AwrHubSourceTypeEnum, 0)
	for _, v := range mappingAwrHubSourceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAwrHubSourceTypeEnumStringValues Enumerates the set of values in String for AwrHubSourceTypeEnum
func GetAwrHubSourceTypeEnumStringValues() []string {
	return []string{
		"ADW_S",
		"ATP_S",
		"ADW_D",
		"ATP_D",
		"EXTERNAL_PDB",
		"EXTERNAL_NONCDB",
		"COMANAGED_VM_CDB",
		"COMANAGED_VM_PDB",
		"COMANAGED_VM_NONCDB",
		"COMANAGED_BM_CDB",
		"COMANAGED_BM_PDB",
		"COMANAGED_BM_NONCDB",
		"COMANAGED_EXACS_CDB",
		"COMANAGED_EXACS_PDB",
		"COMANAGED_EXACS_NONCDB",
		"UNDEFINED",
	}
}

// GetMappingAwrHubSourceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAwrHubSourceTypeEnum(val string) (AwrHubSourceTypeEnum, bool) {
	enum, ok := mappingAwrHubSourceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
