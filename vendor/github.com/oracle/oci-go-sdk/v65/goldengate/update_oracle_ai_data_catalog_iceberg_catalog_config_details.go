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

// UpdateOracleAiDataCatalogIcebergCatalogConfigDetails The information to update an Oracle AI Data Catalog configuration of given type used in an Iceberg connection.
type UpdateOracleAiDataCatalogIcebergCatalogConfigDetails interface {
}

type updateoracleaidatacatalogicebergcatalogconfigdetails struct {
	JsonData          []byte
	CatalogConfigType string `json:"catalogConfigType"`
}

// UnmarshalJSON unmarshals json
func (m *updateoracleaidatacatalogicebergcatalogconfigdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerupdateoracleaidatacatalogicebergcatalogconfigdetails updateoracleaidatacatalogicebergcatalogconfigdetails
	s := struct {
		Model Unmarshalerupdateoracleaidatacatalogicebergcatalogconfigdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.CatalogConfigType = s.Model.CatalogConfigType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *updateoracleaidatacatalogicebergcatalogconfigdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.CatalogConfigType {
	case "DB_CREDENTIALS":
		mm := UpdateDbCredentialsOracleAiDataCatalogIcebergCatalogConfigDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OCI_GOLDENGATE":
		mm := UpdateOciGoldenGateOracleAiDataCatalogIcebergCatalogConfigDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for UpdateOracleAiDataCatalogIcebergCatalogConfigDetails: %s.", m.CatalogConfigType)
		return *m, nil
	}
}

func (m updateoracleaidatacatalogicebergcatalogconfigdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m updateoracleaidatacatalogicebergcatalogconfigdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
