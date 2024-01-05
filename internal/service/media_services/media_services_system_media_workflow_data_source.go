// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package media_services

import (
	"context"
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_media_services "github.com/oracle/oci-go-sdk/v65/mediaservices"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func MediaServicesSystemMediaWorkflowDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularMediaServicesSystemMediaWorkflow,
		Schema: map[string]*schema.Schema{
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			// Computed
			"items": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"parameters": {
							Type:             schema.TypeString,
							Optional:         true,
							Computed:         true,
							Elem:             schema.TypeString,
							DiffSuppressFunc: tfresource.JsonStringDiffSuppressFunction,
						},
						"tasks": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"enable_parameter_reference": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"enable_when_referenced_parameter_equals": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem:     schema.TypeString,
									},
									"key": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"parameters": {
										Type:             schema.TypeString,
										Required:         true,
										DiffSuppressFunc: tfresource.JsonStringDiffSuppressFunction,
									},
									"prerequisites": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"version": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func readSingularMediaServicesSystemMediaWorkflow(d *schema.ResourceData, m interface{}) error {
	sync := &MediaServicesSystemMediaWorkflowDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).MediaServicesClient()

	return tfresource.ReadResource(sync)
}

type MediaServicesSystemMediaWorkflowDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_media_services.MediaServicesClient
	Res    *oci_media_services.ListSystemMediaWorkflowsResponse
}

func (s *MediaServicesSystemMediaWorkflowDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *MediaServicesSystemMediaWorkflowDataSourceCrud) Get() error {
	request := oci_media_services.ListSystemMediaWorkflowsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "media_services")

	response, err := s.Client.ListSystemMediaWorkflows(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *MediaServicesSystemMediaWorkflowDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("MediaServicesSystemMediaWorkflowDataSource-", MediaServicesSystemMediaWorkflowDataSource(), s.D))

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, SystemMediaWorkflowToMap(item))
	}
	s.D.Set("items", items)

	return nil
}

func SystemMediaWorkflowToMap(obj oci_media_services.SystemMediaWorkflow) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	jsonStr, err := json.Marshal(obj.Parameters)
	if err == nil {
		result["parameters"] = string(jsonStr)
	}

	tasks := []interface{}{}
	for _, item := range obj.Tasks {
		tasks = append(tasks, MediaWorkflowTaskToMap(item))
	}
	result["tasks"] = tasks

	return result
}
