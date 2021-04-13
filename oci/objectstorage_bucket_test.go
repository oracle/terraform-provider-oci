// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v39/common"
	oci_object_storage "github.com/oracle/oci-go-sdk/v39/objectstorage"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	BucketRequiredOnlyResource = BucketResourceDependencies +
		generateResourceFromRepresentationMap("oci_objectstorage_bucket", "test_bucket", Required, Create, bucketRepresentation)

	BucketResourceConfig = BucketResourceDependencies +
		generateResourceFromRepresentationMap("oci_objectstorage_bucket", "test_bucket", Optional, Update, bucketRepresentation)

	// Based on Bucket name specifications used in Object Storage Lifecycle policy
	testBucketName  = randomStringOrHttpReplayValue(32, charset, "bucket")
	testBucketName2 = testBucketName + "2"

	bucketSingularDataSourceRepresentation = map[string]interface{}{
		"name":      Representation{repType: Required, create: testBucketName2},
		"namespace": Representation{repType: Required, create: `${data.oci_objectstorage_namespace.test_namespace.namespace}`},
	}

	bucketDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"namespace":      Representation{repType: Required, create: `${data.oci_objectstorage_namespace.test_namespace.namespace}`},
		"filter":         RepresentationGroup{Required, bucketDataSourceFilterRepresentation}}
	bucketDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `name`},
		"values": Representation{repType: Required, create: []string{`${oci_objectstorage_bucket.test_bucket.name}`}},
	}

	bucketRepresentation = map[string]interface{}{
		"compartment_id":        Representation{repType: Required, create: `${var.compartment_id}`},
		"name":                  Representation{repType: Required, create: testBucketName, update: testBucketName2},
		"namespace":             Representation{repType: Required, create: `${data.oci_objectstorage_namespace.test_namespace.namespace}`},
		"access_type":           Representation{repType: Optional, create: `NoPublicAccess`, update: `ObjectRead`},
		"defined_tags":          Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":         Representation{repType: Optional, create: map[string]string{"Department": "Finance"}, update: map[string]string{"Department": "Accounting"}},
		"kms_key_id":            Representation{repType: Optional, create: `${lookup(data.oci_kms_keys.test_keys_dependency.keys[0], "id")}`},
		"metadata":              Representation{repType: Optional, create: map[string]string{"content-type": "text/plain"}, update: map[string]string{"content-type": "text/xml"}},
		"object_events_enabled": Representation{repType: Optional, create: `false`, update: `true`},
		"storage_tier":          Representation{repType: Optional, create: `Standard`},
		"versioning":            Representation{repType: Optional, create: `Enabled`, update: `Disabled`},
	}

	BucketResourceDependencies = generateDataSourceFromRepresentationMap("oci_objectstorage_namespace", "test_namespace", Required, Create, namespaceSingularDataSourceRepresentation) +
		DefinedTagsDependencies +
		KeyResourceDependencyConfig
)

func TestObjectStorageBucketResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestObjectStorageBucketResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentId2 := getEnvSettingWithBlankDefault("compartment_id_for_update")
	compartmentId2VariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId2)

	resourceName := "oci_objectstorage_bucket.test_bucket"
	datasourceName := "data.oci_objectstorage_bucket_summaries.test_buckets"
	singularDatasourceName := "data.oci_objectstorage_bucket.test_bucket"

	var resId, resId2 string
	// Save TF content to create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	saveConfigContent(config+compartmentIdVariableStr+BucketResourceDependencies+
		generateResourceFromRepresentationMap("oci_objectstorage_bucket", "test_bucket", Optional, Create, bucketRepresentation), "objectstorage", "bucket", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckObjectStorageBucketDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + BucketResourceDependencies +
					generateResourceFromRepresentationMap("oci_objectstorage_bucket", "test_bucket", Required, Create, bucketRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "bucket_id"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "name", testBucketName),
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
				Config: config + compartmentIdVariableStr + BucketResourceDependencies +
					generateResourceFromRepresentationMap("oci_objectstorage_bucket", "test_bucket", Optional, Create, bucketRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "access_type", "NoPublicAccess"),
					resource.TestCheckResourceAttrSet(resourceName, "bucket_id"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(resourceName, "created_by"),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "etag"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "kms_key_id"),
					resource.TestCheckResourceAttr(resourceName, "metadata.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "name", testBucketName),
					resource.TestCheckResourceAttrSet(resourceName, "namespace"),
					resource.TestCheckResourceAttr(resourceName, "object_events_enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "storage_tier", "Standard"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
					resource.TestCheckResourceAttr(resourceName, "versioning", "Enabled"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
							if errExport := testExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
								return errExport
							}
						}
						return err
					},
				),
			},
			{
				Config: config + compartmentIdVariableStr + BucketResourceDependencies +
					generateResourceFromRepresentationMap("oci_objectstorage_bucket", "test_bucket", Optional, Create, bucketRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "approximate_count"),
					resource.TestCheckResourceAttrSet(resourceName, "approximate_size"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},
			// verify updates to compartment
			{
				Config: config + compartmentId2VariableStr + BucketResourceDependencies +
					generateResourceFromRepresentationMap("oci_objectstorage_bucket", "test_bucket", Optional, Create, bucketRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "access_type", "NoPublicAccess"),
					resource.TestCheckResourceAttrSet(resourceName, "bucket_id"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId2),
					resource.TestCheckResourceAttrSet(resourceName, "created_by"),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "etag"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "metadata.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "name", testBucketName),
					resource.TestCheckResourceAttrSet(resourceName, "namespace"),
					resource.TestCheckResourceAttr(resourceName, "storage_tier", "Standard"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
					resource.TestCheckResourceAttrSet(resourceName, "approximate_count"),
					resource.TestCheckResourceAttrSet(resourceName, "approximate_size"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("Resource recreated when it was supposed to be updated.")
						}
						return err
					},
				),
			},
			// verify updates to updatable parameters
			{
				Config: config + compartmentIdVariableStr + BucketResourceDependencies +
					generateResourceFromRepresentationMap("oci_objectstorage_bucket", "test_bucket", Optional, Update, bucketRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "access_type", "ObjectRead"),
					resource.TestCheckResourceAttrSet(resourceName, "bucket_id"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(resourceName, "created_by"),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "etag"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "kms_key_id"),
					resource.TestCheckResourceAttr(resourceName, "metadata.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "name", testBucketName2),
					resource.TestCheckResourceAttrSet(resourceName, "namespace"),
					resource.TestCheckResourceAttr(resourceName, "object_events_enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "storage_tier", "Standard"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
					resource.TestCheckResourceAttr(resourceName, "versioning", "Disabled"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						// The id changes when the name changes
						if resId == resId2 {
							return fmt.Errorf("Resource updated when it was supposed to be recreated.")
						}
						return err
					},
				),
			},
			// verify datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_objectstorage_bucket_summaries", "test_buckets", Optional, Update, bucketDataSourceRepresentation) +
					compartmentIdVariableStr + BucketResourceDependencies +
					generateResourceFromRepresentationMap("oci_objectstorage_bucket", "test_bucket", Optional, Update, bucketRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(datasourceName, "namespace"),

					resource.TestCheckResourceAttr(datasourceName, "bucket_summaries.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "bucket_summaries.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(datasourceName, "bucket_summaries.0.created_by"),
					resource.TestCheckResourceAttr(datasourceName, "bucket_summaries.0.defined_tags.%", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "bucket_summaries.0.etag"),
					resource.TestCheckResourceAttr(datasourceName, "bucket_summaries.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(datasourceName, "bucket_summaries.0.name", testBucketName2),
					resource.TestCheckResourceAttrSet(datasourceName, "bucket_summaries.0.namespace"),
					resource.TestCheckResourceAttrSet(datasourceName, "bucket_summaries.0.time_created"),
				),
			},
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_objectstorage_bucket", "test_bucket", Required, Create, bucketSingularDataSourceRepresentation) +
					compartmentIdVariableStr + BucketResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(singularDatasourceName, "name", testBucketName2),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "namespace"),

					resource.TestCheckResourceAttr(singularDatasourceName, "access_type", "ObjectRead"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "approximate_count"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "approximate_size"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "bucket_id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "created_by"),
					resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "etag"),
					resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "metadata.%", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "name", testBucketName2),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "namespace"),
					// This is difficult to test because TF is eager in creating the datasource and gives stale results.
					// If a depends_on is added, we get an error like "After applying this step and refreshing, the plan was not empty:"
					//resource.TestCheckResourceAttrSet(singularDatasourceName, "object_lifecycle_policy_etag"),
					resource.TestCheckResourceAttr(singularDatasourceName, "object_events_enabled", "true"),
					resource.TestCheckResourceAttr(singularDatasourceName, "storage_tier", "Standard"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
					resource.TestCheckResourceAttr(singularDatasourceName, "versioning", "Disabled"),
				),
			},
			// remove singular datasource from previous step so that it doesn't conflict with import tests
			{
				Config: config + compartmentIdVariableStr + BucketResourceConfig,
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

func testAccCheckObjectStorageBucketDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).objectStorageClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_objectstorage_bucket" {
			noResourceFound = false
			request := oci_object_storage.GetBucketRequest{}

			if value, ok := rs.Primary.Attributes["name"]; ok {
				request.BucketName = &value
			}

			if value, ok := rs.Primary.Attributes["namespace"]; ok {
				request.NamespaceName = &value
			}

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "object_storage")

			_, err := client.GetBucket(context.Background(), request)

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

func init() {
	if DependencyGraph == nil {
		initDependencyGraph()
	}
	if !inSweeperExcludeList("ObjectStorageBucket") {
		resource.AddTestSweepers("ObjectStorageBucket", &resource.Sweeper{
			Name:         "ObjectStorageBucket",
			Dependencies: DependencyGraph["bucket"],
			F:            sweepObjectStorageBucketResource,
		})
	}
}

func sweepObjectStorageBucketResource(compartment string) error {
	objectStorageClient := GetTestClients(&schema.ResourceData{}).objectStorageClient()
	bucketIds, err := getBucketIds(compartment)
	if err != nil {
		return err
	}
	for _, bucketId := range bucketIds {
		if ok := SweeperDefaultResourceId[bucketId]; !ok {
			deleteBucketRequest := oci_object_storage.DeleteBucketRequest{}

			deleteBucketRequest.RequestMetadata.RetryPolicy = getRetryPolicy(true, "object_storage")
			_, error := objectStorageClient.DeleteBucket(context.Background(), deleteBucketRequest)
			if error != nil {
				fmt.Printf("Error deleting Bucket %s %s, It is possible that the resource is already deleted. Please verify manually \n", bucketId, error)
				continue
			}
		}
	}
	return nil
}

func getBucketIds(compartment string) ([]string, error) {
	ids := getResourceIdsToSweep(compartment, "BucketId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	objectStorageClient := GetTestClients(&schema.ResourceData{}).objectStorageClient()

	listBucketsRequest := oci_object_storage.ListBucketsRequest{}
	listBucketsRequest.CompartmentId = &compartmentId

	namespaces, error := getNamespaces(compartment)
	if error != nil {
		return resourceIds, fmt.Errorf("Error getting namespace required for Bucket resource requests \n")
	}
	for _, namespace := range namespaces {
		listBucketsRequest.NamespaceName = &namespace

		listBucketsResponse, err := objectStorageClient.ListBuckets(context.Background(), listBucketsRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting Bucket list for compartment id : %s , %s \n", compartmentId, err)
		}
		for _, bucket := range listBucketsResponse.Items {
			id := *bucket.Name
			resourceIds = append(resourceIds, id)
			addResourceIdToSweeperResourceIdMap(compartmentId, "BucketId", id)
		}

	}
	return resourceIds, nil
}
