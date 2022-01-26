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

// CustomPropertyTypeUsage Object which describes the indivial object stats for every custom property
type CustomPropertyTypeUsage struct {

	// Unique type key identifier
	TypeId *string `mandatory:"false" json:"typeId"`

	// Name of the type associated with
	TypeName *string `mandatory:"false" json:"typeName"`

	// Number of objects associated with this type
	Count *int `mandatory:"false" json:"count"`

	// If an OCI Event will be emitted when the custom property is modified.
	IsEventEnabled *bool `mandatory:"false" json:"isEventEnabled"`
}

func (m CustomPropertyTypeUsage) String() string {
	return common.PointerString(m)
}
