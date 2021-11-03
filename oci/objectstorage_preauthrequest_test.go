// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v54/common"
	oci_object_storage "github.com/oracle/oci-go-sdk/v54/objectstorage"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	testPreAuthBucketName = RandomStringOrHttpReplayValue(32, Charset, "bucketPreAuth")

	expirationTimeForPar = time.Now().UTC().AddDate(0, 0, 1).Truncate(time.Millisecond)

	PreauthenticatedRequestRequiredOnlyResource = PreauthenticatedRequestResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_objectstorage_preauthrequest", "test_preauthenticated_request", Required, Create, preauthenticatedRequestRepresentation)

	PreauthenticatedRequestResourceConfig = PreauthenticatedRequestResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_objectstorage_preauthrequest", "test_preauthenticated_request", Optional, Update, preauthenticatedRequestRepresentation)

	preauthenticatedRequestSingularDataSourceRepresentation = map[string]interface{}{
		"bucket":    Representation{RepType: Required, Create: `${oci_objectstorage_bucket.test_bucket.name}`},
		"namespace": Representation{RepType: Required, Create: `${oci_objectstorage_bucket.test_bucket.namespace}`},
		"par_id":    Representation{RepType: Required, Create: `${oci_objectstorage_preauthrequest.test_preauthenticated_request.par_id}`},
	}

	preauthenticatedRequestDataSourceRepresentation = map[string]interface{}{
		"bucket":             Representation{RepType: Required, Create: `${oci_objectstorage_bucket.test_bucket.name}`},
		"namespace":          Representation{RepType: Required, Create: `${oci_objectstorage_bucket.test_bucket.namespace}`},
		"object_name_prefix": Representation{RepType: Optional, Create: `my-test-object`},
		"filter":             RepresentationGroup{Required, preauthenticatedRequestDataSourceFilterRepresentation}}
	preauthenticatedRequestDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{RepType: Required, Create: `id`},
		"values": Representation{RepType: Required, Create: []string{`${oci_objectstorage_preauthrequest.test_preauthenticated_request.par_id}`}},
	}

	preauthenticatedRequestRepresentation = map[string]interface{}{
		"access_type":           Representation{RepType: Required, Create: `AnyObjectWrite`, Update: `ObjectRead`},
		"bucket":                Representation{RepType: Required, Create: `${oci_objectstorage_bucket.test_bucket.name}`},
		"name":                  Representation{RepType: Required, Create: `-tf-par`},
		"namespace":             Representation{RepType: Required, Create: `${oci_objectstorage_bucket.test_bucket.namespace}`},
		"time_expires":          Representation{RepType: Required, Create: expirationTimeForPar.Format(time.RFC3339Nano)},
		"object":                Representation{RepType: Optional, Create: `my-test-object-1`},
		"bucket_listing_action": Representation{RepType: Optional, Create: ``},
	}

	PreauthenticatedRequestResourceDependencies = GenerateResourceFromRepresentationMap("oci_objectstorage_bucket", "test_bucket", Required, Create, GetUpdatedRepresentationCopy("name", Representation{RepType: Required, Create: testPreAuthBucketName}, bucketRepresentation)) +
		GenerateDataSourceFromRepresentationMap("oci_objectstorage_namespace", "test_namespace", Optional, Create, namespaceSingularDataSourceRepresentation) +
		GenerateResourceFromRepresentationMap("oci_objectstorage_object", "test_object", Required, Create, objectRepresentation)
)

// issue-routing-tag: object_storage/default
func TestObjectStoragePreauthenticatedRequestResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestObjectStoragePreauthenticatedRequestResource_basic")
	defer httpreplay.SaveScenario()

	config := ProviderTestConfig()

	compartmentId := GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_objectstorage_preauthrequest.test_preauthenticated_request"
	datasourceName := "data.oci_objectstorage_preauthrequests.test_preauthenticated_requests"
	singularDatasourceName := "data.oci_objectstorage_preauthrequest.test_preauthenticated_request"

	var resId string

	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	SaveConfigContent(config+compartmentIdVariableStr+PreauthenticatedRequestResourceDependencies+
		GenerateResourceFromRepresentationMap("oci_objectstorage_preauthrequest", "test_preauthenticated_request", Optional, Create, preauthenticatedRequestRepresentation), "objectstorage", "preauthenticatedRequest", t)

	ResourceTest(t, testAccCheckObjectStoragePreauthenticatedRequestDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + PreauthenticatedRequestResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_objectstorage_preauthrequest", "test_preauthenticated_request", Required, Create, preauthenticatedRequestRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "access_type", "AnyObjectWrite"),
				resource.TestCheckResourceAttr(resourceName, "bucket", testPreAuthBucketName),
				resource.TestCheckResourceAttr(resourceName, "name", "-tf-par"),
				resource.TestCheckResourceAttrSet(resourceName, "namespace"),
				resource.TestCheckResourceAttr(resourceName, "time_expires", expirationTimeForPar.Format(time.RFC3339Nano)),
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + PreauthenticatedRequestResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + PreauthenticatedRequestResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_objectstorage_preauthrequest", "test_preauthenticated_request", Optional, Update, preauthenticatedRequestRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "access_type", "ObjectRead"),
				resource.TestCheckResourceAttr(resourceName, "bucket_listing_action", ""),
				resource.TestCheckResourceAttrSet(resourceName, "access_uri"),
				resource.TestCheckResourceAttr(resourceName, "bucket", testPreAuthBucketName),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "name", "-tf-par"),
				resource.TestCheckResourceAttrSet(resourceName, "namespace"),
				resource.TestCheckResourceAttr(resourceName, "object", "my-test-object-1"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "time_expires", expirationTimeForPar.Format(time.RFC3339Nano)),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_objectstorage_preauthrequests", "test_preauthenticated_requests", Optional, Update, preauthenticatedRequestDataSourceRepresentation) +
				compartmentIdVariableStr + PreauthenticatedRequestResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_objectstorage_preauthrequest", "test_preauthenticated_request", Optional, Update, preauthenticatedRequestRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "bucket", testPreAuthBucketName),
				resource.TestCheckResourceAttrSet(datasourceName, "namespace"),
				resource.TestCheckResourceAttr(datasourceName, "object_name_prefix", "my-test-object"),

				resource.TestCheckResourceAttr(datasourceName, "preauthenticated_requests.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "preauthenticated_requests.0.access_type", "ObjectRead"),
				resource.TestCheckResourceAttr(datasourceName, "preauthenticated_requests.0.bucket_listing_action", ""),
				resource.TestCheckResourceAttrSet(datasourceName, "preauthenticated_requests.0.id"),
				resource.TestCheckResourceAttr(datasourceName, "preauthenticated_requests.0.name", "-tf-par"),
				resource.TestCheckResourceAttr(datasourceName, "preauthenticated_requests.0.object", "my-test-object-1"),
				resource.TestCheckResourceAttrSet(datasourceName, "preauthenticated_requests.0.time_created"),
				resource.TestCheckResourceAttr(datasourceName, "preauthenticated_requests.0.time_expires", expirationTimeForPar.String()),
			),
		},
		// verify singular datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_objectstorage_preauthrequest", "test_preauthenticated_request", Required, Create, preauthenticatedRequestSingularDataSourceRepresentation) +
				compartmentIdVariableStr + PreauthenticatedRequestResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_objectstorage_preauthrequest", "test_preauthenticated_request", Optional, Update, preauthenticatedRequestRepresentation),

			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(singularDatasourceName, "bucket", testPreAuthBucketName),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "namespace"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "par_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "access_type", "ObjectRead"),
				resource.TestCheckResourceAttr(singularDatasourceName, "bucket_listing_action", ""),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "name", "-tf-par"),
				resource.TestCheckResourceAttr(singularDatasourceName, "object", "my-test-object-1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttr(singularDatasourceName, "time_expires", expirationTimeForPar.String()),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + PreauthenticatedRequestResourceConfig,
		},
		//verify resource import
		{
			Config:            config,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"access_uri",
				"time_expires",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckObjectStoragePreauthenticatedRequestDestroy(s *terraform.State) error {
	noResourceFound := true
	client := TestAccProvider.Meta().(*OracleClients).objectStorageClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_objectstorage_preauthrequest" {
			noResourceFound = false
			request := oci_object_storage.GetPreauthenticatedRequestRequest{}

			if value, ok := rs.Primary.Attributes["bucket"]; ok {
				request.BucketName = &value
			}

			if value, ok := rs.Primary.Attributes["namespace"]; ok {
				request.NamespaceName = &value
			}

			bucket, namespace, parId, er := parsePreauthenticatedRequestCompositeId(rs.Primary.ID)
			if er == nil {
				request.BucketName = &bucket
				request.NamespaceName = &namespace
				request.ParId = &parId
			} else {
				log.Printf("[WARN] Get() unable to parse current ID: %s", rs.Primary.ID)
			}

			request.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "object_storage")

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

func init() {
	if DependencyGraph == nil {
		InitDependencyGraph()
	}
	if !InSweeperExcludeList("ObjectStoragePreauthenticatedRequest") {
		resource.AddTestSweepers("ObjectStoragePreauthenticatedRequest", &resource.Sweeper{
			Name:         "ObjectStoragePreauthenticatedRequest",
			Dependencies: DependencyGraph["preauthenticatedRequest"],
			F:            sweepObjectStoragePreauthenticatedRequestResource,
		})
	}
}

func sweepObjectStoragePreauthenticatedRequestResource(compartment string) error {
	objectStorageClient := GetTestClients(&schema.ResourceData{}).objectStorageClient()
	preauthenticatedRequestIds, err := getPreauthenticatedRequestIds(compartment)
	if err != nil {
		return err
	}
	for _, preauthenticatedRequestId := range preauthenticatedRequestIds {
		if ok := SweeperDefaultResourceId[preauthenticatedRequestId]; !ok {
			deletePreauthenticatedRequestRequest := oci_object_storage.DeletePreauthenticatedRequestRequest{}

			deletePreauthenticatedRequestRequest.ParId = &preauthenticatedRequestId

			deletePreauthenticatedRequestRequest.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "object_storage")
			_, error := objectStorageClient.DeletePreauthenticatedRequest(context.Background(), deletePreauthenticatedRequestRequest)
			if error != nil {
				fmt.Printf("Error deleting PreauthenticatedRequest %s %s, It is possible that the resource is already deleted. Please verify manually \n", preauthenticatedRequestId, error)
				continue
			}
		}
	}
	return nil
}

func getPreauthenticatedRequestIds(compartment string) ([]string, error) {
	ids := GetResourceIdsToSweep(compartment, "PreauthenticatedRequestId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	objectStorageClient := GetTestClients(&schema.ResourceData{}).objectStorageClient()

	listPreauthenticatedRequestsRequest := oci_object_storage.ListPreauthenticatedRequestsRequest{}

	buckets, error := getBucketIds(compartment)
	if error != nil {
		return resourceIds, fmt.Errorf("Error getting bucket required for PreauthenticatedRequest resource requests \n")
	}
	for _, bucket := range buckets {
		listPreauthenticatedRequestsRequest.BucketName = &bucket

		namespaces, error := getNamespaces(compartment)
		if error != nil {
			return resourceIds, fmt.Errorf("Error getting namespace required for PreauthenticatedRequest resource requests \n")
		}
		for _, namespace := range namespaces {
			listPreauthenticatedRequestsRequest.NamespaceName = &namespace

			listPreauthenticatedRequestsResponse, err := objectStorageClient.ListPreauthenticatedRequests(context.Background(), listPreauthenticatedRequestsRequest)

			if err != nil {
				return resourceIds, fmt.Errorf("Error getting PreauthenticatedRequest list for compartment id : %s , %s \n", compartmentId, err)
			}
			for _, preauthenticatedRequest := range listPreauthenticatedRequestsResponse.Items {
				id := *preauthenticatedRequest.Id
				resourceIds = append(resourceIds, id)
				AddResourceIdToSweeperResourceIdMap(compartmentId, "PreauthenticatedRequestId", id)
			}

		}
	}
	return resourceIds, nil
}
