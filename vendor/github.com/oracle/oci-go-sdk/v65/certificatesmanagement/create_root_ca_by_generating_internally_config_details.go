// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// CreateRootCaByGeneratingInternallyConfigDetails The details for creating a private root certificate authority (CA).
type CreateRootCaByGeneratingInternallyConfigDetails struct {
	Subject *CertificateSubject `mandatory:"true" json:"subject"`

	// The name of the CA version. When the value is not null, a name is unique across versions of a given CA.
	VersionName *string `mandatory:"false" json:"versionName"`

	Validity *Validity `mandatory:"false" json:"validity"`

	// The algorithm used to sign public key certificates that the CA issues.
	SigningAlgorithm SignatureAlgorithmEnum `mandatory:"false" json:"signingAlgorithm,omitempty"`
}

// GetVersionName returns VersionName
func (m CreateRootCaByGeneratingInternallyConfigDetails) GetVersionName() *string {
	return m.VersionName
}

func (m CreateRootCaByGeneratingInternallyConfigDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateRootCaByGeneratingInternallyConfigDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingSignatureAlgorithmEnum(string(m.SigningAlgorithm)); !ok && m.SigningAlgorithm != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SigningAlgorithm: %s. Supported values are: %s.", m.SigningAlgorithm, strings.Join(GetSignatureAlgorithmEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateRootCaByGeneratingInternallyConfigDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateRootCaByGeneratingInternallyConfigDetails CreateRootCaByGeneratingInternallyConfigDetails
	s := struct {
		DiscriminatorParam string `json:"configType"`
		MarshalTypeCreateRootCaByGeneratingInternallyConfigDetails
	}{
		"ROOT_CA_GENERATED_INTERNALLY",
		(MarshalTypeCreateRootCaByGeneratingInternallyConfigDetails)(m),
	}

	return json.Marshal(&s)
}
