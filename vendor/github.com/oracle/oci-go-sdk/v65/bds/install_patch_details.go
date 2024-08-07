// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// InstallPatchDetails The reqeust body while installing a patch to a cluster.
type InstallPatchDetails struct {

	// The version of the patch to be installed.
	Version *string `mandatory:"true" json:"version"`

	// Base-64 encoded password for the cluster admin user.
	ClusterAdminPassword *string `mandatory:"true" json:"clusterAdminPassword"`

	PatchingConfig OdhPatchingConfig `mandatory:"false" json:"patchingConfig"`
}

func (m InstallPatchDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m InstallPatchDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *InstallPatchDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		PatchingConfig       odhpatchingconfig `json:"patchingConfig"`
		Version              *string           `json:"version"`
		ClusterAdminPassword *string           `json:"clusterAdminPassword"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
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

	m.ClusterAdminPassword = model.ClusterAdminPassword

	return
}
