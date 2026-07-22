// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Analytics API
//
// Analytics API.
//

package analytics

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// InstanceResourceGroupSummary Resource Groups for Instance
type InstanceResourceGroupSummary struct {

	// Unique identifier and name of resource group.  Must be unique within the instance
	Id *string `mandatory:"true" json:"id"`

	// Meaningful name of resource group for end user
	ResourceName *string `mandatory:"true" json:"resourceName"`

	// The capacity (in OCPU's) to be allocated for this resource.
	Capacity *int `mandatory:"true" json:"capacity"`

	// Meaningful name of resource group for end user
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Optional description of the resource group
	Description *string `mandatory:"false" json:"description"`
}

func (m InstanceResourceGroupSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m InstanceResourceGroupSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
