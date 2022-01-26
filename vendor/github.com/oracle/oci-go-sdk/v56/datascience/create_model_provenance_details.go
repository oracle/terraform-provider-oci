// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Science API
//
// Use the Data Science API to organize your data science work, access data and computing resources, and build, train, deploy and manage models and model deployments. For more information, see Data Science (https://docs.oracle.com/iaas/data-science/using/data-science.htm).
//

package datascience

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// CreateModelProvenanceDetails Model provenance gives data scientists information about the origin of their model. This information allows data scientists to reproduce the development environment in which the model was trained.
type CreateModelProvenanceDetails struct {

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

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of a training session(Job or NotebookSession) in which the model was trained. It is used for model reproducibility purposes.
	TrainingId *string `mandatory:"false" json:"trainingId"`
}

func (m CreateModelProvenanceDetails) String() string {
	return common.PointerString(m)
}
