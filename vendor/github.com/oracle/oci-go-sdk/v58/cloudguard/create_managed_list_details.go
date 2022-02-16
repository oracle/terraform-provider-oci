// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Guard APIs
//
// A description of the Cloud Guard APIs
//

package cloudguard

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// CreateManagedListDetails Create ManagedList
type CreateManagedListDetails struct {

	// ManagedList display name
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Compartment Identifier
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// OCID of the Source ManagedList
	SourceManagedListId *string `mandatory:"false" json:"sourceManagedListId"`

	// ManagedList description
	Description *string `mandatory:"false" json:"description"`

	// type of the list
	ListType ManagedListTypeEnum `mandatory:"false" json:"listType,omitempty"`

	// List of ManagedListItem
	ListItems []string `mandatory:"false" json:"listItems"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateManagedListDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateManagedListDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingManagedListTypeEnum(string(m.ListType)); !ok && m.ListType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ListType: %s. Supported values are: %s.", m.ListType, strings.Join(GetManagedListTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
