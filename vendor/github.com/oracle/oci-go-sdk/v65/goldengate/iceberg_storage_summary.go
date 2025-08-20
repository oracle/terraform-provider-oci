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

// IcebergStorageSummary Summary of the storage of given type used in an Iceberg connection.
type IcebergStorageSummary interface {
}

type icebergstoragesummary struct {
	JsonData    []byte
	StorageType string `json:"storageType"`
}

// UnmarshalJSON unmarshals json
func (m *icebergstoragesummary) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalericebergstoragesummary icebergstoragesummary
	s := struct {
		Model Unmarshalericebergstoragesummary
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.StorageType = s.Model.StorageType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *icebergstoragesummary) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.StorageType {
	case "GOOGLE_CLOUD_STORAGE":
		mm := GoogleCloudStorageIcebergStorageSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "AZURE_DATA_LAKE_STORAGE":
		mm := AzureDataLakeStorageIcebergStorageSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "AMAZON_S3":
		mm := AmazonS3IcebergStorageSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for IcebergStorageSummary: %s.", m.StorageType)
		return *m, nil
	}
}

func (m icebergstoragesummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m icebergstoragesummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
