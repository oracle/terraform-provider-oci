// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Integration API
//
// Oracle Integration API.
//

package integration

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// CreateCustomEndpointDetails Details for a custom endpoint for the integration instance (update).
type CreateCustomEndpointDetails struct {

	// A custom hostname to be used for the integration instance URL, in FQDN format.
	Hostname *string `mandatory:"true" json:"hostname"`

	// Optional OCID of a vault/secret containing a private SSL certificate bundle to be used for the custom hostname.
	// All certificates should be stored in a single base64 encoded secret
	// Note the update will fail if this is not a valid certificate.
	CertificateSecretId *string `mandatory:"false" json:"certificateSecretId"`
}

func (m CreateCustomEndpointDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateCustomEndpointDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
