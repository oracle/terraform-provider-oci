// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DevOps API
//
// Use the DevOps API to create DevOps projects, configure code repositories,  add artifacts to deploy, build and test software applications, configure  target deployment environments, and deploy software applications.  For more information, see DevOps (https://docs.cloud.oracle.com/Content/devops/using/home.htm).
//

package devops

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v56/common"
)

// AbsoluteWaitCriteriaSummary Specifies the absolute wait criteria, user can specify fixed length of wait duration.
type AbsoluteWaitCriteriaSummary struct {

	// The absolute wait duration. Minimum wait duration must be 5 seconds. Maximum wait duration can be up to 2 days.
	WaitDuration *string `mandatory:"false" json:"waitDuration"`
}

func (m AbsoluteWaitCriteriaSummary) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m AbsoluteWaitCriteriaSummary) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeAbsoluteWaitCriteriaSummary AbsoluteWaitCriteriaSummary
	s := struct {
		DiscriminatorParam string `json:"waitType"`
		MarshalTypeAbsoluteWaitCriteriaSummary
	}{
		"ABSOLUTE_WAIT",
		(MarshalTypeAbsoluteWaitCriteriaSummary)(m),
	}

	return json.Marshal(&s)
}
