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

// GenAiUpdateDetail Details required for existing GenAi instance to be updated.
type GenAiUpdateDetail struct {

	// Instance id of the exisitng GenAi instance to be updated.
	InstanceId *string `mandatory:"true" json:"instanceId"`

	// No of replicas of base model to be used for hosting.
	UnitCount *int `mandatory:"false" json:"unitCount"`

	// List of endpoints to be provisioned new or updated if existing for the GenAi dedicated cluster.
	Endpoints []EndpointDetails `mandatory:"false" json:"endpoints"`
}

func (m GenAiUpdateDetail) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m GenAiUpdateDetail) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
