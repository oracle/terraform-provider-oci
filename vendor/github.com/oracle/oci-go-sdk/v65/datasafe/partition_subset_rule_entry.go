// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// PartitionSubsetRuleEntry Defines a subsetting rule based on partitions and sub-partitions to filter rows
type PartitionSubsetRuleEntry struct {

	// A list of partition names which are to be part of the subset data
	PartitionsList []string `mandatory:"false" json:"partitionsList"`

	// A list of sub-partition names which are to be part of the subset data. The sub-partition names should have the partition name also, separated by a dot
	SubPartitionsList []string `mandatory:"false" json:"subPartitionsList"`
}

func (m PartitionSubsetRuleEntry) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PartitionSubsetRuleEntry) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m PartitionSubsetRuleEntry) MarshalJSON() (buff []byte, e error) {
	type MarshalTypePartitionSubsetRuleEntry PartitionSubsetRuleEntry
	s := struct {
		DiscriminatorParam string `json:"ruleType"`
		MarshalTypePartitionSubsetRuleEntry
	}{
		"PARTITION",
		(MarshalTypePartitionSubsetRuleEntry)(m),
	}

	return json.Marshal(&s)
}
