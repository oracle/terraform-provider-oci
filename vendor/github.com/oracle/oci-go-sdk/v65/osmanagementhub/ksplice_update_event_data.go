// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OS Management Hub API
//
// Use the OS Management Hub API to manage and monitor updates and patches for instances in OCI, your private data center, or 3rd-party clouds.
// For more information, see Overview of OS Management Hub (https://docs.cloud.oracle.com/iaas/osmh/doc/overview.htm).
//

package osmanagementhub

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// KspliceUpdateEventData Provides additional information for a Ksplice update event.
type KspliceUpdateEventData struct {

	// The type of Ksplice update.
	OperationType KspliceUpdateEventDataOperationTypeEnum `mandatory:"true" json:"operationType"`

	// Status of the Ksplice update.
	Status EventStatusEnum `mandatory:"true" json:"status"`

	AdditionalDetails *WorkRequestEventDataAdditionalDetails `mandatory:"false" json:"additionalDetails"`
}

func (m KspliceUpdateEventData) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m KspliceUpdateEventData) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingKspliceUpdateEventDataOperationTypeEnum(string(m.OperationType)); !ok && m.OperationType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OperationType: %s. Supported values are: %s.", m.OperationType, strings.Join(GetKspliceUpdateEventDataOperationTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingEventStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetEventStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// KspliceUpdateEventDataOperationTypeEnum Enum with underlying type: string
type KspliceUpdateEventDataOperationTypeEnum string

// Set of constants representing the allowable values for KspliceUpdateEventDataOperationTypeEnum
const (
	KspliceUpdateEventDataOperationTypeKernel    KspliceUpdateEventDataOperationTypeEnum = "UPDATE_KSPLICE_KERNEL"
	KspliceUpdateEventDataOperationTypeUserspace KspliceUpdateEventDataOperationTypeEnum = "UPDATE_KSPLICE_USERSPACE"
)

var mappingKspliceUpdateEventDataOperationTypeEnum = map[string]KspliceUpdateEventDataOperationTypeEnum{
	"UPDATE_KSPLICE_KERNEL":    KspliceUpdateEventDataOperationTypeKernel,
	"UPDATE_KSPLICE_USERSPACE": KspliceUpdateEventDataOperationTypeUserspace,
}

var mappingKspliceUpdateEventDataOperationTypeEnumLowerCase = map[string]KspliceUpdateEventDataOperationTypeEnum{
	"update_ksplice_kernel":    KspliceUpdateEventDataOperationTypeKernel,
	"update_ksplice_userspace": KspliceUpdateEventDataOperationTypeUserspace,
}

// GetKspliceUpdateEventDataOperationTypeEnumValues Enumerates the set of values for KspliceUpdateEventDataOperationTypeEnum
func GetKspliceUpdateEventDataOperationTypeEnumValues() []KspliceUpdateEventDataOperationTypeEnum {
	values := make([]KspliceUpdateEventDataOperationTypeEnum, 0)
	for _, v := range mappingKspliceUpdateEventDataOperationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetKspliceUpdateEventDataOperationTypeEnumStringValues Enumerates the set of values in String for KspliceUpdateEventDataOperationTypeEnum
func GetKspliceUpdateEventDataOperationTypeEnumStringValues() []string {
	return []string{
		"UPDATE_KSPLICE_KERNEL",
		"UPDATE_KSPLICE_USERSPACE",
	}
}

// GetMappingKspliceUpdateEventDataOperationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingKspliceUpdateEventDataOperationTypeEnum(val string) (KspliceUpdateEventDataOperationTypeEnum, bool) {
	enum, ok := mappingKspliceUpdateEventDataOperationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
