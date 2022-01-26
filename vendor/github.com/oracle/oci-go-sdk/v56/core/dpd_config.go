// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Core Services API
//
// Use the Core Services API to manage resources such as virtual cloud networks (VCNs),
// compute instances, and block storage volumes. For more information, see the console
// documentation for the Networking (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.cloud.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm), and
// Block Volume (https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/overview.htm) services.
//

package core

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// DpdConfig DPD Configuration Details
type DpdConfig struct {

	// dpd mode
	DpdMode DpdConfigDpdModeEnum `mandatory:"false" json:"dpdMode,omitempty"`

	// DPD Timeout in seconds.
	DpdTimeoutInSec *int `mandatory:"false" json:"dpdTimeoutInSec"`
}

func (m DpdConfig) String() string {
	return common.PointerString(m)
}

// DpdConfigDpdModeEnum Enum with underlying type: string
type DpdConfigDpdModeEnum string

// Set of constants representing the allowable values for DpdConfigDpdModeEnum
const (
	DpdConfigDpdModeInitiateAndRespond DpdConfigDpdModeEnum = "INITIATE_AND_RESPOND"
	DpdConfigDpdModeRespondOnly        DpdConfigDpdModeEnum = "RESPOND_ONLY"
)

var mappingDpdConfigDpdMode = map[string]DpdConfigDpdModeEnum{
	"INITIATE_AND_RESPOND": DpdConfigDpdModeInitiateAndRespond,
	"RESPOND_ONLY":         DpdConfigDpdModeRespondOnly,
}

// GetDpdConfigDpdModeEnumValues Enumerates the set of values for DpdConfigDpdModeEnum
func GetDpdConfigDpdModeEnumValues() []DpdConfigDpdModeEnum {
	values := make([]DpdConfigDpdModeEnum, 0)
	for _, v := range mappingDpdConfigDpdMode {
		values = append(values, v)
	}
	return values
}
