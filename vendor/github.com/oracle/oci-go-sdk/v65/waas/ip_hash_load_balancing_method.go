// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Web Application Acceleration and Security Services API
//
// OCI Web Application Acceleration and Security Services
//

package waas

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// IpHashLoadBalancingMethod An object that represents the `ip-hash` load balancing method.
type IpHashLoadBalancingMethod struct {
}

func (m IpHashLoadBalancingMethod) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m IpHashLoadBalancingMethod) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m IpHashLoadBalancingMethod) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeIpHashLoadBalancingMethod IpHashLoadBalancingMethod
	s := struct {
		DiscriminatorParam string `json:"method"`
		MarshalTypeIpHashLoadBalancingMethod
	}{
		"IP_HASH",
		(MarshalTypeIpHashLoadBalancingMethod)(m),
	}

	return json.Marshal(&s)
}
