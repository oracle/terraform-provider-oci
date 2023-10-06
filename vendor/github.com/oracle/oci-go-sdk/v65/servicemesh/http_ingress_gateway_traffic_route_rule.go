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

// HttpIngressGatewayTrafficRouteRule Rule for routing incoming ingress gateway traffic with HTTP protocol
type HttpIngressGatewayTrafficRouteRule struct {

	// The destination of the request.
	Destinations []VirtualServiceTrafficRuleTarget `mandatory:"true" json:"destinations"`

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

	// The maximum duration in milliseconds for the upstream service to respond to a request.
	// If provided, the timeout value overrides the default timeout of 15 seconds for the HTTP based route rules, and disabled (no timeout) when 'isGrpc' is true.
	// The value 0 (zero) indicates that the timeout is disabled.
	// For streaming responses from the upstream service, consider either keeping the timeout disabled or set a sufficiently high value.
	RequestTimeoutInMs *int64 `mandatory:"false" json:"requestTimeoutInMs"`

	// Match type for the route
	PathType HttpIngressGatewayTrafficRouteRulePathTypeEnum `mandatory:"false" json:"pathType,omitempty"`
}

//GetIngressGatewayHost returns IngressGatewayHost
func (m HttpIngressGatewayTrafficRouteRule) GetIngressGatewayHost() *IngressGatewayHostRef {
	return m.IngressGatewayHost
}

//GetDestinations returns Destinations
func (m HttpIngressGatewayTrafficRouteRule) GetDestinations() []VirtualServiceTrafficRuleTarget {
	return m.Destinations
}

func (m HttpIngressGatewayTrafficRouteRule) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m HttpIngressGatewayTrafficRouteRule) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingHttpIngressGatewayTrafficRouteRulePathTypeEnum(string(m.PathType)); !ok && m.PathType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PathType: %s. Supported values are: %s.", m.PathType, strings.Join(GetHttpIngressGatewayTrafficRouteRulePathTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m HttpIngressGatewayTrafficRouteRule) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeHttpIngressGatewayTrafficRouteRule HttpIngressGatewayTrafficRouteRule
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeHttpIngressGatewayTrafficRouteRule
	}{
		"HTTP",
		(MarshalTypeHttpIngressGatewayTrafficRouteRule)(m),
	}

	return json.Marshal(&s)
}

// HttpIngressGatewayTrafficRouteRulePathTypeEnum Enum with underlying type: string
type HttpIngressGatewayTrafficRouteRulePathTypeEnum string

// Set of constants representing the allowable values for HttpIngressGatewayTrafficRouteRulePathTypeEnum
const (
	HttpIngressGatewayTrafficRouteRulePathTypePrefix HttpIngressGatewayTrafficRouteRulePathTypeEnum = "PREFIX"
)

var mappingHttpIngressGatewayTrafficRouteRulePathTypeEnum = map[string]HttpIngressGatewayTrafficRouteRulePathTypeEnum{
	"PREFIX": HttpIngressGatewayTrafficRouteRulePathTypePrefix,
}

var mappingHttpIngressGatewayTrafficRouteRulePathTypeEnumLowerCase = map[string]HttpIngressGatewayTrafficRouteRulePathTypeEnum{
	"prefix": HttpIngressGatewayTrafficRouteRulePathTypePrefix,
}

// GetHttpIngressGatewayTrafficRouteRulePathTypeEnumValues Enumerates the set of values for HttpIngressGatewayTrafficRouteRulePathTypeEnum
func GetHttpIngressGatewayTrafficRouteRulePathTypeEnumValues() []HttpIngressGatewayTrafficRouteRulePathTypeEnum {
	values := make([]HttpIngressGatewayTrafficRouteRulePathTypeEnum, 0)
	for _, v := range mappingHttpIngressGatewayTrafficRouteRulePathTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetHttpIngressGatewayTrafficRouteRulePathTypeEnumStringValues Enumerates the set of values in String for HttpIngressGatewayTrafficRouteRulePathTypeEnum
func GetHttpIngressGatewayTrafficRouteRulePathTypeEnumStringValues() []string {
	return []string{
		"PREFIX",
	}
}

// GetMappingHttpIngressGatewayTrafficRouteRulePathTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingHttpIngressGatewayTrafficRouteRulePathTypeEnum(val string) (HttpIngressGatewayTrafficRouteRulePathTypeEnum, bool) {
	enum, ok := mappingHttpIngressGatewayTrafficRouteRulePathTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
