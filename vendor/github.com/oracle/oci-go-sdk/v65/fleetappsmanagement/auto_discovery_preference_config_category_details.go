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

// AutoDiscoveryPreferenceConfigCategoryDetails Auto Discovery Preference
// Defines product-wise auto-discovery preferences for resources in a tenancy or compartment.
// This preference takes precedence over the global auto-discovery preference defined in onboarding.
type AutoDiscoveryPreferenceConfigCategoryDetails struct {

	// Enable or disable auto-discovery in the specified scope for the specified product.
	IsDisabled *bool `mandatory:"true" json:"isDisabled"`

	Product *ConfigAssociationDetails `mandatory:"true" json:"product"`

	ScheduleDetails *AutoDiscoveryScheduleDetails `mandatory:"false" json:"scheduleDetails"`
}

func (m AutoDiscoveryPreferenceConfigCategoryDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AutoDiscoveryPreferenceConfigCategoryDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m AutoDiscoveryPreferenceConfigCategoryDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeAutoDiscoveryPreferenceConfigCategoryDetails AutoDiscoveryPreferenceConfigCategoryDetails
	s := struct {
		DiscriminatorParam string `json:"configCategory"`
		MarshalTypeAutoDiscoveryPreferenceConfigCategoryDetails
	}{
		"AUTO_DISCOVERY_PREFERENCE",
		(MarshalTypeAutoDiscoveryPreferenceConfigCategoryDetails)(m),
	}

	return json.Marshal(&s)
}
