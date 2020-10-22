// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	oci_apigateway "github.com/oracle/oci-go-sdk/v31/apigateway"
	"github.com/oracle/oci-go-sdk/v31/common"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	ApiGatewayCertificateRequiredOnlyResource = ApiGatewayCertificateResourceDependencies +
		generateResourceFromRepresentationMap("oci_apigateway_certificate", "test_certificate", Required, Create, apiGatewaycertificateRepresentation)

	CertificateResourceConfig = ApiGatewayCertificateResourceDependencies +
		generateResourceFromRepresentationMap("oci_apigateway_certificate", "test_certificate", Optional, Update, apiGatewaycertificateRepresentation)

	apiGatewaycertificateSingularDataSourceRepresentation = map[string]interface{}{
		"certificate_id": Representation{repType: Required, create: `${oci_apigateway_certificate.test_certificate.id}`},
	}

	apiGatewaycertificateDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"display_name":   Representation{repType: Optional, create: `displayName`, update: `displayName2`},
		"state":          Representation{repType: Optional, create: `ACTIVE`},
		"filter":         RepresentationGroup{Required, apiGatewaycertificateDataSourceFilterRepresentation}}
	apiGatewaycertificateDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_apigateway_certificate.test_certificate.id}`}},
	}

	apiGatewaycertificateRepresentation = map[string]interface{}{
		"certificate":               Representation{repType: Required, create: `-----BEGIN CERTIFICATE-----\nMIIFajCCA1KgAwIBAgIJAO/er1f/5rCIMA0GCSqGSIb3DQEBBQUAMFkxCzAJBgNV\nBAYTAlVTMRMwEQYDVQQIDApXYXNoaW5ndG9uMRAwDgYDVQQHDAdTZWF0dGxlMQ8w\nDQYDVQQKDAZPcmFjbGUxEjAQBgNVBAMMCXRlcnJhZm9ybTAeFw0yMDExMDMxNTA4\nMTBaFw0zMDExMDExNTA4MTBaMFkxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApXYXNo\naW5ndG9uMRAwDgYDVQQHDAdTZWF0dGxlMQ8wDQYDVQQKDAZPcmFjbGUxEjAQBgNV\nBAMMCXRlcnJhZm9ybTCCAiIwDQYJKoZIhvcNAQEBBQADggIPADCCAgoCggIBAMqU\nU9S8dDvQ1KsGwIyyoHO5MkkaUiRyc0Ts3gNb2yu410ruMFRod/FfmoSAlokcloQZ\nZntKXezVB6U6JrTIn1HfU7YCHFhcmMSjfo0kSD80g//kIWQDvCLwynE2XZktCh2k\nOtN4Vg4fptPnx0SpksT4QH8EbMXo0KQa/i5rTAqDJlyyVJGRzIVgAbnn2WgB6PHg\n7L29Hvr9nYFZwnJUMu6wlPQ0FLJ/Fl3L49ySfewEQB2gAAZpVXgEAqeugpCJY1nL\ndc91jX329WYKT6lfR0HR6vpVXrq4xp+kdGCJOdOQPi3ALHy3eN/XfE9kjHZ1cwCb\nyVjMX6rabMAnM7RrgLu+NNkDn5+eBzaeYfHlvmmLvurLi3LcK//7zOf2TKjJGaPa\nqLWB7zA4PTuJJEPwpZ0LGgcHYjEVBDTeHRTvblhEA8gFCYZLaIXnVjmN1WEOhRoE\nfojMAJ38T2q186wRrPRvBte8TOdDcetgh1d6otcWHRU7cJFhk+KKS1Ba991q9z9w\naNQ4RKumC9tPQtUfuXA1jcAq229WiqOyJM3ImZ0sc9qN5oIxvElU9sPx5DdSvjnd\n9jQRrwfZVi06dVTTI7FACgx8Xb5LaZyey36oZG8hgVw7JcdTsCg+idPVowgu2khV\n9wMsdTIK6fYT1tWmaRkK5CaSAdQHDAhq8Sf32yeHAgMBAAGjNTAzMDEGA1UdEQQq\nMCiCE2xvY2FsLm5ld2lzbS5jb20uYXWCEXd3dy5uZXdpc20uY29tLmF1MA0GCSqG\nSIb3DQEBBQUAA4ICAQBoYhrq4ULsxkqP9ATBbuKUBoJa5LzkjmOhaxhQAd7k/7mI\nhQjM6ZI5sZlGjCsPChthtgfVXIKOy0pk9IFbQDZfDv0jpMkFsKqbweYVcDx+m8rT\nfpoZde4kLham/fuui58fg8EJiYSC6pYG9VV6kw5rfJPCd4GfUX4+n0xFdY8pEsjT\notEL/mjnkeIdrM1pn5X18YFbaV41+3S0ogHBTrgqogbYe32FXwKmqlhuMjvOK06B\nUH7alavSV7cgO3WbX+tx9uFFhIiT7nV4ZppsoDxctHOmJ3JRc5MJLROA9TjJ+/mX\nJfnUvNcetdymW9RVXEqV55HSC2CXVdV+8aRXX1CENGIN0Mbj9FmEZaNbWea52aas\nDQDOOYot4YV21vSLWtciAw+F3xebn1agCu1yid6MJV8zrF+0hzE06a9aN63NkohV\nbfmxFop9z6zKKDawOzTnF5AbnZgDbIR5Hn/eGzca1Re7PWXGtSwIDdx2GXbej/Ko\nKLaX8hMT8ihBW+xCAT/QXSXo5iiBYj8aRlos6w9+7H4igBDCbyoPuXoQ2ab97HTd\nln1R+FqTIfltbiBlXfpor0kH2G2U1Hp7+rUkFD5ip/m7G30yHzFfKZbb3hfVzFkj\nxLaV+YHIYn2fNNNwbJaanlH1B9vk2Cyq2btNc+t+oEPa9JfWykMxBfJdtQK1YA==\n-----END CERTIFICATE-----`},
		"compartment_id":            Representation{repType: Required, create: `${var.compartment_id}`},
		"private_key":               Representation{repType: Required, create: `-----BEGIN RSA PRIVATE KEY-----\n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\n-----END RSA PRIVATE KEY-----`},
		"defined_tags":              Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":              Representation{repType: Optional, create: `displayName`, update: `displayName2`},
		"freeform_tags":             Representation{repType: Optional, create: map[string]string{"Department": "Finance"}, update: map[string]string{"Department": "Accounting"}},
		"intermediate_certificates": Representation{repType: Optional, create: `-----BEGIN CERTIFICATE-----\nMIIDajCCAlKgAwIBAgIJAPG9xit0XaTkMA0GCSqGSIb3DQEBBQUAMFkxCzAJBgNV\nBAYTAlVTMRMwEQYDVQQIDApXYXNoaW5ndG9uMRAwDgYDVQQHDAdTZWF0dGxlMQ8w\nDQYDVQQKDAZPcmFjbGUxEjAQBgNVBAMMCXRlcnJhZm9ybTAeFw0yMDExMDQxNDEy\nMDdaFw0zMDExMDIxNDEyMDdaMFkxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApXYXNo\naW5ndG9uMRAwDgYDVQQHDAdTZWF0dGxlMQ8wDQYDVQQKDAZPcmFjbGUxEjAQBgNV\nBAMMCXRlcnJhZm9ybTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBANqW\n9G055WH2n67KCA9pfyqlfwEATFDicGXoV2oWM07M2MYC/dMvpq6J/AAuB6029fFK\ns65nFT5I3sN9i+M4NZS+yPJ48IYBQ1ghW31dH6Ot0fiGep7TLHzCfq7cLK6gTyq/\n5GyzY4MOQCT/tv8ggbTjb2HndaRnbhyxur6pNn58ZS1hIDVoCcahSbddgPzo7wc9\n9F56WgAOhPyOKtIqCtLtdhimBvd1TdFaSDG9CqKgL1qXgSqJxB+Cv3kRL3vuCyT7\ncYW59iKkA0Yat6xPdEN67SQOybfOsZ/s0lcJTLvUuIWGdX2YZ8/Nc6u09Dc2BLJ+\npbA7dGjBNvnvvw+yGAMCAwEAAaM1MDMwMQYDVR0RBCowKIITbG9jYWwubmV3aXNt\nLmNvbS5hdYIRd3d3Lm5ld2lzbS5jb20uYXUwDQYJKoZIhvcNAQEFBQADggEBALdf\n0UenqIGhSVpi6XOUclwEG4AS0mDPWOl3DKrWJKzmt9BllL+BYd11XmIPJfDxwilf\n9SOO0Be7+zNUQcW7B9P+cZoTe8kN5M7sVmIBpzpBv2GYZYNvAJYqAJH0eqHyj4f2\n9ZrZYeMBJHR9pSdwvVaqCP4imDOaWmI6Yf5TS+sn66E3m9XypSlW60D6vCO76qgG\nd8WK7kcifNe/SDEROyQHr5x7BQS+V9Cega2RSisWCOxRnncwDtYbX69wq1pcJVb9\naZofrcDLJ4rIVIFMvQA7eCA40Vq3DZ7pnkMVI3qXwCXq3ZijbUOfyV59RB2e3xgP\nwr9WYo6QNA0+V+5GOcw=\n-----END CERTIFICATE-----`},
	}

	ApiGatewayCertificateResourceDependencies = DefinedTagsDependencies
)

func TestApigatewayCertificateResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestApigatewayCertificateResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := getEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_apigateway_certificate.test_certificate"
	datasourceName := "data.oci_apigateway_certificates.test_certificates"
	singularDatasourceName := "data.oci_apigateway_certificate.test_certificate"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckApigatewayCertificateDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + ApiGatewayCertificateResourceDependencies +
					generateResourceFromRepresentationMap("oci_apigateway_certificate", "test_certificate", Required, Create, apiGatewaycertificateRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "certificate", "-----BEGIN CERTIFICATE-----\nMIIFajCCA1KgAwIBAgIJAO/er1f/5rCIMA0GCSqGSIb3DQEBBQUAMFkxCzAJBgNV\nBAYTAlVTMRMwEQYDVQQIDApXYXNoaW5ndG9uMRAwDgYDVQQHDAdTZWF0dGxlMQ8w\nDQYDVQQKDAZPcmFjbGUxEjAQBgNVBAMMCXRlcnJhZm9ybTAeFw0yMDExMDMxNTA4\nMTBaFw0zMDExMDExNTA4MTBaMFkxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApXYXNo\naW5ndG9uMRAwDgYDVQQHDAdTZWF0dGxlMQ8wDQYDVQQKDAZPcmFjbGUxEjAQBgNV\nBAMMCXRlcnJhZm9ybTCCAiIwDQYJKoZIhvcNAQEBBQADggIPADCCAgoCggIBAMqU\nU9S8dDvQ1KsGwIyyoHO5MkkaUiRyc0Ts3gNb2yu410ruMFRod/FfmoSAlokcloQZ\nZntKXezVB6U6JrTIn1HfU7YCHFhcmMSjfo0kSD80g//kIWQDvCLwynE2XZktCh2k\nOtN4Vg4fptPnx0SpksT4QH8EbMXo0KQa/i5rTAqDJlyyVJGRzIVgAbnn2WgB6PHg\n7L29Hvr9nYFZwnJUMu6wlPQ0FLJ/Fl3L49ySfewEQB2gAAZpVXgEAqeugpCJY1nL\ndc91jX329WYKT6lfR0HR6vpVXrq4xp+kdGCJOdOQPi3ALHy3eN/XfE9kjHZ1cwCb\nyVjMX6rabMAnM7RrgLu+NNkDn5+eBzaeYfHlvmmLvurLi3LcK//7zOf2TKjJGaPa\nqLWB7zA4PTuJJEPwpZ0LGgcHYjEVBDTeHRTvblhEA8gFCYZLaIXnVjmN1WEOhRoE\nfojMAJ38T2q186wRrPRvBte8TOdDcetgh1d6otcWHRU7cJFhk+KKS1Ba991q9z9w\naNQ4RKumC9tPQtUfuXA1jcAq229WiqOyJM3ImZ0sc9qN5oIxvElU9sPx5DdSvjnd\n9jQRrwfZVi06dVTTI7FACgx8Xb5LaZyey36oZG8hgVw7JcdTsCg+idPVowgu2khV\n9wMsdTIK6fYT1tWmaRkK5CaSAdQHDAhq8Sf32yeHAgMBAAGjNTAzMDEGA1UdEQQq\nMCiCE2xvY2FsLm5ld2lzbS5jb20uYXWCEXd3dy5uZXdpc20uY29tLmF1MA0GCSqG\nSIb3DQEBBQUAA4ICAQBoYhrq4ULsxkqP9ATBbuKUBoJa5LzkjmOhaxhQAd7k/7mI\nhQjM6ZI5sZlGjCsPChthtgfVXIKOy0pk9IFbQDZfDv0jpMkFsKqbweYVcDx+m8rT\nfpoZde4kLham/fuui58fg8EJiYSC6pYG9VV6kw5rfJPCd4GfUX4+n0xFdY8pEsjT\notEL/mjnkeIdrM1pn5X18YFbaV41+3S0ogHBTrgqogbYe32FXwKmqlhuMjvOK06B\nUH7alavSV7cgO3WbX+tx9uFFhIiT7nV4ZppsoDxctHOmJ3JRc5MJLROA9TjJ+/mX\nJfnUvNcetdymW9RVXEqV55HSC2CXVdV+8aRXX1CENGIN0Mbj9FmEZaNbWea52aas\nDQDOOYot4YV21vSLWtciAw+F3xebn1agCu1yid6MJV8zrF+0hzE06a9aN63NkohV\nbfmxFop9z6zKKDawOzTnF5AbnZgDbIR5Hn/eGzca1Re7PWXGtSwIDdx2GXbej/Ko\nKLaX8hMT8ihBW+xCAT/QXSXo5iiBYj8aRlos6w9+7H4igBDCbyoPuXoQ2ab97HTd\nln1R+FqTIfltbiBlXfpor0kH2G2U1Hp7+rUkFD5ip/m7G30yHzFfKZbb3hfVzFkj\nxLaV+YHIYn2fNNNwbJaanlH1B9vk2Cyq2btNc+t+oEPa9JfWykMxBfJdtQK1YA==\n-----END CERTIFICATE-----"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "private_key", "-----BEGIN RSA PRIVATE KEY-----\n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\n-----END RSA PRIVATE KEY-----"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + ApiGatewayCertificateResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + ApiGatewayCertificateResourceDependencies +
					generateResourceFromRepresentationMap("oci_apigateway_certificate", "test_certificate", Optional, Create, apiGatewaycertificateRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "certificate"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "intermediate_certificates"),
					resource.TestCheckResourceAttrSet(resourceName, "private_key"),
					resource.TestCheckResourceAttrSet(resourceName, "subject_names.0"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
					resource.TestCheckResourceAttrSet(resourceName, "time_not_valid_after"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
							if errExport := testExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
								return errExport
							}
						}
						return err
					},
				),
			},

			// verify update to the compartment (the compartment will be switched back in the next step)
			{
				Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + ApiGatewayCertificateResourceDependencies +
					generateResourceFromRepresentationMap("oci_apigateway_certificate", "test_certificate", Optional, Create,
						representationCopyWithNewProperties(apiGatewaycertificateRepresentation, map[string]interface{}{
							"compartment_id": Representation{repType: Required, create: `${var.compartment_id_for_update}`},
						})),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "certificate"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "intermediate_certificates"),
					resource.TestCheckResourceAttrSet(resourceName, "private_key"),
					resource.TestCheckResourceAttrSet(resourceName, "subject_names.0"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
					resource.TestCheckResourceAttrSet(resourceName, "time_not_valid_after"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("resource recreated when it was supposed to be updated")
						}
						return err
					},
				),
			},

			// verify updates to updatable parameters
			{
				Config: config + compartmentIdVariableStr + ApiGatewayCertificateResourceDependencies +
					generateResourceFromRepresentationMap("oci_apigateway_certificate", "test_certificate", Optional, Update, apiGatewaycertificateRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "certificate"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "intermediate_certificates"),
					resource.TestCheckResourceAttrSet(resourceName, "private_key"),
					resource.TestCheckResourceAttrSet(resourceName, "subject_names.0"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
					resource.TestCheckResourceAttrSet(resourceName, "time_not_valid_after"),

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
				Config: config +
					generateDataSourceFromRepresentationMap("oci_apigateway_certificates", "test_certificates", Optional, Update, apiGatewaycertificateDataSourceRepresentation) +
					compartmentIdVariableStr + ApiGatewayCertificateResourceDependencies +
					generateResourceFromRepresentationMap("oci_apigateway_certificate", "test_certificate", Optional, Update, apiGatewaycertificateRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

					resource.TestCheckResourceAttr(datasourceName, "certificate_collection.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "certificate_collection.0.items.#", "1"),
				),
			},
			// verify singular datasource
			{
				Config: config +
					generateResourceFromRepresentationMap("oci_apigateway_certificate", "test_certificate", Optional, Update, apiGatewaycertificateRepresentation) +
					generateDataSourceFromRepresentationMap("oci_apigateway_certificate", "test_certificate", Required, Create, apiGatewaycertificateSingularDataSourceRepresentation) +
					compartmentIdVariableStr + ApiGatewayCertificateResourceDependencies,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "certificate_id"),
					resource.TestCheckResourceAttrSet(resourceName, "certificate"),
					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "intermediate_certificates"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "subject_names.0"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_not_valid_after"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				),
			},
			// remove singular datasource from previous step so that it doesn't conflict with import tests
			{
				Config: config + compartmentIdVariableStr + CertificateResourceConfig,
			},
			// verify resource import
			{
				Config:            config,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"private_key",
				},
				ResourceName: resourceName,
			},
		},
	})
}

func testAccCheckApigatewayCertificateDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).apiGatewayClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_apigateway_certificate" {
			noResourceFound = false
			request := oci_apigateway.GetCertificateRequest{}

			tmp := rs.Primary.ID
			request.CertificateId = &tmp

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "apigateway")

			response, err := client.GetCertificate(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_apigateway.CertificateLifecycleStateDeleted): true,
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
	if DependencyGraph == nil {
		initDependencyGraph()
	}
	if !inSweeperExcludeList("ApigatewayCertificate") {
		resource.AddTestSweepers("ApigatewayCertificate", &resource.Sweeper{
			Name:         "ApigatewayCertificate",
			Dependencies: DependencyGraph["certificate"],
			F:            sweepApigatewayCertificateResource,
		})
	}
}

func sweepApigatewayCertificateResource(compartment string) error {
	apiGatewayClient := GetTestClients(&schema.ResourceData{}).apiGatewayClient()
	certificateIds, err := getApiGatewayCertificateIds(compartment)
	if err != nil {
		return err
	}
	for _, certificateId := range certificateIds {
		if ok := SweeperDefaultResourceId[certificateId]; !ok {
			deleteCertificateRequest := oci_apigateway.DeleteCertificateRequest{}

			deleteCertificateRequest.CertificateId = &certificateId

			deleteCertificateRequest.RequestMetadata.RetryPolicy = getRetryPolicy(true, "apigateway")
			_, error := apiGatewayClient.DeleteCertificate(context.Background(), deleteCertificateRequest)
			if error != nil {
				fmt.Printf("Error deleting Certificate %s %s, It is possible that the resource is already deleted. Please verify manually \n", certificateId, error)
				continue
			}
			waitTillCondition(testAccProvider, &certificateId, apiGatewayCertificateSweepWaitCondition, time.Duration(3*time.Minute),
				apiGatewayCertificateSweepResponseFetchOperation, "apigateway", true)
		}
	}
	return nil
}

func getApiGatewayCertificateIds(compartment string) ([]string, error) {
	ids := getResourceIdsToSweep(compartment, "CertificateId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	apiGatewayClient := GetTestClients(&schema.ResourceData{}).apiGatewayClient()

	listCertificatesRequest := oci_apigateway.ListCertificatesRequest{}
	listCertificatesRequest.CompartmentId = &compartmentId
	listCertificatesRequest.LifecycleState = oci_apigateway.CertificateLifecycleStateActive
	listCertificatesResponse, err := apiGatewayClient.ListCertificates(context.Background(), listCertificatesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting Certificate list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, certificate := range listCertificatesResponse.Items {
		id := *certificate.Id
		resourceIds = append(resourceIds, id)
		addResourceIdToSweeperResourceIdMap(compartmentId, "CertificateId", id)
	}
	return resourceIds, nil
}

func apiGatewayCertificateSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if certificateResponse, ok := response.Response.(oci_apigateway.GetCertificateResponse); ok {
		return certificateResponse.LifecycleState != oci_apigateway.CertificateLifecycleStateDeleted
	}
	return false
}

func apiGatewayCertificateSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.apiGatewayClient().GetCertificate(context.Background(), oci_apigateway.GetCertificateRequest{
		CertificateId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
