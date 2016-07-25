resource "baremetal_server" "foo" {
    address = "hashicorp.com"
}

resource "baremetal_identity_user" "users" {
  compartment_id = "TBD.TBD.TBD"
	description = "A test user"
}
