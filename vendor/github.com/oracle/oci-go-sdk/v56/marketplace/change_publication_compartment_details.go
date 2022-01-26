// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Marketplace Service API
//
// Manage applications in Oracle Cloud Infrastructure Marketplace.
//

package marketplace

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// ChangePublicationCompartmentDetails The model for the parameters needed move a publication from one compartment to another.
type ChangePublicationCompartmentDetails struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment where you want to move the publication.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`
}

func (m ChangePublicationCompartmentDetails) String() string {
	return common.PointerString(m)
}
