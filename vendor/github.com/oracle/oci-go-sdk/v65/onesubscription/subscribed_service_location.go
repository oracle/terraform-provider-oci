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

// SubscribedServiceLocation Address location.
type SubscribedServiceLocation struct {

	// Address first line.
	Address1 *string `mandatory:"false" json:"address1"`

	// Address second line.
	Address2 *string `mandatory:"false" json:"address2"`

	// Postal code.
	PostalCode *string `mandatory:"false" json:"postalCode"`

	// City.
	City *string `mandatory:"false" json:"city"`

	// Country.
	Country *string `mandatory:"false" json:"country"`

	// Region.
	Region *string `mandatory:"false" json:"region"`

	// Region.
	TcaLocationId *int64 `mandatory:"false" json:"tcaLocationId"`
}

func (m SubscribedServiceLocation) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SubscribedServiceLocation) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
