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

// UpdatePropertySetOracleDatabaseExternalAuthenticationDetails Contains the update details of an Oracle Database External Authentication property set
type UpdatePropertySetOracleDatabaseExternalAuthenticationDetails struct {
	IdentityProvider PropertySetOracleDatabaseExternalAuthenticationIdentityProvider `mandatory:"true" json:"identityProvider"`
}

func (m UpdatePropertySetOracleDatabaseExternalAuthenticationDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdatePropertySetOracleDatabaseExternalAuthenticationDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m UpdatePropertySetOracleDatabaseExternalAuthenticationDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdatePropertySetOracleDatabaseExternalAuthenticationDetails UpdatePropertySetOracleDatabaseExternalAuthenticationDetails
	s := struct {
		DiscriminatorParam string `json:"key"`
		MarshalTypeUpdatePropertySetOracleDatabaseExternalAuthenticationDetails
	}{
		"ORACLE_DATABASE_EXTERNAL_AUTHENTICATION",
		(MarshalTypeUpdatePropertySetOracleDatabaseExternalAuthenticationDetails)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *UpdatePropertySetOracleDatabaseExternalAuthenticationDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		IdentityProvider propertysetoracledatabaseexternalauthenticationidentityprovider `json:"identityProvider"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	nn, e = model.IdentityProvider.UnmarshalPolymorphicJSON(model.IdentityProvider.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.IdentityProvider = nn.(PropertySetOracleDatabaseExternalAuthenticationIdentityProvider)
	} else {
		m.IdentityProvider = nil
	}

	return
}
