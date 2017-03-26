resource "baremetal_core_instance" "MongoDBBastion" {
    availability_domain = "${lookup(data.baremetal_identity_availability_domains.ADs.availability_domains[0],"name")}"
    compartment_id = "${var.compartment_ocid}"
    display_name = "MongoDB-Bastion"
    image = "${lookup(data.baremetal_core_images.OLImageOCID.images[0], "id")}" 
    shape = "${var.BastionShape}"
    subnet_id = "${baremetal_core_subnet.BastionSubnetAD1.id}"
    metadata {
        ssh_authorized_keys = "${var.ssh_public_key}"
	user_data = "${base64encode(file(var.BastionBootStrap))}" 
    }
}

resource "baremetal_core_instance" "MongoDBAD1" {
    availability_domain = "${lookup(data.baremetal_identity_availability_domains.ADs.availability_domains[0],"name")}"
    compartment_id = "${var.compartment_ocid}"
    display_name = "MongoDBAD1"
    image = "${lookup(data.baremetal_core_images.OLImageOCID.images[0], "id")}"
    shape = "${var.MongoDBShape}"
    subnet_id = "${baremetal_core_subnet.PrivSubnetAD1.id}"
    metadata {
        ssh_authorized_keys = "${var.ssh_public_key}"
        user_data = "${base64encode(file(var.MongoDBBootStrap))}"
    }
}

resource "baremetal_core_instance" "MongoDBAD2" {
    availability_domain = "${lookup(data.baremetal_identity_availability_domains.ADs.availability_domains[1],"name")}"
    compartment_id = "${var.compartment_ocid}"
    display_name = "MongoDBAD2"
    image = "${lookup(data.baremetal_core_images.OLImageOCID.images[0], "id")}"
    shape = "${var.MongoDBShape}"
    subnet_id = "${baremetal_core_subnet.PrivSubnetAD2.id}"
    metadata {
        ssh_authorized_keys = "${var.ssh_public_key}"
        user_data = "${base64encode(file(var.MongoDBBootStrap))}"
    }
}
