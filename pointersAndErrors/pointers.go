package pointersanderrors

// By doing this we're making a new type and we can declare methods on them.
// This can be very useful when you want to add some domain specific functionality on top of existing types.
type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
