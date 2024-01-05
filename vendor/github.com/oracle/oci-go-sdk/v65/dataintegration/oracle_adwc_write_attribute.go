// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Integration API
//
// Use the Data Integration API to organize your data integration projects, create data flows, pipelines and tasks, and then publish, schedule, and run tasks that extract, transform, and load data. For more information, see Data Integration (https://docs.oracle.com/iaas/data-integration/home.htm).
//

package dataintegration

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// OracleAdwcWriteAttribute Properties to configure writing to Oracle Autonomous Data Warehouse Cloud.
type OracleAdwcWriteAttribute struct {

	// The bucket name for the attribute.
	BucketName *string `mandatory:"false" json:"bucketName"`

	// The file name for the attribute.
	StagingFileName *string `mandatory:"false" json:"stagingFileName"`

	StagingDataAsset DataAsset `mandatory:"false" json:"stagingDataAsset"`

	StagingConnection Connection `mandatory:"false" json:"stagingConnection"`
}

func (m OracleAdwcWriteAttribute) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OracleAdwcWriteAttribute) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m OracleAdwcWriteAttribute) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeOracleAdwcWriteAttribute OracleAdwcWriteAttribute
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeOracleAdwcWriteAttribute
	}{
		"ORACLEADWCWRITEATTRIBUTE",
		(MarshalTypeOracleAdwcWriteAttribute)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *OracleAdwcWriteAttribute) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		BucketName        *string    `json:"bucketName"`
		StagingFileName   *string    `json:"stagingFileName"`
		StagingDataAsset  dataasset  `json:"stagingDataAsset"`
		StagingConnection connection `json:"stagingConnection"`
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

	return
}
