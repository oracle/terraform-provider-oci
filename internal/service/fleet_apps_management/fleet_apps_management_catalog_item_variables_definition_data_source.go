// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package fleet_apps_management

import (
	"context"
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_fleet_apps_management "github.com/oracle/oci-go-sdk/v65/fleetappsmanagement"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func FleetAppsManagementCatalogItemVariablesDefinitionDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: readSingularFleetAppsManagementCatalogItemVariablesDefinitionWithContext,
		Schema: map[string]*schema.Schema{
			"catalog_item_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
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
			"schema_document": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"can_allow_view_state": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"groupings": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"array": {
										// Type:     schema.TypeMap,
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"title": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"variables": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"visible": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
								},
							},
						},
						"informational_text": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"instructions": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"locale": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"logo_url": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"output_groups": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"outputs": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"title": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"outputs": {
							// Type:     schema.TypeMap,
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
									"display_text": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"format": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"is_sensitive": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"title": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"value": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"visible": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"package_version": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"primary_output_button": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"schema_version": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"source": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"reference": {
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
						"stack_description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"title": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"troubleshooting": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"variable_groups": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"title": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"variables": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"visible": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"variables": {
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
			"system_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
		},
	}
}

func readSingularFleetAppsManagementCatalogItemVariablesDefinitionWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &FleetAppsManagementCatalogItemVariablesDefinitionDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FleetAppsManagementCatalogClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type FleetAppsManagementCatalogItemVariablesDefinitionDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_fleet_apps_management.FleetAppsManagementCatalogClient
	Res    *oci_fleet_apps_management.GetCatalogItemVariablesDefinitionResponse
}

func (s *FleetAppsManagementCatalogItemVariablesDefinitionDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *FleetAppsManagementCatalogItemVariablesDefinitionDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_fleet_apps_management.GetCatalogItemVariablesDefinitionRequest{}

	if catalogItemId, ok := s.D.GetOkExists("catalog_item_id"); ok {
		tmp := catalogItemId.(string)
		request.CatalogItemId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "fleet_apps_management")

	response, err := s.Client.GetCatalogItemVariablesDefinition(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *FleetAppsManagementCatalogItemVariablesDefinitionDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("FleetAppsManagementCatalogItemVariablesDefinitionDataSource-", FleetAppsManagementCatalogItemVariablesDefinitionDataSource(), s.D))

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.SchemaDocument != nil {
		s.D.Set("schema_document", []interface{}{SchemaDocumentToMap(s.Res.SchemaDocument)})
	} else {
		s.D.Set("schema_document", nil)
	}

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	return nil
}

func OutputGroupToMap(obj oci_fleet_apps_management.OutputGroup) map[string]interface{} {
	result := map[string]interface{}{}

	result["outputs"] = obj.Outputs

	if obj.Title != nil {
		result["title"] = string(*obj.Title)
	}

	return result
}

func SchemaDocumentToMap(obj *oci_fleet_apps_management.SchemaDocument) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CanAllowViewState != nil {
		result["can_allow_view_state"] = bool(*obj.CanAllowViewState)
	}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.Groupings != nil {
		result["groupings"] = []interface{}{VariableGroupsToMap(obj.Groupings)}
	}

	if obj.InformationalText != nil {
		result["informational_text"] = string(*obj.InformationalText)
	}

	if obj.Instructions != nil {
		result["instructions"] = string(*obj.Instructions)
	}

	result["locale"] = string(obj.Locale)

	if obj.LogoUrl != nil {
		result["logo_url"] = string(*obj.LogoUrl)
	}

	outputGroups := []interface{}{}
	for _, item := range obj.OutputGroups {
		outputGroups = append(outputGroups, OutputGroupToMap(item))
	}
	result["output_groups"] = outputGroups

	result["outputs"] = obj.Outputs

	if obj.PackageVersion != nil {
		result["package_version"] = string(*obj.PackageVersion)
	}

	if obj.PrimaryOutputButton != nil {
		result["primary_output_button"] = string(*obj.PrimaryOutputButton)
	}

	result["schema_version"] = string(obj.SchemaVersion)

	if obj.Source != nil {
		result["source"] = []interface{}{StackSourceToMap(obj.Source)}
	}

	if obj.StackDescription != nil {
		result["stack_description"] = string(*obj.StackDescription)
	}

	if obj.Title != nil {
		result["title"] = string(*obj.Title)
	}

	if obj.Troubleshooting != nil {
		result["troubleshooting"] = string(*obj.Troubleshooting)
	}

	variableGroups := []interface{}{}
	for _, item := range obj.VariableGroups {
		variableGroups = append(variableGroups, VariableGroupToMap(item))
	}
	result["variable_groups"] = variableGroups

	if obj.Variables != nil {
		tmp, _ := json.Marshal(obj.Variables)
		result["variables"] = string(tmp)
	}

	if obj.Version != nil {
		result["version"] = string(*obj.Version)
	}

	return result
}

func StackSourceToMap(obj *oci_fleet_apps_management.StackSource) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Reference != nil {
		result["reference"] = string(*obj.Reference)
	}

	result["type"] = string(obj.Type)

	return result
}

func VariableGroupToMap(obj oci_fleet_apps_management.VariableGroup) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Title != nil {
		result["title"] = string(*obj.Title)
	}

	result["variables"] = obj.Variables

	if obj.Visible != nil {
		result["visible"] = string(*obj.Visible)
	}

	return result
}

func VariableGroupsToMap(obj *oci_fleet_apps_management.VariableGroups) map[string]interface{} {
	result := map[string]interface{}{}

	result["array"] = obj.Array

	return result
}

// func (s *FleetAppsManagementCatalogItemVariablesDefinitionDataSourceCrud) mapToobject(fieldKeyFormat string) (oci_fleet_apps_management.Object, error) {
// 	result := oci_fleet_apps_management.Object{}

// 	return result, nil
// }
