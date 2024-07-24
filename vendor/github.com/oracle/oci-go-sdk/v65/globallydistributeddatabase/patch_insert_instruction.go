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

// PatchInsertInstruction An operation that inserts a value into an array, shifting array items as necessary and handling NOT_FOUND exceptions by creating the implied containing structure.
type PatchInsertInstruction struct {

	// The set of values to which the operation applies as a JMESPath expression (https://jmespath.org/specification.html) for evaluation against the context resource.
	// An operation fails if the selection yields an exception, except as otherwise specified.
	// Note that comparisons involving non-primitive values (objects or arrays) are not supported and will always evaluate to false.
	Selection *string `mandatory:"true" json:"selection"`

	// A value to be inserted into the target.
	Value *interface{} `mandatory:"true" json:"value"`
}

// GetSelection returns Selection
func (m PatchInsertInstruction) GetSelection() *string {
	return m.Selection
}

func (m PatchInsertInstruction) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PatchInsertInstruction) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m PatchInsertInstruction) MarshalJSON() (buff []byte, e error) {
	type MarshalTypePatchInsertInstruction PatchInsertInstruction
	s := struct {
		DiscriminatorParam string `json:"operation"`
		MarshalTypePatchInsertInstruction
	}{
		"INSERT",
		(MarshalTypePatchInsertInstruction)(m),
	}

	return json.Marshal(&s)
}
