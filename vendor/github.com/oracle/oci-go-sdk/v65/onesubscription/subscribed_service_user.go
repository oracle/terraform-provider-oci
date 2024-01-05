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

// SubscribedServiceUser User.
type SubscribedServiceUser struct {

	// Name.
	Name *string `mandatory:"false" json:"name"`

	// Username.
	Username *string `mandatory:"false" json:"username"`

	// First name.
	FirstName *string `mandatory:"false" json:"firstName"`

	// Last name.
	LastName *string `mandatory:"false" json:"lastName"`

	// Email.
	Email *string `mandatory:"false" json:"email"`

	// TCA contact ID.
	TcaContactId *int64 `mandatory:"false" json:"tcaContactId"`

	// TCA customer account site ID.
	TcaCustAccntSiteId *int64 `mandatory:"false" json:"tcaCustAccntSiteId"`

	// TCA party ID.
	TcaPartyId *int64 `mandatory:"false" json:"tcaPartyId"`
}

func (m SubscribedServiceUser) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SubscribedServiceUser) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
