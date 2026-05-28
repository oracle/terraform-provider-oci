// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// GoldenGate API
//
// Use the Oracle Cloud Infrastructure GoldenGate APIs to perform data replication operations.
//

package goldengate

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateAiModelAuthDetails The information about new authentication details for an AI Model connection.
type CreateAiModelAuthDetails interface {
}

type createaimodelauthdetails struct {
	JsonData []byte
	AuthType string `json:"authType"`
}

// UnmarshalJSON unmarshals json
func (m *createaimodelauthdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreateaimodelauthdetails createaimodelauthdetails
	s := struct {
		Model Unmarshalercreateaimodelauthdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.AuthType = s.Model.AuthType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *createaimodelauthdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.AuthType {
	case "API_KEY":
		mm := CreateApiKeyAiModelAuthDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OCI_GEN_AI":
		mm := CreateOciGenAiModelAuthDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for CreateAiModelAuthDetails: %s.", m.AuthType)
		return *m, nil
	}
}

func (m createaimodelauthdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m createaimodelauthdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
