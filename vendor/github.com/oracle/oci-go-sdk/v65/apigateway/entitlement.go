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
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Entitlement A usage plan entitlement, comprising of rate limits, quotas and the deployments they are applied to.
type Entitlement struct {

	// An entitlement name, unique within a usage plan.
	Name *string `mandatory:"true" json:"name"`

	// A user-friendly description. To provide some insight about the resource.
	// Avoid entering confidential information.
	Description *string `mandatory:"false" json:"description"`

	RateLimit *RateLimit `mandatory:"false" json:"rateLimit"`

	Quota *Quota `mandatory:"false" json:"quota"`

	// A collection of targeted deployments that the entitlement will be applied to.
	Targets []EntitlementTarget `mandatory:"false" json:"targets"`
}

func (m Entitlement) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Entitlement) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
