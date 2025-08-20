// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Marketplace Service API
//
// Use the Marketplace API to manage applications in Oracle Cloud Infrastructure Marketplace. For more information, see Overview of Marketplace (https://docs.oracle.com/iaas/Content/Marketplace/Concepts/marketoverview.htm)
//

package marketplace

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// MarketplaceMetadataPublicKeySummary Model that contains signed marketplace instance metadata and associated signature in JWT form
type MarketplaceMetadataPublicKeySummary struct {

	// algorithm for public key (i.e. RS256)
	KeyAlgorithm *string `mandatory:"true" json:"keyAlgorithm"`

	// key type (i.e. RSA)
	KeyType *string `mandatory:"true" json:"keyType"`

	// how key is to be used
	KeyUse *string `mandatory:"true" json:"keyUse"`

	// unique id that maps to public certificate, directs user which certificate to use to verfiy
	KeyId *string `mandatory:"true" json:"keyId"`

	// base64 encoded exponent for public key
	Exponent *string `mandatory:"true" json:"exponent"`

	// RSA public modulus
	Modulus *string `mandatory:"true" json:"modulus"`

	// chain of certificates used to sign JWT
	CertificateChain []string `mandatory:"true" json:"certificateChain"`

	// unique identifier of associated X509 certificate
	CertificateThumbprint *string `mandatory:"true" json:"certificateThumbprint"`
}

func (m MarketplaceMetadataPublicKeySummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MarketplaceMetadataPublicKeySummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
