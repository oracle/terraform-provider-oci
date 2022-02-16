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

// FieldName Name of a form field.
type FieldName struct {

	// Name of the field.
	Name *string `mandatory:"true" json:"name"`

	// Confidence score between 0 to 1.
	Confidence *float32 `mandatory:"false" json:"confidence"`

	BoundingPolygon *BoundingPolygon `mandatory:"false" json:"boundingPolygon"`

	// Indexes of the words in the field name.
	WordIndexes []int `mandatory:"false" json:"wordIndexes"`
}

func (m FieldName) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m FieldName) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
