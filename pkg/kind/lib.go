package kind

import (
	"github.com/magefile/mage/sh"
	"github.com/reaperes/tok/pkg/util"
)

func CreateCluster() {
	err := sh.Run(
		"kind", "create", "cluster",
		"--image", "kindest/node:v1.18.15",
		"--wait", "600s",
	)
	util.PanicOnError(err)
}

func DeleteCluster() {
	_ = sh.Run("kind", "delete", "cluster")
}

func GetKubeconfig() {
	sh.Run("kind", "")
}
