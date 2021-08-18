// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v46/common"
	oci_email "github.com/oracle/oci-go-sdk/v46/email"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	DkimRequiredOnlyResource = DkimResourceDependencies +
		generateResourceFromRepresentationMap("oci_email_dkim", "test_dkim", Required, Create, dkimRepresentation)

	DkimResourceConfig = DkimResourceDependencies +
		generateResourceFromRepresentationMap("oci_email_dkim", "test_dkim", Optional, Update, dkimRepresentation)

	dkimSingularDataSourceRepresentation = map[string]interface{}{
		"dkim_id": Representation{repType: Required, create: `${oci_email_dkim.test_dkim.id}`},
	}

	dkimDataSourceRepresentation = map[string]interface{}{
		"email_domain_id": Representation{repType: Required, create: `${oci_email_email_domain.test_email_domain.id}`},
		"id":              Representation{repType: Optional, create: `${oci_email_dkim.test_dkim.id}`},
		"name":            Representation{repType: Optional, create: `testselector1`},
		"state":           Representation{repType: Optional, create: `NEEDS_ATTENTION`},
		"filter":          RepresentationGroup{Required, dkimDataSourceFilterRepresentation}}
	dkimDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_email_dkim.test_dkim.id}`}},
	}

	dkimRepresentation = map[string]interface{}{
		"email_domain_id": Representation{repType: Required, create: `${oci_email_email_domain.test_email_domain.id}`},
		"defined_tags":    Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`},
		"description":     Representation{repType: Optional, create: `description`, update: `description`},
		"freeform_tags":   Representation{repType: Optional, create: map[string]string{"Department": "Finance"}, update: map[string]string{"Department": "Finance"}},
		"name":            Representation{repType: Optional, create: `testselector1`},
	}

	DkimResourceDependencies = generateResourceFromRepresentationMap("oci_email_email_domain", "test_email_domain", Required, Create, emailDomainRepresentation) +
		DefinedTagsDependencies
)

func TestEmailDkimResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestEmailDkimResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_email_dkim.test_dkim"
	datasourceName := "data.oci_email_dkims.test_dkims"
	singularDatasourceName := "data.oci_email_dkim.test_dkim"

	var resId, resId2 string
	// Save TF content to create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	saveConfigContent(config+compartmentIdVariableStr+DkimResourceDependencies+
		generateResourceFromRepresentationMap("oci_email_dkim", "test_dkim", Optional, Create, dkimRepresentation), "email", "dkim", t)

	ResourceTest(t, testAccCheckEmailDkimDestroy, []resource.TestStep{
		//verify create
		{
			Config: config + compartmentIdVariableStr + DkimResourceDependencies +
				generateResourceFromRepresentationMap("oci_email_dkim", "test_dkim", Required, Create, dkimRepresentation),
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttrSet(resourceName, "email_domain_id"),

				func(s *terraform.State) (err error) {
					resId, err = fromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		// delete before next create
		{
			Config: config + compartmentIdVariableStr + DkimResourceDependencies,
		},
		// verify create with optionals
		{
			Config: config + compartmentIdVariableStr + DkimResourceDependencies +
				generateResourceFromRepresentationMap("oci_email_dkim", "test_dkim", Optional, Create, dkimRepresentation),
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttrSet(resourceName, "email_domain_id"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "name", "testselector1"),

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

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + DkimResourceDependencies +
				generateResourceFromRepresentationMap("oci_email_dkim", "test_dkim", Optional, Update, dkimRepresentation),
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttrSet(resourceName, "email_domain_id"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "name", "testselector1"),

				func(s *terraform.State) (err error) {
					resId2, err = fromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		//verify datasource
		{
			Config: config +
				generateDataSourceFromRepresentationMap("oci_email_dkims", "test_dkims", Optional, Update, dkimDataSourceRepresentation) +
				compartmentIdVariableStr + DkimResourceDependencies +
				generateResourceFromRepresentationMap("oci_email_dkim", "test_dkim", Optional, Update, dkimRepresentation),
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttrSet(datasourceName, "email_domain_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttr(datasourceName, "name", "testselector1"),
				resource.TestCheckResourceAttr(datasourceName, "state", "NEEDS_ATTENTION"),

				resource.TestCheckResourceAttr(datasourceName, "dkim_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "dkim_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				generateDataSourceFromRepresentationMap("oci_email_dkim", "test_dkim", Required, Create, dkimSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DkimResourceConfig,
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "dkim_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "cname_record_value"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "dns_subdomain_name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "name", "testselector1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "txt_record_value"),
			),
		},
		//remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + DkimResourceConfig,
		},
		//verify resource import
		{
			Config:                  config,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckEmailDkimDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).emailClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_email_dkim" {
			noResourceFound = false
			request := oci_email.GetDkimRequest{}

			tmp := rs.Primary.ID
			request.DkimId = &tmp

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "email")

			response, err := client.GetDkim(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_email.DkimLifecycleStateDeleted): true,
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
	if DependencyGraph == nil {
		initDependencyGraph()
	}
	if !inSweeperExcludeList("EmailDkim") {
		resource.AddTestSweepers("EmailDkim", &resource.Sweeper{
			Name:         "EmailDkim",
			Dependencies: DependencyGraph["dkim"],
			F:            sweepEmailDkimResource,
		})
	}
}

func sweepEmailDkimResource(compartment string) error {
	emailClient := GetTestClients(&schema.ResourceData{}).emailClient()
	dkimIds, err := getDkimIds(compartment)
	if err != nil {
		return err
	}
	for _, dkimId := range dkimIds {
		if ok := SweeperDefaultResourceId[dkimId]; !ok {
			deleteDkimRequest := oci_email.DeleteDkimRequest{}

			deleteDkimRequest.DkimId = &dkimId

			deleteDkimRequest.RequestMetadata.RetryPolicy = getRetryPolicy(true, "email")
			_, error := emailClient.DeleteDkim(context.Background(), deleteDkimRequest)
			if error != nil {
				fmt.Printf("Error deleting Dkim %s %s, It is possible that the resource is already deleted. Please verify manually \n", dkimId, error)
				continue
			}
			waitTillCondition(testAccProvider, &dkimId, dkimSweepWaitCondition, time.Duration(3*time.Minute),
				dkimSweepResponseFetchOperation, "email", true)
		}
	}
	return nil
}

func getDkimIds(compartment string) ([]string, error) {
	ids := getResourceIdsToSweep(compartment, "DkimId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	emailClient := GetTestClients(&schema.ResourceData{}).emailClient()

	listDkimsRequest := oci_email.ListDkimsRequest{}
	listDkimsRequest.Id = &compartmentId

	emailDomainIds, error := getEmailDomainIds(compartment)
	if error != nil {
		return resourceIds, fmt.Errorf("Error getting emailDomainId required for Dkim resource requests \n")
	}
	for _, emailDomainId := range emailDomainIds {
		listDkimsRequest.EmailDomainId = &emailDomainId

		listDkimsRequest.LifecycleState = oci_email.DkimLifecycleStateDeleted
		listDkimsResponse, err := emailClient.ListDkims(context.Background(), listDkimsRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting Dkim list for compartment id : %s , %s \n", compartmentId, err)
		}
		for _, dkim := range listDkimsResponse.Items {
			id := *dkim.Id
			resourceIds = append(resourceIds, id)
			addResourceIdToSweeperResourceIdMap(compartmentId, "DkimId", id)
		}

	}
	return resourceIds, nil
}

func dkimSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if dkimResponse, ok := response.Response.(oci_email.GetDkimResponse); ok {
		return dkimResponse.LifecycleState != oci_email.DkimLifecycleStateDeleted
	}
	return false
}

func dkimSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.emailClient().GetDkim(context.Background(), oci_email.GetDkimRequest{
		DkimId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
