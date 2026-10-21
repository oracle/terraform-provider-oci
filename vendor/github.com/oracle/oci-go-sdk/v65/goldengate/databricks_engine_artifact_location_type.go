// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// GoldenGate API
//
// Use the Oracle Cloud Infrastructure GoldenGate APIs to perform data replication operations.
//

package goldengate

import (
	"strings"
)

// DatabricksEngineArtifactLocationTypeEnum Enum with underlying type: string
type DatabricksEngineArtifactLocationTypeEnum string

// Set of constants representing the allowable values for DatabricksEngineArtifactLocationTypeEnum
const (
	DatabricksEngineArtifactLocationTypeUcVolume             DatabricksEngineArtifactLocationTypeEnum = "UC_VOLUME"
	DatabricksEngineArtifactLocationTypeOciObjectStorage     DatabricksEngineArtifactLocationTypeEnum = "OCI_OBJECT_STORAGE"
	DatabricksEngineArtifactLocationTypeAzureDateLakeStorage DatabricksEngineArtifactLocationTypeEnum = "AZURE_DATE_LAKE_STORAGE"
)

var mappingDatabricksEngineArtifactLocationTypeEnum = map[string]DatabricksEngineArtifactLocationTypeEnum{
	"UC_VOLUME":               DatabricksEngineArtifactLocationTypeUcVolume,
	"OCI_OBJECT_STORAGE":      DatabricksEngineArtifactLocationTypeOciObjectStorage,
	"AZURE_DATE_LAKE_STORAGE": DatabricksEngineArtifactLocationTypeAzureDateLakeStorage,
}

var mappingDatabricksEngineArtifactLocationTypeEnumLowerCase = map[string]DatabricksEngineArtifactLocationTypeEnum{
	"uc_volume":               DatabricksEngineArtifactLocationTypeUcVolume,
	"oci_object_storage":      DatabricksEngineArtifactLocationTypeOciObjectStorage,
	"azure_date_lake_storage": DatabricksEngineArtifactLocationTypeAzureDateLakeStorage,
}

// GetDatabricksEngineArtifactLocationTypeEnumValues Enumerates the set of values for DatabricksEngineArtifactLocationTypeEnum
func GetDatabricksEngineArtifactLocationTypeEnumValues() []DatabricksEngineArtifactLocationTypeEnum {
	values := make([]DatabricksEngineArtifactLocationTypeEnum, 0)
	for _, v := range mappingDatabricksEngineArtifactLocationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDatabricksEngineArtifactLocationTypeEnumStringValues Enumerates the set of values in String for DatabricksEngineArtifactLocationTypeEnum
func GetDatabricksEngineArtifactLocationTypeEnumStringValues() []string {
	return []string{
		"UC_VOLUME",
		"OCI_OBJECT_STORAGE",
		"AZURE_DATE_LAKE_STORAGE",
	}
}

// GetMappingDatabricksEngineArtifactLocationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDatabricksEngineArtifactLocationTypeEnum(val string) (DatabricksEngineArtifactLocationTypeEnum, bool) {
	enum, ok := mappingDatabricksEngineArtifactLocationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
