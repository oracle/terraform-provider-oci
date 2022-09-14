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
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// WildcardSelectionKey Information around the Wildcard expression matching the value for selector of a dynamic authentication/ routing branch.
type WildcardSelectionKey struct {

	// Name assigned to the branch.
	Name *string `mandatory:"true" json:"name"`

	// String describing the expression with wildcards.
	Expression *string `mandatory:"true" json:"expression"`

	// Information regarding whether this is the default branch.
	IsDefault *bool `mandatory:"false" json:"isDefault"`
}

//GetIsDefault returns IsDefault
func (m WildcardSelectionKey) GetIsDefault() *bool {
	return m.IsDefault
}

//GetName returns Name
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
