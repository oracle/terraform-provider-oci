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

// HttpIngressGatewayTrafficRouteRuleDetails Rule for routing incoming ingress gateway traffic with HTTP protocol
type HttpIngressGatewayTrafficRouteRuleDetails struct {

	// The destination of the request.
	Destinations []VirtualServiceTrafficRuleTargetDetails `mandatory:"true" json:"destinations"`

	IngressGatewayHost *IngressGatewayHostRef `mandatory:"false" json:"ingressGatewayHost"`

	// Route to match
	Path *string `mandatory:"false" json:"path"`

	// If true, the rule will check that the content-type header has a application/grpc
	// or one of the various application/grpc+ values.
	IsGrpc *bool `mandatory:"false" json:"isGrpc"`

	// If true, the hostname will be rewritten to the target virtual deployment's DNS hostname.
	IsHostRewriteEnabled *bool `mandatory:"false" json:"isHostRewriteEnabled"`

	// If true, the matched path prefix will be rewritten to '/' before being directed to the target virtual deployment.
	IsPathRewriteEnabled *bool `mandatory:"false" json:"isPathRewriteEnabled"`

	// Match type for the route
	PathType HttpIngressGatewayTrafficRouteRuleDetailsPathTypeEnum `mandatory:"false" json:"pathType,omitempty"`
}

//GetIngressGatewayHost returns IngressGatewayHost
func (m HttpIngressGatewayTrafficRouteRuleDetails) GetIngressGatewayHost() *IngressGatewayHostRef {
	return m.IngressGatewayHost
}

//GetDestinations returns Destinations
func (m HttpIngressGatewayTrafficRouteRuleDetails) GetDestinations() []VirtualServiceTrafficRuleTargetDetails {
	return m.Destinations
}

func (m HttpIngressGatewayTrafficRouteRuleDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m HttpIngressGatewayTrafficRouteRuleDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingHttpIngressGatewayTrafficRouteRuleDetailsPathTypeEnum(string(m.PathType)); !ok && m.PathType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PathType: %s. Supported values are: %s.", m.PathType, strings.Join(GetHttpIngressGatewayTrafficRouteRuleDetailsPathTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m HttpIngressGatewayTrafficRouteRuleDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeHttpIngressGatewayTrafficRouteRuleDetails HttpIngressGatewayTrafficRouteRuleDetails
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeHttpIngressGatewayTrafficRouteRuleDetails
	}{
		"HTTP",
		(MarshalTypeHttpIngressGatewayTrafficRouteRuleDetails)(m),
	}

	return json.Marshal(&s)
}

// HttpIngressGatewayTrafficRouteRuleDetailsPathTypeEnum Enum with underlying type: string
type HttpIngressGatewayTrafficRouteRuleDetailsPathTypeEnum string

// Set of constants representing the allowable values for HttpIngressGatewayTrafficRouteRuleDetailsPathTypeEnum
const (
	HttpIngressGatewayTrafficRouteRuleDetailsPathTypePrefix HttpIngressGatewayTrafficRouteRuleDetailsPathTypeEnum = "PREFIX"
)

var mappingHttpIngressGatewayTrafficRouteRuleDetailsPathTypeEnum = map[string]HttpIngressGatewayTrafficRouteRuleDetailsPathTypeEnum{
	"PREFIX": HttpIngressGatewayTrafficRouteRuleDetailsPathTypePrefix,
}

var mappingHttpIngressGatewayTrafficRouteRuleDetailsPathTypeEnumLowerCase = map[string]HttpIngressGatewayTrafficRouteRuleDetailsPathTypeEnum{
	"prefix": HttpIngressGatewayTrafficRouteRuleDetailsPathTypePrefix,
}

// GetHttpIngressGatewayTrafficRouteRuleDetailsPathTypeEnumValues Enumerates the set of values for HttpIngressGatewayTrafficRouteRuleDetailsPathTypeEnum
func GetHttpIngressGatewayTrafficRouteRuleDetailsPathTypeEnumValues() []HttpIngressGatewayTrafficRouteRuleDetailsPathTypeEnum {
	values := make([]HttpIngressGatewayTrafficRouteRuleDetailsPathTypeEnum, 0)
	for _, v := range mappingHttpIngressGatewayTrafficRouteRuleDetailsPathTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetHttpIngressGatewayTrafficRouteRuleDetailsPathTypeEnumStringValues Enumerates the set of values in String for HttpIngressGatewayTrafficRouteRuleDetailsPathTypeEnum
func GetHttpIngressGatewayTrafficRouteRuleDetailsPathTypeEnumStringValues() []string {
	return []string{
		"PREFIX",
	}
}

// GetMappingHttpIngressGatewayTrafficRouteRuleDetailsPathTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingHttpIngressGatewayTrafficRouteRuleDetailsPathTypeEnum(val string) (HttpIngressGatewayTrafficRouteRuleDetailsPathTypeEnum, bool) {
	enum, ok := mappingHttpIngressGatewayTrafficRouteRuleDetailsPathTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
