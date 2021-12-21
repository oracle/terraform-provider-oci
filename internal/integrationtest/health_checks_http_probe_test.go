// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	"github.com/terraform-providers/terraform-provider-oci/internal/resourcediscovery"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"
)

var (
	HttpProbeRequiredOnlyResource = HttpProbeResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_health_checks_http_probe", "test_http_probe", acctest.Required, acctest.Create, httpProbeRepresentation)

	httpProbeRepresentation = map[string]interface{}{
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"protocol":            acctest.Representation{RepType: acctest.Required, Create: `HTTP`},
		"targets":             acctest.Representation{RepType: acctest.Required, Create: []string{`www.oracle.com`}},
		"headers":             acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"headers": "headers"}, Update: map[string]string{"headers2": "headers2"}},
		"method":              acctest.Representation{RepType: acctest.Optional, Create: `GET`},
		"path":                acctest.Representation{RepType: acctest.Optional, Create: `/`},
		"port":                acctest.Representation{RepType: acctest.Optional, Create: `80`},
		"timeout_in_seconds":  acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"vantage_point_names": acctest.Representation{RepType: acctest.Optional, Create: []string{`goo-chs`}},
	}

	HttpProbeResourceDependencies = ""
)

// issue-routing-tag: health_checks/default
func TestHealthChecksHttpProbeResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestHealthChecksHttpProbeResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_health_checks_http_probe.test_http_probe"

	var resId string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+HttpProbeResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_health_checks_http_probe", "test_http_probe", acctest.Optional, acctest.Create, httpProbeRepresentation), "healthchecks", "httpProbe", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + HttpProbeResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_health_checks_http_probe", "test_http_probe", acctest.Required, acctest.Create, httpProbeRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "protocol", "HTTP"),
				resource.TestCheckResourceAttr(resourceName, "targets.#", "1"),
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + HttpProbeResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + HttpProbeResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_health_checks_http_probe", "test_http_probe", acctest.Optional, acctest.Create, httpProbeRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "headers.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "method", "GET"),
				resource.TestCheckResourceAttr(resourceName, "path", "/"),
				resource.TestCheckResourceAttr(resourceName, "port", "80"),
				resource.TestCheckResourceAttr(resourceName, "protocol", "HTTP"),
				resource.TestCheckResourceAttr(resourceName, "targets.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "timeout_in_seconds", "10"),
				resource.TestCheckResourceAttr(resourceName, "vantage_point_names.#", "1"),

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
	})
}
