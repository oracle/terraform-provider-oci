// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Certificates Service Management API
//
// API for managing certificates.
//

package certificatesmanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// PublicCaCertificateValidity An object that describes a period of time during which an entity is valid. If this is not provided when you create a certificate, the validity of the issuing CA is used.
type PublicCaCertificateValidity struct {

	// The desired validity period for the certificate.
	// Example: `DAYS_90`
	ValidityInDays PublicCertificateValidityEnum `mandatory:"true" json:"validityInDays"`
}

func (m PublicCaCertificateValidity) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PublicCaCertificateValidity) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingPublicCertificateValidityEnum(string(m.ValidityInDays)); !ok && m.ValidityInDays != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ValidityInDays: %s. Supported values are: %s.", m.ValidityInDays, strings.Join(GetPublicCertificateValidityEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
