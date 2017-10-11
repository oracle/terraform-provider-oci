KVM Installation module
=======================

This Module is responsible for setting up a custom KVM based Virtual Machines on top of a KVM hypervisor on Oracle Cloud Infrastructure. KVM domain is installed as part of this module.


### Reference:
1. [OCI - Installing and Configuring KVM on Bare Metal
Instances with Multi-VNIC](https://docs.us-phoenix-1.oraclecloud.com/Content/Resources/Assets/installing_kvm_multi_vnics.pdf)

2. [OCI - About Secondary Vnics](https://docs.us-phoenix-1.oraclecloud.com/Content/Network/Tasks/managingVNICs.htm)

3. [Terraform template_file data source](https://www.terraform.io/docs/providers/template/d/file.html)

4. [Terraform Modules](https://www.terraform.io/docs/modules/index.html)


### Module required Variables

#### private_key
ssh private key. Example:
private_key="SXSDKFJSKJAFKAJFLAJFLKAF"

#### host
IP Address of the Host (Bare Metal Instance). Example:
host="122.123.123.5"

#### qcow2_image_url
URL of the Guest VM qcow2 file. Example:
qcow2_image_url="https://my-image-url.com/my-image-file-1.0.x.qcow2"

#### qcow2_image_target_path
Location (full path) where the image will be uploaded to on the bare metal host. Example:
qcow2_image_target_path="/kvm-imgs/"

#### qcow2_image_filename
Filename of the image file. Example:
qcow2_image_filename="my-image-file-1.0.x.qcow2"

#### kvm_guest_domain_name
The KVM guest domain name. Example:
kvm_guest_domain_name="MyGuestDomain"

#### kvm_guest_memory
The amount of memory allocated to the Guest VM. Example:
kvm_guest_memory="16384"

#### kvm_guest_vcpu
The amount of VCPU allocated to the Guest VM. Example:
kvm_guest_vcpu="8"

#### kvm_guest_os_type
The OS type. Example:
kvm_guest_os_type="linux"

#### kvm_guest_vnc_port
The VNC port number for accessing the Guest VM. Example:
kvm_guest_vnc_port="5901"

#### kvm_guest_vnc_pwd
The VNC password for accessing the Guest VM. Example:
kvm_guest_vnc_pwd="Test123"

#### kvm_guest_vnic_mac_address
The secondary vnic MAC Address where the Guest VM will be hosted. Example:
"${data.oci_core_vnic.kvm-guest-mgmt-vnic.mac_address}"

#### kvm_guest_vnic_id
The secondary vnic ID where the Guest VM will be hosted. Example:
kvm_guest_vnic_id="${data.oci_core_vnic.kvm-guest-mgmt-vnic.id}"

#### kvm_guest_emulation_mode
The guest vm emulation mode. Example:
kvm_guest_emulation_mode="virtio"



### Sample Usage

```
module "setup-kvm-hypervisor" {
  source = "./modules/kvm-hypervisor"
  private_key = "${file(var.ssh_private_key_path)}"
  host        = "${oci_core_instance.KVM-HOST.public_ip}"
  qcow2_image_url = "${var.kvm_image_url}"
  qcow2_image_target_path = "${var.kvm_image_path}"
  qcow2_image_filename = "${var.kvm_image_name}"
  kvm_guest_domain_name  = "${var.kvm_guest_domain_name}"
  kvm_guest_memory = "${var.kvm_guest_memory}"
  kvm_guest_vcpu = "${var.kvm_guest_vcpu}"
  kvm_guest_os_type = "${var.kvm_guest_os_type}"
  kvm_guest_vnc_port = "${var.kvm_guest_vnc_port}"
  kvm_guest_vnc_pwd = "${var.kvm_guest_vnc_pwd}"

  kvm_guest_vnic_mac_address = "${data.oci_core_vnic.kvm-guest-mgmt-vnic.mac_address}"
  kvm_guest_vnic_id = "${data.oci_core_vnic.kvm-guest-mgmt-vnic.id}"
  kvm_guest_emulation_mode = "${var.kvm_emulation_mode}"

}
```


## Module Template files

#### main.tf
Remote exec resources responsible for the installation and setup of KVM on the Oracle Linux instance.

#### variables.tf
Module variables as described above.



## Module Data Source template_file

#### generate_virsh_attach.sh.tp
Data source template file used to generate a shell script dynamically based on the instance runtime variables. The generated script retrieves vnic data from the metadata repository.



## Shell Scrips
#### /scripts/configure-secondary-vnics.sh
Shell script to setup secondary vnics at OS level based on the metadata repository.

#### /scripts/install-kvm-dependencies.sh
Shell script responsible to install all the linux packages for KVM.

#### /scripts/secondary_vnics.service
Linux Service definition for setting up all the secondary vnics in terms of instance reboot since these data are not persisted.
