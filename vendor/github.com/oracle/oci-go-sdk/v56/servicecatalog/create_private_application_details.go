// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Service Catalog API
//
// Manage solutions in Oracle Cloud Infrastructure Service Catalog.
//

package servicecatalog

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v56/common"
)

// CreatePrivateApplicationDetails The model for the parameters needed to create a private application.
type CreatePrivateApplicationDetails struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment where you want to create the private application.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The name of the private application.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// A short description of the private application.
	ShortDescription *string `mandatory:"true" json:"shortDescription"`

	PackageDetails CreatePrivateApplicationPackage `mandatory:"true" json:"packageDetails"`

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

func (m CreatePrivateApplicationDetails) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *CreatePrivateApplicationDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		LongDescription       *string                           `json:"longDescription"`
		LogoFileBase64Encoded *string                           `json:"logoFileBase64Encoded"`
		DefinedTags           map[string]map[string]interface{} `json:"definedTags"`
		FreeformTags          map[string]string                 `json:"freeformTags"`
		CompartmentId         *string                           `json:"compartmentId"`
		DisplayName           *string                           `json:"displayName"`
		ShortDescription      *string                           `json:"shortDescription"`
		PackageDetails        createprivateapplicationpackage   `json:"packageDetails"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.LongDescription = model.LongDescription

	m.LogoFileBase64Encoded = model.LogoFileBase64Encoded

	m.DefinedTags = model.DefinedTags

	m.FreeformTags = model.FreeformTags

	m.CompartmentId = model.CompartmentId

	m.DisplayName = model.DisplayName

	m.ShortDescription = model.ShortDescription

	nn, e = model.PackageDetails.UnmarshalPolymorphicJSON(model.PackageDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.PackageDetails = nn.(CreatePrivateApplicationPackage)
	} else {
		m.PackageDetails = nil
	}

	return
}
