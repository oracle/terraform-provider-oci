// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Stack Monitoring API
//
// Stack Monitoring API.
//

package stackmonitoring

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateAndPropagateTagsDetails The information about monitored resource tags. Request will fail if at least one of
// freeformTags or definedTags are not specified.
// Provided tags will be added or updated in the existing list of tags for the affected resources.
// Resources to be updated are identified based on association types specified.
// If association types are not specified, then tags will be updated only for the current resource.
type UpdateAndPropagateTagsDetails struct {

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Association types that will be traversed recursively starting from the current resource,
	// to identify resources for which the tags will be updated.
	// If no association type is specified, only current resource will be updated.
	// Default is empty list, which means no related resources will be updated.
	AssociationTypes []string `mandatory:"false" json:"associationTypes"`
}

func (m UpdateAndPropagateTagsDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateAndPropagateTagsDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
