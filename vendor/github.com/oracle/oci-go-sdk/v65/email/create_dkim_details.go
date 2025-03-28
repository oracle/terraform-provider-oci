// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Email Delivery API
//
// Use the Email Delivery API to do the necessary set up to send high-volume and application-generated emails through the OCI Email Delivery service.
// For more information, see Overview of the Email Delivery Service (https://docs.oracle.com/iaas/Content/Email/Concepts/overview.htm).
//  **Note:** Write actions (POST, UPDATE, DELETE) may take several minutes to propagate and be reflected by the API.
//  If a subsequent read request fails to reflect your changes, wait a few minutes and try again.
//

package email

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateDkimDetails Properties to create a new DKIM.
// The new object will be created in the same compartment as the EmailDomain.
type CreateDkimDetails struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the EmailDomain for this DKIM.
	EmailDomainId *string `mandatory:"true" json:"emailDomainId"`

	// The DKIM selector. This selector is required to be globally unique for this email domain.
	// If you do not provide the selector, we will generate one for you.
	// If you do provide the selector, we suggest adding a short region indicator
	// to differentiate from your signing of emails in other regions you might be subscribed to.
	// Selectors limited to ASCII characters can use alphanumeric, dash ("-"), and dot (".") characters.
	// Non-ASCII selector names should adopt IDNA2008 normalization (RFC 5891-5892).
	// Avoid entering confidential information.
	// Example: `mydomain-phx-20210228`
	Name *string `mandatory:"false" json:"name"`

	// A string that describes the details about the DKIM. It does not have to be unique,
	// and you can change it. Avoid entering confidential information.
	Description *string `mandatory:"false" json:"description"`

	// The DKIM RSA Private Key in Privacy-Enhanced Mail (PEM) format. It is a text-based representation of the private key used for signing email messages.
	PrivateKey *string `mandatory:"false" json:"privateKey"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateDkimDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateDkimDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
