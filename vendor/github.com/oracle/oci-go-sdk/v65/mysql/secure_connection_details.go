// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// MySQL Database Service API
//
// The API for the MySQL Database Service
//

package mysql

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SecureConnectionDetails Secure connection configuration details.
type SecureConnectionDetails struct {

	// Select whether to use MySQL Database Service-managed certificate (SYSTEM) or your own certificate (BYOC).
	CertificateGenerationType CertificateGenerationTypeEnum `mandatory:"true" json:"certificateGenerationType"`

	// The OCID of the certificate to use.
	CertificateId *string `mandatory:"false" json:"certificateId"`
}

func (m SecureConnectionDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SecureConnectionDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCertificateGenerationTypeEnum(string(m.CertificateGenerationType)); !ok && m.CertificateGenerationType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CertificateGenerationType: %s. Supported values are: %s.", m.CertificateGenerationType, strings.Join(GetCertificateGenerationTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
