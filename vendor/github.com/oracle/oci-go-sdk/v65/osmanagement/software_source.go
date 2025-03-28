// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OS Management API
//
// API for the OS Management service. Use these API operations for working
// with Managed instances and Managed instance groups.
//

package osmanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SoftwareSource A software source contains a collection of packages
type SoftwareSource struct {

	// OCID for the Software Source
	Id *string `mandatory:"true" json:"id"`

	// OCID for the Compartment
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// User friendly name for the software source
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Type of the Software Source
	RepoType *string `mandatory:"true" json:"repoType"`

	// URL for the repostiory
	Url *string `mandatory:"true" json:"url"`

	// Information specified by the user about the software source
	Description *string `mandatory:"false" json:"description"`

	// The architecture type supported by the Software Source
	ArchType ArchTypesEnum `mandatory:"false" json:"archType,omitempty"`

	// OCID for the parent software source, if there is one
	ParentId *string `mandatory:"false" json:"parentId"`

	// Display name the parent software source, if there is one
	ParentName *string `mandatory:"false" json:"parentName"`

	// The yum repository checksum type used by this software source
	ChecksumType ChecksumTypesEnum `mandatory:"false" json:"checksumType,omitempty"`

	// Name of the person maintaining this software source
	MaintainerName *string `mandatory:"false" json:"maintainerName"`

	// Email address of the person maintaining this software source
	MaintainerEmail *string `mandatory:"false" json:"maintainerEmail"`

	// Phone number of the person maintaining this software source
	MaintainerPhone *string `mandatory:"false" json:"maintainerPhone"`

	// URL of the GPG key for this software source
	GpgKeyUrl *string `mandatory:"false" json:"gpgKeyUrl"`

	// ID of the GPG key for this software source
	GpgKeyId *string `mandatory:"false" json:"gpgKeyId"`

	// Fingerprint of the GPG key for this software source
	GpgKeyFingerprint *string `mandatory:"false" json:"gpgKeyFingerprint"`

	// status of the software source.
	Status SoftwareSourceStatusEnum `mandatory:"false" json:"status,omitempty"`

	// The current state of the Software Source.
	LifecycleState LifecycleStatesEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// Number of packages
	Packages *int `mandatory:"false" json:"packages"`

	// list of the Managed Instances associated with this Software Sources
	AssociatedManagedInstances []Id `mandatory:"false" json:"associatedManagedInstances"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m SoftwareSource) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SoftwareSource) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingArchTypesEnum(string(m.ArchType)); !ok && m.ArchType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ArchType: %s. Supported values are: %s.", m.ArchType, strings.Join(GetArchTypesEnumStringValues(), ",")))
	}
	if _, ok := GetMappingChecksumTypesEnum(string(m.ChecksumType)); !ok && m.ChecksumType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ChecksumType: %s. Supported values are: %s.", m.ChecksumType, strings.Join(GetChecksumTypesEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSoftwareSourceStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetSoftwareSourceStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingLifecycleStatesEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetLifecycleStatesEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SoftwareSourceStatusEnum Enum with underlying type: string
type SoftwareSourceStatusEnum string

// Set of constants representing the allowable values for SoftwareSourceStatusEnum
const (
	SoftwareSourceStatusNormal      SoftwareSourceStatusEnum = "NORMAL"
	SoftwareSourceStatusUnreachable SoftwareSourceStatusEnum = "UNREACHABLE"
	SoftwareSourceStatusError       SoftwareSourceStatusEnum = "ERROR"
	SoftwareSourceStatusWarning     SoftwareSourceStatusEnum = "WARNING"
)

var mappingSoftwareSourceStatusEnum = map[string]SoftwareSourceStatusEnum{
	"NORMAL":      SoftwareSourceStatusNormal,
	"UNREACHABLE": SoftwareSourceStatusUnreachable,
	"ERROR":       SoftwareSourceStatusError,
	"WARNING":     SoftwareSourceStatusWarning,
}

var mappingSoftwareSourceStatusEnumLowerCase = map[string]SoftwareSourceStatusEnum{
	"normal":      SoftwareSourceStatusNormal,
	"unreachable": SoftwareSourceStatusUnreachable,
	"error":       SoftwareSourceStatusError,
	"warning":     SoftwareSourceStatusWarning,
}

// GetSoftwareSourceStatusEnumValues Enumerates the set of values for SoftwareSourceStatusEnum
func GetSoftwareSourceStatusEnumValues() []SoftwareSourceStatusEnum {
	values := make([]SoftwareSourceStatusEnum, 0)
	for _, v := range mappingSoftwareSourceStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetSoftwareSourceStatusEnumStringValues Enumerates the set of values in String for SoftwareSourceStatusEnum
func GetSoftwareSourceStatusEnumStringValues() []string {
	return []string{
		"NORMAL",
		"UNREACHABLE",
		"ERROR",
		"WARNING",
	}
}

// GetMappingSoftwareSourceStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSoftwareSourceStatusEnum(val string) (SoftwareSourceStatusEnum, bool) {
	enum, ok := mappingSoftwareSourceStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
