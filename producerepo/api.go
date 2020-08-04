package producerepo

type ProduceRepository interface {
	StoreProduct(name string, id int)
	FindProductByID(id int)
}

func New(env string) ProduceRepository {
	switch env {
	case "aliyun":
		return aliCloudProductRepo{}
	case "local-mysql":
		return mysqlProductRepo{}
	}

	return mockProduceRepo{}
}
