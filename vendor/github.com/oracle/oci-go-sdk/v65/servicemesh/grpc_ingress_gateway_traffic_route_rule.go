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

// GrpcIngressGatewayTrafficRouteRule Rule for routing incoming Virtual Service traffic with GRPC protocol
type GrpcIngressGatewayTrafficRouteRule struct {

	// The destination of the request.
	Destinations []VirtualServiceTrafficRuleTarget `mandatory:"true" json:"destinations"`

	IngressGatewayHost *IngressGatewayHostRef `mandatory:"false" json:"ingressGatewayHost"`

	// An object that represents the data to match from the request.
	Metadata []StringMatch `mandatory:"false" json:"metadata"`

	// The fully qualified domain name for the service to match from the request.
	ServiceName *string `mandatory:"false" json:"serviceName"`

	// The method name to match from the request. If you specify a methodName, you must also specify a serviceName.
	MethodName *string `mandatory:"false" json:"methodName"`

	// Port to match for routing traffic.
	Port *int `mandatory:"false" json:"port"`

	// If true, the hostname will be rewritten to the target virtual deployment's DNS hostname.
	IsHostRewriteEnabled *bool `mandatory:"false" json:"isHostRewriteEnabled"`
}

//GetIngressGatewayHost returns IngressGatewayHost
func (m GrpcIngressGatewayTrafficRouteRule) GetIngressGatewayHost() *IngressGatewayHostRef {
	return m.IngressGatewayHost
}

//GetDestinations returns Destinations
func (m GrpcIngressGatewayTrafficRouteRule) GetDestinations() []VirtualServiceTrafficRuleTarget {
	return m.Destinations
}

func (m GrpcIngressGatewayTrafficRouteRule) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m GrpcIngressGatewayTrafficRouteRule) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m GrpcIngressGatewayTrafficRouteRule) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeGrpcIngressGatewayTrafficRouteRule GrpcIngressGatewayTrafficRouteRule
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeGrpcIngressGatewayTrafficRouteRule
	}{
		"GRPC",
		(MarshalTypeGrpcIngressGatewayTrafficRouteRule)(m),
	}

	return json.Marshal(&s)
}
