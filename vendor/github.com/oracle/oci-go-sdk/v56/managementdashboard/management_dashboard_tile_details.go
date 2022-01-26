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

// ManagementDashboardTileDetails Properties of the dashboard tile representing a saved search.
// Tiles are laid out in a twelve column grid system with (0,0) at upper left corner.
type ManagementDashboardTileDetails struct {

	// Display name of the saved search.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// ID of the saved search.
	SavedSearchId *string `mandatory:"true" json:"savedSearchId"`

	// Tile's row number.
	Row *int `mandatory:"true" json:"row"`

	// Tile's column number.
	Column *int `mandatory:"true" json:"column"`

	// The number of rows the tile occupies.
	Height *int `mandatory:"true" json:"height"`

	// The number of columns the tile occupies.
	Width *int `mandatory:"true" json:"width"`

	// JSON that contains internationalization options.
	Nls *interface{} `mandatory:"true" json:"nls"`

	// JSON that contains user interface options.
	UiConfig *interface{} `mandatory:"true" json:"uiConfig"`

	// Array of JSON that contain data source options.
	DataConfig []interface{} `mandatory:"true" json:"dataConfig"`

	// Current state of the saved search.
	State ManagementDashboardTileDetailsStateEnum `mandatory:"true" json:"state"`

	// Drill-down configuration to define the destination of a drill-down action.
	DrilldownConfig *interface{} `mandatory:"true" json:"drilldownConfig"`

	// Specifies the saved search parameters values
	ParametersMap *interface{} `mandatory:"false" json:"parametersMap"`
}

func (m ManagementDashboardTileDetails) String() string {
	return common.PointerString(m)
}

// ManagementDashboardTileDetailsStateEnum Enum with underlying type: string
type ManagementDashboardTileDetailsStateEnum string

// Set of constants representing the allowable values for ManagementDashboardTileDetailsStateEnum
const (
	ManagementDashboardTileDetailsStateDeleted      ManagementDashboardTileDetailsStateEnum = "DELETED"
	ManagementDashboardTileDetailsStateUnauthorized ManagementDashboardTileDetailsStateEnum = "UNAUTHORIZED"
	ManagementDashboardTileDetailsStateDefault      ManagementDashboardTileDetailsStateEnum = "DEFAULT"
)

var mappingManagementDashboardTileDetailsState = map[string]ManagementDashboardTileDetailsStateEnum{
	"DELETED":      ManagementDashboardTileDetailsStateDeleted,
	"UNAUTHORIZED": ManagementDashboardTileDetailsStateUnauthorized,
	"DEFAULT":      ManagementDashboardTileDetailsStateDefault,
}

// GetManagementDashboardTileDetailsStateEnumValues Enumerates the set of values for ManagementDashboardTileDetailsStateEnum
func GetManagementDashboardTileDetailsStateEnumValues() []ManagementDashboardTileDetailsStateEnum {
	values := make([]ManagementDashboardTileDetailsStateEnum, 0)
	for _, v := range mappingManagementDashboardTileDetailsState {
		values = append(values, v)
	}
	return values
}
