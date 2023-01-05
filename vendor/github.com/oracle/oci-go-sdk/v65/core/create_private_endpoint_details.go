// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Core Services API
//
// Use the Core Services API to manage resources such as virtual cloud networks (VCNs),
// compute instances, and block storage volumes. For more information, see the console
// documentation for the Networking (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.cloud.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm), and
// Block Volume (https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/overview.htm) services.
// The required permissions are documented in the
// Details for the Core Services (https://docs.cloud.oracle.com/iaas/Content/Identity/Reference/corepolicyreference.htm) article.
//

package core

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreatePrivateEndpointDetails Details to create a private endpoint.
type CreatePrivateEndpointDetails struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment to contain the
	// private endpoint.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the endpoint service that will be
	// associated with the private endpoint.
	EndpointServiceId *string `mandatory:"true" json:"endpointServiceId"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the customer's
	// subnet where the private endpoint VNIC will reside.
	SubnetId *string `mandatory:"true" json:"subnetId"`

	// Defined tags for this resource. Each key is predefined and scoped to a
	// namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// A description of this private endpoint.
	Description *string `mandatory:"false" json:"description"`

	// The three-label FQDN to use for the private endpoint. The customer VCN's DNS records are
	// updated with this FQDN.
	// For important information about how this attribute is used, see the discussion
	// of DNS and FQDNs in PrivateEndpoint.
	// Example: `xyz.oraclecloud.com`
	EndpointFqdn *string `mandatory:"false" json:"endpointFqdn"`

	// A list of additional FQDNs that you can provide along with endpointFqdn. These FQDNs are added to the
	// customer VCN's DNS record. For more information, see the discussion of DNS and FQDNs in
	// PrivateEndpoint.
	AdditionalFqdns []string `mandatory:"false" json:"additionalFqdns"`

	// A list of the OCIDs of the network security groups (NSGs) to add the private endpoint's VNIC to.
	// For more information about NSGs, see
	// NetworkSecurityGroup.
	NsgIds []string `mandatory:"false" json:"nsgIds"`

	// The private IP address to assign to this private endpoint. If you provide a value,
	// it must be an available IP address in the customer's subnet. If it's not available, an error
	// is returned.
	// If you do not provide a value, an available IP address in the subnet is automatically chosen.
	PrivateEndpointIp *string `mandatory:"false" json:"privateEndpointIp"`

	// Custom shape of the VNIC that is used when provisioning private endpoint. If the value is empty
	// then default shape is used.
	// Allowed Values:
	//   - PE_8G: 8G VNIC shape will be used to provision private endpoint.
	//   - PE_25G: 25G VNIC shape will be used to provision private endpoint.
	//   - PE_50G: 50G VNIC shape will be used to provision private endpoint.
	PrivateEndpointVnicShape CreatePrivateEndpointDetailsPrivateEndpointVnicShapeEnum `mandatory:"false" json:"privateEndpointVnicShape,omitempty"`
}

func (m CreatePrivateEndpointDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreatePrivateEndpointDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingCreatePrivateEndpointDetailsPrivateEndpointVnicShapeEnum(string(m.PrivateEndpointVnicShape)); !ok && m.PrivateEndpointVnicShape != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PrivateEndpointVnicShape: %s. Supported values are: %s.", m.PrivateEndpointVnicShape, strings.Join(GetCreatePrivateEndpointDetailsPrivateEndpointVnicShapeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreatePrivateEndpointDetailsPrivateEndpointVnicShapeEnum Enum with underlying type: string
type CreatePrivateEndpointDetailsPrivateEndpointVnicShapeEnum string

// Set of constants representing the allowable values for CreatePrivateEndpointDetailsPrivateEndpointVnicShapeEnum
const (
	CreatePrivateEndpointDetailsPrivateEndpointVnicShape8g  CreatePrivateEndpointDetailsPrivateEndpointVnicShapeEnum = "PE_8G"
	CreatePrivateEndpointDetailsPrivateEndpointVnicShape25g CreatePrivateEndpointDetailsPrivateEndpointVnicShapeEnum = "PE_25G"
	CreatePrivateEndpointDetailsPrivateEndpointVnicShape50g CreatePrivateEndpointDetailsPrivateEndpointVnicShapeEnum = "PE_50G"
)

var mappingCreatePrivateEndpointDetailsPrivateEndpointVnicShapeEnum = map[string]CreatePrivateEndpointDetailsPrivateEndpointVnicShapeEnum{
	"PE_8G":  CreatePrivateEndpointDetailsPrivateEndpointVnicShape8g,
	"PE_25G": CreatePrivateEndpointDetailsPrivateEndpointVnicShape25g,
	"PE_50G": CreatePrivateEndpointDetailsPrivateEndpointVnicShape50g,
}

var mappingCreatePrivateEndpointDetailsPrivateEndpointVnicShapeEnumLowerCase = map[string]CreatePrivateEndpointDetailsPrivateEndpointVnicShapeEnum{
	"pe_8g":  CreatePrivateEndpointDetailsPrivateEndpointVnicShape8g,
	"pe_25g": CreatePrivateEndpointDetailsPrivateEndpointVnicShape25g,
	"pe_50g": CreatePrivateEndpointDetailsPrivateEndpointVnicShape50g,
}

// GetCreatePrivateEndpointDetailsPrivateEndpointVnicShapeEnumValues Enumerates the set of values for CreatePrivateEndpointDetailsPrivateEndpointVnicShapeEnum
func GetCreatePrivateEndpointDetailsPrivateEndpointVnicShapeEnumValues() []CreatePrivateEndpointDetailsPrivateEndpointVnicShapeEnum {
	values := make([]CreatePrivateEndpointDetailsPrivateEndpointVnicShapeEnum, 0)
	for _, v := range mappingCreatePrivateEndpointDetailsPrivateEndpointVnicShapeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreatePrivateEndpointDetailsPrivateEndpointVnicShapeEnumStringValues Enumerates the set of values in String for CreatePrivateEndpointDetailsPrivateEndpointVnicShapeEnum
func GetCreatePrivateEndpointDetailsPrivateEndpointVnicShapeEnumStringValues() []string {
	return []string{
		"PE_8G",
		"PE_25G",
		"PE_50G",
	}
}

// GetMappingCreatePrivateEndpointDetailsPrivateEndpointVnicShapeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreatePrivateEndpointDetailsPrivateEndpointVnicShapeEnum(val string) (CreatePrivateEndpointDetailsPrivateEndpointVnicShapeEnum, bool) {
	enum, ok := mappingCreatePrivateEndpointDetailsPrivateEndpointVnicShapeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
