package certificates_management

import (
	oci_certificates_management "github.com/oracle/oci-go-sdk/v65/certificatesmanagement"
	"time"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	exportCertificatesManagementCertificateAuthorityHints.ProcessDiscoveredResourcesFn = processCertificateAuthorities
	exportCertificatesManagementCertificateHints.ProcessDiscoveredResourcesFn = processCertificates
	tf_export.RegisterCompartmentGraphs("certificates_management", certificatesManagementResourceGraph)
}

// Custom overrides for generating composite IDs within the resource discovery framework

func processCertificateAuthorities(ctx *tf_export.ResourceDiscoveryContext, resources []*tf_export.OCIResource) ([]*tf_export.OCIResource, error) {
	for _, resource := range resources {
		certificateAuthorityConfigMap := map[string]interface{}{}
		if configType, ok := resource.SourceAttributes["config_type"].(string); ok {
			certificateAuthorityConfigMap["config_type"] = configType
		}

		if subjects, ok := resource.SourceAttributes["subject"].([]interface{}); ok {
			if subject, ok := subjects[0].(map[string]interface{}); ok {
				certificateAuthorityConfigMap["subject"] = []interface{}{subject}
			}
		}

		if issuerCertificateAuthorityId, ok := resource.SourceAttributes["issuer_certificate_authority_id"].(string); ok {
			certificateAuthorityConfigMap["issuer_certificate_authority_id"] = issuerCertificateAuthorityId
		}

		if signingAlgorithm, ok := resource.SourceAttributes["signing_algorithm"].(string); ok {
			certificateAuthorityConfigMap["signing_algorithm"] = signingAlgorithm
		}

		if currentVersions, ok := resource.SourceAttributes["current_version"].([]interface{}); ok {
			if currentVersion, ok := currentVersions[0].(map[string]interface{}); ok {
				if validity, ok := currentVersion["validity"].([]interface{}); ok {
					validityMap := map[string]interface{}{}
					if timeOfValidityNotAfter, ok := validity[0].(map[string]interface{})["time_of_validity_not_after"]; ok {
						// Check if time is already in RFC3339Nano format
						tmp, err := time.Parse(time.RFC3339Nano, timeOfValidityNotAfter.(string))
						if err != nil {
							// parse time using format in time.String()
							tmp, err = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", timeOfValidityNotAfter.(string))
							if err != nil {
								return resources, err
							}
							// Format to RFC3339Nano
							validityMap["time_of_validity_not_after"] = tmp.Format(time.RFC3339Nano)
						}
					}
					if timeOfValidityNotBefore, ok := validity[0].(map[string]interface{})["time_of_validity_not_before"]; ok {
						// Check if time is already in RFC3339Nano format
						tmp, err := time.Parse(time.RFC3339Nano, timeOfValidityNotBefore.(string))
						if err != nil {
							// parse time using format in time.String()
							tmp, err = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", timeOfValidityNotBefore.(string))
							if err != nil {
								return resources, err
							}
							// Format to RFC3339Nano
							validityMap["time_of_validity_not_before"] = tmp.Format(time.RFC3339Nano)
						}
					}

					certificateAuthorityConfigMap["validity"] = []interface{}{validityMap}
				}

				if versionName, ok := currentVersion["version_name"].(string); ok {
					certificateAuthorityConfigMap["version_name"] = versionName
				}
			}
		}

		resource.SourceAttributes["certificate_authority_config"] = []interface{}{certificateAuthorityConfigMap}
	}

	return resources, nil
}

func processCertificates(ctx *tf_export.ResourceDiscoveryContext, resources []*tf_export.OCIResource) ([]*tf_export.OCIResource, error) {
	for _, resource := range resources {
		certificateConfigMap := map[string]interface{}{}
		if configType, ok := resource.SourceAttributes["config_type"].(string); ok {
			certificateConfigMap["config_type"] = configType
		}

		if profileType, ok := resource.SourceAttributes["certificate_profile_type"].(string); ok {
			certificateConfigMap["certificate_profile_type"] = profileType
		}

		if csrPem, ok := resource.SourceAttributes["csr_pem"].(string); ok {
			certificateConfigMap["csr_pem"] = csrPem
		}

		if issuerCertificateAuthorityId, ok := resource.SourceAttributes["issuer_certificate_authority_id"].(string); ok {
			certificateConfigMap["issuer_certificate_authority_id"] = issuerCertificateAuthorityId
		}

		if keyAlgorithm, ok := resource.SourceAttributes["key_algorithm"].(string); ok {
			certificateConfigMap["key_algorithm"] = keyAlgorithm
		}

		if signatureAlgorithm, ok := resource.SourceAttributes["signature_algorithm"].(string); ok {
			certificateConfigMap["signature_algorithm"] = signatureAlgorithm
		}

		if subjects, ok := resource.SourceAttributes["subject"].([]interface{}); ok {
			if subject, ok := subjects[0].(map[string]interface{}); ok {
				certificateConfigMap["subject"] = []interface{}{subject}
			}
		}

		if currentVersions, ok := resource.SourceAttributes["current_version"].([]interface{}); ok {
			if currentVersion, ok := currentVersions[0].(map[string]interface{}); ok {
				if validity, ok := currentVersion["validity"].([]interface{}); ok {
					validityMap := map[string]interface{}{}
					if timeOfValidityNotAfter, ok := validity[0].(map[string]interface{})["time_of_validity_not_after"]; ok {
						// Check if time is already in RFC3339Nano format
						tmp, err := time.Parse(time.RFC3339Nano, timeOfValidityNotAfter.(string))
						if err != nil {
							// parse time using format in time.String()
							tmp, err = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", timeOfValidityNotAfter.(string))
							if err != nil {
								return resources, err
							}
							// Format to RFC3339Nano
							validityMap["time_of_validity_not_after"] = tmp.Format(time.RFC3339Nano)
						}
					}
					if timeOfValidityNotBefore, ok := validity[0].(map[string]interface{})["time_of_validity_not_before"]; ok {
						// Check if time is already in RFC3339Nano format
						tmp, err := time.Parse(time.RFC3339Nano, timeOfValidityNotBefore.(string))
						if err != nil {
							// parse time using format in time.String()
							tmp, err = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", timeOfValidityNotBefore.(string))
							if err != nil {
								return resources, err
							}
							// Format to RFC3339Nano
							validityMap["time_of_validity_not_before"] = tmp.Format(time.RFC3339Nano)
						}
					}

					certificateConfigMap["validity"] = []interface{}{validityMap}
				}

				if versionName, ok := currentVersion["version_name"].(string); ok {
					certificateConfigMap["version_name"] = versionName
				}

				if subjectAlternativeNames, ok := currentVersion["subject_alternative_names"].([]interface{}); ok {
					tmp := []interface{}{}
					for _, item := range subjectAlternativeNames {
						tmp = append(tmp, item)
					}
					certificateConfigMap["subject_alternative_names"] = tmp
				}
			}
		}

		resource.SourceAttributes["certificate_config"] = []interface{}{certificateConfigMap}
	}

	return resources, nil
}

// Hints for discovering and exporting this resource to configuration and state files
var exportCertificatesManagementCaBundleHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_certificates_management_ca_bundle",
	DatasourceClass:        "oci_certificates_management_ca_bundles",
	DatasourceItemsAttr:    "ca_bundle_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "ca_bundle",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_certificates_management.CaBundleLifecycleStateActive),
	},
}

var exportCertificatesManagementCertificateAuthorityHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_certificates_management_certificate_authority",
	DatasourceClass:        "oci_certificates_management_certificate_authorities",
	DatasourceItemsAttr:    "certificate_authority_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "certificate_authority",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_certificates_management.CertificateAuthorityLifecycleStateActive),
	},
}

var exportCertificatesManagementCertificateHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_certificates_management_certificate",
	DatasourceClass:        "oci_certificates_management_certificates",
	DatasourceItemsAttr:    "certificate_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "certificate",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_certificates_management.CertificateLifecycleStateActive),
	},
}

var certificatesManagementResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportCertificatesManagementCaBundleHints},
		{TerraformResourceHints: exportCertificatesManagementCertificateAuthorityHints},
		{TerraformResourceHints: exportCertificatesManagementCertificateHints},
	},
}
