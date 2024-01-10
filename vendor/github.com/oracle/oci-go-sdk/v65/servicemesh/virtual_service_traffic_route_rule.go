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

// VirtualServiceTrafficRouteRule Rule for routing incoming virtual service traffic to a version.
type VirtualServiceTrafficRouteRule interface {

	// The destination of the request.
	GetDestinations() []VirtualDeploymentTrafficRuleTarget
}

type virtualservicetrafficrouterule struct {
	JsonData     []byte
	Destinations []VirtualDeploymentTrafficRuleTarget `mandatory:"true" json:"destinations"`
	Type         string                               `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *virtualservicetrafficrouterule) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalervirtualservicetrafficrouterule virtualservicetrafficrouterule
	s := struct {
		Model Unmarshalervirtualservicetrafficrouterule
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Destinations = s.Model.Destinations
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *virtualservicetrafficrouterule) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "TCP":
		mm := TcpVirtualServiceTrafficRouteRule{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "TLS_PASSTHROUGH":
		mm := TlsPassthroughVirtualServiceTrafficRouteRule{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "HTTP":
		mm := HttpVirtualServiceTrafficRouteRule{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for VirtualServiceTrafficRouteRule: %s.", m.Type)
		return *m, nil
	}
}

// GetDestinations returns Destinations
func (m virtualservicetrafficrouterule) GetDestinations() []VirtualDeploymentTrafficRuleTarget {
	return m.Destinations
}

func (m virtualservicetrafficrouterule) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m virtualservicetrafficrouterule) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// VirtualServiceTrafficRouteRuleTypeEnum Enum with underlying type: string
type VirtualServiceTrafficRouteRuleTypeEnum string

// Set of constants representing the allowable values for VirtualServiceTrafficRouteRuleTypeEnum
const (
	VirtualServiceTrafficRouteRuleTypeHttp           VirtualServiceTrafficRouteRuleTypeEnum = "HTTP"
	VirtualServiceTrafficRouteRuleTypeTlsPassthrough VirtualServiceTrafficRouteRuleTypeEnum = "TLS_PASSTHROUGH"
	VirtualServiceTrafficRouteRuleTypeTcp            VirtualServiceTrafficRouteRuleTypeEnum = "TCP"
)

var mappingVirtualServiceTrafficRouteRuleTypeEnum = map[string]VirtualServiceTrafficRouteRuleTypeEnum{
	"HTTP":            VirtualServiceTrafficRouteRuleTypeHttp,
	"TLS_PASSTHROUGH": VirtualServiceTrafficRouteRuleTypeTlsPassthrough,
	"TCP":             VirtualServiceTrafficRouteRuleTypeTcp,
}

var mappingVirtualServiceTrafficRouteRuleTypeEnumLowerCase = map[string]VirtualServiceTrafficRouteRuleTypeEnum{
	"http":            VirtualServiceTrafficRouteRuleTypeHttp,
	"tls_passthrough": VirtualServiceTrafficRouteRuleTypeTlsPassthrough,
	"tcp":             VirtualServiceTrafficRouteRuleTypeTcp,
}

// GetVirtualServiceTrafficRouteRuleTypeEnumValues Enumerates the set of values for VirtualServiceTrafficRouteRuleTypeEnum
func GetVirtualServiceTrafficRouteRuleTypeEnumValues() []VirtualServiceTrafficRouteRuleTypeEnum {
	values := make([]VirtualServiceTrafficRouteRuleTypeEnum, 0)
	for _, v := range mappingVirtualServiceTrafficRouteRuleTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetVirtualServiceTrafficRouteRuleTypeEnumStringValues Enumerates the set of values in String for VirtualServiceTrafficRouteRuleTypeEnum
func GetVirtualServiceTrafficRouteRuleTypeEnumStringValues() []string {
	return []string{
		"HTTP",
		"TLS_PASSTHROUGH",
		"TCP",
	}
}

// GetMappingVirtualServiceTrafficRouteRuleTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingVirtualServiceTrafficRouteRuleTypeEnum(val string) (VirtualServiceTrafficRouteRuleTypeEnum, bool) {
	enum, ok := mappingVirtualServiceTrafficRouteRuleTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
