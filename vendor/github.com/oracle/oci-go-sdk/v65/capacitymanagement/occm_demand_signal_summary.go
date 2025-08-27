// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OCI Control Center Capacity Management API
//
// OCI Control Center (OCC) Capacity Management enables you to manage capacity requests in realms where OCI Control Center Capacity Management is available. For more information, see OCI Control Center (https://docs.oracle.com/iaas/Content/control-center/home.htm).
//

package capacitymanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// OccmDemandSignalSummary A summary model for the occm demand signal.
type OccmDemandSignalSummary struct {

	// The OCID of the demand signal.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the tenancy from which the request to create the demand signal was made.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The display name of the demand signal.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The different states associated with a demand signal.
	// CREATED -> A demand signal is by default created in this state.
	// SUBMITTED -> Once you have reviewed the details of the demand signal, you can transition it to SUBMITTED state so that OCI can start working on it.
	// DELETED -> You can delete a demand signal as long as it is in either CREATED or SUBMITTED state.
	// IN_PROGRESS -> Once OCI starts working on a given demand signal. They transition it to IN_PROGRESS.
	// REJECTED -> OCI can transition the demand signal to this state if all the demand signal items of that demand signal are declined.
	// COMPLETED -> OCI will transition the demand signal to COMPLETED state once the quantities which OCI committed to deliver to you has been delivered.
	LifecycleDetails OccmDemandSignalLifecycleDetailsEnum `mandatory:"true" json:"lifecycleDetails"`

	// The current lifecycle state of the demand signal.
	LifecycleState OccmDemandSignalLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The time when the demand signal was created.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time when the demand signal was last updated.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// A short description about the demand signal.
	Description *string `mandatory:"false" json:"description"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m OccmDemandSignalSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OccmDemandSignalSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingOccmDemandSignalLifecycleDetailsEnum(string(m.LifecycleDetails)); !ok && m.LifecycleDetails != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleDetails: %s. Supported values are: %s.", m.LifecycleDetails, strings.Join(GetOccmDemandSignalLifecycleDetailsEnumStringValues(), ",")))
	}
	if _, ok := GetMappingOccmDemandSignalLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetOccmDemandSignalLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
