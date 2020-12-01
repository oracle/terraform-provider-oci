// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Integration API
//
// Use the Data Integration Service APIs to perform common extract, load, and transform (ETL) tasks.
//

package dataintegration

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v30/common"
)

// OracleAtpWriteAttributes Properties to configure when writing to Oracle Autonomous Data Warehouse Cloud.
type OracleAtpWriteAttributes struct {

	// The type of the abstract write attribute.
	ModelType OracleAtpWriteAttributesModelTypeEnum `mandatory:"true" json:"modelType"`

	// The bucket name for the attribute.
	BucketName *string `mandatory:"false" json:"bucketName"`

	// The file name for the attribute.
	StagingFileName *string `mandatory:"false" json:"stagingFileName"`

	StagingDataAsset DataAsset `mandatory:"false" json:"stagingDataAsset"`

	StagingConnection Connection `mandatory:"false" json:"stagingConnection"`
}

func (m OracleAtpWriteAttributes) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *OracleAtpWriteAttributes) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		BucketName        *string                               `json:"bucketName"`
		StagingFileName   *string                               `json:"stagingFileName"`
		StagingDataAsset  dataasset                             `json:"stagingDataAsset"`
		StagingConnection connection                            `json:"stagingConnection"`
		ModelType         OracleAtpWriteAttributesModelTypeEnum `json:"modelType"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.BucketName = model.BucketName

	m.StagingFileName = model.StagingFileName

	nn, e = model.StagingDataAsset.UnmarshalPolymorphicJSON(model.StagingDataAsset.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.StagingDataAsset = nn.(DataAsset)
	} else {
		m.StagingDataAsset = nil
	}

	nn, e = model.StagingConnection.UnmarshalPolymorphicJSON(model.StagingConnection.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.StagingConnection = nn.(Connection)
	} else {
		m.StagingConnection = nil
	}

	m.ModelType = model.ModelType

	return
}

// OracleAtpWriteAttributesModelTypeEnum Enum with underlying type: string
type OracleAtpWriteAttributesModelTypeEnum string

// Set of constants representing the allowable values for OracleAtpWriteAttributesModelTypeEnum
const (
	OracleAtpWriteAttributesModelTypeOraclewriteattribute     OracleAtpWriteAttributesModelTypeEnum = "ORACLEWRITEATTRIBUTE"
	OracleAtpWriteAttributesModelTypeOracleatpwriteattribute  OracleAtpWriteAttributesModelTypeEnum = "ORACLEATPWRITEATTRIBUTE"
	OracleAtpWriteAttributesModelTypeOracleadwcwriteattribute OracleAtpWriteAttributesModelTypeEnum = "ORACLEADWCWRITEATTRIBUTE"
	OracleAtpWriteAttributesModelTypeOracleWriteAttribute     OracleAtpWriteAttributesModelTypeEnum = "ORACLE_WRITE_ATTRIBUTE"
	OracleAtpWriteAttributesModelTypeOracleAtpWriteAttribute  OracleAtpWriteAttributesModelTypeEnum = "ORACLE_ATP_WRITE_ATTRIBUTE"
	OracleAtpWriteAttributesModelTypeOracleAdwcWriteAttribute OracleAtpWriteAttributesModelTypeEnum = "ORACLE_ADWC_WRITE_ATTRIBUTE"
)

var mappingOracleAtpWriteAttributesModelType = map[string]OracleAtpWriteAttributesModelTypeEnum{
	"ORACLEWRITEATTRIBUTE":        OracleAtpWriteAttributesModelTypeOraclewriteattribute,
	"ORACLEATPWRITEATTRIBUTE":     OracleAtpWriteAttributesModelTypeOracleatpwriteattribute,
	"ORACLEADWCWRITEATTRIBUTE":    OracleAtpWriteAttributesModelTypeOracleadwcwriteattribute,
	"ORACLE_WRITE_ATTRIBUTE":      OracleAtpWriteAttributesModelTypeOracleWriteAttribute,
	"ORACLE_ATP_WRITE_ATTRIBUTE":  OracleAtpWriteAttributesModelTypeOracleAtpWriteAttribute,
	"ORACLE_ADWC_WRITE_ATTRIBUTE": OracleAtpWriteAttributesModelTypeOracleAdwcWriteAttribute,
}

// GetOracleAtpWriteAttributesModelTypeEnumValues Enumerates the set of values for OracleAtpWriteAttributesModelTypeEnum
func GetOracleAtpWriteAttributesModelTypeEnumValues() []OracleAtpWriteAttributesModelTypeEnum {
	values := make([]OracleAtpWriteAttributesModelTypeEnum, 0)
	for _, v := range mappingOracleAtpWriteAttributesModelType {
		values = append(values, v)
	}
	return values
}
