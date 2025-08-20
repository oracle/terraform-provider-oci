// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Core Services API
//
// Use the Core Services API to manage resources such as virtual cloud networks (VCNs),
// compute instances, and block storage volumes. For more information, see the console
// documentation for the Networking (https://docs.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm), and
// Block Volume (https://docs.oracle.com/iaas/Content/Block/Concepts/overview.htm) services.
// The required permissions are documented in the
// Details for the Core Services (https://docs.oracle.com/iaas/Content/Identity/Reference/corepolicyreference.htm) article.
//

package core

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// RouteReflectorRoute Disintermediated routes
type RouteReflectorRoute struct {

	// Identity of routing prefix.
	Nlri *string `mandatory:"true" json:"nlri"`

	// Next-hop IP address for the route.
	NextHop *string `mandatory:"false" json:"nextHop"`

	// MPLS label associated with the route.
	MplsLabel *int `mandatory:"false" json:"mplsLabel"`

	// BGP Local Preference attribute.
	LocalPreference *int `mandatory:"false" json:"localPreference"`

	// Standard BGP community values.
	Communities []string `mandatory:"false" json:"communities"`

	// Large community information.
	LargeCommunityInfo []string `mandatory:"false" json:"largeCommunityInfo"`

	// The common or attachment specific label
	RouteTarget *string `mandatory:"false" json:"routeTarget"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DRG route table assigned to the DRG attachment.
	DrgRouteTableId *string `mandatory:"false" json:"drgRouteTableId"`

	// Prefix of the route
	Prefix *string `mandatory:"false" json:"prefix"`

	// Indicates if it is public
	IsPublic *bool `mandatory:"false" json:"isPublic"`
}

func (m RouteReflectorRoute) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RouteReflectorRoute) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
