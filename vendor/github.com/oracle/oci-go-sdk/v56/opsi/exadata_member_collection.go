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

// ExadataMemberCollection Partial definition of the exadata insight resource.
type ExadataMemberCollection struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Exadata insight.
	ExadataInsightId *string `mandatory:"true" json:"exadataInsightId"`

	// The Exadata system name. If the Exadata systems managed by Enterprise Manager, the name is unique amongst the Exadata systems managed by the same Enterprise Manager.
	ExadataName *string `mandatory:"true" json:"exadataName"`

	// The user-friendly name for the Exadata system. The name does not have to be unique.
	ExadataDisplayName *string `mandatory:"true" json:"exadataDisplayName"`

	// Operations Insights internal representation of the the Exadata system type.
	ExadataType ExadataTypeEnum `mandatory:"true" json:"exadataType"`

	// Exadata rack type.
	ExadataRackType ExadataRackTypeEnum `mandatory:"true" json:"exadataRackType"`

	// Collection of Exadata members
	Items []ExadataMemberSummary `mandatory:"true" json:"items"`
}

func (m ExadataMemberCollection) String() string {
	return common.PointerString(m)
}
