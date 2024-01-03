// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Science API
//
// Use the Data Science API to organize your data science work, access data and computing resources, and build, train, deploy and manage models and model deployments. For more information, see Data Science (https://docs.oracle.com/iaas/data-science/using/data-science.htm).
//

package datascience

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// JobStartupProbe The probe indicates whether the application/code within the Job container is started
type JobStartupProbe struct {

	// The commands to run in the target container to perform the startup probe
	Command []string `mandatory:"false" json:"command"`

	// Number of seconds how often the container should perform a startup probe
	PeriodInSeconds *int `mandatory:"false" json:"periodInSeconds"`

	// How many times the Job will try before giving up when a probe fails.
	FailureThreshold *int `mandatory:"false" json:"failureThreshold"`

	// Number of seconds after the container has started before a startup probe is initiated.
	InitialDelayInSeconds *int `mandatory:"false" json:"initialDelayInSeconds"`
}

func (m JobStartupProbe) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m JobStartupProbe) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
