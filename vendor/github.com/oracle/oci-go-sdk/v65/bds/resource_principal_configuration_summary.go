// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// ResourcePrincipalConfigurationSummary Resource Principal Session Token Details.
type ResourcePrincipalConfigurationSummary struct {

	// the ID of the ResourcePrincipalConfiguration.
	Id *string `mandatory:"true" json:"id"`

	// the OCID of the bdsInstance which is the parent resource id.
	BdsInstanceId *string `mandatory:"true" json:"bdsInstanceId"`

	// a user-friendly name. only ascii alphanumeric characters with no spaces allowed. the name does not have to be unique, and it may be changed. avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// the state of the ResourcePrincipalConfiguration.
	LifecycleState ResourcePrincipalConfigurationLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// the time the ResourcePrincipalConfiguration was created, shown as an rfc 3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// the time the ResourcePrincipalConfiguration was updated, shown as an rfc 3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// the time the resource principal session token was refreshed, shown as an rfc 3339 formatted datetime string.
	TimeTokenRefreshed *common.SDKTime `mandatory:"false" json:"timeTokenRefreshed"`

	// the time the resource principal session token will expired, shown as an rfc 3339 formatted datetime string.
	TimeTokenExpiry *common.SDKTime `mandatory:"false" json:"timeTokenExpiry"`
}

func (m ResourcePrincipalConfigurationSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ResourcePrincipalConfigurationSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingResourcePrincipalConfigurationLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetResourcePrincipalConfigurationLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
