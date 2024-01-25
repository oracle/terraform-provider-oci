// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Digital Assistant Service Instance API
//
// API to create and maintain Oracle Digital Assistant service instances.
//

package oda

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CompositeEntity Metadata for a composite entity.
type CompositeEntity struct {

	// Unique immutable identifier that was assigned when the resource was created.
	Id *string `mandatory:"true" json:"id"`

	// The entity name. This must be unique within the parent resource.
	Name *string `mandatory:"true" json:"name"`

	// List of entity attributes.
	Attributes []EntityAttribute `mandatory:"true" json:"attributes"`

	// List of entity actions.
	Actions []EntityAction `mandatory:"false" json:"actions"`

	NaturalLanguageMapping *EntityNaturalLanguageMapping `mandatory:"false" json:"naturalLanguageMapping"`
}

// GetId returns Id
func (m CompositeEntity) GetId() *string {
	return m.Id
}

// GetName returns Name
func (m CompositeEntity) GetName() *string {
	return m.Name
}

func (m CompositeEntity) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CompositeEntity) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CompositeEntity) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCompositeEntity CompositeEntity
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeCompositeEntity
	}{
		"COMPOSITE",
		(MarshalTypeCompositeEntity)(m),
	}

	return json.Marshal(&s)
}
