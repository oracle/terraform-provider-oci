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

// ServiceDetailResponse Details of the cloud service.
type ServiceDetailResponse struct {

	// ID for the service instance.
	InstanceId *string `mandatory:"false" json:"instanceId"`

	// ID for the service
	ServiceId *string `mandatory:"false" json:"serviceId"`

	// name of the service
	DisplayName *string `mandatory:"false" json:"displayName"`

	// name of the service
	CurrentArtifactPath *string `mandatory:"false" json:"currentArtifactPath"`

	// name of the cloud service
	ServiceType *string `mandatory:"false" json:"serviceType"`

	// url for the service
	ServiceUrl *string `mandatory:"false" json:"serviceUrl"`

	// state of the service
	Status *string `mandatory:"false" json:"status"`

	AdditionalDetails *AdditionalDetails `mandatory:"false" json:"additionalDetails"`
}

func (m ServiceDetailResponse) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ServiceDetailResponse) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
