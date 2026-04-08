// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OS Management Hub API
//
// Use the OS Management Hub API to manage and monitor updates and patches for instances in OCI, your private data center, or 3rd-party clouds.
// For more information, see Overview of OS Management Hub (https://docs.oracle.com/iaas/osmh/doc/overview.htm).
//

package osmanagementhub

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DefinedTag Defined tag assigned to a resource.
type DefinedTag struct {

	// The key of the tag.
	Key *string `mandatory:"false" json:"key"`

	// The value associated with the tag key.
	Value *string `mandatory:"false" json:"value"`

	// The namespace of the tag.
	Namespace *string `mandatory:"false" json:"namespace"`
}

// GetKey returns Key
func (m DefinedTag) GetKey() *string {
	return m.Key
}

// GetValue returns Value
func (m DefinedTag) GetValue() *string {
	return m.Value
}

func (m DefinedTag) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DefinedTag) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m DefinedTag) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDefinedTag DefinedTag
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeDefinedTag
	}{
		"DEFINED",
		(MarshalTypeDefinedTag)(m),
	}

	return json.Marshal(&s)
}
