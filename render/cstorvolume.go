package render

import (
	"context"

	"github.com/weaveworks/scope/probe/kubernetes"
	"github.com/weaveworks/scope/report"
)

//CStorVolumeRenderer is a Renderer which produces a renderable openebs CV.
var CStorVolumeRenderer = cStorVolumeRenderer{}

//cStorVolumeRenderer is a Renderer to render CStor Volumes.
type cStorVolumeRenderer struct{}

//Render renders the CV.
func (v cStorVolumeRenderer) Render(ctx context.Context, rpt report.Report) Nodes {
	cStorNodes := make(report.Nodes)

	for cvID, cvNode := range rpt.CStorVolume.Nodes {
		cStorNodes[cvID] = cvNode
	}

	for cspID, cspNode := range rpt.CStorPool.Nodes {
		cStorNodes[cspID] = cspNode
	}

	for cspiID, cspiNode := range rpt.CStorPoolInstance.Nodes {
		cStorNodes[cspiID] = cspiNode
	}

	for cvrID, cvrNode := range rpt.CStorVolumeReplica.Nodes {
		cStorVolume, _ := cvrNode.Latest.Lookup(kubernetes.CStorVolumeName)
		cStorVolumeNodeID := report.MakeCStorVolumeNodeID(cStorVolume)

		if cvNode, ok := cStorNodes[cStorVolumeNodeID]; ok {
			cvNode.Adjacency = cvNode.Adjacency.Add(cvrID)
			cvNode.Children = cvNode.Children.Add(cvrNode)
			cStorNodes[cStorVolumeNodeID] = cvNode
		}

		cStorPoolUID, ok := cvrNode.Latest.Lookup(kubernetes.CStorPoolUID)
		if ok {
			cStorPoolNodeID := report.MakeCStorPoolNodeID(cStorPoolUID)
			if cStorPoolNode, ok := cStorNodes[cStorPoolNodeID]; ok {
				cvrNode.Children = cvrNode.Children.Add(cStorPoolNode)
			}
		}

		cStorPoolInstanceUID, ok := cvrNode.Latest.Lookup(kubernetes.CStorPoolInstanceUID)
		if ok {
			cStorPoolInstanceNodeID := report.MakeCStorPoolInstanceNodeID(cStorPoolInstanceUID)
			if cStorPoolInstanceNode, ok := cStorNodes[cStorPoolInstanceNodeID]; ok {
				cvrNode.Children = cvrNode.Children.Add(cStorPoolInstanceNode)
			}
		}

		cStorNodes[cvrID] = cvrNode
	}
	return Nodes{Nodes: cStorNodes}
}
