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

// FunctionsFunctionsRuntimeVersionsDataSource defines the list Functions runtime versions lookup schema.
func FunctionsFunctionsRuntimeVersionsDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: readFunctionsFunctionsRuntimeVersionsWithContext,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"functions_runtime_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"functions_runtime_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"functions_runtime_version_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"is_current_version": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"language_version": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"os_version": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"functions_runtime_version_collection": {
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
									"functions_runtime_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"language_version": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"metadata": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"os_version": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"state": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"supported_architectures": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
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

// readFunctionsFunctionsRuntimeVersionsWithContext wires Terraform reads to the runtime versions CRUD.
func readFunctionsFunctionsRuntimeVersionsWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &FunctionsFunctionsRuntimeVersionsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FunctionsManagementClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type FunctionsFunctionsRuntimeVersionsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_functions.FunctionsManagementClient
	Res    *oci_functions.ListFunctionsRuntimeVersionsResponse
}

func (s *FunctionsFunctionsRuntimeVersionsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

// GetWithContext calls OCI to list Functions runtime versions and follows pagination.
func (s *FunctionsFunctionsRuntimeVersionsDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_functions.ListFunctionsRuntimeVersionsRequest{}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if functionsRuntimeId, ok := s.D.GetOkExists("functions_runtime_id"); ok {
		tmp := functionsRuntimeId.(string)
		request.FunctionsRuntimeId = &tmp
	}

	if functionsRuntimeName, ok := s.D.GetOkExists("functions_runtime_name"); ok {
		tmp := functionsRuntimeName.(string)
		request.FunctionsRuntimeName = &tmp
	}

	if functionsRuntimeVersionId, ok := s.D.GetOkExists("functions_runtime_version_id"); ok {
		tmp := functionsRuntimeVersionId.(string)
		request.FunctionsRuntimeVersionId = &tmp
	}

	if isCurrentVersion, ok := s.D.GetOkExists("is_current_version"); ok {
		tmp := isCurrentVersion.(bool)
		request.IsCurrentVersion = &tmp
	}

	if languageVersion, ok := s.D.GetOkExists("language_version"); ok {
		tmp := languageVersion.(string)
		request.LanguageVersion = &tmp
	}

	if osVersion, ok := s.D.GetOkExists("os_version"); ok {
		tmp := osVersion.(string)
		request.OsVersion = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_functions.FunctionsRuntimeVersionLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "functions")

	response, err := s.Client.ListFunctionsRuntimeVersions(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListFunctionsRuntimeVersions(ctx, request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

// SetData maps listed runtime versions into Terraform state and applies Terraform-side filters.
func (s *FunctionsFunctionsRuntimeVersionsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("FunctionsFunctionsRuntimeVersionsDataSource-", FunctionsFunctionsRuntimeVersionsDataSource(), s.D))
	resources := []map[string]interface{}{}
	functionsRuntimeVersion := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, FunctionsRuntimeVersionSummaryToMap(item))
	}
	functionsRuntimeVersion["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, FunctionsFunctionsRuntimeVersionsDataSource().Schema["functions_runtime_version_collection"].Elem.(*schema.Resource).Schema)
		functionsRuntimeVersion["items"] = items
	}

	resources = append(resources, functionsRuntimeVersion)
	if err := s.D.Set("functions_runtime_version_collection", resources); err != nil {
		return err
	}

	return nil
}

// FunctionsRuntimeVersionSummaryToMap converts an OCI runtime version summary into a Terraform item map.
func FunctionsRuntimeVersionSummaryToMap(obj oci_functions.FunctionsRuntimeVersionSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.FunctionsRuntimeId != nil {
		result["functions_runtime_id"] = string(*obj.FunctionsRuntimeId)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.LanguageVersion != nil {
		result["language_version"] = string(*obj.LanguageVersion)
	}

	if obj.Metadata != nil {
		result["metadata"] = string(*obj.Metadata)
	}

	if obj.OsVersion != nil {
		result["os_version"] = string(*obj.OsVersion)
	}

	result["state"] = string(obj.LifecycleState)

	result["supported_architectures"] = obj.SupportedArchitectures

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	return result
}
