// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Core Services API
//
// Use the Core Services API to manage resources such as virtual cloud networks (VCNs),
// compute instances, and block storage volumes. For more information, see the console
// documentation for the Networking (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.cloud.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm), and
// Block Volume (https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/overview.htm) services.
//

package core

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// TopologyRoutesToRelationshipDetails Defines route rule details for a `routesTo` relationship.
type TopologyRoutesToRelationshipDetails struct {

	// The destinationType can be set to one of two values:
	// * Use `CIDR_BLOCK` if the rule's `destination` is an IP address range in CIDR notation.
	// * Use `SERVICE_CIDR_BLOCK` if the rule's `destination` is the `cidrBlock` value for a Service.
	DestinationType *string `mandatory:"true" json:"destinationType"`

	// An IP address range in CIDR notation or the `cidrBlock` value for a Service.
	Destination *string `mandatory:"true" json:"destination"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the routing table that contains the route rule.
	RouteTableId *string `mandatory:"true" json:"routeTableId"`

	// A route rule can be `STATIC` if manually added to the route table or `DYNAMIC` if imported from another route table.
	RouteType TopologyRoutesToRelationshipDetailsRouteTypeEnum `mandatory:"false" json:"routeType,omitempty"`
}

func (m TopologyRoutesToRelationshipDetails) String() string {
	return common.PointerString(m)
}

// TopologyRoutesToRelationshipDetailsRouteTypeEnum Enum with underlying type: string
type TopologyRoutesToRelationshipDetailsRouteTypeEnum string

// Set of constants representing the allowable values for TopologyRoutesToRelationshipDetailsRouteTypeEnum
const (
	TopologyRoutesToRelationshipDetailsRouteTypeStatic  TopologyRoutesToRelationshipDetailsRouteTypeEnum = "STATIC"
	TopologyRoutesToRelationshipDetailsRouteTypeDynamic TopologyRoutesToRelationshipDetailsRouteTypeEnum = "DYNAMIC"
)

var mappingTopologyRoutesToRelationshipDetailsRouteType = map[string]TopologyRoutesToRelationshipDetailsRouteTypeEnum{
	"STATIC":  TopologyRoutesToRelationshipDetailsRouteTypeStatic,
	"DYNAMIC": TopologyRoutesToRelationshipDetailsRouteTypeDynamic,
}

// GetTopologyRoutesToRelationshipDetailsRouteTypeEnumValues Enumerates the set of values for TopologyRoutesToRelationshipDetailsRouteTypeEnum
func GetTopologyRoutesToRelationshipDetailsRouteTypeEnumValues() []TopologyRoutesToRelationshipDetailsRouteTypeEnum {
	values := make([]TopologyRoutesToRelationshipDetailsRouteTypeEnum, 0)
	for _, v := range mappingTopologyRoutesToRelationshipDetailsRouteType {
		values = append(values, v)
	}
	return values
}
