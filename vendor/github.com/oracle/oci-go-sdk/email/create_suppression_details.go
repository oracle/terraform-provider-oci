// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Email Delivery API
//
// API for the Email Delivery service. Use this API to send high-volume, application-generated
// emails. For more information, see Overview of the Email Delivery Service (https://docs.cloud.oracle.com/iaas/Content/Email/Concepts/overview.htm).
//

package email

import (
	"github.com/oracle/oci-go-sdk/common"
)

// CreateSuppressionDetails The details needed for creating a single suppression.
type CreateSuppressionDetails struct {

	// The OCID of the compartment to contain the suppression. Since
	// suppressions are at the customer level, this must be the tenancy
	// OCID.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The recipient email address of the suppression.
	EmailAddress *string `mandatory:"true" json:"emailAddress"`
}

func (m CreateSuppressionDetails) String() string {
	return common.PointerString(m)
}
