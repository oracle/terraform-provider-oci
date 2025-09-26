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

func IotDigitalTwinInstanceContentDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: readSingularIotDigitalTwinInstanceContentWithContext,
		Schema: map[string]*schema.Schema{
			"digital_twin_instance_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"should_include_metadata": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			// Computed
		},
	}
}

func readSingularIotDigitalTwinInstanceContentWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &IotDigitalTwinInstanceContentDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IotClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type IotDigitalTwinInstanceContentDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_iot.IotClient
	Res    *oci_iot.GetDigitalTwinInstanceContentResponse
}

func (s *IotDigitalTwinInstanceContentDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *IotDigitalTwinInstanceContentDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_iot.GetDigitalTwinInstanceContentRequest{}

	if digitalTwinInstanceId, ok := s.D.GetOkExists("digital_twin_instance_id"); ok {
		tmp := digitalTwinInstanceId.(string)
		request.DigitalTwinInstanceId = &tmp
	}

	if shouldIncludeMetadata, ok := s.D.GetOkExists("should_include_metadata"); ok {
		tmp := shouldIncludeMetadata.(bool)
		request.ShouldIncludeMetadata = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "iot")

	response, err := s.Client.GetDigitalTwinInstanceContent(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *IotDigitalTwinInstanceContentDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("IotDigitalTwinInstanceContentDataSource-", IotDigitalTwinInstanceContentDataSource(), s.D))

	return nil
}
