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

// AnyOfSelectionKey When dynamically routing and dynamically authenticating requests, the route or authentication server associated with a set of selection keys is used if the context variable in an incoming request exactly matches one of the keys in the set.
type AnyOfSelectionKey struct {

	// Name assigned to the branch.
	Name *string `mandatory:"true" json:"name"`

	// Specifies whether to use the route or authentication server associated with this selection key as the default. The default is used if the value of a context variable in an incoming request does not match any of the other selection key values when dynamically routing and dynamically authenticating requests.
	IsDefault *bool `mandatory:"false" json:"isDefault"`

	// The set of selection keys to match with the context variable in an incoming request. If the context variable exactly matches one of the keys in the set, the request is sent to the route or authentication server associated with the set.
	Values []string `mandatory:"false" json:"values"`
}

// GetIsDefault returns IsDefault
func (m AnyOfSelectionKey) GetIsDefault() *bool {
	return m.IsDefault
}

// GetName returns Name
func (m AnyOfSelectionKey) GetName() *string {
	return m.Name
}

func (m AnyOfSelectionKey) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AnyOfSelectionKey) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m AnyOfSelectionKey) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeAnyOfSelectionKey AnyOfSelectionKey
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeAnyOfSelectionKey
	}{
		"ANY_OF",
		(MarshalTypeAnyOfSelectionKey)(m),
	}

	return json.Marshal(&s)
}
