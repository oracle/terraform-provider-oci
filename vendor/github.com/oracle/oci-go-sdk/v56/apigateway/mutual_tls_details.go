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
	"github.com/oracle/oci-go-sdk/v56/common"
)

// MutualTlsDetails Properties used to configure client mTLS verification when API Consumer makes connection to the gateway.
type MutualTlsDetails struct {

	// Determines whether to enable client verification when API Consumer makes connection to the gateway.
	IsVerifiedCertificateRequired *bool `mandatory:"false" json:"isVerifiedCertificateRequired"`

	// Allowed list of CN or SAN which will be used for verification of certificate.
	AllowedSans []string `mandatory:"false" json:"allowedSans"`
}

func (m MutualTlsDetails) String() string {
	return common.PointerString(m)
}
