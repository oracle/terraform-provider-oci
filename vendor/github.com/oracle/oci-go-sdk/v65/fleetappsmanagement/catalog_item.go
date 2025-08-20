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

// CatalogItem A description of a CatalogItem resource.
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized, talk to
// an administrator. If you're an administrator who needs to write policies to give users access, see
// Getting Started with Policies (https://docs.oracle.com/iaas/Content/Identity/policiesgs/get-started-with-policies.htm).
type CatalogItem struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the catalog.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Config source type Eg: STACK_TEMPLATE_CATALOG_SOURCE, PAR_CATALOG_SOURCE, GIT_CATALOG_SOURCE, MARKETPLACE_CATALOG_SOURCE.
	ConfigSourceType CatalogItemConfigSourceTypeEnum `mandatory:"true" json:"configSourceType"`

	// Description about the catalog item.
	Description *string `mandatory:"true" json:"description"`

	// The catalog listing Id.
	ListingId *string `mandatory:"true" json:"listingId"`

	// The catalog package version.
	ListingVersion *string `mandatory:"true" json:"listingVersion"`

	// A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Config package type Eg: TF_PACKAGE, NON_TF_PACKAGE, CONFIG_FILE.
	PackageType CatalogItemPackageTypeEnum `mandatory:"true" json:"packageType"`

	// The current state of the CatalogItem.
	LifecycleState CatalogItemLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The details of lifecycle state CatalogItem.
	LifecycleDetails *string `mandatory:"true" json:"lifecycleDetails"`

	// The date and time the CatalogItem was created, in the format defined by RFC 3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The date and time the CatalogItem was updated, in the format defined by RFC 3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The indicator to append Public Items from the root compartment to any query, when set to TRUE.
	ShouldListPublicItems *bool `mandatory:"false" json:"shouldListPublicItems"`

	CatalogSourcePayload CatalogSourcePayload `mandatory:"false" json:"catalogSourcePayload"`

	CatalogResultPayload CatalogResultPayload `mandatory:"false" json:"catalogResultPayload"`

	// Version description about the catalog item.
	VersionDescription *string `mandatory:"false" json:"versionDescription"`

	// Short description about the catalog item.
	ShortDescription *string `mandatory:"false" json:"shortDescription"`

	// The date and time the CatalogItem was released, in the format defined by RFC 3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeReleased *common.SDKTime `mandatory:"false" json:"timeReleased"`

	// The date and time the CatalogItem was last checked by backfill job, in the format defined by
	// RFC 3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeBackfillLastChecked *common.SDKTime `mandatory:"false" json:"timeBackfillLastChecked"`

	// The date and time the CatalogItem was last checked, in the format defined by
	// RFC 3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeLastChecked *common.SDKTime `mandatory:"false" json:"timeLastChecked"`

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

func (m CatalogItem) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CatalogItem) ValidateEnumValue() (bool, error) {
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
func (m *CatalogItem) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		ShouldListPublicItems   *bool                             `json:"shouldListPublicItems"`
		CatalogSourcePayload    catalogsourcepayload              `json:"catalogSourcePayload"`
		CatalogResultPayload    catalogresultpayload              `json:"catalogResultPayload"`
		VersionDescription      *string                           `json:"versionDescription"`
		ShortDescription        *string                           `json:"shortDescription"`
		TimeReleased            *common.SDKTime                   `json:"timeReleased"`
		TimeBackfillLastChecked *common.SDKTime                   `json:"timeBackfillLastChecked"`
		TimeLastChecked         *common.SDKTime                   `json:"timeLastChecked"`
		FreeformTags            map[string]string                 `json:"freeformTags"`
		DefinedTags             map[string]map[string]interface{} `json:"definedTags"`
		SystemTags              map[string]map[string]interface{} `json:"systemTags"`
		Id                      *string                           `json:"id"`
		CompartmentId           *string                           `json:"compartmentId"`
		ConfigSourceType        CatalogItemConfigSourceTypeEnum   `json:"configSourceType"`
		Description             *string                           `json:"description"`
		ListingId               *string                           `json:"listingId"`
		ListingVersion          *string                           `json:"listingVersion"`
		DisplayName             *string                           `json:"displayName"`
		PackageType             CatalogItemPackageTypeEnum        `json:"packageType"`
		LifecycleState          CatalogItemLifecycleStateEnum     `json:"lifecycleState"`
		LifecycleDetails        *string                           `json:"lifecycleDetails"`
		TimeCreated             *common.SDKTime                   `json:"timeCreated"`
		TimeUpdated             *common.SDKTime                   `json:"timeUpdated"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.ShouldListPublicItems = model.ShouldListPublicItems

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

	m.VersionDescription = model.VersionDescription

	m.ShortDescription = model.ShortDescription

	m.TimeReleased = model.TimeReleased

	m.TimeBackfillLastChecked = model.TimeBackfillLastChecked

	m.TimeLastChecked = model.TimeLastChecked

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

	return
}

// CatalogItemConfigSourceTypeEnum Enum with underlying type: string
type CatalogItemConfigSourceTypeEnum string

// Set of constants representing the allowable values for CatalogItemConfigSourceTypeEnum
const (
	CatalogItemConfigSourceTypeParCatalogSource           CatalogItemConfigSourceTypeEnum = "PAR_CATALOG_SOURCE"
	CatalogItemConfigSourceTypeGitCatalogSource           CatalogItemConfigSourceTypeEnum = "GIT_CATALOG_SOURCE"
	CatalogItemConfigSourceTypeMarketplaceCatalogSource   CatalogItemConfigSourceTypeEnum = "MARKETPLACE_CATALOG_SOURCE"
	CatalogItemConfigSourceTypeStackTemplateCatalogSource CatalogItemConfigSourceTypeEnum = "STACK_TEMPLATE_CATALOG_SOURCE"
)

var mappingCatalogItemConfigSourceTypeEnum = map[string]CatalogItemConfigSourceTypeEnum{
	"PAR_CATALOG_SOURCE":            CatalogItemConfigSourceTypeParCatalogSource,
	"GIT_CATALOG_SOURCE":            CatalogItemConfigSourceTypeGitCatalogSource,
	"MARKETPLACE_CATALOG_SOURCE":    CatalogItemConfigSourceTypeMarketplaceCatalogSource,
	"STACK_TEMPLATE_CATALOG_SOURCE": CatalogItemConfigSourceTypeStackTemplateCatalogSource,
}

var mappingCatalogItemConfigSourceTypeEnumLowerCase = map[string]CatalogItemConfigSourceTypeEnum{
	"par_catalog_source":            CatalogItemConfigSourceTypeParCatalogSource,
	"git_catalog_source":            CatalogItemConfigSourceTypeGitCatalogSource,
	"marketplace_catalog_source":    CatalogItemConfigSourceTypeMarketplaceCatalogSource,
	"stack_template_catalog_source": CatalogItemConfigSourceTypeStackTemplateCatalogSource,
}

// GetCatalogItemConfigSourceTypeEnumValues Enumerates the set of values for CatalogItemConfigSourceTypeEnum
func GetCatalogItemConfigSourceTypeEnumValues() []CatalogItemConfigSourceTypeEnum {
	values := make([]CatalogItemConfigSourceTypeEnum, 0)
	for _, v := range mappingCatalogItemConfigSourceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCatalogItemConfigSourceTypeEnumStringValues Enumerates the set of values in String for CatalogItemConfigSourceTypeEnum
func GetCatalogItemConfigSourceTypeEnumStringValues() []string {
	return []string{
		"PAR_CATALOG_SOURCE",
		"GIT_CATALOG_SOURCE",
		"MARKETPLACE_CATALOG_SOURCE",
		"STACK_TEMPLATE_CATALOG_SOURCE",
	}
}

// GetMappingCatalogItemConfigSourceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCatalogItemConfigSourceTypeEnum(val string) (CatalogItemConfigSourceTypeEnum, bool) {
	enum, ok := mappingCatalogItemConfigSourceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// CatalogItemPackageTypeEnum Enum with underlying type: string
type CatalogItemPackageTypeEnum string

// Set of constants representing the allowable values for CatalogItemPackageTypeEnum
const (
	CatalogItemPackageTypeTfPackage    CatalogItemPackageTypeEnum = "TF_PACKAGE"
	CatalogItemPackageTypeNonTfPackage CatalogItemPackageTypeEnum = "NON_TF_PACKAGE"
	CatalogItemPackageTypeConfigFile   CatalogItemPackageTypeEnum = "CONFIG_FILE"
)

var mappingCatalogItemPackageTypeEnum = map[string]CatalogItemPackageTypeEnum{
	"TF_PACKAGE":     CatalogItemPackageTypeTfPackage,
	"NON_TF_PACKAGE": CatalogItemPackageTypeNonTfPackage,
	"CONFIG_FILE":    CatalogItemPackageTypeConfigFile,
}

var mappingCatalogItemPackageTypeEnumLowerCase = map[string]CatalogItemPackageTypeEnum{
	"tf_package":     CatalogItemPackageTypeTfPackage,
	"non_tf_package": CatalogItemPackageTypeNonTfPackage,
	"config_file":    CatalogItemPackageTypeConfigFile,
}

// GetCatalogItemPackageTypeEnumValues Enumerates the set of values for CatalogItemPackageTypeEnum
func GetCatalogItemPackageTypeEnumValues() []CatalogItemPackageTypeEnum {
	values := make([]CatalogItemPackageTypeEnum, 0)
	for _, v := range mappingCatalogItemPackageTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCatalogItemPackageTypeEnumStringValues Enumerates the set of values in String for CatalogItemPackageTypeEnum
func GetCatalogItemPackageTypeEnumStringValues() []string {
	return []string{
		"TF_PACKAGE",
		"NON_TF_PACKAGE",
		"CONFIG_FILE",
	}
}

// GetMappingCatalogItemPackageTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCatalogItemPackageTypeEnum(val string) (CatalogItemPackageTypeEnum, bool) {
	enum, ok := mappingCatalogItemPackageTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// CatalogItemLifecycleStateEnum Enum with underlying type: string
type CatalogItemLifecycleStateEnum string

// Set of constants representing the allowable values for CatalogItemLifecycleStateEnum
const (
	CatalogItemLifecycleStateCreating CatalogItemLifecycleStateEnum = "CREATING"
	CatalogItemLifecycleStateUpdating CatalogItemLifecycleStateEnum = "UPDATING"
	CatalogItemLifecycleStateActive   CatalogItemLifecycleStateEnum = "ACTIVE"
	CatalogItemLifecycleStateDeleting CatalogItemLifecycleStateEnum = "DELETING"
	CatalogItemLifecycleStateDeleted  CatalogItemLifecycleStateEnum = "DELETED"
	CatalogItemLifecycleStateFailed   CatalogItemLifecycleStateEnum = "FAILED"
)

var mappingCatalogItemLifecycleStateEnum = map[string]CatalogItemLifecycleStateEnum{
	"CREATING": CatalogItemLifecycleStateCreating,
	"UPDATING": CatalogItemLifecycleStateUpdating,
	"ACTIVE":   CatalogItemLifecycleStateActive,
	"DELETING": CatalogItemLifecycleStateDeleting,
	"DELETED":  CatalogItemLifecycleStateDeleted,
	"FAILED":   CatalogItemLifecycleStateFailed,
}

var mappingCatalogItemLifecycleStateEnumLowerCase = map[string]CatalogItemLifecycleStateEnum{
	"creating": CatalogItemLifecycleStateCreating,
	"updating": CatalogItemLifecycleStateUpdating,
	"active":   CatalogItemLifecycleStateActive,
	"deleting": CatalogItemLifecycleStateDeleting,
	"deleted":  CatalogItemLifecycleStateDeleted,
	"failed":   CatalogItemLifecycleStateFailed,
}

// GetCatalogItemLifecycleStateEnumValues Enumerates the set of values for CatalogItemLifecycleStateEnum
func GetCatalogItemLifecycleStateEnumValues() []CatalogItemLifecycleStateEnum {
	values := make([]CatalogItemLifecycleStateEnum, 0)
	for _, v := range mappingCatalogItemLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetCatalogItemLifecycleStateEnumStringValues Enumerates the set of values in String for CatalogItemLifecycleStateEnum
func GetCatalogItemLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingCatalogItemLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCatalogItemLifecycleStateEnum(val string) (CatalogItemLifecycleStateEnum, bool) {
	enum, ok := mappingCatalogItemLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
