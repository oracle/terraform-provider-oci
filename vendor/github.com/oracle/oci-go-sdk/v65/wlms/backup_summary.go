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

// BackupSummary The information of a backup for the server.
type BackupSummary struct {

	// The unique identifier of the backup.
	// **Note:** Not an OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	Id *string `mandatory:"true" json:"id"`

	// The type of the backup.
	Type BackupTypeEnum `mandatory:"true" json:"type"`

	// The location of the backup. For backups of type LOCAL_FILE this is the absolute path of the backup file.
	BackupLocation *string `mandatory:"true" json:"backupLocation"`

	// The managed instance ID of the server for which the backup was created.
	ManagedInstanceId *string `mandatory:"false" json:"managedInstanceId"`

	// The type of content of the backup.
	ContentType BackupContentTypeEnum `mandatory:"false" json:"contentType,omitempty"`

	// The date and time when the backup was created (in RFC 3339 (https://tools.ietf.org/rfc/rfc3339) format).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`
}

func (m BackupSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m BackupSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingBackupTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetBackupTypeEnumStringValues(), ",")))
	}

	if _, ok := GetMappingBackupContentTypeEnum(string(m.ContentType)); !ok && m.ContentType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ContentType: %s. Supported values are: %s.", m.ContentType, strings.Join(GetBackupContentTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
