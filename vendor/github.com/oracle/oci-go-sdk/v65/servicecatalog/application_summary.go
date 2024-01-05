// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Service Catalog API
//
// Manage solutions in Oracle Cloud Infrastructure Service Catalog.
//

package servicecatalog

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ApplicationSummary The model for summary of an application in service catalog.
type ApplicationSummary struct {

	// Identifier of the application from a service catalog.
	EntityId *string `mandatory:"true" json:"entityId"`

	// The type of an application in the service catalog.
	EntityType *string `mandatory:"true" json:"entityType"`

	// The name that service catalog should use to display this application.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Indicates whether the application is featured.
	IsFeatured *bool `mandatory:"false" json:"isFeatured"`

	Publisher *PublisherSummary `mandatory:"false" json:"publisher"`

	// A short description of the application.
	ShortDescription *string `mandatory:"false" json:"shortDescription"`

	Logo *UploadData `mandatory:"false" json:"logo"`

	// Summary of the pricing types available across all packages in the application.
	PricingType PricingTypeEnumEnum `mandatory:"false" json:"pricingType,omitempty"`

	// The type of the packages withing the application.
	PackageType PackageTypeEnumEnum `mandatory:"false" json:"packageType,omitempty"`
}

func (m ApplicationSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ApplicationSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingPricingTypeEnumEnum(string(m.PricingType)); !ok && m.PricingType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PricingType: %s. Supported values are: %s.", m.PricingType, strings.Join(GetPricingTypeEnumEnumStringValues(), ",")))
	}
	if _, ok := GetMappingPackageTypeEnumEnum(string(m.PackageType)); !ok && m.PackageType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PackageType: %s. Supported values are: %s.", m.PackageType, strings.Join(GetPackageTypeEnumEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
