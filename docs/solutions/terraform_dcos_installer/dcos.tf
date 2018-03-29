data "oci_identity_availability_domains" "ADs" {
    compartment_id = "${var.tenancy_ocid}"
}

### VCN
module "vcn" {
    source              = "./network/vcn"
    compartment_ocid    = "${var.compartment_ocid}"
    tenancy_ocid        = "${var.compartment_ocid}"
    network_cidrs	= "${var.network_cidrs}"
}

### Instances
# Bootstrap Instance
module "dcos_bootstrap" {
    source              = "./instances/bootstrap"
    BootstrapAD         = "${var.BootstrapAD}"
    compartment_ocid    = "${var.compartment_ocid}"
    tenancy_ocid        = "${var.compartment_ocid}"
    dcos_cluster_name   = "${var.dcos_cluster_name}"
    image               = "${var.InstanceImageOCID[var.region]}"
    shape               = "${var.BootstrapInstanceShape}"
    dcos_installer_url  = "${var.dcos_installer_url}"
    ssh_public_key      = "${var.ssh_public_key}"
    ssh_private_key     = "${var.ssh_private_key}"
    subnets	        = "${module.vcn.subnets}"
}

module "dcos_master_ad1" {
    source              = "./instances/master"
    count               = "${var.dcos_master_ad1_count}"
    availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[0],"name")}"
    compartment_ocid    = "${var.compartment_ocid}"
    tenancy_ocid        = "${var.compartment_ocid}"
    dcos_cluster_name   = "${var.dcos_cluster_name}"
    image               = "${var.InstanceImageOCID[var.region]}"
    shape               = "${var.MasterInstanceShape}"
    subnet_id           = "${module.vcn.subnet_ad1_id}" 
    ssh_public_key      = "${var.ssh_public_key}"
    ssh_private_key     = "${var.ssh_private_key}"
    display_name_prefix = "ad1"
}

module "dcos_master_ad2" {
    source              = "./instances/master"
    count               = "${var.dcos_master_ad2_count}"
    availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[1],"name")}"
    compartment_ocid    = "${var.compartment_ocid}"
    tenancy_ocid        = "${var.compartment_ocid}"
    dcos_cluster_name   = "${var.dcos_cluster_name}"
    image               = "${var.InstanceImageOCID[var.region]}"
    shape               = "${var.MasterInstanceShape}"
    subnet_id           = "${module.vcn.subnet_ad2_id}" 
    ssh_public_key      = "${var.ssh_public_key}"
    ssh_private_key     = "${var.ssh_private_key}"
    display_name_prefix = "ad2"
}

module "dcos_master_ad3" {
    source              = "./instances/master"
    count               = "${var.dcos_master_ad3_count}"
    availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[2],"name")}"
    compartment_ocid    = "${var.compartment_ocid}"
    tenancy_ocid        = "${var.compartment_ocid}"
    dcos_cluster_name   = "${var.dcos_cluster_name}"
    image               = "${var.InstanceImageOCID[var.region]}"
    shape               = "${var.MasterInstanceShape}"
    subnet_id           = "${module.vcn.subnet_ad3_id}" 
    ssh_public_key      = "${var.ssh_public_key}"
    ssh_private_key     = "${var.ssh_private_key}"
    display_name_prefix = "ad3"
}


module  "dcos_agent_ad1" {
    source              = "./instances/agent"
    count               = "${var.dcos_agent_ad1_count}"
    availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[0],"name")}"
    compartment_ocid    = "${var.compartment_ocid}"
    tenancy_ocid        = "${var.compartment_ocid}"
    dcos_cluster_name   = "${var.dcos_cluster_name}"
    image               = "${var.InstanceImageOCID[var.region]}"
    shape               = "${var.AgentInstanceShape}"
    subnet_id           = "${module.vcn.subnet_ad1_id}" 
    ssh_public_key      = "${var.ssh_public_key}"
    ssh_private_key     = "${var.ssh_private_key}"
    display_name_prefix = "ad1"
    dcos_bootstrap_instance_id = "${module.dcos_bootstrap.instance_id}"
}

module  "dcos_agent_ad2" {
    source              = "./instances/agent"
    count               = "${var.dcos_agent_ad2_count}"
    availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[1],"name")}"
    compartment_ocid    = "${var.compartment_ocid}"
    tenancy_ocid        = "${var.compartment_ocid}"
    dcos_cluster_name   = "${var.dcos_cluster_name}"
    image               = "${var.InstanceImageOCID[var.region]}"
    shape               = "${var.AgentInstanceShape}"
    subnet_id           = "${module.vcn.subnet_ad2_id}" 
    ssh_public_key      = "${var.ssh_public_key}"
    ssh_private_key     = "${var.ssh_private_key}"
    display_name_prefix = "ad2"
    dcos_bootstrap_instance_id = "${module.dcos_bootstrap.instance_id}"
}

module  "dcos_agent_ad3" {
    source              = "./instances/agent"
    count               = "${var.dcos_agent_ad3_count}"
    availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[2],"name")}"
    compartment_ocid    = "${var.compartment_ocid}"
    tenancy_ocid        = "${var.compartment_ocid}"
    dcos_cluster_name   = "${var.dcos_cluster_name}"
    image               = "${var.InstanceImageOCID[var.region]}"
    shape               = "${var.AgentInstanceShape}"
    subnet_id           = "${module.vcn.subnet_ad3_id}" 
    ssh_public_key      = "${var.ssh_public_key}"
    ssh_private_key     = "${var.ssh_private_key}"
    display_name_prefix = "ad3"
    dcos_bootstrap_instance_id = "${module.dcos_bootstrap.instance_id}"
}

module  "dcos_public_agent_ad1" {
    source              = "./instances/public-agent"
    count               = "${var.dcos_public_agent_ad1_count}"
    availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[0],"name")}"
    compartment_ocid    = "${var.compartment_ocid}"
    tenancy_ocid        = "${var.compartment_ocid}"
    dcos_cluster_name   = "${var.dcos_cluster_name}"
    image               = "${var.InstanceImageOCID[var.region]}"
    shape               = "${var.PublicAgentInstanceShape}"
    subnet_id           = "${module.vcn.subnet_ad1_id}" 
    ssh_public_key      = "${var.ssh_public_key}"
    ssh_private_key     = "${var.ssh_private_key}"
    display_name_prefix = "ad1"
    dcos_bootstrap_instance_id = "${module.dcos_bootstrap.instance_id}"
}

module  "dcos_public_agent_ad2" {
    source              = "./instances/public-agent"
    count               = "${var.dcos_public_agent_ad2_count}"
    availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[1],"name")}"
    compartment_ocid    = "${var.compartment_ocid}"
    tenancy_ocid        = "${var.compartment_ocid}"
    dcos_cluster_name   = "${var.dcos_cluster_name}"
    image               = "${var.InstanceImageOCID[var.region]}"
    shape               = "${var.PublicAgentInstanceShape}"
    subnet_id           = "${module.vcn.subnet_ad2_id}" 
    ssh_public_key      = "${var.ssh_public_key}"
    ssh_private_key     = "${var.ssh_private_key}"
    display_name_prefix = "ad2"
    dcos_bootstrap_instance_id = "${module.dcos_bootstrap.instance_id}"
}


module  "dcos_public_agent_ad3" {
    source              = "./instances/public-agent"
    count               = "${var.dcos_public_agent_ad3_count}"
    availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[2],"name")}"
    compartment_ocid    = "${var.compartment_ocid}"
    tenancy_ocid        = "${var.compartment_ocid}"
    dcos_cluster_name   = "${var.dcos_cluster_name}"
    image               = "${var.InstanceImageOCID[var.region]}"
    shape               = "${var.PublicAgentInstanceShape}"
    subnet_id           = "${module.vcn.subnet_ad3_id}" 
    ssh_public_key      = "${var.ssh_public_key}"
    ssh_private_key     = "${var.ssh_private_key}"
    display_name_prefix = "ad3"
    dcos_bootstrap_instance_id = "${module.dcos_bootstrap.instance_id}"
}

module  "dcos_gpu_agent_ad1" {
    source              = "./instances/gpu-agent"
    count               = "${var.dcos_gpu_agent_ad1_count}"
    availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[0],"name")}"
    compartment_ocid    = "${var.compartment_ocid}"
    tenancy_ocid        = "${var.compartment_ocid}"
    dcos_cluster_name   = "${var.dcos_cluster_name}"
    image               = "${var.GPUInstanceImageOCID[var.region]}"
    shape               = "${var.GPUInstanceShape}"
    subnet_id           = "${module.vcn.subnet_ad1_id}" 
    ssh_public_key      = "${var.ssh_public_key}"
    ssh_private_key     = "${var.ssh_private_key}"
    display_name_prefix = "ad1"
    dcos_bootstrap_instance_id = "${module.dcos_bootstrap.instance_id}"
}

module  "dcos_gpu_agent_ad2" {
    source              = "./instances/gpu-agent"
    count               = "${var.dcos_gpu_agent_ad2_count}"
    availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[1],"name")}"
    compartment_ocid      = "${var.compartment_ocid}"
    tenancy_ocid        = "${var.compartment_ocid}"
    dcos_cluster_name   = "${var.dcos_cluster_name}"
    image               = "${var.GPUInstanceImageOCID[var.region]}"
    shape               = "${var.GPUInstanceShape}"
    subnet_id           = "${module.vcn.subnet_ad2_id}" 
    ssh_public_key      = "${var.ssh_public_key}"
    ssh_private_key     = "${var.ssh_private_key}"
    display_name_prefix = "ad2"
    dcos_bootstrap_instance_id = "${module.dcos_bootstrap.instance_id}"
}

module  "dcos_gpu_agent_ad3" {
    source              = "./instances/gpu-agent"
    count               = "${var.dcos_gpu_agent_ad3_count}"
    availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[2],"name")}"
    compartment_ocid    = "${var.compartment_ocid}"
    tenancy_ocid        = "${var.compartment_ocid}"
    dcos_cluster_name   = "${var.dcos_cluster_name}"
    image               = "${var.GPUInstanceImageOCID[var.region]}"
    shape               = "${var.GPUInstanceShape}"
    subnet_id           = "${module.vcn.subnet_ad3_id}" 
    ssh_public_key      = "${var.ssh_public_key}"
    ssh_private_key     = "${var.ssh_private_key}"
    display_name_prefix = "ad3"
    dcos_bootstrap_instance_id = "${module.dcos_bootstrap.instance_id}"
}

