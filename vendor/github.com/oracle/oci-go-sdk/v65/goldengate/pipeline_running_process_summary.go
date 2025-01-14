// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// GoldenGate API
//
// Use the Oracle Cloud Infrastructure GoldenGate APIs to perform data replication operations.
//

package goldengate

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// PipelineRunningProcessSummary Each replication process and their summary details.
type PipelineRunningProcessSummary struct {

	// An object's Display Name.
	Name *string `mandatory:"true" json:"name"`

	// The type of process running in a replication. For example, Extract or Replicat. This option applies when retrieving running processes.
	ProcessType ProcessTypeEnum `mandatory:"true" json:"processType"`

	// The status of the Extract or Replicat process. This option applies when retrieving running processes.
	Status ProcessStatusTypeEnum `mandatory:"true" json:"status"`

	// The latency, in seconds, of a process running in a replication. This option applies when retrieving running processes.
	LastRecordLagInSeconds *float32 `mandatory:"true" json:"lastRecordLagInSeconds"`

	// The date and time the last record was processed by an Extract or Replicat. This option applies when retrieving running processes.
	// The format is defined by RFC3339 (https://tools.ietf.org/html/rfc3339), such as `2024-07-25T21:10:29.600Z`.
	TimeLastProcessed *common.SDKTime `mandatory:"true" json:"timeLastProcessed"`
}

func (m PipelineRunningProcessSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PipelineRunningProcessSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingProcessTypeEnum(string(m.ProcessType)); !ok && m.ProcessType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ProcessType: %s. Supported values are: %s.", m.ProcessType, strings.Join(GetProcessTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingProcessStatusTypeEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetProcessStatusTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
