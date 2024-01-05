// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OS Management API
//
// API for the OS Management service. Use these API operations for working
// with Managed instances and Managed instance groups.
//

package osmanagement

import (
	"strings"
)

// ContentAvailabilityEnum Enum with underlying type: string
type ContentAvailabilityEnum string

// Set of constants representing the allowable values for ContentAvailabilityEnum
const (
	ContentAvailabilityNotAvailable                        ContentAvailabilityEnum = "NOT_AVAILABLE"
	ContentAvailabilityAvailableOnInstance                 ContentAvailabilityEnum = "AVAILABLE_ON_INSTANCE"
	ContentAvailabilityAvailableOnService                  ContentAvailabilityEnum = "AVAILABLE_ON_SERVICE"
	ContentAvailabilityAvailableOnInstanceAndService       ContentAvailabilityEnum = "AVAILABLE_ON_INSTANCE_AND_SERVICE"
	ContentAvailabilityAvailableOnInstanceUploadInProgress ContentAvailabilityEnum = "AVAILABLE_ON_INSTANCE_UPLOAD_IN_PROGRESS"
)

var mappingContentAvailabilityEnum = map[string]ContentAvailabilityEnum{
	"NOT_AVAILABLE":                            ContentAvailabilityNotAvailable,
	"AVAILABLE_ON_INSTANCE":                    ContentAvailabilityAvailableOnInstance,
	"AVAILABLE_ON_SERVICE":                     ContentAvailabilityAvailableOnService,
	"AVAILABLE_ON_INSTANCE_AND_SERVICE":        ContentAvailabilityAvailableOnInstanceAndService,
	"AVAILABLE_ON_INSTANCE_UPLOAD_IN_PROGRESS": ContentAvailabilityAvailableOnInstanceUploadInProgress,
}

var mappingContentAvailabilityEnumLowerCase = map[string]ContentAvailabilityEnum{
	"not_available":                            ContentAvailabilityNotAvailable,
	"available_on_instance":                    ContentAvailabilityAvailableOnInstance,
	"available_on_service":                     ContentAvailabilityAvailableOnService,
	"available_on_instance_and_service":        ContentAvailabilityAvailableOnInstanceAndService,
	"available_on_instance_upload_in_progress": ContentAvailabilityAvailableOnInstanceUploadInProgress,
}

// GetContentAvailabilityEnumValues Enumerates the set of values for ContentAvailabilityEnum
func GetContentAvailabilityEnumValues() []ContentAvailabilityEnum {
	values := make([]ContentAvailabilityEnum, 0)
	for _, v := range mappingContentAvailabilityEnum {
		values = append(values, v)
	}
	return values
}

// GetContentAvailabilityEnumStringValues Enumerates the set of values in String for ContentAvailabilityEnum
func GetContentAvailabilityEnumStringValues() []string {
	return []string{
		"NOT_AVAILABLE",
		"AVAILABLE_ON_INSTANCE",
		"AVAILABLE_ON_SERVICE",
		"AVAILABLE_ON_INSTANCE_AND_SERVICE",
		"AVAILABLE_ON_INSTANCE_UPLOAD_IN_PROGRESS",
	}
}

// GetMappingContentAvailabilityEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingContentAvailabilityEnum(val string) (ContentAvailabilityEnum, bool) {
	enum, ok := mappingContentAvailabilityEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
