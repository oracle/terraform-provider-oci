// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// PrivateServiceAccess Control Plane API
//
// Use the PrivateServiceAccess Control Plane API to manage privateServiceAccess.
//

package psa

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreatePsaEndpointServiceDetails description: Details for creating a PsaEndpoint service.
type CreatePsaEndpointServiceDetails struct {

	// The Psa Endpoint Service's ID(this should be a unique identifier provided by service teams).
	ServiceId *string `mandatory:"true" json:"serviceId"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Basic description of the OCI service.
	Description *string `mandatory:"true" json:"description"`

	// OCID referencing a corresponding EndpointService(PA-CP) resource for this service
	PaEndpointServiceId *string `mandatory:"true" json:"paEndpointServiceId"`

	// OCID of the service's compartment to own the resource
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Security attributes (https://docs.oracle.com/iaas/Content/zero-trust-packet-routing/zpr-artifacts.htm#security-attributes) are labels
	// for a resource that can be referenced in a Zero Trust Packet Routing (https://docs.oracle.com/iaas/Content/zero-trust-packet-routing/overview.htm)
	// (ZPR) policy to control access to ZPR-supported resources.
	// Example: `{"Oracle-DataSecurity-ZPR": {"MaxEgressCount": {"value":"42","mode":"audit"}}}`
	SecurityAttributes map[string]map[string]interface{} `mandatory:"false" json:"securityAttributes"`

	// For SPLAT services, one or more SPLAT's service names that will be exposed by PSAs for this service.
	SplatServiceNames []string `mandatory:"false" json:"splatServiceNames"`

	// List of FQDNs to be used by customers to access the service. These FQDNs will be registered in customer's
	// VCNs DNS. If provided, PSA will have an IPv4 address.
	PublicFqdns []string `mandatory:"false" json:"publicFqdns"`

	// List of IPv6-enabled FQDNs to be used by customers to access the service. These FQDNs will be
	// registered in customer's VCNs DNS. If provided, PSA will have an IPv6 address.
	PublicFqdnsV6 []string `mandatory:"false" json:"publicFqdnsV6"`
}

func (m CreatePsaEndpointServiceDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreatePsaEndpointServiceDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
