// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Science API
//
// Use the Data Science APIs to organize your data science work, access data and computing resources, and build, train, deploy, and manage models on Oracle Cloud.
//

package datascience

import (
	"github.com/oracle/oci-go-sdk/common"
)

// ModelProvenance Model provenance gives data scientists information about the origin of their model. This information allows data scientists to reproduce the development environment in which the model was trained.
type ModelProvenance struct {

	// For model reproducibility purposes. URL of the git repository associated with model training.
	RepositoryUrl *string `mandatory:"false" json:"repositoryUrl"`

	// For model reproducibility purposes. Branch of the git repository associated with model training.
	GitBranch *string `mandatory:"false" json:"gitBranch"`

	// For model reproducibility purposes. Commit ID of the git repository associated with model training.
	GitCommit *string `mandatory:"false" json:"gitCommit"`

	// For model reproducibility purposes. Path to model artifacts.
	ScriptDir *string `mandatory:"false" json:"scriptDir"`

	// For model reproducibility purposes. Path to the python script or notebook in which the model was trained."
	TrainingScript *string `mandatory:"false" json:"trainingScript"`
}

func (m ModelProvenance) String() string {
	return common.PointerString(m)
}
