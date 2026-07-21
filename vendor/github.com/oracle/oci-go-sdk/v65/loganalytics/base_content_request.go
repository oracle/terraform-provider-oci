// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// BaseContentRequest The base class to use for the content generation with LoganAI
type BaseContentRequest interface {
}

type basecontentrequest struct {
	JsonData    []byte
	ContentKind string `json:"contentKind"`
}

// UnmarshalJSON unmarshals json
func (m *basecontentrequest) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerbasecontentrequest basecontentrequest
	s := struct {
		Model Unmarshalerbasecontentrequest
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ContentKind = s.Model.ContentKind

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *basecontentrequest) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ContentKind {
	case "DASHBOARD_CONTENT":
		mm := DashboardMetadataContentRequest{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ELIGIBLE_BUSINESS_INSIGHTS":
		mm := EligibleInsightsContentRequest{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for BaseContentRequest: %s.", m.ContentKind)
		return *m, nil
	}
}

func (m basecontentrequest) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m basecontentrequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
