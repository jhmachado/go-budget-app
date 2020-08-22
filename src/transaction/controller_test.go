package transaction

import (
    "net/http"
    "net/http/httptest"
	"testing"
	"reflect"
	"strings"
)

func TestGetTransactionsRouteShouldReturnASuccessfullStatusCodeWithAJsonArray(t *testing.T) {
	req, err := http.NewRequest("GET", "/transactions", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetTransactions)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("The route returned a wrong status code: got %v want %v", status, http.StatusOK)
	}

	if ctype := rr.Header().Get("Content-Type"); ctype != "application/json" {
        t.Errorf("The Content-Type header does not match: got %v want %v",
            ctype, "application/json")
	}
	
	expectedBody := `[{"title":"Salário","amount":1220,"type":0,"created_at":"2020-08-22T13:43:00Z"}]`
	if reflect.DeepEqual(rr.Body.String(), expectedBody) {
        t.Errorf("The route returned an unexpected body: got %v want %v",
            rr.Body.String(), expectedBody)
    }
}

func TestGetTransactionsRouteWithTheWrongMethodShouldReturnAnErrorResponse(t *testing.T) {
	req, err := http.NewRequest("POST", "/transactions", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetTransactions)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusMethodNotAllowed {
		t.Errorf("The route returned a wrong status code: got %v want %v", status, http.StatusMethodNotAllowed)
	}

	expectedBody := ""
	if rr.Body.String() != expectedBody {
        t.Error("The method not allowed error should have an empty body")
    }
}

func TestCreateTransactionShouldReturnACreatedHttpStatusCode(t *testing.T) {
	requestBody := strings.NewReader(`{"title": "Salário","amount": 1220,"type": 0,"created_at": "2020-08-22T13:43:00Z"}`)
	req, err := http.NewRequest("POST", "/transactions/create", requestBody)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateTransaction)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("The route returned a wrong status code: got %v want %v", status, http.StatusCreated)
	}

	expectedBody := `{"response": "A new transaction was created"}`
	if rr.Body.String() != expectedBody {
        t.Errorf("The route returned an unexpected body: got %v want %v",
            rr.Body.String(), expectedBody)
    }
}

func TestCreateTransactionWithTheWrongMethodShouldReturnAnErrorResponse(t *testing.T) {
	req, err := http.NewRequest("PUT", "/transactions/create", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateTransaction)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusMethodNotAllowed {
		t.Errorf("The route returned a wrong status code: got %v want %v", status, http.StatusMethodNotAllowed)
	}

	expectedBody := ""
	if rr.Body.String() != expectedBody {
        t.Error("The method not allowed error should have an empty body")
    }
}
