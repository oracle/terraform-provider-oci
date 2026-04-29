// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fusion Applications Environment Management API
//
// Use the Fusion Applications Environment Management API to manage the environments where your Fusion Applications run. For more information, see the Fusion Applications Environment Management documentation (https://docs.oracle.com/iaas/Content/fusion-applications/home.htm).
//

package fusionapps

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UploadVanityDomainCertificateDetails Add vanity domain with certificate
type UploadVanityDomainCertificateDetails struct {

	// Identify if this is certificate for pod LB or Akamai
	IsAkamai *bool `mandatory:"true" json:"isAkamai"`

	// Fully qualified host name
	CommonName *string `mandatory:"true" json:"commonName"`

	// Public certificate PEM
	CertificatePem *string `mandatory:"true" json:"certificatePem"`

	// Certificate chain PEM, including intermediate and root
	CertificateChainPem *string `mandatory:"false" json:"certificateChainPem"`
}

func (m UploadVanityDomainCertificateDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UploadVanityDomainCertificateDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
