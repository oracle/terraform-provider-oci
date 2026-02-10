// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Migration API
//
// Use the Oracle Cloud Infrastructure Database Migration APIs to perform database migration operations.
//

package databasemigration

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ObjectMetadata Metadata of object.
type ObjectMetadata struct {

	// The field that stores the owner of the object.
	SchemaOwnerColumn *string `mandatory:"true" json:"schemaOwnerColumn"`

	// The field that stores the name of the object.
	ObjectNameColumn *string `mandatory:"true" json:"objectNameColumn"`

	// The field that stores the fixed type of the object.
	ObjectTypeFixed *string `mandatory:"false" json:"objectTypeFixed"`

	// The field that stores the type of the object.
	ObjectTypeColumn *string `mandatory:"false" json:"objectTypeColumn"`
}

func (m ObjectMetadata) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ObjectMetadata) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
