// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Web Application Acceleration and Security Services API
//
// OCI Web Application Acceleration and Security Services
//

package waas

import (
	"github.com/oracle/oci-go-sdk/common"
)

// ChangeWaasPolicyCompartmentDetails The representation of ChangeWaasPolicyCompartmentDetails
type ChangeWaasPolicyCompartmentDetails struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment into which the resource should be moved.
	// For information about moving resources between compartments, see Moving Resources to a Different Compartment (https://docs.cloud.oracle.com/iaas/Content/Identity/Tasks/managingcompartments.htm#moveRes).
	CompartmentId *string `mandatory:"true" json:"compartmentId"`
}

func (m ChangeWaasPolicyCompartmentDetails) String() string {
	return common.PointerString(m)
}
