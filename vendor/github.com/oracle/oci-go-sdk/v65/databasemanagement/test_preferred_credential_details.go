// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to monitor and manage resources such as
// Oracle Databases, MySQL Databases, and External Database Systems.
// For more information, see Database Management (https://docs.oracle.com/iaas/database-management/home.htm).
//

package databasemanagement

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// TestPreferredCredentialDetails The status of the preferred credential test. The status is 'SUCCEEDED' if the preferred credential is working else the status is 'FAILED'.
type TestPreferredCredentialDetails interface {
}

type testpreferredcredentialdetails struct {
	JsonData []byte
	Type     string `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *testpreferredcredentialdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalertestpreferredcredentialdetails testpreferredcredentialdetails
	s := struct {
		Model Unmarshalertestpreferredcredentialdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *testpreferredcredentialdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "NAMED_CREDENTIAL":
		mm := TestNamedPreferredCredentialDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "BASIC":
		mm := TestBasicPreferredCredentialDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for TestPreferredCredentialDetails: %s.", m.Type)
		return *m, nil
	}
}

func (m testpreferredcredentialdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m testpreferredcredentialdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
