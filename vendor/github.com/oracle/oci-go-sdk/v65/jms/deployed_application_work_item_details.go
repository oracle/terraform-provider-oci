// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Java Management Service Fleets API
//
// The APIs for the Fleet Management (https://docs.oracle.com/en-us/iaas/jms/doc/fleet-management.html) feature of Java Management Service to monitor and manage the usage of Java in your enterprise. Use these APIs to manage fleets, configure managed instances to report to fleets, and gain insights into the Java workloads running on these instances by carrying out basic and advanced features.
//

package jms

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DeployedApplicationWorkItemDetails The work item details for deployed application related information.
type DeployedApplicationWorkItemDetails struct {

	// The unique key of the deployed application.
	DeployedApplicationKey *string `mandatory:"true" json:"deployedApplicationKey"`

	// The deployed application name.
	DeployedApplicationName *string `mandatory:"true" json:"deployedApplicationName"`

	// The unique key of the deployed application installation.
	DeployedApplicationInstallationKey *string `mandatory:"false" json:"deployedApplicationInstallationKey"`

	// The full path on which deployed application installation was detected.
	DeployedApplicationInstallationPath *string `mandatory:"false" json:"deployedApplicationInstallationPath"`

	// The work item type.
	WorkItemType WorkItemTypeEnum `mandatory:"false" json:"workItemType,omitempty"`
}

// GetWorkItemType returns WorkItemType
func (m DeployedApplicationWorkItemDetails) GetWorkItemType() WorkItemTypeEnum {
	return m.WorkItemType
}

func (m DeployedApplicationWorkItemDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DeployedApplicationWorkItemDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingWorkItemTypeEnum(string(m.WorkItemType)); !ok && m.WorkItemType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for WorkItemType: %s. Supported values are: %s.", m.WorkItemType, strings.Join(GetWorkItemTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m DeployedApplicationWorkItemDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDeployedApplicationWorkItemDetails DeployedApplicationWorkItemDetails
	s := struct {
		DiscriminatorParam string `json:"kind"`
		MarshalTypeDeployedApplicationWorkItemDetails
	}{
		"DEPLOYED_APPLICATION",
		(MarshalTypeDeployedApplicationWorkItemDetails)(m),
	}

	return json.Marshal(&s)
}
