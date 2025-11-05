// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DataIntelligences Control Plane API
//
// Use the DataIntelligences Control Plane API to manage dataIntelligences.
//

package dif

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DeployArtifactsDetails The data to create a DataIntelligence.
type DeployArtifactsDetails struct {

	// List of templates to be onboarded for the stack.
	StackTemplates []StackTemplateEnum `mandatory:"true" json:"stackTemplates"`

	// List of services to be onboarded for the stack.
	Services []ServiceEnum `mandatory:"true" json:"services"`

	// Subnet id for the Private Endpoint creation for artifact deployment.
	SubnetId *string `mandatory:"false" json:"subnetId"`

	// ADB artifact details if adb is included in the services.
	Adb []AdbArtifactsDetail `mandatory:"false" json:"adb"`

	// GGCS artifact details if ggcs is included in the services.
	Ggcs []GgcsArtifactsDetail `mandatory:"false" json:"ggcs"`

	// Dataflow artifact details if dataflow is included in the services.
	Dataflow []DataflowArtifactsDetail `mandatory:"false" json:"dataflow"`
}

func (m DeployArtifactsDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DeployArtifactsDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	for _, val := range m.StackTemplates {
		if _, ok := GetMappingStackTemplateEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for StackTemplates: %s. Supported values are: %s.", val, strings.Join(GetStackTemplateEnumStringValues(), ",")))
		}
	}

	for _, val := range m.Services {
		if _, ok := GetMappingServiceEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Services: %s. Supported values are: %s.", val, strings.Join(GetServiceEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
