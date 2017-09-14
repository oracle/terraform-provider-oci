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
