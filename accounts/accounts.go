package accounts

import (
	"errors"
	"fmt"
)

//Account is bank struct
type Account struct {
	owner   string
	balance int
}

//에러에 대한 변수를 정의하면 err로 시작하는 변수명으로 해야한다.
var errNoMoney = errors.New("Can't withdraw")

//NewAccount creates Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

//Deposit x amount on your account
func (a *Account) Deposit(amount int) {
	//receiver로 Account의 method로 작동한다.
	//receiver로 사용되는 변수는 소문자로 시작해야한다.
	//a Account는 이 함수를 호출한 오브젝트의 복사본을 가져와 담는다,
	//여기서는 복사본이 아닌 호출한 오브젝트 자체의 값을 변경해야 하므로 Account가 아닌 *Account가 와야된다.
	//정리를 하자면 set 메서드는 원래 오브젝트를 가져와서 수정해야하기 때문에 *Account 이며
	// get 메서드는 원래 오브젝트의 복사본을 가져와도 되므로 Account가 와야한다.
	a.balance += amount
}

//Balance of your account
func (a Account) Balance() int {
	return a.balance
}

//Withdraw x amount from your account
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errNoMoney
	}
	a.balance -= amount

	//method에서 문제가 있으면 error를 리턴하기 때문에 method 리턴값으로 error를 정의했다.
	//문제가 없으면 nil를 리턴하면 된다.
	return nil
}

//ChangeOwner of the account
func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

//Owner of the account
func (a Account) Owner() string {
	return a.owner
}

//String of the account
func (a Account) String() string {
	return fmt.Sprint(a.Owner(), "'s account. \nHas: ", a.Balance())
}
