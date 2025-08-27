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

// UpdateIcebergStorageDetails The information to update a storage of given type used in an Iceberg connection.
type UpdateIcebergStorageDetails interface {
}

type updateicebergstoragedetails struct {
	JsonData    []byte
	StorageType string `json:"storageType"`
}

// UnmarshalJSON unmarshals json
func (m *updateicebergstoragedetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerupdateicebergstoragedetails updateicebergstoragedetails
	s := struct {
		Model Unmarshalerupdateicebergstoragedetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.StorageType = s.Model.StorageType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *updateicebergstoragedetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.StorageType {
	case "AMAZON_S3":
		mm := UpdateAmazonS3IcebergStorageDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "GOOGLE_CLOUD_STORAGE":
		mm := UpdateGoogleCloudStorageIcebergStorageDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "AZURE_DATA_LAKE_STORAGE":
		mm := UpdateAzureDataLakeStorageIcebergStorageDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for UpdateIcebergStorageDetails: %s.", m.StorageType)
		return *m, nil
	}
}

func (m updateicebergstoragedetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m updateicebergstoragedetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
