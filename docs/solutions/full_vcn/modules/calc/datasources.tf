# Calculate the list of net numbers for use by the template when having customized
# subnet masks
data "external" "netnum" {
	program = [ "/usr/local/bin/python2.7", "${path.module}/tfnetnum.py"]
	query {
		masks = "${join(",", var.subnet["subnet_masks"])}"
	}
}