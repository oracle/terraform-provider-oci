// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DevOps API
//
// Use the DevOps APIs to create a DevOps project to group the pipelines,  add reference to target deployment environments, add artifacts to deploy,  and create deployment pipelines needed to deploy your software.
//

package devops

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v51/common"
)

// CreateAbsoluteWaitCriteriaDetails Specifies the absolute wait criteria, user can specify fixed length of wait duration.
type CreateAbsoluteWaitCriteriaDetails struct {

	// The absolute wait duration.
	// Minimum waitDuration should be 5 seconds.
	// Maximum waitDuration can be up to 2 days.
	WaitDuration *string `mandatory:"true" json:"waitDuration"`
}

func (m CreateAbsoluteWaitCriteriaDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m CreateAbsoluteWaitCriteriaDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateAbsoluteWaitCriteriaDetails CreateAbsoluteWaitCriteriaDetails
	s := struct {
		DiscriminatorParam string `json:"waitType"`
		MarshalTypeCreateAbsoluteWaitCriteriaDetails
	}{
		"ABSOLUTE_WAIT",
		(MarshalTypeCreateAbsoluteWaitCriteriaDetails)(m),
	}

	return json.Marshal(&s)
}
