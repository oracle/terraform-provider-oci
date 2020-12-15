// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Advisor API
//
// APIs for managing Cloud Advisor. Cloud Advisor provides recommendations that help you maximize cost savings and improve the security posture of your tenancy.
//

package optimizer

import (
	"github.com/oracle/oci-go-sdk/v31/common"
)

// SupportedLevels Optional. The profile levels supported by a recommendation.
// For example, profile level values could be `Low`, `Medium`, and `High`.
// Not all recommendations support this field.
type SupportedLevels struct {

	// The list of supported levels.
	Items []SupportedLevel `mandatory:"false" json:"items"`
}

func (m SupportedLevels) String() string {
	return common.PointerString(m)
}
