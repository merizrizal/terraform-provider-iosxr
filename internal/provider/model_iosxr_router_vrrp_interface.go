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
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type RouterVRRPInterface struct {
	Device             types.String `tfsdk:"device"`
	Id                 types.String `tfsdk:"id"`
	DeleteMode         types.String `tfsdk:"delete_mode"`
	InterfaceName      types.String `tfsdk:"interface_name"`
	MacRefresh         types.Int64  `tfsdk:"mac_refresh"`
	DelayMinimum       types.Int64  `tfsdk:"delay_minimum"`
	DelayReload        types.Int64  `tfsdk:"delay_reload"`
	BfdMinimumInterval types.Int64  `tfsdk:"bfd_minimum_interval"`
	BfdMultiplier      types.Int64  `tfsdk:"bfd_multiplier"`
}

type RouterVRRPInterfaceData struct {
	Device             types.String `tfsdk:"device"`
	Id                 types.String `tfsdk:"id"`
	InterfaceName      types.String `tfsdk:"interface_name"`
	MacRefresh         types.Int64  `tfsdk:"mac_refresh"`
	DelayMinimum       types.Int64  `tfsdk:"delay_minimum"`
	DelayReload        types.Int64  `tfsdk:"delay_reload"`
	BfdMinimumInterval types.Int64  `tfsdk:"bfd_minimum_interval"`
	BfdMultiplier      types.Int64  `tfsdk:"bfd_multiplier"`
}

func (data RouterVRRPInterface) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XR-um-router-vrrp-cfg:/router/vrrp/interfaces/interface[interface-name=%s]", data.InterfaceName.ValueString())
}

func (data RouterVRRPInterfaceData) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XR-um-router-vrrp-cfg:/router/vrrp/interfaces/interface[interface-name=%s]", data.InterfaceName.ValueString())
}

func (data RouterVRRPInterface) toBody(ctx context.Context) string {
	body := "{}"
	if !data.InterfaceName.IsNull() && !data.InterfaceName.IsUnknown() {
		body, _ = sjson.Set(body, "interface-name", data.InterfaceName.ValueString())
	}
	if !data.MacRefresh.IsNull() && !data.MacRefresh.IsUnknown() {
		body, _ = sjson.Set(body, "mac-refresh", strconv.FormatInt(data.MacRefresh.ValueInt64(), 10))
	}
	if !data.DelayMinimum.IsNull() && !data.DelayMinimum.IsUnknown() {
		body, _ = sjson.Set(body, "delay.minimum", strconv.FormatInt(data.DelayMinimum.ValueInt64(), 10))
	}
	if !data.DelayReload.IsNull() && !data.DelayReload.IsUnknown() {
		body, _ = sjson.Set(body, "delay.reload", strconv.FormatInt(data.DelayReload.ValueInt64(), 10))
	}
	if !data.BfdMinimumInterval.IsNull() && !data.BfdMinimumInterval.IsUnknown() {
		body, _ = sjson.Set(body, "bfd.minimum-interval", strconv.FormatInt(data.BfdMinimumInterval.ValueInt64(), 10))
	}
	if !data.BfdMultiplier.IsNull() && !data.BfdMultiplier.IsUnknown() {
		body, _ = sjson.Set(body, "bfd.multiplier", strconv.FormatInt(data.BfdMultiplier.ValueInt64(), 10))
	}
	return body
}

func (data *RouterVRRPInterface) updateFromBody(ctx context.Context, res []byte) {
	if value := gjson.GetBytes(res, "mac-refresh"); value.Exists() && !data.MacRefresh.IsNull() {
		data.MacRefresh = types.Int64Value(value.Int())
	} else {
		data.MacRefresh = types.Int64Null()
	}
	if value := gjson.GetBytes(res, "delay.minimum"); value.Exists() && !data.DelayMinimum.IsNull() {
		data.DelayMinimum = types.Int64Value(value.Int())
	} else {
		data.DelayMinimum = types.Int64Null()
	}
	if value := gjson.GetBytes(res, "delay.reload"); value.Exists() && !data.DelayReload.IsNull() {
		data.DelayReload = types.Int64Value(value.Int())
	} else {
		data.DelayReload = types.Int64Null()
	}
	if value := gjson.GetBytes(res, "bfd.minimum-interval"); value.Exists() && !data.BfdMinimumInterval.IsNull() {
		data.BfdMinimumInterval = types.Int64Value(value.Int())
	} else {
		data.BfdMinimumInterval = types.Int64Null()
	}
	if value := gjson.GetBytes(res, "bfd.multiplier"); value.Exists() && !data.BfdMultiplier.IsNull() {
		data.BfdMultiplier = types.Int64Value(value.Int())
	} else {
		data.BfdMultiplier = types.Int64Null()
	}
}

func (data *RouterVRRPInterfaceData) fromBody(ctx context.Context, res []byte) {
	if value := gjson.GetBytes(res, "mac-refresh"); value.Exists() {
		data.MacRefresh = types.Int64Value(value.Int())
	}
	if value := gjson.GetBytes(res, "delay.minimum"); value.Exists() {
		data.DelayMinimum = types.Int64Value(value.Int())
	}
	if value := gjson.GetBytes(res, "delay.reload"); value.Exists() {
		data.DelayReload = types.Int64Value(value.Int())
	}
	if value := gjson.GetBytes(res, "bfd.minimum-interval"); value.Exists() {
		data.BfdMinimumInterval = types.Int64Value(value.Int())
	}
	if value := gjson.GetBytes(res, "bfd.multiplier"); value.Exists() {
		data.BfdMultiplier = types.Int64Value(value.Int())
	}
}

func (data *RouterVRRPInterface) getDeletedItems(ctx context.Context, state RouterVRRPInterface) []string {
	deletedItems := make([]string, 0)
	if !state.MacRefresh.IsNull() && data.MacRefresh.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/mac-refresh", state.getPath()))
	}
	if !state.DelayMinimum.IsNull() && data.DelayMinimum.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/delay/minimum", state.getPath()))
	}
	if !state.DelayReload.IsNull() && data.DelayReload.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/delay/reload", state.getPath()))
	}
	if !state.BfdMinimumInterval.IsNull() && data.BfdMinimumInterval.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/bfd/minimum-interval", state.getPath()))
	}
	if !state.BfdMultiplier.IsNull() && data.BfdMultiplier.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/bfd/multiplier", state.getPath()))
	}
	return deletedItems
}

func (data *RouterVRRPInterface) getEmptyLeafsDelete(ctx context.Context) []string {
	emptyLeafsDelete := make([]string, 0)
	return emptyLeafsDelete
}

func (data *RouterVRRPInterface) getDeletePaths(ctx context.Context) []string {
	var deletePaths []string
	if !data.MacRefresh.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/mac-refresh", data.getPath()))
	}
	if !data.DelayMinimum.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/delay/minimum", data.getPath()))
	}
	if !data.DelayReload.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/delay/reload", data.getPath()))
	}
	if !data.BfdMinimumInterval.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/bfd/minimum-interval", data.getPath()))
	}
	if !data.BfdMultiplier.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/bfd/multiplier", data.getPath()))
	}
	return deletePaths
}
