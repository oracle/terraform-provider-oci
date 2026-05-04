// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Tools Runtime API
//
// Use the Database Tools Runtime API to connect to databases through Database Tools Connections.
//

package databasetoolsruntime

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// PropertySetOracleDatabaseExternalAuthenticationIdentityProvider External identity type provider
type PropertySetOracleDatabaseExternalAuthenticationIdentityProvider interface {
}

type propertysetoracledatabaseexternalauthenticationidentityprovider struct {
	JsonData []byte
	Type     string `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *propertysetoracledatabaseexternalauthenticationidentityprovider) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerpropertysetoracledatabaseexternalauthenticationidentityprovider propertysetoracledatabaseexternalauthenticationidentityprovider
	s := struct {
		Model Unmarshalerpropertysetoracledatabaseexternalauthenticationidentityprovider
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *propertysetoracledatabaseexternalauthenticationidentityprovider) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "OCI_IAM":
		mm := PropertySetOracleDatabaseExternalAuthenticationIdentityProviderOciIam{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "AZURE_AD":
		mm := PropertySetOracleDatabaseExternalAuthenticationIdentityProviderAzureAd{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "NONE":
		mm := PropertySetOracleDatabaseExternalAuthenticationIdentityProviderNone{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for PropertySetOracleDatabaseExternalAuthenticationIdentityProvider: %s.", m.Type)
		return *m, nil
	}
}

func (m propertysetoracledatabaseexternalauthenticationidentityprovider) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m propertysetoracledatabaseexternalauthenticationidentityprovider) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
