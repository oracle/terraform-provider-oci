package provider

import (
	"crypto/tls"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"testing"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/terraform-provider-oci/internal/globalvar"
	"github.com/oracle/terraform-provider-oci/internal/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// ensure a custom domain can be targeted and expected http client settings are preserved
// issue-routing-tag: terraform/default
func TestUnitBuildClientConfigureFn_withDomainNameOverride_inIsolation(t *testing.T) {

	// --- Child process: Run actual test and exit the flow after running test logic
	if subProcessFlag := os.Getenv("GO_TEST_ISOLATED_SUBPROCESS"); subProcessFlag != "" {
		t.Setenv(globalvar.DomainNameOverrideEnv, "0r4-c10ud.com")
		t.Setenv(globalvar.HasCorrectDomainNameEnv, "")
		t.Setenv(globalvar.OciRegionMetadataEnv, `{"realmKey":"OC150","realmDomainComponent":"oraclecloud150.sol.mar","regionKey":"mars","regionIdentifier":"solar-mars-1"}`)

		assert.Equal(t, "0r4-c10ud.com", utils.GetEnvSettingWithBlankDefault(globalvar.DomainNameOverrideEnv))
		assert.Equal(t, `{"realmKey":"OC150","realmDomainComponent":"oraclecloud150.sol.mar","regionKey":"mars","regionIdentifier":"solar-mars-1"}`, utils.GetEnvSettingWithBlankDefault(globalvar.OciRegionMetadataEnv))
		configProvider := oci_common.NewRawConfigurationProvider(testTenancyOCID, testUserOCID, "solar-mars-1", testKeyFingerPrint, testPrivateKey, nil)
		httpClient := BuildHttpClient()
		configureClientFn, err := BuildConfigureClientFn(configProvider, httpClient)
		assert.NoError(t, err)

		baseClient := &oci_common.BaseClient{}
		baseClient.Host = oci_common.StringToRegion("mars").Endpoint("megatron")
		assert.Equal(t, `megatron.solar-mars-1.oraclecloud150.sol.mar`, baseClient.Host)
		err = configureClientFn(baseClient)
		assert.NoError(t, err)

		// verify transport settings are unchanged
		tr := httpClient.Transport.(*http.Transport)
		assert.NotNil(t, tr.TLSClientConfig)
		assert.Equal(t, uint16(tls.VersionTLS12), tr.TLSClientConfig.MinVersion, "expected min tls 1.2")
		assert.NotNil(t, tr.Proxy, "expected http.ProxyFromEnvironment fn")
		assert.Nil(t, tr.TLSClientConfig.RootCAs)

		// verify url has expected domain
		assert.Equal(t, `megatron.solar-mars-1.0r4-c10ud.com`, baseClient.Host)

		// verify subdomains are preserved
		baseClient = &oci_common.BaseClient{}
		baseClient.Host = "avnzdivwaadfa-management.kms.solar-mars-1.oraclecloud150.sol.mar"
		err = configureClientFn(baseClient)
		assert.NoError(t, err)
		assert.Equal(t, `avnzdivwaadfa-management.kms.solar-mars-1.0r4-c10ud.com`, baseClient.Host)

		// verify non-match preserves original url
		baseClient = &oci_common.BaseClient{}
		baseClient.Host = "DUMMY_ENDPOINT"
		err = configureClientFn(baseClient)
		assert.NoError(t, err)
		assert.Equal(t, `DUMMY_ENDPOINT`, baseClient.Host)

		os.Exit(0)
	}

	//Parent process: Make the test run in a new process for isolation.
	exe, err := os.Executable()
	if err != nil {
		t.Fatal(err)
	}

	name := t.Name()
	cmd := exec.Command(exe, "-test.v", "-test.run", "^"+name+"$")
	cmd.Env = append(os.Environ(), "GO_TEST_ISOLATED_SUBPROCESS="+name)
	combinedStdOutStdErr, err := cmd.CombinedOutput()
	outputStr := string(combinedStdOutStdErr)
	if err != nil {
		t.Fatalf("%s failed: %v\n%s", name, err, outputStr)
	}
	if combinedStdOutStdErr != nil &&
		(strings.Contains(strings.ToLower(outputStr), "error") || strings.Contains(strings.ToLower(outputStr), "fail")) {
		t.Fatalf("%s failed: \n%s", name, string(combinedStdOutStdErr))
	}
}

// ensure a custom domain that has already override with more than 2 dots can be targeted and expected http client settings are preserved
// issue-routing-tag: terraform/default
func TestUnitBuildClientConfigureFn_withDomainNameOverrideAndCorrectDomainName_inIsolation(t *testing.T) {

	// --- Child process: Run actual test and exit the flow after running test logic
	if subProcessFlag := os.Getenv("GO_TEST_ISOLATED_SUBPROCESS"); subProcessFlag != "" {
		t.Setenv(globalvar.DomainNameOverrideEnv, "oc.0r4-c10ud.com")
		t.Setenv(globalvar.HasCorrectDomainNameEnv, "oc.0r4-c10ud.com")
		assert.Equal(t, "oc.0r4-c10ud.com", utils.GetEnvSettingWithBlankDefault(globalvar.DomainNameOverrideEnv))
		assert.Equal(t, "oc.0r4-c10ud.com", utils.GetEnvSettingWithBlankDefault(globalvar.HasCorrectDomainNameEnv))
		configProvider := oci_common.DefaultConfigProvider()
		httpClient := BuildHttpClient()
		configureClientFn, err := BuildConfigureClientFn(configProvider, httpClient)
		assert.NoError(t, err)

		baseClient := &oci_common.BaseClient{}
		baseClient.Host = "https://svc.region.oc.0r4-c10ud.com"
		err = configureClientFn(baseClient)
		assert.NoError(t, err)

		// verify transport settings are unchanged
		tr := httpClient.Transport.(*http.Transport)
		assert.NotNil(t, tr.TLSClientConfig)
		assert.Equal(t, uint16(tls.VersionTLS12), tr.TLSClientConfig.MinVersion, "expected min tls 1.2")
		assert.NotNil(t, tr.Proxy, "expected http.ProxyFromEnvironment fn")
		assert.Nil(t, tr.TLSClientConfig.RootCAs)

		// verify url has expected domain
		assert.Equal(t, `https://svc.region.oc.0r4-c10ud.com`, baseClient.Host)

		// verify subdomains are preserved
		baseClient = &oci_common.BaseClient{}
		baseClient.Host = "avnzdivwaadfa-management.kms.us-phoenix-1.oraclecloud.com"
		err = configureClientFn(baseClient)
		assert.NoError(t, err)
		assert.Equal(t, `avnzdivwaadfa-management.kms.us-phoenix-1.oc.0r4-c10ud.com`, baseClient.Host)

		// verify non-match preserves original url
		baseClient = &oci_common.BaseClient{}
		baseClient.Host = "DUMMY_ENDPOINT"
		err = configureClientFn(baseClient)
		assert.NoError(t, err)
		assert.Equal(t, `DUMMY_ENDPOINT`, baseClient.Host)

		os.Exit(0)
	}

	//Parent process: Make the test run in a new process for isolation.
	exe, err := os.Executable()
	if err != nil {
		t.Fatal(err)
	}

	name := t.Name()
	cmd := exec.Command(exe, "-test.v", "-test.run", "^"+name+"$")
	cmd.Env = append(os.Environ(), "GO_TEST_ISOLATED_SUBPROCESS="+name)
	combinedStdOutStdErr, err := cmd.CombinedOutput()
	outputStr := string(combinedStdOutStdErr)
	if err != nil {
		t.Fatalf("%s failed: %v\n%s", name, err, outputStr)
	}
	if combinedStdOutStdErr != nil &&
		(strings.Contains(strings.ToLower(outputStr), "error") || strings.Contains(strings.ToLower(outputStr), "fail")) {
		t.Fatalf("%s failed: \n%s", name, string(combinedStdOutStdErr))
	}
}

// Ensure that domain override doesnt overrides domain when custom endpoint template with static domain is used for end point creation.
func TestUnitBuildClientConfigureFn_withDomainNameOverride_andCustomServiceEndpointTemplate_withoutRealmDomainInTemplate_inIsolation(t *testing.T) {
	// --- Child process: Run actual test and exit the flow after running test logic
	if subProcessFlag := os.Getenv("GO_TEST_ISOLATED_SUBPROCESS"); subProcessFlag != "" {
		domainOverride := "orac.cloud80.com"
		region := "solar-mars-1"
		expectedHost := "adm.solar-mars-1.oci.tel"

		t.Setenv(globalvar.DomainNameOverrideEnv, domainOverride)

		configurationProvider := oci_common.NewRawConfigurationProvider(
			testTenancyOCID, testUserOCID, region,
			testKeyFingerPrint, testPrivateKey, nil)

		configure, err := BuildConfigureClientFn(configurationProvider, http.DefaultClient)
		require.NoError(t, err)

		baseClient := &oci_common.BaseClient{
			Host: oci_common.StringToRegion("solar-mars-1").EndpointForTemplate("megatron", "adm.{region}.oci.tel"),
		}
		require.NoError(t, configure(baseClient))

		assert.Equal(t, expectedHost, baseClient.Host)

		os.Exit(0)
	}

	//Parent process: Make the test run in a new process for isolation.
	exe, err := os.Executable()
	if err != nil {
		t.Fatal(err)
	}

	name := t.Name()
	cmd := exec.Command(exe, "-test.v", "-test.run", "^"+name+"$")
	cmd.Env = append(os.Environ(), "GO_TEST_ISOLATED_SUBPROCESS="+name)
	combinedStdOutStdErr, err := cmd.CombinedOutput()
	outputStr := string(combinedStdOutStdErr)
	if err != nil {
		t.Fatalf("%s failed: %v\n%s", name, err, outputStr)
	}
	if combinedStdOutStdErr != nil &&
		(strings.Contains(strings.ToLower(outputStr), "error") || strings.Contains(strings.ToLower(outputStr), "fail")) {
		t.Fatalf("%s failed: \n%s", name, string(combinedStdOutStdErr))
	}
}

// Ensure that domain override doesn't override realm domain when realm specific endpoint template with static domain is used.
func TestUnitBuildClientConfigureFn_withDomainNameOverride_andRealmSpecificEndpointEnabled_inIsolation(t *testing.T) {

	// --- Child process: Run actual test and exit the flow after running test logic
	if subProcessFlag := os.Getenv("GO_TEST_ISOLATED_SUBPROCESS"); subProcessFlag != "" {
		domainOverride := "orac.cloud80.com"
		region := "solar-mars-1"
		objectStorageNamespace := "abcdefghijklmnop"

		realmSpecificEndpointTemplate := objectStorageNamespace + ".objectstorage.{region}.oci.customer-oci.com"
		expectedHost := objectStorageNamespace + "." + "objectstorage" + "." + region + "." + "oci.customer-oci.com"

		t.Setenv(globalvar.DomainNameOverrideEnv, domainOverride)

		configurationProvider := oci_common.NewRawConfigurationProvider(
			testTenancyOCID, testUserOCID, region,
			testKeyFingerPrint, testPrivateKey, nil)

		configure, err := BuildConfigureClientFn(configurationProvider, http.DefaultClient)
		require.NoError(t, err)

		baseClient := &oci_common.BaseClient{
			Host: oci_common.StringToRegion("solar-mars-1").EndpointForTemplate("megatron", realmSpecificEndpointTemplate),
		}
		require.NoError(t, configure(baseClient))

		assert.Equal(t, expectedHost, baseClient.Host)

		os.Exit(0)
	}

	//Parent process: Make the test run in a new process for isolation.
	exe, err := os.Executable()
	if err != nil {
		t.Fatal(err)
	}

	name := t.Name()
	cmd := exec.Command(exe, "-test.v", "-test.run", "^"+name+"$")
	cmd.Env = append(os.Environ(), "GO_TEST_ISOLATED_SUBPROCESS="+name)
	combinedStdOutStdErr, err := cmd.CombinedOutput()
	outputStr := string(combinedStdOutStdErr)
	if err != nil {
		t.Fatalf("%s failed: %v\n%s", name, err, outputStr)
	}
	if combinedStdOutStdErr != nil &&
		(strings.Contains(strings.ToLower(outputStr), "error") || strings.Contains(strings.ToLower(outputStr), "fail")) {
		t.Fatalf("%s failed: \n%s", name, string(combinedStdOutStdErr))
	}
}

func TestBuildClientConfigureFn_hostMutation_DomainOverride_inIsolation(t *testing.T) {
	tests := []struct {
		name           string
		regionMetaData string
		override       string
		expectedHost   string
	}{
		{"with region metadata and realm domain same as override",
			`{"realmKey":"OC150","realmDomainComponent":"oc140.usa.sec","regionKey":"mars","regionIdentifier":"solar-mars-1"}`,
			"oc140.usa.sec",
			"megatron.solar-mars-1.oc140.usa.sec"},
		{"with region metadata and realm domain not same as override",
			`{"realmKey":"OC150","realmDomainComponent":"oraclecloud150.sol.mar","regionKey":"mars","regionIdentifier":"solar-mars-1"}`,
			"oc140.usa.sec",
			"megatron.solar-mars-1.oc140.usa.sec"},
		{"no region metadata",
			"",
			"oc140.usa.sec",
			"megatron.solar-mars-1.oc140.usa.sec"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			// --- Child process: Run actual test and exit the flow after running test logic
			if subProcessFlag := os.Getenv("GO_TEST_ISOLATED_SUBPROCESS"); subProcessFlag != "" {
				t.Setenv(globalvar.DomainNameOverrideEnv, tc.override)
				t.Setenv(globalvar.OciRegionMetadataEnv, tc.regionMetaData)

				configurationProvider := oci_common.NewRawConfigurationProvider(
					testTenancyOCID, testUserOCID, "solar-mars-1",
					testKeyFingerPrint, testPrivateKey, nil)

				configure, err := BuildConfigureClientFn(configurationProvider, http.DefaultClient)
				require.NoError(t, err)

				baseClient := &oci_common.BaseClient{
					Host: oci_common.StringToRegion("solar-mars-1").Endpoint("megatron"),
				}
				require.NoError(t, configure(baseClient))

				assert.Equal(t, tc.expectedHost, baseClient.Host)

				os.Exit(0)
			}

			//Parent process: Make the test run in a new process for isolation.
			exe, err := os.Executable()
			if err != nil {
				t.Fatal(err)
			}

			name := t.Name()
			cmd := exec.Command(exe, "-test.v", "-test.run", "^"+name+"$")
			cmd.Env = append(os.Environ(), "GO_TEST_ISOLATED_SUBPROCESS="+name)
			combinedStdOutStdErr, err := cmd.CombinedOutput()
			outputStr := string(combinedStdOutStdErr)
			if err != nil {
				t.Fatalf("%s failed: %v\n%s", name, err, outputStr)
			}
			if combinedStdOutStdErr != nil &&
				(strings.Contains(strings.ToLower(outputStr), "error") || strings.Contains(strings.ToLower(outputStr), "fail")) {
				t.Fatalf("%s failed: \n%s", name, string(combinedStdOutStdErr))
			}
		})
	}
}
