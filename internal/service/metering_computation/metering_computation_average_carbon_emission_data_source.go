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

func MeteringComputationAverageCarbonEmissionDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularMeteringComputationAverageCarbonEmission,
		Schema: map[string]*schema.Schema{
			"sku_part_number": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"average_carbon_emission": {
				Type:     schema.TypeFloat,
				Computed: true,
			},
		},
	}
}

func readSingularMeteringComputationAverageCarbonEmission(d *schema.ResourceData, m interface{}) error {
	sync := &MeteringComputationAverageCarbonEmissionDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).UsageapiClient()

	return tfresource.ReadResource(sync)
}

type MeteringComputationAverageCarbonEmissionDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_metering_computation.UsageapiClient
	Res    *oci_metering_computation.RequestAverageCarbonEmissionResponse
}

func (s *MeteringComputationAverageCarbonEmissionDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *MeteringComputationAverageCarbonEmissionDataSourceCrud) Get() error {
	request := oci_metering_computation.RequestAverageCarbonEmissionRequest{}

	if skuPartNumber, ok := s.D.GetOkExists("sku_part_number"); ok {
		tmp := skuPartNumber.(string)
		request.SkuPartNumber = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "metering_computation")

	response, err := s.Client.RequestAverageCarbonEmission(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *MeteringComputationAverageCarbonEmissionDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("MeteringComputationAverageCarbonEmissionDataSource-", MeteringComputationAverageCarbonEmissionDataSource(), s.D))

	s.D.Set("average_carbon_emission", *s.Res.AverageCarbonEmission.AverageCarbonEmission)

	return nil
}
