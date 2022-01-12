// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Bastion API
//
// Oracle Cloud Infrastructure Bastion provides restricted and time-limited access to target resources that don't have public endpoints. Through the configuration of a bastion, you can let authorized users connect from specific IP addresses to target resources by way of Secure Shell (SSH) sessions hosted on the bastion.
//

package bastion

// OperationTypeEnum Enum with underlying type: string
type OperationTypeEnum string

// Set of constants representing the allowable values for OperationTypeEnum
const (
	OperationTypeCreateBastion OperationTypeEnum = "CREATE_BASTION"
	OperationTypeUpdateBastion OperationTypeEnum = "UPDATE_BASTION"
	OperationTypeDeleteBastion OperationTypeEnum = "DELETE_BASTION"
	OperationTypeCreateSession OperationTypeEnum = "CREATE_SESSION"
	OperationTypeDeleteSession OperationTypeEnum = "DELETE_SESSION"
)

var mappingOperationType = map[string]OperationTypeEnum{
	"CREATE_BASTION": OperationTypeCreateBastion,
	"UPDATE_BASTION": OperationTypeUpdateBastion,
	"DELETE_BASTION": OperationTypeDeleteBastion,
	"CREATE_SESSION": OperationTypeCreateSession,
	"DELETE_SESSION": OperationTypeDeleteSession,
}

// GetOperationTypeEnumValues Enumerates the set of values for OperationTypeEnum
func GetOperationTypeEnumValues() []OperationTypeEnum {
	values := make([]OperationTypeEnum, 0)
	for _, v := range mappingOperationType {
		values = append(values, v)
	}
	return values
}
