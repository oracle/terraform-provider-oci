# Gets a list of Availability Domains
data "baremetal_identity_availability_domains" "ADs" {
  compartment_id = "${var.tenancy_ocid}"
}

# Get DB node details
data "baremetal_database_db_node" "DBNodeDetails" { 
db_node_id = "${baremetal_database_db_system.TFDBNode.id}" 
}

# Gets the OCID of the first (default) vNIC
data "baremetal_core_vnic" "DBNodeVnic" {
vnic_id = "${lookup(data.baremetal_database_db_node.DBNodeDetails.vnic_id,"vnic_id")}"
}

