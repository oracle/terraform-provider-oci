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
