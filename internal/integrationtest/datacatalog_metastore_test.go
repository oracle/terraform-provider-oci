// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/oracle/oci-go-sdk/v65/common"
	oci_datacatalog "github.com/oracle/oci-go-sdk/v65/datacatalog"
	"github.com/oracle/oci-go-sdk/v65/objectstorage"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DatacatalogMetastoreRequiredOnlyResource = DatacatalogMetastoreResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_datacatalog_metastore", "test_metastore", acctest.Required, acctest.Create, DatacatalogMetastoreRepresentation)

	DatacatalogMetastoreResourceConfig = DatacatalogMetastoreResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_datacatalog_metastore", "test_metastore", acctest.Optional, acctest.Update, DatacatalogMetastoreRepresentation)

	DatacatalogMetastoreSingularDataSourceRepresentation = map[string]interface{}{
		"metastore_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_datacatalog_metastore.test_metastore.id}`},
	}

	DatacatalogMetastoreDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: DatacatalogMetastoreDataSourceFilterRepresentation}}
	DatacatalogMetastoreDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_datacatalog_metastore.test_metastore.id}`}},
	}

	//Changes for retrieving the ObjectStorageURIs in the runtime
	//Representation map to Create ObjectStorage bucket
	bucketRepresentationMetastore = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"name":           acctest.Representation{RepType: acctest.Required, Create: `tf_metastore_objectstoragebucket`},
		"namespace":      acctest.Representation{RepType: acctest.Required, Create: `${data.oci_objectstorage_namespace.test_namespace.namespace}`},
	}

	//Representation map to Create directories inside the ObjectStorage bucket above
	objectRepresentationDefault = map[string]interface{}{
		"bucket":    acctest.Representation{RepType: acctest.Required, Create: `${oci_objectstorage_bucket.test_bucket.name}`},
		"namespace": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_objectstorage_namespace.test_namespace.namespace}`},
		"object":    acctest.Representation{RepType: acctest.Required, Create: "Default/"},
	}
	objectRepresentationExternal = map[string]interface{}{
		"bucket":    acctest.Representation{RepType: acctest.Required, Create: `${oci_objectstorage_bucket.test_bucket.name}`},
		"namespace": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_objectstorage_namespace.test_namespace.namespace}`},
		"object":    acctest.Representation{RepType: acctest.Required, Create: "External/"},
	}

	//Object storage values required to Create URI i.e bucket name, namespace
	objectstoragenamespace = "${data.oci_objectstorage_namespace.test_namespace.namespace}"
	objectstoragebucket    = "tf_metastore_objectstoragebucket"

	//Trying to Create the URI(oci://bucket@namespace/sub-dir/) from objectstorage objects outputs
	defaultExternalTableLocationvar = "oci://" + objectstoragebucket + "@" + objectstoragenamespace + "/" + "External"
	defaultManagedTableLocationvar  = "oci://" + objectstoragebucket + "@" + objectstoragenamespace + "/" + "Default"

	DatacatalogMetastoreRepresentation = map[string]interface{}{
		"compartment_id":                  acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"default_external_table_location": acctest.Representation{RepType: acctest.Required, Create: defaultExternalTableLocationvar},
		"default_managed_table_location":  acctest.Representation{RepType: acctest.Required, Create: defaultManagedTableLocationvar},
		"defined_tags":                    acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":                    acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":                   acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"lifecycle":                       acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreMetastoreDefinedTagsnSystemTagsChangesRepresentation},
	}
	ignoreMetastoreDefinedTagsnSystemTagsChangesRepresentation = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`, `system_tags`}},
	}

	//Changes made for Create dependency resources
	DatacatalogMetastoreResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_objectstorage_bucket", "test_bucket", acctest.Required, acctest.Create, bucketRepresentationMetastore) +
		acctest.GenerateDataSourceFromRepresentationMap("oci_objectstorage_namespace", "test_namespace", acctest.Optional, acctest.Create, ObjectStorageObjectStorageNamespaceSingularDataSourceRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_objectstorage_object", "object1", acctest.Required, acctest.Create, objectRepresentationDefault) +
		acctest.GenerateResourceFromRepresentationMap("oci_objectstorage_object", "object2", acctest.Required, acctest.Create, objectRepresentationExternal) +
		DefinedTagsDependencies
)

// issue-routing-tag: datacatalog/default
func TestDatacatalogMetastoreResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatacatalogMetastoreResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	//Create the ObjectStorageURIs
	objectstoragenamespace = getObjectStorageNamespace(compartmentId)
	defaultExternalTableLocation := "oci://" + objectstoragebucket + "@" + objectstoragenamespace + "/" + "External"
	defaultManagedTableLocation := "oci://" + objectstoragebucket + "@" + objectstoragenamespace + "/" + "Default"

	resourceName := "oci_datacatalog_metastore.test_metastore"
	datasourceName := "data.oci_datacatalog_metastores.test_metastores"
	singularDatasourceName := "data.oci_datacatalog_metastore.test_metastore"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DatacatalogMetastoreResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_datacatalog_metastore", "test_metastore", acctest.Optional, acctest.Create, DatacatalogMetastoreRepresentation), "datacatalog", "metastore", t)

	acctest.ResourceTest(t, testAccCheckDatacatalogMetastoreDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DatacatalogMetastoreResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datacatalog_metastore", "test_metastore", acctest.Required, acctest.Create, DatacatalogMetastoreRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "default_external_table_location", defaultExternalTableLocation),
				resource.TestCheckResourceAttr(resourceName, "default_managed_table_location", defaultManagedTableLocation),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DatacatalogMetastoreResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DatacatalogMetastoreResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datacatalog_metastore", "test_metastore", acctest.Optional, acctest.Create, DatacatalogMetastoreRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "default_external_table_location", defaultExternalTableLocation),
				resource.TestCheckResourceAttr(resourceName, "default_managed_table_location", defaultManagedTableLocation),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),

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

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + DatacatalogMetastoreResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datacatalog_metastore", "test_metastore", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(DatacatalogMetastoreRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "default_external_table_location", defaultExternalTableLocation),
				resource.TestCheckResourceAttr(resourceName, "default_managed_table_location", defaultManagedTableLocation),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + DatacatalogMetastoreResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datacatalog_metastore", "test_metastore", acctest.Optional, acctest.Update, DatacatalogMetastoreRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "default_external_table_location", defaultExternalTableLocation),
				resource.TestCheckResourceAttr(resourceName, "default_managed_table_location", defaultManagedTableLocation),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_datacatalog_metastores", "test_metastores", acctest.Optional, acctest.Update, DatacatalogMetastoreDataSourceRepresentation) +
				compartmentIdVariableStr + DatacatalogMetastoreResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datacatalog_metastore", "test_metastore", acctest.Optional, acctest.Update, DatacatalogMetastoreRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "metastores.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "metastores.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "metastores.0.display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "metastores.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "metastores.0.id"),
				resource.TestCheckResourceAttr(datasourceName, "metastores.0.locks.#", "0"),
				resource.TestCheckResourceAttrSet(datasourceName, "metastores.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "metastores.0.time_created"),
				resource.TestCheckResourceAttrSet(datasourceName, "metastores.0.time_updated"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_datacatalog_metastore", "test_metastore", acctest.Required, acctest.Create, DatacatalogMetastoreSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DatacatalogMetastoreResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "metastore_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "default_external_table_location", defaultExternalTableLocation),
				resource.TestCheckResourceAttr(singularDatasourceName, "default_managed_table_location", defaultManagedTableLocation),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "locks.#", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "system_tags.%", "0"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + DatacatalogMetastoreRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckDatacatalogMetastoreDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DataCatalogClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_datacatalog_metastore" {
			noResourceFound = false
			request := oci_datacatalog.GetMetastoreRequest{}

			tmp := rs.Primary.ID
			request.MetastoreId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "datacatalog")

			response, err := client.GetMetastore(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_datacatalog.LifecycleStateDeleted): true,
				}
				if _, ok := deletedLifecycleStates[string(response.LifecycleState)]; !ok {
					//resource lifecycle state is not in expected deleted lifecycle states.
					return fmt.Errorf("resource lifecycle state: %s is not in expected deleted lifecycle states", response.LifecycleState)
				}
				//resource lifecycle state is in expected deleted lifecycle states. continue with next one.
				continue
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
	if !acctest.InSweeperExcludeList("DatacatalogMetastore") {
		resource.AddTestSweepers("DatacatalogMetastore", &resource.Sweeper{
			Name:         "DatacatalogMetastore",
			Dependencies: acctest.DependencyGraph["metastore"],
			F:            sweepDatacatalogMetastoreResource,
		})
	}
}

func sweepDatacatalogMetastoreResource(compartment string) error {
	dataCatalogClient := acctest.GetTestClients(&schema.ResourceData{}).DataCatalogClient()
	metastoreIds, err := getDatacatalogMetastoreIds(compartment)
	if err != nil {
		return err
	}
	for _, metastoreId := range metastoreIds {
		if ok := acctest.SweeperDefaultResourceId[metastoreId]; !ok {
			deleteMetastoreRequest := oci_datacatalog.DeleteMetastoreRequest{}

			deleteMetastoreRequest.MetastoreId = &metastoreId

			deleteMetastoreRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "datacatalog")
			_, error := dataCatalogClient.DeleteMetastore(context.Background(), deleteMetastoreRequest)
			if error != nil {
				fmt.Printf("Error deleting Metastore %s %s, It is possible that the resource is already deleted. Please verify manually \n", metastoreId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &metastoreId, DatacatalogMetastoreSweepWaitCondition, time.Duration(3*time.Minute),
				DatacatalogMetastoreSweepResponseFetchOperation, "datacatalog", true)
		}
	}
	return nil
}

func getDatacatalogMetastoreIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "MetastoreId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	dataCatalogClient := acctest.GetTestClients(&schema.ResourceData{}).DataCatalogClient()

	listMetastoresRequest := oci_datacatalog.ListMetastoresRequest{}
	listMetastoresRequest.CompartmentId = &compartmentId
	listMetastoresRequest.LifecycleState = oci_datacatalog.ListMetastoresLifecycleStateActive
	listMetastoresResponse, err := dataCatalogClient.ListMetastores(context.Background(), listMetastoresRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting Metastore list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, metastore := range listMetastoresResponse.Items {
		id := *metastore.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "MetastoreId", id)
		//acctest.SweeperDefaultResourceId[*metastore.DefaultExternalTableLocation] = true
		//acctest.SweeperDefaultResourceId[*metastore.DefaultManagedTableLocation] = true

	}
	return resourceIds, nil
}

func DatacatalogMetastoreSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if metastoreResponse, ok := response.Response.(oci_datacatalog.GetMetastoreResponse); ok {
		return metastoreResponse.LifecycleState != oci_datacatalog.LifecycleStateDeleted
	}
	return false
}

func DatacatalogMetastoreSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DataCatalogClient().GetMetastore(context.Background(), oci_datacatalog.GetMetastoreRequest{
		MetastoreId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}

// Function to get ObjectStorage Namespace
func getObjectStorageNamespace(compartmentId string) string {
	compartment := compartmentId
	objectStorageClient := acctest.GetTestClients(&schema.ResourceData{}).ObjectStorageClient()
	ctx := context.Background()
	request := objectstorage.GetNamespaceRequest{CompartmentId: common.String(compartment)}
	r, err := objectStorageClient.GetNamespace(ctx, request)
	if err != nil {
		err := fmt.Errorf("Error getting namespace : %v", err)
		fmt.Println(err.Error())
	}
	return *r.Value
}
