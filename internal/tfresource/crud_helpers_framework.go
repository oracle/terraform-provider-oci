// Copyright (c) 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package tfresource

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"reflect"
	"sort"
	"sync"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	oci_load_balancer "github.com/oracle/oci-go-sdk/v65/loadbalancer"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

type BaseCrudFW struct {
	Context          *context.Context
	Request          interface{}
	Response         interface{}
	RequestState     *tfsdk.State
	ResponseState    *tfsdk.State
	Mutex            *sync.Mutex
	OperationTimeout time.Duration
}

func (s *BaseCrudFW) GetOperationTimeout() time.Duration {
	return s.OperationTimeout
}

func (s *BaseCrudFW) VoidState() {
	s.ResponseState.RemoveResource(*s.Context)
}

func (s *BaseCrudFW) State() string {
	var stateValue string

	diags := s.ResponseState.GetAttribute(*s.Context, path.Root("state"), &stateValue)
	if diags.HasError() {
		return ""
	}
	return stateValue
}

// Default implementation, used in conjunction with State()
func (s *BaseCrudFW) setState(sync StatefulResource) error {
	v := reflectValueOf(sync).Elem()
	for _, key := range []string{"Res", "Resource", "WorkRequest"} {
		// Yes, this "valid"ation is terrible
		if resourceReferenceValue := v.FieldByName(key); resourceReferenceValue.IsValid() {
			if resourceValue := resourceReferenceValue.Elem(); resourceValue.IsValid() {
				// In rare cases, the kind for "Res" is an interface (e.g. if the resource itself is
				// a polymorphic type, opposed to a field on the resource). Use Elem() to get the value
				// the interface contains, otherwise the ".FieldByName()" method will throw.
				if resourceValue.Kind() == reflect.Interface {
					resourceValue = resourceValue.Elem()
				}

				if stateValue := resourceValue.FieldByName("LifecycleState"); stateValue.IsValid() {
					currentState := stateValue.String()
					log.Printf("[DEBUG] BaseCrud.setState: state: %#v", currentState)
					diags := s.ResponseState.SetAttribute(*s.Context, path.Root("state"), currentState)
					if diags.HasError() {
						return DiagnosticsToError(diags)
					} else {
						return nil
					}
				} else if stateValue := resourceValue.FieldByName("State"); stateValue.IsValid() {
					currentState := stateValue.String()
					log.Printf("[DEBUG] BaseCrud.setState: state: %#v", currentState)
					diags := s.ResponseState.SetAttribute(*s.Context, path.Root("state"), currentState)
					if diags.HasError() {
						return DiagnosticsToError(diags)
					} else {
						return nil
					}
				}
			}
		}
	}

	return nil
}

func GenerateFrameworkDataSourceHashID(idPrefix string, ctx context.Context, state tfsdk.State) string {
	if ctx == nil {
		return ""
	}

	attrs := state.Schema.GetAttributes()

	var buf bytes.Buffer
	// sort keys of the map
	keys := make([]string, 0, len(attrs))
	for key := range attrs {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	for _, key := range keys {
		// parse schema field is user input
		attr := attrs[key]
		if !(attr.IsRequired() || attr.IsOptional()) {
			continue
		}

		// Ignoring TypeList, TypeSet and TypeMap
		if isCollectionType(attr) {
			continue
		}

		var value types.String
		diags := state.GetAttribute(ctx, path.Root(key), &value)
		if diags.HasError() {
			log.Printf("[DEBUG] Error while reading attribute from state file %s", DiagnosticsToError(diags))
			return ""
		}

		valueStr := value.ValueString()
		if valueStr != "" {
			buf.WriteString(fmt.Sprintf("%v-", valueStr))
		}
	}
	return fmt.Sprintf("%s%d", idPrefix, utils.GetStringHashcode(buf.String()))
}

func isCollectionType(attr schema.Attribute) bool {
	if _, ok := attr.(schema.ListAttribute); ok {
		return true
	} else if _, ok := attr.(schema.SetAttribute); ok {
		return true
	} else if _, ok := attr.(schema.MapAttribute); ok {
		return true
	}
	return false
}

func CreateResourceFw(sync ResourceCreator) error {
	if synchronizedResource, ok := sync.(SynchronizedResource); ok {
		if mutex := synchronizedResource.GetMutex(); mutex != nil {
			mutex.Lock()
			defer mutex.Unlock()
		}
	}

	if e := sync.Create(); e != nil {
		return HandleError(sync, e)
	}

	if stateful, ok := sync.(StatefullyCreatedResourceFw); ok {
		if e := waitForStateRefreshVar(stateful, stateful.GetOperationTimeout(), "creation", stateful.CreatedPending(), stateful.CreatedTarget()); e != nil {
			if stateful.State() == FAILED {
				// Remove resource from state if asynchronous work request has failed so that it is recreated on next apply
				// TODO: automatic retry on WorkRequestFailed
				sync.VoidState()
			}

			//We need to SetData() here because if there is an error or timeout in the wait for state after the Create() was successful we want to store the resource in the statefile to avoid dangling resources
			if setDataErr := sync.SetData(); setDataErr != nil {
				log.Printf("[ERROR] error setting data after WaitForStateRefresh() error: %v", setDataErr)
			}

			return e
		}
	}

	//d.SetId(sync.ID())
	if e := sync.SetData(); e != nil {
		return e
	}

	if ew, waitOK := sync.(ExtraWaitPostCreateDelete); waitOK {
		time.Sleep(ew.ExtraWaitPostCreateDelete())
	}

	return nil
}

func ReadResourceFw(sync ResourceReader) error {
	if e := sync.Get(); e != nil {
		log.Printf("ERROR IN GET: %v\n", e.Error())
		handleMissingResourceError(sync, &e)
		return HandleError(sync, e)
	}

	if e := sync.SetData(); e != nil {
		return e
	}

	// Remove resource from state if it has been terminated so that it is recreated on next apply
	if dr, ok := sync.(StatefullyDeletedResource); ok {
		for _, target := range dr.DeletedTarget() {
			if dr.State() == target && dr.State() != string(oci_load_balancer.WorkRequestLifecycleStateSucceeded) {
				dr.VoidState()
				return nil
			}
		}
	}

	return nil
}

func UpdateResourceFw(sync ResourceUpdater) error {
	if synchronizedResource, ok := sync.(SynchronizedResource); ok {
		if mutex := synchronizedResource.GetMutex(); mutex != nil {
			mutex.Lock()
			defer mutex.Unlock()
		}
	}

	//d.Partial(true)
	if e := sync.Update(); e != nil {

		return HandleError(sync, e)
	}
	//d.Partial(false)

	if stateful, ok := sync.(StatefullyUpdatedResourceFw); ok {
		if e := waitForStateRefreshVar(stateful, stateful.GetOperationTimeout(), "update", stateful.UpdatedPending(), stateful.UpdatedTarget()); e != nil {

			return e
		}
	}

	if e := sync.SetData(); e != nil {
		return e
	}

	return nil
}

// DeleteResource requests a Delete(). If the resource deletes
// statefully (not immediately), poll State to ensure:
// () -> Pending -> Deleted.
// Finally, sets the ResourceData state to empty.
func DeleteResourceFw(sync ResourceDeleter) error {
	if synchronizedResource, ok := sync.(SynchronizedResource); ok {
		if mutex := synchronizedResource.GetMutex(); mutex != nil {
			mutex.Lock()
			defer mutex.Unlock()
		}
	}

	if e := sync.Delete(); e != nil {
		handleMissingResourceError(sync, &e)
		return HandleError(sync, e)
	}

	if stateful, ok := sync.(StatefullyDeletedResourceFw); ok {
		if e := waitForStateRefreshVar(stateful, stateful.GetOperationTimeout(), "deletion", stateful.DeletedPending(), stateful.DeletedTarget()); e != nil {
			handleMissingResourceError(sync, &e)
			return e
		}
	}

	if ew, waitOK := sync.(ExtraWaitPostCreateDelete); waitOK {
		time.Sleep(ew.ExtraWaitPostCreateDelete())
	}

	if ew, waitOK := sync.(ExtraWaitPostDelete); waitOK {
		time.Sleep(ew.ExtraWaitPostDelete())
	}

	return nil
}

// Helper function to wait for Update to reach terminal state before doing another Update
// Useful in situations where more than one Update is needed and prior Update needs to complete
func WaitForUpdatedStateFw(sync ResourceUpdater) error {
	if stateful, ok := sync.(StatefullyUpdatedResourceFw); ok {
		if e := waitForStateRefreshVar(stateful, stateful.GetOperationTimeout(), "update", stateful.UpdatedPending(), stateful.UpdatedTarget()); e != nil {
			return e
		}
	}

	return nil
}
