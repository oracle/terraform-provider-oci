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

// BulkReplaceServiceCatalogAssociationsDetails The model to replace service catalog associations in bulk.
type BulkReplaceServiceCatalogAssociationsDetails struct {

	// Collection of CreateServiceCatalogAssociationDetails for bulk operation.
	Items []CreateServiceCatalogAssociationDetails `mandatory:"true" json:"items"`
}

func (m BulkReplaceServiceCatalogAssociationsDetails) String() string {
	return common.PointerString(m)
}
