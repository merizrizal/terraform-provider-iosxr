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

type L2VPN struct {
	Device                     types.String          `tfsdk:"device"`
	Id                         types.String          `tfsdk:"id"`
	DeleteMode                 types.String          `tfsdk:"delete_mode"`
	Description                types.String          `tfsdk:"description"`
	RouterId                   types.String          `tfsdk:"router_id"`
	LoadBalancingFlowSrcDstMac types.Bool            `tfsdk:"load_balancing_flow_src_dst_mac"`
	LoadBalancingFlowSrcDstIp  types.Bool            `tfsdk:"load_balancing_flow_src_dst_ip"`
	XconnectGroups             []L2VPNXconnectGroups `tfsdk:"xconnect_groups"`
}

type L2VPNData struct {
	Device                     types.String          `tfsdk:"device"`
	Id                         types.String          `tfsdk:"id"`
	Description                types.String          `tfsdk:"description"`
	RouterId                   types.String          `tfsdk:"router_id"`
	LoadBalancingFlowSrcDstMac types.Bool            `tfsdk:"load_balancing_flow_src_dst_mac"`
	LoadBalancingFlowSrcDstIp  types.Bool            `tfsdk:"load_balancing_flow_src_dst_ip"`
	XconnectGroups             []L2VPNXconnectGroups `tfsdk:"xconnect_groups"`
}
type L2VPNXconnectGroups struct {
	GroupName types.String `tfsdk:"group_name"`
}

func (data L2VPN) getPath() string {
	return "Cisco-IOS-XR-um-l2vpn-cfg:/l2vpn"
}

func (data L2VPNData) getPath() string {
	return "Cisco-IOS-XR-um-l2vpn-cfg:/l2vpn"
}

func (data L2VPN) toBody(ctx context.Context) string {
	body := "{}"
	if !data.Description.IsNull() && !data.Description.IsUnknown() {
		body, _ = sjson.Set(body, "description", data.Description.ValueString())
	}
	if !data.RouterId.IsNull() && !data.RouterId.IsUnknown() {
		body, _ = sjson.Set(body, "router-id", data.RouterId.ValueString())
	}
	if !data.LoadBalancingFlowSrcDstMac.IsNull() && !data.LoadBalancingFlowSrcDstMac.IsUnknown() {
		if data.LoadBalancingFlowSrcDstMac.ValueBool() {
			body, _ = sjson.Set(body, "load-balancing.flow.src-dst-mac", map[string]string{})
		}
	}
	if !data.LoadBalancingFlowSrcDstIp.IsNull() && !data.LoadBalancingFlowSrcDstIp.IsUnknown() {
		if data.LoadBalancingFlowSrcDstIp.ValueBool() {
			body, _ = sjson.Set(body, "load-balancing.flow.src-dst-ip", map[string]string{})
		}
	}
	if len(data.XconnectGroups) > 0 {
		body, _ = sjson.Set(body, "xconnect.groups.group", []interface{}{})
		for index, item := range data.XconnectGroups {
			if !item.GroupName.IsNull() && !item.GroupName.IsUnknown() {
				body, _ = sjson.Set(body, "xconnect.groups.group"+"."+strconv.Itoa(index)+"."+"group-name", item.GroupName.ValueString())
			}
		}
	}
	return body
}

func (data *L2VPN) updateFromBody(ctx context.Context, res []byte) {
	if value := gjson.GetBytes(res, "description"); value.Exists() && !data.Description.IsNull() {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	if value := gjson.GetBytes(res, "router-id"); value.Exists() && !data.RouterId.IsNull() {
		data.RouterId = types.StringValue(value.String())
	} else {
		data.RouterId = types.StringNull()
	}
	if value := gjson.GetBytes(res, "load-balancing.flow.src-dst-mac"); !data.LoadBalancingFlowSrcDstMac.IsNull() {
		if value.Exists() {
			data.LoadBalancingFlowSrcDstMac = types.BoolValue(true)
		} else {
			data.LoadBalancingFlowSrcDstMac = types.BoolValue(false)
		}
	} else {
		data.LoadBalancingFlowSrcDstMac = types.BoolNull()
	}
	if value := gjson.GetBytes(res, "load-balancing.flow.src-dst-ip"); !data.LoadBalancingFlowSrcDstIp.IsNull() {
		if value.Exists() {
			data.LoadBalancingFlowSrcDstIp = types.BoolValue(true)
		} else {
			data.LoadBalancingFlowSrcDstIp = types.BoolValue(false)
		}
	} else {
		data.LoadBalancingFlowSrcDstIp = types.BoolNull()
	}
	for i := range data.XconnectGroups {
		keys := [...]string{"group-name"}
		keyValues := [...]string{data.XconnectGroups[i].GroupName.ValueString()}

		var r gjson.Result
		gjson.GetBytes(res, "xconnect.groups.group").ForEach(
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
		if value := r.Get("group-name"); value.Exists() && !data.XconnectGroups[i].GroupName.IsNull() {
			data.XconnectGroups[i].GroupName = types.StringValue(value.String())
		} else {
			data.XconnectGroups[i].GroupName = types.StringNull()
		}
	}
}

func (data *L2VPNData) fromBody(ctx context.Context, res []byte) {
	if value := gjson.GetBytes(res, "description"); value.Exists() {
		data.Description = types.StringValue(value.String())
	}
	if value := gjson.GetBytes(res, "router-id"); value.Exists() {
		data.RouterId = types.StringValue(value.String())
	}
	if value := gjson.GetBytes(res, "load-balancing.flow.src-dst-mac"); value.Exists() {
		data.LoadBalancingFlowSrcDstMac = types.BoolValue(true)
	} else {
		data.LoadBalancingFlowSrcDstMac = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "load-balancing.flow.src-dst-ip"); value.Exists() {
		data.LoadBalancingFlowSrcDstIp = types.BoolValue(true)
	} else {
		data.LoadBalancingFlowSrcDstIp = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "xconnect.groups.group"); value.Exists() {
		data.XconnectGroups = make([]L2VPNXconnectGroups, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := L2VPNXconnectGroups{}
			if cValue := v.Get("group-name"); cValue.Exists() {
				item.GroupName = types.StringValue(cValue.String())
			}
			data.XconnectGroups = append(data.XconnectGroups, item)
			return true
		})
	}
}

func (data *L2VPN) getDeletedItems(ctx context.Context, state L2VPN) []string {
	deletedItems := make([]string, 0)
	if !state.Description.IsNull() && data.Description.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/description", state.getPath()))
	}
	if !state.RouterId.IsNull() && data.RouterId.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/router-id", state.getPath()))
	}
	if !state.LoadBalancingFlowSrcDstMac.IsNull() && data.LoadBalancingFlowSrcDstMac.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/load-balancing/flow/src-dst-mac", state.getPath()))
	}
	if !state.LoadBalancingFlowSrcDstIp.IsNull() && data.LoadBalancingFlowSrcDstIp.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/load-balancing/flow/src-dst-ip", state.getPath()))
	}
	for i := range state.XconnectGroups {
		keys := [...]string{"group-name"}
		stateKeyValues := [...]string{state.XconnectGroups[i].GroupName.ValueString()}
		keyString := ""
		for ki := range keys {
			keyString += "[" + keys[ki] + "=" + stateKeyValues[ki] + "]"
		}

		emptyKeys := true
		if !reflect.ValueOf(state.XconnectGroups[i].GroupName.ValueString()).IsZero() {
			emptyKeys = false
		}
		if emptyKeys {
			continue
		}

		found := false
		for j := range data.XconnectGroups {
			found = true
			if state.XconnectGroups[i].GroupName.ValueString() != data.XconnectGroups[j].GroupName.ValueString() {
				found = false
			}
			if found {
				break
			}
		}
		if !found {
			deletedItems = append(deletedItems, fmt.Sprintf("%v/xconnect/groups/group%v", state.getPath(), keyString))
		}
	}
	return deletedItems
}

func (data *L2VPN) getEmptyLeafsDelete(ctx context.Context) []string {
	emptyLeafsDelete := make([]string, 0)
	if !data.LoadBalancingFlowSrcDstMac.IsNull() && !data.LoadBalancingFlowSrcDstMac.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/load-balancing/flow/src-dst-mac", data.getPath()))
	}
	if !data.LoadBalancingFlowSrcDstIp.IsNull() && !data.LoadBalancingFlowSrcDstIp.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/load-balancing/flow/src-dst-ip", data.getPath()))
	}
	for i := range data.XconnectGroups {
		keys := [...]string{"group-name"}
		keyValues := [...]string{data.XconnectGroups[i].GroupName.ValueString()}
		keyString := ""
		for ki := range keys {
			keyString += "[" + keys[ki] + "=" + keyValues[ki] + "]"
		}
	}
	return emptyLeafsDelete
}

func (data *L2VPN) getDeletePaths(ctx context.Context) []string {
	var deletePaths []string
	if !data.Description.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/description", data.getPath()))
	}
	if !data.RouterId.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/router-id", data.getPath()))
	}
	if !data.LoadBalancingFlowSrcDstMac.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/load-balancing/flow/src-dst-mac", data.getPath()))
	}
	if !data.LoadBalancingFlowSrcDstIp.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/load-balancing/flow/src-dst-ip", data.getPath()))
	}
	for i := range data.XconnectGroups {
		keys := [...]string{"group-name"}
		keyValues := [...]string{data.XconnectGroups[i].GroupName.ValueString()}

		keyString := ""
		for ki := range keys {
			keyString += "[" + keys[ki] + "=" + keyValues[ki] + "]"
		}
		deletePaths = append(deletePaths, fmt.Sprintf("%v/xconnect/groups/group%v", data.getPath(), keyString))
	}
	return deletePaths
}
