// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// GoldenGate API
//
// Use the Oracle Cloud Infrastructure GoldenGate APIs to perform data replication operations.
//

package goldengate

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateOracleAiDataCatalogIcebergCatalogDetails The information about a new Oracle AI Data Catalog catalog used in the Iceberg connection.
type CreateOracleAiDataCatalogIcebergCatalogDetails struct {
	CatalogConfig CreateOracleAiDataCatalogIcebergCatalogConfigDetails `mandatory:"true" json:"catalogConfig"`
}

func (m CreateOracleAiDataCatalogIcebergCatalogDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateOracleAiDataCatalogIcebergCatalogDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateOracleAiDataCatalogIcebergCatalogDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateOracleAiDataCatalogIcebergCatalogDetails CreateOracleAiDataCatalogIcebergCatalogDetails
	s := struct {
		DiscriminatorParam string `json:"catalogType"`
		MarshalTypeCreateOracleAiDataCatalogIcebergCatalogDetails
	}{
		"OADC",
		(MarshalTypeCreateOracleAiDataCatalogIcebergCatalogDetails)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *CreateOracleAiDataCatalogIcebergCatalogDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		CatalogConfig createoracleaidatacatalogicebergcatalogconfigdetails `json:"catalogConfig"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	nn, e = model.CatalogConfig.UnmarshalPolymorphicJSON(model.CatalogConfig.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.CatalogConfig = nn.(CreateOracleAiDataCatalogIcebergCatalogConfigDetails)
	} else {
		m.CatalogConfig = nil
	}

	return
}
