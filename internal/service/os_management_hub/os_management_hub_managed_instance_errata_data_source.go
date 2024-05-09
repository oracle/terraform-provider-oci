// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package os_management_hub

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_os_management_hub "github.com/oracle/oci-go-sdk/v65/osmanagementhub"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func OsManagementHubManagedInstanceErrataDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readOsManagementHubManagedInstanceErrata,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"classification_type": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"managed_instance_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"name_contains": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"managed_instance_erratum_summary_collection": {
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
									"advisory_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"packages": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"architecture": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"display_name": {
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
												"version": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"related_cves": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"synopsis": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_issued": {
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

func readOsManagementHubManagedInstanceErrata(d *schema.ResourceData, m interface{}) error {
	sync := &OsManagementHubManagedInstanceErrataDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ManagedInstanceClient()

	return tfresource.ReadResource(sync)
}

type OsManagementHubManagedInstanceErrataDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_os_management_hub.ManagedInstanceClient
	Res    *oci_os_management_hub.ListManagedInstanceErrataResponse
}

func (s *OsManagementHubManagedInstanceErrataDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OsManagementHubManagedInstanceErrataDataSourceCrud) Get() error {
	request := oci_os_management_hub.ListManagedInstanceErrataRequest{}

	if classificationType, ok := s.D.GetOkExists("classification_type"); ok {
		interfaces := classificationType.([]interface{})
		tmp := make([]oci_os_management_hub.ClassificationTypesEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_os_management_hub.ClassificationTypesEnum(interfaces[i].(string))
			}
		}
		if len(tmp) != 0 || s.D.HasChange("classification_type") {
			request.ClassificationType = tmp
		}
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if managedInstanceId, ok := s.D.GetOkExists("managed_instance_id"); ok {
		tmp := managedInstanceId.(string)
		request.ManagedInstanceId = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		interfaces := name.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("name") {
			request.Name = tmp
		}
	}

	if nameContains, ok := s.D.GetOkExists("name_contains"); ok {
		tmp := nameContains.(string)
		request.NameContains = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "os_management_hub")

	response, err := s.Client.ListManagedInstanceErrata(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListManagedInstanceErrata(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *OsManagementHubManagedInstanceErrataDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("OsManagementHubManagedInstanceErrataDataSource-", OsManagementHubManagedInstanceErrataDataSource(), s.D))
	resources := []map[string]interface{}{}
	managedInstanceErrata := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, ManagedInstanceErratumSummaryToMap(item))
	}
	managedInstanceErrata["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, OsManagementHubManagedInstanceErrataDataSource().Schema["managed_instance_erratum_summary_collection"].Elem.(*schema.Resource).Schema)
		managedInstanceErrata["items"] = items
	}

	resources = append(resources, managedInstanceErrata)
	if err := s.D.Set("managed_instance_erratum_summary_collection", resources); err != nil {
		return err
	}

	return nil
}

func ManagedInstanceErratumSummaryToMap(obj oci_os_management_hub.ManagedInstanceErratumSummary) map[string]interface{} {
	result := map[string]interface{}{}

	result["advisory_type"] = string(obj.AdvisoryType)

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	packages := []interface{}{}
	for _, item := range obj.Packages {
		packages = append(packages, PackageNameSummaryToMap(item))
	}
	result["packages"] = packages

	result["related_cves"] = obj.RelatedCves

	if obj.Synopsis != nil {
		result["synopsis"] = string(*obj.Synopsis)
	}

	if obj.TimeIssued != nil {
		result["time_issued"] = obj.TimeIssued.String()
	}

	return result
}

func PackageNameSummaryToMap(obj oci_os_management_hub.PackageNameSummary) map[string]interface{} {
	result := map[string]interface{}{}

	result["architecture"] = string(obj.Architecture)

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.Type != nil {
		result["type"] = string(*obj.Type)
	}

	if obj.Version != nil {
		result["version"] = string(*obj.Version)
	}

	return result
}
