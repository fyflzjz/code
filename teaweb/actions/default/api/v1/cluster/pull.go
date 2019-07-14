package cluster

import (
	"github.com/TeaWeb/code/teacluster"
	"github.com/TeaWeb/code/teaconfigs"
	"github.com/TeaWeb/code/teaweb/actions/default/api/apiutils"
	"github.com/iwind/TeaGo/actions"
)

type PullAction actions.Action

// 从cluster pull
func (this *PullAction) RunGet(params struct{}) {
	node := teaconfigs.SharedNodeConfig()
	if node == nil {
		apiutils.Fail(this, "the node has not been configured")
		return
	}

	if !teacluster.ClusterManager.IsActive() {
		apiutils.Fail(this, "the node is not connecting to cluster")
		return
	}

	if !node.IsMaster() {
		teacluster.ClusterManager.PullItems()
		teacluster.ClusterManager.SetIsChanged(false)
	}

	apiutils.SuccessOK(this)
}
