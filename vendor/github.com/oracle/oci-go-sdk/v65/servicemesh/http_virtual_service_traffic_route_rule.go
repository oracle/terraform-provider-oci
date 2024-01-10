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

// HttpVirtualServiceTrafficRouteRule Rule for routing incoming Virtual Service traffic with HTTP protocol
type HttpVirtualServiceTrafficRouteRule struct {

	// The destination of the request.
	Destinations []VirtualDeploymentTrafficRuleTarget `mandatory:"true" json:"destinations"`

	// Route to match
	Path *string `mandatory:"false" json:"path"`

	// If true, the rule will check that the content-type header has a application/grpc
	// or one of the various application/grpc+ values.
	IsGrpc *bool `mandatory:"false" json:"isGrpc"`

	// The maximum duration in milliseconds for the target service to respond to a request.
	// If provided, the timeout value overrides the default timeout of 15 seconds for the HTTP based route rules, and disabled (no timeout) when 'isGrpc' is true.
	// The value 0 (zero) indicates that the timeout is disabled.
	// For streaming responses from the target service, consider either keeping the timeout disabled or set a sufficiently high value.
	RequestTimeoutInMs *int64 `mandatory:"false" json:"requestTimeoutInMs"`

	// Match type for the route
	PathType HttpVirtualServiceTrafficRouteRulePathTypeEnum `mandatory:"false" json:"pathType,omitempty"`
}

// GetDestinations returns Destinations
func (m HttpVirtualServiceTrafficRouteRule) GetDestinations() []VirtualDeploymentTrafficRuleTarget {
	return m.Destinations
}

func (m HttpVirtualServiceTrafficRouteRule) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m HttpVirtualServiceTrafficRouteRule) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingHttpVirtualServiceTrafficRouteRulePathTypeEnum(string(m.PathType)); !ok && m.PathType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PathType: %s. Supported values are: %s.", m.PathType, strings.Join(GetHttpVirtualServiceTrafficRouteRulePathTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m HttpVirtualServiceTrafficRouteRule) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeHttpVirtualServiceTrafficRouteRule HttpVirtualServiceTrafficRouteRule
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeHttpVirtualServiceTrafficRouteRule
	}{
		"HTTP",
		(MarshalTypeHttpVirtualServiceTrafficRouteRule)(m),
	}

	return json.Marshal(&s)
}

// HttpVirtualServiceTrafficRouteRulePathTypeEnum Enum with underlying type: string
type HttpVirtualServiceTrafficRouteRulePathTypeEnum string

// Set of constants representing the allowable values for HttpVirtualServiceTrafficRouteRulePathTypeEnum
const (
	HttpVirtualServiceTrafficRouteRulePathTypePrefix HttpVirtualServiceTrafficRouteRulePathTypeEnum = "PREFIX"
)

var mappingHttpVirtualServiceTrafficRouteRulePathTypeEnum = map[string]HttpVirtualServiceTrafficRouteRulePathTypeEnum{
	"PREFIX": HttpVirtualServiceTrafficRouteRulePathTypePrefix,
}

var mappingHttpVirtualServiceTrafficRouteRulePathTypeEnumLowerCase = map[string]HttpVirtualServiceTrafficRouteRulePathTypeEnum{
	"prefix": HttpVirtualServiceTrafficRouteRulePathTypePrefix,
}

// GetHttpVirtualServiceTrafficRouteRulePathTypeEnumValues Enumerates the set of values for HttpVirtualServiceTrafficRouteRulePathTypeEnum
func GetHttpVirtualServiceTrafficRouteRulePathTypeEnumValues() []HttpVirtualServiceTrafficRouteRulePathTypeEnum {
	values := make([]HttpVirtualServiceTrafficRouteRulePathTypeEnum, 0)
	for _, v := range mappingHttpVirtualServiceTrafficRouteRulePathTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetHttpVirtualServiceTrafficRouteRulePathTypeEnumStringValues Enumerates the set of values in String for HttpVirtualServiceTrafficRouteRulePathTypeEnum
func GetHttpVirtualServiceTrafficRouteRulePathTypeEnumStringValues() []string {
	return []string{
		"PREFIX",
	}
}

// GetMappingHttpVirtualServiceTrafficRouteRulePathTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingHttpVirtualServiceTrafficRouteRulePathTypeEnum(val string) (HttpVirtualServiceTrafficRouteRulePathTypeEnum, bool) {
	enum, ok := mappingHttpVirtualServiceTrafficRouteRulePathTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
