// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// File Storage API
//
// Use the File Storage service API to manage file systems, mount targets, and snapshots.
// For more information, see Overview of File Storage (https://docs.cloud.oracle.com/iaas/Content/File/Concepts/filestorageoverview.htm).
//

package filestorage

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ShareSet The FSS SMB Share Set contains the SMB configuration data of a Mount Target.
// A set of file systems to provide SMB share through one mount target.
// Composed of zero or more share resources.
type ShareSet struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment that contains the share set.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the virtual cloud network (VCN) the share set is in.
	VcnId *string `mandatory:"true" json:"vcnId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the mount target the share set bind to.
	MountTargetId *string `mandatory:"true" json:"mountTargetId"`

	// A comment of the share set. It does not have to be unique, and it is changeable.
	// Avoid entering confidential information.
	// Example: `My share set`
	Comment *string `mandatory:"true" json:"comment"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the share set.
	Id *string `mandatory:"true" json:"id"`

	// A customer-provided DNS name. This name plays a critical role in
	// establishing the server's place in the customer SMB security hierarchy.
	// For example, if an SMB server has a DNS name of
	// register5.store34.california.usa.marks-hats.com, then this particular
	// server is part of the store34.california.usa.marks-hats.com security
	// domain which in turn is part of the california.usa.marks-hats.com which
	// in turn is part of the usa.marks-hats.com
	// which in turn is part of the marks-hats.com security domain.
	// Must be unique across all FQDNs in the subnet and comply
	// with RFC 952 (https://tools.ietf.org/html/rfc952)
	// and RFC 1123 (https://tools.ietf.org/html/rfc1123).
	CustomFqdn *string `mandatory:"true" json:"customFqdn"`

	// Every SMB server (i.e. each mount target) needs a Netbios name in
	// addition to its fqdn (fully qualified domain name). Normally,
	// the Netbios name is simply the hostname portion of the fqdn.
	// This doesn't work when multiple computers have the same hostname.
	// For example, a computer called orange.colors.com and a computer
	// called orange.fruit.org can interfere with each other if they both
	// use orange as their Netbios name. To avoid problems, at least one
	// computer can be configured to have a Netbios name that is
	// not its hostname.
	NetBiosName *string `mandatory:"true" json:"netBiosName"`

	// A read-only property for the connection status between the
	// mount target and the customer-provided domain controller which
	// is the domain based on the customFQDN.
	DomainConnectionStatus ShareSetDomainConnectionStatusEnum `mandatory:"true" json:"domainConnectionStatus"`

	// The current state of the share set.
	LifecycleState ShareSetLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time the share set was created, expressed
	// in RFC 3339 (https://tools.ietf.org/rfc/rfc3339) timestamp format.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// A read-only property that emits the status of the join domain operation
	// of the mount target share set to a domain controller for SMB access.
	JoinDomainResult *string `mandatory:"true" json:"joinDomainResult"`

	// The availability domain the share set is in. May be unset
	// as a blank or NULL value.
	// Example: `Uocm:PHX-AD-1`
	AvailabilityDomain *string `mandatory:"false" json:"availabilityDomain"`

	// Turn on this flag to allow unsigned SMB traffic.
	IsUnsignedTrafficAllowed *bool `mandatory:"false" json:"isUnsignedTrafficAllowed"`

	// Describes the mount target's policy on SMB encryption
	SmbEncryption ShareSetSmbEncryptionEnum `mandatory:"false" json:"smbEncryption,omitempty"`

	// A read-only field that's only populated when the join domain operation FAILED.
	JoinDomainErrorMessage *string `mandatory:"false" json:"joinDomainErrorMessage"`

	// The organizational unit (OU) is a container in an Active Directory that can
	// hold user accounts, service accounts, computer accounts, and other OUs and
	// this parameter specifies the OU that the mount target will join within the
	// AD domain. You can then assign administrators to specific OUs, and apply
	// group policy to enforce targeted configuration settings.
	OrganizationalUnit *string `mandatory:"false" json:"organizationalUnit"`
}

func (m ShareSet) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ShareSet) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingShareSetDomainConnectionStatusEnum(string(m.DomainConnectionStatus)); !ok && m.DomainConnectionStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DomainConnectionStatus: %s. Supported values are: %s.", m.DomainConnectionStatus, strings.Join(GetShareSetDomainConnectionStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingShareSetLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetShareSetLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingShareSetSmbEncryptionEnum(string(m.SmbEncryption)); !ok && m.SmbEncryption != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SmbEncryption: %s. Supported values are: %s.", m.SmbEncryption, strings.Join(GetShareSetSmbEncryptionEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ShareSetSmbEncryptionEnum Enum with underlying type: string
type ShareSetSmbEncryptionEnum string

// Set of constants representing the allowable values for ShareSetSmbEncryptionEnum
const (
	ShareSetSmbEncryptionNever     ShareSetSmbEncryptionEnum = "NEVER"
	ShareSetSmbEncryptionSupported ShareSetSmbEncryptionEnum = "SUPPORTED"
	ShareSetSmbEncryptionRequired  ShareSetSmbEncryptionEnum = "REQUIRED"
)

var mappingShareSetSmbEncryptionEnum = map[string]ShareSetSmbEncryptionEnum{
	"NEVER":     ShareSetSmbEncryptionNever,
	"SUPPORTED": ShareSetSmbEncryptionSupported,
	"REQUIRED":  ShareSetSmbEncryptionRequired,
}

var mappingShareSetSmbEncryptionEnumLowerCase = map[string]ShareSetSmbEncryptionEnum{
	"never":     ShareSetSmbEncryptionNever,
	"supported": ShareSetSmbEncryptionSupported,
	"required":  ShareSetSmbEncryptionRequired,
}

// GetShareSetSmbEncryptionEnumValues Enumerates the set of values for ShareSetSmbEncryptionEnum
func GetShareSetSmbEncryptionEnumValues() []ShareSetSmbEncryptionEnum {
	values := make([]ShareSetSmbEncryptionEnum, 0)
	for _, v := range mappingShareSetSmbEncryptionEnum {
		values = append(values, v)
	}
	return values
}

// GetShareSetSmbEncryptionEnumStringValues Enumerates the set of values in String for ShareSetSmbEncryptionEnum
func GetShareSetSmbEncryptionEnumStringValues() []string {
	return []string{
		"NEVER",
		"SUPPORTED",
		"REQUIRED",
	}
}

// GetMappingShareSetSmbEncryptionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingShareSetSmbEncryptionEnum(val string) (ShareSetSmbEncryptionEnum, bool) {
	enum, ok := mappingShareSetSmbEncryptionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ShareSetDomainConnectionStatusEnum Enum with underlying type: string
type ShareSetDomainConnectionStatusEnum string

// Set of constants representing the allowable values for ShareSetDomainConnectionStatusEnum
const (
	ShareSetDomainConnectionStatusSuccess                           ShareSetDomainConnectionStatusEnum = "SUCCESS"
	ShareSetDomainConnectionStatusFailed                            ShareSetDomainConnectionStatusEnum = "FAILED"
	ShareSetDomainConnectionStatusJoinDomainNeeded                  ShareSetDomainConnectionStatusEnum = "JOIN_DOMAIN_NEEDED"
	ShareSetDomainConnectionStatusJoinDomainStarted                 ShareSetDomainConnectionStatusEnum = "JOIN_DOMAIN_STARTED"
	ShareSetDomainConnectionStatusUnableToLocateDomainController    ShareSetDomainConnectionStatusEnum = "UNABLE_TO_LOCATE_DOMAIN_CONTROLLER"
	ShareSetDomainConnectionStatusUnableToConnectToDomainController ShareSetDomainConnectionStatusEnum = "UNABLE_TO_CONNECT_TO_DOMAIN_CONTROLLER"
	ShareSetDomainConnectionStatusUnableToLocateKerberosServer      ShareSetDomainConnectionStatusEnum = "UNABLE_TO_LOCATE_KERBEROS_SERVER"
	ShareSetDomainConnectionStatusUnableToConnectToKerberosServer   ShareSetDomainConnectionStatusEnum = "UNABLE_TO_CONNECT_TO_KERBEROS_SERVER"
	ShareSetDomainConnectionStatusBadCredential                     ShareSetDomainConnectionStatusEnum = "BAD_CREDENTIAL"
)

var mappingShareSetDomainConnectionStatusEnum = map[string]ShareSetDomainConnectionStatusEnum{
	"SUCCESS":                                ShareSetDomainConnectionStatusSuccess,
	"FAILED":                                 ShareSetDomainConnectionStatusFailed,
	"JOIN_DOMAIN_NEEDED":                     ShareSetDomainConnectionStatusJoinDomainNeeded,
	"JOIN_DOMAIN_STARTED":                    ShareSetDomainConnectionStatusJoinDomainStarted,
	"UNABLE_TO_LOCATE_DOMAIN_CONTROLLER":     ShareSetDomainConnectionStatusUnableToLocateDomainController,
	"UNABLE_TO_CONNECT_TO_DOMAIN_CONTROLLER": ShareSetDomainConnectionStatusUnableToConnectToDomainController,
	"UNABLE_TO_LOCATE_KERBEROS_SERVER":       ShareSetDomainConnectionStatusUnableToLocateKerberosServer,
	"UNABLE_TO_CONNECT_TO_KERBEROS_SERVER":   ShareSetDomainConnectionStatusUnableToConnectToKerberosServer,
	"BAD_CREDENTIAL":                         ShareSetDomainConnectionStatusBadCredential,
}

var mappingShareSetDomainConnectionStatusEnumLowerCase = map[string]ShareSetDomainConnectionStatusEnum{
	"success":                                ShareSetDomainConnectionStatusSuccess,
	"failed":                                 ShareSetDomainConnectionStatusFailed,
	"join_domain_needed":                     ShareSetDomainConnectionStatusJoinDomainNeeded,
	"join_domain_started":                    ShareSetDomainConnectionStatusJoinDomainStarted,
	"unable_to_locate_domain_controller":     ShareSetDomainConnectionStatusUnableToLocateDomainController,
	"unable_to_connect_to_domain_controller": ShareSetDomainConnectionStatusUnableToConnectToDomainController,
	"unable_to_locate_kerberos_server":       ShareSetDomainConnectionStatusUnableToLocateKerberosServer,
	"unable_to_connect_to_kerberos_server":   ShareSetDomainConnectionStatusUnableToConnectToKerberosServer,
	"bad_credential":                         ShareSetDomainConnectionStatusBadCredential,
}

// GetShareSetDomainConnectionStatusEnumValues Enumerates the set of values for ShareSetDomainConnectionStatusEnum
func GetShareSetDomainConnectionStatusEnumValues() []ShareSetDomainConnectionStatusEnum {
	values := make([]ShareSetDomainConnectionStatusEnum, 0)
	for _, v := range mappingShareSetDomainConnectionStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetShareSetDomainConnectionStatusEnumStringValues Enumerates the set of values in String for ShareSetDomainConnectionStatusEnum
func GetShareSetDomainConnectionStatusEnumStringValues() []string {
	return []string{
		"SUCCESS",
		"FAILED",
		"JOIN_DOMAIN_NEEDED",
		"JOIN_DOMAIN_STARTED",
		"UNABLE_TO_LOCATE_DOMAIN_CONTROLLER",
		"UNABLE_TO_CONNECT_TO_DOMAIN_CONTROLLER",
		"UNABLE_TO_LOCATE_KERBEROS_SERVER",
		"UNABLE_TO_CONNECT_TO_KERBEROS_SERVER",
		"BAD_CREDENTIAL",
	}
}

// GetMappingShareSetDomainConnectionStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingShareSetDomainConnectionStatusEnum(val string) (ShareSetDomainConnectionStatusEnum, bool) {
	enum, ok := mappingShareSetDomainConnectionStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ShareSetLifecycleStateEnum Enum with underlying type: string
type ShareSetLifecycleStateEnum string

// Set of constants representing the allowable values for ShareSetLifecycleStateEnum
const (
	ShareSetLifecycleStateCreating ShareSetLifecycleStateEnum = "CREATING"
	ShareSetLifecycleStateActive   ShareSetLifecycleStateEnum = "ACTIVE"
	ShareSetLifecycleStateDeleting ShareSetLifecycleStateEnum = "DELETING"
	ShareSetLifecycleStateDeleted  ShareSetLifecycleStateEnum = "DELETED"
)

var mappingShareSetLifecycleStateEnum = map[string]ShareSetLifecycleStateEnum{
	"CREATING": ShareSetLifecycleStateCreating,
	"ACTIVE":   ShareSetLifecycleStateActive,
	"DELETING": ShareSetLifecycleStateDeleting,
	"DELETED":  ShareSetLifecycleStateDeleted,
}

var mappingShareSetLifecycleStateEnumLowerCase = map[string]ShareSetLifecycleStateEnum{
	"creating": ShareSetLifecycleStateCreating,
	"active":   ShareSetLifecycleStateActive,
	"deleting": ShareSetLifecycleStateDeleting,
	"deleted":  ShareSetLifecycleStateDeleted,
}

// GetShareSetLifecycleStateEnumValues Enumerates the set of values for ShareSetLifecycleStateEnum
func GetShareSetLifecycleStateEnumValues() []ShareSetLifecycleStateEnum {
	values := make([]ShareSetLifecycleStateEnum, 0)
	for _, v := range mappingShareSetLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetShareSetLifecycleStateEnumStringValues Enumerates the set of values in String for ShareSetLifecycleStateEnum
func GetShareSetLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
	}
}

// GetMappingShareSetLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingShareSetLifecycleStateEnum(val string) (ShareSetLifecycleStateEnum, bool) {
	enum, ok := mappingShareSetLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
