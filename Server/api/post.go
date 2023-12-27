package api

import (
	"encoding/json"
	"femm/data"
	"net/http"
)


func Post(w http.ResponseWriter, r *http.Request){
	if r.Method == http.MethodPost{
		var requestItem data.Exhibition
		err := json.NewDecoder(r.Body).Decode(&requestItem)
		if err!= nil{
			http.Error(w, err.Error, http.StatusBadRequest)
			return
		}
		data.Add(exhibiton)
		w.Write([]byte("OK"))

	}else{
		http.Error(w, "Unsupported Method", http.StatusMethodNotAllowed)
	}
}
