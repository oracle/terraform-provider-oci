// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Batch API
//
// Use the Batch Control Plane API to encapsulate and manage all aspects of computationally intensive jobs.
//

package batch

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// JobPriorityConfiguration Job priority configuration to instruct the service on how to use priority tags.
type JobPriorityConfiguration struct {

	// Name of the corresponding tag namespace.
	TagNamespace *string `mandatory:"true" json:"tagNamespace"`

	// Name of the tag key.
	TagKey *string `mandatory:"true" json:"tagKey"`

	// Weight associated with the tag key. Percentage point is the unit of measurement.
	Weight *int `mandatory:"true" json:"weight"`

	// Mapping of tag value to its priority.
	Values map[string]int `mandatory:"true" json:"values"`
}

func (m JobPriorityConfiguration) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m JobPriorityConfiguration) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
