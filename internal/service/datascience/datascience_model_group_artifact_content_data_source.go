// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package datascience

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_datascience "github.com/oracle/oci-go-sdk/v65/datascience"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatascienceModelGroupArtifactContentDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: readSingularDatascienceModelGroupArtifactContentWithContext,
		Schema: map[string]*schema.Schema{
			"model_group_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"range": {
				Type:     schema.TypeString,
				Optional: true,
			},
			// Computed
		},
	}
}

func readSingularDatascienceModelGroupArtifactContentWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DatascienceModelGroupArtifactContentDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataScienceClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type DatascienceModelGroupArtifactContentDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_datascience.DataScienceClient
	Res    *oci_datascience.GetModelGroupArtifactContentResponse
}

func (s *DatascienceModelGroupArtifactContentDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatascienceModelGroupArtifactContentDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_datascience.GetModelGroupArtifactContentRequest{}

	if modelGroupId, ok := s.D.GetOkExists("model_group_id"); ok {
		tmp := modelGroupId.(string)
		request.ModelGroupId = &tmp
	}

	if range_, ok := s.D.GetOkExists("range"); ok {
		tmp := range_.(string)
		request.Range = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "datascience")

	response, err := s.Client.GetModelGroupArtifactContent(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatascienceModelGroupArtifactContentDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatascienceModelGroupArtifactContentDataSource-", DatascienceModelGroupArtifactContentDataSource(), s.D))

	return nil
}
