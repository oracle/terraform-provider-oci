// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Language API
//
// OCI Language Service solutions can help enterprise customers integrate AI into their products immediately using our proven,
// pre-trained and custom models or containers, without a need to set up an house team of AI and ML experts.
// This allows enterprises to focus on business drivers and development work rather than AI and ML operations, which shortens the time to market.
//

package ailanguage

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// JobSummary sub set of Job details data which need returns in list API
type JobSummary struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the job.
	Id *string `mandatory:"true" json:"id"`

	// A user-friendly display name for the job.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment where you want to create the job.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// A short description of the job.
	Description *string `mandatory:"false" json:"description"`

	// The current state of the Speech Job.
	LifecycleState JobLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// How much progress the operation has made, vs the total amount of work that must be performed.
	PercentComplete *int `mandatory:"false" json:"percentComplete"`

	// Total number of documents given as input for prediction. For CSV this signifies number of rows and for TXT this signifies number of files.
	TotalDocuments *int `mandatory:"false" json:"totalDocuments"`

	// Number of documents still to process. For CSV this signifies number of rows and for TXT this signifies number of files.
	PendingDocuments *int `mandatory:"false" json:"pendingDocuments"`

	// Number of documents processed for prediction. For CSV this signifies number of rows and for TXT this signifies number of files.
	CompletedDocuments *int `mandatory:"false" json:"completedDocuments"`

	// Number of documents failed for prediction. For CSV this signifies number of rows and for TXT this signifies number of files.
	FailedDocuments *int `mandatory:"false" json:"failedDocuments"`

	// warnings count
	WarningsCount *int `mandatory:"false" json:"warningsCount"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the user who created the job.
	CreatedBy *string `mandatory:"false" json:"createdBy"`

	// Job accepted time.
	TimeAccepted *common.SDKTime `mandatory:"false" json:"timeAccepted"`

	// Job started time.
	TimeStarted *common.SDKTime `mandatory:"false" json:"timeStarted"`

	// Job finished time.
	TimeCompleted *common.SDKTime `mandatory:"false" json:"timeCompleted"`
}

func (m JobSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m JobSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingJobLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetJobLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
