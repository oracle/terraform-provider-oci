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

// ValueListEntity Metadata for a value list entity.
type ValueListEntity struct {

	// Unique immutable identifier that was assigned when the resource was created.
	Id *string `mandatory:"true" json:"id"`

	// The entity name. This must be unique within the parent resource.
	Name *string `mandatory:"true" json:"name"`

	// List of values for a value list entity.
	Values []StaticEntityValue `mandatory:"true" json:"values"`
}

// GetId returns Id
func (m ValueListEntity) GetId() *string {
	return m.Id
}

// GetName returns Name
func (m ValueListEntity) GetName() *string {
	return m.Name
}

func (m ValueListEntity) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ValueListEntity) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m ValueListEntity) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeValueListEntity ValueListEntity
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeValueListEntity
	}{
		"ENUM_VALUES",
		(MarshalTypeValueListEntity)(m),
	}

	return json.Marshal(&s)
}
