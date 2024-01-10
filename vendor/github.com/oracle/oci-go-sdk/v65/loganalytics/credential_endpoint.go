// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CredentialEndpoint The endpoint from where to fetch a credential, for example, the OAuth 2.0 token.
type CredentialEndpoint struct {

	// The credential endpoint name.
	Name *string `mandatory:"true" json:"name"`

	Request *EndpointRequest `mandatory:"true" json:"request"`

	// The credential endpoint description.
	Description *string `mandatory:"false" json:"description"`

	// The credential endpoint model.
	Model *string `mandatory:"false" json:"model"`

	// The endpoint unique identifier.
	EndpointId *int64 `mandatory:"false" json:"endpointId"`

	Response *EndpointResponse `mandatory:"false" json:"response"`

	Proxy *EndpointProxy `mandatory:"false" json:"proxy"`
}

func (m CredentialEndpoint) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CredentialEndpoint) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
