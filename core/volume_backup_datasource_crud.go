// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package core

import (
	"time"

	"github.com/MustWin/baremetal-sdk-go"

	"github.com/oracle/terraform-provider-baremetal/options"

	"github.com/oracle/terraform-provider-baremetal/crud"
)

type VolumeBackupDatasourceCrud struct {
	crud.BaseCrud
	Res *baremetal.ListVolumeBackups
}

func (s *VolumeBackupDatasourceCrud) Get() (e error) {
	compartmentID := s.D.Get("compartment_id").(string)

	opts := &baremetal.ListBackupsOptions{}
	options.SetListOptions(s.D, &opts.ListOptions)
	if val, ok := s.D.GetOk("volume_id"); ok {
		opts.VolumeID = val.(string)
	}

	s.Res = &baremetal.ListVolumeBackups{
		VolumeBackups: []baremetal.VolumeBackup{},
	}

	for {
		var list *baremetal.ListVolumeBackups
		if list, e = s.Client.ListVolumeBackups(compartmentID, opts); e != nil {
			break
		}

		s.Res.VolumeBackups = append(s.Res.VolumeBackups, list.VolumeBackups...)

		if hasNextPage := options.SetNextPageOption(list.NextPage, &opts.ListOptions.PageListOptions); !hasNextPage {
			break
		}
	}

	return
}

func (s *VolumeBackupDatasourceCrud) SetData() {
	if s.Res != nil {
		s.D.SetId(time.Now().UTC().String())
		volumes := []map[string]interface{}{}
		for _, v := range s.Res.VolumeBackups {
			vol := map[string]interface{}{
				"compartment_id":        v.CompartmentID,
				"display_name":          v.DisplayName,
				"id":                    v.ID,
				"state":                 v.State,
				"size_in_mbs":           v.SizeInMBs,
				"time_created":          v.TimeCreated.String(),
				"time_request_received": v.TimeRequestReceived.String(),
				"unique_size_in_mbs":    v.UniqueSizeInMBs,
				"volume_id":             v.VolumeID,
			}
			volumes = append(volumes, vol)
		}
		s.D.Set("volume_backups", volumes)
	}
	return
}
