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

var (
	ObjectLifecyclePolicyRequiredOnlyResource = ObjectLifecyclePolicyResourceDependencies + `
resource "oci_objectstorage_object_lifecycle_policy" "test_object_lifecycle_policy" {
	#Required
	bucket = "${oci_objectstorage_bucket.test_bucket.name}"
	namespace = "${data.oci_objectstorage_namespace.t.namespace}"
}
`

	ObjectLifecyclePolicyResourceConfig = ObjectLifecyclePolicyResourceDependencies + `
resource "oci_objectstorage_object_lifecycle_policy" "test_object_lifecycle_policy" {
	#Required
	bucket = "${oci_objectstorage_bucket.test_bucket.name}"
	namespace = "${data.oci_objectstorage_namespace.t.namespace}"

	#Optional
	rules {
		#Required
		action = "${var.object_lifecycle_policy_rules_action}"
		is_enabled = "${var.object_lifecycle_policy_rules_is_enabled}"
		name = "${var.object_lifecycle_policy_rules_name}"
		time_amount = "${var.object_lifecycle_policy_rules_time_amount}"
		time_unit = "${var.object_lifecycle_policy_rules_time_unit}"

		#Optional
		object_name_filter {

			#Optional
			inclusion_prefixes = "${var.object_lifecycle_policy_rules_object_name_filter_inclusion_prefixes}"
		}
	}
}
`
	ObjectLifecyclePolicyPropertyVariables = `
variable "object_lifecycle_policy_bucket" { default = "my-test-1" }
variable "object_lifecycle_policy_namespace" { default = "namespace" }
variable "object_lifecycle_policy_rules_action" { default = "ARCHIVE" }
variable "object_lifecycle_policy_rules_is_enabled" { default = false }
variable "object_lifecycle_policy_rules_name" { default = "sampleRule" }
variable "object_lifecycle_policy_rules_object_name_filter_inclusion_prefixes" { default = ["my-test-1","my-test-2"] }
variable "object_lifecycle_policy_rules_time_amount" { default = 10 }
variable "object_lifecycle_policy_rules_time_unit" { default = "DAYS" }

`
	ObjectLifecyclePolicyResourceDependencies = BucketRequiredOnlyResource
)

func TestObjectStorageObjectLifecyclePolicyResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_objectstorage_object_lifecycle_policy.test_object_lifecycle_policy"

	singularDatasourceName := "data.oci_objectstorage_object_lifecycle_policy.test_object_lifecycle_policy"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckObjectStorageObjectLifecyclePolicyDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + ObjectLifecyclePolicyPropertyVariables + compartmentIdVariableStr + ObjectLifecyclePolicyRequiredOnlyResource,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "bucket", "my-test-1"),
					resource.TestCheckResourceAttr(resourceName, "rules.#", "0"),
					resource.TestCheckResourceAttrSet(resourceName, "namespace"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + ObjectLifecyclePolicyResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + ObjectLifecyclePolicyPropertyVariables + compartmentIdVariableStr + ObjectLifecyclePolicyResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "bucket", "my-test-1"),
					resource.TestCheckResourceAttrSet(resourceName, "namespace"),
					resource.TestCheckResourceAttr(resourceName, "rules.#", "1"),
					CheckResourceSetContainsElementWithProperties(resourceName, "rules", map[string]string{
						"action":                                    "ARCHIVE",
						"is_enabled":                                "false",
						"name":                                      "sampleRule",
						"object_name_filter.#":                      "1",
						"object_name_filter.0.inclusion_prefixes.#": "2",
						"time_amount":                               "10",
						"time_unit":                                 "DAYS",
					},
						[]string{}),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// verify updates to updatable parameters
			{
				Config: config + `
variable "object_lifecycle_policy_bucket" { default = "my-test-1" }
variable "object_lifecycle_policy_namespace" { default = "namespace" }
variable "object_lifecycle_policy_rules_action" { default = "DELETE" }
variable "object_lifecycle_policy_rules_is_enabled" { default = true }
variable "object_lifecycle_policy_rules_name" { default = "name2" }
variable "object_lifecycle_policy_rules_object_name_filter_inclusion_prefixes" { default = ["my-test-1","my-test-2","my-test-3"] }
variable "object_lifecycle_policy_rules_time_amount" { default = 11 }
variable "object_lifecycle_policy_rules_time_unit" { default = "YEARS" }

                ` + compartmentIdVariableStr + ObjectLifecyclePolicyResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "bucket", "my-test-1"),
					resource.TestCheckResourceAttrSet(resourceName, "namespace"),
					resource.TestCheckResourceAttr(resourceName, "rules.#", "1"),
					CheckResourceSetContainsElementWithProperties(resourceName, "rules", map[string]string{
						"action":                                    "DELETE",
						"is_enabled":                                "true",
						"name":                                      "name2",
						"object_name_filter.#":                      "1",
						"object_name_filter.0.inclusion_prefixes.#": "3",
						"time_amount":                               "11",
						"time_unit":                                 "YEARS",
					},
						[]string{}),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("Resource recreated when it was supposed to be updated.")
						}
						return err
					},
				),
			},
			// verify singular datasource
			{
				Config: config + `
variable "object_lifecycle_policy_bucket" { default = "my-test-1" }
variable "object_lifecycle_policy_namespace" { default = "namespace" }
variable "object_lifecycle_policy_rules_action" { default = "DELETE" }
variable "object_lifecycle_policy_rules_is_enabled" { default = true }
variable "object_lifecycle_policy_rules_name" { default = "name2" }
variable "object_lifecycle_policy_rules_object_name_filter_inclusion_prefixes" { default = ["my-test-1","my-test-2","my-test-3"] }
variable "object_lifecycle_policy_rules_time_amount" { default = 11 }
variable "object_lifecycle_policy_rules_time_unit" { default = "YEARS" }

data "oci_objectstorage_object_lifecycle_policy" "test_object_lifecycle_policy" {
	#Required
	bucket = "${var.object_lifecycle_policy_bucket}"
	namespace = "${data.oci_objectstorage_namespace.t.namespace}"
}
                ` + compartmentIdVariableStr + ObjectLifecyclePolicyResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(singularDatasourceName, "bucket", "my-test-1"),
					resource.TestCheckResourceAttrSet(resourceName, "namespace"),
					resource.TestCheckResourceAttr(singularDatasourceName, "rules.#", "1"),
					CheckResourceSetContainsElementWithProperties(singularDatasourceName, "rules", map[string]string{},
						[]string{}),

					resource.TestCheckResourceAttr(singularDatasourceName, "rules.#", "1"),
					CheckResourceSetContainsElementWithProperties(singularDatasourceName, "rules", map[string]string{
						"action":                                    "DELETE",
						"is_enabled":                                "true",
						"name":                                      "name2",
						"object_name_filter.#":                      "1",
						"object_name_filter.0.inclusion_prefixes.#": "3",
						"time_amount":                               "11",
						"time_unit":                                 "YEARS",
					},
						[]string{}),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				),
			},
			// remove singular datasource from previous step so that it doesn't conflict with import tests
			{
				Config: config + `
variable "object_lifecycle_policy_bucket" { default = "my-test-1" }
variable "object_lifecycle_policy_namespace" { default = "namespace" }
variable "object_lifecycle_policy_rules_action" { default = "DELETE" }
variable "object_lifecycle_policy_rules_is_enabled" { default = true }
variable "object_lifecycle_policy_rules_name" { default = "name2" }
variable "object_lifecycle_policy_rules_object_name_filter_inclusion_prefixes" { default = ["my-test-1","my-test-2","my-test-3"] }
variable "object_lifecycle_policy_rules_time_amount" { default = 11 }
variable "object_lifecycle_policy_rules_time_unit" { default = "YEARS" }

                ` + compartmentIdVariableStr + ObjectLifecyclePolicyResourceConfig,
			},
			// verify resource import
			{
				Config:                  config,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
				ResourceName:            resourceName,
			},
		},
	})
}

func TestObjectStorageObjectLifecyclePolicyResource_validations(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_objectstorage_object_lifecycle_policy.test_object_lifecycle_policy"

	var resId string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckObjectStorageObjectLifecyclePolicyDestroy,
		Steps: []resource.TestStep{
			// verify baseline create
			{
				Config: config + compartmentIdVariableStr + ObjectLifecyclePolicyResourceConfig + `
variable "object_lifecycle_policy_bucket" { default = "my-test-1" }
variable "object_lifecycle_policy_namespace" { default = "namespace" }
variable "object_lifecycle_policy_rules_action" { default = "ARCHIVE" }
variable "object_lifecycle_policy_rules_is_enabled" { default = false }
variable "object_lifecycle_policy_rules_name" { default = "sampleRule" }
variable "object_lifecycle_policy_rules_object_name_filter_inclusion_prefixes" { default = ["my-test-1","my-test-2"] }
variable "object_lifecycle_policy_rules_time_amount" { default = 10 }
variable "object_lifecycle_policy_rules_time_unit" { default = "DAYS" }
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "bucket", "my-test-1"),
					resource.TestCheckResourceAttr(resourceName, "rules.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "namespace"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},
			// change order of inclusion prefixes
			{
				Config: config + compartmentIdVariableStr + ObjectLifecyclePolicyResourceConfig + `
variable "object_lifecycle_policy_bucket" { default = "my-test-1" }
variable "object_lifecycle_policy_namespace" { default = "namespace" }
variable "object_lifecycle_policy_rules_action" { default = "ARCHIVE" }
variable "object_lifecycle_policy_rules_is_enabled" { default = false }
variable "object_lifecycle_policy_rules_name" { default = "sampleRule" }
variable "object_lifecycle_policy_rules_object_name_filter_inclusion_prefixes" { default = ["my-test-2", "my-test-1"] }
variable "object_lifecycle_policy_rules_time_amount" { default = 10 }
variable "object_lifecycle_policy_rules_time_unit" { default = "DAYS" }
`,
				PlanOnly:           true,
				ExpectNonEmptyPlan: false,
			},
			// Remove inclusion prefixes to see plan has changed
			{
				Config: config + compartmentIdVariableStr + ObjectLifecyclePolicyResourceConfig + `
variable "object_lifecycle_policy_bucket" { default = "my-test-1" }
variable "object_lifecycle_policy_namespace" { default = "namespace" }
variable "object_lifecycle_policy_rules_action" { default = "ARCHIVE" }
variable "object_lifecycle_policy_rules_is_enabled" { default = false }
variable "object_lifecycle_policy_rules_name" { default = "sampleRule" }
variable "object_lifecycle_policy_rules_object_name_filter_inclusion_prefixes" { default = [] }
variable "object_lifecycle_policy_rules_time_amount" { default = 10 }
variable "object_lifecycle_policy_rules_time_unit" { default = "DAYS" }
`,
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func testAccCheckObjectStorageObjectLifecyclePolicyDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).objectStorageClient
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_objectstorage_object_lifecycle_policy" {
			noResourceFound = false
			request := oci_object_storage.GetObjectLifecyclePolicyRequest{}

			if value, ok := rs.Primary.Attributes["bucket"]; ok {
				request.BucketName = &value
			}

			if value, ok := rs.Primary.Attributes["namespace"]; ok {
				request.NamespaceName = &value
			}

			_, err := client.GetObjectLifecyclePolicy(context.Background(), request)

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
