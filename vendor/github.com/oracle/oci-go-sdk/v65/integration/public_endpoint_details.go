// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Integration API
//
// Oracle Integration API.
//

package integration

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// PublicEndpointDetails Public endpoint configuration details.
type PublicEndpointDetails struct {

	// Source IP addresses or IP address ranges ingress rules. (ex: "168.122.59.5", "10.20.30.0/26")
	// An invalid IP or CIDR block will result in a 400 response.
	AllowlistedHttpIps []string `mandatory:"false" json:"allowlistedHttpIps"`

	// Virtual Cloud Networks allowed to access this network endpoint.
	AllowlistedHttpVcns []VirtualCloudNetwork `mandatory:"false" json:"allowlistedHttpVcns"`

	// The Integration service's VCN is allow-listed to allow integrations to call back into other integrations
	IsIntegrationVcnAllowlisted *bool `mandatory:"false" json:"isIntegrationVcnAllowlisted"`
}

func (m PublicEndpointDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PublicEndpointDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
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
