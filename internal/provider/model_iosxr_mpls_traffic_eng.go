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

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type MPLSTrafficEng struct {
	Device     types.String `tfsdk:"device"`
	Id         types.String `tfsdk:"id"`
	DeleteMode types.String `tfsdk:"delete_mode"`
	TrafficEng types.Bool   `tfsdk:"traffic_eng"`
}

type MPLSTrafficEngData struct {
	Device     types.String `tfsdk:"device"`
	Id         types.String `tfsdk:"id"`
	TrafficEng types.Bool   `tfsdk:"traffic_eng"`
}

func (data MPLSTrafficEng) getPath() string {
	return "Cisco-IOS-XR-um-mpls-te-cfg:/mpls"
}

func (data MPLSTrafficEngData) getPath() string {
	return "Cisco-IOS-XR-um-mpls-te-cfg:/mpls"
}

func (data MPLSTrafficEng) toBody(ctx context.Context) string {
	body := "{}"
	if !data.TrafficEng.IsNull() && !data.TrafficEng.IsUnknown() {
		if data.TrafficEng.ValueBool() {
			body, _ = sjson.Set(body, "traffic-eng", map[string]string{})
		}
	}
	return body
}

func (data *MPLSTrafficEng) updateFromBody(ctx context.Context, res []byte) {
	if value := gjson.GetBytes(res, "traffic-eng"); !data.TrafficEng.IsNull() {
		if value.Exists() {
			data.TrafficEng = types.BoolValue(true)
		} else {
			data.TrafficEng = types.BoolValue(false)
		}
	} else {
		data.TrafficEng = types.BoolNull()
	}
}

func (data *MPLSTrafficEng) fromBody(ctx context.Context, res []byte) {
	if value := gjson.GetBytes(res, "traffic-eng"); value.Exists() {
		data.TrafficEng = types.BoolValue(true)
	} else {
		data.TrafficEng = types.BoolValue(false)
	}
}

func (data *MPLSTrafficEngData) fromBody(ctx context.Context, res []byte) {
	if value := gjson.GetBytes(res, "traffic-eng"); value.Exists() {
		data.TrafficEng = types.BoolValue(true)
	} else {
		data.TrafficEng = types.BoolValue(false)
	}
}

func (data *MPLSTrafficEng) getDeletedItems(ctx context.Context, state MPLSTrafficEng) []string {
	deletedItems := make([]string, 0)
	if !state.TrafficEng.IsNull() && data.TrafficEng.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/traffic-eng", state.getPath()))
	}
	return deletedItems
}

func (data *MPLSTrafficEng) getEmptyLeafsDelete(ctx context.Context) []string {
	emptyLeafsDelete := make([]string, 0)
	if !data.TrafficEng.IsNull() && !data.TrafficEng.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/traffic-eng", data.getPath()))
	}
	return emptyLeafsDelete
}

func (data *MPLSTrafficEng) getDeletePaths(ctx context.Context) []string {
	var deletePaths []string
	if !data.TrafficEng.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/traffic-eng", data.getPath()))
	}
	return deletePaths
}