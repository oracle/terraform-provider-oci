// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
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
	"github.com/oracle/oci-go-sdk/v64/common"
	"strings"
)

// MlJobStepRun Detail of each MLJobStepRun
type MlJobStepRun struct {

	// The date and time the pipeline step run was started in the timestamp format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeStarted *common.SDKTime `mandatory:"true" json:"timeStarted"`

	// The name of the step being run
	StepName *string `mandatory:"true" json:"stepName"`

	// The date and time the pipeline step run finshed executing in the timestamp format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeFinished *common.SDKTime `mandatory:"false" json:"timeFinished"`

	// Details of the state of the step run.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the job run triggered for this step run
	JobRunId *string `mandatory:"false" json:"jobRunId"`

	// The state of the step run.
	LifecycleState StepRunLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`
}

//GetTimeStarted returns TimeStarted
func (m MlJobStepRun) GetTimeStarted() *common.SDKTime {
	return m.TimeStarted
}

//GetTimeFinished returns TimeFinished
func (m MlJobStepRun) GetTimeFinished() *common.SDKTime {
	return m.TimeFinished
}

//GetStepName returns StepName
func (m MlJobStepRun) GetStepName() *string {
	return m.StepName
}

//GetLifecycleState returns LifecycleState
func (m MlJobStepRun) GetLifecycleState() StepRunLifecycleStateEnum {
	return m.LifecycleState
}

//GetLifecycleDetails returns LifecycleDetails
func (m MlJobStepRun) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

func (m MlJobStepRun) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MlJobStepRun) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingStepRunLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetStepRunLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m MlJobStepRun) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeMlJobStepRun MlJobStepRun
	s := struct {
		DiscriminatorParam string `json:"stepType"`
		MarshalTypeMlJobStepRun
	}{
		"ML_JOB",
		(MarshalTypeMlJobStepRun)(m),
	}

	return json.Marshal(&s)
}
