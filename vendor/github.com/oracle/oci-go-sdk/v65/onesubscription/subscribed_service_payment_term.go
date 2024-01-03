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

// SubscribedServicePaymentTerm Payment Term details
type SubscribedServicePaymentTerm struct {

	// Payment Term name
	Name *string `mandatory:"false" json:"name"`

	// Payment Term value
	Value *string `mandatory:"false" json:"value"`

	// Payment term Description
	Description *string `mandatory:"false" json:"description"`

	// Payment term active flag
	IsActive *bool `mandatory:"false" json:"isActive"`

	// Payment term last update date
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// User that created the Payment term
	CreatedBy *string `mandatory:"false" json:"createdBy"`

	// Payment term last update date
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// User that updated the Payment term
	UpdatedBy *string `mandatory:"false" json:"updatedBy"`
}

func (m SubscribedServicePaymentTerm) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SubscribedServicePaymentTerm) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
