package main

import (
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm"
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm/types"
)

func main() {
	proxywasm.SetVMContext(&vmContext{})
}

type vmContext struct {
	types.DefaultVMContext
}

func (*vmContext) NewPluginContext(contextID uint32) types.PluginContext {
	return &pluginContext{}
}

type pluginContext struct {
	types.DefaultPluginContext
}

func (ctx *pluginContext) NewTcpContext(contextID uint32) types.TcpContext {
	return &networkContext{}
}

type networkContext struct {
	types.DefaultTcpContext
}

func (ctx *networkContext) OnDownstreamData(dataSize int, endOfStream bool) types.Action {
	addr, err := proxywasm.GetProperty([]string{"source", "address"})
	if err != nil {
		proxywasm.LogCriticalf("failed to get request data: %v", err)
		return types.ActionContinue
	}
	proxywasm.LogCriticalf("source host address: %s", string(addr))

	port, err := proxywasm.GetProperty([]string{"source", "port"})
	if err != nil {
		proxywasm.LogCriticalf("failed to get request data: %v", err)
		return types.ActionContinue
	}
	proxywasm.LogCriticalf("source host port: %s", string(port))

	daddr, err := proxywasm.GetProperty([]string{"destination", "address"})
	if err != nil {
		proxywasm.LogCriticalf("failed to get request data: %v", err)
		return types.ActionContinue
	}
	proxywasm.LogCriticalf("destination address: %s", string(daddr))

	dport, err := proxywasm.GetProperty([]string{"destination", "port"})
	if err != nil {
		proxywasm.LogCriticalf("failed to get request data: %v", err)
		return types.ActionContinue
	}
	proxywasm.LogCriticalf("destination port: %s", string(dport))

	ulocal, err := proxywasm.GetProperty([]string{"upstream", "local_address"})
	if err != nil {
		proxywasm.LogCriticalf("failed to get request data: %v", err)
		return types.ActionContinue
	}
	proxywasm.LogCriticalf("upstream local address: %s", string(ulocal))

	uaddr, err := proxywasm.GetProperty([]string{"upstream", "address"})
	if err != nil {
		proxywasm.LogCriticalf("failed to get request data: %v", err)
		return types.ActionContinue
	}
	proxywasm.LogCriticalf("upstream address: %s", string(uaddr))

	uport, err := proxywasm.GetProperty([]string{"upstream", "port"})
	if err != nil {
		proxywasm.LogCriticalf("failed to get request data: %v", err)
		return types.ActionContinue
	}
	proxywasm.LogCriticalf("upstream port: %s", string(uport))

	node, err := proxywasm.GetProperty([]string{"node", "id"})
	if err != nil {
		proxywasm.LogCriticalf("failed to get request data: %v", err)
		return types.ActionContinue
	}
	proxywasm.LogCriticalf("nnnoooodddde id: %s", string(node))

	nodeWorkloadName, err := proxywasm.GetProperty([]string{"node", "metadata", "WORKLOAD_NAME"})
	if err != nil {
		proxywasm.LogCriticalf("failed to get request data: %v", err)
		return types.ActionContinue
	}
	proxywasm.LogCriticalf("nnnoooodddde downstream WORKLOAD_NAME: %s", string(nodeWorkloadName))

	nodeWorkloadLabels, err := proxywasm.GetProperty([]string{"node", "metadata", "LABELS"})
	if err != nil {
		proxywasm.LogCriticalf("failed to get request data: %v", err)
		return types.ActionContinue
	}
	proxywasm.LogCriticalf("nnnoooodddde downstream WORKLOAD_NAME: %s", string(nodeWorkloadLabels))

	data, err := proxywasm.GetDownstreamData(0, dataSize)
	proxywasm.AppendDownstreamData([]byte{})
	if err != nil && err != types.ErrorStatusNotFound {
		proxywasm.LogCriticalf("failed to get downstream data: %v", err)
		return types.ActionContinue
	}
	proxywasm.LogCriticalf("downstream data, size %d, data: %s", dataSize, string(data))
	return types.ActionContinue
}

func (ctx *networkContext) OnUpstreamData(dataSize int, endOfStream bool) types.Action {
	ret, err := proxywasm.GetProperty([]string{"upstream", "address"})
	if err != nil {
		proxywasm.LogCriticalf("failed to get upstream data: %v", err)
		return types.ActionContinue
	}

	proxywasm.LogCriticalf("remote address: %s", string(ret))
	nodeWorkloadName, err := proxywasm.GetProperty([]string{"node", "metadata", "WORKLOAD_NAME"})
	if err != nil {
		proxywasm.LogCriticalf("failed to get request data: %v", err)
		return types.ActionContinue
	}
	proxywasm.LogCriticalf("nnnoooodddde upstream WORKLOAD_NAME: %s", string(nodeWorkloadName))

	nodeWorkloadLabels, err := proxywasm.GetProperty([]string{"node", "metadata", "LABELS"})
	if err != nil {
		proxywasm.LogCriticalf("failed to get request data: %v", err)
		return types.ActionContinue
	}
	proxywasm.LogCriticalf("nnnoooodddde upstream WORKLOAD_NAME: %s", string(nodeWorkloadLabels))

	data, err := proxywasm.GetUpstreamData(0, dataSize)
	if err != nil && err != types.ErrorStatusNotFound {
		proxywasm.LogCritical(err.Error())
	}
	proxywasm.LogCriticalf("upstream data, size %d, data: %s", dataSize, string(data))
	return types.ActionContinue
}
