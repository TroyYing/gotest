package balance

type Balancer interface {
	DoBalance(insts []*Instance) (inst *Instance, err error)
}
