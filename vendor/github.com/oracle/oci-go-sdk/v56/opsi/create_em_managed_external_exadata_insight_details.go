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
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v56/common"
)

// CreateEmManagedExternalExadataInsightDetails The information about the Exadata system to be analyzed. If memberEntityDetails is not specified, the the Enterprise Manager entity (e.g. databases and hosts) associated with an Exadata system will be placed in the same compartment as the Exadata system.
type CreateEmManagedExternalExadataInsightDetails struct {

	// Compartment Identifier of Exadata insight
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Enterprise Manager Unique Identifier
	EnterpriseManagerIdentifier *string `mandatory:"true" json:"enterpriseManagerIdentifier"`

	// OPSI Enterprise Manager Bridge OCID
	EnterpriseManagerBridgeId *string `mandatory:"true" json:"enterpriseManagerBridgeId"`

	// Enterprise Manager Entity Unique Identifier
	EnterpriseManagerEntityIdentifier *string `mandatory:"true" json:"enterpriseManagerEntityIdentifier"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	MemberEntityDetails []CreateEmManagedExternalExadataMemberEntityDetails `mandatory:"false" json:"memberEntityDetails"`

	// Set to true to enable automatic enablement and disablement of related targets from Enterprise Manager. New resources (e.g. Database Insights) will be placed in the same compartment as the related Exadata Insight.
	IsAutoSyncEnabled *bool `mandatory:"false" json:"isAutoSyncEnabled"`
}

//GetCompartmentId returns CompartmentId
func (m CreateEmManagedExternalExadataInsightDetails) GetCompartmentId() *string {
	return m.CompartmentId
}

//GetFreeformTags returns FreeformTags
func (m CreateEmManagedExternalExadataInsightDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m CreateEmManagedExternalExadataInsightDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m CreateEmManagedExternalExadataInsightDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m CreateEmManagedExternalExadataInsightDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateEmManagedExternalExadataInsightDetails CreateEmManagedExternalExadataInsightDetails
	s := struct {
		DiscriminatorParam string `json:"entitySource"`
		MarshalTypeCreateEmManagedExternalExadataInsightDetails
	}{
		"EM_MANAGED_EXTERNAL_EXADATA",
		(MarshalTypeCreateEmManagedExternalExadataInsightDetails)(m),
	}

	return json.Marshal(&s)
}
