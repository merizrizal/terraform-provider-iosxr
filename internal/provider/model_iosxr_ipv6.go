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

type IPv6 struct {
	Device                           types.String `tfsdk:"device"`
	Id                               types.String `tfsdk:"id"`
	DeleteMode                       types.String `tfsdk:"delete_mode"`
	HopLimit                         types.Int64  `tfsdk:"hop_limit"`
	IcmpErrorInterval                types.Int64  `tfsdk:"icmp_error_interval"`
	IcmpErrorIntervalBucketSize      types.Int64  `tfsdk:"icmp_error_interval_bucket_size"`
	SourceRoute                      types.Bool   `tfsdk:"source_route"`
	AssemblerTimeout                 types.Int64  `tfsdk:"assembler_timeout"`
	AssemblerMaxPackets              types.Int64  `tfsdk:"assembler_max_packets"`
	AssemblerReassemblerDropEnable   types.Bool   `tfsdk:"assembler_reassembler_drop_enable"`
	AssemblerFragHdrIncompleteEnable types.Bool   `tfsdk:"assembler_frag_hdr_incomplete_enable"`
	AssemblerOverlapFragDropEnable   types.Bool   `tfsdk:"assembler_overlap_frag_drop_enable"`
	PathMtuEnable                    types.Bool   `tfsdk:"path_mtu_enable"`
	PathMtuTimeout                   types.Int64  `tfsdk:"path_mtu_timeout"`
}

type IPv6Data struct {
	Device                           types.String `tfsdk:"device"`
	Id                               types.String `tfsdk:"id"`
	HopLimit                         types.Int64  `tfsdk:"hop_limit"`
	IcmpErrorInterval                types.Int64  `tfsdk:"icmp_error_interval"`
	IcmpErrorIntervalBucketSize      types.Int64  `tfsdk:"icmp_error_interval_bucket_size"`
	SourceRoute                      types.Bool   `tfsdk:"source_route"`
	AssemblerTimeout                 types.Int64  `tfsdk:"assembler_timeout"`
	AssemblerMaxPackets              types.Int64  `tfsdk:"assembler_max_packets"`
	AssemblerReassemblerDropEnable   types.Bool   `tfsdk:"assembler_reassembler_drop_enable"`
	AssemblerFragHdrIncompleteEnable types.Bool   `tfsdk:"assembler_frag_hdr_incomplete_enable"`
	AssemblerOverlapFragDropEnable   types.Bool   `tfsdk:"assembler_overlap_frag_drop_enable"`
	PathMtuEnable                    types.Bool   `tfsdk:"path_mtu_enable"`
	PathMtuTimeout                   types.Int64  `tfsdk:"path_mtu_timeout"`
}

func (data IPv6) getPath() string {
	return "Cisco-IOS-XR-um-ipv6-cfg:/ipv6"
}

func (data IPv6Data) getPath() string {
	return "Cisco-IOS-XR-um-ipv6-cfg:/ipv6"
}

func (data IPv6) toBody(ctx context.Context) string {
	body := "{}"
	if !data.HopLimit.IsNull() && !data.HopLimit.IsUnknown() {
		body, _ = sjson.Set(body, "hop-limit", strconv.FormatInt(data.HopLimit.ValueInt64(), 10))
	}
	if !data.IcmpErrorInterval.IsNull() && !data.IcmpErrorInterval.IsUnknown() {
		body, _ = sjson.Set(body, "icmp.error-interval.interval-time", strconv.FormatInt(data.IcmpErrorInterval.ValueInt64(), 10))
	}
	if !data.IcmpErrorIntervalBucketSize.IsNull() && !data.IcmpErrorIntervalBucketSize.IsUnknown() {
		body, _ = sjson.Set(body, "icmp.error-interval.bucket-size", strconv.FormatInt(data.IcmpErrorIntervalBucketSize.ValueInt64(), 10))
	}
	if !data.SourceRoute.IsNull() && !data.SourceRoute.IsUnknown() {
		if data.SourceRoute.ValueBool() {
			body, _ = sjson.Set(body, "source-route", map[string]string{})
		}
	}
	if !data.AssemblerTimeout.IsNull() && !data.AssemblerTimeout.IsUnknown() {
		body, _ = sjson.Set(body, "assembler.timeout", strconv.FormatInt(data.AssemblerTimeout.ValueInt64(), 10))
	}
	if !data.AssemblerMaxPackets.IsNull() && !data.AssemblerMaxPackets.IsUnknown() {
		body, _ = sjson.Set(body, "assembler.max-packets", strconv.FormatInt(data.AssemblerMaxPackets.ValueInt64(), 10))
	}
	if !data.AssemblerReassemblerDropEnable.IsNull() && !data.AssemblerReassemblerDropEnable.IsUnknown() {
		if data.AssemblerReassemblerDropEnable.ValueBool() {
			body, _ = sjson.Set(body, "assembler.reassembler-drop.enable", map[string]string{})
		}
	}
	if !data.AssemblerFragHdrIncompleteEnable.IsNull() && !data.AssemblerFragHdrIncompleteEnable.IsUnknown() {
		if data.AssemblerFragHdrIncompleteEnable.ValueBool() {
			body, _ = sjson.Set(body, "assembler.frag-hdr-incomplete.enable", map[string]string{})
		}
	}
	if !data.AssemblerOverlapFragDropEnable.IsNull() && !data.AssemblerOverlapFragDropEnable.IsUnknown() {
		if data.AssemblerOverlapFragDropEnable.ValueBool() {
			body, _ = sjson.Set(body, "assembler.overlap-frag-drop.enable", map[string]string{})
		}
	}
	if !data.PathMtuEnable.IsNull() && !data.PathMtuEnable.IsUnknown() {
		if data.PathMtuEnable.ValueBool() {
			body, _ = sjson.Set(body, "path-mtu.enable", map[string]string{})
		}
	}
	if !data.PathMtuTimeout.IsNull() && !data.PathMtuTimeout.IsUnknown() {
		body, _ = sjson.Set(body, "path-mtu.timeout", strconv.FormatInt(data.PathMtuTimeout.ValueInt64(), 10))
	}
	return body
}

func (data *IPv6) updateFromBody(ctx context.Context, res []byte) {
	if value := gjson.GetBytes(res, "hop-limit"); value.Exists() && !data.HopLimit.IsNull() {
		data.HopLimit = types.Int64Value(value.Int())
	} else {
		data.HopLimit = types.Int64Null()
	}
	if value := gjson.GetBytes(res, "icmp.error-interval.interval-time"); value.Exists() && !data.IcmpErrorInterval.IsNull() {
		data.IcmpErrorInterval = types.Int64Value(value.Int())
	} else {
		data.IcmpErrorInterval = types.Int64Null()
	}
	if value := gjson.GetBytes(res, "icmp.error-interval.bucket-size"); value.Exists() && !data.IcmpErrorIntervalBucketSize.IsNull() {
		data.IcmpErrorIntervalBucketSize = types.Int64Value(value.Int())
	} else {
		data.IcmpErrorIntervalBucketSize = types.Int64Null()
	}
	if value := gjson.GetBytes(res, "source-route"); !data.SourceRoute.IsNull() {
		if value.Exists() {
			data.SourceRoute = types.BoolValue(true)
		} else {
			data.SourceRoute = types.BoolValue(false)
		}
	} else {
		data.SourceRoute = types.BoolNull()
	}
	if value := gjson.GetBytes(res, "assembler.timeout"); value.Exists() && !data.AssemblerTimeout.IsNull() {
		data.AssemblerTimeout = types.Int64Value(value.Int())
	} else {
		data.AssemblerTimeout = types.Int64Null()
	}
	if value := gjson.GetBytes(res, "assembler.max-packets"); value.Exists() && !data.AssemblerMaxPackets.IsNull() {
		data.AssemblerMaxPackets = types.Int64Value(value.Int())
	} else {
		data.AssemblerMaxPackets = types.Int64Null()
	}
	if value := gjson.GetBytes(res, "assembler.reassembler-drop.enable"); !data.AssemblerReassemblerDropEnable.IsNull() {
		if value.Exists() {
			data.AssemblerReassemblerDropEnable = types.BoolValue(true)
		} else {
			data.AssemblerReassemblerDropEnable = types.BoolValue(false)
		}
	} else {
		data.AssemblerReassemblerDropEnable = types.BoolNull()
	}
	if value := gjson.GetBytes(res, "assembler.frag-hdr-incomplete.enable"); !data.AssemblerFragHdrIncompleteEnable.IsNull() {
		if value.Exists() {
			data.AssemblerFragHdrIncompleteEnable = types.BoolValue(true)
		} else {
			data.AssemblerFragHdrIncompleteEnable = types.BoolValue(false)
		}
	} else {
		data.AssemblerFragHdrIncompleteEnable = types.BoolNull()
	}
	if value := gjson.GetBytes(res, "assembler.overlap-frag-drop.enable"); !data.AssemblerOverlapFragDropEnable.IsNull() {
		if value.Exists() {
			data.AssemblerOverlapFragDropEnable = types.BoolValue(true)
		} else {
			data.AssemblerOverlapFragDropEnable = types.BoolValue(false)
		}
	} else {
		data.AssemblerOverlapFragDropEnable = types.BoolNull()
	}
	if value := gjson.GetBytes(res, "path-mtu.enable"); !data.PathMtuEnable.IsNull() {
		if value.Exists() {
			data.PathMtuEnable = types.BoolValue(true)
		} else {
			data.PathMtuEnable = types.BoolValue(false)
		}
	} else {
		data.PathMtuEnable = types.BoolNull()
	}
	if value := gjson.GetBytes(res, "path-mtu.timeout"); value.Exists() && !data.PathMtuTimeout.IsNull() {
		data.PathMtuTimeout = types.Int64Value(value.Int())
	} else {
		data.PathMtuTimeout = types.Int64Null()
	}
}

func (data *IPv6Data) fromBody(ctx context.Context, res []byte) {
	if value := gjson.GetBytes(res, "hop-limit"); value.Exists() {
		data.HopLimit = types.Int64Value(value.Int())
	}
	if value := gjson.GetBytes(res, "icmp.error-interval.interval-time"); value.Exists() {
		data.IcmpErrorInterval = types.Int64Value(value.Int())
	}
	if value := gjson.GetBytes(res, "icmp.error-interval.bucket-size"); value.Exists() {
		data.IcmpErrorIntervalBucketSize = types.Int64Value(value.Int())
	}
	if value := gjson.GetBytes(res, "source-route"); value.Exists() {
		data.SourceRoute = types.BoolValue(true)
	} else {
		data.SourceRoute = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "assembler.timeout"); value.Exists() {
		data.AssemblerTimeout = types.Int64Value(value.Int())
	}
	if value := gjson.GetBytes(res, "assembler.max-packets"); value.Exists() {
		data.AssemblerMaxPackets = types.Int64Value(value.Int())
	}
	if value := gjson.GetBytes(res, "assembler.reassembler-drop.enable"); value.Exists() {
		data.AssemblerReassemblerDropEnable = types.BoolValue(true)
	} else {
		data.AssemblerReassemblerDropEnable = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "assembler.frag-hdr-incomplete.enable"); value.Exists() {
		data.AssemblerFragHdrIncompleteEnable = types.BoolValue(true)
	} else {
		data.AssemblerFragHdrIncompleteEnable = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "assembler.overlap-frag-drop.enable"); value.Exists() {
		data.AssemblerOverlapFragDropEnable = types.BoolValue(true)
	} else {
		data.AssemblerOverlapFragDropEnable = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "path-mtu.enable"); value.Exists() {
		data.PathMtuEnable = types.BoolValue(true)
	} else {
		data.PathMtuEnable = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "path-mtu.timeout"); value.Exists() {
		data.PathMtuTimeout = types.Int64Value(value.Int())
	}
}

func (data *IPv6) getDeletedItems(ctx context.Context, state IPv6) []string {
	deletedItems := make([]string, 0)
	if !state.HopLimit.IsNull() && data.HopLimit.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/hop-limit", state.getPath()))
	}
	if !state.IcmpErrorInterval.IsNull() && data.IcmpErrorInterval.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/icmp/error-interval", state.getPath()))
	}
	if !state.IcmpErrorIntervalBucketSize.IsNull() && data.IcmpErrorIntervalBucketSize.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/icmp/error-interval", state.getPath()))
	}
	if !state.SourceRoute.IsNull() && data.SourceRoute.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/source-route", state.getPath()))
	}
	if !state.AssemblerTimeout.IsNull() && data.AssemblerTimeout.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/assembler/timeout", state.getPath()))
	}
	if !state.AssemblerMaxPackets.IsNull() && data.AssemblerMaxPackets.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/assembler/max-packets", state.getPath()))
	}
	if !state.AssemblerReassemblerDropEnable.IsNull() && data.AssemblerReassemblerDropEnable.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/assembler/reassembler-drop/enable", state.getPath()))
	}
	if !state.AssemblerFragHdrIncompleteEnable.IsNull() && data.AssemblerFragHdrIncompleteEnable.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/assembler/frag-hdr-incomplete/enable", state.getPath()))
	}
	if !state.AssemblerOverlapFragDropEnable.IsNull() && data.AssemblerOverlapFragDropEnable.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/assembler/overlap-frag-drop/enable", state.getPath()))
	}
	if !state.PathMtuEnable.IsNull() && data.PathMtuEnable.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/path-mtu/enable", state.getPath()))
	}
	if !state.PathMtuTimeout.IsNull() && data.PathMtuTimeout.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/path-mtu/timeout", state.getPath()))
	}
	return deletedItems
}

func (data *IPv6) getEmptyLeafsDelete(ctx context.Context) []string {
	emptyLeafsDelete := make([]string, 0)
	if !data.SourceRoute.IsNull() && !data.SourceRoute.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/source-route", data.getPath()))
	}
	if !data.AssemblerReassemblerDropEnable.IsNull() && !data.AssemblerReassemblerDropEnable.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/assembler/reassembler-drop/enable", data.getPath()))
	}
	if !data.AssemblerFragHdrIncompleteEnable.IsNull() && !data.AssemblerFragHdrIncompleteEnable.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/assembler/frag-hdr-incomplete/enable", data.getPath()))
	}
	if !data.AssemblerOverlapFragDropEnable.IsNull() && !data.AssemblerOverlapFragDropEnable.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/assembler/overlap-frag-drop/enable", data.getPath()))
	}
	if !data.PathMtuEnable.IsNull() && !data.PathMtuEnable.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/path-mtu/enable", data.getPath()))
	}
	return emptyLeafsDelete
}

func (data *IPv6) getDeletePaths(ctx context.Context) []string {
	var deletePaths []string
	if !data.HopLimit.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/hop-limit", data.getPath()))
	}
	if !data.IcmpErrorInterval.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/icmp/error-interval", data.getPath()))
	}
	if !data.IcmpErrorIntervalBucketSize.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/icmp/error-interval", data.getPath()))
	}
	if !data.SourceRoute.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/source-route", data.getPath()))
	}
	if !data.AssemblerTimeout.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/assembler/timeout", data.getPath()))
	}
	if !data.AssemblerMaxPackets.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/assembler/max-packets", data.getPath()))
	}
	if !data.AssemblerReassemblerDropEnable.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/assembler/reassembler-drop/enable", data.getPath()))
	}
	if !data.AssemblerFragHdrIncompleteEnable.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/assembler/frag-hdr-incomplete/enable", data.getPath()))
	}
	if !data.AssemblerOverlapFragDropEnable.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/assembler/overlap-frag-drop/enable", data.getPath()))
	}
	if !data.PathMtuEnable.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/path-mtu/enable", data.getPath()))
	}
	if !data.PathMtuTimeout.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/path-mtu/timeout", data.getPath()))
	}
	return deletePaths
}
