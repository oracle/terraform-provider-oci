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

// EntityAttribute Attribute of an entity
type EntityAttribute struct {

	// The name of an entity attribute
	Name *string `mandatory:"true" json:"name"`

	// The type of an entity attribute
	Type EntityAttributeTypeEnum `mandatory:"true" json:"type"`

	NaturalLanguageMapping *EntityAttributeNaturalLanguageMapping `mandatory:"false" json:"naturalLanguageMapping"`

	// Is the entity attribute multi-value
	IsMultiValue *bool `mandatory:"false" json:"isMultiValue"`

	// Is the entity attribute a fuzzy match
	IsFuzzyMatch *bool `mandatory:"false" json:"isFuzzyMatch"`

	// Are comparisons inverted in the entity attribute
	IsInvertComparisons *bool `mandatory:"false" json:"isInvertComparisons"`

	// Temporal preference of an attribute
	TemporalPreference TemporalPreferenceEnum `mandatory:"false" json:"temporalPreference,omitempty"`

	// Name of referenced entity.
	EntityName *string `mandatory:"false" json:"entityName"`
}

func (m EntityAttribute) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m EntityAttribute) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingEntityAttributeTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetEntityAttributeTypeEnumStringValues(), ",")))
	}

	if _, ok := GetMappingTemporalPreferenceEnum(string(m.TemporalPreference)); !ok && m.TemporalPreference != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TemporalPreference: %s. Supported values are: %s.", m.TemporalPreference, strings.Join(GetTemporalPreferenceEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
