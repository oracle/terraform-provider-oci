// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fleet Application Management Service API
//
// Fleet Application Management provides a centralized platform to help you automate resource management tasks, validate patch compliance, and enhance operational efficiency across an enterprise.
//

package fleetappsmanagement

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateCatalogItemDetails The data to create a CatalogItem.
type CreateCatalogItemDetails struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Config source type Eg: STACK_TEMPLATE_CATALOG_SOURCE, PAR_CATALOG_SOURCE, GIT_CATALOG_SOURCE, MARKETPLACE_CATALOG_SOURCE.
	ConfigSourceType CatalogItemConfigSourceTypeEnum `mandatory:"true" json:"configSourceType"`

	// The description of the CatalogItem.
	Description *string `mandatory:"true" json:"description"`

	// The CatalogItem name.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Config package type Eg: TF_PACKAGE, NON_TF_PACKAGE, CONFIG_FILE.
	PackageType CatalogItemPackageTypeEnum `mandatory:"true" json:"packageType"`

	// Version description about the catalog item.
	VersionDescription *string `mandatory:"false" json:"versionDescription"`

	// Short description about the catalog item.
	ShortDescription *string `mandatory:"false" json:"shortDescription"`

	// The date and time the CatalogItem was released, in the format defined by RFC 3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeReleased *common.SDKTime `mandatory:"false" json:"timeReleased"`

	CatalogSourcePayload CatalogSourcePayload `mandatory:"false" json:"catalogSourcePayload"`

	// The catalog listing Id.
	ListingId *string `mandatory:"false" json:"listingId"`

	// The catalog package version.
	ListingVersion *string `mandatory:"false" json:"listingVersion"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateCatalogItemDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateCatalogItemDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCatalogItemConfigSourceTypeEnum(string(m.ConfigSourceType)); !ok && m.ConfigSourceType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ConfigSourceType: %s. Supported values are: %s.", m.ConfigSourceType, strings.Join(GetCatalogItemConfigSourceTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCatalogItemPackageTypeEnum(string(m.PackageType)); !ok && m.PackageType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PackageType: %s. Supported values are: %s.", m.PackageType, strings.Join(GetCatalogItemPackageTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *CreateCatalogItemDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		VersionDescription   *string                           `json:"versionDescription"`
		ShortDescription     *string                           `json:"shortDescription"`
		TimeReleased         *common.SDKTime                   `json:"timeReleased"`
		CatalogSourcePayload catalogsourcepayload              `json:"catalogSourcePayload"`
		ListingId            *string                           `json:"listingId"`
		ListingVersion       *string                           `json:"listingVersion"`
		FreeformTags         map[string]string                 `json:"freeformTags"`
		DefinedTags          map[string]map[string]interface{} `json:"definedTags"`
		CompartmentId        *string                           `json:"compartmentId"`
		ConfigSourceType     CatalogItemConfigSourceTypeEnum   `json:"configSourceType"`
		Description          *string                           `json:"description"`
		DisplayName          *string                           `json:"displayName"`
		PackageType          CatalogItemPackageTypeEnum        `json:"packageType"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.VersionDescription = model.VersionDescription

	m.ShortDescription = model.ShortDescription

	m.TimeReleased = model.TimeReleased

	nn, e = model.CatalogSourcePayload.UnmarshalPolymorphicJSON(model.CatalogSourcePayload.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.CatalogSourcePayload = nn.(CatalogSourcePayload)
	} else {
		m.CatalogSourcePayload = nil
	}

	m.ListingId = model.ListingId

	m.ListingVersion = model.ListingVersion

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.CompartmentId = model.CompartmentId

	m.ConfigSourceType = model.ConfigSourceType

	m.Description = model.Description

	m.DisplayName = model.DisplayName

	m.PackageType = model.PackageType

	return
}
