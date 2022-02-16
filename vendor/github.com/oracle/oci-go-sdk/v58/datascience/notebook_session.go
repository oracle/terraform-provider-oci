// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Science API
//
// Use the Data Science API to organize your data science work, access data and computing resources, and build, train, deploy and manage models and model deployments. For more information, see Data Science (https://docs.oracle.com/iaas/data-science/using/data-science.htm).
//

package datascience

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// NotebookSession Notebook sessions are interactive coding environments for data scientists.
type NotebookSession struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the notebook session.
	Id *string `mandatory:"true" json:"id"`

	// The date and time the resource was created in the timestamp format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: 2019-08-25T21:10:29.41Z
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// A user-friendly display name for the resource. It does not have to be unique and can be modified. Avoid entering confidential information.
	// Example: `My NotebookSession`
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the project associated with the notebook session.
	ProjectId *string `mandatory:"true" json:"projectId"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the user who created the notebook session.
	CreatedBy *string `mandatory:"true" json:"createdBy"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the notebook session's compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The state of the notebook session.
	LifecycleState NotebookSessionLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	NotebookSessionConfigurationDetails *NotebookSessionConfigurationDetails `mandatory:"false" json:"notebookSessionConfigurationDetails"`

	// The URL to interact with the notebook session.
	NotebookSessionUrl *string `mandatory:"false" json:"notebookSessionUrl"`

	// Details about the state of the notebook session.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m NotebookSession) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m NotebookSession) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingNotebookSessionLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetNotebookSessionLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
