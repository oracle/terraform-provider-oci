// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OpenSearch Service API
//
// The OpenSearch service API provides access to OCI Search Service with OpenSearch.
//

package opensearch

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ResourcePrincipalTokenResponse Contains a Resource Principal Token (RPT) and Service Principal Session Token (SPST), both of
// which are needed to obtain a Resource Principal Session Token (RPST) from Identity Data Plane.
type ResourcePrincipalTokenResponse struct {

	// A Resource Principal Token
	ResourcePrincipalToken *string `mandatory:"true" json:"resourcePrincipalToken"`

	// A Service Principal Session Token
	ServicePrincipalSessionToken *string `mandatory:"true" json:"servicePrincipalSessionToken"`
}

func (m ResourcePrincipalTokenResponse) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ResourcePrincipalTokenResponse) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
