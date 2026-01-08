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

// RequestEmailSubdomainCsrDetails Email subdomain certificate request detail
type RequestEmailSubdomainCsrDetails struct {

	// fully qualified host name
	CommonName *string `mandatory:"true" json:"commonName"`

	// company name
	OrganizationName *string `mandatory:"true" json:"organizationName"`

	// company section
	OrganizationUnit *string `mandatory:"true" json:"organizationUnit"`

	// city
	Locality *string `mandatory:"true" json:"locality"`

	// state or province
	State *string `mandatory:"true" json:"state"`

	// country name
	Country *string `mandatory:"true" json:"country"`

	// email address
	EmailAddress *string `mandatory:"true" json:"emailAddress"`

	// subject alternative names, comma separated
	Sans *string `mandatory:"false" json:"sans"`
}

func (m RequestEmailSubdomainCsrDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RequestEmailSubdomainCsrDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
