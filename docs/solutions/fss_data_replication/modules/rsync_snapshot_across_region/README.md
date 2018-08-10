# FSS File System data-sync Across Regions

This repository contains a [Terraform](https://www.terraform.io/) module responsible to deploy a host on [Oracle Cloud Infrastructure (OCI)](https://cloud.oracle.com/en_US/cloud-infrastructure) and replicate the data across two [Oracle Cloud Infrastructure File Storage Service (FSS)](https://docs.us-phoenix-1.oraclecloud.com/Content/File/Concepts/filestorageoverview.htm) shared [File Systems](https://docs.us-phoenix-1.oraclecloud.com/Content/File/Tasks/creatingfilesystems.htm) located on different [OCI Region](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/regions.htm). This module is designed to create a [Snapshot](https://docs.us-phoenix-1.oraclecloud.com/Content/File/Concepts/filestorageoverview.htm#concepts) and copy the data directly from the `source` FSS File System Snapshot folder to a `destination` FSS File System using cron job in conjunction with [rsync](https://en.wikipedia.org/wiki/Rsync).

Notes:

1. This process is recommended for data syncing any sort of transactional data or frequent modified data since the copy is originated from the snapshot.

2. Snapshots provide a consistent, point-in-time view of your file system. Each snapshot reflects only data that changed from the previous snapshot.

3. Remote Peering is recommended to establish network connectivity between two VCNs located on different regions.

4. There is a lock mechanism (flock) to avoid running another rsync process while the previous one has not completed.

5. rsync runs with the following parameters. Example:
 `rsync -aHAXxv --numeric-ids --delete /mnt/src-fs/ /mnt/dst-fs/ `

* List of parameters:

  * `a`: This parameter means you want recursion and want to preserve almost everything.
  * `H`: This tells rsync to look for hard-linked files in the transfer and link together the corresponding files on the receiving side.
  * `A`: This option causes rsync to update the destination ACLs to be the same as the source ACLs.
  * `X`: This option causes rsync to update the remote extended attributes to be the same as the local ones.
  * `x`: This tells rsync to avoid crossing a filesystem boundary when recursing. This does not limit the user's ability to specify items to copy from multiple filesystems, just rsync's recursion through the hierarchy of each directory that the user specified, and also the analogous recursion on the receiving side during deletion.
  * `v`: This option increases the amount of information the daemon logs during its startup phase.
  * `--numeric-ids`: With this option rsync will transfer numeric group and user IDs rather than using user and group names and mapping them at both ends.
  * `--delete`: This tells rsync to delete extraneous files from the receiving side (ones that aren't on the sending side), but only for the directories that are being synchronized.

## How to use this Module

This folder defines a [Terraform module](https://www.terraform.io/docs/modules/usage.html), which you can use in your code by adding a `module` configuration and setting its `source` parameter to URL of this folder:

```hcl
module "rsync_snapshot_across_ashburn_phoenix" {
  # TODO: Update this to the final URL
  source = "git::ssh://git@orahub.oraclecorp.com/pts-cloud-dev/terraform-modules//terraform-oci-fss/modules/rsync_snapshot_across_region"
  

  providers = {
      "oci.src" = "oci.iad"
      "oci.dst" = "oci.phx"
    }

  # Specify SSH Private and Public Keys
  ssh_private_key_path = "${var.ssh_private_key_path}"
  ssh_public_key_path  = "${var.ssh_public_key_path}"

  #Specify IAM Settings
  compartment_id = "${var.compartment_id}"

  #Specify Network Settings for Source Host
  src_availability_domain = "${lookup(data.oci_identity_availability_domains.ads_iad.availability_domains[0],"name")}"
  src_subnet_id           = "${oci_core_subnet.subnet_iad_ad1.id}"

  #Specify Instance Settings for Source Host
  src_instance_hostname   = "rsync_snap_reg_iad"
  src_instance_shape      = "VM.Standard2.16"
  src_instance_image = "${var.instance_image_ocid[var.region_iad]}"

  #Specify Network Settings for Target Host
  dst_availability_domain = "${lookup(data.oci_identity_availability_domains.ads_phx.availability_domains[0],"name")}"
  dst_subnet_id           = "${oci_core_subnet.subnet_phx_ad1.id}"

  #Specify Instance Settings for Target Host
  dst_instance_hostname   = "rsync_snap_reg_phx"
  dst_instance_shape      = "VM.Standard2.16"
  dst_instance_image = "${var.instance_image_ocid[var.region_phx]}"

  //FSS configuration
  src_export_path             = "${var.src_export_path}"
  src_mount_target_private_ip = "${local.src_mt_private_ip_iad_ad1}"

  dst_export_path             = "${var.dst_export_path}"
  dst_mount_target_private_ip = "${local.dst_mt_private_ip_phx_ad1}"

  data_sync_frequency        = "*/5 * * * *"
  snapshot_frequency = "*/30 * * * *"
}
```

Input Parameters
----------------

* `source`: Use this parameter to specify the URL of the FSS rsync filesystem (local) module. The double slash (`//`) is intentional and required. Terraform uses it to specify subfolders within a Git repo (see [module sources](https://www.terraform.io/docs/modules/sources.html)).

* `providers oci.src`: Specify the [provider alias](https://www.terraform.io/docs/configuration/providers.html) used to map the source region.
* `providers oci.dst`: Specify the [provider alias](https://www.terraform.io/docs/configuration/providers.html) used to map the target region.
* `src_instance_hostname`: Use this parameter to specify the hostname of the FSS Data Sync source host.

* `src_instance_shape`: Use this parameter to specify the [Compute Shape](https://docs.us-phoenix-1.oraclecloud.com/Content/Compute/Concepts/computeoverview.htm) of the FSS Data Sync source host. Network Bandwidth available for the instance depends on the Compute Shape.

* `src_instance_image`: Use this parameter to specify the [OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the [OCI Image](https://docs.us-phoenix-1.oraclecloud.com/images/) used to run the FSS Rsync source host.

* `dst_instance_hostname`: Use this parameter to specify the hostname of the FSS Data Sync target host, located on a different VCN/Region from the source host.

* `dst_instance_shape`: Use this parameter to specify the [Compute Shape](https://docs.us-phoenix-1.oraclecloud.com/Content/Compute/Concepts/computeoverview.htm) of the FSS Data Sync target host, located on a different VCN/Region from the source host. Network Bandwidth available for the instance depends on the Compute Shape.

* `dst_instance_image`: Use this parameter to specify the [OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the [OCI Image](https://docs.us-phoenix-1.oraclecloud.com/images/) used to run the FSS Rsync source host, located on a different VCN/Region from the source host.

* `src_export_path`: Use this parameter to specify the *source* FSS File System [Export Path](https://docs.us-phoenix-1.oraclecloud.com/Content/File/Concepts/filestorageoverview.htm#concepts) used by the data sync process. Export path is appended to the mount target IP address and is used to mount to the file system. Example: `/src_fs_demo-iad-ad1`.

* `src_mount_target_private_ip`: Use this parameter to specify the *source* FSS File System [Mount Target](https://docs.us-phoenix-1.oraclecloud.com/Content/File/Concepts/filestorageoverview.htm#concepts) IP address used by the data sync process.

* `dst_export_path`: Use this parameter to specify the *destination or target* FSS File System [Export Path](https://docs.us-phoenix-1.oraclecloud.com/Content/File/Concepts/filestorageoverview.htm#concepts) used by the data sync process. Export path is appended to the mount target IP address and is used to mount to the file system. Example: `/dest_fs_demo-iad-ad1`.

* `dst_mount_target_private_ip`: Use this parameter to specify the *destination (target)* FSS File System [Mount Target](https://docs.us-phoenix-1.oraclecloud.com/Content/File/Concepts/filestorageoverview.htm#concepts) IP address used by the data sync process.

* `data_sync_frequency`: Use this parameter to specify the frequency that data sync will take place. Syntax is based on Linux [Cron job](https://en.wikipedia.org/wiki/Cron). Example: `*/30 * * * *` - this will run the data sync job every 30 minutes.

* `snapshot_frequency`: Use this parameter to specify the frequency that a Snapshot will take place. Syntax is based on Linux [Cron job](https://en.wikipedia.org/wiki/Cron). Example: `@hourly` - this will run generate a snapshot every hour. Snapshot folder will be created based on the following naming standard: ```tf-fss-snapshot-`date -u +%Y-%m-%dT%H` ```

You can find the other optional parameters in [variables.tf](variables.tf).

Check out the [main example](../../README.md) for fully-working sample code.

Output Parameters
-----------------

* `data_sync_src_public_ip`: Public IP of the data sync source host.
* `data_sync_src_private_ip`: Private IP of the data sync source host.
* `data_sync_dst_public_ip`: Public IP of the data sync target host.
* `data_sync_dst_private_ip`: Private IP of the data sync target host.

Example: 

```hcl
module "rsync_snapshot_across_region" {

# TODO: Update this to the final URL
  source = "git::ssh://git@orahub.oraclecorp.com/pts-cloud-dev/terraform-modules//terraform-oci-fss/modules/rsync_snapshot_across_region"
  ... 
}

output "src_public_ip" {
  value = "${module.rsync_snapshot_across_region.data_sync_src_public_ip}"
}

output "src_private_ip" {
  value = "${module.rsync_snapshot_across_region.data_sync_src_private_ip}"


output "dst_public_ip" {
  value = "${module.rsync_snapshot_across_region.data_sync_dst_public_ip}"
}

output "dst_private_ip" {
  value = "${module.rsync_snapshot_across_region.data_sync_dst_private_ip}"

```

## How do you connect to the Data Sync host?

User can SSH to the data sync host. Both Data Sync Host Private and Public IP addresses are available in the module output. An [SSH Key Pair](https://docs.us-phoenix-1.oraclecloud.com/Content/GSG/Tasks/creatingkeys.htm) should be associated with the Data Sync host for SSH access. The key values are associated with the input variables `ssh_public_key_path` & `ssh_private_key_path` . Make sure you have a security list that opens the SSH port (default: 22).

## What's included in this module?

This module creates two Compute Instances on different regions (which includes VCN, Availability Domain and Subnet). Rsync and cron jobs are configured during the bootstrap process.

## What's NOT included in this module?

This module does NOT handle the following items:

* [Network Settings](#network-settings)
* [File Storage](#file-storage)
* [IAM](#iam)
* [Remote Peering](#remote-peering)

### Network Settings

This module assumes you've already created your network topology (Source and Target VCNs, Subnets, route tables &  Security Lists). You will need to pass in the the relevant info about your network topology (e.g. `subnet_id`, `compartment_id`, `availability_domain`  as input variables to this module.

### File Storage

This module assumes you've already created both `source` and `target` File Systems on different [Regions](https://docs.us-phoenix-1.oraclecloud.com/Content/File/Concepts/filestorageoverview.htm#regions).

### IAM

This module does not creates OCI Compartments, Identity User, Group and Policy.

### Remote Peering

This module does not creates OCI [Remote Peering](https://docs.us-phoenix-1.oraclecloud.com/Content/Network/Tasks/remoteVCNpeering.htm)