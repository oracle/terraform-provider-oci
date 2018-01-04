# DO NOT ALTER THIS FILE

data "oci_identity_compartments" "compartment" {
	compartment_id = "${var.tenancy_ocid}"
	filter {
		name = "name"
		values = [ "${var.build_env["compartment"]}" ]
	}
}

data "oci_identity_availability_domains" "ad" {
	compartment_id = "${var.tenancy_ocid}"
	filter {
		name = "name"
		values = [ "\\w*-${upper(var.build_env["ad"])}" ]
		regex = true
	}
}

data "oci_core_virtual_networks" "vcn" {
	compartment_id = "${data.oci_identity_compartments.compartment.compartments.0.id}"
	filter {
		name = "display_name"
		values = [ "${var.build_env["vcn"]}" ]
	}
}

data "oci_core_subnets" "subnet" {
	compartment_id = "${data.oci_identity_compartments.compartment.compartments.0.id}"
	vcn_id = "${data.oci_core_virtual_networks.vcn.virtual_networks.0.id}"
	filter {
		name = "display_name"
		values = [ "${var.build_env["subnet"]}" ]
	}

}

# Gets the OCID of the image. This technique is for example purposes only. 
# The results of oci_core_images may change over time for Oracle-provided images, 
# so the only sure way to get the correct OCID is to supply it directly.

data "oci_core_images" "image" {
	compartment_id = "${data.oci_identity_compartments.compartment.compartments.0.id}"
	operating_system = "${var.ipxe_instance["os"]}"
	operating_system_version = "${var.ipxe_instance["os-version"]}"
	filter {
		name = "display_name"
		values = [ ".*-${var.ipxe_instance["os-version"]}-2.*" ]
		regex = true
	}
}

data "external" "ipxe_gen" {
	program = [ "/bin/bash", "./ipxe_gen.sh"]
	query = {
		tenancy_ocid         = "${var.tenancy_ocid}"
  		user_ocid            = "${var.user_ocid}"
 		private_key_path     = "${var.private_key_path}"
		private_key_password = "${var.private_key_password}"
		region               = "${var.region}"
		ssh_public_key		 = "${var.ssh_public_key}"
		os_short_name		 = "rhel74"
		rhel_user			 = "${var.rhel_account["user_name"]}"
		rhel_pw			 = "${var.rhel_account["password"]}"
		zeros_ocid		 = "${var.region_all_zeros_ocid[var.region]}"
		iso_url			 = "${var.iso_url}"
	}
}