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
	"github.com/oracle/oci-go-sdk/common"
)

// RenameQueryParameterPolicyItem The value will be a copy of the original value of the source parameter and will not be affected by any other
// transformation policies applied to that parameter.
type RenameQueryParameterPolicyItem struct {

	// The original case-sensitive name of the query parameter.  This name must be unique across transformation
	// policies.
	From *string `mandatory:"true" json:"from"`

	// The new name of the query parameter.  This name must be unique across transformation policies.
	To *string `mandatory:"true" json:"to"`
}

func (m RenameQueryParameterPolicyItem) String() string {
	return common.PointerString(m)
}
