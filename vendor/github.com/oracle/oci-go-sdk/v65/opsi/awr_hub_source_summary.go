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

// AwrHubSourceSummary Awr hub source object
type AwrHubSourceSummary struct {

	// The name of the Awr Hub source database.
	Name *string `mandatory:"true" json:"name"`

	// AWR Hub OCID
	AwrHubId *string `mandatory:"true" json:"awrHubId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// source type of the database
	Type AwrHubSourceTypeEnum `mandatory:"true" json:"type"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Awr Hub source database.
	Id *string `mandatory:"true" json:"id"`

	// The shorted string of the Awr Hub source database identifier.
	AwrHubOpsiSourceId *string `mandatory:"true" json:"awrHubOpsiSourceId"`

	// Opsi Mailbox URL based on the Awr Hub and Awr Hub source.
	SourceMailBoxUrl *string `mandatory:"true" json:"sourceMailBoxUrl"`

	// The time at which the resource was first created. An RFC3339 formatted datetime string
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// the current state of the source database
	LifecycleState AwrHubSourceLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Indicates the status of a source database in Operations Insights
	Status AwrHubSourceStatusEnum `mandatory:"true" json:"status"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the database id.
	AssociatedResourceId *string `mandatory:"false" json:"associatedResourceId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the database id.
	AssociatedOpsiId *string `mandatory:"false" json:"associatedOpsiId"`

	// The time at which the resource was last updated. An RFC3339 formatted datetime string
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	// This is `true` if the source databse is registered with a Awr Hub, otherwise `false`
	IsRegisteredWithAwrHub *bool `mandatory:"false" json:"isRegisteredWithAwrHub"`

	// DatabaseId of the Source database for which AWR Data will be uploaded to AWR Hub.
	AwrSourceDatabaseId *string `mandatory:"false" json:"awrSourceDatabaseId"`

	// The minimum snapshot identifier of the source database for which AWR data is uploaded to AWR Hub.
	MinSnapshotIdentifier *float32 `mandatory:"false" json:"minSnapshotIdentifier"`

	// The maximum snapshot identifier of the source database for which AWR data is uploaded to AWR Hub.
	MaxSnapshotIdentifier *float32 `mandatory:"false" json:"maxSnapshotIdentifier"`

	// The time at which the earliest snapshot was generated in the source database for which data is uploaded to AWR Hub. An RFC3339 formatted datetime string
	TimeFirstSnapshotGenerated *common.SDKTime `mandatory:"false" json:"timeFirstSnapshotGenerated"`

	// The time at which the latest snapshot was generated in the source database for which data is uploaded to AWR Hub. An RFC3339 formatted datetime string
	TimeLastSnapshotGenerated *common.SDKTime `mandatory:"false" json:"timeLastSnapshotGenerated"`

	// Number of hours since last AWR snapshots import happened from the Source database.
	HoursSinceLastImport *float64 `mandatory:"false" json:"hoursSinceLastImport"`
}

func (m AwrHubSourceSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AwrHubSourceSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAwrHubSourceTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetAwrHubSourceTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAwrHubSourceLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetAwrHubSourceLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAwrHubSourceStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetAwrHubSourceStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
