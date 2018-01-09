package controller

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"io/ioutil"
	"os"
	"bytes"
	"encoding/json"
	"strconv"
	"time"

	"github.com/adrenalinchik/gowebapp/test"
	"github.com/adrenalinchik/gowebapp/model"
)

// TestMain is responsible for preconditions and postconditions in test suit.
// Before tests are started it creates logger, database and save test data to db.
// After all test are finished successfully or not it deletes db and close connection to db instance.
func TestMain(m *testing.M) {
	test.Setup()
	code := m.Run()
	test.Shutdown()
	os.Exit(code)
}

func TestCreateOwner(t *testing.T) {
	tt := []struct {
		name  string
		email string
		err   string
	}{
		{name: "save owner", email: "test@gmail.com"},
		{name: "save owner with the same email", email: "test@gmail.com", err: "owner with such email is already in the system"},
		{name: "save invalid owner", email: "", err: "owner validation is failed"},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			owner := &model.Owner{
				0,
				"Taras",
				"Fihurnyak",
				model.MALE,
				time.Now(),
				tc.email,
				model.ACTIVE,
				nil,
			}
			r, _ := json.Marshal(owner)
			reg, err := http.NewRequest("POST", "localhost:8080/owner/create", bytes.NewReader(r))
			if err != nil {
				t.Fatalf("could not created request: %v", err)
			}
			rec := httptest.NewRecorder()
			CreateOwner(rec, reg)
			res := rec.Result()
			defer res.Body.Close()

			b, err := ioutil.ReadAll(res.Body)
			if err != nil {
				t.Fatalf("could not read response: %v", err)
			}
			if tc.err != "" {
				if res.StatusCode != http.StatusUnprocessableEntity {
					t.Errorf("expected status 422; got %v", res.StatusCode)
				}
				if msg := string(bytes.TrimSpace(b)); msg != tc.err {
					t.Errorf("expected message %q; got %q", tc.err, msg)
				}
				return
			}
			if res.StatusCode != http.StatusOK {
				t.Errorf("expected status OK; got %v", res.Status)
			}
			ownerResult := &model.Owner{}
			json.Unmarshal([]byte(b), ownerResult)
			if !ownerResult.Valid() {
				t.Error("owner is invalid")
			}

			if ownerResult.Email != tc.email {
				t.Errorf("expected email %v; got %v", tc.email, owner.Email)
			}
		})
	}
}

func TestGetOwner(t *testing.T) {
	tt := []struct {
		name string
		id   string
		err  string
	}{
		{name: "get owner with id=1", id: "1"},
		{name: "missing value", id: "", err: "id field is empty"},
		{name: "id not a number", id: "x", err: "not a number: x"},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			reg, err := http.NewRequest("GET", "localhost:8080/owner?id="+tc.id, nil)
			if err != nil {
				t.Fatalf("could not created request: %v", err)
			}
			rec := httptest.NewRecorder()
			GetOwner(rec, reg)
			res := rec.Result()
			defer res.Body.Close()

			b, err := ioutil.ReadAll(res.Body)
			if err != nil {
				t.Fatalf("could not read response: %v", err)
			}
			if tc.err != "" {
				if res.StatusCode != http.StatusBadRequest {
					t.Errorf("expected status Bad Request; got %v", res.StatusCode)
				}
				if msg := string(bytes.TrimSpace(b)); msg != tc.err {
					t.Errorf("expected message %q; got %q", tc.err, msg)
				}
				return
			}
			if res.StatusCode != http.StatusOK {
				t.Errorf("expected status OK; got %v", res.Status)
			}
			owner := &model.Owner{}
			json.Unmarshal([]byte(b), owner)
			if !owner.Valid() {
				t.Error("owner is invalid")
			}

			id, _ := strconv.Atoi(tc.id)
			if owner.Id != int64(id) {
				t.Errorf("expected id %d; got %v", id, owner.Id)
			}
		})
	}
}
