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

// UpdateCopyObjectRequestDetails Properties used in copy object request update operations.
type UpdateCopyObjectRequestDetails struct {

	// The status of the object.
	Status UpdateCopyObjectRequestDetailsStatusEnum `mandatory:"false" json:"status,omitempty"`
}

func (m UpdateCopyObjectRequestDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateCopyObjectRequestDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingUpdateCopyObjectRequestDetailsStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetUpdateCopyObjectRequestDetailsStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UpdateCopyObjectRequestDetailsStatusEnum Enum with underlying type: string
type UpdateCopyObjectRequestDetailsStatusEnum string

// Set of constants representing the allowable values for UpdateCopyObjectRequestDetailsStatusEnum
const (
	UpdateCopyObjectRequestDetailsStatusTerminating UpdateCopyObjectRequestDetailsStatusEnum = "TERMINATING"
)

var mappingUpdateCopyObjectRequestDetailsStatusEnum = map[string]UpdateCopyObjectRequestDetailsStatusEnum{
	"TERMINATING": UpdateCopyObjectRequestDetailsStatusTerminating,
}

var mappingUpdateCopyObjectRequestDetailsStatusEnumLowerCase = map[string]UpdateCopyObjectRequestDetailsStatusEnum{
	"terminating": UpdateCopyObjectRequestDetailsStatusTerminating,
}

// GetUpdateCopyObjectRequestDetailsStatusEnumValues Enumerates the set of values for UpdateCopyObjectRequestDetailsStatusEnum
func GetUpdateCopyObjectRequestDetailsStatusEnumValues() []UpdateCopyObjectRequestDetailsStatusEnum {
	values := make([]UpdateCopyObjectRequestDetailsStatusEnum, 0)
	for _, v := range mappingUpdateCopyObjectRequestDetailsStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateCopyObjectRequestDetailsStatusEnumStringValues Enumerates the set of values in String for UpdateCopyObjectRequestDetailsStatusEnum
func GetUpdateCopyObjectRequestDetailsStatusEnumStringValues() []string {
	return []string{
		"TERMINATING",
	}
}

// GetMappingUpdateCopyObjectRequestDetailsStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateCopyObjectRequestDetailsStatusEnum(val string) (UpdateCopyObjectRequestDetailsStatusEnum, bool) {
	enum, ok := mappingUpdateCopyObjectRequestDetailsStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
