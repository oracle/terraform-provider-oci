// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"github.com/oracle/oci-go-sdk/v30/common"
)

// LogAnalyticsParserMetaPluginParameter LogAnalyticsParserMetaPluginParameter
type LogAnalyticsParserMetaPluginParameter struct {

	// parameter description
	Description *string `mandatory:"false" json:"description"`

	// parameter internal name
	Name *string `mandatory:"false" json:"name"`

	// is mandatory flag
	IsMandatory *bool `mandatory:"false" json:"isMandatory"`

	// is repeatable flag
	IsRepeatable *bool `mandatory:"false" json:"isRepeatable"`

	// plugin internal name
	PluginName *string `mandatory:"false" json:"pluginName"`

	// parameter type
	Type *string `mandatory:"false" json:"type"`
}

func (m LogAnalyticsParserMetaPluginParameter) String() string {
	return common.PointerString(m)
}
