###############################
#### Environment variables ####
###############################

tenancy_ocid="ocid1.tenancy.oc1..abcabcababccbcbaabcabababcbabababcbababcbabababacbcabcabcabac4uq"
compartment_ocid="ocid1.tenancy.oc1..abcabcababccbcbaabcabababcbabababcbababcbabababacbcabcabcabac4uq"
user_ocid="ocid1.user.oc1..aaaaaaaaubbbbbbbbbbccccccccccdddddd64fizdmbjblpq"
fingerprint="aa:bb:cc:dd:5e:ff:ee:47:9b:3c:5f:c5:69:02:1b:aa"
ssh_private_key_path="~/.ssh/id_rsa"
ssh_public_key_path="~/.ssh/id_rsa.pub"
region="us-phoenix-1"
#region="us-ashburn-1"

#######################################
#### Instance definition variables ####
#######################################

prefix = "sample-nestedkvm"
#availability_domain number For AD1 uses 1. For AD2, uses 2, For AD3, uses 3
availability_domain = "1"
#Only VM Shapes are supported
instance_shape = "VM.Standard1.8"
vcn_cidr_block = "10.0.0.0/16"
kvm_host_subnet_cidr_block = "10.0.10.0/24"


##############################
#### KVM related settings ####
##############################

#URL of your image file (you can place your image in the object storage!)
kvm_image_url = "<my-qcow2-image-url>"
kvm_image_name = "my-qcow2-image.qcow2"
kvm_image_path = "/kvm-imgs/"

kvm_guest_domain_name = "MyKVMDomain"
kvm_guest_memory = "16384"
kvm_guest_vcpu = "8"
kvm_emulation_mode = "virtio"
kvm_guest_os_type = "linux"
kvm_guest_vnc_port = "5901"
kvm_guest_vnc_pwd = "Test123"
