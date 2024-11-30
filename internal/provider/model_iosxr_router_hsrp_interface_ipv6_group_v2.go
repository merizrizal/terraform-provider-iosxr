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

type RouterHSRPInterfaceIPv6GroupV2 struct {
	Device                                     types.String                                    `tfsdk:"device"`
	Id                                         types.String                                    `tfsdk:"id"`
	DeleteMode                                 types.String                                    `tfsdk:"delete_mode"`
	InterfaceName                              types.String                                    `tfsdk:"interface_name"`
	GroupId                                    types.Int64                                     `tfsdk:"group_id"`
	Name                                       types.String                                    `tfsdk:"name"`
	MacAddress                                 types.String                                    `tfsdk:"mac_address"`
	TimersHoldTime                             types.Int64                                     `tfsdk:"timers_hold_time"`
	TimersHoldTime2                            types.Int64                                     `tfsdk:"timers_hold_time2"`
	TimersMsec                                 types.Int64                                     `tfsdk:"timers_msec"`
	TimersMsec2                                types.Int64                                     `tfsdk:"timers_msec2"`
	PreemptDelay                               types.Int64                                     `tfsdk:"preempt_delay"`
	Priority                                   types.Int64                                     `tfsdk:"priority"`
	BfdFastDetectPeerIpv6                      types.String                                    `tfsdk:"bfd_fast_detect_peer_ipv6"`
	BfdFastDetectPeerInterface                 types.String                                    `tfsdk:"bfd_fast_detect_peer_interface"`
	TrackObjects                               []RouterHSRPInterfaceIPv6GroupV2TrackObjects    `tfsdk:"track_objects"`
	TrackInterfaces                            []RouterHSRPInterfaceIPv6GroupV2TrackInterfaces `tfsdk:"track_interfaces"`
	Addresses                                  []RouterHSRPInterfaceIPv6GroupV2Addresses       `tfsdk:"addresses"`
	AddressLinkLocalAutoconfig                 types.Bool                                      `tfsdk:"address_link_local_autoconfig"`
	AddressLinkLocalAutoconfigLegacyCompatible types.Bool                                      `tfsdk:"address_link_local_autoconfig_legacy_compatible"`
	AddressLinkLocalIpv6Address                types.String                                    `tfsdk:"address_link_local_ipv6_address"`
}

type RouterHSRPInterfaceIPv6GroupV2Data struct {
	Device                                     types.String                                    `tfsdk:"device"`
	Id                                         types.String                                    `tfsdk:"id"`
	InterfaceName                              types.String                                    `tfsdk:"interface_name"`
	GroupId                                    types.Int64                                     `tfsdk:"group_id"`
	Name                                       types.String                                    `tfsdk:"name"`
	MacAddress                                 types.String                                    `tfsdk:"mac_address"`
	TimersHoldTime                             types.Int64                                     `tfsdk:"timers_hold_time"`
	TimersHoldTime2                            types.Int64                                     `tfsdk:"timers_hold_time2"`
	TimersMsec                                 types.Int64                                     `tfsdk:"timers_msec"`
	TimersMsec2                                types.Int64                                     `tfsdk:"timers_msec2"`
	PreemptDelay                               types.Int64                                     `tfsdk:"preempt_delay"`
	Priority                                   types.Int64                                     `tfsdk:"priority"`
	BfdFastDetectPeerIpv6                      types.String                                    `tfsdk:"bfd_fast_detect_peer_ipv6"`
	BfdFastDetectPeerInterface                 types.String                                    `tfsdk:"bfd_fast_detect_peer_interface"`
	TrackObjects                               []RouterHSRPInterfaceIPv6GroupV2TrackObjects    `tfsdk:"track_objects"`
	TrackInterfaces                            []RouterHSRPInterfaceIPv6GroupV2TrackInterfaces `tfsdk:"track_interfaces"`
	Addresses                                  []RouterHSRPInterfaceIPv6GroupV2Addresses       `tfsdk:"addresses"`
	AddressLinkLocalAutoconfig                 types.Bool                                      `tfsdk:"address_link_local_autoconfig"`
	AddressLinkLocalAutoconfigLegacyCompatible types.Bool                                      `tfsdk:"address_link_local_autoconfig_legacy_compatible"`
	AddressLinkLocalIpv6Address                types.String                                    `tfsdk:"address_link_local_ipv6_address"`
}
type RouterHSRPInterfaceIPv6GroupV2TrackObjects struct {
	ObjectName        types.String `tfsdk:"object_name"`
	PriorityDecrement types.Int64  `tfsdk:"priority_decrement"`
}
type RouterHSRPInterfaceIPv6GroupV2TrackInterfaces struct {
	TrackName         types.String `tfsdk:"track_name"`
	PriorityDecrement types.Int64  `tfsdk:"priority_decrement"`
}
type RouterHSRPInterfaceIPv6GroupV2Addresses struct {
	Address types.String `tfsdk:"address"`
}

func (data RouterHSRPInterfaceIPv6GroupV2) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XR-um-router-hsrp-cfg:/router/hsrp/interfaces/interface[interface-name=%s]/address-family/ipv6/hsrp/group-number-version-2s/group-number-version-2[group-number-version-2-id=%v]", data.InterfaceName.ValueString(), data.GroupId.ValueInt64())
}

func (data RouterHSRPInterfaceIPv6GroupV2Data) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XR-um-router-hsrp-cfg:/router/hsrp/interfaces/interface[interface-name=%s]/address-family/ipv6/hsrp/group-number-version-2s/group-number-version-2[group-number-version-2-id=%v]", data.InterfaceName.ValueString(), data.GroupId.ValueInt64())
}

func (data RouterHSRPInterfaceIPv6GroupV2) toBody(ctx context.Context) string {
	body := "{}"
	if !data.GroupId.IsNull() && !data.GroupId.IsUnknown() {
		body, _ = sjson.Set(body, "group-number-version-2-id", strconv.FormatInt(data.GroupId.ValueInt64(), 10))
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	if !data.MacAddress.IsNull() && !data.MacAddress.IsUnknown() {
		body, _ = sjson.Set(body, "mac-address", data.MacAddress.ValueString())
	}
	if !data.TimersHoldTime.IsNull() && !data.TimersHoldTime.IsUnknown() {
		body, _ = sjson.Set(body, "timers.hold-time", strconv.FormatInt(data.TimersHoldTime.ValueInt64(), 10))
	}
	if !data.TimersHoldTime2.IsNull() && !data.TimersHoldTime2.IsUnknown() {
		body, _ = sjson.Set(body, "timers.hold-time2", strconv.FormatInt(data.TimersHoldTime2.ValueInt64(), 10))
	}
	if !data.TimersMsec.IsNull() && !data.TimersMsec.IsUnknown() {
		body, _ = sjson.Set(body, "timers.msec", strconv.FormatInt(data.TimersMsec.ValueInt64(), 10))
	}
	if !data.TimersMsec2.IsNull() && !data.TimersMsec2.IsUnknown() {
		body, _ = sjson.Set(body, "timers.msec2", strconv.FormatInt(data.TimersMsec2.ValueInt64(), 10))
	}
	if !data.PreemptDelay.IsNull() && !data.PreemptDelay.IsUnknown() {
		body, _ = sjson.Set(body, "preempt.delay", strconv.FormatInt(data.PreemptDelay.ValueInt64(), 10))
	}
	if !data.Priority.IsNull() && !data.Priority.IsUnknown() {
		body, _ = sjson.Set(body, "priority", strconv.FormatInt(data.Priority.ValueInt64(), 10))
	}
	if !data.BfdFastDetectPeerIpv6.IsNull() && !data.BfdFastDetectPeerIpv6.IsUnknown() {
		body, _ = sjson.Set(body, "bfd.fast-detect.peer.ipv6", data.BfdFastDetectPeerIpv6.ValueString())
	}
	if !data.BfdFastDetectPeerInterface.IsNull() && !data.BfdFastDetectPeerInterface.IsUnknown() {
		body, _ = sjson.Set(body, "bfd.fast-detect.peer.interface", data.BfdFastDetectPeerInterface.ValueString())
	}
	if !data.AddressLinkLocalAutoconfig.IsNull() && !data.AddressLinkLocalAutoconfig.IsUnknown() {
		if data.AddressLinkLocalAutoconfig.ValueBool() {
			body, _ = sjson.Set(body, "address.link-local.autoconfig", map[string]string{})
		}
	}
	if !data.AddressLinkLocalAutoconfigLegacyCompatible.IsNull() && !data.AddressLinkLocalAutoconfigLegacyCompatible.IsUnknown() {
		if data.AddressLinkLocalAutoconfigLegacyCompatible.ValueBool() {
			body, _ = sjson.Set(body, "address.link-local.autoconfig.legacy-compatible", map[string]string{})
		}
	}
	if !data.AddressLinkLocalIpv6Address.IsNull() && !data.AddressLinkLocalIpv6Address.IsUnknown() {
		body, _ = sjson.Set(body, "address.link-local.ipv6-address", data.AddressLinkLocalIpv6Address.ValueString())
	}
	if len(data.TrackObjects) > 0 {
		body, _ = sjson.Set(body, "track-objects.track-object", []interface{}{})
		for index, item := range data.TrackObjects {
			if !item.ObjectName.IsNull() && !item.ObjectName.IsUnknown() {
				body, _ = sjson.Set(body, "track-objects.track-object"+"."+strconv.Itoa(index)+"."+"object-name", item.ObjectName.ValueString())
			}
			if !item.PriorityDecrement.IsNull() && !item.PriorityDecrement.IsUnknown() {
				body, _ = sjson.Set(body, "track-objects.track-object"+"."+strconv.Itoa(index)+"."+"priority-decrement", strconv.FormatInt(item.PriorityDecrement.ValueInt64(), 10))
			}
		}
	}
	if len(data.TrackInterfaces) > 0 {
		body, _ = sjson.Set(body, "track-interfaces.track-interface", []interface{}{})
		for index, item := range data.TrackInterfaces {
			if !item.TrackName.IsNull() && !item.TrackName.IsUnknown() {
				body, _ = sjson.Set(body, "track-interfaces.track-interface"+"."+strconv.Itoa(index)+"."+"track-name", item.TrackName.ValueString())
			}
			if !item.PriorityDecrement.IsNull() && !item.PriorityDecrement.IsUnknown() {
				body, _ = sjson.Set(body, "track-interfaces.track-interface"+"."+strconv.Itoa(index)+"."+"priority-decrement", strconv.FormatInt(item.PriorityDecrement.ValueInt64(), 10))
			}
		}
	}
	if len(data.Addresses) > 0 {
		body, _ = sjson.Set(body, "address.globals.global", []interface{}{})
		for index, item := range data.Addresses {
			if !item.Address.IsNull() && !item.Address.IsUnknown() {
				body, _ = sjson.Set(body, "address.globals.global"+"."+strconv.Itoa(index)+"."+"address", item.Address.ValueString())
			}
		}
	}
	return body
}

func (data *RouterHSRPInterfaceIPv6GroupV2) updateFromBody(ctx context.Context, res []byte) {
	if value := gjson.GetBytes(res, "name"); value.Exists() && !data.Name.IsNull() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := gjson.GetBytes(res, "mac-address"); value.Exists() && !data.MacAddress.IsNull() {
		data.MacAddress = types.StringValue(value.String())
	} else {
		data.MacAddress = types.StringNull()
	}
	if value := gjson.GetBytes(res, "timers.hold-time"); value.Exists() && !data.TimersHoldTime.IsNull() {
		data.TimersHoldTime = types.Int64Value(value.Int())
	} else {
		data.TimersHoldTime = types.Int64Null()
	}
	if value := gjson.GetBytes(res, "timers.hold-time2"); value.Exists() && !data.TimersHoldTime2.IsNull() {
		data.TimersHoldTime2 = types.Int64Value(value.Int())
	} else {
		data.TimersHoldTime2 = types.Int64Null()
	}
	if value := gjson.GetBytes(res, "timers.msec"); value.Exists() && !data.TimersMsec.IsNull() {
		data.TimersMsec = types.Int64Value(value.Int())
	} else {
		data.TimersMsec = types.Int64Null()
	}
	if value := gjson.GetBytes(res, "timers.msec2"); value.Exists() && !data.TimersMsec2.IsNull() {
		data.TimersMsec2 = types.Int64Value(value.Int())
	} else {
		data.TimersMsec2 = types.Int64Null()
	}
	if value := gjson.GetBytes(res, "preempt.delay"); value.Exists() && !data.PreemptDelay.IsNull() {
		data.PreemptDelay = types.Int64Value(value.Int())
	} else {
		data.PreemptDelay = types.Int64Null()
	}
	if value := gjson.GetBytes(res, "priority"); value.Exists() && !data.Priority.IsNull() {
		data.Priority = types.Int64Value(value.Int())
	} else {
		data.Priority = types.Int64Null()
	}
	if value := gjson.GetBytes(res, "bfd.fast-detect.peer.ipv6"); value.Exists() && !data.BfdFastDetectPeerIpv6.IsNull() {
		data.BfdFastDetectPeerIpv6 = types.StringValue(value.String())
	} else {
		data.BfdFastDetectPeerIpv6 = types.StringNull()
	}
	if value := gjson.GetBytes(res, "bfd.fast-detect.peer.interface"); value.Exists() && !data.BfdFastDetectPeerInterface.IsNull() {
		data.BfdFastDetectPeerInterface = types.StringValue(value.String())
	} else {
		data.BfdFastDetectPeerInterface = types.StringNull()
	}
	for i := range data.TrackObjects {
		keys := [...]string{"object-name"}
		keyValues := [...]string{data.TrackObjects[i].ObjectName.ValueString()}

		var r gjson.Result
		gjson.GetBytes(res, "track-objects.track-object").ForEach(
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
	for i := range data.TrackInterfaces {
		keys := [...]string{"track-name"}
		keyValues := [...]string{data.TrackInterfaces[i].TrackName.ValueString()}

		var r gjson.Result
		gjson.GetBytes(res, "track-interfaces.track-interface").ForEach(
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
		if value := r.Get("track-name"); value.Exists() && !data.TrackInterfaces[i].TrackName.IsNull() {
			data.TrackInterfaces[i].TrackName = types.StringValue(value.String())
		} else {
			data.TrackInterfaces[i].TrackName = types.StringNull()
		}
		if value := r.Get("priority-decrement"); value.Exists() && !data.TrackInterfaces[i].PriorityDecrement.IsNull() {
			data.TrackInterfaces[i].PriorityDecrement = types.Int64Value(value.Int())
		} else {
			data.TrackInterfaces[i].PriorityDecrement = types.Int64Null()
		}
	}
	for i := range data.Addresses {
		keys := [...]string{"address"}
		keyValues := [...]string{data.Addresses[i].Address.ValueString()}

		var r gjson.Result
		gjson.GetBytes(res, "address.globals.global").ForEach(
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
		if value := r.Get("address"); value.Exists() && !data.Addresses[i].Address.IsNull() {
			data.Addresses[i].Address = types.StringValue(value.String())
		} else {
			data.Addresses[i].Address = types.StringNull()
		}
	}
	if value := gjson.GetBytes(res, "address.link-local.autoconfig"); !data.AddressLinkLocalAutoconfig.IsNull() {
		if value.Exists() {
			data.AddressLinkLocalAutoconfig = types.BoolValue(true)
		} else {
			data.AddressLinkLocalAutoconfig = types.BoolValue(false)
		}
	} else {
		data.AddressLinkLocalAutoconfig = types.BoolNull()
	}
	if value := gjson.GetBytes(res, "address.link-local.autoconfig.legacy-compatible"); !data.AddressLinkLocalAutoconfigLegacyCompatible.IsNull() {
		if value.Exists() {
			data.AddressLinkLocalAutoconfigLegacyCompatible = types.BoolValue(true)
		} else {
			data.AddressLinkLocalAutoconfigLegacyCompatible = types.BoolValue(false)
		}
	} else {
		data.AddressLinkLocalAutoconfigLegacyCompatible = types.BoolNull()
	}
	if value := gjson.GetBytes(res, "address.link-local.ipv6-address"); value.Exists() && !data.AddressLinkLocalIpv6Address.IsNull() {
		data.AddressLinkLocalIpv6Address = types.StringValue(value.String())
	} else {
		data.AddressLinkLocalIpv6Address = types.StringNull()
	}
}

func (data *RouterHSRPInterfaceIPv6GroupV2) fromBody(ctx context.Context, res []byte) {
	if value := gjson.GetBytes(res, "name"); value.Exists() {
		data.Name = types.StringValue(value.String())
	}
	if value := gjson.GetBytes(res, "mac-address"); value.Exists() {
		data.MacAddress = types.StringValue(value.String())
	}
	if value := gjson.GetBytes(res, "timers.hold-time"); value.Exists() {
		data.TimersHoldTime = types.Int64Value(value.Int())
	}
	if value := gjson.GetBytes(res, "timers.hold-time2"); value.Exists() {
		data.TimersHoldTime2 = types.Int64Value(value.Int())
	}
	if value := gjson.GetBytes(res, "timers.msec"); value.Exists() {
		data.TimersMsec = types.Int64Value(value.Int())
	}
	if value := gjson.GetBytes(res, "timers.msec2"); value.Exists() {
		data.TimersMsec2 = types.Int64Value(value.Int())
	}
	if value := gjson.GetBytes(res, "preempt.delay"); value.Exists() {
		data.PreemptDelay = types.Int64Value(value.Int())
	}
	if value := gjson.GetBytes(res, "priority"); value.Exists() {
		data.Priority = types.Int64Value(value.Int())
	}
	if value := gjson.GetBytes(res, "bfd.fast-detect.peer.ipv6"); value.Exists() {
		data.BfdFastDetectPeerIpv6 = types.StringValue(value.String())
	}
	if value := gjson.GetBytes(res, "bfd.fast-detect.peer.interface"); value.Exists() {
		data.BfdFastDetectPeerInterface = types.StringValue(value.String())
	}
	if value := gjson.GetBytes(res, "track-objects.track-object"); value.Exists() {
		data.TrackObjects = make([]RouterHSRPInterfaceIPv6GroupV2TrackObjects, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := RouterHSRPInterfaceIPv6GroupV2TrackObjects{}
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
	if value := gjson.GetBytes(res, "track-interfaces.track-interface"); value.Exists() {
		data.TrackInterfaces = make([]RouterHSRPInterfaceIPv6GroupV2TrackInterfaces, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := RouterHSRPInterfaceIPv6GroupV2TrackInterfaces{}
			if cValue := v.Get("track-name"); cValue.Exists() {
				item.TrackName = types.StringValue(cValue.String())
			}
			if cValue := v.Get("priority-decrement"); cValue.Exists() {
				item.PriorityDecrement = types.Int64Value(cValue.Int())
			}
			data.TrackInterfaces = append(data.TrackInterfaces, item)
			return true
		})
	}
	if value := gjson.GetBytes(res, "address.globals.global"); value.Exists() {
		data.Addresses = make([]RouterHSRPInterfaceIPv6GroupV2Addresses, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := RouterHSRPInterfaceIPv6GroupV2Addresses{}
			if cValue := v.Get("address"); cValue.Exists() {
				item.Address = types.StringValue(cValue.String())
			}
			data.Addresses = append(data.Addresses, item)
			return true
		})
	}
	if value := gjson.GetBytes(res, "address.link-local.autoconfig"); value.Exists() {
		data.AddressLinkLocalAutoconfig = types.BoolValue(true)
	} else {
		data.AddressLinkLocalAutoconfig = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "address.link-local.autoconfig.legacy-compatible"); value.Exists() {
		data.AddressLinkLocalAutoconfigLegacyCompatible = types.BoolValue(true)
	} else {
		data.AddressLinkLocalAutoconfigLegacyCompatible = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "address.link-local.ipv6-address"); value.Exists() {
		data.AddressLinkLocalIpv6Address = types.StringValue(value.String())
	}
}

func (data *RouterHSRPInterfaceIPv6GroupV2Data) fromBody(ctx context.Context, res []byte) {
	if value := gjson.GetBytes(res, "name"); value.Exists() {
		data.Name = types.StringValue(value.String())
	}
	if value := gjson.GetBytes(res, "mac-address"); value.Exists() {
		data.MacAddress = types.StringValue(value.String())
	}
	if value := gjson.GetBytes(res, "timers.hold-time"); value.Exists() {
		data.TimersHoldTime = types.Int64Value(value.Int())
	}
	if value := gjson.GetBytes(res, "timers.hold-time2"); value.Exists() {
		data.TimersHoldTime2 = types.Int64Value(value.Int())
	}
	if value := gjson.GetBytes(res, "timers.msec"); value.Exists() {
		data.TimersMsec = types.Int64Value(value.Int())
	}
	if value := gjson.GetBytes(res, "timers.msec2"); value.Exists() {
		data.TimersMsec2 = types.Int64Value(value.Int())
	}
	if value := gjson.GetBytes(res, "preempt.delay"); value.Exists() {
		data.PreemptDelay = types.Int64Value(value.Int())
	}
	if value := gjson.GetBytes(res, "priority"); value.Exists() {
		data.Priority = types.Int64Value(value.Int())
	}
	if value := gjson.GetBytes(res, "bfd.fast-detect.peer.ipv6"); value.Exists() {
		data.BfdFastDetectPeerIpv6 = types.StringValue(value.String())
	}
	if value := gjson.GetBytes(res, "bfd.fast-detect.peer.interface"); value.Exists() {
		data.BfdFastDetectPeerInterface = types.StringValue(value.String())
	}
	if value := gjson.GetBytes(res, "track-objects.track-object"); value.Exists() {
		data.TrackObjects = make([]RouterHSRPInterfaceIPv6GroupV2TrackObjects, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := RouterHSRPInterfaceIPv6GroupV2TrackObjects{}
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
	if value := gjson.GetBytes(res, "track-interfaces.track-interface"); value.Exists() {
		data.TrackInterfaces = make([]RouterHSRPInterfaceIPv6GroupV2TrackInterfaces, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := RouterHSRPInterfaceIPv6GroupV2TrackInterfaces{}
			if cValue := v.Get("track-name"); cValue.Exists() {
				item.TrackName = types.StringValue(cValue.String())
			}
			if cValue := v.Get("priority-decrement"); cValue.Exists() {
				item.PriorityDecrement = types.Int64Value(cValue.Int())
			}
			data.TrackInterfaces = append(data.TrackInterfaces, item)
			return true
		})
	}
	if value := gjson.GetBytes(res, "address.globals.global"); value.Exists() {
		data.Addresses = make([]RouterHSRPInterfaceIPv6GroupV2Addresses, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := RouterHSRPInterfaceIPv6GroupV2Addresses{}
			if cValue := v.Get("address"); cValue.Exists() {
				item.Address = types.StringValue(cValue.String())
			}
			data.Addresses = append(data.Addresses, item)
			return true
		})
	}
	if value := gjson.GetBytes(res, "address.link-local.autoconfig"); value.Exists() {
		data.AddressLinkLocalAutoconfig = types.BoolValue(true)
	} else {
		data.AddressLinkLocalAutoconfig = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "address.link-local.autoconfig.legacy-compatible"); value.Exists() {
		data.AddressLinkLocalAutoconfigLegacyCompatible = types.BoolValue(true)
	} else {
		data.AddressLinkLocalAutoconfigLegacyCompatible = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "address.link-local.ipv6-address"); value.Exists() {
		data.AddressLinkLocalIpv6Address = types.StringValue(value.String())
	}
}

func (data *RouterHSRPInterfaceIPv6GroupV2) getDeletedItems(ctx context.Context, state RouterHSRPInterfaceIPv6GroupV2) []string {
	deletedItems := make([]string, 0)
	if !state.Name.IsNull() && data.Name.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/name", state.getPath()))
	}
	if !state.MacAddress.IsNull() && data.MacAddress.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/mac-address", state.getPath()))
	}
	if !state.TimersHoldTime.IsNull() && data.TimersHoldTime.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/timers/hold-time", state.getPath()))
	}
	if !state.TimersHoldTime2.IsNull() && data.TimersHoldTime2.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/timers/hold-time2", state.getPath()))
	}
	if !state.TimersMsec.IsNull() && data.TimersMsec.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/timers/msec", state.getPath()))
	}
	if !state.TimersMsec2.IsNull() && data.TimersMsec2.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/timers/msec2", state.getPath()))
	}
	if !state.PreemptDelay.IsNull() && data.PreemptDelay.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/preempt/delay", state.getPath()))
	}
	if !state.Priority.IsNull() && data.Priority.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/priority", state.getPath()))
	}
	if !state.BfdFastDetectPeerIpv6.IsNull() && data.BfdFastDetectPeerIpv6.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/bfd/fast-detect/peer/ipv6", state.getPath()))
	}
	if !state.BfdFastDetectPeerInterface.IsNull() && data.BfdFastDetectPeerInterface.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/bfd/fast-detect/peer/interface", state.getPath()))
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
					deletedItems = append(deletedItems, fmt.Sprintf("%v/track-objects/track-object%v/priority-decrement", state.getPath(), keyString))
				}
				break
			}
		}
		if !found {
			deletedItems = append(deletedItems, fmt.Sprintf("%v/track-objects/track-object%v", state.getPath(), keyString))
		}
	}
	for i := range state.TrackInterfaces {
		keys := [...]string{"track-name"}
		stateKeyValues := [...]string{state.TrackInterfaces[i].TrackName.ValueString()}
		keyString := ""
		for ki := range keys {
			keyString += "[" + keys[ki] + "=" + stateKeyValues[ki] + "]"
		}

		emptyKeys := true
		if !reflect.ValueOf(state.TrackInterfaces[i].TrackName.ValueString()).IsZero() {
			emptyKeys = false
		}
		if emptyKeys {
			continue
		}

		found := false
		for j := range data.TrackInterfaces {
			found = true
			if state.TrackInterfaces[i].TrackName.ValueString() != data.TrackInterfaces[j].TrackName.ValueString() {
				found = false
			}
			if found {
				if !state.TrackInterfaces[i].PriorityDecrement.IsNull() && data.TrackInterfaces[j].PriorityDecrement.IsNull() {
					deletedItems = append(deletedItems, fmt.Sprintf("%v/track-interfaces/track-interface%v/priority-decrement", state.getPath(), keyString))
				}
				break
			}
		}
		if !found {
			deletedItems = append(deletedItems, fmt.Sprintf("%v/track-interfaces/track-interface%v", state.getPath(), keyString))
		}
	}
	for i := range state.Addresses {
		keys := [...]string{"address"}
		stateKeyValues := [...]string{state.Addresses[i].Address.ValueString()}
		keyString := ""
		for ki := range keys {
			keyString += "[" + keys[ki] + "=" + stateKeyValues[ki] + "]"
		}

		emptyKeys := true
		if !reflect.ValueOf(state.Addresses[i].Address.ValueString()).IsZero() {
			emptyKeys = false
		}
		if emptyKeys {
			continue
		}

		found := false
		for j := range data.Addresses {
			found = true
			if state.Addresses[i].Address.ValueString() != data.Addresses[j].Address.ValueString() {
				found = false
			}
			if found {
				break
			}
		}
		if !found {
			deletedItems = append(deletedItems, fmt.Sprintf("%v/address/globals/global%v", state.getPath(), keyString))
		}
	}
	if !state.AddressLinkLocalAutoconfig.IsNull() && data.AddressLinkLocalAutoconfig.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/address/link-local/autoconfig", state.getPath()))
	}
	if !state.AddressLinkLocalAutoconfigLegacyCompatible.IsNull() && data.AddressLinkLocalAutoconfigLegacyCompatible.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/address/link-local/autoconfig/legacy-compatible", state.getPath()))
	}
	if !state.AddressLinkLocalIpv6Address.IsNull() && data.AddressLinkLocalIpv6Address.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/address/link-local/ipv6-address", state.getPath()))
	}
	return deletedItems
}

func (data *RouterHSRPInterfaceIPv6GroupV2) getEmptyLeafsDelete(ctx context.Context) []string {
	emptyLeafsDelete := make([]string, 0)
	for i := range data.TrackObjects {
		keys := [...]string{"object-name"}
		keyValues := [...]string{data.TrackObjects[i].ObjectName.ValueString()}
		keyString := ""
		for ki := range keys {
			keyString += "[" + keys[ki] + "=" + keyValues[ki] + "]"
		}
	}
	for i := range data.TrackInterfaces {
		keys := [...]string{"track-name"}
		keyValues := [...]string{data.TrackInterfaces[i].TrackName.ValueString()}
		keyString := ""
		for ki := range keys {
			keyString += "[" + keys[ki] + "=" + keyValues[ki] + "]"
		}
	}
	for i := range data.Addresses {
		keys := [...]string{"address"}
		keyValues := [...]string{data.Addresses[i].Address.ValueString()}
		keyString := ""
		for ki := range keys {
			keyString += "[" + keys[ki] + "=" + keyValues[ki] + "]"
		}
	}
	if !data.AddressLinkLocalAutoconfig.IsNull() && !data.AddressLinkLocalAutoconfig.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/address/link-local/autoconfig", data.getPath()))
	}
	if !data.AddressLinkLocalAutoconfigLegacyCompatible.IsNull() && !data.AddressLinkLocalAutoconfigLegacyCompatible.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/address/link-local/autoconfig/legacy-compatible", data.getPath()))
	}
	return emptyLeafsDelete
}

func (data *RouterHSRPInterfaceIPv6GroupV2) getDeletePaths(ctx context.Context) []string {
	var deletePaths []string
	if !data.Name.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/name", data.getPath()))
	}
	if !data.MacAddress.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/mac-address", data.getPath()))
	}
	if !data.TimersHoldTime.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/timers/hold-time", data.getPath()))
	}
	if !data.TimersHoldTime2.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/timers/hold-time2", data.getPath()))
	}
	if !data.TimersMsec.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/timers/msec", data.getPath()))
	}
	if !data.TimersMsec2.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/timers/msec2", data.getPath()))
	}
	if !data.PreemptDelay.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/preempt/delay", data.getPath()))
	}
	if !data.Priority.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/priority", data.getPath()))
	}
	if !data.BfdFastDetectPeerIpv6.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/bfd/fast-detect/peer/ipv6", data.getPath()))
	}
	if !data.BfdFastDetectPeerInterface.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/bfd/fast-detect/peer/interface", data.getPath()))
	}
	for i := range data.TrackObjects {
		keys := [...]string{"object-name"}
		keyValues := [...]string{data.TrackObjects[i].ObjectName.ValueString()}

		keyString := ""
		for ki := range keys {
			keyString += "[" + keys[ki] + "=" + keyValues[ki] + "]"
		}
		deletePaths = append(deletePaths, fmt.Sprintf("%v/track-objects/track-object%v", data.getPath(), keyString))
	}
	for i := range data.TrackInterfaces {
		keys := [...]string{"track-name"}
		keyValues := [...]string{data.TrackInterfaces[i].TrackName.ValueString()}

		keyString := ""
		for ki := range keys {
			keyString += "[" + keys[ki] + "=" + keyValues[ki] + "]"
		}
		deletePaths = append(deletePaths, fmt.Sprintf("%v/track-interfaces/track-interface%v", data.getPath(), keyString))
	}
	for i := range data.Addresses {
		keys := [...]string{"address"}
		keyValues := [...]string{data.Addresses[i].Address.ValueString()}

		keyString := ""
		for ki := range keys {
			keyString += "[" + keys[ki] + "=" + keyValues[ki] + "]"
		}
		deletePaths = append(deletePaths, fmt.Sprintf("%v/address/globals/global%v", data.getPath(), keyString))
	}
	if !data.AddressLinkLocalAutoconfig.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/address/link-local/autoconfig", data.getPath()))
	}
	if !data.AddressLinkLocalAutoconfigLegacyCompatible.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/address/link-local/autoconfig/legacy-compatible", data.getPath()))
	}
	if !data.AddressLinkLocalIpv6Address.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/address/link-local/ipv6-address", data.getPath()))
	}
	return deletePaths
}
