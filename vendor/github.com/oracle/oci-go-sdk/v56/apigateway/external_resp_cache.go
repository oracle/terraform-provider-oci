// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
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
	"github.com/oracle/oci-go-sdk/v56/common"
)

// ExternalRespCache Connection details for an external RESP based cache store for Response Caching.
type ExternalRespCache struct {

	// The set of cache store members to connect to. At present only a single server is supported.
	Servers []ResponseCacheRespServer `mandatory:"true" json:"servers"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Oracle Vault Service secret resource.
	AuthenticationSecretId *string `mandatory:"true" json:"authenticationSecretId"`

	// The version number of the authentication secret to use.
	AuthenticationSecretVersionNumber *int64 `mandatory:"true" json:"authenticationSecretVersionNumber"`

	// Defines if the connection should be over SSL.
	IsSslEnabled *bool `mandatory:"false" json:"isSslEnabled"`

	// Defines whether or not to uphold SSL verification.
	IsSslVerifyDisabled *bool `mandatory:"false" json:"isSslVerifyDisabled"`

	// Defines the timeout for establishing a connection with the Response Cache.
	ConnectTimeoutInMs *int `mandatory:"false" json:"connectTimeoutInMs"`

	// Defines the timeout for reading data from the Response Cache.
	ReadTimeoutInMs *int `mandatory:"false" json:"readTimeoutInMs"`

	// Defines the timeout for transmitting data to the Response Cache.
	SendTimeoutInMs *int `mandatory:"false" json:"sendTimeoutInMs"`
}

func (m ExternalRespCache) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m ExternalRespCache) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeExternalRespCache ExternalRespCache
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeExternalRespCache
	}{
		"EXTERNAL_RESP_CACHE",
		(MarshalTypeExternalRespCache)(m),
	}

	return json.Marshal(&s)
}
