package api

import (
	"encoding/json"
	"femm/data"
	"fmt"
	"net/http"
	"strconv"
)

func Get(w http.ResponseWriter, r*http.Request){
	w.Header().Set("Content-Type", "application/json")
	id:=r.URL.Query()["id"]
	if id!=nil{
		_id, err:= strconv.Atoi(id[0])
		if err == nil && _id <len(data.GetAll()){
			item := data.GetAll()[_id]
			fmt.Println(item.Title,"_______________-",item.Image )
			json.NewEncoder(w).Encode(item)//Encoder is erroring for some reason need to look in detail why later!
		}else{
			http.Error(w, "Invalud Exhibition", http.StatusBadRequest)
		}
	}
	json.NewEncoder(w).Encode(data.GetAll())

}
