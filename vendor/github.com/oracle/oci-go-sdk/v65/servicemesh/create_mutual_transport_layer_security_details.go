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

// CreateMutualTransportLayerSecurityDetails The mTLS authentication mode to use when receiving requests from other virtual services or ingress gateways within the mesh.
type CreateMutualTransportLayerSecurityDetails struct {

	// DISABLED: Connection is not tunneled.
	// PERMISSIVE: Connection can be either plaintext or an mTLS tunnel.
	// STRICT: Connection is an mTLS tunnel.  Clients without a valid certificate will be rejected.
	Mode MutualTransportLayerSecurityModeEnum `mandatory:"true" json:"mode"`

	// The number of days the mTLS certificate is valid.  This value should be less than the Maximum Validity Duration
	// for Certificates (Days) setting on the Certificate Authority associated with this Mesh.  The certificate will
	// be automatically renewed after 2/3 of the validity period, so a certificate with a maximum validity of 45 days
	// will be renewed every 30 days.
	MaximumValidity *int `mandatory:"false" json:"maximumValidity"`
}

func (m CreateMutualTransportLayerSecurityDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateMutualTransportLayerSecurityDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingMutualTransportLayerSecurityModeEnum(string(m.Mode)); !ok && m.Mode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Mode: %s. Supported values are: %s.", m.Mode, strings.Join(GetMutualTransportLayerSecurityModeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
