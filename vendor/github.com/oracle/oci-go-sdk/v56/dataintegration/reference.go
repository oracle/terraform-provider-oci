// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Integration API
//
// Use the Data Integration API to organize your data integration projects, create data flows, pipelines and tasks, and then publish, schedule, and run tasks that extract, transform, and load data. For more information, see Data Integration (https://docs.oracle.com/iaas/data-integration/home.htm).
//

package dataintegration

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// Reference Reference contains application configuration information.
type Reference struct {

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
	Type ReferenceTypeEnum `mandatory:"false" json:"type,omitempty"`

	// The new reference object to use instead of the original reference. For example, this can be a data asset reference.
	TargetObject *interface{} `mandatory:"false" json:"targetObject"`

	// The application key of the reference object.
	ApplicationKey *string `mandatory:"false" json:"applicationKey"`

	// List of published objects where this is used.
	UsedBy []ReferenceUsedBy `mandatory:"false" json:"usedBy"`

	// List of references that are dependent on this reference.
	ChildReferences []ChildReference `mandatory:"false" json:"childReferences"`
}

func (m Reference) String() string {
	return common.PointerString(m)
}

// ReferenceTypeEnum Enum with underlying type: string
type ReferenceTypeEnum string

// Set of constants representing the allowable values for ReferenceTypeEnum
const (
	ReferenceTypeOracleDataAsset              ReferenceTypeEnum = "ORACLE_DATA_ASSET"
	ReferenceTypeOracleObjectStorageDataAsset ReferenceTypeEnum = "ORACLE_OBJECT_STORAGE_DATA_ASSET"
	ReferenceTypeOracleAtpDataAsset           ReferenceTypeEnum = "ORACLE_ATP_DATA_ASSET"
	ReferenceTypeOracleAdwcDataAsset          ReferenceTypeEnum = "ORACLE_ADWC_DATA_ASSET"
	ReferenceTypeMysqlDataAsset               ReferenceTypeEnum = "MYSQL_DATA_ASSET"
	ReferenceTypeGenericJdbcDataAsset         ReferenceTypeEnum = "GENERIC_JDBC_DATA_ASSET"
)

var mappingReferenceType = map[string]ReferenceTypeEnum{
	"ORACLE_DATA_ASSET":                ReferenceTypeOracleDataAsset,
	"ORACLE_OBJECT_STORAGE_DATA_ASSET": ReferenceTypeOracleObjectStorageDataAsset,
	"ORACLE_ATP_DATA_ASSET":            ReferenceTypeOracleAtpDataAsset,
	"ORACLE_ADWC_DATA_ASSET":           ReferenceTypeOracleAdwcDataAsset,
	"MYSQL_DATA_ASSET":                 ReferenceTypeMysqlDataAsset,
	"GENERIC_JDBC_DATA_ASSET":          ReferenceTypeGenericJdbcDataAsset,
}

// GetReferenceTypeEnumValues Enumerates the set of values for ReferenceTypeEnum
func GetReferenceTypeEnumValues() []ReferenceTypeEnum {
	values := make([]ReferenceTypeEnum, 0)
	for _, v := range mappingReferenceType {
		values = append(values, v)
	}
	return values
}
