package main

import "go-interface/producerepo"

func main() {
	env := "aliyun"
	repo := producerepo.New(env)
	repo.StoreProduct("kong", 105)
	repo.FindProductByID(105)
}
