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

// IngressGatewayAccessPolicyTarget Ingress gateway target that virtual services in mesh receive traffic from.
type IngressGatewayAccessPolicyTarget struct {

	// The OCID of the ingress gateway resource.
	IngressGatewayId *string `mandatory:"true" json:"ingressGatewayId"`
}

func (m IngressGatewayAccessPolicyTarget) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m IngressGatewayAccessPolicyTarget) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m IngressGatewayAccessPolicyTarget) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeIngressGatewayAccessPolicyTarget IngressGatewayAccessPolicyTarget
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeIngressGatewayAccessPolicyTarget
	}{
		"INGRESS_GATEWAY",
		(MarshalTypeIngressGatewayAccessPolicyTarget)(m),
	}

	return json.Marshal(&s)
}
