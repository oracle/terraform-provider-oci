// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	PingProbeRequiredOnlyResource = PingProbeResourceDependencies +
		generateResourceFromRepresentationMap("oci_health_checks_ping_probe", "test_ping_probe", Required, Create, pingProbeRepresentation)

	pingProbeRepresentation = map[string]interface{}{
		"compartment_id":      Representation{repType: Required, create: `${var.compartment_id}`},
		"protocol":            Representation{repType: Required, create: `TCP`},
		"targets":             Representation{repType: Required, create: []string{`www.oracle.com`}},
		"port":                Representation{repType: Optional, create: `80`},
		"timeout_in_seconds":  Representation{repType: Optional, create: `10`},
		"vantage_point_names": Representation{repType: Optional, create: []string{`goo-chs`}},
	}

	PingProbeResourceDependencies = ""
)

func TestHealthChecksPingProbeResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestHealthChecksPingProbeResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_health_checks_ping_probe.test_ping_probe"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + PingProbeResourceDependencies +
					generateResourceFromRepresentationMap("oci_health_checks_ping_probe", "test_ping_probe", Required, Create, pingProbeRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "protocol", "TCP"),
					resource.TestCheckResourceAttr(resourceName, "targets.#", "1"),
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + PingProbeResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + PingProbeResourceDependencies +
					generateResourceFromRepresentationMap("oci_health_checks_ping_probe", "test_ping_probe", Optional, Create, pingProbeRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "port", "80"),
					resource.TestCheckResourceAttr(resourceName, "protocol", "TCP"),
					resource.TestCheckResourceAttr(resourceName, "targets.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "timeout_in_seconds", "10"),
					resource.TestCheckResourceAttr(resourceName, "vantage_point_names.#", "1"),
				),
			},
		},
	})
}
