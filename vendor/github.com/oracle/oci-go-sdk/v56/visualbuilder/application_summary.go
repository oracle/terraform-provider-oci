// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Visual Builder API
//
// Oracle Visual Builder enables developers to quickly build web and mobile applications. With a visual development environment that makes it easy to connect to Oracle data and third-party REST services, developers can build modern, consumer-grade applications in a fraction of the time it would take in other tools.
// The Visual Builder Instance Management API allows users to create and manage a Visual Builder instance.
//

package visualbuilder

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// ApplicationSummary Summary of the Vb Instance's applications.
type ApplicationSummary struct {

	// Unique identifier of the application.
	Id *string `mandatory:"true" json:"id"`

	// Project identifier.
	ProjectId *string `mandatory:"true" json:"projectId"`

	// Version of deployed application.
	Version *string `mandatory:"true" json:"version"`

	// Represents the deployment state of the application.
	State ApplicationSummaryStateEnum `mandatory:"true" json:"state"`
}

func (m ApplicationSummary) String() string {
	return common.PointerString(m)
}

// ApplicationSummaryStateEnum Enum with underlying type: string
type ApplicationSummaryStateEnum string

// Set of constants representing the allowable values for ApplicationSummaryStateEnum
const (
	ApplicationSummaryStateStage ApplicationSummaryStateEnum = "STAGE"
	ApplicationSummaryStateLive  ApplicationSummaryStateEnum = "LIVE"
)

var mappingApplicationSummaryState = map[string]ApplicationSummaryStateEnum{
	"STAGE": ApplicationSummaryStateStage,
	"LIVE":  ApplicationSummaryStateLive,
}

// GetApplicationSummaryStateEnumValues Enumerates the set of values for ApplicationSummaryStateEnum
func GetApplicationSummaryStateEnumValues() []ApplicationSummaryStateEnum {
	values := make([]ApplicationSummaryStateEnum, 0)
	for _, v := range mappingApplicationSummaryState {
		values = append(values, v)
	}
	return values
}
