// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Resource Scheduler API
//
// Use the Resource scheduler API to manage schedules, to perform actions on a collection of resources.
//

package resourcescheduler

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Parameters This is the user input parameters to use when acting on the resource.
//
//	{
//	    "parameters": [
//	        {
//	            "parameterType": "BODY",
//	            "value": {
//	                "ip": "192.168.44.44",
//	                "memory": "1024",
//	                "synced_folders": [
//	                    {
//	                        "host_path": "data/",
//	                        "guest_path": "/var/www",
//	                        "type": "default"
//	                    }
//	                ],
//	                "forwarded_ports": []
//	            }
//	        },
//	        {
//	            "parameterType": "PATH",
//	            "value": {
//	                "compartmentId": "ocid1.compartment.oc1..xxxxx",
//	                "instanceId": "ocid1.vcn.oc1..yyyy"
//	            }
//	        },
//	        {
//	            "parameterType": "QUERY",
//	            "value": {
//	                "limit": "10",
//	                "tenantId": "ocid1.tenant.oc1..zzzz"
//	            }
//	        },
//	        {
//	            "parameterType": "HEADER",
//	            "value": {
//	              "token": "xxxx"
//	            }
//	        }
//	    ]
//	}
type Parameters struct {

	// This is an array of user input parameters.
	Parameters []Parameter `mandatory:"false" json:"parameters"`
}

func (m Parameters) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Parameters) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *Parameters) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Parameters []parameter `json:"parameters"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Parameters = make([]Parameter, len(model.Parameters))
	for i, n := range model.Parameters {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.Parameters[i] = nn.(Parameter)
		} else {
			m.Parameters[i] = nil
		}
	}
	return
}
