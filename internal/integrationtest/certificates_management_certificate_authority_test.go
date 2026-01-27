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

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	rootCaName                                 = "test-root-ca-" + utils.RandomString(10, utils.CharsetWithoutDigits)
	externallyManagedRootCA                    = "test-externally-managed-root-ca-" + utils.RandomString(10, utils.CharsetWithoutDigits)
	subCAManagedInternallyIssuedExternally     = "test-sub-ca-" + utils.RandomString(10, utils.CharsetWithoutDigits)
	subCAManagedInternallyIssuedExternallyName = "test-sub-managed-internally-issued-external-ca-" + utils.RandomString(10, utils.CharsetWithoutDigits)
	notBeforeCreate                            = time.Now().UTC().AddDate(0, 0, 8).Truncate(time.Millisecond).Format(time.RFC3339Nano)
	notBeforeUpdate                            = time.Now().UTC().AddDate(0, 0, 30).Truncate(time.Millisecond).Format(time.RFC3339Nano)
	notAfterCreate                             = time.Now().UTC().AddDate(5, 0, 0).Truncate(time.Millisecond).Format(time.RFC3339Nano)
	notAfterUpdate                             = time.Now().UTC().AddDate(7, 0, 0).Truncate(time.Millisecond).Format(time.RFC3339Nano)

	cmKmsKeyId            = utils.GetEnvSettingWithBlankDefault("kms_key_ocid")
	cmKmsKeyIdVariableStr = fmt.Sprintf("variable \"kms_key_ocid\" { default = \"%s\" }\n", cmKmsKeyId)

	issuerCaId            = utils.GetEnvSettingWithBlankDefault("issuer_ca_ocid")
	issuerCaIdVariableStr = fmt.Sprintf("variable \"issuer_ca_ocid\" { default = \"%s\" }\n", issuerCaId)

	CertificatesManagementCertificateAuthorityRequiredOnlyResource = acctest.GenerateResourceFromRepresentationMap("oci_certificates_management_certificate_authority", "test_certificate_authority", acctest.Required, acctest.Create, certificateAuthorityRepresentationRoot)

	CertificatesManagementRootManagedExternallyResource = acctest.GenerateResourceFromRepresentationMap("oci_certificates_management_certificate_authority", "test_certificate_authority_root_managed_externally", acctest.Optional, acctest.Create, certificateAuthorityRepresentationRootManagedExternally)

	CertificatesManagementCertificateAuthorityPathLengthTestResource = acctest.GenerateResourceFromRepresentationMap("oci_certificates_management_certificate_authority", "test_certificate_authority_path", acctest.Required, acctest.Create, acctest.RepresentationCopyWithNewProperties(certificateAuthorityRepresentationRoot, map[string]interface{}{
		"certificate_authority_rules": acctest.RepresentationGroup{RepType: acctest.Required, Group: certificateAuthorityCertificateAuthorityRules_IssuanceRuleRepresentation},
		"description":                 acctest.Representation{RepType: acctest.Required, Create: `description_create`},
	}))

	CertificatesManagementCertificateAuthorityPathLengthTestUpdateResource = acctest.GenerateResourceFromRepresentationMap("oci_certificates_management_certificate_authority", "test_certificate_authority", acctest.Required, acctest.Update, acctest.RepresentationCopyWithNewProperties(certificateAuthorityRepresentationRoot, map[string]interface{}{
		"certificate_authority_rules": acctest.RepresentationGroup{RepType: acctest.Required, Group: certificateAuthorityCertificateAuthorityRules_IssuanceRuleRepresentation},
		"description":                 acctest.Representation{RepType: acctest.Required, Update: `description_update`},
	}))

	CertificatesManagementCertificateAuthorityResourceConfig = CertificatesManagementCertificateAuthorityResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_certificates_management_certificate_authority", "test_certificate_authority", acctest.Optional, acctest.Update, certificateAuthorityRepresentationSub)

	CertificatesManagementcertificateAuthoritySingularDataSourceRepresentation = map[string]interface{}{
		"certificate_authority_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_certificates_management_certificate_authority.test_certificate_authority.id}`},
	}

	certificateAuthorityDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"name":           acctest.Representation{RepType: acctest.Optional, Create: subCAManagedInternallyIssuedExternally},
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
		"certificate_authority_rules":         acctest.RepresentationGroup{RepType: acctest.Optional, Group: certificateAuthorityCertificateAuthorityRules_ExpiryRuleRepresentation},
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
		"certificate_authority_rules":         acctest.RepresentationGroup{RepType: acctest.Optional, Group: certificateAuthorityCertificateAuthorityRules_ExpiryRuleRepresentation},
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

	/** ROOT MANAGED EXTERNALLY CONFIGS **/

	certificateAuthorityConfigRepresentationRootManagedExternally_V1 = map[string]interface{}{
		"config_type":     acctest.Representation{RepType: acctest.Required, Create: `ROOT_CA_MANAGED_EXTERNALLY`},
		"certificate_pem": acctest.Representation{RepType: acctest.Required, Create: `${var.root_certificate_pem}`},
		"version_name":    acctest.Representation{RepType: acctest.Required, Create: `version_1`},
	}

	certificateAuthorityConfigRepresentationRootManagedExternally_V2 = map[string]interface{}{
		"config_type":     acctest.Representation{RepType: acctest.Required, Create: `ROOT_CA_MANAGED_EXTERNALLY`},
		"certificate_pem": acctest.Representation{RepType: acctest.Required, Create: `${var.root_certificate_pem}`},
		"version_name":    acctest.Representation{RepType: acctest.Required, Create: `version_2`},
	}

	certificateAuthorityRepresentationRootManagedExternally = map[string]interface{}{
		"certificate_authority_config": acctest.RepresentationGroup{RepType: acctest.Required, Group: certificateAuthorityConfigRepresentationRootManagedExternally_V1},
		"compartment_id":               acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"name":                         acctest.Representation{RepType: acctest.Required, Create: externallyManagedRootCA},
		"certificate_authority_rules":  acctest.RepresentationGroup{RepType: acctest.Optional, Group: certificateAuthorityCertificateAuthorityRules_ExpiryRuleRepresentation},
		"defined_tags":                 acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":                  acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"external_key_description":     acctest.Representation{RepType: acctest.Required, Create: `external_key_root_description`, Update: `external_key_root_description2`},
		"freeform_tags":                acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"lifecycle":                    acctest.RepresentationGroup{RepType: acctest.Required, Group: certificatesManagementIgnoreChangesRepresentation},
	}

	/** SUBORDINATE CONFIGS **/
	certificateAuthorityRepresentationSub = map[string]interface{}{
		"certificate_authority_config":        acctest.RepresentationGroup{RepType: acctest.Required, Group: certificateAuthorityCertificateAuthorityConfigRepresentationSub},
		"compartment_id":                      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"kms_key_id":                          acctest.Representation{RepType: acctest.Required, Create: cmKmsKeyId},
		"name":                                acctest.Representation{RepType: acctest.Required, Create: subCAManagedInternallyIssuedExternally},
		"certificate_authority_rules":         acctest.RepresentationGroup{RepType: acctest.Optional, Group: certificateAuthorityCertificateAuthorityRules_ExpiryRuleRepresentation},
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

	/** SUBORDINATE MANAGED INTERNALLY ISSUED BY EXTERNAL CA CONFIGS **/
	certificateAuthorityRepresentationSubManagedInternally = map[string]interface{}{
		"certificate_authority_config": acctest.RepresentationGroup{RepType: acctest.Required, Group: certificateAuthorityCertificateAuthorityConfigRepresentationSubManagedInternally},
		"compartment_id":               acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"kms_key_id":                   acctest.Representation{RepType: acctest.Required, Create: cmKmsKeyId},
		"name":                         acctest.Representation{RepType: acctest.Required, Create: subCAManagedInternallyIssuedExternallyName},
		"certificate_authority_rules":  acctest.RepresentationGroup{RepType: acctest.Optional, Group: certificateAuthorityCertificateAuthorityRules_ExpiryRuleCreateRepresentation},
		"lifecycle":                    acctest.RepresentationGroup{RepType: acctest.Required, Group: certificatesManagementIgnoreChangesRepresentation},
	}

	certificateAuthorityCertificateAuthorityConfigRepresentationSubManagedInternally = map[string]interface{}{
		"config_type":                     acctest.Representation{RepType: acctest.Required, Create: `SUBORDINATE_CA_MANAGED_INTERNALLY_ISSUED_BY_EXTERNAL_CA`},
		"subject":                         acctest.RepresentationGroup{RepType: acctest.Required, Group: certificatesManagementCertificateSubjectRepresentation},
		"issuer_certificate_authority_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_certificates_management_certificate_authority.test_certificate_authority_root_managed_externally.id}`},
		"signing_algorithm":               acctest.Representation{RepType: acctest.Optional, Create: `SHA256_WITH_RSA`},
		"version_name":                    acctest.Representation{RepType: acctest.Optional, Create: `versionName`, Update: `versionName2`},
		"action_details":                  acctest.RepresentationGroup{RepType: acctest.Required, Group: certificatesManagementSubCaActionDetailsUpdateCertRepresentation},
	}

	certificatesManagementSubCaActionDetailsUpdateCertRepresentation = map[string]interface{}{
		"action_type":     acctest.Representation{RepType: acctest.Required, Update: `UPDATE_CERTIFICATE`},
		"certificate_pem": acctest.Representation{RepType: acctest.Required, Update: `${var.signed_cert_pem}`},
	}

	certificateAuthorityCertificateAuthorityGenerateCsrConfigRepresentationSubManagedInternally = map[string]interface{}{
		"config_type":                     acctest.Representation{RepType: acctest.Required, Create: `SUBORDINATE_CA_MANAGED_INTERNALLY_ISSUED_BY_EXTERNAL_CA`},
		"subject":                         acctest.RepresentationGroup{RepType: acctest.Required, Group: certificatesManagementCertificateSubjectRepresentation},
		"issuer_certificate_authority_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_certificates_management_certificate_authority.test_certificate_authority_root_managed_externally.id}`},
		"signing_algorithm":               acctest.Representation{RepType: acctest.Optional, Create: `SHA256_WITH_RSA`},
		"version_name":                    acctest.Representation{RepType: acctest.Optional, Create: `versionName`, Update: `versionName2`},
		"action_details":                  acctest.RepresentationGroup{RepType: acctest.Required, Group: certificatesManagementSubCaActionDetailsGenerateCsrRepresentation},
	}

	certificatesManagementSubCaActionDetailsGenerateCsrRepresentation = map[string]interface{}{
		"action_type": acctest.Representation{RepType: acctest.Required, Update: `GENERATE_CSR`},
	}

	certificateAuthorityVersionDataSourceRepresentationSubManagedInternally = map[string]interface{}{
		"certificate_authority_id":             acctest.Representation{RepType: acctest.Required, Create: `${oci_certificates_management_certificate_authority.test_certificate_authority_sub_byoca.id}`},
		"certificate_authority_version_number": acctest.Representation{RepType: acctest.Optional, Create: `1`},
	}

	certificateAuthorityVersionDataSourceRepresentationSubManagedInternallyV2 = map[string]interface{}{
		"certificate_authority_id":             acctest.Representation{RepType: acctest.Required, Create: `${oci_certificates_management_certificate_authority.test_certificate_authority_sub_byoca.id}`},
		"certificate_authority_version_number": acctest.Representation{RepType: acctest.Optional, Create: `2`},
	}

	certificateAuthorityVersionDataSourceRepresentationRootManagedExternallyV2 = map[string]interface{}{
		"certificate_authority_id":             acctest.Representation{RepType: acctest.Required, Create: `${oci_certificates_management_certificate_authority.test_certificate_authority_root_managed_externally.id}`},
		"certificate_authority_version_number": acctest.Representation{RepType: acctest.Optional, Create: `2`},
	}

	certificateAuthorityCertificateAuthorityRules_ExpiryRuleRepresentation = map[string]interface{}{
		"rule_type": acctest.Representation{RepType: acctest.Required, Create: `CERTIFICATE_AUTHORITY_ISSUANCE_EXPIRY_RULE`, Update: `CERTIFICATE_AUTHORITY_ISSUANCE_EXPIRY_RULE`},
		"certificate_authority_max_validity_duration": acctest.Representation{RepType: acctest.Optional, Create: `P1000D`, Update: `P1500D`},
		"leaf_certificate_max_validity_duration":      acctest.Representation{RepType: acctest.Optional, Create: `P500D`, Update: `P400D`},
	}

	certificateAuthorityCertificateAuthorityRules_ExpiryRuleCreateRepresentation = map[string]interface{}{
		"rule_type": acctest.Representation{RepType: acctest.Required, Create: `CERTIFICATE_AUTHORITY_ISSUANCE_EXPIRY_RULE`},
		"certificate_authority_max_validity_duration": acctest.Representation{RepType: acctest.Optional, Create: `P1000D`},
		"leaf_certificate_max_validity_duration":      acctest.Representation{RepType: acctest.Optional, Create: `P500D`},
	}

	certificateAuthorityCertificateAuthorityRules_IssuanceRuleRepresentation = map[string]interface{}{
		"rule_type":              acctest.Representation{RepType: acctest.Required, Create: `CERTIFICATE_AUTHORITY_ISSUANCE_RULE`},
		"path_length_constraint": acctest.Representation{RepType: acctest.Required, Create: `5`},
		"name_constraint":        acctest.RepresentationGroup{RepType: acctest.Required, Group: nameConstraintRepresentation},
	}

	certificateAuthorityCertificateAuthorityRules_IssuanceRuleRepresentationDiffSuppressCase = map[string]interface{}{
		"rule_type":       acctest.Representation{RepType: acctest.Required, Create: `CERTIFICATE_AUTHORITY_ISSUANCE_RULE`},
		"name_constraint": acctest.RepresentationGroup{RepType: acctest.Required, Group: nameConstraintRepresentationAddOn},
	}

	byocaRootCertificateAuthorityCertificateAuthorityRules_IssuanceRulePathLengthRepresentation = map[string]interface{}{
		"rule_type":              acctest.Representation{RepType: acctest.Required, Create: `CERTIFICATE_AUTHORITY_ISSUANCE_RULE`},
		"path_length_constraint": acctest.Representation{RepType: acctest.Required, Create: `10`},
		"name_constraint":        acctest.RepresentationGroup{RepType: acctest.Required, Group: byocaNameConstraintRepresentation},
	}

	byocaSubCertificateAuthorityCertificateAuthorityRules_IssuanceRulePathLengthRepresentation = map[string]interface{}{
		"rule_type":              acctest.Representation{RepType: acctest.Required, Create: `CERTIFICATE_AUTHORITY_ISSUANCE_RULE`},
		"path_length_constraint": acctest.Representation{RepType: acctest.Required, Create: `9`},
		"name_constraint":        acctest.RepresentationGroup{RepType: acctest.Required, Group: byocaNameConstraintRepresentation},
	}

	subCertificateAuthorityCertificateAuthorityRules_IssuanceRuleRepresentation = map[string]interface{}{
		"rule_type":              acctest.Representation{RepType: acctest.Required, Create: `CERTIFICATE_AUTHORITY_ISSUANCE_RULE`},
		"path_length_constraint": acctest.Representation{RepType: acctest.Required, Create: `4`},
		"name_constraint":        acctest.RepresentationGroup{RepType: acctest.Required, Group: nameConstraintRepresentation},
	}

	certificateAuthorityCertificateRevocationListDetailsRepresentation = map[string]interface{}{
		"object_storage_config": acctest.RepresentationGroup{RepType: acctest.Required, Group: certificateAuthorityCertificateRevocationListDetailsObjectStorageConfigRepresentation},
		"custom_formatted_urls": acctest.Representation{RepType: acctest.Optional, Create: []string{`https://www.example.com/crl.bin`}},
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
		"user_id":                      acctest.Representation{RepType: acctest.Optional, Create: `Neytiri`},
	}
	certificateAuthorityCertificateAuthorityConfigValidityRepresentation = map[string]interface{}{
		"time_of_validity_not_after":  acctest.Representation{RepType: acctest.Optional, Create: notAfterCreate, Update: notAfterUpdate},
		"time_of_validity_not_before": acctest.Representation{RepType: acctest.Optional, Create: notBeforeCreate, Update: notBeforeUpdate},
	}

	certificateAuthorityCertificateRevocationListDetailsObjectStorageConfigRepresentation = map[string]interface{}{
		"object_storage_bucket_name":        acctest.Representation{RepType: acctest.Required, Create: `${oci_objectstorage_bucket.test_bucket.name}`},
		"object_storage_object_name_format": acctest.Representation{RepType: acctest.Required, Create: subCAManagedInternallyIssuedExternally + "versionName{}.crl"},
		"object_storage_namespace":          acctest.Representation{RepType: acctest.Optional, Create: `${oci_objectstorage_bucket.test_bucket.namespace}`},
	}

	certificatesManagementIgnoreChangesRepresentation = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`}},
	}

	certificatesManagementCertificateAuthorityCertificateAuthorityRulesNameConstraintExcludedSubtreeRepresentation = map[string]interface{}{
		"type":  acctest.Representation{RepType: acctest.Required, Create: `DNS`},
		"value": acctest.Representation{RepType: acctest.Required, Create: `exampledenied.com`},
	}

	certificatesManagementCertificateAuthorityCertificateAuthorityRulesNameConstraintExcludedSubtreeRepresentationAddOn = map[string]interface{}{
		"type":  acctest.Representation{RepType: acctest.Required, Create: `DNS`},
		"value": acctest.Representation{RepType: acctest.Required, Create: `examplebaddenied.com`},
	}

	certificatesManagementCertificateAuthorityCertificateAuthorityRulesNameConstraintPermittedSubtreeRepresentation = map[string]interface{}{
		"type":  acctest.Representation{RepType: acctest.Required, Create: `DNS`},
		"value": acctest.Representation{RepType: acctest.Required, Create: `example.com`},
	}

	certificatesManagementCertificateAuthorityCertificateAuthorityRulesNameConstraintPermittedSubtreeRepresentationAddOn = map[string]interface{}{
		"type":  acctest.Representation{RepType: acctest.Required, Create: `DNS`},
		"value": acctest.Representation{RepType: acctest.Required, Create: `sub.example.com`},
	}

	nameConstraintRepresentation = map[string]interface{}{
		"permitted_subtree": acctest.RepresentationGroup{RepType: acctest.Required, Group: certificatesManagementCertificateAuthorityCertificateAuthorityRulesNameConstraintPermittedSubtreeRepresentation},
		"excluded_subtree":  acctest.RepresentationGroup{RepType: acctest.Required, Group: certificatesManagementCertificateAuthorityCertificateAuthorityRulesNameConstraintExcludedSubtreeRepresentation},
	}

	nameConstraintRepresentationAddOn = map[string]interface{}{
		"permitted_subtree": acctest.RepresentationGroup{RepType: acctest.Required, Group: certificatesManagementCertificateAuthorityCertificateAuthorityRulesNameConstraintPermittedSubtreeRepresentationAddOn},
		"excluded_subtree":  acctest.RepresentationGroup{RepType: acctest.Required, Group: certificatesManagementCertificateAuthorityCertificateAuthorityRulesNameConstraintExcludedSubtreeRepresentationAddOn},
	}

	byocaCertificatesManagementCertificateAuthorityCertificateAuthorityRulesNameConstraintExcludedSubtreeRepresentation = map[string]interface{}{
		"type":  acctest.Representation{RepType: acctest.Required, Create: `DNS`},
		"value": acctest.Representation{RepType: acctest.Required, Create: `bad.com`},
	}
	byocaCertificatesManagementCertificateAuthorityCertificateAuthorityRulesNameConstraintPermittedSubtreeRepresentation = map[string]interface{}{
		"type":  acctest.Representation{RepType: acctest.Required, Create: `DNS`},
		"value": acctest.Representation{RepType: acctest.Required, Create: `example.com`},
	}

	byocaNameConstraintRepresentation = map[string]interface{}{
		"permitted_subtree": acctest.RepresentationGroup{RepType: acctest.Required, Group: byocaCertificatesManagementCertificateAuthorityCertificateAuthorityRulesNameConstraintPermittedSubtreeRepresentation},
		"excluded_subtree":  acctest.RepresentationGroup{RepType: acctest.Required, Group: byocaCertificatesManagementCertificateAuthorityCertificateAuthorityRulesNameConstraintExcludedSubtreeRepresentation},
	}

	CertificatesManagementCertificateAuthorityResourceDependencies = DefinedTagsDependencies +
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
					resource.TestCheckResourceAttr(resourceName, "certificate_revocation_list_details.0.object_storage_config.0.object_storage_object_name_format", subCAManagedInternallyIssuedExternally+"versionName{}.crl"),
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
					resource.TestCheckResourceAttr(resourceName, "name", subCAManagedInternallyIssuedExternally),
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
					resource.TestCheckResourceAttr(resourceName, "certificate_revocation_list_details.0.object_storage_config.0.object_storage_object_name_format", subCAManagedInternallyIssuedExternally+"versionName{}.crl"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
					resource.TestCheckResourceAttrSet(resourceName, "config_type"),
					resource.TestCheckResourceAttr(resourceName, "description", "description"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "kms_key_id"),
					resource.TestCheckResourceAttr(resourceName, "name", subCAManagedInternallyIssuedExternally),
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
					resource.TestCheckResourceAttr(resourceName, "certificate_revocation_list_details.0.object_storage_config.0.object_storage_object_name_format", subCAManagedInternallyIssuedExternally+"versionName{}.crl"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(resourceName, "config_type"),
					resource.TestCheckResourceAttr(resourceName, "description", "description2"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "kms_key_id"),
					resource.TestCheckResourceAttr(resourceName, "name", subCAManagedInternallyIssuedExternally),
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
					resource.TestCheckResourceAttr(datasourceName, "name", subCAManagedInternallyIssuedExternally),
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
					resource.TestCheckResourceAttr(singularDatasourceName, "certificate_revocation_list_details.0.object_storage_config.0.object_storage_namespace", "id4jtm77o185"),
					resource.TestCheckResourceAttr(singularDatasourceName, "certificate_revocation_list_details.0.object_storage_config.0.object_storage_object_name_format", subCAManagedInternallyIssuedExternally+"versionName{}.crl"),
					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "config_type"),
					resource.TestCheckResourceAttr(singularDatasourceName, "current_version.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
					resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "name", subCAManagedInternallyIssuedExternally),
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

func TestCertificatesManagementCertificateAuthorityResource_IssuanceRuleTests(t *testing.T) {
	httpreplay.SetScenario("TestCertificatesManagementCertificateAuthorityResource_IssuanceRuleTests")
	defer httpreplay.SaveScenario()

	provider := acctest.TestAccProvider
	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_certificates_management_certificate_authority.test_certificate_authority_path"
	subCAPathLengthResourceName := "oci_certificates_management_certificate_authority.test_sub_certificate_authority"

	var resId, resId2 string
	// Save TF content to create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+CertificatesManagementCertificateAuthorityResourceDependencies, "certificatesmanagement", "certificateAuthority", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { acctest.PreCheck(t) },
		Providers: map[string]*schema.Provider{
			"oci": provider,
		},
		Steps: []resource.TestStep{

			// verify create with ca_issuance_rules -> config will match the service returned value so no diff
			{
				Config: config + compartmentIdVariableStr + CertificatesManagementCertificateAuthorityResourceDependencies + CertificatesManagementCertificateAuthorityPathLengthTestResource,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "certificate_authority_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "certificate_authority_config.0.config_type", "ROOT_CA_GENERATED_INTERNALLY"),
					resource.TestCheckResourceAttr(resourceName, "certificate_authority_config.0.subject.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "certificate_authority_config.0.subject.0.common_name", "www.example.com"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(resourceName, "kms_key_id"),
					resource.TestCheckResourceAttr(resourceName, "name", rootCaName),
					resource.TestCheckResourceAttr(resourceName, "certificate_authority_rules.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "certificate_authority_rules.0.path_length_constraint", "5"),
					resource.TestCheckResourceAttr(resourceName, "certificate_authority_rules.0.rule_type", "CERTIFICATE_AUTHORITY_ISSUANCE_RULE"),
					resource.TestCheckResourceAttr(resourceName, "certificate_authority_rules.0.name_constraint.#", "1"),
					func(s *terraform.State) (err error) {
						resId, err = acctest.FromInstanceState(s, resourceName, "id")
						fmt.Printf("variable \"resId\" { default = \"%s\" }\n", resId)
						return err
					},
				),
			},
			//step 2: setup to verify when sub ca doesn't have a ca_issuance_rule but root has it then next steps plan change should not show that as a diff
			{

				Config: config + compartmentIdVariableStr + CertificatesManagementCertificateAuthorityResourceDependencies + CertificatesManagementCertificateAuthorityPathLengthTestResource + acctest.GenerateResourceFromRepresentationMap("oci_certificates_management_certificate_authority", "test_sub_certificate_authority", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(certificateAuthorityRepresentationSub, map[string]interface{}{
						"certificate_authority_config": acctest.RepresentationGroup{
							RepType: acctest.Required, Group: map[string]interface{}{
								"config_type":                     acctest.Representation{RepType: acctest.Required, Create: `SUBORDINATE_CA_ISSUED_BY_INTERNAL_CA`},
								"subject":                         acctest.RepresentationGroup{RepType: acctest.Required, Group: certificatesManagementCertificateSubjectRepresentation},
								"issuer_certificate_authority_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_certificates_management_certificate_authority.test_certificate_authority_path.id}`},
								"signing_algorithm":               acctest.Representation{RepType: acctest.Optional, Create: `SHA256_WITH_RSA`},
								"validity":                        acctest.RepresentationGroup{RepType: acctest.Optional, Group: certificateAuthorityCertificateAuthorityConfigValidityRepresentation},
								"version_name":                    acctest.Representation{RepType: acctest.Optional, Create: `versionNamePathLength`},
							},
						},
					})),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(subCAPathLengthResourceName, "certificate_authority_config.#", "1"),
					resource.TestCheckResourceAttr(subCAPathLengthResourceName, "certificate_authority_config.0.config_type", "SUBORDINATE_CA_ISSUED_BY_INTERNAL_CA"),
					resource.TestCheckResourceAttrSet(subCAPathLengthResourceName, "certificate_authority_config.0.issuer_certificate_authority_id"),
					resource.TestCheckResourceAttr(subCAPathLengthResourceName, "certificate_authority_config.0.signing_algorithm", "SHA256_WITH_RSA"),
					resource.TestCheckResourceAttr(subCAPathLengthResourceName, "certificate_authority_config.0.subject.#", "1"),
					resource.TestCheckResourceAttr(subCAPathLengthResourceName, "certificate_authority_config.0.subject.0.common_name", "www.example.com"),
					resource.TestCheckResourceAttr(subCAPathLengthResourceName, "certificate_authority_rules.#", "2"),
					resource.TestCheckResourceAttr(subCAPathLengthResourceName, "certificate_authority_rules.0.rule_type", "CERTIFICATE_AUTHORITY_ISSUANCE_EXPIRY_RULE"),
					resource.TestCheckResourceAttr(subCAPathLengthResourceName, "certificate_authority_rules.0.certificate_authority_max_validity_duration", "P1000D"),
					resource.TestCheckResourceAttr(subCAPathLengthResourceName, "certificate_authority_rules.0.leaf_certificate_max_validity_duration", "P500D"),
					resource.TestCheckResourceAttr(subCAPathLengthResourceName, "certificate_authority_rules.1.rule_type", "CERTIFICATE_AUTHORITY_ISSUANCE_RULE"),
					resource.TestCheckResourceAttr(subCAPathLengthResourceName, "certificate_authority_rules.1.path_length_constraint", "4"),
					resource.TestCheckResourceAttr(subCAPathLengthResourceName, "certificate_authority_rules.1.name_constraint.#", "1"),
					func(s *terraform.State) (err error) {
						resId, err = acctest.FromInstanceState(s, subCAPathLengthResourceName, "id")
						return err
					},
				),
				ExpectNonEmptyPlan: false,
			},
			//Step 3: Update of CA rules; Try removing expiry rule and retain issuance rule verifies if user updates the expiry rule config the update should go through as expected with issuance rule getting injected by the provider
			{
				Config: config + compartmentIdVariableStr + CertificatesManagementCertificateAuthorityResourceDependencies + CertificatesManagementCertificateAuthorityPathLengthTestResource + acctest.GenerateResourceFromRepresentationMap("oci_certificates_management_certificate_authority", "test_sub_certificate_authority", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(certificateAuthorityRepresentationSub, map[string]interface{}{
						"certificate_authority_config": acctest.RepresentationGroup{RepType: acctest.Required, Group: map[string]interface{}{
							"config_type":                     acctest.Representation{RepType: acctest.Required, Create: `SUBORDINATE_CA_ISSUED_BY_INTERNAL_CA`},
							"subject":                         acctest.RepresentationGroup{RepType: acctest.Required, Group: certificatesManagementCertificateSubjectRepresentation},
							"issuer_certificate_authority_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_certificates_management_certificate_authority.test_certificate_authority_path.id}`},
							"signing_algorithm":               acctest.Representation{RepType: acctest.Optional, Create: `SHA256_WITH_RSA`},
							"validity":                        acctest.RepresentationGroup{RepType: acctest.Optional, Group: certificateAuthorityCertificateAuthorityConfigValidityRepresentation},
							"version_name":                    acctest.Representation{RepType: acctest.Optional, Create: `versionNamePathLengthUpdate`}}},
						"certificate_authority_rules": acctest.RepresentationGroup{RepType: acctest.Optional, Group: map[string]interface{}{
							"rule_type": acctest.Representation{RepType: acctest.Required, Create: `CERTIFICATE_AUTHORITY_ISSUANCE_EXPIRY_RULE`, Update: `CERTIFICATE_AUTHORITY_ISSUANCE_EXPIRY_RULE`},
							"certificate_authority_max_validity_duration": acctest.Representation{RepType: acctest.Optional, Create: `P100D`, Update: `P1500D`},
							"leaf_certificate_max_validity_duration":      acctest.Representation{RepType: acctest.Optional, Create: `P50D`, Update: `P400D`},
						}},
					})),

				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(subCAPathLengthResourceName, "certificate_authority_config.#", "1"),
					resource.TestCheckResourceAttr(subCAPathLengthResourceName, "certificate_authority_config.0.config_type", "SUBORDINATE_CA_ISSUED_BY_INTERNAL_CA"),
					resource.TestCheckResourceAttr(subCAPathLengthResourceName, "certificate_authority_rules.#", "2"),
					resource.TestCheckResourceAttr(subCAPathLengthResourceName, "certificate_authority_rules.1.path_length_constraint", "4"),
					resource.TestCheckResourceAttr(subCAPathLengthResourceName, "certificate_authority_rules.1.name_constraint.#", "1"),
					resource.TestCheckResourceAttr(subCAPathLengthResourceName, "certificate_authority_rules.1.rule_type", "CERTIFICATE_AUTHORITY_ISSUANCE_RULE"),
					resource.TestCheckResourceAttr(subCAPathLengthResourceName, "certificate_authority_rules.1.certificate_authority_max_validity_duration", ""),
					resource.TestCheckResourceAttr(subCAPathLengthResourceName, "certificate_authority_rules.1.leaf_certificate_max_validity_duration", ""),
					resource.TestCheckResourceAttr(subCAPathLengthResourceName, "certificate_authority_rules.0.path_length_constraint", "0"),
					resource.TestCheckResourceAttr(subCAPathLengthResourceName, "certificate_authority_rules.0.name_constraint.#", "0"),
					resource.TestCheckResourceAttr(subCAPathLengthResourceName, "certificate_authority_rules.0.rule_type", "CERTIFICATE_AUTHORITY_ISSUANCE_EXPIRY_RULE"),
					resource.TestCheckResourceAttr(subCAPathLengthResourceName, "certificate_authority_rules.0.certificate_authority_max_validity_duration", "P100D"),
					resource.TestCheckResourceAttr(subCAPathLengthResourceName, "certificate_authority_rules.0.leaf_certificate_max_validity_duration", "P50D"),
					func(s *terraform.State) (err error) {
						resId2, err = acctest.FromInstanceState(s, subCAPathLengthResourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("Resource recreated when it was supposed to be updated.")
						}
						return err
					},
				),
			},
		},
	})
}

func TestCertificatesManagementCertificateAuthorityResource_BYOCA(t *testing.T) {
	httpreplay.SetScenario("TestCertificatesManagementCertificateAuthorityResource_BYOCA")
	defer httpreplay.SaveScenario()

	provider := acctest.TestAccProvider
	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	rootCertificatePem := utils.GetEnvSettingWithBlankDefault("root_managed_externally_cert_pem")
	rootCertificatePemVariableStr := fmt.Sprintf("variable \"root_certificate_pem\" { default = \"%s\" }\n", rootCertificatePem)

	//We should be good until 2035 but If subCsr changes then we have to resign the csr and update the signedSubCertPemVariableStr
	subCsrPem := "-----BEGIN CERTIFICATE REQUEST-----\nMIIHhDCCBmwCAQAwggFiMRgwFgYDVQQDDA93d3cuZXhhbXBsZS5jb20xCzAJBgNV\nBAYTAnVzMSMwIQYDVQQuExpkaXN0aW5ndWlzaGVkTmFtZVF1YWxpZmllcjEfMB0G\nCgmSJomT8ixkARkWD2RvbWFpbkNvbXBvbmVudDELMAkGA1UELAwCSlIxDDAKBgNV\nBCoMA1NpcjEMMAoGA1UEKwwDSEFNMRAwDgYDVQQHDAdTZWF0dGxlMQwwCgYDVQQK\nDANPQ0kxGTAXBgNVBAsMEFNlY3VyaXR5UHJvZHVjdHMxEjAQBgNVBEEMCXBzZXVk\nb255bTEVMBMGA1UEBRMMc2VyaWFsTnVtYmVyMRMwEQYDVQQIDApXYXNoaW5ndG9u\nMRgwFgYDVQQJDA8xMjMgTWFpbiBTdHJlZXQxDTALBgNVBAQMBExhc3QxDTALBgNV\nBAwMBExvcmQxFzAVBgoJkiaJk/IsZAEBDAdOZXl0aXJpMIIBIjANBgkqhkiG9w0B\nAQEFAAOCAQ8AMIIBCgKCAQEAsy1DjkwGWHcuqvgB4i+iJ4exABSGZs8F5kKvDqoe\ngoJ+myHlxppyyGdLnIqrhnuj0xMEjDf33iJn5PISg9ry+EzZzhdqrahTJoaixqvt\n4nA29BYNA7Y8SgVchxUeuMibUJCmX8bVlmjp32t46rR9xZJ6IenBlygdKOmBFIzT\nPz9ROZ+hpTUYGxAR4JrwRdbOyDRDExfToOspr0CbDdrNspBTX7qy94mmNrr0WfPw\nCjf69FQYDZNVgJhaFAfnLH5bkCxePPueEq4ZNQpjoGchwbi5NpsOUcS9sHMyzns7\nq6Cxk2ws2tTHOX8CxQT5ewXW1I2HNtXIT7TG49Jlrm0iMwIDAQABoIID2TCCA9UG\nCSqGSIb3DQEJDjGCA8YwggPCMA4GA1UdDwEB/wQEAwICBDASBgNVHRMBAf8ECDAG\nAQH/AgEJMCoGA1UdHgEB/wQgMB6gDzANggtleGFtcGxlLmNvbaELMAmCB2JhZC5j\nb20wggI3BgNVHSMEggIuMIICKoCCAiYwggIiMA0GCSqGSIb3DQEBAQUAA4ICDwAw\nggIKAoICAQC5AYaf0Mu2x9go+egvO+FGNkGmr2N4ZnnlFfjaaSFsU+LFQ6cr3je/\nv++tRpjMrqv+QWnINUixyXzxToCzQhz5ZF+9lAHjvwlo4bQ8UNeEpP966RjrgqPW\nburAdHX7c1UsbkM3esdGTbF0QqBdhGRt39ZswToS5OemAQmYbbsw6VaKDNiGe+iH\ndCGHZGFTX2PZ+zvdnjkBfy4fAfemo1QTwtLYt25YOUrE3/oQ7w9WMeFmQa/5SuIt\n/pw8bj3O/qGoxvF1uJQ5cq2LdoWLiZr6BCoUKWXcKSCo+3XGnD9bcpOB2XUmuOhM\nSs5f+JK0yA18iMXkz4jyQYv4nIq8C2wbTLU6Vxg/untvmOC40a0U8x8lH74TX6Ts\nhc35OkyGFvRk2mZNGf52w34G7gX6dtLWFsm8M80qiwenwD5XNEZZKnD4smCOKcny\neYKgsG0RXmvi1Eb2YHTtEFxVUJcJHBAFeAuM7t7Ohe+MUe52WP+zSkIbGZ7hBHDk\nM/py3lL9bSBhYXG8OFziiu/Idba/Cb2JCcZajJCGr9Oh9WN+RG0x/WzeR1aOf0nT\nnHt3GrP10POflkDar/kj2R/B1zo2Cvr3wLJUB/ggZndfeVq5L3ttWVg4NI/jqtUw\nmZX+hnX5vPjOJLgbq+uVMJ29RngW4E/CgyqfIdL0sgqg6DR0zmXhgwIDAQABMIIB\nMwYDVR0OBIIBKgSCASYwggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQCz\nLUOOTAZYdy6q+AHiL6Inh7EAFIZmzwXmQq8Oqh6Cgn6bIeXGmnLIZ0uciquGe6PT\nEwSMN/feImfk8hKD2vL4TNnOF2qtqFMmhqLGq+3icDb0Fg0DtjxKBVyHFR64yJtQ\nkKZfxtWWaOnfa3jqtH3Fknoh6cGXKB0o6YEUjNM/P1E5n6GlNRgbEBHgmvBF1s7I\nNEMTF9Og6ymvQJsN2s2ykFNfurL3iaY2uvRZ8/AKN/r0VBgNk1WAmFoUB+csfluQ\nLF48+54Srhk1CmOgZyHBuLk2mw5RxL2wczLOezuroLGTbCza1Mc5fwLFBPl7BdbU\njYc21chPtMbj0mWubSIzAgMBAAEwDQYJKoZIhvcNAQELBQADggEBAF0JNLksYahf\nugvrpY6mwaguBleUzopM23ktgRzAPAZiuLBh4mddi7E+kLq9v6dwxuOsc4371qy9\nivDzo57ih+fK728zzh4o2CH595tC5mWZl3kzFrlFYe0Fk0/1a7ibReQUqpS2kWpF\naxmAPIB561ik5rnWzHEeeulzf6MUvlE08Y+TpxFyMk7toJSzB9vYsoKKZdRLlwh8\nw+IyxWYzKcg31T8hkQ+n5zABZiSuFsual0B/pizjmhbgT4ccz33f/Dtcs4zfxWrj\nHc9dHyMwL2yhPJXKIydNgYIeMFzFVGzed5m9AM/hE1W/kQBhePdeBDUW995xxccq\nUREkYpbRA08=\n-----END CERTIFICATE REQUEST-----\n"
	signedSubCertPemVariableStr := fmt.Sprintf("variable \"signed_cert_pem\" { default = \"%s\" }\n", "-----BEGIN CERTIFICATE-----MIIGADCCA+igAwIBAgIUeZ8AzcHZ19Qayr2IYSoUTD+Ydy4wDQYJKoZIhvcNAQELBQAwYjELMAkGA1UEBhMCVVMxDjAMBgNVBAgMBVN0YXRlMQ0wCwYDVQQHDARDaXR5MQ4wDAYDVQQKDAVNeU9yZzEPMA0GA1UECwwGTXlVbml0MRMwEQYDVQQDDApNeSBSb290IENBMB4XDTI1MDUwMTIzMDU0NFoXDTM0MDIwMzIzMDU0NFowggFiMRgwFgYDVQQDDA93d3cuZXhhbXBsZS5jb20xCzAJBgNVBAYTAnVzMSMwIQYDVQQuExpkaXN0aW5ndWlzaGVkTmFtZVF1YWxpZmllcjEfMB0GCgmSJomT8ixkARkWD2RvbWFpbkNvbXBvbmVudDELMAkGA1UELAwCSlIxDDAKBgNVBCoMA1NpcjEMMAoGA1UEKwwDSEFNMRAwDgYDVQQHDAdTZWF0dGxlMQwwCgYDVQQKDANPQ0kxGTAXBgNVBAsMEFNlY3VyaXR5UHJvZHVjdHMxEjAQBgNVBEEMCXBzZXVkb255bTEVMBMGA1UEBRMMc2VyaWFsTnVtYmVyMRMwEQYDVQQIDApXYXNoaW5ndG9uMRgwFgYDVQQJDA8xMjMgTWFpbiBTdHJlZXQxDTALBgNVBAQMBExhc3QxDTALBgNVBAwMBExvcmQxFzAVBgoJkiaJk/IsZAEBDAdOZXl0aXJpMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAsy1DjkwGWHcuqvgB4i+iJ4exABSGZs8F5kKvDqoegoJ+myHlxppyyGdLnIqrhnuj0xMEjDf33iJn5PISg9ry+EzZzhdqrahTJoaixqvt4nA29BYNA7Y8SgVchxUeuMibUJCmX8bVlmjp32t46rR9xZJ6IenBlygdKOmBFIzTPz9ROZ+hpTUYGxAR4JrwRdbOyDRDExfToOspr0CbDdrNspBTX7qy94mmNrr0WfPwCjf69FQYDZNVgJhaFAfnLH5bkCxePPueEq4ZNQpjoGchwbi5NpsOUcS9sHMyzns7q6Cxk2ws2tTHOX8CxQT5ewXW1I2HNtXIT7TG49Jlrm0iMwIDAQABo4GrMIGoMB0GA1UdDgQWBBQA55iUzZzchtKwH80IrYM2OA67rjAfBgNVHSMEGDAWgBTyriDG4J4TCGs6PSSd3ImWU7F0SDASBgNVHRMBAf8ECDAGAQH/AgEJMA4GA1UdDwEB/wQEAwIBhjAWBgNVHREEDzANggtleGFtcGxlLmNvbTAqBgNVHR4BAf8EIDAeoA8wDYILZXhhbXBsZS5jb22hCzAJggdiYWQuY29tMA0GCSqGSIb3DQEBCwUAA4ICAQB0peLu7qivTdcrHVuBxn3QDvLCdDcY9Vd2/O7JG3zpPK22qTAL74gx0ZYyBlhTZBKfB892VRMfyoARP0Bh6sqlTa8aPJaxL3ol1vtc+jvgKqCyLi4KcpI3rMexr4DHccebPCFzKZZ7oGT2NUHPKS9AdZPRdpNLSVgwdRWWL34qRIndPBU7e+/19MSTlrJr4OskWeeXIJynBOtCkX5pBflGhNOzzx3harVijTIzz4QDBtz6iazutBSS40GkQmLWRKvjOkGeiH/RDmjXyq5xUUMeyx5Y9kzkSXpynco+oTkxMwxBUGOYbjmAU6t7CRt5J5YMQrWe7ptdRivT1FvWu244BSft4DCFfJtCPlWrysw1x1D6Ll1dG21ZVjh4MQDCLbGm8T1ku5fNqUlVz24NQbtQXUWpTXremlktXRkYHj5x4Rqli5qN9o+48YPVBgI2DMFTwNOH8PQZ3TiCgEEk40RlHEnFrPLIT6F4ZHl+pbLyx8KGPLcYr8k6CVBw1oRwHlg5S4c/H654GZNwlL6f8QQyzvlwsq30cPX0NfNEy9TOSzSRPSU9FNttd6QNMnBNX8NvA8HDPXj51uleIv8h4KUKRIglnu0Jea8X9OvVw43uldpJ4Ys8tKIwz4nQ3iXt65w0xkDl8jVlQ4xDaN8jbOYhcSFqFQSAkgW2QKQ0C9zPxw==-----END CERTIFICATE-----")

	resourceName := "oci_certificates_management_certificate_authority.test_certificate_authority_root_managed_externally"
	rootCAV2DataSourceName := "data.oci_certificates_management_certificate_authority_version.test_certificate_authority_root_byoca_v2"
	subCAResourceName := "oci_certificates_management_certificate_authority.test_certificate_authority_sub_byoca"
	subCAV1DataSourceName := "data.oci_certificates_management_certificate_authority_version.test_certificate_authority_sub_byoca_v1"
	subCAV2DataSourceName := "data.oci_certificates_management_certificate_authority_version.test_certificate_authority_sub_byoca_v2"
	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { acctest.PreCheck(t) },
		Providers: map[string]*schema.Provider{
			"oci": provider,
		},
		Steps: []resource.TestStep{

			// verify create with ca rules populated from the input cert pem
			{
				Config: config + compartmentIdVariableStr + rootCertificatePemVariableStr + CertificatesManagementCertificateAuthorityResourceDependencies + CertificatesManagementRootManagedExternallyResource,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "certificate_authority_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "certificate_authority_config.0.config_type", "ROOT_CA_MANAGED_EXTERNALLY"),
					resource.TestCheckResourceAttr(resourceName, "certificate_authority_config.0.certificate_pem", rootCertificatePem),
					resource.TestCheckResourceAttr(resourceName, "external_key_description", "external_key_root_description"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "name", externallyManagedRootCA),
					resource.TestCheckResourceAttr(resourceName, "certificate_authority_rules.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "certificate_authority_rules.0.rule_type", "CERTIFICATE_AUTHORITY_ISSUANCE_EXPIRY_RULE"),
					resource.TestCheckResourceAttr(resourceName, "certificate_authority_rules.0.certificate_authority_max_validity_duration", "P1000D"),
					resource.TestCheckResourceAttr(resourceName, "certificate_authority_rules.0.leaf_certificate_max_validity_duration", "P500D"),
					resource.TestCheckResourceAttr(resourceName, "certificate_authority_rules.1.path_length_constraint", "10"), //rules value depends on input cert pem
					resource.TestCheckResourceAttr(resourceName, "certificate_authority_rules.1.name_constraint.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "certificate_authority_rules.1.rule_type", "CERTIFICATE_AUTHORITY_ISSUANCE_RULE"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(resourceName, "config_type"),
					resource.TestCheckResourceAttrSet(resourceName, "current_version.0.certificate_authority_id"),
					resource.TestCheckResourceAttrSet(resourceName, "current_version.0.serial_number"),
					resource.TestCheckResourceAttr(resourceName, "current_version.0.stages.#", "2"),
					resource.TestCheckResourceAttrSet(resourceName, "current_version.0.time_created"),
					resource.TestCheckResourceAttr(resourceName, "current_version.0.validity.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "current_version.0.validity.0.time_of_validity_not_after"),
					resource.TestCheckResourceAttrSet(resourceName, "current_version.0.validity.0.time_of_validity_not_before"),
					resource.TestCheckResourceAttrSet(resourceName, "current_version.0.version_number"),
					resource.TestCheckResourceAttrSet(resourceName, "signing_algorithm"),
					resource.TestCheckResourceAttrSet(resourceName, "subject.#"),
					resource.TestCheckResourceAttrSet(resourceName, "subject.0.common_name"),
					resource.TestCheckResourceAttr(resourceName, "external_key_description", "external_key_root_description"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "name", externallyManagedRootCA),
					resource.TestCheckResourceAttr(resourceName, "state", "ACTIVE"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
					func(s *terraform.State) (err error) {
						resId, err = acctest.FromInstanceState(s, resourceName, "id")
						fmt.Printf("variable \"resId\" { default = \"%s\" }\n", resId)
						return err
					},
				),
				ExpectNonEmptyPlan: false,
			},
			//delete the issuance_expiry rule while making sure newly added issuance_rule in config doesn't cause an issue
			//mimics user made update by adding issuance rule in config and deleting expiry rule
			{

				Config: config + compartmentIdVariableStr + rootCertificatePemVariableStr + CertificatesManagementCertificateAuthorityResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_certificates_management_certificate_authority", "test_certificate_authority_root_managed_externally", acctest.Optional, acctest.Create,
						acctest.RepresentationCopyWithNewProperties(certificateAuthorityRepresentationRootManagedExternally, map[string]interface{}{
							"certificate_authority_rules": acctest.RepresentationGroup{RepType: acctest.Required, Group: byocaRootCertificateAuthorityCertificateAuthorityRules_IssuanceRulePathLengthRepresentation},
						})),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", externallyManagedRootCA),
					resource.TestCheckResourceAttr(resourceName, "certificate_authority_rules.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "certificate_authority_rules.0.path_length_constraint", "10"), //rules value depends on input cert pem
					resource.TestCheckResourceAttr(resourceName, "certificate_authority_rules.0.name_constraint.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "certificate_authority_rules.0.rule_type", "CERTIFICATE_AUTHORITY_ISSUANCE_RULE"),

					func(s *terraform.State) (err error) {
						resId2, err = acctest.FromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("Resource recreated when it was supposed to be updated.")
						}
						return err
					},
				),
				ExpectNonEmptyPlan: false,
			},
			//add a sub managed internally issued by external CA and preserve the same root CA config from previous step
			//Get the CA version 1 to fetch the csr;
			{

				Config: config + compartmentIdVariableStr + rootCertificatePemVariableStr + CertificatesManagementCertificateAuthorityResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_certificates_management_certificate_authority", "test_certificate_authority_root_managed_externally", acctest.Optional, acctest.Create,
						acctest.RepresentationCopyWithNewProperties(certificateAuthorityRepresentationRootManagedExternally, map[string]interface{}{
							"certificate_authority_rules": acctest.RepresentationGroup{RepType: acctest.Required, Group: byocaRootCertificateAuthorityCertificateAuthorityRules_IssuanceRulePathLengthRepresentation},
						})) + acctest.GenerateResourceFromRepresentationMap("oci_certificates_management_certificate_authority", "test_certificate_authority_sub_byoca", acctest.Optional, acctest.Create, certificateAuthorityRepresentationSubManagedInternally) +
					acctest.GenerateDataSourceFromRepresentationMap("oci_certificates_management_certificate_authority_version", "test_certificate_authority_sub_byoca_v1", acctest.Optional, acctest.Create, certificateAuthorityVersionDataSourceRepresentationSubManagedInternally),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(subCAResourceName, "name", subCAManagedInternallyIssuedExternallyName),
					resource.TestCheckResourceAttr(subCAResourceName, "certificate_authority_rules.#", "2"),
					resource.TestCheckResourceAttr(subCAResourceName, "certificate_authority_rules.0.rule_type", "CERTIFICATE_AUTHORITY_ISSUANCE_EXPIRY_RULE"),
					resource.TestCheckResourceAttr(subCAResourceName, "certificate_authority_rules.0.certificate_authority_max_validity_duration", "P1000D"),
					resource.TestCheckResourceAttr(subCAResourceName, "certificate_authority_rules.0.leaf_certificate_max_validity_duration", "P500D"),
					resource.TestCheckResourceAttr(subCAResourceName, "certificate_authority_rules.1.path_length_constraint", "9"), //rules value depends on input cert pem
					resource.TestCheckResourceAttr(subCAResourceName, "certificate_authority_rules.1.name_constraint.#", "1"),
					resource.TestCheckResourceAttr(subCAResourceName, "certificate_authority_rules.1.rule_type", "CERTIFICATE_AUTHORITY_ISSUANCE_RULE"),
					resource.TestCheckResourceAttr(subCAResourceName, "state", "PENDING_ACTIVATION"),
					resource.TestCheckResourceAttr(subCAV1DataSourceName, "csr_pem", subCsrPem), //if this doesn't match then csr changed hence certificate needs to be resigned before updating in the next step

					func(s *terraform.State) (err error) {
						resId, err = acctest.FromInstanceState(s, subCAResourceName, "id")
						return err
					},
				),
				ExpectNonEmptyPlan: false,
			},
			//Update the signed_cert for version 1 of the CA to make CA  active; no change to ca rules in both root and sub
			{

				Config: config + compartmentIdVariableStr + rootCertificatePemVariableStr + signedSubCertPemVariableStr + CertificatesManagementCertificateAuthorityResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_certificates_management_certificate_authority", "test_certificate_authority_root_managed_externally", acctest.Optional, acctest.Create,
						acctest.RepresentationCopyWithNewProperties(certificateAuthorityRepresentationRootManagedExternally, map[string]interface{}{
							"certificate_authority_rules": acctest.RepresentationGroup{RepType: acctest.Required, Group: byocaRootCertificateAuthorityCertificateAuthorityRules_IssuanceRulePathLengthRepresentation},
						})) + acctest.GenerateDataSourceFromRepresentationMap("oci_certificates_management_certificate_authority_version", "test_certificate_authority_sub_byoca_v1", acctest.Optional, acctest.Create, certificateAuthorityVersionDataSourceRepresentationSubManagedInternally) +
					acctest.GenerateResourceFromRepresentationMap("oci_certificates_management_certificate_authority", "test_certificate_authority_sub_byoca", acctest.Optional, acctest.Update, certificateAuthorityRepresentationSubManagedInternally),

				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(subCAResourceName, "name", subCAManagedInternallyIssuedExternallyName),
					resource.TestCheckResourceAttr(subCAResourceName, "certificate_authority_rules.#", "2"),
					resource.TestCheckResourceAttr(subCAResourceName, "certificate_authority_rules.0.rule_type", "CERTIFICATE_AUTHORITY_ISSUANCE_EXPIRY_RULE"),
					resource.TestCheckResourceAttr(subCAResourceName, "certificate_authority_rules.0.certificate_authority_max_validity_duration", "P1000D"),
					resource.TestCheckResourceAttr(subCAResourceName, "certificate_authority_rules.0.leaf_certificate_max_validity_duration", "P500D"),
					resource.TestCheckResourceAttr(subCAResourceName, "certificate_authority_rules.1.path_length_constraint", "9"), //rules value depends on input cert pem
					resource.TestCheckResourceAttr(subCAResourceName, "certificate_authority_rules.1.name_constraint.#", "1"),
					resource.TestCheckResourceAttr(subCAResourceName, "certificate_authority_rules.1.rule_type", "CERTIFICATE_AUTHORITY_ISSUANCE_RULE"),
					resource.TestCheckResourceAttr(subCAV1DataSourceName, "csr_pem", subCsrPem),
					resource.TestCheckResourceAttr(subCAResourceName, "current_version.0.stages.#", "2"),
					resource.TestCheckResourceAttr(subCAResourceName, "state", "ACTIVE"),
					func(s *terraform.State) (err error) {
						resId2, err = acctest.FromInstanceState(s, subCAResourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("Resource recreated when it was supposed to be updated.")
						}
						return err
					},
				),
				ExpectNonEmptyPlan: false,
			},
			//Renew new version 2 of both the rootCA and subCA, change the ca_rules in subCA to delete expiry rule and preserve same config for root CA
			{

				Config: config + compartmentIdVariableStr + rootCertificatePemVariableStr + signedSubCertPemVariableStr + CertificatesManagementCertificateAuthorityResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_certificates_management_certificate_authority", "test_certificate_authority_root_managed_externally", acctest.Optional, acctest.Create,
						acctest.RepresentationCopyWithNewProperties(certificateAuthorityRepresentationRootManagedExternally, map[string]interface{}{
							"certificate_authority_config": acctest.RepresentationGroup{RepType: acctest.Required, Group: certificateAuthorityConfigRepresentationRootManagedExternally_V2},
							"certificate_authority_rules":  acctest.RepresentationGroup{RepType: acctest.Required, Group: byocaRootCertificateAuthorityCertificateAuthorityRules_IssuanceRulePathLengthRepresentation},
						})) + acctest.GenerateDataSourceFromRepresentationMap("oci_certificates_management_certificate_authority_version", "test_certificate_authority_sub_byoca_v1", acctest.Optional, acctest.Create, certificateAuthorityVersionDataSourceRepresentationSubManagedInternally) +
					acctest.GenerateResourceFromRepresentationMap("oci_certificates_management_certificate_authority", "test_certificate_authority_sub_byoca", acctest.Optional, acctest.Update, acctest.RepresentationCopyWithNewProperties(certificateAuthorityRepresentationSubManagedInternally, map[string]interface{}{
						"certificate_authority_config": acctest.RepresentationGroup{RepType: acctest.Required, Group: certificateAuthorityCertificateAuthorityGenerateCsrConfigRepresentationSubManagedInternally},
						"certificate_authority_rules":  acctest.RepresentationGroup{RepType: acctest.Required, Group: byocaSubCertificateAuthorityCertificateAuthorityRules_IssuanceRulePathLengthRepresentation},
					})) + acctest.GenerateDataSourceFromRepresentationMap("oci_certificates_management_certificate_authority_version", "test_certificate_authority_sub_byoca_v2", acctest.Optional, acctest.Create, certificateAuthorityVersionDataSourceRepresentationSubManagedInternallyV2) +
					acctest.GenerateDataSourceFromRepresentationMap("oci_certificates_management_certificate_authority_version", "test_certificate_authority_root_byoca_v2", acctest.Optional, acctest.Create, certificateAuthorityVersionDataSourceRepresentationRootManagedExternallyV2),

				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(subCAResourceName, "name", subCAManagedInternallyIssuedExternallyName),
					resource.TestCheckResourceAttr(subCAResourceName, "certificate_authority_rules.#", "1"),
					resource.TestCheckResourceAttr(subCAResourceName, "certificate_authority_rules.0.path_length_constraint", "9"), //rules value depends on input cert pem
					resource.TestCheckResourceAttr(subCAResourceName, "certificate_authority_rules.0.name_constraint.#", "1"),
					resource.TestCheckResourceAttr(subCAResourceName, "certificate_authority_rules.0.rule_type", "CERTIFICATE_AUTHORITY_ISSUANCE_RULE"),
					resource.TestCheckResourceAttr(subCAResourceName, "current_version.0.stages.#", "1"), //current version is not latest anymore
					resource.TestCheckResourceAttr(subCAResourceName, "state", "ACTIVE"),
					resource.TestCheckResourceAttr(subCAV1DataSourceName, "csr_pem", subCsrPem),
					resource.TestCheckResourceAttr(subCAV1DataSourceName, "stages.#", "1"),
					resource.TestCheckResourceAttr(subCAV2DataSourceName, "csr_pem", subCsrPem),
					resource.TestCheckResourceAttr(subCAV2DataSourceName, "stages.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "current_version.0.stages.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "state", "ACTIVE"),
					resource.TestCheckResourceAttr(rootCAV2DataSourceName, "stages.#", "2"),

					func(s *terraform.State) (err error) {
						resId2, err = acctest.FromInstanceState(s, subCAResourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("Resource recreated when it was supposed to be updated.")
						}
						return err
					},
				),
				ExpectNonEmptyPlan: false,
			},
		},
	})
}

func TestCertificatesManagementCertificateAuthorityResource_ConstraintsDiffSuppressTest(t *testing.T) {
	httpreplay.SetScenario("TestCertificatesManagementCertificateAuthorityResource_ConstraintsDiffSuppressTest")
	defer httpreplay.SaveScenario()

	provider := acctest.TestAccProvider
	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_certificates_management_certificate_authority.test_certificate_authority_path"
	subCAPathLengthResourceName := "oci_certificates_management_certificate_authority.test_sub_certificate_authority"

	var resId, resId2 string
	// Save TF content to create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+CertificatesManagementCertificateAuthorityResourceDependencies, "certificatesmanagement", "certificateAuthority", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { acctest.PreCheck(t) },
		Providers: map[string]*schema.Provider{
			"oci": provider,
		},
		Steps: []resource.TestStep{

			// verify create with issuance rules
			{
				Config: config + compartmentIdVariableStr + CertificatesManagementCertificateAuthorityResourceDependencies + CertificatesManagementCertificateAuthorityPathLengthTestResource,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "certificate_authority_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "certificate_authority_config.0.config_type", "ROOT_CA_GENERATED_INTERNALLY"),
					resource.TestCheckResourceAttr(resourceName, "certificate_authority_config.0.subject.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "certificate_authority_config.0.subject.0.common_name", "www.example.com"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(resourceName, "kms_key_id"),
					resource.TestCheckResourceAttr(resourceName, "name", rootCaName),
					resource.TestCheckResourceAttr(resourceName, "certificate_authority_rules.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "certificate_authority_rules.0.path_length_constraint", "5"),
					resource.TestCheckResourceAttr(resourceName, "certificate_authority_rules.0.rule_type", "CERTIFICATE_AUTHORITY_ISSUANCE_RULE"),
					resource.TestCheckResourceAttr(resourceName, "certificate_authority_rules.0.name_constraint.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "certificate_authority_rules.0.name_constraint.0.permitted_subtree.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "certificate_authority_rules.0.name_constraint.0.excluded_subtree.#", "1"),
					func(s *terraform.State) (err error) {
						resId, err = acctest.FromInstanceState(s, resourceName, "id")
						fmt.Printf("variable \"resId\" { default = \"%s\" }\n", resId)
						return err
					},
				),
				ExpectNonEmptyPlan: false,
			},
			//step 2: setup to verify when sub ca also has issuance rule and issuer has one,even when the sub CA's config doesn't match state diff is suppressed
			{

				Config: config + compartmentIdVariableStr + CertificatesManagementCertificateAuthorityResourceDependencies + CertificatesManagementCertificateAuthorityPathLengthTestResource + acctest.GenerateResourceFromRepresentationMap("oci_certificates_management_certificate_authority", "test_sub_certificate_authority", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(certificateAuthorityRepresentationSub, map[string]interface{}{
						"certificate_authority_config": acctest.RepresentationGroup{
							RepType: acctest.Required, Group: map[string]interface{}{
								"config_type":                     acctest.Representation{RepType: acctest.Required, Create: `SUBORDINATE_CA_ISSUED_BY_INTERNAL_CA`},
								"subject":                         acctest.RepresentationGroup{RepType: acctest.Required, Group: certificatesManagementCertificateSubjectRepresentation},
								"issuer_certificate_authority_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_certificates_management_certificate_authority.test_certificate_authority_path.id}`},
								"signing_algorithm":               acctest.Representation{RepType: acctest.Optional, Create: `SHA256_WITH_RSA`},
								"validity":                        acctest.RepresentationGroup{RepType: acctest.Optional, Group: certificateAuthorityCertificateAuthorityConfigValidityRepresentation},
								"version_name":                    acctest.Representation{RepType: acctest.Optional, Create: `versionNamePathLength`},
							},
						},
						"certificate_authority_rules": acctest.RepresentationGroup{RepType: acctest.Optional, Group: certificateAuthorityCertificateAuthorityRules_IssuanceRuleRepresentationDiffSuppressCase},
					})),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(subCAPathLengthResourceName, "certificate_authority_config.#", "1"),
					resource.TestCheckResourceAttr(subCAPathLengthResourceName, "certificate_authority_config.0.config_type", "SUBORDINATE_CA_ISSUED_BY_INTERNAL_CA"),
					resource.TestCheckResourceAttrSet(subCAPathLengthResourceName, "certificate_authority_config.0.issuer_certificate_authority_id"),
					resource.TestCheckResourceAttr(subCAPathLengthResourceName, "certificate_authority_config.0.signing_algorithm", "SHA256_WITH_RSA"),
					resource.TestCheckResourceAttr(subCAPathLengthResourceName, "certificate_authority_config.0.subject.#", "1"),
					resource.TestCheckResourceAttr(subCAPathLengthResourceName, "certificate_authority_config.0.subject.0.common_name", "www.example.com"),
					resource.TestCheckResourceAttr(subCAPathLengthResourceName, "certificate_authority_rules.#", "1"),
					resource.TestCheckResourceAttr(subCAPathLengthResourceName, "certificate_authority_rules.0.rule_type", "CERTIFICATE_AUTHORITY_ISSUANCE_RULE"),
					resource.TestCheckResourceAttr(subCAPathLengthResourceName, "certificate_authority_rules.0.path_length_constraint", "4"),
					resource.TestCheckResourceAttr(subCAPathLengthResourceName, "certificate_authority_rules.0.name_constraint.0.permitted_subtree.#", "1"),
					resource.TestCheckResourceAttr(subCAPathLengthResourceName, "certificate_authority_rules.0.name_constraint.0.excluded_subtree.#", "2"),
					resource.TestCheckResourceAttr(subCAPathLengthResourceName, "certificate_authority_rules.0.name_constraint.#", "1"),
					func(s *terraform.State) (err error) {
						resId, err = acctest.FromInstanceState(s, subCAPathLengthResourceName, "id")
						return err
					},
				),
				ExpectNonEmptyPlan: false,
			},
			//step 3: setup to verify when sub ca had a path length in config file but now that is deleted and replaced with issuance_Expiry rule
			//the update goes though for expiry rule with provider injected issuance rule
			{

				Config: config + compartmentIdVariableStr + CertificatesManagementCertificateAuthorityResourceDependencies + CertificatesManagementCertificateAuthorityPathLengthTestResource + acctest.GenerateResourceFromRepresentationMap("oci_certificates_management_certificate_authority", "test_sub_certificate_authority", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(certificateAuthorityRepresentationSub, map[string]interface{}{
						"certificate_authority_config": acctest.RepresentationGroup{
							RepType: acctest.Required, Group: map[string]interface{}{
								"config_type":                     acctest.Representation{RepType: acctest.Required, Create: `SUBORDINATE_CA_ISSUED_BY_INTERNAL_CA`},
								"subject":                         acctest.RepresentationGroup{RepType: acctest.Required, Group: certificatesManagementCertificateSubjectRepresentation},
								"issuer_certificate_authority_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_certificates_management_certificate_authority.test_certificate_authority_path.id}`},
								"signing_algorithm":               acctest.Representation{RepType: acctest.Optional, Create: `SHA256_WITH_RSA`},
								"validity":                        acctest.RepresentationGroup{RepType: acctest.Optional, Group: certificateAuthorityCertificateAuthorityConfigValidityRepresentation},
								"version_name":                    acctest.Representation{RepType: acctest.Optional, Create: `versionNamePathLength`},
							},
						},
						"certificate_authority_rules": acctest.RepresentationGroup{RepType: acctest.Optional, Group: map[string]interface{}{
							"rule_type": acctest.Representation{RepType: acctest.Required, Create: `CERTIFICATE_AUTHORITY_ISSUANCE_EXPIRY_RULE`, Update: `CERTIFICATE_AUTHORITY_ISSUANCE_EXPIRY_RULE`},
							"certificate_authority_max_validity_duration": acctest.Representation{RepType: acctest.Optional, Create: `P101D`, Update: `P1500D`},
							"leaf_certificate_max_validity_duration":      acctest.Representation{RepType: acctest.Optional, Create: `P51D`, Update: `P400D`},
						}},
					})),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(subCAPathLengthResourceName, "certificate_authority_config.#", "1"),
					resource.TestCheckResourceAttr(subCAPathLengthResourceName, "certificate_authority_config.0.config_type", "SUBORDINATE_CA_ISSUED_BY_INTERNAL_CA"),
					resource.TestCheckResourceAttrSet(subCAPathLengthResourceName, "certificate_authority_config.0.issuer_certificate_authority_id"),
					resource.TestCheckResourceAttr(subCAPathLengthResourceName, "certificate_authority_config.0.signing_algorithm", "SHA256_WITH_RSA"),
					resource.TestCheckResourceAttr(subCAPathLengthResourceName, "certificate_authority_config.0.subject.#", "1"),
					resource.TestCheckResourceAttr(subCAPathLengthResourceName, "certificate_authority_config.0.subject.0.common_name", "www.example.com"),
					resource.TestCheckResourceAttr(subCAPathLengthResourceName, "certificate_authority_rules.#", "2"),
					resource.TestCheckResourceAttr(subCAPathLengthResourceName, "certificate_authority_rules.1.rule_type", "CERTIFICATE_AUTHORITY_ISSUANCE_RULE"),
					resource.TestCheckResourceAttr(subCAPathLengthResourceName, "certificate_authority_rules.1.path_length_constraint", "4"),
					resource.TestCheckResourceAttr(subCAPathLengthResourceName, "certificate_authority_rules.1.name_constraint.#", "1"),
					resource.TestCheckResourceAttr(subCAPathLengthResourceName, "certificate_authority_rules.1.name_constraint.0.permitted_subtree.#", "1"),
					resource.TestCheckResourceAttr(subCAPathLengthResourceName, "certificate_authority_rules.1.name_constraint.0.excluded_subtree.#", "2"),
					resource.TestCheckResourceAttr(subCAPathLengthResourceName, "certificate_authority_rules.0.rule_type", "CERTIFICATE_AUTHORITY_ISSUANCE_EXPIRY_RULE"),
					resource.TestCheckResourceAttr(subCAPathLengthResourceName, "certificate_authority_rules.0.certificate_authority_max_validity_duration", "P101D"),
					resource.TestCheckResourceAttr(subCAPathLengthResourceName, "certificate_authority_rules.0.leaf_certificate_max_validity_duration", "P51D"),
					func(s *terraform.State) (err error) {
						resId2, err = acctest.FromInstanceState(s, subCAPathLengthResourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("Resource recreated when it was supposed to be updated.")
						}
						return err
					},
				),
				ExpectNonEmptyPlan: false,
			},
		},
	})
}
