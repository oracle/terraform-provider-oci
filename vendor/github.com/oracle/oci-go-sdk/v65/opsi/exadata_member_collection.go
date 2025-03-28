// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Ops Insights API
//
// Use the Ops Insights API to perform data extraction operations to obtain database
// resource utilization, performance statistics, and reference information. For more information,
// see About Oracle Cloud Infrastructure Ops Insights (https://docs.oracle.com/iaas/en-us/iaas/operations-insights/doc/operations-insights.html).
//

package opsi

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ExadataMemberCollection Partial definition of the exadata insight resource.
type ExadataMemberCollection struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Exadata insight.
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

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExadataMemberCollection) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingExadataTypeEnum(string(m.ExadataType)); !ok && m.ExadataType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ExadataType: %s. Supported values are: %s.", m.ExadataType, strings.Join(GetExadataTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingExadataRackTypeEnum(string(m.ExadataRackType)); !ok && m.ExadataRackType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ExadataRackType: %s. Supported values are: %s.", m.ExadataRackType, strings.Join(GetExadataRackTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
