// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// LogAnalyticsParserMetaPluginParameter LogAnalyticsParserMetaPluginParameter
type LogAnalyticsParserMetaPluginParameter struct {

	// The parameter description.
	Description *string `mandatory:"false" json:"description"`

	// The parameter internal name.
	Name *string `mandatory:"false" json:"name"`

	// A flag indicating whether or not the parameter is mandatory.
	IsMandatory *bool `mandatory:"false" json:"isMandatory"`

	// A flag indicating whether or not the parameter is repeatable.
	IsRepeatable *bool `mandatory:"false" json:"isRepeatable"`

	// The plugin internal name.
	PluginName *string `mandatory:"false" json:"pluginName"`

	// The parameter type.
	Type *string `mandatory:"false" json:"type"`
}

func (m LogAnalyticsParserMetaPluginParameter) String() string {
	return common.PointerString(m)
}
