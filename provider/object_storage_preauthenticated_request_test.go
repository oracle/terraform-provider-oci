// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/oracle/oci-go-sdk/common"
	oci_object_storage "github.com/oracle/oci-go-sdk/objectstorage"
)

const (
	PreauthenticatedRequestRequiredOnlyResource = PreauthenticatedRequestResourceDependencies + `
resource "oci_objectstorage_preauthrequest" "test_preauthenticated_request" {
	#Required
	access_type = "AnyObjectWrite"
	bucket = "${var.preauthenticated_request_bucket}"
	name = "${var.preauthenticated_request_name}"
	namespace = "${oci_objectstorage_bucket.test_bucket.namespace}"
	time_expires = "${var.preauthenticated_request_time_expires}"
}
`

	PreauthenticatedRequestResourceConfig = PreauthenticatedRequestResourceDependencies + `
resource "oci_objectstorage_preauthrequest" "test_preauthenticated_request" {
	#Required
	access_type = "${var.preauthenticated_request_access_type}"
	bucket = "${var.preauthenticated_request_bucket}"
	name = "${var.preauthenticated_request_name}"
	namespace = "${oci_objectstorage_bucket.test_bucket.namespace}"
	time_expires = "${var.preauthenticated_request_time_expires}"

	#Optional
	object = "${var.preauthenticated_request_object}"
}
`
	PreauthenticatedRequestPropertyVariables = `
variable "preauthenticated_request_access_type" { default = "ObjectRead" }
variable "preauthenticated_request_bucket" { default = "my-test-1" }
variable "preauthenticated_request_name" { default = "-tf-par" }
variable "preauthenticated_request_namespace" { default = "namespace" }
variable "preauthenticated_request_object" { default = "my-test-object-1" }
variable "preauthenticated_request_object_name_prefix" { default = "my-test-object" }
variable "preauthenticated_request_time_expires" { default = "2020-01-01T00:00:00Z" }

`
	PreauthenticatedRequestResourceDependencies = ObjectRequiredOnlyResource + ObjectPropertyVariables
)

func TestObjectStoragePreauthenticatedRequestResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getRequiredEnvSetting("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_objectstorage_preauthrequest.test_preauthenticated_request"
	datasourceName := "data.oci_objectstorage_preauthrequests.test_preauthenticated_requests"
	singularDatasourceName := "data.oci_objectstorage_preauthrequest.test_preauthenticated_request"

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckObjectStoragePreauthenticatedRequestDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + PreauthenticatedRequestPropertyVariables + compartmentIdVariableStr + PreauthenticatedRequestRequiredOnlyResource,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "access_type", "AnyObjectWrite"),
					resource.TestCheckResourceAttr(resourceName, "bucket", "my-test-1"),
					resource.TestCheckResourceAttr(resourceName, "name", "-tf-par"),
					resource.TestCheckResourceAttrSet(resourceName, "namespace"),
					resource.TestCheckResourceAttr(resourceName, "time_expires", "2020-01-01T00:00:00Z"),
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + PreauthenticatedRequestResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + PreauthenticatedRequestPropertyVariables + compartmentIdVariableStr + PreauthenticatedRequestResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "access_type", "ObjectRead"),
					resource.TestCheckResourceAttr(resourceName, "bucket", "my-test-1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "name", "-tf-par"),
					resource.TestCheckResourceAttrSet(resourceName, "namespace"),
					resource.TestCheckResourceAttr(resourceName, "object", "my-test-object-1"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
					resource.TestCheckResourceAttr(resourceName, "time_expires", "2020-01-01T00:00:00Z"),
				),
			},

			// verify datasource
			{
				Config: config + `
variable "preauthenticated_request_access_type" { default = "ObjectRead" }
variable "preauthenticated_request_bucket" { default = "my-test-1" }
variable "preauthenticated_request_name" { default = "-tf-par" }
variable "preauthenticated_request_namespace" { default = "namespace" }
variable "preauthenticated_request_object" { default = "my-test-object-1" }
variable "preauthenticated_request_object_name_prefix" { default = "my-test-object" }
variable "preauthenticated_request_time_expires" { default = "2020-01-01T00:00:00Z" }

data "oci_objectstorage_preauthrequests" "test_preauthenticated_requests" {
	#Required
	bucket = "${var.preauthenticated_request_bucket}"
	namespace = "${oci_objectstorage_bucket.test_bucket.namespace}"

	#Optional
	object_name_prefix = "${var.preauthenticated_request_object_name_prefix}"

    filter {
    	name = "id"
    	values = ["${oci_objectstorage_preauthrequest.test_preauthenticated_request.id}"]
    }
}
                ` + compartmentIdVariableStr + PreauthenticatedRequestResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "bucket", "my-test-1"),
					resource.TestCheckResourceAttrSet(datasourceName, "namespace"),
					resource.TestCheckResourceAttr(datasourceName, "object_name_prefix", "my-test-object"),

					resource.TestCheckResourceAttr(datasourceName, "preauthenticated_requests.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "preauthenticated_requests.0.access_type", "ObjectRead"),
					resource.TestCheckResourceAttrSet(datasourceName, "preauthenticated_requests.0.id"),
					resource.TestCheckResourceAttr(datasourceName, "preauthenticated_requests.0.name", "-tf-par"),
					resource.TestCheckResourceAttr(datasourceName, "preauthenticated_requests.0.object", "my-test-object-1"),
					resource.TestCheckResourceAttrSet(datasourceName, "preauthenticated_requests.0.time_created"),
					resource.TestCheckResourceAttr(datasourceName, "preauthenticated_requests.0.time_expires", "2020-01-01 00:00:00 +0000 UTC"),
				),
			},
			// verify singular datasource
			{
				Config: config + `
variable "preauthenticated_request_access_type" { default = "ObjectRead" }
variable "preauthenticated_request_bucket" { default = "my-test-1" }
variable "preauthenticated_request_name" { default = "-tf-par" }
variable "preauthenticated_request_namespace" { default = "namespace" }
variable "preauthenticated_request_object" { default = "my-test-object-1" }
variable "preauthenticated_request_object_name_prefix" { default = "my-test-object" }
variable "preauthenticated_request_time_expires" { default = "2020-01-01T00:00:00Z" }

data "oci_objectstorage_preauthrequest" "test_preauthenticated_request" {
	#Required
	bucket = "${var.preauthenticated_request_bucket}"
	namespace = "${oci_objectstorage_bucket.test_bucket.namespace}"
	par_id = "${oci_objectstorage_preauthrequest.test_preauthenticated_request.id}"
}
                ` + compartmentIdVariableStr + PreauthenticatedRequestResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(singularDatasourceName, "bucket", "my-test-1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "namespace"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "par_id"),

					resource.TestCheckResourceAttr(singularDatasourceName, "access_type", "ObjectRead"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "name", "-tf-par"),
					resource.TestCheckResourceAttr(singularDatasourceName, "object", "my-test-object-1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
					resource.TestCheckResourceAttr(singularDatasourceName, "time_expires", "2020-01-01 00:00:00 +0000 UTC"),
				),
			},
		},
	})
}

func testAccCheckObjectStoragePreauthenticatedRequestDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).objectStorageClient
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_objectstorage_preauthenticated_request" {
			noResourceFound = false
			request := oci_object_storage.GetPreauthenticatedRequestRequest{}

			if value, ok := rs.Primary.Attributes["bucket"]; ok {
				request.BucketName = &value
			}

			if value, ok := rs.Primary.Attributes["namespace"]; ok {
				request.NamespaceName = &value
			}

			tmp := rs.Primary.ID
			request.ParId = &tmp

			_, err := client.GetPreauthenticatedRequest(context.Background(), request)

			if err == nil {
				return fmt.Errorf("resource still exists")
			}
			//Verify that exception is for '404 not found'.
			if failure, isServiceError := common.IsServiceError(err); !isServiceError || failure.GetHTTPStatusCode() != 404 {
				return err
			}
		}
	}
	if noResourceFound {
		return fmt.Errorf("at least one resource was expected from the state file, but could not be found")
	}

	return nil
}
