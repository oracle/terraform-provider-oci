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

// ConfigWorkRequestResult The result of the submitted config work request.
type ConfigWorkRequestResult interface {
}

type configworkrequestresult struct {
	JsonData        []byte
	WorkRequestKind string `json:"workRequestKind"`
}

// UnmarshalJSON unmarshals json
func (m *configworkrequestresult) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerconfigworkrequestresult configworkrequestresult
	s := struct {
		Model Unmarshalerconfigworkrequestresult
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.WorkRequestKind = s.Model.WorkRequestKind

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *configworkrequestresult) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.WorkRequestKind {
	case "GENERATE_AI_CONTENT":
		mm := ConfigWorkRequestGenAiContentResult{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for ConfigWorkRequestResult: %s.", m.WorkRequestKind)
		return *m, nil
	}
}

func (m configworkrequestresult) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m configworkrequestresult) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
