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

// Profile The comprehensive information about the user profiles available on a given target.
// It includes details such as profile name, failed login attempts, password reuse time, password verification function,
// password verification function implementation code snippet, sessions per user, connect time inactive account time,
// password lock time, cpu usage per session, target id, and compartment id.
type Profile struct {

	// The OCID of the user assessment corresponding to the target under consideration.
	UserAssessmentId *string `mandatory:"true" json:"userAssessmentId"`

	// The name of the profile.
	ProfileName *string `mandatory:"true" json:"profileName"`

	// The OCID of the compartment that contains the user assessment.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// The OCID of the target database.
	TargetId *string `mandatory:"false" json:"targetId"`

	// The number of users that have a given profile.
	UserCount *int `mandatory:"false" json:"userCount"`

	// Maximum times the user is allowed in fail login before the user account is locked.
	FailedLoginAttempts *string `mandatory:"false" json:"failedLoginAttempts"`

	// Name of the PL/SQL that can be used for password verification.
	PasswordVerificationFunction *string `mandatory:"false" json:"passwordVerificationFunction"`

	// Details about the PL/SQL that can be used for password verification.
	PasswordVerificationFunctionDetails *string `mandatory:"false" json:"passwordVerificationFunctionDetails"`

	// Number of days the user account remains locked after failed login.
	PasswordLockTime *string `mandatory:"false" json:"passwordLockTime"`

	// Number of days the password is valid before expiry.
	PasswordLifeTime *string `mandatory:"false" json:"passwordLifeTime"`

	// Number of day after the user can use the already used password.
	PasswordReuseMax *string `mandatory:"false" json:"passwordReuseMax"`

	// Number of days before which a password cannot be reused.
	PasswordReuseTime *string `mandatory:"false" json:"passwordReuseTime"`

	// Number of days the password rollover is allowed. Minimum value can be 1/24 day (1 hour) to 60 days.
	PasswordRolloverTime *string `mandatory:"false" json:"passwordRolloverTime"`

	// Number of grace days for user to change password.
	PasswordGraceTime *string `mandatory:"false" json:"passwordGraceTime"`

	// Represents if the profile is created by user.
	IsUserCreated *bool `mandatory:"false" json:"isUserCreated"`

	// Specify the number of concurrent sessions to which you want to limit the user.
	SessionsPerUser *string `mandatory:"false" json:"sessionsPerUser"`

	// The permitted periods of continuous inactive time during a session, expressed in minutes.
	// Long-running queries and other operations are not subject to this limit.
	InactiveAccountTime *string `mandatory:"false" json:"inactiveAccountTime"`

	// Specify the total elapsed time limit for a session, expressed in minutes.
	ConnectTime *string `mandatory:"false" json:"connectTime"`

	// Specify the permitted periods of continuous inactive time during a  session, expressed in minutes.
	IdleTime *string `mandatory:"false" json:"idleTime"`

	// Specify the total resource cost for a session, expressed in service units. Oracle Database calculates the total
	// service units as a weighted sum of CPU_PER_SESSION, CONNECT_TIME, LOGICAL_READS_PER_SESSION, and PRIVATE_SGA.
	CompositeLimit *string `mandatory:"false" json:"compositeLimit"`

	// Specify the CPU time limit for a call (a parse, execute, or fetch), expressed in hundredths of seconds.
	CpuPerCall *string `mandatory:"false" json:"cpuPerCall"`

	// Specify the CPU time limit for a session, expressed in hundredth of seconds.
	CpuPerSession *string `mandatory:"false" json:"cpuPerSession"`

	// Specify the permitted the number of data blocks read for a call to process a SQL statement (a parse, execute, or fetch).
	LogicalReadsPerCall *string `mandatory:"false" json:"logicalReadsPerCall"`

	// Specify the permitted number of data blocks read in a session, including blocks read from memory and disk.
	LogicalReadsPerSession *string `mandatory:"false" json:"logicalReadsPerSession"`

	// Specify the amount of private space a session can allocate in the shared pool of the system global area (SGA),
	// expressed in bytes.
	PrivateSga *string `mandatory:"false" json:"privateSga"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m Profile) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Profile) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
