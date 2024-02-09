// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Digital Assistant Service Instance API
//
// API to create and maintain Oracle Digital Assistant service instances.
//

package oda

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// EntityActionArgument Argument of an entity action
type EntityActionArgument struct {

	// Name of an entity action argument
	Name *string `mandatory:"true" json:"name"`

	// Type of an entity action argument
	Type EntityAttributeTypeEnum `mandatory:"true" json:"type"`

	// Metatype of an entity action argument
	MetaType *string `mandatory:"false" json:"metaType"`

	NaturalLanguageMapping *EntityActionArgumentNaturalLanguageMapping `mandatory:"false" json:"naturalLanguageMapping"`

	// Is the entity action argument multi-value
	IsMultiValue *bool `mandatory:"false" json:"isMultiValue"`

	// Name of referenced entity.
	EntityName *string `mandatory:"false" json:"entityName"`
}

func (m EntityActionArgument) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m EntityActionArgument) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingEntityAttributeTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetEntityAttributeTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
