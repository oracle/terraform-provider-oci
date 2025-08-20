// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// DeployedApplicationMigrationTaskDetails The task details with deployed application migration related information.
type DeployedApplicationMigrationTaskDetails struct {
	DeployedApplicationMigrationTaskRequest *RequestDeployedApplicationMigrationAnalysesDetails `mandatory:"false" json:"deployedApplicationMigrationTaskRequest"`
}

func (m DeployedApplicationMigrationTaskDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DeployedApplicationMigrationTaskDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m DeployedApplicationMigrationTaskDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDeployedApplicationMigrationTaskDetails DeployedApplicationMigrationTaskDetails
	s := struct {
		DiscriminatorParam string `json:"taskType"`
		MarshalTypeDeployedApplicationMigrationTaskDetails
	}{
		"DEPLOYED_APPLICATION_MIGRATION",
		(MarshalTypeDeployedApplicationMigrationTaskDetails)(m),
	}

	return json.Marshal(&s)
}
