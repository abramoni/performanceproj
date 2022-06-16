package main

import (

	"presentation/mypackages"
)


func main(){
	data := mypackages.GenerateJson()
	mypackages.StoreInElasticDb(data,"longivitycluster")
}