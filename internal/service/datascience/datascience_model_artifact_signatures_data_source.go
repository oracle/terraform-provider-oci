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

func DatascienceModelArtifactSignaturesDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: readDatascienceModelArtifactSignaturesWithContext,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"digest": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"model_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"model_artifact_signatures": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(DatascienceModelArtifactSignatureResource()),
			},
		},
	}
}

func readDatascienceModelArtifactSignaturesWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DatascienceModelArtifactSignaturesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataScienceClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type DatascienceModelArtifactSignaturesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_datascience.DataScienceClient
	Res    *oci_datascience.ListModelArtifactSignaturesResponse
}

func (s *DatascienceModelArtifactSignaturesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatascienceModelArtifactSignaturesDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_datascience.ListModelArtifactSignaturesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if digest, ok := s.D.GetOkExists("digest"); ok {
		tmp := digest.(string)
		request.Digest = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if modelId, ok := s.D.GetOkExists("model_id"); ok {
		tmp := modelId.(string)
		request.ModelId = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_datascience.ListModelArtifactSignaturesLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "datascience")

	response, err := s.Client.ListModelArtifactSignatures(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListModelArtifactSignatures(ctx, request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatascienceModelArtifactSignaturesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatascienceModelArtifactSignaturesDataSource-", DatascienceModelArtifactSignaturesDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		modelArtifactSignature := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
			"model_id":       *r.ModelId,
		}

		if r.CreatedBy != nil {
			modelArtifactSignature["created_by"] = *r.CreatedBy
		}

		if r.DefinedTags != nil {
			modelArtifactSignature["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		if r.Digest != nil {
			modelArtifactSignature["digest"] = *r.Digest
		}

		if r.DisplayName != nil {
			modelArtifactSignature["display_name"] = *r.DisplayName
		}

		modelArtifactSignature["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			modelArtifactSignature["id"] = *r.Id
		}

		modelArtifactSignature["state"] = r.LifecycleState

		if r.SystemTags != nil {
			modelArtifactSignature["system_tags"] = tfresource.SystemTagsToMap(r.SystemTags)
		}

		if r.TimeCreated != nil {
			modelArtifactSignature["time_created"] = r.TimeCreated.String()
		}

		if r.TimeUpdated != nil {
			modelArtifactSignature["time_updated"] = r.TimeUpdated.String()
		}

		resources = append(resources, modelArtifactSignature)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, DatascienceModelArtifactSignaturesDataSource().Schema["model_artifact_signatures"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("model_artifact_signatures", resources); err != nil {
		return err
	}

	return nil
}
