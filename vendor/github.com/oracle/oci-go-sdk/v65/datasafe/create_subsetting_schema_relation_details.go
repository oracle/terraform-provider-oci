// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateSubsettingSchemaRelationDetails Summary of a relationship between tables in the subsetting for a subsetting policy.
type CreateSubsettingSchemaRelationDetails struct {

	// The database schema that contains the parent subsetting table
	ParentSchemaName *string `mandatory:"true" json:"parentSchemaName"`

	// The name of the parent subsetting table
	ParentObjectName *string `mandatory:"true" json:"parentObjectName"`

	// Unique identifiers identifying the parents columns in the relation.
	ParentColumns []string `mandatory:"true" json:"parentColumns"`

	// The database schema that contains the child subsetting table
	ChildSchemaName *string `mandatory:"true" json:"childSchemaName"`

	// The name of the child subsetting table
	ChildObjectName *string `mandatory:"true" json:"childObjectName"`

	// Unique identifiers identifying the child columns in the relation.
	ChildColumns []string `mandatory:"true" json:"childColumns"`

	// The key that identifies the parent subsetting table in this relation.
	ParentObjectKey *string `mandatory:"false" json:"parentObjectKey"`

	// The key that identifies the child subsetting table in this relation.
	ChildObjectKey *string `mandatory:"false" json:"childObjectKey"`
}

func (m CreateSubsettingSchemaRelationDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateSubsettingSchemaRelationDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
