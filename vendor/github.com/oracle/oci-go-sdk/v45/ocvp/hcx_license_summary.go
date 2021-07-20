// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud VMware Solution API
//
// Use this API to manage your Oracle Cloud VMware Solution (https://docs.cloud.oracle.com/iaas/Content/VMware/Concepts/ocvsoverview.htm).
//

package ocvp

import (
	"github.com/oracle/oci-go-sdk/v45/common"
)

// HcxLicenseSummary HCX on-premise license information summary.
type HcxLicenseSummary struct {

	// HCX on-premise license key value
	ActivationKey *string `mandatory:"true" json:"activationKey"`

	// status of HCX on-premise license
	Status HcxLicenseStatusEnum `mandatory:"true" json:"status"`

	// Name of the system that consumed the HCX on-premise license
	SystemName *string `mandatory:"false" json:"systemName"`
}

func (m HcxLicenseSummary) String() string {
	return common.PointerString(m)
}
