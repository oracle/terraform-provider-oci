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
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// VirtualCircuitDrgAttachmentNetworkCreateDetails The representation of VirtualCircuitDrgAttachmentNetworkCreateDetails
type VirtualCircuitDrgAttachmentNetworkCreateDetails struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of compartment that contains the Virtual Circuit.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the network attached to the DRG.
	Id *string `mandatory:"false" json:"id"`

	// The BGP ASN to use for the Virtual Circuit's route target
	RegionalOciAsn *string `mandatory:"false" json:"regionalOciAsn"`

	// Whether the Fast Connect exists through an edge pop region.
	// Example: `true`
	IsEdgePop *bool `mandatory:"false" json:"isEdgePop"`

	// The OCI region name
	RegionName *string `mandatory:"false" json:"regionName"`

	// Boolean flag that determines wether all traffic over the VCs is encrypted.
	// Example: `true`
	TransportOnlyMode *bool `mandatory:"false" json:"transportOnlyMode"`

	// Determines Throughput capacity of Virtual Circuit.
	// Example: `400G`
	Throughput VirtualCircuitDrgAttachmentNetworkCreateDetailsThroughputEnum `mandatory:"false" json:"throughput,omitempty"`
}

//GetId returns Id
func (m VirtualCircuitDrgAttachmentNetworkCreateDetails) GetId() *string {
	return m.Id
}

func (m VirtualCircuitDrgAttachmentNetworkCreateDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m VirtualCircuitDrgAttachmentNetworkCreateDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingVirtualCircuitDrgAttachmentNetworkCreateDetailsThroughputEnum(string(m.Throughput)); !ok && m.Throughput != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Throughput: %s. Supported values are: %s.", m.Throughput, strings.Join(GetVirtualCircuitDrgAttachmentNetworkCreateDetailsThroughputEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m VirtualCircuitDrgAttachmentNetworkCreateDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeVirtualCircuitDrgAttachmentNetworkCreateDetails VirtualCircuitDrgAttachmentNetworkCreateDetails
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeVirtualCircuitDrgAttachmentNetworkCreateDetails
	}{
		"VIRTUAL_CIRCUIT",
		(MarshalTypeVirtualCircuitDrgAttachmentNetworkCreateDetails)(m),
	}

	return json.Marshal(&s)
}

// VirtualCircuitDrgAttachmentNetworkCreateDetailsThroughputEnum Enum with underlying type: string
type VirtualCircuitDrgAttachmentNetworkCreateDetailsThroughputEnum string

// Set of constants representing the allowable values for VirtualCircuitDrgAttachmentNetworkCreateDetailsThroughputEnum
const (
	VirtualCircuitDrgAttachmentNetworkCreateDetailsThroughput400g VirtualCircuitDrgAttachmentNetworkCreateDetailsThroughputEnum = "400G"
)

var mappingVirtualCircuitDrgAttachmentNetworkCreateDetailsThroughputEnum = map[string]VirtualCircuitDrgAttachmentNetworkCreateDetailsThroughputEnum{
	"400G": VirtualCircuitDrgAttachmentNetworkCreateDetailsThroughput400g,
}

var mappingVirtualCircuitDrgAttachmentNetworkCreateDetailsThroughputEnumLowerCase = map[string]VirtualCircuitDrgAttachmentNetworkCreateDetailsThroughputEnum{
	"400g": VirtualCircuitDrgAttachmentNetworkCreateDetailsThroughput400g,
}

// GetVirtualCircuitDrgAttachmentNetworkCreateDetailsThroughputEnumValues Enumerates the set of values for VirtualCircuitDrgAttachmentNetworkCreateDetailsThroughputEnum
func GetVirtualCircuitDrgAttachmentNetworkCreateDetailsThroughputEnumValues() []VirtualCircuitDrgAttachmentNetworkCreateDetailsThroughputEnum {
	values := make([]VirtualCircuitDrgAttachmentNetworkCreateDetailsThroughputEnum, 0)
	for _, v := range mappingVirtualCircuitDrgAttachmentNetworkCreateDetailsThroughputEnum {
		values = append(values, v)
	}
	return values
}

// GetVirtualCircuitDrgAttachmentNetworkCreateDetailsThroughputEnumStringValues Enumerates the set of values in String for VirtualCircuitDrgAttachmentNetworkCreateDetailsThroughputEnum
func GetVirtualCircuitDrgAttachmentNetworkCreateDetailsThroughputEnumStringValues() []string {
	return []string{
		"400G",
	}
}

// GetMappingVirtualCircuitDrgAttachmentNetworkCreateDetailsThroughputEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingVirtualCircuitDrgAttachmentNetworkCreateDetailsThroughputEnum(val string) (VirtualCircuitDrgAttachmentNetworkCreateDetailsThroughputEnum, bool) {
	enum, ok := mappingVirtualCircuitDrgAttachmentNetworkCreateDetailsThroughputEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
