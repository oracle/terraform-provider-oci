// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Application Dependency Management API
//
// Use the Application Dependency Management API to create knowledge bases and vulnerability audits.  For more information, see ADM (https://docs.cloud.oracle.com/Content/application-dependency-management/home.htm).
//

package adm

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ScmConfiguration A configuration for the Source Code Management tool/platform used by a remediation recipe.
type ScmConfiguration interface {

	// The branch used by ADM to patch vulnerabilities.
	GetBranch() *string

	// If true, the Pull Request (PR) will be merged after the verify stage completes successfully
	// If false, the PR with the proposed changes must be reviewed and manually merged.
	GetIsAutomergeEnabled() *bool

	// The location of the build file relative to the root of the repository. Only Maven build files (POM) are currently supported.
	// If this property is not specified, ADM will use the build file located at the root of the repository.
	GetBuildFileLocation() *string
}

type scmconfiguration struct {
	JsonData           []byte
	BuildFileLocation  *string `mandatory:"false" json:"buildFileLocation"`
	Branch             *string `mandatory:"true" json:"branch"`
	IsAutomergeEnabled *bool   `mandatory:"true" json:"isAutomergeEnabled"`
	ScmType            string  `json:"scmType"`
}

// UnmarshalJSON unmarshals json
func (m *scmconfiguration) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerscmconfiguration scmconfiguration
	s := struct {
		Model Unmarshalerscmconfiguration
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Branch = s.Model.Branch
	m.IsAutomergeEnabled = s.Model.IsAutomergeEnabled
	m.BuildFileLocation = s.Model.BuildFileLocation
	m.ScmType = s.Model.ScmType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *scmconfiguration) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ScmType {
	case "OCI_CODE_REPOSITORY":
		mm := OciCodeRepositoryConfiguration{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "EXTERNAL_SCM":
		mm := ExternalScmConfiguration{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for ScmConfiguration: %s.", m.ScmType)
		return *m, nil
	}
}

// GetBuildFileLocation returns BuildFileLocation
func (m scmconfiguration) GetBuildFileLocation() *string {
	return m.BuildFileLocation
}

// GetBranch returns Branch
func (m scmconfiguration) GetBranch() *string {
	return m.Branch
}

// GetIsAutomergeEnabled returns IsAutomergeEnabled
func (m scmconfiguration) GetIsAutomergeEnabled() *bool {
	return m.IsAutomergeEnabled
}

func (m scmconfiguration) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m scmconfiguration) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ScmConfigurationScmTypeEnum Enum with underlying type: string
type ScmConfigurationScmTypeEnum string

// Set of constants representing the allowable values for ScmConfigurationScmTypeEnum
const (
	ScmConfigurationScmTypeOciCodeRepository ScmConfigurationScmTypeEnum = "OCI_CODE_REPOSITORY"
	ScmConfigurationScmTypeExternalScm       ScmConfigurationScmTypeEnum = "EXTERNAL_SCM"
)

var mappingScmConfigurationScmTypeEnum = map[string]ScmConfigurationScmTypeEnum{
	"OCI_CODE_REPOSITORY": ScmConfigurationScmTypeOciCodeRepository,
	"EXTERNAL_SCM":        ScmConfigurationScmTypeExternalScm,
}

var mappingScmConfigurationScmTypeEnumLowerCase = map[string]ScmConfigurationScmTypeEnum{
	"oci_code_repository": ScmConfigurationScmTypeOciCodeRepository,
	"external_scm":        ScmConfigurationScmTypeExternalScm,
}

// GetScmConfigurationScmTypeEnumValues Enumerates the set of values for ScmConfigurationScmTypeEnum
func GetScmConfigurationScmTypeEnumValues() []ScmConfigurationScmTypeEnum {
	values := make([]ScmConfigurationScmTypeEnum, 0)
	for _, v := range mappingScmConfigurationScmTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetScmConfigurationScmTypeEnumStringValues Enumerates the set of values in String for ScmConfigurationScmTypeEnum
func GetScmConfigurationScmTypeEnumStringValues() []string {
	return []string{
		"OCI_CODE_REPOSITORY",
		"EXTERNAL_SCM",
	}
}

// GetMappingScmConfigurationScmTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingScmConfigurationScmTypeEnum(val string) (ScmConfigurationScmTypeEnum, bool) {
	enum, ok := mappingScmConfigurationScmTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
