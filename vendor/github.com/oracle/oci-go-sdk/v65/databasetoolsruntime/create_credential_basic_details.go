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

// CreateCredentialBasicDetails Details for the Credential for the Basic type.
type CreateCredentialBasicDetails struct {

	// The credential_name to be created
	Key *string `mandatory:"true" json:"key"`

	// The username for the new credential.
	UserName *string `mandatory:"true" json:"userName"`

	// The password for the new credential.
	Password *string `mandatory:"true" json:"password"`
}

// GetKey returns Key
func (m CreateCredentialBasicDetails) GetKey() *string {
	return m.Key
}

func (m CreateCredentialBasicDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateCredentialBasicDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateCredentialBasicDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateCredentialBasicDetails CreateCredentialBasicDetails
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeCreateCredentialBasicDetails
	}{
		"BASIC",
		(MarshalTypeCreateCredentialBasicDetails)(m),
	}

	return json.Marshal(&s)
}
