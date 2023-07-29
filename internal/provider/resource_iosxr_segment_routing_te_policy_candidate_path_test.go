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

func TestAccIosxrSegmentRoutingTEPolicyCandidatePath(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_segment_routing_te_policy_candidate_path.test", "path_index", "100"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_segment_routing_te_policy_candidate_path.test", "path_infos.0.type", "dynamic"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_segment_routing_te_policy_candidate_path.test", "path_infos.0.pcep", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_segment_routing_te_policy_candidate_path.test", "path_infos.0.metric_type", "igp"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_segment_routing_te_policy_candidate_path.test", "path_infos.0.hop_type", "mpls"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_segment_routing_te_policy_candidate_path.test", "path_infos.0.segment_list_name", "dynamic"))
	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccIosxrSegmentRoutingTEPolicyCandidatePathConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccIosxrSegmentRoutingTEPolicyCandidatePathConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
	})
	steps = append(steps, resource.TestStep{
		ResourceName:  "iosxr_segment_routing_te_policy_candidate_path.test",
		ImportState:   true,
		ImportStateId: "Cisco-IOS-XR-segment-routing-ms-cfg:/sr/Cisco-IOS-XR-infra-xtc-agent-cfg:traffic-engineering/Cisco-IOS-XR-infra-xtc-agent-cfg:policies/Cisco-IOS-XR-infra-xtc-agent-cfg:policy[policy-name=POLICY1]/Cisco-IOS-XR-infra-xtc-agent-cfg:candidate-paths/Cisco-IOS-XR-infra-xtc-agent-cfg:preferences/Cisco-IOS-XR-infra-xtc-agent-cfg:preference[path-index=%!d(string=100)]",
	})
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps:                    steps,
	})
}

func testAccIosxrSegmentRoutingTEPolicyCandidatePathConfig_minimum() string {
	config := `resource "iosxr_segment_routing_te_policy_candidate_path" "test" {` + "\n"
	config += `	policy_name = "POLICY1"` + "\n"
	config += `	path_index = 100` + "\n"
	config += `}` + "\n"
	return config
}

func testAccIosxrSegmentRoutingTEPolicyCandidatePathConfig_all() string {
	config := `resource "iosxr_segment_routing_te_policy_candidate_path" "test" {` + "\n"
	config += `	policy_name = "POLICY1"` + "\n"
	config += `	path_index = 100` + "\n"
	config += `	path_infos = [{` + "\n"
	config += `		type = "dynamic"` + "\n"
	config += `		pcep = true` + "\n"
	config += `		metric_type = "igp"` + "\n"
	config += `		hop_type = "mpls"` + "\n"
	config += `		segment_list_name = "dynamic"` + "\n"
	config += `		}]` + "\n"
	config += `}` + "\n"
	return config
}
