// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"log"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	ResolverEndpointRequiredOnlyResource = ResolverEndpointResourceDependencies +
		generateResourceFromRepresentationMap("oci_dns_resolver_endpoint", "test_resolver_endpoint", Required, Create, resolverEndpointRepresentation)

	ResolverEndpointResourceConfig = ResolverEndpointResourceDependencies +
		generateResourceFromRepresentationMap("oci_dns_resolver_endpoint", "test_resolver_endpoint", Optional, Update, resolverEndpointRepresentation)

	resolverEndpointSingularDataSourceRepresentation = map[string]interface{}{
		"resolver_endpoint_name": Representation{repType: Required, create: `${oci_dns_resolver_endpoint.test_resolver_endpoint.name}`},
		"resolver_id":            Representation{repType: Required, create: `${data.oci_core_vcn_dns_resolver_association.test_vcn_dns_resolver_association.dns_resolver_id}`},
		"scope":                  Representation{repType: Required, create: `PRIVATE`},
	}

	resolverEndpointDataSourceRepresentation = map[string]interface{}{
		"resolver_id": Representation{repType: Required, create: `${data.oci_core_vcn_dns_resolver_association.test_vcn_dns_resolver_association.dns_resolver_id}`},
		"name":        Representation{repType: Optional, create: `endpointName`},
		"scope":       Representation{repType: Required, create: `PRIVATE`},
		"state":       Representation{repType: Optional, create: `ACTIVE`},
		"filter":      RepresentationGroup{Required, resolverEndpointDataSourceFilterRepresentation}}

	resolverEndpointDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `name`},
		"values": Representation{repType: Required, create: []string{`${oci_dns_resolver_endpoint.test_resolver_endpoint.name}`}},
	}

	resolverEndpointRepresentation = map[string]interface{}{
		"is_forwarding":      Representation{repType: Required, create: `true`},
		"is_listening":       Representation{repType: Required, create: `false`},
		"name":               Representation{repType: Required, create: `endpointName`},
		"resolver_id":        Representation{repType: Required, create: `${oci_dns_resolver.test_resolver.id}`},
		"subnet_id":          Representation{repType: Required, create: `${oci_core_subnet.test_subnet.id}`},
		"endpoint_type":      Representation{repType: Optional, create: `VNIC`},
		"forwarding_address": Representation{repType: Optional, create: `10.0.0.5`},
		"scope":              Representation{repType: Required, create: `PRIVATE`},
		"nsg_ids":            Representation{repType: Optional, create: []string{`${oci_core_network_security_group.test_network_security_group.id}`}},
	}

	resolverEndpointRepresentationWithoutNsgId = representationCopyWithRemovedProperties(resolverEndpointRepresentation, []string{"nsg_ids"})

	ResolverEndpointResourceDependencies = ResolverResourceDependencies +
		generateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", Required, Create, subnetRepresentation) +
		generateResourceFromRepresentationMap("oci_core_network_security_group", "test_network_security_group", Required, Create, networkSecurityGroupRepresentation)
)

func TestDnsResolverEndpointResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDnsResolverEndpointResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_dns_resolver_endpoint.test_resolver_endpoint"

	datasourceName := "data.oci_dns_resolver_endpoints.test_resolver_endpoints"
	singularDatasourceName := "data.oci_dns_resolver_endpoint.test_resolver_endpoint"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// create dependencies
			{
				Config: config + compartmentIdVariableStr + ResolverEndpointResourceDependencies,
				Check: func(s *terraform.State) (err error) {
					log.Printf("[DEBUG] Wait for 2 minutes for oci_core_vcn resource to get created")
					time.Sleep(2 * time.Minute)
					return nil
				},
			},
			// verify create
			{
				Config: config + compartmentIdVariableStr + ResolverEndpointResourceDependencies +
					generateDataSourceFromRepresentationMap("oci_core_vcn_dns_resolver_association", "test_vcn_dns_resolver_association", Required, Create, vcnDnsResolverAssociationSingularDataSourceRepresentation) +
					generateResourceFromRepresentationMap("oci_dns_resolver", "test_resolver", Required, Create, resolverRepresentation) +
					generateResourceFromRepresentationMap("oci_dns_resolver_endpoint", "test_resolver_endpoint", Required, Create, resolverEndpointRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "is_forwarding", "true"),
					resource.TestCheckResourceAttr(resourceName, "is_listening", "false"),
					resource.TestCheckResourceAttr(resourceName, "name", "endpointName"),
					resource.TestCheckResourceAttrSet(resourceName, "resolver_id"),

					func(s *terraform.State) (err error) {
						_, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},
			// delete before next create
			{
				Config: config + compartmentIdVariableStr + ResolverEndpointResourceDependencies +
					generateDataSourceFromRepresentationMap("oci_core_vcn_dns_resolver_association", "test_vcn_dns_resolver_association", Required, Create, vcnDnsResolverAssociationSingularDataSourceRepresentation) +
					generateResourceFromRepresentationMap("oci_dns_resolver", "test_resolver", Required, Create, resolverRepresentation),
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + ResolverEndpointResourceDependencies +
					generateDataSourceFromRepresentationMap("oci_core_vcn_dns_resolver_association", "test_vcn_dns_resolver_association", Required, Create, vcnDnsResolverAssociationSingularDataSourceRepresentation) +
					generateResourceFromRepresentationMap("oci_dns_resolver", "test_resolver", Required, Create, resolverRepresentation) +
					generateResourceFromRepresentationMap("oci_dns_resolver_endpoint", "test_resolver_endpoint", Optional, Create, resolverEndpointRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
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
			// verify updates to updatable parameters
			{
				Config: config + compartmentIdVariableStr + ResolverEndpointResourceDependencies +
					generateDataSourceFromRepresentationMap("oci_core_vcn_dns_resolver_association", "test_vcn_dns_resolver_association", Required, Update, vcnDnsResolverAssociationSingularDataSourceRepresentation) +
					generateResourceFromRepresentationMap("oci_dns_resolver", "test_resolver", Required, Create, resolverRepresentation) +
					generateResourceFromRepresentationMap("oci_dns_resolver_endpoint", "test_resolver_endpoint", Optional, Update, resolverEndpointRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
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
				Config: config + compartmentIdVariableStr + ResolverEndpointResourceDependencies +
					generateDataSourceFromRepresentationMap("oci_core_vcn_dns_resolver_association", "test_vcn_dns_resolver_association", Required, Create, vcnDnsResolverAssociationSingularDataSourceRepresentation) +
					generateDataSourceFromRepresentationMap("oci_dns_resolver_endpoints", "test_resolver_endpoints", Optional, Update, resolverEndpointDataSourceRepresentation) +
					generateResourceFromRepresentationMap("oci_dns_resolver", "test_resolver", Required, Create, resolverRepresentation) +
					generateResourceFromRepresentationMap("oci_dns_resolver_endpoint", "test_resolver_endpoint", Optional, Update, resolverEndpointRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
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
					generateDataSourceFromRepresentationMap("oci_core_vcn_dns_resolver_association", "test_vcn_dns_resolver_association", Required, Create, vcnDnsResolverAssociationSingularDataSourceRepresentation) +
					generateDataSourceFromRepresentationMap("oci_dns_resolver_endpoint", "test_resolver_endpoint", Optional, Update, resolverEndpointSingularDataSourceRepresentation) +
					generateResourceFromRepresentationMap("oci_dns_resolver", "test_resolver", Required, Create, resolverRepresentation) +
					generateResourceFromRepresentationMap("oci_dns_resolver_endpoint", "test_resolver_endpoint", Optional, Update, resolverEndpointRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
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
					generateDataSourceFromRepresentationMap("oci_core_vcn_dns_resolver_association", "test_vcn_dns_resolver_association", Required, Create, vcnDnsResolverAssociationSingularDataSourceRepresentation) +
					generateResourceFromRepresentationMap("oci_dns_resolver", "test_resolver", Required, Create, resolverRepresentation) +
					generateResourceFromRepresentationMap("oci_dns_resolver_endpoint", "test_resolver_endpoint", Optional, Update, resolverEndpointRepresentation),
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
