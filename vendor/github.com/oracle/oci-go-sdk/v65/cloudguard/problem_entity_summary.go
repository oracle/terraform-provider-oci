// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Guard and Security Zones API
//
// Use the Cloud Guard and Security Zones API to automate processes that you would otherwise perform through the Cloud Guard Console or the Security Zones Console. For more information on these services, see the Cloud Guard (https://docs.cloud.oracle.com/iaas/cloud-guard/home.htm) and Security Zones (https://docs.cloud.oracle.com/iaas/security-zone/home.htm) documentation.
// **Note:** For Cloud Guard, you can perform Create, Update, and Delete operations only from the reporting region of your Cloud Guard tenancy. You can perform Read operations from any region.
//

package cloudguard

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ProblemEntitySummary The information about problem entities details of DataSource for a CloudGuard Problem.
type ProblemEntitySummary struct {

	// Data source problem entities region
	Regions []string `mandatory:"true" json:"regions"`

	// Data source problem entities first detected time
	TimeFirstDetected *common.SDKTime `mandatory:"true" json:"timeFirstDetected"`

	// Attached problem id
	ProblemId *string `mandatory:"true" json:"problemId"`

	// Data source problem entities last detected time
	TimeLastDetected *common.SDKTime `mandatory:"true" json:"timeLastDetected"`

	// Log result query url for a data source query
	ResultUrl *string `mandatory:"false" json:"resultUrl"`

	// List of event related to a DataSource
	EntityDetails []EntityDetails `mandatory:"false" json:"entityDetails"`
}

func (m ProblemEntitySummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ProblemEntitySummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
