// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// UsageApi API
//
// A description of the UsageApi API.
//

package usage

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// CreateRedeemableUserDetails It contains a list of email Ids to be added to the redeemable users.
type CreateRedeemableUserDetails struct {

	// The list of email ids to be added to the redeemable users list.
	Items []RedeemableUser `mandatory:"false" json:"items"`
}

func (m CreateRedeemableUserDetails) String() string {
	return common.PointerString(m)
}
