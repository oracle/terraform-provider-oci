// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud Migrations API
//
// A description of the Oracle Cloud Migrations API.
//

package cloudmigrations

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// OlvmCertificate The TLS certificate in case of a TLS connection.
type OlvmCertificate struct {

	// Free text containing comments about this object.
	Comment *string `mandatory:"false" json:"comment"`

	// Content of the certificate.
	Content *string `mandatory:"false" json:"content"`

	// A human-readable description in plain text.
	CertificateDescription *string `mandatory:"false" json:"certificateDescription"`

	// A unique identifier.
	Id *string `mandatory:"false" json:"id"`

	// A human-readable name in plain text.
	Name *string `mandatory:"false" json:"name"`

	// Organization of the certificate
	Organization *string `mandatory:"false" json:"organization"`

	// Subject of the certificate
	Subject *string `mandatory:"false" json:"subject"`
}

func (m OlvmCertificate) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OlvmCertificate) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
