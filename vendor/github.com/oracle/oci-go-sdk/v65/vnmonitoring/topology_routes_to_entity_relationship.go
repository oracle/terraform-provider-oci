// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Network Monitoring API
//
// Use the Network Monitoring API to troubleshoot routing and security issues for resources such as virtual cloud networks (VCNs) and compute instances. For more information, see the console
// documentation for the Network Path Analyzer (https://docs.oracle.com/iaas/Content/Network/Concepts/path_analyzer.htm) tool.
//

package vnmonitoring

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// TopologyRoutesToEntityRelationship Defines the `routesTo` relationship between virtual network topology entities. A `RoutesTo` relationship
// is defined when a routing table and a routing rule  are used to govern how to route traffic
// from one entity to another. For example, a DRG might have a routing rule to send certain traffic to an LPG.
type TopologyRoutesToEntityRelationship struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the first entity in the relationship.
	Id1 *string `mandatory:"true" json:"id1"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the second entity in the relationship.
	Id2 *string `mandatory:"true" json:"id2"`

	RouteRuleDetails *TopologyRoutesToRelationshipDetails `mandatory:"true" json:"routeRuleDetails"`
}

// GetId1 returns Id1
func (m TopologyRoutesToEntityRelationship) GetId1() *string {
	return m.Id1
}

// GetId2 returns Id2
func (m TopologyRoutesToEntityRelationship) GetId2() *string {
	return m.Id2
}

func (m TopologyRoutesToEntityRelationship) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TopologyRoutesToEntityRelationship) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m TopologyRoutesToEntityRelationship) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeTopologyRoutesToEntityRelationship TopologyRoutesToEntityRelationship
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeTopologyRoutesToEntityRelationship
	}{
		"ROUTES_TO",
		(MarshalTypeTopologyRoutesToEntityRelationship)(m),
	}

	return json.Marshal(&s)
}
