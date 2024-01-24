// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package osmanagement

import (
	"bytes"
	"context"
	"fmt"
	"log"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_osmanagement "github.com/oracle/oci-go-sdk/v65/osmanagement"
)

func OsmanagementManagedInstanceManagementResource() *schema.Resource {
	return &schema.Resource{
		Timeouts: tfresource.DefaultTimeout,
		Create:   createOsmanagementManagedInstanceManagement,
		Read:     readOsmanagementManagedInstanceManagement,
		Update:   updateOsmanagementManagedInstanceManagement,
		Delete:   deleteOsmanagementManagedInstanceManagement,
		Schema: map[string]*schema.Schema{
			"managed_instance_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Optional
			"child_software_sources": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Set:      softwareSourceHashCodeForSets,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"last_boot": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"last_checkin": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"managed_instance_groups": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Set:      managedInstanceGroupsHashCodeForSets,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"display_name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"os_kernel_version": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"os_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"os_version": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"parent_software_source": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Set:      softwareSourceHashCodeForSets,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"updates_available": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func createOsmanagementManagedInstanceManagement(d *schema.ResourceData, m interface{}) error {
	sync := &OsmanagementManagedInstanceManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OsManagementClient()

	return tfresource.CreateResource(d, sync)
}

func readOsmanagementManagedInstanceManagement(d *schema.ResourceData, m interface{}) error {
	sync := &OsmanagementManagedInstanceManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OsManagementClient()

	return tfresource.ReadResource(sync)
}

func updateOsmanagementManagedInstanceManagement(d *schema.ResourceData, m interface{}) error {
	sync := &OsmanagementManagedInstanceManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OsManagementClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteOsmanagementManagedInstanceManagement(d *schema.ResourceData, m interface{}) error {
	return nil
}

type OsmanagementManagedInstanceManagementResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_osmanagement.OsManagementClient
	Res                    *oci_osmanagement.GetManagedInstanceResponse
	DisableNotFoundRetries bool
}

func (s *OsmanagementManagedInstanceManagementResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *OsmanagementManagedInstanceManagementResourceCrud) Create() error {
	e := s.Get()
	if e != nil {
		return e
	}

	var managedInstanceId *string
	if id, ok := s.D.GetOkExists("managed_instance_id"); ok {
		tmp := id.(string)
		managedInstanceId = &tmp
	}

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

	// detach old existing parent software sources and attach new parent software sources present in config
	e = s.createParentSoftwareSource(managedInstanceId)
	if e != nil {
		return e
	}

	// Update resource response
	e = s.Get()
	if e != nil {
		return e
	}

	// detach old existing child software sources and attach new child software sources present in config
	e = s.createChildSoftwareSources(managedInstanceId)
	if e != nil {
		return e
	}

	// Update the resource response
	e = s.Get()
	if e != nil {
		return e
	}

	// detach old existing managed instance groups and attach new managed instance groups present in config
	e = s.createManagedInstanceGroups(managedInstanceId)
	if e != nil {
		return e
	}

	return nil
}

func (s *OsmanagementManagedInstanceManagementResourceCrud) Update() error {

	var managedInstanceId *string
	if id, ok := s.D.GetOkExists("managed_instance_id"); ok {
		tmp := id.(string)
		managedInstanceId = &tmp
	}

	defer func() {
		// get latest state of the instance
		err := s.Get()
		if err != nil {
			log.Printf("[ERROR] unable to invoke GET() after UPDATE '%v'", err)
		}
		// write latest state
		if err := s.SetData(); err != nil {
			log.Printf("[ERROR] unable to invoke setData() '%v'", err)
		}
	}()

	err := s.updateParentSoftwareSource(managedInstanceId)
	if err != nil {
		return err
	}
	err = s.updateChildSoftwareSources(managedInstanceId)
	if err != nil {
		return err
	}
	err = s.updateManagedInstanceGroups(managedInstanceId)
	if err != nil {
		return err
	}

	return s.Get()
}

func (s *OsmanagementManagedInstanceManagementResourceCrud) Get() error {
	request := oci_osmanagement.GetManagedInstanceRequest{}

	if managedInstanceId, ok := s.D.GetOkExists("managed_instance_id"); ok {
		tmp := managedInstanceId.(string)
		request.ManagedInstanceId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "osmanagement")

	response, err := s.Client.GetManagedInstance(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *OsmanagementManagedInstanceManagementResourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	childSoftwareSources := []interface{}{}
	for _, item := range s.Res.ChildSoftwareSources {
		childSoftwareSources = append(childSoftwareSources, SoftwareSourceIdToMap(&item))
	}
	s.D.Set("child_software_sources", schema.NewSet(softwareSourceHashCodeForSets, childSoftwareSources))

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.LastBoot != nil {
		s.D.Set("last_boot", *s.Res.LastBoot)
	}

	if s.Res.LastCheckin != nil {
		s.D.Set("last_checkin", *s.Res.LastCheckin)
	}

	managedInstanceGroups := []interface{}{}
	for _, item := range s.Res.ManagedInstanceGroups {
		managedInstanceGroups = append(managedInstanceGroups, IdToMap(item))
	}
	s.D.Set("managed_instance_groups", schema.NewSet(managedInstanceGroupsHashCodeForSets, managedInstanceGroups))

	if s.Res.OsKernelVersion != nil {
		s.D.Set("os_kernel_version", *s.Res.OsKernelVersion)
	}

	if s.Res.OsName != nil {
		s.D.Set("os_name", *s.Res.OsName)
	}

	if s.Res.OsVersion != nil {
		s.D.Set("os_version", *s.Res.OsVersion)
	}

	if s.Res.ParentSoftwareSource != nil {
		s.D.Set("parent_software_source", schema.NewSet(softwareSourceHashCodeForSets, []interface{}{SoftwareSourceIdToMap(s.Res.ParentSoftwareSource)}))
	} else {
		s.D.Set("parent_software_source", nil)
	}

	s.D.Set("status", s.Res.Status)

	if s.Res.UpdatesAvailable != nil {
		s.D.Set("updates_available", *s.Res.UpdatesAvailable)
	}

	return nil
}

func (s *OsmanagementManagedInstanceManagementResourceCrud) mapToSoftwareSourceId(fieldKeyFormat string) (oci_osmanagement.SoftwareSourceId, error) {
	result := oci_osmanagement.SoftwareSourceId{}

	if displayName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := displayName.(string)
		result.Name = &tmp
	}

	if id, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "id")); ok {
		tmp := id.(string)
		result.Id = &tmp
	}

	return result, nil
}

// Converting raw set data from state diff to DetachParentSoftwareSourceFromManagedInstanceDetails
func mapToDetachParentSoftwareSourceFromManagedInstanceDetails(parentSoftwareSource map[string]interface{}) oci_osmanagement.DetachParentSoftwareSourceFromManagedInstanceDetails {
	result := oci_osmanagement.DetachParentSoftwareSourceFromManagedInstanceDetails{}

	if id, ok := parentSoftwareSource["id"]; ok {
		tmp := id.(string)
		result.SoftwareSourceId = &tmp
	}

	return result
}

// Converting raw set data from state diff to AttachParentSoftwareSourceToManagedInstanceDetails
func mapToAttachParentSoftwareSourceToManagedInstanceDetails(parentSoftwareSource map[string]interface{}) oci_osmanagement.AttachParentSoftwareSourceToManagedInstanceDetails {
	result := oci_osmanagement.AttachParentSoftwareSourceToManagedInstanceDetails{}

	if id, ok := parentSoftwareSource["id"]; ok {
		tmp := id.(string)
		result.SoftwareSourceId = &tmp
	}

	return result
}

// Converting raw set data from state diff to DetachParentSoftwareSourceFromManagedInstanceDetails
func mapToDetachChildSoftwareSourceFromManagedInstanceDetails(childSoftwareSource map[string]interface{}) oci_osmanagement.DetachChildSoftwareSourceFromManagedInstanceDetails {
	result := oci_osmanagement.DetachChildSoftwareSourceFromManagedInstanceDetails{}

	if id, ok := childSoftwareSource["id"]; ok {
		tmp := id.(string)
		result.SoftwareSourceId = &tmp
	}

	return result
}

// Converting raw set data from state diff to AttachParentSoftwareSourceToManagedInstanceDetails
func mapToAttachChildSoftwareSourceToManagedInstanceDetails(childSoftwareSource map[string]interface{}) oci_osmanagement.AttachChildSoftwareSourceToManagedInstanceDetails {
	result := oci_osmanagement.AttachChildSoftwareSourceToManagedInstanceDetails{}

	if id, ok := childSoftwareSource["id"]; ok {
		tmp := id.(string)
		result.SoftwareSourceId = &tmp
	}

	return result
}

// Converting raw set data from state diff to DetachParentSoftwareSourceFromManagedInstanceDetails
func mapTomanagedInstanceGroupId(managedInstanceGroup map[string]interface{}) *string {
	var result *string
	if id, ok := managedInstanceGroup["id"]; ok {
		tmp := id.(string)
		result = &tmp
	}

	return result
}

func softwareSourceHashCodeForSets(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})
	if id, ok := m["id"]; ok && id != "" {
		buf.WriteString(fmt.Sprintf("%v-", id))
	}

	if name, ok := m["name"]; ok && name != "" {
		buf.WriteString(fmt.Sprintf("%v-", name))
	}
	return utils.GetStringHashcode(buf.String())
}

func (s *OsmanagementManagedInstanceManagementResourceCrud) createParentSoftwareSource(managedInstanceId *string) error {
	// detach old groups
	if s.Res.ParentSoftwareSource != nil {
		oldSourceToDetach := []interface{}{SoftwareSourceIdToMap(s.Res.ParentSoftwareSource)}
		for _, oldSourceMap := range oldSourceToDetach {
			opp := mapToDetachParentSoftwareSourceFromManagedInstanceDetails(oldSourceMap.(map[string]interface{}))
			detachParentSourceRequest := oci_osmanagement.DetachParentSoftwareSourceFromManagedInstanceRequest{}
			detachParentSourceRequest.ManagedInstanceId = managedInstanceId
			detachParentSourceRequest.DetachParentSoftwareSourceFromManagedInstanceDetails = opp

			_, detachErr := s.Client.DetachParentSoftwareSourceFromManagedInstance(context.Background(), detachParentSourceRequest)
			if detachErr != nil {
				return fmt.Errorf("failed to detach parent software source, error: %v", detachErr)
			}
		}
	}

	if newParentSource, ok := s.D.GetOkExists("parent_software_source"); ok {
		newSourceToAttach := newParentSource.(*schema.Set).List()
		for _, newSourceMap := range newSourceToAttach {
			opp := mapToAttachParentSoftwareSourceToManagedInstanceDetails(newSourceMap.(map[string]interface{}))
			attachParentSourceRequest := oci_osmanagement.AttachParentSoftwareSourceToManagedInstanceRequest{}
			attachParentSourceRequest.ManagedInstanceId = managedInstanceId
			attachParentSourceRequest.AttachParentSoftwareSourceToManagedInstanceDetails = opp

			_, attachErr := s.Client.AttachParentSoftwareSourceToManagedInstance(context.Background(), attachParentSourceRequest)
			if attachErr != nil {
				return fmt.Errorf("failed to attach parent software source, error: %v", attachErr)
			}
		}
	}
	return nil
}

func (s *OsmanagementManagedInstanceManagementResourceCrud) updateParentSoftwareSource(managedInstanceId *string) error {
	// check if any Update is made to parent software source
	if _, ok := s.D.GetOkExists("parent_software_source"); ok && s.D.HasChange("parent_software_source") {
		o, n := s.D.GetChange("parent_software_source")
		if o == nil {
			o = new(schema.Set)
		}
		if n == nil {
			n = new(schema.Set)
		}

		os := o.(*schema.Set)
		ns := n.(*schema.Set)

		newSourceToAttach := ns.Difference(os).List()
		oldSourceToDetach := os.Difference(ns).List()

		for _, oldSourceMap := range oldSourceToDetach {
			opp := mapToDetachParentSoftwareSourceFromManagedInstanceDetails(oldSourceMap.(map[string]interface{}))
			detachParentSourceRequest := oci_osmanagement.DetachParentSoftwareSourceFromManagedInstanceRequest{}
			detachParentSourceRequest.ManagedInstanceId = managedInstanceId
			detachParentSourceRequest.DetachParentSoftwareSourceFromManagedInstanceDetails = opp

			_, detachErr := s.Client.DetachParentSoftwareSourceFromManagedInstance(context.Background(), detachParentSourceRequest)
			if detachErr != nil {
				return fmt.Errorf("failed to detach parent software source, error: %v", detachErr)
			}
		}

		for _, newSourceMap := range newSourceToAttach {
			opp := mapToAttachParentSoftwareSourceToManagedInstanceDetails(newSourceMap.(map[string]interface{}))
			attachParentSourceRequest := oci_osmanagement.AttachParentSoftwareSourceToManagedInstanceRequest{}
			attachParentSourceRequest.ManagedInstanceId = managedInstanceId
			attachParentSourceRequest.AttachParentSoftwareSourceToManagedInstanceDetails = opp

			_, attachErr := s.Client.AttachParentSoftwareSourceToManagedInstance(context.Background(), attachParentSourceRequest)
			if attachErr != nil {
				return fmt.Errorf("failed to attach parent software source, error: %v", attachErr)
			}
		}
	}
	return nil
}

func (s *OsmanagementManagedInstanceManagementResourceCrud) createChildSoftwareSources(managedInstanceId *string) error {
	// detach old child software sources
	childSoftwareSources := []interface{}{}
	for _, item := range s.Res.ChildSoftwareSources {
		childSoftwareSources = append(childSoftwareSources, SoftwareSourceIdToMap(&item))
	}

	for _, oldChildSourceMap := range childSoftwareSources {
		opp := mapToDetachChildSoftwareSourceFromManagedInstanceDetails(oldChildSourceMap.(map[string]interface{}))
		detachChildSourceRequest := oci_osmanagement.DetachChildSoftwareSourceFromManagedInstanceRequest{}
		detachChildSourceRequest.ManagedInstanceId = managedInstanceId
		detachChildSourceRequest.DetachChildSoftwareSourceFromManagedInstanceDetails = opp

		_, detachErr := s.Client.DetachChildSoftwareSourceFromManagedInstance(context.Background(), detachChildSourceRequest)
		if detachErr != nil {
			return fmt.Errorf("failed to detach child software source, error: %v", detachErr)
		}
	}

	// attach new child software sources
	if newChildSources, ok := s.D.GetOkExists("child_software_sources"); ok {
		newSourceToAttach := newChildSources.(*schema.Set).List()
		for _, newSourceMap := range newSourceToAttach {
			opp := mapToAttachChildSoftwareSourceToManagedInstanceDetails(newSourceMap.(map[string]interface{}))
			attachChildSourceRequest := oci_osmanagement.AttachChildSoftwareSourceToManagedInstanceRequest{}
			attachChildSourceRequest.ManagedInstanceId = managedInstanceId
			attachChildSourceRequest.AttachChildSoftwareSourceToManagedInstanceDetails = opp

			_, attachErr := s.Client.AttachChildSoftwareSourceToManagedInstance(context.Background(), attachChildSourceRequest)
			if attachErr != nil {
				return fmt.Errorf("failed to attach child software source, error: %v", attachErr)
			}
		}
	}
	return nil
}

func (s *OsmanagementManagedInstanceManagementResourceCrud) updateChildSoftwareSources(managedInstanceId *string) error {
	if _, ok := s.D.GetOkExists("child_software_sources"); ok && s.D.HasChange("child_software_sources") {
		o, n := s.D.GetChange("child_software_sources")
		if o == nil {
			o = new(schema.Set)
		}
		if n == nil {
			n = new(schema.Set)
		}

		os := o.(*schema.Set)
		ns := n.(*schema.Set)

		newSourceToAttach := ns.Difference(os).List()
		oldSourceToDetach := os.Difference(ns).List()

		for _, oldSourceMap := range oldSourceToDetach {
			opp := mapToDetachChildSoftwareSourceFromManagedInstanceDetails(oldSourceMap.(map[string]interface{}))
			detachChildSourceRequest := oci_osmanagement.DetachChildSoftwareSourceFromManagedInstanceRequest{}
			detachChildSourceRequest.ManagedInstanceId = managedInstanceId
			detachChildSourceRequest.DetachChildSoftwareSourceFromManagedInstanceDetails = opp

			_, detachErr := s.Client.DetachChildSoftwareSourceFromManagedInstance(context.Background(), detachChildSourceRequest)
			if detachErr != nil {
				return fmt.Errorf("failed to detach child software source, error: %v", detachErr)
			}
		}

		for _, newSourceMap := range newSourceToAttach {
			opp := mapToAttachChildSoftwareSourceToManagedInstanceDetails(newSourceMap.(map[string]interface{}))
			attachChildSourceRequest := oci_osmanagement.AttachChildSoftwareSourceToManagedInstanceRequest{}
			attachChildSourceRequest.ManagedInstanceId = managedInstanceId
			attachChildSourceRequest.AttachChildSoftwareSourceToManagedInstanceDetails = opp

			_, attachErr := s.Client.AttachChildSoftwareSourceToManagedInstance(context.Background(), attachChildSourceRequest)
			if attachErr != nil {
				return fmt.Errorf("failed to attach child software source, error: %v", attachErr)
			}
		}
	}
	return nil
}

func (s *OsmanagementManagedInstanceManagementResourceCrud) createManagedInstanceGroups(managedInstanceId *string) error {
	// detach old managed groups
	oldGroups := []interface{}{}
	for _, item := range s.Res.ManagedInstanceGroups {
		oldGroups = append(oldGroups, IdToMap(item))
	}

	for _, oldGroupMap := range oldGroups {
		managedInstanceGroupId := mapTomanagedInstanceGroupId(oldGroupMap.(map[string]interface{}))
		detachManagedInstanceFromGroupRequest := oci_osmanagement.DetachManagedInstanceFromManagedInstanceGroupRequest{}
		detachManagedInstanceFromGroupRequest.ManagedInstanceId = managedInstanceId
		detachManagedInstanceFromGroupRequest.ManagedInstanceGroupId = managedInstanceGroupId

		_, detachErr := s.Client.DetachManagedInstanceFromManagedInstanceGroup(context.Background(), detachManagedInstanceFromGroupRequest)
		if detachErr != nil {
			return fmt.Errorf("failed to detach managed instance from managed instance Group request, error: %v", detachErr)
		}
	}

	// attach old managed groups
	if newManagedGroupsToAttach, ok := s.D.GetOkExists("managed_instance_groups"); ok {
		newSourceToAttach := newManagedGroupsToAttach.(*schema.Set).List()
		for _, newGroupMap := range newSourceToAttach {
			managedInstanceGroupId := mapTomanagedInstanceGroupId(newGroupMap.(map[string]interface{}))
			attachManagedInstanceToGroupRequest := oci_osmanagement.AttachManagedInstanceToManagedInstanceGroupRequest{}
			attachManagedInstanceToGroupRequest.ManagedInstanceId = managedInstanceId
			attachManagedInstanceToGroupRequest.ManagedInstanceGroupId = managedInstanceGroupId

			_, attachErr := s.Client.AttachManagedInstanceToManagedInstanceGroup(context.Background(), attachManagedInstanceToGroupRequest)
			if attachErr != nil {
				return fmt.Errorf("failed to attach managed instance to  managed instance Group, error: %v", attachErr)
			}
		}
	}
	return nil
}

func (s *OsmanagementManagedInstanceManagementResourceCrud) updateManagedInstanceGroups(managedInstanceId *string) error {
	// check if any Update is made to managed instance groups
	if _, ok := s.D.GetOkExists("managed_instance_groups"); ok && s.D.HasChange("managed_instance_groups") {
		o, n := s.D.GetChange("managed_instance_groups")
		if o == nil {
			o = new(schema.Set)
		}
		if n == nil {
			n = new(schema.Set)
		}

		os := o.(*schema.Set)
		ns := n.(*schema.Set)

		newSourceToAttach := ns.Difference(os).List()
		oldSourceToDetach := os.Difference(ns).List()

		for _, oldGroupMap := range oldSourceToDetach {
			managedInstanceGroupId := mapTomanagedInstanceGroupId(oldGroupMap.(map[string]interface{}))
			detachManagedInstanceFromGroupRequest := oci_osmanagement.DetachManagedInstanceFromManagedInstanceGroupRequest{}
			detachManagedInstanceFromGroupRequest.ManagedInstanceId = managedInstanceId
			detachManagedInstanceFromGroupRequest.ManagedInstanceGroupId = managedInstanceGroupId

			_, detachErr := s.Client.DetachManagedInstanceFromManagedInstanceGroup(context.Background(), detachManagedInstanceFromGroupRequest)
			if detachErr != nil {
				return fmt.Errorf("failed to detach managed instance from managed instance Group request, error: %v", detachErr)
			}
		}

		for _, newGroupMap := range newSourceToAttach {
			managedInstanceGroupId := mapTomanagedInstanceGroupId(newGroupMap.(map[string]interface{}))
			attachManagedInstanceToGroupRequest := oci_osmanagement.AttachManagedInstanceToManagedInstanceGroupRequest{}
			attachManagedInstanceToGroupRequest.ManagedInstanceId = managedInstanceId
			attachManagedInstanceToGroupRequest.ManagedInstanceGroupId = managedInstanceGroupId

			_, attachErr := s.Client.AttachManagedInstanceToManagedInstanceGroup(context.Background(), attachManagedInstanceToGroupRequest)
			if attachErr != nil {
				return fmt.Errorf("failed to attach parent software source, error: %v", attachErr)
			}
		}
	}
	return nil
}

func managedInstanceGroupsHashCodeForSets(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})
	if displayName, ok := m["display_name"]; ok && displayName != "" {
		buf.WriteString(fmt.Sprintf("%v-", displayName))
	}
	if id, ok := m["id"]; ok && id != "" {
		buf.WriteString(fmt.Sprintf("%v-", id))
	}
	return utils.GetStringHashcode(buf.String())
}

func (s *OsmanagementManagedInstanceManagementResourceCrud) mapToId(fieldKeyFormat string) (oci_osmanagement.Id, error) {
	result := oci_osmanagement.Id{}

	if displayName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "display_name")); ok {
		tmp := displayName.(string)
		result.DisplayName = &tmp
	}

	if id, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "id")); ok {
		tmp := id.(string)
		result.Id = &tmp
	}

	return result, nil
}
