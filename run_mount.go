/*
   Copyright The containerd Authors.

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package main

import (
	"github.com/containerd/containerd/oci"
	"github.com/containerd/nerdctl/pkg/mountutil"
	"github.com/containerd/nerdctl/pkg/strutil"
	"github.com/opencontainers/runtime-spec/specs-go"
	"github.com/urfave/cli/v2"
)

func generateMountOpts(clicontext *cli.Context) ([]oci.SpecOpts, error) {
	var opts []oci.SpecOpts

	if flagVSlice := strutil.DedupeStrSlice(clicontext.StringSlice("v")); len(flagVSlice) > 0 {
		volStore, err := getVolumeStore(clicontext)
		if err != nil {
			return nil, err
		}
		ociMounts := make([]specs.Mount, len(flagVSlice))
		for i, v := range flagVSlice {
			m, err := mountutil.ParseFlagV(v, volStore)
			if err != nil {
				return nil, err
			}
			ociMounts[i] = *m
		}
		opts = append(opts, oci.WithMounts(ociMounts))
	}

	return opts, nil
}
