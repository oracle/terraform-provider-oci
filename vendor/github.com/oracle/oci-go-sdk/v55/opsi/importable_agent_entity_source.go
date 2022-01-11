// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Operations Insights API
//
// Use the Operations Insights API to perform data extraction operations to obtain database
// resource utilization, performance statistics, and reference information. For more information,
// see About Oracle Cloud Infrastructure Operations Insights (https://docs.cloud.oracle.com/en-us/iaas/operations-insights/doc/operations-insights.html).
//

package opsi

// ImportableAgentEntitySourceEnum Enum with underlying type: string
type ImportableAgentEntitySourceEnum string

// Set of constants representing the allowable values for ImportableAgentEntitySourceEnum
const (
	ImportableAgentEntitySourceMacsManagedExternalHost ImportableAgentEntitySourceEnum = "MACS_MANAGED_EXTERNAL_HOST"
)

var mappingImportableAgentEntitySource = map[string]ImportableAgentEntitySourceEnum{
	"MACS_MANAGED_EXTERNAL_HOST": ImportableAgentEntitySourceMacsManagedExternalHost,
}

// GetImportableAgentEntitySourceEnumValues Enumerates the set of values for ImportableAgentEntitySourceEnum
func GetImportableAgentEntitySourceEnumValues() []ImportableAgentEntitySourceEnum {
	values := make([]ImportableAgentEntitySourceEnum, 0)
	for _, v := range mappingImportableAgentEntitySource {
		values = append(values, v)
	}
	return values
}
