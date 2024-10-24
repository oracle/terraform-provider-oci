// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fleet Application Management Service API
//
// Fleet Application Management provides a centralized platform to help you automate resource management tasks, validate patch compliance, and enhance operational efficiency across an enterprise.
//

package fleetappsmanagement

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// PatchLevelSelectionDetails Patch Level Details.
// If you set the policy rule by selecting this option and provide the Patch level as Latest, Fleet Application Management calculates the compliance status of the product in the following ways:
//   - Reports the patch process as compliant for the software identified or targets discovered at the Latest and Latest-1 version.
//   - Reports the patch process as noncompliant for the software identified or targets discovered at the Latest-2 version.
type PatchLevelSelectionDetails struct {

	// Patch Name.
	PatchLevel PatchLevelSelectionDetailsPatchLevelEnum `mandatory:"true" json:"patchLevel"`
}

func (m PatchLevelSelectionDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PatchLevelSelectionDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingPatchLevelSelectionDetailsPatchLevelEnum(string(m.PatchLevel)); !ok && m.PatchLevel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PatchLevel: %s. Supported values are: %s.", m.PatchLevel, strings.Join(GetPatchLevelSelectionDetailsPatchLevelEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m PatchLevelSelectionDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypePatchLevelSelectionDetails PatchLevelSelectionDetails
	s := struct {
		DiscriminatorParam string `json:"selectionType"`
		MarshalTypePatchLevelSelectionDetails
	}{
		"PATCH_LEVEL",
		(MarshalTypePatchLevelSelectionDetails)(m),
	}

	return json.Marshal(&s)
}

// PatchLevelSelectionDetailsPatchLevelEnum Enum with underlying type: string
type PatchLevelSelectionDetailsPatchLevelEnum string

// Set of constants representing the allowable values for PatchLevelSelectionDetailsPatchLevelEnum
const (
	PatchLevelSelectionDetailsPatchLevelLatest         PatchLevelSelectionDetailsPatchLevelEnum = "LATEST"
	PatchLevelSelectionDetailsPatchLevelLatestMinusOne PatchLevelSelectionDetailsPatchLevelEnum = "LATEST_MINUS_ONE"
	PatchLevelSelectionDetailsPatchLevelLatestMinusTwo PatchLevelSelectionDetailsPatchLevelEnum = "LATEST_MINUS_TWO"
)

var mappingPatchLevelSelectionDetailsPatchLevelEnum = map[string]PatchLevelSelectionDetailsPatchLevelEnum{
	"LATEST":           PatchLevelSelectionDetailsPatchLevelLatest,
	"LATEST_MINUS_ONE": PatchLevelSelectionDetailsPatchLevelLatestMinusOne,
	"LATEST_MINUS_TWO": PatchLevelSelectionDetailsPatchLevelLatestMinusTwo,
}

var mappingPatchLevelSelectionDetailsPatchLevelEnumLowerCase = map[string]PatchLevelSelectionDetailsPatchLevelEnum{
	"latest":           PatchLevelSelectionDetailsPatchLevelLatest,
	"latest_minus_one": PatchLevelSelectionDetailsPatchLevelLatestMinusOne,
	"latest_minus_two": PatchLevelSelectionDetailsPatchLevelLatestMinusTwo,
}

// GetPatchLevelSelectionDetailsPatchLevelEnumValues Enumerates the set of values for PatchLevelSelectionDetailsPatchLevelEnum
func GetPatchLevelSelectionDetailsPatchLevelEnumValues() []PatchLevelSelectionDetailsPatchLevelEnum {
	values := make([]PatchLevelSelectionDetailsPatchLevelEnum, 0)
	for _, v := range mappingPatchLevelSelectionDetailsPatchLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetPatchLevelSelectionDetailsPatchLevelEnumStringValues Enumerates the set of values in String for PatchLevelSelectionDetailsPatchLevelEnum
func GetPatchLevelSelectionDetailsPatchLevelEnumStringValues() []string {
	return []string{
		"LATEST",
		"LATEST_MINUS_ONE",
		"LATEST_MINUS_TWO",
	}
}

// GetMappingPatchLevelSelectionDetailsPatchLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPatchLevelSelectionDetailsPatchLevelEnum(val string) (PatchLevelSelectionDetailsPatchLevelEnum, bool) {
	enum, ok := mappingPatchLevelSelectionDetailsPatchLevelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
