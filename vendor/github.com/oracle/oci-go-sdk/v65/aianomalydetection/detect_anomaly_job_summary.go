// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Anomaly Detection API
//
// OCI AI Service solutions can help Enterprise customers integrate AI into their products immediately by using our proven,
// pre-trained/custom models or containers, and without a need to set up in house team of AI and ML experts.
// This allows enterprises to focus on business drivers and development work rather than AI/ML operations, shortening the time to market.
//

package aianomalydetection

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DetectAnomalyJobSummary Anomaly Job summary contains minimal information for asynchronous inference of anomalies
// returned in list response.
type DetectAnomalyJobSummary struct {

	// Id of the job.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the compartment that starts the job.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID of the trained model.
	ModelId *string `mandatory:"true" json:"modelId"`

	// Job accepted time
	TimeAccepted *common.SDKTime `mandatory:"true" json:"timeAccepted"`

	// The current state of the batch document job.
	LifecycleState DetectAnomalyJobLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Detect anomaly job display name.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Detect anomaly job description.
	Description *string `mandatory:"false" json:"description"`

	// The OCID of the project.
	ProjectId *string `mandatory:"false" json:"projectId"`

	// Job started time
	TimeStarted *common.SDKTime `mandatory:"false" json:"timeStarted"`

	// Job finished time
	TimeFinished *common.SDKTime `mandatory:"false" json:"timeFinished"`

	// The current state details of the batch document job.
	LifecycleStateDetails *string `mandatory:"false" json:"lifecycleStateDetails"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{ "orcl-cloud": { "free-tier-retained": "true" } }`
	SystemTags map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m DetectAnomalyJobSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DetectAnomalyJobSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDetectAnomalyJobLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetDetectAnomalyJobLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
