// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package os_management_hub

import (
	"context"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_os_management_hub "github.com/oracle/oci-go-sdk/v65/osmanagementhub"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func OsManagementHubErrataDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readOsManagementHubErrata,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"advisory_severity": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"advisory_type": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"classification_type": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"compartment_id": {
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
			"os_family": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"time_issue_date_end": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"time_issue_date_start": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"erratum_collection": {
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
									"advisory_severity": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"advisory_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"classification_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"description": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"from": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"os_families": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
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
												"checksum": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"checksum_type": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"display_name": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"is_latest": {
													Type:     schema.TypeBool,
													Computed: true,
												},
												"name": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"os_families": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"software_sources": {
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
															"id": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"is_mandatory_for_autonomous_linux": {
																Type:     schema.TypeBool,
																Computed: true,
															},
															"software_source_type": {
																Type:     schema.TypeString,
																Computed: true,
															},
														},
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
									"references": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"related_cves": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"repositories": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"solution": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"synopsis": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_issued": {
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

func readOsManagementHubErrata(d *schema.ResourceData, m interface{}) error {
	sync := &OsManagementHubErrataDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).SoftwareSourceClient()

	return tfresource.ReadResource(sync)
}

type OsManagementHubErrataDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_os_management_hub.SoftwareSourceClient
	Res    *oci_os_management_hub.ListErrataResponse
}

func (s *OsManagementHubErrataDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OsManagementHubErrataDataSourceCrud) Get() error {
	request := oci_os_management_hub.ListErrataRequest{}

	if advisorySeverity, ok := s.D.GetOkExists("advisory_severity"); ok {
		interfaces := advisorySeverity.([]interface{})
		tmp := make([]oci_os_management_hub.AdvisorySeverityEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_os_management_hub.AdvisorySeverityEnum(interfaces[i].(string))
			}
		}
		request.AdvisorySeverity = tmp
	}

	if advisoryType, ok := s.D.GetOkExists("advisory_type"); ok {
		interfaces := advisoryType.([]interface{})
		tmp := make([]oci_os_management_hub.AdvisoryTypesEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_os_management_hub.AdvisoryTypesEnum(interfaces[i].(string))
			}
		}
		request.AdvisoryType = tmp
	}

	if classificationType, ok := s.D.GetOkExists("advisory_type"); ok {
		interfaces := classificationType.([]interface{})
		tmp := make([]oci_os_management_hub.ClassificationTypesEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_os_management_hub.ClassificationTypesEnum(interfaces[i].(string))
			}
		}
		request.ClassificationType = tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		interfaces := name.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		request.Name = tmp
	}

	if nameContains, ok := s.D.GetOkExists("name_contains"); ok {
		tmp := nameContains.(string)
		request.NameContains = &tmp
	}

	if osFamily, ok := s.D.GetOkExists("os_family"); ok {
		request.OsFamily = oci_os_management_hub.ListErrataOsFamilyEnum(osFamily.(string))
	}

	if timeIssueDateEnd, ok := s.D.GetOkExists("time_issue_date_end"); ok {
		tmp, err := time.Parse(time.RFC3339, timeIssueDateEnd.(string))
		if err != nil {
			return err
		}
		request.TimeIssueDateEnd = &oci_common.SDKTime{Time: tmp}
	}

	if timeIssueDateStart, ok := s.D.GetOkExists("time_issue_date_start"); ok {
		tmp, err := time.Parse(time.RFC3339, timeIssueDateStart.(string))
		if err != nil {
			return err
		}
		request.TimeIssueDateStart = &oci_common.SDKTime{Time: tmp}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "os_management_hub")

	response, err := s.Client.ListErrata(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListErrata(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *OsManagementHubErrataDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("OsManagementHubErrataDataSource-", OsManagementHubErrataDataSource(), s.D))
	resources := []map[string]interface{}{}
	erratum := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, ErratumSummaryToMap(item))
	}
	erratum["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, OsManagementHubErrataDataSource().Schema["erratum_collection"].Elem.(*schema.Resource).Schema)
		erratum["items"] = items
	}

	resources = append(resources, erratum)
	if err := s.D.Set("erratum_collection", resources); err != nil {
		return err
	}

	return nil
}

func ErratumSummaryToMap(obj oci_os_management_hub.ErratumSummary) map[string]interface{} {
	result := map[string]interface{}{}

	result["advisory_severity"] = string(obj.AdvisorySeverity)

	result["advisory_type"] = string(obj.AdvisoryType)

	result["classification_type"] = string(obj.ClassificationType)

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	result["os_families"] = obj.OsFamilies

	result["related_cves"] = obj.RelatedCves

	if obj.Synopsis != nil {
		result["synopsis"] = string(*obj.Synopsis)
	}

	if obj.TimeIssued != nil {
		result["time_issued"] = obj.TimeIssued.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	return result
}
