// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Audit API
//
// API for the Audit Service. Use this API for compliance monitoring in your tenancy.
// For more information, see Overview of Audit (https://docs.cloud.oracle.com/iaas/Content/Audit/Concepts/auditoverview.htm).
// **Tip**: This API is good for queries, but not bulk-export operations.
//

package audit

import (
	"github.com/oracle/oci-go-sdk/common"
)

// Configuration The retention period setting, specified in days. For more information, see Setting Audit
// Log Retention Period (https://docs.cloud.oracle.com/iaas/Content/Audit/Tasks/settingretentionperiod.htm).
type Configuration struct {

	// The retention period setting, specified in days. The minimum is 90, the maximum 365.
	// Example: `90`
	RetentionPeriodDays *int `mandatory:"false" json:"retentionPeriodDays"`
}

func (m Configuration) String() string {
	return common.PointerString(m)
}
