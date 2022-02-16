// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	oci_database "github.com/oracle/oci-go-sdk/v58/database"
	oci_work_requests "github.com/oracle/oci-go-sdk/v58/workrequests"
)

func DatabaseCloudAutonomousVmClusterResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDatabaseCloudAutonomousVmCluster,
		Read:     readDatabaseCloudAutonomousVmCluster,
		Update:   updateDatabaseCloudAutonomousVmCluster,
		Delete:   deleteDatabaseCloudAutonomousVmCluster,
		Schema: map[string]*schema.Schema{
			// Required
			"cloud_exadata_infrastructure_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"subnet_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"license_model": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"nsg_ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Set:      utils.LiteralTypeHashCodeForSets,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			// Computed
			"availability_domain": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"cpu_core_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"data_storage_size_in_gb": {
				Type:     schema.TypeFloat,
				Computed: true,
			},
			"data_storage_size_in_tbs": {
				Type:     schema.TypeFloat,
				Computed: true,
			},
			"domain": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"hostname": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"last_maintenance_run_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"last_update_history_entry_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"memory_size_in_gbs": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"next_maintenance_run_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"node_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"ocpu_count": {
				Type:     schema.TypeFloat,
				Computed: true,
			},
			"shape": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"rotate_ords_certs_trigger": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"rotate_ssl_certs_trigger": {
				Type:     schema.TypeBool,
				Optional: true,
			},
		},
	}
}

func createDatabaseCloudAutonomousVmCluster(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseCloudAutonomousVmClusterResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	if e := tfresource.CreateResource(d, sync); e != nil {
		return e
	}

	if rotateOrdsCertsTrigger, ok := sync.D.GetOkExists("rotate_ords_certs_trigger"); ok {
		tmp := rotateOrdsCertsTrigger.(bool)
		if tmp {
			err := sync.RotateAutonomousCloudVmClusterOrdsCerts()
			if err != nil {
				return err
			}
		}
	}

	if rotateSslCertsTrigger, ok := sync.D.GetOkExists("rotate_ssl_certs_trigger"); ok {
		tmp := rotateSslCertsTrigger.(bool)
		if tmp {
			err := sync.RotateAutonomousCloudVmClusterSslCerts()
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func readDatabaseCloudAutonomousVmCluster(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseCloudAutonomousVmClusterResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

func updateDatabaseCloudAutonomousVmCluster(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseCloudAutonomousVmClusterResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	if rotateOrdsCertsTrigger, ok := sync.D.GetOkExists("rotate_ords_certs_trigger"); ok && sync.D.HasChange("rotate_ords_certs_trigger") {
		tmp := rotateOrdsCertsTrigger.(bool)
		if tmp {
			err := sync.RotateAutonomousCloudVmClusterOrdsCerts()
			if err != nil {
				return err
			}
		}
	}

	if rotateSslCertsTrigger, ok := sync.D.GetOkExists("rotate_ssl_certs_trigger"); ok && sync.D.HasChange("rotate_ssl_certs_trigger") {
		tmp := rotateSslCertsTrigger.(bool)
		if tmp {
			err := sync.RotateAutonomousCloudVmClusterSslCerts()
			if err != nil {
				return err
			}
		}
	}

	return tfresource.UpdateResource(d, sync)
}

func deleteDatabaseCloudAutonomousVmCluster(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseCloudAutonomousVmClusterResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.DisableNotFoundRetries = true
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.DeleteResource(d, sync)
}

type DatabaseCloudAutonomousVmClusterResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_database.DatabaseClient
	Res                    *oci_database.CloudAutonomousVmCluster
	DisableNotFoundRetries bool
	WorkRequestClient      *oci_work_requests.WorkRequestClient
}

func (s *DatabaseCloudAutonomousVmClusterResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DatabaseCloudAutonomousVmClusterResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_database.CloudAutonomousVmClusterLifecycleStateProvisioning),
	}
}

func (s *DatabaseCloudAutonomousVmClusterResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_database.CloudAutonomousVmClusterLifecycleStateAvailable),
	}
}

func (s *DatabaseCloudAutonomousVmClusterResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_database.CloudAutonomousVmClusterLifecycleStateTerminating),
	}
}

func (s *DatabaseCloudAutonomousVmClusterResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_database.CloudAutonomousVmClusterLifecycleStateTerminated),
	}
}

func (s *DatabaseCloudAutonomousVmClusterResourceCrud) Create() error {
	request := oci_database.CreateCloudAutonomousVmClusterRequest{}

	if cloudExadataInfrastructureId, ok := s.D.GetOkExists("cloud_exadata_infrastructure_id"); ok {
		tmp := cloudExadataInfrastructureId.(string)
		request.CloudExadataInfrastructureId = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if licenseModel, ok := s.D.GetOkExists("license_model"); ok {
		request.LicenseModel = oci_database.CreateCloudAutonomousVmClusterDetailsLicenseModelEnum(licenseModel.(string))
	}

	if nsgIds, ok := s.D.GetOkExists("nsg_ids"); ok {
		set := nsgIds.(*schema.Set)
		interfaces := set.List()
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("nsg_ids") {
			request.NsgIds = tmp
		}
	}

	if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
		tmp := subnetId.(string)
		request.SubnetId = &tmp
	}

	request.RequestMetadata.RetryPolicy = nil //tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.CreateCloudAutonomousVmCluster(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.CloudAutonomousVmCluster

	workId := response.OpcWorkRequestId
	if workId != nil {
		identifier, err := tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "cloudAutonomousVmCluster", oci_work_requests.WorkRequestResourceActionTypeCreated, s.D.Timeout(schema.TimeoutCreate), s.DisableNotFoundRetries)
		if identifier != nil {
			s.D.SetId(*identifier)
		}
		if err != nil {
			return err
		}
	}

	return s.Get()
}

func (s *DatabaseCloudAutonomousVmClusterResourceCrud) Get() error {
	request := oci_database.GetCloudAutonomousVmClusterRequest{}

	tmp := s.D.Id()
	request.CloudAutonomousVmClusterId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.GetCloudAutonomousVmCluster(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.CloudAutonomousVmCluster
	return nil
}

func (s *DatabaseCloudAutonomousVmClusterResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_database.UpdateCloudAutonomousVmClusterRequest{}

	tmp := s.D.Id()
	request.CloudAutonomousVmClusterId = &tmp

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if licenseModel, ok := s.D.GetOkExists("license_model"); ok {
		request.LicenseModel = oci_database.UpdateCloudAutonomousVmClusterDetailsLicenseModelEnum(licenseModel.(string))
	}

	if nsgIds, ok := s.D.GetOkExists("nsg_ids"); ok {
		set := nsgIds.(*schema.Set)
		interfaces := set.List()
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("nsg_ids") {
			request.NsgIds = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.UpdateCloudAutonomousVmCluster(context.Background(), request)
	if err != nil {
		return err
	}
	s.Res = &response.CloudAutonomousVmCluster

	workId := response.OpcWorkRequestId
	if workId != nil {
		_, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "cloudAutonomousVmCluster", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries)
		if err != nil {
			return err
		}
	}

	return s.Get()
}

func (s *DatabaseCloudAutonomousVmClusterResourceCrud) Delete() error {
	request := oci_database.DeleteCloudAutonomousVmClusterRequest{}

	tmp := s.D.Id()
	request.CloudAutonomousVmClusterId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.DeleteCloudAutonomousVmCluster(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	if workId != nil {
		_, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "cloudAutonomousVmCluster", oci_work_requests.WorkRequestResourceActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *DatabaseCloudAutonomousVmClusterResourceCrud) SetData() error {
	if s.Res.AvailabilityDomain != nil {
		s.D.Set("availability_domain", *s.Res.AvailabilityDomain)
	}

	if s.Res.CloudExadataInfrastructureId != nil {
		s.D.Set("cloud_exadata_infrastructure_id", *s.Res.CloudExadataInfrastructureId)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.CpuCoreCount != nil {
		s.D.Set("cpu_core_count", *s.Res.CpuCoreCount)
	}

	if s.Res.DataStorageSizeInGBs != nil {
		s.D.Set("data_storage_size_in_gb", *s.Res.DataStorageSizeInGBs)
	}

	if s.Res.DataStorageSizeInTBs != nil {
		s.D.Set("data_storage_size_in_tbs", *s.Res.DataStorageSizeInTBs)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.Domain != nil {
		s.D.Set("domain", *s.Res.Domain)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.Hostname != nil {
		s.D.Set("hostname", *s.Res.Hostname)
	}

	if s.Res.LastMaintenanceRunId != nil {
		s.D.Set("last_maintenance_run_id", *s.Res.LastMaintenanceRunId)
	}

	if s.Res.LastUpdateHistoryEntryId != nil {
		s.D.Set("last_update_history_entry_id", *s.Res.LastUpdateHistoryEntryId)
	}

	s.D.Set("license_model", s.Res.LicenseModel)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.MemorySizeInGBs != nil {
		s.D.Set("memory_size_in_gbs", *s.Res.MemorySizeInGBs)
	}

	if s.Res.NextMaintenanceRunId != nil {
		s.D.Set("next_maintenance_run_id", *s.Res.NextMaintenanceRunId)
	}

	if s.Res.NodeCount != nil {
		s.D.Set("node_count", *s.Res.NodeCount)
	}

	nsgIds := []interface{}{}
	for _, item := range s.Res.NsgIds {
		nsgIds = append(nsgIds, item)
	}
	s.D.Set("nsg_ids", schema.NewSet(utils.LiteralTypeHashCodeForSets, nsgIds))

	if s.Res.OcpuCount != nil {
		s.D.Set("ocpu_count", *s.Res.OcpuCount)
	}

	if s.Res.Shape != nil {
		s.D.Set("shape", *s.Res.Shape)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SubnetId != nil {
		s.D.Set("subnet_id", *s.Res.SubnetId)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}

func (s *DatabaseCloudAutonomousVmClusterResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_database.ChangeCloudAutonomousVmClusterCompartmentRequest{}

	idTmp := s.D.Id()
	changeCompartmentRequest.CloudAutonomousVmClusterId = &idTmp

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.ChangeCloudAutonomousVmClusterCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	if workId != nil {
		_, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "cloudAutonomousVmCluster", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *DatabaseCloudAutonomousVmClusterResourceCrud) RotateAutonomousCloudVmClusterOrdsCerts() error {
	request := oci_database.RotateCloudAutonomousVmClusterOrdsCertsRequest{}

	tmp := s.D.Id()
	request.CloudAutonomousVmClusterId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.RotateCloudAutonomousVmClusterOrdsCerts(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	_, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "database", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries)
	if err != nil {
		return err
	}

	return nil
}

func (s *DatabaseCloudAutonomousVmClusterResourceCrud) RotateAutonomousCloudVmClusterSslCerts() error {
	request := oci_database.RotateCloudAutonomousVmClusterSslCertsRequest{}

	tmp := s.D.Id()
	request.CloudAutonomousVmClusterId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.RotateCloudAutonomousVmClusterSslCerts(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	_, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "database", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries)
	if err != nil {
		return err
	}

	return nil
}
