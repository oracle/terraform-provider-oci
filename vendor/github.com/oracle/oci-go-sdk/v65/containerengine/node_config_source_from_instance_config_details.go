// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Kubernetes Engine API
//
// API for the Kubernetes Engine service (also known as the Container Engine for Kubernetes service). Use this API to build, deploy,
// and manage cloud-native applications. For more information, see
// Overview of Kubernetes Engine (https://docs.oracle.com/iaas/Content/ContEng/Concepts/contengoverview.htm).
//

package containerengine

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// NodeConfigSourceFromInstanceConfigDetails Configure nodes from existing InstanceConfiguration object.
type NodeConfigSourceFromInstanceConfigDetails struct {

	// The ocid of the InstanceConfiguration for the nodepool
	InstanceConfigurationId *string `mandatory:"true" json:"instanceConfigurationId"`
}

func (m NodeConfigSourceFromInstanceConfigDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m NodeConfigSourceFromInstanceConfigDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m NodeConfigSourceFromInstanceConfigDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeNodeConfigSourceFromInstanceConfigDetails NodeConfigSourceFromInstanceConfigDetails
	s := struct {
		DiscriminatorParam string `json:"configurationSource"`
		MarshalTypeNodeConfigSourceFromInstanceConfigDetails
	}{
		"INSTANCE_CONFIG",
		(MarshalTypeNodeConfigSourceFromInstanceConfigDetails)(m),
	}

	return json.Marshal(&s)
}
