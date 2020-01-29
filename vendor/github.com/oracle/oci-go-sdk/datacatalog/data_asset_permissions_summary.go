// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Data Catalog API
//
// Use the Data Catalog APIs to collect, organize, find, access, understand, enrich, and activate technical, business, and operational metadata.
//

package datacatalog

import (
	"github.com/oracle/oci-go-sdk/common"
)

// DataAssetPermissionsSummary Permissions object for data assets.
type DataAssetPermissionsSummary struct {

	// An array of permissions.
	UserPermissions []string `mandatory:"false" json:"userPermissions"`

	// The unique key of the parent data asset.
	DataAssetKey *string `mandatory:"false" json:"dataAssetKey"`
}

func (m DataAssetPermissionsSummary) String() string {
	return common.PointerString(m)
}
