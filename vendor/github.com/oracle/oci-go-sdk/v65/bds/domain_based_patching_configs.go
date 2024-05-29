// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Big Data Service API
//
// REST API for Oracle Big Data Service. Use this API to build, deploy, and manage fully elastic Big Data Service clusters. Build on Hadoop, Spark and Data Science distributions, which can be fully integrated with existing enterprise data in Oracle Database and Oracle applications.
//

package bds

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DomainBasedPatchingConfigs This patching config will patch all nodes in one AD/FD at a time.
type DomainBasedPatchingConfigs struct {

	// The wait time between AD/FD in seconds.
	WaitTimeBetweenDomainInSeconds *int `mandatory:"true" json:"waitTimeBetweenDomainInSeconds"`

	// Acceptable number of failed-to-be-patched nodes in each domain. The maximum number of failed-to-patch nodes cannot exceed 20% of the number of nodes.
	ToleranceThresholdPerDomain *int `mandatory:"true" json:"toleranceThresholdPerDomain"`
}

func (m DomainBasedPatchingConfigs) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DomainBasedPatchingConfigs) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m DomainBasedPatchingConfigs) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDomainBasedPatchingConfigs DomainBasedPatchingConfigs
	s := struct {
		DiscriminatorParam string `json:"patchingConfigStrategy"`
		MarshalTypeDomainBasedPatchingConfigs
	}{
		"DOMAIN_BASED",
		(MarshalTypeDomainBasedPatchingConfigs)(m),
	}

	return json.Marshal(&s)
}
