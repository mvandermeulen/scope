package report

import (
	"time"

	"github.com/weaveworks/common/mtime"
)

// Node describes a superset of the metadata that probes can collect
// about a given node in a given topology, along with the edges (aka
// adjacency) emanating from the node.
type Node struct {
	ID             string                   `json:"id,omitempty"`
	Topology       string                   `json:"topology,omitempty"`
	NodeTag        string                   `json:"nodeTag,omitempty"`
	Counters       Counters                 `json:"counters,omitempty"`
	Sets           Sets                     `json:"sets,omitempty"`
	Adjacency      IDList                   `json:"adjacency,omitempty"`
	LatestControls NodeControlDataLatestMap `json:"latestControls,omitempty"`
	Latest         StringLatestMap          `json:"latest,omitempty"`
	Metrics        Metrics                  `json:"metrics,omitempty" deepequal:"nil==empty"`
	Parents        Sets                     `json:"parents,omitempty"`
	Children       NodeSet                  `json:"children,omitempty"`
}

// MakeNode creates a new Node with no initial metadata.
func MakeNode(id string) Node {
	return Node{
		ID:             id,
		NodeTag:        "",
		Counters:       MakeCounters(),
		Sets:           MakeSets(),
		Adjacency:      MakeIDList(),
		LatestControls: MakeNodeControlDataLatestMap(),
		Latest:         MakeStringLatestMap(),
		Metrics:        Metrics{},
		Parents:        MakeSets(),
	}
}

// MakeNodeWith creates a new Node with the supplied map.
func MakeNodeWith(id string, m map[string]string) Node {
	return MakeNode(id).WithLatests(m)
}

// WithID returns a fresh copy of n, with ID changed.
func (n Node) WithID(id string) Node {
	n.ID = id
	return n
}

// WithTopology returns a fresh copy of n, with ID changed.
func (n Node) WithTopology(topology string) Node {
	n.Topology = topology
	return n
}

// Before is used for sorting nodes by topology and id
func (n Node) Before(other Node) bool {
	return n.Topology < other.Topology || (n.Topology == other.Topology && n.ID < other.ID)
}

// Equal is used for comparing nodes by topology and id
func (n Node) Equal(other Node) bool {
	return n.Topology == other.Topology && n.ID == other.ID
}

// After is used for sorting nodes by topology and id
func (n Node) After(other Node) bool {
	return other.Topology < n.Topology || (other.Topology == n.Topology && other.ID < n.ID)
}

// WithLatests returns a fresh copy of n, with Metadata m merged in.
func (n Node) WithLatests(m map[string]string) Node {
	ts := mtime.Now()
	n.Latest = n.Latest.addMapEntries(ts, m)
	return n
}

// WithLatest produces a new Node with k mapped to v in the Latest metadata.
func (n Node) WithLatest(k string, ts time.Time, v string) Node {
	n.Latest = n.Latest.Set(k, ts, v)
	return n
}

// WithCounters returns a fresh copy of n, with Counters c merged in.
func (n Node) WithCounters(c map[string]int) Node {
	n.Counters = n.Counters.Merge(Counters{}.fromIntermediate(c))
	return n
}

// WithSet returns a fresh copy of n, with set merged in at key.
func (n Node) WithSet(key string, set StringSet) Node {
	n.Sets = n.Sets.Add(key, set)
	return n
}

// WithSets returns a fresh copy of n, with sets merged in.
func (n Node) WithSets(sets Sets) Node {
	n.Sets = n.Sets.Merge(sets)
	return n
}

// WithMetric returns a fresh copy of n, with metric merged in at key.
func (n Node) WithMetric(key string, metric Metric) Node {
	n.Metrics = n.Metrics.Copy()
	n.Metrics[key] = n.Metrics[key].Merge(metric)
	return n
}

// WithMetrics returns a fresh copy of n, with metrics merged in.
func (n Node) WithMetrics(metrics Metrics) Node {
	n.Metrics = n.Metrics.Merge(metrics)
	return n
}

// WithAdjacent returns a fresh copy of n, with 'a' added to Adjacency
func (n Node) WithAdjacent(a ...string) Node {
	n.Adjacency = n.Adjacency.Add(a...)
	return n
}

// WithLatestActiveControls returns a fresh copy of n, with active controls cs added to LatestControls.
func (n Node) WithLatestActiveControls(cs ...string) Node {
	lcs := map[string]NodeControlData{}
	for _, control := range cs {
		lcs[control] = NodeControlData{}
	}
	return n.WithLatestControls(lcs)
}

// WithLatestControls returns a fresh copy of n, with lcs added to LatestControls.
func (n Node) WithLatestControls(lcs map[string]NodeControlData) Node {
	ts := mtime.Now()
	for k, v := range lcs {
		n.LatestControls = n.LatestControls.Set(k, ts, v)
	}
	return n
}

// WithLatestControl produces a new Node with control added to it
func (n Node) WithLatestControl(control string, ts time.Time, data NodeControlData) Node {
	n.LatestControls = n.LatestControls.Set(control, ts, data)
	return n
}

// WithParent returns a fresh copy of n, with one parent added
func (n Node) WithParent(key, parent string) Node {
	n.Parents = n.Parents.AddString(key, parent)
	return n
}

// WithParents returns a fresh copy of n, with sets merged in.
func (n Node) WithParents(parents Sets) Node {
	n.Parents = n.Parents.Merge(parents)
	return n
}

// PruneParents returns a fresh copy of n, without any parents.
func (n Node) PruneParents() Node {
	n.Parents = MakeSets()
	return n
}

// WithChildren returns a fresh copy of n, with children merged in.
func (n Node) WithChildren(children NodeSet) Node {
	n.Children = n.Children.Merge(children)
	return n
}

// WithChild returns a fresh copy of n, with one child merged in.
func (n Node) WithChild(child Node) Node {
	n.Children = n.Children.Merge(MakeNodeSet(child))
	return n
}

// WithNodeTag retrun a fresh copy of n, with node shape.
func (n Node) WithNodeTag(nodeTag string) Node {
	n.NodeTag = nodeTag
	return n
}

// Merge mergses the individual components of a node and returns a
// fresh node.
func (n Node) Merge(other Node) Node {
	id := n.ID
	if id == "" {
		id = other.ID
	}
	topology := n.Topology
	if topology == "" {
		topology = other.Topology
	} else if other.Topology != "" && topology != other.Topology {
		panic("Cannot merge nodes with different topology types: " + topology + " != " + other.Topology)
	}
	nodeTag := n.NodeTag
	if nodeTag == "" {
		nodeTag = other.NodeTag
	}
	return Node{
		ID:             id,
		NodeTag:        nodeTag,
		Topology:       topology,
		Counters:       n.Counters.Merge(other.Counters),
		Sets:           n.Sets.Merge(other.Sets),
		Adjacency:      n.Adjacency.Merge(other.Adjacency),
		LatestControls: n.LatestControls.Merge(other.LatestControls),
		Latest:         n.Latest.Merge(other.Latest),
		Metrics:        n.Metrics.Merge(other.Metrics),
		Parents:        n.Parents.Merge(other.Parents),
		Children:       n.Children.Merge(other.Children),
	}
}

// UnsafeUnMerge removes data from n that would be added by merging other,
// modifying the original.
// returns true if n.Merge(other) is the same as n
func (n *Node) UnsafeUnMerge(other Node) bool {
	// If it's not the same ID and topology then just bail out
	if n.ID != other.ID || n.Topology != other.Topology {
		return false
	}
	n.ID = ""
	n.Topology = ""
	remove := true
	// We either keep a whole section or drop it if anything changed
	//  - a trade-off of some extra data size in favour of faster simpler code.
	// (in practice, very few values reported by Scope probes do change over time)
	if n.LatestControls.EqualIgnoringTimestamps(other.LatestControls) {
		n.LatestControls = nil
	} else {
		remove = false
	}
	if n.Latest.EqualIgnoringTimestamps(other.Latest) {
		n.Latest = nil
	} else {
		remove = false
	}
	if n.Sets.DeepEqual(other.Sets) {
		n.Sets = MakeSets()
	} else {
		remove = false
	}
	if n.Parents.DeepEqual(other.Parents) {
		n.Parents = MakeSets()
	} else {
		remove = false
	}
	if n.Adjacency.Equal(other.Adjacency) {
		n.Adjacency = nil
	} else {
		remove = false
	}
	// counters and children are not created in the probe so we don't check those
	// metrics don't overlap so just check if we have any
	return remove && len(n.Metrics) == 0
}
