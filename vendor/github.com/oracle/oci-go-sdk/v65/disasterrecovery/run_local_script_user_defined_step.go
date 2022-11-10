// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Full Stack Disaster Recovery API
//
// Use the Full Stack Disaster Recovery (FSDR) API to manage disaster recovery for business applications.
// FSDR is an OCI disaster recovery orchestration and management service that provides comprehensive disaster recovery
// capabilities for all layers of an application stack, including infrastructure, middleware, database, and application.
//

package disasterrecovery

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// RunLocalScriptUserDefinedStep Run Local Script step details.
type RunLocalScriptUserDefinedStep struct {

	// The OCID of the instance where this script or command should be executed.
	// Example: `ocid1.instance.oc1.phx.exampleocid1`
	RunOnInstanceId *string `mandatory:"false" json:"runOnInstanceId"`

	// The region in which the instance is present.
	// Example: `us-phoenix-1`
	RunOnInstanceRegion *string `mandatory:"false" json:"runOnInstanceRegion"`

	// The script name and arguments.
	// Example: `/usr/bin/python3 /home/opc/scripts/my_app_script.py arg1 arg2 arg3`
	ScriptCommand *string `mandatory:"false" json:"scriptCommand"`

	// The userid on the instance to be used for executing the script or command.
	// Example: `opc`
	RunAsUser *string `mandatory:"false" json:"runAsUser"`
}

func (m RunLocalScriptUserDefinedStep) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RunLocalScriptUserDefinedStep) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m RunLocalScriptUserDefinedStep) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeRunLocalScriptUserDefinedStep RunLocalScriptUserDefinedStep
	s := struct {
		DiscriminatorParam string `json:"stepType"`
		MarshalTypeRunLocalScriptUserDefinedStep
	}{
		"RUN_LOCAL_SCRIPT",
		(MarshalTypeRunLocalScriptUserDefinedStep)(m),
	}

	return json.Marshal(&s)
}
