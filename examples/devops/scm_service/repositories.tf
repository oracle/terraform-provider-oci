resource "oci_devops_repository" "test_repository" {
  #Required
  name       = "repositoryName"
  project_id = oci_devops_project.test_project.id

  #Optional
  default_branch = "defaultBranch"
  description    = "description"
  repository_type = "HOSTED"
}

resource "oci_devops_repository" "test_mirrored_repository" {
  #Required
  name       = "repositoryMirroredName"
  project_id = oci_devops_project.test_project.id

  #Optional
  default_branch = "defaultBranch"
  description    = "description"
  repository_type = "MIRRORED"
  mirror_repository_config {
    connector_id = oci_devops_connection.test_connection.id
    repository_url = "https://github.com/Maxrovr/hello-dev"
    trigger_schedule {
      schedule_type = "NONE"
    }
  }
}

resource "oci_devops_repository_mirror" "test_repository_mirror" {
  #Required
  repository_id = oci_devops_repository.test_mirrored_repository.id
}

resource "oci_devops_repository_ref" "test_repository_ref" {
  commit_id = "commitId"
  ref_name = "refName"
  ref_type = "BRANCH"
  repository_id = oci_devops_repository.test_repository.id
}

resource "oci_devops_repository_ref" "test_repository_ref" {
  object_id = "object_id"
  ref_name = "refName"
  ref_type = "TAG"
  repository_id = oci_devops_repository.test_repository.id
}
