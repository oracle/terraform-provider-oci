// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Health Checks API
//
// API for the Health Checks service. Use this API to manage endpoint probes and monitors.
// For more information, see
// Overview of the Health Checks Service (https://docs.oracle.com/iaas/Content/HealthChecks/Concepts/healthchecks.htm).
//

package healthchecks

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Geolocation Geographic information about a vantage point.
type Geolocation struct {

	// An opaque identifier for the geographic location of the vantage point.
	GeoKey *string `mandatory:"false" json:"geoKey"`

	// The ISO 3166-2 code for this location's first-level administrative
	// division, either a US state or Canadian province. Only included for locations
	// in the US or Canada. For a list of codes, see
	// Country Codes (https://www.iso.org/obp/ui/#search).
	AdminDivCode *string `mandatory:"false" json:"adminDivCode"`

	// Common English-language name for the city.
	CityName *string `mandatory:"false" json:"cityName"`

	// The ISO 3166-1 alpha-2 country code. For a list of codes,
	// see Country Codes (https://www.iso.org/obp/ui/#search).
	CountryCode *string `mandatory:"false" json:"countryCode"`

	// The common English-language name for the country.
	CountryName *string `mandatory:"false" json:"countryName"`

	// Degrees north of the Equator.
	Latitude *float32 `mandatory:"false" json:"latitude"`

	// Degrees east of the prime meridian.
	Longitude *float32 `mandatory:"false" json:"longitude"`
}

func (m Geolocation) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Geolocation) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
