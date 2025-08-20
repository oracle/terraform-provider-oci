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

// ResourcesPatchComplianceSummary Summary of resources and their patch compliance.
type ResourcesPatchComplianceSummary struct {

	// Total number of resources.
	TotalResourcesCount *int `mandatory:"false" json:"totalResourcesCount"`

	// Total number of resources that are up to date.
	UpToDateResourcesCount *int `mandatory:"false" json:"upToDateResourcesCount"`

	// Total number of  non-compliant resources.
	NonCompliantResourcesCount *int `mandatory:"false" json:"nonCompliantResourcesCount"`

	// Total number of resources that are not subscribed.
	NotSubscribedResourcesCount *int `mandatory:"false" json:"notSubscribedResourcesCount"`

	// Total number of resources that are not registered to DBLM.
	NotDblmRegisteredResourcesCount *int `mandatory:"false" json:"notDblmRegisteredResourcesCount"`
}

func (m ResourcesPatchComplianceSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ResourcesPatchComplianceSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
