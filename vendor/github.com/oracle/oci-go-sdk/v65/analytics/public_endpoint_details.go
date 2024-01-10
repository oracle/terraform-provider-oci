// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Analytics API
//
// Analytics API.
//

package analytics

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// PublicEndpointDetails Public endpoint configuration details.
type PublicEndpointDetails struct {

	// Source IP addresses or IP address ranges in ingress rules.
	WhitelistedIps []string `mandatory:"false" json:"whitelistedIps"`

	// Virtual Cloud Networks allowed to access this network endpoint.
	WhitelistedVcns []VirtualCloudNetwork `mandatory:"false" json:"whitelistedVcns"`

	// Oracle Cloud Services that are allowed to access this Analytics instance.
	WhitelistedServices []AccessControlServiceTypeEnum `mandatory:"false" json:"whitelistedServices,omitempty"`
}

func (m PublicEndpointDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PublicEndpointDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	for _, val := range m.WhitelistedServices {
		if _, ok := GetMappingAccessControlServiceTypeEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for WhitelistedServices: %s. Supported values are: %s.", val, strings.Join(GetAccessControlServiceTypeEnumStringValues(), ",")))
		}
	}

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
