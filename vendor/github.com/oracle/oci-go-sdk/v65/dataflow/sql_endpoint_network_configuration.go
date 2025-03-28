// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Flow API
//
// Use the Data Flow APIs to run any Apache Spark application at any scale without deploying or managing any infrastructure.
//

package dataflow

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SqlEndpointNetworkConfiguration The network configuration of a SQL Endpoint.
type SqlEndpointNetworkConfiguration interface {
}

type sqlendpointnetworkconfiguration struct {
	JsonData    []byte
	NetworkType string `json:"networkType"`
}

// UnmarshalJSON unmarshals json
func (m *sqlendpointnetworkconfiguration) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalersqlendpointnetworkconfiguration sqlendpointnetworkconfiguration
	s := struct {
		Model Unmarshalersqlendpointnetworkconfiguration
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.NetworkType = s.Model.NetworkType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *sqlendpointnetworkconfiguration) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.NetworkType {
	case "VCN":
		mm := SqlEndpointVcnConfig{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "SECURE_ACCESS":
		mm := SqlEndpointSecureAccessConfig{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for SqlEndpointNetworkConfiguration: %s.", m.NetworkType)
		return *m, nil
	}
}

func (m sqlendpointnetworkconfiguration) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m sqlendpointnetworkconfiguration) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
