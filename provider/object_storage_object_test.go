// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

const (
	ObjectRequiredOnlyResource = ObjectResourceDependencies + `
resource "oci_object_storage_object" "test_object" {
	#Required
	content_length = "${var.object_content_length}"
	bucket = "${var.object_bucket}"
	content = "${var.object_content}"
	namespace = "${var.object_namespace}"
	object = "${var.object_object}"
}
`

	ObjectResourceConfig = ObjectResourceDependencies + `
resource "oci_object_storage_object" "test_object" {
	#Required
	content_length = "${var.object_content_length}"
	bucket = "${var.object_bucket}"
	content = "${var.object_content}"
	namespace = "${var.object_namespace}"
	object = "${var.object_object}"

	#Optional
	content_encoding = "${var.object_content_encoding}"
	content_language = "${var.object_content_language}"
	content_md5 = "${var.object_content_md5}"
	content_type = "${var.object_content_type}"
	metadata = "${var.object_metadata}"
}
`
	ObjectPropertyVariables = `
variable "object_content_encoding" { default = "contentEncoding" }
variable "object_content_language" { default = "contentLanguage" }
variable "object_content_length" { default = 10 }
variable "object_content_md5" { default = "contentMD5" }
variable "object_content_type" { default = "contentType" }
variable "object_bucket" { default = "bucket" }
variable "object_content" { default = "content" }
variable "object_delimiter" { default = "delimiter" }
variable "object_end" { default = "end" }
variable "object_fields" { default = "fields" }
variable "object_metadata" { default = "metadata" }
variable "object_namespace" { default = "namespace" }
variable "object_object" { default = "object" }
variable "object_prefix" { default = "prefix" }
variable "object_start" { default = "start" }

`
	ObjectResourceDependencies = ""
)

func TestObjectStorageObjectResource_basic(t *testing.T) {
	t.Skip("Skipping generated test for now as it has not been worked on.")
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getRequiredEnvSetting("compartment_id_for_create")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	compartmentId2 := getRequiredEnvSetting("compartment_id_for_update")
	compartmentIdVariableStr2 := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId2)

	resourceName := "oci_object_storage_object.test_object"
	datasourceName := "data.oci_object_storage_objects.test_objects"

	var resId, resId2 string

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
					resource.TestCheckResourceAttr(resourceName, "content_length", "10"),
					resource.TestCheckResourceAttr(resourceName, "bucket", "bucket"),
					resource.TestCheckResourceAttr(resourceName, "content", "content"),
					resource.TestCheckResourceAttr(resourceName, "namespace", "namespace"),
					resource.TestCheckResourceAttr(resourceName, "object", "object"),

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
					resource.TestCheckResourceAttr(resourceName, "content_encoding", "contentEncoding"),
					resource.TestCheckResourceAttr(resourceName, "content_language", "contentLanguage"),
					resource.TestCheckResourceAttr(resourceName, "content_length", "10"),
					resource.TestCheckResourceAttr(resourceName, "content_md5", "contentMD5"),
					resource.TestCheckResourceAttr(resourceName, "content_type", "contentType"),
					resource.TestCheckResourceAttr(resourceName, "bucket", "bucket"),
					resource.TestCheckResourceAttr(resourceName, "content", "content"),
					resource.TestCheckResourceAttr(resourceName, "metadata", "metadata"),
					resource.TestCheckResourceAttr(resourceName, "namespace", "namespace"),
					resource.TestCheckResourceAttr(resourceName, "object", "object"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// verify updates to updatable parameters
			{
				Config: config + `
variable "object_content_encoding" { default = "contentEncoding2" }
variable "object_content_language" { default = "contentLanguage2" }
variable "object_content_length" { default = 10 }
variable "object_content_md5" { default = "contentMD52" }
variable "object_content_type" { default = "contentType2" }
variable "object_bucket" { default = "bucket" }
variable "object_content" { default = "content" }
variable "object_delimiter" { default = "delimiter" }
variable "object_end" { default = "end" }
variable "object_fields" { default = "fields" }
variable "object_metadata" { default = "metadata2" }
variable "object_namespace" { default = "namespace" }
variable "object_object" { default = "object" }
variable "object_prefix" { default = "prefix" }
variable "object_start" { default = "start" }

                ` + compartmentIdVariableStr + ObjectResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "content_encoding", "contentEncoding2"),
					resource.TestCheckResourceAttr(resourceName, "content_language", "contentLanguage2"),
					resource.TestCheckResourceAttr(resourceName, "content_length", "10"),
					resource.TestCheckResourceAttr(resourceName, "content_md5", "contentMD52"),
					resource.TestCheckResourceAttr(resourceName, "content_type", "contentType2"),
					resource.TestCheckResourceAttr(resourceName, "bucket", "bucket"),
					resource.TestCheckResourceAttr(resourceName, "content", "content"),
					resource.TestCheckResourceAttr(resourceName, "metadata", "metadata2"),
					resource.TestCheckResourceAttr(resourceName, "namespace", "namespace"),
					resource.TestCheckResourceAttr(resourceName, "object", "object"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("Resource recreated when it was supposed to be updated.")
						}
						return err
					},
				),
			},
			// verify updates to Force New parameters.
			{
				Config: config + `
variable "object_content_encoding" { default = "contentEncoding2" }
variable "object_content_language" { default = "contentLanguage2" }
variable "object_content_length" { default = 11 }
variable "object_content_md5" { default = "contentMD52" }
variable "object_content_type" { default = "contentType2" }
variable "object_bucket" { default = "bucket2" }
variable "object_content" { default = "content2" }
variable "object_delimiter" { default = "delimiter2" }
variable "object_end" { default = "end2" }
variable "object_fields" { default = "fields2" }
variable "object_metadata" { default = "metadata2" }
variable "object_namespace" { default = "namespace2" }
variable "object_object" { default = "object2" }
variable "object_prefix" { default = "prefix2" }
variable "object_start" { default = "start2" }

                ` + compartmentIdVariableStr2 + ObjectResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "content_encoding", "contentEncoding2"),
					resource.TestCheckResourceAttr(resourceName, "content_language", "contentLanguage2"),
					resource.TestCheckResourceAttr(resourceName, "content_length", "11"),
					resource.TestCheckResourceAttr(resourceName, "content_md5", "contentMD52"),
					resource.TestCheckResourceAttr(resourceName, "content_type", "contentType2"),
					resource.TestCheckResourceAttr(resourceName, "bucket", "bucket2"),
					resource.TestCheckResourceAttr(resourceName, "content", "content2"),
					resource.TestCheckResourceAttr(resourceName, "metadata", "metadata2"),
					resource.TestCheckResourceAttr(resourceName, "namespace", "namespace2"),
					resource.TestCheckResourceAttr(resourceName, "object", "object2"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId == resId2 {
							return fmt.Errorf("Resource was expected to be recreated but it wasn't.")
						}
						return err
					},
				),
			},
			// verify datasource
			{
				Config: config + `
variable "object_content_encoding" { default = "contentEncoding2" }
variable "object_content_language" { default = "contentLanguage2" }
variable "object_content_length" { default = 11 }
variable "object_content_md5" { default = "contentMD52" }
variable "object_content_type" { default = "contentType2" }
variable "object_bucket" { default = "bucket2" }
variable "object_content" { default = "content2" }
variable "object_delimiter" { default = "delimiter2" }
variable "object_end" { default = "end2" }
variable "object_fields" { default = "fields2" }
variable "object_metadata" { default = "metadata2" }
variable "object_namespace" { default = "namespace2" }
variable "object_object" { default = "object2" }
variable "object_prefix" { default = "prefix2" }
variable "object_start" { default = "start2" }

data "oci_object_storage_objects" "test_objects" {
	#Required
	bucket = "${var.object_bucket}"
	namespace = "${var.object_namespace}"

	#Optional
	delimiter = "${var.object_delimiter}"
	end = "${var.object_end}"
	fields = "${var.object_fields}"
	prefix = "${var.object_prefix}"
	start = "${var.object_start}"

    filter {
    	name = "id"
    	values = ["${oci_object_storage_object.test_object.id}"]
    }
}
                ` + compartmentIdVariableStr2 + ObjectResourceConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "bucket", "bucket2"),
					resource.TestCheckResourceAttr(datasourceName, "delimiter", "delimiter2"),
					resource.TestCheckResourceAttr(datasourceName, "end", "end2"),
					resource.TestCheckResourceAttr(datasourceName, "fields", "fields2"),
					resource.TestCheckResourceAttr(datasourceName, "namespace", "namespace2"),
					resource.TestCheckResourceAttr(datasourceName, "prefix", "prefix2"),
					resource.TestCheckResourceAttr(datasourceName, "start", "start2"),

					resource.TestCheckResourceAttr(datasourceName, "objects.#", "1"),
				),
			},
		},
	})
}

func TestObjectStorageObjectResource_forcenew(t *testing.T) {
	t.Skip("Skipping generated test for now as it has not been worked on.")
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getRequiredEnvSetting("compartment_id_for_create")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_object_storage_object.test_object"

	var resId string

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify create with optionals
			{
				Config: config + ObjectPropertyVariables + compartmentIdVariableStr + ObjectResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "content_encoding", "contentEncoding"),
					resource.TestCheckResourceAttr(resourceName, "content_language", "contentLanguage"),
					resource.TestCheckResourceAttr(resourceName, "content_length", "10"),
					resource.TestCheckResourceAttr(resourceName, "content_md5", "contentMD5"),
					resource.TestCheckResourceAttr(resourceName, "content_type", "contentType"),
					resource.TestCheckResourceAttr(resourceName, "bucket", "bucket"),
					resource.TestCheckResourceAttr(resourceName, "content", "content"),
					resource.TestCheckResourceAttr(resourceName, "metadata", "metadata"),
					resource.TestCheckResourceAttr(resourceName, "namespace", "namespace"),
					resource.TestCheckResourceAttr(resourceName, "object", "object"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},
			// force new tests, test that changing a parameter would result in creation of a new resource.

		},
	})
}
