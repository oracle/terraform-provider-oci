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

// CertificatesCaBundle CA bundle from Certificates Service that should be used on the gateway for TLS validation
type CertificatesCaBundle struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the resource.
	CaBundleId *string `mandatory:"false" json:"caBundleId"`
}

func (m CertificatesCaBundle) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CertificatesCaBundle) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CertificatesCaBundle) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCertificatesCaBundle CertificatesCaBundle
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeCertificatesCaBundle
	}{
		"CA_BUNDLE",
		(MarshalTypeCertificatesCaBundle)(m),
	}

	return json.Marshal(&s)
}
