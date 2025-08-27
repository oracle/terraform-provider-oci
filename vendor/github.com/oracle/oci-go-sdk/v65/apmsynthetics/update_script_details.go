// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// APM Availability Monitoring API
//
// Use the APM Availability Monitoring API to query Scripts, Monitors, Dedicated Vantage Points and On-Premise Vantage Points resources. For more information, see Application Performance Monitoring (https://docs.oracle.com/iaas/application-performance-monitoring/index.html).
//

package apmsynthetics

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateScriptDetails Details of the request body used to update a script.
// Only Side, JavaScript and Playwright TypeScript content types are supported and content should be in Side, JavaScript and TypeScript formats only.
type UpdateScriptDetails struct {

	// Unique name that can be edited. The name should not contain any confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Content type of script.
	ContentType ContentTypesEnum `mandatory:"false" json:"contentType,omitempty"`

	// The content of the script. It may contain custom-defined tags that can be used for setting dynamic parameters.
	// The format to set dynamic parameters is: `<ORAP><ON>param name</ON><OV>param value</OV><OS>isParamValueSecret(true/false)</OS></ORAP>`.
	// Param value and isParamValueSecret are optional, the default value for isParamValueSecret is false.
	// Examples:
	// With mandatory param name : `<ORAP><ON>param name</ON></ORAP>`
	// With parameter name and value : `<ORAP><ON>param name</ON><OV>param value</OV></ORAP>`
	// Note that the content is valid if it matches the given content type. For example, if the content type is SIDE, then the content should be in Side script format. If the content type is JS, then the content should be in JavaScript format. If the content type is PLAYWRIGHT_TS, then the content should be in TypeScript format.
	Content *string `mandatory:"false" json:"content"`

	// File name of uploaded script content.
	ContentFileName *string `mandatory:"false" json:"contentFileName"`

	// List of script parameters. Example: `[{"paramName": "userid", "paramValue":"testuser", "isSecret": false}]`
	Parameters []ScriptParameter `mandatory:"false" json:"parameters"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m UpdateScriptDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateScriptDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingContentTypesEnum(string(m.ContentType)); !ok && m.ContentType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ContentType: %s. Supported values are: %s.", m.ContentType, strings.Join(GetContentTypesEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
