// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
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

// IngressGatewayTrafficRouteRuleDetails Rule for routing incoming ingress gateway traffic to a virtual service.
type IngressGatewayTrafficRouteRuleDetails interface {

	// The destination of the request.
	GetDestinations() []VirtualServiceTrafficRuleTargetDetails

	GetIngressGatewayHost() *IngressGatewayHostRef
}

type ingressgatewaytrafficrouteruledetails struct {
	JsonData           []byte
	Destinations       []VirtualServiceTrafficRuleTargetDetails `mandatory:"true" json:"destinations"`
	IngressGatewayHost *IngressGatewayHostRef                   `mandatory:"false" json:"ingressGatewayHost"`
	Type               string                                   `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *ingressgatewaytrafficrouteruledetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshaleringressgatewaytrafficrouteruledetails ingressgatewaytrafficrouteruledetails
	s := struct {
		Model Unmarshaleringressgatewaytrafficrouteruledetails
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
func (m *ingressgatewaytrafficrouteruledetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "TLS_PASSTHROUGH":
		mm := TlsPassthroughIngressGatewayTrafficRouteRuleDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "TCP":
		mm := TcpIngressGatewayTrafficRouteRuleDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "HTTP":
		mm := HttpIngressGatewayTrafficRouteRuleDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetDestinations returns Destinations
func (m ingressgatewaytrafficrouteruledetails) GetDestinations() []VirtualServiceTrafficRuleTargetDetails {
	return m.Destinations
}

//GetIngressGatewayHost returns IngressGatewayHost
func (m ingressgatewaytrafficrouteruledetails) GetIngressGatewayHost() *IngressGatewayHostRef {
	return m.IngressGatewayHost
}

func (m ingressgatewaytrafficrouteruledetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ingressgatewaytrafficrouteruledetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// IngressGatewayTrafficRouteRuleDetailsTypeEnum Enum with underlying type: string
type IngressGatewayTrafficRouteRuleDetailsTypeEnum string

// Set of constants representing the allowable values for IngressGatewayTrafficRouteRuleDetailsTypeEnum
const (
	IngressGatewayTrafficRouteRuleDetailsTypeHttp           IngressGatewayTrafficRouteRuleDetailsTypeEnum = "HTTP"
	IngressGatewayTrafficRouteRuleDetailsTypeTlsPassthrough IngressGatewayTrafficRouteRuleDetailsTypeEnum = "TLS_PASSTHROUGH"
	IngressGatewayTrafficRouteRuleDetailsTypeTcp            IngressGatewayTrafficRouteRuleDetailsTypeEnum = "TCP"
)

var mappingIngressGatewayTrafficRouteRuleDetailsTypeEnum = map[string]IngressGatewayTrafficRouteRuleDetailsTypeEnum{
	"HTTP":            IngressGatewayTrafficRouteRuleDetailsTypeHttp,
	"TLS_PASSTHROUGH": IngressGatewayTrafficRouteRuleDetailsTypeTlsPassthrough,
	"TCP":             IngressGatewayTrafficRouteRuleDetailsTypeTcp,
}

var mappingIngressGatewayTrafficRouteRuleDetailsTypeEnumLowerCase = map[string]IngressGatewayTrafficRouteRuleDetailsTypeEnum{
	"http":            IngressGatewayTrafficRouteRuleDetailsTypeHttp,
	"tls_passthrough": IngressGatewayTrafficRouteRuleDetailsTypeTlsPassthrough,
	"tcp":             IngressGatewayTrafficRouteRuleDetailsTypeTcp,
}

// GetIngressGatewayTrafficRouteRuleDetailsTypeEnumValues Enumerates the set of values for IngressGatewayTrafficRouteRuleDetailsTypeEnum
func GetIngressGatewayTrafficRouteRuleDetailsTypeEnumValues() []IngressGatewayTrafficRouteRuleDetailsTypeEnum {
	values := make([]IngressGatewayTrafficRouteRuleDetailsTypeEnum, 0)
	for _, v := range mappingIngressGatewayTrafficRouteRuleDetailsTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetIngressGatewayTrafficRouteRuleDetailsTypeEnumStringValues Enumerates the set of values in String for IngressGatewayTrafficRouteRuleDetailsTypeEnum
func GetIngressGatewayTrafficRouteRuleDetailsTypeEnumStringValues() []string {
	return []string{
		"HTTP",
		"TLS_PASSTHROUGH",
		"TCP",
	}
}

// GetMappingIngressGatewayTrafficRouteRuleDetailsTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingIngressGatewayTrafficRouteRuleDetailsTypeEnum(val string) (IngressGatewayTrafficRouteRuleDetailsTypeEnum, bool) {
	enum, ok := mappingIngressGatewayTrafficRouteRuleDetailsTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
