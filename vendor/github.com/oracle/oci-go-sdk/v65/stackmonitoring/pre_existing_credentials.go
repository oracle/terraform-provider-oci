// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Stack Monitoring API
//
// Stack Monitoring API.
//

package stackmonitoring

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// PreExistingCredentials Pre existing credentials [indicated by the type property in CredentialStore].
type PreExistingCredentials struct {

	// The source type and source name combination, delimited with (.) separator.
	// {source type}.{source name} and source type max char limit is 63.
	Source *string `mandatory:"false" json:"source"`

	// The name of the credential, within the context of the source.
	Name *string `mandatory:"false" json:"name"`

	// The type of the credential ( ex. JMXCreds,DBCreds).
	Type *string `mandatory:"false" json:"type"`

	// The user-specified textual description of the credential.
	Description *string `mandatory:"false" json:"description"`
}

// GetSource returns Source
func (m PreExistingCredentials) GetSource() *string {
	return m.Source
}

// GetName returns Name
func (m PreExistingCredentials) GetName() *string {
	return m.Name
}

// GetType returns Type
func (m PreExistingCredentials) GetType() *string {
	return m.Type
}

// GetDescription returns Description
func (m PreExistingCredentials) GetDescription() *string {
	return m.Description
}

func (m PreExistingCredentials) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PreExistingCredentials) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m PreExistingCredentials) MarshalJSON() (buff []byte, e error) {
	type MarshalTypePreExistingCredentials PreExistingCredentials
	s := struct {
		DiscriminatorParam string `json:"credentialType"`
		MarshalTypePreExistingCredentials
	}{
		"EXISTING",
		(MarshalTypePreExistingCredentials)(m),
	}

	return json.Marshal(&s)
}
