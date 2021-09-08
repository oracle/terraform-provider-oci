// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Java Management Service Query API
//
// API for the Java Management Service. Use this API to view and manage Fleets.
//

package jms

import (
	"github.com/oracle/oci-go-sdk/v47/common"
)

// RequestSummarizedJreUsageDetails Parameters for filtering Java Runtime Usages.
type RequestSummarizedJreUsageDetails struct {

	// The start of the time period during which resources are searched (formatted according to RFC3339). Defaults to current time minus seven days.
	TimeStart *common.SDKTime `mandatory:"false" json:"timeStart"`

	// The end of the time period during which resources are searched (formatted according to RFC3339). Defaults to current time.
	TimeEnd *common.SDKTime `mandatory:"false" json:"timeEnd"`

	// The vendor of the Java Runtime.
	JreVendor *string `mandatory:"false" json:"jreVendor"`

	// The distribution of the Java Runtime.
	JreDistribution *string `mandatory:"false" json:"jreDistribution"`

	// The version of the Java Runtime.
	JreVersion *string `mandatory:"false" json:"jreVersion"`

	// The ID of the related application.
	ApplicationId *string `mandatory:"false" json:"applicationId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the related managed instance.
	ManagedInstanceId *string `mandatory:"false" json:"managedInstanceId"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder SortOrderEnum `mandatory:"false" json:"sortOrder,omitempty"`

	// The field to sort Java Runtime views. Only one sort order may be provided.
	// Default order for _timeFirstSeen_, _timeLastSeen_, _version_, _approximateInstallationCount_,
	// _approximateApplicationCount_ and _approximateManagedInstanceCount_  is **descending**.
	// Default order for _distribution_ and _vendor_ is **ascending**. If no value is specified _timeLastSeen_ is default.
	SortBy JreSortByEnum `mandatory:"false" json:"sortBy,omitempty"`

	// Additional fields to include into the returned model on top of the required ones.
	// This parameter can also include 'approximateApplicationCount', 'approximateInstallationCount' and 'approximateManagedInstanceCount'.
	// For example 'approximateApplicationCount,approximateManagedInstanceCount'.
	Fields []SummarizeJreUsageFieldsEnum `mandatory:"false" json:"fields"`
}

func (m RequestSummarizedJreUsageDetails) String() string {
	return common.PointerString(m)
}
