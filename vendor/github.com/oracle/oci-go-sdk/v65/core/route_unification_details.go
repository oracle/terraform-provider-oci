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

// RouteUnificationDetails Details for Route Unification APIs
type RouteUnificationDetails struct {

	// The attachment type.
	AttachmentType RouteUnificationDetailsAttachmentTypeEnum `mandatory:"true" json:"attachmentType"`
}

func (m RouteUnificationDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RouteUnificationDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingRouteUnificationDetailsAttachmentTypeEnum(string(m.AttachmentType)); !ok && m.AttachmentType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AttachmentType: %s. Supported values are: %s.", m.AttachmentType, strings.Join(GetRouteUnificationDetailsAttachmentTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// RouteUnificationDetailsAttachmentTypeEnum Enum with underlying type: string
type RouteUnificationDetailsAttachmentTypeEnum string

// Set of constants representing the allowable values for RouteUnificationDetailsAttachmentTypeEnum
const (
	RouteUnificationDetailsAttachmentTypeVirtualCircuit RouteUnificationDetailsAttachmentTypeEnum = "VIRTUAL_CIRCUIT"
	RouteUnificationDetailsAttachmentTypeIpsecTunnel    RouteUnificationDetailsAttachmentTypeEnum = "IPSEC_TUNNEL"
)

var mappingRouteUnificationDetailsAttachmentTypeEnum = map[string]RouteUnificationDetailsAttachmentTypeEnum{
	"VIRTUAL_CIRCUIT": RouteUnificationDetailsAttachmentTypeVirtualCircuit,
	"IPSEC_TUNNEL":    RouteUnificationDetailsAttachmentTypeIpsecTunnel,
}

var mappingRouteUnificationDetailsAttachmentTypeEnumLowerCase = map[string]RouteUnificationDetailsAttachmentTypeEnum{
	"virtual_circuit": RouteUnificationDetailsAttachmentTypeVirtualCircuit,
	"ipsec_tunnel":    RouteUnificationDetailsAttachmentTypeIpsecTunnel,
}

// GetRouteUnificationDetailsAttachmentTypeEnumValues Enumerates the set of values for RouteUnificationDetailsAttachmentTypeEnum
func GetRouteUnificationDetailsAttachmentTypeEnumValues() []RouteUnificationDetailsAttachmentTypeEnum {
	values := make([]RouteUnificationDetailsAttachmentTypeEnum, 0)
	for _, v := range mappingRouteUnificationDetailsAttachmentTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetRouteUnificationDetailsAttachmentTypeEnumStringValues Enumerates the set of values in String for RouteUnificationDetailsAttachmentTypeEnum
func GetRouteUnificationDetailsAttachmentTypeEnumStringValues() []string {
	return []string{
		"VIRTUAL_CIRCUIT",
		"IPSEC_TUNNEL",
	}
}

// GetMappingRouteUnificationDetailsAttachmentTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRouteUnificationDetailsAttachmentTypeEnum(val string) (RouteUnificationDetailsAttachmentTypeEnum, bool) {
	enum, ok := mappingRouteUnificationDetailsAttachmentTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
