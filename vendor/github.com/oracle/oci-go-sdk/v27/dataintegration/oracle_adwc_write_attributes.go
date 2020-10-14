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
	"github.com/oracle/oci-go-sdk/v27/common"
)

// OracleAdwcWriteAttributes Properties to configure when writing to Oracle Autonomous Data Warehouse Cloud.
type OracleAdwcWriteAttributes struct {

	// The type of the abstract write attribute.
	ModelType OracleAdwcWriteAttributesModelTypeEnum `mandatory:"true" json:"modelType"`

	// The bucket name for the attribute.
	BucketName *string `mandatory:"false" json:"bucketName"`

	// The file name for the attribute.
	StagingFileName *string `mandatory:"false" json:"stagingFileName"`

	StagingDataAsset DataAsset `mandatory:"false" json:"stagingDataAsset"`

	StagingConnection Connection `mandatory:"false" json:"stagingConnection"`
}

func (m OracleAdwcWriteAttributes) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *OracleAdwcWriteAttributes) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		BucketName        *string                                `json:"bucketName"`
		StagingFileName   *string                                `json:"stagingFileName"`
		StagingDataAsset  dataasset                              `json:"stagingDataAsset"`
		StagingConnection connection                             `json:"stagingConnection"`
		ModelType         OracleAdwcWriteAttributesModelTypeEnum `json:"modelType"`
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

// OracleAdwcWriteAttributesModelTypeEnum Enum with underlying type: string
type OracleAdwcWriteAttributesModelTypeEnum string

// Set of constants representing the allowable values for OracleAdwcWriteAttributesModelTypeEnum
const (
	OracleAdwcWriteAttributesModelTypeOraclewriteattribute     OracleAdwcWriteAttributesModelTypeEnum = "ORACLEWRITEATTRIBUTE"
	OracleAdwcWriteAttributesModelTypeOracleatpwriteattribute  OracleAdwcWriteAttributesModelTypeEnum = "ORACLEATPWRITEATTRIBUTE"
	OracleAdwcWriteAttributesModelTypeOracleadwcwriteattribute OracleAdwcWriteAttributesModelTypeEnum = "ORACLEADWCWRITEATTRIBUTE"
	OracleAdwcWriteAttributesModelTypeOracleWriteAttribute     OracleAdwcWriteAttributesModelTypeEnum = "ORACLE_WRITE_ATTRIBUTE"
	OracleAdwcWriteAttributesModelTypeOracleAtpWriteAttribute  OracleAdwcWriteAttributesModelTypeEnum = "ORACLE_ATP_WRITE_ATTRIBUTE"
	OracleAdwcWriteAttributesModelTypeOracleAdwcWriteAttribute OracleAdwcWriteAttributesModelTypeEnum = "ORACLE_ADWC_WRITE_ATTRIBUTE"
)

var mappingOracleAdwcWriteAttributesModelType = map[string]OracleAdwcWriteAttributesModelTypeEnum{
	"ORACLEWRITEATTRIBUTE":        OracleAdwcWriteAttributesModelTypeOraclewriteattribute,
	"ORACLEATPWRITEATTRIBUTE":     OracleAdwcWriteAttributesModelTypeOracleatpwriteattribute,
	"ORACLEADWCWRITEATTRIBUTE":    OracleAdwcWriteAttributesModelTypeOracleadwcwriteattribute,
	"ORACLE_WRITE_ATTRIBUTE":      OracleAdwcWriteAttributesModelTypeOracleWriteAttribute,
	"ORACLE_ATP_WRITE_ATTRIBUTE":  OracleAdwcWriteAttributesModelTypeOracleAtpWriteAttribute,
	"ORACLE_ADWC_WRITE_ATTRIBUTE": OracleAdwcWriteAttributesModelTypeOracleAdwcWriteAttribute,
}

// GetOracleAdwcWriteAttributesModelTypeEnumValues Enumerates the set of values for OracleAdwcWriteAttributesModelTypeEnum
func GetOracleAdwcWriteAttributesModelTypeEnumValues() []OracleAdwcWriteAttributesModelTypeEnum {
	values := make([]OracleAdwcWriteAttributesModelTypeEnum, 0)
	for _, v := range mappingOracleAdwcWriteAttributesModelType {
		values = append(values, v)
	}
	return values
}
