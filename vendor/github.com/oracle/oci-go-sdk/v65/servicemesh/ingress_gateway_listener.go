// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
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

// IngressGatewayListener Listener configuration.
type IngressGatewayListener struct {

	// Type of protocol used.
	Protocol IngressGatewayListenerProtocolEnum `mandatory:"true" json:"protocol"`

	// Port on which ingress gateway is listening.
	Port *int `mandatory:"true" json:"port"`

	Tls *IngressListenerTlsConfig `mandatory:"false" json:"tls"`
}

func (m IngressGatewayListener) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m IngressGatewayListener) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingIngressGatewayListenerProtocolEnum(string(m.Protocol)); !ok && m.Protocol != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Protocol: %s. Supported values are: %s.", m.Protocol, strings.Join(GetIngressGatewayListenerProtocolEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// IngressGatewayListenerProtocolEnum Enum with underlying type: string
type IngressGatewayListenerProtocolEnum string

// Set of constants representing the allowable values for IngressGatewayListenerProtocolEnum
const (
	IngressGatewayListenerProtocolHttp           IngressGatewayListenerProtocolEnum = "HTTP"
	IngressGatewayListenerProtocolTlsPassthrough IngressGatewayListenerProtocolEnum = "TLS_PASSTHROUGH"
	IngressGatewayListenerProtocolTcp            IngressGatewayListenerProtocolEnum = "TCP"
)

var mappingIngressGatewayListenerProtocolEnum = map[string]IngressGatewayListenerProtocolEnum{
	"HTTP":            IngressGatewayListenerProtocolHttp,
	"TLS_PASSTHROUGH": IngressGatewayListenerProtocolTlsPassthrough,
	"TCP":             IngressGatewayListenerProtocolTcp,
}

var mappingIngressGatewayListenerProtocolEnumLowerCase = map[string]IngressGatewayListenerProtocolEnum{
	"http":            IngressGatewayListenerProtocolHttp,
	"tls_passthrough": IngressGatewayListenerProtocolTlsPassthrough,
	"tcp":             IngressGatewayListenerProtocolTcp,
}

// GetIngressGatewayListenerProtocolEnumValues Enumerates the set of values for IngressGatewayListenerProtocolEnum
func GetIngressGatewayListenerProtocolEnumValues() []IngressGatewayListenerProtocolEnum {
	values := make([]IngressGatewayListenerProtocolEnum, 0)
	for _, v := range mappingIngressGatewayListenerProtocolEnum {
		values = append(values, v)
	}
	return values
}

// GetIngressGatewayListenerProtocolEnumStringValues Enumerates the set of values in String for IngressGatewayListenerProtocolEnum
func GetIngressGatewayListenerProtocolEnumStringValues() []string {
	return []string{
		"HTTP",
		"TLS_PASSTHROUGH",
		"TCP",
	}
}

// GetMappingIngressGatewayListenerProtocolEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingIngressGatewayListenerProtocolEnum(val string) (IngressGatewayListenerProtocolEnum, bool) {
	enum, ok := mappingIngressGatewayListenerProtocolEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
