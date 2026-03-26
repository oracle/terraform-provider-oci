// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Ops Insights API
//
// Use the Ops Insights API to perform data extraction operations to obtain database
// resource utilization, performance statistics, and reference information. For more information,
// see About Oracle Cloud Infrastructure Ops Insights (https://docs.oracle.com/iaas/en-us/iaas/operations-insights/doc/operations-insights.html).
//

package opsi

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ChargebackPlanDetails Object containing chargeback plan details
type ChargebackPlanDetails struct {

	// OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of OPSI Chargeback plan resource.
	PlanId *string `mandatory:"false" json:"planId"`

	// Chargeback Plan type of the chargeback entity. For an Exadata it can be WEIGHTED_ALLOCATION, EQUAL_ALLOCATION, UNUSED_ALLOCATION.
	PlanType *string `mandatory:"false" json:"planType"`

	// The date and time the chargeback plan was enabled on the resource, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeEnabled *common.SDKTime `mandatory:"false" json:"timeEnabled"`
}

func (m ChargebackPlanDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ChargebackPlanDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
