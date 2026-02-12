// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fleet Application Management Service API
//
// Fleet Application Management provides a centralized platform to help you automate resource management tasks, validate patch compliance, and enhance operational efficiency across an enterprise.
//

package fleetappsmanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ReleaseDate Filters and adds patches based on release date range and latest flag.
type ReleaseDate struct {

	// End time when the patch group is not applicable to any patchGroup locked fleets.
	TimeApplicableTo *common.SDKTime `mandatory:"true" json:"timeApplicableTo"`

	// Include latest patch per product (LATEST) or all patches (ALL) within the date range. Default is LATEST.
	InclusionCriteria ReleaseDateInclusionCriteriaEnum `mandatory:"true" json:"inclusionCriteria"`

	// Start time when the patch group is applicable for the patchGroup locked fleets.
	TimeApplicableFrom *common.SDKTime `mandatory:"false" json:"timeApplicableFrom"`
}

func (m ReleaseDate) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ReleaseDate) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingReleaseDateInclusionCriteriaEnum(string(m.InclusionCriteria)); !ok && m.InclusionCriteria != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for InclusionCriteria: %s. Supported values are: %s.", m.InclusionCriteria, strings.Join(GetReleaseDateInclusionCriteriaEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ReleaseDateInclusionCriteriaEnum Enum with underlying type: string
type ReleaseDateInclusionCriteriaEnum string

// Set of constants representing the allowable values for ReleaseDateInclusionCriteriaEnum
const (
	ReleaseDateInclusionCriteriaLatest ReleaseDateInclusionCriteriaEnum = "LATEST"
	ReleaseDateInclusionCriteriaAll    ReleaseDateInclusionCriteriaEnum = "ALL"
)

var mappingReleaseDateInclusionCriteriaEnum = map[string]ReleaseDateInclusionCriteriaEnum{
	"LATEST": ReleaseDateInclusionCriteriaLatest,
	"ALL":    ReleaseDateInclusionCriteriaAll,
}

var mappingReleaseDateInclusionCriteriaEnumLowerCase = map[string]ReleaseDateInclusionCriteriaEnum{
	"latest": ReleaseDateInclusionCriteriaLatest,
	"all":    ReleaseDateInclusionCriteriaAll,
}

// GetReleaseDateInclusionCriteriaEnumValues Enumerates the set of values for ReleaseDateInclusionCriteriaEnum
func GetReleaseDateInclusionCriteriaEnumValues() []ReleaseDateInclusionCriteriaEnum {
	values := make([]ReleaseDateInclusionCriteriaEnum, 0)
	for _, v := range mappingReleaseDateInclusionCriteriaEnum {
		values = append(values, v)
	}
	return values
}

// GetReleaseDateInclusionCriteriaEnumStringValues Enumerates the set of values in String for ReleaseDateInclusionCriteriaEnum
func GetReleaseDateInclusionCriteriaEnumStringValues() []string {
	return []string{
		"LATEST",
		"ALL",
	}
}

// GetMappingReleaseDateInclusionCriteriaEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingReleaseDateInclusionCriteriaEnum(val string) (ReleaseDateInclusionCriteriaEnum, bool) {
	enum, ok := mappingReleaseDateInclusionCriteriaEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
