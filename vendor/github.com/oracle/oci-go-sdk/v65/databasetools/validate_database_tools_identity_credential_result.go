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

// ValidateDatabaseToolsIdentityCredentialResult Identity validation result.
type ValidateDatabaseToolsIdentityCredentialResult interface {

	// A short code that defines the result of the validation, meant for programmatic parsing. The value OK indicates that the validation was successful.
	GetCode() *string

	// A human-readable message that describes the result of the validation.
	GetMessage() *string

	// A human-readable message that describes possible causes for the validation error.
	GetCause() *string

	// A human-readable message that suggests a remedial action to resolve the validation error.
	GetAction() *string
}

type validatedatabasetoolsidentitycredentialresult struct {
	JsonData []byte
	Cause    *string `mandatory:"false" json:"cause"`
	Action   *string `mandatory:"false" json:"action"`
	Code     *string `mandatory:"true" json:"code"`
	Message  *string `mandatory:"true" json:"message"`
	Type     string  `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *validatedatabasetoolsidentitycredentialresult) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalervalidatedatabasetoolsidentitycredentialresult validatedatabasetoolsidentitycredentialresult
	s := struct {
		Model Unmarshalervalidatedatabasetoolsidentitycredentialresult
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Code = s.Model.Code
	m.Message = s.Model.Message
	m.Cause = s.Model.Cause
	m.Action = s.Model.Action
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *validatedatabasetoolsidentitycredentialresult) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "ORACLE_DATABASE_RESOURCE_PRINCIPAL":
		mm := ValidateDatabaseToolsIdentityCredentialOracleDatabaseResourcePrincipalResult{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for ValidateDatabaseToolsIdentityCredentialResult: %s.", m.Type)
		return *m, nil
	}
}

// GetCause returns Cause
func (m validatedatabasetoolsidentitycredentialresult) GetCause() *string {
	return m.Cause
}

// GetAction returns Action
func (m validatedatabasetoolsidentitycredentialresult) GetAction() *string {
	return m.Action
}

// GetCode returns Code
func (m validatedatabasetoolsidentitycredentialresult) GetCode() *string {
	return m.Code
}

// GetMessage returns Message
func (m validatedatabasetoolsidentitycredentialresult) GetMessage() *string {
	return m.Message
}

func (m validatedatabasetoolsidentitycredentialresult) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m validatedatabasetoolsidentitycredentialresult) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
