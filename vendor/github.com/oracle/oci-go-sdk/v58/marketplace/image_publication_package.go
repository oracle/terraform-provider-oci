// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Marketplace Service API
//
// Use the Marketplace API to manage applications in Oracle Cloud Infrastructure Marketplace. For more information, see Overview of Marketplace (https://docs.cloud.oracle.com/Content/Marketplace/Concepts/marketoverview.htm)
//

package marketplace

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// ImagePublicationPackage A publication package for image publications.
type ImagePublicationPackage struct {

	// The ID of the listing that the specified package belongs to.
	ListingId *string `mandatory:"true" json:"listingId"`

	// The package version.
	Version *string `mandatory:"true" json:"version"`

	// A description of the package.
	Description *string `mandatory:"false" json:"description"`

	// The unique identifier for the package resource.
	ResourceId *string `mandatory:"false" json:"resourceId"`

	// The date and time the publication package was created, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339)
	// timestamp format.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	OperatingSystem *OperatingSystem `mandatory:"false" json:"operatingSystem"`

	// The ID of the listing resource associated with this publication package. For more information, see AppCatalogListing (https://docs.cloud.oracle.com/en-us/iaas/api/#/en/iaas/latest/AppCatalogListing/) in the Core Services API.
	AppCatalogListingId *string `mandatory:"false" json:"appCatalogListingId"`

	// The resource version of the listing resource associated with this publication package.
	AppCatalogListingResourceVersion *string `mandatory:"false" json:"appCatalogListingResourceVersion"`

	// The ID of the image that corresponds to the package.
	ImageId *string `mandatory:"false" json:"imageId"`
}

//GetDescription returns Description
func (m ImagePublicationPackage) GetDescription() *string {
	return m.Description
}

//GetListingId returns ListingId
func (m ImagePublicationPackage) GetListingId() *string {
	return m.ListingId
}

//GetVersion returns Version
func (m ImagePublicationPackage) GetVersion() *string {
	return m.Version
}

//GetResourceId returns ResourceId
func (m ImagePublicationPackage) GetResourceId() *string {
	return m.ResourceId
}

//GetTimeCreated returns TimeCreated
func (m ImagePublicationPackage) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

//GetOperatingSystem returns OperatingSystem
func (m ImagePublicationPackage) GetOperatingSystem() *OperatingSystem {
	return m.OperatingSystem
}

func (m ImagePublicationPackage) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ImagePublicationPackage) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m ImagePublicationPackage) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeImagePublicationPackage ImagePublicationPackage
	s := struct {
		DiscriminatorParam string `json:"packageType"`
		MarshalTypeImagePublicationPackage
	}{
		"IMAGE",
		(MarshalTypeImagePublicationPackage)(m),
	}

	return json.Marshal(&s)
}
