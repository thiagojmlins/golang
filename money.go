package main

type Dollar struct {
	amount int
}

func (d Dollar) Times (multiplier int) Dollar {
	return Dollar{amount: d.amount * multiplier}
}