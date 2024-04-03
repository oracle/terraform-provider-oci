// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// UpdateRemotePeeringConnectionInternalDetails A request to update the remote peering connection based on the peering
// status of the peer region. For Oracle internal use only
type UpdateRemotePeeringConnectionInternalDetails struct {

	// For Oracle internal use only.
	PeerDrgId *string `mandatory:"true" json:"peerDrgId"`

	// For Oracle internal use only.
	PeerDrgRouteTarget *string `mandatory:"true" json:"peerDrgRouteTarget"`

	// For Oracle internal use only.
	PeerId *string `mandatory:"true" json:"peerId"`

	// For Oracle internal use only.
	PeerRegionName *string `mandatory:"true" json:"peerRegionName"`

	// For Oracle internal use only.
	PeerTenancyId *string `mandatory:"true" json:"peerTenancyId"`

	// For Oracle internal use only.
	PeeringStateOfOtherRegion UpdateRemotePeeringConnectionInternalDetailsPeeringStateOfOtherRegionEnum `mandatory:"true" json:"peeringStateOfOtherRegion"`

	// For Oracle internal use only.
	PeerGfcRouteTargets []GfcRouteTarget `mandatory:"false" json:"peerGfcRouteTargets"`

	// For Oracle internal use only.
	IsPromoting *bool `mandatory:"false" json:"isPromoting"`

	// For Oracle internal use only.
	PeerIngressVip *string `mandatory:"false" json:"peerIngressVip"`
}

func (m UpdateRemotePeeringConnectionInternalDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateRemotePeeringConnectionInternalDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingUpdateRemotePeeringConnectionInternalDetailsPeeringStateOfOtherRegionEnum(string(m.PeeringStateOfOtherRegion)); !ok && m.PeeringStateOfOtherRegion != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PeeringStateOfOtherRegion: %s. Supported values are: %s.", m.PeeringStateOfOtherRegion, strings.Join(GetUpdateRemotePeeringConnectionInternalDetailsPeeringStateOfOtherRegionEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UpdateRemotePeeringConnectionInternalDetailsPeeringStateOfOtherRegionEnum Enum with underlying type: string
type UpdateRemotePeeringConnectionInternalDetailsPeeringStateOfOtherRegionEnum string

// Set of constants representing the allowable values for UpdateRemotePeeringConnectionInternalDetailsPeeringStateOfOtherRegionEnum
const (
	UpdateRemotePeeringConnectionInternalDetailsPeeringStateOfOtherRegionInitiated UpdateRemotePeeringConnectionInternalDetailsPeeringStateOfOtherRegionEnum = "INITIATED"
	UpdateRemotePeeringConnectionInternalDetailsPeeringStateOfOtherRegionRevoked   UpdateRemotePeeringConnectionInternalDetailsPeeringStateOfOtherRegionEnum = "REVOKED"
)

var mappingUpdateRemotePeeringConnectionInternalDetailsPeeringStateOfOtherRegionEnum = map[string]UpdateRemotePeeringConnectionInternalDetailsPeeringStateOfOtherRegionEnum{
	"INITIATED": UpdateRemotePeeringConnectionInternalDetailsPeeringStateOfOtherRegionInitiated,
	"REVOKED":   UpdateRemotePeeringConnectionInternalDetailsPeeringStateOfOtherRegionRevoked,
}

var mappingUpdateRemotePeeringConnectionInternalDetailsPeeringStateOfOtherRegionEnumLowerCase = map[string]UpdateRemotePeeringConnectionInternalDetailsPeeringStateOfOtherRegionEnum{
	"initiated": UpdateRemotePeeringConnectionInternalDetailsPeeringStateOfOtherRegionInitiated,
	"revoked":   UpdateRemotePeeringConnectionInternalDetailsPeeringStateOfOtherRegionRevoked,
}

// GetUpdateRemotePeeringConnectionInternalDetailsPeeringStateOfOtherRegionEnumValues Enumerates the set of values for UpdateRemotePeeringConnectionInternalDetailsPeeringStateOfOtherRegionEnum
func GetUpdateRemotePeeringConnectionInternalDetailsPeeringStateOfOtherRegionEnumValues() []UpdateRemotePeeringConnectionInternalDetailsPeeringStateOfOtherRegionEnum {
	values := make([]UpdateRemotePeeringConnectionInternalDetailsPeeringStateOfOtherRegionEnum, 0)
	for _, v := range mappingUpdateRemotePeeringConnectionInternalDetailsPeeringStateOfOtherRegionEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateRemotePeeringConnectionInternalDetailsPeeringStateOfOtherRegionEnumStringValues Enumerates the set of values in String for UpdateRemotePeeringConnectionInternalDetailsPeeringStateOfOtherRegionEnum
func GetUpdateRemotePeeringConnectionInternalDetailsPeeringStateOfOtherRegionEnumStringValues() []string {
	return []string{
		"INITIATED",
		"REVOKED",
	}
}

// GetMappingUpdateRemotePeeringConnectionInternalDetailsPeeringStateOfOtherRegionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateRemotePeeringConnectionInternalDetailsPeeringStateOfOtherRegionEnum(val string) (UpdateRemotePeeringConnectionInternalDetailsPeeringStateOfOtherRegionEnum, bool) {
	enum, ok := mappingUpdateRemotePeeringConnectionInternalDetailsPeeringStateOfOtherRegionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
