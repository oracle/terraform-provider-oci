// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Java Management Service Query API
//
// API for the Java Management Service. Use this API to view and manage Fleets.
//

package jms

import (
	"github.com/oracle/oci-go-sdk/v42/common"
)

// RequestSummarizedApplicationUsageDetails Parameters for filtering applications.
type RequestSummarizedApplicationUsageDetails struct {

	// The start of the time period during which resources are searched (formatted according to RFC3339). Defaults to current time minus seven days.
	TimeStart *common.SDKTime `mandatory:"false" json:"timeStart"`

	// The end of the time period during which resources are searched (formatted according to RFC3339). Defaults to current time.
	TimeEnd *common.SDKTime `mandatory:"false" json:"timeEnd"`

	// The display name of the application.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The installation path of the related installation.
	InstallationPath *string `mandatory:"false" json:"installationPath"`

	// The vendor of the related Java Runtime.
	JreVendor *string `mandatory:"false" json:"jreVendor"`

	// The distribution of the related Java Runtime.
	JreDistribution *string `mandatory:"false" json:"jreDistribution"`

	// The version of the related Java Runtime.
	JreVersion *string `mandatory:"false" json:"jreVersion"`

	// The ID of the application.
	ApplicationId *string `mandatory:"false" json:"applicationId"`

	// The way the application was started.
	ApplicationType *string `mandatory:"false" json:"applicationType"`

	// The ID of the related managed instance.
	ManagedInstanceId *string `mandatory:"false" json:"managedInstanceId"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder SortOrderEnum `mandatory:"false" json:"sortOrder,omitempty"`

	// The field to sort application views. Only one sort order may be provided.
	// Default order for _timeFirstSeen_, _timeLastSeen_, _approximateJreCount_, _approximateInstallationCount_
	// and _approximateManagedInstanceCount_  is **descending**.
	// Default order for _displayName_ is **ascending**.
	// If no value is specified _timeLastSeen_ is default.
	SortBy ApplicationSortByEnum `mandatory:"false" json:"sortBy,omitempty"`

	// Additional fields to include into the returned model on top of the required ones.
	// This parameter can also include 'approximateJreCount', 'approximateInstallationCount' and 'approximateManagedInstanceCount'.
	// For example 'approximateJreCount,approximateInstallationCount'.
	Fields []SummarizeApplicationUsageFieldsEnum `mandatory:"false" json:"fields"`
}

func (m RequestSummarizedApplicationUsageDetails) String() string {
	return common.PointerString(m)
}
