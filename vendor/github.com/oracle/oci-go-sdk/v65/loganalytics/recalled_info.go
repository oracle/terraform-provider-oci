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

// RecalledInfo This is the information about data recalled
type RecalledInfo struct {

	// This is the status of the recall
	Status RecallStatusEnum `mandatory:"true" json:"status"`

	// This is the purpose of the recall
	Purpose *string `mandatory:"true" json:"purpose"`

	// This is the query associated with the recall
	QueryString *string `mandatory:"true" json:"queryString"`

	// This is the list of logsets associated with the recall
	LogSets *string `mandatory:"true" json:"logSets"`

	// This is the id for the recalled data collection
	CollectionId *int64 `mandatory:"true" json:"collectionId"`

	// This is the recalled date start time
	TimeRecalledDataStarted *common.SDKTime `mandatory:"true" json:"timeRecalledDataStarted"`

	// This is the recalled data end time
	TimeRecalledDataEnded *common.SDKTime `mandatory:"true" json:"timeRecalledDataEnded"`
}

func (m RecalledInfo) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RecalledInfo) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingRecallStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetRecallStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
