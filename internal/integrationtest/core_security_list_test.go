// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v58/common"
	oci_core "github.com/oracle/oci-go-sdk/v58/core"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	tf_client "github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/resourcediscovery"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"
)

var (
	SecurityListRequiredOnlyResource = SecurityListResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_core_security_list", "test_security_list", acctest.Required, acctest.Create, securityListRepresentation)

	securityListDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `MyPrivateSubnetSecurityList`, Update: `displayName2`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `AVAILABLE`},
		"vcn_id":         acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_vcn.test_vcn.id}`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: securityListDataSourceFilterRepresentation}}
	securityListDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_core_security_list.test_security_list.id}`}},
	}

	securityListRepresentation = map[string]interface{}{
		"compartment_id":         acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"vcn_id":                 acctest.Representation{RepType: acctest.Required, Create: `${oci_core_vcn.test_vcn.id}`},
		"defined_tags":           acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":           acctest.Representation{RepType: acctest.Optional, Create: `MyPrivateSubnetSecurityList`, Update: `displayName2`},
		"egress_security_rules":  []acctest.RepresentationGroup{{RepType: acctest.Required, Group: securityListEgressSecurityRulesICMPRepresentation}, {RepType: acctest.Optional, Group: securityListEgressSecurityRulesTCPRepresentation}, {RepType: acctest.Optional, Group: securityListEgressSecurityRulesUDPRepresentation}},
		"freeform_tags":          acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"ingress_security_rules": []acctest.RepresentationGroup{{RepType: acctest.Required, Group: securityListIngressSecurityRulesICMPRepresentation}, {RepType: acctest.Optional, Group: securityListIngressSecurityRulesTCPRepresentation}, {RepType: acctest.Optional, Group: securityListIngressSecurityRulesUDPRepresentation}},
	}
	securityListEgressSecurityRulesICMPRepresentation = map[string]interface{}{
		"destination":      acctest.Representation{RepType: acctest.Required, Create: `10.0.2.0/24`, Update: `${lookup(data.oci_core_services.test_services.services[0], "cidr_block")}`},
		"protocol":         acctest.Representation{RepType: acctest.Required, Create: `1`},
		"description":      acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"destination_type": acctest.Representation{RepType: acctest.Optional, Create: `CIDR_BLOCK`, Update: `SERVICE_CIDR_BLOCK`},
		"icmp_options":     acctest.RepresentationGroup{RepType: acctest.Optional, Group: securityListEgressSecurityRulesIcmpOptionsRepresentation},
		"stateless":        acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
	}
	securityListEgressSecurityRulesTCPRepresentation = map[string]interface{}{
		"destination":      acctest.Representation{RepType: acctest.Required, Create: `10.0.2.0/24`, Update: `${lookup(data.oci_core_services.test_services.services[0], "cidr_block")}`},
		"protocol":         acctest.Representation{RepType: acctest.Required, Create: `6`},
		"destination_type": acctest.Representation{RepType: acctest.Optional, Create: `CIDR_BLOCK`, Update: `SERVICE_CIDR_BLOCK`},
		"stateless":        acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"tcp_options":      acctest.RepresentationGroup{RepType: acctest.Optional, Group: securityListEgressSecurityRulesTcpOptionsRepresentation},
	}
	securityListEgressSecurityRulesUDPRepresentation = map[string]interface{}{
		"destination":      acctest.Representation{RepType: acctest.Required, Create: `10.0.2.0/24`, Update: `${lookup(data.oci_core_services.test_services.services[0], "cidr_block")}`},
		"protocol":         acctest.Representation{RepType: acctest.Required, Create: `17`},
		"destination_type": acctest.Representation{RepType: acctest.Optional, Create: `CIDR_BLOCK`, Update: `SERVICE_CIDR_BLOCK`},
		"stateless":        acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"udp_options":      acctest.RepresentationGroup{RepType: acctest.Optional, Group: securityListEgressSecurityRulesUdpOptionsRepresentation},
	}
	securityListIngressSecurityRulesICMPRepresentation = map[string]interface{}{
		"protocol":     acctest.Representation{RepType: acctest.Required, Create: `1`},
		"description":  acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"source":       acctest.Representation{RepType: acctest.Required, Create: `10.0.1.0/24`, Update: `${lookup(data.oci_core_services.test_services.services[0], "cidr_block")}`},
		"icmp_options": acctest.RepresentationGroup{RepType: acctest.Optional, Group: securityListIngressSecurityRulesIcmpOptionsRepresentation},
		"source_type":  acctest.Representation{RepType: acctest.Optional, Create: `CIDR_BLOCK`, Update: `SERVICE_CIDR_BLOCK`},
		"stateless":    acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
	}
	securityListIngressSecurityRulesTCPRepresentation = map[string]interface{}{
		"protocol":    acctest.Representation{RepType: acctest.Required, Create: `6`},
		"source":      acctest.Representation{RepType: acctest.Required, Create: `10.0.1.0/24`, Update: `${lookup(data.oci_core_services.test_services.services[0], "cidr_block")}`},
		"source_type": acctest.Representation{RepType: acctest.Optional, Create: `CIDR_BLOCK`, Update: `SERVICE_CIDR_BLOCK`},
		"stateless":   acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"tcp_options": acctest.RepresentationGroup{RepType: acctest.Optional, Group: securityListIngressSecurityRulesTcpOptionsRepresentation},
	}
	securityListIngressSecurityRulesUDPRepresentation = map[string]interface{}{
		"protocol":    acctest.Representation{RepType: acctest.Required, Create: `17`},
		"source":      acctest.Representation{RepType: acctest.Required, Create: `10.0.1.0/24`, Update: `${lookup(data.oci_core_services.test_services.services[0], "cidr_block")}`},
		"source_type": acctest.Representation{RepType: acctest.Optional, Create: `CIDR_BLOCK`, Update: `SERVICE_CIDR_BLOCK`},
		"stateless":   acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"udp_options": acctest.RepresentationGroup{RepType: acctest.Optional, Group: securityListIngressSecurityRulesUdpOptionsRepresentation},
	}
	securityListEgressSecurityRulesIcmpOptionsRepresentation = map[string]interface{}{
		"type": acctest.Representation{RepType: acctest.Required, Create: `3`},
		"code": acctest.Representation{RepType: acctest.Optional, Create: `4`, Update: `0`},
	}
	securityListEgressSecurityRulesTcpOptionsRepresentation = map[string]interface{}{
		"max":               acctest.Representation{RepType: acctest.Optional, Create: `1521`, Update: `1522`},
		"min":               acctest.Representation{RepType: acctest.Optional, Create: `1521`, Update: `1522`},
		"source_port_range": acctest.RepresentationGroup{RepType: acctest.Optional, Group: securityListEgressSecurityRulesTcpOptionsSourcePortRangeRepresentation},
	}
	securityListEgressSecurityRulesUdpOptionsRepresentation = map[string]interface{}{
		"max":               acctest.Representation{RepType: acctest.Optional, Create: `1521`, Update: `1522`},
		"min":               acctest.Representation{RepType: acctest.Optional, Create: `1521`, Update: `1522`},
		"source_port_range": acctest.RepresentationGroup{RepType: acctest.Optional, Group: securityListEgressSecurityRulesUdpOptionsSourcePortRangeRepresentation},
	}
	securityListIngressSecurityRulesIcmpOptionsRepresentation = map[string]interface{}{
		"type": acctest.Representation{RepType: acctest.Required, Create: `3`},
		"code": acctest.Representation{RepType: acctest.Optional, Create: `4`, Update: `0`},
	}
	securityListIngressSecurityRulesTcpOptionsRepresentation = map[string]interface{}{
		"max":               acctest.Representation{RepType: acctest.Optional, Create: `1521`, Update: `1522`},
		"min":               acctest.Representation{RepType: acctest.Optional, Create: `1521`, Update: `1522`},
		"source_port_range": acctest.RepresentationGroup{RepType: acctest.Optional, Group: securityListIngressSecurityRulesTcpOptionsSourcePortRangeRepresentation},
	}
	securityListIngressSecurityRulesUdpOptionsRepresentation = map[string]interface{}{
		"max":               acctest.Representation{RepType: acctest.Optional, Create: `1521`, Update: `1522`},
		"min":               acctest.Representation{RepType: acctest.Optional, Create: `1521`, Update: `1522`},
		"source_port_range": acctest.RepresentationGroup{RepType: acctest.Optional, Group: securityListIngressSecurityRulesUdpOptionsSourcePortRangeRepresentation},
	}
	securityListEgressSecurityRulesTcpOptionsSourcePortRangeRepresentation = map[string]interface{}{
		"max": acctest.Representation{RepType: acctest.Required, Create: `1521`, Update: `1522`},
		"min": acctest.Representation{RepType: acctest.Required, Create: `1521`, Update: `1522`},
	}
	securityListEgressSecurityRulesUdpOptionsSourcePortRangeRepresentation = map[string]interface{}{
		"max": acctest.Representation{RepType: acctest.Required, Create: `1521`, Update: `1522`},
		"min": acctest.Representation{RepType: acctest.Required, Create: `1521`, Update: `1522`},
	}
	securityListIngressSecurityRulesTcpOptionsSourcePortRangeRepresentation = map[string]interface{}{
		"max": acctest.Representation{RepType: acctest.Required, Create: `1521`, Update: `1522`},
		"min": acctest.Representation{RepType: acctest.Required, Create: `1521`, Update: `1522`},
	}
	securityListIngressSecurityRulesUdpOptionsSourcePortRangeRepresentation = map[string]interface{}{
		"max": acctest.Representation{RepType: acctest.Required, Create: `1521`, Update: `1522`},
		"min": acctest.Representation{RepType: acctest.Required, Create: `1521`, Update: `1522`},
	}

	SecurityListResourceDependencies = acctest.GenerateDataSourceFromRepresentationMap("oci_core_services", "test_services", acctest.Required, acctest.Create, serviceDataSourceRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, vcnRepresentation) +
		DefinedTagsDependencies
)

// issue-routing-tag: core/virtualNetwork
func TestCoreSecurityListResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreSecurityListResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_core_security_list.test_security_list"
	datasourceName := "data.oci_core_security_lists.test_security_lists"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+SecurityListResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_core_security_list", "test_security_list", acctest.Optional, acctest.Create, securityListRepresentation), "core", "securityList", t)

	acctest.ResourceTest(t, testAccCheckCoreSecurityListDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + SecurityListResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_security_list", "test_security_list", acctest.Required, acctest.Create, securityListRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "egress_security_rules.#", "1"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "egress_security_rules", map[string]string{
					"destination": "10.0.2.0/24",
					"protocol":    "1",
				},
					[]string{}),
				resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.#", "1"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "ingress_security_rules", map[string]string{
					"protocol": "1",
					"source":   "10.0.1.0/24",
				},
					[]string{}),
				resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + SecurityListResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + SecurityListResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_security_list", "test_security_list", acctest.Optional, acctest.Create, securityListRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "MyPrivateSubnetSecurityList"),
				resource.TestCheckResourceAttr(resourceName, "egress_security_rules.#", "3"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "egress_security_rules", map[string]string{
					"destination":         "10.0.2.0/24",
					"description":         "description",
					"destination_type":    "CIDR_BLOCK",
					"icmp_options.#":      "1",
					"icmp_options.0.code": "4",
					"icmp_options.0.type": "3",
					"protocol":            "1",
					"stateless":           "false",
				},
					[]string{}),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "egress_security_rules", map[string]string{
					"destination":                           "10.0.2.0/24",
					"destination_type":                      "CIDR_BLOCK",
					"protocol":                              "6",
					"stateless":                             "false",
					"tcp_options.#":                         "1",
					"tcp_options.0.max":                     "1521",
					"tcp_options.0.min":                     "1521",
					"tcp_options.0.source_port_range.#":     "1",
					"tcp_options.0.source_port_range.0.max": "1521",
					"tcp_options.0.source_port_range.0.min": "1521",
				},
					[]string{}),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "egress_security_rules", map[string]string{
					"destination":                           "10.0.2.0/24",
					"destination_type":                      "CIDR_BLOCK",
					"protocol":                              "17",
					"stateless":                             "false",
					"udp_options.#":                         "1",
					"udp_options.0.max":                     "1521",
					"udp_options.0.min":                     "1521",
					"udp_options.0.source_port_range.#":     "1",
					"udp_options.0.source_port_range.0.max": "1521",
					"udp_options.0.source_port_range.0.min": "1521",
				},
					[]string{}),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.#", "3"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "ingress_security_rules", map[string]string{
					"icmp_options.#":      "1",
					"icmp_options.0.code": "4",
					"icmp_options.0.type": "3",
					"description":         "description",
					"protocol":            "1",
					"source":              "10.0.1.0/24",
					"source_type":         "CIDR_BLOCK",
					"stateless":           "false",
				},
					[]string{}),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "ingress_security_rules", map[string]string{
					"protocol":                              "6",
					"source":                                "10.0.1.0/24",
					"source_type":                           "CIDR_BLOCK",
					"stateless":                             "false",
					"tcp_options.#":                         "1",
					"tcp_options.0.max":                     "1521",
					"tcp_options.0.min":                     "1521",
					"tcp_options.0.source_port_range.#":     "1",
					"tcp_options.0.source_port_range.0.max": "1521",
					"tcp_options.0.source_port_range.0.min": "1521",
				},
					[]string{}),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "ingress_security_rules", map[string]string{
					"protocol":                              "17",
					"source":                                "10.0.1.0/24",
					"source_type":                           "CIDR_BLOCK",
					"stateless":                             "false",
					"udp_options.#":                         "1",
					"udp_options.0.max":                     "1521",
					"udp_options.0.min":                     "1521",
					"udp_options.0.source_port_range.#":     "1",
					"udp_options.0.source_port_range.0.max": "1521",
					"udp_options.0.source_port_range.0.min": "1521",
				},
					[]string{}),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),

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

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + SecurityListResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_security_list", "test_security_list", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(securityListRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "display_name", "MyPrivateSubnetSecurityList"),
				resource.TestCheckResourceAttr(resourceName, "egress_security_rules.#", "3"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "egress_security_rules", map[string]string{
					"description":         "description",
					"destination":         "10.0.2.0/24",
					"destination_type":    "CIDR_BLOCK",
					"icmp_options.#":      "1",
					"icmp_options.0.code": "4",
					"icmp_options.0.type": "3",
					"protocol":            "1",
					"stateless":           "false",
				},
					[]string{}),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.#", "3"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "ingress_security_rules", map[string]string{
					"description":         "description",
					"icmp_options.#":      "1",
					"icmp_options.0.code": "4",
					"icmp_options.0.type": "3",
					"protocol":            "1",
					"source":              "10.0.1.0/24",
					"source_type":         "CIDR_BLOCK",
					"stateless":           "false",
				},
					[]string{}),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),

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
			Config: config + compartmentIdVariableStr + SecurityListResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_security_list", "test_security_list", acctest.Optional, acctest.Update, securityListRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "egress_security_rules.#", "3"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "egress_security_rules", map[string]string{
					"description":         "description2",
					"destination_type":    "SERVICE_CIDR_BLOCK",
					"icmp_options.#":      "1",
					"icmp_options.0.code": "0",
					"icmp_options.0.type": "3",
					"protocol":            "1",
					"stateless":           "true",
				},
					[]string{
						"destination",
					}),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "egress_security_rules", map[string]string{
					"destination_type":                      "SERVICE_CIDR_BLOCK",
					"protocol":                              "6",
					"stateless":                             "true",
					"tcp_options.#":                         "1",
					"tcp_options.0.max":                     "1522",
					"tcp_options.0.min":                     "1522",
					"tcp_options.0.source_port_range.#":     "1",
					"tcp_options.0.source_port_range.0.max": "1522",
					"tcp_options.0.source_port_range.0.min": "1522",
				},
					[]string{
						"destination",
					}),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "egress_security_rules", map[string]string{
					"destination_type":                      "SERVICE_CIDR_BLOCK",
					"protocol":                              "17",
					"stateless":                             "true",
					"udp_options.#":                         "1",
					"udp_options.0.max":                     "1522",
					"udp_options.0.min":                     "1522",
					"udp_options.0.source_port_range.#":     "1",
					"udp_options.0.source_port_range.0.max": "1522",
					"udp_options.0.source_port_range.0.min": "1522",
				},
					[]string{
						"destination",
					}),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.#", "3"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "ingress_security_rules", map[string]string{
					"description":         "description2",
					"icmp_options.#":      "1",
					"icmp_options.0.code": "0",
					"icmp_options.0.type": "3",
					"protocol":            "1",
					"source_type":         "SERVICE_CIDR_BLOCK",
					"stateless":           "true",
				},
					[]string{
						"source",
					}),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "ingress_security_rules", map[string]string{
					"protocol":                              "6",
					"source_type":                           "SERVICE_CIDR_BLOCK",
					"stateless":                             "true",
					"tcp_options.#":                         "1",
					"tcp_options.0.max":                     "1522",
					"tcp_options.0.min":                     "1522",
					"tcp_options.0.source_port_range.#":     "1",
					"tcp_options.0.source_port_range.0.max": "1522",
					"tcp_options.0.source_port_range.0.min": "1522",
				},
					[]string{
						"source",
					}),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "ingress_security_rules", map[string]string{
					"protocol":                              "17",
					"source_type":                           "SERVICE_CIDR_BLOCK",
					"stateless":                             "true",
					"udp_options.#":                         "1",
					"udp_options.0.max":                     "1522",
					"udp_options.0.min":                     "1522",
					"udp_options.0.source_port_range.#":     "1",
					"udp_options.0.source_port_range.0.max": "1522",
					"udp_options.0.source_port_range.0.min": "1522",
				},
					[]string{
						"source",
					}),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_security_lists", "test_security_lists", acctest.Optional, acctest.Update, securityListDataSourceRepresentation) +
				compartmentIdVariableStr + SecurityListResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_security_list", "test_security_list", acctest.Optional, acctest.Update, securityListRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),
				resource.TestCheckResourceAttrSet(datasourceName, "vcn_id"),

				resource.TestCheckResourceAttr(datasourceName, "security_lists.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "security_lists.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "security_lists.0.display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "security_lists.0.egress_security_rules.#", "3"),
				acctest.CheckResourceSetContainsElementWithProperties(datasourceName, "security_lists.0.egress_security_rules", map[string]string{
					"description":         "description2",
					"destination_type":    "SERVICE_CIDR_BLOCK",
					"icmp_options.#":      "1",
					"icmp_options.0.code": "0",
					"icmp_options.0.type": "3",
					"protocol":            "1",
					"stateless":           "true",
				},
					[]string{
						"destination",
					}),
				acctest.CheckResourceSetContainsElementWithProperties(datasourceName, "security_lists.0.egress_security_rules", map[string]string{
					"destination_type":                      "SERVICE_CIDR_BLOCK",
					"protocol":                              "6",
					"stateless":                             "true",
					"tcp_options.#":                         "1",
					"tcp_options.0.max":                     "1522",
					"tcp_options.0.min":                     "1522",
					"tcp_options.0.source_port_range.#":     "1",
					"tcp_options.0.source_port_range.0.max": "1522",
					"tcp_options.0.source_port_range.0.min": "1522",
				},
					[]string{
						"destination",
					}),
				acctest.CheckResourceSetContainsElementWithProperties(datasourceName, "security_lists.0.egress_security_rules", map[string]string{
					"destination_type":                      "SERVICE_CIDR_BLOCK",
					"protocol":                              "17",
					"stateless":                             "true",
					"udp_options.#":                         "1",
					"udp_options.0.max":                     "1522",
					"udp_options.0.min":                     "1522",
					"udp_options.0.source_port_range.#":     "1",
					"udp_options.0.source_port_range.0.max": "1522",
					"udp_options.0.source_port_range.0.min": "1522",
				},
					[]string{
						"destination",
					}),
				resource.TestCheckResourceAttr(datasourceName, "security_lists.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "security_lists.0.id"),
				resource.TestCheckResourceAttr(datasourceName, "security_lists.0.ingress_security_rules.#", "3"),
				acctest.CheckResourceSetContainsElementWithProperties(datasourceName, "security_lists.0.ingress_security_rules", map[string]string{
					"description":         "description2",
					"icmp_options.#":      "1",
					"icmp_options.0.code": "0",
					"icmp_options.0.type": "3",
					"protocol":            "1",
					"source_type":         "SERVICE_CIDR_BLOCK",
					"stateless":           "true",
				},
					[]string{
						"source",
					}),
				acctest.CheckResourceSetContainsElementWithProperties(datasourceName, "security_lists.0.ingress_security_rules", map[string]string{
					"protocol":                              "6",
					"source_type":                           "SERVICE_CIDR_BLOCK",
					"stateless":                             "true",
					"tcp_options.#":                         "1",
					"tcp_options.0.max":                     "1522",
					"tcp_options.0.min":                     "1522",
					"tcp_options.0.source_port_range.#":     "1",
					"tcp_options.0.source_port_range.0.max": "1522",
					"tcp_options.0.source_port_range.0.min": "1522",
				},
					[]string{
						"source",
					}),
				acctest.CheckResourceSetContainsElementWithProperties(datasourceName, "security_lists.0.ingress_security_rules", map[string]string{
					"protocol":                              "17",
					"source_type":                           "SERVICE_CIDR_BLOCK",
					"stateless":                             "true",
					"udp_options.#":                         "1",
					"udp_options.0.max":                     "1522",
					"udp_options.0.min":                     "1522",
					"udp_options.0.source_port_range.#":     "1",
					"udp_options.0.source_port_range.0.max": "1522",
					"udp_options.0.source_port_range.0.min": "1522",
				},
					[]string{
						"source",
					}),
				resource.TestCheckResourceAttrSet(datasourceName, "security_lists.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "security_lists.0.time_created"),
				resource.TestCheckResourceAttrSet(datasourceName, "security_lists.0.vcn_id"),
			),
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

func testAccCheckCoreSecurityListDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).VirtualNetworkClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_core_security_list" {
			noResourceFound = false
			request := oci_core.GetSecurityListRequest{}

			tmp := rs.Primary.ID
			request.SecurityListId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "core")

			response, err := client.GetSecurityList(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_core.SecurityListLifecycleStateTerminated): true,
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
	if !acctest.InSweeperExcludeList("CoreSecurityList") {
		resource.AddTestSweepers("CoreSecurityList", &resource.Sweeper{
			Name:         "CoreSecurityList",
			Dependencies: acctest.DependencyGraph["securityList"],
			F:            sweepCoreSecurityListResource,
		})
	}
}

func sweepCoreSecurityListResource(compartment string) error {
	virtualNetworkClient := acctest.GetTestClients(&schema.ResourceData{}).VirtualNetworkClient()
	securityListIds, err := getSecurityListIds(compartment)
	if err != nil {
		return err
	}
	for _, securityListId := range securityListIds {
		if ok := acctest.SweeperDefaultResourceId[securityListId]; !ok {
			deleteSecurityListRequest := oci_core.DeleteSecurityListRequest{}

			deleteSecurityListRequest.SecurityListId = &securityListId

			deleteSecurityListRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "core")
			_, error := virtualNetworkClient.DeleteSecurityList(context.Background(), deleteSecurityListRequest)
			if error != nil {
				fmt.Printf("Error deleting SecurityList %s %s, It is possible that the resource is already deleted. Please verify manually \n", securityListId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &securityListId, securityListSweepWaitCondition, time.Duration(3*time.Minute),
				securityListSweepResponseFetchOperation, "core", true)
		}
	}
	return nil
}

func getSecurityListIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "SecurityListId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	virtualNetworkClient := acctest.GetTestClients(&schema.ResourceData{}).VirtualNetworkClient()

	listSecurityListsRequest := oci_core.ListSecurityListsRequest{}
	listSecurityListsRequest.CompartmentId = &compartmentId
	listSecurityListsRequest.LifecycleState = oci_core.SecurityListLifecycleStateAvailable
	listSecurityListsResponse, err := virtualNetworkClient.ListSecurityLists(context.Background(), listSecurityListsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting SecurityList list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, securityList := range listSecurityListsResponse.Items {
		id := *securityList.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "SecurityListId", id)
	}
	return resourceIds, nil
}

func securityListSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if securityListResponse, ok := response.Response.(oci_core.GetSecurityListResponse); ok {
		return securityListResponse.LifecycleState != oci_core.SecurityListLifecycleStateTerminated
	}
	return false
}

func securityListSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.VirtualNetworkClient().GetSecurityList(context.Background(), oci_core.GetSecurityListRequest{
		SecurityListId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
