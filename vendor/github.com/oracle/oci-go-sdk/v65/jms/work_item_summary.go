// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Java Management Service Fleets API
//
// The APIs for the Fleet Management (https://docs.oracle.com/en-us/iaas/jms/doc/fleet-management.html) feature of Java Management Service to monitor and manage the usage of Java in your enterprise. Use these APIs to manage fleets, configure managed instances to report to fleets, and gain insights into the Java workloads running on these instances by carrying out basic and advanced features.
//

package jms

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// WorkItemSummary Work item to complete a work request.
type WorkItemSummary struct {

	// The unique ID of ths work item.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the work request created this work item.
	WorkRequestId *string `mandatory:"true" json:"workRequestId"`

	InstallationSite *InstallationSite `mandatory:"true" json:"installationSite"`

	Details WorkItemDetails `mandatory:"true" json:"details"`

	// The status of the work item.
	Status WorkItemStatusEnum `mandatory:"true" json:"status"`

	// Number of times this work item is retried.
	RetryCount *int `mandatory:"true" json:"retryCount"`

	// The date and time the work item was last updated. (formatted according to RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339)).
	TimeLastUpdated *common.SDKTime `mandatory:"false" json:"timeLastUpdated"`
}

func (m WorkItemSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m WorkItemSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingWorkItemStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetWorkItemStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *WorkItemSummary) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		TimeLastUpdated  *common.SDKTime    `json:"timeLastUpdated"`
		Id               *string            `json:"id"`
		WorkRequestId    *string            `json:"workRequestId"`
		InstallationSite *InstallationSite  `json:"installationSite"`
		Details          workitemdetails    `json:"details"`
		Status           WorkItemStatusEnum `json:"status"`
		RetryCount       *int               `json:"retryCount"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.TimeLastUpdated = model.TimeLastUpdated

	m.Id = model.Id

	m.WorkRequestId = model.WorkRequestId

	m.InstallationSite = model.InstallationSite

	nn, e = model.Details.UnmarshalPolymorphicJSON(model.Details.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Details = nn.(WorkItemDetails)
	} else {
		m.Details = nil
	}

	m.Status = model.Status

	m.RetryCount = model.RetryCount

	return
}
