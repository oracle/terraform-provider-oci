// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Full Stack Disaster Recovery API
//
// Use the Full Stack Disaster Recovery (DR) API to manage disaster recovery for business applications.
// Full Stack DR is an OCI disaster recovery orchestration and management service that provides comprehensive disaster
// recovery capabilities for all layers of an application stack, including infrastructure, middleware, database,
// and application.
//

package disasterrecovery

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DrProtectionGroupMemberNetworkLoadBalancer The properties for a network load balancer member of a DR protection group.
type DrProtectionGroupMemberNetworkLoadBalancer struct {

	// The OCID of the member.
	// Example: `ocid1.instance.oc1..uniqueID`
	MemberId *string `mandatory:"true" json:"memberId"`

	// The OCID of the destination network load balancer.
	// The backend sets in this destination network load balancer are updated during DR.
	// Example: `ocid1.networkloadbalancer.oc1..uniqueID`
	DestinationNetworkLoadBalancerId *string `mandatory:"false" json:"destinationNetworkLoadBalancerId"`

	// A list of backend set mappings that are used to transfer or update backends during DR.
	BackendSetMappings []NetworkLoadBalancerBackendSetMapping `mandatory:"false" json:"backendSetMappings"`
}

// GetMemberId returns MemberId
func (m DrProtectionGroupMemberNetworkLoadBalancer) GetMemberId() *string {
	return m.MemberId
}

func (m DrProtectionGroupMemberNetworkLoadBalancer) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DrProtectionGroupMemberNetworkLoadBalancer) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m DrProtectionGroupMemberNetworkLoadBalancer) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDrProtectionGroupMemberNetworkLoadBalancer DrProtectionGroupMemberNetworkLoadBalancer
	s := struct {
		DiscriminatorParam string `json:"memberType"`
		MarshalTypeDrProtectionGroupMemberNetworkLoadBalancer
	}{
		"NETWORK_LOAD_BALANCER",
		(MarshalTypeDrProtectionGroupMemberNetworkLoadBalancer)(m),
	}

	return json.Marshal(&s)
}
