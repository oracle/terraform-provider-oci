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

// PatchReleaseDateSelectionDetails Days since patch release date details.
// If you set the policy rule by selecting this option and entering the number of days since release, Fleet Application Management calculates the number of days passed after the release date and reports the product as noncompliant if the days exceed the number provided in Days since release
type PatchReleaseDateSelectionDetails struct {

	// Days passed since patch release.
	DaysSinceRelease *int `mandatory:"true" json:"daysSinceRelease"`
}

func (m PatchReleaseDateSelectionDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PatchReleaseDateSelectionDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m PatchReleaseDateSelectionDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypePatchReleaseDateSelectionDetails PatchReleaseDateSelectionDetails
	s := struct {
		DiscriminatorParam string `json:"selectionType"`
		MarshalTypePatchReleaseDateSelectionDetails
	}{
		"PATCH_RELEASE_DATE",
		(MarshalTypePatchReleaseDateSelectionDetails)(m),
	}

	return json.Marshal(&s)
}
