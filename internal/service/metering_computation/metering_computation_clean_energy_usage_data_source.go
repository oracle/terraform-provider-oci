// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package metering_computation

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_metering_computation "github.com/oracle/oci-go-sdk/v65/usageapi"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func MeteringComputationCleanEnergyUsageDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularMeteringComputationCleanEnergyUsage,
		Schema: map[string]*schema.Schema{
			"ad": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"region": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"usage": {
				Type:     schema.TypeFloat,
				Computed: true,
			},
		},
	}
}

func readSingularMeteringComputationCleanEnergyUsage(d *schema.ResourceData, m interface{}) error {
	sync := &MeteringComputationCleanEnergyUsageDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).UsageapiClient()

	return tfresource.ReadResource(sync)
}

type MeteringComputationCleanEnergyUsageDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_metering_computation.UsageapiClient
	Res    *oci_metering_computation.RequestCleanEnergyUsageResponse
}

func (s *MeteringComputationCleanEnergyUsageDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *MeteringComputationCleanEnergyUsageDataSourceCrud) Get() error {
	request := oci_metering_computation.RequestCleanEnergyUsageRequest{}

	if ad, ok := s.D.GetOkExists("ad"); ok {
		tmp := ad.(string)
		request.Ad = &tmp
	}

	if region, ok := s.D.GetOkExists("region"); ok {
		tmp := region.(string)
		request.Region = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "metering_computation")

	response, err := s.Client.RequestCleanEnergyUsage(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *MeteringComputationCleanEnergyUsageDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("MeteringComputationCleanEnergyUsageDataSource-", MeteringComputationCleanEnergyUsageDataSource(), s.D))

	if s.Res.Ad != nil {
		s.D.Set("ad", *s.Res.Ad)
	}

	if s.Res.Usage != nil {
		s.D.Set("usage", *s.Res.Usage)
	}

	return nil
}
