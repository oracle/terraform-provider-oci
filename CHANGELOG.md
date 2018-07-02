# Change Log
All notable changes to this project are documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/).

## 2.1.13 - 2018-07-02

### Added
- Add defined and freeform tags to applicable resources, see [Tagging Resources](https://github.com/oracle/terraform-provider-oci/blob/master/docs/Tagging%20Resources.md)
- Manage defined tags
- Filter by tags in data sources
- Support health status datasources for load balancer, backends, and backend sets
- Object Storage Buckets supports [storage tier](https://docs.cloud.oracle.com/iaas/Content/Object/Tasks/managingbuckets.htm) settings.
- Object Storage Objects can be renamed.
- Object Storage Objects data source supports specifying a `delimiter`.
- DBsystems supports update. This allows scaling up the cpu_core_count in and the data_storage_size_in_gb.
- Create backups from a database.
- Support creating a DBSystem from a Database backup.
- Support db_system_id for db_versions data source.
- The db_system_shapes data source results now include information about max/min node count, and min core count supported by the relevant shape.
- Assign backup policies to volumes.
- Support additional ways of finding a Public IP via custom Public IP data source.
- Ability to create and manage console connections.

### Changed
- Object Storage Object's attributes other than `name` are now marked `forceNew`. This is consistent with the behavior of the service as defined [here](https://docs.cloud.oracle.com/iaas/api/#/en/objectstorage/20160918/Object/PutObject).

### Fixed
- Multiple updates on Object Storage Object's metadata used to cause contents of the file to get overwritten by its md5 value.
- DBSystems cpu_core_count was made optional as the service ignores it when you provide a VM shape. [#517](https://github.com/oracle/terraform-provider-oci/issues/517), [#539](https://github.com/oracle/terraform-provider-oci/issues/539).


## 2.1.12 - 2018-06-14

### Added
- Support importing images from object store or external sources.
- Updated Terraform Provider to use LaunchDbSystemDetails to provision DbSystem resource.
- Fix orphaned load balancer backend on port change [#519](https://github.com/oracle/terraform-provider-oci/issues/519).
- Fix to example in Route Tables documentation file.
- Added support for AuthToken Resource (replacement of deprecated SwiftPasswords) in Identity Service.
- Added support for Volume Group and Volume Group Backup.
- HCL syntax highlighting in docs
- Nil checks for time properties to avoid panic

## 2.1.10 - 2018-05-24

### Added
- Support for dynamic group resources and data sources
- Support for object storage namespace metadata resources and data sources
- Support for region subscription data sources

## 2.1.9 - 2018-05-17

### Added
- Added support for customer secret keys. More details can be found [here](https://github.com/oracle/terraform-provider-oci/tree/master/docs/identity/customer_secret_keys.md).
- Added boot volume attachments data source. More details can be found [here](https://github.com/oracle/terraform-provider-oci/tree/master/docs/core/boot_volume_attachments.md).
- Added region data source. More details can be found [here](https://github.com/oracle/terraform-provider-oci/tree/master/docs/identity/regions.md).
- Added tenancy data source. More details can be found [here](https://github.com/oracle/terraform-provider-oci/tree/master/docs/identity/tenancies.md).


## 2.1.8 - 2018-05-10

### Added
- Added support for remote VCN peering. More details can be found [here](https://github.com/oracle/terraform-provider-oci/blob/master/docs/core/remote_peering_connections.md), and an example [here](https://github.com/oracle/terraform-provider-oci/blob/master/docs/examples/networking/remote_vcn_peering_full).
- Added a data source for boot volumes. More details can be found [here](https://github.com/oracle/terraform-provider-oci/blob/master/docs/core/boot_volumes.md).

### Fixed
- Fixed a crash that can occur when using the `oci_identity_api_key` resource and editing the API key outside of Terraform.


## 2.1.7 - 2018-05-03

### Added
- Added support for virtual host names for Load balancer listeners. See [listeners](https://github.com/oracle/terraform-provider-oci/blob/master/docs/load_balancer/listeners.md), [hostnames](https://github.com/oracle/terraform-provider-oci/blob/master/docs/load_balancer/hostnames.md) for more details. 

## 2.1.6 - 2018-04-26

### Added
- New features for images -
     - Image launch mode can be specified when creating an image
     - The image size can be read from image resources and data sources
     - Image data sources can query using a “shape” filter
- New features for boot volumes -     
     - Custom instance boot volume sizes can be specified at launch time
     - Launch options can be read from instance and image resources and data sources
- New features for block volumes -
     - Volume attachments can enable CHAP authentication for iSCSI attachments
     - Volume attachments can be specified as read-only
     - Paravirtualized volume attachments can be created
     - Volume backups can specify whether a full or incremental backup type should be created
 - Filters support all Terraform primitives (string, bool, int, float)
 - Imports for Load Balancer resource are now enabled
 
### Fixed
- Fixed policy version_date bug (#508)
     
## 2.1.5 - 2018-04-12

### Added
- New features for Instances
    - Add “preserve_boot_volume” attribute for preserving attached boot volume on destroy.
    - Add “source_details” attribute for specifying either an image or an existing boot volume when launching.
    - More details can be found [here](https://github.com/oracle/terraform-provider-oci/blob/master/docs/core/instances.md).
- Added support for Local VCN Peering. More details can be found [here](https://github.com/oracle/terraform-provider-oci/blob/master/docs/core/local_peering_gateways.md).
- DNS service integration: adds Zone and Record resources, datasources, documentation and basic examples. More details can be found [here](https://github.com/oracle/terraform-provider-oci/tree/master/docs/dns).

### Deprecated 
- Instances: The “image” attribute is now deprecated. Please use the “source_details” with “source_type” set to “image” instead.

## 2.1.4 - 2018-04-09

### Added
- Add support for Public IPs. More details can be found [here](https://github.com/oracle/terraform-provider-oci/blob/master/docs/core/public_ips.md).

## 2.1.3 - 2018-03-29

### Added
- Added export set resource to File Storage Service. Users can now update FSSTAT related parameters on the export set resource.

### Notes
- Support a new resource name for load balancer backend set that is consistent with other resources. The new name is 'oci_load_balancer_backend_set'. The previous usage of 'oci_load_balancer_backendset' is still supported.

## 2.1.2 - 2018-03-26

### Added
- File Storage Service: Allows management of NFS filesystems, mount targets, exports, and snapshots. (#440)
More details can be found [here](https://github.com/oracle/terraform-provider-oci/blob/master/docs/file_storage).
- Load Balancer PathRouteSets: Added support for load balancer request routing using [path route sets](https://github.com/oracle/terraform-provider-oci/blob/master/docs/load_balancer/path_route_sets.md). (#434)
- Load Balancer Listeners: Added [connection_configuration](https://github.com/oracle/terraform-provider-oci/blob/master/docs/load_balancer/listeners.md) attribute for specifying idle timeouts. (#425)
- Instance Principals: Allows Terraform OCI provider running within an authorized instance to reach Oracle Cloud Infrastructure services.
More details can be found [here](https://github.com/oracle/terraform-provider-oci/blob/master/docs/Writing%20Terraform%20configurations%20for%20OCI.md).

### Fixed
- Load Balancer Certificates: `passphrase` and `private_key` attributes are now marked as Sensitive. (#447)
- Load Balancer work request failures now include extra error details from the service.

## 2.1.1 - 2018-03-14

### Fixed
- VolumeAttachment: Handle unsupported attachment types. If an unsupported attachment type is returned by the service, the SDK's base interface is used to populate common fields.
- Instances: Add missing state field to datasource.

## 2.1.0 - 2018-03-08
More details for the changes introduced in 2.1.0 can be found [here](https://github.com/oracle/terraform-provider-oci/wiki/Details-for-v2.1.0-Release)

### Added
- [Client side filtering](https://github.com/oracle/terraform-provider-oci/blob/master/docs/Filters.md) is now enabled for all data sources that return a list.
- Some Core data sources now support server side filtering by `display_name` and `state`.
- New optional parameters and fields have been added to existing resources and data sources to support new functionality added by the services.
- Documentation files have been updated and improved. Documentation files for resources and data sources of the same entity have now been consolidated into one file.

### Deprecated
- `limit` and `page` parameters in data sources have been deprecated. All list data sources loop through all the pages and return one aggregated list.  
- The `time_modified` field was deprecated from a few resources as it is no longer set by the service.

### Fixed
- Updates to fields in `oci_objectstorage_preauthrequest` resource will force the destruction and recreation of the resource. Updates to fields in this resource had no effect earlier.
- Updating some fields resulted in nothing happening. This has been fixed.
- Unexpected destruction and recreation of `oci_objectstorage_object` was fixed by constraining all keys in the `metadata` map to be lower case.

### Notes
- With this release we started using the new official [OCI Go SDK](https://github.com/oracle/oci-go-sdk). Widespread changes to the source code were needed to make this happen.
- Removing optional parameters from a created resource will not result in a difference and the value for that field will remain as it was. If you want to reset the field to the default value returned by the service for that field you will have to taint the resource to destroy it and recreate it. 
- If upgrading the OCI provider from v1.x.x, see [this wiki](https://github.com/oracle/terraform-provider-oci/wiki/Oracle-Terraform-Provider-Name-Change) for migration steps.
- See docs for this version [here](https://github.com/oracle/terraform-provider-oci/tree/v2.1.0).

## 2.0.7 - 2018-02-08

### Added
- NA

### Fixed
- Correctly resolve Load Balancer and Listener creation failures so plans can be reapplied (#414 and #430).
- Allow Object Storage Buckets to be renamed in plans by implementing the correct ForceNew behavior (#424).

### Notes
- If upgrading from v1, see [this wiki](https://github.com/oracle/terraform-provider-oci/wiki/Oracle-Terraform-Provider-Name-Change) for migration steps.
- See docs for this version [here](https://github.com/oracle/terraform-provider-oci/tree/v2.0.7).

## 2.0.6 - 2018-01-08

### Added
- A minimum of TLS 1.2 is now enforced by the provider (#394)

### Fixed
- Fixed an issue where importing a default resource would leave the manage_default_resource_id empty in the state file during import of default resources (#393, #379)

### Notes
- If upgrading from v1, see [this wiki](https://github.com/oracle/terraform-provider-oci/wiki/Oracle-Terraform-Provider-Name-Change) for migration steps.
- See docs for this version [here](https://github.com/oracle/terraform-provider-oci/tree/v2.0.6).

## 2.0.5 - 2017-12-14

### Added
- Enhanced security options by adding support for source port range under security list rules. This can be specified in "tcp_options" and "udp_options" (#340).
- Allow configuration of default resources under VCNs (#374). See more details about this feature [here](https://github.com/oracle/terraform-provider-oci/blob/master/docs/Managing%20Default%20Resources.md).

### Fixed
- Fixed bug wherein policy was not destroyed and recreated when compartment is changed (#389)
- Fixed errors with terraform import because of missing vcn_id in `*.tfstate` files (internet_gateway, route_tables, dhcp_options) (#388, #379)
- Fixed error where same retry token was being used for multiple requests in some development environments when auto retries were activated (Issue #170)

### Notes
- Code refactoring was done as part of this release. Go source file names have changed, the `provider` directory has been added. Should not impact the users in any way.
- If upgrading from v1, see [this wiki](https://github.com/oracle/terraform-provider-oci/wiki/Oracle-Terraform-Provider-Name-Change) for migration steps.
- See docs for this version [here](https://github.com/oracle/terraform-provider-oci/tree/v2.0.5).

## 2.0.4 - 2017-11-2

### Added
- Host header and version to signing (#340)
- Support for block volume fast clones (#347)

### Fixed
- Examples of "oci_core_images" data source now filter on "display_name" to accommodate changes to available images (#342 and #345)

### Notes
- If upgrading from v1, see [this wiki](https://github.com/oracle/terraform-provider-oci/wiki/Oracle-Terraform-Provider-Name-Change) for migration steps.
- See docs for this version [here](https://github.com/oracle/terraform-provider-oci/tree/v2.0.4).

## 2.0.3 - 2017-10-26

### Added
- Filters for most core, IAM, and Load Balancer data sources. See [docs/Filters.md](https://github.com/oracle/terraform-provider-oci/blob/master/docs/Filters.md) for details.
- Support for Virtual Machine (VM) DB Systems
- Support for Bring Your Own License (BYOL) licensing model for DB Systems

### Notes
- If upgrading from v1, see [this wiki](https://github.com/oracle/terraform-provider-oci/wiki/Oracle-Terraform-Provider-Name-Change) for migration steps.
- See docs for this version [here](https://github.com/oracle/terraform-provider-oci/tree/v2.0.3).

## 2.0.2 - 2017-10-12

### Fixed
- Optimize service error retry behavior (#179)
- Object store fixes (#225)
- Properly handle version date in policies, ignore format changes when diffing (#230)
- Ignore case for DNS Labels (#279)
- Oci-tool migration tool fixes (#298) (#292)

### Added
- Support update and refresh on Instance and Vnic details
- File upload example
- Block volumes support for size in gigabytes (#297)
- Support for compartment renaming (#250)

### Changed
- Handle and log URL parsing errors (#277)
- Minor update to bmcs-go-sdk license
- Acceptance test refinements

### Notes
- If upgrading from v1, see [this wiki](https://github.com/oracle/terraform-provider-oci/wiki/Oracle-Terraform-Provider-Name-Change) for migration steps.
- See docs for this version [here](https://github.com/oracle/terraform-provider-oci/tree/2.0.2).

## 2.0.1 - 2017-9-26

### Fixed
- Resources are now removed from the state file if in a "terminated" state so that it is recreated on an apply (#113)
- Enable empty route rules (#68)
- Fix import of Subnet prohibit_public_ip_on_vnic
- Adds pagination to all IAM data sources
- General fixes for plans including compartments as a resource

### Added
- VNIC skip_source_dest_check property

### Notes
- If upgrading from v1, see [this wiki](https://github.com/oracle/terraform-provider-oci/wiki/Oracle-Terraform-Provider-Name-Change) for migration steps.
- See docs for this version [here](https://github.com/oracle/terraform-provider-oci/tree/v2.0.1).

## 2.0.0 - 2017-9-13

### Changed
- Changes name from terraform-provider-baremetal to terraform-provider-oci. See [this wiki](https://github.com/oracle/terraform-provider-oci/wiki/Oracle-Terraform-Provider-Name-Change) on migration steps and associated migration tool usage instructions.

### Added
* Support for Secondary Private IPs

### Notes
- If upgrading from v1, see [this wiki](https://github.com/oracle/terraform-provider-oci/wiki/Oracle-Terraform-Provider-Name-Change) for migration steps.
- See docs for this version [here](https://github.com/oracle/terraform-provider-oci/tree/v2.0.0).

## Earlier Versions
- For earlier versions, see [releases](https://github.com/oracle/terraform-provider-oci/releases).
