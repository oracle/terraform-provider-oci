// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// RotateAutonomousVmClusterSslCertsDetails Details for configuring the ORDS certificates on Autonomous Exadata VM Cluster
type RotateAutonomousVmClusterSslCertsDetails struct {

	// Specify SYSTEM to use Oracle-managed certificates. Specify BYOC when you want to bring your own certificate.
	CertificateGenerationType RotateAutonomousVmClusterSslCertsDetailsCertificateGenerationTypeEnum `mandatory:"true" json:"certificateGenerationType"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the certificate to use.
	CertificateId *string `mandatory:"false" json:"certificateId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the certificate authority.
	CertificateAuthorityId *string `mandatory:"false" json:"certificateAuthorityId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the certificate bundle.
	CaBundleId *string `mandatory:"false" json:"caBundleId"`
}

func (m RotateAutonomousVmClusterSslCertsDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RotateAutonomousVmClusterSslCertsDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingRotateAutonomousVmClusterSslCertsDetailsCertificateGenerationTypeEnum(string(m.CertificateGenerationType)); !ok && m.CertificateGenerationType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CertificateGenerationType: %s. Supported values are: %s.", m.CertificateGenerationType, strings.Join(GetRotateAutonomousVmClusterSslCertsDetailsCertificateGenerationTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// RotateAutonomousVmClusterSslCertsDetailsCertificateGenerationTypeEnum Enum with underlying type: string
type RotateAutonomousVmClusterSslCertsDetailsCertificateGenerationTypeEnum string

// Set of constants representing the allowable values for RotateAutonomousVmClusterSslCertsDetailsCertificateGenerationTypeEnum
const (
	RotateAutonomousVmClusterSslCertsDetailsCertificateGenerationTypeSystem RotateAutonomousVmClusterSslCertsDetailsCertificateGenerationTypeEnum = "SYSTEM"
	RotateAutonomousVmClusterSslCertsDetailsCertificateGenerationTypeByoc   RotateAutonomousVmClusterSslCertsDetailsCertificateGenerationTypeEnum = "BYOC"
)

var mappingRotateAutonomousVmClusterSslCertsDetailsCertificateGenerationTypeEnum = map[string]RotateAutonomousVmClusterSslCertsDetailsCertificateGenerationTypeEnum{
	"SYSTEM": RotateAutonomousVmClusterSslCertsDetailsCertificateGenerationTypeSystem,
	"BYOC":   RotateAutonomousVmClusterSslCertsDetailsCertificateGenerationTypeByoc,
}

var mappingRotateAutonomousVmClusterSslCertsDetailsCertificateGenerationTypeEnumLowerCase = map[string]RotateAutonomousVmClusterSslCertsDetailsCertificateGenerationTypeEnum{
	"system": RotateAutonomousVmClusterSslCertsDetailsCertificateGenerationTypeSystem,
	"byoc":   RotateAutonomousVmClusterSslCertsDetailsCertificateGenerationTypeByoc,
}

// GetRotateAutonomousVmClusterSslCertsDetailsCertificateGenerationTypeEnumValues Enumerates the set of values for RotateAutonomousVmClusterSslCertsDetailsCertificateGenerationTypeEnum
func GetRotateAutonomousVmClusterSslCertsDetailsCertificateGenerationTypeEnumValues() []RotateAutonomousVmClusterSslCertsDetailsCertificateGenerationTypeEnum {
	values := make([]RotateAutonomousVmClusterSslCertsDetailsCertificateGenerationTypeEnum, 0)
	for _, v := range mappingRotateAutonomousVmClusterSslCertsDetailsCertificateGenerationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetRotateAutonomousVmClusterSslCertsDetailsCertificateGenerationTypeEnumStringValues Enumerates the set of values in String for RotateAutonomousVmClusterSslCertsDetailsCertificateGenerationTypeEnum
func GetRotateAutonomousVmClusterSslCertsDetailsCertificateGenerationTypeEnumStringValues() []string {
	return []string{
		"SYSTEM",
		"BYOC",
	}
}

// GetMappingRotateAutonomousVmClusterSslCertsDetailsCertificateGenerationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRotateAutonomousVmClusterSslCertsDetailsCertificateGenerationTypeEnum(val string) (RotateAutonomousVmClusterSslCertsDetailsCertificateGenerationTypeEnum, bool) {
	enum, ok := mappingRotateAutonomousVmClusterSslCertsDetailsCertificateGenerationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
