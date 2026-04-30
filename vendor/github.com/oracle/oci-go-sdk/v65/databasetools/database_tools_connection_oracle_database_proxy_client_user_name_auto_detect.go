// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Tools API
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

// DatabaseToolsConnectionOracleDatabaseProxyClientUserNameAutoDetect Proxy client information for user name auto detect based proxy authentication.
type DatabaseToolsConnectionOracleDatabaseProxyClientUserNameAutoDetect struct {

	// A list of database roles for the client. These roles are enabled if the proxy is authorized to use the roles on behalf of the client.
	Roles []string `mandatory:"false" json:"roles"`
}

func (m DatabaseToolsConnectionOracleDatabaseProxyClientUserNameAutoDetect) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DatabaseToolsConnectionOracleDatabaseProxyClientUserNameAutoDetect) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m DatabaseToolsConnectionOracleDatabaseProxyClientUserNameAutoDetect) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDatabaseToolsConnectionOracleDatabaseProxyClientUserNameAutoDetect DatabaseToolsConnectionOracleDatabaseProxyClientUserNameAutoDetect
	s := struct {
		DiscriminatorParam string `json:"proxyAuthenticationType"`
		MarshalTypeDatabaseToolsConnectionOracleDatabaseProxyClientUserNameAutoDetect
	}{
		"USER_NAME_AUTO_DETECT",
		(MarshalTypeDatabaseToolsConnectionOracleDatabaseProxyClientUserNameAutoDetect)(m),
	}

	return json.Marshal(&s)
}
