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

// SearchResult The search result object is the definition of an element that is returned as part of search. It contains basic
// information about the object such as key, name and description. The search result also contains the list of tags
// for each object along with other contextual information like the data asset root, folder, or entity parents.
type SearchResult struct {

	// Unique key of the object returned as part of the search result.
	Key *string `mandatory:"false" json:"key"`

	// Name of the object.
	Name *string `mandatory:"false" json:"name"`

	// Detailed description of the object.
	Description *string `mandatory:"false" json:"description"`

	// The date and time the result object was created, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2019-03-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The date and time the result object was updated, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2019-03-25T21:10:29.600Z`
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Array of the tags associated with this object.
	TagSummary []SearchTagSummary `mandatory:"false" json:"tagSummary"`

	// Array of the terms associated with this object.
	TermSummary []SearchTermSummary `mandatory:"false" json:"termSummary"`

	// Name of the object type.
	TypeName *string `mandatory:"false" json:"typeName"`

	// Name of the external object type in the host data asset. For example, column, field, table, view, or file.
	ExternalTypeName *string `mandatory:"false" json:"externalTypeName"`

	// Data type of the object if the object is an attribute. Null otherwise.
	ExternalDataType *string `mandatory:"false" json:"externalDataType"`

	// Unique key of the data asset that is the root parent of this object.
	DataAssetKey *string `mandatory:"false" json:"dataAssetKey"`

	// Type name of the data asset. For example, Oracle, MySQL or Oracle Object Storage.
	DataAssetType *string `mandatory:"false" json:"dataAssetType"`

	// Name of the data asset that is the root parent of this object.
	DataAssetName *string `mandatory:"false" json:"dataAssetName"`

	// Unique key of the folder object if this object is a sub folder, entity, or attribute.
	FolderKey *string `mandatory:"false" json:"folderKey"`

	// Type name of the folder. For example, schema, directory, or topic.
	FolderType *string `mandatory:"false" json:"folderType"`

	// Name of the parent folder object if this object is a sub folder, entity, or attribute.
	FolderName *string `mandatory:"false" json:"folderName"`

	// Unique key of the entity object if this object is an attribute.
	Entitykey *string `mandatory:"false" json:"entitykey"`

	// Type name of the entity. For example, table, view, external table, file, or object.
	EntityType *string `mandatory:"false" json:"entityType"`

	// Name of the parent entity object if this object is an attribute.
	EntityName *string `mandatory:"false" json:"entityName"`

	// Unique id of the parent glossary.
	GlossaryKey *string `mandatory:"false" json:"glossaryKey"`

	// Name of the parent glossary if this object is a term.
	GlossaryName *string `mandatory:"false" json:"glossaryName"`

	// This terms parent term key. Will be null if the term has no parent term.
	ParentTermKey *string `mandatory:"false" json:"parentTermKey"`

	// Name of the parent term. Will be null if the term has no parent term.
	ParentTermName *string `mandatory:"false" json:"parentTermName"`

	// OCID of the user who created the resource.
	CreatedById *string `mandatory:"false" json:"createdById"`

	// OCID of the user who updated the resource.
	UpdatedById *string `mandatory:"false" json:"updatedById"`

	// Absolute path of this resource, which could be a term, folder, entity etc, usually resolvable to this resource through a namespace hierarchy.
	Path *string `mandatory:"false" json:"path"`

	// Optional user friendly business name of the data object. If set, this supplements the harvested display name of the object.
	BusinessName *string `mandatory:"false" json:"businessName"`

	// The current state of the data object.
	LifecycleState LifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// Type name of the attribute. For example - complex, primitive, or array.
	AttributeType *string `mandatory:"false" json:"attributeType"`

	// Expression for logical entities against which names of dataObjects will be matched.
	Expression *string `mandatory:"false" json:"expression"`

	// Custom properties defined by users.
	CustomProperties []FacetedSearchCustomProperty `mandatory:"false" json:"customProperties"`

	// A map of maps that contains the properties which are specific to the element type in the search result.
	// The map keys are category names and the values are maps of property name to property value. Every property
	// is contained inside of a category. Most element types have required properties within the "default" category.
	// Example: `{"properties": { "default": { "key1": "value1"}}}`
	Properties map[string]map[string]string `mandatory:"false" json:"properties"`
}

func (m SearchResult) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SearchResult) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
