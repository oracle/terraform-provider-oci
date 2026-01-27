// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Certificates Service Management API
//
// API for managing certificates.
//

package certificatesmanagement

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateRootCaManagedExternallyConfigDetails The configuration details for creating an externally managed private root certificate authority (CA) issued by an external CA.
type CreateRootCaManagedExternallyConfigDetails struct {

	// The externally signed certificate (in PEM format) for the imported root CA.
	CertificatePem *string `mandatory:"true" json:"certificatePem"`

	// The name of the CA version. When the value is not null, a name is unique across versions of a given CA.
	VersionName *string `mandatory:"false" json:"versionName"`
}

// GetVersionName returns VersionName
func (m CreateRootCaManagedExternallyConfigDetails) GetVersionName() *string {
	return m.VersionName
}

func (m CreateRootCaManagedExternallyConfigDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateRootCaManagedExternallyConfigDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateRootCaManagedExternallyConfigDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateRootCaManagedExternallyConfigDetails CreateRootCaManagedExternallyConfigDetails
	s := struct {
		DiscriminatorParam string `json:"configType"`
		MarshalTypeCreateRootCaManagedExternallyConfigDetails
	}{
		"ROOT_CA_MANAGED_EXTERNALLY",
		(MarshalTypeCreateRootCaManagedExternallyConfigDetails)(m),
	}

	return json.Marshal(&s)
}
