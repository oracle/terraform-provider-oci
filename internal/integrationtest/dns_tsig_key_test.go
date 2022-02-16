// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	tf_client "github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/resourcediscovery"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v58/common"
	oci_dns "github.com/oracle/oci-go-sdk/v58/dns"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	TsigKeyRequiredOnlyResource = TsigKeyResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_dns_tsig_key", "test_tsig_key", acctest.Required, acctest.Create, tsigKeyRepresentation)

	TsigKeyResourceConfig = TsigKeyResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_dns_tsig_key", "test_tsig_key", acctest.Optional, acctest.Update, tsigKeyRepresentation)

	tsigKeySingularDataSourceRepresentation = map[string]interface{}{
		"tsig_key_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_dns_tsig_key.test_tsig_key.id}`},
	}

	tsigKeyName                     = utils.RandomString(7, utils.CharsetWithoutDigits) + "." + utils.RandomString(8, utils.CharsetWithoutDigits)
	tsigKeyDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_dns_tsig_key.test_tsig_key.id}`},
		"name":           acctest.Representation{RepType: acctest.Optional, Create: tsigKeyName},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: tsigKeyDataSourceFilterRepresentation}}
	tsigKeyDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_dns_tsig_key.test_tsig_key.id}`}},
	}

	tsigKeyRepresentation = map[string]interface{}{
		"algorithm":      acctest.Representation{RepType: acctest.Required, Create: `hmac-sha1`},
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"name":           acctest.Representation{RepType: acctest.Required, Create: tsigKeyName},
		"secret":         acctest.Representation{RepType: acctest.Required, Create: `c2VjcmV0`},
		"defined_tags":   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"freeformTags": "freeformTags"}, Update: map[string]string{"freeformTags2": "freeformTags2"}},
	}

	TsigKeyResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: dns/default
func TestDnsTsigKeyResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDnsTsigKeyResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_dns_tsig_key.test_tsig_key"
	datasourceName := "data.oci_dns_tsig_keys.test_tsig_keys"
	singularDatasourceName := "data.oci_dns_tsig_key.test_tsig_key"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+TsigKeyResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_dns_tsig_key", "test_tsig_key", acctest.Optional, acctest.Create, tsigKeyRepresentation), "dns", "tsigKey", t)

	acctest.ResourceTest(t, testAccCheckDnsTsigKeyDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + TsigKeyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dns_tsig_key", "test_tsig_key", acctest.Required, acctest.Create, tsigKeyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "algorithm", "hmac-sha1"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "name", tsigKeyName),
				resource.TestCheckResourceAttr(resourceName, "secret", "c2VjcmV0"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + TsigKeyResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + TsigKeyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dns_tsig_key", "test_tsig_key", acctest.Optional, acctest.Create, tsigKeyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "algorithm", "hmac-sha1"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "name", tsigKeyName),
				resource.TestCheckResourceAttr(resourceName, "secret", "c2VjcmV0"),

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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + TsigKeyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dns_tsig_key", "test_tsig_key", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(tsigKeyRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "algorithm", "hmac-sha1"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "name", tsigKeyName),
				resource.TestCheckResourceAttr(resourceName, "secret", "c2VjcmV0"),

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
			Config: config + compartmentIdVariableStr + TsigKeyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dns_tsig_key", "test_tsig_key", acctest.Optional, acctest.Update, tsigKeyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "algorithm", "hmac-sha1"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "name", tsigKeyName),
				resource.TestCheckResourceAttr(resourceName, "secret", "c2VjcmV0"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_dns_tsig_keys", "test_tsig_keys", acctest.Optional, acctest.Update, tsigKeyDataSourceRepresentation) +
				compartmentIdVariableStr + TsigKeyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dns_tsig_key", "test_tsig_key", acctest.Optional, acctest.Update, tsigKeyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttr(datasourceName, "name", tsigKeyName),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "tsig_keys.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "tsig_keys.0.algorithm", "hmac-sha1"),
				resource.TestCheckResourceAttr(datasourceName, "tsig_keys.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "tsig_keys.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "tsig_keys.0.id"),
				resource.TestCheckResourceAttr(datasourceName, "tsig_keys.0.name", tsigKeyName),
				resource.TestCheckResourceAttrSet(datasourceName, "tsig_keys.0.self"),
				resource.TestCheckResourceAttrSet(datasourceName, "tsig_keys.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "tsig_keys.0.time_created"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_dns_tsig_key", "test_tsig_key", acctest.Required, acctest.Create, tsigKeySingularDataSourceRepresentation) +
				compartmentIdVariableStr + TsigKeyResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "tsig_key_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "algorithm", "hmac-sha1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "name", tsigKeyName),
				resource.TestCheckResourceAttr(singularDatasourceName, "secret", "c2VjcmV0"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "self"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + TsigKeyResourceConfig,
		},
		// verify resource import
		{
			Config:                  config,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckDnsTsigKeyDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DnsClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_dns_tsig_key" {
			noResourceFound = false
			request := oci_dns.GetTsigKeyRequest{}

			tmp := rs.Primary.ID
			request.TsigKeyId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "dns")

			response, err := client.GetTsigKey(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_dns.TsigKeyLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("DnsTsigKey") {
		resource.AddTestSweepers("DnsTsigKey", &resource.Sweeper{
			Name:         "DnsTsigKey",
			Dependencies: acctest.DependencyGraph["tsigKey"],
			F:            sweepDnsTsigKeyResource,
		})
	}
}

func sweepDnsTsigKeyResource(compartment string) error {
	dnsClient := acctest.GetTestClients(&schema.ResourceData{}).DnsClient()
	tsigKeyIds, err := getTsigKeyIds(compartment)
	if err != nil {
		return err
	}
	for _, tsigKeyId := range tsigKeyIds {
		if ok := acctest.SweeperDefaultResourceId[tsigKeyId]; !ok {
			deleteTsigKeyRequest := oci_dns.DeleteTsigKeyRequest{}

			deleteTsigKeyRequest.TsigKeyId = &tsigKeyId

			deleteTsigKeyRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "dns")
			_, error := dnsClient.DeleteTsigKey(context.Background(), deleteTsigKeyRequest)
			if error != nil {
				fmt.Printf("Error deleting TsigKey %s %s, It is possible that the resource is already deleted. Please verify manually \n", tsigKeyId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &tsigKeyId, tsigKeySweepWaitCondition, time.Duration(3*time.Minute),
				tsigKeySweepResponseFetchOperation, "dns", true)
		}
	}
	return nil
}

func getTsigKeyIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "TsigKeyId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	dnsClient := acctest.GetTestClients(&schema.ResourceData{}).DnsClient()

	listTsigKeysRequest := oci_dns.ListTsigKeysRequest{}
	listTsigKeysRequest.CompartmentId = &compartmentId
	listTsigKeysRequest.LifecycleState = oci_dns.TsigKeySummaryLifecycleStateActive
	listTsigKeysResponse, err := dnsClient.ListTsigKeys(context.Background(), listTsigKeysRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting TsigKey list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, tsigKey := range listTsigKeysResponse.Items {
		id := *tsigKey.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "TsigKeyId", id)
	}
	return resourceIds, nil
}

func tsigKeySweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if tsigKeyResponse, ok := response.Response.(oci_dns.GetTsigKeyResponse); ok {
		return tsigKeyResponse.LifecycleState != oci_dns.TsigKeyLifecycleStateDeleted
	}
	return false
}

func tsigKeySweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DnsClient().GetTsigKey(context.Background(), oci_dns.GetTsigKeyRequest{
		TsigKeyId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
