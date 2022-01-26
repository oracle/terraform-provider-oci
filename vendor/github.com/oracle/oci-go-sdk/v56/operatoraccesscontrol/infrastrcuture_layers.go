// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OperatorAccessControl API
//
// Operator Access Control enables you to control the time duration and the actions an Oracle operator can perform on your Exadata Cloud@Customer infrastructure.
// Using logging service, you can view a near real-time audit report of all actions performed by an Oracle operator.
// Use the table of contents and search tool to explore the OperatorAccessControl API.
//

package operatoraccesscontrol

// InfrastrcutureLayersEnum Enum with underlying type: string
type InfrastrcutureLayersEnum string

// Set of constants representing the allowable values for InfrastrcutureLayersEnum
const (
	InfrastrcutureLayersDom0       InfrastrcutureLayersEnum = "DOM0"
	InfrastrcutureLayersCellserver InfrastrcutureLayersEnum = "CELLSERVER"
	InfrastrcutureLayersCps        InfrastrcutureLayersEnum = "CPS"
)

var mappingInfrastrcutureLayers = map[string]InfrastrcutureLayersEnum{
	"DOM0":       InfrastrcutureLayersDom0,
	"CELLSERVER": InfrastrcutureLayersCellserver,
	"CPS":        InfrastrcutureLayersCps,
}

// GetInfrastrcutureLayersEnumValues Enumerates the set of values for InfrastrcutureLayersEnum
func GetInfrastrcutureLayersEnumValues() []InfrastrcutureLayersEnum {
	values := make([]InfrastrcutureLayersEnum, 0)
	for _, v := range mappingInfrastrcutureLayers {
		values = append(values, v)
	}
	return values
}
