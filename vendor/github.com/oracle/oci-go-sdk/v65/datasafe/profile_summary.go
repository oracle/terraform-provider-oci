// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ProfileSummary The summary of information about the user profiles. It includes details such as profile name, failed login attempts,
// sessions per user, inactive account time, password lock time, user created, target id, and the compartment id.
type ProfileSummary struct {

	// The OCID of the latest user assessment corresponding to the target under consideration. A compartment
	// type assessment can also be passed to profiles from all the targets from the corresponding compartment.
	UserAssessmentId *string `mandatory:"true" json:"userAssessmentId"`

	// The OCID of the compartment that contains the user assessment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID of the target database.
	TargetId *string `mandatory:"false" json:"targetId"`

	// The name of the profile.
	ProfileName *string `mandatory:"false" json:"profileName"`

	// The number of users having a given profile.
	UserCount *int `mandatory:"false" json:"userCount"`

	// Maximum times the user is allowed to fail login before the user account is locked.
	FailedLoginAttempts *string `mandatory:"false" json:"failedLoginAttempts"`

	// PL/SQL that can be used for password verification.
	PasswordVerificationFunction *string `mandatory:"false" json:"passwordVerificationFunction"`

	// The maximum number of sessions a user is allowed to create.
	SessionsPerUser *string `mandatory:"false" json:"sessionsPerUser"`

	// The permitted periods of continuous inactive time during a session, expressed in minutes.
	// Long-running queries and other operations are not subjected to this limit.
	InactiveAccountTime *string `mandatory:"false" json:"inactiveAccountTime"`

	// Number of days the user account remains locked after failed login
	PasswordLockTime *string `mandatory:"false" json:"passwordLockTime"`

	// Represents if the profile is created by user.
	IsUserCreated *bool `mandatory:"false" json:"isUserCreated"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m ProfileSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ProfileSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
