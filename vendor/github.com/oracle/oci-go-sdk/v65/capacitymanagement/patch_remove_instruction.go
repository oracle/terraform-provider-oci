// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OciControlCenterCp API
//
// A description of the OciControlCenterCp API
//

package capacitymanagement

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// PatchRemoveInstruction An operation that deletes items, ignoring NOT_FOUND exceptions.
// To avoid referential errors if an item's descendant is also in the selection, items of the selection are processed in order of decreasing depth.
type PatchRemoveInstruction struct {

	// The set of values to which the operation applies as a JMESPath expression (https://jmespath.org/specification.html) for evaluation against the context resource.
	// An operation fails if the selection yields an exception, except as otherwise specified.
	// Note that comparisons involving non-primitive values (objects or arrays) are not supported and will always evaluate to false.
	Selection *string `mandatory:"true" json:"selection"`
}

// GetSelection returns Selection
func (m PatchRemoveInstruction) GetSelection() *string {
	return m.Selection
}

func (m PatchRemoveInstruction) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PatchRemoveInstruction) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m PatchRemoveInstruction) MarshalJSON() (buff []byte, e error) {
	type MarshalTypePatchRemoveInstruction PatchRemoveInstruction
	s := struct {
		DiscriminatorParam string `json:"operation"`
		MarshalTypePatchRemoveInstruction
	}{
		"REMOVE",
		(MarshalTypePatchRemoveInstruction)(m),
	}

	return json.Marshal(&s)
}
