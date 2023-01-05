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

// PostHeartBeatDetails Configuration details to registrate a hcs dp host.
// **Caution:** Oracle recommends that you avoid using any confidential information when you supply string values using the API.
type PostHeartBeatDetails struct {

	// The host Id for the hsc dp host
	//      Example: `10.241.46.137`
	HostId *string `mandatory:"false" json:"hostId"`

	// The ad of the hsc dp host
	// Example: `phx-ad-1`
	AvailabilityDomain *string `mandatory:"false" json:"availabilityDomain"`

	// Host Load.
	HostLoad *int `mandatory:"false" json:"hostLoad"`
}

func (m PostHeartBeatDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PostHeartBeatDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
