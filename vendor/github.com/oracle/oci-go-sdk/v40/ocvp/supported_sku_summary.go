// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud VMware Solution API
//
// Use this API to manage your Oracle Cloud VMware Solution (https://docs.cloud.oracle.com/iaas/Content/VMware/Concepts/ocvsoverview.htm).
//

package ocvp

import (
	"github.com/oracle/oci-go-sdk/v40/common"
)

// SupportedSkuSummary A specific SKU. Oracle Cloud Infrastructure VMware Solution supports the following billing interval SKUs:
// HOUR, MONTH, ONE_YEAR, and THREE_YEARS.
type SupportedSkuSummary struct {

	// name of SKU
	Name SkuEnum `mandatory:"true" json:"name"`
}

func (m SupportedSkuSummary) String() string {
	return common.PointerString(m)
}
