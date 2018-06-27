// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_database "github.com/oracle/oci-go-sdk/database"

	"github.com/oracle/terraform-provider-oci/crud"
)

func DbSystemsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDbSystems,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"backup_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"limit": {
				Type:       schema.TypeInt,
				Optional:   true,
				Deprecated: crud.FieldDeprecated("limit"),
			},
			"page": {
				Type:       schema.TypeString,
				Optional:   true,
				Deprecated: crud.FieldDeprecated("page"),
			},
			"db_systems": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     DbSystemResource(),
			},
		},
	}
}

func readDbSystems(d *schema.ResourceData, m interface{}) error {
	sync := &DbSystemsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).databaseClient

	return crud.ReadResource(sync)
}

type DbSystemsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.ListDbSystemsResponse
}

func (s *DbSystemsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DbSystemsDataSourceCrud) Get() error {
	request := oci_database.ListDbSystemsRequest{}

	if backupId, ok := s.D.GetOkExists("backup_id"); ok {
		tmp := backupId.(string)
		request.BackupId = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if limit, ok := s.D.GetOkExists("limit"); ok {
		tmp := limit.(int)
		request.Limit = &tmp
	}

	if page, ok := s.D.GetOkExists("page"); ok {
		tmp := page.(string)
		request.Page = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "database")

	response, err := s.Client.ListDbSystems(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListDbSystems(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DbSystemsDataSourceCrud) SetData() {
	if s.Res == nil {
		return
	}

	s.D.SetId(crud.GenerateDataSourceID())
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		dbSystem := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.AvailabilityDomain != nil {
			dbSystem["availability_domain"] = *r.AvailabilityDomain
		}

		if r.BackupSubnetId != nil {
			dbSystem["backup_subnet_id"] = *r.BackupSubnetId
		}

		if r.ClusterName != nil {
			dbSystem["cluster_name"] = *r.ClusterName
		}

		if r.CpuCoreCount != nil {
			dbSystem["cpu_core_count"] = *r.CpuCoreCount
		}

		if r.DataStoragePercentage != nil {
			dbSystem["data_storage_percentage"] = *r.DataStoragePercentage
		}

		if r.DataStorageSizeInGBs != nil {
			dbSystem["data_storage_size_in_gb"] = *r.DataStorageSizeInGBs
		}

		dbSystem["database_edition"] = r.DatabaseEdition

		if r.DefinedTags != nil {
			dbSystem["defined_tags"] = definedTagsToMap(r.DefinedTags)
		}

		dbSystem["disk_redundancy"] = r.DiskRedundancy

		if r.DisplayName != nil {
			dbSystem["display_name"] = *r.DisplayName
		}

		if r.Domain != nil {
			dbSystem["domain"] = *r.Domain
		}

		dbSystem["freeform_tags"] = r.FreeformTags

		if r.Hostname != nil {
			dbSystem["hostname"] = *r.Hostname
		}

		if r.Id != nil {
			dbSystem["id"] = *r.Id
		}

		if r.LastPatchHistoryEntryId != nil {
			dbSystem["last_patch_history_entry_id"] = *r.LastPatchHistoryEntryId
		}

		dbSystem["license_model"] = r.LicenseModel

		if r.LifecycleDetails != nil {
			dbSystem["lifecycle_details"] = *r.LifecycleDetails
		}

		if r.ListenerPort != nil {
			dbSystem["listener_port"] = *r.ListenerPort
		}

		if r.NodeCount != nil {
			dbSystem["node_count"] = *r.NodeCount
		}

		if r.RecoStorageSizeInGB != nil {
			dbSystem["reco_storage_size_in_gb"] = *r.RecoStorageSizeInGB
		}

		if r.ScanDnsRecordId != nil {
			dbSystem["scan_dns_record_id"] = *r.ScanDnsRecordId
		}

		dbSystem["scan_ip_ids"] = r.ScanIpIds

		if r.Shape != nil {
			dbSystem["shape"] = *r.Shape
		}

		dbSystem["ssh_public_keys"] = r.SshPublicKeys

		dbSystem["state"] = r.LifecycleState

		if r.SubnetId != nil {
			dbSystem["subnet_id"] = *r.SubnetId
		}

		if r.TimeCreated != nil {
			dbSystem["time_created"] = r.TimeCreated.String()
		}

		if r.Version != nil {
			dbSystem["version"] = *r.Version
		}

		dbSystem["vip_ids"] = r.VipIds

		resources = append(resources, dbSystem)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, DbSystemsDataSource().Schema["db_systems"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("db_systems", resources); err != nil {
		panic(err)
	}

	return
}
