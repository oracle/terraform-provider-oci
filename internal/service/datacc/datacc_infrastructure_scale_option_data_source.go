// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package datacc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_datacc "github.com/oracle/oci-go-sdk/v65/datacc"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DataccInfrastructureScaleOptionDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: readSingularDataccInfrastructureScaleOptionWithContext,
		Schema: map[string]*schema.Schema{
			"infrastructure_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"possible_ssd_configurations": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func readSingularDataccInfrastructureScaleOptionWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DataccInfrastructureScaleOptionDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BaseinfraClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type DataccInfrastructureScaleOptionDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_datacc.BaseinfraClient
	Res    *oci_datacc.GetInfrastructureScaleOptionResponse
}

func (s *DataccInfrastructureScaleOptionDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataccInfrastructureScaleOptionDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_datacc.GetInfrastructureScaleOptionRequest{}

	if infrastructureId, ok := s.D.GetOkExists("infrastructure_id"); ok {
		tmp := infrastructureId.(string)
		request.InfrastructureId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "datacc")

	response, err := s.Client.GetInfrastructureScaleOption(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DataccInfrastructureScaleOptionDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DataccInfrastructureScaleOptionDataSource-", DataccInfrastructureScaleOptionDataSource(), s.D))

	s.D.Set("possible_ssd_configurations", s.Res.PossibleSsdConfigurations)

	return nil
}
