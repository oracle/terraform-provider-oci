// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Load Balancing API
//
// API for the Load Balancing service. Use this API to manage load balancers, backend sets, and related items. For more
// information, see Overview of Load Balancing (https://docs.cloud.oracle.com/iaas/Content/Balance/Concepts/balanceoverview.htm).
//

package loadbalancer

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// IpBasedMaxConnectionsRule An object that represents the action of configuring an IP max connection rule which specifies
// how many connections IPs can make to a listener.
type IpBasedMaxConnectionsRule struct {

	// The maximum number of connections that the any IP can make to a listener unless the IP is mentioned
	// in maxConnections. If no defaultMaxConnections is specified the default is unlimited.
	DefaultMaxConnections *int `mandatory:"false" json:"defaultMaxConnections"`

	// An array of IPs that have a maxConnection setting different than the default and what
	// that maxConnection setting is
	IpMaxConnections []IpMaxConnections `mandatory:"false" json:"ipMaxConnections"`
}

func (m IpBasedMaxConnectionsRule) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m IpBasedMaxConnectionsRule) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m IpBasedMaxConnectionsRule) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeIpBasedMaxConnectionsRule IpBasedMaxConnectionsRule
	s := struct {
		DiscriminatorParam string `json:"action"`
		MarshalTypeIpBasedMaxConnectionsRule
	}{
		"IP_BASED_MAX_CONNECTIONS",
		(MarshalTypeIpBasedMaxConnectionsRule)(m),
	}

	return json.Marshal(&s)
}
