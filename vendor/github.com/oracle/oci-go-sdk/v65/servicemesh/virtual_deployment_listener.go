// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Service Mesh API
//
// Use the Service Mesh API to manage mesh, virtual service, access policy and other mesh related items.
//

package servicemesh

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// VirtualDeploymentListener Listener configuration for a virtual deployment.
type VirtualDeploymentListener struct {

	// Type of protocol used in virtual deployment.
	Protocol VirtualDeploymentListenerProtocolEnum `mandatory:"true" json:"protocol"`

	// Port in which virtual deployment is running.
	Port *int `mandatory:"true" json:"port"`

	// The maximum duration in milliseconds for the deployed service to respond to an incoming request through the listener.
	// If provided, the timeout value overrides the default timeout of 15 seconds for the HTTP/HTTP2 listeners, and disabled (no timeout) for the GRPC listeners. The value 0 (zero) indicates that the timeout is disabled.
	// The timeout cannot be configured for the TCP and TLS_PASSTHROUGH listeners.
	// For streaming responses from the deployed service, consider either keeping the timeout disabled or set a sufficiently high value.
	RequestTimeoutInMs *int64 `mandatory:"false" json:"requestTimeoutInMs"`

	// The maximum duration in milliseconds for which the request's stream may be idle. The value 0 (zero) indicates that the timeout is disabled.
	IdleTimeoutInMs *int64 `mandatory:"false" json:"idleTimeoutInMs"`
}

func (m VirtualDeploymentListener) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m VirtualDeploymentListener) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingVirtualDeploymentListenerProtocolEnum(string(m.Protocol)); !ok && m.Protocol != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Protocol: %s. Supported values are: %s.", m.Protocol, strings.Join(GetVirtualDeploymentListenerProtocolEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// VirtualDeploymentListenerProtocolEnum Enum with underlying type: string
type VirtualDeploymentListenerProtocolEnum string

// Set of constants representing the allowable values for VirtualDeploymentListenerProtocolEnum
const (
	VirtualDeploymentListenerProtocolHttp           VirtualDeploymentListenerProtocolEnum = "HTTP"
	VirtualDeploymentListenerProtocolTlsPassthrough VirtualDeploymentListenerProtocolEnum = "TLS_PASSTHROUGH"
	VirtualDeploymentListenerProtocolTcp            VirtualDeploymentListenerProtocolEnum = "TCP"
	VirtualDeploymentListenerProtocolHttp2          VirtualDeploymentListenerProtocolEnum = "HTTP2"
	VirtualDeploymentListenerProtocolGrpc           VirtualDeploymentListenerProtocolEnum = "GRPC"
)

var mappingVirtualDeploymentListenerProtocolEnum = map[string]VirtualDeploymentListenerProtocolEnum{
	"HTTP":            VirtualDeploymentListenerProtocolHttp,
	"TLS_PASSTHROUGH": VirtualDeploymentListenerProtocolTlsPassthrough,
	"TCP":             VirtualDeploymentListenerProtocolTcp,
	"HTTP2":           VirtualDeploymentListenerProtocolHttp2,
	"GRPC":            VirtualDeploymentListenerProtocolGrpc,
}

var mappingVirtualDeploymentListenerProtocolEnumLowerCase = map[string]VirtualDeploymentListenerProtocolEnum{
	"http":            VirtualDeploymentListenerProtocolHttp,
	"tls_passthrough": VirtualDeploymentListenerProtocolTlsPassthrough,
	"tcp":             VirtualDeploymentListenerProtocolTcp,
	"http2":           VirtualDeploymentListenerProtocolHttp2,
	"grpc":            VirtualDeploymentListenerProtocolGrpc,
}

// GetVirtualDeploymentListenerProtocolEnumValues Enumerates the set of values for VirtualDeploymentListenerProtocolEnum
func GetVirtualDeploymentListenerProtocolEnumValues() []VirtualDeploymentListenerProtocolEnum {
	values := make([]VirtualDeploymentListenerProtocolEnum, 0)
	for _, v := range mappingVirtualDeploymentListenerProtocolEnum {
		values = append(values, v)
	}
	return values
}

// GetVirtualDeploymentListenerProtocolEnumStringValues Enumerates the set of values in String for VirtualDeploymentListenerProtocolEnum
func GetVirtualDeploymentListenerProtocolEnumStringValues() []string {
	return []string{
		"HTTP",
		"TLS_PASSTHROUGH",
		"TCP",
		"HTTP2",
		"GRPC",
	}
}

// GetMappingVirtualDeploymentListenerProtocolEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingVirtualDeploymentListenerProtocolEnum(val string) (VirtualDeploymentListenerProtocolEnum, bool) {
	enum, ok := mappingVirtualDeploymentListenerProtocolEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
