// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Web Application Acceleration and Security Services API
//
// OCI Web Application Acceleration and Security Services
//

package waas

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateCertificateDetails The data used to create a new SSL certificate.
// **Warning:** Oracle recommends that you avoid using any confidential information when you supply string values using the API.
type CreateCertificateDetails struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to create the SSL certificate.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The data of the SSL certificate.
	//
	// **Note:** Many SSL certificate providers require an intermediate certificate chain to ensure a trusted status.
	// If your SSL certificate requires an intermediate certificate chain, please append the intermediate certificate
	// key in the `certificateData` field after the leaf certificate issued by the SSL certificate provider. If you
	// are unsure if your certificate requires an intermediate certificate chain, see your certificate
	// provider's documentation.
	//
	// The example below shows an intermediate certificate appended to a leaf certificate.
	CertificateData *string `mandatory:"true" json:"certificateData"`

	// The private key of the SSL certificate.
	PrivateKeyData *string `mandatory:"true" json:"privateKeyData"`

	// A user-friendly name for the SSL certificate. The name can be changed and does not need to be unique.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Set to `true` if the SSL certificate is self-signed.
	IsTrustVerificationDisabled *bool `mandatory:"false" json:"isTrustVerificationDisabled"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateCertificateDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateCertificateDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
