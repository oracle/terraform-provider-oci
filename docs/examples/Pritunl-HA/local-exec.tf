resource "null_resource" "mongokey" {
	depends_on = ["oci_core_instance.MongoP","oci_core_instance.MongoR1","oci_core_instance.MongoR2"]	
	provisioner "local-exec" {
		command="openssl rand -base64 756 > mongokey"
	}
}

