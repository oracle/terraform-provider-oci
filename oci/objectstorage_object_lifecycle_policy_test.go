// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/oracle/oci-go-sdk/v25/common"
	oci_object_storage "github.com/oracle/oci-go-sdk/v25/objectstorage"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	ObjectLifecyclePolicyRequiredOnlyResource = ObjectLifecyclePolicyResourceDependencies +
		generateResourceFromRepresentationMap("oci_objectstorage_object_lifecycle_policy", "test_object_lifecycle_policy", Required, Create, objectLifecyclePolicyRepresentation)

	ObjectLifecyclePolicyResourceConfig = ObjectLifecyclePolicyResourceDependencies +
		generateResourceFromRepresentationMap("oci_objectstorage_object_lifecycle_policy", "test_object_lifecycle_policy", Optional, Update, objectLifecyclePolicyRepresentation)

	bucketName  = randomStringOrHttpReplayValue(32, charset, "bucket1")
	bucketName2 = randomStringOrHttpReplayValue(32, charset, "bucket2")
	bucketName3 = randomStringOrHttpReplayValue(32, charset, "bucket3")

	objectLifecyclePolicySingularDataSourceRepresentation = map[string]interface{}{
		"bucket":    Representation{repType: Required, create: bucketName},
		"namespace": Representation{repType: Required, create: `${data.oci_objectstorage_namespace.test_namespace.namespace}`},
	}

	objectLifecyclePolicyRepresentation = map[string]interface{}{
		"bucket":    Representation{repType: Required, create: bucketName},
		"namespace": Representation{repType: Required, create: `${data.oci_objectstorage_namespace.test_namespace.namespace}`},
		"rules":     RepresentationGroup{Optional, objectLifecyclePolicyRulesRepresentation},
	}
	objectLifecyclePolicyRulesRepresentation = map[string]interface{}{
		"action":             Representation{repType: Required, create: `ARCHIVE`, update: `DELETE`},
		"is_enabled":         Representation{repType: Required, create: `false`, update: `true`},
		"name":               Representation{repType: Required, create: `sampleRule`, update: `name2`},
		"time_amount":        Representation{repType: Required, create: `10`, update: `11`},
		"time_unit":          Representation{repType: Required, create: `DAYS`, update: `YEARS`},
		"object_name_filter": RepresentationGroup{Optional, objectLifecyclePolicyRulesObjectNameFilterRepresentation},
		"target":             Representation{repType: Optional, create: `objects`},
	}
	objectLifecyclePolicyRulesObjectNameFilterRepresentation = map[string]interface{}{
		"exclusion_patterns": Representation{repType: Optional, create: []string{`exclusionPattern1`, `exclusionPattern2`}, update: []string{`exclusionPattern1`, `exclusionPattern2`, `exclusionPattern3`}},
		"inclusion_patterns": Representation{repType: Optional, create: []string{`inclusionPattern1`, `inclusionPattern2`}, update: []string{`inclusionPattern1`, `inclusionPattern2`, `inclusionPattern3`}},
		"inclusion_prefixes": Representation{repType: Optional, create: []string{bucketName, bucketName2}, update: []string{bucketName, bucketName2, bucketName3}},
	}

	ObjectLifecyclePolicyResourceDependencies = generateResourceFromRepresentationMap("oci_objectstorage_bucket", "test_bucket", Required, Create, getUpdatedRepresentationCopy("name", Representation{repType: Required, create: bucketName}, bucketRepresentation)) +
		generateDataSourceFromRepresentationMap("oci_objectstorage_namespace", "test_namespace", Required, Create, namespaceSingularDataSourceRepresentation)
)

func TestObjectStorageObjectLifecyclePolicyResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestObjectStorageObjectLifecyclePolicyResource_basic")
	defer httpreplay.SaveScenario()

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
				Config: config + compartmentIdVariableStr + ObjectLifecyclePolicyResourceDependencies +
					generateResourceFromRepresentationMap("oci_objectstorage_object_lifecycle_policy", "test_object_lifecycle_policy", Required, Create, objectLifecyclePolicyRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "bucket", bucketName),
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
				Config: config + compartmentIdVariableStr + ObjectLifecyclePolicyResourceDependencies +
					generateResourceFromRepresentationMap("oci_objectstorage_object_lifecycle_policy", "test_object_lifecycle_policy", Optional, Create, objectLifecyclePolicyRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "bucket", bucketName),
					resource.TestCheckResourceAttrSet(resourceName, "namespace"),
					resource.TestCheckResourceAttr(resourceName, "rules.#", "1"),
					CheckResourceSetContainsElementWithProperties(resourceName, "rules", map[string]string{
						"action":               "ARCHIVE",
						"is_enabled":           "false",
						"name":                 "sampleRule",
						"object_name_filter.#": "1",
						"object_name_filter.0.inclusion_prefixes.#": "2",
						"object_name_filter.0.exclusion_patterns.#": "2",
						"object_name_filter.0.inclusion_patterns.#": "2",
						"target":      "objects",
						"time_amount": "10",
						"time_unit":   "DAYS",
					},
						[]string{}),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "false")); isEnableExportCompartment {
							if errExport := testExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
								return errExport
							}
						}
						return err
					},
				),
			},

			// verify updates to updatable parameters
			{
				Config: config + compartmentIdVariableStr + ObjectLifecyclePolicyResourceDependencies +
					generateResourceFromRepresentationMap("oci_objectstorage_object_lifecycle_policy", "test_object_lifecycle_policy", Optional, Update, objectLifecyclePolicyRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "bucket", bucketName),
					resource.TestCheckResourceAttrSet(resourceName, "namespace"),
					resource.TestCheckResourceAttr(resourceName, "rules.#", "1"),
					CheckResourceSetContainsElementWithProperties(resourceName, "rules", map[string]string{
						"action":               "DELETE",
						"is_enabled":           "true",
						"name":                 "name2",
						"object_name_filter.#": "1",
						"object_name_filter.0.inclusion_prefixes.#": "3",
						"object_name_filter.0.exclusion_patterns.#": "3",
						"object_name_filter.0.inclusion_patterns.#": "3",
						"target":      "objects",
						"time_amount": "11",
						"time_unit":   "YEARS",
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
				Config: config +
					generateDataSourceFromRepresentationMap("oci_objectstorage_object_lifecycle_policy", "test_object_lifecycle_policy", Required, Create, objectLifecyclePolicySingularDataSourceRepresentation) +
					compartmentIdVariableStr + ObjectLifecyclePolicyResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(singularDatasourceName, "bucket", bucketName),
					resource.TestCheckResourceAttrSet(resourceName, "namespace"),
					resource.TestCheckResourceAttr(singularDatasourceName, "rules.#", "1"),
					CheckResourceSetContainsElementWithProperties(singularDatasourceName, "rules", map[string]string{},
						[]string{}),

					resource.TestCheckResourceAttr(singularDatasourceName, "rules.#", "1"),
					CheckResourceSetContainsElementWithProperties(singularDatasourceName, "rules", map[string]string{
						"action":               "DELETE",
						"is_enabled":           "true",
						"name":                 "name2",
						"object_name_filter.#": "1",
						"object_name_filter.0.inclusion_prefixes.#": "3",
						"object_name_filter.0.exclusion_patterns.#": "3",
						"object_name_filter.0.inclusion_patterns.#": "3",
						"target":      "objects",
						"time_amount": "11",
						"time_unit":   "YEARS",
					},
						[]string{}),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				),
			},
			// remove singular datasource from previous step so that it doesn't conflict with import tests
			{
				Config: config + compartmentIdVariableStr + ObjectLifecyclePolicyResourceConfig,
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

func testAccCheckObjectStorageObjectLifecyclePolicyDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).objectStorageClient()
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

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "object_storage")

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
