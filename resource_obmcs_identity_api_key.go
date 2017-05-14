// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"github.com/MustWin/baremetal-sdk-go"
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/oracle/terraform-provider-baremetal/client"
	"github.com/oracle/terraform-provider-baremetal/crud"
	"errors"
	"regexp"
)

func APIKeyResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: crud.DefaultTimeout,
		Create:   createAPIKey,
		Read:     readAPIKey,
		Delete:   deleteAPIKey,
		Schema: map[string]*schema.Schema{
			"fingerprint": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"inactive_status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"key_value": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					r := regexp.MustCompile("\\s")
					strippedOld := r.ReplaceAllString(old, "")
					strippedNew := r.ReplaceAllString(new, "")
					if strippedOld == strippedNew {
						return true
					}
					return false
				},
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"user_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func createAPIKey(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	sync := &APIKeyResourceCrud{}
	sync.D = d
	sync.Client = client
	return crud.CreateResource(d, sync)
}

func readAPIKey(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	sync := &APIKeyResourceCrud{}
	sync.D = d
	sync.Client = client
	return crud.ReadResource(sync)
}

func deleteAPIKey(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	sync := &APIKeyResourceCrud{}
	sync.D = d
	sync.Client = client
	return crud.DeleteResource(d, sync)
}

type APIKeyResourceCrud struct {
	crud.BaseCrud
	Res *baremetal.APIKey
}

func (s *APIKeyResourceCrud) ID() string {
	return s.Res.KeyID
}

func (s *APIKeyResourceCrud) CreatedPending() []string {
	return []string{baremetal.ResourceCreating}
}

func (s *APIKeyResourceCrud) CreatedTarget() []string {
	return []string{baremetal.ResourceActive}
}

func (s *APIKeyResourceCrud) DeletedPending() []string {
	return []string{baremetal.ResourceDeleting}
}

func (s *APIKeyResourceCrud) DeletedTarget() []string {
	return []string{baremetal.ResourceDeleted}
}

func (s *APIKeyResourceCrud) State() string {
	return s.Res.State
}

func (s *APIKeyResourceCrud) Create() (e error) {
	userID := s.D.Get("user_id").(string)
	key := s.D.Get("key_value").(string)

	s.Res, e = s.Client.UploadAPIKey(userID, key, nil)

	return
}

func (s *APIKeyResourceCrud) Get() (e error) {
	userID := s.D.Get("user_id").(string)
	fingerprint := s.D.Get("fingerprint").(string)

	var res *baremetal.ListAPIKeyResponses
	if res, e = s.Client.ListAPIKeys(userID); e != nil {
		return
	}

	// The API does not provide a Get(user_id, fingerprint) method.
	// Loop through the list of keys and try to find by fingerprint.
	for _, val := range res.Keys {
		if val.Fingerprint == fingerprint {
			s.Res = &val
			return nil
		}
	}

	return errors.New("Specified APIKEY does not exist")
}

func (s *APIKeyResourceCrud) SetData() {
	s.D.Set("fingerprint", s.Res.Fingerprint)
	s.D.Set("key_value", s.Res.KeyValue)
	s.D.Set("state", s.Res.State)
	s.D.Set("time_created", s.Res.TimeCreated.String())
	s.D.Set("user_id", s.Res.UserID)
}

func (s *APIKeyResourceCrud) Delete() (e error) {
	userID := s.D.Get("user_id").(string)
	fingerprint := s.D.Get("fingerprint").(string)
	return s.Client.DeleteAPIKey(userID, fingerprint, nil)
}
