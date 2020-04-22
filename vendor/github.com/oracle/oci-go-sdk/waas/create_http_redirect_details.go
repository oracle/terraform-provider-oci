// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Web Application Acceleration and Security Services API
//
// OCI Web Application Acceleration and Security Services
//

package waas

import (
	"github.com/oracle/oci-go-sdk/common"
)

// CreateHttpRedirectDetails The details of a HTTP Redirect configured to redirect traffic from one hostname to another.
// **Warning:** Oracle recommends that you avoid using any confidential information when you supply string values using the API.
type CreateHttpRedirectDetails struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the HTTP Redirects compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The domain from which traffic will be redirected.
	Domain *string `mandatory:"true" json:"domain"`

	// The redirect target object including all the redirect data.
	Target *HttpRedirectTarget `mandatory:"true" json:"target"`

	// The user-friendly name of the HTTP Redirect. The name can be changed and does not need to be unique.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The response code returned for the redirect to the client. For more information, see RFC 7231 (https://tools.ietf.org/html/rfc7231#section-6.4).
	ResponseCode *int `mandatory:"false" json:"responseCode"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateHttpRedirectDetails) String() string {
	return common.PointerString(m)
}
