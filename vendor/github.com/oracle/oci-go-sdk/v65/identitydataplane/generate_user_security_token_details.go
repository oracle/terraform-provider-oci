// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Identity and Access Management Data Plane API
//
// APIs for managing identity data plane services. For example, use this API to create a scoped-access security token. To manage identity domains (for example, creating or deleting an identity domain) or to manage resources (for example, users and groups) within the default identity domain, see IAM API (https://docs.oracle.com/iaas/api/#/en/identity/).
//

package identitydataplane

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// GenerateUserSecurityTokenDetails Request parameters in body for obtaining a user principal session token (UPST) for self.
type GenerateUserSecurityTokenDetails struct {

	// The user-owned public key in PEM format that corresponds to the RSA key pair used for signing requests.
	// The user also owns the corresponding private key. This public key will be put inside the user
	// security token by the auth service after successful validation of the request.
	PublicKey *string `mandatory:"true" json:"publicKey"`

	// User session expiration in minutes to which the requested user principal session token (UPST) is bounded.
	// Valid values are from 5 to 60 for all realms.
	SessionExpirationInMinutes *int `mandatory:"false" json:"sessionExpirationInMinutes"`
}

func (m GenerateUserSecurityTokenDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m GenerateUserSecurityTokenDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
