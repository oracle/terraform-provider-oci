// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Catalog API
//
// Use the Data Catalog APIs to collect, organize, find, access, understand, enrich, and activate technical, business, and operational metadata.
// For more information, see Data Catalog (https://docs.oracle.com/iaas/data-catalog/home.htm).
//

package datacatalog

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// FacetedSearchStringFilterRequest Object with string filter criteria
type FacetedSearchStringFilterRequest struct {

	// String/boolean/numerical field name that needs to be filtered by.
	// Acceptable field names: CatalogType, AttributeType, FolderType, DataAssetType, CreatedBy, UpdatedBy, Term, Tag, DataAssetName, LifeCycleState.
	Field *string `mandatory:"false" json:"field"`

	// Array of values that the search results needs to be filtered by.
	// Acceptable values for field 'CatalogType': DataAsset, Folder, DataEntity, Attribute, Term, Category, Glossary, Pattern, Job, Schedule, CustomProperty.
	// For other fields, acceptable values can be derived by inspecting the data object.
	Values []string `mandatory:"false" json:"values"`
}

func (m FacetedSearchStringFilterRequest) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m FacetedSearchStringFilterRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
