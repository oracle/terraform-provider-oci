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
	"github.com/oracle/oci-go-sdk/v56/common"
)

// UpdateManagementSavedSearchDetails Properties of a saved search.  Saved search ID must not be provided.
type UpdateManagementSavedSearchDetails struct {

	// Display name of the saved search.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// ID of the service (for example log-analytics) that owns the saved search. Each service has a unique ID.
	ProviderId *string `mandatory:"false" json:"providerId"`

	// Version of the service that owns this saved search.
	ProviderVersion *string `mandatory:"false" json:"providerVersion"`

	// Name of the service (for example, Logging Analytics) that owns the saved search.
	ProviderName *string `mandatory:"false" json:"providerName"`

	// OCID of the compartment in which the saved search resides.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// Determines whether the saved search is an Out-of-the-Box (OOB) saved search. Note that OOB saved searches are only provided by Oracle and cannot be modified.
	IsOobSavedSearch *bool `mandatory:"false" json:"isOobSavedSearch"`

	// Description of the saved search.
	Description *string `mandatory:"false" json:"description"`

	// JSON that contains internationalization options.
	Nls *interface{} `mandatory:"false" json:"nls"`

	// Determines how the saved search is displayed in a dashboard.
	Type SavedSearchTypesEnum `mandatory:"false" json:"type,omitempty"`

	// JSON that contains user interface options.
	UiConfig *interface{} `mandatory:"false" json:"uiConfig"`

	// Array of JSON that contain data source options.
	DataConfig []interface{} `mandatory:"false" json:"dataConfig"`

	// Screen image of the saved search.
	ScreenImage *string `mandatory:"false" json:"screenImage"`

	// Version of the metadata.
	MetadataVersion *string `mandatory:"false" json:"metadataVersion"`

	// Reference to the HTML file of the widget.
	WidgetTemplate *string `mandatory:"false" json:"widgetTemplate"`

	// Reference to the view model of the widget.
	WidgetVM *string `mandatory:"false" json:"widgetVM"`

	// Defines parameters for the saved search.
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

func (m UpdateManagementSavedSearchDetails) String() string {
	return common.PointerString(m)
}
