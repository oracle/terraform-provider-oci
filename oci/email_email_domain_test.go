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
	"github.com/oracle/oci-go-sdk/v47/common"
	oci_email "github.com/oracle/oci-go-sdk/v47/email"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	EmailDomainRequiredOnlyResource = EmailDomainResourceDependencies +
		generateResourceFromRepresentationMap("oci_email_email_domain", "test_email_domain", Required, Create, emailDomainRepresentation)

	EmailDomainResourceConfig = EmailDomainResourceDependencies +
		generateResourceFromRepresentationMap("oci_email_email_domain", "test_email_domain", Optional, Update, emailDomainRepresentation)

	emailDomainSingularDataSourceRepresentation = map[string]interface{}{
		"email_domain_id": Representation{repType: Required, create: `${oci_email_email_domain.test_email_domain.id}`},
	}

	randomDomain = randomString(8, charsetLowerCaseWithoutDigits) + ".email.us-phoenix-1.oci.oc-test.com"

	emailDomainDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"id":             Representation{repType: Optional, create: `${oci_email_email_domain.test_email_domain.id}`},
		"name":           Representation{repType: Optional, create: randomDomain},
		"state":          Representation{repType: Optional, create: `ACTIVE`},
		"filter":         RepresentationGroup{Required, emailDomainDataSourceFilterRepresentation}}
	emailDomainDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_email_email_domain.test_email_domain.id}`}},
	}

	emailDomainRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"name":           Representation{repType: Required, create: randomDomain},
		"defined_tags":   Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":    Representation{repType: Optional, create: `description`, update: `description2`},
		"freeform_tags":  Representation{repType: Optional, create: map[string]string{"Department": "Finance"}, update: map[string]string{"Department": "Accounting"}},
	}

	EmailDomainResourceDependencies = DefinedTagsDependencies
)

func TestEmailEmailDomainResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestEmailEmailDomainResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := getEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_email_email_domain.test_email_domain"
	datasourceName := "data.oci_email_email_domains.test_email_domains"
	singularDatasourceName := "data.oci_email_email_domain.test_email_domain"

	var resId, resId2 string
	// Save TF content to create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	saveConfigContent(config+compartmentIdVariableStr+EmailDomainResourceDependencies+
		generateResourceFromRepresentationMap("oci_email_email_domain", "test_email_domain", Optional, Create, emailDomainRepresentation), "email", "emailDomain", t)

	ResourceTest(t, testAccCheckEmailEmailDomainDestroy, []resource.TestStep{
		// verify create
		{
			Config: config + compartmentIdVariableStr + EmailDomainResourceDependencies +
				generateResourceFromRepresentationMap("oci_email_email_domain", "test_email_domain", Required, Create, emailDomainRepresentation),
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "name", randomDomain),

				func(s *terraform.State) (err error) {
					resId, err = fromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next create
		{
			Config: config + compartmentIdVariableStr + EmailDomainResourceDependencies,
		},
		// verify create with optionals
		{
			Config: config + compartmentIdVariableStr + EmailDomainResourceDependencies +
				generateResourceFromRepresentationMap("oci_email_email_domain", "test_email_domain", Optional, Create, emailDomainRepresentation),
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "name", randomDomain),

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

		// verify update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + EmailDomainResourceDependencies +
				generateResourceFromRepresentationMap("oci_email_email_domain", "test_email_domain", Optional, Create,
					representationCopyWithNewProperties(emailDomainRepresentation, map[string]interface{}{
						"compartment_id": Representation{repType: Required, create: `${var.compartment_id_for_update}`},
					})),
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "name", randomDomain),

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
			Config: config + compartmentIdVariableStr + EmailDomainResourceDependencies +
				generateResourceFromRepresentationMap("oci_email_email_domain", "test_email_domain", Optional, Update, emailDomainRepresentation),
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "name", randomDomain),

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
				generateDataSourceFromRepresentationMap("oci_email_email_domains", "test_email_domains", Optional, Update, emailDomainDataSourceRepresentation) +
				compartmentIdVariableStr + EmailDomainResourceDependencies +
				generateResourceFromRepresentationMap("oci_email_email_domain", "test_email_domain", Optional, Update, emailDomainRepresentation),
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "name", randomDomain),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "email_domain_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "email_domain_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				generateDataSourceFromRepresentationMap("oci_email_email_domain", "test_email_domain", Required, Create, emailDomainSingularDataSourceRepresentation) +
				compartmentIdVariableStr + EmailDomainResourceConfig,
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "email_domain_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_spf"),
				resource.TestCheckResourceAttr(singularDatasourceName, "name", randomDomain),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + EmailDomainResourceConfig,
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

func testAccCheckEmailEmailDomainDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).emailClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_email_email_domain" {
			noResourceFound = false
			request := oci_email.GetEmailDomainRequest{}

			tmp := rs.Primary.ID
			request.EmailDomainId = &tmp

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "email")

			response, err := client.GetEmailDomain(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_email.EmailDomainLifecycleStateDeleted): true,
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
	if !inSweeperExcludeList("EmailEmailDomain") {
		resource.AddTestSweepers("EmailEmailDomain", &resource.Sweeper{
			Name:         "EmailEmailDomain",
			Dependencies: DependencyGraph["emailDomain"],
			F:            sweepEmailEmailDomainResource,
		})
	}
}

func sweepEmailEmailDomainResource(compartment string) error {
	emailClient := GetTestClients(&schema.ResourceData{}).emailClient()
	emailDomainIds, err := getEmailDomainIds(compartment)
	if err != nil {
		return err
	}
	for _, emailDomainId := range emailDomainIds {
		if ok := SweeperDefaultResourceId[emailDomainId]; !ok {
			deleteEmailDomainRequest := oci_email.DeleteEmailDomainRequest{}

			deleteEmailDomainRequest.EmailDomainId = &emailDomainId

			deleteEmailDomainRequest.RequestMetadata.RetryPolicy = getRetryPolicy(true, "email")
			_, error := emailClient.DeleteEmailDomain(context.Background(), deleteEmailDomainRequest)
			if error != nil {
				fmt.Printf("Error deleting EmailDomain %s %s, It is possible that the resource is already deleted. Please verify manually \n", emailDomainId, error)
				continue
			}
			waitTillCondition(testAccProvider, &emailDomainId, emailDomainSweepWaitCondition, time.Duration(3*time.Minute),
				emailDomainSweepResponseFetchOperation, "email", true)
		}
	}
	return nil
}

func getEmailDomainIds(compartment string) ([]string, error) {
	ids := getResourceIdsToSweep(compartment, "EmailDomainId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	emailClient := GetTestClients(&schema.ResourceData{}).emailClient()

	listEmailDomainsRequest := oci_email.ListEmailDomainsRequest{}
	listEmailDomainsRequest.CompartmentId = &compartmentId
	listEmailDomainsRequest.LifecycleState = oci_email.EmailDomainLifecycleStateActive
	listEmailDomainsResponse, err := emailClient.ListEmailDomains(context.Background(), listEmailDomainsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting EmailDomain list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, emailDomain := range listEmailDomainsResponse.Items {
		id := *emailDomain.Id
		resourceIds = append(resourceIds, id)
		addResourceIdToSweeperResourceIdMap(compartmentId, "EmailDomainId", id)
	}
	return resourceIds, nil
}

func emailDomainSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is ACTIVE beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if emailDomainResponse, ok := response.Response.(oci_email.GetEmailDomainResponse); ok {
		return emailDomainResponse.LifecycleState != oci_email.EmailDomainLifecycleStateDeleted
	}
	return false
}

func emailDomainSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.emailClient().GetEmailDomain(context.Background(), oci_email.GetEmailDomainRequest{
		EmailDomainId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
