// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Logging Management API
//
// Use the Logging Management API to create, read, list, update, move and delete
// log groups, log objects, log saved searches, and agent configurations.
// For more information, see Logging Overview (https://docs.oracle.com/iaas/Content/Logging/Concepts/loggingoverview.htm).
//

package logging

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// LogPipelineDestinationResponse Response for Log Pipeline destinations.
type LogPipelineDestinationResponse interface {

	// Name of Log Pipeline destination.
	GetName() *string
}

type logpipelinedestinationresponse struct {
	JsonData                []byte
	Name                    *string `mandatory:"false" json:"name"`
	PipelineDestinationType string  `json:"pipelineDestinationType"`
}

// UnmarshalJSON unmarshals json
func (m *logpipelinedestinationresponse) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerlogpipelinedestinationresponse logpipelinedestinationresponse
	s := struct {
		Model Unmarshalerlogpipelinedestinationresponse
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Name = s.Model.Name
	m.PipelineDestinationType = s.Model.PipelineDestinationType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *logpipelinedestinationresponse) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.PipelineDestinationType {
	case "OBJECT_STORAGE":
		mm := LogPipelineDestinationObjectStorageResponse{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "LOGGING_SEARCH":
		mm := LogPipelineDestinationLoggingSearchResponse{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "STREAMING_OSS":
		mm := LogPipelineDestinationStreamingOssResponse{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "MANAGED_KAFKA":
		mm := LogPipelineDestinationManagedKafkaResponse{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "LOGGING":
		mm := LogPipelineDestinationLoggingResponse{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for LogPipelineDestinationResponse: %s.", m.PipelineDestinationType)
		return *m, nil
	}
}

// GetName returns Name
func (m logpipelinedestinationresponse) GetName() *string {
	return m.Name
}

func (m logpipelinedestinationresponse) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m logpipelinedestinationresponse) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
