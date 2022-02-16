// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// VisionService API
//
// A description of the VisionService API.
//

package aivision

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// NormalizedVertex An (x, y) coordinate in the image with dimensions normalized from zero to one.
// The origin is at top left, with the positive x-axis pointing right and the positive y-axis pointing down.
// The bottom right corner is at (1, 1).
type NormalizedVertex struct {

	// X axis normalized coordinate.
	X *float64 `mandatory:"true" json:"x"`

	// Y axis normalized coordinate.
	Y *float64 `mandatory:"true" json:"y"`
}

func (m NormalizedVertex) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m NormalizedVertex) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
