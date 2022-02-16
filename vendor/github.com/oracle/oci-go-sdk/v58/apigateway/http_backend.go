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
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// HttpBackend Send the request to an HTTP backend.
type HttpBackend struct {
	Url *string `mandatory:"true" json:"url"`

	// Defines a timeout for establishing a connection with a proxied server.
	ConnectTimeoutInSeconds *float32 `mandatory:"false" json:"connectTimeoutInSeconds"`

	// Defines a timeout for reading a response from the proxied server.
	ReadTimeoutInSeconds *float32 `mandatory:"false" json:"readTimeoutInSeconds"`

	// Defines a timeout for transmitting a request to the proxied server.
	SendTimeoutInSeconds *float32 `mandatory:"false" json:"sendTimeoutInSeconds"`

	// Defines whether or not to uphold SSL verification.
	IsSslVerifyDisabled *bool `mandatory:"false" json:"isSslVerifyDisabled"`
}

func (m HttpBackend) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m HttpBackend) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m HttpBackend) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeHttpBackend HttpBackend
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeHttpBackend
	}{
		"HTTP_BACKEND",
		(MarshalTypeHttpBackend)(m),
	}

	return json.Marshal(&s)
}
