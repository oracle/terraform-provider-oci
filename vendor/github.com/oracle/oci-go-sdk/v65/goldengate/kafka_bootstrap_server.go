// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// GoldenGate API
//
// Use the Oracle Cloud Infrastructure GoldenGate APIs to perform data replication operations.
//

package goldengate

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// KafkaBootstrapServer Represents a Kafka bootstrap server with host name, optional port defaults to 9092.
// Deprecated: privateIp is returned only for legacy deployments created with it.
// On update payloads, privateIp can only be set to empty (cleared).
type KafkaBootstrapServer struct {

	// The name or address of a host.
	Host *string `mandatory:"true" json:"host"`

	// The port of an endpoint usually specified for a connection.
	Port *int `mandatory:"false" json:"port"`

	// This property is not available when creating connections. For existing deprecated connections having this value set, the value cannot be updated; set it to empty.
	// For deprecated connections created with this field in the past, either the private IP had to be specified in the connectionString or host field, or the host name had to be resolvable in the target VCN.
	PrivateIp *string `mandatory:"false" json:"privateIp"`
}

func (m KafkaBootstrapServer) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m KafkaBootstrapServer) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
