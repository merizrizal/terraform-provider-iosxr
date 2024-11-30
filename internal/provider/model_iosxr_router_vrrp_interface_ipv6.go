// Copyright © 2023 Cisco Systems, Inc. and its affiliates.
// All rights reserved.
//
// Licensed under the Mozilla Public License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	https://mozilla.org/MPL/2.0/
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: MPL-2.0

// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"reflect"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type RouterVRRPInterfaceIPv6 struct {
	Device                         types.String                             `tfsdk:"device"`
	Id                             types.String                             `tfsdk:"id"`
	DeleteMode                     types.String                             `tfsdk:"delete_mode"`
	InterfaceName                  types.String                             `tfsdk:"interface_name"`
	VrrpId                         types.Int64                              `tfsdk:"vrrp_id"`
	GlobalAddresses                types.String                             `tfsdk:"global_addresses"`
	AddressLinklocal               types.String                             `tfsdk:"address_linklocal"`
	AddressLinklocalAutoconfig     types.Bool                               `tfsdk:"address_linklocal_autoconfig"`
	Priority                       types.Int64                              `tfsdk:"priority"`
	Name                           types.String                             `tfsdk:"name"`
	TimerAdvertisementSeconds      types.Int64                              `tfsdk:"timer_advertisement_seconds"`
	TimerAdvertisementMilliseconds types.Int64                              `tfsdk:"timer_advertisement_milliseconds"`
	TimerForce                     types.Bool                               `tfsdk:"timer_force"`
	PreemptDisable                 types.Bool                               `tfsdk:"preempt_disable"`
	PreemptDelay                   types.Int64                              `tfsdk:"preempt_delay"`
	AcceptModeDisable              types.Bool                               `tfsdk:"accept_mode_disable"`
	TrackInterfaces                []RouterVRRPInterfaceIPv6TrackInterfaces `tfsdk:"track_interfaces"`
	TrackObjects                   []RouterVRRPInterfaceIPv6TrackObjects    `tfsdk:"track_objects"`
	BfdFastDetectPeerIpv6          types.String                             `tfsdk:"bfd_fast_detect_peer_ipv6"`
}

type RouterVRRPInterfaceIPv6Data struct {
	Device                         types.String                             `tfsdk:"device"`
	Id                             types.String                             `tfsdk:"id"`
	InterfaceName                  types.String                             `tfsdk:"interface_name"`
	VrrpId                         types.Int64                              `tfsdk:"vrrp_id"`
	GlobalAddresses                types.String                             `tfsdk:"global_addresses"`
	AddressLinklocal               types.String                             `tfsdk:"address_linklocal"`
	AddressLinklocalAutoconfig     types.Bool                               `tfsdk:"address_linklocal_autoconfig"`
	Priority                       types.Int64                              `tfsdk:"priority"`
	Name                           types.String                             `tfsdk:"name"`
	TimerAdvertisementSeconds      types.Int64                              `tfsdk:"timer_advertisement_seconds"`
	TimerAdvertisementMilliseconds types.Int64                              `tfsdk:"timer_advertisement_milliseconds"`
	TimerForce                     types.Bool                               `tfsdk:"timer_force"`
	PreemptDisable                 types.Bool                               `tfsdk:"preempt_disable"`
	PreemptDelay                   types.Int64                              `tfsdk:"preempt_delay"`
	AcceptModeDisable              types.Bool                               `tfsdk:"accept_mode_disable"`
	TrackInterfaces                []RouterVRRPInterfaceIPv6TrackInterfaces `tfsdk:"track_interfaces"`
	TrackObjects                   []RouterVRRPInterfaceIPv6TrackObjects    `tfsdk:"track_objects"`
	BfdFastDetectPeerIpv6          types.String                             `tfsdk:"bfd_fast_detect_peer_ipv6"`
}
type RouterVRRPInterfaceIPv6TrackInterfaces struct {
	InterfaceName     types.String `tfsdk:"interface_name"`
	PriorityDecrement types.Int64  `tfsdk:"priority_decrement"`
}
type RouterVRRPInterfaceIPv6TrackObjects struct {
	ObjectName        types.String `tfsdk:"object_name"`
	PriorityDecrement types.Int64  `tfsdk:"priority_decrement"`
}

func (data RouterVRRPInterfaceIPv6) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XR-um-router-vrrp-cfg:/router/vrrp/interfaces/interface[interface-name=%s]/address-family/ipv6/vrrps/vrrp[vrrp-id=%d]", data.InterfaceName.ValueString(), data.VrrpId.ValueInt64())
}

func (data RouterVRRPInterfaceIPv6Data) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XR-um-router-vrrp-cfg:/router/vrrp/interfaces/interface[interface-name=%s]/address-family/ipv6/vrrps/vrrp[vrrp-id=%d]", data.InterfaceName.ValueString(), data.VrrpId.ValueInt64())
}

func (data RouterVRRPInterfaceIPv6) toBody(ctx context.Context) string {
	body := "{}"
	if !data.VrrpId.IsNull() && !data.VrrpId.IsUnknown() {
		body, _ = sjson.Set(body, "vrrp-id", strconv.FormatInt(data.VrrpId.ValueInt64(), 10))
	}
	if !data.GlobalAddresses.IsNull() && !data.GlobalAddresses.IsUnknown() {
		body, _ = sjson.Set(body, "address.global.global-address.address", data.GlobalAddresses.ValueString())
	}
	if !data.AddressLinklocal.IsNull() && !data.AddressLinklocal.IsUnknown() {
		body, _ = sjson.Set(body, "address.linklocal.linklocal-address", data.AddressLinklocal.ValueString())
	}
	if !data.AddressLinklocalAutoconfig.IsNull() && !data.AddressLinklocalAutoconfig.IsUnknown() {
		if data.AddressLinklocalAutoconfig.ValueBool() {
			body, _ = sjson.Set(body, "address.linklocal.autoconfig", map[string]string{})
		}
	}
	if !data.Priority.IsNull() && !data.Priority.IsUnknown() {
		body, _ = sjson.Set(body, "priority", strconv.FormatInt(data.Priority.ValueInt64(), 10))
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	if !data.TimerAdvertisementSeconds.IsNull() && !data.TimerAdvertisementSeconds.IsUnknown() {
		body, _ = sjson.Set(body, "timer.advertisement-time-in-seconds", strconv.FormatInt(data.TimerAdvertisementSeconds.ValueInt64(), 10))
	}
	if !data.TimerAdvertisementMilliseconds.IsNull() && !data.TimerAdvertisementMilliseconds.IsUnknown() {
		body, _ = sjson.Set(body, "timer.advertisement-time-in-milliseconds", strconv.FormatInt(data.TimerAdvertisementMilliseconds.ValueInt64(), 10))
	}
	if !data.TimerForce.IsNull() && !data.TimerForce.IsUnknown() {
		if data.TimerForce.ValueBool() {
			body, _ = sjson.Set(body, "timer.force", map[string]string{})
		}
	}
	if !data.PreemptDisable.IsNull() && !data.PreemptDisable.IsUnknown() {
		if data.PreemptDisable.ValueBool() {
			body, _ = sjson.Set(body, "preempt.disable", map[string]string{})
		}
	}
	if !data.PreemptDelay.IsNull() && !data.PreemptDelay.IsUnknown() {
		body, _ = sjson.Set(body, "preempt.delay", strconv.FormatInt(data.PreemptDelay.ValueInt64(), 10))
	}
	if !data.AcceptModeDisable.IsNull() && !data.AcceptModeDisable.IsUnknown() {
		if data.AcceptModeDisable.ValueBool() {
			body, _ = sjson.Set(body, "accept-mode.disable", map[string]string{})
		}
	}
	if !data.BfdFastDetectPeerIpv6.IsNull() && !data.BfdFastDetectPeerIpv6.IsUnknown() {
		body, _ = sjson.Set(body, "bfd.fast-detect.peer.ipv6", data.BfdFastDetectPeerIpv6.ValueString())
	}
	if len(data.TrackInterfaces) > 0 {
		body, _ = sjson.Set(body, "track.interfaces.interface", []interface{}{})
		for index, item := range data.TrackInterfaces {
			if !item.InterfaceName.IsNull() && !item.InterfaceName.IsUnknown() {
				body, _ = sjson.Set(body, "track.interfaces.interface"+"."+strconv.Itoa(index)+"."+"interface-name", item.InterfaceName.ValueString())
			}
			if !item.PriorityDecrement.IsNull() && !item.PriorityDecrement.IsUnknown() {
				body, _ = sjson.Set(body, "track.interfaces.interface"+"."+strconv.Itoa(index)+"."+"priority-decrement", strconv.FormatInt(item.PriorityDecrement.ValueInt64(), 10))
			}
		}
	}
	if len(data.TrackObjects) > 0 {
		body, _ = sjson.Set(body, "track.objects.object", []interface{}{})
		for index, item := range data.TrackObjects {
			if !item.ObjectName.IsNull() && !item.ObjectName.IsUnknown() {
				body, _ = sjson.Set(body, "track.objects.object"+"."+strconv.Itoa(index)+"."+"object-name", item.ObjectName.ValueString())
			}
			if !item.PriorityDecrement.IsNull() && !item.PriorityDecrement.IsUnknown() {
				body, _ = sjson.Set(body, "track.objects.object"+"."+strconv.Itoa(index)+"."+"priority-decrement", strconv.FormatInt(item.PriorityDecrement.ValueInt64(), 10))
			}
		}
	}
	return body
}

func (data *RouterVRRPInterfaceIPv6) updateFromBody(ctx context.Context, res []byte) {
	if value := gjson.GetBytes(res, "address.global.global-address.address"); value.Exists() && !data.GlobalAddresses.IsNull() {
		data.GlobalAddresses = types.StringValue(value.String())
	} else {
		data.GlobalAddresses = types.StringNull()
	}
	if value := gjson.GetBytes(res, "address.linklocal.linklocal-address"); value.Exists() && !data.AddressLinklocal.IsNull() {
		data.AddressLinklocal = types.StringValue(value.String())
	} else {
		data.AddressLinklocal = types.StringNull()
	}
	if value := gjson.GetBytes(res, "address.linklocal.autoconfig"); !data.AddressLinklocalAutoconfig.IsNull() {
		if value.Exists() {
			data.AddressLinklocalAutoconfig = types.BoolValue(true)
		} else {
			data.AddressLinklocalAutoconfig = types.BoolValue(false)
		}
	} else {
		data.AddressLinklocalAutoconfig = types.BoolNull()
	}
	if value := gjson.GetBytes(res, "priority"); value.Exists() && !data.Priority.IsNull() {
		data.Priority = types.Int64Value(value.Int())
	} else {
		data.Priority = types.Int64Null()
	}
	if value := gjson.GetBytes(res, "name"); value.Exists() && !data.Name.IsNull() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := gjson.GetBytes(res, "timer.advertisement-time-in-seconds"); value.Exists() && !data.TimerAdvertisementSeconds.IsNull() {
		data.TimerAdvertisementSeconds = types.Int64Value(value.Int())
	} else {
		data.TimerAdvertisementSeconds = types.Int64Null()
	}
	if value := gjson.GetBytes(res, "timer.advertisement-time-in-milliseconds"); value.Exists() && !data.TimerAdvertisementMilliseconds.IsNull() {
		data.TimerAdvertisementMilliseconds = types.Int64Value(value.Int())
	} else {
		data.TimerAdvertisementMilliseconds = types.Int64Null()
	}
	if value := gjson.GetBytes(res, "timer.force"); !data.TimerForce.IsNull() {
		if value.Exists() {
			data.TimerForce = types.BoolValue(true)
		} else {
			data.TimerForce = types.BoolValue(false)
		}
	} else {
		data.TimerForce = types.BoolNull()
	}
	if value := gjson.GetBytes(res, "preempt.disable"); !data.PreemptDisable.IsNull() {
		if value.Exists() {
			data.PreemptDisable = types.BoolValue(true)
		} else {
			data.PreemptDisable = types.BoolValue(false)
		}
	} else {
		data.PreemptDisable = types.BoolNull()
	}
	if value := gjson.GetBytes(res, "preempt.delay"); value.Exists() && !data.PreemptDelay.IsNull() {
		data.PreemptDelay = types.Int64Value(value.Int())
	} else {
		data.PreemptDelay = types.Int64Null()
	}
	if value := gjson.GetBytes(res, "accept-mode.disable"); !data.AcceptModeDisable.IsNull() {
		if value.Exists() {
			data.AcceptModeDisable = types.BoolValue(true)
		} else {
			data.AcceptModeDisable = types.BoolValue(false)
		}
	} else {
		data.AcceptModeDisable = types.BoolNull()
	}
	for i := range data.TrackInterfaces {
		keys := [...]string{"interface-name"}
		keyValues := [...]string{data.TrackInterfaces[i].InterfaceName.ValueString()}

		var r gjson.Result
		gjson.GetBytes(res, "track.interfaces.interface").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					if v.Get(keys[ik]).String() == keyValues[ik] {
						found = true
						continue
					}
					found = false
					break
				}
				if found {
					r = v
					return false
				}
				return true
			},
		)
		if value := r.Get("interface-name"); value.Exists() && !data.TrackInterfaces[i].InterfaceName.IsNull() {
			data.TrackInterfaces[i].InterfaceName = types.StringValue(value.String())
		} else {
			data.TrackInterfaces[i].InterfaceName = types.StringNull()
		}
		if value := r.Get("priority-decrement"); value.Exists() && !data.TrackInterfaces[i].PriorityDecrement.IsNull() {
			data.TrackInterfaces[i].PriorityDecrement = types.Int64Value(value.Int())
		} else {
			data.TrackInterfaces[i].PriorityDecrement = types.Int64Null()
		}
	}
	for i := range data.TrackObjects {
		keys := [...]string{"object-name"}
		keyValues := [...]string{data.TrackObjects[i].ObjectName.ValueString()}

		var r gjson.Result
		gjson.GetBytes(res, "track.objects.object").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					if v.Get(keys[ik]).String() == keyValues[ik] {
						found = true
						continue
					}
					found = false
					break
				}
				if found {
					r = v
					return false
				}
				return true
			},
		)
		if value := r.Get("object-name"); value.Exists() && !data.TrackObjects[i].ObjectName.IsNull() {
			data.TrackObjects[i].ObjectName = types.StringValue(value.String())
		} else {
			data.TrackObjects[i].ObjectName = types.StringNull()
		}
		if value := r.Get("priority-decrement"); value.Exists() && !data.TrackObjects[i].PriorityDecrement.IsNull() {
			data.TrackObjects[i].PriorityDecrement = types.Int64Value(value.Int())
		} else {
			data.TrackObjects[i].PriorityDecrement = types.Int64Null()
		}
	}
	if value := gjson.GetBytes(res, "bfd.fast-detect.peer.ipv6"); value.Exists() && !data.BfdFastDetectPeerIpv6.IsNull() {
		data.BfdFastDetectPeerIpv6 = types.StringValue(value.String())
	} else {
		data.BfdFastDetectPeerIpv6 = types.StringNull()
	}
}

func (data *RouterVRRPInterfaceIPv6) fromBody(ctx context.Context, res []byte) {
	if value := gjson.GetBytes(res, "address.global.global-address.address"); value.Exists() {
		data.GlobalAddresses = types.StringValue(value.String())
	}
	if value := gjson.GetBytes(res, "address.linklocal.linklocal-address"); value.Exists() {
		data.AddressLinklocal = types.StringValue(value.String())
	}
	if value := gjson.GetBytes(res, "address.linklocal.autoconfig"); value.Exists() {
		data.AddressLinklocalAutoconfig = types.BoolValue(true)
	} else {
		data.AddressLinklocalAutoconfig = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "priority"); value.Exists() {
		data.Priority = types.Int64Value(value.Int())
	}
	if value := gjson.GetBytes(res, "name"); value.Exists() {
		data.Name = types.StringValue(value.String())
	}
	if value := gjson.GetBytes(res, "timer.advertisement-time-in-seconds"); value.Exists() {
		data.TimerAdvertisementSeconds = types.Int64Value(value.Int())
	}
	if value := gjson.GetBytes(res, "timer.advertisement-time-in-milliseconds"); value.Exists() {
		data.TimerAdvertisementMilliseconds = types.Int64Value(value.Int())
	}
	if value := gjson.GetBytes(res, "timer.force"); value.Exists() {
		data.TimerForce = types.BoolValue(true)
	} else {
		data.TimerForce = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "preempt.disable"); value.Exists() {
		data.PreemptDisable = types.BoolValue(true)
	} else {
		data.PreemptDisable = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "preempt.delay"); value.Exists() {
		data.PreemptDelay = types.Int64Value(value.Int())
	}
	if value := gjson.GetBytes(res, "accept-mode.disable"); value.Exists() {
		data.AcceptModeDisable = types.BoolValue(true)
	} else {
		data.AcceptModeDisable = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "track.interfaces.interface"); value.Exists() {
		data.TrackInterfaces = make([]RouterVRRPInterfaceIPv6TrackInterfaces, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := RouterVRRPInterfaceIPv6TrackInterfaces{}
			if cValue := v.Get("interface-name"); cValue.Exists() {
				item.InterfaceName = types.StringValue(cValue.String())
			}
			if cValue := v.Get("priority-decrement"); cValue.Exists() {
				item.PriorityDecrement = types.Int64Value(cValue.Int())
			}
			data.TrackInterfaces = append(data.TrackInterfaces, item)
			return true
		})
	}
	if value := gjson.GetBytes(res, "track.objects.object"); value.Exists() {
		data.TrackObjects = make([]RouterVRRPInterfaceIPv6TrackObjects, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := RouterVRRPInterfaceIPv6TrackObjects{}
			if cValue := v.Get("object-name"); cValue.Exists() {
				item.ObjectName = types.StringValue(cValue.String())
			}
			if cValue := v.Get("priority-decrement"); cValue.Exists() {
				item.PriorityDecrement = types.Int64Value(cValue.Int())
			}
			data.TrackObjects = append(data.TrackObjects, item)
			return true
		})
	}
	if value := gjson.GetBytes(res, "bfd.fast-detect.peer.ipv6"); value.Exists() {
		data.BfdFastDetectPeerIpv6 = types.StringValue(value.String())
	}
}

func (data *RouterVRRPInterfaceIPv6Data) fromBody(ctx context.Context, res []byte) {
	if value := gjson.GetBytes(res, "address.global.global-address.address"); value.Exists() {
		data.GlobalAddresses = types.StringValue(value.String())
	}
	if value := gjson.GetBytes(res, "address.linklocal.linklocal-address"); value.Exists() {
		data.AddressLinklocal = types.StringValue(value.String())
	}
	if value := gjson.GetBytes(res, "address.linklocal.autoconfig"); value.Exists() {
		data.AddressLinklocalAutoconfig = types.BoolValue(true)
	} else {
		data.AddressLinklocalAutoconfig = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "priority"); value.Exists() {
		data.Priority = types.Int64Value(value.Int())
	}
	if value := gjson.GetBytes(res, "name"); value.Exists() {
		data.Name = types.StringValue(value.String())
	}
	if value := gjson.GetBytes(res, "timer.advertisement-time-in-seconds"); value.Exists() {
		data.TimerAdvertisementSeconds = types.Int64Value(value.Int())
	}
	if value := gjson.GetBytes(res, "timer.advertisement-time-in-milliseconds"); value.Exists() {
		data.TimerAdvertisementMilliseconds = types.Int64Value(value.Int())
	}
	if value := gjson.GetBytes(res, "timer.force"); value.Exists() {
		data.TimerForce = types.BoolValue(true)
	} else {
		data.TimerForce = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "preempt.disable"); value.Exists() {
		data.PreemptDisable = types.BoolValue(true)
	} else {
		data.PreemptDisable = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "preempt.delay"); value.Exists() {
		data.PreemptDelay = types.Int64Value(value.Int())
	}
	if value := gjson.GetBytes(res, "accept-mode.disable"); value.Exists() {
		data.AcceptModeDisable = types.BoolValue(true)
	} else {
		data.AcceptModeDisable = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "track.interfaces.interface"); value.Exists() {
		data.TrackInterfaces = make([]RouterVRRPInterfaceIPv6TrackInterfaces, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := RouterVRRPInterfaceIPv6TrackInterfaces{}
			if cValue := v.Get("interface-name"); cValue.Exists() {
				item.InterfaceName = types.StringValue(cValue.String())
			}
			if cValue := v.Get("priority-decrement"); cValue.Exists() {
				item.PriorityDecrement = types.Int64Value(cValue.Int())
			}
			data.TrackInterfaces = append(data.TrackInterfaces, item)
			return true
		})
	}
	if value := gjson.GetBytes(res, "track.objects.object"); value.Exists() {
		data.TrackObjects = make([]RouterVRRPInterfaceIPv6TrackObjects, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := RouterVRRPInterfaceIPv6TrackObjects{}
			if cValue := v.Get("object-name"); cValue.Exists() {
				item.ObjectName = types.StringValue(cValue.String())
			}
			if cValue := v.Get("priority-decrement"); cValue.Exists() {
				item.PriorityDecrement = types.Int64Value(cValue.Int())
			}
			data.TrackObjects = append(data.TrackObjects, item)
			return true
		})
	}
	if value := gjson.GetBytes(res, "bfd.fast-detect.peer.ipv6"); value.Exists() {
		data.BfdFastDetectPeerIpv6 = types.StringValue(value.String())
	}
}

func (data *RouterVRRPInterfaceIPv6) getDeletedItems(ctx context.Context, state RouterVRRPInterfaceIPv6) []string {
	deletedItems := make([]string, 0)
	if !state.GlobalAddresses.IsNull() && data.GlobalAddresses.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/address/global/global-address/address", state.getPath()))
	}
	if !state.AddressLinklocal.IsNull() && data.AddressLinklocal.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/address/linklocal/linklocal-address", state.getPath()))
	}
	if !state.AddressLinklocalAutoconfig.IsNull() && data.AddressLinklocalAutoconfig.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/address/linklocal/autoconfig", state.getPath()))
	}
	if !state.Priority.IsNull() && data.Priority.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/priority", state.getPath()))
	}
	if !state.Name.IsNull() && data.Name.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/name", state.getPath()))
	}
	if !state.TimerAdvertisementSeconds.IsNull() && data.TimerAdvertisementSeconds.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/timer/advertisement-time-in-seconds", state.getPath()))
	}
	if !state.TimerAdvertisementMilliseconds.IsNull() && data.TimerAdvertisementMilliseconds.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/timer/advertisement-time-in-milliseconds", state.getPath()))
	}
	if !state.TimerForce.IsNull() && data.TimerForce.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/timer/force", state.getPath()))
	}
	if !state.PreemptDisable.IsNull() && data.PreemptDisable.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/preempt/disable", state.getPath()))
	}
	if !state.PreemptDelay.IsNull() && data.PreemptDelay.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/preempt/delay", state.getPath()))
	}
	if !state.AcceptModeDisable.IsNull() && data.AcceptModeDisable.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/accept-mode/disable", state.getPath()))
	}
	for i := range state.TrackInterfaces {
		keys := [...]string{"interface-name"}
		stateKeyValues := [...]string{state.TrackInterfaces[i].InterfaceName.ValueString()}
		keyString := ""
		for ki := range keys {
			keyString += "[" + keys[ki] + "=" + stateKeyValues[ki] + "]"
		}

		emptyKeys := true
		if !reflect.ValueOf(state.TrackInterfaces[i].InterfaceName.ValueString()).IsZero() {
			emptyKeys = false
		}
		if emptyKeys {
			continue
		}

		found := false
		for j := range data.TrackInterfaces {
			found = true
			if state.TrackInterfaces[i].InterfaceName.ValueString() != data.TrackInterfaces[j].InterfaceName.ValueString() {
				found = false
			}
			if found {
				if !state.TrackInterfaces[i].PriorityDecrement.IsNull() && data.TrackInterfaces[j].PriorityDecrement.IsNull() {
					deletedItems = append(deletedItems, fmt.Sprintf("%v/track/interfaces/interface%v/priority-decrement", state.getPath(), keyString))
				}
				break
			}
		}
		if !found {
			deletedItems = append(deletedItems, fmt.Sprintf("%v/track/interfaces/interface%v", state.getPath(), keyString))
		}
	}
	for i := range state.TrackObjects {
		keys := [...]string{"object-name"}
		stateKeyValues := [...]string{state.TrackObjects[i].ObjectName.ValueString()}
		keyString := ""
		for ki := range keys {
			keyString += "[" + keys[ki] + "=" + stateKeyValues[ki] + "]"
		}

		emptyKeys := true
		if !reflect.ValueOf(state.TrackObjects[i].ObjectName.ValueString()).IsZero() {
			emptyKeys = false
		}
		if emptyKeys {
			continue
		}

		found := false
		for j := range data.TrackObjects {
			found = true
			if state.TrackObjects[i].ObjectName.ValueString() != data.TrackObjects[j].ObjectName.ValueString() {
				found = false
			}
			if found {
				if !state.TrackObjects[i].PriorityDecrement.IsNull() && data.TrackObjects[j].PriorityDecrement.IsNull() {
					deletedItems = append(deletedItems, fmt.Sprintf("%v/track/objects/object%v/priority-decrement", state.getPath(), keyString))
				}
				break
			}
		}
		if !found {
			deletedItems = append(deletedItems, fmt.Sprintf("%v/track/objects/object%v", state.getPath(), keyString))
		}
	}
	if !state.BfdFastDetectPeerIpv6.IsNull() && data.BfdFastDetectPeerIpv6.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/bfd/fast-detect/peer/ipv6", state.getPath()))
	}
	return deletedItems
}

func (data *RouterVRRPInterfaceIPv6) getEmptyLeafsDelete(ctx context.Context) []string {
	emptyLeafsDelete := make([]string, 0)
	if !data.AddressLinklocalAutoconfig.IsNull() && !data.AddressLinklocalAutoconfig.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/address/linklocal/autoconfig", data.getPath()))
	}
	if !data.TimerForce.IsNull() && !data.TimerForce.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/timer/force", data.getPath()))
	}
	if !data.PreemptDisable.IsNull() && !data.PreemptDisable.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/preempt/disable", data.getPath()))
	}
	if !data.AcceptModeDisable.IsNull() && !data.AcceptModeDisable.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/accept-mode/disable", data.getPath()))
	}
	for i := range data.TrackInterfaces {
		keys := [...]string{"interface-name"}
		keyValues := [...]string{data.TrackInterfaces[i].InterfaceName.ValueString()}
		keyString := ""
		for ki := range keys {
			keyString += "[" + keys[ki] + "=" + keyValues[ki] + "]"
		}
	}
	for i := range data.TrackObjects {
		keys := [...]string{"object-name"}
		keyValues := [...]string{data.TrackObjects[i].ObjectName.ValueString()}
		keyString := ""
		for ki := range keys {
			keyString += "[" + keys[ki] + "=" + keyValues[ki] + "]"
		}
	}
	return emptyLeafsDelete
}

func (data *RouterVRRPInterfaceIPv6) getDeletePaths(ctx context.Context) []string {
	var deletePaths []string
	if !data.GlobalAddresses.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/address/global/global-address/address", data.getPath()))
	}
	if !data.AddressLinklocal.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/address/linklocal/linklocal-address", data.getPath()))
	}
	if !data.AddressLinklocalAutoconfig.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/address/linklocal/autoconfig", data.getPath()))
	}
	if !data.Priority.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/priority", data.getPath()))
	}
	if !data.Name.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/name", data.getPath()))
	}
	if !data.TimerAdvertisementSeconds.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/timer/advertisement-time-in-seconds", data.getPath()))
	}
	if !data.TimerAdvertisementMilliseconds.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/timer/advertisement-time-in-milliseconds", data.getPath()))
	}
	if !data.TimerForce.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/timer/force", data.getPath()))
	}
	if !data.PreemptDisable.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/preempt/disable", data.getPath()))
	}
	if !data.PreemptDelay.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/preempt/delay", data.getPath()))
	}
	if !data.AcceptModeDisable.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/accept-mode/disable", data.getPath()))
	}
	for i := range data.TrackInterfaces {
		keys := [...]string{"interface-name"}
		keyValues := [...]string{data.TrackInterfaces[i].InterfaceName.ValueString()}

		keyString := ""
		for ki := range keys {
			keyString += "[" + keys[ki] + "=" + keyValues[ki] + "]"
		}
		deletePaths = append(deletePaths, fmt.Sprintf("%v/track/interfaces/interface%v", data.getPath(), keyString))
	}
	for i := range data.TrackObjects {
		keys := [...]string{"object-name"}
		keyValues := [...]string{data.TrackObjects[i].ObjectName.ValueString()}

		keyString := ""
		for ki := range keys {
			keyString += "[" + keys[ki] + "=" + keyValues[ki] + "]"
		}
		deletePaths = append(deletePaths, fmt.Sprintf("%v/track/objects/object%v", data.getPath(), keyString))
	}
	if !data.BfdFastDetectPeerIpv6.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/bfd/fast-detect/peer/ipv6", data.getPath()))
	}
	return deletePaths
}
