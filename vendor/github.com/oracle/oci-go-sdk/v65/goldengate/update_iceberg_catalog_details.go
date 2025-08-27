// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// UpdateIcebergCatalogDetails The information to update a catalog of given type used in an Iceberg connection.
type UpdateIcebergCatalogDetails interface {
}

type updateicebergcatalogdetails struct {
	JsonData    []byte
	CatalogType string `json:"catalogType"`
}

// UnmarshalJSON unmarshals json
func (m *updateicebergcatalogdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerupdateicebergcatalogdetails updateicebergcatalogdetails
	s := struct {
		Model Unmarshalerupdateicebergcatalogdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.CatalogType = s.Model.CatalogType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *updateicebergcatalogdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.CatalogType {
	case "REST":
		mm := UpdateRestIcebergCatalogDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "POLARIS":
		mm := UpdatePolarisIcebergCatalogDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "NESSIE":
		mm := UpdateNessieIcebergCatalogDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "HADOOP":
		mm := UpdateHadoopIcebergCatalogDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "GLUE":
		mm := UpdateGlueIcebergCatalogDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for UpdateIcebergCatalogDetails: %s.", m.CatalogType)
		return *m, nil
	}
}

func (m updateicebergcatalogdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m updateicebergcatalogdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
