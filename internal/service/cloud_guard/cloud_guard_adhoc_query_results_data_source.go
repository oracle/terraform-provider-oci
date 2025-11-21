// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

//resource not exposed to user through Terraform, but generated.
//Hence TF team suggested to keep the file commented as codeGen patch build fails if file not present

package cloud_guard

/*import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_cloud_guard "github.com/oracle/oci-go-sdk/v65/cloudguard"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func CloudGuardAdhocQueryResultsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readCloudGuardAdhocQueryResults,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"adhoc_query_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"adhoc_query_result_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"items": {
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
									"error_message": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"freeform_tags": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem:     schema.TypeString,
									},
									"host_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"region": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"result": {
										Type:     schema.TypeList,
										Computed: true,
										Elem:     schema.TypeString,
									},
									"result_count": {
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
									"time_submitted": {
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

func readCloudGuardAdhocQueryResults(d *schema.ResourceData, m interface{}) error {
	sync := &CloudGuardAdhocQueryResultsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).CloudGuardClient()

	return tfresource.ReadResource(sync)
}

type CloudGuardAdhocQueryResultsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_cloud_guard.CloudGuardClient
	Res    *oci_cloud_guard.ListAdhocQueryResultsResponse
}

func (s *CloudGuardAdhocQueryResultsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CloudGuardAdhocQueryResultsDataSourceCrud) Get() error {
	request := oci_cloud_guard.ListAdhocQueryResultsRequest{}

	if adhocQueryId, ok := s.D.GetOkExists("adhoc_query_id"); ok {
		tmp := adhocQueryId.(string)
		request.AdhocQueryId = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "cloud_guard")

	response, err := s.Client.ListAdhocQueryResults(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListAdhocQueryResults(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *CloudGuardAdhocQueryResultsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("CloudGuardAdhocQueryResultsDataSource-", CloudGuardAdhocQueryResultsDataSource(), s.D))
	resources := []map[string]interface{}{}
	adhocQueryResult := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, AdhocQueryResultSummaryToMap(item))
	}
	adhocQueryResult["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, CloudGuardAdhocQueryResultsDataSource().Schema["adhoc_query_result_collection"].Elem.(*schema.Resource).Schema)
		adhocQueryResult["items"] = items
	}

	resources = append(resources, adhocQueryResult)
	if err := s.D.Set("adhoc_query_result_collection", resources); err != nil {
		return err
	}

	return nil
}

func AdhocQueryResultSummaryToMap(obj oci_cloud_guard.AdhocQueryResultSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.ErrorMessage != nil {
		result["error_message"] = string(*obj.ErrorMessage)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.HostId != nil {
		result["host_id"] = string(*obj.HostId)
	}

	if obj.Region != nil {
		result["region"] = string(*obj.Region)
	}

	result := []interface{}{}
	for _, item := range obj.Result {
		result = append(result, objectToMap(item))
	}
	result["result"] = result

	if obj.ResultCount != nil {
		result["result_count"] = strconv.FormatInt(*obj.ResultCount, 10)
	}

	result["state"] = string(obj.State)

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TimeSubmitted != nil {
		result["time_submitted"] = obj.TimeSubmitted.String()
	}

	return result
}

func objectToMap(obj oci_cloud_guard.Object) map[string]interface{} {
	result := map[string]interface{}{}

	return result
}*/
