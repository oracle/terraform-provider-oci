// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// PatchReleaseDateRangeSelectionDetails Patch Release Date Range Selection Details.
// If this option is selected to create rule, the latest patch from that patch release date range for the product and product version is used to calculate compliance.
type PatchReleaseDateRangeSelectionDetails struct {

	// Patch release date.
	TimeTo *common.SDKTime `mandatory:"true" json:"timeTo"`

	// Patch release date.
	TimeFrom *common.SDKTime `mandatory:"false" json:"timeFrom"`

	// Pattern to match and support suffix and prefix "*" pattern.
	Pattern *string `mandatory:"false" json:"pattern"`
}

func (m PatchReleaseDateRangeSelectionDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PatchReleaseDateRangeSelectionDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m PatchReleaseDateRangeSelectionDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypePatchReleaseDateRangeSelectionDetails PatchReleaseDateRangeSelectionDetails
	s := struct {
		DiscriminatorParam string `json:"selectionType"`
		MarshalTypePatchReleaseDateRangeSelectionDetails
	}{
		"PATCH_RELEASE_DATE_RANGE",
		(MarshalTypePatchReleaseDateRangeSelectionDetails)(m),
	}

	return json.Marshal(&s)
}
