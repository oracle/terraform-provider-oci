// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Operations Insights API
//
// Use the Operations Insights API to perform data extraction operations to obtain database
// resource utilization, performance statistics, and reference information. For more information,
// see About Oracle Cloud Infrastructure Operations Insights (https://docs.cloud.oracle.com/en-us/iaas/operations-insights/doc/operations-insights.html).
//

package opsi

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// ImportableEnterpriseManagerEntity An Enterprise Manager entity that can be imported into Operations Insights.
type ImportableEnterpriseManagerEntity struct {

	// Enterprise Manager Unique Identifier
	EnterpriseManagerIdentifier *string `mandatory:"true" json:"enterpriseManagerIdentifier"`

	// Enterprise Manager Entity Name
	EnterpriseManagerEntityName *string `mandatory:"true" json:"enterpriseManagerEntityName"`

	// Enterprise Manager Entity Type
	EnterpriseManagerEntityType *string `mandatory:"true" json:"enterpriseManagerEntityType"`

	// Enterprise Manager Entity Unique Identifier
	EnterpriseManagerEntityIdentifier *string `mandatory:"true" json:"enterpriseManagerEntityIdentifier"`

	// Operations Insights internal representation of the resource type.
	OpsiEntityType *string `mandatory:"false" json:"opsiEntityType"`
}

func (m ImportableEnterpriseManagerEntity) String() string {
	return common.PointerString(m)
}
