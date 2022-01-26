// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Service Catalog API
//
// Manage solutions in Oracle Cloud Infrastructure Service Catalog.
//

package servicecatalog

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// UpdatePrivateApplicationDetails The model for the parameters needed to update a private application.
type UpdatePrivateApplicationDetails struct {

	// The name of the private application.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// A short description of the private application.
	ShortDescription *string `mandatory:"false" json:"shortDescription"`

	// A long description of the private application.
	LongDescription *string `mandatory:"false" json:"longDescription"`

	// Base64-encoded logo to use as the private application icon.
	// Template icon file requirements: PNG format, 50 KB maximum, 130 x 130 pixels.
	LogoFileBase64Encoded *string `mandatory:"false" json:"logoFileBase64Encoded"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`
}

func (m UpdatePrivateApplicationDetails) String() string {
	return common.PointerString(m)
}
