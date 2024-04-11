package balancer

type HashSpace struct {
	partitionNumber uint
	nodes []node
}