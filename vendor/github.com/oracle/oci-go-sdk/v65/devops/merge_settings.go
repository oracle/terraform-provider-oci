// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DevOps API
//
// Use the DevOps API to create DevOps projects, configure code repositories,  add artifacts to deploy, build and test software applications, configure  target deployment environments, and deploy software applications.  For more information, see DevOps (https://docs.oracle.com/iaas/Content/devops/using/home.htm).
//

package devops

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// MergeSettings Enabled and disabled merge strategies for a project or repository, also contains a default strategy.
type MergeSettings struct {

	// Default type of merge strategy associated with the a Project or Repository.
	DefaultMergeStrategy MergeStrategyEnum `mandatory:"true" json:"defaultMergeStrategy"`

	// List of merge strategies which are allowed for a Project or Repository.
	AllowedMergeStrategies []MergeStrategyEnum `mandatory:"true" json:"allowedMergeStrategies"`
}

func (m MergeSettings) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MergeSettings) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingMergeStrategyEnum(string(m.DefaultMergeStrategy)); !ok && m.DefaultMergeStrategy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DefaultMergeStrategy: %s. Supported values are: %s.", m.DefaultMergeStrategy, strings.Join(GetMergeStrategyEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
