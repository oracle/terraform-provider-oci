// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DevOps API
//
// Use the DevOps API to create DevOps projects, configure code repositories,  add artifacts to deploy, build and test software applications, configure  target deployment environments, and deploy software applications.  For more information, see DevOps (https://docs.oracle.com/iaas/Content/devops/using/home.htm).
//

package devops

import (
	"strings"
)

// OperationTypeEnum Enum with underlying type: string
type OperationTypeEnum string

// Set of constants representing the allowable values for OperationTypeEnum
const (
	OperationTypeCreateProject                           OperationTypeEnum = "CREATE_PROJECT"
	OperationTypeUpdateProject                           OperationTypeEnum = "UPDATE_PROJECT"
	OperationTypeDeleteProject                           OperationTypeEnum = "DELETE_PROJECT"
	OperationTypeMoveProject                             OperationTypeEnum = "MOVE_PROJECT"
	OperationTypeCreateDeployPipeline                    OperationTypeEnum = "CREATE_DEPLOY_PIPELINE"
	OperationTypeUpdateDeployPipeline                    OperationTypeEnum = "UPDATE_DEPLOY_PIPELINE"
	OperationTypeDeleteDeployPipeline                    OperationTypeEnum = "DELETE_DEPLOY_PIPELINE"
	OperationTypeCreateDeployStage                       OperationTypeEnum = "CREATE_DEPLOY_STAGE"
	OperationTypeUpdateDeployStage                       OperationTypeEnum = "UPDATE_DEPLOY_STAGE"
	OperationTypeDeleteDeployStage                       OperationTypeEnum = "DELETE_DEPLOY_STAGE"
	OperationTypeCreateDeployArtifact                    OperationTypeEnum = "CREATE_DEPLOY_ARTIFACT"
	OperationTypeUpdateDeployArtifact                    OperationTypeEnum = "UPDATE_DEPLOY_ARTIFACT"
	OperationTypeDeleteDeployArtifact                    OperationTypeEnum = "DELETE_DEPLOY_ARTIFACT"
	OperationTypeCreateDeployEnvironment                 OperationTypeEnum = "CREATE_DEPLOY_ENVIRONMENT"
	OperationTypeUpdateDeployEnvironment                 OperationTypeEnum = "UPDATE_DEPLOY_ENVIRONMENT"
	OperationTypeDeleteDeployEnvironment                 OperationTypeEnum = "DELETE_DEPLOY_ENVIRONMENT"
	OperationTypeCreateDeployment                        OperationTypeEnum = "CREATE_DEPLOYMENT"
	OperationTypeUpdateDeployment                        OperationTypeEnum = "UPDATE_DEPLOYMENT"
	OperationTypeDeleteDeployment                        OperationTypeEnum = "DELETE_DEPLOYMENT"
	OperationTypeCreateBuildPipeline                     OperationTypeEnum = "CREATE_BUILD_PIPELINE"
	OperationTypeUpdateBuildPipeline                     OperationTypeEnum = "UPDATE_BUILD_PIPELINE"
	OperationTypeDeleteBuildPipeline                     OperationTypeEnum = "DELETE_BUILD_PIPELINE"
	OperationTypeCreateBuildPipelineStage                OperationTypeEnum = "CREATE_BUILD_PIPELINE_STAGE"
	OperationTypeUpdateBuildPipelineStage                OperationTypeEnum = "UPDATE_BUILD_PIPELINE_STAGE"
	OperationTypeDeleteBuildPipelineStage                OperationTypeEnum = "DELETE_BUILD_PIPELINE_STAGE"
	OperationTypeCreateConnection                        OperationTypeEnum = "CREATE_CONNECTION"
	OperationTypeUpdateConnection                        OperationTypeEnum = "UPDATE_CONNECTION"
	OperationTypeDeleteConnection                        OperationTypeEnum = "DELETE_CONNECTION"
	OperationTypeCreateTrigger                           OperationTypeEnum = "CREATE_TRIGGER"
	OperationTypeUpdateTrigger                           OperationTypeEnum = "UPDATE_TRIGGER"
	OperationTypeDeleteTrigger                           OperationTypeEnum = "DELETE_TRIGGER"
	OperationTypeExecuteTrigger                          OperationTypeEnum = "EXECUTE_TRIGGER"
	OperationTypeCreateRepository                        OperationTypeEnum = "CREATE_REPOSITORY"
	OperationTypeUpdateRepository                        OperationTypeEnum = "UPDATE_REPOSITORY"
	OperationTypeDeleteRepository                        OperationTypeEnum = "DELETE_REPOSITORY"
	OperationTypeMirrorRepository                        OperationTypeEnum = "MIRROR_REPOSITORY"
	OperationTypeForkRepository                          OperationTypeEnum = "FORK_REPOSITORY"
	OperationTypeSyncForkRepository                      OperationTypeEnum = "SYNC_FORK_REPOSITORY"
	OperationTypeScheduleCascadingProjectDeletion        OperationTypeEnum = "SCHEDULE_CASCADING_PROJECT_DELETION"
	OperationTypeCancelScheduledCascadingProjectDeletion OperationTypeEnum = "CANCEL_SCHEDULED_CASCADING_PROJECT_DELETION"
	OperationTypeCreatePullRequest                       OperationTypeEnum = "CREATE_PULL_REQUEST"
	OperationTypeUpdatePullRequest                       OperationTypeEnum = "UPDATE_PULL_REQUEST"
	OperationTypeDeletePullRequest                       OperationTypeEnum = "DELETE_PULL_REQUEST"
	OperationTypeMovePullRequest                         OperationTypeEnum = "MOVE_PULL_REQUEST"
	OperationTypeMergePullRequest                        OperationTypeEnum = "MERGE_PULL_REQUEST"
)

var mappingOperationTypeEnum = map[string]OperationTypeEnum{
	"CREATE_PROJECT":                              OperationTypeCreateProject,
	"UPDATE_PROJECT":                              OperationTypeUpdateProject,
	"DELETE_PROJECT":                              OperationTypeDeleteProject,
	"MOVE_PROJECT":                                OperationTypeMoveProject,
	"CREATE_DEPLOY_PIPELINE":                      OperationTypeCreateDeployPipeline,
	"UPDATE_DEPLOY_PIPELINE":                      OperationTypeUpdateDeployPipeline,
	"DELETE_DEPLOY_PIPELINE":                      OperationTypeDeleteDeployPipeline,
	"CREATE_DEPLOY_STAGE":                         OperationTypeCreateDeployStage,
	"UPDATE_DEPLOY_STAGE":                         OperationTypeUpdateDeployStage,
	"DELETE_DEPLOY_STAGE":                         OperationTypeDeleteDeployStage,
	"CREATE_DEPLOY_ARTIFACT":                      OperationTypeCreateDeployArtifact,
	"UPDATE_DEPLOY_ARTIFACT":                      OperationTypeUpdateDeployArtifact,
	"DELETE_DEPLOY_ARTIFACT":                      OperationTypeDeleteDeployArtifact,
	"CREATE_DEPLOY_ENVIRONMENT":                   OperationTypeCreateDeployEnvironment,
	"UPDATE_DEPLOY_ENVIRONMENT":                   OperationTypeUpdateDeployEnvironment,
	"DELETE_DEPLOY_ENVIRONMENT":                   OperationTypeDeleteDeployEnvironment,
	"CREATE_DEPLOYMENT":                           OperationTypeCreateDeployment,
	"UPDATE_DEPLOYMENT":                           OperationTypeUpdateDeployment,
	"DELETE_DEPLOYMENT":                           OperationTypeDeleteDeployment,
	"CREATE_BUILD_PIPELINE":                       OperationTypeCreateBuildPipeline,
	"UPDATE_BUILD_PIPELINE":                       OperationTypeUpdateBuildPipeline,
	"DELETE_BUILD_PIPELINE":                       OperationTypeDeleteBuildPipeline,
	"CREATE_BUILD_PIPELINE_STAGE":                 OperationTypeCreateBuildPipelineStage,
	"UPDATE_BUILD_PIPELINE_STAGE":                 OperationTypeUpdateBuildPipelineStage,
	"DELETE_BUILD_PIPELINE_STAGE":                 OperationTypeDeleteBuildPipelineStage,
	"CREATE_CONNECTION":                           OperationTypeCreateConnection,
	"UPDATE_CONNECTION":                           OperationTypeUpdateConnection,
	"DELETE_CONNECTION":                           OperationTypeDeleteConnection,
	"CREATE_TRIGGER":                              OperationTypeCreateTrigger,
	"UPDATE_TRIGGER":                              OperationTypeUpdateTrigger,
	"DELETE_TRIGGER":                              OperationTypeDeleteTrigger,
	"EXECUTE_TRIGGER":                             OperationTypeExecuteTrigger,
	"CREATE_REPOSITORY":                           OperationTypeCreateRepository,
	"UPDATE_REPOSITORY":                           OperationTypeUpdateRepository,
	"DELETE_REPOSITORY":                           OperationTypeDeleteRepository,
	"MIRROR_REPOSITORY":                           OperationTypeMirrorRepository,
	"FORK_REPOSITORY":                             OperationTypeForkRepository,
	"SYNC_FORK_REPOSITORY":                        OperationTypeSyncForkRepository,
	"SCHEDULE_CASCADING_PROJECT_DELETION":         OperationTypeScheduleCascadingProjectDeletion,
	"CANCEL_SCHEDULED_CASCADING_PROJECT_DELETION": OperationTypeCancelScheduledCascadingProjectDeletion,
	"CREATE_PULL_REQUEST":                         OperationTypeCreatePullRequest,
	"UPDATE_PULL_REQUEST":                         OperationTypeUpdatePullRequest,
	"DELETE_PULL_REQUEST":                         OperationTypeDeletePullRequest,
	"MOVE_PULL_REQUEST":                           OperationTypeMovePullRequest,
	"MERGE_PULL_REQUEST":                          OperationTypeMergePullRequest,
}

var mappingOperationTypeEnumLowerCase = map[string]OperationTypeEnum{
	"create_project":                              OperationTypeCreateProject,
	"update_project":                              OperationTypeUpdateProject,
	"delete_project":                              OperationTypeDeleteProject,
	"move_project":                                OperationTypeMoveProject,
	"create_deploy_pipeline":                      OperationTypeCreateDeployPipeline,
	"update_deploy_pipeline":                      OperationTypeUpdateDeployPipeline,
	"delete_deploy_pipeline":                      OperationTypeDeleteDeployPipeline,
	"create_deploy_stage":                         OperationTypeCreateDeployStage,
	"update_deploy_stage":                         OperationTypeUpdateDeployStage,
	"delete_deploy_stage":                         OperationTypeDeleteDeployStage,
	"create_deploy_artifact":                      OperationTypeCreateDeployArtifact,
	"update_deploy_artifact":                      OperationTypeUpdateDeployArtifact,
	"delete_deploy_artifact":                      OperationTypeDeleteDeployArtifact,
	"create_deploy_environment":                   OperationTypeCreateDeployEnvironment,
	"update_deploy_environment":                   OperationTypeUpdateDeployEnvironment,
	"delete_deploy_environment":                   OperationTypeDeleteDeployEnvironment,
	"create_deployment":                           OperationTypeCreateDeployment,
	"update_deployment":                           OperationTypeUpdateDeployment,
	"delete_deployment":                           OperationTypeDeleteDeployment,
	"create_build_pipeline":                       OperationTypeCreateBuildPipeline,
	"update_build_pipeline":                       OperationTypeUpdateBuildPipeline,
	"delete_build_pipeline":                       OperationTypeDeleteBuildPipeline,
	"create_build_pipeline_stage":                 OperationTypeCreateBuildPipelineStage,
	"update_build_pipeline_stage":                 OperationTypeUpdateBuildPipelineStage,
	"delete_build_pipeline_stage":                 OperationTypeDeleteBuildPipelineStage,
	"create_connection":                           OperationTypeCreateConnection,
	"update_connection":                           OperationTypeUpdateConnection,
	"delete_connection":                           OperationTypeDeleteConnection,
	"create_trigger":                              OperationTypeCreateTrigger,
	"update_trigger":                              OperationTypeUpdateTrigger,
	"delete_trigger":                              OperationTypeDeleteTrigger,
	"execute_trigger":                             OperationTypeExecuteTrigger,
	"create_repository":                           OperationTypeCreateRepository,
	"update_repository":                           OperationTypeUpdateRepository,
	"delete_repository":                           OperationTypeDeleteRepository,
	"mirror_repository":                           OperationTypeMirrorRepository,
	"fork_repository":                             OperationTypeForkRepository,
	"sync_fork_repository":                        OperationTypeSyncForkRepository,
	"schedule_cascading_project_deletion":         OperationTypeScheduleCascadingProjectDeletion,
	"cancel_scheduled_cascading_project_deletion": OperationTypeCancelScheduledCascadingProjectDeletion,
	"create_pull_request":                         OperationTypeCreatePullRequest,
	"update_pull_request":                         OperationTypeUpdatePullRequest,
	"delete_pull_request":                         OperationTypeDeletePullRequest,
	"move_pull_request":                           OperationTypeMovePullRequest,
	"merge_pull_request":                          OperationTypeMergePullRequest,
}

// GetOperationTypeEnumValues Enumerates the set of values for OperationTypeEnum
func GetOperationTypeEnumValues() []OperationTypeEnum {
	values := make([]OperationTypeEnum, 0)
	for _, v := range mappingOperationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetOperationTypeEnumStringValues Enumerates the set of values in String for OperationTypeEnum
func GetOperationTypeEnumStringValues() []string {
	return []string{
		"CREATE_PROJECT",
		"UPDATE_PROJECT",
		"DELETE_PROJECT",
		"MOVE_PROJECT",
		"CREATE_DEPLOY_PIPELINE",
		"UPDATE_DEPLOY_PIPELINE",
		"DELETE_DEPLOY_PIPELINE",
		"CREATE_DEPLOY_STAGE",
		"UPDATE_DEPLOY_STAGE",
		"DELETE_DEPLOY_STAGE",
		"CREATE_DEPLOY_ARTIFACT",
		"UPDATE_DEPLOY_ARTIFACT",
		"DELETE_DEPLOY_ARTIFACT",
		"CREATE_DEPLOY_ENVIRONMENT",
		"UPDATE_DEPLOY_ENVIRONMENT",
		"DELETE_DEPLOY_ENVIRONMENT",
		"CREATE_DEPLOYMENT",
		"UPDATE_DEPLOYMENT",
		"DELETE_DEPLOYMENT",
		"CREATE_BUILD_PIPELINE",
		"UPDATE_BUILD_PIPELINE",
		"DELETE_BUILD_PIPELINE",
		"CREATE_BUILD_PIPELINE_STAGE",
		"UPDATE_BUILD_PIPELINE_STAGE",
		"DELETE_BUILD_PIPELINE_STAGE",
		"CREATE_CONNECTION",
		"UPDATE_CONNECTION",
		"DELETE_CONNECTION",
		"CREATE_TRIGGER",
		"UPDATE_TRIGGER",
		"DELETE_TRIGGER",
		"EXECUTE_TRIGGER",
		"CREATE_REPOSITORY",
		"UPDATE_REPOSITORY",
		"DELETE_REPOSITORY",
		"MIRROR_REPOSITORY",
		"FORK_REPOSITORY",
		"SYNC_FORK_REPOSITORY",
		"SCHEDULE_CASCADING_PROJECT_DELETION",
		"CANCEL_SCHEDULED_CASCADING_PROJECT_DELETION",
		"CREATE_PULL_REQUEST",
		"UPDATE_PULL_REQUEST",
		"DELETE_PULL_REQUEST",
		"MOVE_PULL_REQUEST",
		"MERGE_PULL_REQUEST",
	}
}

// GetMappingOperationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOperationTypeEnum(val string) (OperationTypeEnum, bool) {
	enum, ok := mappingOperationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
