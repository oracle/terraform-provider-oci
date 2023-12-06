// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// PGSQL Control Plane API
//
// A description of the PGSQL Control Plane API
//

package psql

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// PatchReplaceInstruction An operation that "puts" a value, replacing every item of the selection with it, or creating it if the selection is empty.
// NOT_FOUND exceptions are handled by creating the implied containing structure (but note that this may put the target in an invalid state,
// which can be prevented by use of precondition operations).
// To avoid referential errors if an item's descendant is also in the selection, items of the selection are processed in order of decreasing depth.
type PatchReplaceInstruction struct {

	// The set of values to which the operation applies as a JMESPath expression (https://jmespath.org/specification.html) for evaluation against the context resource.
	// An operation fails if the selection yields an exception, except as otherwise specified.
	// Note that comparisons involving non-primitive values (objects or arrays) are not supported and will always evaluate to false.
	Selection *string `mandatory:"true" json:"selection"`

	// A value to be added into the target.
	Value *interface{} `mandatory:"true" json:"value"`
}

// GetSelection returns Selection
func (m PatchReplaceInstruction) GetSelection() *string {
	return m.Selection
}

func (m PatchReplaceInstruction) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PatchReplaceInstruction) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m PatchReplaceInstruction) MarshalJSON() (buff []byte, e error) {
	type MarshalTypePatchReplaceInstruction PatchReplaceInstruction
	s := struct {
		DiscriminatorParam string `json:"operation"`
		MarshalTypePatchReplaceInstruction
	}{
		"REPLACE",
		(MarshalTypePatchReplaceInstruction)(m),
	}

	return json.Marshal(&s)
}
