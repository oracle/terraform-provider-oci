// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Tools
//
// Database Tools APIs to manage Connections and Private Endpoints.
//

package databasetools

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// DatabaseToolsPrivateEndpointReverseConnectionConfiguration Reverse connection configuration details of Private Endpoint.
type DatabaseToolsPrivateEndpointReverseConnectionConfiguration struct {

	// A list of IP addresses in the customer VCN to be used as the source IPs for reverse connection packets
	// traveling from the service's VCN to the customer's VCN.
	ReverseConnectionsSourceIps []DatabaseToolsPrivateEndpointReverseConnectionsSourceIp `mandatory:"false" json:"reverseConnectionsSourceIps"`
}

func (m DatabaseToolsPrivateEndpointReverseConnectionConfiguration) String() string {
	return common.PointerString(m)
}
