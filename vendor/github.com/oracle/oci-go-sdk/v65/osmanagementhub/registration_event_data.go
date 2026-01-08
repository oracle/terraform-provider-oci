// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OS Management Hub API
//
// Use the OS Management Hub API to manage and monitor updates and patches for instances in OCI, your private data center, or 3rd-party clouds.
// For more information, see Overview of OS Management Hub (https://docs.oracle.com/iaas/osmh/doc/overview.htm).
//

package osmanagementhub

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// RegistrationEventData Events related to the instance registration.
type RegistrationEventData struct {

	// The specific type of Registration Event.
	RegistrationEventType RegistrationEventTypeEnum `mandatory:"true" json:"registrationEventType"`

	// Status of registration.
	Status EventStatusEnum `mandatory:"false" json:"status,omitempty"`

	MatchedRepository *MatchedRepository `mandatory:"false" json:"matchedRepository"`

	AdditionalDetails *WorkRequestEventDataAdditionalDetails `mandatory:"false" json:"additionalDetails"`
}

func (m RegistrationEventData) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RegistrationEventData) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingRegistrationEventTypeEnum(string(m.RegistrationEventType)); !ok && m.RegistrationEventType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RegistrationEventType: %s. Supported values are: %s.", m.RegistrationEventType, strings.Join(GetRegistrationEventTypeEnumStringValues(), ",")))
	}

	if _, ok := GetMappingEventStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetEventStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
