// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"

	oci_database "github.com/oracle/oci-go-sdk/database"
)

func DatabaseAutonomousDatabaseResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: DefaultTimeout,
		Create:   createDatabaseAutonomousDatabase,
		Read:     readDatabaseAutonomousDatabase,
		Update:   updateDatabaseAutonomousDatabase,
		Delete:   deleteDatabaseAutonomousDatabase,
		Schema: map[string]*schema.Schema{
			// Required
			"admin_password": {
				Type:      schema.TypeString,
				Required:  true,
				Sensitive: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"cpu_core_count": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"data_storage_size_in_tbs": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"db_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"db_workload": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: definedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"display_name": {
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
				ForceNew: true,
			},

			// Computed
			"connection_strings": {
				Type:     schema.TypeList,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"all_connection_strings": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem:     schema.TypeString,
						},
						"high": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"low": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"medium": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"db_version": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"service_console_url": {
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
		},
	}
}

func createDatabaseAutonomousDatabase(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseAutonomousDatabaseResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).databaseClient

	return CreateResource(d, sync)
}

func readDatabaseAutonomousDatabase(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseAutonomousDatabaseResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).databaseClient

	return ReadResource(sync)
}

func updateDatabaseAutonomousDatabase(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseAutonomousDatabaseResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).databaseClient

	return UpdateResource(d, sync)
}

func deleteDatabaseAutonomousDatabase(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseAutonomousDatabaseResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).databaseClient
	sync.DisableNotFoundRetries = true

	return DeleteResource(d, sync)
}

type DatabaseAutonomousDatabaseResourceCrud struct {
	BaseCrud
	Client                 *oci_database.DatabaseClient
	Res                    *oci_database.AutonomousDatabase
	DisableNotFoundRetries bool
}

func (s *DatabaseAutonomousDatabaseResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DatabaseAutonomousDatabaseResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_database.AutonomousDatabaseLifecycleStateProvisioning),
		string(oci_database.AutonomousDatabaseLifecycleStateStarting),
	}
}

func (s *DatabaseAutonomousDatabaseResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_database.AutonomousDatabaseLifecycleStateAvailable),
	}
}

func (s *DatabaseAutonomousDatabaseResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_database.AutonomousDatabaseLifecycleStateTerminating),
		string(oci_database.AutonomousDatabaseLifecycleStateUnavailable),
	}
}

func (s *DatabaseAutonomousDatabaseResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_database.AutonomousDatabaseLifecycleStateTerminated),
	}
}

func (s *DatabaseAutonomousDatabaseResourceCrud) UpdatedPending() []string {
	return []string{
		string(oci_database.AutonomousDatabaseLifecycleStateProvisioning),
		string(oci_database.AutonomousDatabaseLifecycleStateUnavailable),
		string(oci_database.AutonomousDatabaseLifecycleStateScaleInProgress),
	}
}

func (s *DatabaseAutonomousDatabaseResourceCrud) UpdatedTarget() []string {
	return []string{
		string(oci_database.AutonomousDatabaseLifecycleStateAvailable),
	}
}

func (s *DatabaseAutonomousDatabaseResourceCrud) Create() error {
	request := oci_database.CreateAutonomousDatabaseRequest{}

	if adminPassword, ok := s.D.GetOkExists("admin_password"); ok {
		tmp := adminPassword.(string)
		request.AdminPassword = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if cpuCoreCount, ok := s.D.GetOkExists("cpu_core_count"); ok {
		tmp := cpuCoreCount.(int)
		request.CpuCoreCount = &tmp
	}

	if dataStorageSizeInTBs, ok := s.D.GetOkExists("data_storage_size_in_tbs"); ok {
		tmp := dataStorageSizeInTBs.(int)
		request.DataStorageSizeInTBs = &tmp
	}

	if dbName, ok := s.D.GetOkExists("db_name"); ok {
		tmp := dbName.(string)
		request.DbName = &tmp
	}

	if dbWorkload, ok := s.D.GetOkExists("db_workload"); ok {
		request.DbWorkload = oci_database.CreateAutonomousDatabaseDetailsDbWorkloadEnum(dbWorkload.(string))
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := mapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = objectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if licenseModel, ok := s.D.GetOkExists("license_model"); ok {
		request.LicenseModel = oci_database.CreateAutonomousDatabaseDetailsLicenseModelEnum(licenseModel.(string))
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.CreateAutonomousDatabase(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.AutonomousDatabase
	return nil
}

func (s *DatabaseAutonomousDatabaseResourceCrud) Get() error {
	request := oci_database.GetAutonomousDatabaseRequest{}

	tmp := s.D.Id()
	request.AutonomousDatabaseId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.GetAutonomousDatabase(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.AutonomousDatabase
	return nil
}

func (s *DatabaseAutonomousDatabaseResourceCrud) Update() error {
	request := oci_database.UpdateAutonomousDatabaseRequest{}

	// @CODEGEN 09/2018: Cannot update the password and scale the Autonomous Transaction Processing in same request, only include changed properties in request
	if adminPassword, ok := s.D.GetOkExists("admin_password"); ok && s.D.HasChange("admin_password") {
		tmp := adminPassword.(string)
		request.AdminPassword = &tmp
	}

	tmp := s.D.Id()
	request.AutonomousDatabaseId = &tmp

	if cpuCoreCount, ok := s.D.GetOkExists("cpu_core_count"); ok && s.D.HasChange("cpu_core_count") {
		tmp := cpuCoreCount.(int)
		request.CpuCoreCount = &tmp
	}

	if dataStorageSizeInTBs, ok := s.D.GetOkExists("data_storage_size_in_tbs"); ok && s.D.HasChange("data_storage_size_in_tbs") {
		tmp := dataStorageSizeInTBs.(int)
		request.DataStorageSizeInTBs = &tmp
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := mapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = objectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.UpdateAutonomousDatabase(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.AutonomousDatabase
	return nil
}

func (s *DatabaseAutonomousDatabaseResourceCrud) Delete() error {
	request := oci_database.DeleteAutonomousDatabaseRequest{}

	tmp := s.D.Id()
	request.AutonomousDatabaseId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "database")

	_, err := s.Client.DeleteAutonomousDatabase(context.Background(), request)
	return err
}

func (s *DatabaseAutonomousDatabaseResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.ConnectionStrings != nil {
		s.D.Set("connection_strings", []interface{}{AutonomousDatabaseConnectionStringsToMap(s.Res.ConnectionStrings)})
	} else {
		s.D.Set("connection_strings", nil)
	}

	if s.Res.CpuCoreCount != nil {
		s.D.Set("cpu_core_count", *s.Res.CpuCoreCount)
	}

	if s.Res.DataStorageSizeInTBs != nil {
		s.D.Set("data_storage_size_in_tbs", *s.Res.DataStorageSizeInTBs)
	}

	if s.Res.DbName != nil {
		s.D.Set("db_name", *s.Res.DbName)
	}

	if s.Res.DbVersion != nil {
		s.D.Set("db_version", *s.Res.DbVersion)
	}

	s.D.Set("db_workload", s.Res.DbWorkload)

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", definedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	s.D.Set("license_model", s.Res.LicenseModel)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.ServiceConsoleUrl != nil {
		s.D.Set("service_console_url", *s.Res.ServiceConsoleUrl)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}

func AutonomousDatabaseConnectionStringsToMap(obj *oci_database.AutonomousDatabaseConnectionStrings) map[string]interface{} {
	result := map[string]interface{}{}

	result["all_connection_strings"] = obj.AllConnectionStrings

	if obj.High != nil {
		result["high"] = string(*obj.High)
	}

	if obj.Low != nil {
		result["low"] = string(*obj.Low)
	}

	if obj.Medium != nil {
		result["medium"] = string(*obj.Medium)
	}

	return result
}
