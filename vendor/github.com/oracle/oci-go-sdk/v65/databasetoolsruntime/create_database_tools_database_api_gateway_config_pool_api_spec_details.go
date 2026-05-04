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

// CreateDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecDetails The content of a Database Tools database API gateway config API spec sub resource to be created.
type CreateDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecDetails interface {

	// A user-friendly name. Does not have to be unique, and it’s changeable. Avoid entering confidential information.
	GetDisplayName() *string

	// The content of a string-escaped Open API spec in JSON format.
	GetContent() *string
}

type createdatabasetoolsdatabaseapigatewayconfigpoolapispecdetails struct {
	JsonData    []byte
	DisplayName *string `mandatory:"true" json:"displayName"`
	Content     *string `mandatory:"true" json:"content"`
	Type        string  `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *createdatabasetoolsdatabaseapigatewayconfigpoolapispecdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreatedatabasetoolsdatabaseapigatewayconfigpoolapispecdetails createdatabasetoolsdatabaseapigatewayconfigpoolapispecdetails
	s := struct {
		Model Unmarshalercreatedatabasetoolsdatabaseapigatewayconfigpoolapispecdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.DisplayName = s.Model.DisplayName
	m.Content = s.Model.Content
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *createdatabasetoolsdatabaseapigatewayconfigpoolapispecdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "DEFAULT":
		mm := CreateDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecDefaultDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for CreateDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecDetails: %s.", m.Type)
		return *m, nil
	}
}

// GetDisplayName returns DisplayName
func (m createdatabasetoolsdatabaseapigatewayconfigpoolapispecdetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetContent returns Content
func (m createdatabasetoolsdatabaseapigatewayconfigpoolapispecdetails) GetContent() *string {
	return m.Content
}

func (m createdatabasetoolsdatabaseapigatewayconfigpoolapispecdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m createdatabasetoolsdatabaseapigatewayconfigpoolapispecdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
