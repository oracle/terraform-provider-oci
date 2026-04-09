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

// UpdatePublicConnectivityAddonDetails The data to update a KafkaClusterAddon.
type UpdatePublicConnectivityAddonDetails struct {

	// A list of CIDR ranges for ingress/egress traffic.
	NetworkCidrs []string `mandatory:"true" json:"networkCidrs"`

	// A unique user-friendly name. Avoid entering confidential information.
	Description *string `mandatory:"false" json:"description"`
}

// GetDescription returns Description
func (m UpdatePublicConnectivityAddonDetails) GetDescription() *string {
	return m.Description
}

func (m UpdatePublicConnectivityAddonDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdatePublicConnectivityAddonDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m UpdatePublicConnectivityAddonDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdatePublicConnectivityAddonDetails UpdatePublicConnectivityAddonDetails
	s := struct {
		DiscriminatorParam string `json:"addonType"`
		MarshalTypeUpdatePublicConnectivityAddonDetails
	}{
		"PUBLICCONNECTIVITY",
		(MarshalTypeUpdatePublicConnectivityAddonDetails)(m),
	}

	return json.Marshal(&s)
}
