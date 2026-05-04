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

// UpdatePropertySetDetails The base definition for a property set update
type UpdatePropertySetDetails interface {
}

type updatepropertysetdetails struct {
	JsonData []byte
	Key      string `json:"key"`
}

// UnmarshalJSON unmarshals json
func (m *updatepropertysetdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerupdatepropertysetdetails updatepropertysetdetails
	s := struct {
		Model Unmarshalerupdatepropertysetdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Key = s.Model.Key

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *updatepropertysetdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Key {
	case "APEX_DOCUMENT_GENERATOR":
		mm := UpdatePropertySetApexDocumentGeneratorDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ORACLE_DATABASE_EXTERNAL_AUTHENTICATION":
		mm := UpdatePropertySetOracleDatabaseExternalAuthenticationDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "APEX_FA_INTEGRATION":
		mm := UpdatePropertySetApexFaIntegrationDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for UpdatePropertySetDetails: %s.", m.Key)
		return *m, nil
	}
}

func (m updatepropertysetdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m updatepropertysetdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
