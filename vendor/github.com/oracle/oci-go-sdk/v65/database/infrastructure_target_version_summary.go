// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
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

// InfrastructureTargetVersionSummary The target Exadata Infrastructure system software version for an infrastructure resource. Applies to Exadata Cloud@Customer and Exadata Cloud instances only.
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized,
// talk to an administrator. If you're an administrator who needs to write policies to give users access,
// see Getting Started with Policies (https://docs.cloud.oracle.com/Content/Identity/Concepts/policygetstarted.htm).
type InfrastructureTargetVersionSummary struct {

	// The history entry of the target system software version for the database server patching operation.
	TargetDbVersionHistoryEntry []string `mandatory:"true" json:"targetDbVersionHistoryEntry"`

	// The history entry of the target storage cell system software version for the storage cell patching operation.
	TargetStorageVersionHistoryEntry []string `mandatory:"true" json:"targetStorageVersionHistoryEntry"`

	// The resource type of the target Exadata infrastructure resource that will receive the system software update.
	TargetResourceType InfrastructureTargetVersionSummaryTargetResourceTypeEnum `mandatory:"false" json:"targetResourceType,omitempty"`

	// The OCID of the target Exadata Infrastructure resource that will receive the maintenance update.
	TargetResourceId *string `mandatory:"false" json:"targetResourceId"`
}

func (m InfrastructureTargetVersionSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m InfrastructureTargetVersionSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingInfrastructureTargetVersionSummaryTargetResourceTypeEnum(string(m.TargetResourceType)); !ok && m.TargetResourceType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TargetResourceType: %s. Supported values are: %s.", m.TargetResourceType, strings.Join(GetInfrastructureTargetVersionSummaryTargetResourceTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// InfrastructureTargetVersionSummaryTargetResourceTypeEnum Enum with underlying type: string
type InfrastructureTargetVersionSummaryTargetResourceTypeEnum string

// Set of constants representing the allowable values for InfrastructureTargetVersionSummaryTargetResourceTypeEnum
const (
	InfrastructureTargetVersionSummaryTargetResourceTypeExadataDbSystem            InfrastructureTargetVersionSummaryTargetResourceTypeEnum = "EXADATA_DB_SYSTEM"
	InfrastructureTargetVersionSummaryTargetResourceTypeCloudExadataInfrastructure InfrastructureTargetVersionSummaryTargetResourceTypeEnum = "CLOUD_EXADATA_INFRASTRUCTURE"
	InfrastructureTargetVersionSummaryTargetResourceTypeExaccInfrastructure        InfrastructureTargetVersionSummaryTargetResourceTypeEnum = "EXACC_INFRASTRUCTURE"
)

var mappingInfrastructureTargetVersionSummaryTargetResourceTypeEnum = map[string]InfrastructureTargetVersionSummaryTargetResourceTypeEnum{
	"EXADATA_DB_SYSTEM":            InfrastructureTargetVersionSummaryTargetResourceTypeExadataDbSystem,
	"CLOUD_EXADATA_INFRASTRUCTURE": InfrastructureTargetVersionSummaryTargetResourceTypeCloudExadataInfrastructure,
	"EXACC_INFRASTRUCTURE":         InfrastructureTargetVersionSummaryTargetResourceTypeExaccInfrastructure,
}

var mappingInfrastructureTargetVersionSummaryTargetResourceTypeEnumLowerCase = map[string]InfrastructureTargetVersionSummaryTargetResourceTypeEnum{
	"exadata_db_system":            InfrastructureTargetVersionSummaryTargetResourceTypeExadataDbSystem,
	"cloud_exadata_infrastructure": InfrastructureTargetVersionSummaryTargetResourceTypeCloudExadataInfrastructure,
	"exacc_infrastructure":         InfrastructureTargetVersionSummaryTargetResourceTypeExaccInfrastructure,
}

// GetInfrastructureTargetVersionSummaryTargetResourceTypeEnumValues Enumerates the set of values for InfrastructureTargetVersionSummaryTargetResourceTypeEnum
func GetInfrastructureTargetVersionSummaryTargetResourceTypeEnumValues() []InfrastructureTargetVersionSummaryTargetResourceTypeEnum {
	values := make([]InfrastructureTargetVersionSummaryTargetResourceTypeEnum, 0)
	for _, v := range mappingInfrastructureTargetVersionSummaryTargetResourceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetInfrastructureTargetVersionSummaryTargetResourceTypeEnumStringValues Enumerates the set of values in String for InfrastructureTargetVersionSummaryTargetResourceTypeEnum
func GetInfrastructureTargetVersionSummaryTargetResourceTypeEnumStringValues() []string {
	return []string{
		"EXADATA_DB_SYSTEM",
		"CLOUD_EXADATA_INFRASTRUCTURE",
		"EXACC_INFRASTRUCTURE",
	}
}

// GetMappingInfrastructureTargetVersionSummaryTargetResourceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingInfrastructureTargetVersionSummaryTargetResourceTypeEnum(val string) (InfrastructureTargetVersionSummaryTargetResourceTypeEnum, bool) {
	enum, ok := mappingInfrastructureTargetVersionSummaryTargetResourceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
