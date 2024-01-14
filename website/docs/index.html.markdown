---
layout: "oci"
page_title: "Provider: Oracle Cloud Infrastructure"
sidebar_current: "docs-oci-index"
description: |-
The Oracle Cloud Infrastructure provider is used to interact with the resources supported by the Oracle Cloud Infrastructure services. The provider must be configured with credentials for the Oracle Cloud Account.
---

# Oracle Cloud Infrastructure Provider

The Oracle Cloud Infrastructure (OCI) provider allows you to use Terraform to interact with [Oracle Cloud Infrastructure](https://cloud.oracle.com/cloud-infrastructure) resources. Wherever you use a Terraform distribution you can use the OCI Terraform provider, including [Terraform Cloud](https://www.terraform.io/docs/cloud/index.html) and the OCI [Resource Manager](#resource-manager).

To learn the basics of Terraform using this provider, follow the
hands-on [get started tutorials](https://developer.hashicorp.com/tutorials/terraform/infrastructure-as-code?in=terraform/oci-get-started) on HashiCorp's Learn platform.

## Terraform CLI

For details on configuring the OCI Terraform provider and using it with the Terraform CLI, refer to the [official OCI Terraform provider documentation](https://docs.oracle.com/en-us/iaas/Content/API/SDKDocs/terraform.htm), which includes:

- How to [install the provider](https://docs.oracle.com/en-us/iaas/Content/API/SDKDocs/terraforminstallation.htm)
- How to [configure the provider](https://docs.oracle.com/en-us/iaas/Content/API/SDKDocs/terraformproviderconfiguration.htm)
- Ways you can [describe your infrastructure](https://docs.oracle.com/en-us/iaas/Content/API/SDKDocs/terraformdescribingyourinfrastructure.htm)
- How to [apply your configurations](https://docs.oracle.com/en-us/iaas/Content/API/SDKDocs/terraformapplyingyourconfigurations.htm)
- [Example configurations and solutions](https://docs.oracle.com/en-us/iaas/Content/API/SDKDocs/terraformexamples.htm)
- Troubleshooting [basics](https://docs.oracle.com/en-us/iaas/Content/API/SDKDocs/terraformtroubleshootingbasics.htm) and [common issues](https://docs.oracle.com/en-us/iaas/Content/API/SDKDocs/terraformtroubleshooting.htm)

As you write your configuration files, use the left navigation panel on this page to access detailed information about each supported resource and data source.

## Resource Manager

The Oracle Cloud Infrastructure [Resource Manager](https://docs.oracle.com/en-us/iaas/Content/ResourceManager/Concepts/landing.htm#ResourceManager) is an Oracle-managed service that is based on Terraform and uses Terraform configuration files to automate deployment and operations for the OCI resources supported by the OCI Terraform provider.

Resource Manager allows you to share and manage infrastructure configurations and state files across multiple teams and platforms. This infrastructure management can't be done with local Terraform installations and Oracle Terraform modules alone. See the [Overview of Resource Manager](https://docs.oracle.com/en-us/iaas/Content/ResourceManager/Concepts/resourcemanager.htm) for more information.