// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
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

// DpHost Configuration details of a hcs dp host.
// **Caution:** Oracle recommends that you avoid using any confidential information when you supply string values using the API.
type DpHost struct {

	// The host Id for the hsc dp host
	//      Example: `10.241.46.137`
	HostId *string `mandatory:"true" json:"hostId"`

	// The name for the hsc dp host
	// Example: `vcndp-vnicaas-02301.node.ad2.r1`
	DpHostName *string `mandatory:"false" json:"dpHostName"`

	// The pod Id of the hsc dp host
	// Example: `172.24.255.59`
	PodId *string `mandatory:"false" json:"podId"`

	// The tor Id of the hsc dp host
	// Example: `10.241.46.129`
	TorId *string `mandatory:"false" json:"torId"`

	// The current state of the hsc dp host
	// Example: `Quarantined`
	State *string `mandatory:"false" json:"state"`

	// Host Capacity.
	HostCapacity *int `mandatory:"false" json:"hostCapacity"`

	// The date and time the host was heartbeated, in the format defined by RFC3339.
	// Example: `2020-05-01T18:28:11+00:00`
	TimeStampHeartBeat *common.SDKTime `mandatory:"false" json:"timeStampHeartBeat"`

	// The date and time the data was created, in the format defined by RFC3339.
	// Example: `2020-05-01T18:28:11+00:00`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The date and time the data was created, in the format defined by RFC3339.
	// Example: `2020-05-01T18:28:11+00:00`
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`
}

func (m DpHost) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DpHost) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
