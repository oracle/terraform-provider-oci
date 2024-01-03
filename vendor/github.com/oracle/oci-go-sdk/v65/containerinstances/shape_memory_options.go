// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Container Instance API
//
// A description of the Container Instance API
//

package containerinstances

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ShapeMemoryOptions For a flexible shape, the amount of memory available for container instances that use this shape.
type ShapeMemoryOptions struct {

	// The minimum amount of memory (GB).
	MinInGBs *float32 `mandatory:"true" json:"minInGBs"`

	// The maximum amount of memory (GB).
	MaxInGBs *float32 `mandatory:"true" json:"maxInGBs"`

	// The default amount of memory per OCPU available for this shape (GB).
	DefaultPerOcpuInGBs *float32 `mandatory:"true" json:"defaultPerOcpuInGBs"`

	// The minimum amount of memory per OCPU available for this shape (GB).
	MinPerOcpuInGBs *float32 `mandatory:"true" json:"minPerOcpuInGBs"`

	// The maximum amount of memory per OCPU available for this shape (GB).
	MaxPerOcpuInGBs *float32 `mandatory:"true" json:"maxPerOcpuInGBs"`
}

func (m ShapeMemoryOptions) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ShapeMemoryOptions) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
