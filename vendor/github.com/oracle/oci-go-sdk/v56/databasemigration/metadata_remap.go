// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Migration API
//
// Use the Oracle Cloud Infrastructure Database Migration APIs to perform database migration operations.
//

package databasemigration

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// MetadataRemap Defines remapping to be applied to objects as they are processed.
// Refer to METADATA_REMAP Procedure  (https://docs.oracle.com/en/database/oracle/oracle-database/19/arpls/DBMS_DATAPUMP.html#GUID-0FC32790-91E6-4781-87A3-229DE024CB3D)
type MetadataRemap struct {

	// Type of remap. Refer to METADATA_REMAP Procedure  (https://docs.oracle.com/en/database/oracle/oracle-database/19/arpls/DBMS_DATAPUMP.html#GUID-0FC32790-91E6-4781-87A3-229DE024CB3D)
	Type MetadataRemapTypeEnum `mandatory:"true" json:"type"`

	// Specifies the value which needs to be reset.
	OldValue *string `mandatory:"true" json:"oldValue"`

	// Specifies the new value that oldValue should be translated into.
	NewValue *string `mandatory:"true" json:"newValue"`
}

func (m MetadataRemap) String() string {
	return common.PointerString(m)
}

// MetadataRemapTypeEnum Enum with underlying type: string
type MetadataRemapTypeEnum string

// Set of constants representing the allowable values for MetadataRemapTypeEnum
const (
	MetadataRemapTypeSchema     MetadataRemapTypeEnum = "SCHEMA"
	MetadataRemapTypeTablespace MetadataRemapTypeEnum = "TABLESPACE"
	MetadataRemapTypeDatafile   MetadataRemapTypeEnum = "DATAFILE"
	MetadataRemapTypeTable      MetadataRemapTypeEnum = "TABLE"
)

var mappingMetadataRemapType = map[string]MetadataRemapTypeEnum{
	"SCHEMA":     MetadataRemapTypeSchema,
	"TABLESPACE": MetadataRemapTypeTablespace,
	"DATAFILE":   MetadataRemapTypeDatafile,
	"TABLE":      MetadataRemapTypeTable,
}

// GetMetadataRemapTypeEnumValues Enumerates the set of values for MetadataRemapTypeEnum
func GetMetadataRemapTypeEnumValues() []MetadataRemapTypeEnum {
	values := make([]MetadataRemapTypeEnum, 0)
	for _, v := range mappingMetadataRemapType {
		values = append(values, v)
	}
	return values
}
