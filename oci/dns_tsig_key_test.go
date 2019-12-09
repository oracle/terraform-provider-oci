// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package oci

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/oracle/oci-go-sdk/common"
	oci_dns "github.com/oracle/oci-go-sdk/dns"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	TsigKeyRequiredOnlyResource = TsigKeyResourceDependencies +
		generateResourceFromRepresentationMap("oci_dns_tsig_key", "test_tsig_key", Required, Create, tsigKeyRepresentation)

	TsigKeyResourceConfig = TsigKeyResourceDependencies +
		generateResourceFromRepresentationMap("oci_dns_tsig_key", "test_tsig_key", Optional, Update, tsigKeyRepresentation)

	tsigKeySingularDataSourceRepresentation = map[string]interface{}{
		"tsig_key_id": Representation{repType: Required, create: `${oci_dns_tsig_key.test_tsig_key.id}`},
	}

	tsigKeyName                     = randomString(15, charsetWithoutDigits)
	tsigKeyDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"id":             Representation{repType: Optional, create: `${oci_dns_tsig_key.test_tsig_key.id}`},
		"name":           Representation{repType: Optional, create: tsigKeyName},
		"state":          Representation{repType: Optional, create: `ACTIVE`},
		"filter":         RepresentationGroup{Required, tsigKeyDataSourceFilterRepresentation}}
	tsigKeyDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_dns_tsig_key.test_tsig_key.id}`}},
	}

	tsigKeyRepresentation = map[string]interface{}{
		"algorithm":      Representation{repType: Required, create: `hmac-sha1`},
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"name":           Representation{repType: Required, create: tsigKeyName},
		"secret":         Representation{repType: Required, create: `c2VjcmV0`},
		"defined_tags":   Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":  Representation{repType: Optional, create: map[string]string{"freeformTags": "freeformTags"}, update: map[string]string{"freeformTags2": "freeformTags2"}},
	}

	TsigKeyResourceDependencies = DefinedTagsDependencies
)

func TestDnsTsigKeyResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDnsTsigKeyResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := getEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_dns_tsig_key.test_tsig_key"
	datasourceName := "data.oci_dns_tsig_keys.test_tsig_keys"
	singularDatasourceName := "data.oci_dns_tsig_key.test_tsig_key"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckDnsTsigKeyDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + TsigKeyResourceDependencies +
					generateResourceFromRepresentationMap("oci_dns_tsig_key", "test_tsig_key", Required, Create, tsigKeyRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "algorithm", "hmac-sha1"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "name", tsigKeyName),
					resource.TestCheckResourceAttr(resourceName, "secret", "c2VjcmV0"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + TsigKeyResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + TsigKeyResourceDependencies +
					generateResourceFromRepresentationMap("oci_dns_tsig_key", "test_tsig_key", Optional, Create, tsigKeyRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "algorithm", "hmac-sha1"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "name", tsigKeyName),
					resource.TestCheckResourceAttr(resourceName, "secret", "c2VjcmV0"),

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

			// verify update to the compartment (the compartment will be switched back in the next step)
			{
				Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + TsigKeyResourceDependencies +
					generateResourceFromRepresentationMap("oci_dns_tsig_key", "test_tsig_key", Optional, Create,
						representationCopyWithNewProperties(tsigKeyRepresentation, map[string]interface{}{
							"compartment_id": Representation{repType: Required, create: `${var.compartment_id_for_update}`},
						})),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "algorithm", "hmac-sha1"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "name", tsigKeyName),
					resource.TestCheckResourceAttr(resourceName, "secret", "c2VjcmV0"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
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
					generateResourceFromRepresentationMap("oci_dns_tsig_key", "test_tsig_key", Optional, Update, tsigKeyRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "algorithm", "hmac-sha1"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "name", tsigKeyName),
					resource.TestCheckResourceAttr(resourceName, "secret", "c2VjcmV0"),

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
				Config: config +
					generateDataSourceFromRepresentationMap("oci_dns_tsig_keys", "test_tsig_keys", Optional, Update, tsigKeyDataSourceRepresentation) +
					compartmentIdVariableStr + TsigKeyResourceDependencies +
					generateResourceFromRepresentationMap("oci_dns_tsig_key", "test_tsig_key", Optional, Update, tsigKeyRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(datasourceName, "id"),
					resource.TestCheckResourceAttr(datasourceName, "name", tsigKeyName),
					resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

					resource.TestCheckResourceAttr(datasourceName, "tsig_keys.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "tsig_keys.0.algorithm", "hmac-sha1"),
					resource.TestCheckResourceAttr(datasourceName, "tsig_keys.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "tsig_keys.0.defined_tags.%", "1"),
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
					generateDataSourceFromRepresentationMap("oci_dns_tsig_key", "test_tsig_key", Required, Create, tsigKeySingularDataSourceRepresentation) +
					compartmentIdVariableStr + TsigKeyResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "tsig_key_id"),

					resource.TestCheckResourceAttr(singularDatasourceName, "algorithm", "hmac-sha1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
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
		},
	})
}

func testAccCheckDnsTsigKeyDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).dnsClient
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_dns_tsig_key" {
			noResourceFound = false
			request := oci_dns.GetTsigKeyRequest{}

			tmp := rs.Primary.ID
			request.TsigKeyId = &tmp

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "dns")

			_, err := client.GetTsigKey(context.Background(), request)

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
	if !inSweeperExcludeList("DnsTsigKey") {
		resource.AddTestSweepers("DnsTsigKey", &resource.Sweeper{
			Name:         "DnsTsigKey",
			Dependencies: DependencyGraph["tsigKey"],
			F:            sweepDnsTsigKeyResource,
		})
	}
}

func sweepDnsTsigKeyResource(compartment string) error {
	dnsClient := GetTestClients(&schema.ResourceData{}).dnsClient
	tsigKeyIds, err := getTsigKeyIds(compartment)
	if err != nil {
		return err
	}
	for _, tsigKeyId := range tsigKeyIds {
		if ok := SweeperDefaultResourceId[tsigKeyId]; !ok {
			deleteTsigKeyRequest := oci_dns.DeleteTsigKeyRequest{}

			deleteTsigKeyRequest.TsigKeyId = &tsigKeyId

			deleteTsigKeyRequest.RequestMetadata.RetryPolicy = getRetryPolicy(true, "dns")
			_, error := dnsClient.DeleteTsigKey(context.Background(), deleteTsigKeyRequest)
			if error != nil {
				fmt.Printf("Error deleting TsigKey %s %s, It is possible that the resource is already deleted. Please verify manually \n", tsigKeyId, error)
				continue
			}
		}
	}
	return nil
}

func getTsigKeyIds(compartment string) ([]string, error) {
	ids := getResourceIdsToSweep(compartment, "TsigKeyId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	dnsClient := GetTestClients(&schema.ResourceData{}).dnsClient

	listTsigKeysRequest := oci_dns.ListTsigKeysRequest{}
	listTsigKeysRequest.CompartmentId = &compartmentId
	listTsigKeysResponse, err := dnsClient.ListTsigKeys(context.Background(), listTsigKeysRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting TsigKey list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, tsigKey := range listTsigKeysResponse.Items {
		id := *tsigKey.Id
		resourceIds = append(resourceIds, id)
		addResourceIdToSweeperResourceIdMap(compartmentId, "TsigKeyId", id)
	}
	return resourceIds, nil
}
