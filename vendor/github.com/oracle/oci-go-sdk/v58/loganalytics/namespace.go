// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// Namespace This is the namespace details of a tenancy in Logan Analytics application
type Namespace struct {

	// This is the namespace name of a tenancy
	NamespaceName *string `mandatory:"true" json:"namespaceName"`

	// The is the tenancy ID
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// This indicates if the tenancy is onboarded to Logging Analytics
	IsOnboarded *bool `mandatory:"true" json:"isOnboarded"`

	// This indicates if the log set feature is enabled for the tenancy
	IsLogSetEnabled *bool `mandatory:"false" json:"isLogSetEnabled"`

	// This indicates if data has ever been ingested for the tenancy in Logging Analytics
	IsDataEverIngested *bool `mandatory:"false" json:"isDataEverIngested"`
}

func (m Namespace) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Namespace) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
