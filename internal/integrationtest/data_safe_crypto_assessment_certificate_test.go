// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DataSafeCryptoAssessmentCertificateDataSourceRepresentation = map[string]interface{}{
		"compartment_id":   acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"assessment_id":    acctest.Representation{RepType: acctest.Optional, Create: `${var.crypto_assessment_id}`},
		"certificate_type": acctest.Representation{RepType: acctest.Optional, Create: []string{`SERVER`}},
	}

	DataSafeCryptoAssessmentTrustedCertificateDataSourceRepresentation = map[string]interface{}{
		"compartment_id":   acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"assessment_id":    acctest.Representation{RepType: acctest.Optional, Create: `${var.crypto_assessment_id}`},
		"certificate_type": acctest.Representation{RepType: acctest.Optional, Create: []string{`TRUSTED`}},
	}

	DataSafeCryptoAssessmentUserCertificateDataSourceRepresentation = map[string]interface{}{
		"compartment_id":   acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"assessment_id":    acctest.Representation{RepType: acctest.Optional, Create: `${var.crypto_assessment_id}`},
		"certificate_type": acctest.Representation{RepType: acctest.Optional, Create: []string{`USER`}},
	}

	DataSafeCryptoAssessmentCertificateResourceConfig = ""
)

func testCheckCertificateDetailsIfPresent(datasourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[datasourceName]
		if !ok {
			return fmt.Errorf("not found: %s", datasourceName)
		}

		itemsCountStr := rs.Primary.Attributes["crypto_assessment_certificate_collection.0.items.#"]
		if itemsCountStr == "" {
			return fmt.Errorf("missing item count for %s", datasourceName)
		}

		itemsCount, err := strconv.Atoi(itemsCountStr)
		if err != nil {
			return err
		}

		if itemsCount == 0 {
			return nil
		}

		if rs.Primary.Attributes["crypto_assessment_certificate_collection.0.items.0.assessment_id"] == "" {
			return fmt.Errorf("expected assessment_id to be set")
		}
		if rs.Primary.Attributes["crypto_assessment_certificate_collection.0.items.0.certificate_type"] == "" {
			return fmt.Errorf("expected certificate_type to be set")
		}
		if rs.Primary.Attributes["crypto_assessment_certificate_collection.0.items.0.compartment_id"] == "" {
			return fmt.Errorf("expected compartment_id to be set")
		}
		if rs.Primary.Attributes["crypto_assessment_certificate_collection.0.items.0.time_valid_from"] == "" {
			return fmt.Errorf("expected time_valid_from to be set")
		}
		if rs.Primary.Attributes["crypto_assessment_certificate_collection.0.items.0.time_valid_until"] == "" {
			return fmt.Errorf("expected time_valid_until to be set")
		}

		return nil
	}
}

// issue-routing-tag: data_safe/default
func TestDataSafeCryptoAssessmentCertificateResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataSafeCryptoAssessmentCertificateResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	cryptoAssessmentId := utils.GetEnvSettingWithBlankDefault("crypto_assessment_id")
	cryptoAssessmentIdVariableStr := fmt.Sprintf("variable \"crypto_assessment_id\" { default = \"%s\" }\n", cryptoAssessmentId)

	datasourceName := "data.oci_data_safe_crypto_assessment_certificates.test_crypto_assessment_certificates"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource for SERVER certificates
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_crypto_assessment_certificates", "test_crypto_assessment_certificates", acctest.Required, acctest.Create, DataSafeCryptoAssessmentCertificateDataSourceRepresentation) +
				compartmentIdVariableStr + cryptoAssessmentIdVariableStr + DataSafeCryptoAssessmentCertificateResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "crypto_assessment_certificate_collection.#"),
				testCheckCertificateDetailsIfPresent(datasourceName),
			),
		},
		// verify datasource for USER certificates
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_crypto_assessment_certificates", "test_crypto_assessment_certificates", acctest.Required, acctest.Create, DataSafeCryptoAssessmentUserCertificateDataSourceRepresentation) +
				compartmentIdVariableStr + cryptoAssessmentIdVariableStr + DataSafeCryptoAssessmentCertificateResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "crypto_assessment_certificate_collection.#"),
				testCheckCertificateDetailsIfPresent(datasourceName),
			),
		},
		// verify datasource for TRUSTED certificates
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_crypto_assessment_certificates", "test_crypto_assessment_certificates", acctest.Required, acctest.Create, DataSafeCryptoAssessmentTrustedCertificateDataSourceRepresentation) +
				compartmentIdVariableStr + cryptoAssessmentIdVariableStr + DataSafeCryptoAssessmentCertificateResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "crypto_assessment_certificate_collection.#"),
				testCheckCertificateDetailsIfPresent(datasourceName),
			),
		},
	})
}
