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
	oci_email "github.com/oracle/oci-go-sdk/v65/email"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	ignoreEmailReturnPathDefinedTagsChangesRepresentation = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`freeform_tags`, `defined_tags`}},
	}
	EmailEmailReturnPathRequiredOnlyResource = EmailEmailReturnPathResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_email_email_return_path", "test_email_return_path", acctest.Required, acctest.Create, EmailEmailReturnPathRepresentation)

	EmailEmailReturnPathResourceConfig = EmailEmailReturnPathResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_email_email_return_path", "test_email_return_path", acctest.Optional, acctest.Update, EmailEmailReturnPathRepresentation)

	EmailEmailReturnPathSingularDataSourceRepresentation = map[string]interface{}{
		"email_return_path_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_email_email_return_path.test_email_return_path.id}`},
	}

	EmailEmailReturnPathDataSourceRepresentation = map[string]interface{}{
		"compartment_id":     acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"id":                 acctest.Representation{RepType: acctest.Optional, Create: `${oci_email_email_return_path.test_email_return_path.id}`},
		"name":               acctest.Representation{RepType: acctest.Optional, Create: "tfrp." + `${oci_email_email_domain.test_email_domain.name}`},
		"parent_resource_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_email_email_domain.test_email_domain.id}`},
		"state":              acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":             acctest.RepresentationGroup{RepType: acctest.Required, Group: EmailEmailReturnPathDataSourceFilterRepresentation}}
	EmailEmailReturnPathDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_email_email_return_path.test_email_return_path.id}`}},
	}

	EmailEmailReturnPathRepresentation = map[string]interface{}{
		"parent_resource_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_email_email_domain.test_email_domain.id}`},
		"defined_tags":       acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":        acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"freeform_tags":      acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"name":               acctest.Representation{RepType: acctest.Optional, Create: "tfrp." + `${oci_email_email_domain.test_email_domain.name}`},
		"lifecycle":          acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreEmailReturnPathDefinedTagsChangesRepresentation},
	}

	EmailEmailReturnPathResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_email_email_domain", "test_email_domain", acctest.Required, acctest.Create, acctest.RepresentationCopyWithNewProperties(EmailEmailDomainRepresentation, map[string]interface{}{"name": acctest.Representation{RepType: acctest.Required, Create: "objdomain.email.ap-mumbai-1.oci.oc-test.com"}, "domain_verification_id": acctest.Representation{RepType: acctest.Required, Create: utils.GetEnvSettingWithBlankDefault("domain_verification_id")}})) + DefinedTagsDependencies
)

// issue-routing-tag: email/default
func TestEmailEmailReturnPathResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestEmailEmailReturnPathResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_email_email_return_path.test_email_return_path"
	datasourceName := "data.oci_email_email_return_paths.test_email_return_paths"
	singularDatasourceName := "data.oci_email_email_return_path.test_email_return_path"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+EmailEmailReturnPathResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_email_email_return_path", "test_email_return_path", acctest.Optional, acctest.Create, EmailEmailReturnPathRepresentation), "email", "emailReturnPath", t)

	acctest.ResourceTest(t, testAccCheckEmailEmailReturnPathDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + EmailEmailReturnPathResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_email_email_return_path", "test_email_return_path", acctest.Required, acctest.Create, EmailEmailReturnPathRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "parent_resource_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + EmailEmailReturnPathResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + EmailEmailReturnPathResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_email_email_return_path", "test_email_return_path", acctest.Optional, acctest.Create, EmailEmailReturnPathRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "name", "tfrp.objdomain.email.ap-mumbai-1.oci.oc-test.com"),
				resource.TestCheckResourceAttrSet(resourceName, "parent_resource_id"),

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

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + EmailEmailReturnPathResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_email_email_return_path", "test_email_return_path", acctest.Optional, acctest.Update, EmailEmailReturnPathRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "name", "tfrp.objdomain.email.ap-mumbai-1.oci.oc-test.com"),
				resource.TestCheckResourceAttrSet(resourceName, "parent_resource_id"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_email_email_return_paths", "test_email_return_paths", acctest.Optional, acctest.Update, EmailEmailReturnPathDataSourceRepresentation) +
				compartmentIdVariableStr + EmailEmailReturnPathResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_email_email_return_path", "test_email_return_path", acctest.Optional, acctest.Update, EmailEmailReturnPathRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttr(datasourceName, "name", "tfrp.objdomain.email.ap-mumbai-1.oci.oc-test.com"),
				resource.TestCheckResourceAttrSet(datasourceName, "parent_resource_id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "email_return_path_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "email_return_path_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_email_email_return_path", "test_email_return_path", acctest.Required, acctest.Create, EmailEmailReturnPathSingularDataSourceRepresentation) +
				compartmentIdVariableStr + EmailEmailReturnPathResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "email_return_path_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "cname_record_value"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "dns_subdomain_name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "name", "tfrp.objdomain.email.ap-mumbai-1.oci.oc-test.com"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + EmailEmailReturnPathRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckEmailEmailReturnPathDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).EmailClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_email_email_return_path" {
			noResourceFound = false
			request := oci_email.GetEmailReturnPathRequest{}

			tmp := rs.Primary.ID
			request.EmailReturnPathId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "email")

			response, err := client.GetEmailReturnPath(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_email.EmailReturnPathLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("EmailEmailReturnPath") {
		resource.AddTestSweepers("EmailEmailReturnPath", &resource.Sweeper{
			Name:         "EmailEmailReturnPath",
			Dependencies: acctest.DependencyGraph["emailReturnPath"],
			F:            sweepEmailEmailReturnPathResource,
		})
	}
}

func sweepEmailEmailReturnPathResource(compartment string) error {
	emailClient := acctest.GetTestClients(&schema.ResourceData{}).EmailClient()
	emailReturnPathIds, err := getEmailEmailReturnPathIds(compartment)
	if err != nil {
		return err
	}
	for _, emailReturnPathId := range emailReturnPathIds {
		if ok := acctest.SweeperDefaultResourceId[emailReturnPathId]; !ok {
			deleteEmailReturnPathRequest := oci_email.DeleteEmailReturnPathRequest{}

			deleteEmailReturnPathRequest.EmailReturnPathId = &emailReturnPathId

			deleteEmailReturnPathRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "email")
			_, error := emailClient.DeleteEmailReturnPath(context.Background(), deleteEmailReturnPathRequest)
			if error != nil {
				fmt.Printf("Error deleting EmailReturnPath %s %s, It is possible that the resource is already deleted. Please verify manually \n", emailReturnPathId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &emailReturnPathId, EmailEmailReturnPathSweepWaitCondition, time.Duration(3*time.Minute),
				EmailEmailReturnPathSweepResponseFetchOperation, "email", true)
		}
	}
	return nil
}

func getEmailEmailReturnPathIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "EmailReturnPathId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	emailClient := acctest.GetTestClients(&schema.ResourceData{}).EmailClient()

	listEmailReturnPathsRequest := oci_email.ListEmailReturnPathsRequest{}
	listEmailReturnPathsRequest.CompartmentId = &compartmentId
	listEmailReturnPathsRequest.LifecycleState = oci_email.EmailReturnPathLifecycleStateActive
	listEmailReturnPathsResponse, err := emailClient.ListEmailReturnPaths(context.Background(), listEmailReturnPathsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting EmailReturnPath list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, emailReturnPath := range listEmailReturnPathsResponse.Items {
		id := *emailReturnPath.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "EmailReturnPathId", id)
	}
	return resourceIds, nil
}

func EmailEmailReturnPathSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if emailReturnPathResponse, ok := response.Response.(oci_email.GetEmailReturnPathResponse); ok {
		return emailReturnPathResponse.LifecycleState != oci_email.EmailReturnPathLifecycleStateDeleted
	}
	return false
}

func EmailEmailReturnPathSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.EmailClient().GetEmailReturnPath(context.Background(), oci_email.GetEmailReturnPathRequest{
		EmailReturnPathId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
