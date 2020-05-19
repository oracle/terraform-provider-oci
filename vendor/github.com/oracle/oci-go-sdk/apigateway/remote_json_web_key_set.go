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

// RemoteJsonWebKeySet A set of public keys that is retrieved at run-time from a remote location
// to verify the JWT signature. The set should only contain JWK-formatted
// keys.
type RemoteJsonWebKeySet struct {

	// The uri from which to retrieve the key. It must be accessible
	// without authentication.
	Uri *string `mandatory:"true" json:"uri"`

	// Defines whether or not to uphold SSL verification.
	IsSslVerifyDisabled *bool `mandatory:"false" json:"isSslVerifyDisabled"`

	// The duration for which the JWKS should be cached before it is
	// fetched again.
	MaxCacheDurationInHours *int `mandatory:"false" json:"maxCacheDurationInHours"`
}

func (m RemoteJsonWebKeySet) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m RemoteJsonWebKeySet) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeRemoteJsonWebKeySet RemoteJsonWebKeySet
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeRemoteJsonWebKeySet
	}{
		"REMOTE_JWKS",
		(MarshalTypeRemoteJsonWebKeySet)(m),
	}

	return json.Marshal(&s)
}
