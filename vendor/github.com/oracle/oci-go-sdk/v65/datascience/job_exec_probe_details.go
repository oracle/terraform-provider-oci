// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Science API
//
// Use the Data Science API to organize your data science work, access data and computing resources, and build, train, deploy and manage models and model deployments. For more information, see Data Science (https://docs.oracle.com/iaas/data-science/using/data-science.htm).
//

package datascience

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// JobExecProbeDetails Runs a command in the job run to check whether application is healthy or not.
type JobExecProbeDetails struct {

	// The commands to run in the target job run to perform the startup probe
	Command []string `mandatory:"true" json:"command"`

	// Number of seconds how often the job run should perform a startup probe
	PeriodInSeconds *int `mandatory:"false" json:"periodInSeconds"`

	// How many times the job will try before giving up when a probe fails.
	FailureThreshold *int `mandatory:"false" json:"failureThreshold"`

	// Number of seconds after the job run has started before a startup probe is initiated.
	InitialDelayInSeconds *int `mandatory:"false" json:"initialDelayInSeconds"`
}

func (m JobExecProbeDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m JobExecProbeDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m JobExecProbeDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeJobExecProbeDetails JobExecProbeDetails
	s := struct {
		DiscriminatorParam string `json:"jobProbeCheckType"`
		MarshalTypeJobExecProbeDetails
	}{
		"EXEC",
		(MarshalTypeJobExecProbeDetails)(m),
	}

	return json.Marshal(&s)
}
