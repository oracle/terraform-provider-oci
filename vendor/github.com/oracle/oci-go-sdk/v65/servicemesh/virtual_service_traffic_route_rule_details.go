// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
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

// VirtualServiceTrafficRouteRuleDetails Rule for routing incoming virtual service traffic to a version.
type VirtualServiceTrafficRouteRuleDetails interface {

	// The destination of the request.
	GetDestinations() []VirtualDeploymentTrafficRuleTargetDetails
}

type virtualservicetrafficrouteruledetails struct {
	JsonData     []byte
	Destinations []VirtualDeploymentTrafficRuleTargetDetails `mandatory:"true" json:"destinations"`
	Type         string                                      `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *virtualservicetrafficrouteruledetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalervirtualservicetrafficrouteruledetails virtualservicetrafficrouteruledetails
	s := struct {
		Model Unmarshalervirtualservicetrafficrouteruledetails
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
func (m *virtualservicetrafficrouteruledetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "TCP":
		mm := TcpVirtualServiceTrafficRouteRuleDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "TLS_PASSTHROUGH":
		mm := TlsPassthroughVirtualServiceTrafficRouteRuleDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "HTTP":
		mm := HttpVirtualServiceTrafficRouteRuleDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetDestinations returns Destinations
func (m virtualservicetrafficrouteruledetails) GetDestinations() []VirtualDeploymentTrafficRuleTargetDetails {
	return m.Destinations
}

func (m virtualservicetrafficrouteruledetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m virtualservicetrafficrouteruledetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// VirtualServiceTrafficRouteRuleDetailsTypeEnum Enum with underlying type: string
type VirtualServiceTrafficRouteRuleDetailsTypeEnum string

// Set of constants representing the allowable values for VirtualServiceTrafficRouteRuleDetailsTypeEnum
const (
	VirtualServiceTrafficRouteRuleDetailsTypeHttp           VirtualServiceTrafficRouteRuleDetailsTypeEnum = "HTTP"
	VirtualServiceTrafficRouteRuleDetailsTypeTlsPassthrough VirtualServiceTrafficRouteRuleDetailsTypeEnum = "TLS_PASSTHROUGH"
	VirtualServiceTrafficRouteRuleDetailsTypeTcp            VirtualServiceTrafficRouteRuleDetailsTypeEnum = "TCP"
)

var mappingVirtualServiceTrafficRouteRuleDetailsTypeEnum = map[string]VirtualServiceTrafficRouteRuleDetailsTypeEnum{
	"HTTP":            VirtualServiceTrafficRouteRuleDetailsTypeHttp,
	"TLS_PASSTHROUGH": VirtualServiceTrafficRouteRuleDetailsTypeTlsPassthrough,
	"TCP":             VirtualServiceTrafficRouteRuleDetailsTypeTcp,
}

var mappingVirtualServiceTrafficRouteRuleDetailsTypeEnumLowerCase = map[string]VirtualServiceTrafficRouteRuleDetailsTypeEnum{
	"http":            VirtualServiceTrafficRouteRuleDetailsTypeHttp,
	"tls_passthrough": VirtualServiceTrafficRouteRuleDetailsTypeTlsPassthrough,
	"tcp":             VirtualServiceTrafficRouteRuleDetailsTypeTcp,
}

// GetVirtualServiceTrafficRouteRuleDetailsTypeEnumValues Enumerates the set of values for VirtualServiceTrafficRouteRuleDetailsTypeEnum
func GetVirtualServiceTrafficRouteRuleDetailsTypeEnumValues() []VirtualServiceTrafficRouteRuleDetailsTypeEnum {
	values := make([]VirtualServiceTrafficRouteRuleDetailsTypeEnum, 0)
	for _, v := range mappingVirtualServiceTrafficRouteRuleDetailsTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetVirtualServiceTrafficRouteRuleDetailsTypeEnumStringValues Enumerates the set of values in String for VirtualServiceTrafficRouteRuleDetailsTypeEnum
func GetVirtualServiceTrafficRouteRuleDetailsTypeEnumStringValues() []string {
	return []string{
		"HTTP",
		"TLS_PASSTHROUGH",
		"TCP",
	}
}

// GetMappingVirtualServiceTrafficRouteRuleDetailsTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingVirtualServiceTrafficRouteRuleDetailsTypeEnum(val string) (VirtualServiceTrafficRouteRuleDetailsTypeEnum, bool) {
	enum, ok := mappingVirtualServiceTrafficRouteRuleDetailsTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
