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

var mappingReferenceSummaryType = map[string]ReferenceSummaryTypeEnum{
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
	for _, v := range mappingReferenceSummaryType {
		values = append(values, v)
	}
	return values
}
