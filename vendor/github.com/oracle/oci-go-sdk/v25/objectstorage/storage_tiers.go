// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Object Storage Service API
//
// Common set of Object Storage and Archive Storage APIs for managing buckets, objects, and related resources.
// For more information, see Overview of Object Storage (https://docs.cloud.oracle.com/Content/Object/Concepts/objectstorageoverview.htm) and
// Overview of Archive Storage (https://docs.cloud.oracle.com/Content/Archive/Concepts/archivestorageoverview.htm).
//

package objectstorage

// StorageTiersEnum Enum with underlying type: string
type StorageTiersEnum string

// Set of constants representing the allowable values for StorageTiersEnum
const (
	StorageTiersStandard         StorageTiersEnum = "Standard"
	StorageTiersInfrequentAccess StorageTiersEnum = "InfrequentAccess"
	StorageTiersArchive          StorageTiersEnum = "Archive"
)

var mappingStorageTiers = map[string]StorageTiersEnum{
	"Standard":         StorageTiersStandard,
	"InfrequentAccess": StorageTiersInfrequentAccess,
	"Archive":          StorageTiersArchive,
}

// GetStorageTiersEnumValues Enumerates the set of values for StorageTiersEnum
func GetStorageTiersEnumValues() []StorageTiersEnum {
	values := make([]StorageTiersEnum, 0)
	for _, v := range mappingStorageTiers {
		values = append(values, v)
	}
	return values
}
