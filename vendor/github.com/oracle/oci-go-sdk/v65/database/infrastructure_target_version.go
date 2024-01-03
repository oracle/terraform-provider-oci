// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// InfrastructureTargetVersion Infrastructure target version details.
type InfrastructureTargetVersion struct {

	// The history entry of the target system software version for the database server patching operation.
	TargetDbVersionHistoryEntry []string `mandatory:"true" json:"targetDbVersionHistoryEntry"`

	// The history entry of the target storage cell system software version for the storage cell patching operation.
	TargetStorageVersionHistoryEntry []string `mandatory:"true" json:"targetStorageVersionHistoryEntry"`

	// The resource type of the target Exadata infrastructure resource that will receive the system software update.
	TargetResourceType InfrastructureTargetVersionTargetResourceTypeEnum `mandatory:"false" json:"targetResourceType,omitempty"`

	// The OCID of the target Exadata Infrastructure resource that will receive the maintenance update.
	TargetResourceId *string `mandatory:"false" json:"targetResourceId"`
}

func (m InfrastructureTargetVersion) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m InfrastructureTargetVersion) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingInfrastructureTargetVersionTargetResourceTypeEnum(string(m.TargetResourceType)); !ok && m.TargetResourceType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TargetResourceType: %s. Supported values are: %s.", m.TargetResourceType, strings.Join(GetInfrastructureTargetVersionTargetResourceTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// InfrastructureTargetVersionTargetResourceTypeEnum Enum with underlying type: string
type InfrastructureTargetVersionTargetResourceTypeEnum string

// Set of constants representing the allowable values for InfrastructureTargetVersionTargetResourceTypeEnum
const (
	InfrastructureTargetVersionTargetResourceTypeExadataDbSystem            InfrastructureTargetVersionTargetResourceTypeEnum = "EXADATA_DB_SYSTEM"
	InfrastructureTargetVersionTargetResourceTypeCloudExadataInfrastructure InfrastructureTargetVersionTargetResourceTypeEnum = "CLOUD_EXADATA_INFRASTRUCTURE"
	InfrastructureTargetVersionTargetResourceTypeExaccInfrastructure        InfrastructureTargetVersionTargetResourceTypeEnum = "EXACC_INFRASTRUCTURE"
)

var mappingInfrastructureTargetVersionTargetResourceTypeEnum = map[string]InfrastructureTargetVersionTargetResourceTypeEnum{
	"EXADATA_DB_SYSTEM":            InfrastructureTargetVersionTargetResourceTypeExadataDbSystem,
	"CLOUD_EXADATA_INFRASTRUCTURE": InfrastructureTargetVersionTargetResourceTypeCloudExadataInfrastructure,
	"EXACC_INFRASTRUCTURE":         InfrastructureTargetVersionTargetResourceTypeExaccInfrastructure,
}

var mappingInfrastructureTargetVersionTargetResourceTypeEnumLowerCase = map[string]InfrastructureTargetVersionTargetResourceTypeEnum{
	"exadata_db_system":            InfrastructureTargetVersionTargetResourceTypeExadataDbSystem,
	"cloud_exadata_infrastructure": InfrastructureTargetVersionTargetResourceTypeCloudExadataInfrastructure,
	"exacc_infrastructure":         InfrastructureTargetVersionTargetResourceTypeExaccInfrastructure,
}

// GetInfrastructureTargetVersionTargetResourceTypeEnumValues Enumerates the set of values for InfrastructureTargetVersionTargetResourceTypeEnum
func GetInfrastructureTargetVersionTargetResourceTypeEnumValues() []InfrastructureTargetVersionTargetResourceTypeEnum {
	values := make([]InfrastructureTargetVersionTargetResourceTypeEnum, 0)
	for _, v := range mappingInfrastructureTargetVersionTargetResourceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetInfrastructureTargetVersionTargetResourceTypeEnumStringValues Enumerates the set of values in String for InfrastructureTargetVersionTargetResourceTypeEnum
func GetInfrastructureTargetVersionTargetResourceTypeEnumStringValues() []string {
	return []string{
		"EXADATA_DB_SYSTEM",
		"CLOUD_EXADATA_INFRASTRUCTURE",
		"EXACC_INFRASTRUCTURE",
	}
}

// GetMappingInfrastructureTargetVersionTargetResourceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingInfrastructureTargetVersionTargetResourceTypeEnum(val string) (InfrastructureTargetVersionTargetResourceTypeEnum, bool) {
	enum, ok := mappingInfrastructureTargetVersionTargetResourceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
