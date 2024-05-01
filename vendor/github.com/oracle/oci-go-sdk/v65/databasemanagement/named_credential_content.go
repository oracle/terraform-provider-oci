// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to monitor and manage resources such as
// Oracle Databases, MySQL Databases, and External Database Systems.
// For more information, see Database Management (https://docs.cloud.oracle.com/iaas/database-management/home.htm).
//

package databasemanagement

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// NamedCredentialContent The details of the named credential.
type NamedCredentialContent interface {
}

type namedcredentialcontent struct {
	JsonData       []byte
	CredentialType string `json:"credentialType"`
}

// UnmarshalJSON unmarshals json
func (m *namedcredentialcontent) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalernamedcredentialcontent namedcredentialcontent
	s := struct {
		Model Unmarshalernamedcredentialcontent
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.CredentialType = s.Model.CredentialType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *namedcredentialcontent) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.CredentialType {
	case "BASIC":
		mm := BasicNamedCredentialContent{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for NamedCredentialContent: %s.", m.CredentialType)
		return *m, nil
	}
}

func (m namedcredentialcontent) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m namedcredentialcontent) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
