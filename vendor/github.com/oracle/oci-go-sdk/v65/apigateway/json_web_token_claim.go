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
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// JsonWebTokenClaim An individual JWT claim.
type JsonWebTokenClaim struct {

	// Name of the claim.
	Key *string `mandatory:"true" json:"key"`

	// The list of acceptable values for a given claim.
	// If this value is "null" or empty and "isRequired" set to "true", then
	// the presence of this claim in the JWT is validated.
	Values []string `mandatory:"false" json:"values"`

	// Whether the claim is required to be present in the JWT or not. If set
	// to "false", the claim values will be matched only if the claim is
	// present in the JWT.
	IsRequired *bool `mandatory:"false" json:"isRequired"`
}

func (m JsonWebTokenClaim) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m JsonWebTokenClaim) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
