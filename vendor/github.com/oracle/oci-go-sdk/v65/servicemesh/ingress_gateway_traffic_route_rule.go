// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Service Mesh API
//
// Use the Service Mesh API to manage mesh, virtual service, access policy and other mesh related items.
//

package servicemesh

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// IngressGatewayTrafficRouteRule Rule for routing incoming ingress gateway traffic to a virtual service.
type IngressGatewayTrafficRouteRule interface {

	// The destination of the request.
	GetDestinations() []VirtualServiceTrafficRuleTarget

	GetIngressGatewayHost() *IngressGatewayHostRef
}

type ingressgatewaytrafficrouterule struct {
	JsonData           []byte
	IngressGatewayHost *IngressGatewayHostRef            `mandatory:"false" json:"ingressGatewayHost"`
	Destinations       []VirtualServiceTrafficRuleTarget `mandatory:"true" json:"destinations"`
	Type               string                            `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *ingressgatewaytrafficrouterule) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshaleringressgatewaytrafficrouterule ingressgatewaytrafficrouterule
	s := struct {
		Model Unmarshaleringressgatewaytrafficrouterule
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Destinations = s.Model.Destinations
	m.IngressGatewayHost = s.Model.IngressGatewayHost
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *ingressgatewaytrafficrouterule) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "HTTP":
		mm := HttpIngressGatewayTrafficRouteRule{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "TLS_PASSTHROUGH":
		mm := TlsPassthroughIngressGatewayTrafficRouteRule{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "TCP":
		mm := TcpIngressGatewayTrafficRouteRule{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for IngressGatewayTrafficRouteRule: %s.", m.Type)
		return *m, nil
	}
}

// GetIngressGatewayHost returns IngressGatewayHost
func (m ingressgatewaytrafficrouterule) GetIngressGatewayHost() *IngressGatewayHostRef {
	return m.IngressGatewayHost
}

// GetDestinations returns Destinations
func (m ingressgatewaytrafficrouterule) GetDestinations() []VirtualServiceTrafficRuleTarget {
	return m.Destinations
}

func (m ingressgatewaytrafficrouterule) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ingressgatewaytrafficrouterule) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// IngressGatewayTrafficRouteRuleTypeEnum Enum with underlying type: string
type IngressGatewayTrafficRouteRuleTypeEnum string

// Set of constants representing the allowable values for IngressGatewayTrafficRouteRuleTypeEnum
const (
	IngressGatewayTrafficRouteRuleTypeHttp           IngressGatewayTrafficRouteRuleTypeEnum = "HTTP"
	IngressGatewayTrafficRouteRuleTypeTlsPassthrough IngressGatewayTrafficRouteRuleTypeEnum = "TLS_PASSTHROUGH"
	IngressGatewayTrafficRouteRuleTypeTcp            IngressGatewayTrafficRouteRuleTypeEnum = "TCP"
)

var mappingIngressGatewayTrafficRouteRuleTypeEnum = map[string]IngressGatewayTrafficRouteRuleTypeEnum{
	"HTTP":            IngressGatewayTrafficRouteRuleTypeHttp,
	"TLS_PASSTHROUGH": IngressGatewayTrafficRouteRuleTypeTlsPassthrough,
	"TCP":             IngressGatewayTrafficRouteRuleTypeTcp,
}

var mappingIngressGatewayTrafficRouteRuleTypeEnumLowerCase = map[string]IngressGatewayTrafficRouteRuleTypeEnum{
	"http":            IngressGatewayTrafficRouteRuleTypeHttp,
	"tls_passthrough": IngressGatewayTrafficRouteRuleTypeTlsPassthrough,
	"tcp":             IngressGatewayTrafficRouteRuleTypeTcp,
}

// GetIngressGatewayTrafficRouteRuleTypeEnumValues Enumerates the set of values for IngressGatewayTrafficRouteRuleTypeEnum
func GetIngressGatewayTrafficRouteRuleTypeEnumValues() []IngressGatewayTrafficRouteRuleTypeEnum {
	values := make([]IngressGatewayTrafficRouteRuleTypeEnum, 0)
	for _, v := range mappingIngressGatewayTrafficRouteRuleTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetIngressGatewayTrafficRouteRuleTypeEnumStringValues Enumerates the set of values in String for IngressGatewayTrafficRouteRuleTypeEnum
func GetIngressGatewayTrafficRouteRuleTypeEnumStringValues() []string {
	return []string{
		"HTTP",
		"TLS_PASSTHROUGH",
		"TCP",
	}
}

// GetMappingIngressGatewayTrafficRouteRuleTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingIngressGatewayTrafficRouteRuleTypeEnum(val string) (IngressGatewayTrafficRouteRuleTypeEnum, bool) {
	enum, ok := mappingIngressGatewayTrafficRouteRuleTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
