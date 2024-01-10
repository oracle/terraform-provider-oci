// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Document Understanding API
//
// Document AI helps customers perform various analysis on their documents. If a customer has lots of documents, they can process them in batch using asynchronous API endpoints.
//

package aidocument

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ProcessorConfig The configuration of a processor.
type ProcessorConfig interface {
}

type processorconfig struct {
	JsonData      []byte
	ProcessorType string `json:"processorType"`
}

// UnmarshalJSON unmarshals json
func (m *processorconfig) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerprocessorconfig processorconfig
	s := struct {
		Model Unmarshalerprocessorconfig
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ProcessorType = s.Model.ProcessorType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *processorconfig) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ProcessorType {
	case "GENERAL":
		mm := GeneralProcessorConfig{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for ProcessorConfig: %s.", m.ProcessorType)
		return *m, nil
	}
}

func (m processorconfig) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m processorconfig) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
