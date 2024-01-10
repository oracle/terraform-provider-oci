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

	// It defines the visualization type of the widget saved search, the UI options of that visualization type, the binding of data to the visualization.
	UiConfig *interface{} `mandatory:"true" json:"uiConfig"`

	// It defines how data is fetched. A functional saved search needs a valid dataConfig. See examples on how it can be constructed for various data sources.
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

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ManagementDashboardTileDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingManagementDashboardTileDetailsStateEnum(string(m.State)); !ok && m.State != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for State: %s. Supported values are: %s.", m.State, strings.Join(GetManagementDashboardTileDetailsStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ManagementDashboardTileDetailsStateEnum Enum with underlying type: string
type ManagementDashboardTileDetailsStateEnum string

// Set of constants representing the allowable values for ManagementDashboardTileDetailsStateEnum
const (
	ManagementDashboardTileDetailsStateDeleted      ManagementDashboardTileDetailsStateEnum = "DELETED"
	ManagementDashboardTileDetailsStateUnauthorized ManagementDashboardTileDetailsStateEnum = "UNAUTHORIZED"
	ManagementDashboardTileDetailsStateDefault      ManagementDashboardTileDetailsStateEnum = "DEFAULT"
)

var mappingManagementDashboardTileDetailsStateEnum = map[string]ManagementDashboardTileDetailsStateEnum{
	"DELETED":      ManagementDashboardTileDetailsStateDeleted,
	"UNAUTHORIZED": ManagementDashboardTileDetailsStateUnauthorized,
	"DEFAULT":      ManagementDashboardTileDetailsStateDefault,
}

var mappingManagementDashboardTileDetailsStateEnumLowerCase = map[string]ManagementDashboardTileDetailsStateEnum{
	"deleted":      ManagementDashboardTileDetailsStateDeleted,
	"unauthorized": ManagementDashboardTileDetailsStateUnauthorized,
	"default":      ManagementDashboardTileDetailsStateDefault,
}

// GetManagementDashboardTileDetailsStateEnumValues Enumerates the set of values for ManagementDashboardTileDetailsStateEnum
func GetManagementDashboardTileDetailsStateEnumValues() []ManagementDashboardTileDetailsStateEnum {
	values := make([]ManagementDashboardTileDetailsStateEnum, 0)
	for _, v := range mappingManagementDashboardTileDetailsStateEnum {
		values = append(values, v)
	}
	return values
}

// GetManagementDashboardTileDetailsStateEnumStringValues Enumerates the set of values in String for ManagementDashboardTileDetailsStateEnum
func GetManagementDashboardTileDetailsStateEnumStringValues() []string {
	return []string{
		"DELETED",
		"UNAUTHORIZED",
		"DEFAULT",
	}
}

// GetMappingManagementDashboardTileDetailsStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingManagementDashboardTileDetailsStateEnum(val string) (ManagementDashboardTileDetailsStateEnum, bool) {
	enum, ok := mappingManagementDashboardTileDetailsStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
