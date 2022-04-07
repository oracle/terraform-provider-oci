resource "oci_devops_build_pipeline" "test_build_pipeline" {
  #Required
  project_id = oci_devops_project.test_project.id

  #Optional
  build_pipeline_parameters {
    #Required
    items {
      #Required
      name = "name"

      #Optional
      default_value = "defaultValue"
      description   = "description"
    }
  }
  description   = "description"
  display_name  = "displayName"
}

resource "oci_devops_build_pipeline_stage" "test_build_pipeline_stage" {
  #Required
  build_pipeline_id = oci_devops_build_pipeline.test_build_pipeline.id
  build_pipeline_stage_predecessor_collection {
    #Required
    items {
      #Required
      id = oci_devops_build_pipeline.test_build_pipeline.id
    }
  }
  build_pipeline_stage_type = "TRIGGER_DEPLOYMENT_PIPELINE"

  deploy_pipeline_id                 = oci_devops_deploy_pipeline.test_deploy_pipeline.id
  description                        = "description"
  display_name                       = "displayName"
  is_pass_all_parameters_enabled     = "false"

}

resource "oci_devops_build_pipeline_stage" "test_build_pipeline_wait_stage" {
  #Required
  build_pipeline_id = oci_devops_build_pipeline.test_build_pipeline.id
  build_pipeline_stage_predecessor_collection {
    #Required
    items {
      #Required
      id = oci_devops_build_pipeline.test_build_pipeline.id
    }
  }
  build_pipeline_stage_type = "WAIT"

  deploy_pipeline_id                 = oci_devops_deploy_pipeline.test_deploy_pipeline.id
  description                        = "description"
  display_name                       = "displayName"
  wait_criteria {
    wait_duration = "PT10S"
    wait_type = "ABSOLUTE_WAIT"
  }
}

resource "oci_devops_build_pipeline_stage" "test_build_pipeline_build_github_stage" {
  #Required
  build_pipeline_id = oci_devops_build_pipeline.test_build_pipeline.id
  build_pipeline_stage_predecessor_collection {
    #Required
    items {
      #Required
      id = oci_devops_build_pipeline.test_build_pipeline.id
    }
  }
  build_pipeline_stage_type = "BUILD"

  deploy_pipeline_id                 = oci_devops_deploy_pipeline.test_deploy_pipeline.id
  description                        = "description"
  display_name                       = "displayName"
  build_spec_file                    = "build_spec.yml"
  image                              = "OL7_X86_64_STANDARD_10"
  primary_build_source               = "primaryBuildSource"
  stage_execution_timeout_in_seconds = "10"
  build_source_collection {
    items {
      connection_type   = "GITHUB"
      branch            = "branch"
      connection_id     = oci_devops_connection.test_connection.id
      name              = "primaryBuildSource"
      repository_url    = "repositoryUrl"
    }
  }
}

resource "oci_devops_build_pipeline_stage" "test_build_pipeline_build_gitlab_stage" {
  #Required
  build_pipeline_id = oci_devops_build_pipeline.test_build_pipeline.id
  build_pipeline_stage_predecessor_collection {
    #Required
    items {
      #Required
      id = oci_devops_build_pipeline.test_build_pipeline.id
    }
  }
  build_pipeline_stage_type = "BUILD"

  deploy_pipeline_id                 = oci_devops_deploy_pipeline.test_deploy_pipeline.id
  description                        = "description"
  display_name                       = "displayName"
  build_spec_file                    = "build_spec.yml"
  image                              = "OL7_X86_64_STANDARD_10"
  primary_build_source               = "primaryBuildSource"
  stage_execution_timeout_in_seconds = "10"
  build_source_collection {
    items {
      connection_type   = "GITLAB"
      branch            = "branch"
      connection_id     = oci_devops_connection.test_gitlab_connection.id
      name              = "primaryBuildSource"
      repository_url    = "repositoryUrl"
    }
  }
}

resource "oci_devops_build_pipeline_stage" "test_build_pipeline_build_code_repo_stage" {
  #Required
  build_pipeline_id = oci_devops_build_pipeline.test_build_pipeline.id
  build_pipeline_stage_predecessor_collection {
    #Required
    items {
      #Required
      id = oci_devops_build_pipeline.test_build_pipeline.id
    }
  }
  build_pipeline_stage_type = "BUILD"

  deploy_pipeline_id                 = oci_devops_deploy_pipeline.test_deploy_pipeline.id
  description                        = "description"
  display_name                       = "displayName"
  build_spec_file                    = "build_spec.yml"
  image                              = "OL7_X86_64_STANDARD_10"
  primary_build_source               = "primaryBuildSource"
  stage_execution_timeout_in_seconds = "10"
  build_source_collection {
    items {
      connection_type   = "DEVOPS_CODE_REPOSITORY"
      branch            = "branch"
      repository_id     = oci_devops_repository.test_repository.id
      name              = "primaryBuildSource"
      repository_url    = "repositoryUrl"
    }
  }
}

resource "oci_devops_build_pipeline_stage" "test_build_pipeline_deliver_uim_artifact_stage" {
  #Required
  build_pipeline_id = oci_devops_build_pipeline.test_build_pipeline.id
  build_pipeline_stage_predecessor_collection {
    #Required
    items {
      #Required
      id = oci_devops_build_pipeline.test_build_pipeline.id
    }
  }
  build_pipeline_stage_type = "DELIVER_ARTIFACT"

  deploy_pipeline_id                 = oci_devops_deploy_pipeline.test_deploy_pipeline.id
  description                        = "description"
  display_name                       = "displayName"
  deliver_artifact_collection {
    items {
      artifact_id   = oci_devops_deploy_artifact.test_deploy_generic_artifact.id
      artifact_name = "artifactName2"
    }
  }
}

resource "oci_devops_build_pipeline_stage" "test_build_pipeline_deliver_ocir_artifact_stage" {
  #Required
  build_pipeline_id = oci_devops_build_pipeline.test_build_pipeline.id
  build_pipeline_stage_predecessor_collection {
    #Required
    items {
      #Required
      id = oci_devops_build_pipeline.test_build_pipeline.id
    }
  }
  build_pipeline_stage_type = "DELIVER_ARTIFACT"

  deploy_pipeline_id                 = oci_devops_deploy_pipeline.test_deploy_pipeline.id
  description                        = "description"
  display_name                       = "displayName"
  deliver_artifact_collection {
    items {
      artifact_id   = oci_devops_deploy_artifact.test_deploy_ocir_artifact.id
      artifact_name = "artifactName2"
    }
  }
}

resource "oci_devops_build_run" "test_build_run" {
  #Required
  build_pipeline_id = oci_devops_build_pipeline.test_build_pipeline.id

  #Optional
  build_run_arguments {
    #Required
    items {
      #Required
      name  = "name"
      value = "value"
    }
  }
  commit_info {
    #Required
    commit_hash       = "commitHash"
    repository_branch = "repositoryBranch"
    repository_url    = "repositoryUrl"
  }
  display_name  = "displayName"
}

resource "oci_devops_connection" "test_connection" {
  #Required
  access_token    = var.github_access_token_vault_id // Change it to connection_access_token
  connection_type = "GITHUB_ACCESS_TOKEN"
  project_id      = oci_devops_project.test_project.id

  #Optional
  description   = "description"
  display_name  = "display_name"
}

resource "oci_devops_connection" "test_gitlab_connection" {
  #Required
  access_token    = var.github_access_token_vault_id // Change it to connection_access_token
  connection_type = "GITLAB_ACCESS_TOKEN"
  project_id      = oci_devops_project.test_project.id

  #Optional
  description   = "description"
  display_name  = "display_name"
}

resource "oci_devops_trigger" "test_trigger" {
  #Required
  actions {
    #Required
    build_pipeline_id = oci_devops_build_pipeline.test_build_pipeline.id
    type              = "TRIGGER_BUILD_PIPELINE"

    #Optional
    filter {
      #Required
      trigger_source = "GITHUB"

      #Optional
      events = ["PUSH"]
      include {

        #Optional
        base_ref = "baseRef"
        head_ref = "headRef"
      }
    }
  }
  project_id     = oci_devops_project.test_project.id
  trigger_source = "GITHUB"

  #Optional
  description   = "description"
  display_name  = "displayName"
}

resource "oci_devops_trigger" "test_gitlab_trigger" {
  #Required
  actions {
    #Required
    build_pipeline_id = oci_devops_build_pipeline.test_build_pipeline.id
    type              = "TRIGGER_BUILD_PIPELINE"

    #Optional
    filter {
      #Required
      trigger_source = "GITLAB"

      #Optional
      events = ["PUSH"]
      include {

        #Optional
        base_ref = "baseRef"
        head_ref = "headRef"
      }
    }
  }
  project_id     = oci_devops_project.test_project.id
  trigger_source = "GITLAB"

  #Optional
  description   = "description"
  display_name  = "displayName"
}

resource "oci_devops_trigger" "test_code_repo_trigger" {
  #Required
  actions {
    #Required
    build_pipeline_id = oci_devops_build_pipeline.test_build_pipeline.id
    type              = "TRIGGER_BUILD_PIPELINE"

    #Optional
    filter {
      #Required
      trigger_source = "DEVOPS_CODE_REPOSITORY"

      #Optional
      events = ["PUSH"]
      include {

        #Optional
        base_ref = "baseRef"
        head_ref = "headRef"
      }
    }
  }
  project_id     = oci_devops_project.test_project.id
  trigger_source = "DEVOPS_CODE_REPOSITORY"

  #Optional
  description   = "description"
  display_name  = "displayName"
  repository_id = oci_devops_repository.test_repository.id
}