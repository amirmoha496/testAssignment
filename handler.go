package main

import (
	"fmt"
	"net/http"
	"strconv"
)

var sm ServiceManager = ServiceManager{}

//HandlerManager class provides methods to handle Incoming requests for different Endpoints
type HandlerManager struct {
}

//EncodePasswordHandler Handler to handle request for POST /hash
func (hm HandlerManager) EncodePasswordHandler(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		res.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(res, "Method not allowed")
		return
	}

	req.ParseForm()
	pwd := req.Form.Get("password")
	if pwd == "" {
		res.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(res, "Parameter \"password\" is missing")
		return
	}

	k := sm.HashPassword(pwd)

	fmt.Fprintf(res, strconv.FormatInt(k, 10))
	return
}

//GetHashHandler Handler to handle request for GET /hash/{id}
func (hm HandlerManager) GetHashHandler(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		res.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(res, "Method not allowed")
		return
	}

	sha2, err := sm.GetHashForID(1)

	if err != nil {
		res.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(res, "Requested Hash Not Found")
		return
	}

	fmt.Fprintf(res, sha2)
	return
}

//StatsHandler Handler to handle request for GET /stats
func (hm HandlerManager) StatsHandler(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		res.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(res, "Method not allowed")
		return
	}

	out, err := sm.GetStatistics()
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(res, "Failed to retrive Stats:"+err.Error())
		return
	}

	fmt.Fprintf(res, out)
	return
}
