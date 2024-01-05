// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Integration API
//
// Use the Data Integration API to organize your data integration projects, create data flows, pipelines and tasks, and then publish, schedule, and run tasks that extract, transform, and load data. For more information, see Data Integration (https://docs.oracle.com/iaas/data-integration/home.htm).
//

package dataintegration

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateImportRequestDetails Properties used in import object request update operations.
type UpdateImportRequestDetails struct {

	// The status of the object.
	Status UpdateImportRequestDetailsStatusEnum `mandatory:"false" json:"status,omitempty"`
}

func (m UpdateImportRequestDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateImportRequestDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingUpdateImportRequestDetailsStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetUpdateImportRequestDetailsStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UpdateImportRequestDetailsStatusEnum Enum with underlying type: string
type UpdateImportRequestDetailsStatusEnum string

// Set of constants representing the allowable values for UpdateImportRequestDetailsStatusEnum
const (
	UpdateImportRequestDetailsStatusTerminating UpdateImportRequestDetailsStatusEnum = "TERMINATING"
)

var mappingUpdateImportRequestDetailsStatusEnum = map[string]UpdateImportRequestDetailsStatusEnum{
	"TERMINATING": UpdateImportRequestDetailsStatusTerminating,
}

var mappingUpdateImportRequestDetailsStatusEnumLowerCase = map[string]UpdateImportRequestDetailsStatusEnum{
	"terminating": UpdateImportRequestDetailsStatusTerminating,
}

// GetUpdateImportRequestDetailsStatusEnumValues Enumerates the set of values for UpdateImportRequestDetailsStatusEnum
func GetUpdateImportRequestDetailsStatusEnumValues() []UpdateImportRequestDetailsStatusEnum {
	values := make([]UpdateImportRequestDetailsStatusEnum, 0)
	for _, v := range mappingUpdateImportRequestDetailsStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateImportRequestDetailsStatusEnumStringValues Enumerates the set of values in String for UpdateImportRequestDetailsStatusEnum
func GetUpdateImportRequestDetailsStatusEnumStringValues() []string {
	return []string{
		"TERMINATING",
	}
}

// GetMappingUpdateImportRequestDetailsStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateImportRequestDetailsStatusEnum(val string) (UpdateImportRequestDetailsStatusEnum, bool) {
	enum, ok := mappingUpdateImportRequestDetailsStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
