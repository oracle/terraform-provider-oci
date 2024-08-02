#Step 6 - Example
variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "compartment_ocid" {}
variable "region" {}
variable "default_branch" {}
variable "baseVersion" {
  default = "master"
}
variable "baseVersion_fork" {
  default = "main"
}
variable "targetVersion" {
  default = "<target>"
}
variable "connection_access_token" {
  default = ""
}

resource "random_string" "topicname" {
  length  = 10
  special = false
}

resource "random_string" "projectname" {
  length  = 10
  special = false
}

resource "oci_ons_notification_topic" "test_notification_topic" {
  #Required
  compartment_id = var.compartment_ocid
  name           = join("", ["A", random_string.topicname.result])
}

# Create Project
resource "oci_devops_project" "test_project" {
  #Required
  compartment_id = var.compartment_ocid
  name = join("", ["A", random_string.projectname.result])
  notification_config {
    #Required
    topic_id = oci_ons_notification_topic.test_notification_topic.id
  }
}

resource "oci_devops_connection" "test_connection" {
  #Required
  connection_type = "GITHUB_ACCESS_TOKEN"
  project_id      = oci_devops_project.test_project.id
  access_token = var.connection_access_token

  #Optional
  description   = "description"
  display_name  = "display_name"
}

# Create HOSTED repository
resource "oci_devops_repository" "test_repository" {
  #Required
  name       = "repositoryName"
  project_id = oci_devops_project.test_project.id
  repository_type = "HOSTED"

  #Optional
  default_branch = var.default_branch
  description    = "description"
}

# Retrieve List of commits in HOSTED repository
data "oci_devops_repository_commits" "test_repository_commits" {
  #Required
  repository_id = oci_devops_repository.test_repository.id
}

# Create new branch in HOSTED repository
resource "oci_devops_repository_ref" "test_repository_ref" {
  commit_id = lookup(data.oci_devops_repository_commits.test_repository_commits.repository_commit_collection[0].items[0], "commit_id")
  ref_name = "<refName>"
  ref_type = "BRANCH"
  repository_id = oci_devops_repository.test_repository.id

  lifecycle {
    ignore_changes = [
      defined_tags,
      freeform_tags
    ]
  }
}

# Creating new resource FORKED repository
resource "oci_devops_repository" "test_repository_fork" {
  #Required
  name      = "forkRepositoryName"
  project_id = oci_devops_project.test_project.id
  repository_type = "FORKED"

  #Optional
  description = "description"
  parent_repository_id = oci_devops_repository.test_repository.id
}

# ListCommitsDiff
data "oci_devops_repository_diffs" "test_repository_diff" {
  #Required
  base_version = var.baseVersion_fork
  repository_id = oci_devops_repository.test_repository_fork.id
  target_version = var.targetVersion

  #Optional
  is_comparison_from_merge_base = false
  target_repository_id = oci_devops_repository.test_repository.id
}

output "commit_diffs" {
  value = data.oci_devops_repository_diffs.test_repository_diff
}