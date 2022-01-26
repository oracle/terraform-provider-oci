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

// ChangeServiceCatalogCompartmentDetails The model for the parameters needed move a service catalog from one compartment to another.
type ChangeServiceCatalogCompartmentDetails struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment where you want to move the service catalog.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`
}

func (m ChangeServiceCatalogCompartmentDetails) String() string {
	return common.PointerString(m)
}
