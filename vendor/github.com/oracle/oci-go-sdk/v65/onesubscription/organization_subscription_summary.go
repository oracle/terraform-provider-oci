// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OneSubscription APIs
//
// OneSubscription APIs
//

package onesubscription

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// OrganizationSubscriptionSummary Subscription summary
type OrganizationSubscriptionSummary struct {

	// SPM internal Subscription ID
	Id *string `mandatory:"true" json:"id"`

	// Customer friendly service name provided by PRG
	ServiceName *string `mandatory:"false" json:"serviceName"`

	// Subscription Type i.e. IAAS,SAAS,PAAS
	Type *string `mandatory:"false" json:"type"`

	// Status of the plan
	Status *string `mandatory:"false" json:"status"`

	// Represents the date when the first service of the subscription was activated
	TimeStart *common.SDKTime `mandatory:"false" json:"timeStart"`

	// Represents the date when the last service of the subscription ends
	TimeEnd *common.SDKTime `mandatory:"false" json:"timeEnd"`

	Currency *OrgnizationSubsCurrency `mandatory:"false" json:"currency"`

	// Total aggregate TCLV of all lines for the subscription including expired, active, and signed
	TotalValue *string `mandatory:"false" json:"totalValue"`
}

func (m OrganizationSubscriptionSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OrganizationSubscriptionSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
