// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Guard APIs
//
// A description of the Cloud Guard APIs
//

package cloudguard

// ActorTypeEnum Enum with underlying type: string
type ActorTypeEnum string

// Set of constants representing the allowable values for ActorTypeEnum
const (
	ActorTypeCloudGuardService ActorTypeEnum = "CLOUD_GUARD_SERVICE"
	ActorTypeCorrelation       ActorTypeEnum = "CORRELATION"
	ActorTypeResponder         ActorTypeEnum = "RESPONDER"
	ActorTypeUser              ActorTypeEnum = "USER"
)

var mappingActorType = map[string]ActorTypeEnum{
	"CLOUD_GUARD_SERVICE": ActorTypeCloudGuardService,
	"CORRELATION":         ActorTypeCorrelation,
	"RESPONDER":           ActorTypeResponder,
	"USER":                ActorTypeUser,
}

// GetActorTypeEnumValues Enumerates the set of values for ActorTypeEnum
func GetActorTypeEnumValues() []ActorTypeEnum {
	values := make([]ActorTypeEnum, 0)
	for _, v := range mappingActorType {
		values = append(values, v)
	}
	return values
}
