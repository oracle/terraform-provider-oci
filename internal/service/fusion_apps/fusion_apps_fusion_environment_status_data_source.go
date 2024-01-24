// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package fusion_apps

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_fusion_apps "github.com/oracle/oci-go-sdk/v65/fusionapps"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func FusionAppsFusionEnvironmentStatusDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularFusionAppsFusionEnvironmentStatus,
		Schema: map[string]*schema.Schema{
			"fusion_environment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularFusionAppsFusionEnvironmentStatus(d *schema.ResourceData, m interface{}) error {
	sync := &FusionAppsFusionEnvironmentStatusDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FusionApplicationsClient()

	return tfresource.ReadResource(sync)
}

type FusionAppsFusionEnvironmentStatusDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_fusion_apps.FusionApplicationsClient
	Res    *oci_fusion_apps.GetFusionEnvironmentStatusResponse
}

func (s *FusionAppsFusionEnvironmentStatusDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *FusionAppsFusionEnvironmentStatusDataSourceCrud) Get() error {
	request := oci_fusion_apps.GetFusionEnvironmentStatusRequest{}

	if fusionEnvironmentId, ok := s.D.GetOkExists("fusion_environment_id"); ok {
		tmp := fusionEnvironmentId.(string)
		request.FusionEnvironmentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "fusion_apps")

	response, err := s.Client.GetFusionEnvironmentStatus(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *FusionAppsFusionEnvironmentStatusDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("FusionAppsFusionEnvironmentStatusDataSource-", FusionAppsFusionEnvironmentStatusDataSource(), s.D))

	s.D.Set("status", s.Res.Status)

	return nil
}
