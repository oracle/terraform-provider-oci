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

// PlainTextCredentials Plain text credentials [indicated by the type property in CredentialStore].
type PlainTextCredentials struct {

	// The credential properties list. Credential property values will be either
	// in plain text format or encrypted for encrypted credentials.
	Properties []CredentialProperty `mandatory:"true" json:"properties"`

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
func (m PlainTextCredentials) GetSource() *string {
	return m.Source
}

// GetName returns Name
func (m PlainTextCredentials) GetName() *string {
	return m.Name
}

// GetType returns Type
func (m PlainTextCredentials) GetType() *string {
	return m.Type
}

// GetDescription returns Description
func (m PlainTextCredentials) GetDescription() *string {
	return m.Description
}

func (m PlainTextCredentials) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PlainTextCredentials) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m PlainTextCredentials) MarshalJSON() (buff []byte, e error) {
	type MarshalTypePlainTextCredentials PlainTextCredentials
	s := struct {
		DiscriminatorParam string `json:"credentialType"`
		MarshalTypePlainTextCredentials
	}{
		"PLAINTEXT",
		(MarshalTypePlainTextCredentials)(m),
	}

	return json.Marshal(&s)
}
