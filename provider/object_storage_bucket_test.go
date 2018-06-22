// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

const (
	BucketRequiredOnlyResource = BucketResourceDependencies + `
resource "oci_objectstorage_bucket" "test_bucket" {
	#Required
	compartment_id = "${var.compartment_id}"
	name = "${var.bucket_name}"
	namespace = "${data.oci_objectstorage_namespace.t.namespace}"
}
`

	BucketResourceConfig = BucketResourceDependencies + `
resource "oci_objectstorage_bucket" "test_bucket" {
	#Required
	compartment_id = "${var.compartment_id}"
	name = "${var.bucket_name}"
	namespace = "${data.oci_objectstorage_namespace.t.namespace}"

	#Optional
	access_type = "${var.bucket_access_type}"
	defined_tags = "${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "${var.bucket_defined_tags_value}")}"
	freeform_tags = "${var.bucket_freeform_tags}"
	metadata = "${var.bucket_metadata}"
	storage_tier = "${var.bucket_storage_tier}"
}
`
	BucketPropertyVariables = `
variable "bucket_access_type" { default = "NoPublicAccess" }
variable "bucket_defined_tags_value" { default = "value" }
variable "bucket_freeform_tags" { default = {"Department"= "Finance"} }
variable "bucket_metadata" { default = {"content-type" = "text/plain"} }
variable "bucket_name" { default = "my-test-1" }
variable "bucket_storage_tier" { default = "Standard" }

`
	BucketResourceDependencies = DefinedTagsDependencies + `
data "oci_objectstorage_namespace" "t" {
}
`
)

func TestObjectStorageBucketResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getRequiredEnvSetting("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_objectstorage_bucket.test_bucket"
	datasourceName := "data.oci_objectstorage_bucket_summaries.test_buckets"

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
				Config:            config + BucketPropertyVariables + compartmentIdVariableStr + BucketRequiredOnlyResource,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "name", "my-test-1"),
					resource.TestCheckResourceAttrSet(resourceName, "namespace"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + BucketResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + BucketPropertyVariables + compartmentIdVariableStr + BucketResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "access_type", "NoPublicAccess"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(resourceName, "created_by"),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "etag"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "metadata.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "name", "my-test-1"),
					resource.TestCheckResourceAttrSet(resourceName, "namespace"),
					resource.TestCheckResourceAttr(resourceName, "storage_tier", "Standard"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// verify updates to updatable parameters
			{
				Config: config + `
variable "bucket_access_type" { default = "ObjectRead" }
variable "bucket_defined_tags_value" { default = "updatedValue" }
variable "bucket_freeform_tags" { default = {"Department"= "Accounting"} }
variable "bucket_metadata" { default = {"content-type" = "text/xml"} }
variable "bucket_name" { default = "name2" }
variable "bucket_storage_tier" { default = "Standard" }

                ` + compartmentIdVariableStr + BucketResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "access_type", "ObjectRead"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(resourceName, "created_by"),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "etag"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "metadata.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "name", "name2"),
					resource.TestCheckResourceAttrSet(resourceName, "namespace"),
					resource.TestCheckResourceAttr(resourceName, "storage_tier", "Standard"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						// Reverse the check till we have absorbed the changes
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
variable "bucket_access_type" { default = "ObjectRead" }
variable "bucket_defined_tags_value" { default = "updatedValue" }
variable "bucket_freeform_tags" { default = {"Department"= "Accounting"} }
variable "bucket_metadata" { default = {"content-type" = "text/xml"} }
variable "bucket_name" { default = "name2" }
variable "bucket_storage_tier" { default = "Standard" }

data "oci_objectstorage_bucket_summaries" "test_buckets" {
	#Required
	compartment_id = "${var.compartment_id}"
	namespace = "${data.oci_objectstorage_namespace.t.namespace}"

    filter {
    	name = "name"
    	values = ["${oci_objectstorage_bucket.test_bucket.name}"]
    }
}
                ` + compartmentIdVariableStr + BucketResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(datasourceName, "namespace"),

					resource.TestCheckResourceAttr(datasourceName, "bucket_summaries.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "bucket_summaries.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(datasourceName, "bucket_summaries.0.created_by"),
					resource.TestCheckResourceAttr(datasourceName, "bucket_summaries.0.defined_tags.%", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "bucket_summaries.0.etag"),
					resource.TestCheckResourceAttr(datasourceName, "bucket_summaries.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(datasourceName, "bucket_summaries.0.name", "name2"),
					resource.TestCheckResourceAttrSet(datasourceName, "bucket_summaries.0.namespace"),
					resource.TestCheckResourceAttrSet(datasourceName, "bucket_summaries.0.time_created"),
				),
			},
		},
	})
}
