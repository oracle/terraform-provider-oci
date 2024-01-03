// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Web Application Firewall (WAF) API
//
// API for the Web Application Firewall service.
// Use this API to manage regional Web App Firewalls and corresponding policies for protecting HTTP services.
//

package waf

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CollaborativeCapabilityWeight Defines how much a contributing capability contributes towards the action threshold of a collaborative protection capability.
type CollaborativeCapabilityWeight struct {

	// Unique key of contributing protection capability.
	Key *string `mandatory:"true" json:"key"`

	// The display name of contributing protection capability.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The weight of contributing protection capability.
	Weight *int `mandatory:"true" json:"weight"`
}

func (m CollaborativeCapabilityWeight) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CollaborativeCapabilityWeight) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
