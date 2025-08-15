// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package dbmulticloud

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_dbmulticloud "github.com/oracle/oci-go-sdk/v65/dbmulticloud"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DbmulticloudOracleDbAzureKeysDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDbmulticloudOracleDbAzureKeys,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"oracle_db_azure_key_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"oracle_db_azure_vault_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"oracle_db_azure_key_summary_collection": {
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
									"azure_key_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"compartment_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
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
									"id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"key_properties": {
										Type:     schema.TypeMap,
										Computed: true,
									},
									"last_modification": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"lifecycle_state_details": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"oracle_db_azure_vault_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"resource_type": {
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

func readDbmulticloudOracleDbAzureKeys(d *schema.ResourceData, m interface{}) error {
	sync := &DbmulticloudOracleDbAzureKeysDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OracleDbAzureKeyClient()

	return tfresource.ReadResource(sync)
}

type DbmulticloudOracleDbAzureKeysDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_dbmulticloud.OracleDbAzureKeyClient
	Res    *oci_dbmulticloud.ListOracleDbAzureKeysResponse
}

func (s *DbmulticloudOracleDbAzureKeysDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DbmulticloudOracleDbAzureKeysDataSourceCrud) Get() error {
	request := oci_dbmulticloud.ListOracleDbAzureKeysRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if oracleDbAzureKeyId, ok := s.D.GetOkExists("id"); ok {
		tmp := oracleDbAzureKeyId.(string)
		request.OracleDbAzureKeyId = &tmp
	}

	if oracleDbAzureVaultId, ok := s.D.GetOkExists("oracle_db_azure_vault_id"); ok {
		tmp := oracleDbAzureVaultId.(string)
		request.OracleDbAzureVaultId = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_dbmulticloud.OracleDbAzureKeyLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "dbmulticloud")

	response, err := s.Client.ListOracleDbAzureKeys(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListOracleDbAzureKeys(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DbmulticloudOracleDbAzureKeysDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DbmulticloudOracleDbAzureKeysDataSource-", DbmulticloudOracleDbAzureKeysDataSource(), s.D))
	resources := []map[string]interface{}{}
	oracleDbAzureKey := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, OracleDbAzureKeySummaryToMap(item))
	}
	oracleDbAzureKey["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DbmulticloudOracleDbAzureKeysDataSource().Schema["oracle_db_azure_key_summary_collection"].Elem.(*schema.Resource).Schema)
		oracleDbAzureKey["items"] = items
	}

	resources = append(resources, oracleDbAzureKey)
	if err := s.D.Set("oracle_db_azure_key_summary_collection", resources); err != nil {
		return err
	}

	return nil
}

func OracleDbAzureKeySummaryToMap(obj oci_dbmulticloud.OracleDbAzureKeySummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AzureKeyId != nil {
		result["azure_key_id"] = string(*obj.AzureKeyId)
	}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	//if obj.KeyProperties != nil {
	//	result["key_properties"] = []interface{}{obj.KeyProperties}
	//}
	if obj.KeyProperties != nil {
		keyPropsInterface := *obj.KeyProperties

		keyPropsMap, ok := keyPropsInterface.(map[string]interface{})
		if !ok {
			return result
		}

		result["key_properties"] = keyPropsMap
	}

	if obj.LastModification != nil {
		result["last_modification"] = string(*obj.LastModification)
	}

	if obj.LifecycleStateDetails != nil {
		result["lifecycle_state_details"] = string(*obj.LifecycleStateDetails)
	}

	if obj.OracleDbAzureVaultId != nil {
		result["oracle_db_azure_vault_id"] = string(*obj.OracleDbAzureVaultId)
	}

	if obj.ResourceType != nil {
		result["resource_type"] = string(*obj.ResourceType)
	}

	result["state"] = string(obj.LifecycleState)

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
