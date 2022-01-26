// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Integration API
//
// Oracle Integration API.
//

package integration

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v56/common"
)

// PublicEndpointDetails Public endpoint configuration details.
type PublicEndpointDetails struct {

	// Source IP addresses or IP address ranges ingress rules.
	AllowlistedHttpIps []string `mandatory:"false" json:"allowlistedHttpIps"`

	// Virtual Cloud Networks allowed to access this network endpoint.
	AllowlistedHttpVcns []VirtualCloudNetwork `mandatory:"false" json:"allowlistedHttpVcns"`

	// The Integration service's VCN is allow-listed to allow integrations to call back into other integrations
	IsIntegrationVcnAllowlisted *bool `mandatory:"false" json:"isIntegrationVcnAllowlisted"`
}

func (m PublicEndpointDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m PublicEndpointDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypePublicEndpointDetails PublicEndpointDetails
	s := struct {
		DiscriminatorParam string `json:"networkEndpointType"`
		MarshalTypePublicEndpointDetails
	}{
		"PUBLIC",
		(MarshalTypePublicEndpointDetails)(m),
	}

	return json.Marshal(&s)
}
