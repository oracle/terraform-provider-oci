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

// OracleAiDataCatalogIcebergCatalog Represents an Oracle AI Data Catalog catalog used in the Iceberg connection.
type OracleAiDataCatalogIcebergCatalog struct {
	CatalogConfig OracleAiDataCatalogIcebergCatalogConfig `mandatory:"true" json:"catalogConfig"`
}

func (m OracleAiDataCatalogIcebergCatalog) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OracleAiDataCatalogIcebergCatalog) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m OracleAiDataCatalogIcebergCatalog) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeOracleAiDataCatalogIcebergCatalog OracleAiDataCatalogIcebergCatalog
	s := struct {
		DiscriminatorParam string `json:"catalogType"`
		MarshalTypeOracleAiDataCatalogIcebergCatalog
	}{
		"OADC",
		(MarshalTypeOracleAiDataCatalogIcebergCatalog)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *OracleAiDataCatalogIcebergCatalog) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		CatalogConfig oracleaidatacatalogicebergcatalogconfig `json:"catalogConfig"`
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
		m.CatalogConfig = nn.(OracleAiDataCatalogIcebergCatalogConfig)
	} else {
		m.CatalogConfig = nil
	}

	return
}
