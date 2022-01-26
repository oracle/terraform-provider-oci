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

// PrivateApplicationPackageSummary The model for a summary of a private application package.
type PrivateApplicationPackageSummary struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the private application package.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the private application where the package is hosted.
	PrivateApplicationId *string `mandatory:"true" json:"privateApplicationId"`

	// The version of the specified package.
	Version *string `mandatory:"true" json:"version"`

	// The type of the package.
	PackageType PackageTypeEnumEnum `mandatory:"true" json:"packageType"`

	// The date and time the private application package was created, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339)
	// timestamp format.
	// Example: `2021-05-27T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The display name of the specified package.
	DisplayName *string `mandatory:"false" json:"displayName"`
}

func (m PrivateApplicationPackageSummary) String() string {
	return common.PointerString(m)
}
