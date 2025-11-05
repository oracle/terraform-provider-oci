// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DataIntelligences Control Plane API
//
// Use the DataIntelligences Control Plane API to manage dataIntelligences.
//

package dif

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AdbDetail Details to create an Oracle Autonomous Database.
type AdbDetail struct {

	// Id for the adw instance.
	InstanceId *string `mandatory:"true" json:"instanceId"`

	// DB Workload to be used with ADB. Accepted values are OLTP, DW.
	DbWorkload DbWorkloadEnum `mandatory:"true" json:"dbWorkload"`

	// The compute amount (ECPUs) available to the database.
	Ecpu *int `mandatory:"true" json:"ecpu"`

	// The size, in terabytes, of the data volume that will be created and attached to the database.
	DataStorageSizeInTBs *int `mandatory:"true" json:"dataStorageSizeInTBs"`

	// The OCI vault secret [/Content/General/Concepts/identifiers.htm]OCID for admin password.
	AdminPasswordId *string `mandatory:"true" json:"adminPasswordId"`

	// A valid Oracle Database version for Autonomous Database.
	DbVersion *string `mandatory:"true" json:"dbVersion"`

	// Specifies if the Autonomous Database requires mTLS connections.
	IsMtlsConnectionRequired *bool `mandatory:"false" json:"isMtlsConnectionRequired"`

	// The OCID of the subnet the Autonomous Database is associated with.
	SubnetId *string `mandatory:"false" json:"subnetId"`

	// This is an array of CIDR (classless inter-domain routing) notations for a subnet or VCN OCID (virtual cloud network Oracle Cloud ID). Allowed only when subnetId is provided (private ADB).
	ToolsPublicAccess *string `mandatory:"false" json:"toolsPublicAccess"`

	// If true then subnetId should not be provided.
	IsPublic *bool `mandatory:"false" json:"isPublic"`
}

func (m AdbDetail) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AdbDetail) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDbWorkloadEnum(string(m.DbWorkload)); !ok && m.DbWorkload != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DbWorkload: %s. Supported values are: %s.", m.DbWorkload, strings.Join(GetDbWorkloadEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
