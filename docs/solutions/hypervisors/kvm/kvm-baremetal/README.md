KVM Virtualization
===============================

The Bring Your Own KVM (BYO-KVM) capability of OCI is designed to support customers who need to run operating systems that the platform doesn’t support natively. For a list of natively supported Operating Systems you can refer to the official [documentation](https://docs.us-phoenix-1.oraclecloud.com/Content/Compute/References/images.htm).

In this example we demonstrate the automatic installation of a Guest Virtual Machine on a Bare Metal instance running KVM in Oracle Cloud Infrastructure (OCI). It can be easily adapted to automate the installation of any OS that KVM [supports](https://www.linux-kvm.org/page/Guest_Support_Status).

We don't provide any specific Guest VM on this example. You are responsible to specify the VM file. You can search for a kvm/qcow2 image file on Google or search on some specialized Linux sites such as tuxfixer.com or you can just convert your own image into a qcow2 format.

Note that the Terraform template is expecting to download the image from a particular URL. Because of that, we recommend you to [upload the target image to OCI object storage](https://docs.us-phoenix-1.oraclecloud.com/Content/GSG/Tasks/addingbuckets.htm#two), create a [pre-authenticated request](https://docs.us-phoenix-1.oraclecloud.com/Content/Object/Tasks/managingobjects.htm#par), and use it as the image URL.

This sample code will be responsible to perform the following tasks:

- Setup of all required network components: VCN, Internet Gateway, subnet, and security lists.

- Spin up a [Bare Metal](https://docs.us-phoenix-1.oraclecloud.com/Content/Compute/Concepts/computeoverview.htm) Instance based on Oracle Linux 7.x image.

- Attach [a secondary VNIC](https://docs.us-phoenix-1.oraclecloud.com/Content/Network/Tasks/managingVNICs.htm) to the Bare Metal instance, configured as a vlan-tagged interface.

-	Installation of the KVM hypervisor provided by the Linux Kernel provided as a [Terraform module](https://www.terraform.io/docs/modules/usage.html) (reusable artifact). This process requires instance restart due to the kernel changes.

-	Download of the image file.

- Create a KVM domain for the Guest VM.

- Attach the vlan-tagged interface (in PCI-passthrough mode) as network interfaces to the guest VM.


Requirements
------------

- Access to OCI environment
- [Generated OCI keys and OCIDs](https://docs.us-phoenix-1.oraclecloud.com/Content/API/Concepts/apisigningkey.htm)
- Terraform 0.10.x+


### Usage

- Update `terraform.tfvars` with the required information. Terraform automatically loads them to populate variables, but you can also use the -var-file flag directly to specify a file. These files are the same syntax as Terraform configuration files. And like Terraform configuration files, these files can also be JSON.

We don't recommend saving usernames and passwords to version control, but you can create a local secret variables file and use -var-file to load it.

Most examples use the same set of environment variables so you only need to do this once. There's a sample file available on this sample code for your reference in addition to the snippet below:

```
###############################
#### Environment variables ####
###############################
tenancy_ocid="<tenancy OCID>"
compartment_ocid="<compartment OCID>"
user_ocid="<tenancy OCID>"
fingerprint="<PEM key fingerprint>"
ssh_private_key_path="<path to the ssh private key to ssh to the instance>"
ssh_public_key_path="<path to the ssh public key to setup on the instance>"
region="<OCI region>"

#######################################
#### Instance definition variables ####
#######################################

#Customer Identifier to be used before the resources name
customer_name = "mycustomer"

#availability_domain number - For AD1 uses 1. For AD2, uses 2, For AD3, uses 3
availability_domain = "1"

#Only BM Shapes are supported
instance_shape = "BM.Standard1.36"


##############################
#### KVM related settings ####
##############################

#URL of your image file (you can place your image in the object storage!)
kvm_image_url = "<my-qcow2-image-url>"
kvm_image_name = "my-qcow2-image.qcow2"
kvm_image_path = "/kvm-imgs/"

kvm_guest_domain_name = "MyDomain"
kvm_guest_memory = "16384"
kvm_guest_vcpu = "8"
kvm_emulation_mode = "virtio"
kvm_guest_os_type = "linux"
kvm_guest_vnc_port = "5901"
kvm_guest_vnc_pwd = "Test123"
```

- After saving your environment variables, you should run terraform following the same order as specified below:

- Run `terraform get` to bring the module files into the current workspace

- Run `terraform init` to download all the providers/plugins files

- Run `terraform plan`

- Run `terraform apply`

- Create a Tunnel over SSH to establish a VNC connection to access the guest VM. As a result of the `terraform apply`, you will get the command line used for creating the SSH tunnel.

```
ssh -i <ssh-private-key> -L <vnc_port>:localhost:<vnc_port> opc@<kvm_host_public_ip>
```

- Setup your Guest Instance. Access it through a [VNC](https://en.wikipedia.org/wiki/Virtual_Network_Computing) session

`localhost:kvm_guest_vnc_port`


- Finally, for access any application/tool within your Guest VM, you can retrieve the KVM Guest public IP running the terraform output command:

```
$ terraform output KVM_GUEST_PUBLIC_IP
```
