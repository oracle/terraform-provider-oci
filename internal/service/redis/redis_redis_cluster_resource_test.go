// Copyright (c) 2026, Oracle and/or its affiliates.
// Licensed under the Mozilla Public License v2.0

package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_redis "github.com/oracle/oci-go-sdk/v65/redis"

	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func TestRedisRedisClusterUpdateVersionAndConfigSet(t *testing.T) {
	tests := []struct {
		name       string
		changes    map[string]string
		requests   []map[string]interface{}
		failUpdate bool
	}{
		{
			name:    "version and config set in one request",
			changes: map[string]string{"software_version": "VALKEY_8_1", "oci_cache_config_set_id": "new-config-set"},
			requests: []map[string]interface{}{
				{"softwareVersion": "VALKEY_8_1", "ociCacheConfigSetId": "new-config-set"},
			},
		},
		{
			name:     "config set only",
			changes:  map[string]string{"oci_cache_config_set_id": "compatible-config-set"},
			requests: []map[string]interface{}{{"ociCacheConfigSetId": "compatible-config-set"}},
		},
		{
			name:     "version only does not resend unchanged config set",
			changes:  map[string]string{"software_version": "VALKEY_8_1"},
			requests: []map[string]interface{}{{"softwareVersion": "VALKEY_8_1"}},
		},
		{
			name: "unchanged fields send no update",
		},
		{
			name:    "display name remains separate from migration",
			changes: map[string]string{"display_name": "updated-name", "software_version": "VALKEY_8_1", "oci_cache_config_set_id": "new-config-set"},
			requests: []map[string]interface{}{
				{"displayName": "updated-name"},
				{"softwareVersion": "VALKEY_8_1", "ociCacheConfigSetId": "new-config-set"},
			},
		},
		{
			name:       "migration failure is returned without a second update",
			changes:    map[string]string{"software_version": "VALKEY_8_1", "oci_cache_config_set_id": "new-config-set"},
			requests:   []map[string]interface{}{{"softwareVersion": "VALKEY_8_1", "ociCacheConfigSetId": "new-config-set"}},
			failUpdate: true,
		},
		{
			name:       "earlier update failure prevents migration",
			changes:    map[string]string{"display_name": "updated-name", "software_version": "VALKEY_8_1", "oci_cache_config_set_id": "new-config-set"},
			requests:   []map[string]interface{}{{"displayName": "updated-name"}},
			failUpdate: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			state := map[string]string{
				"display_name": "original-name", "software_version": "VALKEY_7_2",
				"oci_cache_config_set_id": "old-config-set", "primary_cluster_id": "",
			}
			diff := map[string]*terraform.ResourceAttrDiff{}
			for key, value := range tt.changes {
				diff[key] = &terraform.ResourceAttrDiff{Old: state[key], New: value}
			}
			data, err := schema.InternalMap(RedisRedisClusterResource().Schema).Data(
				&terraform.InstanceState{ID: "cluster-test-id", Attributes: state},
				&terraform.InstanceDiff{Attributes: diff},
			)
			if err != nil {
				t.Fatal(err)
			}

			var requests []map[string]interface{}
			workRequests, reads := 0, 0
			cluster := map[string]interface{}{
				"id": "cluster-test-id", "lifecycleState": "ACTIVE",
				"softwareVersion": "VALKEY_7_2", "ociCacheConfigSetId": "old-config-set",
			}
			client := &oci_redis.RedisClusterClient{BaseClient: oci_common.DefaultBaseClientWithSigner(redisUpdateTestSigner{})}
			client.Host = "https://redis.invalid"
			client.BasePath = "/20220315"
			// All SDK requests are handled in memory: no credentials or OCI endpoints are used.
			client.HTTPClient = redisUpdateTestTransport(func(r *http.Request) (*http.Response, error) {
				status, body := http.StatusOK, "{}"
				headers := http.Header{"Content-Type": []string{"application/json"}}
				switch {
				case r.Method == http.MethodPut && strings.HasSuffix(r.URL.Path, "/redisClusters/cluster-test-id"):
					var fields map[string]interface{}
					if err := json.NewDecoder(r.Body).Decode(&fields); err != nil {
						return nil, err
					}
					// The SDK serializes some unset optional fields as null.
					for key, value := range fields {
						if value == nil {
							delete(fields, key)
						}
					}
					requests = append(requests, fields)
					if tt.failUpdate {
						status = http.StatusBadRequest
						body = `{"code":"InvalidParameter","message":"test update rejected"}`
					} else {
						for key, value := range fields {
							cluster[key] = value
						}
						headers.Set("opc-work-request-id", "work-request-test-id")
					}
				case r.Method == http.MethodGet && strings.HasSuffix(r.URL.Path, "/workRequests/work-request-test-id"):
					workRequests++
					body = `{"id":"work-request-test-id","status":"SUCCEEDED","timeFinished":"2026-01-01T00:00:00Z","resources":[{"entityType":"cluster","actionType":"UPDATED","identifier":"cluster-test-id"}]}`
				case r.Method == http.MethodGet && strings.HasSuffix(r.URL.Path, "/redisClusters/cluster-test-id"):
					reads++
					payload, err := json.Marshal(cluster)
					if err != nil {
						return nil, err
					}
					body = string(payload)
				default:
					return nil, fmt.Errorf("unexpected request: %s %s", r.Method, r.URL.Path)
				}
				return &http.Response{StatusCode: status, Header: headers, Body: io.NopCloser(strings.NewReader(body)), Request: r}, nil
			})
			crud := RedisRedisClusterResourceCrud{BaseCrud: tfresource.BaseCrud{D: data}, Client: client}
			err = crud.UpdateWithContext(context.Background())
			if tt.failUpdate {
				if err == nil || !strings.Contains(err.Error(), "test update rejected") {
					t.Fatalf("expected update error, got %v", err)
				}
			} else if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(requests, tt.requests) {
				t.Fatalf("update requests = %#v; want %#v", requests, tt.requests)
			}
			wantReads := len(tt.requests)
			if tt.failUpdate {
				wantReads = 0
			}
			if workRequests != wantReads || reads != wantReads {
				t.Fatalf("work-request polls=%d, cluster reads=%d; want %d each", workRequests, reads, wantReads)
			}
		})
	}
}

type redisUpdateTestTransport func(*http.Request) (*http.Response, error)

func (f redisUpdateTestTransport) Do(r *http.Request) (*http.Response, error) { return f(r) }

type redisUpdateTestSigner struct{}

func (redisUpdateTestSigner) Sign(*http.Request) error { return nil }
