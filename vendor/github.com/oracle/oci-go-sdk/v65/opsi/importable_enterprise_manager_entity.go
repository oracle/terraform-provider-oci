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

	// Ops Insights internal representation of the resource type.
	OpsiEntityType *string `mandatory:"false" json:"opsiEntityType"`
}

func (m ImportableEnterpriseManagerEntity) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ImportableEnterpriseManagerEntity) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
