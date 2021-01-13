// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Advisor API
//
// APIs for managing Cloud Advisor. Cloud Advisor provides recommendations that help you maximize cost savings and improve the security posture of your tenancy.
//

package optimizer

import (
	"github.com/oracle/oci-go-sdk/v32/common"
)

// Action Details about the recommended action.
// **Caution:** Avoid using any confidential information when you use the API to supply string values.
type Action struct {

	// The status of the resource action.
	Type ActionTypeEnum `mandatory:"true" json:"type"`

	// Text describing the recommended action.
	Description *string `mandatory:"true" json:"description"`

	// The URL path to documentation that explains how to perform the action.
	Url *string `mandatory:"true" json:"url"`
}

func (m Action) String() string {
	return common.PointerString(m)
}
