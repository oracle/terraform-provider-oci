// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// OceInstance API
//
// Oracle Content and Experience is a cloud-based content hub to drive omni-channel content management and accelerate experience delivery
//

package oce

import (
	"github.com/oracle/oci-go-sdk/common"
)

// DeleteOceInstanceDetails The information about the resource to be deleted.
type DeleteOceInstanceDetails struct {

	// IDCS access token identifying a stripe and service administrator user
	IdcsAccessToken *string `mandatory:"true" json:"idcsAccessToken"`
}

func (m DeleteOceInstanceDetails) String() string {
	return common.PointerString(m)
}
