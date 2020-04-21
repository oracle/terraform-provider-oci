// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// API Gateway API
//
// API for the API Gateway service. Use this API to manage gateways, deployments, and related items.
// For more information, see
// Overview of API Gateway (https://docs.cloud.oracle.com/iaas/Content/APIGateway/Concepts/apigatewayoverview.htm).
//

package apigateway

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// CustomAuthenticationPolicy Use a function to validate a custom header or query parameter sent with the request authentication.
// A valid policy must specify either tokenHeader or tokenQueryParam.
type CustomAuthenticationPolicy struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Oracle Functions function resource.
	FunctionId *string `mandatory:"true" json:"functionId"`

	// Whether an unauthenticated user may access the API. Must be "true" to enable ANONYMOUS
	// route authorization.
	IsAnonymousAccessAllowed *bool `mandatory:"false" json:"isAnonymousAccessAllowed"`

	// The name of the header containing the authentication token.
	TokenHeader *string `mandatory:"false" json:"tokenHeader"`

	// The name of the query parameter containing the authentication token.
	TokenQueryParam *string `mandatory:"false" json:"tokenQueryParam"`
}

//GetIsAnonymousAccessAllowed returns IsAnonymousAccessAllowed
func (m CustomAuthenticationPolicy) GetIsAnonymousAccessAllowed() *bool {
	return m.IsAnonymousAccessAllowed
}

func (m CustomAuthenticationPolicy) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m CustomAuthenticationPolicy) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCustomAuthenticationPolicy CustomAuthenticationPolicy
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeCustomAuthenticationPolicy
	}{
		"CUSTOM_AUTHENTICATION",
		(MarshalTypeCustomAuthenticationPolicy)(m),
	}

	return json.Marshal(&s)
}
