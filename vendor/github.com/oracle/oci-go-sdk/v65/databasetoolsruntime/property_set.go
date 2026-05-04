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

// PropertySet The base definition for a property set
type PropertySet interface {

	// Indicates whether the property set is mutable or not
	GetIsMutable() *bool
}

type propertyset struct {
	JsonData  []byte
	IsMutable *bool  `mandatory:"true" json:"isMutable"`
	Key       string `json:"key"`
}

// UnmarshalJSON unmarshals json
func (m *propertyset) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerpropertyset propertyset
	s := struct {
		Model Unmarshalerpropertyset
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.IsMutable = s.Model.IsMutable
	m.Key = s.Model.Key

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *propertyset) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Key {
	case "APEX":
		mm := PropertySetApex{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "APEX_FA_INTEGRATION":
		mm := PropertySetApexFaIntegration{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "APEX_DOCUMENT_GENERATOR":
		mm := PropertySetApexDocumentGenerator{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ORACLE_DATABASE_EXTERNAL_AUTHENTICATION":
		mm := PropertySetOracleDatabaseExternalAuthentication{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for PropertySet: %s.", m.Key)
		return *m, nil
	}
}

// GetIsMutable returns IsMutable
func (m propertyset) GetIsMutable() *bool {
	return m.IsMutable
}

func (m propertyset) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m propertyset) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
