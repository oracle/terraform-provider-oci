// Copyright (c) 2017, 2026, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_opensearch "github.com/oracle/oci-go-sdk/v65/opensearch"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

// Public, self-signed test CA, valid from 2020 to 2120. No private key is needed:
// this test verifies trust configuration, not connectivity to a remote source.
const opensearchRemoteReindexTestCaPEM = `-----BEGIN CERTIFICATE-----
MIIC9TCCAd2gAwIBAgIUV9u9AT/oZKnWMdo7BtovLfAD1RcwDQYJKoZIhvcNAQEL
BQAwITEfMB0GA1UEAwwWUmVtb3RlIHJlaW5kZXggdGVzdCBDQTAgFw0yMDAxMDEw
MDAwMDBaGA8yMTIwMDEwMTAwMDAwMFowITEfMB0GA1UEAwwWUmVtb3RlIHJlaW5k
ZXggdGVzdCBDQTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBALIza97k
ilRd8+HJGX7SNE1+i4v7wyHzROYnvedUOD1+rS80mhYNA59GPareDH/8GV/L90KF
79PAyctIaMPvy47t4g9gHGo26r0nmb9rZ+PLMAjPt6og2YFZK5XHqbL1xdy1lWil
caMbgtqJo7AKQMR55ETcu/U1p5aU7PQW1qAWKct4kZIRdiNa+vfXVppxrJiIJmHE
uaG5vKrN4wf46JuyrfILO9wzy46nQguujyfhfU/0h3C//AaChAk6EXb/HdBpHCEB
FzwRT/HmeHunwTBrIwMl5mvcBkZSk1DNf0FwBwaxCTY5sQUO1Ex5KGwzc6dQiW15
Ph5HSXbGvqZb0PMCAwEAAaMjMCEwDwYDVR0TAQH/BAUwAwEB/zAOBgNVHQ8BAf8E
BAMCAYYwDQYJKoZIhvcNAQELBQADggEBAD4g7LO8cebtXvEClouqZnr0vIpr/2il
aRmd+B+8uEipsJYlZL6yH+BrbqQ5by8f6dSH+hnlK9V7gG9g6jlaLKSiVkuzNJO9
QoOVqjs58J2pwKB27Tgw4kL4x0XGITb/5AW9rU98MynWOUXjGPkL8B76MP0YJCSH
tVOpl39WR6K79/n7pVPwH9jSo7IV+2pePbPpNqGACi1nU0HBUdO20ir0WU9CnW8p
25ulCQnSDY+NiwWiRcNkmPdmnHYHeZowIwpHtgrjtmgeLdWwmy/8JA7Xz4zkmAOt
YfeQ9sBryHXXzM9alVwUQhKy1Xzpb+kRNa5hm99YGqG1dMQrPJh81y0=
-----END CERTIFICATE-----
`

// A different public CA makes replacement exercise a real trust change, not
// just a second bundle OCID containing the same certificate.
const opensearchRemoteReindexReplacementCaPEM = `-----BEGIN CERTIFICATE-----
MIIDATCCAemgAwIBAgIUcLUA7iAUADNUik1fx+C16vIepTQwDQYJKoZIhvcNAQEL
BQAwJzElMCMGA1UEAwwcT3RoZXIgcmVtb3RlIHJlaW5kZXggdGVzdCBDQTAgFw0y
MDAxMDEwMDAwMDBaGA8yMTIwMDEwMTAwMDAwMFowJzElMCMGA1UEAwwcT3RoZXIg
cmVtb3RlIHJlaW5kZXggdGVzdCBDQTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCC
AQoCggEBAPYAr2YU0HWENs8Tcqq/tW+v7DZ500nE0KDzpSFiwhJAU5fY6rAckDbl
yyTgVsGp1qNqUyf5no9o9/nfxtS+skXLGRbcmfBELDbOzrPD6dEWKhXZuFQAjYh5
ck4jhddGznZspryfQ7wduB86/YtUOjXO0Oaaz0b5SbezrbIuxxku2nJh2PeKavXH
WWrYeEgK2NmbhsCQujA3V8zfKWUshd0nl0fNEPKOIIxcyRSu2TvPUM3VquK+Q5Kz
QOEu0GzQzVE0pycbSzJzkGN3mHZiRrPxmcqVi/tQSEKUG8nuzX4Qxaec5dkFvfbX
swExAVrRIG2GoYoTHnejisLlk/LdHmECAwEAAaMjMCEwDwYDVR0TAQH/BAUwAwEB
/zAOBgNVHQ8BAf8EBAMCAYYwDQYJKoZIhvcNAQELBQADggEBAGnr8ipk5jFlKe81
DNMA4cpF2ljviRay+4XzD4qRGMPtbmPz6Egn3ygLwS77je4i12dWUyGZh0d1SOf8
qTH7Q2SS+zE1/ESd0k+YdONALMFY2B9OwvXomwZT1AeRSwCQPlHrHSxGXGxB9wm/
9nCgQW1tsfWQIshayNefjBzEwREaKvPRmtxvuIdVFkK6Hpih/3jgFe6F8qV63v2g
fxHDIFAm3Uf+LvijmFFmuvjXjIm2SYo6rI/r/skP9+lPRSmnNYJNI/3QApqgijah
snZVZOXNGGyc+wVYSWgYj+HsdEFAURFfJjsKM1A9+Tm8dH2/vW5SeM7M4f1eAR1B
oTM2LCI=
-----END CERTIFICATE-----
`

const (
	opensearchRemoteReindexClusterResource = "oci_opensearch_opensearch_cluster.test_remote_reindex_cluster"
	opensearchRemoteReindexCaResource      = "oci_certificates_management_ca_bundle.test_remote_reindex_ca_bundle"
	opensearchRemoteReindexOtherCaResource = "oci_certificates_management_ca_bundle.test_remote_reindex_other_ca_bundle"
)

type opensearchRemoteReindexSource struct {
	caBundleResource string
	endpoint         string
	// nil omits the optional field.
	resolutionIP *string
}

type opensearchRemoteReindexConfigOptions struct {
	// No sources omits trust (preserve on update), unless clearSources is explicit.
	sources          []opensearchRemoteReindexSource
	clearSources     bool
	replacementCA    bool
	dataNodeCount    int
	dataNodeOCPUs    int
	dataNodeMemoryGB int
	tagged           bool
}

// issue-routing-tag: opensearch/default
func TestOpensearchOpensearchClusterResource_remoteReindex_basic(t *testing.T) {
	httpreplay.SetScenario("TestOpensearchOpensearchClusterResource_remoteReindex_basic")
	defer httpreplay.SaveScenario()

	const clusterResource = "oci_opensearch_opensearch_cluster.test_remote_reindex_cluster"
	const caBundleResource = "oci_certificates_management_ca_bundle.test_remote_reindex_ca_bundle"
	const sourceAttribute = "remote_reindex_trust_config.0.remote_sources.0."
	// Reserved example hostname/address: no actual remote reindex request is sent.
	const endpoint = "https://remote-reindex.example.com:9200"
	const resolutionIP = "192.0.2.10"

	compartmentID := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	name := "tf-remote-reindex-" + utils.RandomStringOrHttpReplayValue(10, utils.CharsetLowerCaseWithoutDigits, "test")
	config := testAccOpensearchRemoteReindexBasicConfig(compartmentID, name, endpoint, resolutionIP)
	acctest.SaveConfigContent(config, "opensearch", "remoteReindex", t)

	// The cluster's reference to the CA bundle ensures it is destroyed first,
	// allowing MP to release the association before Terraform deletes the bundle.
	checkDestroy := resource.ComposeAggregateTestCheckFunc(
		testAccCheckOpensearchOpensearchClusterDestroy,
		testAccCheckCertificatesManagementCaBundleDestroy,
	)
	acctest.ResourceTest(t, checkDestroy, []resource.TestStep{
		{
			Config: config,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(caBundleResource, "id"),
				resource.TestCheckResourceAttr(caBundleResource, "compartment_id", compartmentID),
				resource.TestCheckResourceAttr(caBundleResource, "name", name),
				resource.TestCheckResourceAttr(caBundleResource, "state", "ACTIVE"),
				resource.TestCheckResourceAttrSet(clusterResource, "id"),
				resource.TestCheckResourceAttr(clusterResource, "compartment_id", compartmentID),
				resource.TestCheckResourceAttr(clusterResource, "display_name", name),
				resource.TestCheckResourceAttr(clusterResource, "state", "ACTIVE"),
				resource.TestCheckResourceAttr(clusterResource, "remote_reindex_trust_config.#", "1"),
				resource.TestCheckResourceAttr(clusterResource, "remote_reindex_trust_config.0.remote_sources.#", "1"),
				resource.TestCheckResourceAttr(clusterResource, sourceAttribute+"endpoint", endpoint),
				resource.TestCheckResourceAttr(clusterResource, sourceAttribute+"endpoint_resolution_ip", resolutionIP),
				resource.TestCheckResourceAttrPair(clusterResource, sourceAttribute+"ca_bundle_id", caBundleResource, "id"),
				testAccCheckOpensearchRemoteReindexTrust(clusterResource, caBundleResource, endpoint, resolutionIP),
			),
		},
	})
}

// issue-routing-tag: opensearch/default
func TestOpensearchOpensearchClusterResource_remoteReindex_lifecycle(t *testing.T) {
	httpreplay.SetScenario("TestOpensearchOpensearchClusterResource_remoteReindex_lifecycle")
	defer httpreplay.SaveScenario()
	compartmentID, name := testAccOpensearchRemoteReindexTestSettings("lifecycle")
	source := opensearchRemoteReindexSource{
		caBundleResource: opensearchRemoteReindexCaResource,
		endpoint:         "https://source-a.example.com:9200",
		resolutionIP:     common.String("192.0.2.10"),
	}
	updated := source
	updated.endpoint = "https://source-b.example.com:9443"
	updated.resolutionIP = common.String("192.0.2.11")
	rotated := updated
	rotated.caBundleResource = opensearchRemoteReindexOtherCaResource
	additional := rotated
	additional.endpoint = "https://source-c.example.com:9200"
	additional.resolutionIP = common.String("192.0.2.12")

	var clusterID string
	var steps []resource.TestStep
	// Each mutation changes only remote_reindex_trust_config. Changing a tag or
	// display name here would hide a missing provider update trigger.
	for _, stage := range []struct {
		name    string
		sources []opensearchRemoteReindexSource
	}{
		{name: "create_without_trust", sources: nil},
		{name: "enable_on_existing_cluster", sources: []opensearchRemoteReindexSource{source}},
		{name: "change_endpoint_and_ip", sources: []opensearchRemoteReindexSource{updated}},
		{name: "replace_ca_bundle", sources: []opensearchRemoteReindexSource{rotated}},
		{name: "add_source_sharing_ca", sources: []opensearchRemoteReindexSource{rotated, additional}},
		{name: "remove_one_source", sources: []opensearchRemoteReindexSource{additional}},
	} {
		config := testAccOpensearchRemoteReindexConfig(compartmentID, name, opensearchRemoteReindexConfigOptions{
			sources: stage.sources, replacementCA: true,
		})
		acctest.SaveConfigContent(config, "opensearch", "remoteReindex_"+stage.name, t)
		steps = append(steps, resource.TestStep{
			Config: config,
			Check: resource.ComposeAggregateTestCheckFunc(
				testAccCheckOpensearchRemoteReindexSameCluster(&clusterID),
				testAccCheckOpensearchRemoteReindexState(opensearchRemoteReindexClusterResource, "", stage.sources),
				testAccCheckOpensearchRemoteReindexSources(opensearchRemoteReindexClusterResource, stage.sources),
			),
		})
	}
	acctest.ResourceTest(t, testAccCheckOpensearchRemoteReindexDestroy, steps)
}

// issue-routing-tag: opensearch/default
func TestOpensearchOpensearchClusterResource_remoteReindex_removal(t *testing.T) {
	httpreplay.SetScenario("TestOpensearchOpensearchClusterResource_remoteReindex_removal")
	defer httpreplay.SaveScenario()
	compartmentID, name := testAccOpensearchRemoteReindexTestSettings("removal")
	sources := []opensearchRemoteReindexSource{
		{caBundleResource: opensearchRemoteReindexCaResource, endpoint: "https://source.example.com:9200", resolutionIP: common.String("192.0.2.10")},
	}
	configured := testAccOpensearchRemoteReindexConfig(compartmentID, name, opensearchRemoteReindexConfigOptions{sources: sources})
	cleared := testAccOpensearchRemoteReindexConfig(compartmentID, name, opensearchRemoteReindexConfigOptions{clearSources: true})
	acctest.SaveConfigContent(cleared, "opensearch", "remoteReindex_removal", t)
	var clusterID string
	configuredChecks := resource.ComposeAggregateTestCheckFunc(
		testAccCheckOpensearchRemoteReindexSameCluster(&clusterID),
		testAccCheckOpensearchRemoteReindexState(opensearchRemoteReindexClusterResource, "", sources),
		testAccCheckOpensearchRemoteReindexSources(opensearchRemoteReindexClusterResource, sources),
	)
	acctest.ResourceTest(t, testAccCheckOpensearchRemoteReindexDestroy, []resource.TestStep{
		{Config: configured, Check: configuredChecks},
		{
			// Keep the parent block and explicitly send remote_sources = [].
			// Omitting the parent block would instead test preservation.
			Config: cleared,
			Check: resource.ComposeAggregateTestCheckFunc(
				testAccCheckOpensearchRemoteReindexSameCluster(&clusterID),
				testAccCheckOpensearchRemoteReindexClearedState(opensearchRemoteReindexClusterResource),
				testAccCheckOpensearchRemoteReindexSources(opensearchRemoteReindexClusterResource, nil),
			),
		},
		// Read-back normalization must not cause a perpetual removal plan.
		{Config: cleared, PlanOnly: true, ExpectNonEmptyPlan: false},
		{Config: configured, Check: configuredChecks},
		{Config: configured, PlanOnly: true, ExpectNonEmptyPlan: false},
	})
}

// issue-routing-tag: opensearch/default
func TestOpensearchOpensearchClusterResource_remoteReindex_readImport(t *testing.T) {
	httpreplay.SetScenario("TestOpensearchOpensearchClusterResource_remoteReindex_readImport")
	defer httpreplay.SaveScenario()
	compartmentID, name := testAccOpensearchRemoteReindexTestSettings("read-import")
	// The API returns sources sorted by normalized endpoint; keep this TypeList
	// in that order. Test shared and distinct bundles, IPv4, IPv6, and no alias.
	sources := []opensearchRemoteReindexSource{
		{caBundleResource: opensearchRemoteReindexCaResource, endpoint: "https://source-a.example.com:9200", resolutionIP: common.String("192.0.2.10")},
		{caBundleResource: opensearchRemoteReindexCaResource, endpoint: "https://source-b.example.com:9200", resolutionIP: common.String("2001:db8::10")},
		{caBundleResource: opensearchRemoteReindexOtherCaResource, endpoint: "https://source-c.example.com:9200"},
	}
	config := testAccOpensearchRemoteReindexConfig(compartmentID, name, opensearchRemoteReindexConfigOptions{
		sources: sources, replacementCA: true,
	})
	readConfig := config + `
data "oci_opensearch_opensearch_cluster" "remote_reindex" {
  opensearch_cluster_id = oci_opensearch_opensearch_cluster.test_remote_reindex_cluster.id
}
data "oci_opensearch_opensearch_clusters" "remote_reindex" {
  compartment_id = var.compartment_id
  id             = oci_opensearch_opensearch_cluster.test_remote_reindex_cluster.id
  state          = "ACTIVE"
}
`
	acctest.SaveConfigContent(readConfig, "opensearch", "remoteReindex_readImport", t)
	const singular = "data.oci_opensearch_opensearch_cluster.remote_reindex"
	const plural = "data.oci_opensearch_opensearch_clusters.remote_reindex"
	var clusterID string
	acctest.ResourceTest(t, testAccCheckOpensearchRemoteReindexDestroy, []resource.TestStep{
		{
			Config: config,
			Check: resource.ComposeAggregateTestCheckFunc(
				testAccCheckOpensearchRemoteReindexSameCluster(&clusterID),
				testAccCheckOpensearchRemoteReindexState(opensearchRemoteReindexClusterResource, "", sources),
				testAccCheckOpensearchRemoteReindexSources(opensearchRemoteReindexClusterResource, sources),
			),
		},
		{
			Config: readConfig,
			Check: resource.ComposeAggregateTestCheckFunc(
				testAccCheckOpensearchRemoteReindexSameCluster(&clusterID),
				resource.TestCheckResourceAttrPair(singular, "id", opensearchRemoteReindexClusterResource, "id"),
				testAccCheckOpensearchRemoteReindexState(singular, "", sources),
				resource.TestCheckResourceAttr(plural, "opensearch_cluster_collection.0.items.#", "1"),
				resource.TestCheckResourceAttrPair(plural, "opensearch_cluster_collection.0.items.0.id", opensearchRemoteReindexClusterResource, "id"),
				testAccCheckOpensearchRemoteReindexState(plural, "opensearch_cluster_collection.0.items.0.", sources),
			),
		},
		{
			Config: readConfig, ResourceName: opensearchRemoteReindexClusterResource,
			ImportState: true, ImportStateVerify: true,
			// Do not ignore remote_reindex_trust_config during import verification.
		},
		{Config: readConfig, PlanOnly: true, ExpectNonEmptyPlan: false},
	})
}

// issue-routing-tag: opensearch/default
func TestOpensearchOpensearchClusterResource_remoteReindex_preservation(t *testing.T) {
	httpreplay.SetScenario("TestOpensearchOpensearchClusterResource_remoteReindex_preservation")
	defer httpreplay.SaveScenario()
	compartmentID, name := testAccOpensearchRemoteReindexTestSettings("preservation")
	sources := []opensearchRemoteReindexSource{
		{caBundleResource: opensearchRemoteReindexCaResource, endpoint: "https://source-a.example.com:9200", resolutionIP: common.String("192.0.2.10")},
	}
	var clusterID string
	var steps []resource.TestStep
	// This deliberately slow test is separately selectable from create/read and
	// validation. Verify preservation after unrelated updates and both resize APIs.
	for _, options := range []opensearchRemoteReindexConfigOptions{
		{sources: sources, dataNodeCount: 1, dataNodeOCPUs: 2, dataNodeMemoryGB: 20},
		{sources: sources, tagged: true, dataNodeCount: 1, dataNodeOCPUs: 2, dataNodeMemoryGB: 20},
		{sources: nil, tagged: true, dataNodeCount: 1, dataNodeOCPUs: 2, dataNodeMemoryGB: 20},
		{sources: nil, tagged: true, dataNodeCount: 2, dataNodeOCPUs: 2, dataNodeMemoryGB: 20},
		{sources: nil, tagged: true, dataNodeCount: 2, dataNodeOCPUs: 4, dataNodeMemoryGB: 40},
	} {
		steps = append(steps, resource.TestStep{
			Config: testAccOpensearchRemoteReindexConfig(compartmentID, name, options),
			Check: resource.ComposeAggregateTestCheckFunc(
				testAccCheckOpensearchRemoteReindexSameCluster(&clusterID),
				testAccCheckOpensearchRemoteReindexState(opensearchRemoteReindexClusterResource, "", sources),
				testAccCheckOpensearchRemoteReindexSources(opensearchRemoteReindexClusterResource, sources),
				resource.TestCheckResourceAttr(opensearchRemoteReindexClusterResource, "data_node_count", strconv.Itoa(options.dataNodeCount)),
				resource.TestCheckResourceAttr(opensearchRemoteReindexClusterResource, "data_node_host_ocpu_count", strconv.Itoa(options.dataNodeOCPUs)),
				resource.TestCheckResourceAttr(opensearchRemoteReindexClusterResource, "data_node_host_memory_gb", strconv.Itoa(options.dataNodeMemoryGB)),
			),
		})
	}
	acctest.ResourceTest(t, testAccCheckOpensearchRemoteReindexDestroy, steps)
}

// issue-routing-tag: opensearch/default
func TestOpensearchOpensearchClusterResource_remoteReindex_validation(t *testing.T) {
	source := opensearchRemoteReindexSource{
		caBundleResource: opensearchRemoteReindexCaResource,
		endpoint:         "https://source.example.com:9200",
		resolutionIP:     common.String("192.0.2.10"),
	}
	nonHTTPS, invalidIP, literalHost := source, source, source
	nonHTTPS.endpoint = "http://source.example.com:9200"
	invalidIP.resolutionIP = common.String("not-an-ip")
	literalHost.endpoint = "https://192.0.2.10:9200"
	duplicate := source
	duplicate.endpoint = "https://SOURCE.example.com:9443"
	tooMany := make([]opensearchRemoteReindexSource, 11)
	for i := range tooMany {
		tooMany[i] = source
		tooMany[i].endpoint = fmt.Sprintf("https://source-%02d.example.com:9200", i)
	}
	for _, tc := range []struct {
		name    string
		sources []opensearchRemoteReindexSource
		message string
		// Overrides one valid reference with deliberately malformed input.
		invalidCaBundle bool
	}{
		{name: "non_https_endpoint", sources: []opensearchRemoteReindexSource{nonHTTPS}, message: "endpoint must be an HTTPS URL"},
		{name: "invalid_resolution_ip", sources: []opensearchRemoteReindexSource{invalidIP}, message: "endpointResolutionIp must be a valid IPv4 or IPv6 address"},
		{name: "ip_endpoint_with_alias", sources: []opensearchRemoteReindexSource{literalHost}, message: "endpoint must use a DNS hostname"},
		{name: "duplicate_hostname_different_case_and_port", sources: []opensearchRemoteReindexSource{source, duplicate}, message: "Duplicate remote source hostname"},
		{name: "too_many_sources", sources: tooMany, message: "At most 10 source configurations"},
		{name: "invalid_ca_bundle_ocid", sources: []opensearchRemoteReindexSource{source}, message: "must be a valid OCI CA bundle OCID", invalidCaBundle: true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			httpreplay.SetScenario("TestOpensearchOpensearchClusterResource_remoteReindex_validation_" + tc.name)
			defer httpreplay.SaveScenario()
			compartmentID, name := testAccOpensearchRemoteReindexTestSettings("validation")
			config := testAccOpensearchRemoteReindexConfig(compartmentID, name, opensearchRemoteReindexConfigOptions{sources: tc.sources})
			if tc.invalidCaBundle {
				config = strings.Replace(config, "ca_bundle_id = "+opensearchRemoteReindexCaResource+".id", `ca_bundle_id = "not-a-ca-bundle-ocid"`, 1)
			}
			acctest.ResourceTest(t, testAccCheckOpensearchRemoteReindexDestroy, []resource.TestStep{
				{
					Config: config,
					// Require the intended API validation error, not authentication,
					// capacity, network, or unrelated Terraform failures.
					ExpectError: regexp.MustCompile(`(?s)400-InvalidParameter.*` + strings.ReplaceAll(regexp.QuoteMeta(tc.message), " ", `\s+`)),
				},
			})
		})
	}
}

func testAccOpensearchRemoteReindexTestSettings(suffix string) (string, string) {
	return utils.GetEnvSettingWithBlankDefault("compartment_ocid"),
		"tf-remote-reindex-" + suffix + "-" + utils.RandomStringOrHttpReplayValue(10, utils.CharsetLowerCaseWithoutDigits, "test")
}

func testAccCheckOpensearchRemoteReindexSameCluster(expectedID *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		id, err := acctest.FromInstanceState(s, opensearchRemoteReindexClusterResource, "id")
		if err != nil {
			return err
		}
		if id == "" {
			return fmt.Errorf("remote-reindex cluster ID is empty")
		}
		if *expectedID == "" {
			*expectedID = id
		} else if *expectedID != id {
			return fmt.Errorf("remote-reindex update replaced cluster %s with %s", *expectedID, id)
		}
		return nil
	}
}

func testAccCheckOpensearchRemoteReindexState(resourceName, prefix string, sources []opensearchRemoteReindexSource) resource.TestCheckFunc {
	trustPrefix := prefix + "remote_reindex_trust_config."
	checks := []resource.TestCheckFunc{resource.TestCheckResourceAttr(resourceName, prefix+"state", "ACTIVE")}
	if len(sources) == 0 {
		checks = append(checks, resource.TestCheckResourceAttr(resourceName, trustPrefix+"#", "0"))
	} else {
		checks = append(checks,
			resource.TestCheckResourceAttr(resourceName, trustPrefix+"#", "1"),
			resource.TestCheckResourceAttr(resourceName, trustPrefix+"0.remote_sources.#", strconv.Itoa(len(sources))),
		)
		for i, source := range sources {
			key := fmt.Sprintf("%s0.remote_sources.%d.", trustPrefix, i)
			checks = append(checks,
				resource.TestCheckResourceAttr(resourceName, key+"endpoint", source.endpoint),
				resource.TestCheckResourceAttrPair(resourceName, key+"ca_bundle_id", source.caBundleResource, "id"),
			)
			if source.resolutionIP != nil {
				checks = append(checks, resource.TestCheckResourceAttr(resourceName, key+"endpoint_resolution_ip", *source.resolutionIP))
			} else {
				checks = append(checks, func(s *terraform.State) error {
					rs, ok := s.RootModule().Resources[resourceName]
					if !ok || rs.Primary == nil {
						return fmt.Errorf("resource %s missing from state", resourceName)
					}
					// SDK state can represent an omitted optional string as absent or empty.
					if ip := rs.Primary.Attributes[key+"endpoint_resolution_ip"]; ip != "" {
						return fmt.Errorf("%s: expected no resolution IP for %s, got %q", resourceName, key, ip)
					}
					return nil
				})
			}
		}
	}
	return resource.ComposeAggregateTestCheckFunc(checks...)
}

func testAccCheckOpensearchRemoteReindexClearedState(resourceName string) resource.TestCheckFunc {
	return resource.ComposeAggregateTestCheckFunc(
		resource.TestCheckResourceAttr(resourceName, "state", "ACTIVE"),
		func(s *terraform.State) error {
			rs, ok := s.RootModule().Resources[resourceName]
			if !ok || rs.Primary == nil {
				return fmt.Errorf("resource %s missing from state", resourceName)
			}
			attrs := rs.Primary.Attributes
			const trustCount = "remote_reindex_trust_config.#"
			const sourcePrefix = "remote_reindex_trust_config.0.remote_sources."
			// The API may represent cleared trust as null or as an empty source
			// collection. Both are cleared, but neither may retain source values.
			switch attrs[trustCount] {
			case "0":
				if count := attrs[sourcePrefix+"#"]; count != "" && count != "0" {
					return fmt.Errorf("%s: absent trust retains %s remote sources", resourceName, count)
				}
			case "1":
				if attrs[sourcePrefix+"#"] != "0" {
					return fmt.Errorf("%s: cleared trust must have zero remote sources", resourceName)
				}
			default:
				return fmt.Errorf("%s: unexpected cleared trust count %q", resourceName, attrs[trustCount])
			}
			for key, value := range attrs {
				if strings.HasPrefix(key, sourcePrefix) && key != sourcePrefix+"#" && value != "" {
					return fmt.Errorf("%s: cleared trust retains stale attribute %s", resourceName, key)
				}
			}
			return nil
		},
	)
}

func testAccCheckOpensearchRemoteReindexDestroy(s *terraform.State) error {
	// Negative-create tests may have only dependencies in state. Still verify
	// any cluster that unexpectedly made it into state was destroyed as well.
	types := map[string]bool{}
	for _, rs := range s.RootModule().Resources {
		if rs.Primary != nil && rs.Primary.ID != "" {
			types[rs.Type] = true
		}
	}
	var checks []resource.TestCheckFunc
	if types["oci_opensearch_opensearch_cluster"] {
		checks = append(checks, testAccCheckOpensearchOpensearchClusterDestroy)
	}
	if types["oci_certificates_management_ca_bundle"] {
		checks = append(checks, testAccCheckCertificatesManagementCaBundleDestroy)
	}
	return resource.ComposeAggregateTestCheckFunc(checks...)(s)
}

// Independently read the API so passing checks cannot rely only on configured
// values retained in Terraform state.
func testAccCheckOpensearchRemoteReindexTrust(clusterResource, caBundleResource, endpoint, resolutionIP string) resource.TestCheckFunc {
	return testAccCheckOpensearchRemoteReindexSources(clusterResource, []opensearchRemoteReindexSource{
		{caBundleResource: caBundleResource, endpoint: endpoint, resolutionIP: common.String(resolutionIP)},
	})
}

func testAccCheckOpensearchRemoteReindexSources(clusterResource string, expected []opensearchRemoteReindexSource) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		clusterID, err := acctest.FromInstanceState(s, clusterResource, "id")
		if err != nil {
			return err
		}
		client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).OpensearchClusterClient()
		response, err := client.GetOpensearchCluster(context.Background(), oci_opensearch.GetOpensearchClusterRequest{
			OpensearchClusterId: common.String(clusterID),
		})
		if err != nil {
			return fmt.Errorf("reading remote-reindex test cluster %s: %w", clusterID, err)
		}
		if response.LifecycleState != oci_opensearch.OpensearchClusterLifecycleStateActive {
			return fmt.Errorf("cluster %s: expected ACTIVE, got %s", clusterID, response.LifecycleState)
		}
		trust := response.RemoteReindexTrustConfig
		var actual []oci_opensearch.RemoteReindexSourceDetails
		if trust != nil {
			actual = trust.RemoteSources
		}
		if len(actual) != len(expected) {
			return fmt.Errorf("cluster %s: expected %d API remote-reindex sources, got %d", clusterID, len(expected), len(actual))
		}
		// Match by endpoint, not API ordering. Terraform's ordered state is checked separately.
		byEndpoint := make(map[string]oci_opensearch.RemoteReindexSourceDetails, len(actual))
		for _, source := range actual {
			if source.Endpoint == nil {
				return fmt.Errorf("cluster %s: API returned a source without an endpoint", clusterID)
			}
			if _, duplicate := byEndpoint[*source.Endpoint]; duplicate {
				return fmt.Errorf("cluster %s: API returned duplicate endpoint %q", clusterID, *source.Endpoint)
			}
			byEndpoint[*source.Endpoint] = source
		}
		for _, want := range expected {
			caBundleID, err := acctest.FromInstanceState(s, want.caBundleResource, "id")
			if err != nil {
				return err
			}
			source, ok := byEndpoint[want.endpoint]
			if !ok {
				return fmt.Errorf("cluster %s: API missing remote-reindex endpoint %q", clusterID, want.endpoint)
			}
			if source.CaBundleId == nil || *source.CaBundleId != caBundleID {
				return fmt.Errorf("cluster %s: API CA bundle ID for %s does not match CA bundle %s", clusterID, want.endpoint, caBundleID)
			}
			if want.resolutionIP == nil {
				if source.EndpointResolutionIp != nil {
					return fmt.Errorf("cluster %s: API resolution IP for %s must be absent, got %q", clusterID, want.endpoint, *source.EndpointResolutionIp)
				}
			} else if source.EndpointResolutionIp == nil || *source.EndpointResolutionIp != *want.resolutionIP {
				return fmt.Errorf("cluster %s: API resolution IP for %s does not match %q", clusterID, want.endpoint, *want.resolutionIP)
			}
		}
		return nil
	}
}

func testAccOpensearchRemoteReindexBasicConfig(compartmentID, name, endpoint, resolutionIP string) string {
	return testAccOpensearchRemoteReindexConfig(compartmentID, name, opensearchRemoteReindexConfigOptions{
		sources: []opensearchRemoteReindexSource{
			{caBundleResource: opensearchRemoteReindexCaResource, endpoint: endpoint, resolutionIP: common.String(resolutionIP)},
		},
	})
}

func testAccOpensearchRemoteReindexConfig(compartmentID, name string, options opensearchRemoteReindexConfigOptions) string {
	// Keep DEV-specific shapes in test configuration, not the production provider.
	nodeHostShape := utils.GetEnvSettingWithDefault("opensearch_node_host_shape", "VM.Standard.E3.Flex")
	trustConfig := ""
	if options.clearSources {
		trustConfig = "remote_reindex_trust_config {\n remote_sources = []\n}\n"
	} else if len(options.sources) > 0 {
		trustConfig = "remote_reindex_trust_config {\n"
		for _, source := range options.sources {
			trustConfig += fmt.Sprintf("remote_sources {\n ca_bundle_id = %s.id\n endpoint = %q\n", source.caBundleResource, source.endpoint)
			if source.resolutionIP != nil {
				trustConfig += fmt.Sprintf("endpoint_resolution_ip = %q\n", *source.resolutionIP)
			}
			trustConfig += "}\n"
		}
		trustConfig += "}\n"
	}
	otherCA := ""
	dependencies := opensearchRemoteReindexCaResource
	if options.replacementCA {
		dependencies += ", " + opensearchRemoteReindexOtherCaResource
		otherCA = fmt.Sprintf(`
resource "oci_certificates_management_ca_bundle" "test_remote_reindex_other_ca_bundle" {
  compartment_id = var.compartment_id
  name           = %q
  ca_bundle_pem  = %q
}
`, name+"-replacement", opensearchRemoteReindexReplacementCaPEM)
	}
	if options.dataNodeCount == 0 {
		options.dataNodeCount = 1
	}
	if options.dataNodeOCPUs == 0 {
		options.dataNodeOCPUs = 2
	}
	if options.dataNodeMemoryGB == 0 {
		options.dataNodeMemoryGB = 20
	}
	tags := ""
	if options.tagged {
		tags = `freeform_tags = { AcceptanceTest = "remote-reindex-preserved" }`
	}
	return acctest.ProviderTestConfig() + fmt.Sprintf(`
variable "compartment_id" {
  default = %[1]q
}

resource "oci_certificates_management_ca_bundle" "test_remote_reindex_ca_bundle" {
  compartment_id = var.compartment_id
  name           = %[2]q
  ca_bundle_pem  = %[4]q
}

%[5]s

resource "oci_core_vcn" "test_remote_reindex_vcn" {
  compartment_id = var.compartment_id
  cidr_block     = "10.0.0.0/16"
  lifecycle {
    ignore_changes = [defined_tags]
  }
}

resource "oci_core_subnet" "test_remote_reindex_subnet" {
  compartment_id = var.compartment_id
  cidr_block     = "10.0.0.0/24"
  vcn_id         = oci_core_vcn.test_remote_reindex_vcn.id
  lifecycle {
    ignore_changes = [defined_tags]
  }
}

resource "oci_opensearch_opensearch_cluster" "test_remote_reindex_cluster" {
  # Keep cleanup ordering even when Optional+Computed trust is omitted in a later step.
  depends_on                         = [%[11]s]
  compartment_id                     = var.compartment_id
  display_name                       = %[2]q
  software_version                   = "2.11.0"
  security_mode                      = "DISABLED"
  master_node_count                  = 1
  master_node_host_type              = "FLEX"
  master_node_host_shape             = %[3]q
  master_node_host_ocpu_count        = 1
  master_node_host_memory_gb         = 20
  data_node_count                    = %[7]d
  data_node_host_type                = "FLEX"
  data_node_host_shape               = %[3]q
  data_node_host_ocpu_count          = %[8]d
  data_node_host_memory_gb           = %[9]d
  data_node_storage_gb               = 50
  opendashboard_node_count           = 1
  opendashboard_node_host_shape      = %[3]q
  opendashboard_node_host_ocpu_count = 2
  opendashboard_node_host_memory_gb  = 10
  subnet_compartment_id              = var.compartment_id
  subnet_id                          = oci_core_subnet.test_remote_reindex_subnet.id
  vcn_compartment_id                 = var.compartment_id
  vcn_id                             = oci_core_vcn.test_remote_reindex_vcn.id

  %[6]s
  %[10]s

  lifecycle {
    ignore_changes = [system_tags, defined_tags]
  }
}
`, compartmentID, name, nodeHostShape, opensearchRemoteReindexTestCaPEM, otherCA, trustConfig,
		options.dataNodeCount, options.dataNodeOCPUs, options.dataNodeMemoryGB, tags, dependencies)
}
