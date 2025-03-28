// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// TcpIngressGatewayTrafficRouteRuleDetails Rule for routing incoming ingress gateway traffic with TCP protocol.
type TcpIngressGatewayTrafficRouteRuleDetails struct {

	// The destination of the request.
	Destinations []VirtualServiceTrafficRuleTargetDetails `mandatory:"true" json:"destinations"`

	IngressGatewayHost *IngressGatewayHostRef `mandatory:"false" json:"ingressGatewayHost"`
}

// GetIngressGatewayHost returns IngressGatewayHost
func (m TcpIngressGatewayTrafficRouteRuleDetails) GetIngressGatewayHost() *IngressGatewayHostRef {
	return m.IngressGatewayHost
}

// GetDestinations returns Destinations
func (m TcpIngressGatewayTrafficRouteRuleDetails) GetDestinations() []VirtualServiceTrafficRuleTargetDetails {
	return m.Destinations
}

func (m TcpIngressGatewayTrafficRouteRuleDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TcpIngressGatewayTrafficRouteRuleDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m TcpIngressGatewayTrafficRouteRuleDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeTcpIngressGatewayTrafficRouteRuleDetails TcpIngressGatewayTrafficRouteRuleDetails
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeTcpIngressGatewayTrafficRouteRuleDetails
	}{
		"TCP",
		(MarshalTypeTcpIngressGatewayTrafficRouteRuleDetails)(m),
	}

	return json.Marshal(&s)
}
