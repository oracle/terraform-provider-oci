// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fleet Application Management Service API
//
// Fleet Application Management Service API. Use this API to for all FAMS related activities.
// To manage fleets,view complaince report for the Fleet,scedule patches and other lifecycle activities
//

package fleetappsmanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DiscoveredTarget A target that is discovered by the Software discovery process.
type DiscoveredTarget struct {

	// OCID of the Target.
	TargetId *string `mandatory:"true" json:"targetId"`

	// Target Name.
	TargetName *string `mandatory:"true" json:"targetName"`

	// Product that the target belongs to.
	Product *string `mandatory:"true" json:"product"`

	// Unique key that identify the resource that target belongs to.
	ResourceId *string `mandatory:"true" json:"resourceId"`

	// Current version of Target
	Version *string `mandatory:"false" json:"version"`
}

func (m DiscoveredTarget) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DiscoveredTarget) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
