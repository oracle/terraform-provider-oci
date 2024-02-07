// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Media Services API
//
// Media Services (includes Media Flow and Media Streams) is a fully managed service for processing media (video) source content. Use Media Flow and Media Streams to transcode and package digital video using configurable workflows and stream video outputs.
// Use the Media Services API to configure media workflows and run Media Flow jobs, create distribution channels, ingest assets, create Preview URLs and play assets. For more information, see Media Flow (https://docs.cloud.oracle.com/iaas/Content/dms-mediaflow/home.htm) and Media Streams (https://docs.cloud.oracle.com/iaas/Content/dms-mediastream/home.htm).
//

package mediaservices

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// RemoveLockDetails Request payload to remove lock to the resource.
type RemoveLockDetails struct {

	// Type of the lock.
	Type RemoveLockDetailsTypeEnum `mandatory:"true" json:"type"`

	// The compartment ID of the lock.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The ID of the resource that is locking this resource. Indicates that deleting this resource will remove the lock.
	RelatedResourceId *string `mandatory:"false" json:"relatedResourceId"`

	// A message added by the creator of the lock. This is typically used to give an
	// indication of why the resource is locked.
	Message *string `mandatory:"false" json:"message"`

	// When the lock was created.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`
}

func (m RemoveLockDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RemoveLockDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingRemoveLockDetailsTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetRemoveLockDetailsTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// RemoveLockDetailsTypeEnum Enum with underlying type: string
type RemoveLockDetailsTypeEnum string

// Set of constants representing the allowable values for RemoveLockDetailsTypeEnum
const (
	RemoveLockDetailsTypeFull   RemoveLockDetailsTypeEnum = "FULL"
	RemoveLockDetailsTypeDelete RemoveLockDetailsTypeEnum = "DELETE"
)

var mappingRemoveLockDetailsTypeEnum = map[string]RemoveLockDetailsTypeEnum{
	"FULL":   RemoveLockDetailsTypeFull,
	"DELETE": RemoveLockDetailsTypeDelete,
}

var mappingRemoveLockDetailsTypeEnumLowerCase = map[string]RemoveLockDetailsTypeEnum{
	"full":   RemoveLockDetailsTypeFull,
	"delete": RemoveLockDetailsTypeDelete,
}

// GetRemoveLockDetailsTypeEnumValues Enumerates the set of values for RemoveLockDetailsTypeEnum
func GetRemoveLockDetailsTypeEnumValues() []RemoveLockDetailsTypeEnum {
	values := make([]RemoveLockDetailsTypeEnum, 0)
	for _, v := range mappingRemoveLockDetailsTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetRemoveLockDetailsTypeEnumStringValues Enumerates the set of values in String for RemoveLockDetailsTypeEnum
func GetRemoveLockDetailsTypeEnumStringValues() []string {
	return []string{
		"FULL",
		"DELETE",
	}
}

// GetMappingRemoveLockDetailsTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRemoveLockDetailsTypeEnum(val string) (RemoveLockDetailsTypeEnum, bool) {
	enum, ok := mappingRemoveLockDetailsTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
