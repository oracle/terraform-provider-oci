// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Network Monitoring API
//
// Use the Network Monitoring API to troubleshoot routing and security issues for resources such as virtual cloud networks (VCNs) and compute instances. For more information, see the console
// documentation for the Network Path Analyzer (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/path_analyzer.htm) tool.
//

package vnmonitoring

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// TopologyContainsEntityRelationship Defines the `contains` relationship between virtual network topology entities. A `Contains` relationship
// is defined when an entity fully owns, contains or manages another entity.
// For example, a subnet is contained and managed in the scope of a VCN, therefore a VCN has a
// `contains` relationship to a subnet.
type TopologyContainsEntityRelationship struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the first entity in the relationship.
	Id1 *string `mandatory:"true" json:"id1"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the second entity in the relationship.
	Id2 *string `mandatory:"true" json:"id2"`
}

// GetId1 returns Id1
func (m TopologyContainsEntityRelationship) GetId1() *string {
	return m.Id1
}

// GetId2 returns Id2
func (m TopologyContainsEntityRelationship) GetId2() *string {
	return m.Id2
}

func (m TopologyContainsEntityRelationship) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TopologyContainsEntityRelationship) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m TopologyContainsEntityRelationship) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeTopologyContainsEntityRelationship TopologyContainsEntityRelationship
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeTopologyContainsEntityRelationship
	}{
		"CONTAINS",
		(MarshalTypeTopologyContainsEntityRelationship)(m),
	}

	return json.Marshal(&s)
}
