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

// SubsettingSchemaObjectSummary Summary of a table included in the schema for a subsetting policy
type SubsettingSchemaObjectSummary struct {

	// The unique key that identifies a subsetting table. The key is numeric and unique within a subsetting policy
	Key *string `mandatory:"true" json:"key"`

	// The database schema that contains the subsetting table
	SchemaName *string `mandatory:"true" json:"schemaName"`

	// The name of the database object
	ObjectName *string `mandatory:"true" json:"objectName"`

	// The type of the database object that contains the subsetting table
	ObjectType ObjectTypeEnum `mandatory:"false" json:"objectType,omitempty"`

	// The initial number of rows in this object/table
	InitialRowCount *int64 `mandatory:"false" json:"initialRowCount"`

	// Indicates if the table stats are stale. This can be used to judge the accuracy of initialRowCount
	IsStatsStale *bool `mandatory:"false" json:"isStatsStale"`

	// The date and time the subsetting table was created, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The date and time the subsetting table was last updated, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`
}

func (m SubsettingSchemaObjectSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SubsettingSchemaObjectSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingObjectTypeEnum(string(m.ObjectType)); !ok && m.ObjectType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ObjectType: %s. Supported values are: %s.", m.ObjectType, strings.Join(GetObjectTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
