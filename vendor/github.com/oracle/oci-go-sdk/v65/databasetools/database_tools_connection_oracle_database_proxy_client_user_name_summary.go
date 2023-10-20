// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
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

// DatabaseToolsConnectionOracleDatabaseProxyClientUserNameSummary Proxy client information for user name based proxy authentication.
type DatabaseToolsConnectionOracleDatabaseProxyClientUserNameSummary struct {

	// The user name.
	UserName *string `mandatory:"true" json:"userName"`

	UserPassword DatabaseToolsUserPasswordSummary `mandatory:"false" json:"userPassword"`

	// A list of database roles for the client. These roles are enabled if the proxy is authorized to use the roles on behalf of the client.
	Roles []string `mandatory:"false" json:"roles"`
}

func (m DatabaseToolsConnectionOracleDatabaseProxyClientUserNameSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DatabaseToolsConnectionOracleDatabaseProxyClientUserNameSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m DatabaseToolsConnectionOracleDatabaseProxyClientUserNameSummary) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDatabaseToolsConnectionOracleDatabaseProxyClientUserNameSummary DatabaseToolsConnectionOracleDatabaseProxyClientUserNameSummary
	s := struct {
		DiscriminatorParam string `json:"proxyAuthenticationType"`
		MarshalTypeDatabaseToolsConnectionOracleDatabaseProxyClientUserNameSummary
	}{
		"USER_NAME",
		(MarshalTypeDatabaseToolsConnectionOracleDatabaseProxyClientUserNameSummary)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *DatabaseToolsConnectionOracleDatabaseProxyClientUserNameSummary) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		UserPassword databasetoolsuserpasswordsummary `json:"userPassword"`
		Roles        []string                         `json:"roles"`
		UserName     *string                          `json:"userName"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	nn, e = model.UserPassword.UnmarshalPolymorphicJSON(model.UserPassword.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.UserPassword = nn.(DatabaseToolsUserPasswordSummary)
	} else {
		m.UserPassword = nil
	}

	m.Roles = make([]string, len(model.Roles))
	copy(m.Roles, model.Roles)
	m.UserName = model.UserName

	return
}
