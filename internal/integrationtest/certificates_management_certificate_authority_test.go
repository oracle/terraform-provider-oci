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
	rootCaName = "test-root-ca-" + utils.RandomString(10, utils.CharsetWithoutDigits)
	subCaName  = "test-sub-ca-" + utils.RandomString(10, utils.CharsetWithoutDigits)

	notBeforeCreate = time.Now().UTC().AddDate(0, 0, 8).Truncate(time.Millisecond).Format(time.RFC3339Nano)
	notBeforeUpdate = time.Now().UTC().AddDate(0, 0, 30).Truncate(time.Millisecond).Format(time.RFC3339Nano)
	notAfterCreate  = time.Now().UTC().AddDate(10, 0, 0).Truncate(time.Millisecond).Format(time.RFC3339Nano)
	notAfterUpdate  = time.Now().UTC().AddDate(12, 0, 0).Truncate(time.Millisecond).Format(time.RFC3339Nano)

	cmKmsKeyId            = utils.GetEnvSettingWithBlankDefault("kms_key_ocid")
	cmKmsKeyIdVariableStr = fmt.Sprintf("variable \"kms_key_ocid\" { default = \"%s\" }\n", cmKmsKeyId)

	issuerCaId            = utils.GetEnvSettingWithBlankDefault("issuer_ca_ocid")
	issuerCaIdVariableStr = fmt.Sprintf("variable \"issuer_ca_ocid\" { default = \"%s\" }\n", issuerCaId)

	CertificatesManagementCertificateAuthorityRequiredOnlyResource = acctest.GenerateResourceFromRepresentationMap("oci_certificates_management_certificate_authority", "test_certificate_authority", acctest.Required, acctest.Create, certificateAuthorityRepresentationRoot)

	CertificatesManagementCertificateAuthorityResourceConfig = CertificatesManagementCertificateAuthorityResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_certificates_management_certificate_authority", "test_certificate_authority", acctest.Optional, acctest.Update, certificateAuthorityRepresentationSub)

	CertificatesManagementcertificateAuthoritySingularDataSourceRepresentation = map[string]interface{}{
		"certificate_authority_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_certificates_management_certificate_authority.test_certificate_authority.id}`},
	}

	certificateAuthorityDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"name":           acctest.Representation{RepType: acctest.Optional, Create: subCaName},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: certificateAuthorityDataSourceFilterRepresentation},
	}
	certificateAuthorityDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_certificates_management_certificate_authority.test_certificate_authority.id}`}},
	}

	certificateAuthorityRepresentation = map[string]interface{}{
		"certificate_authority_config":        acctest.RepresentationGroup{RepType: acctest.Required, Group: certificateAuthorityCertificateAuthorityConfigRepresentationRoot},
		"compartment_id":                      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"kms_key_id":                          acctest.Representation{RepType: acctest.Required, Create: cmKmsKeyId},
		"name":                                acctest.Representation{RepType: acctest.Required, Create: rootCaName},
		"certificate_authority_rules":         acctest.RepresentationGroup{RepType: acctest.Optional, Group: certificateAuthorityCertificateAuthorityRulesRepresentation},
		"certificate_revocation_list_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: certificateAuthorityCertificateRevocationListDetailsRepresentation},
		"defined_tags":                        acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":                         acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"freeform_tags":                       acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
	}

	/** ROOT CONFIGS **/
	certificateAuthorityRepresentationRoot = map[string]interface{}{
		"certificate_authority_config":        acctest.RepresentationGroup{RepType: acctest.Required, Group: certificateAuthorityCertificateAuthorityConfigRepresentationRoot},
		"compartment_id":                      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"kms_key_id":                          acctest.Representation{RepType: acctest.Required, Create: cmKmsKeyId},
		"name":                                acctest.Representation{RepType: acctest.Required, Create: rootCaName},
		"certificate_authority_rules":         acctest.RepresentationGroup{RepType: acctest.Optional, Group: certificateAuthorityCertificateAuthorityRulesRepresentation},
		"certificate_revocation_list_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: certificateAuthorityCertificateRevocationListDetailsRepresentation},
		"defined_tags":                        acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":                         acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"freeform_tags":                       acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"lifecycle":                           acctest.RepresentationGroup{RepType: acctest.Required, Group: certificatesManagementIgnoreChangesRepresentation},
	}

	certificateAuthorityCertificateAuthorityConfigRepresentationRoot = map[string]interface{}{
		"config_type":       acctest.Representation{RepType: acctest.Required, Create: `ROOT_CA_GENERATED_INTERNALLY`},
		"subject":           acctest.RepresentationGroup{RepType: acctest.Required, Group: certificatesManagementCertificateSubjectRepresentation},
		"signing_algorithm": acctest.Representation{RepType: acctest.Optional, Create: `SHA256_WITH_RSA`},
		"validity":          acctest.RepresentationGroup{RepType: acctest.Optional, Group: certificateAuthorityCertificateAuthorityConfigValidityRepresentation},
		"version_name":      acctest.Representation{RepType: acctest.Optional, Create: `versionName`, Update: `versionName2`},
	}

	/** SUBORDINATE CONFIGS **/
	certificateAuthorityRepresentationSub = map[string]interface{}{
		"certificate_authority_config":        acctest.RepresentationGroup{RepType: acctest.Required, Group: certificateAuthorityCertificateAuthorityConfigRepresentationSub},
		"compartment_id":                      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"kms_key_id":                          acctest.Representation{RepType: acctest.Required, Create: cmKmsKeyId},
		"name":                                acctest.Representation{RepType: acctest.Required, Create: subCaName},
		"certificate_authority_rules":         acctest.RepresentationGroup{RepType: acctest.Optional, Group: certificateAuthorityCertificateAuthorityRulesRepresentation},
		"certificate_revocation_list_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: certificateAuthorityCertificateRevocationListDetailsRepresentation},
		"defined_tags":                        acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":                         acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"freeform_tags":                       acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"lifecycle":                           acctest.RepresentationGroup{RepType: acctest.Required, Group: certificatesManagementIgnoreChangesRepresentation},
	}

	certificateAuthorityCertificateAuthorityConfigRepresentationSub = map[string]interface{}{
		"config_type":                     acctest.Representation{RepType: acctest.Required, Create: `SUBORDINATE_CA_ISSUED_BY_INTERNAL_CA`},
		"subject":                         acctest.RepresentationGroup{RepType: acctest.Required, Group: certificatesManagementCertificateSubjectRepresentation},
		"issuer_certificate_authority_id": acctest.Representation{RepType: acctest.Optional, Create: issuerCaId},
		"signing_algorithm":               acctest.Representation{RepType: acctest.Optional, Create: `SHA256_WITH_RSA`},
		"validity":                        acctest.RepresentationGroup{RepType: acctest.Optional, Group: certificateAuthorityCertificateAuthorityConfigValidityRepresentation},
		"version_name":                    acctest.Representation{RepType: acctest.Optional, Create: `versionName`, Update: `versionName2`},
	}

	certificateAuthorityCertificateAuthorityRulesRepresentation = map[string]interface{}{
		"rule_type": acctest.Representation{RepType: acctest.Required, Create: `CERTIFICATE_AUTHORITY_ISSUANCE_EXPIRY_RULE`, Update: `CERTIFICATE_AUTHORITY_ISSUANCE_EXPIRY_RULE`},
		"certificate_authority_max_validity_duration": acctest.Representation{RepType: acctest.Optional, Create: `P1000D`, Update: `P1500D`},
		"leaf_certificate_max_validity_duration":      acctest.Representation{RepType: acctest.Optional, Create: `P500D`, Update: `P400D`},
	}
	certificateAuthorityCertificateRevocationListDetailsRepresentation = map[string]interface{}{
		"object_storage_config": acctest.RepresentationGroup{RepType: acctest.Required, Group: certificateAuthorityCertificateRevocationListDetailsObjectStorageConfigRepresentation},
		"custom_formatted_urls": acctest.Representation{RepType: acctest.Optional, Create: []string{`https://www.example.com/crl.bin`}, Update: []string{`https://www.example.com/crl2.bin`}},
	}
	certificatesManagementCertificateSubjectRepresentation = map[string]interface{}{
		"common_name":                  acctest.Representation{RepType: acctest.Required, Create: `www.example.com`},
		"country":                      acctest.Representation{RepType: acctest.Optional, Create: `US`},
		"distinguished_name_qualifier": acctest.Representation{RepType: acctest.Optional, Create: `distinguishedNameQualifier`},
		"domain_component":             acctest.Representation{RepType: acctest.Optional, Create: `domainComponent`},
		"generation_qualifier":         acctest.Representation{RepType: acctest.Optional, Create: `JR`},
		"given_name":                   acctest.Representation{RepType: acctest.Optional, Create: `Sir`},
		"initials":                     acctest.Representation{RepType: acctest.Optional, Create: `HAM`},
		"locality_name":                acctest.Representation{RepType: acctest.Optional, Create: `Seattle`},
		"organization":                 acctest.Representation{RepType: acctest.Optional, Create: `OCI`},
		"organizational_unit":          acctest.Representation{RepType: acctest.Optional, Create: `SecurityProducts`},
		"pseudonym":                    acctest.Representation{RepType: acctest.Optional, Create: `pseudonym`},
		"serial_number":                acctest.Representation{RepType: acctest.Optional, Create: `serialNumber`},
		"state_or_province_name":       acctest.Representation{RepType: acctest.Optional, Create: `Washington`},
		"street":                       acctest.Representation{RepType: acctest.Optional, Create: `123 Main Street`},
		"surname":                      acctest.Representation{RepType: acctest.Optional, Create: `Last`},
		"title":                        acctest.Representation{RepType: acctest.Optional, Create: `Lord`},
		"user_id":                      acctest.Representation{RepType: acctest.Optional, Create: `${oci_identity_user.test_user.id}`},
	}
	certificateAuthorityCertificateAuthorityConfigValidityRepresentation = map[string]interface{}{
		"time_of_validity_not_after":  acctest.Representation{RepType: acctest.Optional, Create: notAfterCreate, Update: notAfterUpdate},
		"time_of_validity_not_before": acctest.Representation{RepType: acctest.Optional, Create: notBeforeCreate, Update: notBeforeUpdate},
	}

	certificateAuthorityCertificateRevocationListDetailsObjectStorageConfigRepresentation = map[string]interface{}{
		"object_storage_bucket_name":        acctest.Representation{RepType: acctest.Required, Create: `${oci_objectstorage_bucket.test_bucket.name}`},
		"object_storage_object_name_format": acctest.Representation{RepType: acctest.Required, Create: subCaName + "versionName{}.crl", Update: subCaName + "versionName2{}.crl"},
		"object_storage_namespace":          acctest.Representation{RepType: acctest.Optional, Create: `${oci_objectstorage_bucket.test_bucket.namespace}`},
	}

	certificatesManagementIgnoreChangesRepresentation = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`}},
	}

	CertificatesManagementCertificateAuthorityResourceDependencies = DefinedTagsDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_user", "test_user", acctest.Required, acctest.Create, IdentityUserRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_objectstorage_bucket", "test_bucket", acctest.Required, acctest.Create, ObjectStorageBucketRepresentation) +
		acctest.GenerateDataSourceFromRepresentationMap("oci_objectstorage_namespace", "test_namespace", acctest.Optional, acctest.Create, ObjectStorageObjectStorageNamespaceSingularDataSourceRepresentation)
)

func TestCertificatesManagementCertificateAuthorityResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCertificatesManagementCertificateAuthorityResource_basic")
	defer httpreplay.SaveScenario()

	provider := acctest.TestAccProvider
	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_certificates_management_certificate_authority.test_certificate_authority"
	datasourceName := "data.oci_certificates_management_certificate_authorities.test_certificate_authorities"
	singularDatasourceName := "data.oci_certificates_management_certificate_authority.test_certificate_authority"

	var resId, resId2 string
	// Save TF content to create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+CertificatesManagementCertificateAuthorityResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_certificates_management_certificate_authority", "test_certificate_authority", acctest.Optional, acctest.Create, certificateAuthorityRepresentationSub), "certificatesmanagement", "certificateAuthority", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { acctest.PreCheck(t) },
		Providers: map[string]*schema.Provider{
			"oci": provider,
		},
		Steps: []resource.TestStep{

			// verify create
			{
				Config: config + compartmentIdVariableStr + CertificatesManagementCertificateAuthorityRequiredOnlyResource,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "certificate_authority_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "certificate_authority_config.0.config_type", "ROOT_CA_GENERATED_INTERNALLY"),
					resource.TestCheckResourceAttr(resourceName, "certificate_authority_config.0.subject.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "certificate_authority_config.0.subject.0.common_name", "www.example.com"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(resourceName, "kms_key_id"),
					resource.TestCheckResourceAttr(resourceName, "name", rootCaName),

					func(s *terraform.State) (err error) {
						resId, err = acctest.FromInstanceState(s, resourceName, "id")
						fmt.Printf("variable \"resId\" { default = \"%s\" }\n", resId)
						return err
					},
				),
			},

			//delete before next create
			{
				Config: config + compartmentIdVariableStr + CertificatesManagementCertificateAuthorityResourceDependencies,
			},

			//verify create with optionals
			{
				Config: config + compartmentIdVariableStr + CertificatesManagementCertificateAuthorityResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_certificates_management_certificate_authority", "test_certificate_authority", acctest.Optional, acctest.Create, certificateAuthorityRepresentationSub),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "certificate_authority_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "certificate_authority_config.0.config_type", "SUBORDINATE_CA_ISSUED_BY_INTERNAL_CA"),
					resource.TestCheckResourceAttrSet(resourceName, "certificate_authority_config.0.issuer_certificate_authority_id"),
					resource.TestCheckResourceAttr(resourceName, "certificate_authority_config.0.signing_algorithm", "SHA256_WITH_RSA"),
					resource.TestCheckResourceAttr(resourceName, "certificate_authority_config.0.subject.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "certificate_authority_config.0.subject.0.common_name", "www.example.com"),
					resource.TestCheckResourceAttr(resourceName, "certificate_authority_config.0.subject.0.country", "US"),
					resource.TestCheckResourceAttr(resourceName, "certificate_authority_config.0.subject.0.distinguished_name_qualifier", "distinguishedNameQualifier"),
					resource.TestCheckResourceAttr(resourceName, "certificate_authority_config.0.subject.0.domain_component", "domainComponent"),
					resource.TestCheckResourceAttr(resourceName, "certificate_authority_config.0.subject.0.generation_qualifier", "JR"),
					resource.TestCheckResourceAttr(resourceName, "certificate_authority_config.0.subject.0.given_name", "Sir"),
					resource.TestCheckResourceAttr(resourceName, "certificate_authority_config.0.subject.0.initials", "HAM"),
					resource.TestCheckResourceAttr(resourceName, "certificate_authority_config.0.subject.0.locality_name", "Seattle"),
					resource.TestCheckResourceAttr(resourceName, "certificate_authority_config.0.subject.0.organization", "OCI"),
					resource.TestCheckResourceAttr(resourceName, "certificate_authority_config.0.subject.0.organizational_unit", "SecurityProducts"),
					resource.TestCheckResourceAttr(resourceName, "certificate_authority_config.0.subject.0.pseudonym", "pseudonym"),
					resource.TestCheckResourceAttr(resourceName, "certificate_authority_config.0.subject.0.serial_number", "serialNumber"),
					resource.TestCheckResourceAttr(resourceName, "certificate_authority_config.0.subject.0.state_or_province_name", "Washington"),
					resource.TestCheckResourceAttr(resourceName, "certificate_authority_config.0.subject.0.street", "123 Main Street"),
					resource.TestCheckResourceAttr(resourceName, "certificate_authority_config.0.subject.0.surname", "Last"),
					resource.TestCheckResourceAttr(resourceName, "certificate_authority_config.0.subject.0.title", "Lord"),
					resource.TestCheckResourceAttrSet(resourceName, "certificate_authority_config.0.subject.0.user_id"),
					resource.TestCheckResourceAttr(resourceName, "certificate_authority_config.0.validity.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "certificate_authority_config.0.validity.0.time_of_validity_not_after", notAfterCreate),
					resource.TestCheckResourceAttr(resourceName, "certificate_authority_config.0.validity.0.time_of_validity_not_before", notBeforeCreate),
					resource.TestCheckResourceAttr(resourceName, "certificate_authority_config.0.version_name", "versionName"),
					resource.TestCheckResourceAttr(resourceName, "certificate_authority_rules.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "certificate_authority_rules.0.certificate_authority_max_validity_duration", "P1000D"),
					resource.TestCheckResourceAttr(resourceName, "certificate_authority_rules.0.leaf_certificate_max_validity_duration", "P500D"),
					resource.TestCheckResourceAttr(resourceName, "certificate_authority_rules.0.rule_type", "CERTIFICATE_AUTHORITY_ISSUANCE_EXPIRY_RULE"),
					resource.TestCheckResourceAttr(resourceName, "certificate_revocation_list_details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "certificate_revocation_list_details.0.custom_formatted_urls.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "certificate_revocation_list_details.0.object_storage_config.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "certificate_revocation_list_details.0.object_storage_config.0.object_storage_bucket_name"),
					resource.TestCheckResourceAttrSet(resourceName, "certificate_revocation_list_details.0.object_storage_config.0.object_storage_namespace"),
					resource.TestCheckResourceAttr(resourceName, "certificate_revocation_list_details.0.object_storage_config.0.object_storage_object_name_format", subCaName+"versionName{}.crl"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(resourceName, "config_type"),
					resource.TestCheckResourceAttrSet(resourceName, "current_version.0.certificate_authority_id"),
					//resource.TestCheckResourceAttrSet(resourceName, "current_version.0.issuer_ca_version_number"),
					resource.TestCheckResourceAttrSet(resourceName, "current_version.0.serial_number"),
					resource.TestCheckResourceAttr(resourceName, "current_version.0.stages.#", "2"),
					resource.TestCheckResourceAttrSet(resourceName, "current_version.0.time_created"),
					resource.TestCheckResourceAttr(resourceName, "current_version.0.validity.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "current_version.0.validity.0.time_of_validity_not_after"),
					resource.TestCheckResourceAttrSet(resourceName, "current_version.0.validity.0.time_of_validity_not_before"),
					resource.TestCheckResourceAttrSet(resourceName, "current_version.0.version_name"),
					resource.TestCheckResourceAttrSet(resourceName, "current_version.0.version_number"),
					resource.TestCheckResourceAttrSet(resourceName, "issuer_certificate_authority_id"),
					resource.TestCheckResourceAttrSet(resourceName, "signing_algorithm"),
					resource.TestCheckResourceAttrSet(resourceName, "subject.#"),
					resource.TestCheckResourceAttrSet(resourceName, "subject.0.common_name"),
					resource.TestCheckResourceAttrSet(resourceName, "subject.0.country"),
					resource.TestCheckResourceAttrSet(resourceName, "subject.0.distinguished_name_qualifier"),
					resource.TestCheckResourceAttrSet(resourceName, "subject.0.domain_component"),
					resource.TestCheckResourceAttrSet(resourceName, "subject.0.generation_qualifier"),
					resource.TestCheckResourceAttrSet(resourceName, "subject.0.given_name"),
					resource.TestCheckResourceAttrSet(resourceName, "subject.0.initials"),
					resource.TestCheckResourceAttrSet(resourceName, "subject.0.locality_name"),
					resource.TestCheckResourceAttrSet(resourceName, "subject.0.organization"),
					resource.TestCheckResourceAttrSet(resourceName, "subject.0.organizational_unit"),
					resource.TestCheckResourceAttrSet(resourceName, "subject.0.pseudonym"),
					resource.TestCheckResourceAttrSet(resourceName, "subject.0.serial_number"),
					resource.TestCheckResourceAttrSet(resourceName, "subject.0.state_or_province_name"),
					resource.TestCheckResourceAttrSet(resourceName, "subject.0.street"),
					resource.TestCheckResourceAttrSet(resourceName, "subject.0.surname"),
					resource.TestCheckResourceAttrSet(resourceName, "subject.0.title"),
					resource.TestCheckResourceAttrSet(resourceName, "subject.0.user_id"),
					resource.TestCheckResourceAttr(resourceName, "description", "description"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "kms_key_id"),
					resource.TestCheckResourceAttr(resourceName, "name", subCaName),
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
				Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + CertificatesManagementCertificateAuthorityResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_certificates_management_certificate_authority", "test_certificate_authority", acctest.Optional, acctest.Create,
						acctest.RepresentationCopyWithNewProperties(certificateAuthorityRepresentationSub, map[string]interface{}{
							"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
						})),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "certificate_authority_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "certificate_authority_config.0.config_type", "SUBORDINATE_CA_ISSUED_BY_INTERNAL_CA"),
					resource.TestCheckResourceAttrSet(resourceName, "certificate_authority_config.0.issuer_certificate_authority_id"),
					resource.TestCheckResourceAttr(resourceName, "certificate_authority_config.0.signing_algorithm", "SHA256_WITH_RSA"),
					resource.TestCheckResourceAttr(resourceName, "certificate_authority_config.0.subject.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "certificate_authority_config.0.subject.0.common_name", "www.example.com"),
					resource.TestCheckResourceAttr(resourceName, "certificate_authority_config.0.subject.0.country", "US"),
					resource.TestCheckResourceAttr(resourceName, "certificate_authority_config.0.subject.0.distinguished_name_qualifier", "distinguishedNameQualifier"),
					resource.TestCheckResourceAttr(resourceName, "certificate_authority_config.0.subject.0.domain_component", "domainComponent"),
					resource.TestCheckResourceAttr(resourceName, "certificate_authority_config.0.subject.0.generation_qualifier", "JR"),
					resource.TestCheckResourceAttr(resourceName, "certificate_authority_config.0.subject.0.given_name", "Sir"),
					resource.TestCheckResourceAttr(resourceName, "certificate_authority_config.0.subject.0.initials", "HAM"),
					resource.TestCheckResourceAttr(resourceName, "certificate_authority_config.0.subject.0.locality_name", "Seattle"),
					resource.TestCheckResourceAttr(resourceName, "certificate_authority_config.0.subject.0.organization", "OCI"),
					resource.TestCheckResourceAttr(resourceName, "certificate_authority_config.0.subject.0.organizational_unit", "SecurityProducts"),
					resource.TestCheckResourceAttr(resourceName, "certificate_authority_config.0.subject.0.pseudonym", "pseudonym"),
					resource.TestCheckResourceAttr(resourceName, "certificate_authority_config.0.subject.0.serial_number", "serialNumber"),
					resource.TestCheckResourceAttr(resourceName, "certificate_authority_config.0.subject.0.state_or_province_name", "Washington"),
					resource.TestCheckResourceAttr(resourceName, "certificate_authority_config.0.subject.0.street", "123 Main Street"),
					resource.TestCheckResourceAttr(resourceName, "certificate_authority_config.0.subject.0.surname", "Last"),
					resource.TestCheckResourceAttr(resourceName, "certificate_authority_config.0.subject.0.title", "Lord"),
					resource.TestCheckResourceAttrSet(resourceName, "certificate_authority_config.0.subject.0.user_id"),
					resource.TestCheckResourceAttr(resourceName, "certificate_authority_config.0.validity.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "certificate_authority_config.0.validity.0.time_of_validity_not_after", notAfterCreate),
					resource.TestCheckResourceAttr(resourceName, "certificate_authority_config.0.validity.0.time_of_validity_not_before", notBeforeCreate),
					resource.TestCheckResourceAttr(resourceName, "certificate_authority_config.0.version_name", "versionName"),
					resource.TestCheckResourceAttr(resourceName, "certificate_authority_rules.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "certificate_authority_rules.0.certificate_authority_max_validity_duration", "P1000D"),
					resource.TestCheckResourceAttr(resourceName, "certificate_authority_rules.0.leaf_certificate_max_validity_duration", "P500D"),
					resource.TestCheckResourceAttr(resourceName, "certificate_authority_rules.0.rule_type", "CERTIFICATE_AUTHORITY_ISSUANCE_EXPIRY_RULE"),
					resource.TestCheckResourceAttr(resourceName, "certificate_revocation_list_details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "certificate_revocation_list_details.0.custom_formatted_urls.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "certificate_revocation_list_details.0.object_storage_config.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "certificate_revocation_list_details.0.object_storage_config.0.object_storage_bucket_name"),
					resource.TestCheckResourceAttrSet(resourceName, "certificate_revocation_list_details.0.object_storage_config.0.object_storage_namespace"),
					resource.TestCheckResourceAttr(resourceName, "certificate_revocation_list_details.0.object_storage_config.0.object_storage_object_name_format", subCaName+"versionName{}.crl"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
					resource.TestCheckResourceAttrSet(resourceName, "config_type"),
					resource.TestCheckResourceAttr(resourceName, "description", "description"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "kms_key_id"),
					resource.TestCheckResourceAttr(resourceName, "name", subCaName),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

					func(s *terraform.State) (err error) {
						resId2, err = acctest.FromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("resource recreated when it was supposed to be updated. resId %s and resId2 %s", resId, resId2)
						}
						return err
					},
				),
			},

			// verify updates to updatable parameters, VALIDITY RULES
			{
				Config: config + compartmentIdVariableStr + CertificatesManagementCertificateAuthorityResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_certificates_management_certificate_authority", "test_certificate_authority", acctest.Optional, acctest.Update, certificateAuthorityRepresentationSub),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "certificate_authority_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "certificate_authority_config.0.config_type", "SUBORDINATE_CA_ISSUED_BY_INTERNAL_CA"),
					resource.TestCheckResourceAttrSet(resourceName, "certificate_authority_config.0.issuer_certificate_authority_id"),
					resource.TestCheckResourceAttr(resourceName, "certificate_authority_config.0.signing_algorithm", "SHA256_WITH_RSA"),
					resource.TestCheckResourceAttr(resourceName, "certificate_authority_config.0.subject.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "certificate_authority_config.0.subject.0.common_name", "www.example.com"),
					resource.TestCheckResourceAttr(resourceName, "certificate_authority_config.0.subject.0.country", "US"),
					resource.TestCheckResourceAttr(resourceName, "certificate_authority_config.0.subject.0.distinguished_name_qualifier", "distinguishedNameQualifier"),
					resource.TestCheckResourceAttr(resourceName, "certificate_authority_config.0.subject.0.domain_component", "domainComponent"),
					resource.TestCheckResourceAttr(resourceName, "certificate_authority_config.0.subject.0.generation_qualifier", "JR"),
					resource.TestCheckResourceAttr(resourceName, "certificate_authority_config.0.subject.0.given_name", "Sir"),
					resource.TestCheckResourceAttr(resourceName, "certificate_authority_config.0.subject.0.initials", "HAM"),
					resource.TestCheckResourceAttr(resourceName, "certificate_authority_config.0.subject.0.locality_name", "Seattle"),
					resource.TestCheckResourceAttr(resourceName, "certificate_authority_config.0.subject.0.organization", "OCI"),
					resource.TestCheckResourceAttr(resourceName, "certificate_authority_config.0.subject.0.organizational_unit", "SecurityProducts"),
					resource.TestCheckResourceAttr(resourceName, "certificate_authority_config.0.subject.0.pseudonym", "pseudonym"),
					resource.TestCheckResourceAttr(resourceName, "certificate_authority_config.0.subject.0.serial_number", "serialNumber"),
					resource.TestCheckResourceAttr(resourceName, "certificate_authority_config.0.subject.0.state_or_province_name", "Washington"),
					resource.TestCheckResourceAttr(resourceName, "certificate_authority_config.0.subject.0.street", "123 Main Street"),
					resource.TestCheckResourceAttr(resourceName, "certificate_authority_config.0.subject.0.surname", "Last"),
					resource.TestCheckResourceAttr(resourceName, "certificate_authority_config.0.subject.0.title", "Lord"),
					resource.TestCheckResourceAttrSet(resourceName, "certificate_authority_config.0.subject.0.user_id"),
					resource.TestCheckResourceAttr(resourceName, "certificate_authority_config.0.validity.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "certificate_authority_config.0.validity.0.time_of_validity_not_after", notAfterUpdate),
					resource.TestCheckResourceAttr(resourceName, "certificate_authority_config.0.validity.0.time_of_validity_not_before", notBeforeUpdate),
					resource.TestCheckResourceAttr(resourceName, "certificate_authority_config.0.version_name", "versionName2"),
					resource.TestCheckResourceAttr(resourceName, "certificate_authority_rules.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "certificate_authority_rules.0.certificate_authority_max_validity_duration", "P1500D"),
					resource.TestCheckResourceAttr(resourceName, "certificate_authority_rules.0.leaf_certificate_max_validity_duration", "P400D"),
					resource.TestCheckResourceAttr(resourceName, "certificate_authority_rules.0.rule_type", "CERTIFICATE_AUTHORITY_ISSUANCE_EXPIRY_RULE"),
					resource.TestCheckResourceAttr(resourceName, "certificate_revocation_list_details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "certificate_revocation_list_details.0.custom_formatted_urls.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "certificate_revocation_list_details.0.object_storage_config.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "certificate_revocation_list_details.0.object_storage_config.0.object_storage_bucket_name"),
					resource.TestCheckResourceAttrSet(resourceName, "certificate_revocation_list_details.0.object_storage_config.0.object_storage_namespace"),
					resource.TestCheckResourceAttr(resourceName, "certificate_revocation_list_details.0.object_storage_config.0.object_storage_object_name_format", subCaName+"versionName2{}.crl"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(resourceName, "config_type"),
					resource.TestCheckResourceAttr(resourceName, "description", "description2"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "kms_key_id"),
					resource.TestCheckResourceAttr(resourceName, "name", subCaName),
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
					acctest.GenerateDataSourceFromRepresentationMap("oci_certificates_management_certificate_authorities", "test_certificate_authorities", acctest.Optional, acctest.Update, certificateAuthorityDataSourceRepresentation) +
					compartmentIdVariableStr + CertificatesManagementCertificateAuthorityResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_certificates_management_certificate_authority", "test_certificate_authority", acctest.Optional, acctest.Update, certificateAuthorityRepresentationSub),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "name", subCaName),
					resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

					resource.TestCheckResourceAttr(datasourceName, "certificate_authority_collection.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "certificate_authority_collection.0.items.#", "1"),
				),
			},
			// verify singular datasource
			{
				Config: config +
					acctest.GenerateDataSourceFromRepresentationMap("oci_certificates_management_certificate_authority", "test_certificate_authority", acctest.Required, acctest.Create, CertificatesManagementcertificateAuthoritySingularDataSourceRepresentation) +
					compartmentIdVariableStr + CertificatesManagementCertificateAuthorityResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "certificate_authority_id"),

					resource.TestCheckResourceAttr(singularDatasourceName, "certificate_authority_rules.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "certificate_authority_rules.0.certificate_authority_max_validity_duration", "P1500D"),
					resource.TestCheckResourceAttr(singularDatasourceName, "certificate_authority_rules.0.leaf_certificate_max_validity_duration", "P400D"),
					resource.TestCheckResourceAttr(singularDatasourceName, "certificate_authority_rules.0.rule_type", "CERTIFICATE_AUTHORITY_ISSUANCE_EXPIRY_RULE"),
					resource.TestCheckResourceAttr(singularDatasourceName, "certificate_revocation_list_details.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "certificate_revocation_list_details.0.custom_formatted_urls.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "certificate_revocation_list_details.0.object_storage_config.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "certificate_revocation_list_details.0.object_storage_config.0.object_storage_namespace", "ocivlt"),
					resource.TestCheckResourceAttr(singularDatasourceName, "certificate_revocation_list_details.0.object_storage_config.0.object_storage_object_name_format", subCaName+"versionName2{}.crl"),
					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "config_type"),
					resource.TestCheckResourceAttr(singularDatasourceName, "current_version.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
					resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "name", subCaName),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "signing_algorithm"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
					resource.TestCheckResourceAttr(singularDatasourceName, "subject.#", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				),
			},
			// verify resource import
			{
				Config:            config + CertificatesManagementCertificateAuthorityRequiredOnlyResource,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"certificate_authority_config",
				},
				ResourceName: resourceName,
			},
		},
	})
}
