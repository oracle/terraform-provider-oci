package tfresource

import "testing"

func TestUnitADs(t *testing.T) {
	wantedString := `
	data "oci_identity_availability_domains" "t" {
		compartment_id = "${var.compartment_id}"
	}`
	tests := []struct {
		name string
		want string
	}{
		{
			name: "Test returned value is as expected for ADs",
			want: wantedString,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TestADs(); got != tt.want {
				t.Errorf("TestADs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnitImage1(t *testing.T) {
	wantedString := `
	variable "InstanceImageOCID" {
	  type = "map"
	  default = {
		// See https://docs.us-phoenix-1.oraclecloud.com/images/
		// Oracle-provided image "Oracle-Linux-7.5-2018.10.16-0"
		us-phoenix-1 = "ocid1.image.oc1.phx.aaaaaaaadjnj3da72bztpxinmqpih62c2woscbp6l3wjn36by2cvmdhjub6a"
		us-ashburn-1 = "ocid1.image.oc1.iad.aaaaaaaawufnve5jxze4xf7orejupw5iq3pms6cuadzjc7klojix6vmk42va"
		eu-frankfurt-1 = "ocid1.image.oc1.eu-frankfurt-1.aaaaaaaagbrvhganmn7awcr7plaaf5vhabmzhx763z5afiitswjwmzh7upna"
		uk-london-1 = "ocid1.image.oc1.uk-london-1.aaaaaaaajwtut4l7fo3cvyraate6erdkyf2wdk5vpk6fp6ycng3dv2y3ymvq"
	  }
	}`
	tests := []struct {
		name string
		want string
	}{
		{
			name: "Test returned value is as expected for Images",
			want: wantedString,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TestImage1(); got != tt.want {
				t.Errorf("TestImage1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnitInstance1(t *testing.T) {
	wantedString := `
	resource "oci_core_instance" "t" {
		availability_domain = "${data.oci_identity_availability_domains.t.availability_domains.0.name}"
		compartment_id = "${var.compartment_id}"
		subnet_id = "${oci_core_subnet.t.id}"
		image = "${var.InstanceImageOCID[var.region]}"
		shape = "VM.Standard2.1"
		metadata = {}
		timeouts {
			create = "15m"
		}
	}`
	tests := []struct {
		name string
		want string
	}{
		{
			name: "Test returned value is as expected for Images",
			want: wantedString,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TestInstance1(); got != tt.want {
				t.Errorf("TestInstance1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnitSubnet1(t *testing.T) {
	wantedString := `
	resource "oci_core_subnet" "t" {
		compartment_id      = "${var.compartment_id}"
		vcn_id              = "${oci_core_virtual_network.t.id}"
		availability_domain = "${lookup(data.oci_identity_availability_domains.t.availability_domains[0],"name")}"
		route_table_id      = "${oci_core_virtual_network.t.default_route_table_id}"
		security_list_ids = ["${oci_core_virtual_network.t.default_security_list_id}"]
		dhcp_options_id     = "${oci_core_virtual_network.t.default_dhcp_options_id}"
		cidr_block          = "10.0.1.0/24"
		display_name        = "-tf-subnet"
		dns_label           = "subnetdns"
	}`
	tests := []struct {
		name string
		want string
	}{
		{
			name: "Test returned value is as expected for Subnet",
			want: wantedString,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TestSubnet1(); got != tt.want {
				t.Errorf("TestSubnet1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnitVCN1(t *testing.T) {
	wantedString := `
	resource "oci_core_virtual_network" "t" {
		compartment_id = "${var.compartment_id}"
		cidr_block = "10.0.0.0/16"
		display_name = "-tf-vcn"
		dns_label    = "vcndns"
	}`
	tests := []struct {
		name string
		want string
	}{
		{
			name: "Test returned value is as expected for VCN",
			want: wantedString,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TestVCN1(); got != tt.want {
				t.Errorf("TestVCN1() = %v, want %v", got, tt.want)
			}
		})
	}
}
