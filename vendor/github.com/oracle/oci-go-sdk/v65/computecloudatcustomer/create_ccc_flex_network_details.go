// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Compute Cloud@Customer API
//
// Use the Compute Cloud@Customer API to manage Compute Cloud@Customer infrastructures and upgrade schedules.
// For more information see Compute Cloud@Customer documentation (https://docs.oracle.com/iaas/compute-cloud-at-customer/home.htm).
//

package computecloudatcustomer

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateCccFlexNetworkDetails The configuration details for creating Compute Cloud@Customer flexNetwork.
type CreateCccFlexNetworkDetails struct {

	// The name that will be used to display the Compute Cloud@Customer flexNetwork
	// in the Oracle Cloud Console. Does not have to be unique and can be changed.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The Compute Cloud@Customer Infrastructure OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm),
	// which is immutable on creation.
	InfrastructureId *string `mandatory:"true" json:"infrastructureId"`

	// The defined CIDR range for the storage device network connection.  This will accommodate IPv4 and IPv6
	// addresses.  IPv4 min is 0.0.0.0 = 7, but the equivalent IPv6 min is :: = 2.  The max including a prefix
	// notation is 1111:1111:1111:1111:1111:1111:1111:1111/128 = 43.
	NetworkCidr *string `mandatory:"true" json:"networkCidr"`

	// An IP address defined in the networkCidr CIDR range utilized to reach the subnet where the VM resides.
	// This will accommodate IPv4 and IPv6 addresses.
	SpineVip *string `mandatory:"true" json:"spineVip"`

	// An IP address within the networkCidr CIDR range configured on the infrastructure spine switch 1 directly
	// attached to the subnet on the storage device. This will accommodate IPv4 and IPv6 addresses.
	Spine1Ip *string `mandatory:"true" json:"spine1Ip"`

	// An IP address within the networkCidr CIDR range configured on the infrastructure spine switch 2 directly
	// attached to the subnet on the storage device. This will accommodate IPv4 and IPv6 addresses.
	Spine2Ip *string `mandatory:"true" json:"spine2Ip"`

	// An array of switch ports to be used to establish the network connection.  The ports are represented
	// using a string to support breakout mode.  As an example if switch port 7 is a 100Gbps port, it can
	// be configured in breakout mode as four (4) smaller 25Gbps ports, 7/1, 7/2, 7/3 and 7/4.  The minimum
	// entry will be a single port X, where the maximum entry is XX/X.  This is a required field. The
	// minimum number of entries is 1 while the maximum number of entries is 4.
	Ports []string `mandatory:"true" json:"ports"`

	// A mutable client-meaningful text description of the Compute Cloud@Customer flexNetwork.
	// Avoid entering confidential information.
	Description *string `mandatory:"false" json:"description"`

	// The virtual local area network.  This value is optional.
	Vlan *int `mandatory:"false" json:"vlan"`

	// The gateway address for this network. This will accommodate IPv4 and IPv6 addresses.
	// This value is optional.
	GatewayIp *string `mandatory:"false" json:"gatewayIp"`

	// The speed of the network connection in gigabits per second.  This is optional.
	SpeedInGbps *int `mandatory:"false" json:"speedInGbps"`

	// When set to true, this optional field indicates that other resources within the on-premises
	// network can be allowed to connect to this network through the spine switches of the appliance.
	IsNetworkAdvertised *bool `mandatory:"false" json:"isNetworkAdvertised"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateCccFlexNetworkDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateCccFlexNetworkDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
