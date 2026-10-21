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

// UpdateDatabricksEngineOciObjectStorageConfigDetails Configuration for updating an OCI Object Storage Databricks Engine artifact location.
type UpdateDatabricksEngineOciObjectStorageConfigDetails interface {
}

type updatedatabricksengineociobjectstorageconfigdetails struct {
	JsonData           []byte
	AuthenticationType string `json:"authenticationType"`
}

// UnmarshalJSON unmarshals json
func (m *updatedatabricksengineociobjectstorageconfigdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerupdatedatabricksengineociobjectstorageconfigdetails updatedatabricksengineociobjectstorageconfigdetails
	s := struct {
		Model Unmarshalerupdatedatabricksengineociobjectstorageconfigdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.AuthenticationType = s.Model.AuthenticationType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *updatedatabricksengineociobjectstorageconfigdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.AuthenticationType {
	case "CUSTOMER_SECRET_KEY":
		mm := UpdateDatabricksEngineSecretKeyOciObjectStorageConfigDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "USER_PRINCIPAL":
		mm := UpdateDatabricksEngineUserPrincipalOciObjectStorageConfigDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "PRE_AUTHENTICATED_REQUEST":
		mm := UpdateDatabricksEngineParOciObjectStorageConfigDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "RESOURCE_PRINCIPAL":
		mm := UpdateDatabricksEngineResourcePrincipalOciObjectStorageConfigDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for UpdateDatabricksEngineOciObjectStorageConfigDetails: %s.", m.AuthenticationType)
		return *m, nil
	}
}

func (m updatedatabricksengineociobjectstorageconfigdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m updatedatabricksengineociobjectstorageconfigdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
