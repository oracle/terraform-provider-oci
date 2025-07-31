// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// API Gateway API
//
// API for the API Gateway service. Use this API to manage gateways, deployments, and related items.
// For more information, see
// Overview of API Gateway (https://docs.oracle.com/iaas/Content/APIGateway/Concepts/apigatewayoverview.htm).
//

package apigateway

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// WebSocketProxyBackend Proxy a WebSocket session to a backend WebSocket server.
type WebSocketProxyBackend struct {

	// The url of the backend WebSocket server.
	Url *string `mandatory:"true" json:"url"`

	// Defines a timeout for establishing a connection with the backend server.
	ConnectTimeoutInSeconds *float32 `mandatory:"false" json:"connectTimeoutInSeconds"`

	// Defines the amount of time a WebSocket connection can be idle (i.e. no bytes received and no bytes sent).
	IdleTimeoutInSeconds *float32 `mandatory:"false" json:"idleTimeoutInSeconds"`

	// Defines whether or not to uphold SSL verification.
	IsSslVerifyDisabled *bool `mandatory:"false" json:"isSslVerifyDisabled"`
}

func (m WebSocketProxyBackend) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m WebSocketProxyBackend) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m WebSocketProxyBackend) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeWebSocketProxyBackend WebSocketProxyBackend
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeWebSocketProxyBackend
	}{
		"WEBSOCKET_PROXY_BACKEND",
		(MarshalTypeWebSocketProxyBackend)(m),
	}

	return json.Marshal(&s)
}
