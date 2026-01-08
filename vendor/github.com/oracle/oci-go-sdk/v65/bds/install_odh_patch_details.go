// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Big Data Service API
//
// REST API for Oracle Big Data Service. Use this API to build, deploy, and manage fully elastic Big Data Service clusters. Build on Hadoop, Spark and Data Science distributions, which can be fully integrated with existing enterprise data in Oracle Database and Oracle applications.
//

package bds

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// InstallOdhPatchDetails The reqeust body while installing a ODH patch to a cluster.
type InstallOdhPatchDetails struct {

	// The version of the ODH patch to be installed.
	Version *string `mandatory:"true" json:"version"`

	// The display name of the ODH patch to be installed.
	OdhPatchName *string `mandatory:"true" json:"odhPatchName"`

	// The patch Url of the ODH patch to be installed.
	PaUrl *string `mandatory:"true" json:"paUrl"`

	// The md5Hash of the ODH patch to be installed.
	Md5Hash *string `mandatory:"true" json:"md5Hash"`

	// The algorithm for the checkSum used for the ODH patch.
	CheckSumAlgo InstallOdhPatchDetailsCheckSumAlgoEnum `mandatory:"false" json:"checkSumAlgo,omitempty"`

	// The checkSum of the ODH patch to be installed.
	CheckSum *string `mandatory:"false" json:"checkSum"`

	// Cluster Admin Password
	ClusterAdminPassword *string `mandatory:"false" json:"clusterAdminPassword"`

	// The secretId for the clusterAdminPassword.
	SecretId *string `mandatory:"false" json:"secretId"`

	// The flag to check if the ODH patch can be installed.
	IsCompatibilityCheck *bool `mandatory:"false" json:"isCompatibilityCheck"`

	PatchingConfig OdhPatchingConfig `mandatory:"false" json:"patchingConfig"`
}

func (m InstallOdhPatchDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m InstallOdhPatchDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingInstallOdhPatchDetailsCheckSumAlgoEnum(string(m.CheckSumAlgo)); !ok && m.CheckSumAlgo != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CheckSumAlgo: %s. Supported values are: %s.", m.CheckSumAlgo, strings.Join(GetInstallOdhPatchDetailsCheckSumAlgoEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *InstallOdhPatchDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		CheckSumAlgo         InstallOdhPatchDetailsCheckSumAlgoEnum `json:"checkSumAlgo"`
		CheckSum             *string                                `json:"checkSum"`
		ClusterAdminPassword *string                                `json:"clusterAdminPassword"`
		SecretId             *string                                `json:"secretId"`
		IsCompatibilityCheck *bool                                  `json:"isCompatibilityCheck"`
		PatchingConfig       odhpatchingconfig                      `json:"patchingConfig"`
		Version              *string                                `json:"version"`
		OdhPatchName         *string                                `json:"odhPatchName"`
		PaUrl                *string                                `json:"paUrl"`
		Md5Hash              *string                                `json:"md5Hash"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.CheckSumAlgo = model.CheckSumAlgo

	m.CheckSum = model.CheckSum

	m.ClusterAdminPassword = model.ClusterAdminPassword

	m.SecretId = model.SecretId

	m.IsCompatibilityCheck = model.IsCompatibilityCheck

	nn, e = model.PatchingConfig.UnmarshalPolymorphicJSON(model.PatchingConfig.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.PatchingConfig = nn.(OdhPatchingConfig)
	} else {
		m.PatchingConfig = nil
	}

	m.Version = model.Version

	m.OdhPatchName = model.OdhPatchName

	m.PaUrl = model.PaUrl

	m.Md5Hash = model.Md5Hash

	return
}

// InstallOdhPatchDetailsCheckSumAlgoEnum Enum with underlying type: string
type InstallOdhPatchDetailsCheckSumAlgoEnum string

// Set of constants representing the allowable values for InstallOdhPatchDetailsCheckSumAlgoEnum
const (
	InstallOdhPatchDetailsCheckSumAlgoSha256 InstallOdhPatchDetailsCheckSumAlgoEnum = "SHA256"
)

var mappingInstallOdhPatchDetailsCheckSumAlgoEnum = map[string]InstallOdhPatchDetailsCheckSumAlgoEnum{
	"SHA256": InstallOdhPatchDetailsCheckSumAlgoSha256,
}

var mappingInstallOdhPatchDetailsCheckSumAlgoEnumLowerCase = map[string]InstallOdhPatchDetailsCheckSumAlgoEnum{
	"sha256": InstallOdhPatchDetailsCheckSumAlgoSha256,
}

// GetInstallOdhPatchDetailsCheckSumAlgoEnumValues Enumerates the set of values for InstallOdhPatchDetailsCheckSumAlgoEnum
func GetInstallOdhPatchDetailsCheckSumAlgoEnumValues() []InstallOdhPatchDetailsCheckSumAlgoEnum {
	values := make([]InstallOdhPatchDetailsCheckSumAlgoEnum, 0)
	for _, v := range mappingInstallOdhPatchDetailsCheckSumAlgoEnum {
		values = append(values, v)
	}
	return values
}

// GetInstallOdhPatchDetailsCheckSumAlgoEnumStringValues Enumerates the set of values in String for InstallOdhPatchDetailsCheckSumAlgoEnum
func GetInstallOdhPatchDetailsCheckSumAlgoEnumStringValues() []string {
	return []string{
		"SHA256",
	}
}

// GetMappingInstallOdhPatchDetailsCheckSumAlgoEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingInstallOdhPatchDetailsCheckSumAlgoEnum(val string) (InstallOdhPatchDetailsCheckSumAlgoEnum, bool) {
	enum, ok := mappingInstallOdhPatchDetailsCheckSumAlgoEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
