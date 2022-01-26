// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// ProfileDetails The details of a particular profile
type ProfileDetails struct {

	// The number of users using this profile.
	NumUsers *int `mandatory:"false" json:"numUsers"`

	// The value of the CONNECT_TIME resource parameter.
	ConnectTime *string `mandatory:"false" json:"connectTime"`

	// The value of the FAILED_LOGIN_ATTEMPTS password parameter.
	FailedLoginAttempts *string `mandatory:"false" json:"failedLoginAttempts"`

	// The value of the IDLE_TIME resource parameter.
	IdleTime *string `mandatory:"false" json:"idleTime"`

	// The value of the INACTIVE_ACCOUNT_TIME password parameter.
	InactiveAccountTime *string `mandatory:"false" json:"inactiveAccountTime"`

	// The value of the PASSWORD_GRACE_TIME password parameter.
	PasswordGraceTime *string `mandatory:"false" json:"passwordGraceTime"`

	// The value of the PASSWORD_LIFE_TIME password parameter.
	PasswordLifeTime *string `mandatory:"false" json:"passwordLifeTime"`

	// The value of the PASSWORD_LOCK_TIME password parameter.
	PasswordLockTime *string `mandatory:"false" json:"passwordLockTime"`

	// The value of the PASSWORD_REUSE_TIME password parameter.
	PasswordReuseTime *string `mandatory:"false" json:"passwordReuseTime"`

	// The value of the PASSWORD_REUSE_MAX resource parameter.
	PasswordReuseMax *string `mandatory:"false" json:"passwordReuseMax"`

	// The value of the PASSWORD_VERIFY_FUNCTION resource.
	PasswordVerifyFunction *string `mandatory:"false" json:"passwordVerifyFunction"`
}

func (m ProfileDetails) String() string {
	return common.PointerString(m)
}
