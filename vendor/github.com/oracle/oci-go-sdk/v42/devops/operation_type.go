// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DevOps API
//
// Use the DevOps APIs to create a DevOps project to group the pipelines,  add reference to target deployment environments, add artifacts to deploy,  and create deployment pipelines needed to deploy your software.
//

package devops

// OperationTypeEnum Enum with underlying type: string
type OperationTypeEnum string

// Set of constants representing the allowable values for OperationTypeEnum
const (
	OperationTypeCreateProject           OperationTypeEnum = "CREATE_PROJECT"
	OperationTypeUpdateProject           OperationTypeEnum = "UPDATE_PROJECT"
	OperationTypeDeleteProject           OperationTypeEnum = "DELETE_PROJECT"
	OperationTypeMoveProject             OperationTypeEnum = "MOVE_PROJECT"
	OperationTypeCreateDeployPipeline    OperationTypeEnum = "CREATE_DEPLOY_PIPELINE"
	OperationTypeUpdateDeployPipeline    OperationTypeEnum = "UPDATE_DEPLOY_PIPELINE"
	OperationTypeDeleteDeployPipeline    OperationTypeEnum = "DELETE_DEPLOY_PIPELINE"
	OperationTypeCreateDeployStage       OperationTypeEnum = "CREATE_DEPLOY_STAGE"
	OperationTypeUpdateDeployStage       OperationTypeEnum = "UPDATE_DEPLOY_STAGE"
	OperationTypeDeleteDeployStage       OperationTypeEnum = "DELETE_DEPLOY_STAGE"
	OperationTypeCreateDeployArtifact    OperationTypeEnum = "CREATE_DEPLOY_ARTIFACT"
	OperationTypeUpdateDeployArtifact    OperationTypeEnum = "UPDATE_DEPLOY_ARTIFACT"
	OperationTypeDeleteDeployArtifact    OperationTypeEnum = "DELETE_DEPLOY_ARTIFACT"
	OperationTypeCreateDeployEnvironment OperationTypeEnum = "CREATE_DEPLOY_ENVIRONMENT"
	OperationTypeUpdateDeployEnvironment OperationTypeEnum = "UPDATE_DEPLOY_ENVIRONMENT"
	OperationTypeDeleteDeployEnvironment OperationTypeEnum = "DELETE_DEPLOY_ENVIRONMENT"
	OperationTypeCreateDeployment        OperationTypeEnum = "CREATE_DEPLOYMENT"
	OperationTypeUpdateDeployment        OperationTypeEnum = "UPDATE_DEPLOYMENT"
	OperationTypeDeleteDeployment        OperationTypeEnum = "DELETE_DEPLOYMENT"
)

var mappingOperationType = map[string]OperationTypeEnum{
	"CREATE_PROJECT":            OperationTypeCreateProject,
	"UPDATE_PROJECT":            OperationTypeUpdateProject,
	"DELETE_PROJECT":            OperationTypeDeleteProject,
	"MOVE_PROJECT":              OperationTypeMoveProject,
	"CREATE_DEPLOY_PIPELINE":    OperationTypeCreateDeployPipeline,
	"UPDATE_DEPLOY_PIPELINE":    OperationTypeUpdateDeployPipeline,
	"DELETE_DEPLOY_PIPELINE":    OperationTypeDeleteDeployPipeline,
	"CREATE_DEPLOY_STAGE":       OperationTypeCreateDeployStage,
	"UPDATE_DEPLOY_STAGE":       OperationTypeUpdateDeployStage,
	"DELETE_DEPLOY_STAGE":       OperationTypeDeleteDeployStage,
	"CREATE_DEPLOY_ARTIFACT":    OperationTypeCreateDeployArtifact,
	"UPDATE_DEPLOY_ARTIFACT":    OperationTypeUpdateDeployArtifact,
	"DELETE_DEPLOY_ARTIFACT":    OperationTypeDeleteDeployArtifact,
	"CREATE_DEPLOY_ENVIRONMENT": OperationTypeCreateDeployEnvironment,
	"UPDATE_DEPLOY_ENVIRONMENT": OperationTypeUpdateDeployEnvironment,
	"DELETE_DEPLOY_ENVIRONMENT": OperationTypeDeleteDeployEnvironment,
	"CREATE_DEPLOYMENT":         OperationTypeCreateDeployment,
	"UPDATE_DEPLOYMENT":         OperationTypeUpdateDeployment,
	"DELETE_DEPLOYMENT":         OperationTypeDeleteDeployment,
}

// GetOperationTypeEnumValues Enumerates the set of values for OperationTypeEnum
func GetOperationTypeEnumValues() []OperationTypeEnum {
	values := make([]OperationTypeEnum, 0)
	for _, v := range mappingOperationType {
		values = append(values, v)
	}
	return values
}
