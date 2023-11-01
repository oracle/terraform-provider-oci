// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SecurityPolicyEntryState The resource represents the state of a specific entry type deployment on a target.
type SecurityPolicyEntryState struct {

	// Unique id of the security policy entry state.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the security policy entry type associated.
	SecurityPolicyEntryId *string `mandatory:"true" json:"securityPolicyEntryId"`

	// The current deployment status of the security policy deployment and the security policy entry associated.
	DeploymentStatus SecurityPolicyEntryStateDeploymentStatusEnum `mandatory:"true" json:"deploymentStatus"`

	// The OCID of the security policy deployment associated.
	SecurityPolicyDeploymentId *string `mandatory:"false" json:"securityPolicyDeploymentId"`

	EntryDetails EntryDetails `mandatory:"false" json:"entryDetails"`
}

func (m SecurityPolicyEntryState) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SecurityPolicyEntryState) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSecurityPolicyEntryStateDeploymentStatusEnum(string(m.DeploymentStatus)); !ok && m.DeploymentStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DeploymentStatus: %s. Supported values are: %s.", m.DeploymentStatus, strings.Join(GetSecurityPolicyEntryStateDeploymentStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *SecurityPolicyEntryState) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		SecurityPolicyDeploymentId *string                                      `json:"securityPolicyDeploymentId"`
		EntryDetails               entrydetails                                 `json:"entryDetails"`
		Id                         *string                                      `json:"id"`
		SecurityPolicyEntryId      *string                                      `json:"securityPolicyEntryId"`
		DeploymentStatus           SecurityPolicyEntryStateDeploymentStatusEnum `json:"deploymentStatus"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.SecurityPolicyDeploymentId = model.SecurityPolicyDeploymentId

	nn, e = model.EntryDetails.UnmarshalPolymorphicJSON(model.EntryDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.EntryDetails = nn.(EntryDetails)
	} else {
		m.EntryDetails = nil
	}

	m.Id = model.Id

	m.SecurityPolicyEntryId = model.SecurityPolicyEntryId

	m.DeploymentStatus = model.DeploymentStatus

	return
}
