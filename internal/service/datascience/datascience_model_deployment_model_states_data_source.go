// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package datascience

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_datascience "github.com/oracle/oci-go-sdk/v65/datascience"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatascienceModelDeploymentModelStatesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatascienceModelDeploymentModelStates,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"inference_key": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"model_deployment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"model_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"project_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"model_deployment_model_states": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"defined_tags": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem:     schema.TypeString,
						},
						"display_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"freeform_tags": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem:     schema.TypeString,
						},
						"inference_key": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"model_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"project_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"state": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"system_tags": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem:     schema.TypeString,
						},
					},
				},
			},
		},
	}
}

func readDatascienceModelDeploymentModelStates(d *schema.ResourceData, m interface{}) error {
	sync := &DatascienceModelDeploymentModelStatesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataScienceClient()

	return tfresource.ReadResource(sync)
}

type DatascienceModelDeploymentModelStatesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_datascience.DataScienceClient
	Res    *oci_datascience.ListModelDeploymentModelStatesResponse
}

func (s *DatascienceModelDeploymentModelStatesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatascienceModelDeploymentModelStatesDataSourceCrud) Get() error {
	request := oci_datascience.ListModelDeploymentModelStatesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if inferenceKey, ok := s.D.GetOkExists("inference_key"); ok {
		tmp := inferenceKey.(string)
		request.InferenceKey = &tmp
	}

	if modelDeploymentId, ok := s.D.GetOkExists("model_id"); ok {
		tmp := modelDeploymentId.(string)
		request.ModelDeploymentId = &tmp
	}

	if modelId, ok := s.D.GetOkExists("model_id"); ok {
		tmp := modelId.(string)
		request.ModelId = &tmp
	}

	if projectId, ok := s.D.GetOkExists("project_id"); ok {
		tmp := projectId.(string)
		request.ProjectId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "datascience")

	response, err := s.Client.ListModelDeploymentModelStates(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListModelDeploymentModelStates(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatascienceModelDeploymentModelStatesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatascienceModelDeploymentModelStatesDataSource-", DatascienceModelDeploymentModelStatesDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		modelDeploymentModelState := map[string]interface{}{}

		if r.DefinedTags != nil {
			modelDeploymentModelState["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		if r.DisplayName != nil {
			modelDeploymentModelState["display_name"] = *r.DisplayName
		}

		modelDeploymentModelState["freeform_tags"] = r.FreeformTags

		if r.InferenceKey != nil {
			modelDeploymentModelState["inference_key"] = *r.InferenceKey
		}

		if r.ModelId != nil {
			modelDeploymentModelState["model_id"] = *r.ModelId
		}

		if r.ProjectId != nil {
			modelDeploymentModelState["project_id"] = *r.ProjectId
		}

		modelDeploymentModelState["state"] = r.State

		if r.SystemTags != nil {
			modelDeploymentModelState["system_tags"] = tfresource.SystemTagsToMap(r.SystemTags)
		}

		resources = append(resources, modelDeploymentModelState)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, DatascienceModelDeploymentModelStatesDataSource().Schema["model_deployment_model_states"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("model_deployment_model_states", resources); err != nil {
		return err
	}

	return nil
}
