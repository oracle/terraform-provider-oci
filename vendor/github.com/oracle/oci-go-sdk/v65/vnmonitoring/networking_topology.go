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

// NetworkingTopology Defines the representation of a virtual network topology for a region.
type NetworkingTopology struct {

	// Lists entities comprising the virtual network topology.
	Entities []interface{} `mandatory:"true" json:"entities"`

	// Lists relationships between entities in the virtual network topology.
	Relationships []TopologyEntityRelationship `mandatory:"true" json:"relationships"`

	// Lists entities that are limited during ingestion.
	// The values for the items in the list are the entity type names of the limitedEntities.
	// Example: `vcn`
	LimitedEntities []string `mandatory:"true" json:"limitedEntities"`

	// Records when the virtual network topology was created, in RFC3339 (https://tools.ietf.org/html/rfc3339) format for date and time.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`
}

// GetEntities returns Entities
func (m NetworkingTopology) GetEntities() []interface{} {
	return m.Entities
}

// GetRelationships returns Relationships
func (m NetworkingTopology) GetRelationships() []TopologyEntityRelationship {
	return m.Relationships
}

// GetLimitedEntities returns LimitedEntities
func (m NetworkingTopology) GetLimitedEntities() []string {
	return m.LimitedEntities
}

// GetTimeCreated returns TimeCreated
func (m NetworkingTopology) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

func (m NetworkingTopology) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m NetworkingTopology) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m NetworkingTopology) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeNetworkingTopology NetworkingTopology
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeNetworkingTopology
	}{
		"NETWORKING",
		(MarshalTypeNetworkingTopology)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *NetworkingTopology) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Entities        []interface{}                `json:"entities"`
		Relationships   []topologyentityrelationship `json:"relationships"`
		LimitedEntities []string                     `json:"limitedEntities"`
		TimeCreated     *common.SDKTime              `json:"timeCreated"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Entities = make([]interface{}, len(model.Entities))
	copy(m.Entities, model.Entities)
	m.Relationships = make([]TopologyEntityRelationship, len(model.Relationships))
	for i, n := range model.Relationships {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.Relationships[i] = nn.(TopologyEntityRelationship)
		} else {
			m.Relationships[i] = nil
		}
	}
	m.LimitedEntities = make([]string, len(model.LimitedEntities))
	copy(m.LimitedEntities, model.LimitedEntities)
	m.TimeCreated = model.TimeCreated

	return
}
