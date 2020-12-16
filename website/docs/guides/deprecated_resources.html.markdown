---
layout: "oci"
page_title: "Deprecated Resources"
sidebar_current: "docs-oci-guide-deprecated_resources"
description: |-
  The Oracle Cloud Infrastructure provider deprecation guide
---

## OCI Terraform Provider Deprecation Guide

This guide covers the list of resources and data sources that have been marked deprecated and their respective suggested replacements, if any.

Resources and Data Sources marked for deprecation will be warned during Terraform `plan` and `apply`  operations.

Resources on path to deprecation may stop working in future, use the respective guide, if available, on how to migrate using the new replacements.

### Deprecated Resources and Data Sources

Resources that have a migration path have deprecation guides available on how to rename and migrate them to their new replacements.

Data Sources do not have deprecation guide as one should be able to directly replace them in their Terraform configuration and refresh the state.

**Legend:** How to use deprecation table below:

| Column                       | Details                                                              |
|------------------------------|----------------------------------------------------------------------|
| Version                      | Provider version in which said resource or data source was marked deprecated |
| Type                         | Type of the deprecated resource or data source |
| Old Deprecated Resource Name | Deprecated resource or data source name |
| New Resource Name            | New resource or data source name that will provide the same functionality |
| Migration?                   | If migration is possible to the new resource through Terraform `state` import |
| Guide                        | Link to deprecation guide on how to rename and migrate to new resource, if applicable |


**IMPORTANT**: Before executing any deprecation guide, please ensure that you have backed up your Terraform state file to avoid any **data loss**.

| Version | Type        | Old Deprecated Resource Name              | New Resource Name                         | Migration? | Guide  |
|---------|-------------|-------------------------------------------|-------------------------------------------|------------|--------|
| 4.7.0   | Data Source | `oci_database_autonomous_database_wallet` | `oci_database_autonomous_database_wallet` | N/A        | N/A    |
| 3.97.0  | Resource    | `oci_dns_records`                         | `oci_dns_rrset`                           | N/A        | N/A    |
| 3.97.0  | Resource    | `oci_dns_record`                          | `oci_dns_rrset`                           | N/A        | N/A    |
| 3.18    | Resource    | `oci_autonomous_data_warehouse`           | `oci_autonomous_database`                 | Yes        | N/A    |
| 3.18    | Data Source | `oci_autonomous_data_warehouse`           | `oci_autonomous_database`                 | N/A        | N/A    |
| 3.18    | Data Source | `oci_autonomous_data_warehouses`          | `oci_autonomous_databases`                | N/A        | N/A    |
| 3.18    | Resource    | `oci_autonomous_data_warehouse_backup`    | `oci_autonomous_database_backup`          | Yes        | N/A    |
| 3.18    | Data Source | `oci_autonomous_data_warehouse_backup`    | `oci_autonomous_database_backup`          | N/A        | N/A    |
| 3.18    | Data Source | `oci_autonomous_data_warehouse_backups`   | `oci_autonomous_database_backups`         | N/A        | N/A    |
| 2.1.12  | Resource    | `oci_swift_password`                      | `oci_identity_auth_token`                 | No         | N/A    |
| 2.1.12  | Data Source | `oci_swift_passwords`                     | `oci_identity_auth_tokens`                | N/A        | N/A    |

### Deprecated Fields

Deprecation notices including for fields can be found in any of the previously released [CHANGELOG](https://github.com/terraform-providers/terraform-provider-oci/blob/master/CHANGELOG.md).
Deprecated fields will be shown as deprecated during Terraform `plan` and `apply` operations.

### Deprecation Message Examples

* Resource - The 'oci_autonomous_data_warehouse' resource has been deprecated. Please use 'oci_autonomous_database' instead.
* Fields - The 'size_in_mbs' field has been deprecated. Please use 'size_in_gbs' instead.
