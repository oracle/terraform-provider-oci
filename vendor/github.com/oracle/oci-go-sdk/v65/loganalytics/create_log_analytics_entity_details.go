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

// CreateLogAnalyticsEntityDetails Details for new log analytics entity to be added.
type CreateLogAnalyticsEntityDetails struct {

	// Log analytics entity name.
	Name *string `mandatory:"true" json:"name"`

	// Compartment Identifier OCID  (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Log analytics entity type name.
	EntityTypeName *string `mandatory:"true" json:"entityTypeName"`

	// The OCID of the Management Agent.
	ManagementAgentId *string `mandatory:"false" json:"managementAgentId"`

	// The OCID of the Cloud resource which this entity is a representation of. This may be blank when the entity
	// represents a non-cloud resource that the customer may have on their premises.
	CloudResourceId *string `mandatory:"false" json:"cloudResourceId"`

	// The timezone region of the log analytics entity.
	TimezoneRegion *string `mandatory:"false" json:"timezoneRegion"`

	// The hostname where the entity represented here is actually present. This would be the output one would get if
	// they run `echo $HOSTNAME` on Linux or an equivalent OS command. This may be different from
	// management agents host since logs may be collected remotely.
	Hostname *string `mandatory:"false" json:"hostname"`

	// This indicates the type of source. It is primarily for Enterprise Manager Repository ID.
	SourceId *string `mandatory:"false" json:"sourceId"`

	// The name/value pairs for parameter values to be used in file patterns specified in log sources.
	Properties map[string]string `mandatory:"false" json:"properties"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The date and time the resource was last discovered, in the format defined by RFC3339.
	TimeLastDiscovered *common.SDKTime `mandatory:"false" json:"timeLastDiscovered"`

	Metadata *LogAnalyticsMetadataDetails `mandatory:"false" json:"metadata"`
}

func (m CreateLogAnalyticsEntityDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateLogAnalyticsEntityDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
