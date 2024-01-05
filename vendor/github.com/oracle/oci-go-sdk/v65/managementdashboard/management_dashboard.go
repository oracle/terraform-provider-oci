// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ManagementDashboard Properties of a dashboard, including dashboard ID.
type ManagementDashboard struct {

	// ID of the dashboard.  Same as id.
	DashboardId *string `mandatory:"true" json:"dashboardId"`

	// ID of the dashboard.  Same as dashboardId.
	Id *string `mandatory:"true" json:"id"`

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

	// User who created the dashboard.
	CreatedBy *string `mandatory:"true" json:"createdBy"`

	// Date and time the dashboard was created.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// User who updated the dashboard.
	UpdatedBy *string `mandatory:"true" json:"updatedBy"`

	// Date and time the dashboard was updated.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The version of the metadata defined in the API. This is maintained and enforced by dashboard server. Currently it is 2.0.
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

	// Array of saved searches in the dashboard.
	SavedSearches []ManagementSavedSearch `mandatory:"true" json:"savedSearches"`

	// State of dashboard.
	LifecycleState LifecycleStatesEnum `mandatory:"true" json:"lifecycleState"`

	// Defines parameters for the dashboard.
	ParametersConfig []interface{} `mandatory:"false" json:"parametersConfig"`

	// Drill-down configuration to define the destination of a drill-down action.
	DrilldownConfig []interface{} `mandatory:"false" json:"drilldownConfig"`

	// Contains configuration for enabling features.
	FeaturesConfig *interface{} `mandatory:"false" json:"featuresConfig"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m ManagementDashboard) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ManagementDashboard) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingLifecycleStatesEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetLifecycleStatesEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
