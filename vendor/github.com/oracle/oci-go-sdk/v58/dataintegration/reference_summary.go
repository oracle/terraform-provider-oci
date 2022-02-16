// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Integration API
//
// Use the Data Integration API to organize your data integration projects, create data flows, pipelines and tasks, and then publish, schedule, and run tasks that extract, transform, and load data. For more information, see Data Integration (https://docs.oracle.com/iaas/data-integration/home.htm).
//

package dataintegration

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// ReferenceSummary This is the reference summary information.
type ReferenceSummary struct {

	// The reference's key, key of the object that is being used by a published object or its dependents.
	Key *string `mandatory:"false" json:"key"`

	// The name of reference object.
	Name *string `mandatory:"false" json:"name"`

	// The identifier of reference object.
	Identifier *string `mandatory:"false" json:"identifier"`

	// The identifier path of reference object.
	IdentifierPath *string `mandatory:"false" json:"identifierPath"`

	// The description of reference object.
	Description *string `mandatory:"false" json:"description"`

	// The type of reference object.
	Type ReferenceSummaryTypeEnum `mandatory:"false" json:"type,omitempty"`

	// The target object referenced. References are made to data assets and child references are made to connections. The type defining this reference is mentioned in the property type.
	TargetObject *interface{} `mandatory:"false" json:"targetObject"`

	// The aggregator of reference object.
	AggregatorKey *string `mandatory:"false" json:"aggregatorKey"`

	// List of published objects where this is used.
	UsedBy []ReferenceUsedBy `mandatory:"false" json:"usedBy"`

	// List of references that are dependent on this reference.
	ChildReferences []ChildReference `mandatory:"false" json:"childReferences"`
}

func (m ReferenceSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ReferenceSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingReferenceSummaryTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetReferenceSummaryTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ReferenceSummaryTypeEnum Enum with underlying type: string
type ReferenceSummaryTypeEnum string

// Set of constants representing the allowable values for ReferenceSummaryTypeEnum
const (
	ReferenceSummaryTypeOracleDataAsset              ReferenceSummaryTypeEnum = "ORACLE_DATA_ASSET"
	ReferenceSummaryTypeOracleObjectStorageDataAsset ReferenceSummaryTypeEnum = "ORACLE_OBJECT_STORAGE_DATA_ASSET"
	ReferenceSummaryTypeOracleAtpDataAsset           ReferenceSummaryTypeEnum = "ORACLE_ATP_DATA_ASSET"
	ReferenceSummaryTypeOracleAdwcDataAsset          ReferenceSummaryTypeEnum = "ORACLE_ADWC_DATA_ASSET"
	ReferenceSummaryTypeMysqlDataAsset               ReferenceSummaryTypeEnum = "MYSQL_DATA_ASSET"
	ReferenceSummaryTypeGenericJdbcDataAsset         ReferenceSummaryTypeEnum = "GENERIC_JDBC_DATA_ASSET"
)

var mappingReferenceSummaryTypeEnum = map[string]ReferenceSummaryTypeEnum{
	"ORACLE_DATA_ASSET":                ReferenceSummaryTypeOracleDataAsset,
	"ORACLE_OBJECT_STORAGE_DATA_ASSET": ReferenceSummaryTypeOracleObjectStorageDataAsset,
	"ORACLE_ATP_DATA_ASSET":            ReferenceSummaryTypeOracleAtpDataAsset,
	"ORACLE_ADWC_DATA_ASSET":           ReferenceSummaryTypeOracleAdwcDataAsset,
	"MYSQL_DATA_ASSET":                 ReferenceSummaryTypeMysqlDataAsset,
	"GENERIC_JDBC_DATA_ASSET":          ReferenceSummaryTypeGenericJdbcDataAsset,
}

// GetReferenceSummaryTypeEnumValues Enumerates the set of values for ReferenceSummaryTypeEnum
func GetReferenceSummaryTypeEnumValues() []ReferenceSummaryTypeEnum {
	values := make([]ReferenceSummaryTypeEnum, 0)
	for _, v := range mappingReferenceSummaryTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetReferenceSummaryTypeEnumStringValues Enumerates the set of values in String for ReferenceSummaryTypeEnum
func GetReferenceSummaryTypeEnumStringValues() []string {
	return []string{
		"ORACLE_DATA_ASSET",
		"ORACLE_OBJECT_STORAGE_DATA_ASSET",
		"ORACLE_ATP_DATA_ASSET",
		"ORACLE_ADWC_DATA_ASSET",
		"MYSQL_DATA_ASSET",
		"GENERIC_JDBC_DATA_ASSET",
	}
}

// GetMappingReferenceSummaryTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingReferenceSummaryTypeEnum(val string) (ReferenceSummaryTypeEnum, bool) {
	mappingReferenceSummaryTypeEnumIgnoreCase := make(map[string]ReferenceSummaryTypeEnum)
	for k, v := range mappingReferenceSummaryTypeEnum {
		mappingReferenceSummaryTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingReferenceSummaryTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
