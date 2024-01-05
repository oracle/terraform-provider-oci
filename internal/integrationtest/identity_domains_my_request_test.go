// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"log"
	"strconv"
	"testing"

	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	IdentityDomainsMyRequestRequiredOnlyResource = IdentityDomainsMyRequestResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_my_request", "test_my_request", acctest.Required, acctest.Create, IdentityDomainsMyRequestRepresentation)

	IdentityDomainsIdentityDomainsMyRequestDataSourceRepresentation = map[string]interface{}{
		"idcs_endpoint":     acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"my_request_count":  acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"my_request_filter": acctest.Representation{RepType: acctest.Optional, Create: ``},
		// Not using `all` because 'approvalDetails' is not requestable when ListMyRequests
		"attribute_sets": acctest.Representation{RepType: acctest.Optional, Create: []string{``}},
		"start_index":    acctest.Representation{RepType: acctest.Optional, Create: `1`},
	}

	IdentityDomainsMyRequestRepresentation = map[string]interface{}{
		"idcs_endpoint":  acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"justification":  acctest.Representation{RepType: acctest.Required, Create: `justification`},
		"requesting":     acctest.RepresentationGroup{RepType: acctest.Required, Group: IdentityDomainsMyRequestRequestingRepresentation},
		"schemas":        acctest.Representation{RepType: acctest.Required, Create: []string{`urn:ietf:params:scim:schemas:oracle:idcs:Request`}},
		"attribute_sets": acctest.Representation{RepType: acctest.Optional, Create: []string{`all`}},
		"tags":           acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityDomainsMyRequestTagsRepresentation},
		"lifecycle":      acctest.RepresentationGroup{RepType: acctest.Optional, Group: ignoreChangeForIdentityDomainsMyRequest},
	}
	ignoreChangeForIdentityDomainsMyRequest = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Optional, Create: []string{
			`tags`, // my_* resource will not return non-default attributes
		}},
	}
	IdentityDomainsMyRequestRequestingRepresentation = map[string]interface{}{
		"type":  acctest.Representation{RepType: acctest.Required, Create: `Group`},
		"value": acctest.Representation{RepType: acctest.Required, Create: `${oci_identity_domains_group.test_group.id}`},
	}
	IdentityDomainsMyRequestTagsRepresentation = map[string]interface{}{
		"key":   acctest.Representation{RepType: acctest.Required, Create: `key`},
		"value": acctest.Representation{RepType: acctest.Required, Create: `value`},
	}

	IdentityDomainsMyRequestResourceDependencies = TestDomainForMyEndpointDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_group", "test_group", acctest.Required, acctest.Create, acctest.RepresentationCopyWithNewProperties(IdentityDomainsGroupRepresentation, map[string]interface{}{"attribute_sets": acctest.Representation{RepType: acctest.Required, Create: []string{`all`}}, "urnietfparamsscimschemasoracleidcsextensionrequestable_group": acctest.RepresentationGroup{RepType: acctest.Required, Group: map[string]interface{}{"requestable": acctest.Representation{RepType: acctest.Required, Create: `true`}}}}))
)

// issue-routing-tag: identity_domains/default
func TestIdentityDomainsMyRequestResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestIdentityDomainsMyRequestResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_identity_domains_my_request.test_my_request"
	datasourceName := "data.oci_identity_domains_my_requests.test_my_requests"

	var resId string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+IdentityDomainsMyRequestResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_my_request", "test_my_request", acctest.Optional, acctest.Create, IdentityDomainsMyRequestRepresentation), "identitydomains", "myRequest", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsMyRequestResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_my_request", "test_my_request", acctest.Required, acctest.Create, IdentityDomainsMyRequestRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttr(resourceName, "justification", "justification"),
				resource.TestCheckResourceAttr(resourceName, "requesting.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "requesting.0.type", "Group"),
				resource.TestCheckResourceAttrSet(resourceName, "requesting.0.value"),
				resource.TestCheckResourceAttr(resourceName, "schemas.#", "1"),
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsMyRequestResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_my_request", "test_my_request", acctest.Optional, acctest.Create, IdentityDomainsMyRequestRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "attribute_sets.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttr(resourceName, "justification", "justification"),
				resource.TestCheckResourceAttr(resourceName, "requesting.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "requesting.0.type", "Group"),
				resource.TestCheckResourceAttrSet(resourceName, "requesting.0.value"),
				resource.TestCheckResourceAttr(resourceName, "schemas.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "status"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					idcsEndpoint, err := acctest.FromInstanceState(s, "data.oci_identity_domain.test_domain", "url")
					if err != nil {
						return err
					}

					compositeId := getIdentityDomainsCompositeId(idcsEndpoint, "myRequests", resId)
					log.Printf("[DEBUG] Composite ID to import: %s", compositeId)
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&compositeId, &compartmentId, resourceName); errExport != nil {
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_domains_my_requests", "test_my_requests", acctest.Optional, acctest.Update, IdentityDomainsIdentityDomainsMyRequestDataSourceRepresentation) +
				compartmentIdVariableStr + IdentityDomainsMyRequestResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_my_request", "test_my_request", acctest.Optional, acctest.Update, IdentityDomainsMyRequestRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "my_request_count", "10"),
				resource.TestCheckResourceAttr(datasourceName, "attribute_sets.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "start_index", "1"),

				resource.TestCheckResourceAttrSet(datasourceName, "my_requests.#"),
				resource.TestCheckResourceAttr(datasourceName, "my_requests.0.schemas.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "total_results"),
			),
		},
	})
}
