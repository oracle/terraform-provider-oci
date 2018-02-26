# Gets a list of Availability Domains
data "oci_identity_availability_domains" "ADs" {
  compartment_id = "${var.tenancy_ocid}"
}

# Get DB node list
data "oci_database_db_nodes" "DBNodeList" {
  compartment_id = "${var.compartment_ocid}"
  db_system_id = "${oci_database_db_system.TFDBNode.id}"
}

# Get DB node details
data "oci_database_db_node" "DBNodeDetails" {
  db_node_id = "${lookup(data.oci_database_db_nodes.DBNodeList.db_nodes[0], "id")}"
}

# Gets the OCID of the first (default) vNIC
data "oci_core_vnic" "DBNodeVnic" {
  vnic_id = "${data.oci_database_db_node.DBNodeDetails.vnic_id}"
}

