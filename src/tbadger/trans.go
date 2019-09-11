package tbadger

import (
	"fmt"
	"github.com/dgraph-io/badger"
	"log"
)

func GetDB() (db *badger.DB) {
	var err error

	db, err = badger.Open(badger.DefaultOptions("../../data/badger"))
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func getData() (kvlist map[string]string) {
	kvlist = make(map[string]string)

	for i := 0; i < 100; i++ {
		kvlist[fmt.Sprintf("key%d", i)] = fmt.Sprintf("value%d", i)
	}

	return
}

func ReadWriteTx() {
	db := GetDB()
	defer db.Close()

	updates := getData()

	err := db.Update(func(txn *badger.Txn) error {
		for k, v := range updates {
			if err := txn.Set([]byte(k), []byte(v)); err!=nil {
				log.Fatal("ReadWriteTx txnset error")
			}
		}
		return nil
	})

	if err != nil{
		log.Fatal("ReadWriteTx end error")
	}
}

func DeleteTx() {
	db := GetDB()
	defer db.Close()

	updates := getData()

	err := db.Update(func(txn *badger.Txn) error {
		for k, _ := range updates {
			if err := txn.Delete([]byte(k)); err!=nil {
				log.Fatal("DeleteTx txndel error")
			}
		}
		return nil
	})

	if err != nil{
		log.Fatal("DeleteTx end error")
	}
}

func ReadOnlyTx() {
	db := GetDB()
	defer db.Close()

	updates := getData()

	err := db.View(func(txn *badger.Txn) error {
		for k, _ := range updates {
			if item,err := txn.Get([]byte(k)); err!=nil {
				log.Fatal("ReadOnlyTx txnget error")
			}else{
				vbyte,_:=item.ValueCopy(nil)

				log.Println(string(vbyte))
			}
		}
		return nil
	})

	if err != nil{
		log.Fatal("ReadOnlyTx end error")
	}
}
