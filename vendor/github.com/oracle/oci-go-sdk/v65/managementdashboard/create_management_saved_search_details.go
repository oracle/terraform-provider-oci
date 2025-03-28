// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// CreateManagementSavedSearchDetails Properties of a saved search.
type CreateManagementSavedSearchDetails struct {

	// Display name of the saved search.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// ID of the service (for example log-analytics) that owns the saved search. Each service has a unique ID.
	ProviderId *string `mandatory:"true" json:"providerId"`

	// The version of the metadata of the provider. This is useful for provider to version its features and metadata. Any newly created saved search (or dashboard) should use providerVersion 3.0.0.
	ProviderVersion *string `mandatory:"true" json:"providerVersion"`

	// The user friendly name of the service (for example, Logging Analytics) that owns the saved search.
	ProviderName *string `mandatory:"true" json:"providerName"`

	// OCID of the compartment in which the saved search resides.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Determines whether the saved search is an Out-of-the-Box (OOB) saved search. Note that OOB saved searches are only provided by Oracle and cannot be modified.
	IsOobSavedSearch *bool `mandatory:"true" json:"isOobSavedSearch"`

	// Description of the saved search.
	Description *string `mandatory:"true" json:"description"`

	// JSON that contains internationalization options.
	Nls *interface{} `mandatory:"true" json:"nls"`

	// Determines how the saved search is displayed in a dashboard.
	Type SavedSearchTypesEnum `mandatory:"true" json:"type"`

	// It defines the visualization type of the widget saved search, the UI options of that visualization type, the binding of data to the visualization.
	UiConfig *interface{} `mandatory:"true" json:"uiConfig"`

	// It defines how data is fetched. A functional saved search needs a valid dataConfig. See examples on how it can be constructed for various data sources.
	DataConfig []interface{} `mandatory:"true" json:"dataConfig"`

	// Screen image of the saved search.
	ScreenImage *string `mandatory:"true" json:"screenImage"`

	// The version of the metadata defined in the API. This is maintained and enforced by dashboard server. Currently it is 2.0.
	MetadataVersion *string `mandatory:"true" json:"metadataVersion"`

	// The UI template that the saved search uses to render itself.
	WidgetTemplate *string `mandatory:"true" json:"widgetTemplate"`

	// The View Model that the saved search uses to render itself.
	WidgetVM *string `mandatory:"true" json:"widgetVM"`

	// ID of the saved search, which must only be provided for Out-of-the-Box (OOB) saved search.
	Id *string `mandatory:"false" json:"id"`

	// Defines parameters for the saved search.
	ParametersConfig []interface{} `mandatory:"false" json:"parametersConfig"`

	// Contains configuration for enabling features.
	FeaturesConfig *interface{} `mandatory:"false" json:"featuresConfig"`

	// Drill-down configuration to define the destination of a drill-down action.
	DrilldownConfig []interface{} `mandatory:"false" json:"drilldownConfig"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateManagementSavedSearchDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateManagementSavedSearchDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSavedSearchTypesEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetSavedSearchTypesEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
