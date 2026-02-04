// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Flow API
//
// Use the Data Flow APIs to run any Apache Spark application at any scale without deploying or managing any infrastructure.
//

package dataflow

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// OAuthConfiguration OAuth configuration settings for a Deployment.
type OAuthConfiguration struct {

	// The issuer (iss) claim for OAuth token validation.
	// e.g. https://identity.oraclecloud.com
	Issuer *string `mandatory:"false" json:"issuer"`

	// The audience (aud) claims for OAuth token validation.
	// Defaults to the agentFlowKey if not specified.
	Audience []string `mandatory:"false" json:"audience"`

	// The JWKS (JSON Web Key Set) URI for OAuth token validation.
	// e.g. https://<tenant-base-url>/admin/v1/SigningCert/jwk
	JwksUri *string `mandatory:"false" json:"jwksUri"`

	// The OAuth client identifier for authentication.
	ClientId *string `mandatory:"false" json:"clientId"`
}

func (m OAuthConfiguration) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OAuthConfiguration) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
