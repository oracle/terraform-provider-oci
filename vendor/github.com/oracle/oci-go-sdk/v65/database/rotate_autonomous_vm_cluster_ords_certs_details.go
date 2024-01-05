// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// RotateAutonomousVmClusterOrdsCertsDetails The details for configuring the SSL certificates on Autonomous VM Cluster
type RotateAutonomousVmClusterOrdsCertsDetails struct {

	// Specify SYSTEM to use Oracle-managed certificates. Specify BYOC when you want to bring your own certificate.
	CertificateGenerationType RotateAutonomousVmClusterOrdsCertsDetailsCertificateGenerationTypeEnum `mandatory:"true" json:"certificateGenerationType"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the certificate to use.
	CertificateId *string `mandatory:"false" json:"certificateId"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the certificate authority.
	CertificateAuthorityId *string `mandatory:"false" json:"certificateAuthorityId"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the certificate bundle.
	CaBundleId *string `mandatory:"false" json:"caBundleId"`
}

func (m RotateAutonomousVmClusterOrdsCertsDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RotateAutonomousVmClusterOrdsCertsDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingRotateAutonomousVmClusterOrdsCertsDetailsCertificateGenerationTypeEnum(string(m.CertificateGenerationType)); !ok && m.CertificateGenerationType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CertificateGenerationType: %s. Supported values are: %s.", m.CertificateGenerationType, strings.Join(GetRotateAutonomousVmClusterOrdsCertsDetailsCertificateGenerationTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// RotateAutonomousVmClusterOrdsCertsDetailsCertificateGenerationTypeEnum Enum with underlying type: string
type RotateAutonomousVmClusterOrdsCertsDetailsCertificateGenerationTypeEnum string

// Set of constants representing the allowable values for RotateAutonomousVmClusterOrdsCertsDetailsCertificateGenerationTypeEnum
const (
	RotateAutonomousVmClusterOrdsCertsDetailsCertificateGenerationTypeSystem RotateAutonomousVmClusterOrdsCertsDetailsCertificateGenerationTypeEnum = "SYSTEM"
	RotateAutonomousVmClusterOrdsCertsDetailsCertificateGenerationTypeByoc   RotateAutonomousVmClusterOrdsCertsDetailsCertificateGenerationTypeEnum = "BYOC"
)

var mappingRotateAutonomousVmClusterOrdsCertsDetailsCertificateGenerationTypeEnum = map[string]RotateAutonomousVmClusterOrdsCertsDetailsCertificateGenerationTypeEnum{
	"SYSTEM": RotateAutonomousVmClusterOrdsCertsDetailsCertificateGenerationTypeSystem,
	"BYOC":   RotateAutonomousVmClusterOrdsCertsDetailsCertificateGenerationTypeByoc,
}

var mappingRotateAutonomousVmClusterOrdsCertsDetailsCertificateGenerationTypeEnumLowerCase = map[string]RotateAutonomousVmClusterOrdsCertsDetailsCertificateGenerationTypeEnum{
	"system": RotateAutonomousVmClusterOrdsCertsDetailsCertificateGenerationTypeSystem,
	"byoc":   RotateAutonomousVmClusterOrdsCertsDetailsCertificateGenerationTypeByoc,
}

// GetRotateAutonomousVmClusterOrdsCertsDetailsCertificateGenerationTypeEnumValues Enumerates the set of values for RotateAutonomousVmClusterOrdsCertsDetailsCertificateGenerationTypeEnum
func GetRotateAutonomousVmClusterOrdsCertsDetailsCertificateGenerationTypeEnumValues() []RotateAutonomousVmClusterOrdsCertsDetailsCertificateGenerationTypeEnum {
	values := make([]RotateAutonomousVmClusterOrdsCertsDetailsCertificateGenerationTypeEnum, 0)
	for _, v := range mappingRotateAutonomousVmClusterOrdsCertsDetailsCertificateGenerationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetRotateAutonomousVmClusterOrdsCertsDetailsCertificateGenerationTypeEnumStringValues Enumerates the set of values in String for RotateAutonomousVmClusterOrdsCertsDetailsCertificateGenerationTypeEnum
func GetRotateAutonomousVmClusterOrdsCertsDetailsCertificateGenerationTypeEnumStringValues() []string {
	return []string{
		"SYSTEM",
		"BYOC",
	}
}

// GetMappingRotateAutonomousVmClusterOrdsCertsDetailsCertificateGenerationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRotateAutonomousVmClusterOrdsCertsDetailsCertificateGenerationTypeEnum(val string) (RotateAutonomousVmClusterOrdsCertsDetailsCertificateGenerationTypeEnum, bool) {
	enum, ok := mappingRotateAutonomousVmClusterOrdsCertsDetailsCertificateGenerationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
