// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	DnsResolverRequiredOnlyResource = DnsResolverResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_dns_resolver", "test_resolver", acctest.Required, acctest.Create, DnsResolverRepresentation)

	DnsResolverResourceConfig = DnsResolverResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_dns_resolver", "test_resolver", acctest.Optional, acctest.Update, DnsResolverRepresentation)

	DnsDnsResolverSingularDataSourceRepresentation = map[string]interface{}{
		"resolver_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_dns_resolver.test_resolver.id}`},
		"scope":       acctest.Representation{RepType: acctest.Optional, Create: `PRIVATE`},
	}

	DnsDnsResolverDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`},
		"id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_dns_resolver.test_resolver.id}`},
		"scope":          acctest.Representation{RepType: acctest.Optional, Create: `PRIVATE`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: DnsResolverDataSourceFilterRepresentation}}

	DnsResolverDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_dns_resolver.test_resolver.id}`}},
	}

	DnsResolverRepresentation = map[string]interface{}{
		"resolver_id":    acctest.Representation{RepType: acctest.Required, Create: `${data.oci_core_vcn_dns_resolver_association.test_vcn_dns_resolver_association.dns_resolver_id}`},
		"attached_views": acctest.RepresentationGroup{RepType: acctest.Optional, Group: DnsResolverAttachedViewsRepresentation},
		"defined_tags":   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"freeformTags": "freeformTags"}, Update: map[string]string{"freeformTags2": "freeformTags2"}},
		"scope":          acctest.Representation{RepType: acctest.Optional, Create: `PRIVATE`},
	}
	DnsResolverAttachedViewsRepresentation = map[string]interface{}{
		"view_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_dns_view.test_view.id}`},
	}

	resolverRepresentationRules = acctest.RepresentationCopyWithNewProperties(DnsResolverRepresentation, map[string]interface{}{
		"rules": []acctest.RepresentationGroup{{RepType: acctest.Optional, Group: resolverRulesItemsRepresentationClientAddr}, {RepType: acctest.Optional, Group: resolverRulesItemsRepresentationQname}},
	})

	resolverRulesItemsRepresentationClientAddr = map[string]interface{}{
		"action":                    acctest.Representation{RepType: acctest.Required, Create: `FORWARD`},
		"destination_addresses":     acctest.Representation{RepType: acctest.Required, Create: []string{`10.0.0.11`}, Update: []string{`10.0.0.12`}},
		"source_endpoint_name":      acctest.Representation{RepType: acctest.Required, Create: `endpointName`},
		"client_address_conditions": acctest.Representation{RepType: acctest.Optional, Create: []string{`192.0.20.0/24`}, Update: []string{`192.0.21.0/24`}},
		"qname_cover_conditions":    acctest.Representation{RepType: acctest.Optional, Update: []string{}},
	}
	resolverRulesItemsRepresentationQname = map[string]interface{}{
		"action":                    acctest.Representation{RepType: acctest.Required, Create: `FORWARD`},
		"destination_addresses":     acctest.Representation{RepType: acctest.Required, Create: []string{`10.0.0.11`}, Update: []string{`10.0.0.12`}},
		"source_endpoint_name":      acctest.Representation{RepType: acctest.Required, Create: `endpointName`},
		"client_address_conditions": acctest.Representation{RepType: acctest.Optional, Create: []string{}},
		"qname_cover_conditions":    acctest.Representation{RepType: acctest.Optional, Create: []string{`internal.example.com`}, Update: []string{`internal2.example.com`}},
	}

	DnsResolverResourceDependencies = DefinedTagsDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, CoreVcnRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_dns_view", "test_view", acctest.Required, acctest.Create, DnsViewRepresentation)
)

// issue-routing-tag: dns/default
func TestDnsResolverResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDnsResolverResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_dns_resolver.test_resolver"
	datasourceName := "data.oci_dns_resolvers.test_resolvers"
	singularDatasourceName := "data.oci_dns_resolver.test_resolver"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DnsResolverResourceDependencies+
		acctest.GenerateDataSourceFromRepresentationMap("oci_core_vcn_dns_resolver_association", "test_vcn_dns_resolver_association", acctest.Required, acctest.Create, CoreCoreVcnDnsResolverAssociationSingularDataSourceRepresentation)+
		acctest.GenerateResourceFromRepresentationMap("oci_dns_resolver", "test_resolver", acctest.Required, acctest.Create, DnsResolverRepresentation), "dns", "resolver", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// Create dependencies
		{
			Config: config + compartmentIdVariableStr + DnsResolverResourceDependencies,
			Check: func(s *terraform.State) (err error) {
				log.Printf("[DEBUG] Wait for 2 minutes for oci_core_vcn resource to get created")
				time.Sleep(2 * time.Minute)
				return nil
			},
		},
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DnsResolverResourceDependencies +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_vcn_dns_resolver_association", "test_vcn_dns_resolver_association", acctest.Required, acctest.Create, CoreCoreVcnDnsResolverAssociationSingularDataSourceRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_dns_resolver", "test_resolver", acctest.Required, acctest.Create, DnsResolverRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "resolver_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DnsResolverResourceDependencies +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_vcn_dns_resolver_association", "test_vcn_dns_resolver_association", acctest.Required, acctest.Create, CoreCoreVcnDnsResolverAssociationSingularDataSourceRepresentation),
		},
		// Create resolver with endpoint
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + DnsResolverResourceDependencies +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_vcn_dns_resolver_association", "test_vcn_dns_resolver_association", acctest.Required, acctest.Create, CoreCoreVcnDnsResolverAssociationSingularDataSourceRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, CoreSubnetRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_dns_resolver", "test_resolver", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(DnsResolverRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})) +
				acctest.GenerateResourceFromRepresentationMap("oci_dns_resolver_endpoint", "test_resolver_endpoint", acctest.Optional, acctest.Create, DnsResolverEndpointRepresentationWithoutNsgId),
		},
		// verify Create with optionals and resolver rules
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + DnsResolverResourceDependencies +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_vcn_dns_resolver_association", "test_vcn_dns_resolver_association", acctest.Required, acctest.Create, CoreCoreVcnDnsResolverAssociationSingularDataSourceRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, CoreSubnetRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_dns_resolver", "test_resolver", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(resolverRepresentationRules, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})) +
				acctest.GenerateResourceFromRepresentationMap("oci_dns_resolver_endpoint", "test_resolver_endpoint", acctest.Optional, acctest.Create, DnsResolverEndpointRepresentationWithoutNsgId),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "attached_views.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "attached_views.0.view_id"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "is_protected"),
				resource.TestCheckResourceAttrSet(resourceName, "resolver_id"),
				resource.TestCheckResourceAttrSet(resourceName, "self"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
				resource.TestCheckResourceAttr(resourceName, "endpoints.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "endpoints.0.forwarding_address", "10.0.0.5"),
				resource.TestCheckResourceAttr(resourceName, "endpoints.0.is_forwarding", "true"),
				resource.TestCheckResourceAttr(resourceName, "endpoints.0.is_listening", "false"),
				resource.TestCheckResourceAttr(resourceName, "rules.#", "2"),
				resource.TestCheckResourceAttr(resourceName, "rules.0.action", "FORWARD"),
				resource.TestCheckResourceAttr(resourceName, "rules.0.client_address_conditions.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "rules.0.client_address_conditions.0", "192.0.20.0/24"),
				resource.TestCheckResourceAttr(resourceName, "rules.0.destination_addresses.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "rules.0.destination_addresses.0", "10.0.0.11"),
				resource.TestCheckResourceAttr(resourceName, "rules.0.qname_cover_conditions.#", "0"),
				resource.TestCheckResourceAttr(resourceName, "rules.0.source_endpoint_name", "endpointName"),
				resource.TestCheckResourceAttr(resourceName, "rules.1.action", "FORWARD"),
				resource.TestCheckResourceAttr(resourceName, "rules.1.client_address_conditions.#", "0"),
				resource.TestCheckResourceAttr(resourceName, "rules.1.destination_addresses.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "rules.1.destination_addresses.0", "10.0.0.11"),
				resource.TestCheckResourceAttr(resourceName, "rules.1.qname_cover_conditions.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "rules.1.qname_cover_conditions.0", "internal.example.com"),
				resource.TestCheckResourceAttr(resourceName, "rules.1.source_endpoint_name", "endpointName"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					// Resource discovery is disabled for Resolver
					//if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
					//	if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
					//		return errExport
					//	}
					//}
					return err
				},
			),
		},
		// verify updates to updatable parameters and add resolver rules
		{
			Config: config + compartmentIdVariableStr + DnsResolverResourceDependencies +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_vcn_dns_resolver_association", "test_vcn_dns_resolver_association", acctest.Required, acctest.Create, CoreCoreVcnDnsResolverAssociationSingularDataSourceRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, CoreSubnetRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_dns_resolver", "test_resolver", acctest.Optional, acctest.Update,
					acctest.RepresentationCopyWithNewProperties(resolverRepresentationRules, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
					})) +
				acctest.GenerateResourceFromRepresentationMap("oci_dns_resolver_endpoint", "test_resolver_endpoint", acctest.Optional, acctest.Create, DnsResolverEndpointRepresentationWithoutNsgId),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "attached_views.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "attached_views.0.view_id"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "is_protected"),
				resource.TestCheckResourceAttrSet(resourceName, "resolver_id"),
				resource.TestCheckResourceAttrSet(resourceName, "self"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
				resource.TestCheckResourceAttr(resourceName, "endpoints.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "endpoints.0.forwarding_address", "10.0.0.5"),
				resource.TestCheckResourceAttr(resourceName, "endpoints.0.is_forwarding", "true"),
				resource.TestCheckResourceAttr(resourceName, "endpoints.0.is_listening", "false"),
				resource.TestCheckResourceAttr(resourceName, "rules.#", "2"),
				resource.TestCheckResourceAttr(resourceName, "rules.0.action", "FORWARD"),
				resource.TestCheckResourceAttr(resourceName, "rules.0.client_address_conditions.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "rules.0.client_address_conditions.0", "192.0.21.0/24"),
				resource.TestCheckResourceAttr(resourceName, "rules.0.destination_addresses.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "rules.0.destination_addresses.0", "10.0.0.12"),
				resource.TestCheckResourceAttr(resourceName, "rules.0.qname_cover_conditions.#", "0"),
				resource.TestCheckResourceAttr(resourceName, "rules.0.source_endpoint_name", "endpointName"),
				resource.TestCheckResourceAttr(resourceName, "rules.1.action", "FORWARD"),
				resource.TestCheckResourceAttr(resourceName, "rules.1.client_address_conditions.#", "0"),
				resource.TestCheckResourceAttr(resourceName, "rules.1.destination_addresses.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "rules.1.destination_addresses.0", "10.0.0.12"),
				resource.TestCheckResourceAttr(resourceName, "rules.1.qname_cover_conditions.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "rules.1.qname_cover_conditions.0", "internal2.example.com"),
				resource.TestCheckResourceAttr(resourceName, "rules.1.source_endpoint_name", "endpointName"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_vcn_dns_resolver_association", "test_vcn_dns_resolver_association", acctest.Required, acctest.Create, CoreCoreVcnDnsResolverAssociationSingularDataSourceRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_dns_resolvers", "test_resolvers", acctest.Optional, acctest.Update, DnsDnsResolverDataSourceRepresentation) +
				compartmentIdVariableStr + DnsResolverResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dns_resolver", "test_resolver", acctest.Optional, acctest.Update, DnsResolverRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(datasourceName, "resolvers.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "resolvers.0.attached_vcn_id"),
				resource.TestCheckResourceAttr(datasourceName, "resolvers.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "resolvers.0.default_view_id"),
				resource.TestCheckResourceAttr(datasourceName, "resolvers.0.display_name", "displayName"),
				resource.TestCheckResourceAttr(datasourceName, "resolvers.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "resolvers.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "resolvers.0.is_protected"),
				resource.TestCheckResourceAttrSet(datasourceName, "resolvers.0.self"),
				resource.TestCheckResourceAttrSet(datasourceName, "resolvers.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "resolvers.0.time_created"),
				resource.TestCheckResourceAttrSet(datasourceName, "resolvers.0.time_updated"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_vcn_dns_resolver_association", "test_vcn_dns_resolver_association", acctest.Required, acctest.Create, CoreCoreVcnDnsResolverAssociationSingularDataSourceRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_dns_resolver", "test_resolver", acctest.Required, acctest.Create, DnsDnsResolverSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DnsResolverResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "resolver_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "attached_vcn_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "attached_views.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "default_view_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_protected"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "self"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:            config + DnsResolverRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"scope",
			},
			ResourceName: resourceName,
		},
	})
}
