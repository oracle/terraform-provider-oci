// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package core

import (
	"time"

	"github.com/MustWin/baremetal-sdk-go"
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/crud"
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/options"
)

type DrgAttachmentDatasourceCrud struct {
	crud.BaseCrud
	Res *baremetal.ListDrgAttachments
}

func (s *DrgAttachmentDatasourceCrud) Get() (e error) {
	compartmentID := s.D.Get("compartment_id").(string)

	opts := &baremetal.ListDrgAttachmentsOptions{}
	options.SetListOptions(s.D, &opts.ListOptions)
	if val, ok := s.D.GetOk("drg_id"); ok {
		opts.DrgID = val.(string)
	}
	if val, ok := s.D.GetOk("vcn_id"); ok {
		opts.VcnID = val.(string)
	}

	s.Res = &baremetal.ListDrgAttachments{
		DrgAttachments: []baremetal.DrgAttachment{},
	}

	for {
		var list *baremetal.ListDrgAttachments
		if list, e = s.Client.ListDrgAttachments(compartmentID, opts); e != nil {
			break
		}

		s.Res.DrgAttachments = append(s.Res.DrgAttachments, list.DrgAttachments...)

		if hasNextPage := options.SetNextPageOption(list.NextPage, &opts.ListOptions.PageListOptions); !hasNextPage {
			break
		}
	}

	return
}

func (s *DrgAttachmentDatasourceCrud) SetData() {
	if s.Res != nil {
		// Important, if you don't have an ID, make one up for your datasource
		// or things will end in tears
		s.D.SetId(time.Now().UTC().String())
		resources := []map[string]string{}
		for _, v := range s.Res.DrgAttachments {
			res := map[string]string{
				"compartment_id": v.CompartmentID,
				"display_name":   v.DisplayName,
				"drg_id":         v.DrgID,
				"id":             v.ID,
				"state":          v.State,
				"time_created":   v.TimeCreated.String(),
				"vcn_id":         v.VcnID,
			}
			resources = append(resources, res)
		}
		s.D.Set("drg_attachments", resources)
	}
	return
}
