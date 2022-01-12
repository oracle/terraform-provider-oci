// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Application Performance Monitoring Synthetic Monitoring API
//
// Use the Application Performance Monitoring Synthetic Monitoring API to query synthetic scripts and monitors.
//

package apmsynthetics

// ContentTypesEnum Enum with underlying type: string
type ContentTypesEnum string

// Set of constants representing the allowable values for ContentTypesEnum
const (
	ContentTypesSide ContentTypesEnum = "SIDE"
	ContentTypesJs   ContentTypesEnum = "JS"
)

var mappingContentTypes = map[string]ContentTypesEnum{
	"SIDE": ContentTypesSide,
	"JS":   ContentTypesJs,
}

// GetContentTypesEnumValues Enumerates the set of values for ContentTypesEnum
func GetContentTypesEnumValues() []ContentTypesEnum {
	values := make([]ContentTypesEnum, 0)
	for _, v := range mappingContentTypes {
		values = append(values, v)
	}
	return values
}
