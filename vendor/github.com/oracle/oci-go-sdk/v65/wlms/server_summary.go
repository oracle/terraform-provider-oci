// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// WebLogic Management Service API
//
// WebLogic Management Service is an OCI service that enables a unified view and management of WebLogic domains
// in Oracle Cloud Infrastructure. Features include on-demand patching of WebLogic domains, rollback of the
// last applied patch, discovery and management of WebLogic instances on a compute host.
//

package wlms

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ServerSummary The summary of servers.
type ServerSummary struct {

	// The name of the server.
	Name *string `mandatory:"true" json:"name"`

	// Whether or not the server is an admin node.
	IsAdmin *bool `mandatory:"true" json:"isAdmin"`

	// The status of the server.
	Status *string `mandatory:"true" json:"status"`

	// The unique identifier of the server.
	// **Note:** Not an OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	Id *string `mandatory:"false" json:"id"`

	// The restart order assigned to the server.
	RestartOrder *int `mandatory:"false" json:"restartOrder"`

	// The middleware path on the server.
	MiddlewarePath *string `mandatory:"false" json:"middlewarePath"`

	// The middleware type on the server.
	MiddlewareType *string `mandatory:"false" json:"middlewareType"`

	// The version of the WebLogic domain of the server
	WeblogicVersion *string `mandatory:"false" json:"weblogicVersion"`

	// The JDK path on the server.
	JdkPath *string `mandatory:"false" json:"jdkPath"`

	// The JDK version on the server.
	JdkVersion *string `mandatory:"false" json:"jdkVersion"`

	// The name of the WebLogic domain to which the server belongs.
	WlsDomainName *string `mandatory:"false" json:"wlsDomainName"`

	// The ID of the WebLogic domain to which the server belongs.
	WlsDomainId *string `mandatory:"false" json:"wlsDomainId"`

	// The path of the WebLogic domain to which the server belongs.
	WlsDomainPath *string `mandatory:"false" json:"wlsDomainPath"`

	// Whether or not the server has installed the latest patches.
	LatestPatchesStatus ServerSummaryLatestPatchesStatusEnum `mandatory:"false" json:"latestPatchesStatus,omitempty"`

	// The patch readiness status of the server.
	PatchReadinessStatus PatchReadinessStatusEnum `mandatory:"false" json:"patchReadinessStatus,omitempty"`

	// The name of the server.
	HostName *string `mandatory:"false" json:"hostName"`

	// The managed instance ID of the server.
	ManagedInstanceId *string `mandatory:"false" json:"managedInstanceId"`

	// The date and time the server was first reported (in RFC 3339 (https://tools.ietf.org/rfc/rfc3339) format).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The date and time the server was last reported (in RFC 3339 (https://tools.ietf.org/rfc/rfc3339) format).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`
}

func (m ServerSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ServerSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingServerSummaryLatestPatchesStatusEnum(string(m.LatestPatchesStatus)); !ok && m.LatestPatchesStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LatestPatchesStatus: %s. Supported values are: %s.", m.LatestPatchesStatus, strings.Join(GetServerSummaryLatestPatchesStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingPatchReadinessStatusEnum(string(m.PatchReadinessStatus)); !ok && m.PatchReadinessStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PatchReadinessStatus: %s. Supported values are: %s.", m.PatchReadinessStatus, strings.Join(GetPatchReadinessStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ServerSummaryLatestPatchesStatusEnum Enum with underlying type: string
type ServerSummaryLatestPatchesStatusEnum string

// Set of constants representing the allowable values for ServerSummaryLatestPatchesStatusEnum
const (
	ServerSummaryLatestPatchesStatusOnLatestPatches  ServerSummaryLatestPatchesStatusEnum = "ON_LATEST_PATCHES"
	ServerSummaryLatestPatchesStatusPatchingRequired ServerSummaryLatestPatchesStatusEnum = "PATCHING_REQUIRED"
	ServerSummaryLatestPatchesStatusUnknown          ServerSummaryLatestPatchesStatusEnum = "UNKNOWN"
)

var mappingServerSummaryLatestPatchesStatusEnum = map[string]ServerSummaryLatestPatchesStatusEnum{
	"ON_LATEST_PATCHES": ServerSummaryLatestPatchesStatusOnLatestPatches,
	"PATCHING_REQUIRED": ServerSummaryLatestPatchesStatusPatchingRequired,
	"UNKNOWN":           ServerSummaryLatestPatchesStatusUnknown,
}

var mappingServerSummaryLatestPatchesStatusEnumLowerCase = map[string]ServerSummaryLatestPatchesStatusEnum{
	"on_latest_patches": ServerSummaryLatestPatchesStatusOnLatestPatches,
	"patching_required": ServerSummaryLatestPatchesStatusPatchingRequired,
	"unknown":           ServerSummaryLatestPatchesStatusUnknown,
}

// GetServerSummaryLatestPatchesStatusEnumValues Enumerates the set of values for ServerSummaryLatestPatchesStatusEnum
func GetServerSummaryLatestPatchesStatusEnumValues() []ServerSummaryLatestPatchesStatusEnum {
	values := make([]ServerSummaryLatestPatchesStatusEnum, 0)
	for _, v := range mappingServerSummaryLatestPatchesStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetServerSummaryLatestPatchesStatusEnumStringValues Enumerates the set of values in String for ServerSummaryLatestPatchesStatusEnum
func GetServerSummaryLatestPatchesStatusEnumStringValues() []string {
	return []string{
		"ON_LATEST_PATCHES",
		"PATCHING_REQUIRED",
		"UNKNOWN",
	}
}

// GetMappingServerSummaryLatestPatchesStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingServerSummaryLatestPatchesStatusEnum(val string) (ServerSummaryLatestPatchesStatusEnum, bool) {
	enum, ok := mappingServerSummaryLatestPatchesStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
