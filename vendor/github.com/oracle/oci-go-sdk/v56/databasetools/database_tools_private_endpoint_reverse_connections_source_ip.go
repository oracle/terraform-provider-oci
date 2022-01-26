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

// DatabaseToolsPrivateEndpointReverseConnectionsSourceIp Source IP information for reverse connection configuration.
type DatabaseToolsPrivateEndpointReverseConnectionsSourceIp struct {

	// The IP address in the customer's VCN to be used as the source IP for reverse connection packets
	// traveling from the customer's VCN to the service's VCN.
	SourceIp *string `mandatory:"false" json:"sourceIp"`
}

func (m DatabaseToolsPrivateEndpointReverseConnectionsSourceIp) String() string {
	return common.PointerString(m)
}
