// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package media_services

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_media_services "github.com/oracle/oci-go-sdk/v65/mediaservices"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func MediaServicesMediaWorkflowJobFactsDataSource() *schema.Resource {
	return &schema.Resource{
		DeprecationMessage: "This data source has been deprecated and is no longer supported.",
		Read:               readMediaServicesMediaWorkflowJobFacts,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"key": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"media_workflow_job_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"media_workflow_job_fact_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"detail": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"key": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"media_workflow_job_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"type": {
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

func readMediaServicesMediaWorkflowJobFacts(d *schema.ResourceData, m interface{}) error {
	sync := &MediaServicesMediaWorkflowJobFactsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).MediaServicesClient()

	return tfresource.ReadResource(sync)
}

type MediaServicesMediaWorkflowJobFactsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_media_services.MediaServicesClient
	//Res    *oci_media_services.ListMediaWorkflowJobFactsResponse
}

func (s *MediaServicesMediaWorkflowJobFactsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *MediaServicesMediaWorkflowJobFactsDataSourceCrud) Get() error {
	//request := oci_media_services.ListMediaWorkflowJobFactsRequest{}
	//
	//if key, ok := s.D.GetOkExists("key"); ok {
	//	intValue := 0
	//	_, _ = fmt.Sscan(key.(string), &intValue)
	//	request.Key = &intValue
	//}
	//
	//if mediaWorkflowJobId, ok := s.D.GetOkExists("media_workflow_job_id"); ok {
	//	tmp := mediaWorkflowJobId.(string)
	//	request.MediaWorkflowJobId = &tmp
	//}
	//
	//if type_, ok := s.D.GetOkExists("type"); ok {
	//	request.Type = oci_media_services.ListMediaWorkflowJobFactsTypeEnum(type_.(string))
	//}
	//
	//request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "media_services")
	//request.RequestMetadata.RetryPolicy.MaximumNumberAttempts = 1
	//
	//response, err := s.Client.ListMediaWorkflowJobFacts(context.Background(), request)
	//if err != nil && response.RawResponse.StatusCode != 404 {
	//	return err
	//}
	//
	//s.Res = &response
	//request.Page = s.Res.OpcNextPage
	//
	//for request.Page != nil {
	//	listResponse, err := s.Client.ListMediaWorkflowJobFacts(context.Background(), request)
	//	if err != nil {
	//		return err
	//	}
	//
	//	s.Res.Items = append(s.Res.Items, listResponse.Items...)
	//	request.Page = listResponse.OpcNextPage
	//}

	return nil
}

func (s *MediaServicesMediaWorkflowJobFactsDataSourceCrud) SetData() error {
	//if s.Res == nil {
	//	return nil
	//}
	//
	//s.D.SetId(tfresource.GenerateDataSourceHashID("MediaServicesMediaWorkflowJobFactsDataSource-", MediaServicesMediaWorkflowJobFactsDataSource(), s.D))
	//resources := []map[string]interface{}{}
	//mediaWorkflowJobFact := map[string]interface{}{}
	//
	//items := []interface{}{}
	//for _, item := range s.Res.Items {
	//	items = append(items, MediaWorkflowJobFactSummaryToMap(item))
	//}
	//mediaWorkflowJobFact["items"] = items
	//
	//if f, fOk := s.D.GetOkExists("filter"); fOk {
	//	items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, MediaServicesMediaWorkflowJobFactsDataSource().Schema["media_workflow_job_fact_collection"].Elem.(*schema.Resource).Schema)
	//	mediaWorkflowJobFact["items"] = items
	//}
	//
	//resources = append(resources, mediaWorkflowJobFact)
	//if err := s.D.Set("media_workflow_job_fact_collection", resources); err != nil {
	//	return err
	//}

	return nil
}

//func MediaWorkflowJobFactSummaryToMap(obj oci_media_services.MediaWorkflowJobFactSummary) map[string]interface{} {
//	result := map[string]interface{}{}
//
//	if obj.Key != nil {
//		result["key"] = strconv.FormatInt(*obj.Key, 10)
//	}
//
//	if obj.MediaWorkflowJobId != nil {
//		result["media_workflow_job_id"] = string(*obj.MediaWorkflowJobId)
//	}
//
//	if obj.Name != nil {
//		result["name"] = string(*obj.Name)
//	}
//
//	if obj.Type != nil {
//		result["type"] = string(*obj.Type)
//	}
//
//	return result
//}
