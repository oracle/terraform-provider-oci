// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	ResolverEndpointRequiredOnlyResource = ResolverEndpointResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_dns_resolver_endpoint", "test_resolver_endpoint", acctest.Required, acctest.Create, resolverEndpointRepresentation)

	ResolverEndpointResourceConfig = ResolverEndpointResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_dns_resolver_endpoint", "test_resolver_endpoint", acctest.Optional, acctest.Update, resolverEndpointRepresentation)

	resolverEndpointSingularDataSourceRepresentation = map[string]interface{}{
		"resolver_endpoint_name": acctest.Representation{RepType: acctest.Required, Create: `${oci_dns_resolver_endpoint.test_resolver_endpoint.name}`},
		"resolver_id":            acctest.Representation{RepType: acctest.Required, Create: `${data.oci_core_vcn_dns_resolver_association.test_vcn_dns_resolver_association.dns_resolver_id}`},
		"scope":                  acctest.Representation{RepType: acctest.Required, Create: `PRIVATE`},
	}

	resolverEndpointDataSourceRepresentation = map[string]interface{}{
		"resolver_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_core_vcn_dns_resolver_association.test_vcn_dns_resolver_association.dns_resolver_id}`},
		"name":        acctest.Representation{RepType: acctest.Optional, Create: `endpointName`},
		"scope":       acctest.Representation{RepType: acctest.Required, Create: `PRIVATE`},
		"state":       acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":      acctest.RepresentationGroup{RepType: acctest.Required, Group: resolverEndpointDataSourceFilterRepresentation}}

	resolverEndpointDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `name`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_dns_resolver_endpoint.test_resolver_endpoint.name}`}},
	}

	resolverEndpointRepresentation = map[string]interface{}{
		"is_forwarding":      acctest.Representation{RepType: acctest.Required, Create: `true`},
		"is_listening":       acctest.Representation{RepType: acctest.Required, Create: `false`},
		"name":               acctest.Representation{RepType: acctest.Required, Create: `endpointName`},
		"resolver_id":        acctest.Representation{RepType: acctest.Required, Create: `${oci_dns_resolver.test_resolver.id}`},
		"subnet_id":          acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_subnet.id}`},
		"endpoint_type":      acctest.Representation{RepType: acctest.Optional, Create: `VNIC`},
		"forwarding_address": acctest.Representation{RepType: acctest.Optional, Create: `10.0.0.5`},
		"scope":              acctest.Representation{RepType: acctest.Required, Create: `PRIVATE`},
		"nsg_ids":            acctest.Representation{RepType: acctest.Optional, Create: []string{`${oci_core_network_security_group.test_network_security_group.id}`}},
	}

	resolverEndpointRepresentationWithoutNsgId = acctest.RepresentationCopyWithRemovedProperties(resolverEndpointRepresentation, []string{"nsg_ids"})

	ResolverEndpointResourceDependencies = ResolverResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, subnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_network_security_group", "test_network_security_group", acctest.Required, acctest.Create, networkSecurityGroupRepresentation)
)

// issue-routing-tag: dns/default
func TestDnsResolverEndpointResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDnsResolverEndpointResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_dns_resolver_endpoint.test_resolver_endpoint"

	datasourceName := "data.oci_dns_resolver_endpoints.test_resolver_endpoints"
	singularDatasourceName := "data.oci_dns_resolver_endpoint.test_resolver_endpoint"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+ResolverEndpointResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_dns_resolver_endpoint", "test_resolver_endpoint", acctest.Optional, acctest.Create, resolverEndpointRepresentation), "dns", "resolverEndpoint", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// Create dependencies
		{
			Config: config + compartmentIdVariableStr + ResolverEndpointResourceDependencies,
			Check: func(s *terraform.State) (err error) {
				log.Printf("[DEBUG] Wait for 2 minutes for oci_core_vcn resource to get created")
				time.Sleep(2 * time.Minute)
				return nil
			},
		},
		// verify Create
		{
			Config: config + compartmentIdVariableStr + ResolverEndpointResourceDependencies +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_vcn_dns_resolver_association", "test_vcn_dns_resolver_association", acctest.Required, acctest.Create, vcnDnsResolverAssociationSingularDataSourceRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_dns_resolver", "test_resolver", acctest.Required, acctest.Create, resolverRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_dns_resolver_endpoint", "test_resolver_endpoint", acctest.Required, acctest.Create, resolverEndpointRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "is_forwarding", "true"),
				resource.TestCheckResourceAttr(resourceName, "is_listening", "false"),
				resource.TestCheckResourceAttr(resourceName, "name", "endpointName"),
				resource.TestCheckResourceAttrSet(resourceName, "resolver_id"),

				func(s *terraform.State) (err error) {
					_, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + ResolverEndpointResourceDependencies +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_vcn_dns_resolver_association", "test_vcn_dns_resolver_association", acctest.Required, acctest.Create, vcnDnsResolverAssociationSingularDataSourceRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_dns_resolver", "test_resolver", acctest.Required, acctest.Create, resolverRepresentation),
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + ResolverEndpointResourceDependencies +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_vcn_dns_resolver_association", "test_vcn_dns_resolver_association", acctest.Required, acctest.Create, vcnDnsResolverAssociationSingularDataSourceRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_dns_resolver", "test_resolver", acctest.Required, acctest.Create, resolverRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_dns_resolver_endpoint", "test_resolver_endpoint", acctest.Optional, acctest.Create, resolverEndpointRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttr(resourceName, "endpoint_type", "VNIC"),
				resource.TestCheckResourceAttr(resourceName, "forwarding_address", "10.0.0.5"),
				resource.TestCheckResourceAttr(resourceName, "is_forwarding", "true"),
				resource.TestCheckResourceAttr(resourceName, "is_listening", "false"),
				resource.TestCheckResourceAttr(resourceName, "name", "endpointName"),
				resource.TestCheckResourceAttrSet(resourceName, "resolver_id"),
				resource.TestCheckResourceAttr(resourceName, "scope", "PRIVATE"),
				resource.TestCheckResourceAttrSet(resourceName, "self"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					// Resource discovery is disabled for Resolver Endpoints
					//if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
					//	if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
					//		return errExport
					//	}
					//}
					return err
				},
			),
		},
		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + ResolverEndpointResourceDependencies +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_vcn_dns_resolver_association", "test_vcn_dns_resolver_association", acctest.Required, acctest.Update, vcnDnsResolverAssociationSingularDataSourceRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_dns_resolver", "test_resolver", acctest.Required, acctest.Create, resolverRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_dns_resolver_endpoint", "test_resolver_endpoint", acctest.Optional, acctest.Update, resolverEndpointRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttr(resourceName, "endpoint_type", "VNIC"),
				resource.TestCheckResourceAttr(resourceName, "forwarding_address", "10.0.0.5"),
				resource.TestCheckResourceAttr(resourceName, "is_forwarding", "true"),
				resource.TestCheckResourceAttr(resourceName, "is_listening", "false"),
				resource.TestCheckResourceAttr(resourceName, "name", "endpointName"),
				resource.TestCheckResourceAttrSet(resourceName, "resolver_id"),
				resource.TestCheckResourceAttr(resourceName, "scope", "PRIVATE"),
				resource.TestCheckResourceAttrSet(resourceName, "self"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

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
			Config: config + compartmentIdVariableStr + ResolverEndpointResourceDependencies +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_vcn_dns_resolver_association", "test_vcn_dns_resolver_association", acctest.Required, acctest.Create, vcnDnsResolverAssociationSingularDataSourceRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_dns_resolver_endpoints", "test_resolver_endpoints", acctest.Optional, acctest.Update, resolverEndpointDataSourceRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_dns_resolver", "test_resolver", acctest.Required, acctest.Create, resolverRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_dns_resolver_endpoint", "test_resolver_endpoint", acctest.Optional, acctest.Update, resolverEndpointRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "name", "endpointName"),
				resource.TestCheckResourceAttrSet(datasourceName, "resolver_id"),
				resource.TestCheckResourceAttr(datasourceName, "scope", "PRIVATE"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttr(datasourceName, "resolver_endpoints.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "resolver_endpoints.0.compartment_id"),
				resource.TestCheckResourceAttr(datasourceName, "resolver_endpoints.0.endpoint_type", "VNIC"),
				resource.TestCheckResourceAttr(datasourceName, "resolver_endpoints.0.forwarding_address", "10.0.0.5"),
				resource.TestCheckResourceAttr(datasourceName, "resolver_endpoints.0.is_forwarding", "true"),
				resource.TestCheckResourceAttr(datasourceName, "resolver_endpoints.0.is_listening", "false"),
				resource.TestCheckResourceAttr(datasourceName, "resolver_endpoints.0.name", "endpointName"),
				resource.TestCheckResourceAttrSet(datasourceName, "resolver_endpoints.0.self"),
				resource.TestCheckResourceAttrSet(datasourceName, "resolver_endpoints.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "resolver_endpoints.0.subnet_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "resolver_endpoints.0.time_created"),
				resource.TestCheckResourceAttrSet(datasourceName, "resolver_endpoints.0.time_updated"),
			),
		},
		// verify singular datasource
		{
			Config: config + compartmentIdVariableStr + ResolverEndpointResourceDependencies +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_vcn_dns_resolver_association", "test_vcn_dns_resolver_association", acctest.Required, acctest.Create, vcnDnsResolverAssociationSingularDataSourceRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_dns_resolver_endpoint", "test_resolver_endpoint", acctest.Optional, acctest.Update, resolverEndpointSingularDataSourceRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_dns_resolver", "test_resolver", acctest.Required, acctest.Create, resolverRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_dns_resolver_endpoint", "test_resolver_endpoint", acctest.Optional, acctest.Update, resolverEndpointRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "resolver_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "scope", "PRIVATE"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "endpoint_type", "VNIC"),
				resource.TestCheckResourceAttr(singularDatasourceName, "forwarding_address", "10.0.0.5"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_forwarding", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_listening", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "name", "endpointName"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "self"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + ResolverEndpointResourceDependencies +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_vcn_dns_resolver_association", "test_vcn_dns_resolver_association", acctest.Required, acctest.Create, vcnDnsResolverAssociationSingularDataSourceRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_dns_resolver", "test_resolver", acctest.Required, acctest.Create, resolverRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_dns_resolver_endpoint", "test_resolver_endpoint", acctest.Optional, acctest.Update, resolverEndpointRepresentation),
		},
		// verify resource import
		{
			Config:            config,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateIdFunc: getResolverEndpointImportId(resourceName),
			ImportStateVerifyIgnore: []string{
				"scope",
			},
			ResourceName: resourceName,
		},
	})
}

func getResolverEndpointImportId(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return "", fmt.Errorf("not found: %s", resourceName)
		}
		return fmt.Sprintf("resolverId/" + rs.Primary.Attributes["resolver_id"] + "/name/" + rs.Primary.Attributes["name"] + "/scope/" + rs.Primary.Attributes["scope"]), nil
	}
}
