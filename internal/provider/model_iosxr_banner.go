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

type Banner struct {
	Device     types.String `tfsdk:"device"`
	Id         types.String `tfsdk:"id"`
	BannerType types.String `tfsdk:"banner_type"`
	Line       types.String `tfsdk:"line"`
}

type BannerData struct {
	Device     types.String `tfsdk:"device"`
	Id         types.String `tfsdk:"id"`
	BannerType types.String `tfsdk:"banner_type"`
	Line       types.String `tfsdk:"line"`
}

func (data Banner) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XR-um-banner-cfg:/banners/banner[banner-type=%s]", data.BannerType.ValueString())
}

func (data BannerData) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XR-um-banner-cfg:/banners/banner[banner-type=%s]", data.BannerType.ValueString())
}

func (data Banner) toBody(ctx context.Context) string {
	body := "{}"
	if !data.BannerType.IsNull() && !data.BannerType.IsUnknown() {
		body, _ = sjson.Set(body, "banner-type", data.BannerType.ValueString())
	}
	if !data.Line.IsNull() && !data.Line.IsUnknown() {
		body, _ = sjson.Set(body, "line", data.Line.ValueString())
	}
	return body
}

func (data *Banner) updateFromBody(ctx context.Context, res []byte) {
	if value := gjson.GetBytes(res, "line"); value.Exists() && !data.Line.IsNull() {
		data.Line = types.StringValue(value.String())
	} else {
		data.Line = types.StringNull()
	}
}

func (data *BannerData) fromBody(ctx context.Context, res []byte) {
	if value := gjson.GetBytes(res, "line"); value.Exists() {
		data.Line = types.StringValue(value.String())
	}
}

func (data *Banner) getDeletedItems(ctx context.Context, state Banner) []string {
	deletedItems := make([]string, 0)
	if !state.Line.IsNull() && data.Line.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/line", state.getPath()))
	}
	return deletedItems
}

func (data *Banner) getEmptyLeafsDelete(ctx context.Context) []string {
	emptyLeafsDelete := make([]string, 0)
	return emptyLeafsDelete
}

func (data *Banner) getDeletePaths(ctx context.Context) []string {
	var deletePaths []string
	if !data.Line.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/line", data.getPath()))
	}
	return deletePaths
}
