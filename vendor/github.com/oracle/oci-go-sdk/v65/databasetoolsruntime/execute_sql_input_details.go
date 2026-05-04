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

// ExecuteSqlInputDetails Contains the details for the SQL statements to execute on the database connection.
type ExecuteSqlInputDetails interface {
}

type executesqlinputdetails struct {
	JsonData []byte
	Type     string `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *executesqlinputdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerexecutesqlinputdetails executesqlinputdetails
	s := struct {
		Model Unmarshalerexecutesqlinputdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *executesqlinputdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "STANDARD":
		mm := ExecuteSqlInputStandardDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "BATCH":
		mm := ExecuteSqlInputBatchDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "SCRIPT":
		mm := ExecuteSqlInputScriptDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for ExecuteSqlInputDetails: %s.", m.Type)
		return *m, nil
	}
}

func (m executesqlinputdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m executesqlinputdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
