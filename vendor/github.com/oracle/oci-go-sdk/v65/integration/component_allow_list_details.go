// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Integration API
//
// Oracle Integration API.
//

package integration

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ComponentAllowListDetails A collection of AllowListedIps and AllowListedVcns for a specific OIC component: global, RT, or DT
type ComponentAllowListDetails struct {

	// Source IP addresses or IP address ranges ingress rules. (ex: "168.122.59.5", "10.20.30.0/26")
	// An invalid IP or CIDR block will result in a 400 response.
	AllowlistedHttpIps []string `mandatory:"false" json:"allowlistedHttpIps"`

	// Virtual Cloud Networks allowed to access this network endpoint.
	AllowlistedHttpVcns []VirtualCloudNetwork `mandatory:"false" json:"allowlistedHttpVcns"`
}

func (m ComponentAllowListDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ComponentAllowListDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
