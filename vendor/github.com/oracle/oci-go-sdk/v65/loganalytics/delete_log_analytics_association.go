// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// DeleteLogAnalyticsAssociation DeleteLogAnalyticsAssociation
type DeleteLogAnalyticsAssociation struct {

	// The agent unique identifier.
	AgentId *string `mandatory:"false" json:"agentId"`

	// The source name.
	SourceName *string `mandatory:"false" json:"sourceName"`

	// The source type internal name.
	SourceTypeName *string `mandatory:"false" json:"sourceTypeName"`

	// The entity unique identifier.
	EntityId *string `mandatory:"false" json:"entityId"`

	// The entity type internal name.
	EntityTypeName *string `mandatory:"false" json:"entityTypeName"`

	// The host name.
	Host *string `mandatory:"false" json:"host"`

	// The log group unique identifier.
	LogGroupId *string `mandatory:"false" json:"logGroupId"`
}

func (m DeleteLogAnalyticsAssociation) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DeleteLogAnalyticsAssociation) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
