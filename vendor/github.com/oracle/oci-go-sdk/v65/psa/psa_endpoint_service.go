// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// PsaEndpointService Required for Oracle services to provide information about their services like publicFQDNs, serviceID,compartment where its created etc.
type PsaEndpointService struct {

	// The Psa Endpoint Service's ID(this should be a unique identifier provided by service teams).
	ServiceId *string `mandatory:"true" json:"serviceId"`

	// The Psa Endpoint Service's displayName(This is the service teams registered name in splat/puffin).
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The private service access's current lifecycle state.
	LifecycleState PsaEndpointServiceLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// OCID referencing a corresponding EndpointService(PA-CP) resource for this service
	PaEndpointServiceId *string `mandatory:"true" json:"paEndpointServiceId"`

	// OCID of the service's compartment to own the resource
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the psa endpoint service.
	Id *string `mandatory:"false" json:"id"`

	// For SPLAT services, one or more SPLAT's service names that will be exposed by PSAs for this service.
	SplatServiceNames []string `mandatory:"false" json:"splatServiceNames"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

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

	// Basic description of the OCI service.
	Description *string `mandatory:"false" json:"description"`

	// List of FQDNs to be used by customers to access the service. These FQDNs will be registered in customer's
	// VCNs DNS. If provided, PSA will have an IPv4 address.
	PublicFqdns []string `mandatory:"false" json:"publicFqdns"`

	// List of IPv6-enabled FQDNs to be used by customers to access the service. These FQDNs will be
	// registered in customer's VCNs DNS. If provided, PSA will have an IPv6 address.
	PublicFqdnsV6 []string `mandatory:"false" json:"publicFqdnsV6"`

	// The date and time the PSA endpoint service was created, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The date and time the PSA endpoint service was created, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`
}

func (m PsaEndpointService) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PsaEndpointService) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingPsaEndpointServiceLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetPsaEndpointServiceLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// PsaEndpointServiceLifecycleStateEnum Enum with underlying type: string
type PsaEndpointServiceLifecycleStateEnum string

// Set of constants representing the allowable values for PsaEndpointServiceLifecycleStateEnum
const (
	PsaEndpointServiceLifecycleStateActive PsaEndpointServiceLifecycleStateEnum = "ACTIVE"
	PsaEndpointServiceLifecycleStateFailed PsaEndpointServiceLifecycleStateEnum = "FAILED"
)

var mappingPsaEndpointServiceLifecycleStateEnum = map[string]PsaEndpointServiceLifecycleStateEnum{
	"ACTIVE": PsaEndpointServiceLifecycleStateActive,
	"FAILED": PsaEndpointServiceLifecycleStateFailed,
}

var mappingPsaEndpointServiceLifecycleStateEnumLowerCase = map[string]PsaEndpointServiceLifecycleStateEnum{
	"active": PsaEndpointServiceLifecycleStateActive,
	"failed": PsaEndpointServiceLifecycleStateFailed,
}

// GetPsaEndpointServiceLifecycleStateEnumValues Enumerates the set of values for PsaEndpointServiceLifecycleStateEnum
func GetPsaEndpointServiceLifecycleStateEnumValues() []PsaEndpointServiceLifecycleStateEnum {
	values := make([]PsaEndpointServiceLifecycleStateEnum, 0)
	for _, v := range mappingPsaEndpointServiceLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetPsaEndpointServiceLifecycleStateEnumStringValues Enumerates the set of values in String for PsaEndpointServiceLifecycleStateEnum
func GetPsaEndpointServiceLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"FAILED",
	}
}

// GetMappingPsaEndpointServiceLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPsaEndpointServiceLifecycleStateEnum(val string) (PsaEndpointServiceLifecycleStateEnum, bool) {
	enum, ok := mappingPsaEndpointServiceLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
