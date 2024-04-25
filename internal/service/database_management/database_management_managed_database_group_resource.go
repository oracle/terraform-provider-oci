// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database_management

import (
	"bytes"
	"context"
	"fmt"
	"log"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_database_management "github.com/oracle/oci-go-sdk/v65/databasemanagement"
)

func DatabaseManagementManagedDatabaseGroupResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDatabaseManagementManagedDatabaseGroup,
		Read:     readDatabaseManagementManagedDatabaseGroup,
		Update:   updateDatabaseManagementManagedDatabaseGroup,
		Delete:   deleteDatabaseManagementManagedDatabaseGroup,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
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

			"managed_databases": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Set:      managedDatabaseHashCodeForSets,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Computed
						"id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"compartment_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"database_sub_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"defined_tags": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem:     schema.TypeString,
						},
						"deployment_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"freeform_tags": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem:     schema.TypeString,
						},
						"database_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"system_tags": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem:     schema.TypeString,
						},
						"time_added": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"workload_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},

			// Computed
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
	}
}

func managedDatabaseHashCodeForSets(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})
	if id, ok := m["id"]; ok && id != "" {
		buf.WriteString(fmt.Sprintf("%v-", id))
	}
	return utils.GetStringHashcode(buf.String())
}

func createDatabaseManagementManagedDatabaseGroup(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementManagedDatabaseGroupResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()

	return tfresource.CreateResource(d, sync)
}

func readDatabaseManagementManagedDatabaseGroup(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementManagedDatabaseGroupResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()

	return tfresource.ReadResource(sync)
}

func updateDatabaseManagementManagedDatabaseGroup(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementManagedDatabaseGroupResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteDatabaseManagementManagedDatabaseGroup(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementManagedDatabaseGroupResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type DatabaseManagementManagedDatabaseGroupResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_database_management.DbManagementClient
	Res                    *oci_database_management.ManagedDatabaseGroup
	DisableNotFoundRetries bool
}

func (s *DatabaseManagementManagedDatabaseGroupResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DatabaseManagementManagedDatabaseGroupResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_database_management.LifecycleStatesCreating),
	}
}

func (s *DatabaseManagementManagedDatabaseGroupResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_database_management.LifecycleStatesActive),
	}
}

func (s *DatabaseManagementManagedDatabaseGroupResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_database_management.LifecycleStatesDeleting),
	}
}

func (s *DatabaseManagementManagedDatabaseGroupResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_database_management.LifecycleStatesDeleted),
	}
}

func (s *DatabaseManagementManagedDatabaseGroupResourceCrud) Create() error {
	defer func() {
		if s.Res != nil {
			managedDatabaseGroupId := *s.Res.Id
			log.Printf("Invoking GET() for '%v'", managedDatabaseGroupId)
			// get latest state of the instance
			err := s.GetManagedDatabaseGroupForManagedDatabaseGroupId(managedDatabaseGroupId)
			if err != nil {
				log.Printf("[ERROR] unable to invoke GET() after CREATE '%v'", err)
			}
			// write latest state
			if err := s.SetData(); err != nil {
				log.Printf("[ERROR] unable to invoke setData() '%v'", err)
			}
		}
	}()

	request := oci_database_management.CreateManagedDatabaseGroupRequest{}

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

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_management")

	response, err := s.Client.CreateManagedDatabaseGroup(context.Background(), request)
	if err != nil {
		return err
	}
	s.Res = &response.ManagedDatabaseGroup

	managedDatabaseGroupId := *response.ManagedDatabaseGroup.Id

	err = s.updateManagedDatabases(&managedDatabaseGroupId)
	if err != nil {
		return err
	}
	return nil
}

func mapToAddManagedDatabaseToManagedDatabaseGroupDetails(managedDatabase map[string]interface{}) oci_database_management.AddManagedDatabaseToManagedDatabaseGroupDetails {
	result := oci_database_management.AddManagedDatabaseToManagedDatabaseGroupDetails{}

	if id, ok := managedDatabase["id"]; ok {
		tmp := id.(string)
		result.ManagedDatabaseId = &tmp
	}

	return result
}

func mapToRemoveManagedDatabaseFromManagedDatabaseGroupDetails(managedDatabase map[string]interface{}) oci_database_management.RemoveManagedDatabaseFromManagedDatabaseGroupDetails {
	result := oci_database_management.RemoveManagedDatabaseFromManagedDatabaseGroupDetails{}

	if id, ok := managedDatabase["id"]; ok {
		tmp := id.(string)
		result.ManagedDatabaseId = &tmp
	}

	return result
}

func (s *DatabaseManagementManagedDatabaseGroupResourceCrud) addManagedDatabaseToManagedDatabaseGroup(managedDatabaseGroupId string, managedDatabaseDetails oci_database_management.AddManagedDatabaseToManagedDatabaseGroupDetails) error {
	request := oci_database_management.AddManagedDatabaseToManagedDatabaseGroupRequest{}
	request.ManagedDatabaseGroupId = &managedDatabaseGroupId
	request.AddManagedDatabaseToManagedDatabaseGroupDetails = managedDatabaseDetails
	_, err := s.Client.AddManagedDatabaseToManagedDatabaseGroup(context.Background(), request)
	if err != nil {
		return err
	}
	return nil
}

func (s *DatabaseManagementManagedDatabaseGroupResourceCrud) removeManagedDatabaseFromManagedDatabaseGroup(managedDatabaseGroupId string, managedDatabaseDetails oci_database_management.RemoveManagedDatabaseFromManagedDatabaseGroupDetails) error {
	request := oci_database_management.RemoveManagedDatabaseFromManagedDatabaseGroupRequest{}
	request.ManagedDatabaseGroupId = &managedDatabaseGroupId
	request.RemoveManagedDatabaseFromManagedDatabaseGroupDetails = managedDatabaseDetails
	_, err := s.Client.RemoveManagedDatabaseFromManagedDatabaseGroup(context.Background(), request)
	if err != nil {
		return err
	}
	return nil
}

// Retrieves the Managed Database Group given the Managed Database Group ID
func (s *DatabaseManagementManagedDatabaseGroupResourceCrud) GetManagedDatabaseGroupForManagedDatabaseGroupId(managedDatabaseGroupId string) error {
	request := oci_database_management.GetManagedDatabaseGroupRequest{}
	request.ManagedDatabaseGroupId = &managedDatabaseGroupId

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_management")

	response, err := s.Client.GetManagedDatabaseGroup(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ManagedDatabaseGroup
	return nil
}

func (s *DatabaseManagementManagedDatabaseGroupResourceCrud) Get() error {
	request := oci_database_management.GetManagedDatabaseGroupRequest{}

	tmp := s.D.Id()
	request.ManagedDatabaseGroupId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_management")

	response, err := s.Client.GetManagedDatabaseGroup(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ManagedDatabaseGroup
	return nil
}

func (s *DatabaseManagementManagedDatabaseGroupResourceCrud) Update() error {
	defer func() {
		// get latest state of the instance
		err := s.Get()
		if err != nil {
			log.Printf("[ERROR] unable to invoke GET() after CREATE '%v'", err)
		}
		// write latest state
		if err := s.SetData(); err != nil {
			log.Printf("[ERROR] unable to invoke setData() '%v'", err)
		}
	}()

	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_database_management.UpdateManagedDatabaseGroupRequest{}

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

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	tmp := s.D.Id()
	request.ManagedDatabaseGroupId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_management")

	response, err := s.Client.UpdateManagedDatabaseGroup(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ManagedDatabaseGroup
	managedDatabaseGroupId := *response.ManagedDatabaseGroup.Id

	err = s.updateManagedDatabases(&managedDatabaseGroupId)
	if err != nil {
		return err
	}

	return nil
}

func (s *DatabaseManagementManagedDatabaseGroupResourceCrud) updateManagedDatabases(managedDatabaseGroupId *string) error {
	if _, ok := s.D.GetOkExists("managed_databases"); ok && s.D.HasChange("managed_databases") {
		o, n := s.D.GetChange("managed_databases")
		if o == nil {
			o = new(schema.Set)
		}
		if n == nil {
			n = new(schema.Set)
		}
		oldManagedDatabaseGroups := o.(*schema.Set)
		newManagedDatabaseGroups := n.(*schema.Set)

		managedDatabasesToAdd := newManagedDatabaseGroups.Difference(oldManagedDatabaseGroups).List()
		managedDatabasesToRemove := oldManagedDatabaseGroups.Difference(newManagedDatabaseGroups).List()

		for _, managedDatabase := range managedDatabasesToRemove {
			remove := mapToRemoveManagedDatabaseFromManagedDatabaseGroupDetails(managedDatabase.(map[string]interface{}))
			err := s.removeManagedDatabaseFromManagedDatabaseGroup(*managedDatabaseGroupId, remove)
			if err != nil {
				return fmt.Errorf("failed to remove Managed Database, error: %v", err)
			}
		}

		for _, managedDatabase := range managedDatabasesToAdd {
			add := mapToAddManagedDatabaseToManagedDatabaseGroupDetails(managedDatabase.(map[string]interface{}))
			err := s.addManagedDatabaseToManagedDatabaseGroup(*managedDatabaseGroupId, add)
			if err != nil {
				return fmt.Errorf("failed to add Managed Database, error: %v", err)
			}
		}
	}
	return nil
}

func (s *DatabaseManagementManagedDatabaseGroupResourceCrud) removeAllManagedDatabases(managedDatabaseGroupId *string) error {
	if managedDatabases, ok := s.D.GetOkExists("managed_databases"); ok {
		managedDatabasesToRemove := managedDatabases.(*schema.Set).List()

		for _, managedDatabase := range managedDatabasesToRemove {
			remove := mapToRemoveManagedDatabaseFromManagedDatabaseGroupDetails(managedDatabase.(map[string]interface{}))
			err := s.removeManagedDatabaseFromManagedDatabaseGroup(*managedDatabaseGroupId, remove)
			if err != nil {
				return fmt.Errorf("failed to remove Managed Database, error: %v", err)
			}
		}
	}
	return nil
}

func (s *DatabaseManagementManagedDatabaseGroupResourceCrud) Delete() error {
	request := oci_database_management.DeleteManagedDatabaseGroupRequest{}

	tmp := s.D.Id()
	request.ManagedDatabaseGroupId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_management")

	mderr := s.removeAllManagedDatabases(request.ManagedDatabaseGroupId)
	if mderr != nil {
		return mderr
	}

	_, err := s.Client.DeleteManagedDatabaseGroup(context.Background(), request)
	return err
}

func (s *DatabaseManagementManagedDatabaseGroupResourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	managedDatabases := []interface{}{}
	for _, item := range s.Res.ManagedDatabases {
		managedDatabases = append(managedDatabases, ChildDatabaseToMap(item))
	}
	s.D.Set("managed_databases", schema.NewSet(managedDatabaseHashCodeForSets, managedDatabases))

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}

func ChildDatabaseToMap(obj oci_database_management.ChildDatabase) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	result["database_sub_type"] = string(obj.DatabaseSubType)

	result["database_type"] = string(obj.DatabaseType)

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	result["deployment_type"] = string(obj.DeploymentType)

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TimeAdded != nil {
		result["time_added"] = obj.TimeAdded.String()
	}

	result["workload_type"] = string(obj.WorkloadType)

	return result
}

func ManagedDatabaseGroupSummaryToMap(obj oci_database_management.ManagedDatabaseGroupSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	return result
}

func (s *DatabaseManagementManagedDatabaseGroupResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_database_management.ChangeManagedDatabaseGroupCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.ManagedDatabaseGroupId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_management")

	_, err := s.Client.ChangeManagedDatabaseGroupCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}
