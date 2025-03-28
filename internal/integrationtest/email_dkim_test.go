// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_email "github.com/oracle/oci-go-sdk/v65/email"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	ignoreDkimBaseDefinedTagsChangesRepresentation = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`}},
	}
	EmailDkimRequiredOnlyResource = EmailDkimResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_email_dkim", "test_dkim", acctest.Required, acctest.Create, EmailDkimRepresentation)

	EmailDkimResourceConfig = EmailDkimResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_email_dkim", "test_dkim", acctest.Optional, acctest.Update, EmailDkimRepresentation)

	EmailEmailDkimSingularDataSourceRepresentation = map[string]interface{}{
		"dkim_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_email_dkim.test_dkim.id}`},
	}

	EmailEmailDkimDataSourceRepresentation = map[string]interface{}{
		"email_domain_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_email_email_domain.test_email_domain.id}`},
		"id":              acctest.Representation{RepType: acctest.Optional, Create: `${oci_email_dkim.test_dkim.id}`},
		"name":            acctest.Representation{RepType: acctest.Optional, Create: `testselector1`},
		"state":           acctest.Representation{RepType: acctest.Optional, Create: `NEEDS_ATTENTION`},
		"filter":          acctest.RepresentationGroup{RepType: acctest.Required, Group: EmailDkimDataSourceFilterRepresentation}}
	EmailDkimDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_email_dkim.test_dkim.id}`}},
	}

	EmailDkimRepresentation = map[string]interface{}{
		"email_domain_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_email_email_domain.test_email_domain.id}`},
		"defined_tags":    acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`},
		"description":     acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description`},
		"freeform_tags":   acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Finance"}},
		"name":            acctest.Representation{RepType: acctest.Optional, Create: `testselector1`},
		"private_key":     acctest.Representation{RepType: acctest.Optional, Create: `MIIEvwIBADANBgkqhkiG9w0BAQEFAASCBKkwggSlAgEAAoIBAQDYOJHqR6O6TofL3eUshL+yDN9sTDSDCmSMe5hapcu1Cj6zjXLbPvEUM+SZd2MrBaAfh5HaiM0U3hQgPX//6zAewT95n5K94qnCo1yKBYpt81lnY5PRcePOE0wuFQ8EqqJNbIPFojYLpoyKazB7ulCiA1y6Zk5ZzUolxNslnRwjWtNidC62ureDbDsE28LVhHron7+s5StR5cMTRUxU1nU3KeOluUS2YwJQzvdAQiDn9PHnG0k7p9+Cam0mlNPgPUKIKmFigx+WBrrI5k1yxyYVjbwTQgl7J7bEz0kZZJLN0Tb1UVfIe+yTDz9iZgMDZ0+spo9FA4UEpQiHsr3Cih6XAgMBAAECggEADtEZlq9gFeVJwvuqt7lFFaNyMA6szooIx+O0Vgnkgrx3/X/k7gdLJ/JCaQ3YWttg1KhKo2oae6KYSHlOazAntQPTILeBq8d7kJ33OnIUwWIar04QlBq03KtP3hyNUemLg7i68vEhuPypxtZ/lQr4eZG3agKWveqmeR4bUmGIyxdO6GaXQMoGo8javYylS02X3uW/jTT4cXWcnslBibQ48CeGS7fuxGHHoZgU0qq3JDlO2KZtOjnmMvwqhAlnCgcd0EX1exsddWXhi2R7ONMRXT/0XG16AVFUumqU5c1E3qNpB+/OR7GO1eIxXsIIN/lBBLblIuETV0naoEeIR0cC4QKBgQD/bEe8cqJlICDDBucNLviTXU2xWHbg4rik1l+NxcN5oRuF0klEklCKPeoiruwFuOxjR4v+E5E6OY/X3lEdVbm7sQpX5rU5Ql9Mkuikniq0gji3aG1fgp1S+rltie1IwIvWWJII5SKrl5zk+K/1F2fiYIeS3xDpBXZAEuYkmDlXcQKBgQDYtZ4v1gGqz2avqSurVgIBKuSy5lQW4RCctGHvaaMMqmE5GipCId/2toIxCm/TTnkw2xy5iLN+TRu1ejQlnXyjbv5rWpg6Ud4ZfS+dspBZKQDndcUxBb8AOEEa3mfNnqtf2fpLNiBHHQUQbQ0n+Mvlk0P9RCIF8cakw8vTUbAihwKBgQDRovB5xehtcWIST68BTgTtBk5G+WGQj+SWPfBtLtySDl3FhuNlqWT7IXfygXjr9WvmzAklPrAahrX9qdUPa5Wma9cecBdbaJNJCOZyffJ09tBQECB8nHo9adUZB3Q+GvJquYUaUd4T0bsO3uGkGlLEJRAAxp5dcoMscZZ8gIfHsQKBgQCwlu6cVGtvYSpwIg5vw9pGUUmjboU/T49V4rebfB1diMlI2sVstmXmiLp2c+bD7cJbyXESzw5U8UiPDVHJBRg3pY60tFIv096ELrrKamOuA/e0pQdW9zO5Xh3PY0R7DbzgkqOe/jA3QrKJvBNoaxjbrrFMhGSlK2uPf+1r+bjpqQKBgQDvSKV73DAKgi0SWreJsBxzJ5qXTj5vcFfu4ySYzoCabfyH6sSokhRUHrzT0V3I+oF0N84IqoDX4CxnNmmu4DvZx30geMSlkFEzY2hPktEcO3nyEgBEFd8PVIp3NqL9u3N8rfQIfy8kudoeZWg3+9sjIFkEiX/HghFf4tt7qcGRwg==`},
		"lifecycle":       acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDkimBaseDefinedTagsChangesRepresentation},
	}

	EmailDkimResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_email_email_domain", "test_email_domain", acctest.Required, acctest.Create, EmailEmailDomainRepresentation) +
		DefinedTagsDependencies
)

func TestEmailDkimResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestEmailDkimResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_email_dkim.test_dkim"
	datasourceName := "data.oci_email_dkims.test_dkims"
	singularDatasourceName := "data.oci_email_dkim.test_dkim"

	var resId, resId2 string
	// Save TF content to Create resource with acctest.Optional, properties. This has to be exactly the same as the config part in the "Create with acctest.Optional,s" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+EmailDkimResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_email_dkim", "test_dkim", acctest.Optional, acctest.Create, EmailDkimRepresentation), "email", "dkim", t)

	acctest.ResourceTest(t, testAccCheckEmailDkimDestroy, []resource.TestStep{
		//verify Create
		{
			Config: config + compartmentIdVariableStr + EmailDkimResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_email_dkim", "test_dkim", acctest.Required, acctest.Create, EmailDkimRepresentation),
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttrSet(resourceName, "email_domain_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + EmailDkimResourceDependencies,
		},
		// verify Create with acctest.Optional,s
		{
			Config: config + compartmentIdVariableStr + EmailDkimResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_email_dkim", "test_dkim", acctest.Optional, acctest.Create, EmailDkimRepresentation),
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttrSet(resourceName, "email_domain_id"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "name", "testselector1"),
				resource.TestCheckResourceAttrSet(resourceName, "private_key"),

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
			Config: config + compartmentIdVariableStr + EmailDkimResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_email_dkim", "test_dkim", acctest.Optional, acctest.Update, EmailDkimRepresentation),
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttrSet(resourceName, "email_domain_id"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "name", "testselector1"),
				resource.TestCheckResourceAttrSet(resourceName, "private_key"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_email_dkims", "test_dkims", acctest.Optional, acctest.Update, EmailEmailDkimDataSourceRepresentation) +
				compartmentIdVariableStr + EmailDkimResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_email_dkim", "test_dkim", acctest.Optional, acctest.Update, EmailDkimRepresentation),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_email_dkim", "test_dkim", acctest.Required, acctest.Create, EmailEmailDkimSingularDataSourceRepresentation) +
				compartmentIdVariableStr + EmailDkimResourceConfig,
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "dkim_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "cname_record_value"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "dns_subdomain_name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_imported"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "key_length"),
				resource.TestCheckResourceAttr(singularDatasourceName, "name", "testselector1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "txt_record_value"),
			),
		},
		// verify resource import
		{
			Config:            config + EmailDkimRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"private_key",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckEmailDkimDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).EmailClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_email_dkim" {
			noResourceFound = false
			request := oci_email.GetDkimRequest{}

			tmp := rs.Primary.ID
			request.DkimId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "email")

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
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("EmailDkim") {
		resource.AddTestSweepers("EmailDkim", &resource.Sweeper{
			Name:         "EmailDkim",
			Dependencies: acctest.DependencyGraph["dkim"],
			F:            sweepEmailDkimResource,
		})
	}
}

func sweepEmailDkimResource(compartment string) error {
	emailClient := acctest.GetTestClients(&schema.ResourceData{}).EmailClient()
	dkimIds, err := getEmailDkimIds(compartment)
	if err != nil {
		return err
	}
	for _, dkimId := range dkimIds {
		if ok := acctest.SweeperDefaultResourceId[dkimId]; !ok {
			deleteDkimRequest := oci_email.DeleteDkimRequest{}

			deleteDkimRequest.DkimId = &dkimId

			deleteDkimRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "email")
			_, error := emailClient.DeleteDkim(context.Background(), deleteDkimRequest)
			if error != nil {
				fmt.Printf("Error deleting Dkim %s %s, It is possible that the resource is already deleted. Please verify manually \n", dkimId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &dkimId, EmailDkimSweepWaitCondition, time.Duration(3*time.Minute),
				EmailDkimSweepResponseFetchOperation, "email", true)
		}
	}
	return nil
}

func getEmailDkimIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "DkimId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	emailClient := acctest.GetTestClients(&schema.ResourceData{}).EmailClient()

	listDkimsRequest := oci_email.ListDkimsRequest{}
	listDkimsRequest.Id = &compartmentId

	emailDomainIds, error := getEmailEmailDomainIds(compartment)
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
			acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "DkimId", id)
		}

	}
	return resourceIds, nil
}

func EmailDkimSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if dkimResponse, ok := response.Response.(oci_email.GetDkimResponse); ok {
		return dkimResponse.LifecycleState != oci_email.DkimLifecycleStateDeleted
	}
	return false
}

func EmailDkimSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.EmailClient().GetDkim(context.Background(), oci_email.GetDkimRequest{
		DkimId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
