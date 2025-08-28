// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OCI Cache API
//
// Use the OCI Cache API to create and manage clusters. A cluster is a memory-based storage solution. For more information, see OCI Cache (https://docs.oracle.com/iaas/Content/ocicache/home.htm).
//

package redis

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// OciCacheDefaultConfigSetSummary Summary of information about an OCI Cache Default Config Set.
type OciCacheDefaultConfigSetSummary struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm#Oracle) of the OCI Cache Default Config Set.
	Id *string `mandatory:"true" json:"id"`

	// The engine version of the OCI Cache Default Config Set.
	SoftwareVersion OciCacheConfigSetSoftwareVersionEnum `mandatory:"true" json:"softwareVersion"`

	// A user-friendly name of the OCI Cache Default Config Set.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The current state of the OCI Cache Default Config Set.
	LifecycleState OciCacheDefaultConfigSetLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// The date and time the configuration was created. An RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339) formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`
}

func (m OciCacheDefaultConfigSetSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OciCacheDefaultConfigSetSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingOciCacheConfigSetSoftwareVersionEnum(string(m.SoftwareVersion)); !ok && m.SoftwareVersion != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SoftwareVersion: %s. Supported values are: %s.", m.SoftwareVersion, strings.Join(GetOciCacheConfigSetSoftwareVersionEnumStringValues(), ",")))
	}

	if _, ok := GetMappingOciCacheDefaultConfigSetLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetOciCacheDefaultConfigSetLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
