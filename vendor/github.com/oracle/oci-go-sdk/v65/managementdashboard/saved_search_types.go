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
	"strings"
)

// SavedSearchTypesEnum Enum with underlying type: string
type SavedSearchTypesEnum string

// Set of constants representing the allowable values for SavedSearchTypesEnum
const (
	SavedSearchTypesSearchShowInDashboard     SavedSearchTypesEnum = "SEARCH_SHOW_IN_DASHBOARD"
	SavedSearchTypesSearchDontShowInDashboard SavedSearchTypesEnum = "SEARCH_DONT_SHOW_IN_DASHBOARD"
	SavedSearchTypesWidgetShowInDashboard     SavedSearchTypesEnum = "WIDGET_SHOW_IN_DASHBOARD"
	SavedSearchTypesWidgetDontShowInDashboard SavedSearchTypesEnum = "WIDGET_DONT_SHOW_IN_DASHBOARD"
	SavedSearchTypesFilterShowInDashboard     SavedSearchTypesEnum = "FILTER_SHOW_IN_DASHBOARD"
	SavedSearchTypesFilterDontShowInDashboard SavedSearchTypesEnum = "FILTER_DONT_SHOW_IN_DASHBOARD"
)

var mappingSavedSearchTypesEnum = map[string]SavedSearchTypesEnum{
	"SEARCH_SHOW_IN_DASHBOARD":      SavedSearchTypesSearchShowInDashboard,
	"SEARCH_DONT_SHOW_IN_DASHBOARD": SavedSearchTypesSearchDontShowInDashboard,
	"WIDGET_SHOW_IN_DASHBOARD":      SavedSearchTypesWidgetShowInDashboard,
	"WIDGET_DONT_SHOW_IN_DASHBOARD": SavedSearchTypesWidgetDontShowInDashboard,
	"FILTER_SHOW_IN_DASHBOARD":      SavedSearchTypesFilterShowInDashboard,
	"FILTER_DONT_SHOW_IN_DASHBOARD": SavedSearchTypesFilterDontShowInDashboard,
}

var mappingSavedSearchTypesEnumLowerCase = map[string]SavedSearchTypesEnum{
	"search_show_in_dashboard":      SavedSearchTypesSearchShowInDashboard,
	"search_dont_show_in_dashboard": SavedSearchTypesSearchDontShowInDashboard,
	"widget_show_in_dashboard":      SavedSearchTypesWidgetShowInDashboard,
	"widget_dont_show_in_dashboard": SavedSearchTypesWidgetDontShowInDashboard,
	"filter_show_in_dashboard":      SavedSearchTypesFilterShowInDashboard,
	"filter_dont_show_in_dashboard": SavedSearchTypesFilterDontShowInDashboard,
}

// GetSavedSearchTypesEnumValues Enumerates the set of values for SavedSearchTypesEnum
func GetSavedSearchTypesEnumValues() []SavedSearchTypesEnum {
	values := make([]SavedSearchTypesEnum, 0)
	for _, v := range mappingSavedSearchTypesEnum {
		values = append(values, v)
	}
	return values
}

// GetSavedSearchTypesEnumStringValues Enumerates the set of values in String for SavedSearchTypesEnum
func GetSavedSearchTypesEnumStringValues() []string {
	return []string{
		"SEARCH_SHOW_IN_DASHBOARD",
		"SEARCH_DONT_SHOW_IN_DASHBOARD",
		"WIDGET_SHOW_IN_DASHBOARD",
		"WIDGET_DONT_SHOW_IN_DASHBOARD",
		"FILTER_SHOW_IN_DASHBOARD",
		"FILTER_DONT_SHOW_IN_DASHBOARD",
	}
}

// GetMappingSavedSearchTypesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSavedSearchTypesEnum(val string) (SavedSearchTypesEnum, bool) {
	enum, ok := mappingSavedSearchTypesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
