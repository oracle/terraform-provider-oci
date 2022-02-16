// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// ManagementDashboard API
//
// API for the Management Dashboard micro-service. Use this API for dashboard and saved search metadata preservation and to perform  tasks such as creating a dashboard, creating a saved search, and obtaining a list of dashboards and saved searches in a compartment.
//
//

package managementdashboard

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// CreateManagementDashboardDetails Properties of a dashboard.  ID of the dashboard must only be provided for Out-of-the-Box (OOB) dashboards.
type CreateManagementDashboardDetails struct {

	// ID of the service (for example, log-analytics) that owns the dashboard. Each service has a unique ID.
	ProviderId *string `mandatory:"true" json:"providerId"`

	// Name of the service (for example, Logging Analytics) that owns the dashboard.
	ProviderName *string `mandatory:"true" json:"providerName"`

	// Version of the service that owns the dashboard.
	ProviderVersion *string `mandatory:"true" json:"providerVersion"`

	// Array of dashboard tiles.
	Tiles []ManagementDashboardTileDetails `mandatory:"true" json:"tiles"`

	// Display name of the dashboard.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Description of the dashboard.
	Description *string `mandatory:"true" json:"description"`

	// OCID of the compartment in which the dashboard resides.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Determines whether the dashboard is an Out-of-the-Box (OOB) dashboard. Note that OOB dashboards are only provided by Oracle and cannot be modified.
	IsOobDashboard *bool `mandatory:"true" json:"isOobDashboard"`

	// Determines whether the dashboard will be displayed in Dashboard Home.
	IsShowInHome *bool `mandatory:"true" json:"isShowInHome"`

	// Version of the metadata.
	MetadataVersion *string `mandatory:"true" json:"metadataVersion"`

	// Determines whether the description of the dashboard is displayed.
	IsShowDescription *bool `mandatory:"true" json:"isShowDescription"`

	// Screen image of the dashboard.
	ScreenImage *string `mandatory:"true" json:"screenImage"`

	// JSON that contains internationalization options.
	Nls *interface{} `mandatory:"true" json:"nls"`

	// JSON that contains user interface options.
	UiConfig *interface{} `mandatory:"true" json:"uiConfig"`

	// Array of JSON that contain data source options.
	DataConfig []interface{} `mandatory:"true" json:"dataConfig"`

	// Type of dashboard. NORMAL denotes a single dashboard and SET denotes a dashboard set.
	Type *string `mandatory:"true" json:"type"`

	// Determines whether the dashboard is set as favorite.
	IsFavorite *bool `mandatory:"true" json:"isFavorite"`

	// ID of the dashboard, which must only be provided for Out-of-the-Box (OOB) dashboards.
	DashboardId *string `mandatory:"false" json:"dashboardId"`

	// Defines parameters for the dashboard.
	ParametersConfig []interface{} `mandatory:"false" json:"parametersConfig"`

	// Drill-down configuration to define the destination of a drill-down action.
	DrilldownConfig []interface{} `mandatory:"false" json:"drilldownConfig"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateManagementDashboardDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateManagementDashboardDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
