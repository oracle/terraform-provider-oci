// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DevOps API
//
// Use the DevOps API to create DevOps projects, configure code repositories,  add artifacts to deploy, build and test software applications, configure  target deployment environments, and deploy software applications.  For more information, see DevOps (https://docs.oracle.com/iaas/Content/devops/using/home.htm).
//

package devops

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// HelmDiffArgument Parameters for all the helm stages passed for retrieving Helm Diff
type HelmDiffArgument struct {

	// Deploy Stage OCID.
	StageId *string `mandatory:"false" json:"stageId"`

	HelmArgSetValues *HelmSetValueCollection `mandatory:"false" json:"helmArgSetValues"`

	HelmArgStringValues *HelmSetValueCollection `mandatory:"false" json:"helmArgStringValues"`

	ValueArtifactContents *ValueArtifactContentCollection `mandatory:"false" json:"valueArtifactContents"`

	// Stage specific values along with the helm chart injected by Shepherd
	HelmStageContents []HelmStageContent `mandatory:"false" json:"helmStageContents"`
}

func (m HelmDiffArgument) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m HelmDiffArgument) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
