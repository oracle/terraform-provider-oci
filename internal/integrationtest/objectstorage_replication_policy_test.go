// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"testing"

	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	tf_client "github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/resourcediscovery"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v56/common"
	oci_object_storage "github.com/oracle/oci-go-sdk/v56/objectstorage"

	tf_objectstorage "github.com/terraform-providers/terraform-provider-oci/internal/service/objectstorage"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	replicationBucketName           = testBucketName + "_replication"
	ReplicationPolicyResourceConfig = ReplicationPolicyResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_objectstorage_replication_policy", "test_replication_policy", acctest.Optional, acctest.Update, replicationPolicyRepresentation)

	replicationPolicySingularDataSourceRepresentation = map[string]interface{}{
		"bucket":         acctest.Representation{RepType: acctest.Required, Create: `${oci_objectstorage_bucket.test_bucket.name}`},
		"namespace":      acctest.Representation{RepType: acctest.Required, Create: `${oci_objectstorage_bucket.test_bucket.namespace}`},
		"replication_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_objectstorage_replication_policies.test_replication_policies.replication_policies.0.id}`},
	}

	replicationPolicyDataSourceRepresentation = map[string]interface{}{
		"bucket":    acctest.Representation{RepType: acctest.Required, Create: `${oci_objectstorage_bucket.test_bucket.name}`},
		"namespace": acctest.Representation{RepType: acctest.Required, Create: `${oci_objectstorage_bucket.test_bucket.namespace}`},
	}

	replicationPolicyRepresentation = map[string]interface{}{
		"bucket":                  acctest.Representation{RepType: acctest.Required, Create: `${oci_objectstorage_bucket.test_bucket.name}`},
		"destination_bucket_name": acctest.Representation{RepType: acctest.Required, Create: `${oci_objectstorage_bucket.test_bucket_replication.name}`},
		"destination_region_name": acctest.Representation{RepType: acctest.Required, Create: `${var.region}`},
		"name":                    acctest.Representation{RepType: acctest.Required, Create: `mypolicy`},
		"namespace":               acctest.Representation{RepType: acctest.Required, Create: `${oci_objectstorage_bucket.test_bucket.namespace}`},
	}

	ReplicationPolicyResourceDependencies = acctest.GenerateDataSourceFromRepresentationMap("oci_identity_regions", "test_regions", acctest.Required, acctest.Create, regionDataSourceRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_objectstorage_bucket", "test_bucket", acctest.Required, acctest.Create, bucketRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_objectstorage_bucket", "test_bucket_replication", acctest.Required, acctest.Create,
			acctest.RepresentationCopyWithNewProperties(bucketRepresentation, map[string]interface{}{
				"name": acctest.Representation{RepType: acctest.Required, Create: replicationBucketName},
			})) + acctest.GenerateDataSourceFromRepresentationMap("oci_objectstorage_namespace", "test_namespace", acctest.Required, acctest.Create, namespaceSingularDataSourceRepresentation)
)

// issue-routing-tag: object_storage/default
func TestObjectStorageReplicationPolicyResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestObjectStorageReplicationPolicyResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_objectstorage_replication_policy.test_replication_policy"
	datasourceName := "data.oci_objectstorage_replication_policies.test_replication_policies"
	singularDatasourceName := "data.oci_objectstorage_replication_policy.test_replication_policy"

	var resId string
	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the Create step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+ReplicationPolicyResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_objectstorage_replication_policy", "test_replication_policy", acctest.Required, acctest.Create, replicationPolicyRepresentation), "objectstorage", "replicationPolicy", t)

	acctest.ResourceTest(t, testAccCheckObjectStorageReplicationPolicyDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + ReplicationPolicyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_objectstorage_replication_policy", "test_replication_policy", acctest.Required, acctest.Create, replicationPolicyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "bucket", testBucketName),
				resource.TestCheckResourceAttrSet(resourceName, "destination_bucket_name"),
				resource.TestCheckResourceAttrSet(resourceName, "destination_region_name"),
				resource.TestCheckResourceAttr(resourceName, "name", "mypolicy"),
				resource.TestCheckResourceAttrSet(resourceName, "namespace"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_objectstorage_replication_policies", "test_replication_policies", acctest.Optional, acctest.Update, replicationPolicyDataSourceRepresentation) +
				compartmentIdVariableStr + ReplicationPolicyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_objectstorage_replication_policy", "test_replication_policy", acctest.Optional, acctest.Update, replicationPolicyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "bucket", testBucketName),

				resource.TestCheckResourceAttr(datasourceName, "replication_policies.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "replication_policies.0.destination_bucket_name"),
				resource.TestCheckResourceAttrSet(datasourceName, "replication_policies.0.destination_region_name"),
				resource.TestCheckResourceAttrSet(datasourceName, "replication_policies.0.id"),
				resource.TestCheckResourceAttr(datasourceName, "replication_policies.0.name", "mypolicy"),
				resource.TestCheckResourceAttrSet(datasourceName, "replication_policies.0.status"),
				resource.TestCheckResourceAttrSet(datasourceName, "replication_policies.0.status_message"),
				resource.TestCheckResourceAttrSet(datasourceName, "replication_policies.0.time_created"),
			),
		},

		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_objectstorage_replication_policies", "test_replication_policies", acctest.Optional, acctest.Update, replicationPolicyDataSourceRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_objectstorage_replication_policy", "test_replication_policy", acctest.Required, acctest.Create, replicationPolicySingularDataSourceRepresentation) +
				compartmentIdVariableStr + ReplicationPolicyResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(singularDatasourceName, "bucket", testBucketName),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "replication_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "name", "mypolicy"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "status"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "status_message"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + ReplicationPolicyResourceConfig,
		},
		// verify resource import
		{
			Config:            config,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"time_last_sync",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckObjectStorageReplicationPolicyDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).ObjectStorageClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_objectstorage_replication_policy" {
			noResourceFound = false
			request := oci_object_storage.GetReplicationPolicyRequest{}

			bucket, namespace, replicationId, err := tf_objectstorage.ParseReplicationPolicyCompositeId(rs.Primary.ID)
			if err == nil {
				request.BucketName = &bucket
				request.NamespaceName = &namespace
				request.ReplicationId = &replicationId
			} else {
				log.Printf("[WARN] Get() unable to parse current ID: %s", rs.Primary.ID)
			}

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "object_storage")

			_, err = client.GetReplicationPolicy(context.Background(), request)

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
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("ObjectStorageReplicationPolicy") {
		resource.AddTestSweepers("ObjectStorageReplicationPolicy", &resource.Sweeper{
			Name:         "ObjectStorageReplicationPolicy",
			Dependencies: acctest.DependencyGraph["replicationPolicy"],
			F:            sweepObjectStorageReplicationPolicyResource,
		})
	}
}

func sweepObjectStorageReplicationPolicyResource(compartment string) error {
	objectStorageClient := acctest.GetTestClients(&schema.ResourceData{}).ObjectStorageClient()
	replicationPolicyIds, err := getReplicationPolicyIds(compartment)
	if err != nil {
		return err
	}
	for _, replicationPolicyId := range replicationPolicyIds {
		if ok := acctest.SweeperDefaultResourceId[replicationPolicyId]; !ok {
			deleteReplicationPolicyRequest := oci_object_storage.DeleteReplicationPolicyRequest{}

			deleteReplicationPolicyRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "object_storage")
			_, error := objectStorageClient.DeleteReplicationPolicy(context.Background(), deleteReplicationPolicyRequest)
			if error != nil {
				fmt.Printf("Error deleting ReplicationPolicy %s %s, It is possible that the resource is already deleted. Please verify manually \n", replicationPolicyId, error)
				continue
			}
		}
	}
	return nil
}

func getReplicationPolicyIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "ReplicationPolicyId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	objectStorageClient := acctest.GetTestClients(&schema.ResourceData{}).ObjectStorageClient()

	listReplicationPoliciesRequest := oci_object_storage.ListReplicationPoliciesRequest{}

	buckets, error := getBucketIds(compartment)
	if error != nil {
		return resourceIds, fmt.Errorf("Error getting bucket required for ReplicationPolicy resource requests \n")
	}
	for _, bucket := range buckets {
		listReplicationPoliciesRequest.BucketName = &bucket

		namespaces, error := getNamespaces(compartment)
		if error != nil {
			return resourceIds, fmt.Errorf("Error getting namespace required for ReplicationPolicy resource requests \n")
		}
		for _, namespace := range namespaces {
			listReplicationPoliciesRequest.NamespaceName = &namespace

			listReplicationPoliciesResponse, err := objectStorageClient.ListReplicationPolicies(context.Background(), listReplicationPoliciesRequest)

			if err != nil {
				return resourceIds, fmt.Errorf("Error getting ReplicationPolicy list for compartment id : %s , %s \n", compartmentId, err)
			}
			for _, replicationPolicy := range listReplicationPoliciesResponse.Items {
				id := *replicationPolicy.Id
				resourceIds = append(resourceIds, id)
				acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "ReplicationPolicyId", id)
			}

		}
	}
	return resourceIds, nil
}
