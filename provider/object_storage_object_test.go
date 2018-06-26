// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"testing"

	"crypto/md5"
	"encoding/hex"
	"regexp"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

const (
	ObjectRequiredOnlyResource = ObjectResourceDependencies + `
resource "oci_objectstorage_object" "test_object" {
	#Required
	bucket = "${oci_objectstorage_bucket.test_bucket.name}"
	namespace = "${oci_objectstorage_bucket.test_bucket.namespace}"
	object = "${var.object_object}"
}
`

	ObjectResourceConfig = ObjectResourceDependencies + `
resource "oci_objectstorage_object" "test_object" {
	#Required
	bucket = "${oci_objectstorage_bucket.test_bucket.name}"
	namespace = "${oci_objectstorage_bucket.test_bucket.namespace}"
	object = "${var.object_object}"

	#Optional
	content = "${var.object_content}"
	content_encoding = "${var.object_content_encoding}"
	content_language = "${var.object_content_language}"
	content_type = "${var.object_content_type}"
	metadata = "${var.object_metadata}"
}
`
	ObjectPropertyVariables = `
variable "object_content_encoding" { default = "identity" }
variable "object_content_language" { default = "en-US" }
variable "object_content_type" { default = "text/plain" }
variable "object_content" { default = "content" }
variable "object_metadata" { default = {"content-type" = "text/plain"} }
variable "object_object" { default = "my-test-object-1" }

`
	ObjectResourceDependencies = BucketRequiredOnlyResource + BucketPropertyVariables
)

func TestObjectStorageObjectResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getRequiredEnvSetting("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_objectstorage_object.test_object"
	datasourceName := "data.oci_objectstorage_objects.test_objects"

	var resId, resId2 string
	hexSum := md5.Sum([]byte("content"))
	md5sum := hex.EncodeToString(hexSum[:])
	hexSum2 := md5.Sum([]byte("<a1>content</a1>"))
	md5sum2 := hex.EncodeToString(hexSum2[:])

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify create
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            config + ObjectPropertyVariables + compartmentIdVariableStr + ObjectRequiredOnlyResource,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "bucket", "my-test-1"),
					resource.TestCheckResourceAttrSet(resourceName, "namespace"),
					resource.TestCheckResourceAttr(resourceName, "object", "my-test-object-1"),
					resource.TestCheckResourceAttr(resourceName, "content_type", "application/octet-stream"),
					// New SDK doesn't set omitted values from response, check they are missing from state.
					resource.TestCheckNoResourceAttr(resourceName, "content"),
					resource.TestCheckNoResourceAttr(resourceName, "content_language"),
					resource.TestCheckNoResourceAttr(resourceName, "content_encoding"),
					resource.TestCheckResourceAttr(resourceName, "content_length", "0"),
					resource.TestCheckResourceAttrSet(resourceName, "content_md5"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + ObjectResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + ObjectPropertyVariables + compartmentIdVariableStr + ObjectResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "content_encoding", "identity"),
					resource.TestCheckResourceAttr(resourceName, "content_language", "en-US"),
					resource.TestCheckResourceAttr(resourceName, "content_length", "7"),
					resource.TestCheckResourceAttrSet(resourceName, "content_md5"),
					resource.TestCheckResourceAttr(resourceName, "content_type", "text/plain"),
					resource.TestCheckResourceAttr(resourceName, "bucket", "my-test-1"),
					resource.TestCheckResourceAttrSet(resourceName, "content"),
					resource.TestCheckResourceAttr(resourceName, "content", md5sum),
					resource.TestCheckResourceAttr(resourceName, "metadata.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "metadata.content-type", "text/plain"),
					resource.TestCheckResourceAttrSet(resourceName, "namespace"),
					resource.TestCheckResourceAttr(resourceName, "object", "my-test-object-1"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// verify validations on metadata key
			{
				Config: config + `
variable "object_content_encoding" { default = "identity" }
variable "object_content_language" { default = "en-CA" }
variable "object_content_type" { default = "text/xml" }
variable "object_content" { default = "<a1>content</a1>" }
variable "object_metadata" { default = {"CONTENT-TYPE" = "text/xml"} }
variable "object_object" { default = "my-test-object-1" }

                ` + compartmentIdVariableStr + ObjectResourceConfig,
				ExpectError: regexp.MustCompile("All 'metadata' keys must be lowercase"),
			},
			// verify updates to updatable parameters
			{
				Config: config + `
variable "object_content_encoding" { default = "identity" }
variable "object_content_language" { default = "en-CA" }
variable "object_content_type" { default = "text/xml" }
variable "object_content" { default = "<a1>content</a1>" }
variable "object_metadata" { default = {"content-type" = "text/xml"} }
variable "object_object" { default = "my-test-object-2" }

                ` + compartmentIdVariableStr + ObjectResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "content_encoding", "identity"),
					resource.TestCheckResourceAttr(resourceName, "content_language", "en-CA"),
					resource.TestCheckResourceAttr(resourceName, "content_length", "16"),
					resource.TestCheckResourceAttrSet(resourceName, "content_md5"),
					resource.TestCheckResourceAttr(resourceName, "content_type", "text/xml"),
					resource.TestCheckResourceAttr(resourceName, "bucket", "my-test-1"),
					resource.TestCheckResourceAttrSet(resourceName, "content"),
					resource.TestCheckResourceAttr(resourceName, "content", md5sum2),
					resource.TestCheckResourceAttr(resourceName, "metadata.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "metadata.content-type", "text/xml"),
					resource.TestCheckResourceAttrSet(resourceName, "namespace"),
					resource.TestCheckResourceAttr(resourceName, "object", "my-test-object-2"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						// @CODEGEN 06/2018: Name is part of the id, and hence id will be updated
						if resId == resId2 {
							return fmt.Errorf("Resource updated when it was supposed to be recreated.")
						}
						return err
					},
				),
			},
			// verify updates to name alone
			{
				Config: config + `
variable "object_content_encoding" { default = "identity" }
variable "object_content_language" { default = "en-CA" }
variable "object_content_type" { default = "text/xml" }
variable "object_content" { default = "<a1>content</a1>" }
variable "object_metadata" { default = {"content-type" = "text/xml"} }
variable "object_object" { default = "my-test-object-3" }

                ` + compartmentIdVariableStr + ObjectResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "content_encoding", "identity"),
					resource.TestCheckResourceAttr(resourceName, "content_language", "en-CA"),
					resource.TestCheckResourceAttr(resourceName, "content_length", "16"),
					resource.TestCheckResourceAttrSet(resourceName, "content_md5"),
					resource.TestCheckResourceAttr(resourceName, "content_type", "text/xml"),
					resource.TestCheckResourceAttr(resourceName, "bucket", "my-test-1"),
					resource.TestCheckResourceAttrSet(resourceName, "content"),
					resource.TestCheckResourceAttr(resourceName, "content", md5sum2),
					resource.TestCheckResourceAttr(resourceName, "metadata.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "metadata.content-type", "text/xml"),
					resource.TestCheckResourceAttrSet(resourceName, "namespace"),
					resource.TestCheckResourceAttr(resourceName, "object", "my-test-object-3"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						// @CODEGEN 06/2018: Name is part of the id, and hence id will be updated
						if resId == resId2 {
							return fmt.Errorf("Resource updated when it was supposed to be recreated.")
						}
						return err
					},
				),
			},
			// verify datasource
			{
				Config: config + `
variable "object_content_encoding" { default = "identity" }
variable "object_content_language" { default = "en-CA" }
variable "object_content_type" { default = "text/xml" }
variable "object_content" { default = "<a1>content</a1>" }
variable "object_metadata" { default = {"content-type" = "text/xml"} }
variable "object_object" { default = "my-test-object-1" }

data "oci_objectstorage_objects" "test_objects" {
	#Required
	bucket = "${oci_objectstorage_bucket.test_bucket.name}"
	namespace = "${oci_objectstorage_bucket.test_bucket.namespace}"

    filter {
    	name = "name"
    	values = ["${oci_objectstorage_object.test_object.object}"]
    }
}
                ` + compartmentIdVariableStr + ObjectResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "bucket", "my-test-1"),
					resource.TestCheckResourceAttrSet(datasourceName, "namespace"),

					resource.TestCheckResourceAttr(datasourceName, "objects.#", "1"),
				),
			},
			// verify datasource for delimiter and prefix
			{
				Config: config + `
variable "object_content_encoding" { default = "identity" }
variable "object_content_language" { default = "en-CA" }
variable "object_content_type" { default = "text/xml" }
variable "object_content" { default = "<a1>content</a1>" }
variable "object_metadata" { default = {"content-type" = "text/xml"} }
variable "object_object" { default = "my-test/object-1" }

data "oci_objectstorage_objects" "test_objects" {
	#Required
	bucket = "${oci_objectstorage_bucket.test_bucket.name}"
	namespace = "${oci_objectstorage_bucket.test_bucket.namespace}"

	#Optional
	delimiter = "/"
	prefix = "my-test"
    filter {
    	name = "name"
    	values = ["${oci_objectstorage_object.test_object.object}"]
    }
}
                ` + compartmentIdVariableStr + ObjectResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "bucket", "my-test-1"),
					resource.TestCheckResourceAttrSet(datasourceName, "namespace"),

					resource.TestCheckResourceAttr(datasourceName, "objects.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "delimiter", "/"),
					resource.TestCheckResourceAttr(datasourceName, "prefix", "my-test"),
				),
			},
		},
	})
}
