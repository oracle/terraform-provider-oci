// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Big Data Service API
//
// REST API for Oracle Big Data Service. Use this API to build, deploy, and manage fully elastic Big Data Service clusters. Build on Hadoop, Spark and Data Science distributions, which can be fully integrated with existing enterprise data in Oracle Database and Oracle applications.
//

package bds

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// BdsApiKeySummary The API key summary.
type BdsApiKeySummary struct {

	// Identifier of the user's API key.
	Id *string `mandatory:"true" json:"id"`

	// User friendly identifier used to uniquely differentiate between different API keys.
	// Only ASCII alphanumeric characters with no spaces allowed.
	KeyAlias *string `mandatory:"true" json:"keyAlias"`

	// The current status of the API key.
	LifecycleState BdsApiKeyLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The name of the region to establish the Object Storage endpoint which was set as part of key creation operation.
	// If no region was provided this will be set to be the same region where the cluster lives. Example us-phoenix-1 .
	DefaultRegion *string `mandatory:"true" json:"defaultRegion"`

	// The time the API key was created, shown as an RFC 3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`
}

func (m BdsApiKeySummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m BdsApiKeySummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingBdsApiKeyLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetBdsApiKeyLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
