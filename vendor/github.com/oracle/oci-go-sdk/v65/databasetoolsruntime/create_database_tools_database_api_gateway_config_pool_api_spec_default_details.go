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

// CreateDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecDefaultDetails The content of a Database Tools database API gateway config API spec sub resource to be created.
type CreateDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecDefaultDetails struct {

	// A user-friendly name. Does not have to be unique, and it’s changeable. Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The content of a string-escaped Open API spec in JSON format.
	Content *string `mandatory:"true" json:"content"`
}

// GetDisplayName returns DisplayName
func (m CreateDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecDefaultDetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetContent returns Content
func (m CreateDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecDefaultDetails) GetContent() *string {
	return m.Content
}

func (m CreateDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecDefaultDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecDefaultDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecDefaultDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecDefaultDetails CreateDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecDefaultDetails
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeCreateDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecDefaultDetails
	}{
		"DEFAULT",
		(MarshalTypeCreateDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecDefaultDetails)(m),
	}

	return json.Marshal(&s)
}
