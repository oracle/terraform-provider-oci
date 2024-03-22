// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Logging Management API
//
// Use the Logging Management API to create, read, list, update, move and delete
// log groups, log objects, log saved searches, and agent configurations.
// For more information, see Logging Overview (https://docs.cloud.oracle.com/iaas/Content/Logging/Concepts/loggingoverview.htm).
//

package logging

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UnifiedAgentLoggingRecordTransformerFilter Logging record transformer filter object mutates/transforms logs.
// Ref: https://docs.fluentd.org/filter/record_transformer
type UnifiedAgentLoggingRecordTransformerFilter struct {

	// Unique name for the filter.
	Name *string `mandatory:"true" json:"name"`

	// Add new key-value pairs in logs
	RecordList []RecordTransformerPair `mandatory:"true" json:"recordList"`

	// When set to true, the full Ruby syntax is enabled in the ${} expression.
	IsRubyEnabled *bool `mandatory:"false" json:"isRubyEnabled"`

	// If true, automatically casts the field types.
	IsAutoTypecastEnabled *bool `mandatory:"false" json:"isAutoTypecastEnabled"`

	// If true, it modifies a new empty hash
	IsRenewRecordEnabled *bool `mandatory:"false" json:"isRenewRecordEnabled"`

	// Overwrites the time of logs with this value, this value must be a Unix timestamp.
	RenewTimeKey *string `mandatory:"false" json:"renewTimeKey"`

	// A list of keys to keep. Only relevant if isRenewRecordEnabled is set to true
	KeepKeys []string `mandatory:"false" json:"keepKeys"`

	// A list of keys to delete
	RemoveKeys []string `mandatory:"false" json:"removeKeys"`
}

// GetName returns Name
func (m UnifiedAgentLoggingRecordTransformerFilter) GetName() *string {
	return m.Name
}

func (m UnifiedAgentLoggingRecordTransformerFilter) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UnifiedAgentLoggingRecordTransformerFilter) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m UnifiedAgentLoggingRecordTransformerFilter) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUnifiedAgentLoggingRecordTransformerFilter UnifiedAgentLoggingRecordTransformerFilter
	s := struct {
		DiscriminatorParam string `json:"filterType"`
		MarshalTypeUnifiedAgentLoggingRecordTransformerFilter
	}{
		"RECORD_TRANSFORMER_FILTER",
		(MarshalTypeUnifiedAgentLoggingRecordTransformerFilter)(m),
	}

	return json.Marshal(&s)
}
