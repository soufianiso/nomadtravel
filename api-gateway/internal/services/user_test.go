package services

import(
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandlerRegister(t *testing.T){
	req, err := http.NewRequest("POST","/register",nil) 
	req.Header.Set("Content-Type","application/json")
	if err != nil{
		t.Fatalf("can't create a request")
	}

	rr := httptest.NewRecorder()	


	hanlder := handleRegister(nil) 
	hanlder.ServeHTTP(rr,req)

	if rr.Code != 200{
		t.Errorf("he status code is wrong %v",rr.Code)
	}




}
