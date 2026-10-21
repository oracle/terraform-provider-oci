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

// UpdateSparkOciObjectStorageConfigDetails Configuration for updating an OCI Object Storage Spark artifact location.
type UpdateSparkOciObjectStorageConfigDetails interface {
}

type updatesparkociobjectstorageconfigdetails struct {
	JsonData           []byte
	AuthenticationType string `json:"authenticationType"`
}

// UnmarshalJSON unmarshals json
func (m *updatesparkociobjectstorageconfigdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerupdatesparkociobjectstorageconfigdetails updatesparkociobjectstorageconfigdetails
	s := struct {
		Model Unmarshalerupdatesparkociobjectstorageconfigdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.AuthenticationType = s.Model.AuthenticationType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *updatesparkociobjectstorageconfigdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.AuthenticationType {
	case "PRE_AUTHENTICATED_REQUEST":
		mm := UpdateSparkParOciObjectStorageConfigDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "CUSTOMER_SECRET_KEY":
		mm := UpdateSparkSecretKeyOciObjectStorageConfigDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "USER_PRINCIPAL":
		mm := UpdateSparkUserPrincipalOciObjectStorageConfigDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "RESOURCE_PRINCIPAL":
		mm := UpdateSparkResourcePrincipalOciObjectStorageConfigDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for UpdateSparkOciObjectStorageConfigDetails: %s.", m.AuthenticationType)
		return *m, nil
	}
}

func (m updatesparkociobjectstorageconfigdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m updatesparkociobjectstorageconfigdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
