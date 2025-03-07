// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package datascience

import (
	"context"

	"github.com/oracle/terraform-provider-oci/internal/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_datascience "github.com/oracle/oci-go-sdk/v65/datascience"

	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatascienceModelCustomMetadataArtifactContentDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularDatascienceModelCustomMetadataArtifactContent,
		Schema: map[string]*schema.Schema{
			"metadatum_key_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"model_id": {
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

func readSingularDatascienceModelCustomMetadataArtifactContent(d *schema.ResourceData, m interface{}) error {
	sync := &DatascienceModelCustomMetadataArtifactContentDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataScienceClient()

	return tfresource.ReadResource(sync)
}

type DatascienceModelCustomMetadataArtifactContentDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_datascience.DataScienceClient
	Res    *oci_datascience.GetModelCustomMetadatumArtifactContentResponse
}

func (s *DatascienceModelCustomMetadataArtifactContentDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatascienceModelCustomMetadataArtifactContentDataSourceCrud) Get() error {
	request := oci_datascience.GetModelCustomMetadatumArtifactContentRequest{}

	if metadatumKeyName, ok := s.D.GetOkExists("metadatum_key_name"); ok {
		tmp := metadatumKeyName.(string)
		request.MetadatumKeyName = &tmp
	}

	if modelId, ok := s.D.GetOkExists("model_id"); ok {
		tmp := modelId.(string)
		request.ModelId = &tmp
	}

	if range_, ok := s.D.GetOkExists("range"); ok {
		tmp := range_.(string)
		request.Range = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "datascience")

	response, err := s.Client.GetModelCustomMetadatumArtifactContent(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatascienceModelCustomMetadataArtifactContentDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatascienceModelCustomMetadataArtifactContentDataSource-", DatascienceModelCustomMetadataArtifactContentDataSource(), s.D))

	return nil
}
