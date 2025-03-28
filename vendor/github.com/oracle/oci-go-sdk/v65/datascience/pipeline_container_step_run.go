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

// PipelineContainerStepRun Detail of each ContainerStepRun.
type PipelineContainerStepRun struct {

	// The date and time the pipeline step run was started in the timestamp format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeStarted *common.SDKTime `mandatory:"true" json:"timeStarted"`

	// The name of the step.
	StepName *string `mandatory:"true" json:"stepName"`

	// The date and time the pipeline step run finshed executing in the timestamp format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeFinished *common.SDKTime `mandatory:"false" json:"timeFinished"`

	// Details of the state of the step run.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The state of the step run.
	LifecycleState PipelineStepRunLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`
}

// GetTimeStarted returns TimeStarted
func (m PipelineContainerStepRun) GetTimeStarted() *common.SDKTime {
	return m.TimeStarted
}

// GetTimeFinished returns TimeFinished
func (m PipelineContainerStepRun) GetTimeFinished() *common.SDKTime {
	return m.TimeFinished
}

// GetStepName returns StepName
func (m PipelineContainerStepRun) GetStepName() *string {
	return m.StepName
}

// GetLifecycleState returns LifecycleState
func (m PipelineContainerStepRun) GetLifecycleState() PipelineStepRunLifecycleStateEnum {
	return m.LifecycleState
}

// GetLifecycleDetails returns LifecycleDetails
func (m PipelineContainerStepRun) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

func (m PipelineContainerStepRun) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PipelineContainerStepRun) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingPipelineStepRunLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetPipelineStepRunLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m PipelineContainerStepRun) MarshalJSON() (buff []byte, e error) {
	type MarshalTypePipelineContainerStepRun PipelineContainerStepRun
	s := struct {
		DiscriminatorParam string `json:"stepType"`
		MarshalTypePipelineContainerStepRun
	}{
		"CONTAINER",
		(MarshalTypePipelineContainerStepRun)(m),
	}

	return json.Marshal(&s)
}
