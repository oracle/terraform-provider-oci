// Copyright (c) 2017, 2026, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package gdp

import (
	"io"
	"net/http"
	"strings"
	"sync"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_gdp "github.com/oracle/oci-go-sdk/v65/gdp"

	"github.com/oracle/terraform-provider-oci/internal/client"
)

type successfulGdpDispatcher struct{}

func (successfulGdpDispatcher) Do(*http.Request) (*http.Response, error) {
	return &http.Response{
		StatusCode: http.StatusOK,
		Header:     make(http.Header),
		Body: io.NopCloser(strings.NewReader(
			`{"id":"ocid1.gdppipeline.oc1.iad.example","lifecycleState":"ACTIVE"}`,
		)),
	}, nil
}

type noOpGdpSigner struct{}

func (noOpGdpSigner) Sign(*http.Request) error { return nil }

func TestUnitGetGdpClientUsesOperationLocalCommercialClient(t *testing.T) {
	const (
		defaultHost    = "https://gdp.us-ashburn-1.oci.oraclecloud.com"
		commercialHost = "https://prod.cp.cdsaas.us-ashburn-1.oci.oraclecloud.com"
	)
	shared := &oci_gdp.GuardedDataPipelineClient{}
	shared.Host = defaultHost
	clients := &client.OracleClients{SdkClientMap: map[string]interface{}{
		"oci_gdp.GuardedDataPipelineClient": shared,
	}}

	const goroutines = 8
	var wg sync.WaitGroup
	wg.Add(goroutines)
	for range goroutines {
		go func() {
			defer wg.Done()
			operationClient := getGdpClient(clients, true)
			if operationClient == shared {
				t.Error("commercial operation reused the shared GDP client")
				return
			}
			if operationClient.Host != commercialHost {
				t.Errorf("commercial host = %q, want %q", operationClient.Host, commercialHost)
			}
			operationClient.Host = "https://operation-local.example.com"
		}()
	}
	wg.Wait()

	if shared.Host != defaultHost {
		t.Fatalf("shared GDP client host = %q, want %q", shared.Host, defaultHost)
	}
	if got := getGdpClient(clients, false); got != shared {
		t.Fatal("US Government operation did not reuse the unmodified default client")
	}
}

func TestUnitGdpReadDoesNotMutateSharedClient(t *testing.T) {
	const (
		defaultHost = "https://gdp.us-ashburn-1.oci.oraclecloud.com"
		goroutines  = 4
		iterations  = 30
	)

	shared := &oci_gdp.GuardedDataPipelineClient{}
	shared.Host = defaultHost
	shared.BasePath = "20230301"
	shared.UserAgent = "terraform-provider-oci-gdp-race-test"
	shared.Signer = noOpGdpSigner{}
	shared.HTTPClient = successfulGdpDispatcher{}
	clients := &client.OracleClients{SdkClientMap: map[string]interface{}{
		"oci_gdp.GuardedDataPipelineClient": shared,
	}}

	resourceData := make([]*schema.ResourceData, goroutines)
	for i := range resourceData {
		resourceData[i] = schema.TestResourceDataRaw(t, GdpGdpPipelineDataSource().Schema, map[string]interface{}{
			"gdp_pipeline_id": "ocid1.gdppipeline.oc1.iad.example",
			"env":             gdpCommercialCode,
		})
	}

	var wg sync.WaitGroup
	errCh := make(chan string, goroutines*iterations)
	start := make(chan struct{})
	wg.Add(goroutines)
	for i := range goroutines {
		go func(d *schema.ResourceData) {
			defer wg.Done()
			<-start
			for range iterations {
				diagnostics := readSingularGdpGdpPipelineWithContext(t.Context(), d, clients)
				if diagnostics.HasError() {
					errCh <- diagnostics[0].Summary
				}
			}
		}(resourceData[i])
	}
	close(start)
	wg.Wait()
	close(errCh)
	for message := range errCh {
		t.Errorf("GDP read returned an error diagnostic: %s", message)
	}

	if shared.Host != defaultHost {
		t.Fatalf("shared GDP client host = %q, want %q", shared.Host, defaultHost)
	}
}
