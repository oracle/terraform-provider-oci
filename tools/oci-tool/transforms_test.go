// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package oci_tool

import (
	"strings"
	"testing"
)

func TestBraceIndexing1(t *testing.T) {
	const str = `{{}}`
	expectStart := 0
	expectEnd := 3
	start, end := indexOpenCloseTokens('{', '}', str)

	if start != expectStart {
		t.Errorf("expected %d, got %d\n", expectStart, start)
	}

	if end != expectEnd {
		t.Errorf("expected %d, got %d\n", expectEnd, end)
	}
}

func TestBraceIndexing2(t *testing.T) {
	const str = ` { { } } `
	expectStart := 1
	expectEnd := 7
	start, end := indexOpenCloseTokens('{', '}', str)

	if start != expectStart {
		t.Errorf("expected %d, got %d\n", expectStart, start)
	}

	if end != expectEnd {
		t.Errorf("expected %d, got %d\n", expectEnd, end)
	}
}

func TestBraceIndexingWithStrings(t *testing.T) {
	const str = ` " " { { } } `
	expectStart := 5
	expectEnd := 11
	start, end := indexOpenCloseTokens('{', '}', str)

	if start != expectStart {
		t.Errorf("expected %d, got %d\n", expectStart, start)
	}

	if end != expectEnd {
		t.Errorf("expected %d, got %d\n", expectEnd, end)
	}
}

func TestBraceIndexingReal1(t *testing.T) {
	const str = `provider "oci" {
  tenancy_ocid = "${var.tenancy_ocid}"
  //user_ocid = "${var.user_ocid}"
  #fingerprint = "${var.fingerprint}"
  /*private_key_path = "${var.private_key_path}"*/
  region = "${var.region}"
}`

	expectStart := 15
	expectEnd := len(str) - 1
	start, end := indexOpenCloseTokens('{', '}', str)

	if start != expectStart {
		t.Errorf("expected %d, got %d\n", expectStart, start)
	}

	if end != expectEnd {
		t.Errorf("expected %d, got %d\n", expectEnd, end)
	}
}

func TestBraceIndexingRealMessy(t *testing.T) {
	const str = `provider "oci"{region = "${var.region}"/*private_key_path="${var.private_key_path}"*/tenancy_ocid="${var.tenancy_ocid}"fingerprint="${var.fingerprint}"//user_ocid="${var.user_ocid}"
}`

	expectStart := 14
	expectEnd := len(str) - 1
	start, end := indexOpenCloseTokens('{', '}', str)

	if start != expectStart {
		t.Errorf("expected %d, got %d\n", expectStart, start)
	}

	if end != expectEnd {
		t.Errorf("expected %d, got %d\n", expectEnd, end)
	}
}

func TestFindOpeningBrace(t *testing.T) {
	const str = `}{}`
	expect := 1
	start, _ := indexOpenCloseTokens('{', '}', str)

	if start != expect {
		t.Errorf("expected %d, got %d\n", expect, start)
	}
}

func TestFindClosingBrace(t *testing.T) {
	const str = `{}}`
	expect := 1
	_, end := indexOpenCloseTokens('{', '}', str)
	if end != expect {
		t.Errorf("expected %d, got %d\n", expect, end)
	}
}

func TestMissingOpeningBrace(t *testing.T) {
	const str = `}}`
	expect := -1
	start, _ := indexOpenCloseTokens('{', '}', str)

	if start != expect {
		t.Errorf("expected %d, got %d\n", expect, start)
	}
}

func TestMissingClosingBrace(t *testing.T) {
	const str = `{{}`
	expect := -1
	_, end := indexOpenCloseTokens('{', '}', str)

	if end != expect {
		t.Errorf("expected %d, got %d\n", expect, end)
	}
}

func TestInlineCommentsIgnored(t *testing.T) {
	const str = `{//}
}`
	expect := len(str) - 1
	_, end := indexOpenCloseTokens('{', '}', str)

	if end != expect {
		t.Errorf("expected %d, got %d\n", expect, end)
	}
}

func TestEscapedQuotesIgnored(t *testing.T) {
	const str = `{"\""}`
	expect := len(str) - 1
	_, end := indexOpenCloseTokens('{', '}', str)

	if end != expect {
		t.Errorf("expected %d, got %d\n", expect, end)
	}
}

func TestQuotedCommentsIgnored(t *testing.T) {
	const str = `{"//"}`
	expect := len(str) - 1
	_, end := indexOpenCloseTokens('{', '}', str)

	if end != expect {
		t.Errorf("expected %d, got %d\n", expect, end)
	}
}

func TestHashCommentsIgnored(t *testing.T) {
	const str = `{#}
   }`
	expect := len(str) - 1
	_, end := indexOpenCloseTokens('{', '}', str)

	if end != expect {
		t.Errorf("expected %d, got %d\n", expect, end)
	}
}

func TestQuotedHashCommentsIgnored(t *testing.T) {
	const str = `{"#"}`
	expect := len(str) - 1
	_, end := indexOpenCloseTokens('{', '}', str)

	if end != expect {
		t.Errorf("expected %d, got %d\n", expect, end)
	}
}

func TestBlockCommentsIgnored(t *testing.T) {
	const str = `{ /* } */ }`
	expect := len(str) - 1
	_, end := indexOpenCloseTokens('{', '}', str)

	if end != expect {
		t.Errorf("expected %d, got %d\n", expect, end)
	}
}

func TestMultilineBlockCommentsIgnored(t *testing.T) {
	const str = `{/*
}
*/}`
	expect := len(str) - 1
	_, end := indexOpenCloseTokens('{', '}', str)

	if end != expect {
		t.Errorf("expected %d, got %d\n", expect, end)
	}
}

func TestTokensIgnored(t *testing.T) {
	const str = `{${1}}`
	expect := len(str) - 1
	_, end := indexOpenCloseTokens('{', '}', str)

	if end != expect {
		t.Errorf("expected %d, got %d\n", expect, end)
	}
}

func TestMultilineTokensIgnored(t *testing.T) {
	const str = `{ ${1
}}`
	expect := len(str) - 1
	_, end := indexOpenCloseTokens('{', '}', str)

	if end != expect {
		t.Errorf("expected %d, got %d\n", expect, end)
	}
}

func TestProviderBlockHasRegion(t *testing.T) {
	const str = `
variable "compartment_ocid" {}
variable "region" {}

provider "oci" {
  tenancy_ocid = "${var.tenancy_ocid}"
  //user_ocid = "${var.user_ocid}"
  fingerprint = "${var.fingerprint}"
  /*private_key_path = "${var.private_key_path}"*/
  region = "${var.region}"
}`
	expect := true
	_, _, _, hasRegion := providerBlockHasRegion(str)

	if hasRegion != expect {
		t.Errorf("expected %t, got %t\n", expect, hasRegion)
	}
}

func TestProviderBlockHasRegionNot(t *testing.T) {
	const str = `
variable "compartment_ocid" {}
variable "region" {}

provider "oci" {
  tenancy_ocid = "${var.tenancy_ocid}"
  //user_ocid = "${var.user_ocid}"
  fingerprint = "${var.fingerprint}"
  /*private_key_path = "${var.private_key_path}"*/
}`
	expect := false
	_, _, _, hasRegion := providerBlockHasRegion(str)

	if hasRegion != expect {
		t.Errorf("expected %t, got %t\n", expect, hasRegion)
	}
}

func TestProviderBlockHasRegionWrongProvider(t *testing.T) {
	const str = `
provider "opc" {
  id = "${var.id}"
}`
	expect := false
	_, _, isOci, _ := providerBlockHasRegion(str)

	if isOci != expect {
		t.Errorf("expected %t, got %t\n", expect, isOci)
	}
}

func TestInsertRegionInProviderBlock(t *testing.T) {
	const str = `
variable "compartment_ocid" {}
variable "region" {}

provider "oci" {
  tenancy_ocid = "${var.tenancy_ocid}"
  //user_ocid = "${var.user_ocid}"
  fingerprint = "${var.fingerprint}"
  /*private_key_path = "${var.private_key_path}"*/
}`
	expect := true
	res, _ := insertRegionInProviderBlock(str)

	regionFound := matchRegion.MatchString(res)
	if regionFound != expect {
		t.Errorf("expected %t, got %t\n", expect, regionFound)
	}
}

func TestInsertRegionInProviderBlockMessy(t *testing.T) {
	const str = `
variable "compartment_ocid" {}
variable "region" {}

provider "oci" {tenancy_ocid = "${var.tenancy_ocid}"
  //user_ocid = "${var.user_ocid}"
  fingerprint = "${var.fingerprint}"
  /*private_key_path = "${var.private_key_path}"*/
}`
	expect := true
	res, _ := insertRegionInProviderBlock(str)

	regionFound := matchRegion.MatchString(res)
	if regionFound != expect {
		t.Errorf("expected %t, got %t\n", expect, regionFound)
	}
}

func TestInsertRegionInProviderBlockNot(t *testing.T) {
	const str = `
variable "compartment_ocid" {}
variable "region" {}

provider "oci" {
  tenancy_ocid = "${var.tenancy_ocid}"
  //user_ocid = "${var.user_ocid}"
  fingerprint = "${var.fingerprint}"
  /*private_key_path = "${var.private_key_path}"*/
  region = "us-phoenix-1"
}`
	expect := len(str)
	res, _ := insertRegionInProviderBlock(str)

	if len(res) != expect {
		t.Errorf("expected string length %d, got %d\n", expect, len(res))
	}
}

func TestInsertRegionInProviderBlockMalformedNoProvider(t *testing.T) {
	const str = `
variable "compartment_ocid" {}
variable "region" {}

"oci" {
  tenancy_ocid = "${var.tenancy_ocid}"
  //user_ocid = "${var.user_ocid}"
  fingerprint = "${var.fingerprint}"
  /*private_key_path = "${var.private_key_path}"*/`

	res, err := insertRegionInProviderBlock(str)

	if res != str {
		t.Errorf("expected original string \n%s\n got \n%s\n", str, res)
	}

	if err == nil {
		t.Errorf("expected error, got nil")
	}
}

func TestInsertRegionInProviderBlockMalformedNoClosingBrace(t *testing.T) {
	const str = `
variable "compartment_ocid" {}
variable "region" {}

provider "oci" {
  tenancy_ocid = "${var.tenancy_ocid}"
  //user_ocid = "${var.user_ocid}"
  fingerprint = "${var.fingerprint}"
  /*private_key_path = "${var.private_key_path}"*/`

	res, err := insertRegionInProviderBlock(str)

	if res != str {
		t.Errorf("expected original string \n%s\n got \n%s\n", str, res)
	}

	if err == nil {
		t.Errorf("expected error, got nil")
	}
}

func TestScanAndUpdateProvider(t *testing.T) {
	const str = `variable "region" {}

provider "oci" {
  tenancy_ocid = "${var.tenancy_ocid}"
  user_ocid = "${var.user_ocid}"
  fingerprint = "${var.fingerprint}"
  private_key_path = "${var.private_key_path}"
}

data "oci_identity_availability_domains" "ADs" {
  compartment_id = "${var.tenancy_ocid}"
}`
	res, err := scanAndUpdateProvider(str)

	if strings.Count(res, "us-phoenix-1") != 1 {
		t.Errorf("expected to find string 'us-phoenix-1' got\n%s\n", res)
	}

	if err != nil {
		t.Errorf("unexpected error\n %s", err)
	}
}

func TestScanAndUpdateProviderMultiple(t *testing.T) {
	const str = `variable "region" {}

provider "oci" {
  tenancy_ocid = "${var.tenancy_ocid}"
  user_ocid = "${var.user_ocid}"
  fingerprint = "${var.fingerprint}"
  private_key_path = "${var.private_key_path}"
}

provider "oci" {
  tenancy_ocid = "${var.tenancy_ocid}"
  user_ocid = "${var.user_ocid}"
  fingerprint = "${var.fingerprint}"
  private_key_path = "${var.private_key_path}"
}

data "oci_identity_availability_domains" "ADs" {
  compartment_id = "${var.tenancy_ocid}"
}`
	res, err := scanAndUpdateProvider(str)

	expect := 2
	actual := strings.Count(res, "us-phoenix-1")

	if expect != actual {
		t.Errorf("expected 'us-phoenix-1' count to be %d, got %d ", expect, actual)
	}

	if err != nil {
		t.Errorf("unexpected error\n %s", err)
	}
}

func TestScanAndUpdateProviderMultipleAliases(t *testing.T) {
	const str = `variable "region" {}

provider "oci" {
  alias = "phx"
  tenancy_ocid = "${var.tenancy_ocid}"
  user_ocid = "${var.user_ocid}"
  fingerprint = "${var.fingerprint}"
  private_key_path = "${var.private_key_path}"
}

provider "oci" {
  tenancy_ocid = "${var.tenancy_ocid}"
  user_ocid = "${var.user_ocid}"
  fingerprint = "${var.fingerprint}"
  private_key_path = "${var.private_key_path}"
  region = "us-ashburn-1"
  alias = "iad"
}

data "oci_identity_availability_domains" "ADs" {
  compartment_id = "${var.tenancy_ocid}"
}`
	res, err := scanAndUpdateProvider(str)

	expect := 1
	actual := strings.Count(res, `region = "us-ashburn-1"`)

	if expect != actual {
		t.Errorf("expected %d 'region =' fields, got %d ", expect, actual)
	}

	if err != nil {
		t.Errorf("unexpected error\n %s", err)
	}
}

func TestReplaceTemplateTokensBasic(t *testing.T) {
	const original = `
data "baremetal_core_images" "OLImageOCID" {
    compartment_id = "${  data.baremetal_core_compartment.Comp.id  }" # extra space
}
resource "baremetal_core_instance" "ExampleInstance" {
  availability_domain = "${lookup(data.baremetal_identity_availability_domains.ADs.availability_domains[var.AD - 1],"name")}"
  compartment_id = "${data.baremetal_core_compartment.Comp.id}"
  display_name = "baremetal_ExampleInstance" # use baremeta_ in a name
  image = "${lookup(data.baremetal_core_images.OLImageOCID.images[0], "id")}"
  shape = "${var.InstanceShape}"
  subnet_id = "${baremetal_core_subnet.ExampleSubnet.id}"
  create_vnic_details {
    subnet_id = "${baremetal_core_subnet.ExampleSubnet.id}"
  }
  metadata {
    ssh_authorized_keys = "${var.ssh_public_key}"
  }
}`

	const expected = `
data "oci_core_images" "OLImageOCID" {
    compartment_id = "${  data.oci_core_compartment.Comp.id  }" # extra space
}
resource "oci_core_instance" "ExampleInstance" {
  availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[var.AD - 1],"name")}"
  compartment_id = "${data.oci_core_compartment.Comp.id}"
  display_name = "baremetal_ExampleInstance" # use baremeta_ in a name
  image = "${lookup(data.oci_core_images.OLImageOCID.images[0], "id")}"
  shape = "${var.InstanceShape}"
  subnet_id = "${oci_core_subnet.ExampleSubnet.id}"
  create_vnic_details {
    subnet_id = "${oci_core_subnet.ExampleSubnet.id}"
  }
  metadata {
    ssh_authorized_keys = "${var.ssh_public_key}"
  }
}`

	actual := replaceTemplateTokens(original)

	if expected != actual {
		t.Errorf("expected %d, got %d ", expected, actual)
	}
}

func TestReplaceTemplateTokensMultiplePerLine(t *testing.T) {
	const original = `
resource "baremetal_identity_policy" "p" {
  name = "-tf-policy"
  id1 = "${baremetal_identity_compartment1.t.id}", id2 = "${baremetal_core_subnet.ExampleSubnet2.id}", id3 = "${baremetal_core_subnet.ExampleSubnet3.id}"
  id4 = "${data.baremetal_identity_compartment1.t.id}", id5 = "${data.baremetal_core_subnet.ExampleSubnet2.id}", id6 = "${lookup(data.baremetal_identity_group.groups[0], "id")}"
  statements = ["Allow group ${baremetal_identity_group.t.name} to read instances in compartment ${baremetal_identity_compartment.t.name}", "Allow group "${lookup(data.baremetal_identity_group.groups[0], "id")}" to read instances in compartment "${lookup(data.baremetal_identity_compartment.compartments[0], "id")}"]
}`

	const expected = `
resource "oci_identity_policy" "p" {
  name = "-tf-policy"
  id1 = "${oci_identity_compartment1.t.id}", id2 = "${oci_core_subnet.ExampleSubnet2.id}", id3 = "${oci_core_subnet.ExampleSubnet3.id}"
  id4 = "${data.oci_identity_compartment1.t.id}", id5 = "${data.oci_core_subnet.ExampleSubnet2.id}", id6 = "${lookup(data.oci_identity_group.groups[0], "id")}"
  statements = ["Allow group ${oci_identity_group.t.name} to read instances in compartment ${oci_identity_compartment.t.name}", "Allow group "${lookup(data.oci_identity_group.groups[0], "id")}" to read instances in compartment "${lookup(data.oci_identity_compartment.compartments[0], "id")}"]
}`

	actual := replaceTemplateTokens(original)

	if expected != actual {
		t.Errorf("expected %d, got %d ", expected, actual)
	}
}
