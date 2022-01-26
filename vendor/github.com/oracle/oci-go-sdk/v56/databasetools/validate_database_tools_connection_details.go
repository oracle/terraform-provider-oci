// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Tools
//
// Database Tools APIs to manage Connections and Private Endpoints.
//

package databasetools

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v56/common"
)

// ValidateDatabaseToolsConnectionDetails Connection validation details.
type ValidateDatabaseToolsConnectionDetails interface {
}

type validatedatabasetoolsconnectiondetails struct {
	JsonData []byte
	Type     string `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *validatedatabasetoolsconnectiondetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalervalidatedatabasetoolsconnectiondetails validatedatabasetoolsconnectiondetails
	s := struct {
		Model Unmarshalervalidatedatabasetoolsconnectiondetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *validatedatabasetoolsconnectiondetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "ORACLE_DATABASE":
		mm := ValidateDatabaseToolsConnectionOracleDatabaseDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

func (m validatedatabasetoolsconnectiondetails) String() string {
	return common.PointerString(m)
}
