// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// LogAnalyticsParserMetaPlugin LogAnalyticsParserMetaPlugin
type LogAnalyticsParserMetaPlugin struct {

	// An array of plugin parameters.
	MetaPluginParameters []LogAnalyticsParserMetaPluginParameter `mandatory:"false" json:"metaPluginParameters"`

	// The plugin description.
	Description *string `mandatory:"false" json:"description"`

	// The plugin display name.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The plugin internal name.
	Name *string `mandatory:"false" json:"name"`
}

func (m LogAnalyticsParserMetaPlugin) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m LogAnalyticsParserMetaPlugin) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
