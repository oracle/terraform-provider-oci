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

// DatabaseToolsConnectionOracleDatabaseProxyClientSummary The proxy client information.
type DatabaseToolsConnectionOracleDatabaseProxyClientSummary interface {
}

type databasetoolsconnectionoracledatabaseproxyclientsummary struct {
	JsonData                []byte
	ProxyAuthenticationType string `json:"proxyAuthenticationType"`
}

// UnmarshalJSON unmarshals json
func (m *databasetoolsconnectionoracledatabaseproxyclientsummary) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerdatabasetoolsconnectionoracledatabaseproxyclientsummary databasetoolsconnectionoracledatabaseproxyclientsummary
	s := struct {
		Model Unmarshalerdatabasetoolsconnectionoracledatabaseproxyclientsummary
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ProxyAuthenticationType = s.Model.ProxyAuthenticationType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *databasetoolsconnectionoracledatabaseproxyclientsummary) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ProxyAuthenticationType {
	case "NO_PROXY":
		mm := DatabaseToolsConnectionOracleDatabaseProxyClientNoProxySummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "USER_NAME":
		mm := DatabaseToolsConnectionOracleDatabaseProxyClientUserNameSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for DatabaseToolsConnectionOracleDatabaseProxyClientSummary: %s.", m.ProxyAuthenticationType)
		return *m, nil
	}
}

func (m databasetoolsconnectionoracledatabaseproxyclientsummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m databasetoolsconnectionoracledatabaseproxyclientsummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
