// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// CreateOnboardingDetails The information about enabling onboarding.
type CreateOnboardingDetails struct {

	// Tenancy OCID
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// A value determining if the Fleet Application Management tagging is enabled or not.
	// Allow Fleet Application Management to tag resources with fleet name using "Oracle$FAMS-Tags.FleetName" tag.
	IsFamsTagEnabled *bool `mandatory:"false" json:"isFamsTagEnabled"`

	// A value determining if the cost tracking tag is enabled or not.
	// Allow Fleet Application Management to tag resources with cost tracking tag using "Oracle$FAMS-Tags.FAMSManaged" tag.
	IsCostTrackingTagEnabled *bool `mandatory:"false" json:"isCostTrackingTagEnabled"`
}

func (m CreateOnboardingDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateOnboardingDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
