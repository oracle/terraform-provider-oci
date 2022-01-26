// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Catalog API
//
// Use the Data Catalog APIs to collect, organize, find, access, understand, enrich, and activate technical, business, and operational metadata.
// For more information, see Data Catalog (https://docs.oracle.com/iaas/data-catalog/home.htm).
//

package datacatalog

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// CustomPropertyGetUsage Details of a single custom property
type CustomPropertyGetUsage struct {

	// Unique Identifier of the attribute which is ID
	Key *string `mandatory:"false" json:"key"`

	// Display name of the custom property
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Description of the custom property
	Description *string `mandatory:"false" json:"description"`

	// The custom property value
	Value *string `mandatory:"false" json:"value"`

	// The data type of the custom property
	DataType CustomPropertyDataTypeEnum `mandatory:"false" json:"dataType,omitempty"`

	// Namespace name of the custom property
	NamespaceName *string `mandatory:"false" json:"namespaceName"`

	// Unique namespace key that is immutable
	NamespaceKey *string `mandatory:"false" json:"namespaceKey"`

	// If this field allows multiple values to be set
	IsMultiValued *bool `mandatory:"false" json:"isMultiValued"`

	// If this field is a hidden field
	IsHidden *bool `mandatory:"false" json:"isHidden"`

	// If this field is a editable field
	IsEditable *bool `mandatory:"false" json:"isEditable"`

	// If this field is displayed in a list view of applicable objects.
	IsShownInList *bool `mandatory:"false" json:"isShownInList"`

	// If an OCI Event will be emitted when the custom property is modified.
	IsEventEnabled *bool `mandatory:"false" json:"isEventEnabled"`

	// Is this property allowed to have list of values
	IsListType *bool `mandatory:"false" json:"isListType"`

	// Allowed values for the custom property if any
	AllowedValues []string `mandatory:"false" json:"allowedValues"`
}

func (m CustomPropertyGetUsage) String() string {
	return common.PointerString(m)
}
