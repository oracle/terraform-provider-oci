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

// DataAsset Data asset representation. A physical store, or stream, of data known to the data catalog and containing
// one or many data entities, possibly in an organized structure of folders. A data asset is often synonymous
// with a 'System', such as a database, or may be a file container or a message stream.
type DataAsset struct {

	// Unique data asset key that is immutable.
	Key *string `mandatory:"true" json:"key"`

	// A user-friendly display name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Detailed description of the data asset.
	Description *string `mandatory:"false" json:"description"`

	// The data catalog's OCID.
	CatalogId *string `mandatory:"false" json:"catalogId"`

	// External URI that can be used to reference the object. Format will differ based on the type of object.
	ExternalKey *string `mandatory:"false" json:"externalKey"`

	// The key of the object type. Type key's can be found via the '/types' endpoint.
	TypeKey *string `mandatory:"false" json:"typeKey"`

	// The current state of the data asset.
	LifecycleState LifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// A message describing the current state in more detail. An object not in ACTIVE state may have functional limitations,
	// see service documentation for details.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The date and time the data asset was created, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2019-03-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The last time that any change was made to the data asset. An RFC3339 (https://tools.ietf.org/html/rfc3339) formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The last time that a harvest was performed on the data asset. An RFC3339 (https://tools.ietf.org/html/rfc3339) formatted datetime string.
	TimeHarvested *common.SDKTime `mandatory:"false" json:"timeHarvested"`

	// OCID of the user who created the data asset.
	CreatedById *string `mandatory:"false" json:"createdById"`

	// OCID of the user who last modified the data asset.
	UpdatedById *string `mandatory:"false" json:"updatedById"`

	// URI to the data asset instance in the API.
	Uri *string `mandatory:"false" json:"uri"`

	// The list of customized properties along with the values for this object
	CustomPropertyMembers []CustomPropertyGetUsage `mandatory:"false" json:"customPropertyMembers"`

	// The list of data selector patterns used in the harvest for this data asset to derive logical entities.
	DataSelectorPatterns []PatternSummary `mandatory:"false" json:"dataSelectorPatterns"`

	// A map of maps that contains the properties which are specific to the asset type. Each data asset type
	// definition defines it's set of required and optional properties. The map keys are category names and the
	// values are maps of property name to property value. Every property is contained inside of a category. Most
	// data assets have required properties within the "default" category.
	// Example: `{"properties": { "default": { "host": "host1", "port": "1521", "database": "orcl"}}}`
	Properties map[string]map[string]string `mandatory:"false" json:"properties"`
}

func (m DataAsset) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DataAsset) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
