// Copyright (c) 2026, Oracle and/or its affiliates.
// Licensed under the Mozilla Public License Version 2.0

// Package oci exposes the supported Go construction boundary for embedding
// Terraform Provider OCI in another Go process.
package oci

import (
	frameworkprovider "github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	internalprovider "github.com/oracle/terraform-provider-oci/internal/provider"
)

// Provider returns a fresh SDKv2 provider instance configured for safe
// in-process embedding.
func Provider() *schema.Provider {
	return internalprovider.NewSDKv2ProviderForInProcess()
}

// ProviderForResources returns a fresh SDKv2 provider containing only the
// named resource schemas and no data-source schemas. The returned resource
// schemas remain isolated from other provider instances.
func ProviderForResources(resourceNames ...string) (*schema.Provider, error) {
	return internalprovider.NewSDKv2ProviderForInProcessResources(resourceNames...)
}

// ProviderForConfiguration returns a fresh SDKv2 provider that can validate
// and configure credentials without retaining resource or data-source schemas.
func ProviderForConfiguration() *schema.Provider {
	return internalprovider.NewSDKv2ProviderForInProcessConfiguration()
}

// New returns a fresh Plugin Framework provider instance configured for safe
// in-process embedding.
func New() frameworkprovider.Provider {
	return internalprovider.NewFrameworkProviderForInProcess()
}

// InProcessFileCredentialFingerprint returns a stable digest of the contents
// of files that can affect an embedded provider configuration. Hosts that cache
// configured provider metadata can include this value in their cache identity
// so replacing a mounted credential file invalidates the cached metadata.
func InProcessFileCredentialFingerprint(config map[string]any) (string, error) {
	return internalprovider.InProcessFileCredentialFingerprint(config)
}
