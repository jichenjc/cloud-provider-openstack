/*
Copyright 2017 The Kubernetes Authors.

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

package cinder

import (
	"github.com/gophercloud/gophercloud/openstack/blockstorage/v3/snapshots"
	"golang.org/x/net/context"
	"k8s.io/cloud-provider-openstack/pkg/csi/cinder/openstack"
)

var FakeCluster = "cluster"
var FakeNodeID = "CSINodeID"
var FakeEndpoint = "tcp://127.0.0.1:10000"
var FakeConfig = "/etc/cloud.conf"
var FakeCtx = context.Background()
var FakeVolName = "CSIVolumeName"
var FakeVolID = "CSIVolumeID"
var FakeSnapshotName = "CSISnapshotName"
var FakeSnapshotID = "261a8b81-3660-43e5-bab8-6470b65ee4e8"
var FakeCapacityGiB = 1
var FakeVolType = ""
var FakeAvailability = "testaz"
var FakeDevicePath = "/dev/xxx"
var FakeTargetPath = "/mnt/cinder"
var FakeStagingTargetPath = "/mnt/globalmount"
var FakeVol1 = openstack.Volume{
	ID:     "261a8b81-3660-43e5-bab8-6470b65ee4e9",
	Name:   "fake-duplicate",
	Status: "available",
	AZ:     "",
}
var FakeVol2 = openstack.Volume{
	ID:     "261a8b81-3660-43e5-bab8-6470b65ee4e9",
	Name:   "fake-duplicate",
	Status: "available",
	AZ:     "",
}
var FakeVol3 = openstack.Volume{
	ID:     "261a8b81-3660-43e5-bab8-6470b65ee4e9",
	Name:   "fake-3",
	Status: "available",
	AZ:     "",
}
var FakeSnapshotRes = snapshots.Snapshot{
	ID:       FakeSnapshotID,
	Name:     "fake-snapshot",
	VolumeID: FakeVolID,
}

var FakeSnapshotsRes = []snapshots.Snapshot{FakeSnapshotRes}
var FakeVolList = []openstack.Volume{FakeVol1, FakeVol3}
var FakeInstanceID = "321a8b81-3660-43e5-bab8-6470b65ee4e8"
