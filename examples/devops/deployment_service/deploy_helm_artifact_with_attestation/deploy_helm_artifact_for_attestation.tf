/*
This example demonstrates the various ways a helm chart artifact can be created and specifically,
the options available for having a signed helm chart verified per
https://helm.sh/docs/topics/provenance/.
*/

variable "tenancy_ocid" {
}

variable "user_ocid" {
}

variable "fingerprint" {
}

variable "private_key_path" {
}

variable "region" {
}

variable "vault_secret_public_key_ocid" {
}

provider "oci" {
  region           = var.region
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
}


# OCID of the DevOps project.
variable "project_id" {
}


locals {
  chart_url = "oci://iad.ocir.io/namespace/repository/chart-name"
  chart_version = "0.0.1"
  vault_secret_public_key_ocid = var.vault_secret_public_key_ocid
  public_key = "" // Add the public key here
}


resource "oci_devops_deploy_artifact" "helm_chart_implicit_no_public_key" {
    #Required
    argument_substitution_mode = "SUBSTITUTE_PLACEHOLDERS"
    deploy_artifact_source {
        #Required
        deploy_artifact_source_type = "HELM_CHART"

        #Optional
        chart_url = local.chart_url
        deploy_artifact_version = local.chart_version
    }
    deploy_artifact_type = "HELM_CHART"
    project_id = var.project_id

    #Optional
    description = "helm chart artifact with no public key"
    display_name = "helmChartWithImplicitNoPublicKey"
}


resource "oci_devops_deploy_artifact" "helm_chart_no_public_key" {
    #Required
    argument_substitution_mode = "SUBSTITUTE_PLACEHOLDERS"
    deploy_artifact_source {
        #Required
        deploy_artifact_source_type = "HELM_CHART"

        #Optional
        chart_url = local.chart_url
        deploy_artifact_version = local.chart_version

        helm_verification_key_source {
            verification_key_source_type = "NONE"
        }
    }
    deploy_artifact_type = "HELM_CHART"
    project_id = var.project_id

    #Optional
    description = "helm chart artifact with no public key"
    display_name = "helmChartWithNoPublicKey"

}



resource "oci_devops_deploy_artifact" "helm_chart_vault_public_key" {
    #Required
    argument_substitution_mode = "SUBSTITUTE_PLACEHOLDERS"
    deploy_artifact_source {
        #Required
        deploy_artifact_source_type = "HELM_CHART"

        #Optional
        chart_url = local.chart_url
        deploy_artifact_version = local.chart_version

        helm_verification_key_source {
            verification_key_source_type = "VAULT_SECRET"
            vault_secret_id = local.vault_secret_public_key_ocid
        }
    }
    deploy_artifact_type = "HELM_CHART"
    project_id = var.project_id

    #Optional
    description = "helm chart artifact with vault public key"
    display_name = "helmChartVaultPublicKey"


}


resource "oci_devops_deploy_artifact" "helm_chart_inline_public_key" {
    #Required
    argument_substitution_mode = "SUBSTITUTE_PLACEHOLDERS"
    deploy_artifact_source {
        #Required
        deploy_artifact_source_type = "HELM_CHART"

        #Optional
        chart_url = local.chart_url
        deploy_artifact_version = local.chart_version

        helm_verification_key_source {
            verification_key_source_type = "INLINE_PUBLIC_KEY"
            current_public_key = local.public_key
            previous_public_key = local.public_key
        }
    }
    deploy_artifact_type = "HELM_CHART"
    project_id = var.project_id

    #Optional
    description = "helm chart artifact with inline public keys"
    display_name = "helmChartInlinePublicKeys"

}

data "oci_devops_deploy_artifact" "retrieve_artifact_with_vault_public_key" {
    deploy_artifact_id = oci_devops_deploy_artifact.helm_chart_vault_public_key.id
}

data "oci_devops_deploy_artifact" "retrieve_artifact_with_inline_public_key" {
    deploy_artifact_id = oci_devops_deploy_artifact.helm_chart_inline_public_key.id
}

data "oci_devops_deploy_artifact" "retrieve_artifact_with_no_public_key" {
    deploy_artifact_id = oci_devops_deploy_artifact.helm_chart_no_public_key.id
}

data "oci_devops_deploy_artifact" "retrieve_artifact_implicit_no_public_key" {
    deploy_artifact_id = oci_devops_deploy_artifact.helm_chart_implicit_no_public_key.id
}


output helm_chart_implicit_no_public_key {
    value = data.oci_devops_deploy_artifact.retrieve_artifact_implicit_no_public_key
}

output helm_chart_no_public_key {
    value = data.oci_devops_deploy_artifact.retrieve_artifact_with_no_public_key
}

output helm_chart_vault_public_key {
    value = data.oci_devops_deploy_artifact.retrieve_artifact_with_vault_public_key
}

output helm_chart_inline_public_key {
    value = data.oci_devops_deploy_artifact.retrieve_artifact_with_inline_public_key
}

