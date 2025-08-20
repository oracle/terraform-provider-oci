// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Tools
//
// Use the Database Tools API to manage connections, private endpoints, and work requests in the Database Tools service.
//

package databasetools

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ValidateDatabaseToolsIdentityCredentialDetails Identity Credential validation details.
type ValidateDatabaseToolsIdentityCredentialDetails interface {
}

type validatedatabasetoolsidentitycredentialdetails struct {
	JsonData []byte
	Type     string `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *validatedatabasetoolsidentitycredentialdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalervalidatedatabasetoolsidentitycredentialdetails validatedatabasetoolsidentitycredentialdetails
	s := struct {
		Model Unmarshalervalidatedatabasetoolsidentitycredentialdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *validatedatabasetoolsidentitycredentialdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "ORACLE_DATABASE_RESOURCE_PRINCIPAL":
		mm := ValidateDatabaseToolsIdentityCredentialOracleDatabaseResourcePrincipalDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for ValidateDatabaseToolsIdentityCredentialDetails: %s.", m.Type)
		return *m, nil
	}
}

func (m validatedatabasetoolsidentitycredentialdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m validatedatabasetoolsidentitycredentialdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
