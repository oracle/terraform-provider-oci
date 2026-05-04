// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Tools Runtime API
//
// Use the Database Tools Runtime API to connect to databases through Database Tools Connections.
//

package databasetoolsruntime

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// PropertySetApex Contains the details of an APEX property set
type PropertySetApex struct {

	// Indicates whether the property set is mutable or not
	IsMutable *bool `mandatory:"true" json:"isMutable"`

	// The version of APEX
	Version *string `mandatory:"false" json:"version"`

	// The APEX engine schema name
	UserKey *string `mandatory:"false" json:"userKey"`
}

// GetIsMutable returns IsMutable
func (m PropertySetApex) GetIsMutable() *bool {
	return m.IsMutable
}

func (m PropertySetApex) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PropertySetApex) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m PropertySetApex) MarshalJSON() (buff []byte, e error) {
	type MarshalTypePropertySetApex PropertySetApex
	s := struct {
		DiscriminatorParam string `json:"key"`
		MarshalTypePropertySetApex
	}{
		"APEX",
		(MarshalTypePropertySetApex)(m),
	}

	return json.Marshal(&s)
}
