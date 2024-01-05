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

// SubscribedServiceBusinessPartner Business partner.
type SubscribedServiceBusinessPartner struct {

	// Commercial name also called customer name.
	Name *string `mandatory:"false" json:"name"`

	// Phonetic name.
	NamePhonetic *string `mandatory:"false" json:"namePhonetic"`

	// TCA customer account number.
	TcaCustAccountNumber *string `mandatory:"false" json:"tcaCustAccountNumber"`

	// The business partner is part of the public sector or not.
	IsPublicSector *bool `mandatory:"false" json:"isPublicSector"`

	// The business partner is chain customer or not.
	IsChainCustomer *bool `mandatory:"false" json:"isChainCustomer"`

	// Customer chain type.
	CustomerChainType *string `mandatory:"false" json:"customerChainType"`

	// TCA party number.
	TcaPartyNumber *string `mandatory:"false" json:"tcaPartyNumber"`

	// TCA party ID.
	TcaPartyId *int64 `mandatory:"false" json:"tcaPartyId"`

	// TCA customer account ID.
	TcaCustomerAccountId *int64 `mandatory:"false" json:"tcaCustomerAccountId"`
}

func (m SubscribedServiceBusinessPartner) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SubscribedServiceBusinessPartner) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
