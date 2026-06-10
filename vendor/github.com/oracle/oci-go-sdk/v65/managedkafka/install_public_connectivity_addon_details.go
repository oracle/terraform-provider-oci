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

// InstallPublicConnectivityAddonDetails The data to install a KafkaClusterAddon.
type InstallPublicConnectivityAddonDetails struct {

	// A unique user-friendly name. Avoid entering confidential information.
	Name *string `mandatory:"true" json:"name"`

	// A list of CIDR's for ingress/egress traffic.
	NetworkCidrs []string `mandatory:"true" json:"networkCidrs"`

	// A brief description of the add on being installed.
	Description *string `mandatory:"false" json:"description"`

	// Authentication mechanism.
	AuthenticationMechanism AuthenticationMechanismEnum `mandatory:"true" json:"authenticationMechanism"`
}

// GetName returns Name
func (m InstallPublicConnectivityAddonDetails) GetName() *string {
	return m.Name
}

// GetDescription returns Description
func (m InstallPublicConnectivityAddonDetails) GetDescription() *string {
	return m.Description
}

func (m InstallPublicConnectivityAddonDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m InstallPublicConnectivityAddonDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingAuthenticationMechanismEnum(string(m.AuthenticationMechanism)); !ok && m.AuthenticationMechanism != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AuthenticationMechanism: %s. Supported values are: %s.", m.AuthenticationMechanism, strings.Join(GetAuthenticationMechanismEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m InstallPublicConnectivityAddonDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeInstallPublicConnectivityAddonDetails InstallPublicConnectivityAddonDetails
	s := struct {
		DiscriminatorParam string `json:"addonType"`
		MarshalTypeInstallPublicConnectivityAddonDetails
	}{
		"PUBLICCONNECTIVITY",
		(MarshalTypeInstallPublicConnectivityAddonDetails)(m),
	}

	return json.Marshal(&s)
}
