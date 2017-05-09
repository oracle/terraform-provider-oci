// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/MustWin/baremetal-sdk-go"
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/oracle/terraform-provider-baremetal/client"
	"github.com/oracle/terraform-provider-baremetal/crud"
)

func LoadBalancerResource() *schema.Resource {
	return &schema.Resource{
		Create: createLoadBalancer,
		Read:   readLoadBalancer,
		Update: updateLoadBalancer,
		Delete: deleteLoadBalancer,
		Schema: map[string]*schema.Schema{
			// Required {
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"shape": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"subnet_ids": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			// }
			// Computed {
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ip_addresses": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
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

func createLoadBalancer(d *schema.ResourceData, m interface{}) (e error) {
	sync := &LoadBalancerResourceCrud{}
	sync.D = d
	sync.Client = m.(client.BareMetalClient)
	return crud.CreateResource(d, sync)
}

func readLoadBalancer(d *schema.ResourceData, m interface{}) (e error) {
	sync := &LoadBalancerResourceCrud{}
	sync.D = d
	sync.Client = m.(client.BareMetalClient)
	return crud.ReadResource(sync)
}

func updateLoadBalancer(d *schema.ResourceData, m interface{}) (e error) {
	sync := &LoadBalancerResourceCrud{}
	sync.D = d
	sync.Client = m.(client.BareMetalClient)
	return crud.UpdateResource(d, sync)
}

func deleteLoadBalancer(d *schema.ResourceData, m interface{}) (e error) {
	sync := &LoadBalancerResourceCrud{}
	sync.D = d
	sync.Client = m.(client.BareMetalClient)
	return crud.DeleteResource(d, sync)
}

// LoadBalancerResourceCrud wraps a baremetal.LoadBalancer to support crud
type LoadBalancerResourceCrud struct {
	crud.BaseCrud
	WorkRequest *baremetal.WorkRequest
	Resource    *baremetal.LoadBalancer
}

// ID delegates to the load balancer ID, falling back to the work request ID
func (s *LoadBalancerResourceCrud) ID() string {
	log.Printf("[DEBUG] lb.LoadBalancerResourceCrud.ID()")
	log.Printf("[DEBUG] lb.LoadBalancerResourceCrud.ID: Resource: %#v", s.Resource)
	if s.Resource != nil && s.Resource.ID != "" {
		log.Printf("[DEBUG] lb.LoadBalancerResourceCrud.ID: Resource.ID: %#v", s.Resource.ID)
		return s.Resource.ID
	}
	log.Printf("[DEBUG] lb.LoadBalancerResourceCrud.ID: WorkRequest: %#v", s.WorkRequest)
	if s.WorkRequest != nil {
		log.Printf("[DEBUG] lb.LoadBalancerResourceCrud.ID: WorkRequest.State: %s", s.WorkRequest.State)
		if s.WorkRequest.State == baremetal.WorkRequestSucceeded {
			log.Printf("[DEBUG] lb.LoadBalancerResourceCrud.ID: WorkRequest.LoadBalancerID: %#v", s.WorkRequest.LoadBalancerID)
			return s.WorkRequest.LoadBalancerID
		} else {
			log.Printf("[DEBUG] lb.LoadBalancerResourceCrud.ID: WorkRequest.ID: %s", s.WorkRequest.ID)
			return s.WorkRequest.ID
		}
	}
	log.Printf("[DEBUG] lb.LoadBalancerResourceCrud.ID: Resource & WorkRequest are nil, returning \"\"")
	return ""
}

// RefreshWorkRequest returns the last updated workRequest
func (s *LoadBalancerResourceCrud) RefreshWorkRequest() (*baremetal.WorkRequest, error) {
	if s.WorkRequest == nil {
		return nil, nil
	}
	wr, err := s.Client.GetWorkRequest(s.WorkRequest.ID, nil)
	if err != nil {
		return nil, err
	}
	s.WorkRequest = wr
	return wr, nil
}

// CreatedPending returns the resource states which qualify as "creating"
func (s *LoadBalancerResourceCrud) CreatedPending() []string {
	return []string{
		baremetal.ResourceWaitingForWorkRequest,
		baremetal.ResourceCreating,
		baremetal.WorkRequestAccepted,
		baremetal.WorkRequestInProgress,
	}
}

// CreatedTarget returns the resource states which qualify as "created"
func (s *LoadBalancerResourceCrud) CreatedTarget() []string {
	return []string{
		baremetal.ResourceActive,
	}
}

// DeletedPending returns the resource states which qualify as "deleting"
func (s *LoadBalancerResourceCrud) DeletedPending() []string {
	return []string{
		baremetal.ResourceWaitingForWorkRequest,
		baremetal.ResourceDeleting,
		baremetal.WorkRequestAccepted,
		baremetal.WorkRequestInProgress,
	}
}

// DeletedTarget returns the resource states which qualify as "deleted"
func (s *LoadBalancerResourceCrud) DeletedTarget() []string {
	return []string{baremetal.ResourceDeleted}
}

// Create makes a request to create a new load balancer from the resourceData
// It should leave the work request set up
func (s *LoadBalancerResourceCrud) Create() (e error) {
	sns := []string{}
	for _, v := range s.D.Get("subnet_ids").([]interface{}) {
		sns = append(sns, v.(string))
	}

	opts := &baremetal.CreateOptions{}
	opts.DisplayName = s.D.Get("display_name").(string)

	workReqID, e := s.Client.CreateLoadBalancer(
		nil,
		nil,
		s.D.Get("compartment_id").(string),
		nil,
		s.D.Get("shape").(string),
		sns,
		opts)

	if e != nil {
		return
	}
	s.WorkRequest, e = s.Client.GetWorkRequest(workReqID, nil)
	s.D.Set("state", s.WorkRequest.State)
	return
}

// Get makes a request to get the load balancer, populating s.Resource.
func (s *LoadBalancerResourceCrud) Get() (e error) {
	// key: {workRequestID} || {loadBalancerID}
	id := s.D.Id()
	log.Printf("[DEBUG] lb.LoadBalancerBackendResource.Get: ID: %#v", id)
	if id == "" {
		panic(fmt.Sprintf("LoadBalancer had empty ID: %#v Resource: %#V", s, s.Resource))
	}
	wr := s.WorkRequest
	log.Printf("[DEBUG] lb.LoadBalancerBackendResource.Get: WorkRequest: %#v", wr)
	state := s.D.Get("state").(string)
	log.Printf("[DEBUG] lb.LoadBalancerBackendResource.Get: State: %#v", state)

	// NOTE: if the id is for a work request, refresh its state and loadBalancerID. then refresh the load balancer
	if strings.HasPrefix(id, "ocid1.loadbalancerworkrequest.") {
		log.Printf("[DEBUG] lb.LoadBalancerBackendResource.Get: ID is for WorkRequest, refreshing")
		s.WorkRequest, e = s.Client.GetWorkRequest(id, nil)
		log.Printf("[DEBUG] lb.LoadBalancerBackendResource.Get: WorkRequest: %#v", s.WorkRequest)
		s.D.Set("state", s.WorkRequest.State)
		if s.WorkRequest.State == baremetal.WorkRequestSucceeded {
			id = s.WorkRequest.LoadBalancerID
			if id == "" {
				panic(fmt.Sprintf("WorkRequest had empty LoadBalancerID: %#v", s.WorkRequest))
			}
			s.D.SetId(id)
			// unset work request on success
			s.WorkRequest = nil
		} else {
			// We do not have a LoadBalancerID, so we short-circuit out
			return

		}
	}

	if !strings.HasPrefix(id, "ocid1.loadbalancer.") {
		panic(fmt.Sprintf("Cannot request loadbalancer with this ID, expected it to begin with \"ocid1.loadbalancer.\", but was: %#v", id))
	}
	log.Printf("[DEBUG] lb.LoadBalancerBackendResource.Get: ID: %#v", id)
	if id == "" {
		panic(fmt.Sprintf("LoadBalancer had empty ID: %#v Resource: %#V", s, s.Resource))
	}
	s.Resource, e = s.Client.GetLoadBalancer(id, nil)

	return
}

// Update makes a request to update the load balancer
func (s *LoadBalancerResourceCrud) Update() (e error) {
	opts := &baremetal.UpdateOptions{}
	if displayName, ok := s.D.GetOk("display_name"); ok {
		opts.DisplayName = displayName.(string)
	}

	var workReqID string
	workReqID, e = s.Client.UpdateLoadBalancer(s.D.Id(), opts)
	if e != nil {
		return
	}
	s.WorkRequest, e = s.Client.GetWorkRequest(workReqID, nil)
	return
}

// SetData populates the resourceData from the model
func (s *LoadBalancerResourceCrud) SetData() {
	// The first time this is called, we haven't actually fetched the resource yet, we just got a work request
	if s.Resource != nil && s.Resource.ID != "" {
		s.D.SetId(s.Resource.ID)
		s.D.Set("compartment_id", s.Resource.CompartmentID)
		s.D.Set("display_name", s.Resource.DisplayName)
		s.D.Set("shape", s.Resource.Shape)
		s.D.Set("subnet_ids", s.Resource.SubnetIDs)
		// Computed
		s.D.Set("id", s.Resource.ID)
		s.D.Set("state", s.Resource.State)
		s.D.Set("time_created", s.Resource.TimeCreated.String())
		ip_addresses := make([]string, len(s.Resource.IPAddresses))
		for i, ad := range s.Resource.IPAddresses {
			ip_addresses[i] = ad.IPAddress
		}
		s.D.Set("ip_addresses", ip_addresses)
	}
}

// Delete makes a request to delete the load balancer
func (s *LoadBalancerResourceCrud) Delete() (e error) {
	var workReqID string
	workReqID, e = s.Client.DeleteLoadBalancer(s.D.Id(), nil)
	if e != nil {
		return
	}
	s.WorkRequest, e = s.Client.GetWorkRequest(workReqID, nil)
	return
}
