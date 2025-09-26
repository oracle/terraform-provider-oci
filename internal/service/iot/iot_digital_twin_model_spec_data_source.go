// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package iot

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_iot "github.com/oracle/oci-go-sdk/v65/iot"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func IotDigitalTwinModelSpecDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: readSingularIotDigitalTwinModelSpecWithContext,
		Schema: map[string]*schema.Schema{
			"digital_twin_model_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
		},
	}
}

func readSingularIotDigitalTwinModelSpecWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &IotDigitalTwinModelSpecDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IotClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type IotDigitalTwinModelSpecDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_iot.IotClient
	Res    *oci_iot.GetDigitalTwinModelSpecResponse
}

func (s *IotDigitalTwinModelSpecDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *IotDigitalTwinModelSpecDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_iot.GetDigitalTwinModelSpecRequest{}

	if digitalTwinModelId, ok := s.D.GetOkExists("digital_twin_model_id"); ok {
		tmp := digitalTwinModelId.(string)
		request.DigitalTwinModelId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "iot")

	response, err := s.Client.GetDigitalTwinModelSpec(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *IotDigitalTwinModelSpecDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("IotDigitalTwinModelSpecDataSource-", IotDigitalTwinModelSpecDataSource(), s.D))

	return nil
}
