// Copyright (c) 2017, 2026, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package kms

import (
	"errors"
	"net"
	"net/url"
	"testing"
	"time"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
)

func TestUnitKmsKeyCreateDNSNotFoundError(t *testing.T) {
	dnsError := &net.DNSError{
		Err:        "no such host",
		Name:       "example-management.kms.us-phoenix-1.oraclecloud.com",
		IsNotFound: true,
	}
	wrappedDNSError := &url.Error{
		Op:  "Post",
		URL: "https://example-management.kms.us-phoenix-1.oraclecloud.com/20180608/keys",
		Err: &net.OpError{
			Op:  "dial",
			Net: "tcp",
			Err: dnsError,
		},
	}

	tests := []struct {
		name     string
		err      error
		expected bool
	}{
		{
			name:     "wrapped DNS not found error",
			err:      wrappedDNSError,
			expected: true,
		},
		{
			name:     "direct DNS not found error",
			err:      dnsError,
			expected: true,
		},
		{
			name: "DNS timeout without not found",
			err: &net.DNSError{
				Err:       "i/o timeout",
				Name:      "example-management.kms.us-phoenix-1.oraclecloud.com",
				IsTimeout: true,
			},
			expected: false,
		},
		{
			name:     "error text alone is not classified",
			err:      errors.New("lookup example-management.kms.us-phoenix-1.oraclecloud.com: no such host"),
			expected: false,
		},
		{
			name:     "nil error",
			err:      nil,
			expected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			response := oci_common.OCIOperationResponse{Error: test.err}
			if actual := isKmsKeyCreateDNSNotFoundError(response); actual != test.expected {
				t.Fatalf("expected DNS not found classification %t, got %t", test.expected, actual)
			}
		})
	}
}

func TestUnitKmsKeyCreateRetryPolicy(t *testing.T) {
	dnsError := &url.Error{
		Op:  "Post",
		URL: "https://example-management.kms.us-phoenix-1.oraclecloud.com/20180608/keys",
		Err: &net.OpError{
			Op:  "dial",
			Net: "tcp",
			Err: &net.DNSError{
				Err:        "no such host",
				Name:       "example-management.kms.us-phoenix-1.oraclecloud.com",
				IsNotFound: true,
			},
		},
	}

	retryPolicy := getKmsKeyCreateRetryPolicy()
	if !retryPolicy.ShouldRetryOperation(oci_common.OCIOperationResponse{Error: dnsError}) {
		t.Fatal("expected KMS key create retry policy to retry DNS not found error")
	}

	if retryPolicy.ShouldRetryOperation(oci_common.OCIOperationResponse{Error: errors.New("invalid endpoint")}) {
		t.Fatal("expected KMS key create retry policy not to retry unrelated error")
	}
}

func TestUnitCapKmsKeyCreateDNSRetryBackoff(t *testing.T) {
	if kmsKeyCreateDNSRetryTimeout != 30*time.Second {
		t.Fatalf("expected KMS key create DNS retry timeout to be 30s, got %s", kmsKeyCreateDNSRetryTimeout)
	}

	tests := []struct {
		name     string
		backoff  time.Duration
		elapsed  time.Duration
		expected time.Duration
	}{
		{
			name:     "backoff within remaining window",
			backoff:  2 * time.Second,
			elapsed:  5 * time.Second,
			expected: 2 * time.Second,
		},
		{
			name:     "backoff capped to remaining window",
			backoff:  20 * time.Second,
			elapsed:  25 * time.Second,
			expected: 5 * time.Second,
		},
		{
			name:     "retry window exhausted",
			backoff:  time.Second,
			elapsed:  30 * time.Second,
			expected: 0,
		},
		{
			name:     "retry window exceeded",
			backoff:  time.Second,
			elapsed:  31 * time.Second,
			expected: 0,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if actual := capKmsKeyCreateDNSRetryBackoff(test.backoff, test.elapsed); actual != test.expected {
				t.Fatalf("expected capped backoff %s, got %s", test.expected, actual)
			}
		})
	}
}
