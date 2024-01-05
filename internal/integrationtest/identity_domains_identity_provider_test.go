// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_identity_domains "github.com/oracle/oci-go-sdk/v65/identitydomains"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	IdentityDomainsIdentityProviderRequiredOnlyResource = IdentityDomainsIdentityProviderResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_identity_provider", "test_identity_provider", acctest.Required, acctest.Create, IdentityDomainsIdentityProviderRepresentation)

	IdentityDomainsIdentityProviderResourceConfig = IdentityDomainsIdentityProviderResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_identity_provider", "test_identity_provider", acctest.Optional, acctest.Update, IdentityDomainsIdentityProviderRepresentation)

	IdentityDomainsIdentityDomainsIdentityProviderSingularDataSourceRepresentation = map[string]interface{}{
		"idcs_endpoint":        acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"identity_provider_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_identity_domains_identity_provider.test_identity_provider.id}`},
		"attribute_sets":       acctest.Representation{RepType: acctest.Optional, Create: []string{`all`}},
	}

	IdentityDomainsIdentityDomainsIdentityProviderDataSourceRepresentation = map[string]interface{}{
		"idcs_endpoint":            acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"identity_provider_count":  acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"identity_provider_filter": acctest.Representation{RepType: acctest.Optional, Create: `partnerName eq \"partnerName2\"`},
		"attribute_sets":           acctest.Representation{RepType: acctest.Optional, Create: []string{`all`}},
		"start_index":              acctest.Representation{RepType: acctest.Optional, Create: `1`},
	}

	IdentityDomainsIdentityProviderRepresentation = map[string]interface{}{
		"enabled":                                acctest.Representation{RepType: acctest.Required, Create: `true`, Update: `false`},
		"idcs_endpoint":                          acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"partner_name":                           acctest.Representation{RepType: acctest.Required, Create: `partnerName`, Update: `partnerName2`},
		"schemas":                                acctest.Representation{RepType: acctest.Required, Create: []string{`urn:ietf:params:scim:schemas:oracle:idcs:IdentityProvider`}},
		"assertion_attribute":                    acctest.Representation{RepType: acctest.Optional, Create: `assertionAttribute`, Update: `assertionAttribute2`},
		"attribute_sets":                         acctest.Representation{RepType: acctest.Optional, Create: []string{`all`}},
		"authn_request_binding":                  acctest.Representation{RepType: acctest.Optional, Create: `Redirect`, Update: `Post`},
		"description":                            acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"encryption_certificate":                 acctest.Representation{RepType: acctest.Optional, Create: `MIIDZjCCAk6gAwIBAgIGAXoTivr8MA0GCSqGSIb3DQEBCwUAMFkxEzARBgoJkiaJk/IsZAEZFgNjb20xFjAUBgoJkiaJk/IsZAEZFgZvcmFjbGUxFTATBgoJkiaJk/IsZAEZFgVjbG91ZDETMBEGA1UEAxMKQ2xvdWQ5Q0EtMjAeFw0yMTA2MTYwNjM4MzJaFw0zMTA2MTYwNjM4MzJaMFsxEzARBgNVBAMTCnNzbERvbWFpbnMxDzANBgNVBAMTBkNsb3VkOTEzMDEGA1UEAwwqaWRjcy1hY2YxM2MzMDYyMTM0NTJiYmQwZjgzZmJkYjM3OWYxZl9rZXlzMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEArC8Sg+4SF3ZKdRjveLNheqFlcXZIDV6N6FjMJEYVLyQLrRg8gAXIRqhpS/Ym9Bj8emyV+P/ZlHxmwPkI8DbJX+gsTg9xXG+DlfDSnNTdEOr90hrmg5D8NN3L0SpLughTnywPRbvk6Iv/21u8vALqQHfJYezBDqS/USpm/EP5+soaOlfGWAUaVnxqLtVC6g8kN3dj4hry3ZcfOZwAriEC2M1TTA8vOd7lkGRwEPjBL/QOXWKJU7LPYArKpuL75+gJteQ/yvDHIjbo7ZGaKGhhiCYcQ0gS1I6Fe3V8Mw3N+m7/TpXwGxyEo7GrR7M/j7jhJto88eINDhBjiQkungCY3QIDAQABozIwMDAPBgNVHQ8BAf8EBQMDB/gAMB0GA1UdDgQWBBQ+PjQeG7CnRrqF9HXop07/756KrzANBgkqhkiG9w0BAQsFAAOCAQEAtQqQhc3tsdTHAlQRDkhnPS3N6x7xBpbjef+HOlta86DA9LSoYHBGnnZiwQHTgRbfUdUNllp7hITKMW942wTTwzY/owPukYoo4UAurOZcNkqFCzYJGqxS5bwQRSQO8iSRGwwooet730tzevcilLWH7gF752unoPBoxgCPNxdxWxxSmHpiaQTUAy00sfba0KQ5SLCkXJHEOXpnQsXhe0qm7f/0Mo51cU9pIChj0Cfc/eEmwbcdPnPOCP4N8je+kqW79b3b90kcyUqMDTJw+4HKK62juXm4SXaePJYzx1Yb1VCSZpwOBjHxudp4WrkRPTXFRX248UtSUI5zVw1QsCmOnw==`},
		"external_id":                            acctest.Representation{RepType: acctest.Optional, Create: `externalId`},
		"icon_url":                               acctest.Representation{RepType: acctest.Optional, Create: `https://something.com/iconUrl.png`, Update: `https://something.com/iconUrl2.png`},
		"idp_sso_url":                            acctest.Representation{RepType: acctest.Required, Create: `https://idpSsoUrl.com`, Update: `https://idpSsoUrl2.com`},
		"include_signing_cert_in_signature":      acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"jit_user_prov_assigned_groups":          acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityDomainsIdentityProviderJitUserProvAssignedGroupsRepresentation},
		"jit_user_prov_attribute_update_enabled": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"jit_user_prov_create_user_enabled":      acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"jit_user_prov_enabled":                  acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"jit_user_prov_group_assertion_attribute_enabled": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"jit_user_prov_group_assignment_method":           acctest.Representation{RepType: acctest.Optional, Create: `Overwrite`, Update: `Merge`},
		"jit_user_prov_group_mapping_mode":                acctest.Representation{RepType: acctest.Optional, Create: `implicit`, Update: `explicit`},
		"jit_user_prov_group_mappings":                    acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityDomainsIdentityProviderJitUserProvGroupMappingsRepresentation},
		"jit_user_prov_group_saml_attribute_name":         acctest.Representation{RepType: acctest.Optional, Create: `jitUserProvGroupSAMLAttributeName`, Update: `jitUserProvGroupSAMLAttributeName2`},
		"jit_user_prov_group_static_list_enabled":         acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"jit_user_prov_ignore_error_on_absent_groups":     acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"logout_binding":                                  acctest.Representation{RepType: acctest.Optional, Create: `Redirect`, Update: `Post`},
		"logout_enabled":                                  acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"logout_request_url":                              acctest.Representation{RepType: acctest.Optional, Create: `https://logoutRequestUrl.com`, Update: `https://logoutRequestUrl2.com`},
		"logout_response_url":                             acctest.Representation{RepType: acctest.Optional, Create: `https://logoutResponseUrl.com`, Update: `https://logoutResponseUrl2.com`},
		"metadata":                                        acctest.Representation{RepType: acctest.Optional, Create: `<md:EntityDescriptor xmlns:md=\"urn:oasis:names:tc:SAML:2.0:metadata\" xmlns:dsig=\"http://www.w3.org/2000/09/xmldsig#\" xmlns:enc=\"http://www.w3.org/2001/04/xmlenc#\" xmlns:mdattr=\"urn:oasis:names:tc:SAML:metadata:attribute\" xmlns:query=\"urn:oasis:names:tc:SAML:metadata:ext:query\" xmlns:saml=\"urn:oasis:names:tc:SAML:2.0:assertion\" xmlns:x500=\"urn:oasis:names:tc:SAML:2.0:profiles:attribute:X500\" xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" ID=\"id-zzU36agM7bKRB32xe6Ronm131S0-\" cacheDuration=\"P3633DT0H0M0S\" entityID=\"https://idcs-acf13c306213452bbd0f83fbdb379f1f.unstable-alpha.identity.oci.oracleiaas.com:443/fed\" validUntil=\"2031-06-16T06:38:32Z\"><dsig:Signature><dsig:SignedInfo><dsig:CanonicalizationMethod Algorithm=\"http://www.w3.org/2001/10/xml-exc-c14n#\"/><dsig:SignatureMethod Algorithm=\"http://www.w3.org/2001/04/xmldsig-more#rsa-sha256\"/><dsig:Reference URI=\"#id-zzU36agM7bKRB32xe6Ronm131S0-\"><dsig:Transforms><dsig:Transform Algorithm=\"http://www.w3.org/2000/09/xmldsig#enveloped-signature\"/><dsig:Transform Algorithm=\"http://www.w3.org/2001/10/xml-exc-c14n#\"/></dsig:Transforms><dsig:DigestMethod Algorithm=\"http://www.w3.org/2001/04/xmlenc#sha256\"/><dsig:DigestValue>NZnYsjLx3UbuL43iu3jo0mJUg/Rv9DTPNB5IQPRaD6g=</dsig:DigestValue></dsig:Reference></dsig:SignedInfo><dsig:SignatureValue>KRIgTD7//x/uT73veS0iGcWWw8uprjd+MtREu3vlbFTk0BNgkeSOYItx2LDQhnHP\nO0zsTmtOHlVIsDXQL3KysHwzYndIuMJtETqEC6NpMw3ZF108IK0eT+o/2xC9u13/\nGq10z/MagGvco1mM/RIzX5e2omGyZcKARiDoeNPwg2znmV0WcifntVqn4Y0rnWM7\no0M5HFHZQEgICdTJbC5d6DwLgfnI4ck505fHNRYLsRqj9IGLukKx9kocSG1xzCye\nHlffU4CDyEA7dptEUH59dZmY0Xy35/aepNc7W6IovWsJ2Otr+qDUp207ZCKuISF0\nMEX5hX5VJzVlHDwxkEcYCA==</dsig:SignatureValue><dsig:KeyInfo><dsig:X509Data><dsig:X509Certificate>MIIDZjCCAk6gAwIBAgIGAXoTivr8MA0GCSqGSIb3DQEBCwUAMFkxEzARBgoJkiaJ\nk/IsZAEZFgNjb20xFjAUBgoJkiaJk/IsZAEZFgZvcmFjbGUxFTATBgoJkiaJk/Is\nZAEZFgVjbG91ZDETMBEGA1UEAxMKQ2xvdWQ5Q0EtMjAeFw0yMTA2MTYwNjM4MzJa\nFw0zMTA2MTYwNjM4MzJaMFsxEzARBgNVBAMTCnNzbERvbWFpbnMxDzANBgNVBAMT\nBkNsb3VkOTEzMDEGA1UEAwwqaWRjcy1hY2YxM2MzMDYyMTM0NTJiYmQwZjgzZmJk\nYjM3OWYxZl9rZXlzMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEArC8S\ng+4SF3ZKdRjveLNheqFlcXZIDV6N6FjMJEYVLyQLrRg8gAXIRqhpS/Ym9Bj8emyV\n+P/ZlHxmwPkI8DbJX+gsTg9xXG+DlfDSnNTdEOr90hrmg5D8NN3L0SpLughTnywP\nRbvk6Iv/21u8vALqQHfJYezBDqS/USpm/EP5+soaOlfGWAUaVnxqLtVC6g8kN3dj\n4hry3ZcfOZwAriEC2M1TTA8vOd7lkGRwEPjBL/QOXWKJU7LPYArKpuL75+gJteQ/\nyvDHIjbo7ZGaKGhhiCYcQ0gS1I6Fe3V8Mw3N+m7/TpXwGxyEo7GrR7M/j7jhJto8\n8eINDhBjiQkungCY3QIDAQABozIwMDAPBgNVHQ8BAf8EBQMDB/gAMB0GA1UdDgQW\nBBQ+PjQeG7CnRrqF9HXop07/756KrzANBgkqhkiG9w0BAQsFAAOCAQEAtQqQhc3t\nsdTHAlQRDkhnPS3N6x7xBpbjef+HOlta86DA9LSoYHBGnnZiwQHTgRbfUdUNllp7\nhITKMW942wTTwzY/owPukYoo4UAurOZcNkqFCzYJGqxS5bwQRSQO8iSRGwwooet7\n30tzevcilLWH7gF752unoPBoxgCPNxdxWxxSmHpiaQTUAy00sfba0KQ5SLCkXJHE\nOXpnQsXhe0qm7f/0Mo51cU9pIChj0Cfc/eEmwbcdPnPOCP4N8je+kqW79b3b90kc\nyUqMDTJw+4HKK62juXm4SXaePJYzx1Yb1VCSZpwOBjHxudp4WrkRPTXFRX248UtS\nUI5zVw1QsCmOnw==</dsig:X509Certificate></dsig:X509Data></dsig:KeyInfo></dsig:Signature><md:IDPSSODescriptor WantAuthnRequestsSigned=\"false\" protocolSupportEnumeration=\"urn:oasis:names:tc:SAML:2.0:protocol\"><md:KeyDescriptor use=\"signing\"><dsig:KeyInfo><dsig:X509Data><dsig:X509Certificate>MIIDZjCCAk6gAwIBAgIGAXoTivr8MA0GCSqGSIb3DQEBCwUAMFkxEzARBgoJkiaJ\nk/IsZAEZFgNjb20xFjAUBgoJkiaJk/IsZAEZFgZvcmFjbGUxFTATBgoJkiaJk/Is\nZAEZFgVjbG91ZDETMBEGA1UEAxMKQ2xvdWQ5Q0EtMjAeFw0yMTA2MTYwNjM4MzJa\nFw0zMTA2MTYwNjM4MzJaMFsxEzARBgNVBAMTCnNzbERvbWFpbnMxDzANBgNVBAMT\nBkNsb3VkOTEzMDEGA1UEAwwqaWRjcy1hY2YxM2MzMDYyMTM0NTJiYmQwZjgzZmJk\nYjM3OWYxZl9rZXlzMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEArC8S\ng+4SF3ZKdRjveLNheqFlcXZIDV6N6FjMJEYVLyQLrRg8gAXIRqhpS/Ym9Bj8emyV\n+P/ZlHxmwPkI8DbJX+gsTg9xXG+DlfDSnNTdEOr90hrmg5D8NN3L0SpLughTnywP\nRbvk6Iv/21u8vALqQHfJYezBDqS/USpm/EP5+soaOlfGWAUaVnxqLtVC6g8kN3dj\n4hry3ZcfOZwAriEC2M1TTA8vOd7lkGRwEPjBL/QOXWKJU7LPYArKpuL75+gJteQ/\nyvDHIjbo7ZGaKGhhiCYcQ0gS1I6Fe3V8Mw3N+m7/TpXwGxyEo7GrR7M/j7jhJto8\n8eINDhBjiQkungCY3QIDAQABozIwMDAPBgNVHQ8BAf8EBQMDB/gAMB0GA1UdDgQW\nBBQ+PjQeG7CnRrqF9HXop07/756KrzANBgkqhkiG9w0BAQsFAAOCAQEAtQqQhc3t\nsdTHAlQRDkhnPS3N6x7xBpbjef+HOlta86DA9LSoYHBGnnZiwQHTgRbfUdUNllp7\nhITKMW942wTTwzY/owPukYoo4UAurOZcNkqFCzYJGqxS5bwQRSQO8iSRGwwooet7\n30tzevcilLWH7gF752unoPBoxgCPNxdxWxxSmHpiaQTUAy00sfba0KQ5SLCkXJHE\nOXpnQsXhe0qm7f/0Mo51cU9pIChj0Cfc/eEmwbcdPnPOCP4N8je+kqW79b3b90kc\nyUqMDTJw+4HKK62juXm4SXaePJYzx1Yb1VCSZpwOBjHxudp4WrkRPTXFRX248UtS\nUI5zVw1QsCmOnw==\n</dsig:X509Certificate><dsig:X509IssuerSerial><dsig:X509IssuerName>CN=Cloud9CA-2, DC=cloud, DC=oracle, DC=com</dsig:X509IssuerName><dsig:X509SerialNumber>1623825513212</dsig:X509SerialNumber></dsig:X509IssuerSerial><dsig:X509SubjectName>CN=idcs-acf13c306213452bbd0f83fbdb379f1f_keys, CN=Cloud9, CN=sslDomains</dsig:X509SubjectName></dsig:X509Data></dsig:KeyInfo></md:KeyDescriptor><md:KeyDescriptor use=\"encryption\"><dsig:KeyInfo><dsig:X509Data><dsig:X509Certificate>MIIDZjCCAk6gAwIBAgIGAXoTivr8MA0GCSqGSIb3DQEBCwUAMFkxEzARBgoJkiaJ\nk/IsZAEZFgNjb20xFjAUBgoJkiaJk/IsZAEZFgZvcmFjbGUxFTATBgoJkiaJk/Is\nZAEZFgVjbG91ZDETMBEGA1UEAxMKQ2xvdWQ5Q0EtMjAeFw0yMTA2MTYwNjM4MzJa\nFw0zMTA2MTYwNjM4MzJaMFsxEzARBgNVBAMTCnNzbERvbWFpbnMxDzANBgNVBAMT\nBkNsb3VkOTEzMDEGA1UEAwwqaWRjcy1hY2YxM2MzMDYyMTM0NTJiYmQwZjgzZmJk\nYjM3OWYxZl9rZXlzMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEArC8S\ng+4SF3ZKdRjveLNheqFlcXZIDV6N6FjMJEYVLyQLrRg8gAXIRqhpS/Ym9Bj8emyV\n+P/ZlHxmwPkI8DbJX+gsTg9xXG+DlfDSnNTdEOr90hrmg5D8NN3L0SpLughTnywP\nRbvk6Iv/21u8vALqQHfJYezBDqS/USpm/EP5+soaOlfGWAUaVnxqLtVC6g8kN3dj\n4hry3ZcfOZwAriEC2M1TTA8vOd7lkGRwEPjBL/QOXWKJU7LPYArKpuL75+gJteQ/\nyvDHIjbo7ZGaKGhhiCYcQ0gS1I6Fe3V8Mw3N+m7/TpXwGxyEo7GrR7M/j7jhJto8\n8eINDhBjiQkungCY3QIDAQABozIwMDAPBgNVHQ8BAf8EBQMDB/gAMB0GA1UdDgQW\nBBQ+PjQeG7CnRrqF9HXop07/756KrzANBgkqhkiG9w0BAQsFAAOCAQEAtQqQhc3t\nsdTHAlQRDkhnPS3N6x7xBpbjef+HOlta86DA9LSoYHBGnnZiwQHTgRbfUdUNllp7\nhITKMW942wTTwzY/owPukYoo4UAurOZcNkqFCzYJGqxS5bwQRSQO8iSRGwwooet7\n30tzevcilLWH7gF752unoPBoxgCPNxdxWxxSmHpiaQTUAy00sfba0KQ5SLCkXJHE\nOXpnQsXhe0qm7f/0Mo51cU9pIChj0Cfc/eEmwbcdPnPOCP4N8je+kqW79b3b90kc\nyUqMDTJw+4HKK62juXm4SXaePJYzx1Yb1VCSZpwOBjHxudp4WrkRPTXFRX248UtS\nUI5zVw1QsCmOnw==\n</dsig:X509Certificate><dsig:X509IssuerSerial><dsig:X509IssuerName>CN=Cloud9CA-2, DC=cloud, DC=oracle, DC=com</dsig:X509IssuerName><dsig:X509SerialNumber>1623825513212</dsig:X509SerialNumber></dsig:X509IssuerSerial><dsig:X509SubjectName>CN=idcs-acf13c306213452bbd0f83fbdb379f1f_keys, CN=Cloud9, CN=sslDomains</dsig:X509SubjectName></dsig:X509Data></dsig:KeyInfo><md:EncryptionMethod Algorithm=\"http://www.w3.org/2001/04/xmlenc#rsa-1_5\"/><md:EncryptionMethod Algorithm=\"http://www.w3.org/2001/04/xmlenc#aes128-cbc\"/><md:EncryptionMethod Algorithm=\"http://www.w3.org/2001/04/xmlenc#aes192-cbc\"/><md:EncryptionMethod Algorithm=\"http://www.w3.org/2001/04/xmlenc#aes256-cbc\"/><md:EncryptionMethod Algorithm=\"http://www.w3.org/2001/04/xmlenc#tripledes-cbc\"/></md:KeyDescriptor><md:SingleLogoutService Binding=\"urn:oasis:names:tc:SAML:2.0:bindings:HTTP-Redirect\" Location=\"https://idcs-acf13c306213452bbd0f83fbdb379f1f.unstable-alpha.identity.oci.oracleiaas.com/fed/v1/idp/slo\" ResponseLocation=\"https://idcs-acf13c306213452bbd0f83fbdb379f1f.unstable-alpha.identity.oci.oracleiaas.com/fed/v1/idp/slo\"/><md:SingleLogoutService Binding=\"urn:oasis:names:tc:SAML:2.0:bindings:HTTP-POST\" Location=\"https://idcs-acf13c306213452bbd0f83fbdb379f1f.unstable-alpha.identity.oci.oracleiaas.com/fed/v1/idp/slo\" ResponseLocation=\"https://idcs-acf13c306213452bbd0f83fbdb379f1f.unstable-alpha.identity.oci.oracleiaas.com/fed/v1/idp/slo\"/><md:SingleSignOnService Binding=\"urn:oasis:names:tc:SAML:2.0:bindings:HTTP-Redirect\" Location=\"https://idcs-acf13c306213452bbd0f83fbdb379f1f.unstable-alpha.identity.oci.oracleiaas.com/fed/v1/idp/sso\"/><md:SingleSignOnService Binding=\"urn:oasis:names:tc:SAML:2.0:bindings:HTTP-POST\" Location=\"https://idcs-acf13c306213452bbd0f83fbdb379f1f.unstable-alpha.identity.oci.oracleiaas.com/fed/v1/idp/sso\"/></md:IDPSSODescriptor><md:SPSSODescriptor AuthnRequestsSigned=\"true\" WantAssertionsSigned=\"true\" protocolSupportEnumeration=\"urn:oasis:names:tc:SAML:2.0:protocol\"><md:KeyDescriptor use=\"signing\"><dsig:KeyInfo><dsig:X509Data><dsig:X509Certificate>MIIDZjCCAk6gAwIBAgIGAXoTivr8MA0GCSqGSIb3DQEBCwUAMFkxEzARBgoJkiaJ\nk/IsZAEZFgNjb20xFjAUBgoJkiaJk/IsZAEZFgZvcmFjbGUxFTATBgoJkiaJk/Is\nZAEZFgVjbG91ZDETMBEGA1UEAxMKQ2xvdWQ5Q0EtMjAeFw0yMTA2MTYwNjM4MzJa\nFw0zMTA2MTYwNjM4MzJaMFsxEzARBgNVBAMTCnNzbERvbWFpbnMxDzANBgNVBAMT\nBkNsb3VkOTEzMDEGA1UEAwwqaWRjcy1hY2YxM2MzMDYyMTM0NTJiYmQwZjgzZmJk\nYjM3OWYxZl9rZXlzMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEArC8S\ng+4SF3ZKdRjveLNheqFlcXZIDV6N6FjMJEYVLyQLrRg8gAXIRqhpS/Ym9Bj8emyV\n+P/ZlHxmwPkI8DbJX+gsTg9xXG+DlfDSnNTdEOr90hrmg5D8NN3L0SpLughTnywP\nRbvk6Iv/21u8vALqQHfJYezBDqS/USpm/EP5+soaOlfGWAUaVnxqLtVC6g8kN3dj\n4hry3ZcfOZwAriEC2M1TTA8vOd7lkGRwEPjBL/QOXWKJU7LPYArKpuL75+gJteQ/\nyvDHIjbo7ZGaKGhhiCYcQ0gS1I6Fe3V8Mw3N+m7/TpXwGxyEo7GrR7M/j7jhJto8\n8eINDhBjiQkungCY3QIDAQABozIwMDAPBgNVHQ8BAf8EBQMDB/gAMB0GA1UdDgQW\nBBQ+PjQeG7CnRrqF9HXop07/756KrzANBgkqhkiG9w0BAQsFAAOCAQEAtQqQhc3t\nsdTHAlQRDkhnPS3N6x7xBpbjef+HOlta86DA9LSoYHBGnnZiwQHTgRbfUdUNllp7\nhITKMW942wTTwzY/owPukYoo4UAurOZcNkqFCzYJGqxS5bwQRSQO8iSRGwwooet7\n30tzevcilLWH7gF752unoPBoxgCPNxdxWxxSmHpiaQTUAy00sfba0KQ5SLCkXJHE\nOXpnQsXhe0qm7f/0Mo51cU9pIChj0Cfc/eEmwbcdPnPOCP4N8je+kqW79b3b90kc\nyUqMDTJw+4HKK62juXm4SXaePJYzx1Yb1VCSZpwOBjHxudp4WrkRPTXFRX248UtS\nUI5zVw1QsCmOnw==\n</dsig:X509Certificate><dsig:X509IssuerSerial><dsig:X509IssuerName>CN=Cloud9CA-2, DC=cloud, DC=oracle, DC=com</dsig:X509IssuerName><dsig:X509SerialNumber>1623825513212</dsig:X509SerialNumber></dsig:X509IssuerSerial><dsig:X509SubjectName>CN=idcs-acf13c306213452bbd0f83fbdb379f1f_keys, CN=Cloud9, CN=sslDomains</dsig:X509SubjectName></dsig:X509Data></dsig:KeyInfo></md:KeyDescriptor><md:KeyDescriptor use=\"encryption\"><dsig:KeyInfo><dsig:X509Data><dsig:X509Certificate>MIIDZjCCAk6gAwIBAgIGAXoTivr8MA0GCSqGSIb3DQEBCwUAMFkxEzARBgoJkiaJ\nk/IsZAEZFgNjb20xFjAUBgoJkiaJk/IsZAEZFgZvcmFjbGUxFTATBgoJkiaJk/Is\nZAEZFgVjbG91ZDETMBEGA1UEAxMKQ2xvdWQ5Q0EtMjAeFw0yMTA2MTYwNjM4MzJa\nFw0zMTA2MTYwNjM4MzJaMFsxEzARBgNVBAMTCnNzbERvbWFpbnMxDzANBgNVBAMT\nBkNsb3VkOTEzMDEGA1UEAwwqaWRjcy1hY2YxM2MzMDYyMTM0NTJiYmQwZjgzZmJk\nYjM3OWYxZl9rZXlzMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEArC8S\ng+4SF3ZKdRjveLNheqFlcXZIDV6N6FjMJEYVLyQLrRg8gAXIRqhpS/Ym9Bj8emyV\n+P/ZlHxmwPkI8DbJX+gsTg9xXG+DlfDSnNTdEOr90hrmg5D8NN3L0SpLughTnywP\nRbvk6Iv/21u8vALqQHfJYezBDqS/USpm/EP5+soaOlfGWAUaVnxqLtVC6g8kN3dj\n4hry3ZcfOZwAriEC2M1TTA8vOd7lkGRwEPjBL/QOXWKJU7LPYArKpuL75+gJteQ/\nyvDHIjbo7ZGaKGhhiCYcQ0gS1I6Fe3V8Mw3N+m7/TpXwGxyEo7GrR7M/j7jhJto8\n8eINDhBjiQkungCY3QIDAQABozIwMDAPBgNVHQ8BAf8EBQMDB/gAMB0GA1UdDgQW\nBBQ+PjQeG7CnRrqF9HXop07/756KrzANBgkqhkiG9w0BAQsFAAOCAQEAtQqQhc3t\nsdTHAlQRDkhnPS3N6x7xBpbjef+HOlta86DA9LSoYHBGnnZiwQHTgRbfUdUNllp7\nhITKMW942wTTwzY/owPukYoo4UAurOZcNkqFCzYJGqxS5bwQRSQO8iSRGwwooet7\n30tzevcilLWH7gF752unoPBoxgCPNxdxWxxSmHpiaQTUAy00sfba0KQ5SLCkXJHE\nOXpnQsXhe0qm7f/0Mo51cU9pIChj0Cfc/eEmwbcdPnPOCP4N8je+kqW79b3b90kc\nyUqMDTJw+4HKK62juXm4SXaePJYzx1Yb1VCSZpwOBjHxudp4WrkRPTXFRX248UtS\nUI5zVw1QsCmOnw==\n</dsig:X509Certificate><dsig:X509IssuerSerial><dsig:X509IssuerName>CN=Cloud9CA-2, DC=cloud, DC=oracle, DC=com</dsig:X509IssuerName><dsig:X509SerialNumber>1623825513212</dsig:X509SerialNumber></dsig:X509IssuerSerial><dsig:X509SubjectName>CN=idcs-acf13c306213452bbd0f83fbdb379f1f_keys, CN=Cloud9, CN=sslDomains</dsig:X509SubjectName></dsig:X509Data></dsig:KeyInfo><md:EncryptionMethod Algorithm=\"http://www.w3.org/2001/04/xmlenc#rsa-1_5\"/><md:EncryptionMethod Algorithm=\"http://www.w3.org/2001/04/xmlenc#aes128-cbc\"/><md:EncryptionMethod Algorithm=\"http://www.w3.org/2001/04/xmlenc#aes192-cbc\"/><md:EncryptionMethod Algorithm=\"http://www.w3.org/2001/04/xmlenc#aes256-cbc\"/><md:EncryptionMethod Algorithm=\"http://www.w3.org/2001/04/xmlenc#tripledes-cbc\"/></md:KeyDescriptor><md:SingleLogoutService Binding=\"urn:oasis:names:tc:SAML:2.0:bindings:HTTP-Redirect\" Location=\"https://idcs-acf13c306213452bbd0f83fbdb379f1f.unstable-alpha.identity.oci.oracleiaas.com/fed/v1/sp/slo\" ResponseLocation=\"https://idcs-acf13c306213452bbd0f83fbdb379f1f.unstable-alpha.identity.oci.oracleiaas.com/fed/v1/sp/slo\"/><md:SingleLogoutService Binding=\"urn:oasis:names:tc:SAML:2.0:bindings:HTTP-POST\" Location=\"https://idcs-acf13c306213452bbd0f83fbdb379f1f.unstable-alpha.identity.oci.oracleiaas.com/fed/v1/sp/slo\" ResponseLocation=\"https://idcs-acf13c306213452bbd0f83fbdb379f1f.unstable-alpha.identity.oci.oracleiaas.com/fed/v1/sp/slo\"/><md:AssertionConsumerService Binding=\"urn:oasis:names:tc:SAML:2.0:bindings:HTTP-POST\" Location=\"https://idcs-acf13c306213452bbd0f83fbdb379f1f.unstable-alpha.identity.oci.oracleiaas.com/fed/v1/sp/sso\" index=\"1\" isDefault=\"true\"/></md:SPSSODescriptor></md:EntityDescriptor>`},
		"name_id_format":                                  acctest.Representation{RepType: acctest.Required, Create: `nameIdFormat`, Update: `nameIdFormat2`},
		"partner_provider_id":                             acctest.Representation{RepType: acctest.Required, Create: `https://idcs-acf13c306213452bbd0f83fbdb379f1f.unstable-alpha.identity.oci.oracleiaas.com:443/fed`},
		"require_force_authn":                             acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"requires_encrypted_assertion":                    acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"saml_ho_krequired":                               acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"service_instance_identifier":                     acctest.Representation{RepType: acctest.Optional, Create: `serviceInstanceIdentifier`, Update: `serviceInstanceIdentifier2`},
		"shown_on_login_page":                             acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"signature_hash_algorithm":                        acctest.Representation{RepType: acctest.Optional, Create: `SHA-1`, Update: `SHA-256`},
		"signing_certificate":                             acctest.Representation{RepType: acctest.Required, Create: `MIIDZjCCAk6gAwIBAgIGAXoTivr8MA0GCSqGSIb3DQEBCwUAMFkxEzARBgoJkiaJk/IsZAEZFgNjb20xFjAUBgoJkiaJk/IsZAEZFgZvcmFjbGUxFTATBgoJkiaJk/IsZAEZFgVjbG91ZDETMBEGA1UEAxMKQ2xvdWQ5Q0EtMjAeFw0yMTA2MTYwNjM4MzJaFw0zMTA2MTYwNjM4MzJaMFsxEzARBgNVBAMTCnNzbERvbWFpbnMxDzANBgNVBAMTBkNsb3VkOTEzMDEGA1UEAwwqaWRjcy1hY2YxM2MzMDYyMTM0NTJiYmQwZjgzZmJkYjM3OWYxZl9rZXlzMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEArC8Sg+4SF3ZKdRjveLNheqFlcXZIDV6N6FjMJEYVLyQLrRg8gAXIRqhpS/Ym9Bj8emyV+P/ZlHxmwPkI8DbJX+gsTg9xXG+DlfDSnNTdEOr90hrmg5D8NN3L0SpLughTnywPRbvk6Iv/21u8vALqQHfJYezBDqS/USpm/EP5+soaOlfGWAUaVnxqLtVC6g8kN3dj4hry3ZcfOZwAriEC2M1TTA8vOd7lkGRwEPjBL/QOXWKJU7LPYArKpuL75+gJteQ/yvDHIjbo7ZGaKGhhiCYcQ0gS1I6Fe3V8Mw3N+m7/TpXwGxyEo7GrR7M/j7jhJto88eINDhBjiQkungCY3QIDAQABozIwMDAPBgNVHQ8BAf8EBQMDB/gAMB0GA1UdDgQWBBQ+PjQeG7CnRrqF9HXop07/756KrzANBgkqhkiG9w0BAQsFAAOCAQEAtQqQhc3tsdTHAlQRDkhnPS3N6x7xBpbjef+HOlta86DA9LSoYHBGnnZiwQHTgRbfUdUNllp7hITKMW942wTTwzY/owPukYoo4UAurOZcNkqFCzYJGqxS5bwQRSQO8iSRGwwooet730tzevcilLWH7gF752unoPBoxgCPNxdxWxxSmHpiaQTUAy00sfba0KQ5SLCkXJHEOXpnQsXhe0qm7f/0Mo51cU9pIChj0Cfc/eEmwbcdPnPOCP4N8je+kqW79b3b90kcyUqMDTJw+4HKK62juXm4SXaePJYzx1Yb1VCSZpwOBjHxudp4WrkRPTXFRX248UtSUI5zVw1QsCmOnw==`},
		"succinct_id":                                     acctest.Representation{RepType: acctest.Optional, Create: `succinctId`},
		"tags":                                            acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityDomainsIdentityProviderTagsRepresentation},
		"type":                                            acctest.Representation{RepType: acctest.Optional, Create: `SAML`},
		"urnietfparamsscimschemasoracleidcsextensionsocial_identity_provider": acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityDomainsIdentityProviderUrnietfparamsscimschemasoracleidcsextensionsocialIdentityProviderRepresentation},
		"urnietfparamsscimschemasoracleidcsextensionx509identity_provider":    acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityDomainsIdentityProviderUrnietfparamsscimschemasoracleidcsextensionx509IdentityProviderRepresentation},
		"user_mapping_method":          acctest.Representation{RepType: acctest.Optional, Create: `NameIDToUserAttribute`, Update: `AssertionAttributeToUserAttribute`},
		"user_mapping_store_attribute": acctest.Representation{RepType: acctest.Optional, Create: `userName`, Update: `name.givenName`},
		"lifecycle":                    acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreChangeForIdentityDomainsIdentityProvider},
	}
	ignoreChangeForIdentityDomainsIdentityProvider = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{
			`schemas`,
		}},
	}
	IdentityDomainsIdentityProviderJitUserProvAssignedGroupsRepresentation = map[string]interface{}{
		"value": acctest.Representation{RepType: acctest.Required, Create: `${oci_identity_domains_group.test_group.id}`},
	}
	IdentityDomainsIdentityProviderJitUserProvGroupMappingsRepresentation = map[string]interface{}{
		"idp_group": acctest.Representation{RepType: acctest.Required, Create: `idpGroup`, Update: `idpGroup2`},
		"value":     acctest.Representation{RepType: acctest.Required, Create: `${oci_identity_domains_group.test_group.id}`},
	}
	IdentityDomainsIdentityProviderTagsRepresentation = map[string]interface{}{
		"key":   acctest.Representation{RepType: acctest.Required, Create: `key`, Update: `key2`},
		"value": acctest.Representation{RepType: acctest.Required, Create: `value`, Update: `value2`},
	}
	IdentityDomainsIdentityProviderUrnietfparamsscimschemasoracleidcsextensionsocialIdentityProviderRepresentation = map[string]interface{}{
		"account_linking_enabled":      acctest.Representation{RepType: acctest.Required, Create: `false`, Update: `true`},
		"consumer_key":                 acctest.Representation{RepType: acctest.Required, Create: `consumerKey`, Update: `consumerKey2`},
		"consumer_secret":              acctest.Representation{RepType: acctest.Required, Create: `consumerSecret`, Update: `consumerSecret2`},
		"registration_enabled":         acctest.Representation{RepType: acctest.Required, Create: `false`, Update: `true`},
		"service_provider_name":        acctest.Representation{RepType: acctest.Required, Create: `serviceProviderName`},
		"access_token_url":             acctest.Representation{RepType: acctest.Optional, Create: `accessTokenUrl`, Update: `accessTokenUrl2`},
		"admin_scope":                  acctest.Representation{RepType: acctest.Optional, Create: []string{`adminScope`}, Update: []string{`adminScope2`}},
		"authz_url":                    acctest.Representation{RepType: acctest.Optional, Create: `authzUrl`, Update: `authzUrl2`},
		"client_credential_in_payload": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"clock_skew_in_seconds":        acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"discovery_url":                acctest.Representation{RepType: acctest.Optional, Create: `discoveryUrl`, Update: `discoveryUrl2`},
		"id_attribute":                 acctest.Representation{RepType: acctest.Optional, Create: `idAttribute`},
		"profile_url":                  acctest.Representation{RepType: acctest.Optional, Create: `profileUrl`, Update: `profileUrl2`},
		"redirect_url":                 acctest.Representation{RepType: acctest.Optional, Create: `redirectUrl`, Update: `redirectUrl2`},
		"scope":                        acctest.Representation{RepType: acctest.Optional, Create: []string{`scope`}, Update: []string{`scope2`}},
		"status":                       acctest.Representation{RepType: acctest.Optional, Create: `created`, Update: `deleted`},
	}
	IdentityDomainsIdentityProviderUrnietfparamsscimschemasoracleidcsextensionx509IdentityProviderRepresentation = map[string]interface{}{
		"cert_match_attribute":               acctest.Representation{RepType: acctest.Required, Create: `certMatchAttribute`, Update: `certMatchAttribute2`},
		"signing_certificate_chain":          acctest.Representation{RepType: acctest.Required, Create: []string{`signingCertificateChain`}, Update: []string{`signingCertificateChain2`}},
		"user_match_attribute":               acctest.Representation{RepType: acctest.Required, Create: `userMatchAttribute`, Update: `userMatchAttribute2`},
		"crl_check_on_ocsp_failure_enabled":  acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"crl_enabled":                        acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"crl_location":                       acctest.Representation{RepType: acctest.Optional, Create: `crlLocation`, Update: `crlLocation2`},
		"crl_reload_duration":                acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"eku_validation_enabled":             acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"eku_values":                         acctest.Representation{RepType: acctest.Optional, Create: []string{`server_auth`}, Update: []string{`client_auth`}},
		"ocsp_allow_unknown_response_status": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"ocsp_enable_signed_response":        acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"ocsp_enabled":                       acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"ocsp_responder_url":                 acctest.Representation{RepType: acctest.Optional, Create: `ocspResponderURL`, Update: `ocspResponderURL2`},
		"ocsp_revalidate_time":               acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"ocsp_server_name":                   acctest.Representation{RepType: acctest.Optional, Create: `ocspServerName`, Update: `ocspServerName2`},
		"ocsp_trust_cert_chain":              acctest.Representation{RepType: acctest.Optional, Create: []string{`ocspTrustCertChain`}, Update: []string{`ocspTrustCertChain2`}},
		"other_cert_match_attribute":         acctest.Representation{RepType: acctest.Optional, Create: `otherCertMatchAttribute`, Update: `otherCertMatchAttribute2`},
	}

	IdentityDomainsIdentityProviderResourceDependencies = TestDomainDependencies + acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_group", "test_group", acctest.Required, acctest.Create, IdentityDomainsGroupRepresentation)
)

// issue-routing-tag: identity_domains/default
func TestIdentityDomainsIdentityProviderResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestIdentityDomainsIdentityProviderResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_identity_domains_identity_provider.test_identity_provider"
	datasourceName := "data.oci_identity_domains_identity_providers.test_identity_providers"
	singularDatasourceName := "data.oci_identity_domains_identity_provider.test_identity_provider"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+IdentityDomainsIdentityProviderResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_identity_provider", "test_identity_provider", acctest.Optional, acctest.Create, IdentityDomainsIdentityProviderRepresentation), "identitydomains", "identityProvider", t)

	print(config + compartmentIdVariableStr + IdentityDomainsIdentityProviderResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_identity_provider", "test_identity_provider", acctest.Optional, acctest.Create, IdentityDomainsIdentityProviderRepresentation))
	acctest.ResourceTest(t, testAccCheckIdentityDomainsIdentityProviderDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsIdentityProviderResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_identity_provider", "test_identity_provider", acctest.Required, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(IdentityDomainsIdentityProviderRepresentation, map[string]interface{}{
						"enabled": acctest.Representation{RepType: acctest.Required, Create: `false`},
					},
					),
				),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "enabled", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttr(resourceName, "partner_name", "partnerName"),
				resource.TestMatchResourceAttr(resourceName, "schemas.#", regexp.MustCompile("[1-9]+")),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsIdentityProviderResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsIdentityProviderResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_identity_provider", "test_identity_provider", acctest.Optional, acctest.Create, IdentityDomainsIdentityProviderRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "assertion_attribute", "assertionAttribute"),
				resource.TestCheckResourceAttr(resourceName, "attribute_sets.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "authn_request_binding", "Redirect"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "encryption_certificate", "MIIDZjCCAk6gAwIBAgIGAXoTivr8MA0GCSqGSIb3DQEBCwUAMFkxEzARBgoJkiaJk/IsZAEZFgNjb20xFjAUBgoJkiaJk/IsZAEZFgZvcmFjbGUxFTATBgoJkiaJk/IsZAEZFgVjbG91ZDETMBEGA1UEAxMKQ2xvdWQ5Q0EtMjAeFw0yMTA2MTYwNjM4MzJaFw0zMTA2MTYwNjM4MzJaMFsxEzARBgNVBAMTCnNzbERvbWFpbnMxDzANBgNVBAMTBkNsb3VkOTEzMDEGA1UEAwwqaWRjcy1hY2YxM2MzMDYyMTM0NTJiYmQwZjgzZmJkYjM3OWYxZl9rZXlzMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEArC8Sg+4SF3ZKdRjveLNheqFlcXZIDV6N6FjMJEYVLyQLrRg8gAXIRqhpS/Ym9Bj8emyV+P/ZlHxmwPkI8DbJX+gsTg9xXG+DlfDSnNTdEOr90hrmg5D8NN3L0SpLughTnywPRbvk6Iv/21u8vALqQHfJYezBDqS/USpm/EP5+soaOlfGWAUaVnxqLtVC6g8kN3dj4hry3ZcfOZwAriEC2M1TTA8vOd7lkGRwEPjBL/QOXWKJU7LPYArKpuL75+gJteQ/yvDHIjbo7ZGaKGhhiCYcQ0gS1I6Fe3V8Mw3N+m7/TpXwGxyEo7GrR7M/j7jhJto88eINDhBjiQkungCY3QIDAQABozIwMDAPBgNVHQ8BAf8EBQMDB/gAMB0GA1UdDgQWBBQ+PjQeG7CnRrqF9HXop07/756KrzANBgkqhkiG9w0BAQsFAAOCAQEAtQqQhc3tsdTHAlQRDkhnPS3N6x7xBpbjef+HOlta86DA9LSoYHBGnnZiwQHTgRbfUdUNllp7hITKMW942wTTwzY/owPukYoo4UAurOZcNkqFCzYJGqxS5bwQRSQO8iSRGwwooet730tzevcilLWH7gF752unoPBoxgCPNxdxWxxSmHpiaQTUAy00sfba0KQ5SLCkXJHEOXpnQsXhe0qm7f/0Mo51cU9pIChj0Cfc/eEmwbcdPnPOCP4N8je+kqW79b3b90kcyUqMDTJw+4HKK62juXm4SXaePJYzx1Yb1VCSZpwOBjHxudp4WrkRPTXFRX248UtSUI5zVw1QsCmOnw=="),
				resource.TestCheckResourceAttrSet(resourceName, "external_id"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "icon_url", "https://something.com/iconUrl.png"),
				resource.TestCheckResourceAttr(resourceName, "idcs_created_by.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttr(resourceName, "idp_sso_url", "https://idpSsoUrl.com"),
				resource.TestCheckResourceAttr(resourceName, "include_signing_cert_in_signature", "false"),
				resource.TestCheckResourceAttr(resourceName, "jit_user_prov_assigned_groups.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "jit_user_prov_assigned_groups.0.value"),
				resource.TestCheckResourceAttr(resourceName, "jit_user_prov_attribute_update_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "jit_user_prov_create_user_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "jit_user_prov_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "jit_user_prov_group_assertion_attribute_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "jit_user_prov_group_assignment_method", "Overwrite"),
				resource.TestCheckResourceAttr(resourceName, "jit_user_prov_group_mapping_mode", "implicit"),
				resource.TestCheckResourceAttr(resourceName, "jit_user_prov_group_mappings.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "jit_user_prov_group_mappings.0.idp_group", "idpGroup"),
				resource.TestCheckResourceAttrSet(resourceName, "jit_user_prov_group_mappings.0.ref"),
				resource.TestCheckResourceAttrSet(resourceName, "jit_user_prov_group_mappings.0.value"),
				resource.TestCheckResourceAttr(resourceName, "jit_user_prov_group_saml_attribute_name", "jitUserProvGroupSAMLAttributeName"),
				resource.TestCheckResourceAttr(resourceName, "jit_user_prov_group_static_list_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "jit_user_prov_ignore_error_on_absent_groups", "false"),
				resource.TestCheckResourceAttr(resourceName, "logout_binding", "Redirect"),
				resource.TestCheckResourceAttr(resourceName, "logout_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "logout_request_url", "https://logoutRequestUrl.com"),
				resource.TestCheckResourceAttr(resourceName, "logout_response_url", "https://logoutResponseUrl.com"),
				resource.TestCheckResourceAttr(resourceName, "metadata", "<md:EntityDescriptor xmlns:md=\"urn:oasis:names:tc:SAML:2.0:metadata\" xmlns:dsig=\"http://www.w3.org/2000/09/xmldsig#\" xmlns:enc=\"http://www.w3.org/2001/04/xmlenc#\" xmlns:mdattr=\"urn:oasis:names:tc:SAML:metadata:attribute\" xmlns:query=\"urn:oasis:names:tc:SAML:metadata:ext:query\" xmlns:saml=\"urn:oasis:names:tc:SAML:2.0:assertion\" xmlns:x500=\"urn:oasis:names:tc:SAML:2.0:profiles:attribute:X500\" xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" ID=\"id-zzU36agM7bKRB32xe6Ronm131S0-\" cacheDuration=\"P3633DT0H0M0S\" entityID=\"https://idcs-acf13c306213452bbd0f83fbdb379f1f.unstable-alpha.identity.oci.oracleiaas.com:443/fed\" validUntil=\"2031-06-16T06:38:32Z\"><dsig:Signature><dsig:SignedInfo><dsig:CanonicalizationMethod Algorithm=\"http://www.w3.org/2001/10/xml-exc-c14n#\"/><dsig:SignatureMethod Algorithm=\"http://www.w3.org/2001/04/xmldsig-more#rsa-sha256\"/><dsig:Reference URI=\"#id-zzU36agM7bKRB32xe6Ronm131S0-\"><dsig:Transforms><dsig:Transform Algorithm=\"http://www.w3.org/2000/09/xmldsig#enveloped-signature\"/><dsig:Transform Algorithm=\"http://www.w3.org/2001/10/xml-exc-c14n#\"/></dsig:Transforms><dsig:DigestMethod Algorithm=\"http://www.w3.org/2001/04/xmlenc#sha256\"/><dsig:DigestValue>NZnYsjLx3UbuL43iu3jo0mJUg/Rv9DTPNB5IQPRaD6g=</dsig:DigestValue></dsig:Reference></dsig:SignedInfo><dsig:SignatureValue>KRIgTD7//x/uT73veS0iGcWWw8uprjd+MtREu3vlbFTk0BNgkeSOYItx2LDQhnHP\nO0zsTmtOHlVIsDXQL3KysHwzYndIuMJtETqEC6NpMw3ZF108IK0eT+o/2xC9u13/\nGq10z/MagGvco1mM/RIzX5e2omGyZcKARiDoeNPwg2znmV0WcifntVqn4Y0rnWM7\no0M5HFHZQEgICdTJbC5d6DwLgfnI4ck505fHNRYLsRqj9IGLukKx9kocSG1xzCye\nHlffU4CDyEA7dptEUH59dZmY0Xy35/aepNc7W6IovWsJ2Otr+qDUp207ZCKuISF0\nMEX5hX5VJzVlHDwxkEcYCA==</dsig:SignatureValue><dsig:KeyInfo><dsig:X509Data><dsig:X509Certificate>MIIDZjCCAk6gAwIBAgIGAXoTivr8MA0GCSqGSIb3DQEBCwUAMFkxEzARBgoJkiaJ\nk/IsZAEZFgNjb20xFjAUBgoJkiaJk/IsZAEZFgZvcmFjbGUxFTATBgoJkiaJk/Is\nZAEZFgVjbG91ZDETMBEGA1UEAxMKQ2xvdWQ5Q0EtMjAeFw0yMTA2MTYwNjM4MzJa\nFw0zMTA2MTYwNjM4MzJaMFsxEzARBgNVBAMTCnNzbERvbWFpbnMxDzANBgNVBAMT\nBkNsb3VkOTEzMDEGA1UEAwwqaWRjcy1hY2YxM2MzMDYyMTM0NTJiYmQwZjgzZmJk\nYjM3OWYxZl9rZXlzMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEArC8S\ng+4SF3ZKdRjveLNheqFlcXZIDV6N6FjMJEYVLyQLrRg8gAXIRqhpS/Ym9Bj8emyV\n+P/ZlHxmwPkI8DbJX+gsTg9xXG+DlfDSnNTdEOr90hrmg5D8NN3L0SpLughTnywP\nRbvk6Iv/21u8vALqQHfJYezBDqS/USpm/EP5+soaOlfGWAUaVnxqLtVC6g8kN3dj\n4hry3ZcfOZwAriEC2M1TTA8vOd7lkGRwEPjBL/QOXWKJU7LPYArKpuL75+gJteQ/\nyvDHIjbo7ZGaKGhhiCYcQ0gS1I6Fe3V8Mw3N+m7/TpXwGxyEo7GrR7M/j7jhJto8\n8eINDhBjiQkungCY3QIDAQABozIwMDAPBgNVHQ8BAf8EBQMDB/gAMB0GA1UdDgQW\nBBQ+PjQeG7CnRrqF9HXop07/756KrzANBgkqhkiG9w0BAQsFAAOCAQEAtQqQhc3t\nsdTHAlQRDkhnPS3N6x7xBpbjef+HOlta86DA9LSoYHBGnnZiwQHTgRbfUdUNllp7\nhITKMW942wTTwzY/owPukYoo4UAurOZcNkqFCzYJGqxS5bwQRSQO8iSRGwwooet7\n30tzevcilLWH7gF752unoPBoxgCPNxdxWxxSmHpiaQTUAy00sfba0KQ5SLCkXJHE\nOXpnQsXhe0qm7f/0Mo51cU9pIChj0Cfc/eEmwbcdPnPOCP4N8je+kqW79b3b90kc\nyUqMDTJw+4HKK62juXm4SXaePJYzx1Yb1VCSZpwOBjHxudp4WrkRPTXFRX248UtS\nUI5zVw1QsCmOnw==</dsig:X509Certificate></dsig:X509Data></dsig:KeyInfo></dsig:Signature><md:IDPSSODescriptor WantAuthnRequestsSigned=\"false\" protocolSupportEnumeration=\"urn:oasis:names:tc:SAML:2.0:protocol\"><md:KeyDescriptor use=\"signing\"><dsig:KeyInfo><dsig:X509Data><dsig:X509Certificate>MIIDZjCCAk6gAwIBAgIGAXoTivr8MA0GCSqGSIb3DQEBCwUAMFkxEzARBgoJkiaJ\nk/IsZAEZFgNjb20xFjAUBgoJkiaJk/IsZAEZFgZvcmFjbGUxFTATBgoJkiaJk/Is\nZAEZFgVjbG91ZDETMBEGA1UEAxMKQ2xvdWQ5Q0EtMjAeFw0yMTA2MTYwNjM4MzJa\nFw0zMTA2MTYwNjM4MzJaMFsxEzARBgNVBAMTCnNzbERvbWFpbnMxDzANBgNVBAMT\nBkNsb3VkOTEzMDEGA1UEAwwqaWRjcy1hY2YxM2MzMDYyMTM0NTJiYmQwZjgzZmJk\nYjM3OWYxZl9rZXlzMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEArC8S\ng+4SF3ZKdRjveLNheqFlcXZIDV6N6FjMJEYVLyQLrRg8gAXIRqhpS/Ym9Bj8emyV\n+P/ZlHxmwPkI8DbJX+gsTg9xXG+DlfDSnNTdEOr90hrmg5D8NN3L0SpLughTnywP\nRbvk6Iv/21u8vALqQHfJYezBDqS/USpm/EP5+soaOlfGWAUaVnxqLtVC6g8kN3dj\n4hry3ZcfOZwAriEC2M1TTA8vOd7lkGRwEPjBL/QOXWKJU7LPYArKpuL75+gJteQ/\nyvDHIjbo7ZGaKGhhiCYcQ0gS1I6Fe3V8Mw3N+m7/TpXwGxyEo7GrR7M/j7jhJto8\n8eINDhBjiQkungCY3QIDAQABozIwMDAPBgNVHQ8BAf8EBQMDB/gAMB0GA1UdDgQW\nBBQ+PjQeG7CnRrqF9HXop07/756KrzANBgkqhkiG9w0BAQsFAAOCAQEAtQqQhc3t\nsdTHAlQRDkhnPS3N6x7xBpbjef+HOlta86DA9LSoYHBGnnZiwQHTgRbfUdUNllp7\nhITKMW942wTTwzY/owPukYoo4UAurOZcNkqFCzYJGqxS5bwQRSQO8iSRGwwooet7\n30tzevcilLWH7gF752unoPBoxgCPNxdxWxxSmHpiaQTUAy00sfba0KQ5SLCkXJHE\nOXpnQsXhe0qm7f/0Mo51cU9pIChj0Cfc/eEmwbcdPnPOCP4N8je+kqW79b3b90kc\nyUqMDTJw+4HKK62juXm4SXaePJYzx1Yb1VCSZpwOBjHxudp4WrkRPTXFRX248UtS\nUI5zVw1QsCmOnw==\n</dsig:X509Certificate><dsig:X509IssuerSerial><dsig:X509IssuerName>CN=Cloud9CA-2, DC=cloud, DC=oracle, DC=com</dsig:X509IssuerName><dsig:X509SerialNumber>1623825513212</dsig:X509SerialNumber></dsig:X509IssuerSerial><dsig:X509SubjectName>CN=idcs-acf13c306213452bbd0f83fbdb379f1f_keys, CN=Cloud9, CN=sslDomains</dsig:X509SubjectName></dsig:X509Data></dsig:KeyInfo></md:KeyDescriptor><md:KeyDescriptor use=\"encryption\"><dsig:KeyInfo><dsig:X509Data><dsig:X509Certificate>MIIDZjCCAk6gAwIBAgIGAXoTivr8MA0GCSqGSIb3DQEBCwUAMFkxEzARBgoJkiaJ\nk/IsZAEZFgNjb20xFjAUBgoJkiaJk/IsZAEZFgZvcmFjbGUxFTATBgoJkiaJk/Is\nZAEZFgVjbG91ZDETMBEGA1UEAxMKQ2xvdWQ5Q0EtMjAeFw0yMTA2MTYwNjM4MzJa\nFw0zMTA2MTYwNjM4MzJaMFsxEzARBgNVBAMTCnNzbERvbWFpbnMxDzANBgNVBAMT\nBkNsb3VkOTEzMDEGA1UEAwwqaWRjcy1hY2YxM2MzMDYyMTM0NTJiYmQwZjgzZmJk\nYjM3OWYxZl9rZXlzMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEArC8S\ng+4SF3ZKdRjveLNheqFlcXZIDV6N6FjMJEYVLyQLrRg8gAXIRqhpS/Ym9Bj8emyV\n+P/ZlHxmwPkI8DbJX+gsTg9xXG+DlfDSnNTdEOr90hrmg5D8NN3L0SpLughTnywP\nRbvk6Iv/21u8vALqQHfJYezBDqS/USpm/EP5+soaOlfGWAUaVnxqLtVC6g8kN3dj\n4hry3ZcfOZwAriEC2M1TTA8vOd7lkGRwEPjBL/QOXWKJU7LPYArKpuL75+gJteQ/\nyvDHIjbo7ZGaKGhhiCYcQ0gS1I6Fe3V8Mw3N+m7/TpXwGxyEo7GrR7M/j7jhJto8\n8eINDhBjiQkungCY3QIDAQABozIwMDAPBgNVHQ8BAf8EBQMDB/gAMB0GA1UdDgQW\nBBQ+PjQeG7CnRrqF9HXop07/756KrzANBgkqhkiG9w0BAQsFAAOCAQEAtQqQhc3t\nsdTHAlQRDkhnPS3N6x7xBpbjef+HOlta86DA9LSoYHBGnnZiwQHTgRbfUdUNllp7\nhITKMW942wTTwzY/owPukYoo4UAurOZcNkqFCzYJGqxS5bwQRSQO8iSRGwwooet7\n30tzevcilLWH7gF752unoPBoxgCPNxdxWxxSmHpiaQTUAy00sfba0KQ5SLCkXJHE\nOXpnQsXhe0qm7f/0Mo51cU9pIChj0Cfc/eEmwbcdPnPOCP4N8je+kqW79b3b90kc\nyUqMDTJw+4HKK62juXm4SXaePJYzx1Yb1VCSZpwOBjHxudp4WrkRPTXFRX248UtS\nUI5zVw1QsCmOnw==\n</dsig:X509Certificate><dsig:X509IssuerSerial><dsig:X509IssuerName>CN=Cloud9CA-2, DC=cloud, DC=oracle, DC=com</dsig:X509IssuerName><dsig:X509SerialNumber>1623825513212</dsig:X509SerialNumber></dsig:X509IssuerSerial><dsig:X509SubjectName>CN=idcs-acf13c306213452bbd0f83fbdb379f1f_keys, CN=Cloud9, CN=sslDomains</dsig:X509SubjectName></dsig:X509Data></dsig:KeyInfo><md:EncryptionMethod Algorithm=\"http://www.w3.org/2001/04/xmlenc#rsa-1_5\"/><md:EncryptionMethod Algorithm=\"http://www.w3.org/2001/04/xmlenc#aes128-cbc\"/><md:EncryptionMethod Algorithm=\"http://www.w3.org/2001/04/xmlenc#aes192-cbc\"/><md:EncryptionMethod Algorithm=\"http://www.w3.org/2001/04/xmlenc#aes256-cbc\"/><md:EncryptionMethod Algorithm=\"http://www.w3.org/2001/04/xmlenc#tripledes-cbc\"/></md:KeyDescriptor><md:SingleLogoutService Binding=\"urn:oasis:names:tc:SAML:2.0:bindings:HTTP-Redirect\" Location=\"https://idcs-acf13c306213452bbd0f83fbdb379f1f.unstable-alpha.identity.oci.oracleiaas.com/fed/v1/idp/slo\" ResponseLocation=\"https://idcs-acf13c306213452bbd0f83fbdb379f1f.unstable-alpha.identity.oci.oracleiaas.com/fed/v1/idp/slo\"/><md:SingleLogoutService Binding=\"urn:oasis:names:tc:SAML:2.0:bindings:HTTP-POST\" Location=\"https://idcs-acf13c306213452bbd0f83fbdb379f1f.unstable-alpha.identity.oci.oracleiaas.com/fed/v1/idp/slo\" ResponseLocation=\"https://idcs-acf13c306213452bbd0f83fbdb379f1f.unstable-alpha.identity.oci.oracleiaas.com/fed/v1/idp/slo\"/><md:SingleSignOnService Binding=\"urn:oasis:names:tc:SAML:2.0:bindings:HTTP-Redirect\" Location=\"https://idcs-acf13c306213452bbd0f83fbdb379f1f.unstable-alpha.identity.oci.oracleiaas.com/fed/v1/idp/sso\"/><md:SingleSignOnService Binding=\"urn:oasis:names:tc:SAML:2.0:bindings:HTTP-POST\" Location=\"https://idcs-acf13c306213452bbd0f83fbdb379f1f.unstable-alpha.identity.oci.oracleiaas.com/fed/v1/idp/sso\"/></md:IDPSSODescriptor><md:SPSSODescriptor AuthnRequestsSigned=\"true\" WantAssertionsSigned=\"true\" protocolSupportEnumeration=\"urn:oasis:names:tc:SAML:2.0:protocol\"><md:KeyDescriptor use=\"signing\"><dsig:KeyInfo><dsig:X509Data><dsig:X509Certificate>MIIDZjCCAk6gAwIBAgIGAXoTivr8MA0GCSqGSIb3DQEBCwUAMFkxEzARBgoJkiaJ\nk/IsZAEZFgNjb20xFjAUBgoJkiaJk/IsZAEZFgZvcmFjbGUxFTATBgoJkiaJk/Is\nZAEZFgVjbG91ZDETMBEGA1UEAxMKQ2xvdWQ5Q0EtMjAeFw0yMTA2MTYwNjM4MzJa\nFw0zMTA2MTYwNjM4MzJaMFsxEzARBgNVBAMTCnNzbERvbWFpbnMxDzANBgNVBAMT\nBkNsb3VkOTEzMDEGA1UEAwwqaWRjcy1hY2YxM2MzMDYyMTM0NTJiYmQwZjgzZmJk\nYjM3OWYxZl9rZXlzMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEArC8S\ng+4SF3ZKdRjveLNheqFlcXZIDV6N6FjMJEYVLyQLrRg8gAXIRqhpS/Ym9Bj8emyV\n+P/ZlHxmwPkI8DbJX+gsTg9xXG+DlfDSnNTdEOr90hrmg5D8NN3L0SpLughTnywP\nRbvk6Iv/21u8vALqQHfJYezBDqS/USpm/EP5+soaOlfGWAUaVnxqLtVC6g8kN3dj\n4hry3ZcfOZwAriEC2M1TTA8vOd7lkGRwEPjBL/QOXWKJU7LPYArKpuL75+gJteQ/\nyvDHIjbo7ZGaKGhhiCYcQ0gS1I6Fe3V8Mw3N+m7/TpXwGxyEo7GrR7M/j7jhJto8\n8eINDhBjiQkungCY3QIDAQABozIwMDAPBgNVHQ8BAf8EBQMDB/gAMB0GA1UdDgQW\nBBQ+PjQeG7CnRrqF9HXop07/756KrzANBgkqhkiG9w0BAQsFAAOCAQEAtQqQhc3t\nsdTHAlQRDkhnPS3N6x7xBpbjef+HOlta86DA9LSoYHBGnnZiwQHTgRbfUdUNllp7\nhITKMW942wTTwzY/owPukYoo4UAurOZcNkqFCzYJGqxS5bwQRSQO8iSRGwwooet7\n30tzevcilLWH7gF752unoPBoxgCPNxdxWxxSmHpiaQTUAy00sfba0KQ5SLCkXJHE\nOXpnQsXhe0qm7f/0Mo51cU9pIChj0Cfc/eEmwbcdPnPOCP4N8je+kqW79b3b90kc\nyUqMDTJw+4HKK62juXm4SXaePJYzx1Yb1VCSZpwOBjHxudp4WrkRPTXFRX248UtS\nUI5zVw1QsCmOnw==\n</dsig:X509Certificate><dsig:X509IssuerSerial><dsig:X509IssuerName>CN=Cloud9CA-2, DC=cloud, DC=oracle, DC=com</dsig:X509IssuerName><dsig:X509SerialNumber>1623825513212</dsig:X509SerialNumber></dsig:X509IssuerSerial><dsig:X509SubjectName>CN=idcs-acf13c306213452bbd0f83fbdb379f1f_keys, CN=Cloud9, CN=sslDomains</dsig:X509SubjectName></dsig:X509Data></dsig:KeyInfo></md:KeyDescriptor><md:KeyDescriptor use=\"encryption\"><dsig:KeyInfo><dsig:X509Data><dsig:X509Certificate>MIIDZjCCAk6gAwIBAgIGAXoTivr8MA0GCSqGSIb3DQEBCwUAMFkxEzARBgoJkiaJ\nk/IsZAEZFgNjb20xFjAUBgoJkiaJk/IsZAEZFgZvcmFjbGUxFTATBgoJkiaJk/Is\nZAEZFgVjbG91ZDETMBEGA1UEAxMKQ2xvdWQ5Q0EtMjAeFw0yMTA2MTYwNjM4MzJa\nFw0zMTA2MTYwNjM4MzJaMFsxEzARBgNVBAMTCnNzbERvbWFpbnMxDzANBgNVBAMT\nBkNsb3VkOTEzMDEGA1UEAwwqaWRjcy1hY2YxM2MzMDYyMTM0NTJiYmQwZjgzZmJk\nYjM3OWYxZl9rZXlzMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEArC8S\ng+4SF3ZKdRjveLNheqFlcXZIDV6N6FjMJEYVLyQLrRg8gAXIRqhpS/Ym9Bj8emyV\n+P/ZlHxmwPkI8DbJX+gsTg9xXG+DlfDSnNTdEOr90hrmg5D8NN3L0SpLughTnywP\nRbvk6Iv/21u8vALqQHfJYezBDqS/USpm/EP5+soaOlfGWAUaVnxqLtVC6g8kN3dj\n4hry3ZcfOZwAriEC2M1TTA8vOd7lkGRwEPjBL/QOXWKJU7LPYArKpuL75+gJteQ/\nyvDHIjbo7ZGaKGhhiCYcQ0gS1I6Fe3V8Mw3N+m7/TpXwGxyEo7GrR7M/j7jhJto8\n8eINDhBjiQkungCY3QIDAQABozIwMDAPBgNVHQ8BAf8EBQMDB/gAMB0GA1UdDgQW\nBBQ+PjQeG7CnRrqF9HXop07/756KrzANBgkqhkiG9w0BAQsFAAOCAQEAtQqQhc3t\nsdTHAlQRDkhnPS3N6x7xBpbjef+HOlta86DA9LSoYHBGnnZiwQHTgRbfUdUNllp7\nhITKMW942wTTwzY/owPukYoo4UAurOZcNkqFCzYJGqxS5bwQRSQO8iSRGwwooet7\n30tzevcilLWH7gF752unoPBoxgCPNxdxWxxSmHpiaQTUAy00sfba0KQ5SLCkXJHE\nOXpnQsXhe0qm7f/0Mo51cU9pIChj0Cfc/eEmwbcdPnPOCP4N8je+kqW79b3b90kc\nyUqMDTJw+4HKK62juXm4SXaePJYzx1Yb1VCSZpwOBjHxudp4WrkRPTXFRX248UtS\nUI5zVw1QsCmOnw==\n</dsig:X509Certificate><dsig:X509IssuerSerial><dsig:X509IssuerName>CN=Cloud9CA-2, DC=cloud, DC=oracle, DC=com</dsig:X509IssuerName><dsig:X509SerialNumber>1623825513212</dsig:X509SerialNumber></dsig:X509IssuerSerial><dsig:X509SubjectName>CN=idcs-acf13c306213452bbd0f83fbdb379f1f_keys, CN=Cloud9, CN=sslDomains</dsig:X509SubjectName></dsig:X509Data></dsig:KeyInfo><md:EncryptionMethod Algorithm=\"http://www.w3.org/2001/04/xmlenc#rsa-1_5\"/><md:EncryptionMethod Algorithm=\"http://www.w3.org/2001/04/xmlenc#aes128-cbc\"/><md:EncryptionMethod Algorithm=\"http://www.w3.org/2001/04/xmlenc#aes192-cbc\"/><md:EncryptionMethod Algorithm=\"http://www.w3.org/2001/04/xmlenc#aes256-cbc\"/><md:EncryptionMethod Algorithm=\"http://www.w3.org/2001/04/xmlenc#tripledes-cbc\"/></md:KeyDescriptor><md:SingleLogoutService Binding=\"urn:oasis:names:tc:SAML:2.0:bindings:HTTP-Redirect\" Location=\"https://idcs-acf13c306213452bbd0f83fbdb379f1f.unstable-alpha.identity.oci.oracleiaas.com/fed/v1/sp/slo\" ResponseLocation=\"https://idcs-acf13c306213452bbd0f83fbdb379f1f.unstable-alpha.identity.oci.oracleiaas.com/fed/v1/sp/slo\"/><md:SingleLogoutService Binding=\"urn:oasis:names:tc:SAML:2.0:bindings:HTTP-POST\" Location=\"https://idcs-acf13c306213452bbd0f83fbdb379f1f.unstable-alpha.identity.oci.oracleiaas.com/fed/v1/sp/slo\" ResponseLocation=\"https://idcs-acf13c306213452bbd0f83fbdb379f1f.unstable-alpha.identity.oci.oracleiaas.com/fed/v1/sp/slo\"/><md:AssertionConsumerService Binding=\"urn:oasis:names:tc:SAML:2.0:bindings:HTTP-POST\" Location=\"https://idcs-acf13c306213452bbd0f83fbdb379f1f.unstable-alpha.identity.oci.oracleiaas.com/fed/v1/sp/sso\" index=\"1\" isDefault=\"true\"/></md:SPSSODescriptor></md:EntityDescriptor>"),
				resource.TestCheckResourceAttr(resourceName, "name_id_format", "nameIdFormat"),
				resource.TestCheckResourceAttr(resourceName, "partner_name", "partnerName"),
				resource.TestCheckResourceAttrSet(resourceName, "partner_provider_id"),
				resource.TestCheckResourceAttr(resourceName, "require_force_authn", "false"),
				resource.TestCheckResourceAttr(resourceName, "requires_encrypted_assertion", "false"),
				resource.TestCheckResourceAttr(resourceName, "saml_ho_krequired", "false"),
				resource.TestMatchResourceAttr(resourceName, "schemas.#", regexp.MustCompile("[1-9]+")),
				resource.TestCheckResourceAttr(resourceName, "service_instance_identifier", "serviceInstanceIdentifier"),
				resource.TestCheckResourceAttr(resourceName, "shown_on_login_page", "false"),
				resource.TestCheckResourceAttr(resourceName, "signature_hash_algorithm", "SHA-1"),
				resource.TestCheckResourceAttr(resourceName, "signing_certificate", "MIIDZjCCAk6gAwIBAgIGAXoTivr8MA0GCSqGSIb3DQEBCwUAMFkxEzARBgoJkiaJk/IsZAEZFgNjb20xFjAUBgoJkiaJk/IsZAEZFgZvcmFjbGUxFTATBgoJkiaJk/IsZAEZFgVjbG91ZDETMBEGA1UEAxMKQ2xvdWQ5Q0EtMjAeFw0yMTA2MTYwNjM4MzJaFw0zMTA2MTYwNjM4MzJaMFsxEzARBgNVBAMTCnNzbERvbWFpbnMxDzANBgNVBAMTBkNsb3VkOTEzMDEGA1UEAwwqaWRjcy1hY2YxM2MzMDYyMTM0NTJiYmQwZjgzZmJkYjM3OWYxZl9rZXlzMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEArC8Sg+4SF3ZKdRjveLNheqFlcXZIDV6N6FjMJEYVLyQLrRg8gAXIRqhpS/Ym9Bj8emyV+P/ZlHxmwPkI8DbJX+gsTg9xXG+DlfDSnNTdEOr90hrmg5D8NN3L0SpLughTnywPRbvk6Iv/21u8vALqQHfJYezBDqS/USpm/EP5+soaOlfGWAUaVnxqLtVC6g8kN3dj4hry3ZcfOZwAriEC2M1TTA8vOd7lkGRwEPjBL/QOXWKJU7LPYArKpuL75+gJteQ/yvDHIjbo7ZGaKGhhiCYcQ0gS1I6Fe3V8Mw3N+m7/TpXwGxyEo7GrR7M/j7jhJto88eINDhBjiQkungCY3QIDAQABozIwMDAPBgNVHQ8BAf8EBQMDB/gAMB0GA1UdDgQWBBQ+PjQeG7CnRrqF9HXop07/756KrzANBgkqhkiG9w0BAQsFAAOCAQEAtQqQhc3tsdTHAlQRDkhnPS3N6x7xBpbjef+HOlta86DA9LSoYHBGnnZiwQHTgRbfUdUNllp7hITKMW942wTTwzY/owPukYoo4UAurOZcNkqFCzYJGqxS5bwQRSQO8iSRGwwooet730tzevcilLWH7gF752unoPBoxgCPNxdxWxxSmHpiaQTUAy00sfba0KQ5SLCkXJHEOXpnQsXhe0qm7f/0Mo51cU9pIChj0Cfc/eEmwbcdPnPOCP4N8je+kqW79b3b90kcyUqMDTJw+4HKK62juXm4SXaePJYzx1Yb1VCSZpwOBjHxudp4WrkRPTXFRX248UtSUI5zVw1QsCmOnw=="),
				resource.TestCheckResourceAttrSet(resourceName, "succinct_id"),
				resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tags.0.key", "key"),
				resource.TestCheckResourceAttr(resourceName, "tags.0.value", "value"),
				resource.TestCheckResourceAttr(resourceName, "type", "SAML"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionsocial_identity_provider.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionsocial_identity_provider.0.access_token_url", "accessTokenUrl"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionsocial_identity_provider.0.account_linking_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionsocial_identity_provider.0.admin_scope.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionsocial_identity_provider.0.authz_url", "authzUrl"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionsocial_identity_provider.0.client_credential_in_payload", "false"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionsocial_identity_provider.0.clock_skew_in_seconds", "10"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionsocial_identity_provider.0.consumer_key", "consumerKey"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionsocial_identity_provider.0.consumer_secret", "consumerSecret"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionsocial_identity_provider.0.discovery_url", "discoveryUrl"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionsocial_identity_provider.0.id_attribute", "idAttribute"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionsocial_identity_provider.0.profile_url", "profileUrl"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionsocial_identity_provider.0.redirect_url", "redirectUrl"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionsocial_identity_provider.0.registration_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionsocial_identity_provider.0.scope.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionsocial_identity_provider.0.service_provider_name", "serviceProviderName"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionsocial_identity_provider.0.status", "created"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionx509identity_provider.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionx509identity_provider.0.cert_match_attribute", "certMatchAttribute"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionx509identity_provider.0.crl_check_on_ocsp_failure_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionx509identity_provider.0.crl_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionx509identity_provider.0.crl_location", "crlLocation"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionx509identity_provider.0.crl_reload_duration", "10"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionx509identity_provider.0.eku_validation_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionx509identity_provider.0.eku_values.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionx509identity_provider.0.ocsp_allow_unknown_response_status", "false"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionx509identity_provider.0.ocsp_enable_signed_response", "false"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionx509identity_provider.0.ocsp_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionx509identity_provider.0.ocsp_responder_url", "ocspResponderURL"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionx509identity_provider.0.ocsp_revalidate_time", "10"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionx509identity_provider.0.ocsp_server_name", "ocspServerName"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionx509identity_provider.0.ocsp_trust_cert_chain.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionx509identity_provider.0.other_cert_match_attribute", "otherCertMatchAttribute"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionx509identity_provider.0.signing_certificate_chain.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionx509identity_provider.0.user_match_attribute", "userMatchAttribute"),
				resource.TestCheckResourceAttr(resourceName, "user_mapping_method", "NameIDToUserAttribute"),
				resource.TestCheckResourceAttr(resourceName, "user_mapping_store_attribute", "userName"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					idcsEndpoint, err := acctest.FromInstanceState(s, "data.oci_identity_domain.test_domain", "url")
					if err != nil {
						return err
					}

					compositeId := getIdentityDomainsCompositeId(idcsEndpoint, "identityProviders", resId)
					log.Printf("[DEBUG] Composite ID to import: %s", compositeId)
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&compositeId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsIdentityProviderResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_identity_provider", "test_identity_provider", acctest.Optional, acctest.Update, IdentityDomainsIdentityProviderRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "assertion_attribute", "assertionAttribute2"),
				resource.TestCheckResourceAttr(resourceName, "attribute_sets.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "authn_request_binding", "Post"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "encryption_certificate", "MIIDZjCCAk6gAwIBAgIGAXoTivr8MA0GCSqGSIb3DQEBCwUAMFkxEzARBgoJkiaJk/IsZAEZFgNjb20xFjAUBgoJkiaJk/IsZAEZFgZvcmFjbGUxFTATBgoJkiaJk/IsZAEZFgVjbG91ZDETMBEGA1UEAxMKQ2xvdWQ5Q0EtMjAeFw0yMTA2MTYwNjM4MzJaFw0zMTA2MTYwNjM4MzJaMFsxEzARBgNVBAMTCnNzbERvbWFpbnMxDzANBgNVBAMTBkNsb3VkOTEzMDEGA1UEAwwqaWRjcy1hY2YxM2MzMDYyMTM0NTJiYmQwZjgzZmJkYjM3OWYxZl9rZXlzMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEArC8Sg+4SF3ZKdRjveLNheqFlcXZIDV6N6FjMJEYVLyQLrRg8gAXIRqhpS/Ym9Bj8emyV+P/ZlHxmwPkI8DbJX+gsTg9xXG+DlfDSnNTdEOr90hrmg5D8NN3L0SpLughTnywPRbvk6Iv/21u8vALqQHfJYezBDqS/USpm/EP5+soaOlfGWAUaVnxqLtVC6g8kN3dj4hry3ZcfOZwAriEC2M1TTA8vOd7lkGRwEPjBL/QOXWKJU7LPYArKpuL75+gJteQ/yvDHIjbo7ZGaKGhhiCYcQ0gS1I6Fe3V8Mw3N+m7/TpXwGxyEo7GrR7M/j7jhJto88eINDhBjiQkungCY3QIDAQABozIwMDAPBgNVHQ8BAf8EBQMDB/gAMB0GA1UdDgQWBBQ+PjQeG7CnRrqF9HXop07/756KrzANBgkqhkiG9w0BAQsFAAOCAQEAtQqQhc3tsdTHAlQRDkhnPS3N6x7xBpbjef+HOlta86DA9LSoYHBGnnZiwQHTgRbfUdUNllp7hITKMW942wTTwzY/owPukYoo4UAurOZcNkqFCzYJGqxS5bwQRSQO8iSRGwwooet730tzevcilLWH7gF752unoPBoxgCPNxdxWxxSmHpiaQTUAy00sfba0KQ5SLCkXJHEOXpnQsXhe0qm7f/0Mo51cU9pIChj0Cfc/eEmwbcdPnPOCP4N8je+kqW79b3b90kcyUqMDTJw+4HKK62juXm4SXaePJYzx1Yb1VCSZpwOBjHxudp4WrkRPTXFRX248UtSUI5zVw1QsCmOnw=="),
				resource.TestCheckResourceAttrSet(resourceName, "external_id"),
				resource.TestCheckResourceAttr(resourceName, "icon_url", "https://something.com/iconUrl2.png"),
				resource.TestCheckResourceAttr(resourceName, "idcs_created_by.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttr(resourceName, "idp_sso_url", "https://idpSsoUrl2.com"),
				resource.TestCheckResourceAttr(resourceName, "include_signing_cert_in_signature", "true"),
				resource.TestCheckResourceAttr(resourceName, "jit_user_prov_assigned_groups.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "jit_user_prov_assigned_groups.0.value"),
				resource.TestCheckResourceAttr(resourceName, "jit_user_prov_attribute_update_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "jit_user_prov_create_user_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "jit_user_prov_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "jit_user_prov_group_assertion_attribute_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "jit_user_prov_group_assignment_method", "Merge"),
				resource.TestCheckResourceAttr(resourceName, "jit_user_prov_group_mapping_mode", "explicit"),
				resource.TestCheckResourceAttr(resourceName, "jit_user_prov_group_mappings.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "jit_user_prov_group_mappings.0.idp_group", "idpGroup2"),
				resource.TestCheckResourceAttrSet(resourceName, "jit_user_prov_group_mappings.0.ref"),
				resource.TestCheckResourceAttrSet(resourceName, "jit_user_prov_group_mappings.0.value"),
				resource.TestCheckResourceAttr(resourceName, "jit_user_prov_group_saml_attribute_name", "jitUserProvGroupSAMLAttributeName2"),
				resource.TestCheckResourceAttr(resourceName, "jit_user_prov_group_static_list_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "jit_user_prov_ignore_error_on_absent_groups", "true"),
				resource.TestCheckResourceAttr(resourceName, "logout_binding", "Post"),
				resource.TestCheckResourceAttr(resourceName, "logout_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "logout_request_url", "https://logoutRequestUrl2.com"),
				resource.TestCheckResourceAttr(resourceName, "logout_response_url", "https://logoutResponseUrl2.com"),
				resource.TestCheckResourceAttr(resourceName, "metadata", "<md:EntityDescriptor xmlns:md=\"urn:oasis:names:tc:SAML:2.0:metadata\" xmlns:dsig=\"http://www.w3.org/2000/09/xmldsig#\" xmlns:enc=\"http://www.w3.org/2001/04/xmlenc#\" xmlns:mdattr=\"urn:oasis:names:tc:SAML:metadata:attribute\" xmlns:query=\"urn:oasis:names:tc:SAML:metadata:ext:query\" xmlns:saml=\"urn:oasis:names:tc:SAML:2.0:assertion\" xmlns:x500=\"urn:oasis:names:tc:SAML:2.0:profiles:attribute:X500\" xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" ID=\"id-zzU36agM7bKRB32xe6Ronm131S0-\" cacheDuration=\"P3633DT0H0M0S\" entityID=\"https://idcs-acf13c306213452bbd0f83fbdb379f1f.unstable-alpha.identity.oci.oracleiaas.com:443/fed\" validUntil=\"2031-06-16T06:38:32Z\"><dsig:Signature><dsig:SignedInfo><dsig:CanonicalizationMethod Algorithm=\"http://www.w3.org/2001/10/xml-exc-c14n#\"/><dsig:SignatureMethod Algorithm=\"http://www.w3.org/2001/04/xmldsig-more#rsa-sha256\"/><dsig:Reference URI=\"#id-zzU36agM7bKRB32xe6Ronm131S0-\"><dsig:Transforms><dsig:Transform Algorithm=\"http://www.w3.org/2000/09/xmldsig#enveloped-signature\"/><dsig:Transform Algorithm=\"http://www.w3.org/2001/10/xml-exc-c14n#\"/></dsig:Transforms><dsig:DigestMethod Algorithm=\"http://www.w3.org/2001/04/xmlenc#sha256\"/><dsig:DigestValue>NZnYsjLx3UbuL43iu3jo0mJUg/Rv9DTPNB5IQPRaD6g=</dsig:DigestValue></dsig:Reference></dsig:SignedInfo><dsig:SignatureValue>KRIgTD7//x/uT73veS0iGcWWw8uprjd+MtREu3vlbFTk0BNgkeSOYItx2LDQhnHP\nO0zsTmtOHlVIsDXQL3KysHwzYndIuMJtETqEC6NpMw3ZF108IK0eT+o/2xC9u13/\nGq10z/MagGvco1mM/RIzX5e2omGyZcKARiDoeNPwg2znmV0WcifntVqn4Y0rnWM7\no0M5HFHZQEgICdTJbC5d6DwLgfnI4ck505fHNRYLsRqj9IGLukKx9kocSG1xzCye\nHlffU4CDyEA7dptEUH59dZmY0Xy35/aepNc7W6IovWsJ2Otr+qDUp207ZCKuISF0\nMEX5hX5VJzVlHDwxkEcYCA==</dsig:SignatureValue><dsig:KeyInfo><dsig:X509Data><dsig:X509Certificate>MIIDZjCCAk6gAwIBAgIGAXoTivr8MA0GCSqGSIb3DQEBCwUAMFkxEzARBgoJkiaJ\nk/IsZAEZFgNjb20xFjAUBgoJkiaJk/IsZAEZFgZvcmFjbGUxFTATBgoJkiaJk/Is\nZAEZFgVjbG91ZDETMBEGA1UEAxMKQ2xvdWQ5Q0EtMjAeFw0yMTA2MTYwNjM4MzJa\nFw0zMTA2MTYwNjM4MzJaMFsxEzARBgNVBAMTCnNzbERvbWFpbnMxDzANBgNVBAMT\nBkNsb3VkOTEzMDEGA1UEAwwqaWRjcy1hY2YxM2MzMDYyMTM0NTJiYmQwZjgzZmJk\nYjM3OWYxZl9rZXlzMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEArC8S\ng+4SF3ZKdRjveLNheqFlcXZIDV6N6FjMJEYVLyQLrRg8gAXIRqhpS/Ym9Bj8emyV\n+P/ZlHxmwPkI8DbJX+gsTg9xXG+DlfDSnNTdEOr90hrmg5D8NN3L0SpLughTnywP\nRbvk6Iv/21u8vALqQHfJYezBDqS/USpm/EP5+soaOlfGWAUaVnxqLtVC6g8kN3dj\n4hry3ZcfOZwAriEC2M1TTA8vOd7lkGRwEPjBL/QOXWKJU7LPYArKpuL75+gJteQ/\nyvDHIjbo7ZGaKGhhiCYcQ0gS1I6Fe3V8Mw3N+m7/TpXwGxyEo7GrR7M/j7jhJto8\n8eINDhBjiQkungCY3QIDAQABozIwMDAPBgNVHQ8BAf8EBQMDB/gAMB0GA1UdDgQW\nBBQ+PjQeG7CnRrqF9HXop07/756KrzANBgkqhkiG9w0BAQsFAAOCAQEAtQqQhc3t\nsdTHAlQRDkhnPS3N6x7xBpbjef+HOlta86DA9LSoYHBGnnZiwQHTgRbfUdUNllp7\nhITKMW942wTTwzY/owPukYoo4UAurOZcNkqFCzYJGqxS5bwQRSQO8iSRGwwooet7\n30tzevcilLWH7gF752unoPBoxgCPNxdxWxxSmHpiaQTUAy00sfba0KQ5SLCkXJHE\nOXpnQsXhe0qm7f/0Mo51cU9pIChj0Cfc/eEmwbcdPnPOCP4N8je+kqW79b3b90kc\nyUqMDTJw+4HKK62juXm4SXaePJYzx1Yb1VCSZpwOBjHxudp4WrkRPTXFRX248UtS\nUI5zVw1QsCmOnw==</dsig:X509Certificate></dsig:X509Data></dsig:KeyInfo></dsig:Signature><md:IDPSSODescriptor WantAuthnRequestsSigned=\"false\" protocolSupportEnumeration=\"urn:oasis:names:tc:SAML:2.0:protocol\"><md:KeyDescriptor use=\"signing\"><dsig:KeyInfo><dsig:X509Data><dsig:X509Certificate>MIIDZjCCAk6gAwIBAgIGAXoTivr8MA0GCSqGSIb3DQEBCwUAMFkxEzARBgoJkiaJ\nk/IsZAEZFgNjb20xFjAUBgoJkiaJk/IsZAEZFgZvcmFjbGUxFTATBgoJkiaJk/Is\nZAEZFgVjbG91ZDETMBEGA1UEAxMKQ2xvdWQ5Q0EtMjAeFw0yMTA2MTYwNjM4MzJa\nFw0zMTA2MTYwNjM4MzJaMFsxEzARBgNVBAMTCnNzbERvbWFpbnMxDzANBgNVBAMT\nBkNsb3VkOTEzMDEGA1UEAwwqaWRjcy1hY2YxM2MzMDYyMTM0NTJiYmQwZjgzZmJk\nYjM3OWYxZl9rZXlzMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEArC8S\ng+4SF3ZKdRjveLNheqFlcXZIDV6N6FjMJEYVLyQLrRg8gAXIRqhpS/Ym9Bj8emyV\n+P/ZlHxmwPkI8DbJX+gsTg9xXG+DlfDSnNTdEOr90hrmg5D8NN3L0SpLughTnywP\nRbvk6Iv/21u8vALqQHfJYezBDqS/USpm/EP5+soaOlfGWAUaVnxqLtVC6g8kN3dj\n4hry3ZcfOZwAriEC2M1TTA8vOd7lkGRwEPjBL/QOXWKJU7LPYArKpuL75+gJteQ/\nyvDHIjbo7ZGaKGhhiCYcQ0gS1I6Fe3V8Mw3N+m7/TpXwGxyEo7GrR7M/j7jhJto8\n8eINDhBjiQkungCY3QIDAQABozIwMDAPBgNVHQ8BAf8EBQMDB/gAMB0GA1UdDgQW\nBBQ+PjQeG7CnRrqF9HXop07/756KrzANBgkqhkiG9w0BAQsFAAOCAQEAtQqQhc3t\nsdTHAlQRDkhnPS3N6x7xBpbjef+HOlta86DA9LSoYHBGnnZiwQHTgRbfUdUNllp7\nhITKMW942wTTwzY/owPukYoo4UAurOZcNkqFCzYJGqxS5bwQRSQO8iSRGwwooet7\n30tzevcilLWH7gF752unoPBoxgCPNxdxWxxSmHpiaQTUAy00sfba0KQ5SLCkXJHE\nOXpnQsXhe0qm7f/0Mo51cU9pIChj0Cfc/eEmwbcdPnPOCP4N8je+kqW79b3b90kc\nyUqMDTJw+4HKK62juXm4SXaePJYzx1Yb1VCSZpwOBjHxudp4WrkRPTXFRX248UtS\nUI5zVw1QsCmOnw==\n</dsig:X509Certificate><dsig:X509IssuerSerial><dsig:X509IssuerName>CN=Cloud9CA-2, DC=cloud, DC=oracle, DC=com</dsig:X509IssuerName><dsig:X509SerialNumber>1623825513212</dsig:X509SerialNumber></dsig:X509IssuerSerial><dsig:X509SubjectName>CN=idcs-acf13c306213452bbd0f83fbdb379f1f_keys, CN=Cloud9, CN=sslDomains</dsig:X509SubjectName></dsig:X509Data></dsig:KeyInfo></md:KeyDescriptor><md:KeyDescriptor use=\"encryption\"><dsig:KeyInfo><dsig:X509Data><dsig:X509Certificate>MIIDZjCCAk6gAwIBAgIGAXoTivr8MA0GCSqGSIb3DQEBCwUAMFkxEzARBgoJkiaJ\nk/IsZAEZFgNjb20xFjAUBgoJkiaJk/IsZAEZFgZvcmFjbGUxFTATBgoJkiaJk/Is\nZAEZFgVjbG91ZDETMBEGA1UEAxMKQ2xvdWQ5Q0EtMjAeFw0yMTA2MTYwNjM4MzJa\nFw0zMTA2MTYwNjM4MzJaMFsxEzARBgNVBAMTCnNzbERvbWFpbnMxDzANBgNVBAMT\nBkNsb3VkOTEzMDEGA1UEAwwqaWRjcy1hY2YxM2MzMDYyMTM0NTJiYmQwZjgzZmJk\nYjM3OWYxZl9rZXlzMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEArC8S\ng+4SF3ZKdRjveLNheqFlcXZIDV6N6FjMJEYVLyQLrRg8gAXIRqhpS/Ym9Bj8emyV\n+P/ZlHxmwPkI8DbJX+gsTg9xXG+DlfDSnNTdEOr90hrmg5D8NN3L0SpLughTnywP\nRbvk6Iv/21u8vALqQHfJYezBDqS/USpm/EP5+soaOlfGWAUaVnxqLtVC6g8kN3dj\n4hry3ZcfOZwAriEC2M1TTA8vOd7lkGRwEPjBL/QOXWKJU7LPYArKpuL75+gJteQ/\nyvDHIjbo7ZGaKGhhiCYcQ0gS1I6Fe3V8Mw3N+m7/TpXwGxyEo7GrR7M/j7jhJto8\n8eINDhBjiQkungCY3QIDAQABozIwMDAPBgNVHQ8BAf8EBQMDB/gAMB0GA1UdDgQW\nBBQ+PjQeG7CnRrqF9HXop07/756KrzANBgkqhkiG9w0BAQsFAAOCAQEAtQqQhc3t\nsdTHAlQRDkhnPS3N6x7xBpbjef+HOlta86DA9LSoYHBGnnZiwQHTgRbfUdUNllp7\nhITKMW942wTTwzY/owPukYoo4UAurOZcNkqFCzYJGqxS5bwQRSQO8iSRGwwooet7\n30tzevcilLWH7gF752unoPBoxgCPNxdxWxxSmHpiaQTUAy00sfba0KQ5SLCkXJHE\nOXpnQsXhe0qm7f/0Mo51cU9pIChj0Cfc/eEmwbcdPnPOCP4N8je+kqW79b3b90kc\nyUqMDTJw+4HKK62juXm4SXaePJYzx1Yb1VCSZpwOBjHxudp4WrkRPTXFRX248UtS\nUI5zVw1QsCmOnw==\n</dsig:X509Certificate><dsig:X509IssuerSerial><dsig:X509IssuerName>CN=Cloud9CA-2, DC=cloud, DC=oracle, DC=com</dsig:X509IssuerName><dsig:X509SerialNumber>1623825513212</dsig:X509SerialNumber></dsig:X509IssuerSerial><dsig:X509SubjectName>CN=idcs-acf13c306213452bbd0f83fbdb379f1f_keys, CN=Cloud9, CN=sslDomains</dsig:X509SubjectName></dsig:X509Data></dsig:KeyInfo><md:EncryptionMethod Algorithm=\"http://www.w3.org/2001/04/xmlenc#rsa-1_5\"/><md:EncryptionMethod Algorithm=\"http://www.w3.org/2001/04/xmlenc#aes128-cbc\"/><md:EncryptionMethod Algorithm=\"http://www.w3.org/2001/04/xmlenc#aes192-cbc\"/><md:EncryptionMethod Algorithm=\"http://www.w3.org/2001/04/xmlenc#aes256-cbc\"/><md:EncryptionMethod Algorithm=\"http://www.w3.org/2001/04/xmlenc#tripledes-cbc\"/></md:KeyDescriptor><md:SingleLogoutService Binding=\"urn:oasis:names:tc:SAML:2.0:bindings:HTTP-Redirect\" Location=\"https://idcs-acf13c306213452bbd0f83fbdb379f1f.unstable-alpha.identity.oci.oracleiaas.com/fed/v1/idp/slo\" ResponseLocation=\"https://idcs-acf13c306213452bbd0f83fbdb379f1f.unstable-alpha.identity.oci.oracleiaas.com/fed/v1/idp/slo\"/><md:SingleLogoutService Binding=\"urn:oasis:names:tc:SAML:2.0:bindings:HTTP-POST\" Location=\"https://idcs-acf13c306213452bbd0f83fbdb379f1f.unstable-alpha.identity.oci.oracleiaas.com/fed/v1/idp/slo\" ResponseLocation=\"https://idcs-acf13c306213452bbd0f83fbdb379f1f.unstable-alpha.identity.oci.oracleiaas.com/fed/v1/idp/slo\"/><md:SingleSignOnService Binding=\"urn:oasis:names:tc:SAML:2.0:bindings:HTTP-Redirect\" Location=\"https://idcs-acf13c306213452bbd0f83fbdb379f1f.unstable-alpha.identity.oci.oracleiaas.com/fed/v1/idp/sso\"/><md:SingleSignOnService Binding=\"urn:oasis:names:tc:SAML:2.0:bindings:HTTP-POST\" Location=\"https://idcs-acf13c306213452bbd0f83fbdb379f1f.unstable-alpha.identity.oci.oracleiaas.com/fed/v1/idp/sso\"/></md:IDPSSODescriptor><md:SPSSODescriptor AuthnRequestsSigned=\"true\" WantAssertionsSigned=\"true\" protocolSupportEnumeration=\"urn:oasis:names:tc:SAML:2.0:protocol\"><md:KeyDescriptor use=\"signing\"><dsig:KeyInfo><dsig:X509Data><dsig:X509Certificate>MIIDZjCCAk6gAwIBAgIGAXoTivr8MA0GCSqGSIb3DQEBCwUAMFkxEzARBgoJkiaJ\nk/IsZAEZFgNjb20xFjAUBgoJkiaJk/IsZAEZFgZvcmFjbGUxFTATBgoJkiaJk/Is\nZAEZFgVjbG91ZDETMBEGA1UEAxMKQ2xvdWQ5Q0EtMjAeFw0yMTA2MTYwNjM4MzJa\nFw0zMTA2MTYwNjM4MzJaMFsxEzARBgNVBAMTCnNzbERvbWFpbnMxDzANBgNVBAMT\nBkNsb3VkOTEzMDEGA1UEAwwqaWRjcy1hY2YxM2MzMDYyMTM0NTJiYmQwZjgzZmJk\nYjM3OWYxZl9rZXlzMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEArC8S\ng+4SF3ZKdRjveLNheqFlcXZIDV6N6FjMJEYVLyQLrRg8gAXIRqhpS/Ym9Bj8emyV\n+P/ZlHxmwPkI8DbJX+gsTg9xXG+DlfDSnNTdEOr90hrmg5D8NN3L0SpLughTnywP\nRbvk6Iv/21u8vALqQHfJYezBDqS/USpm/EP5+soaOlfGWAUaVnxqLtVC6g8kN3dj\n4hry3ZcfOZwAriEC2M1TTA8vOd7lkGRwEPjBL/QOXWKJU7LPYArKpuL75+gJteQ/\nyvDHIjbo7ZGaKGhhiCYcQ0gS1I6Fe3V8Mw3N+m7/TpXwGxyEo7GrR7M/j7jhJto8\n8eINDhBjiQkungCY3QIDAQABozIwMDAPBgNVHQ8BAf8EBQMDB/gAMB0GA1UdDgQW\nBBQ+PjQeG7CnRrqF9HXop07/756KrzANBgkqhkiG9w0BAQsFAAOCAQEAtQqQhc3t\nsdTHAlQRDkhnPS3N6x7xBpbjef+HOlta86DA9LSoYHBGnnZiwQHTgRbfUdUNllp7\nhITKMW942wTTwzY/owPukYoo4UAurOZcNkqFCzYJGqxS5bwQRSQO8iSRGwwooet7\n30tzevcilLWH7gF752unoPBoxgCPNxdxWxxSmHpiaQTUAy00sfba0KQ5SLCkXJHE\nOXpnQsXhe0qm7f/0Mo51cU9pIChj0Cfc/eEmwbcdPnPOCP4N8je+kqW79b3b90kc\nyUqMDTJw+4HKK62juXm4SXaePJYzx1Yb1VCSZpwOBjHxudp4WrkRPTXFRX248UtS\nUI5zVw1QsCmOnw==\n</dsig:X509Certificate><dsig:X509IssuerSerial><dsig:X509IssuerName>CN=Cloud9CA-2, DC=cloud, DC=oracle, DC=com</dsig:X509IssuerName><dsig:X509SerialNumber>1623825513212</dsig:X509SerialNumber></dsig:X509IssuerSerial><dsig:X509SubjectName>CN=idcs-acf13c306213452bbd0f83fbdb379f1f_keys, CN=Cloud9, CN=sslDomains</dsig:X509SubjectName></dsig:X509Data></dsig:KeyInfo></md:KeyDescriptor><md:KeyDescriptor use=\"encryption\"><dsig:KeyInfo><dsig:X509Data><dsig:X509Certificate>MIIDZjCCAk6gAwIBAgIGAXoTivr8MA0GCSqGSIb3DQEBCwUAMFkxEzARBgoJkiaJ\nk/IsZAEZFgNjb20xFjAUBgoJkiaJk/IsZAEZFgZvcmFjbGUxFTATBgoJkiaJk/Is\nZAEZFgVjbG91ZDETMBEGA1UEAxMKQ2xvdWQ5Q0EtMjAeFw0yMTA2MTYwNjM4MzJa\nFw0zMTA2MTYwNjM4MzJaMFsxEzARBgNVBAMTCnNzbERvbWFpbnMxDzANBgNVBAMT\nBkNsb3VkOTEzMDEGA1UEAwwqaWRjcy1hY2YxM2MzMDYyMTM0NTJiYmQwZjgzZmJk\nYjM3OWYxZl9rZXlzMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEArC8S\ng+4SF3ZKdRjveLNheqFlcXZIDV6N6FjMJEYVLyQLrRg8gAXIRqhpS/Ym9Bj8emyV\n+P/ZlHxmwPkI8DbJX+gsTg9xXG+DlfDSnNTdEOr90hrmg5D8NN3L0SpLughTnywP\nRbvk6Iv/21u8vALqQHfJYezBDqS/USpm/EP5+soaOlfGWAUaVnxqLtVC6g8kN3dj\n4hry3ZcfOZwAriEC2M1TTA8vOd7lkGRwEPjBL/QOXWKJU7LPYArKpuL75+gJteQ/\nyvDHIjbo7ZGaKGhhiCYcQ0gS1I6Fe3V8Mw3N+m7/TpXwGxyEo7GrR7M/j7jhJto8\n8eINDhBjiQkungCY3QIDAQABozIwMDAPBgNVHQ8BAf8EBQMDB/gAMB0GA1UdDgQW\nBBQ+PjQeG7CnRrqF9HXop07/756KrzANBgkqhkiG9w0BAQsFAAOCAQEAtQqQhc3t\nsdTHAlQRDkhnPS3N6x7xBpbjef+HOlta86DA9LSoYHBGnnZiwQHTgRbfUdUNllp7\nhITKMW942wTTwzY/owPukYoo4UAurOZcNkqFCzYJGqxS5bwQRSQO8iSRGwwooet7\n30tzevcilLWH7gF752unoPBoxgCPNxdxWxxSmHpiaQTUAy00sfba0KQ5SLCkXJHE\nOXpnQsXhe0qm7f/0Mo51cU9pIChj0Cfc/eEmwbcdPnPOCP4N8je+kqW79b3b90kc\nyUqMDTJw+4HKK62juXm4SXaePJYzx1Yb1VCSZpwOBjHxudp4WrkRPTXFRX248UtS\nUI5zVw1QsCmOnw==\n</dsig:X509Certificate><dsig:X509IssuerSerial><dsig:X509IssuerName>CN=Cloud9CA-2, DC=cloud, DC=oracle, DC=com</dsig:X509IssuerName><dsig:X509SerialNumber>1623825513212</dsig:X509SerialNumber></dsig:X509IssuerSerial><dsig:X509SubjectName>CN=idcs-acf13c306213452bbd0f83fbdb379f1f_keys, CN=Cloud9, CN=sslDomains</dsig:X509SubjectName></dsig:X509Data></dsig:KeyInfo><md:EncryptionMethod Algorithm=\"http://www.w3.org/2001/04/xmlenc#rsa-1_5\"/><md:EncryptionMethod Algorithm=\"http://www.w3.org/2001/04/xmlenc#aes128-cbc\"/><md:EncryptionMethod Algorithm=\"http://www.w3.org/2001/04/xmlenc#aes192-cbc\"/><md:EncryptionMethod Algorithm=\"http://www.w3.org/2001/04/xmlenc#aes256-cbc\"/><md:EncryptionMethod Algorithm=\"http://www.w3.org/2001/04/xmlenc#tripledes-cbc\"/></md:KeyDescriptor><md:SingleLogoutService Binding=\"urn:oasis:names:tc:SAML:2.0:bindings:HTTP-Redirect\" Location=\"https://idcs-acf13c306213452bbd0f83fbdb379f1f.unstable-alpha.identity.oci.oracleiaas.com/fed/v1/sp/slo\" ResponseLocation=\"https://idcs-acf13c306213452bbd0f83fbdb379f1f.unstable-alpha.identity.oci.oracleiaas.com/fed/v1/sp/slo\"/><md:SingleLogoutService Binding=\"urn:oasis:names:tc:SAML:2.0:bindings:HTTP-POST\" Location=\"https://idcs-acf13c306213452bbd0f83fbdb379f1f.unstable-alpha.identity.oci.oracleiaas.com/fed/v1/sp/slo\" ResponseLocation=\"https://idcs-acf13c306213452bbd0f83fbdb379f1f.unstable-alpha.identity.oci.oracleiaas.com/fed/v1/sp/slo\"/><md:AssertionConsumerService Binding=\"urn:oasis:names:tc:SAML:2.0:bindings:HTTP-POST\" Location=\"https://idcs-acf13c306213452bbd0f83fbdb379f1f.unstable-alpha.identity.oci.oracleiaas.com/fed/v1/sp/sso\" index=\"1\" isDefault=\"true\"/></md:SPSSODescriptor></md:EntityDescriptor>"),
				resource.TestCheckResourceAttr(resourceName, "name_id_format", "nameIdFormat2"),
				resource.TestCheckResourceAttr(resourceName, "partner_name", "partnerName2"),
				resource.TestCheckResourceAttrSet(resourceName, "partner_provider_id"),
				resource.TestCheckResourceAttr(resourceName, "require_force_authn", "false"),
				resource.TestCheckResourceAttr(resourceName, "requires_encrypted_assertion", "false"),
				resource.TestCheckResourceAttr(resourceName, "saml_ho_krequired", "false"),
				resource.TestMatchResourceAttr(resourceName, "schemas.#", regexp.MustCompile("[1-9]+")),
				resource.TestCheckResourceAttr(resourceName, "service_instance_identifier", "serviceInstanceIdentifier2"),
				resource.TestCheckResourceAttr(resourceName, "shown_on_login_page", "true"),
				resource.TestCheckResourceAttr(resourceName, "signature_hash_algorithm", "SHA-256"),
				resource.TestCheckResourceAttr(resourceName, "signing_certificate", "MIIDZjCCAk6gAwIBAgIGAXoTivr8MA0GCSqGSIb3DQEBCwUAMFkxEzARBgoJkiaJk/IsZAEZFgNjb20xFjAUBgoJkiaJk/IsZAEZFgZvcmFjbGUxFTATBgoJkiaJk/IsZAEZFgVjbG91ZDETMBEGA1UEAxMKQ2xvdWQ5Q0EtMjAeFw0yMTA2MTYwNjM4MzJaFw0zMTA2MTYwNjM4MzJaMFsxEzARBgNVBAMTCnNzbERvbWFpbnMxDzANBgNVBAMTBkNsb3VkOTEzMDEGA1UEAwwqaWRjcy1hY2YxM2MzMDYyMTM0NTJiYmQwZjgzZmJkYjM3OWYxZl9rZXlzMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEArC8Sg+4SF3ZKdRjveLNheqFlcXZIDV6N6FjMJEYVLyQLrRg8gAXIRqhpS/Ym9Bj8emyV+P/ZlHxmwPkI8DbJX+gsTg9xXG+DlfDSnNTdEOr90hrmg5D8NN3L0SpLughTnywPRbvk6Iv/21u8vALqQHfJYezBDqS/USpm/EP5+soaOlfGWAUaVnxqLtVC6g8kN3dj4hry3ZcfOZwAriEC2M1TTA8vOd7lkGRwEPjBL/QOXWKJU7LPYArKpuL75+gJteQ/yvDHIjbo7ZGaKGhhiCYcQ0gS1I6Fe3V8Mw3N+m7/TpXwGxyEo7GrR7M/j7jhJto88eINDhBjiQkungCY3QIDAQABozIwMDAPBgNVHQ8BAf8EBQMDB/gAMB0GA1UdDgQWBBQ+PjQeG7CnRrqF9HXop07/756KrzANBgkqhkiG9w0BAQsFAAOCAQEAtQqQhc3tsdTHAlQRDkhnPS3N6x7xBpbjef+HOlta86DA9LSoYHBGnnZiwQHTgRbfUdUNllp7hITKMW942wTTwzY/owPukYoo4UAurOZcNkqFCzYJGqxS5bwQRSQO8iSRGwwooet730tzevcilLWH7gF752unoPBoxgCPNxdxWxxSmHpiaQTUAy00sfba0KQ5SLCkXJHEOXpnQsXhe0qm7f/0Mo51cU9pIChj0Cfc/eEmwbcdPnPOCP4N8je+kqW79b3b90kcyUqMDTJw+4HKK62juXm4SXaePJYzx1Yb1VCSZpwOBjHxudp4WrkRPTXFRX248UtSUI5zVw1QsCmOnw=="),
				resource.TestCheckResourceAttrSet(resourceName, "succinct_id"),
				resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tags.0.key", "key2"),
				resource.TestCheckResourceAttr(resourceName, "tags.0.value", "value2"),
				resource.TestCheckResourceAttr(resourceName, "type", "SAML"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionsocial_identity_provider.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionsocial_identity_provider.0.access_token_url", "accessTokenUrl2"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionsocial_identity_provider.0.account_linking_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionsocial_identity_provider.0.admin_scope.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionsocial_identity_provider.0.authz_url", "authzUrl2"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionsocial_identity_provider.0.client_credential_in_payload", "true"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionsocial_identity_provider.0.clock_skew_in_seconds", "11"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionsocial_identity_provider.0.consumer_key", "consumerKey2"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionsocial_identity_provider.0.consumer_secret", "consumerSecret2"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionsocial_identity_provider.0.discovery_url", "discoveryUrl2"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionsocial_identity_provider.0.id_attribute", "idAttribute"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionsocial_identity_provider.0.profile_url", "profileUrl2"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionsocial_identity_provider.0.redirect_url", "redirectUrl2"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionsocial_identity_provider.0.registration_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionsocial_identity_provider.0.scope.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionsocial_identity_provider.0.service_provider_name", "serviceProviderName"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionsocial_identity_provider.0.status", "deleted"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionx509identity_provider.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionx509identity_provider.0.cert_match_attribute", "certMatchAttribute2"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionx509identity_provider.0.crl_check_on_ocsp_failure_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionx509identity_provider.0.crl_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionx509identity_provider.0.crl_location", "crlLocation2"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionx509identity_provider.0.crl_reload_duration", "11"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionx509identity_provider.0.eku_validation_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionx509identity_provider.0.eku_values.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionx509identity_provider.0.ocsp_allow_unknown_response_status", "true"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionx509identity_provider.0.ocsp_enable_signed_response", "true"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionx509identity_provider.0.ocsp_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionx509identity_provider.0.ocsp_responder_url", "ocspResponderURL2"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionx509identity_provider.0.ocsp_revalidate_time", "11"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionx509identity_provider.0.ocsp_server_name", "ocspServerName2"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionx509identity_provider.0.ocsp_trust_cert_chain.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionx509identity_provider.0.other_cert_match_attribute", "otherCertMatchAttribute2"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionx509identity_provider.0.signing_certificate_chain.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "urnietfparamsscimschemasoracleidcsextensionx509identity_provider.0.user_match_attribute", "userMatchAttribute2"),
				resource.TestCheckResourceAttr(resourceName, "user_mapping_method", "AssertionAttributeToUserAttribute"),
				resource.TestCheckResourceAttr(resourceName, "user_mapping_store_attribute", "name.givenName"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_domains_identity_providers", "test_identity_providers", acctest.Optional, acctest.Update, IdentityDomainsIdentityDomainsIdentityProviderDataSourceRepresentation) +
				compartmentIdVariableStr + IdentityDomainsIdentityProviderResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_identity_provider", "test_identity_provider", acctest.Optional, acctest.Update, IdentityDomainsIdentityProviderRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "identity_provider_count", "10"),
				resource.TestCheckResourceAttr(datasourceName, "attribute_sets.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "start_index", "1"),

				resource.TestCheckResourceAttr(datasourceName, "identity_providers.#", "1"),
				resource.TestMatchResourceAttr(datasourceName, "identity_providers.0.schemas.#", regexp.MustCompile("[1-9]+")),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_domains_identity_provider", "test_identity_provider", acctest.Required, acctest.Create, IdentityDomainsIdentityDomainsIdentityProviderSingularDataSourceRepresentation) +
				compartmentIdVariableStr + IdentityDomainsIdentityProviderResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "identity_provider_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "assertion_attribute", "assertionAttribute2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "authn_request_binding", "Post"),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "enabled", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "encryption_certificate", "MIIDZjCCAk6gAwIBAgIGAXoTivr8MA0GCSqGSIb3DQEBCwUAMFkxEzARBgoJkiaJk/IsZAEZFgNjb20xFjAUBgoJkiaJk/IsZAEZFgZvcmFjbGUxFTATBgoJkiaJk/IsZAEZFgVjbG91ZDETMBEGA1UEAxMKQ2xvdWQ5Q0EtMjAeFw0yMTA2MTYwNjM4MzJaFw0zMTA2MTYwNjM4MzJaMFsxEzARBgNVBAMTCnNzbERvbWFpbnMxDzANBgNVBAMTBkNsb3VkOTEzMDEGA1UEAwwqaWRjcy1hY2YxM2MzMDYyMTM0NTJiYmQwZjgzZmJkYjM3OWYxZl9rZXlzMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEArC8Sg+4SF3ZKdRjveLNheqFlcXZIDV6N6FjMJEYVLyQLrRg8gAXIRqhpS/Ym9Bj8emyV+P/ZlHxmwPkI8DbJX+gsTg9xXG+DlfDSnNTdEOr90hrmg5D8NN3L0SpLughTnywPRbvk6Iv/21u8vALqQHfJYezBDqS/USpm/EP5+soaOlfGWAUaVnxqLtVC6g8kN3dj4hry3ZcfOZwAriEC2M1TTA8vOd7lkGRwEPjBL/QOXWKJU7LPYArKpuL75+gJteQ/yvDHIjbo7ZGaKGhhiCYcQ0gS1I6Fe3V8Mw3N+m7/TpXwGxyEo7GrR7M/j7jhJto88eINDhBjiQkungCY3QIDAQABozIwMDAPBgNVHQ8BAf8EBQMDB/gAMB0GA1UdDgQWBBQ+PjQeG7CnRrqF9HXop07/756KrzANBgkqhkiG9w0BAQsFAAOCAQEAtQqQhc3tsdTHAlQRDkhnPS3N6x7xBpbjef+HOlta86DA9LSoYHBGnnZiwQHTgRbfUdUNllp7hITKMW942wTTwzY/owPukYoo4UAurOZcNkqFCzYJGqxS5bwQRSQO8iSRGwwooet730tzevcilLWH7gF752unoPBoxgCPNxdxWxxSmHpiaQTUAy00sfba0KQ5SLCkXJHEOXpnQsXhe0qm7f/0Mo51cU9pIChj0Cfc/eEmwbcdPnPOCP4N8je+kqW79b3b90kcyUqMDTJw+4HKK62juXm4SXaePJYzx1Yb1VCSZpwOBjHxudp4WrkRPTXFRX248UtSUI5zVw1QsCmOnw=="),
				resource.TestCheckResourceAttr(singularDatasourceName, "icon_url", "https://something.com/iconUrl2.png"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "idcs_created_by.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "idp_sso_url", "https://idpSsoUrl2.com"),
				resource.TestCheckResourceAttr(singularDatasourceName, "include_signing_cert_in_signature", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "jit_user_prov_assigned_groups.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "jit_user_prov_assigned_groups.0.value"),
				resource.TestCheckResourceAttr(singularDatasourceName, "jit_user_prov_attribute_update_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "jit_user_prov_create_user_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "jit_user_prov_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "jit_user_prov_group_assertion_attribute_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "jit_user_prov_group_assignment_method", "Merge"),
				resource.TestCheckResourceAttr(singularDatasourceName, "jit_user_prov_group_mapping_mode", "explicit"),
				resource.TestCheckResourceAttr(singularDatasourceName, "jit_user_prov_group_mappings.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "jit_user_prov_group_mappings.0.idp_group", "idpGroup2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "jit_user_prov_group_mappings.0.ref"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "jit_user_prov_group_mappings.0.value"),
				resource.TestCheckResourceAttr(singularDatasourceName, "jit_user_prov_group_saml_attribute_name", "jitUserProvGroupSAMLAttributeName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "jit_user_prov_group_static_list_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "jit_user_prov_ignore_error_on_absent_groups", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "logout_binding", "Post"),
				resource.TestCheckResourceAttr(singularDatasourceName, "logout_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "logout_request_url", "https://logoutRequestUrl2.com"),
				resource.TestCheckResourceAttr(singularDatasourceName, "logout_response_url", "https://logoutResponseUrl2.com"),
				resource.TestCheckResourceAttr(singularDatasourceName, "metadata", "<md:EntityDescriptor xmlns:md=\"urn:oasis:names:tc:SAML:2.0:metadata\" xmlns:dsig=\"http://www.w3.org/2000/09/xmldsig#\" xmlns:enc=\"http://www.w3.org/2001/04/xmlenc#\" xmlns:mdattr=\"urn:oasis:names:tc:SAML:metadata:attribute\" xmlns:query=\"urn:oasis:names:tc:SAML:metadata:ext:query\" xmlns:saml=\"urn:oasis:names:tc:SAML:2.0:assertion\" xmlns:x500=\"urn:oasis:names:tc:SAML:2.0:profiles:attribute:X500\" xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" ID=\"id-zzU36agM7bKRB32xe6Ronm131S0-\" cacheDuration=\"P3633DT0H0M0S\" entityID=\"https://idcs-acf13c306213452bbd0f83fbdb379f1f.unstable-alpha.identity.oci.oracleiaas.com:443/fed\" validUntil=\"2031-06-16T06:38:32Z\"><dsig:Signature><dsig:SignedInfo><dsig:CanonicalizationMethod Algorithm=\"http://www.w3.org/2001/10/xml-exc-c14n#\"/><dsig:SignatureMethod Algorithm=\"http://www.w3.org/2001/04/xmldsig-more#rsa-sha256\"/><dsig:Reference URI=\"#id-zzU36agM7bKRB32xe6Ronm131S0-\"><dsig:Transforms><dsig:Transform Algorithm=\"http://www.w3.org/2000/09/xmldsig#enveloped-signature\"/><dsig:Transform Algorithm=\"http://www.w3.org/2001/10/xml-exc-c14n#\"/></dsig:Transforms><dsig:DigestMethod Algorithm=\"http://www.w3.org/2001/04/xmlenc#sha256\"/><dsig:DigestValue>NZnYsjLx3UbuL43iu3jo0mJUg/Rv9DTPNB5IQPRaD6g=</dsig:DigestValue></dsig:Reference></dsig:SignedInfo><dsig:SignatureValue>KRIgTD7//x/uT73veS0iGcWWw8uprjd+MtREu3vlbFTk0BNgkeSOYItx2LDQhnHP\nO0zsTmtOHlVIsDXQL3KysHwzYndIuMJtETqEC6NpMw3ZF108IK0eT+o/2xC9u13/\nGq10z/MagGvco1mM/RIzX5e2omGyZcKARiDoeNPwg2znmV0WcifntVqn4Y0rnWM7\no0M5HFHZQEgICdTJbC5d6DwLgfnI4ck505fHNRYLsRqj9IGLukKx9kocSG1xzCye\nHlffU4CDyEA7dptEUH59dZmY0Xy35/aepNc7W6IovWsJ2Otr+qDUp207ZCKuISF0\nMEX5hX5VJzVlHDwxkEcYCA==</dsig:SignatureValue><dsig:KeyInfo><dsig:X509Data><dsig:X509Certificate>MIIDZjCCAk6gAwIBAgIGAXoTivr8MA0GCSqGSIb3DQEBCwUAMFkxEzARBgoJkiaJ\nk/IsZAEZFgNjb20xFjAUBgoJkiaJk/IsZAEZFgZvcmFjbGUxFTATBgoJkiaJk/Is\nZAEZFgVjbG91ZDETMBEGA1UEAxMKQ2xvdWQ5Q0EtMjAeFw0yMTA2MTYwNjM4MzJa\nFw0zMTA2MTYwNjM4MzJaMFsxEzARBgNVBAMTCnNzbERvbWFpbnMxDzANBgNVBAMT\nBkNsb3VkOTEzMDEGA1UEAwwqaWRjcy1hY2YxM2MzMDYyMTM0NTJiYmQwZjgzZmJk\nYjM3OWYxZl9rZXlzMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEArC8S\ng+4SF3ZKdRjveLNheqFlcXZIDV6N6FjMJEYVLyQLrRg8gAXIRqhpS/Ym9Bj8emyV\n+P/ZlHxmwPkI8DbJX+gsTg9xXG+DlfDSnNTdEOr90hrmg5D8NN3L0SpLughTnywP\nRbvk6Iv/21u8vALqQHfJYezBDqS/USpm/EP5+soaOlfGWAUaVnxqLtVC6g8kN3dj\n4hry3ZcfOZwAriEC2M1TTA8vOd7lkGRwEPjBL/QOXWKJU7LPYArKpuL75+gJteQ/\nyvDHIjbo7ZGaKGhhiCYcQ0gS1I6Fe3V8Mw3N+m7/TpXwGxyEo7GrR7M/j7jhJto8\n8eINDhBjiQkungCY3QIDAQABozIwMDAPBgNVHQ8BAf8EBQMDB/gAMB0GA1UdDgQW\nBBQ+PjQeG7CnRrqF9HXop07/756KrzANBgkqhkiG9w0BAQsFAAOCAQEAtQqQhc3t\nsdTHAlQRDkhnPS3N6x7xBpbjef+HOlta86DA9LSoYHBGnnZiwQHTgRbfUdUNllp7\nhITKMW942wTTwzY/owPukYoo4UAurOZcNkqFCzYJGqxS5bwQRSQO8iSRGwwooet7\n30tzevcilLWH7gF752unoPBoxgCPNxdxWxxSmHpiaQTUAy00sfba0KQ5SLCkXJHE\nOXpnQsXhe0qm7f/0Mo51cU9pIChj0Cfc/eEmwbcdPnPOCP4N8je+kqW79b3b90kc\nyUqMDTJw+4HKK62juXm4SXaePJYzx1Yb1VCSZpwOBjHxudp4WrkRPTXFRX248UtS\nUI5zVw1QsCmOnw==</dsig:X509Certificate></dsig:X509Data></dsig:KeyInfo></dsig:Signature><md:IDPSSODescriptor WantAuthnRequestsSigned=\"false\" protocolSupportEnumeration=\"urn:oasis:names:tc:SAML:2.0:protocol\"><md:KeyDescriptor use=\"signing\"><dsig:KeyInfo><dsig:X509Data><dsig:X509Certificate>MIIDZjCCAk6gAwIBAgIGAXoTivr8MA0GCSqGSIb3DQEBCwUAMFkxEzARBgoJkiaJ\nk/IsZAEZFgNjb20xFjAUBgoJkiaJk/IsZAEZFgZvcmFjbGUxFTATBgoJkiaJk/Is\nZAEZFgVjbG91ZDETMBEGA1UEAxMKQ2xvdWQ5Q0EtMjAeFw0yMTA2MTYwNjM4MzJa\nFw0zMTA2MTYwNjM4MzJaMFsxEzARBgNVBAMTCnNzbERvbWFpbnMxDzANBgNVBAMT\nBkNsb3VkOTEzMDEGA1UEAwwqaWRjcy1hY2YxM2MzMDYyMTM0NTJiYmQwZjgzZmJk\nYjM3OWYxZl9rZXlzMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEArC8S\ng+4SF3ZKdRjveLNheqFlcXZIDV6N6FjMJEYVLyQLrRg8gAXIRqhpS/Ym9Bj8emyV\n+P/ZlHxmwPkI8DbJX+gsTg9xXG+DlfDSnNTdEOr90hrmg5D8NN3L0SpLughTnywP\nRbvk6Iv/21u8vALqQHfJYezBDqS/USpm/EP5+soaOlfGWAUaVnxqLtVC6g8kN3dj\n4hry3ZcfOZwAriEC2M1TTA8vOd7lkGRwEPjBL/QOXWKJU7LPYArKpuL75+gJteQ/\nyvDHIjbo7ZGaKGhhiCYcQ0gS1I6Fe3V8Mw3N+m7/TpXwGxyEo7GrR7M/j7jhJto8\n8eINDhBjiQkungCY3QIDAQABozIwMDAPBgNVHQ8BAf8EBQMDB/gAMB0GA1UdDgQW\nBBQ+PjQeG7CnRrqF9HXop07/756KrzANBgkqhkiG9w0BAQsFAAOCAQEAtQqQhc3t\nsdTHAlQRDkhnPS3N6x7xBpbjef+HOlta86DA9LSoYHBGnnZiwQHTgRbfUdUNllp7\nhITKMW942wTTwzY/owPukYoo4UAurOZcNkqFCzYJGqxS5bwQRSQO8iSRGwwooet7\n30tzevcilLWH7gF752unoPBoxgCPNxdxWxxSmHpiaQTUAy00sfba0KQ5SLCkXJHE\nOXpnQsXhe0qm7f/0Mo51cU9pIChj0Cfc/eEmwbcdPnPOCP4N8je+kqW79b3b90kc\nyUqMDTJw+4HKK62juXm4SXaePJYzx1Yb1VCSZpwOBjHxudp4WrkRPTXFRX248UtS\nUI5zVw1QsCmOnw==\n</dsig:X509Certificate><dsig:X509IssuerSerial><dsig:X509IssuerName>CN=Cloud9CA-2, DC=cloud, DC=oracle, DC=com</dsig:X509IssuerName><dsig:X509SerialNumber>1623825513212</dsig:X509SerialNumber></dsig:X509IssuerSerial><dsig:X509SubjectName>CN=idcs-acf13c306213452bbd0f83fbdb379f1f_keys, CN=Cloud9, CN=sslDomains</dsig:X509SubjectName></dsig:X509Data></dsig:KeyInfo></md:KeyDescriptor><md:KeyDescriptor use=\"encryption\"><dsig:KeyInfo><dsig:X509Data><dsig:X509Certificate>MIIDZjCCAk6gAwIBAgIGAXoTivr8MA0GCSqGSIb3DQEBCwUAMFkxEzARBgoJkiaJ\nk/IsZAEZFgNjb20xFjAUBgoJkiaJk/IsZAEZFgZvcmFjbGUxFTATBgoJkiaJk/Is\nZAEZFgVjbG91ZDETMBEGA1UEAxMKQ2xvdWQ5Q0EtMjAeFw0yMTA2MTYwNjM4MzJa\nFw0zMTA2MTYwNjM4MzJaMFsxEzARBgNVBAMTCnNzbERvbWFpbnMxDzANBgNVBAMT\nBkNsb3VkOTEzMDEGA1UEAwwqaWRjcy1hY2YxM2MzMDYyMTM0NTJiYmQwZjgzZmJk\nYjM3OWYxZl9rZXlzMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEArC8S\ng+4SF3ZKdRjveLNheqFlcXZIDV6N6FjMJEYVLyQLrRg8gAXIRqhpS/Ym9Bj8emyV\n+P/ZlHxmwPkI8DbJX+gsTg9xXG+DlfDSnNTdEOr90hrmg5D8NN3L0SpLughTnywP\nRbvk6Iv/21u8vALqQHfJYezBDqS/USpm/EP5+soaOlfGWAUaVnxqLtVC6g8kN3dj\n4hry3ZcfOZwAriEC2M1TTA8vOd7lkGRwEPjBL/QOXWKJU7LPYArKpuL75+gJteQ/\nyvDHIjbo7ZGaKGhhiCYcQ0gS1I6Fe3V8Mw3N+m7/TpXwGxyEo7GrR7M/j7jhJto8\n8eINDhBjiQkungCY3QIDAQABozIwMDAPBgNVHQ8BAf8EBQMDB/gAMB0GA1UdDgQW\nBBQ+PjQeG7CnRrqF9HXop07/756KrzANBgkqhkiG9w0BAQsFAAOCAQEAtQqQhc3t\nsdTHAlQRDkhnPS3N6x7xBpbjef+HOlta86DA9LSoYHBGnnZiwQHTgRbfUdUNllp7\nhITKMW942wTTwzY/owPukYoo4UAurOZcNkqFCzYJGqxS5bwQRSQO8iSRGwwooet7\n30tzevcilLWH7gF752unoPBoxgCPNxdxWxxSmHpiaQTUAy00sfba0KQ5SLCkXJHE\nOXpnQsXhe0qm7f/0Mo51cU9pIChj0Cfc/eEmwbcdPnPOCP4N8je+kqW79b3b90kc\nyUqMDTJw+4HKK62juXm4SXaePJYzx1Yb1VCSZpwOBjHxudp4WrkRPTXFRX248UtS\nUI5zVw1QsCmOnw==\n</dsig:X509Certificate><dsig:X509IssuerSerial><dsig:X509IssuerName>CN=Cloud9CA-2, DC=cloud, DC=oracle, DC=com</dsig:X509IssuerName><dsig:X509SerialNumber>1623825513212</dsig:X509SerialNumber></dsig:X509IssuerSerial><dsig:X509SubjectName>CN=idcs-acf13c306213452bbd0f83fbdb379f1f_keys, CN=Cloud9, CN=sslDomains</dsig:X509SubjectName></dsig:X509Data></dsig:KeyInfo><md:EncryptionMethod Algorithm=\"http://www.w3.org/2001/04/xmlenc#rsa-1_5\"/><md:EncryptionMethod Algorithm=\"http://www.w3.org/2001/04/xmlenc#aes128-cbc\"/><md:EncryptionMethod Algorithm=\"http://www.w3.org/2001/04/xmlenc#aes192-cbc\"/><md:EncryptionMethod Algorithm=\"http://www.w3.org/2001/04/xmlenc#aes256-cbc\"/><md:EncryptionMethod Algorithm=\"http://www.w3.org/2001/04/xmlenc#tripledes-cbc\"/></md:KeyDescriptor><md:SingleLogoutService Binding=\"urn:oasis:names:tc:SAML:2.0:bindings:HTTP-Redirect\" Location=\"https://idcs-acf13c306213452bbd0f83fbdb379f1f.unstable-alpha.identity.oci.oracleiaas.com/fed/v1/idp/slo\" ResponseLocation=\"https://idcs-acf13c306213452bbd0f83fbdb379f1f.unstable-alpha.identity.oci.oracleiaas.com/fed/v1/idp/slo\"/><md:SingleLogoutService Binding=\"urn:oasis:names:tc:SAML:2.0:bindings:HTTP-POST\" Location=\"https://idcs-acf13c306213452bbd0f83fbdb379f1f.unstable-alpha.identity.oci.oracleiaas.com/fed/v1/idp/slo\" ResponseLocation=\"https://idcs-acf13c306213452bbd0f83fbdb379f1f.unstable-alpha.identity.oci.oracleiaas.com/fed/v1/idp/slo\"/><md:SingleSignOnService Binding=\"urn:oasis:names:tc:SAML:2.0:bindings:HTTP-Redirect\" Location=\"https://idcs-acf13c306213452bbd0f83fbdb379f1f.unstable-alpha.identity.oci.oracleiaas.com/fed/v1/idp/sso\"/><md:SingleSignOnService Binding=\"urn:oasis:names:tc:SAML:2.0:bindings:HTTP-POST\" Location=\"https://idcs-acf13c306213452bbd0f83fbdb379f1f.unstable-alpha.identity.oci.oracleiaas.com/fed/v1/idp/sso\"/></md:IDPSSODescriptor><md:SPSSODescriptor AuthnRequestsSigned=\"true\" WantAssertionsSigned=\"true\" protocolSupportEnumeration=\"urn:oasis:names:tc:SAML:2.0:protocol\"><md:KeyDescriptor use=\"signing\"><dsig:KeyInfo><dsig:X509Data><dsig:X509Certificate>MIIDZjCCAk6gAwIBAgIGAXoTivr8MA0GCSqGSIb3DQEBCwUAMFkxEzARBgoJkiaJ\nk/IsZAEZFgNjb20xFjAUBgoJkiaJk/IsZAEZFgZvcmFjbGUxFTATBgoJkiaJk/Is\nZAEZFgVjbG91ZDETMBEGA1UEAxMKQ2xvdWQ5Q0EtMjAeFw0yMTA2MTYwNjM4MzJa\nFw0zMTA2MTYwNjM4MzJaMFsxEzARBgNVBAMTCnNzbERvbWFpbnMxDzANBgNVBAMT\nBkNsb3VkOTEzMDEGA1UEAwwqaWRjcy1hY2YxM2MzMDYyMTM0NTJiYmQwZjgzZmJk\nYjM3OWYxZl9rZXlzMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEArC8S\ng+4SF3ZKdRjveLNheqFlcXZIDV6N6FjMJEYVLyQLrRg8gAXIRqhpS/Ym9Bj8emyV\n+P/ZlHxmwPkI8DbJX+gsTg9xXG+DlfDSnNTdEOr90hrmg5D8NN3L0SpLughTnywP\nRbvk6Iv/21u8vALqQHfJYezBDqS/USpm/EP5+soaOlfGWAUaVnxqLtVC6g8kN3dj\n4hry3ZcfOZwAriEC2M1TTA8vOd7lkGRwEPjBL/QOXWKJU7LPYArKpuL75+gJteQ/\nyvDHIjbo7ZGaKGhhiCYcQ0gS1I6Fe3V8Mw3N+m7/TpXwGxyEo7GrR7M/j7jhJto8\n8eINDhBjiQkungCY3QIDAQABozIwMDAPBgNVHQ8BAf8EBQMDB/gAMB0GA1UdDgQW\nBBQ+PjQeG7CnRrqF9HXop07/756KrzANBgkqhkiG9w0BAQsFAAOCAQEAtQqQhc3t\nsdTHAlQRDkhnPS3N6x7xBpbjef+HOlta86DA9LSoYHBGnnZiwQHTgRbfUdUNllp7\nhITKMW942wTTwzY/owPukYoo4UAurOZcNkqFCzYJGqxS5bwQRSQO8iSRGwwooet7\n30tzevcilLWH7gF752unoPBoxgCPNxdxWxxSmHpiaQTUAy00sfba0KQ5SLCkXJHE\nOXpnQsXhe0qm7f/0Mo51cU9pIChj0Cfc/eEmwbcdPnPOCP4N8je+kqW79b3b90kc\nyUqMDTJw+4HKK62juXm4SXaePJYzx1Yb1VCSZpwOBjHxudp4WrkRPTXFRX248UtS\nUI5zVw1QsCmOnw==\n</dsig:X509Certificate><dsig:X509IssuerSerial><dsig:X509IssuerName>CN=Cloud9CA-2, DC=cloud, DC=oracle, DC=com</dsig:X509IssuerName><dsig:X509SerialNumber>1623825513212</dsig:X509SerialNumber></dsig:X509IssuerSerial><dsig:X509SubjectName>CN=idcs-acf13c306213452bbd0f83fbdb379f1f_keys, CN=Cloud9, CN=sslDomains</dsig:X509SubjectName></dsig:X509Data></dsig:KeyInfo></md:KeyDescriptor><md:KeyDescriptor use=\"encryption\"><dsig:KeyInfo><dsig:X509Data><dsig:X509Certificate>MIIDZjCCAk6gAwIBAgIGAXoTivr8MA0GCSqGSIb3DQEBCwUAMFkxEzARBgoJkiaJ\nk/IsZAEZFgNjb20xFjAUBgoJkiaJk/IsZAEZFgZvcmFjbGUxFTATBgoJkiaJk/Is\nZAEZFgVjbG91ZDETMBEGA1UEAxMKQ2xvdWQ5Q0EtMjAeFw0yMTA2MTYwNjM4MzJa\nFw0zMTA2MTYwNjM4MzJaMFsxEzARBgNVBAMTCnNzbERvbWFpbnMxDzANBgNVBAMT\nBkNsb3VkOTEzMDEGA1UEAwwqaWRjcy1hY2YxM2MzMDYyMTM0NTJiYmQwZjgzZmJk\nYjM3OWYxZl9rZXlzMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEArC8S\ng+4SF3ZKdRjveLNheqFlcXZIDV6N6FjMJEYVLyQLrRg8gAXIRqhpS/Ym9Bj8emyV\n+P/ZlHxmwPkI8DbJX+gsTg9xXG+DlfDSnNTdEOr90hrmg5D8NN3L0SpLughTnywP\nRbvk6Iv/21u8vALqQHfJYezBDqS/USpm/EP5+soaOlfGWAUaVnxqLtVC6g8kN3dj\n4hry3ZcfOZwAriEC2M1TTA8vOd7lkGRwEPjBL/QOXWKJU7LPYArKpuL75+gJteQ/\nyvDHIjbo7ZGaKGhhiCYcQ0gS1I6Fe3V8Mw3N+m7/TpXwGxyEo7GrR7M/j7jhJto8\n8eINDhBjiQkungCY3QIDAQABozIwMDAPBgNVHQ8BAf8EBQMDB/gAMB0GA1UdDgQW\nBBQ+PjQeG7CnRrqF9HXop07/756KrzANBgkqhkiG9w0BAQsFAAOCAQEAtQqQhc3t\nsdTHAlQRDkhnPS3N6x7xBpbjef+HOlta86DA9LSoYHBGnnZiwQHTgRbfUdUNllp7\nhITKMW942wTTwzY/owPukYoo4UAurOZcNkqFCzYJGqxS5bwQRSQO8iSRGwwooet7\n30tzevcilLWH7gF752unoPBoxgCPNxdxWxxSmHpiaQTUAy00sfba0KQ5SLCkXJHE\nOXpnQsXhe0qm7f/0Mo51cU9pIChj0Cfc/eEmwbcdPnPOCP4N8je+kqW79b3b90kc\nyUqMDTJw+4HKK62juXm4SXaePJYzx1Yb1VCSZpwOBjHxudp4WrkRPTXFRX248UtS\nUI5zVw1QsCmOnw==\n</dsig:X509Certificate><dsig:X509IssuerSerial><dsig:X509IssuerName>CN=Cloud9CA-2, DC=cloud, DC=oracle, DC=com</dsig:X509IssuerName><dsig:X509SerialNumber>1623825513212</dsig:X509SerialNumber></dsig:X509IssuerSerial><dsig:X509SubjectName>CN=idcs-acf13c306213452bbd0f83fbdb379f1f_keys, CN=Cloud9, CN=sslDomains</dsig:X509SubjectName></dsig:X509Data></dsig:KeyInfo><md:EncryptionMethod Algorithm=\"http://www.w3.org/2001/04/xmlenc#rsa-1_5\"/><md:EncryptionMethod Algorithm=\"http://www.w3.org/2001/04/xmlenc#aes128-cbc\"/><md:EncryptionMethod Algorithm=\"http://www.w3.org/2001/04/xmlenc#aes192-cbc\"/><md:EncryptionMethod Algorithm=\"http://www.w3.org/2001/04/xmlenc#aes256-cbc\"/><md:EncryptionMethod Algorithm=\"http://www.w3.org/2001/04/xmlenc#tripledes-cbc\"/></md:KeyDescriptor><md:SingleLogoutService Binding=\"urn:oasis:names:tc:SAML:2.0:bindings:HTTP-Redirect\" Location=\"https://idcs-acf13c306213452bbd0f83fbdb379f1f.unstable-alpha.identity.oci.oracleiaas.com/fed/v1/sp/slo\" ResponseLocation=\"https://idcs-acf13c306213452bbd0f83fbdb379f1f.unstable-alpha.identity.oci.oracleiaas.com/fed/v1/sp/slo\"/><md:SingleLogoutService Binding=\"urn:oasis:names:tc:SAML:2.0:bindings:HTTP-POST\" Location=\"https://idcs-acf13c306213452bbd0f83fbdb379f1f.unstable-alpha.identity.oci.oracleiaas.com/fed/v1/sp/slo\" ResponseLocation=\"https://idcs-acf13c306213452bbd0f83fbdb379f1f.unstable-alpha.identity.oci.oracleiaas.com/fed/v1/sp/slo\"/><md:AssertionConsumerService Binding=\"urn:oasis:names:tc:SAML:2.0:bindings:HTTP-POST\" Location=\"https://idcs-acf13c306213452bbd0f83fbdb379f1f.unstable-alpha.identity.oci.oracleiaas.com/fed/v1/sp/sso\" index=\"1\" isDefault=\"true\"/></md:SPSSODescriptor></md:EntityDescriptor>"),
				resource.TestCheckResourceAttr(singularDatasourceName, "name_id_format", "nameIdFormat2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "partner_name", "partnerName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "require_force_authn", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "requires_encrypted_assertion", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "saml_ho_krequired", "false"),
				resource.TestMatchResourceAttr(singularDatasourceName, "schemas.#", regexp.MustCompile("[1-9]+")),
				resource.TestCheckResourceAttr(singularDatasourceName, "shown_on_login_page", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "signature_hash_algorithm", "SHA-256"),
				resource.TestCheckResourceAttr(singularDatasourceName, "signing_certificate", "MIIDZjCCAk6gAwIBAgIGAXoTivr8MA0GCSqGSIb3DQEBCwUAMFkxEzARBgoJkiaJk/IsZAEZFgNjb20xFjAUBgoJkiaJk/IsZAEZFgZvcmFjbGUxFTATBgoJkiaJk/IsZAEZFgVjbG91ZDETMBEGA1UEAxMKQ2xvdWQ5Q0EtMjAeFw0yMTA2MTYwNjM4MzJaFw0zMTA2MTYwNjM4MzJaMFsxEzARBgNVBAMTCnNzbERvbWFpbnMxDzANBgNVBAMTBkNsb3VkOTEzMDEGA1UEAwwqaWRjcy1hY2YxM2MzMDYyMTM0NTJiYmQwZjgzZmJkYjM3OWYxZl9rZXlzMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEArC8Sg+4SF3ZKdRjveLNheqFlcXZIDV6N6FjMJEYVLyQLrRg8gAXIRqhpS/Ym9Bj8emyV+P/ZlHxmwPkI8DbJX+gsTg9xXG+DlfDSnNTdEOr90hrmg5D8NN3L0SpLughTnywPRbvk6Iv/21u8vALqQHfJYezBDqS/USpm/EP5+soaOlfGWAUaVnxqLtVC6g8kN3dj4hry3ZcfOZwAriEC2M1TTA8vOd7lkGRwEPjBL/QOXWKJU7LPYArKpuL75+gJteQ/yvDHIjbo7ZGaKGhhiCYcQ0gS1I6Fe3V8Mw3N+m7/TpXwGxyEo7GrR7M/j7jhJto88eINDhBjiQkungCY3QIDAQABozIwMDAPBgNVHQ8BAf8EBQMDB/gAMB0GA1UdDgQWBBQ+PjQeG7CnRrqF9HXop07/756KrzANBgkqhkiG9w0BAQsFAAOCAQEAtQqQhc3tsdTHAlQRDkhnPS3N6x7xBpbjef+HOlta86DA9LSoYHBGnnZiwQHTgRbfUdUNllp7hITKMW942wTTwzY/owPukYoo4UAurOZcNkqFCzYJGqxS5bwQRSQO8iSRGwwooet730tzevcilLWH7gF752unoPBoxgCPNxdxWxxSmHpiaQTUAy00sfba0KQ5SLCkXJHEOXpnQsXhe0qm7f/0Mo51cU9pIChj0Cfc/eEmwbcdPnPOCP4N8je+kqW79b3b90kcyUqMDTJw+4HKK62juXm4SXaePJYzx1Yb1VCSZpwOBjHxudp4WrkRPTXFRX248UtSUI5zVw1QsCmOnw=="),
				resource.TestCheckResourceAttr(singularDatasourceName, "type", "SAML"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionsocial_identity_provider.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionsocial_identity_provider.0.access_token_url", "accessTokenUrl2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionsocial_identity_provider.0.account_linking_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionsocial_identity_provider.0.admin_scope.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionsocial_identity_provider.0.authz_url", "authzUrl2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionsocial_identity_provider.0.client_credential_in_payload", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionsocial_identity_provider.0.clock_skew_in_seconds", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionsocial_identity_provider.0.consumer_key", "consumerKey2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionsocial_identity_provider.0.consumer_secret", "consumerSecret2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionsocial_identity_provider.0.discovery_url", "discoveryUrl2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionsocial_identity_provider.0.id_attribute", "idAttribute"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionsocial_identity_provider.0.profile_url", "profileUrl2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionsocial_identity_provider.0.redirect_url", "redirectUrl2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionsocial_identity_provider.0.registration_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionsocial_identity_provider.0.scope.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionsocial_identity_provider.0.service_provider_name", "serviceProviderName"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionsocial_identity_provider.0.status", "deleted"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionx509identity_provider.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionx509identity_provider.0.cert_match_attribute", "certMatchAttribute2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionx509identity_provider.0.crl_check_on_ocsp_failure_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionx509identity_provider.0.crl_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionx509identity_provider.0.crl_location", "crlLocation2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionx509identity_provider.0.crl_reload_duration", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionx509identity_provider.0.eku_validation_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionx509identity_provider.0.eku_values.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionx509identity_provider.0.ocsp_allow_unknown_response_status", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionx509identity_provider.0.ocsp_enable_signed_response", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionx509identity_provider.0.ocsp_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionx509identity_provider.0.ocsp_responder_url", "ocspResponderURL2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionx509identity_provider.0.ocsp_revalidate_time", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionx509identity_provider.0.ocsp_server_name", "ocspServerName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionx509identity_provider.0.ocsp_trust_cert_chain.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionx509identity_provider.0.other_cert_match_attribute", "otherCertMatchAttribute2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionx509identity_provider.0.signing_certificate_chain.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urnietfparamsscimschemasoracleidcsextensionx509identity_provider.0.user_match_attribute", "userMatchAttribute2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "user_mapping_method", "AssertionAttributeToUserAttribute"),
				resource.TestCheckResourceAttr(singularDatasourceName, "user_mapping_store_attribute", "name.givenName"),
			),
		},
		// verify resource import
		{
			Config:            config + IdentityDomainsIdentityProviderRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateIdFunc: getIdentityDomainsImportIdFn("oci_identity_domains_identity_provider", "identityProviders"),
			ImportStateVerifyIgnore: []string{
				"attribute_sets",
				"attributes",
				"authorization",
				"idcs_endpoint",
				"resource_type_schema_version",
				"tags",
				"service_instance_identifier",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckIdentityDomainsIdentityProviderDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).IdentityDomainsClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_identity_domains_identity_provider" {
			noResourceFound = false
			request := oci_identity_domains.GetIdentityProviderRequest{}

			if value, ok := rs.Primary.Attributes["authorization"]; ok {
				request.Authorization = &value
			}

			if value, ok := rs.Primary.Attributes["idcs_endpoint"]; ok {
				client.Host = value
			}

			tmp := rs.Primary.ID
			request.IdentityProviderId = &tmp

			if value, ok := rs.Primary.Attributes["resource_type_schema_version"]; ok {
				request.ResourceTypeSchemaVersion = &value
			}

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "identity_domains")

			_, err := client.GetIdentityProvider(context.Background(), request)

			if err == nil {
				return fmt.Errorf("resource still exists")
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
	if !acctest.InSweeperExcludeList("IdentityDomainsIdentityProvider") {
		resource.AddTestSweepers("IdentityDomainsIdentityProvider", &resource.Sweeper{
			Name:         "IdentityDomainsIdentityProvider",
			Dependencies: acctest.DependencyGraph["identityProvider"],
			F:            sweepIdentityDomainsIdentityProviderResource,
		})
	}
}

func sweepIdentityDomainsIdentityProviderResource(compartment string) error {
	identityDomainsClient := acctest.GetTestClients(&schema.ResourceData{}).IdentityDomainsClient()
	identityProviderIds, err := getIdentityDomainsIdentityProviderIds(compartment)
	if err != nil {
		return err
	}
	for _, identityProviderId := range identityProviderIds {
		if ok := acctest.SweeperDefaultResourceId[identityProviderId]; !ok {
			deleteIdentityProviderRequest := oci_identity_domains.DeleteIdentityProviderRequest{}

			deleteIdentityProviderRequest.IdentityProviderId = &identityProviderId

			deleteIdentityProviderRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "identity_domains")
			_, error := identityDomainsClient.DeleteIdentityProvider(context.Background(), deleteIdentityProviderRequest)
			if error != nil {
				fmt.Printf("Error deleting IdentityProvider %s %s, It is possible that the resource is already deleted. Please verify manually \n", identityProviderId, error)
				continue
			}
		}
	}
	return nil
}

func getIdentityDomainsIdentityProviderIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "IdentityProviderId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	identityDomainsClient := acctest.GetTestClients(&schema.ResourceData{}).IdentityDomainsClient()

	listIdentityProvidersRequest := oci_identity_domains.ListIdentityProvidersRequest{}
	listIdentityProvidersResponse, err := identityDomainsClient.ListIdentityProviders(context.Background(), listIdentityProvidersRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting IdentityProvider list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, identityProvider := range listIdentityProvidersResponse.Resources {
		id := *identityProvider.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "IdentityProviderId", id)
	}
	return resourceIds, nil
}
