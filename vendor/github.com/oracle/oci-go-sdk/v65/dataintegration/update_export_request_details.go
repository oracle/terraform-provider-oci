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

// UpdateExportRequestDetails Properties used in export object request update operations.
type UpdateExportRequestDetails struct {

	// The status of the object.
	Status UpdateExportRequestDetailsStatusEnum `mandatory:"false" json:"status,omitempty"`
}

func (m UpdateExportRequestDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateExportRequestDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingUpdateExportRequestDetailsStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetUpdateExportRequestDetailsStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UpdateExportRequestDetailsStatusEnum Enum with underlying type: string
type UpdateExportRequestDetailsStatusEnum string

// Set of constants representing the allowable values for UpdateExportRequestDetailsStatusEnum
const (
	UpdateExportRequestDetailsStatusTerminating UpdateExportRequestDetailsStatusEnum = "TERMINATING"
)

var mappingUpdateExportRequestDetailsStatusEnum = map[string]UpdateExportRequestDetailsStatusEnum{
	"TERMINATING": UpdateExportRequestDetailsStatusTerminating,
}

var mappingUpdateExportRequestDetailsStatusEnumLowerCase = map[string]UpdateExportRequestDetailsStatusEnum{
	"terminating": UpdateExportRequestDetailsStatusTerminating,
}

// GetUpdateExportRequestDetailsStatusEnumValues Enumerates the set of values for UpdateExportRequestDetailsStatusEnum
func GetUpdateExportRequestDetailsStatusEnumValues() []UpdateExportRequestDetailsStatusEnum {
	values := make([]UpdateExportRequestDetailsStatusEnum, 0)
	for _, v := range mappingUpdateExportRequestDetailsStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateExportRequestDetailsStatusEnumStringValues Enumerates the set of values in String for UpdateExportRequestDetailsStatusEnum
func GetUpdateExportRequestDetailsStatusEnumStringValues() []string {
	return []string{
		"TERMINATING",
	}
}

// GetMappingUpdateExportRequestDetailsStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateExportRequestDetailsStatusEnum(val string) (UpdateExportRequestDetailsStatusEnum, bool) {
	enum, ok := mappingUpdateExportRequestDetailsStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
