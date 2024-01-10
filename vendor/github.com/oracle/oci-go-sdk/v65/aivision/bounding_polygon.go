// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Vision API
//
// Using Vision, you can upload images to detect and classify objects in them. If you have lots of images, you can process them in batch using asynchronous API endpoints. Vision's features are thematically split between Document AI for document-centric images, and Image Analysis for object and scene-based images. Pretrained models and custom models are supported.
//

package aivision

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// BoundingPolygon The object-bounding polygon box.
type BoundingPolygon struct {

	// An array of normalized points defining the polygon's perimeter, with an implicit segment between subsequent points and between the first and last point.
	// Rectangles are defined with four points. For example, `[{"x": 0, "y": 0}, {"x": 1, "y": 0}, {"x": 1, "y": 0.5}, {"x": 0, "y": 0.5}]` represents the top half of an image.
	NormalizedVertices []NormalizedVertex `mandatory:"true" json:"normalizedVertices"`
}

func (m BoundingPolygon) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m BoundingPolygon) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
