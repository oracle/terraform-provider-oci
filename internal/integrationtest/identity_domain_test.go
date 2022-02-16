// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v58/common"
	oci_identity "github.com/oracle/oci-go-sdk/v58/identity"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	tf_client "github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/resourcediscovery"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"
)

var (
	DomainRequiredOnlyResource = DomainResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domain", "test_domain", acctest.Required, acctest.Create, domainRepresentation)

	DomainResourceConfig = DomainResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domain", "test_domain", acctest.Optional, acctest.Update, domainRepresentation)

	domainSingularDataSourceRepresentation = map[string]interface{}{
		"domain_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_identity_domain.test_domain.id}`},
	}

	domainDataSourceRepresentation = map[string]interface{}{
		"compartment_id":     acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`, Update: `${var.compartment_id}`},
		"display_name":       acctest.Representation{RepType: acctest.Optional, Create: `displayName9`, Update: `displayName9`},
		"is_hidden_on_login": acctest.Representation{RepType: acctest.Optional, Update: `true`},
		"license_type":       acctest.Representation{RepType: acctest.Optional, Create: `external-user`, Update: `external-user`},
		"state":              acctest.Representation{RepType: acctest.Optional, Create: `INACTIVE`},
		"type":               acctest.Representation{RepType: acctest.Optional, Create: string(oci_identity.DomainTypeSecondary), Update: string(oci_identity.DomainTypeSecondary)},
		"filter":             acctest.RepresentationGroup{RepType: acctest.Required, Group: domainDataSourceFilterRepresentation},
	}
	domainDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_identity_domain.test_domain.id}`}},
	}

	domainRepresentation = map[string]interface{}{
		"compartment_id":            acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"description":               acctest.Representation{RepType: acctest.Required, Create: `description`, Update: `description2`},
		"display_name":              acctest.Representation{RepType: acctest.Required, Create: `displayName11`, Update: `displayName9`},
		"home_region":               acctest.Representation{RepType: acctest.Required, Create: `ca-toronto-1`},
		"license_type":              acctest.Representation{RepType: acctest.Required, Create: `free`, Update: `external-user`},
		"admin_email":               acctest.Representation{RepType: acctest.Optional, Create: `adminEmail@test.com`},
		"admin_first_name":          acctest.Representation{RepType: acctest.Optional, Create: `adminFirstName`},
		"admin_last_name":           acctest.Representation{RepType: acctest.Optional, Create: `adminLastName`},
		"admin_user_name":           acctest.Representation{RepType: acctest.Optional, Create: `adminUserName`},
		"freeform_tags":             acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"is_hidden_on_login":        acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"is_notification_bypassed":  acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"is_primary_email_required": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"state":                     acctest.Representation{RepType: acctest.Optional, Update: `INACTIVE`},
	}

	DomainResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: identity/default
func TestIdentityDomainResource_basic(t *testing.T) {
	t.Skip("Skip this test because henosis tenancy is needed")
	httpreplay.SetScenario("TestIdentityDomainResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_identity_domain.test_domain"
	datasourceName := "data.oci_identity_domains.test_domains"
	singularDatasourceName := "data.oci_identity_domain.test_domain"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DomainResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domain", "test_domain", acctest.Optional, acctest.Create, domainRepresentation), "identity", "domain", t)

	acctest.ResourceTest(t, testAccCheckIdentityDomainDestroy, []resource.TestStep{
		// verify Create and deactivate domain
		{
			Config: config + compartmentIdVariableStr + DomainResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domain", "test_domain", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(domainRepresentation, map[string]interface{}{
						"state": acctest.Representation{RepType: acctest.Required, Create: `inactive`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName11"),
				resource.TestCheckResourceAttr(resourceName, "home_region", "ca-toronto-1"),
				resource.TestCheckResourceAttr(resourceName, "license_type", "free"),
				resource.TestCheckResourceAttr(resourceName, "state", "INACTIVE"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DomainResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DomainResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domain", "test_domain", acctest.Optional, acctest.Create, domainRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_email", "adminEmail@test.com"),
				resource.TestCheckResourceAttr(resourceName, "admin_first_name", "adminFirstName"),
				resource.TestCheckResourceAttr(resourceName, "admin_last_name", "adminLastName"),
				resource.TestCheckResourceAttr(resourceName, "admin_user_name", "adminUserName"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName11"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "home_region", "ca-toronto-1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_hidden_on_login", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_notification_bypassed", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_primary_email_required", "false"),
				resource.TestCheckResourceAttr(resourceName, "license_type", "free"),
				resource.TestCheckResourceAttr(resourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "type"),
				resource.TestCheckResourceAttrSet(resourceName, "url"),

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

		// verify update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + DomainResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domain", "test_domain", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(domainRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_email", "adminEmail@test.com"),
				resource.TestCheckResourceAttr(resourceName, "admin_first_name", "adminFirstName"),
				resource.TestCheckResourceAttr(resourceName, "admin_last_name", "adminLastName"),
				resource.TestCheckResourceAttr(resourceName, "admin_user_name", "adminUserName"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName11"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "home_region", "ca-toronto-1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_hidden_on_login", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_notification_bypassed", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_primary_email_required", "false"),
				resource.TestCheckResourceAttr(resourceName, "license_type", "free"),
				resource.TestCheckResourceAttr(resourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "type"),
				resource.TestCheckResourceAttrSet(resourceName, "url"),

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
			Config: config + compartmentIdVariableStr + DomainResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domain", "test_domain", acctest.Optional, acctest.Update, domainRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_email", "adminEmail@test.com"),
				resource.TestCheckResourceAttr(resourceName, "admin_first_name", "adminFirstName"),
				resource.TestCheckResourceAttr(resourceName, "admin_last_name", "adminLastName"),
				resource.TestCheckResourceAttr(resourceName, "admin_user_name", "adminUserName"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName9"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "home_region", "ca-toronto-1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_hidden_on_login", "true"),
				resource.TestCheckResourceAttr(resourceName, "is_notification_bypassed", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_primary_email_required", "false"),
				resource.TestCheckResourceAttr(resourceName, "license_type", "external-user"),
				resource.TestCheckResourceAttr(resourceName, "state", "INACTIVE"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "type"),
				resource.TestCheckResourceAttrSet(resourceName, "url"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_domains", "test_domains", acctest.Optional, acctest.Update, domainDataSourceRepresentation) +
				compartmentIdVariableStr + DomainResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domain", "test_domain", acctest.Optional, acctest.Update,
					acctest.RepresentationCopyWithNewProperties(domainRepresentation, map[string]interface{}{
						"type": acctest.Representation{RepType: acctest.Optional, Update: oci_identity.DomainTypeSecondary},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName9"),
				resource.TestCheckResourceAttr(datasourceName, "is_hidden_on_login", "true"),
				resource.TestCheckResourceAttr(datasourceName, "license_type", "external-user"),
				resource.TestCheckResourceAttr(datasourceName, "state", "INACTIVE"),
				resource.TestCheckResourceAttr(datasourceName, "type", string(oci_identity.DomainTypeSecondary)),

				resource.TestCheckResourceAttr(datasourceName, "domains.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "domains.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "domains.0.description", "description2"),
				resource.TestCheckResourceAttr(datasourceName, "domains.0.display_name", "displayName9"),
				resource.TestCheckResourceAttr(datasourceName, "domains.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(datasourceName, "domains.0.home_region", "ca-toronto-1"),
				resource.TestCheckResourceAttrSet(datasourceName, "domains.0.id"),
				resource.TestCheckResourceAttr(datasourceName, "domains.0.is_hidden_on_login", "true"),
				resource.TestCheckResourceAttr(datasourceName, "domains.0.license_type", "external-user"),
				resource.TestCheckResourceAttr(datasourceName, "domains.0.state", "INACTIVE"),
				resource.TestCheckResourceAttrSet(datasourceName, "domains.0.time_created"),
				resource.TestCheckResourceAttrSet(datasourceName, "domains.0.type"),
				resource.TestCheckResourceAttrSet(datasourceName, "domains.0.url"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_domain", "test_domain", acctest.Required, acctest.Create, domainSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DomainResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "domain_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName9"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "home_region", "ca-toronto-1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_hidden_on_login", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "license_type", "external-user"),
				resource.TestCheckResourceAttr(singularDatasourceName, "state", "INACTIVE"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "type"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "url"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + DomainResourceConfig,
		},
		// verify resource import
		{
			Config:            config,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"admin_email",
				"admin_first_name",
				"admin_last_name",
				"admin_user_name",
				"is_notification_bypassed",
				"is_primary_email_required",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckIdentityDomainDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).IdentityClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_identity_domain" {
			noResourceFound = false
			request := oci_identity.GetDomainRequest{}

			tmp := rs.Primary.ID
			request.DomainId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "identity")

			_, err := client.GetDomain(context.Background(), request)

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
	if !acctest.InSweeperExcludeList("IdentityDomain") {
		resource.AddTestSweepers("IdentityDomain", &resource.Sweeper{
			Name:         "IdentityDomain",
			Dependencies: acctest.DependencyGraph["domain"],
			F:            sweepIdentityDomainResource,
		})
	}
}

func sweepIdentityDomainResource(compartment string) error {
	identityClient := acctest.GetTestClients(&schema.ResourceData{}).IdentityClient()
	domainIds, err := getDomainIds(compartment)
	if err != nil {
		return err
	}
	for _, domainId := range domainIds {
		if ok := acctest.SweeperDefaultResourceId[domainId]; !ok {
			deleteDomainRequest := oci_identity.DeleteDomainRequest{}

			deleteDomainRequest.DomainId = &domainId

			deleteDomainRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "identity")
			_, error := identityClient.DeleteDomain(context.Background(), deleteDomainRequest)
			if error != nil {
				fmt.Printf("Error deleting Domain %s %s, It is possible that the resource is already deleted. Please verify manually \n", domainId, error)
				continue
			}
		}
	}
	return nil
}

func getDomainIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "DomainId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	identityClient := acctest.GetTestClients(&schema.ResourceData{}).IdentityClient()

	listDomainsRequest := oci_identity.ListDomainsRequest{}
	listDomainsRequest.CompartmentId = &compartmentId
	listDomainsResponse, err := identityClient.ListDomains(context.Background(), listDomainsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting Domain list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, domain := range listDomainsResponse.Items {
		id := *domain.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "DomainId", id)
	}
	return resourceIds, nil
}
