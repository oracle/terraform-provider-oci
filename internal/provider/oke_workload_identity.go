// Copyright (c) 2026, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package provider

import (
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_common_auth "github.com/oracle/oci-go-sdk/v65/common/auth"
)

var okeWorkloadIdentityConfigurationProviderFn = oci_common_auth.OkeWorkloadIdentityConfigurationProvider

type okeWorkloadIdentityRegionOverrideProvider struct {
	oci_common_auth.ConfigurationProviderWithClaimAccess
	region string
}

var _ oci_common_auth.ConfigurationProviderWithClaimAccess = (*okeWorkloadIdentityRegionOverrideProvider)(nil)
var _ oci_common.RefreshableConfigurationProvider = (*okeWorkloadIdentityRegionOverrideProvider)(nil)

func (p *okeWorkloadIdentityRegionOverrideProvider) Region() (string, error) {
	return p.region, nil
}

func (p *okeWorkloadIdentityRegionOverrideProvider) Refreshable() bool {
	refreshableProvider, ok := p.ConfigurationProviderWithClaimAccess.(oci_common.RefreshableConfigurationProvider)
	return ok && refreshableProvider.Refreshable()
}

func newOKEWorkloadIdentityConfigurationProvider(targetRegion string) (oci_common_auth.ConfigurationProviderWithClaimAccess, error) {
	configProvider, err := okeWorkloadIdentityConfigurationProviderFn()
	if err != nil {
		return nil, err
	}

	if targetRegion == "" {
		return configProvider, nil
	}

	return &okeWorkloadIdentityRegionOverrideProvider{
		ConfigurationProviderWithClaimAccess: configProvider,
		region:                               targetRegion,
	}, nil
}
