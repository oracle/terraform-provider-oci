// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package functions

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_functions "github.com/oracle/oci-go-sdk/v65/functions"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

// FunctionsFunctionsRuntimesDataSource defines the list Functions runtimes lookup schema.
func FunctionsFunctionsRuntimesDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: readFunctionsFunctionsRuntimesWithContext,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"functions_runtime_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"language": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"name_contains": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"name_starts_with": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"os": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"functions_runtime_collection": {
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
									"current_functions_runtime_version_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"defined_tags": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem:     schema.TypeString,
									},
									"freeform_tags": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem:     schema.TypeString,
									},
									"id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"language": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"metadata": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"os": {
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
									"time_created": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_decommissioned": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_deprecated": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_updated": {
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

// readFunctionsFunctionsRuntimesWithContext wires Terraform reads to the runtimes data source CRUD.
func readFunctionsFunctionsRuntimesWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &FunctionsFunctionsRuntimesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FunctionsManagementClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type FunctionsFunctionsRuntimesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_functions.FunctionsManagementClient
	Res    *oci_functions.ListFunctionsRuntimesResponse
}

func (s *FunctionsFunctionsRuntimesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

// GetWithContext calls OCI to list Functions runtimes and follows pagination.
func (s *FunctionsFunctionsRuntimesDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_functions.ListFunctionsRuntimesRequest{}

	if functionsRuntimeId, ok := s.D.GetOkExists("functions_runtime_id"); ok {
		tmp := functionsRuntimeId.(string)
		request.FunctionsRuntimeId = &tmp
	}

	if language, ok := s.D.GetOkExists("language"); ok {
		tmp := language.(string)
		request.Language = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if nameContains, ok := s.D.GetOkExists("name_contains"); ok {
		tmp := nameContains.(string)
		request.NameContains = &tmp
	}

	if nameStartsWith, ok := s.D.GetOkExists("name_starts_with"); ok {
		tmp := nameStartsWith.(string)
		request.NameStartsWith = &tmp
	}

	if os, ok := s.D.GetOkExists("os"); ok {
		tmp := os.(string)
		request.Os = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_functions.FunctionsRuntimeLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "functions")

	response, err := s.Client.ListFunctionsRuntimes(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListFunctionsRuntimes(ctx, request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

// SetData maps listed runtimes into Terraform state and applies Terraform-side filters.
func (s *FunctionsFunctionsRuntimesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("FunctionsFunctionsRuntimesDataSource-", FunctionsFunctionsRuntimesDataSource(), s.D))
	resources := []map[string]interface{}{}
	functionsRuntime := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, FunctionsRuntimeSummaryToMap(item))
	}
	functionsRuntime["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, FunctionsFunctionsRuntimesDataSource().Schema["functions_runtime_collection"].Elem.(*schema.Resource).Schema)
		functionsRuntime["items"] = items
	}

	resources = append(resources, functionsRuntime)
	if err := s.D.Set("functions_runtime_collection", resources); err != nil {
		return err
	}

	return nil
}

// FunctionsRuntimeSummaryToMap converts an OCI runtime summary into a Terraform item map.
func FunctionsRuntimeSummaryToMap(obj oci_functions.FunctionsRuntimeSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CurrentFunctionsRuntimeVersionId != nil {
		result["current_functions_runtime_version_id"] = string(*obj.CurrentFunctionsRuntimeVersionId)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.Language != nil {
		result["language"] = string(*obj.Language)
	}

	if obj.Metadata != nil {
		result["metadata"] = string(*obj.Metadata)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.Os != nil {
		result["os"] = string(*obj.Os)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeDecommissioned != nil {
		result["time_decommissioned"] = obj.TimeDecommissioned.String()
	}

	if obj.TimeDeprecated != nil {
		result["time_deprecated"] = obj.TimeDeprecated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	return result
}
