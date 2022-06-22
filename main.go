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

type workload struct {
	cluster      string
	namespace    string
	workloadName string
	podName      string
}

func (ctx *networkContext) OnDownstreamData(dataSize int, endOfStream bool) types.Action {
	//addr, err := proxywasm.GetProperty([]string{"source", "address"})
	//if err != nil {
	//	proxywasm.LogCriticalf("failed to get request data: %v", err)
	//	return types.ActionContinue
	//}
	//proxywasm.LogCriticalf("source host address: %s", string(addr))
	//
	//port, err := proxywasm.GetProperty([]string{"source", "port"})
	//if err != nil {
	//	proxywasm.LogCriticalf("failed to get request data: %v", err)
	//	return types.ActionContinue
	//}
	//proxywasm.LogCriticalf("source host port: %s", string(port))
	//
	//daddr, err := proxywasm.GetProperty([]string{"destination", "address"})
	//if err != nil {
	//	proxywasm.LogCriticalf("failed to get request data: %v", err)
	//	return types.ActionContinue
	//}
	//proxywasm.LogCriticalf("destination address: %s", string(daddr))
	//
	//dport, err := proxywasm.GetProperty([]string{"destination", "port"})
	//if err != nil {
	//	proxywasm.LogCriticalf("failed to get request data: %v", err)
	//	return types.ActionContinue
	//}
	//proxywasm.LogCriticalf("destination port: %s", string(dport))
	//
	//ulocal, err := proxywasm.GetProperty([]string{"upstream", "local_address"})
	//if err != nil {
	//	proxywasm.LogCriticalf("failed to get request data: %v", err)
	//	return types.ActionContinue
	//}
	//proxywasm.LogCriticalf("upstream local address: %s", string(ulocal))
	//
	//uaddr, err := proxywasm.GetProperty([]string{"upstream", "address"})
	//if err != nil {
	//	proxywasm.LogCriticalf("failed to get request data: %v", err)
	//	return types.ActionContinue
	//}
	//proxywasm.LogCriticalf("upstream address: %s", string(uaddr))
	//
	//uport, err := proxywasm.GetProperty([]string{"upstream", "port"})
	//if err != nil {
	//	proxywasm.LogCriticalf("failed to get request data: %v", err)
	//	return types.ActionContinue
	//}
	//proxywasm.LogCriticalf("upstream port: %s", string(uport))
	//
	//node, err := proxywasm.GetProperty([]string{"node", "id"})
	//if err != nil {
	//	proxywasm.LogCriticalf("failed to get request data: %v", err)
	//	return types.ActionContinue
	//}
	//proxywasm.LogCriticalf("nnnoooodddde id: %s", string(node))
	//
	//nodeWorkloadName, err := proxywasm.GetProperty([]string{"node", "metadata", "WORKLOAD_NAME"})
	//if err != nil {
	//	proxywasm.LogCriticalf("failed to get request data: %v", err)
	//	return types.ActionContinue
	//}
	//proxywasm.LogCriticalf("nnnoooodddde downstream WORKLOAD_NAME: %s", string(nodeWorkloadName))
	//
	//nodeName, err := proxywasm.GetProperty([]string{"node", "metadata", "NAME"})
	//if err != nil {
	//	proxywasm.LogCriticalf("failed to get request data: %v", err)
	//	return types.ActionContinue
	//}
	//proxywasm.LogCriticalf("nnnoooodddde downstream node NAME: %s", string(nodeName))
	//
	//nodeNamespace, err := proxywasm.GetProperty([]string{"node", "metadata", "NAMESPACE"})
	//if err != nil {
	//	proxywasm.LogCriticalf("failed to get request data: %v", err)
	//	return types.ActionContinue
	//}
	//proxywasm.LogCriticalf("nnnoooodddde downstream node NAMESPACE: %s", string(nodeNamespace))
	//
	//nodeowner, err := proxywasm.GetProperty([]string{"node", "metadata", "OWNER"})
	//if err != nil {
	//	proxywasm.LogCriticalf("failed to get request data: %v", err)
	//	return types.ActionContinue
	//}
	//proxywasm.LogCriticalf("nnnoooodddde downstream node OWNER: %s", string(nodeowner))
	//
	//nodeClusterID, err := proxywasm.GetProperty([]string{"node", "metadata", "CLUSTER_ID"})
	//if err != nil {
	//	proxywasm.LogCriticalf("failed to get request data: %v", err)
	//	return types.ActionContinue
	//}
	//proxywasm.LogCriticalf("nnnoooodddde downstream node ClusterID: %s", string(nodeClusterID))
	//
	//nodeWorkloadLabels, err := proxywasm.GetProperty([]string{"node", "metadata", "LABELS"})
	//if err != nil {
	//	proxywasm.LogCriticalf("failed to get request data: %v", err)
	//	return types.ActionContinue
	//}
	//proxywasm.LogCriticalf("nnnoooodddde downstream WORKLOAD_NAME: %s", string(nodeWorkloadLabels))

	// cluster_id
	// workload_namespace
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

	// this is inbound xxxxx
	//clusterName, err := proxywasm.GetProperty([]string{"cluster_name"})
	//if err != nil {
	//	proxywasm.LogCriticalf("failed to get cluster name: %v", err)
	//	return types.ActionContinue
	//}
	clusterName, err := proxywasm.GetProperty([]string{"node", "cluster"})
	if err != nil {
		proxywasm.LogCriticalf("failed to get cluster name: %v", err)
		return types.ActionContinue
	}
	sourceWorkload := workload{
		cluster:      string(sourceClusterID),
		namespace:    string(sourceWorkloadNamespace),
		workloadName: string(sourceWorkloadName),
		podName:      string(sourcePodName),
	}
	headers := [][2]string{
		generateHeader("cluster", sourceWorkload.cluster),
		generateHeader("workload_namespace", sourceWorkload.workloadName),
		generateHeader("workload_name", sourceWorkload.workloadName),
		generateHeader("pod_name", sourceWorkload.podName),
	}
	proxywasm.LogCriticalf("headers: %v", headers)
	proxywasm.LogCriticalf("cluster: %s", string(clusterName))
	_, err = proxywasm.DispatchHttpCall(string(clusterName), headers, nil, nil, 5000, callBack)
	if err != nil {
		proxywasm.LogCriticalf("dispatch httpcall failed: %v", err)
	}
	return types.ActionContinue
}

func callBack(numHeaders, bodySize, numTrailers int) {
	proxywasm.LogCriticalf("wasm callback...")
}

func generateHeader(key, value string) [2]string {
	return [2]string{key, value}
}
