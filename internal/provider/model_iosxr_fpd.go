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

type FPD struct {
	Device             types.String `tfsdk:"device"`
	Id                 types.String `tfsdk:"id"`
	DeleteMode         types.String `tfsdk:"delete_mode"`
	AutoUpgradeEnable  types.Bool   `tfsdk:"auto_upgrade_enable"`
	AutoUpgradeDisable types.Bool   `tfsdk:"auto_upgrade_disable"`
	AutoReloadEnable   types.Bool   `tfsdk:"auto_reload_enable"`
	AutoReloadDisable  types.Bool   `tfsdk:"auto_reload_disable"`
}

type FPDData struct {
	Device             types.String `tfsdk:"device"`
	Id                 types.String `tfsdk:"id"`
	AutoUpgradeEnable  types.Bool   `tfsdk:"auto_upgrade_enable"`
	AutoUpgradeDisable types.Bool   `tfsdk:"auto_upgrade_disable"`
	AutoReloadEnable   types.Bool   `tfsdk:"auto_reload_enable"`
	AutoReloadDisable  types.Bool   `tfsdk:"auto_reload_disable"`
}

func (data FPD) getPath() string {
	return "Cisco-IOS-XR-um-fpd-cfg:/fpd"
}

func (data FPDData) getPath() string {
	return "Cisco-IOS-XR-um-fpd-cfg:/fpd"
}

func (data FPD) toBody(ctx context.Context) string {
	body := "{}"
	if !data.AutoUpgradeEnable.IsNull() && !data.AutoUpgradeEnable.IsUnknown() {
		if data.AutoUpgradeEnable.ValueBool() {
			body, _ = sjson.Set(body, "auto-upgrade.enable", map[string]string{})
		}
	}
	if !data.AutoUpgradeDisable.IsNull() && !data.AutoUpgradeDisable.IsUnknown() {
		if data.AutoUpgradeDisable.ValueBool() {
			body, _ = sjson.Set(body, "auto-upgrade.disable", map[string]string{})
		}
	}
	if !data.AutoReloadEnable.IsNull() && !data.AutoReloadEnable.IsUnknown() {
		if data.AutoReloadEnable.ValueBool() {
			body, _ = sjson.Set(body, "auto-reload.enable", map[string]string{})
		}
	}
	if !data.AutoReloadDisable.IsNull() && !data.AutoReloadDisable.IsUnknown() {
		if data.AutoReloadDisable.ValueBool() {
			body, _ = sjson.Set(body, "auto-reload.disable", map[string]string{})
		}
	}
	return body
}

func (data *FPD) updateFromBody(ctx context.Context, res []byte) {
	if value := gjson.GetBytes(res, "auto-upgrade.enable"); !data.AutoUpgradeEnable.IsNull() {
		if value.Exists() {
			data.AutoUpgradeEnable = types.BoolValue(true)
		} else {
			data.AutoUpgradeEnable = types.BoolValue(false)
		}
	} else {
		data.AutoUpgradeEnable = types.BoolNull()
	}
	if value := gjson.GetBytes(res, "auto-upgrade.disable"); !data.AutoUpgradeDisable.IsNull() {
		if value.Exists() {
			data.AutoUpgradeDisable = types.BoolValue(true)
		} else {
			data.AutoUpgradeDisable = types.BoolValue(false)
		}
	} else {
		data.AutoUpgradeDisable = types.BoolNull()
	}
	if value := gjson.GetBytes(res, "auto-reload.enable"); !data.AutoReloadEnable.IsNull() {
		if value.Exists() {
			data.AutoReloadEnable = types.BoolValue(true)
		} else {
			data.AutoReloadEnable = types.BoolValue(false)
		}
	} else {
		data.AutoReloadEnable = types.BoolNull()
	}
	if value := gjson.GetBytes(res, "auto-reload.disable"); !data.AutoReloadDisable.IsNull() {
		if value.Exists() {
			data.AutoReloadDisable = types.BoolValue(true)
		} else {
			data.AutoReloadDisable = types.BoolValue(false)
		}
	} else {
		data.AutoReloadDisable = types.BoolNull()
	}
}

func (data *FPDData) fromBody(ctx context.Context, res []byte) {
	if value := gjson.GetBytes(res, "auto-upgrade.enable"); value.Exists() {
		data.AutoUpgradeEnable = types.BoolValue(true)
	} else {
		data.AutoUpgradeEnable = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "auto-upgrade.disable"); value.Exists() {
		data.AutoUpgradeDisable = types.BoolValue(true)
	} else {
		data.AutoUpgradeDisable = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "auto-reload.enable"); value.Exists() {
		data.AutoReloadEnable = types.BoolValue(true)
	} else {
		data.AutoReloadEnable = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "auto-reload.disable"); value.Exists() {
		data.AutoReloadDisable = types.BoolValue(true)
	} else {
		data.AutoReloadDisable = types.BoolValue(false)
	}
}

func (data *FPD) getDeletedItems(ctx context.Context, state FPD) []string {
	deletedItems := make([]string, 0)
	if !state.AutoUpgradeEnable.IsNull() && data.AutoUpgradeEnable.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/auto-upgrade/enable", state.getPath()))
	}
	if !state.AutoUpgradeDisable.IsNull() && data.AutoUpgradeDisable.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/auto-upgrade/disable", state.getPath()))
	}
	if !state.AutoReloadEnable.IsNull() && data.AutoReloadEnable.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/auto-reload/enable", state.getPath()))
	}
	if !state.AutoReloadDisable.IsNull() && data.AutoReloadDisable.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/auto-reload/disable", state.getPath()))
	}
	return deletedItems
}

func (data *FPD) getEmptyLeafsDelete(ctx context.Context) []string {
	emptyLeafsDelete := make([]string, 0)
	if !data.AutoUpgradeEnable.IsNull() && !data.AutoUpgradeEnable.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/auto-upgrade/enable", data.getPath()))
	}
	if !data.AutoUpgradeDisable.IsNull() && !data.AutoUpgradeDisable.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/auto-upgrade/disable", data.getPath()))
	}
	if !data.AutoReloadEnable.IsNull() && !data.AutoReloadEnable.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/auto-reload/enable", data.getPath()))
	}
	if !data.AutoReloadDisable.IsNull() && !data.AutoReloadDisable.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/auto-reload/disable", data.getPath()))
	}
	return emptyLeafsDelete
}

func (data *FPD) getDeletePaths(ctx context.Context) []string {
	var deletePaths []string
	if !data.AutoUpgradeEnable.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/auto-upgrade/enable", data.getPath()))
	}
	if !data.AutoUpgradeDisable.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/auto-upgrade/disable", data.getPath()))
	}
	if !data.AutoReloadEnable.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/auto-reload/enable", data.getPath()))
	}
	if !data.AutoReloadDisable.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/auto-reload/disable", data.getPath()))
	}
	return deletePaths
}
