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

// SubsettedObjectSummary Summary of a subsetted object. A subsetted object is a database table subsetted by a data subsetting request
type SubsettedObjectSummary struct {

	// The name of the schema that contains the subsetted object
	SchemaName *string `mandatory:"true" json:"schemaName"`

	// The name of the object (table or editioning view) subsetted
	ObjectName *string `mandatory:"true" json:"objectName"`

	// The type of the object (table or editioning view) subsetted
	ObjectType ObjectTypeEnum `mandatory:"true" json:"objectType"`

	// The count of rows in the subsetted table before subsetting
	RowCountBeforeSubsetting *int64 `mandatory:"false" json:"rowCountBeforeSubsetting"`

	// The count of rows in the subsetted table after subsetting
	RowCountAfterSubsetting *int64 `mandatory:"false" json:"rowCountAfterSubsetting"`

	// The size of the subsetted table before subsetting in KBs
	SizeBeforeSubsettingInKBs *string `mandatory:"false" json:"sizeBeforeSubsettingInKBs"`

	// The size of the subsetted table after subsetting in KBs
	SizeAfterSubsettingInKBs *string `mandatory:"false" json:"sizeAfterSubsettingInKBs"`
}

func (m SubsettedObjectSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SubsettedObjectSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingObjectTypeEnum(string(m.ObjectType)); !ok && m.ObjectType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ObjectType: %s. Supported values are: %s.", m.ObjectType, strings.Join(GetObjectTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
