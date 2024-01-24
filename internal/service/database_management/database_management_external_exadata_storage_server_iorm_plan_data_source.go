// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database_management

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_database_management "github.com/oracle/oci-go-sdk/v65/databasemanagement"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatabaseManagementExternalExadataStorageServerIormPlanDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularDatabaseManagementExternalExadataStorageServerIormPlan,
		Schema: map[string]*schema.Schema{
			"external_exadata_storage_server_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"db_plan": {
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
									"allocation": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"limit": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"asm_cluster": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"flash_cache_limit": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"flash_cache_min": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"flash_cache_size": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"is_flash_cache_on": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"is_flash_log_on": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"is_pmem_cache_on": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"is_pmem_log_on": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"level": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"pmem_cache_limit": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"pmem_cache_min": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"pmem_cache_size": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"role": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"share": {
										Type:     schema.TypeInt,
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
			"plan_objective": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"plan_status": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularDatabaseManagementExternalExadataStorageServerIormPlan(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementExternalExadataStorageServerIormPlanDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()

	return tfresource.ReadResource(sync)
}

type DatabaseManagementExternalExadataStorageServerIormPlanDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_management.DbManagementClient
	Res    *oci_database_management.GetIormPlanResponse
}

func (s *DatabaseManagementExternalExadataStorageServerIormPlanDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseManagementExternalExadataStorageServerIormPlanDataSourceCrud) Get() error {
	request := oci_database_management.GetIormPlanRequest{}

	if externalExadataStorageServerId, ok := s.D.GetOkExists("external_exadata_storage_server_id"); ok {
		tmp := externalExadataStorageServerId.(string)
		request.ExternalExadataStorageServerId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database_management")

	response, err := s.Client.GetIormPlan(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatabaseManagementExternalExadataStorageServerIormPlanDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseManagementExternalExadataStorageServerIormPlanDataSource-", DatabaseManagementExternalExadataStorageServerIormPlanDataSource(), s.D))

	if s.Res.DbPlan != nil {
		s.D.Set("db_plan", []interface{}{DatabasePlanToMap(s.Res.DbPlan)})
	} else {
		s.D.Set("db_plan", nil)
	}

	s.D.Set("plan_objective", s.Res.PlanObjective)

	s.D.Set("plan_status", s.Res.PlanStatus)

	return nil
}

func DatabasePlanToMap(obj *oci_database_management.DatabasePlan) map[string]interface{} {
	result := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range obj.Items {
		items = append(items, DatabasePlanDirectiveToMap(item))
	}
	result["items"] = items

	return result
}

func DatabasePlanDirectiveToMap(obj oci_database_management.DatabasePlanDirective) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Allocation != nil {
		result["allocation"] = int(*obj.Allocation)
	}

	if obj.Limit != nil {
		result["limit"] = int(*obj.Limit)
	}

	if obj.AsmCluster != nil {
		result["asm_cluster"] = string(*obj.AsmCluster)
	}

	if obj.FlashCacheLimit != nil {
		result["flash_cache_limit"] = string(*obj.FlashCacheLimit)
	}

	if obj.FlashCacheMin != nil {
		result["flash_cache_min"] = string(*obj.FlashCacheMin)
	}

	if obj.FlashCacheSize != nil {
		result["flash_cache_size"] = string(*obj.FlashCacheSize)
	}

	if obj.IsFlashCacheOn != nil {
		result["is_flash_cache_on"] = bool(*obj.IsFlashCacheOn)
	}

	if obj.IsFlashLogOn != nil {
		result["is_flash_log_on"] = bool(*obj.IsFlashLogOn)
	}

	if obj.IsPmemCacheOn != nil {
		result["is_pmem_cache_on"] = bool(*obj.IsPmemCacheOn)
	}

	if obj.IsPmemLogOn != nil {
		result["is_pmem_log_on"] = bool(*obj.IsPmemLogOn)
	}

	if obj.Level != nil {
		result["level"] = int(*obj.Level)
	}

	if obj.Limit != nil {
		result["limit"] = int(*obj.Limit)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.PmemCacheLimit != nil {
		result["pmem_cache_limit"] = string(*obj.PmemCacheLimit)
	}

	if obj.PmemCacheMin != nil {
		result["pmem_cache_min"] = string(*obj.PmemCacheMin)
	}

	if obj.PmemCacheSize != nil {
		result["pmem_cache_size"] = string(*obj.PmemCacheSize)
	}

	result["role"] = string(obj.Role)

	if obj.Share != nil {
		result["share"] = int(*obj.Share)
	}

	result["type"] = string(obj.Type)

	return result
}
