// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Dblm API
//
// A description of the Dblm API
//

package dblm

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ResourceSummary Summary of dblm registered and unregistered resources.
type ResourceSummary struct {

	// The total number of resources.
	TotalResourcesCount *int `mandatory:"true" json:"totalResourcesCount"`

	// The total number of registered resources.
	RegisteredResourcesCount *int `mandatory:"true" json:"registeredResourcesCount"`

	// The total number of resources that are not registered.
	NotRegisteredResourcesCount *int `mandatory:"true" json:"notRegisteredResourcesCount"`

	// Total number of resources that have 1 or more vulnerabilities.
	VulnerableResourcesCount *int `mandatory:"true" json:"vulnerableResourcesCount"`

	// Total number of resources that have 0 vulnerabilities.
	CleanResourcesCount *int `mandatory:"true" json:"cleanResourcesCount"`

	// Total number of resources that contain an error.
	ErrorResourcesCount *int `mandatory:"true" json:"errorResourcesCount"`
}

func (m ResourceSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ResourceSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
