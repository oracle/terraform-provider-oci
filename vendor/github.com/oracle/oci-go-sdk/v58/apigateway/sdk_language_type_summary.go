// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// API Gateway API
//
// API for the API Gateway service. Use this API to manage gateways, deployments, and related items.
// For more information, see
// Overview of API Gateway (https://docs.cloud.oracle.com/iaas/Content/APIGateway/Concepts/apigatewayoverview.htm).
//

package apigateway

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// SdkLanguageTypeSummary SDK target language details.
type SdkLanguageTypeSummary struct {

	// Name of the programming language.
	Name *string `mandatory:"true" json:"name"`

	// Version string of the programming language defined in name.
	Version *string `mandatory:"true" json:"version"`

	// Display name of the target programming language.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Additional details.
	Description *string `mandatory:"false" json:"description"`

	// List of optional configurations that can be used while generating SDK for the given target language.
	Parameters []SdkLanguageOptionalParameters `mandatory:"false" json:"parameters"`
}

func (m SdkLanguageTypeSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SdkLanguageTypeSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
