// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Integration API
//
// Use the Data Integration Service APIs to perform common extract, load, and transform (ETL) tasks.
//

package dataintegration

import (
	"github.com/oracle/oci-go-sdk/v38/common"
)

// OracleReadAttributes Properties to configure reading from an Oracle Database.
type OracleReadAttributes struct {

	// The type of the abstract read attribute.
	ModelType OracleReadAttributesModelTypeEnum `mandatory:"true" json:"modelType"`

	// The fetch size for reading.
	FetchSize *int `mandatory:"false" json:"fetchSize"`
}

func (m OracleReadAttributes) String() string {
	return common.PointerString(m)
}

// OracleReadAttributesModelTypeEnum Enum with underlying type: string
type OracleReadAttributesModelTypeEnum string

// Set of constants representing the allowable values for OracleReadAttributesModelTypeEnum
const (
	OracleReadAttributesModelTypeOraclereadattribute OracleReadAttributesModelTypeEnum = "ORACLEREADATTRIBUTE"
	OracleReadAttributesModelTypeOracleReadAttribute OracleReadAttributesModelTypeEnum = "ORACLE_READ_ATTRIBUTE"
)

var mappingOracleReadAttributesModelType = map[string]OracleReadAttributesModelTypeEnum{
	"ORACLEREADATTRIBUTE":   OracleReadAttributesModelTypeOraclereadattribute,
	"ORACLE_READ_ATTRIBUTE": OracleReadAttributesModelTypeOracleReadAttribute,
}

// GetOracleReadAttributesModelTypeEnumValues Enumerates the set of values for OracleReadAttributesModelTypeEnum
func GetOracleReadAttributesModelTypeEnumValues() []OracleReadAttributesModelTypeEnum {
	values := make([]OracleReadAttributesModelTypeEnum, 0)
	for _, v := range mappingOracleReadAttributesModelType {
		values = append(values, v)
	}
	return values
}
