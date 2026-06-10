// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Streaming with Apache Kafka (OSAK) API
//
// Use Oracle Streaming with Apache Kafka Control Plane API to create/update/delete managed Kafka clusters.
//

package managedkafka

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// PublicConnectivityAddon The data that represents a Public Connectivity Addon
type PublicConnectivityAddon struct {

	// A unique user-friendly name.
	Name *string `mandatory:"true" json:"name"`

	// The bootstrap url of the kafka cluster.
	BootstrapUrl *string `mandatory:"true" json:"bootstrapUrl"`

	// A list of CIDR ranges for ingress/egress traffic.
	NetworkCidrs []string `mandatory:"true" json:"networkCidrs"`

	// Description of the add on
	Description *string `mandatory:"false" json:"description"`

	// The time the addon was created.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The time the addon was updated.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The current state of the KafkaCluster.
	LifecycleState KafkaClusterAddonLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Authentication mechanism.
	AuthenticationMechanism AuthenticationMechanismEnum `mandatory:"true" json:"authenticationMechanism"`
}

// GetName returns Name
func (m PublicConnectivityAddon) GetName() *string {
	return m.Name
}

// GetDescription returns Description
func (m PublicConnectivityAddon) GetDescription() *string {
	return m.Description
}

// GetTimeCreated returns TimeCreated
func (m PublicConnectivityAddon) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m PublicConnectivityAddon) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetLifecycleState returns LifecycleState
func (m PublicConnectivityAddon) GetLifecycleState() KafkaClusterAddonLifecycleStateEnum {
	return m.LifecycleState
}

func (m PublicConnectivityAddon) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PublicConnectivityAddon) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingKafkaClusterAddonLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetKafkaClusterAddonLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAuthenticationMechanismEnum(string(m.AuthenticationMechanism)); !ok && m.AuthenticationMechanism != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AuthenticationMechanism: %s. Supported values are: %s.", m.AuthenticationMechanism, strings.Join(GetAuthenticationMechanismEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m PublicConnectivityAddon) MarshalJSON() (buff []byte, e error) {
	type MarshalTypePublicConnectivityAddon PublicConnectivityAddon
	s := struct {
		DiscriminatorParam string `json:"addonType"`
		MarshalTypePublicConnectivityAddon
	}{
		"PUBLICCONNECTIVITY",
		(MarshalTypePublicConnectivityAddon)(m),
	}

	return json.Marshal(&s)
}
