// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// NetworkLoadBalancer API
//
// This describes the network load balancer API.
//

package networkloadbalancer

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// RegisterHealthCheckServiceInfraDpHostDetails Configuration details to registrate a hcs dp host.
// **Caution:** Oracle recommends that you avoid using any confidential information when you supply string values using the API.
type RegisterHealthCheckServiceInfraDpHostDetails struct {

	// The host Id for the hsc dp host
	//      Example: `10.241.46.137`
	HostId *string `mandatory:"false" json:"hostId"`

	// The name for the hsc dp host
	// Example: `vcndp-vnicaas-02301.node.ad2.r1`
	DpHostName *string `mandatory:"false" json:"dpHostName"`

	// The pod Id of the hsc dp host
	// Example: `172.24.255.59`
	PodId *string `mandatory:"false" json:"podId"`

	// The tor Id of the hsc dp host
	// Example: `10.241.46.129`
	TorId *string `mandatory:"false" json:"torId"`

	// The ad of the hsc dp host
	// Example: `phx-ad-1`
	AvailabilityDomain *string `mandatory:"false" json:"availabilityDomain"`

	// Host Capacity.
	HostCapacity *int `mandatory:"false" json:"hostCapacity"`
}

func (m RegisterHealthCheckServiceInfraDpHostDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RegisterHealthCheckServiceInfraDpHostDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
