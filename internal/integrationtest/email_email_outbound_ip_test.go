// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	EmailEmailOutboundIpDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
	}

	EmailEmailOutboundIpResourceConfig = ""
)

// issue-routing-tag: email/default
func TestEmailEmailOutboundIpResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestEmailEmailOutboundIpResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("tenancy_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	outboundIp := utils.GetEnvSettingWithBlankDefault("OCI_EMAIL_IP_3")
	datasourceName := "data.oci_email_email_outbound_ips.test_email_outbound_ips"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_email_email_outbound_ips", "test_email_outbound_ips", acctest.Required, acctest.Create, EmailEmailOutboundIpDataSourceRepresentation) +
				compartmentIdVariableStr + EmailEmailOutboundIpResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(

				resource.TestCheckResourceAttrSet(datasourceName, "email_outbound_ip_collection.#"),
				resource.TestCheckResourceAttr(datasourceName, "email_outbound_ip_collection.0.items.#", "3"),

				resource.TestCheckResourceAttr(datasourceName, "email_outbound_ip_collection.0.items.2.outbound_ip", outboundIp),
				resource.TestCheckResourceAttr(datasourceName, "email_outbound_ip_collection.0.items.2.state", "ACTIVE"),
			),
		},
	})
}
