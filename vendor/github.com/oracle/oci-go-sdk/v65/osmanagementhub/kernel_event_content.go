// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OS Management Hub API
//
// Use the OS Management Hub API to manage and monitor updates and patches for the operating system environments in your private data centers through a single management console. For more information, see Overview of OS Management Hub (https://docs.cloud.oracle.com/iaas/osmh/doc/overview.htm).
//

package osmanagementhub

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// KernelEventContent Kernel event content
type KernelEventContent struct {

	// Kernel event content location
	ContentLocation *string `mandatory:"true" json:"contentLocation"`

	// content size
	Size *int `mandatory:"false" json:"size"`

	// Crash Content Availability Status:
	//   * NOT_AVAILABLE - content is not available on the instance/resource nor in the service
	//   * AVAILABLE_ON_INSTANCE - content is only available on the instance/resource
	//   * AVAILABLE_ON_SERVICE - content is only available on the service
	//   * AVAILABLE_ON_INSTANCE_AND_SERVICE - content is available both on the instance/resource and the service
	//   * AVAILABLE_ON_INSTANCE_UPLOAD_IN_PROGRESS - content is available on the instance/resource and its upload is in progress
	ContentAvailability KernelEventContentContentAvailabilityEnum `mandatory:"true" json:"contentAvailability"`
}

func (m KernelEventContent) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m KernelEventContent) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingKernelEventContentContentAvailabilityEnum(string(m.ContentAvailability)); !ok && m.ContentAvailability != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ContentAvailability: %s. Supported values are: %s.", m.ContentAvailability, strings.Join(GetKernelEventContentContentAvailabilityEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m KernelEventContent) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeKernelEventContent KernelEventContent
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeKernelEventContent
	}{
		"KERNEL",
		(MarshalTypeKernelEventContent)(m),
	}

	return json.Marshal(&s)
}

// KernelEventContentContentAvailabilityEnum Enum with underlying type: string
type KernelEventContentContentAvailabilityEnum string

// Set of constants representing the allowable values for KernelEventContentContentAvailabilityEnum
const (
	KernelEventContentContentAvailabilityNotAvailable                        KernelEventContentContentAvailabilityEnum = "NOT_AVAILABLE"
	KernelEventContentContentAvailabilityAvailableOnInstance                 KernelEventContentContentAvailabilityEnum = "AVAILABLE_ON_INSTANCE"
	KernelEventContentContentAvailabilityAvailableOnService                  KernelEventContentContentAvailabilityEnum = "AVAILABLE_ON_SERVICE"
	KernelEventContentContentAvailabilityAvailableOnInstanceAndService       KernelEventContentContentAvailabilityEnum = "AVAILABLE_ON_INSTANCE_AND_SERVICE"
	KernelEventContentContentAvailabilityAvailableOnInstanceUploadInProgress KernelEventContentContentAvailabilityEnum = "AVAILABLE_ON_INSTANCE_UPLOAD_IN_PROGRESS"
)

var mappingKernelEventContentContentAvailabilityEnum = map[string]KernelEventContentContentAvailabilityEnum{
	"NOT_AVAILABLE":                            KernelEventContentContentAvailabilityNotAvailable,
	"AVAILABLE_ON_INSTANCE":                    KernelEventContentContentAvailabilityAvailableOnInstance,
	"AVAILABLE_ON_SERVICE":                     KernelEventContentContentAvailabilityAvailableOnService,
	"AVAILABLE_ON_INSTANCE_AND_SERVICE":        KernelEventContentContentAvailabilityAvailableOnInstanceAndService,
	"AVAILABLE_ON_INSTANCE_UPLOAD_IN_PROGRESS": KernelEventContentContentAvailabilityAvailableOnInstanceUploadInProgress,
}

var mappingKernelEventContentContentAvailabilityEnumLowerCase = map[string]KernelEventContentContentAvailabilityEnum{
	"not_available":                            KernelEventContentContentAvailabilityNotAvailable,
	"available_on_instance":                    KernelEventContentContentAvailabilityAvailableOnInstance,
	"available_on_service":                     KernelEventContentContentAvailabilityAvailableOnService,
	"available_on_instance_and_service":        KernelEventContentContentAvailabilityAvailableOnInstanceAndService,
	"available_on_instance_upload_in_progress": KernelEventContentContentAvailabilityAvailableOnInstanceUploadInProgress,
}

// GetKernelEventContentContentAvailabilityEnumValues Enumerates the set of values for KernelEventContentContentAvailabilityEnum
func GetKernelEventContentContentAvailabilityEnumValues() []KernelEventContentContentAvailabilityEnum {
	values := make([]KernelEventContentContentAvailabilityEnum, 0)
	for _, v := range mappingKernelEventContentContentAvailabilityEnum {
		values = append(values, v)
	}
	return values
}

// GetKernelEventContentContentAvailabilityEnumStringValues Enumerates the set of values in String for KernelEventContentContentAvailabilityEnum
func GetKernelEventContentContentAvailabilityEnumStringValues() []string {
	return []string{
		"NOT_AVAILABLE",
		"AVAILABLE_ON_INSTANCE",
		"AVAILABLE_ON_SERVICE",
		"AVAILABLE_ON_INSTANCE_AND_SERVICE",
		"AVAILABLE_ON_INSTANCE_UPLOAD_IN_PROGRESS",
	}
}

// GetMappingKernelEventContentContentAvailabilityEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingKernelEventContentContentAvailabilityEnum(val string) (KernelEventContentContentAvailabilityEnum, bool) {
	enum, ok := mappingKernelEventContentContentAvailabilityEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
