// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Service Mesh API
//
// Use the Service Mesh API to manage mesh, virtual service, access policy and other mesh related items.
//

package servicemesh

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// IngressGatewayMutualTransportLayerSecurityDetails Mutual TLS settings used when sending requests to virtual services within the mesh.
type IngressGatewayMutualTransportLayerSecurityDetails struct {

	// The number of days the mTLS certificate is valid.  This value should be less than the Maximum Validity Duration
	// for Certificates (Days) setting on the Certificate Authority associated with this Mesh.  The certificate will
	// be automatically renewed after 2/3 of the validity period, so a certificate with a maximum validity of 45 days
	// will be renewed every 30 days.
	MaximumValidity *int `mandatory:"false" json:"maximumValidity"`
}

func (m IngressGatewayMutualTransportLayerSecurityDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m IngressGatewayMutualTransportLayerSecurityDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
