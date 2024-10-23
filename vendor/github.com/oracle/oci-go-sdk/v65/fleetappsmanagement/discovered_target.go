// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fleet Application Management Service API
//
// Fleet Application Management provides a centralized platform to help you automate resource management tasks, validate patch compliance, and enhance operational efficiency across an enterprise.
//

package fleetappsmanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DiscoveredTarget A target that is discovered by the Software discovery process.
type DiscoveredTarget struct {

	// ID of the Target. Can be the target name if a separate ID is not available.
	TargetId *string `mandatory:"true" json:"targetId"`

	// Target Name.
	TargetName *string `mandatory:"true" json:"targetName"`

	// Product that the target belongs to.
	Product *string `mandatory:"true" json:"product"`

	// Unique key that identifies the resource that the target belongs to.
	ResourceId *string `mandatory:"true" json:"resourceId"`

	// Current version of the target.
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
