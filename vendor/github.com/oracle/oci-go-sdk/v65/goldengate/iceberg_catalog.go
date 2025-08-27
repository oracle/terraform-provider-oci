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

// IcebergCatalog Represents the catalog of given type used in an Iceberg connection.
type IcebergCatalog interface {
}

type icebergcatalog struct {
	JsonData    []byte
	CatalogType string `json:"catalogType"`
}

// UnmarshalJSON unmarshals json
func (m *icebergcatalog) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalericebergcatalog icebergcatalog
	s := struct {
		Model Unmarshalericebergcatalog
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.CatalogType = s.Model.CatalogType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *icebergcatalog) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.CatalogType {
	case "GLUE":
		mm := GlueIcebergCatalog{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "POLARIS":
		mm := PolarisIcebergCatalog{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "REST":
		mm := RestIcebergCatalog{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "NESSIE":
		mm := NessieIcebergCatalog{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "HADOOP":
		mm := HadoopIcebergCatalog{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for IcebergCatalog: %s.", m.CatalogType)
		return *m, nil
	}
}

func (m icebergcatalog) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m icebergcatalog) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
