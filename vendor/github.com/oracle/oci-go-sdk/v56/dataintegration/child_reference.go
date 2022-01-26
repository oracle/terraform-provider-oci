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

// ChildReference Child reference contains application configuration information.
type ChildReference struct {

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

	// The type of the reference object.
	Type ChildReferenceTypeEnum `mandatory:"false" json:"type,omitempty"`

	// The new reference object to use instead of the original reference. For example, this can be a data asset reference.
	TargetObject *interface{} `mandatory:"false" json:"targetObject"`

	// The aggregator key of the child reference object. For example, this can be a data asset key.
	AggregatorKey *string `mandatory:"false" json:"aggregatorKey"`

	// List of published objects where this is used.
	UsedBy []ReferenceUsedBy `mandatory:"false" json:"usedBy"`
}

func (m ChildReference) String() string {
	return common.PointerString(m)
}

// ChildReferenceTypeEnum Enum with underlying type: string
type ChildReferenceTypeEnum string

// Set of constants representing the allowable values for ChildReferenceTypeEnum
const (
	ChildReferenceTypeOracledbConnection            ChildReferenceTypeEnum = "ORACLEDB_CONNECTION"
	ChildReferenceTypeOracleObjectStorageConnection ChildReferenceTypeEnum = "ORACLE_OBJECT_STORAGE_CONNECTION"
	ChildReferenceTypeOracleAtpConnection           ChildReferenceTypeEnum = "ORACLE_ATP_CONNECTION"
	ChildReferenceTypeOracleAdwcConnection          ChildReferenceTypeEnum = "ORACLE_ADWC_CONNECTION"
	ChildReferenceTypeMysqlConnection               ChildReferenceTypeEnum = "MYSQL_CONNECTION"
	ChildReferenceTypeGenericJdbcConnection         ChildReferenceTypeEnum = "GENERIC_JDBC_CONNECTION"
)

var mappingChildReferenceType = map[string]ChildReferenceTypeEnum{
	"ORACLEDB_CONNECTION":              ChildReferenceTypeOracledbConnection,
	"ORACLE_OBJECT_STORAGE_CONNECTION": ChildReferenceTypeOracleObjectStorageConnection,
	"ORACLE_ATP_CONNECTION":            ChildReferenceTypeOracleAtpConnection,
	"ORACLE_ADWC_CONNECTION":           ChildReferenceTypeOracleAdwcConnection,
	"MYSQL_CONNECTION":                 ChildReferenceTypeMysqlConnection,
	"GENERIC_JDBC_CONNECTION":          ChildReferenceTypeGenericJdbcConnection,
}

// GetChildReferenceTypeEnumValues Enumerates the set of values for ChildReferenceTypeEnum
func GetChildReferenceTypeEnumValues() []ChildReferenceTypeEnum {
	values := make([]ChildReferenceTypeEnum, 0)
	for _, v := range mappingChildReferenceType {
		values = append(values, v)
	}
	return values
}
