// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Analytics API
//
// Analytics API.
//

package analytics

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SetFeatureBundleDetails Input payload for the feature set of an Analytics instance.
type SetFeatureBundleDetails struct {

	// The feature set of an Analytics instance.
	FeatureBundle FeatureBundleEnum `mandatory:"false" json:"featureBundle,omitempty"`
}

func (m SetFeatureBundleDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SetFeatureBundleDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingFeatureBundleEnum(string(m.FeatureBundle)); !ok && m.FeatureBundle != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for FeatureBundle: %s. Supported values are: %s.", m.FeatureBundle, strings.Join(GetFeatureBundleEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
