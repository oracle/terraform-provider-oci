// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DatabaseInsightsDataObjectSummary Summary of a database insights data object.
type DatabaseInsightsDataObjectSummary struct {

	// Unique identifier of OPSI data object.
	Identifier *string `mandatory:"true" json:"identifier"`

	// User-friendly name of OPSI data object.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Description of OPSI data object.
	Description *string `mandatory:"false" json:"description"`

	// Name of the data object, which can be used in data object queries just like how view names are used in a query.
	Name *string `mandatory:"false" json:"name"`

	// Names of all the groups to which the data object belongs to.
	GroupNames []string `mandatory:"false" json:"groupNames"`
}

// GetIdentifier returns Identifier
func (m DatabaseInsightsDataObjectSummary) GetIdentifier() *string {
	return m.Identifier
}

// GetDisplayName returns DisplayName
func (m DatabaseInsightsDataObjectSummary) GetDisplayName() *string {
	return m.DisplayName
}

// GetDescription returns Description
func (m DatabaseInsightsDataObjectSummary) GetDescription() *string {
	return m.Description
}

// GetName returns Name
func (m DatabaseInsightsDataObjectSummary) GetName() *string {
	return m.Name
}

// GetGroupNames returns GroupNames
func (m DatabaseInsightsDataObjectSummary) GetGroupNames() []string {
	return m.GroupNames
}

func (m DatabaseInsightsDataObjectSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DatabaseInsightsDataObjectSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m DatabaseInsightsDataObjectSummary) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDatabaseInsightsDataObjectSummary DatabaseInsightsDataObjectSummary
	s := struct {
		DiscriminatorParam string `json:"dataObjectType"`
		MarshalTypeDatabaseInsightsDataObjectSummary
	}{
		"DATABASE_INSIGHTS_DATA_OBJECT",
		(MarshalTypeDatabaseInsightsDataObjectSummary)(m),
	}

	return json.Marshal(&s)
}
