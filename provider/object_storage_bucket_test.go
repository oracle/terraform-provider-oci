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
resource "oci_object_storage_bucket" "test_bucket" {
	#Required
	compartment_id = "${var.compartment_id}"
	name = "${var.bucket_name}"
	namespace = "${var.bucket_namespace}"
}
`

	BucketResourceConfig = BucketResourceDependencies + `
resource "oci_object_storage_bucket" "test_bucket" {
	#Required
	compartment_id = "${var.compartment_id}"
	name = "${var.bucket_name}"
	namespace = "${var.bucket_namespace}"

	#Optional
	access_type = "${var.bucket_access_type}"
	metadata = "${var.bucket_metadata}"
}
`
	BucketPropertyVariables = `
variable "bucket_access_type" { default = "accessType" }
variable "bucket_metadata" { default = object{} }
variable "bucket_name" { default = "my-test-1" }
variable "bucket_namespace" { default = "example_namespace" }

`
	BucketResourceDependencies = ""
)

func TestObjectStorageBucketResource_basic(t *testing.T) {
	t.Skip("Skipping generated test for now as it has not been worked on.")
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getRequiredEnvSetting("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_object_storage_bucket.test_bucket"
	datasourceName := "data.oci_object_storage_buckets.test_buckets"

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
					resource.TestCheckResourceAttr(resourceName, "namespace", "example_namespace"),

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
					resource.TestCheckResourceAttr(resourceName, "access_type", "accessType"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(resourceName, "created_by"),
					resource.TestCheckResourceAttrSet(resourceName, "etag"),
					//resource.TestCheckResourceAttr(resourceName, "metadata", object{}),
					resource.TestCheckResourceAttr(resourceName, "name", "my-test-1"),
					resource.TestCheckResourceAttr(resourceName, "namespace", "example_namespace"),
					resource.TestCheckResourceAttrSet(resourceName, "namespace"),
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
variable "bucket_access_type" { default = "accessType2" }
variable "bucket_metadata" { default = object{} }
variable "bucket_name" { default = "name2" }
variable "bucket_namespace" { default = "example_namespace" }

                ` + compartmentIdVariableStr + BucketResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "access_type", "accessType2"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(resourceName, "created_by"),
					resource.TestCheckResourceAttrSet(resourceName, "etag"),
					//resource.TestCheckResourceAttr(resourceName, "metadata", object{}),
					resource.TestCheckResourceAttr(resourceName, "name", "name2"),
					resource.TestCheckResourceAttr(resourceName, "namespace", "example_namespace"),
					resource.TestCheckResourceAttrSet(resourceName, "namespace"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("Resource recreated when it was supposed to be updated.")
						}
						return err
					},
				),
			},
			// verify datasource
			{
				Config: config + `
variable "bucket_access_type" { default = "accessType2" }
variable "bucket_metadata" { default = object{} }
variable "bucket_name" { default = "name2" }
variable "bucket_namespace" { default = "example_namespace" }

data "oci_object_storage_buckets" "test_buckets" {
	#Required
	compartment_id = "${var.compartment_id}"
	namespace = "${var.bucket_namespace}"

    filter {
    	name = "id"
    	values = ["${oci_object_storage_bucket.test_bucket.id}"]
    }
}
                ` + compartmentIdVariableStr + BucketResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "namespace", "example_namespace"),

					resource.TestCheckResourceAttr(datasourceName, "bucket_summaries.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "bucket_summaries.0.access_type", "accessType2"),
					resource.TestCheckResourceAttr(datasourceName, "bucket_summaries.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(datasourceName, "bucket_summaries.0.created_by"),
					resource.TestCheckResourceAttrSet(datasourceName, "bucket_summaries.0.etag"),
					resource.TestCheckResourceAttr(datasourceName, "bucket_summaries.0.metadata", "metadata2"),
					resource.TestCheckResourceAttr(datasourceName, "bucket_summaries.0.name", "name2"),
					resource.TestCheckResourceAttr(datasourceName, "bucket_summaries.0.namespace", "example_namespace"),
					resource.TestCheckResourceAttrSet(datasourceName, "bucket_summaries.0.time_created"),
				),
			},
		},
	})
}
