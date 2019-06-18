## 3.30.1 (Unreleased)
## 3.30.0 (June 19, 2019)

### Added
- Support for scheduling KMS key deletion
- Support for moving Volumes, Volume groups, Boot Volumes and corresponding Backups across compartments
- Support for moving Service Gateway resource across Compartments

### Fixed
- Instance `create_vnic_detials` will be fetched for all applicable instance lifecycle states.

## 3.29.0 (June 12, 2019)

### Added
- Support for autonomous transaction database-dedicated, autonomous exadata infrastructures, autonomous container databases and maintenance runs.
- Support for `boot_volume_size_in_gbs` argument in the `oci_instance_configuration` resource 

## 3.28.2 (June 07, 2019)

### Added
- `oci_core_ipsec_connection_tunnel_management` resource to manage IPSec tunnel connection
### Fixed
- `oci_core_ipsec` backward compatibility issue by removing `tunnel_configuration` property, which is reported by https://github.com/terraform-providers/terraform-provider-oci/issues/779

## 3.28.1 (June 05, 2019)

## Notes

- This is a Terraform 0.12 compatible release of this provider.

## 3.28.0 (June 05, 2019)

### Added
- Support for ATP-S autoscaling
- Support for specifying Fault Domains in `launch_details` for `oci_core_instance_configuration` resource
- Support for defined tags and tag namespace deletion

## 3.27.0 (May 29, 2019)

### Added
- Support for moving File Systems and Mount Targets across compartments
- Support for filtering File Storage resources by tags
- Support for getting UI password information

### Notes
- This is the first provider version that supports Terraform v0.12.

## 3.26.0 (May 22, 2019)

### Added
- Support for setting `compartment_id` argument in `object_storage_namespace` data source
- Support BGP dynamic routing and allow customer to input PSK for IPSec tunnels
- ListInstanceConfig/Pools and ListAutoScalingConfiguration return tags

### Fixed
- Fix for dbSystem `db_version` causing unnecessary diffs on subsequent applies
- Fix for database `db_backup_config` causing unnecessary diffs on subsequent applies.

## 3.25.0 (May 15, 2019)

### Added
- Support for recovery window in backup config for Database DbSystem and DbHome resources
- Support for KMS throttling and audit logs

## 3.24.1 (May 07, 2019)

### Fixed
- Fix unhandled error when Security Lists are altered outside Terraform
- Updated `availability_domain` property to be case insensitive

## 3.24.0 (April 24, 2019)

### Added
- Support data source for cost tracking tags
- Singular data sources will reuse resource schema

## 3.23.0 (April 17, 2019)

### Added
- Support for updating `license_model` for `oci_autonomous_database` resource
- Support for updating `static_routes` and new `cpe_local_identifier` in `oci_core_ipsec` resource for improved VPN service usability
- Support for updating `whitelisted_ips` in `autonomous_database`. Note: Cannot be used during creation.
- Support tagging for Dynamic Groups in Identity

## 3.22.0 (April 10, 2019)

### Added
- Support for `compartment_id` filter in `email_senders` and `email_suppressions` data sources
- Support for import in dbHomes and dbSystems

### Fixed
- Backward compatibility for compositeId in Object Storage - Objects and PARs

## 3.21.0 (April 03, 2019)

### Added
- Support for additional dbHomes/databases in a BM Db System
- Support for tags in databases
- Support for updates to database auto_backup_enabled
- Support for provider service keys in Fast Connect Provider Services
- Singular data sources for User, Group, File Storage Snapshot, Private IP and Virtual Cloud Network (VCN).
- Support for authentication policy introduced in v3.18.0 is now generally available.

### Fixed
- Virtual Circuit update failures by handling default values
- Importing `assign_public_ip` for Core vnic attachment

## 3.20.0 (March 27, 2019)

### Added
- Support for importing Buckets and Pre-authenticated requests in Object Storage
- Support glob inclusion and exclusion patterns for object names allowed in Object Storage Lifecycle
- Support for sorting for resources returned in `oci_core_images` data source
- Support for Web Application Acceleration and Security service

### Fixed
- Import functionality for Objects in Object Storage
- Import functionality for Identity Policy

## 3.19.0 (March 20, 2019)

### Added
- Support for cloning of Autonomous Databases
- Support for node metadata in container engine node pool
- Support for Data Guard Associations for databases

## 3.18.0 (March 13, 2019)

### Added
- Add Budget and Alert Rules resources
- Support starting and stopping instances
- Support to create Containerengine Node Pool with Image Id
- Support for customer specified timezone in Database Systems
- Support for creating Autonomous Data Warehouses through Autonomous Database resource `oci_database_autonomous_database` using the field `db_workload`
- Support for Defined Tag defaults through the `oci_identity_tag_default` resource
- Support for updating the compartment on a Tag Namespace
- Support for exadata io resource management config for DB system
- Support `email` attribute for `oci_identity_user` resource
- Support for authentication policy

### Fixed
- Marked oci_identity_ui_password resource as not importable

### Deprecated
- Deprecated Autonomous Data Warehouse resources `oci_database_autonomous_data_warehouse`, the API is now unified with Autonomous Database

## 3.17.0 (March 05, 2019)

### Added
- Add singular Availability Domain data source with related example updates
- Support for Monitoring service
- Adding ability to disable monitoring in instances
- Adding support for Metrics-based Dynamic Auto-scaling
- Support for listing and specifying Fault Domains in Database resources
- Support for Notification service

## 3.16.0 (February 26, 2019)

### Added
- Adding description property to rules in Steering Policies in DNS
- Enable regional Subnets by making Availability Domain optional when creating a Subnet
- Support for Streaming service
- Support for the tagging of applicable KMS resources

### Fixed
- DNS Record now requires domain and rtype as mandatory arguments. Managing DNS record resources now requires DNS_RECORD* level policy entitlements instead of DNS_ZONE*. [Permissions List](https://docs.cloud.oracle.com/iaas/Content/Identity/Reference/dnspolicyreference.htm)

## 3.15.0 (February 12, 2019)

### Added
- Adding support for the tagging of Email Delivery service approved senders
- Support for Health Check Service
- Adding database connection information to the `oci_database_database` and `oci_database_databases` data sources
- Adding support for Steering Policies in DNS

## 3.14.1 (February 05, 2019)

### Fixed
- Timeout should be updatable for the `oci_containerengine_cluster` and `oci_containerengine_node_pool` resources
- Virtual Circuit `public_prefixes` to be updatable and importable. [Issue #700](https://github.com/terraform-providers/terraform-provider-oci/issues/700)

## 3.14.0 (January 29, 2019)

### Added
- Adding support for the database renaming during restore from incremental backup

## 3.13.0 (January 23, 2019)

### Added
- Added singular data source for Object Storage objects

### Fixed
- Fixed an issue where the default retry timeout is zero seconds if `retry_duration_seconds` isn't specified
- Modifying immutable `metadata` fields such as `ssh_authorized_keys` and `user_data` should result in new instances. [Issue #673](https://github.com/terraform-providers/terraform-provider-oci/issues/673)
- Vendored Terraform helper/schema SDK to return matching data type for maps in case of empty state. [Issue #685](https://github.com/terraform-providers/terraform-provider-oci/issues/685)

## 3.12.0 (January 15, 2019)

### Added
- Support for `retry_duration_seconds` option to configure length of retry in the face of HTTP 429 and 500 errors
- Support for custom header insertion, extension, and removal for Load Balancer listener resource
- Support for consistent volume names in the Block Volume attachments

### Fixed
- Retried SDK calls are now jittered to avoid herding of retry requests in high parallelism scenarios
- Fail the initialization of the provider if either of `user_ocid`, `fingerprint`, `private_key`, `private_key_path` or `private_key_password` are specified for `InstancePrincipal` or `InstancePrincipalWithCerts` auth mode.

### Note
- Examples and test updated to use VM.Standard2.1
- Windows example image updated to Windows-Server-2012-R2-Standard-Edition-VM-Gen2-2018.12.12-0

## 3.11.2 (January 10, 2019)

### Fixed
- Reverted previous fix for immutable `metadata` fields `ssh_authorized_keys` and `user_data` that results in new instances due to a crash when using interpolations in TypeMap with customdiff (Issue #685)

## 3.11.1 (January 08, 2019)

### Changed
- LoadBalancer BackendSets to have TypeSet for Backends to avoid out of order diffs

### Fixed
- Regression in handling of failed work-requests to pass the errors to the user and fail the apply
- Removing certificates from load balancer listeners can be done by omitting `ssl_configuration`
- Load balancer resources that are stuck in failed state during deletion can now be deleted after upgrading
- Modifying immutable `metadata` fields such as `ssh_authorized_keys` and `user_data` should result in new instances

## 3.11.0 (December 18, 2018)

### Added
- Support for tagging in `oci_dns_zone`
- New attribute `nameservers` is added to `oci_dns_zone`
- Support for in-transit encryption for paravirtualized boot and data attachment
- Identify latest database version with `oci_databse_db_versions` data source using `is_latest_for_major_version` property
- Support for importing tag. Note tag uses custom Id(import only) format (tagNamespaces/{tagNamespaceId}/tags/{tagName}) to support import.
- Support for provisioning user capabilities for native and federation shadow users
- Support `id` attribute for `oci_identity_availability_domains`
- Support `freeform_attributes` attribute for the `oci_identity_identity_provider`
- Support for `sparse_diskgroup` for Exadata dbsystem

## 3.10.0 (December 11, 2018)

### Added
- Support for attaching Route Table to Subnet. Issue [#270](https://github.com/terraform-providers/terraform-provider-oci/issues/270)

## 3.9.0 (December 04, 2018)

### Added
- Support for the Instance Pools & Instance Configurations
- Support for the Block Volume cross-region backups
- Support for 'approximate_count' and 'approximate_size' for bucket resource

## 3.8.0 (November 28, 2018)

### Added
- Support VCN Transit

## 3.7.0 (November 14, 2018)

### Added
- New parameter `is_hydrated` in `oci_core_volume_groups` resource and data source
- Support for public IP prefixes (CIDRs) up to 31
- Support for tagging in `oci_file_storage_file_system`, `oci_file_storage_mount_target`, and `oci_file_storage_snapshot`

### Changed
- Make `route_table_id`, `dhcp_options_id` in `oci_core_subnet` updatable
- Make `security_list_ids` in `oci_core_subnet` optional and updatable

### Deprecated
- Volumes: The `backup_policy_id` attribute is now deprecated. Backup policy should be assigned through `volume_backup_policy_assignments` resource instead.
- BootVolumes: The `backup_policy_id` attribute is now deprecated. Backup policy should be assigned through `volume_backup_policy_assignments` resource instead.

## 3.6.0 (November 01, 2018)

### Added
- New parameters `db_name` and `state` in `oci_database_database` data source
- New parameters `display_name` and `state` in `oci_database_db_homes` data source
- New parameter `state` parameter in `oci_database_db_nodes` data source
- New parameters `availability_domain`, `display_name`, and `state` in `oci_database_db_systems` data source
- Support for Partner Image Catalog
- Support for Key Management Service
- Support for encrypting the contents of an Object Storage bucket using a Key Management Service key
- Support for specifying a Key Management Service key when launching a compute instance in the Compute service
- Support for specifying a Key Management Service key when backing up or restoring a block storage volume in the Block Volume service
- Support enabling cost tracking for tags using `is_cost_tracking` field
- Support returning maintenance reboot time for compute instances using `time_maintenance_reboot_due` field
- Support nesting and deleting compartments. Compartment delete requires opt in, see compartment documentation

### Fixed
- Data type for properties with type as TypeSet to TypeList in following datasources: `oci_core_route_tables`, `oci_core_security_lists`, `oci_core_volume`, and `oci_core_service_gateways` to allow referencing by indexes in Terraform configs.

## 3.5.0 (October 19, 2018)

### Added
- Support for [Cross Region Copy](https://docs.cloud.oracle.com/iaas/Content/Object/Tasks/copyingobjects.htm) of objects
- Support for object lifecycle policies on a bucket on object storage. See [Using Object Lifecycle Management](https://docs.cloud.oracle.com/iaas/Content/Object/Tasks/usinglifecyclepolicies.htm)
- Support for singular data source for a bucket
- Additional nested field in `oci_database_backups` data source and `oci_database_backup` resource, under the `backups` property called `database_size_in_gbs`
- Support for generating and downloading wallets for Autonomous Database and Autonomous Data Warehouse. See [Connecting to Autonomous Data Warehouse](https://docs.cloud.oracle.com/iaas/Content/Database/Tasks/adwconnecting.htm) for more details.

### Changed
- Nested field in `oci_database_backups` data source and `oci_database_backup` resource, under the `backups` property called `db_data_size_in_mbs` marked as deprecated

## 3.4.0 (October 11, 2018)

### Added
- Support for clone and resize of Boot Volume
- Support for specifying a backup policy at the time of creating a Boot Volume
- Support for offline resizing of Boot Volume
- Support for tagging of Boot Volume
- Support for NAT Gateways
- Support for singular data sources that can query individual Volumes, Subnets, and Instances
- Fields "assigned_entity_id" and "assigned_entity_type" to Public IPs to allow distinguishing Public IPs of the NAT Gateway.

### Fixed
- Importing of volumes with backup policies. Issue [#590](https://github.com/terraform-providers/terraform-provider-oci/issues/590)
- Updating of Virtual Circuits fails with field bgpMd5AuthKey is not supported

## 3.3.0 (October 04, 2018)

### Added
- Support for new Image launch mode: paravirtualization

### Fixed
- Fix logic to prevent unexpected diffs related to numbers. Issue [#607](https://github.com/terraform-providers/terraform-provider-oci/issues/607)

## 3.2.0 (September 28, 2018)

### Added
- Support updating size of offline volumes

### Fixed
- Specifying lifecycle state in container engine cluster datasource properly filters. Issue [#600](https://github.com/terraform-providers/terraform-provider-oci/issues/600)
- Importing the assign_public_ip attribute for instances has the correct default. Issue [#593](https://github.com/terraform-providers/terraform-provider-oci/issues/593)
- ADW and ATP resources destruction still succeeds if the database lifecycle state becomes `Unavailable`

## 3.1.1 (September 21, 2018)

### Fixed
- Fixed bug with load balancer compositeId. Issue [#612](https://github.com/oracle/terraform-provider-oci/issues/612)

## 3.1.0 (September 20, 2018)

### Added
- Support for importing load balancer related resources such as backend, backend set, hostname, listeners, and path route sets
- Support for updating an instance's metadata and extended metadata

## 3.0.0 (September 17, 2018)

### Fixed
- Fixed bug with DNS Records when the user specified more than 50 records in a terraform config. Issue [#581](https://github.com/oracle/terraform-provider-oci/issues/581)

### Notes
- This is the first provider version that can be automatically downloaded and installed with the `terraform init` command.

## 2.2.4 - 2018-09-11

### Added
- Support for Autonomous Data Warehouse and manual backups
- Support for Autonomous Transaction Processing (a.k.a Autonomous Database) and manual backups

## 2.2.3 - 2018-09-06

### Added
- Support for specifying a backup policy at the time of creating a Volume

## 2.2.2 - 2018-08-30

### Added
- Support for listing Fault Domains in an AD and specifying them when launching an Instance


## 2.2.1 - 2018-08-23

### Added
- Support for Boot Volume Backups. See [Boot Volume Backup Resources](https://github.com/oracle/terraform-provider-oci/blob/master/docs/core/boot_volume_backups.md) and [Backing Up a Boot Volume](https://docs.cloud.oracle.com/iaas/Content/Block/Tasks/backingupabootvolume.htm)
- Support for efficient large file uploads in Object Storage using multi-part API by providing `source` path. See [Object Resources](https://github.com/oracle/terraform-provider-oci/blob/master/docs/object_storage/objects.md) and [Using Multipart Uploads](https://docs.cloud.oracle.com/iaas/Content/Object/Tasks/usingmultipartuploads.htm)

## 2.2.0 - 2018-08-09

### Fixed
- Fix to security lists to avoid diffs after an apply in certain cases (#565)

### Added
- Support Audit Events Data Source
- Support for export options in the File Storage service for improved access controls
- Support for tagging on Load Balancer Resource. See [Tagging Resources](https://github.com/oracle/terraform-provider-oci/blob/master/docs/Tagging%20Resources.md)
- Support for large integers (int64) on `oci_core_volume.size_in_gbs`, `load_balancer_listener.idle_timeout_in_seconds`, `oci_file_storage_export_set.max_fs_stat_bytes`, and `oci_file_storage_export_set.max_fs_stat_files` inputs
- Include additional exported attributes related to computed sizes in [VolumeGroup](https://github.com/oracle/terraform-provider-oci/blob/master/docs/core/volume_groups.md) and [VolumeGroupBackup](https://github.com/oracle/terraform-provider-oci/blob/master/docs/core/volume_group_backups.md)

### Notes
- This release updates the OCI Provider code dependencies to Terraform v0.11.7, the result is that users with Terraform binary versions earlier than v0.10.1 will need to update--we recommend using the latest 0.11.x binary

## 2.1.17 - 2018-08-02

### Fixed
- Fix bug that was causing creation of tags and tagging namespaces to fail (#562)

## 2.1.16 - 2018-07-19

### Added
- Support for [Container Engine for Kubernetes](https://docs.cloud.oracle.com/iaas/Content/ContEng/Concepts/contengoverview.htm), adding resources for clusters, node pools, and data source for [kubeconfig](https://docs.cloud.oracle.com/iaas/Content/ContEng/Tasks/contengdownloadkubeconfigfile.htm)
- Support for [FastConnect](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/fastconnect.htm), cross-connect group and virtual circuits resources and data sources

## 2.1.15 - 2018-07-13

### Fixed
- Fix bug introduced in v2.1.14 (#558), failure updating a Route Table's Route Rules when they contain a rule that includes a Service Gateway ID

## 2.1.14 - 2018-07-13

###Notes
_This build contains a known issue where updates to a Route Table's Route Rules (when they contain a rule that includes a Service Gateway ID) fail with a 400 service error code (#558). The issue is fixed in v2.1.15._

### Added
- Ability to create and manage email approved senders, suppressions, and SMTP credentials
- Adding Service Gateway resource and data source, update Route Table and Security List
- Add Audit service configuration resource
- Support Identity Federation

### Changed
- Users may notice larger diffs for Security List's `ingress_security_rules`, `egress_security_rules` and Route Table's `route_rules`. The internal representation has been changed from Lists to Sets, which results in unexpected but innocuous Terraform behavior. See this issue for discussion: https://github.com/hashicorp/terraform/issues/15180
- Default timeout changed from 5 minutes to 15 minutes to accommodate some resources that may take longer to succeed
- Ability to update compartment of an Object Storage Bucket
- Updated Database data source to support tags

### Fixed
- Delete behavior fixed on Load Balancer resources for failed work requests

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
- Added support for remote VCN peering. More details can be found [here](https://github.com/oracle/terraform-provider-oci/blob/master/docs/core/remote_peering_connections.md), and an example [here](https://github.com/oracle/terraform-provider-oci/blob/master/examples/networking/remote_vcn_peering_full).
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
