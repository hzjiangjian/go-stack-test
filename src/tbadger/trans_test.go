package tbadger

import "testing"

func TestBadger(t *testing.T){
	ReadWriteTx()
	ReadOnlyTx()
	DeleteTx()
	ReadWriteTx()
	ReadOnlyTx()
}
