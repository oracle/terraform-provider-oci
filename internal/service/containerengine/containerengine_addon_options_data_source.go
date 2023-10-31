// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package containerengine

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_containerengine "github.com/oracle/oci-go-sdk/v65/containerengine"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func ContainerengineAddonOptionsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readContainerengineAddonOptions,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"addon_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"kubernetes_version": {
				Type:     schema.TypeString,
				Required: true,
			},
			"addon_options": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"addon_group": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"addon_schema_version": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"defined_tags": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem:     schema.TypeString,
						},
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"freeform_tags": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem:     schema.TypeString,
						},
						"is_essential": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"name": {
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
						"versions": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"configurations": {
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
												"display_name": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"is_required": {
													Type:     schema.TypeBool,
													Computed: true,
												},
												"key": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"value": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"description": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"kubernetes_version_filters": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"exact_kubernetes_versions": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"maximum_version": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"minimal_version": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"status": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"version_number": {
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

func readContainerengineAddonOptions(d *schema.ResourceData, m interface{}) error {
	sync := &ContainerengineAddonOptionsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ContainerEngineClient()

	return tfresource.ReadResource(sync)
}

type ContainerengineAddonOptionsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_containerengine.ContainerEngineClient
	Res    *oci_containerengine.ListAddonOptionsResponse
}

func (s *ContainerengineAddonOptionsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ContainerengineAddonOptionsDataSourceCrud) Get() error {
	request := oci_containerengine.ListAddonOptionsRequest{}

	if addonName, ok := s.D.GetOkExists("addon_name"); ok {
		tmp := addonName.(string)
		request.AddonName = &tmp
	}

	if kubernetesVersion, ok := s.D.GetOkExists("kubernetes_version"); ok {
		tmp := kubernetesVersion.(string)
		request.KubernetesVersion = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "containerengine")

	response, err := s.Client.ListAddonOptions(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListAddonOptions(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *ContainerengineAddonOptionsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("ContainerengineAddonOptionsDataSource-", ContainerengineAddonOptionsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		addonOption := map[string]interface{}{}

		if r.AddonGroup != nil {
			addonOption["addon_group"] = *r.AddonGroup
		}

		if r.AddonSchemaVersion != nil {
			addonOption["addon_schema_version"] = *r.AddonSchemaVersion
		}

		if r.DefinedTags != nil {
			addonOption["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		if r.Description != nil {
			addonOption["description"] = *r.Description
		}

		addonOption["freeform_tags"] = r.FreeformTags

		if r.IsEssential != nil {
			addonOption["is_essential"] = *r.IsEssential
		}

		if r.Name != nil {
			addonOption["name"] = *r.Name
		}

		addonOption["state"] = r.LifecycleState

		if r.SystemTags != nil {
			addonOption["system_tags"] = tfresource.SystemTagsToMap(r.SystemTags)
		}

		if r.TimeCreated != nil {
			addonOption["time_created"] = r.TimeCreated.String()
		}

		versions := []interface{}{}
		for _, item := range r.Versions {
			versions = append(versions, AddonVersionsToMap(item))
		}
		addonOption["versions"] = versions

		resources = append(resources, addonOption)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, ContainerengineAddonOptionsDataSource().Schema["addon_options"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("addon_options", resources); err != nil {
		return err
	}

	return nil
}

func AddonVersionConfigurationToMap(obj oci_containerengine.AddonVersionConfiguration) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.IsRequired != nil {
		result["is_required"] = bool(*obj.IsRequired)
	}

	if obj.Key != nil {
		result["key"] = string(*obj.Key)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func AddonVersionsToMap(obj oci_containerengine.AddonVersions) map[string]interface{} {
	result := map[string]interface{}{}

	configurations := []interface{}{}
	for _, item := range obj.Configurations {
		configurations = append(configurations, AddonVersionConfigurationToMap(item))
	}
	result["configurations"] = configurations

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.KubernetesVersionFilters != nil {
		result["kubernetes_version_filters"] = []interface{}{KubernetesVersionsFiltersToMap(obj.KubernetesVersionFilters)}
	}

	result["status"] = string(obj.Status)

	if obj.VersionNumber != nil {
		result["version_number"] = string(*obj.VersionNumber)
	}

	return result
}

func KubernetesVersionsFiltersToMap(obj *oci_containerengine.KubernetesVersionsFilters) map[string]interface{} {
	result := map[string]interface{}{}

	result["exact_kubernetes_versions"] = obj.ExactKubernetesVersions

	if obj.MaximumVersion != nil {
		result["maximum_version"] = string(*obj.MaximumVersion)
	}

	if obj.MinimalVersion != nil {
		result["minimal_version"] = string(*obj.MinimalVersion)
	}

	return result
}
