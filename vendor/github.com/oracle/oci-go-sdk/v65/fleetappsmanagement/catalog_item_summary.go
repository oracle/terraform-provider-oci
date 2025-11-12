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

// CatalogItemSummary Summary information about a CatalogItem.
type CatalogItemSummary struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the catalog.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Config source type Eg: STACK_TEMPLATE_CATALOG_SOURCE, PAR_CATALOG_SOURCE, GIT_CATALOG_SOURCE, MARKETPLACE_CATALOG_SOURCE.
	ConfigSourceType CatalogItemConfigSourceTypeEnum `mandatory:"true" json:"configSourceType"`

	// The description of the catalogItem.
	Description *string `mandatory:"true" json:"description"`

	// The catalog listing Id.
	ListingId *string `mandatory:"true" json:"listingId"`

	// The catalog package version.
	ListingVersion *string `mandatory:"true" json:"listingVersion"`

	// A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Config package type Eg: BASE_PACKAGE, EXTENSION_PACKAGE.
	PackageType CatalogItemPackageTypeEnum `mandatory:"true" json:"packageType"`

	// The current state of the CatalogItem.
	LifecycleState CatalogItemLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The details of lifecycle state CatalogItem.
	LifecycleDetails *string `mandatory:"true" json:"lifecycleDetails"`

	// The date and time the CatalogItem was created, in the format defined by
	// RFC 3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The date and time the CatalogItem was updated, in the format defined by
	// RFC 3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// Version description about the catalog item.
	VersionDescription *string `mandatory:"true" json:"versionDescription"`

	// Short description about the catalog item.
	ShortDescription *string `mandatory:"true" json:"shortDescription"`

	// The date and time the CatalogItem was released, in the format defined by
	// RFC 3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeReleased *common.SDKTime `mandatory:"true" json:"timeReleased"`

	CatalogSourcePayload CatalogSourcePayload `mandatory:"false" json:"catalogSourcePayload"`

	CatalogResultPayload CatalogResultPayload `mandatory:"false" json:"catalogResultPayload"`

	// Indicates if the CatalogItem is immutable or not.
	IsItemLocked *bool `mandatory:"false" json:"isItemLocked"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m CatalogItemSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CatalogItemSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCatalogItemConfigSourceTypeEnum(string(m.ConfigSourceType)); !ok && m.ConfigSourceType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ConfigSourceType: %s. Supported values are: %s.", m.ConfigSourceType, strings.Join(GetCatalogItemConfigSourceTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCatalogItemPackageTypeEnum(string(m.PackageType)); !ok && m.PackageType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PackageType: %s. Supported values are: %s.", m.PackageType, strings.Join(GetCatalogItemPackageTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCatalogItemLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetCatalogItemLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *CatalogItemSummary) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		CatalogSourcePayload catalogsourcepayload              `json:"catalogSourcePayload"`
		CatalogResultPayload catalogresultpayload              `json:"catalogResultPayload"`
		IsItemLocked         *bool                             `json:"isItemLocked"`
		FreeformTags         map[string]string                 `json:"freeformTags"`
		DefinedTags          map[string]map[string]interface{} `json:"definedTags"`
		SystemTags           map[string]map[string]interface{} `json:"systemTags"`
		Id                   *string                           `json:"id"`
		CompartmentId        *string                           `json:"compartmentId"`
		ConfigSourceType     CatalogItemConfigSourceTypeEnum   `json:"configSourceType"`
		Description          *string                           `json:"description"`
		ListingId            *string                           `json:"listingId"`
		ListingVersion       *string                           `json:"listingVersion"`
		DisplayName          *string                           `json:"displayName"`
		PackageType          CatalogItemPackageTypeEnum        `json:"packageType"`
		LifecycleState       CatalogItemLifecycleStateEnum     `json:"lifecycleState"`
		LifecycleDetails     *string                           `json:"lifecycleDetails"`
		TimeCreated          *common.SDKTime                   `json:"timeCreated"`
		TimeUpdated          *common.SDKTime                   `json:"timeUpdated"`
		VersionDescription   *string                           `json:"versionDescription"`
		ShortDescription     *string                           `json:"shortDescription"`
		TimeReleased         *common.SDKTime                   `json:"timeReleased"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	nn, e = model.CatalogSourcePayload.UnmarshalPolymorphicJSON(model.CatalogSourcePayload.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.CatalogSourcePayload = nn.(CatalogSourcePayload)
	} else {
		m.CatalogSourcePayload = nil
	}

	nn, e = model.CatalogResultPayload.UnmarshalPolymorphicJSON(model.CatalogResultPayload.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.CatalogResultPayload = nn.(CatalogResultPayload)
	} else {
		m.CatalogResultPayload = nil
	}

	m.IsItemLocked = model.IsItemLocked

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.SystemTags = model.SystemTags

	m.Id = model.Id

	m.CompartmentId = model.CompartmentId

	m.ConfigSourceType = model.ConfigSourceType

	m.Description = model.Description

	m.ListingId = model.ListingId

	m.ListingVersion = model.ListingVersion

	m.DisplayName = model.DisplayName

	m.PackageType = model.PackageType

	m.LifecycleState = model.LifecycleState

	m.LifecycleDetails = model.LifecycleDetails

	m.TimeCreated = model.TimeCreated

	m.TimeUpdated = model.TimeUpdated

	m.VersionDescription = model.VersionDescription

	m.ShortDescription = model.ShortDescription

	m.TimeReleased = model.TimeReleased

	return
}
