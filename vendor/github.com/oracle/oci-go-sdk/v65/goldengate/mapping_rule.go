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

// MappingRule Mapping rule for source and target schemas for the pipeline data replication.
// For example:
// "{mappingType: INCLUDE, source: HR.*, target: HR.*}" for rule "Include HR.*" which will include all the tables under HR schema.
type MappingRule struct {

	// Defines the exclude/include rules of source and target schemas and tables when replicating from source to target. This option applies when creating and updating a pipeline.
	MappingType MappingTypeEnum `mandatory:"true" json:"mappingType"`

	// The source schema/table combination for replication to target.
	Source *string `mandatory:"true" json:"source"`

	// The target schema/table combination for replication from the source.
	Target *string `mandatory:"false" json:"target"`
}

func (m MappingRule) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MappingRule) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingMappingTypeEnum(string(m.MappingType)); !ok && m.MappingType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for MappingType: %s. Supported values are: %s.", m.MappingType, strings.Join(GetMappingTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
