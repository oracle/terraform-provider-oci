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

// InvokeMlApplicationProviderTriggerScheduleActionDetails invoke ml application trigger details
type InvokeMlApplicationProviderTriggerScheduleActionDetails struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the schedule.
	MlApplicationInstanceViewId *string `mandatory:"true" json:"mlApplicationInstanceViewId"`

	TriggerMlApplicationInstanceViewFlowDetails *TriggerMlApplicationInstanceViewFlowDetails `mandatory:"true" json:"triggerMlApplicationInstanceViewFlowDetails"`
}

func (m InvokeMlApplicationProviderTriggerScheduleActionDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m InvokeMlApplicationProviderTriggerScheduleActionDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m InvokeMlApplicationProviderTriggerScheduleActionDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeInvokeMlApplicationProviderTriggerScheduleActionDetails InvokeMlApplicationProviderTriggerScheduleActionDetails
	s := struct {
		DiscriminatorParam string `json:"httpActionType"`
		MarshalTypeInvokeMlApplicationProviderTriggerScheduleActionDetails
	}{
		"INVOKE_ML_APPLICATION_PROVIDER_TRIGGER",
		(MarshalTypeInvokeMlApplicationProviderTriggerScheduleActionDetails)(m),
	}

	return json.Marshal(&s)
}
