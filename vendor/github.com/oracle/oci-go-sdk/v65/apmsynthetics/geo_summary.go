// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// APM Availability Monitoring API
//
// Use the APM Availability Monitoring API to query Scripts, Monitors, Dedicated Vantage Points and On-Premise Vantage Points resources. For more information, see Application Performance Monitoring (https://docs.oracle.com/iaas/application-performance-monitoring/index.html).
//

package apmsynthetics

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// GeoSummary Geographic summary of a vantage point.
type GeoSummary struct {

	// The ISO 3166-2 code for this location's first-level administrative division, either a US state or Canadian province.
	// Only included for locations in the US or Canada. For a list of codes, see Country Codes.
	AdminDivCode *string `mandatory:"false" json:"adminDivCode"`

	// Common English-language name for the city.
	CityName *string `mandatory:"false" json:"cityName"`

	// The ISO 3166-1 alpha-2 country code. For a list of codes, see Country Codes.
	CountryCode *string `mandatory:"false" json:"countryCode"`

	// The common English-language name for the country.
	CountryName *string `mandatory:"false" json:"countryName"`

	// Degrees north of the equator.
	Latitude *float64 `mandatory:"false" json:"latitude"`

	// Degrees east of the prime meridian.
	Longitude *float64 `mandatory:"false" json:"longitude"`
}

func (m GeoSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m GeoSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
