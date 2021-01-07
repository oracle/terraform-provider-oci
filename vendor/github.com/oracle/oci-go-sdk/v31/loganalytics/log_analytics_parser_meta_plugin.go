// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"github.com/oracle/oci-go-sdk/v31/common"
)

// LogAnalyticsParserMetaPlugin LogAnalyticsParserMetaPlugin
type LogAnalyticsParserMetaPlugin struct {

	// parameter list
	MetaPluginParameters []LogAnalyticsParserMetaPluginParameter `mandatory:"false" json:"metaPluginParameters"`

	// plugin description
	Description *string `mandatory:"false" json:"description"`

	// plugin display name
	DisplayName *string `mandatory:"false" json:"displayName"`

	// plugin internal name
	Name *string `mandatory:"false" json:"name"`
}

func (m LogAnalyticsParserMetaPlugin) String() string {
	return common.PointerString(m)
}
