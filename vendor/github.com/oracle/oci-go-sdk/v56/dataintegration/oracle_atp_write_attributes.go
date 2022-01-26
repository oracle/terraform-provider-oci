// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Integration API
//
// Use the Data Integration API to organize your data integration projects, create data flows, pipelines and tasks, and then publish, schedule, and run tasks that extract, transform, and load data. For more information, see Data Integration (https://docs.oracle.com/iaas/data-integration/home.htm).
//

package dataintegration

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v56/common"
)

// OracleAtpWriteAttributes Properties to configure when writing to Oracle Autonomous Data Warehouse Cloud.
type OracleAtpWriteAttributes struct {
	BucketSchema *Schema `mandatory:"false" json:"bucketSchema"`

	// The file name for the attribute.
	StagingFileName *string `mandatory:"false" json:"stagingFileName"`

	StagingDataAsset DataAsset `mandatory:"false" json:"stagingDataAsset"`

	StagingConnection Connection `mandatory:"false" json:"stagingConnection"`
}

func (m OracleAtpWriteAttributes) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m OracleAtpWriteAttributes) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeOracleAtpWriteAttributes OracleAtpWriteAttributes
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeOracleAtpWriteAttributes
	}{
		"ORACLE_ATP_WRITE_ATTRIBUTE",
		(MarshalTypeOracleAtpWriteAttributes)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *OracleAtpWriteAttributes) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		BucketSchema      *Schema    `json:"bucketSchema"`
		StagingFileName   *string    `json:"stagingFileName"`
		StagingDataAsset  dataasset  `json:"stagingDataAsset"`
		StagingConnection connection `json:"stagingConnection"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.BucketSchema = model.BucketSchema

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

	return
}
