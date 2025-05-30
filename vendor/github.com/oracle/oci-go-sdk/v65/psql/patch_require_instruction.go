// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// PGSQL Control Plane API
//
// Use the OCI Database with PostgreSQL API to manage resources such as database systems, database nodes, backups, and configurations.
// For information, see the user guide documentation for the service (https://docs.oracle.com/iaas/Content/postgresql/home.htm).
//

package psql

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// PatchRequireInstruction A precondition operation that requires a selection to be non-empty, and optionally to include an item with a specified value
// (useful for asserting that a value exists before attempting to update it, avoiding accidental creation).
// It fails if the selection is empty, or if value is provided and no item of the selection matches it.
type PatchRequireInstruction struct {

	// The set of values to which the operation applies as a JMESPath expression (https://jmespath.org/specification.html) for evaluation against the context resource.
	// An operation fails if the selection yields an exception, except as otherwise specified.
	// Note that comparisons involving non-primitive values (objects or arrays) are not supported and will always evaluate to false.
	Selection *string `mandatory:"true" json:"selection"`

	// A value to be compared against each item of the selection.
	// If this value is an object, then it matches any item that would be unaffected by applying this value as a merge operation.
	// Otherwise, it matches any item to which it is equal according to the rules of JSON Schema (https://tools.ietf.org/html/draft-handrews-json-schema-00#section-4.2.3).
	Value *interface{} `mandatory:"false" json:"value"`
}

// GetSelection returns Selection
func (m PatchRequireInstruction) GetSelection() *string {
	return m.Selection
}

func (m PatchRequireInstruction) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PatchRequireInstruction) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m PatchRequireInstruction) MarshalJSON() (buff []byte, e error) {
	type MarshalTypePatchRequireInstruction PatchRequireInstruction
	s := struct {
		DiscriminatorParam string `json:"operation"`
		MarshalTypePatchRequireInstruction
	}{
		"REQUIRE",
		(MarshalTypePatchRequireInstruction)(m),
	}

	return json.Marshal(&s)
}
