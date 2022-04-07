// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Streaming API
//
// Use the Streaming API to produce and consume messages, create streams and stream pools, and manage related items. For more information, see Streaming (https://docs.cloud.oracle.com/Content/Streaming/Concepts/streamingoverview.htm).
//

package streaming

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// PrivateEndpointDetails Optional parameters if a private stream pool is requested.
type PrivateEndpointDetails struct {

	// If specified, the stream pool will be private and only accessible from inside that subnet.
	// Producing-to and consuming-from a stream inside a private stream pool can also only be done from inside the subnet.
	// That value cannot be changed.
	SubnetId *string `mandatory:"false" json:"subnetId"`

	// The optional private IP you want to be associated with your private stream pool.
	// That parameter can only be specified when the subnetId parameter is set. It cannot be changed.
	// The private IP needs to be part of the CIDR range of the specified subnetId or the creation will fail.
	// If not specified a random IP inside the subnet will be chosen.
	// After the stream pool is created, a custom FQDN, pointing to this private IP, is created.
	// The FQDN is then used to access the service instead of the private IP.
	PrivateEndpointIp *string `mandatory:"false" json:"privateEndpointIp"`

	// The optional list of network security groups to be used with the private endpoint of the stream pool.
	// That value cannot be changed.
	NsgIds []string `mandatory:"false" json:"nsgIds"`
}

func (m PrivateEndpointDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PrivateEndpointDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
