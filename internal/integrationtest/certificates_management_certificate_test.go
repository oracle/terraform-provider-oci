// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	requiredCertName         = "test-required-cert-" + utils.RandomString(10, utils.CharsetWithoutDigits)
	certNameForOptionalTests = "test-optional-cert-" + utils.RandomString(10, utils.CharsetWithoutDigits)

	certNotBeforeCreate = time.Now().UTC().AddDate(0, 0, 1).Truncate(time.Millisecond).Format(time.RFC3339Nano)
	certNotBeforeUpdate = time.Now().UTC().AddDate(0, 0, 7).Truncate(time.Millisecond).Format(time.RFC3339Nano)
	certNotAfterCreate  = time.Now().UTC().AddDate(1, 0, 1).Truncate(time.Millisecond).Format(time.RFC3339Nano)
	certNotAfterUpdate  = time.Now().UTC().AddDate(1, 0, 7).Truncate(time.Millisecond).Format(time.RFC3339Nano)

	csrPem = "-----BEGIN CERTIFICATE REQUEST-----\nMIICzzCCAbcCAQAwgYkxCzAJBgNVBAYTAnVzMRMwEQYDVQQIDApXYXNoaW5ndG9u\nMREwDwYDVQQHDAhCZWxsZXZ1ZTEMMAoGA1UECgwDT0NJMQwwCgYDVQQLDANTRUMx\nGDAWBgNVBAMMD3d3dy50ZXN0Y3NyLmNvbTEcMBoGCSqGSIb3DQEJARYNdGVzdEB0\nZXN0LmNvbTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBALrlpiuehIMz\nre0pa+G8WjPaltQOKsA5LZWFUmnkRvy9zGtVKXmcZ70m8+SIIW0svwhDENtvl6wA\nPfQX8G+9XelIzyUv62EG9PwcSqW7rLpq8VygH/6fbYLU/LG8wuIjVq8yU8AQSvcJ\nBSNEiPMxN7qmXviKXHyY8A3ramgY+2wIfqNPRfrM/udPGd8xls/MviqCXLl3rT5c\nLDwCOi+rqOgWwt9wkOzAVNQWtwCaGaed1j98T3QQtYjeSVeV/HaGUquK9nare0dF\n49SEgRuCpVSeD/PHN5lz1YE99bioEiqfdATw05xZ08wTezgscCKfuaMMAXDF2Q0N\nqMNaynCFuuECAwEAAaAAMA0GCSqGSIb3DQEBCwUAA4IBAQBXsLUdghNHVHM54DRJ\nwNJdICec9HX5OuwHP1C/QpwwcEeO4lIKTxww4WWwReJVIoUewgGaxYLh+izp40bP\no+wFR0NDC9eBaRtnBZiPYn8bVA2PBl9z0VS2+AsQHR9hlKH0G3iU+C0O2wVSJpQW\nAIU3c/2CNg9GCTWmJE+Jd8dTd21WDlARBKw7GXwToHQL7J3vY+2/S2e6hdbxh9aG\n7ZEOYssclMtVZxKPdGp9l3JbwBrk/9hv3kxFZvbtxjnn3ah5Sas+YwIewn2VPSwY\nfoewMRgyYZKUFmw/T3YH2ony9ouT8mgPMJys72iVXO+Ri31VdL7AVQHMA0uQEpKM\n+Pn4\n-----END CERTIFICATE REQUEST-----"

	CertificatesManagementCertificateRequiredOnlyResource = acctest.GenerateResourceFromRepresentationMap("oci_certificates_management_certificate", "test_certificate", acctest.Required, acctest.Create, certificatesManagementCertificateRepresentation)

	CertificatesManagementCertificateResourceConfig = CertificatesManagementCertificateResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_certificates_management_certificate", "test_certificate", acctest.Optional, acctest.Update, certificatesManagementCertificateRepresentation)

	CertificatesManagementcertificateSingularDataSourceRepresentation = map[string]interface{}{
		"certificate_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_certificates_management_certificate.test_certificate.id}`},
	}

	certificatesManagementCertificateDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"name":           acctest.Representation{RepType: acctest.Optional, Create: certNameForOptionalTests},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: certificatesManagementCertificateDataSourceFilterRepresentation}}
	certificatesManagementCertificateDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_certificates_management_certificate.test_certificate.id}`}},
	}

	certificatesManagementCertificateRepresentation = map[string]interface{}{
		"certificate_config": acctest.RepresentationGroup{RepType: acctest.Required, Group: certificateCertificateConfigRepresentationInternal},
		"compartment_id":     acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"name":               acctest.Representation{RepType: acctest.Required, Create: certNameForOptionalTests},
		"certificate_rules":  acctest.RepresentationGroup{RepType: acctest.Optional, Group: certificateCertificateRulesRepresentation},
		"defined_tags":       acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":        acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"freeform_tags":      acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
	}
	certificateCertificateConfigRepresentation = map[string]interface{}{
		"config_type":                     acctest.Representation{RepType: acctest.Required, Create: `ISSUED_BY_INTERNAL_CA`},
		"cert_chain_pem":                  acctest.Representation{RepType: acctest.Optional, Create: `certChainPem`, Update: `certChainPem2`},
		"certificate_pem":                 acctest.Representation{RepType: acctest.Optional, Create: `certificatePem`, Update: `certificatePem2`},
		"certificate_profile_type":        acctest.Representation{RepType: acctest.Optional, Create: `TLS_SERVER_OR_CLIENT`},
		"csr_pem":                         acctest.Representation{RepType: acctest.Optional, Create: `csrPem`, Update: `csrPem2`},
		"issuer_certificate_authority_id": acctest.Representation{RepType: acctest.Optional, Create: issuerCaId},
		"key_algorithm":                   acctest.Representation{RepType: acctest.Optional, Create: `RSA2048`},
		"private_key_pem":                 acctest.Representation{RepType: acctest.Optional, Create: `privateKeyPem`, Update: `privateKeyPem2`},
		"private_key_pem_passphrase":      acctest.Representation{RepType: acctest.Optional, Create: `privateKeyPemPassphrase`, Update: `privateKeyPemPassphrase2`},
		"signature_algorithm":             acctest.Representation{RepType: acctest.Optional, Create: `SHA256_WITH_RSA`},
		"subject":                         acctest.RepresentationGroup{RepType: acctest.Optional, Group: certificatesManagementCertificateSubjectRepresentation},
		"subject_alternative_names":       acctest.RepresentationGroup{RepType: acctest.Optional, Group: certificateCertificateConfigSubjectAlternativeNamesRepresentation},
		"validity":                        acctest.RepresentationGroup{RepType: acctest.Optional, Group: certificateCertificateConfigValidityRepresentation},
		"version_name":                    acctest.Representation{RepType: acctest.Optional, Create: `versionName`, Update: `versionName2`},
		"lifecycle":                       acctest.RepresentationGroup{RepType: acctest.Required, Group: certificatesManagementIgnoreChangesRepresentation},
	}

	// Internal managed cert config
	certificateCertificateConfigRepresentationInternal = map[string]interface{}{
		"config_type":                     acctest.Representation{RepType: acctest.Required, Create: `ISSUED_BY_INTERNAL_CA`, Update: `ISSUED_BY_INTERNAL_CA`},
		"certificate_profile_type":        acctest.Representation{RepType: acctest.Required, Create: `TLS_SERVER_OR_CLIENT`},
		"issuer_certificate_authority_id": acctest.Representation{RepType: acctest.Required, Create: issuerCaId},
		"key_algorithm":                   acctest.Representation{RepType: acctest.Optional, Create: `RSA2048`},
		"signature_algorithm":             acctest.Representation{RepType: acctest.Optional, Create: `SHA256_WITH_RSA`},
		"subject":                         acctest.RepresentationGroup{RepType: acctest.Required, Group: certificatesManagementCertificateSubjectRepresentation},
		"subject_alternative_names":       acctest.RepresentationGroup{RepType: acctest.Optional, Group: certificateCertificateConfigSubjectAlternativeNamesRepresentation},
		"validity":                        acctest.RepresentationGroup{RepType: acctest.Optional, Group: certificateCertificateConfigValidityRepresentation},
		"version_name":                    acctest.Representation{RepType: acctest.Optional, Create: `versionName`, Update: `versionName2`},
	}

	certificateCertificateConfigRepresentationCSR = map[string]interface{}{
		"config_type":                     acctest.Representation{RepType: acctest.Required, Create: `MANAGED_EXTERNALLY_ISSUED_BY_INTERNAL_CA`, Update: `MANAGED_EXTERNALLY_ISSUED_BY_INTERNAL_CA`},
		"csr_pem":                         acctest.Representation{RepType: acctest.Required, Create: `-----BEGIN CERTIFICATE REQUEST-----\nMIICzzCCAbcCAQAwgYkxCzAJBgNVBAYTAnVzMRMwEQYDVQQIDApXYXNoaW5ndG9u\nMREwDwYDVQQHDAhCZWxsZXZ1ZTEMMAoGA1UECgwDT0NJMQwwCgYDVQQLDANTRUMx\nGDAWBgNVBAMMD3d3dy50ZXN0Y3NyLmNvbTEcMBoGCSqGSIb3DQEJARYNdGVzdEB0\nZXN0LmNvbTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBALrlpiuehIMz\nre0pa+G8WjPaltQOKsA5LZWFUmnkRvy9zGtVKXmcZ70m8+SIIW0svwhDENtvl6wA\nPfQX8G+9XelIzyUv62EG9PwcSqW7rLpq8VygH/6fbYLU/LG8wuIjVq8yU8AQSvcJ\nBSNEiPMxN7qmXviKXHyY8A3ramgY+2wIfqNPRfrM/udPGd8xls/MviqCXLl3rT5c\nLDwCOi+rqOgWwt9wkOzAVNQWtwCaGaed1j98T3QQtYjeSVeV/HaGUquK9nare0dF\n49SEgRuCpVSeD/PHN5lz1YE99bioEiqfdATw05xZ08wTezgscCKfuaMMAXDF2Q0N\nqMNaynCFuuECAwEAAaAAMA0GCSqGSIb3DQEBCwUAA4IBAQBXsLUdghNHVHM54DRJ\nwNJdICec9HX5OuwHP1C/QpwwcEeO4lIKTxww4WWwReJVIoUewgGaxYLh+izp40bP\no+wFR0NDC9eBaRtnBZiPYn8bVA2PBl9z0VS2+AsQHR9hlKH0G3iU+C0O2wVSJpQW\nAIU3c/2CNg9GCTWmJE+Jd8dTd21WDlARBKw7GXwToHQL7J3vY+2/S2e6hdbxh9aG\n7ZEOYssclMtVZxKPdGp9l3JbwBrk/9hv3kxFZvbtxjnn3ah5Sas+YwIewn2VPSwY\nfoewMRgyYZKUFmw/T3YH2ony9ouT8mgPMJys72iVXO+Ri31VdL7AVQHMA0uQEpKM\n+Pn4\n-----END CERTIFICATE REQUEST-----`, Update: `csrPem2`},
		"issuer_certificate_authority_id": acctest.Representation{RepType: acctest.Required, Create: issuerCaId},
		"validity":                        acctest.RepresentationGroup{RepType: acctest.Optional, Group: certificateCertificateConfigValidityRepresentation},
		"version_name":                    acctest.Representation{RepType: acctest.Optional, Create: `versionName`, Update: `versionName2`},
	}

	certificateCertificateRulesRepresentation = map[string]interface{}{
		"advance_renewal_period": acctest.Representation{RepType: acctest.Required, Create: `P30D`, Update: `P45D`},
		"renewal_interval":       acctest.Representation{RepType: acctest.Required, Create: `P365D`, Update: `P100D`},
		"rule_type":              acctest.Representation{RepType: acctest.Required, Create: `CERTIFICATE_RENEWAL_RULE`},
	}

	certificateCertificateConfigSubjectAlternativeNamesRepresentation = map[string]interface{}{
		"type":  acctest.Representation{RepType: acctest.Optional, Create: `DNS`},
		"value": acctest.Representation{RepType: acctest.Optional, Create: `www.oracle.com`},
	}
	certificateCertificateConfigValidityRepresentation = map[string]interface{}{
		"time_of_validity_not_after":  acctest.Representation{RepType: acctest.Required, Create: certNotAfterCreate, Update: certNotAfterUpdate},
		"time_of_validity_not_before": acctest.Representation{RepType: acctest.Optional, Create: certNotBeforeCreate, Update: certNotBeforeUpdate},
	}

	CertificatesManagementCertificateResourceDependencies = DefinedTagsDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_user", "test_user", acctest.Required, acctest.Create, IdentityUserRepresentation)
)

func TestCertificatesManagementCertificateResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCertificatesManagementCertificateResource_basic")
	defer httpreplay.SaveScenario()

	provider := acctest.TestAccProvider
	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_certificates_management_certificate.test_certificate"
	datasourceName := "data.oci_certificates_management_certificates.test_certificates"
	singularDatasourceName := "data.oci_certificates_management_certificate.test_certificate"

	var resId, resId2 string
	// Save TF content to create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+CertificatesManagementCertificateResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_certificates_management_certificate", "test_certificate", acctest.Optional, acctest.Create, certificatesManagementCertificateRepresentation), "certificatesmanagement", "certificate", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { acctest.PreCheck(t) },
		Providers: map[string]*schema.Provider{
			"oci": provider,
		},
		Steps: []resource.TestStep{

			// verify create on CSR cert
			{
				Config: config + compartmentIdVariableStr +
					acctest.GenerateResourceFromRepresentationMap("oci_certificates_management_certificate", "test_certificate", acctest.Required, acctest.Create,
						acctest.RepresentationCopyWithNewProperties(certificatesManagementCertificateRepresentation, map[string]interface{}{
							"certificate_config": acctest.RepresentationGroup{RepType: acctest.Required, Group: certificateCertificateConfigRepresentationCSR},
							"name":               acctest.Representation{RepType: acctest.Required, Create: requiredCertName},
						})),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "certificate_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.config_type", "MANAGED_EXTERNALLY_ISSUED_BY_INTERNAL_CA"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.csr_pem", csrPem),
					resource.TestCheckResourceAttrSet(resourceName, "certificate_config.0.issuer_certificate_authority_id"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "name", requiredCertName),

					func(s *terraform.State) (err error) {
						resId, err = acctest.FromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + CertificatesManagementCertificateResourceDependencies,
			},

			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + CertificatesManagementCertificateResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_certificates_management_certificate", "test_certificate", acctest.Optional, acctest.Create, certificatesManagementCertificateRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "certificate_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.certificate_profile_type", "TLS_SERVER_OR_CLIENT"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.config_type", "ISSUED_BY_INTERNAL_CA"),
					resource.TestCheckResourceAttrSet(resourceName, "certificate_config.0.issuer_certificate_authority_id"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.key_algorithm", "RSA2048"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.signature_algorithm", "SHA256_WITH_RSA"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.subject.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.subject.0.common_name", "www.example.com"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.subject.0.country", "US"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.subject.0.distinguished_name_qualifier", "distinguishedNameQualifier"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.subject.0.domain_component", "domainComponent"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.subject.0.generation_qualifier", "JR"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.subject.0.given_name", "Sir"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.subject.0.initials", "HAM"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.subject.0.locality_name", "Seattle"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.subject.0.organization", "OCI"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.subject.0.organizational_unit", "SecurityProducts"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.subject.0.pseudonym", "pseudonym"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.subject.0.serial_number", "serialNumber"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.subject.0.state_or_province_name", "Washington"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.subject.0.street", "123 Main Street"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.subject.0.surname", "Last"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.subject.0.title", "Lord"),
					resource.TestCheckResourceAttrSet(resourceName, "certificate_config.0.subject.0.user_id"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.subject_alternative_names.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.subject_alternative_names.0.type", "DNS"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.subject_alternative_names.0.value", "www.oracle.com"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.validity.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.validity.0.time_of_validity_not_after", certNotAfterCreate),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.validity.0.time_of_validity_not_before", certNotBeforeCreate),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.version_name", "versionName"),
					resource.TestCheckResourceAttr(resourceName, "certificate_rules.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "certificate_rules.0.advance_renewal_period", "P30D"),
					resource.TestCheckResourceAttr(resourceName, "certificate_rules.0.renewal_interval", "P365D"),
					resource.TestCheckResourceAttr(resourceName, "certificate_rules.0.rule_type", "CERTIFICATE_RENEWAL_RULE"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(resourceName, "config_type"),
					resource.TestCheckResourceAttrSet(resourceName, "current_version.0.certificate_id"),
					resource.TestCheckResourceAttrSet(resourceName, "current_version.0.issuer_ca_version_number"),
					resource.TestCheckResourceAttrSet(resourceName, "current_version.0.serial_number"),
					resource.TestCheckResourceAttr(resourceName, "current_version.0.stages.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "current_version.0.subject_alternative_names.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "current_version.0.subject_alternative_names.0.type", "DNS"),
					resource.TestCheckResourceAttr(resourceName, "current_version.0.subject_alternative_names.0.value", "www.oracle.com"),
					resource.TestCheckResourceAttrSet(resourceName, "current_version.0.time_created"),
					resource.TestCheckResourceAttr(resourceName, "current_version.0.validity.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "current_version.0.validity.0.time_of_validity_not_after"),
					resource.TestCheckResourceAttrSet(resourceName, "current_version.0.validity.0.time_of_validity_not_before"),
					resource.TestCheckResourceAttr(resourceName, "current_version.0.version_name", "versionName"),
					resource.TestCheckResourceAttrSet(resourceName, "current_version.0.version_number"),
					resource.TestCheckResourceAttrSet(resourceName, "issuer_certificate_authority_id"),
					resource.TestCheckResourceAttr(resourceName, "key_algorithm", "RSA2048"),
					resource.TestCheckResourceAttr(resourceName, "signature_algorithm", "SHA256_WITH_RSA"),
					resource.TestCheckResourceAttr(resourceName, "subject.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "subject.0.common_name", "www.example.com"),
					resource.TestCheckResourceAttrSet(resourceName, "subject.0.country"),
					resource.TestCheckResourceAttr(resourceName, "subject.0.distinguished_name_qualifier", "distinguishedNameQualifier"),
					resource.TestCheckResourceAttr(resourceName, "subject.0.domain_component", "domainComponent"),
					resource.TestCheckResourceAttr(resourceName, "subject.0.generation_qualifier", "JR"),
					resource.TestCheckResourceAttr(resourceName, "subject.0.given_name", "Sir"),
					resource.TestCheckResourceAttr(resourceName, "subject.0.initials", "HAM"),
					resource.TestCheckResourceAttr(resourceName, "subject.0.locality_name", "Seattle"),
					resource.TestCheckResourceAttr(resourceName, "subject.0.organization", "OCI"),
					resource.TestCheckResourceAttr(resourceName, "subject.0.organizational_unit", "SecurityProducts"),
					resource.TestCheckResourceAttr(resourceName, "subject.0.pseudonym", "pseudonym"),
					resource.TestCheckResourceAttr(resourceName, "subject.0.serial_number", "serialNumber"),
					resource.TestCheckResourceAttr(resourceName, "subject.0.state_or_province_name", "Washington"),
					resource.TestCheckResourceAttr(resourceName, "subject.0.street", "123 Main Street"),
					resource.TestCheckResourceAttr(resourceName, "subject.0.surname", "Last"),
					resource.TestCheckResourceAttr(resourceName, "subject.0.title", "Lord"),
					resource.TestCheckResourceAttrSet(resourceName, "subject.0.user_id"),
					resource.TestCheckResourceAttr(resourceName, "description", "description"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "name", certNameForOptionalTests),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
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

			// verify update to the compartment (the compartment will be switched back in the next step)
			{
				Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + CertificatesManagementCertificateResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_certificates_management_certificate", "test_certificate", acctest.Optional, acctest.Create,
						acctest.RepresentationCopyWithNewProperties(certificatesManagementCertificateRepresentation, map[string]interface{}{
							"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
						})),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "certificate_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.certificate_profile_type", "TLS_SERVER_OR_CLIENT"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.config_type", "ISSUED_BY_INTERNAL_CA"),
					resource.TestCheckResourceAttrSet(resourceName, "certificate_config.0.issuer_certificate_authority_id"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.key_algorithm", "RSA2048"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.signature_algorithm", "SHA256_WITH_RSA"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.subject.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.subject.0.common_name", "www.example.com"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.subject.0.country", "US"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.subject.0.distinguished_name_qualifier", "distinguishedNameQualifier"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.subject.0.domain_component", "domainComponent"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.subject.0.generation_qualifier", "JR"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.subject.0.given_name", "Sir"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.subject.0.initials", "HAM"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.subject.0.locality_name", "Seattle"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.subject.0.organization", "OCI"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.subject.0.organizational_unit", "SecurityProducts"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.subject.0.pseudonym", "pseudonym"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.subject.0.serial_number", "serialNumber"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.subject.0.state_or_province_name", "Washington"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.subject.0.street", "123 Main Street"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.subject.0.surname", "Last"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.subject.0.title", "Lord"),
					resource.TestCheckResourceAttrSet(resourceName, "certificate_config.0.subject.0.user_id"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.subject_alternative_names.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.subject_alternative_names.0.type", "DNS"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.subject_alternative_names.0.value", "www.oracle.com"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.validity.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.validity.0.time_of_validity_not_after", certNotAfterCreate),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.validity.0.time_of_validity_not_before", certNotBeforeCreate),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.version_name", "versionName"),
					resource.TestCheckResourceAttr(resourceName, "certificate_rules.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "certificate_rules.0.advance_renewal_period", "P30D"),
					resource.TestCheckResourceAttr(resourceName, "certificate_rules.0.renewal_interval", "P365D"),
					resource.TestCheckResourceAttr(resourceName, "certificate_rules.0.rule_type", "CERTIFICATE_RENEWAL_RULE"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
					resource.TestCheckResourceAttrSet(resourceName, "config_type"),
					resource.TestCheckResourceAttr(resourceName, "description", "description"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "name", certNameForOptionalTests),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

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
				Config: config + compartmentIdVariableStr + CertificatesManagementCertificateResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_certificates_management_certificate", "test_certificate", acctest.Optional, acctest.Update, certificatesManagementCertificateRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "certificate_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.certificate_profile_type", "TLS_SERVER_OR_CLIENT"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.config_type", "ISSUED_BY_INTERNAL_CA"),
					resource.TestCheckResourceAttrSet(resourceName, "certificate_config.0.issuer_certificate_authority_id"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.key_algorithm", "RSA2048"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.signature_algorithm", "SHA256_WITH_RSA"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.subject.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.subject.0.common_name", "www.example.com"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.subject.0.country", "US"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.subject.0.distinguished_name_qualifier", "distinguishedNameQualifier"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.subject.0.domain_component", "domainComponent"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.subject.0.generation_qualifier", "JR"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.subject.0.given_name", "Sir"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.subject.0.initials", "HAM"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.subject.0.locality_name", "Seattle"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.subject.0.organization", "OCI"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.subject.0.organizational_unit", "SecurityProducts"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.subject.0.pseudonym", "pseudonym"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.subject.0.serial_number", "serialNumber"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.subject.0.state_or_province_name", "Washington"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.subject.0.street", "123 Main Street"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.subject.0.surname", "Last"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.subject.0.title", "Lord"),
					resource.TestCheckResourceAttrSet(resourceName, "certificate_config.0.subject.0.user_id"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.subject_alternative_names.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.subject_alternative_names.0.type", "DNS"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.subject_alternative_names.0.value", "www.oracle.com"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.validity.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.validity.0.time_of_validity_not_after", certNotAfterUpdate),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.validity.0.time_of_validity_not_before", certNotBeforeUpdate),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.version_name", "versionName2"),
					resource.TestCheckResourceAttr(resourceName, "certificate_rules.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "certificate_rules.0.advance_renewal_period", "P45D"),
					resource.TestCheckResourceAttr(resourceName, "certificate_rules.0.renewal_interval", "P100D"),
					resource.TestCheckResourceAttr(resourceName, "certificate_rules.0.rule_type", "CERTIFICATE_RENEWAL_RULE"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(resourceName, "config_type"),
					resource.TestCheckResourceAttr(resourceName, "description", "description2"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "name", certNameForOptionalTests),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

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
					acctest.GenerateDataSourceFromRepresentationMap("oci_certificates_management_certificates", "test_certificates", acctest.Optional, acctest.Create, certificatesManagementCertificateDataSourceRepresentation) +
					compartmentIdVariableStr + CertificatesManagementCertificateResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_certificates_management_certificate", "test_certificate", acctest.Optional, acctest.Update, certificatesManagementCertificateRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "name", certNameForOptionalTests),
					resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

					resource.TestCheckResourceAttr(datasourceName, "certificate_collection.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "certificate_collection.0.items.#", "1"),
				),
			},

			// verify singular datasource
			{
				Config: config +
					acctest.GenerateDataSourceFromRepresentationMap("oci_certificates_management_certificate", "test_certificate", acctest.Required, acctest.Create, CertificatesManagementcertificateSingularDataSourceRepresentation) +
					compartmentIdVariableStr + CertificatesManagementCertificateResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "certificate_id"),

					resource.TestCheckResourceAttrSet(singularDatasourceName, "certificate_profile_type"),
					resource.TestCheckResourceAttr(singularDatasourceName, "certificate_rules.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "certificate_rules.0.advance_renewal_period", "P45D"),
					resource.TestCheckResourceAttr(singularDatasourceName, "certificate_rules.0.renewal_interval", "P100D"),
					resource.TestCheckResourceAttr(singularDatasourceName, "certificate_rules.0.rule_type", "CERTIFICATE_RENEWAL_RULE"),
					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "config_type"),
					resource.TestCheckResourceAttr(singularDatasourceName, "current_version.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
					resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "key_algorithm"),
					resource.TestCheckResourceAttr(singularDatasourceName, "name", certNameForOptionalTests),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "signature_algorithm"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
					resource.TestCheckResourceAttr(singularDatasourceName, "subject.#", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				),
			},
			// verify resource import
			{
				Config:            config + CertificatesManagementCertificateRequiredOnlyResource,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"certificate_config",
				},
				ResourceName: resourceName,
			},
		},
	})
}
