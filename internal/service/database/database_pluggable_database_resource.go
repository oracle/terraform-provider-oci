// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	oci_database "github.com/oracle/oci-go-sdk/v56/database"
)

func DatabasePluggableDatabaseResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDatabasePluggableDatabase,
		Read:     readDatabasePluggableDatabase,
		Update:   updateDatabasePluggableDatabase,
		Delete:   deleteDatabasePluggableDatabase,
		Schema: map[string]*schema.Schema{
			// Required
			"container_database_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"pdb_name": {
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
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"pdb_admin_password": {
				Type:      schema.TypeString,
				Optional:  true,
				Computed:  true,
				ForceNew:  true,
				Sensitive: true,
			},
			"should_pdb_admin_account_be_locked": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"tde_wallet_password": {
				Type:      schema.TypeString,
				Optional:  true,
				Computed:  true,
				ForceNew:  true,
				Sensitive: true,
			},

			// Computed
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
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
						"pdb_default": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"pdb_ip_default": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"is_restricted": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"open_mode": {
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

func createDatabasePluggableDatabase(d *schema.ResourceData, m interface{}) error {
	sync := &DatabasePluggableDatabaseResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.CreateResource(d, sync)
}

func readDatabasePluggableDatabase(d *schema.ResourceData, m interface{}) error {
	sync := &DatabasePluggableDatabaseResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

func updateDatabasePluggableDatabase(d *schema.ResourceData, m interface{}) error {
	sync := &DatabasePluggableDatabaseResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteDatabasePluggableDatabase(d *schema.ResourceData, m interface{}) error {
	sync := &DatabasePluggableDatabaseResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type DatabasePluggableDatabaseResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_database.DatabaseClient
	Res                    *oci_database.PluggableDatabase
	DisableNotFoundRetries bool
}

func (s *DatabasePluggableDatabaseResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DatabasePluggableDatabaseResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_database.PluggableDatabaseLifecycleStateProvisioning),
	}
}

func (s *DatabasePluggableDatabaseResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_database.PluggableDatabaseLifecycleStateAvailable),
	}
}

func (s *DatabasePluggableDatabaseResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_database.PluggableDatabaseLifecycleStateTerminating),
	}
}

func (s *DatabasePluggableDatabaseResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_database.PluggableDatabaseLifecycleStateTerminated),
	}
}

func (s *DatabasePluggableDatabaseResourceCrud) Create() error {
	request := oci_database.CreatePluggableDatabaseRequest{}

	if containerDatabaseId, ok := s.D.GetOkExists("container_database_id"); ok {
		tmp := containerDatabaseId.(string)
		request.ContainerDatabaseId = &tmp
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if pdbAdminPassword, ok := s.D.GetOkExists("pdb_admin_password"); ok {
		tmp := pdbAdminPassword.(string)
		request.PdbAdminPassword = &tmp
	}

	if pdbName, ok := s.D.GetOkExists("pdb_name"); ok {
		tmp := pdbName.(string)
		request.PdbName = &tmp
	}

	if shouldPdbAdminAccountBeLocked, ok := s.D.GetOkExists("should_pdb_admin_account_be_locked"); ok {
		tmp := shouldPdbAdminAccountBeLocked.(bool)
		request.ShouldPdbAdminAccountBeLocked = &tmp
	}

	if tdeWalletPassword, ok := s.D.GetOkExists("tde_wallet_password"); ok {
		tmp := tdeWalletPassword.(string)
		request.TdeWalletPassword = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.CreatePluggableDatabase(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.PluggableDatabase
	return nil
}

func (s *DatabasePluggableDatabaseResourceCrud) Get() error {
	request := oci_database.GetPluggableDatabaseRequest{}

	tmp := s.D.Id()
	request.PluggableDatabaseId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.GetPluggableDatabase(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.PluggableDatabase
	return nil
}

func (s *DatabasePluggableDatabaseResourceCrud) Update() error {
	request := oci_database.UpdatePluggableDatabaseRequest{}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	tmp := s.D.Id()
	request.PluggableDatabaseId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.UpdatePluggableDatabase(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.PluggableDatabase
	return nil
}

func (s *DatabasePluggableDatabaseResourceCrud) Delete() error {
	request := oci_database.DeletePluggableDatabaseRequest{}

	tmp := s.D.Id()
	request.PluggableDatabaseId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	_, err := s.Client.DeletePluggableDatabase(context.Background(), request)
	return err
}

func (s *DatabasePluggableDatabaseResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.ConnectionStrings != nil {
		s.D.Set("connection_strings", []interface{}{PluggableDatabaseConnectionStringsToMap(s.Res.ConnectionStrings)})
	} else {
		s.D.Set("connection_strings", nil)
	}

	if s.Res.ContainerDatabaseId != nil {
		s.D.Set("container_database_id", *s.Res.ContainerDatabaseId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IsRestricted != nil {
		s.D.Set("is_restricted", *s.Res.IsRestricted)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	s.D.Set("open_mode", s.Res.OpenMode)

	if s.Res.PdbName != nil {
		s.D.Set("pdb_name", *s.Res.PdbName)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}

func PluggableDatabaseConnectionStringsToMap(obj *oci_database.PluggableDatabaseConnectionStrings) map[string]interface{} {
	result := map[string]interface{}{}

	result["all_connection_strings"] = obj.AllConnectionStrings

	if obj.PdbDefault != nil {
		result["pdb_default"] = string(*obj.PdbDefault)
	}

	if obj.PdbIpDefault != nil {
		result["pdb_ip_default"] = string(*obj.PdbIpDefault)
	}

	return result
}
