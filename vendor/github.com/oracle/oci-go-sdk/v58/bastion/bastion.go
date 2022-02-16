// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Bastion API
//
// Oracle Cloud Infrastructure Bastion provides restricted and time-limited access to target resources that don't have public endpoints. Through the configuration of a bastion, you can let authorized users connect from specific IP addresses to target resources by way of Secure Shell (SSH) sessions hosted on the bastion.
//

package bastion

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// Bastion A bastion resource. A bastion provides secured, public access to target resources in the cloud that you cannot otherwise reach from the internet. A bastion resides in a public subnet and establishes the network infrastructure needed to connect a user to a target resource in a private subnet.
type Bastion struct {

	// The type of bastion.
	BastionType *string `mandatory:"true" json:"bastionType"`

	// The unique identifier (OCID) of the bastion, which can't be changed after creation.
	Id *string `mandatory:"true" json:"id"`

	// The name of the bastion, which can't be changed after creation.
	Name *string `mandatory:"true" json:"name"`

	// The unique identifier (OCID) of the compartment where the bastion is located.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The unique identifier (OCID) of the virtual cloud network (VCN) that the bastion connects to.
	TargetVcnId *string `mandatory:"true" json:"targetVcnId"`

	// The unique identifier (OCID) of the subnet that the bastion connects to.
	TargetSubnetId *string `mandatory:"true" json:"targetSubnetId"`

	// The maximum amount of time that any session on the bastion can remain active.
	MaxSessionTtlInSeconds *int `mandatory:"true" json:"maxSessionTtlInSeconds"`

	// The time the bastion was created. Format is defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2020-01-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The current state of the bastion.
	LifecycleState BastionLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The phonebook entry of the customer's team, which can't be changed after creation. Not applicable to `standard` bastions.
	PhoneBookEntry *string `mandatory:"false" json:"phoneBookEntry"`

	// A list of address ranges in CIDR notation that you want to allow to connect to sessions hosted by this bastion.
	ClientCidrBlockAllowList []string `mandatory:"false" json:"clientCidrBlockAllowList"`

	// A list of IP addresses of the hosts that the bastion has access to. Not applicable to `standard` bastions.
	StaticJumpHostIpAddresses []string `mandatory:"false" json:"staticJumpHostIpAddresses"`

	// The private IP address of the created private endpoint.
	PrivateEndpointIpAddress *string `mandatory:"false" json:"privateEndpointIpAddress"`

	// The maximum number of active sessions allowed on the bastion.
	MaxSessionsAllowed *int `mandatory:"false" json:"maxSessionsAllowed"`

	// The time the bastion was updated. Format is defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2020-01-25T21:10:29.600Z`
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// A message describing the current state in more detail.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m Bastion) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Bastion) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingBastionLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetBastionLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
