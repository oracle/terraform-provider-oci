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

// CreateSourceApplicationInfo The information about the application.
type CreateSourceApplicationInfo struct {

	// The OCID of the workspace containing the application. This allows cross workspace deployment to publish an application from a different workspace into the current workspace specified in this operation.
	WorkspaceId *string `mandatory:"false" json:"workspaceId"`

	// The source application key to use when creating the application.
	ApplicationKey *string `mandatory:"false" json:"applicationKey"`

	// Parameter to specify the link between SOURCE and TARGET application after copying. CONNECTED    - Indicate that TARGET application is conneced to SOURCE and can be synced after copy. DISCONNECTED - Indicate that TARGET application is not conneced to SOURCE and can evolve independently.
	CopyType CreateSourceApplicationInfoCopyTypeEnum `mandatory:"false" json:"copyType,omitempty"`
}

func (m CreateSourceApplicationInfo) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateSourceApplicationInfo) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingCreateSourceApplicationInfoCopyTypeEnum(string(m.CopyType)); !ok && m.CopyType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CopyType: %s. Supported values are: %s.", m.CopyType, strings.Join(GetCreateSourceApplicationInfoCopyTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateSourceApplicationInfoCopyTypeEnum Enum with underlying type: string
type CreateSourceApplicationInfoCopyTypeEnum string

// Set of constants representing the allowable values for CreateSourceApplicationInfoCopyTypeEnum
const (
	CreateSourceApplicationInfoCopyTypeConnected    CreateSourceApplicationInfoCopyTypeEnum = "CONNECTED"
	CreateSourceApplicationInfoCopyTypeDisconnected CreateSourceApplicationInfoCopyTypeEnum = "DISCONNECTED"
)

var mappingCreateSourceApplicationInfoCopyTypeEnum = map[string]CreateSourceApplicationInfoCopyTypeEnum{
	"CONNECTED":    CreateSourceApplicationInfoCopyTypeConnected,
	"DISCONNECTED": CreateSourceApplicationInfoCopyTypeDisconnected,
}

var mappingCreateSourceApplicationInfoCopyTypeEnumLowerCase = map[string]CreateSourceApplicationInfoCopyTypeEnum{
	"connected":    CreateSourceApplicationInfoCopyTypeConnected,
	"disconnected": CreateSourceApplicationInfoCopyTypeDisconnected,
}

// GetCreateSourceApplicationInfoCopyTypeEnumValues Enumerates the set of values for CreateSourceApplicationInfoCopyTypeEnum
func GetCreateSourceApplicationInfoCopyTypeEnumValues() []CreateSourceApplicationInfoCopyTypeEnum {
	values := make([]CreateSourceApplicationInfoCopyTypeEnum, 0)
	for _, v := range mappingCreateSourceApplicationInfoCopyTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateSourceApplicationInfoCopyTypeEnumStringValues Enumerates the set of values in String for CreateSourceApplicationInfoCopyTypeEnum
func GetCreateSourceApplicationInfoCopyTypeEnumStringValues() []string {
	return []string{
		"CONNECTED",
		"DISCONNECTED",
	}
}

// GetMappingCreateSourceApplicationInfoCopyTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateSourceApplicationInfoCopyTypeEnum(val string) (CreateSourceApplicationInfoCopyTypeEnum, bool) {
	enum, ok := mappingCreateSourceApplicationInfoCopyTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
