// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// LogAnalyticsEntitySummary Summary of a log analytics entity.
type LogAnalyticsEntitySummary struct {

	// The log analytics entity OCID. This ID is a reference used by log analytics features and it represents
	// a resource that is provisioned and managed by the customer on their premises or on the cloud.
	Id *string `mandatory:"true" json:"id"`

	// Log analytics entity name.
	Name *string `mandatory:"true" json:"name"`

	// Compartment Identifier OCID  (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Log analytics entity type name.
	EntityTypeName *string `mandatory:"true" json:"entityTypeName"`

	// Internal name for the log analytics entity type.
	EntityTypeInternalName *string `mandatory:"true" json:"entityTypeInternalName"`

	// The current state of the log analytics entity.
	LifecycleState EntityLifecycleStatesEnum `mandatory:"true" json:"lifecycleState"`

	// lifecycleDetails has additional information regarding substeps such as management agent plugin deployment.
	LifecycleDetails *string `mandatory:"true" json:"lifecycleDetails"`

	// The date and time the resource was created, in the format defined by RFC3339.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The date and time the resource was last updated, in the format defined by RFC3339.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The OCID of the Management Agent.
	ManagementAgentId *string `mandatory:"false" json:"managementAgentId"`

	// The OCID of the Cloud resource which this entity is a representation of. This may be blank when the entity
	// represents a non-cloud resource that the customer may have on their premises.
	CloudResourceId *string `mandatory:"false" json:"cloudResourceId"`

	// The timezone region of the log analytics entity.
	TimezoneRegion *string `mandatory:"false" json:"timezoneRegion"`

	// The Boolean flag to indicate if logs are collected for an entity for log analytics usage.
	AreLogsCollected *bool `mandatory:"false" json:"areLogsCollected"`

	// This indicates the type of source. It is primarily for Enterprise Manager Repository ID.
	SourceId *string `mandatory:"false" json:"sourceId"`

	CreationSource *CreationSource `mandatory:"false" json:"creationSource"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m LogAnalyticsEntitySummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m LogAnalyticsEntitySummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingEntityLifecycleStatesEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetEntityLifecycleStatesEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
