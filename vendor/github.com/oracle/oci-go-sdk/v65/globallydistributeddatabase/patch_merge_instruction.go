// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Globally Distributed Database
//
// Use the Globally Distributed Database service APIs to create and manage distributed databases.
//

package globallydistributeddatabase

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// PatchMergeInstruction An operation that recursively updates items of the selection, or adding the value if the selection is empty.
// If the value is not an object, it is used directly, otherwise each key-value member is used
// to create or update a member of the same name in the target and the same process is applied recursively for each object-typed value
// (similar to RFC 7396 (https://tools.ietf.org/html/rfc7396#section-2) JSON Merge Patch, except that null values are copied
// rather than transformed into deletions).
// NOT_FOUND exceptions are handled by creating the implied containing structure.
// To avoid referential errors if an item's descendant is also in the selection, items of the selection are processed in order of decreasing depth.
type PatchMergeInstruction struct {

	// The set of values to which the operation applies as a JMESPath expression (https://jmespath.org/specification.html) for evaluation against the context resource.
	// An operation fails if the selection yields an exception, except as otherwise specified.
	// Note that comparisons involving non-primitive values (objects or arrays) are not supported and will always evaluate to false.
	Selection *string `mandatory:"true" json:"selection"`

	// A value to be merged into the target.
	Value *interface{} `mandatory:"false" json:"value"`
}

// GetSelection returns Selection
func (m PatchMergeInstruction) GetSelection() *string {
	return m.Selection
}

func (m PatchMergeInstruction) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PatchMergeInstruction) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m PatchMergeInstruction) MarshalJSON() (buff []byte, e error) {
	type MarshalTypePatchMergeInstruction PatchMergeInstruction
	s := struct {
		DiscriminatorParam string `json:"operation"`
		MarshalTypePatchMergeInstruction
	}{
		"MERGE",
		(MarshalTypePatchMergeInstruction)(m),
	}

	return json.Marshal(&s)
}
