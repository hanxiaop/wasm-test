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
	return &pluginContext{counter: proxywasm.DefineCounterMetric("proxy_wasm_go.connection_counter")}
}

type pluginContext struct {
	// Embed the default plugin context here,
	// so that we don't need to reimplement all the methods.
	types.DefaultPluginContext
	counter proxywasm.MetricCounter
}

func (ctx *pluginContext) NewTcpContext(contextID uint32) types.TcpContext {
	return &networkContext{counter: ctx.counter}
}

type networkContext struct {
	// Embed the default tcp context here,
	// so that we don't need to reimplement all the methods.
	types.DefaultTcpContext
	counter proxywasm.MetricCounter
}

func (ctx *networkContext) OnNewConnection() types.Action {
	proxywasm.LogInfo("new connection!")
	return types.ActionContinue
}

func (ctx *networkContext) OnStreamDone() {
	ctx.counter.Increment(1)
	proxywasm.LogInfo("connection complete!")
}

func (ctx *networkContext) OnDownstreamClose(types.PeerType) {
	proxywasm.LogInfo("downstream connection close!")
	return
}

func (ctx *networkContext) OnDownstreamData(dataSize int, endOfStream bool) types.Action {
	if dataSize == 0 {
		return types.ActionContinue
	}

	sourceClusterID, err := proxywasm.GetProperty([]string{"downstream_peer", "cluster_id"})
	if err != nil {
		proxywasm.LogCriticalf("failed to get downstream cluster id: %v", err)
		return types.ActionContinue
	}
	sourceWorkloadNamespace, err := proxywasm.GetProperty([]string{"downstream_peer", "namespace"})
	if err != nil {
		proxywasm.LogCriticalf("failed to get downstream workload namespace: %v", err)
		return types.ActionContinue
	}
	sourceWorkloadName, err := proxywasm.GetProperty([]string{"downstream_peer", "workload_name"})
	if err != nil {
		proxywasm.LogCriticalf("failed to get downstream workload name: %v", err)
		return types.ActionContinue
	}
	sourcePodName, err := proxywasm.GetProperty([]string{"downstream_peer", "name"})
	if err != nil {
		proxywasm.LogCriticalf("failed to get downstream pod name: %v", err)
		return types.ActionContinue
	}

	// this is the Envoy inbound to send request to
	clusterName, err := proxywasm.GetProperty([]string{"cluster_name"})
	if err != nil {
		proxywasm.LogCriticalf("failed to get cluster name: %v", err)
		return types.ActionContinue
	}
	//proxywasm.LogCriticalf("downstream cluster id: %v", sourceClusterID)
	//proxywasm.LogCriticalf("downstream workload namespace: %v", sourceWorkloadNamespace)
	//proxywasm.LogCriticalf("downstream workload name: %v", sourceWorkloadName)
	//proxywasm.LogCriticalf("downstream pod name: %v", sourcePodName)
	//proxywasm.LogCriticalf("downstream cluster name: %v", clusterName)

	workload := ""
	for k, v := range map[string]string{
		"cluster":   string(sourceClusterID),
		"namespace": string(sourceWorkloadNamespace),
		"name":      string(sourceWorkloadName),
		"pod":       string(sourcePodName),
	} {
		workload += k + "=" + v + ","
	}
	workload = workload[:len(workload)-1]
	//headers := [][2]string{
	//	{":method", "GET"},
	//	{"workload", workload},
	//	{":authority", string(destinationAddr)},
	//	{":path", "/"},
	//	{"accept", "*/*"},
	//}

	//proxywasm.LogInfof("headers: %+v", headers)
	proxywasm.LogCriticalf("workload: %s", workload)
	proxywasm.LogCriticalf("cluster: %s", string(clusterName))
	////_, err = proxywasm.DispatchHttpCall(string(clusterName), headers, nil, nil, 5000, func(numHeaders, bodySize, numTrailers int) {
	////	proxywasm.LogCriticalf("wasm callback...")
	////})
	//if err != nil {
	//	proxywasm.LogCriticalf("dispatch httpcall failed: %v", err)
	//	return types.ActionContinue
	//}
	return types.ActionContinue
}
