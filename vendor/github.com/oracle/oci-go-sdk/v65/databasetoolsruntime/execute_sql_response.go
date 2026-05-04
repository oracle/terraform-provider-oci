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

// ExecuteSqlResponse Contains the details for the SQL statements to execute on the database connection.
type ExecuteSqlResponse interface {
	GetEnv() *ExecuteSqlResponseEnv

	// Script version
	GetVersion() *string
}

type executesqlresponse struct {
	JsonData []byte
	Env      *ExecuteSqlResponseEnv `mandatory:"false" json:"env"`
	Version  *string                `mandatory:"false" json:"version"`
	Type     string                 `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *executesqlresponse) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerexecutesqlresponse executesqlresponse
	s := struct {
		Model Unmarshalerexecutesqlresponse
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Env = s.Model.Env
	m.Version = s.Model.Version
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *executesqlresponse) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "STANDARD":
		mm := ExecuteSqlResponseStandard{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "SCRIPT":
		mm := ExecuteSqlResponseScript{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "BATCH":
		mm := ExecuteSqlResponseBatch{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for ExecuteSqlResponse: %s.", m.Type)
		return *m, nil
	}
}

// GetEnv returns Env
func (m executesqlresponse) GetEnv() *ExecuteSqlResponseEnv {
	return m.Env
}

// GetVersion returns Version
func (m executesqlresponse) GetVersion() *string {
	return m.Version
}

func (m executesqlresponse) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m executesqlresponse) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
