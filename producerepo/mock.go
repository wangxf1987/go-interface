package producerepo

import "fmt"

type mockProduceRepo struct {
}

func (m mockProduceRepo) StoreProduct(name string, id int) {
	fmt.Println("mocking the StoreProduct func")
}

func (m mockProduceRepo) FindProductByID(id int) {
	fmt.Println("mocking the FindProduceByID func")
}
