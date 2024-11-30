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
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIosxrBFD(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_bfd.test", "echo_disable", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_bfd.test", "echo_latency_detect_percentage", "200"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_bfd.test", "echo_latency_detect_count", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_bfd.test", "echo_startup_validate_force", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_bfd.test", "echo_ipv4_source", "10.1.1.1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_bfd.test", "echo_ipv4_bundle_per_member_preferred_minimum_interval", "200"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_bfd.test", "trap_singlehop_pre_mapped", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_bfd.test", "multipath_locations.0.location_name", "0/0/CPU0"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_bfd.test", "multihop_ttl_drop_threshold", "200"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_bfd.test", "dampening_initial_wait", "3600"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_bfd.test", "dampening_secondary_wait", "3200"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_bfd.test", "dampening_maximum_wait", "3100"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_bfd.test", "dampening_threshold", "60000"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_bfd.test", "dampening_extensions_down_monitoring", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_bfd.test", "dampening_disable", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_bfd.test", "dampening_bundle_member_l3_only_mode", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_bfd.test", "dampening_bundle_member_initial_wait", "5184"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_bfd.test", "dampening_bundle_member_secondary_wait", "6184"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_bfd.test", "dampening_bundle_member_maximum_wait", "7184"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_bfd.test", "bundle_coexistence_bob_blb_inherit", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_bfd.test", "bundle_coexistence_bob_blb_logical", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_bfd.test", "interfaces.0.interface_name", "GigabitEthernet0/0/0/0"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_bfd.test", "interfaces.0.echo_disable", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_bfd.test", "interfaces.0.echo_ipv4_source", "12.1.1.1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_bfd.test", "interfaces.0.ipv6_checksum_disable", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_bfd.test", "interfaces.0.disable", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_bfd.test", "interfaces.0.local_address", "33.33.31.1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_bfd.test", "interfaces.0.tx_interval", "3200"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_bfd.test", "interfaces.0.rx_interval", "4200"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_bfd.test", "interfaces.0.multiplier", "40"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_bfd.test", "ipv6_checksum_disable", "true"))
	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccIosxrBFDConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccIosxrBFDConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
	})
	steps = append(steps, resource.TestStep{
		ResourceName:  "iosxr_bfd.test",
		ImportState:   true,
		ImportStateId: "",
		Check:         resource.ComposeTestCheckFunc(checks...),
	})
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps:                    steps,
	})
}

func testAccIosxrBFDConfig_minimum() string {
	config := `resource "iosxr_bfd" "test" {` + "\n"
	config += `}` + "\n"
	return config
}

func testAccIosxrBFDConfig_all() string {
	config := `resource "iosxr_bfd" "test" {` + "\n"
	config += `	echo_disable = true` + "\n"
	config += `	echo_latency_detect_percentage = 200` + "\n"
	config += `	echo_latency_detect_count = 10` + "\n"
	config += `	echo_startup_validate_force = true` + "\n"
	config += `	echo_ipv4_source = "10.1.1.1"` + "\n"
	config += `	echo_ipv4_bundle_per_member_preferred_minimum_interval = 200` + "\n"
	config += `	trap_singlehop_pre_mapped = true` + "\n"
	config += `	multipath_locations = [{` + "\n"
	config += `		location_name = "0/0/CPU0"` + "\n"
	config += `		}]` + "\n"
	config += `	multihop_ttl_drop_threshold = 200` + "\n"
	config += `	dampening_initial_wait = 3600` + "\n"
	config += `	dampening_secondary_wait = 3200` + "\n"
	config += `	dampening_maximum_wait = 3100` + "\n"
	config += `	dampening_threshold = 60000` + "\n"
	config += `	dampening_extensions_down_monitoring = true` + "\n"
	config += `	dampening_disable = true` + "\n"
	config += `	dampening_bundle_member_l3_only_mode = true` + "\n"
	config += `	dampening_bundle_member_initial_wait = 5184` + "\n"
	config += `	dampening_bundle_member_secondary_wait = 6184` + "\n"
	config += `	dampening_bundle_member_maximum_wait = 7184` + "\n"
	config += `	bundle_coexistence_bob_blb_inherit = false` + "\n"
	config += `	bundle_coexistence_bob_blb_logical = true` + "\n"
	config += `	interfaces = [{` + "\n"
	config += `		interface_name = "GigabitEthernet0/0/0/0"` + "\n"
	config += `		echo_disable = true` + "\n"
	config += `		echo_ipv4_source = "12.1.1.1"` + "\n"
	config += `		ipv6_checksum_disable = true` + "\n"
	config += `		disable = true` + "\n"
	config += `		local_address = "33.33.31.1"` + "\n"
	config += `		tx_interval = 3200` + "\n"
	config += `		rx_interval = 4200` + "\n"
	config += `		multiplier = 40` + "\n"
	config += `		}]` + "\n"
	config += `	ipv6_checksum_disable = true` + "\n"
	config += `}` + "\n"
	return config
}