output "cidr_list" {
	value = "${join(",",data.template_file.subnet.*.rendered)}"
}