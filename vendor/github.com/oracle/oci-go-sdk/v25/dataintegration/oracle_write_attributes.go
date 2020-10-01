// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Integration API
//
// Use the Data Integration Service APIs to perform common extract, load, and transform (ETL) tasks.
//

package dataintegration

import (
	"github.com/oracle/oci-go-sdk/v25/common"
)

// OracleWriteAttributes Properties to configure when writing to an Oracle Database.
type OracleWriteAttributes struct {

	// The type of the abstract write attribute.
	ModelType OracleWriteAttributesModelTypeEnum `mandatory:"true" json:"modelType"`

	// The batch size for writing.
	BatchSize *int `mandatory:"false" json:"batchSize"`

	// Specifies whether to truncate.
	IsTruncate *bool `mandatory:"false" json:"isTruncate"`

	// Specifies the isolation level.
	IsolationLevel *string `mandatory:"false" json:"isolationLevel"`
}

func (m OracleWriteAttributes) String() string {
	return common.PointerString(m)
}

// OracleWriteAttributesModelTypeEnum Enum with underlying type: string
type OracleWriteAttributesModelTypeEnum string

// Set of constants representing the allowable values for OracleWriteAttributesModelTypeEnum
const (
	OracleWriteAttributesModelTypeOraclewriteattribute     OracleWriteAttributesModelTypeEnum = "ORACLEWRITEATTRIBUTE"
	OracleWriteAttributesModelTypeOracleatpwriteattribute  OracleWriteAttributesModelTypeEnum = "ORACLEATPWRITEATTRIBUTE"
	OracleWriteAttributesModelTypeOracleadwcwriteattribute OracleWriteAttributesModelTypeEnum = "ORACLEADWCWRITEATTRIBUTE"
	OracleWriteAttributesModelTypeOracleWriteAttribute     OracleWriteAttributesModelTypeEnum = "ORACLE_WRITE_ATTRIBUTE"
	OracleWriteAttributesModelTypeOracleAtpWriteAttribute  OracleWriteAttributesModelTypeEnum = "ORACLE_ATP_WRITE_ATTRIBUTE"
	OracleWriteAttributesModelTypeOracleAdwcWriteAttribute OracleWriteAttributesModelTypeEnum = "ORACLE_ADWC_WRITE_ATTRIBUTE"
)

var mappingOracleWriteAttributesModelType = map[string]OracleWriteAttributesModelTypeEnum{
	"ORACLEWRITEATTRIBUTE":        OracleWriteAttributesModelTypeOraclewriteattribute,
	"ORACLEATPWRITEATTRIBUTE":     OracleWriteAttributesModelTypeOracleatpwriteattribute,
	"ORACLEADWCWRITEATTRIBUTE":    OracleWriteAttributesModelTypeOracleadwcwriteattribute,
	"ORACLE_WRITE_ATTRIBUTE":      OracleWriteAttributesModelTypeOracleWriteAttribute,
	"ORACLE_ATP_WRITE_ATTRIBUTE":  OracleWriteAttributesModelTypeOracleAtpWriteAttribute,
	"ORACLE_ADWC_WRITE_ATTRIBUTE": OracleWriteAttributesModelTypeOracleAdwcWriteAttribute,
}

// GetOracleWriteAttributesModelTypeEnumValues Enumerates the set of values for OracleWriteAttributesModelTypeEnum
func GetOracleWriteAttributesModelTypeEnumValues() []OracleWriteAttributesModelTypeEnum {
	values := make([]OracleWriteAttributesModelTypeEnum, 0)
	for _, v := range mappingOracleWriteAttributesModelType {
		values = append(values, v)
	}
	return values
}
