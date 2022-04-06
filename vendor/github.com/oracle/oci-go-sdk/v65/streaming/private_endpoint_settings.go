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

// PrivateEndpointSettings Optional settings if the stream pool is private.
type PrivateEndpointSettings struct {

	// The subnet id from which the private stream pool can be accessed.
	// Trying to access the streams from another network location will result in an error.
	SubnetId *string `mandatory:"false" json:"subnetId"`

	// The private IP associated with the stream pool in the associated subnetId.
	// The stream pool's FQDN resolves to that IP and should be used - instead of the private IP - in order to not trigger any TLS issues.
	PrivateEndpointIp *string `mandatory:"false" json:"privateEndpointIp"`

	// The optional list of network security groups that are associated with the private endpoint of the stream pool.
	NsgIds []string `mandatory:"false" json:"nsgIds"`
}

func (m PrivateEndpointSettings) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PrivateEndpointSettings) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
