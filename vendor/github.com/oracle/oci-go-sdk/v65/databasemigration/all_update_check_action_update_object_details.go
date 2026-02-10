// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Migration API
//
// Use the Oracle Cloud Infrastructure Database Migration APIs to perform database migration operations.
//

package databasemigration

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AllUpdateCheckActionUpdateObjectDetails Exclude all check objects.
type AllUpdateCheckActionUpdateObjectDetails struct {

	// Flag showing the action on the object.
	IsExclude *bool `mandatory:"false" json:"isExclude"`
}

func (m AllUpdateCheckActionUpdateObjectDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AllUpdateCheckActionUpdateObjectDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m AllUpdateCheckActionUpdateObjectDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeAllUpdateCheckActionUpdateObjectDetails AllUpdateCheckActionUpdateObjectDetails
	s := struct {
		DiscriminatorParam string `json:"kind"`
		MarshalTypeAllUpdateCheckActionUpdateObjectDetails
	}{
		"ALL_OBJECTS",
		(MarshalTypeAllUpdateCheckActionUpdateObjectDetails)(m),
	}

	return json.Marshal(&s)
}
