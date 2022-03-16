// Code generated by bpf2go; DO NOT EDIT.
//go:build 386 || amd64 || amd64p32 || arm || arm64 || mips64le || mips64p32le || mipsle || ppc64le || riscv64
// +build 386 amd64 amd64p32 arm arm64 mips64le mips64p32le mipsle ppc64le riscv64

package bpf2go

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"

	"github.com/cilium/ebpf"
)

// LoadXdpCluster returns the embedded CollectionSpec for XdpCluster.
func LoadXdpCluster() (*ebpf.CollectionSpec, error) {
	reader := bytes.NewReader(_XdpClusterBytes)
	spec, err := ebpf.LoadCollectionSpecFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("can't load XdpCluster: %w", err)
	}

	return spec, err
}

// LoadXdpClusterObjects loads XdpCluster and converts it into a struct.
//
// The following types are suitable as obj argument:
//
//     *XdpClusterObjects
//     *XdpClusterPrograms
//     *XdpClusterMaps
//
// See ebpf.CollectionSpec.LoadAndAssign documentation for details.
func LoadXdpClusterObjects(obj interface{}, opts *ebpf.CollectionOptions) error {
	spec, err := LoadXdpCluster()
	if err != nil {
		return err
	}

	return spec.LoadAndAssign(obj, opts)
}

// XdpClusterSpecs contains maps and programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type XdpClusterSpecs struct {
	XdpClusterProgramSpecs
	XdpClusterMapSpecs
}

// XdpClusterSpecs contains programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type XdpClusterProgramSpecs struct {
	ClusterManager *ebpf.ProgramSpec `ebpf:"cluster_manager"`
}

// XdpClusterMapSpecs contains maps before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type XdpClusterMapSpecs struct {
	Cluster      *ebpf.MapSpec `ebpf:"cluster"`
	Endpoint     *ebpf.MapSpec `ebpf:"endpoint"`
	Filter       *ebpf.MapSpec `ebpf:"filter"`
	FilterChain  *ebpf.MapSpec `ebpf:"filter_chain"`
	Listener     *ebpf.MapSpec `ebpf:"listener"`
	Loadbalance  *ebpf.MapSpec `ebpf:"loadbalance"`
	TailCallCtx  *ebpf.MapSpec `ebpf:"tail_call_ctx"`
	TailCallProg *ebpf.MapSpec `ebpf:"tail_call_prog"`
	TupleCt      *ebpf.MapSpec `ebpf:"tuple_ct"`
}

// XdpClusterObjects contains all objects after they have been loaded into the kernel.
//
// It can be passed to LoadXdpClusterObjects or ebpf.CollectionSpec.LoadAndAssign.
type XdpClusterObjects struct {
	XdpClusterPrograms
	XdpClusterMaps
}

func (o *XdpClusterObjects) Close() error {
	return _XdpClusterClose(
		&o.XdpClusterPrograms,
		&o.XdpClusterMaps,
	)
}

// XdpClusterMaps contains all maps after they have been loaded into the kernel.
//
// It can be passed to LoadXdpClusterObjects or ebpf.CollectionSpec.LoadAndAssign.
type XdpClusterMaps struct {
	Cluster      *ebpf.Map `ebpf:"cluster"`
	Endpoint     *ebpf.Map `ebpf:"endpoint"`
	Filter       *ebpf.Map `ebpf:"filter"`
	FilterChain  *ebpf.Map `ebpf:"filter_chain"`
	Listener     *ebpf.Map `ebpf:"listener"`
	Loadbalance  *ebpf.Map `ebpf:"loadbalance"`
	TailCallCtx  *ebpf.Map `ebpf:"tail_call_ctx"`
	TailCallProg *ebpf.Map `ebpf:"tail_call_prog"`
	TupleCt      *ebpf.Map `ebpf:"tuple_ct"`
}

func (m *XdpClusterMaps) Close() error {
	return _XdpClusterClose(
		m.Cluster,
		m.Endpoint,
		m.Filter,
		m.FilterChain,
		m.Listener,
		m.Loadbalance,
		m.TailCallCtx,
		m.TailCallProg,
		m.TupleCt,
	)
}

// XdpClusterPrograms contains all programs after they have been loaded into the kernel.
//
// It can be passed to LoadXdpClusterObjects or ebpf.CollectionSpec.LoadAndAssign.
type XdpClusterPrograms struct {
	ClusterManager *ebpf.Program `ebpf:"cluster_manager"`
}

func (p *XdpClusterPrograms) Close() error {
	return _XdpClusterClose(
		p.ClusterManager,
	)
}

func _XdpClusterClose(closers ...io.Closer) error {
	for _, closer := range closers {
		if err := closer.Close(); err != nil {
			return err
		}
	}
	return nil
}

// Do not access this directly.
//go:embed xdpcluster_bpfel.o
var _XdpClusterBytes []byte
