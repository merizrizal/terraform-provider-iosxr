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

type RouterISISInterfaceAddressFamily struct {
	Device                     types.String                                                 `tfsdk:"device"`
	Id                         types.String                                                 `tfsdk:"id"`
	DeleteMode                 types.String                                                 `tfsdk:"delete_mode"`
	ProcessId                  types.String                                                 `tfsdk:"process_id"`
	InterfaceName              types.String                                                 `tfsdk:"interface_name"`
	AfName                     types.String                                                 `tfsdk:"af_name"`
	SafName                    types.String                                                 `tfsdk:"saf_name"`
	FastReroutePerPrefixLevels []RouterISISInterfaceAddressFamilyFastReroutePerPrefixLevels `tfsdk:"fast_reroute_per_prefix_levels"`
	Tag                        types.Int64                                                  `tfsdk:"tag"`
	PrefixSidAbsolute          types.Int64                                                  `tfsdk:"prefix_sid_absolute"`
	PrefixSidNFlagClear        types.Bool                                                   `tfsdk:"prefix_sid_n_flag_clear"`
	AdvertisePrefixRoutePolicy types.String                                                 `tfsdk:"advertise_prefix_route_policy"`
	PrefixSidIndex             types.Int64                                                  `tfsdk:"prefix_sid_index"`
	PrefixSidStrictSpfAbsolute types.Int64                                                  `tfsdk:"prefix_sid_strict_spf_absolute"`
	Metric                     types.Int64                                                  `tfsdk:"metric"`
	MetricMaximum              types.Bool                                                   `tfsdk:"metric_maximum"`
	MetricLevels               []RouterISISInterfaceAddressFamilyMetricLevels               `tfsdk:"metric_levels"`
}

type RouterISISInterfaceAddressFamilyData struct {
	Device                     types.String                                                 `tfsdk:"device"`
	Id                         types.String                                                 `tfsdk:"id"`
	ProcessId                  types.String                                                 `tfsdk:"process_id"`
	InterfaceName              types.String                                                 `tfsdk:"interface_name"`
	AfName                     types.String                                                 `tfsdk:"af_name"`
	SafName                    types.String                                                 `tfsdk:"saf_name"`
	FastReroutePerPrefixLevels []RouterISISInterfaceAddressFamilyFastReroutePerPrefixLevels `tfsdk:"fast_reroute_per_prefix_levels"`
	Tag                        types.Int64                                                  `tfsdk:"tag"`
	PrefixSidAbsolute          types.Int64                                                  `tfsdk:"prefix_sid_absolute"`
	PrefixSidNFlagClear        types.Bool                                                   `tfsdk:"prefix_sid_n_flag_clear"`
	AdvertisePrefixRoutePolicy types.String                                                 `tfsdk:"advertise_prefix_route_policy"`
	PrefixSidIndex             types.Int64                                                  `tfsdk:"prefix_sid_index"`
	PrefixSidStrictSpfAbsolute types.Int64                                                  `tfsdk:"prefix_sid_strict_spf_absolute"`
	Metric                     types.Int64                                                  `tfsdk:"metric"`
	MetricMaximum              types.Bool                                                   `tfsdk:"metric_maximum"`
	MetricLevels               []RouterISISInterfaceAddressFamilyMetricLevels               `tfsdk:"metric_levels"`
}
type RouterISISInterfaceAddressFamilyFastReroutePerPrefixLevels struct {
	LevelId types.Int64 `tfsdk:"level_id"`
	TiLfa   types.Bool  `tfsdk:"ti_lfa"`
}
type RouterISISInterfaceAddressFamilyMetricLevels struct {
	LevelId types.Int64 `tfsdk:"level_id"`
	Metric  types.Int64 `tfsdk:"metric"`
	Maximum types.Bool  `tfsdk:"maximum"`
}

func (data RouterISISInterfaceAddressFamily) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XR-um-router-isis-cfg:/router/isis/processes/process[process-id=%s]/interfaces/interface[interface-name=%s]/address-families/address-family[af-name=%s][saf-name=%s]", data.ProcessId.ValueString(), data.InterfaceName.ValueString(), data.AfName.ValueString(), data.SafName.ValueString())
}

func (data RouterISISInterfaceAddressFamilyData) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XR-um-router-isis-cfg:/router/isis/processes/process[process-id=%s]/interfaces/interface[interface-name=%s]/address-families/address-family[af-name=%s][saf-name=%s]", data.ProcessId.ValueString(), data.InterfaceName.ValueString(), data.AfName.ValueString(), data.SafName.ValueString())
}

func (data RouterISISInterfaceAddressFamily) toBody(ctx context.Context) string {
	body := "{}"
	if !data.AfName.IsNull() && !data.AfName.IsUnknown() {
		body, _ = sjson.Set(body, "af-name", data.AfName.ValueString())
	}
	if !data.SafName.IsNull() && !data.SafName.IsUnknown() {
		body, _ = sjson.Set(body, "saf-name", data.SafName.ValueString())
	}
	if !data.Tag.IsNull() && !data.Tag.IsUnknown() {
		body, _ = sjson.Set(body, "tag.interface-tag", strconv.FormatInt(data.Tag.ValueInt64(), 10))
	}
	if !data.PrefixSidAbsolute.IsNull() && !data.PrefixSidAbsolute.IsUnknown() {
		body, _ = sjson.Set(body, "prefix-sid.sid.absolute.sid-value", strconv.FormatInt(data.PrefixSidAbsolute.ValueInt64(), 10))
	}
	if !data.PrefixSidNFlagClear.IsNull() && !data.PrefixSidNFlagClear.IsUnknown() {
		if data.PrefixSidNFlagClear.ValueBool() {
			body, _ = sjson.Set(body, "prefix-sid.sid.n-flag-clear", map[string]string{})
		}
	}
	if !data.AdvertisePrefixRoutePolicy.IsNull() && !data.AdvertisePrefixRoutePolicy.IsUnknown() {
		body, _ = sjson.Set(body, "advertise.prefix.route-policy", data.AdvertisePrefixRoutePolicy.ValueString())
	}
	if !data.PrefixSidIndex.IsNull() && !data.PrefixSidIndex.IsUnknown() {
		body, _ = sjson.Set(body, "prefix-sid.sid.index.sid-index", strconv.FormatInt(data.PrefixSidIndex.ValueInt64(), 10))
	}
	if !data.PrefixSidStrictSpfAbsolute.IsNull() && !data.PrefixSidStrictSpfAbsolute.IsUnknown() {
		body, _ = sjson.Set(body, "prefix-sid.strict-spf.absolute.sid-value", strconv.FormatInt(data.PrefixSidStrictSpfAbsolute.ValueInt64(), 10))
	}
	if !data.Metric.IsNull() && !data.Metric.IsUnknown() {
		body, _ = sjson.Set(body, "metric.default-metric", strconv.FormatInt(data.Metric.ValueInt64(), 10))
	}
	if !data.MetricMaximum.IsNull() && !data.MetricMaximum.IsUnknown() {
		if data.MetricMaximum.ValueBool() {
			body, _ = sjson.Set(body, "metric.maximum", map[string]string{})
		}
	}
	if len(data.FastReroutePerPrefixLevels) > 0 {
		body, _ = sjson.Set(body, "fast-reroute.per-prefix.per-prefix.levels.level", []interface{}{})
		for index, item := range data.FastReroutePerPrefixLevels {
			if !item.LevelId.IsNull() && !item.LevelId.IsUnknown() {
				body, _ = sjson.Set(body, "fast-reroute.per-prefix.per-prefix.levels.level"+"."+strconv.Itoa(index)+"."+"level-id", strconv.FormatInt(item.LevelId.ValueInt64(), 10))
			}
			if !item.TiLfa.IsNull() && !item.TiLfa.IsUnknown() {
				if item.TiLfa.ValueBool() {
					body, _ = sjson.Set(body, "fast-reroute.per-prefix.per-prefix.levels.level"+"."+strconv.Itoa(index)+"."+"ti-lfa", map[string]string{})
				}
			}
		}
	}
	if len(data.MetricLevels) > 0 {
		body, _ = sjson.Set(body, "metric.levels.level", []interface{}{})
		for index, item := range data.MetricLevels {
			if !item.LevelId.IsNull() && !item.LevelId.IsUnknown() {
				body, _ = sjson.Set(body, "metric.levels.level"+"."+strconv.Itoa(index)+"."+"level-id", strconv.FormatInt(item.LevelId.ValueInt64(), 10))
			}
			if !item.Metric.IsNull() && !item.Metric.IsUnknown() {
				body, _ = sjson.Set(body, "metric.levels.level"+"."+strconv.Itoa(index)+"."+"default-metric", strconv.FormatInt(item.Metric.ValueInt64(), 10))
			}
			if !item.Maximum.IsNull() && !item.Maximum.IsUnknown() {
				if item.Maximum.ValueBool() {
					body, _ = sjson.Set(body, "metric.levels.level"+"."+strconv.Itoa(index)+"."+"maximum", map[string]string{})
				}
			}
		}
	}
	return body
}

func (data *RouterISISInterfaceAddressFamily) updateFromBody(ctx context.Context, res []byte) {
	for i := range data.FastReroutePerPrefixLevels {
		keys := [...]string{"level-id"}
		keyValues := [...]string{strconv.FormatInt(data.FastReroutePerPrefixLevels[i].LevelId.ValueInt64(), 10)}

		var r gjson.Result
		gjson.GetBytes(res, "fast-reroute.per-prefix.per-prefix.levels.level").ForEach(
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
		if value := r.Get("level-id"); value.Exists() && !data.FastReroutePerPrefixLevels[i].LevelId.IsNull() {
			data.FastReroutePerPrefixLevels[i].LevelId = types.Int64Value(value.Int())
		} else {
			data.FastReroutePerPrefixLevels[i].LevelId = types.Int64Null()
		}
		if value := r.Get("ti-lfa"); !data.FastReroutePerPrefixLevels[i].TiLfa.IsNull() {
			if value.Exists() {
				data.FastReroutePerPrefixLevels[i].TiLfa = types.BoolValue(true)
			} else {
				data.FastReroutePerPrefixLevels[i].TiLfa = types.BoolValue(false)
			}
		} else {
			data.FastReroutePerPrefixLevels[i].TiLfa = types.BoolNull()
		}
	}
	if value := gjson.GetBytes(res, "tag.interface-tag"); value.Exists() && !data.Tag.IsNull() {
		data.Tag = types.Int64Value(value.Int())
	} else {
		data.Tag = types.Int64Null()
	}
	if value := gjson.GetBytes(res, "prefix-sid.sid.absolute.sid-value"); value.Exists() && !data.PrefixSidAbsolute.IsNull() {
		data.PrefixSidAbsolute = types.Int64Value(value.Int())
	} else {
		data.PrefixSidAbsolute = types.Int64Null()
	}
	if value := gjson.GetBytes(res, "prefix-sid.sid.n-flag-clear"); !data.PrefixSidNFlagClear.IsNull() {
		if value.Exists() {
			data.PrefixSidNFlagClear = types.BoolValue(true)
		} else {
			data.PrefixSidNFlagClear = types.BoolValue(false)
		}
	} else {
		data.PrefixSidNFlagClear = types.BoolNull()
	}
	if value := gjson.GetBytes(res, "advertise.prefix.route-policy"); value.Exists() && !data.AdvertisePrefixRoutePolicy.IsNull() {
		data.AdvertisePrefixRoutePolicy = types.StringValue(value.String())
	} else {
		data.AdvertisePrefixRoutePolicy = types.StringNull()
	}
	if value := gjson.GetBytes(res, "prefix-sid.sid.index.sid-index"); value.Exists() && !data.PrefixSidIndex.IsNull() {
		data.PrefixSidIndex = types.Int64Value(value.Int())
	} else {
		data.PrefixSidIndex = types.Int64Null()
	}
	if value := gjson.GetBytes(res, "prefix-sid.strict-spf.absolute.sid-value"); value.Exists() && !data.PrefixSidStrictSpfAbsolute.IsNull() {
		data.PrefixSidStrictSpfAbsolute = types.Int64Value(value.Int())
	} else {
		data.PrefixSidStrictSpfAbsolute = types.Int64Null()
	}
	if value := gjson.GetBytes(res, "metric.default-metric"); value.Exists() && !data.Metric.IsNull() {
		data.Metric = types.Int64Value(value.Int())
	} else {
		data.Metric = types.Int64Null()
	}
	if value := gjson.GetBytes(res, "metric.maximum"); !data.MetricMaximum.IsNull() {
		if value.Exists() {
			data.MetricMaximum = types.BoolValue(true)
		} else {
			data.MetricMaximum = types.BoolValue(false)
		}
	} else {
		data.MetricMaximum = types.BoolNull()
	}
	for i := range data.MetricLevels {
		keys := [...]string{"level-id"}
		keyValues := [...]string{strconv.FormatInt(data.MetricLevels[i].LevelId.ValueInt64(), 10)}

		var r gjson.Result
		gjson.GetBytes(res, "metric.levels.level").ForEach(
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
		if value := r.Get("level-id"); value.Exists() && !data.MetricLevels[i].LevelId.IsNull() {
			data.MetricLevels[i].LevelId = types.Int64Value(value.Int())
		} else {
			data.MetricLevels[i].LevelId = types.Int64Null()
		}
		if value := r.Get("default-metric"); value.Exists() && !data.MetricLevels[i].Metric.IsNull() {
			data.MetricLevels[i].Metric = types.Int64Value(value.Int())
		} else {
			data.MetricLevels[i].Metric = types.Int64Null()
		}
		if value := r.Get("maximum"); !data.MetricLevels[i].Maximum.IsNull() {
			if value.Exists() {
				data.MetricLevels[i].Maximum = types.BoolValue(true)
			} else {
				data.MetricLevels[i].Maximum = types.BoolValue(false)
			}
		} else {
			data.MetricLevels[i].Maximum = types.BoolNull()
		}
	}
}

func (data *RouterISISInterfaceAddressFamilyData) fromBody(ctx context.Context, res []byte) {
	if value := gjson.GetBytes(res, "fast-reroute.per-prefix.per-prefix.levels.level"); value.Exists() {
		data.FastReroutePerPrefixLevels = make([]RouterISISInterfaceAddressFamilyFastReroutePerPrefixLevels, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := RouterISISInterfaceAddressFamilyFastReroutePerPrefixLevels{}
			if cValue := v.Get("level-id"); cValue.Exists() {
				item.LevelId = types.Int64Value(cValue.Int())
			}
			if cValue := v.Get("ti-lfa"); cValue.Exists() {
				item.TiLfa = types.BoolValue(true)
			} else {
				item.TiLfa = types.BoolValue(false)
			}
			data.FastReroutePerPrefixLevels = append(data.FastReroutePerPrefixLevels, item)
			return true
		})
	}
	if value := gjson.GetBytes(res, "tag.interface-tag"); value.Exists() {
		data.Tag = types.Int64Value(value.Int())
	}
	if value := gjson.GetBytes(res, "prefix-sid.sid.absolute.sid-value"); value.Exists() {
		data.PrefixSidAbsolute = types.Int64Value(value.Int())
	}
	if value := gjson.GetBytes(res, "prefix-sid.sid.n-flag-clear"); value.Exists() {
		data.PrefixSidNFlagClear = types.BoolValue(true)
	} else {
		data.PrefixSidNFlagClear = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "advertise.prefix.route-policy"); value.Exists() {
		data.AdvertisePrefixRoutePolicy = types.StringValue(value.String())
	}
	if value := gjson.GetBytes(res, "prefix-sid.sid.index.sid-index"); value.Exists() {
		data.PrefixSidIndex = types.Int64Value(value.Int())
	}
	if value := gjson.GetBytes(res, "prefix-sid.strict-spf.absolute.sid-value"); value.Exists() {
		data.PrefixSidStrictSpfAbsolute = types.Int64Value(value.Int())
	}
	if value := gjson.GetBytes(res, "metric.default-metric"); value.Exists() {
		data.Metric = types.Int64Value(value.Int())
	}
	if value := gjson.GetBytes(res, "metric.maximum"); value.Exists() {
		data.MetricMaximum = types.BoolValue(true)
	} else {
		data.MetricMaximum = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "metric.levels.level"); value.Exists() {
		data.MetricLevels = make([]RouterISISInterfaceAddressFamilyMetricLevels, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := RouterISISInterfaceAddressFamilyMetricLevels{}
			if cValue := v.Get("level-id"); cValue.Exists() {
				item.LevelId = types.Int64Value(cValue.Int())
			}
			if cValue := v.Get("default-metric"); cValue.Exists() {
				item.Metric = types.Int64Value(cValue.Int())
			}
			if cValue := v.Get("maximum"); cValue.Exists() {
				item.Maximum = types.BoolValue(true)
			} else {
				item.Maximum = types.BoolValue(false)
			}
			data.MetricLevels = append(data.MetricLevels, item)
			return true
		})
	}
}

func (data *RouterISISInterfaceAddressFamily) getDeletedItems(ctx context.Context, state RouterISISInterfaceAddressFamily) []string {
	deletedItems := make([]string, 0)
	for i := range state.FastReroutePerPrefixLevels {
		keys := [...]string{"level-id"}
		stateKeyValues := [...]string{strconv.FormatInt(state.FastReroutePerPrefixLevels[i].LevelId.ValueInt64(), 10)}
		keyString := ""
		for ki := range keys {
			keyString += "[" + keys[ki] + "=" + stateKeyValues[ki] + "]"
		}

		emptyKeys := true
		if !reflect.ValueOf(state.FastReroutePerPrefixLevels[i].LevelId.ValueInt64()).IsZero() {
			emptyKeys = false
		}
		if emptyKeys {
			continue
		}

		found := false
		for j := range data.FastReroutePerPrefixLevels {
			found = true
			if state.FastReroutePerPrefixLevels[i].LevelId.ValueInt64() != data.FastReroutePerPrefixLevels[j].LevelId.ValueInt64() {
				found = false
			}
			if found {
				if !state.FastReroutePerPrefixLevels[i].TiLfa.IsNull() && data.FastReroutePerPrefixLevels[j].TiLfa.IsNull() {
					deletedItems = append(deletedItems, fmt.Sprintf("%v/fast-reroute/per-prefix/per-prefix/levels/level%v/ti-lfa", state.getPath(), keyString))
				}
				break
			}
		}
		if !found {
			deletedItems = append(deletedItems, fmt.Sprintf("%v/fast-reroute/per-prefix/per-prefix/levels/level%v", state.getPath(), keyString))
		}
	}
	if !state.Tag.IsNull() && data.Tag.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/tag/interface-tag", state.getPath()))
	}
	if !state.PrefixSidAbsolute.IsNull() && data.PrefixSidAbsolute.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/prefix-sid/sid/absolute/sid-value", state.getPath()))
	}
	if !state.PrefixSidNFlagClear.IsNull() && data.PrefixSidNFlagClear.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/prefix-sid/sid/n-flag-clear", state.getPath()))
	}
	if !state.AdvertisePrefixRoutePolicy.IsNull() && data.AdvertisePrefixRoutePolicy.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/advertise/prefix/route-policy", state.getPath()))
	}
	if !state.PrefixSidIndex.IsNull() && data.PrefixSidIndex.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/prefix-sid/sid/index/sid-index", state.getPath()))
	}
	if !state.PrefixSidStrictSpfAbsolute.IsNull() && data.PrefixSidStrictSpfAbsolute.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/prefix-sid/strict-spf/absolute/sid-value", state.getPath()))
	}
	if !state.Metric.IsNull() && data.Metric.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/metric/default-metric", state.getPath()))
	}
	if !state.MetricMaximum.IsNull() && data.MetricMaximum.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/metric/maximum", state.getPath()))
	}
	for i := range state.MetricLevels {
		keys := [...]string{"level-id"}
		stateKeyValues := [...]string{strconv.FormatInt(state.MetricLevels[i].LevelId.ValueInt64(), 10)}
		keyString := ""
		for ki := range keys {
			keyString += "[" + keys[ki] + "=" + stateKeyValues[ki] + "]"
		}

		emptyKeys := true
		if !reflect.ValueOf(state.MetricLevels[i].LevelId.ValueInt64()).IsZero() {
			emptyKeys = false
		}
		if emptyKeys {
			continue
		}

		found := false
		for j := range data.MetricLevels {
			found = true
			if state.MetricLevels[i].LevelId.ValueInt64() != data.MetricLevels[j].LevelId.ValueInt64() {
				found = false
			}
			if found {
				if !state.MetricLevels[i].Metric.IsNull() && data.MetricLevels[j].Metric.IsNull() {
					deletedItems = append(deletedItems, fmt.Sprintf("%v/metric/levels/level%v/default-metric", state.getPath(), keyString))
				}
				if !state.MetricLevels[i].Maximum.IsNull() && data.MetricLevels[j].Maximum.IsNull() {
					deletedItems = append(deletedItems, fmt.Sprintf("%v/metric/levels/level%v/maximum", state.getPath(), keyString))
				}
				break
			}
		}
		if !found {
			deletedItems = append(deletedItems, fmt.Sprintf("%v/metric/levels/level%v", state.getPath(), keyString))
		}
	}
	return deletedItems
}

func (data *RouterISISInterfaceAddressFamily) getEmptyLeafsDelete(ctx context.Context) []string {
	emptyLeafsDelete := make([]string, 0)
	for i := range data.FastReroutePerPrefixLevels {
		keys := [...]string{"level-id"}
		keyValues := [...]string{strconv.FormatInt(data.FastReroutePerPrefixLevels[i].LevelId.ValueInt64(), 10)}
		keyString := ""
		for ki := range keys {
			keyString += "[" + keys[ki] + "=" + keyValues[ki] + "]"
		}
		if !data.FastReroutePerPrefixLevels[i].TiLfa.IsNull() && !data.FastReroutePerPrefixLevels[i].TiLfa.ValueBool() {
			emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/fast-reroute/per-prefix/per-prefix/levels/level%v/ti-lfa", data.getPath(), keyString))
		}
	}
	if !data.PrefixSidNFlagClear.IsNull() && !data.PrefixSidNFlagClear.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/prefix-sid/sid/n-flag-clear", data.getPath()))
	}
	if !data.MetricMaximum.IsNull() && !data.MetricMaximum.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/metric/maximum", data.getPath()))
	}
	for i := range data.MetricLevels {
		keys := [...]string{"level-id"}
		keyValues := [...]string{strconv.FormatInt(data.MetricLevels[i].LevelId.ValueInt64(), 10)}
		keyString := ""
		for ki := range keys {
			keyString += "[" + keys[ki] + "=" + keyValues[ki] + "]"
		}
		if !data.MetricLevels[i].Maximum.IsNull() && !data.MetricLevels[i].Maximum.ValueBool() {
			emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/metric/levels/level%v/maximum", data.getPath(), keyString))
		}
	}
	return emptyLeafsDelete
}

func (data *RouterISISInterfaceAddressFamily) getDeletePaths(ctx context.Context) []string {
	var deletePaths []string
	for i := range data.FastReroutePerPrefixLevels {
		keys := [...]string{"level-id"}
		keyValues := [...]string{strconv.FormatInt(data.FastReroutePerPrefixLevels[i].LevelId.ValueInt64(), 10)}

		keyString := ""
		for ki := range keys {
			keyString += "[" + keys[ki] + "=" + keyValues[ki] + "]"
		}
		deletePaths = append(deletePaths, fmt.Sprintf("%v/fast-reroute/per-prefix/per-prefix/levels/level%v", data.getPath(), keyString))
	}
	if !data.Tag.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/tag/interface-tag", data.getPath()))
	}
	if !data.PrefixSidAbsolute.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/prefix-sid/sid/absolute/sid-value", data.getPath()))
	}
	if !data.PrefixSidNFlagClear.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/prefix-sid/sid/n-flag-clear", data.getPath()))
	}
	if !data.AdvertisePrefixRoutePolicy.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/advertise/prefix/route-policy", data.getPath()))
	}
	if !data.PrefixSidIndex.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/prefix-sid/sid/index/sid-index", data.getPath()))
	}
	if !data.PrefixSidStrictSpfAbsolute.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/prefix-sid/strict-spf/absolute/sid-value", data.getPath()))
	}
	if !data.Metric.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/metric/default-metric", data.getPath()))
	}
	if !data.MetricMaximum.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/metric/maximum", data.getPath()))
	}
	for i := range data.MetricLevels {
		keys := [...]string{"level-id"}
		keyValues := [...]string{strconv.FormatInt(data.MetricLevels[i].LevelId.ValueInt64(), 10)}

		keyString := ""
		for ki := range keys {
			keyString += "[" + keys[ki] + "=" + keyValues[ki] + "]"
		}
		deletePaths = append(deletePaths, fmt.Sprintf("%v/metric/levels/level%v", data.getPath(), keyString))
	}
	return deletePaths
}
