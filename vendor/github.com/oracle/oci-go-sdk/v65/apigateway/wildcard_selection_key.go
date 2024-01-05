// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// WildcardSelectionKey When dynamically routing and dynamically authenticating requests, the route or authentication server associated with a selection key containing a wildcard is used if the context variable in an incoming request matches that key.
type WildcardSelectionKey struct {

	// Name assigned to the branch.
	Name *string `mandatory:"true" json:"name"`

	// A selection key string containing a wildcard to match with the context variable in an incoming request. If the context variable matches the string, the request is sent to the route or authentication server associated with the selection key. Valid wildcards are '*' (zero or more characters) and '+' (one or more characters). The string can only contain one wildcard, and the wildcard must be at the start or the end of the string.
	Expression *string `mandatory:"true" json:"expression"`

	// Specifies whether to use the route or authentication server associated with this selection key as the default. The default is used if the value of a context variable in an incoming request does not match any of the other selection key values when dynamically routing and dynamically authenticating requests.
	IsDefault *bool `mandatory:"false" json:"isDefault"`
}

// GetIsDefault returns IsDefault
func (m WildcardSelectionKey) GetIsDefault() *bool {
	return m.IsDefault
}

// GetName returns Name
func (m WildcardSelectionKey) GetName() *string {
	return m.Name
}

func (m WildcardSelectionKey) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m WildcardSelectionKey) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m WildcardSelectionKey) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeWildcardSelectionKey WildcardSelectionKey
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeWildcardSelectionKey
	}{
		"WILDCARD",
		(MarshalTypeWildcardSelectionKey)(m),
	}

	return json.Marshal(&s)
}
