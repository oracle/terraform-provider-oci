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

// PrivateApplicationSummary Brief data about an application or a solution, which lives inside the tenancy and may be included into service catalogs.
type PrivateApplicationSummary struct {

	// The lifecycle state of the private application.
	LifecycleState PrivateApplicationLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment where the private application resides.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the private application.
	Id *string `mandatory:"true" json:"id"`

	// The name of the private application.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Type of the packages, which are hosted by the private application.
	PackageType PackageTypeEnumEnum `mandatory:"true" json:"packageType"`

	// The date and time the private application was created, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339)
	// timestamp format.
	// Example: `2021-05-27T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// A short description of the private application.
	ShortDescription *string `mandatory:"false" json:"shortDescription"`

	Logo *UploadData `mandatory:"false" json:"logo"`
}

func (m PrivateApplicationSummary) String() string {
	return common.PointerString(m)
}
