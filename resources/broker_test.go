package resources_test

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/vschettino/exfin/db"
	m "github.com/vschettino/exfin/models"
	"github.com/vschettino/exfin/tests"
	"testing"
)

func CreateBrokerRequest() map[string]string {
	return map[string]string{
		"name":        "My Account",
		"institution": "That Broker Cia",
	}
}

var BrokersPath = "/brokers"

func TestBrokerUnauthorized(t *testing.T) {
	w, _ := tests.MakeGET(BrokersPath, "")
	assert.Equal(t, 401, w.Code)
}

func TestGetBrokersSuccessful(t *testing.T) {
	tests.GenerateToken("admin@exfin.org", 1)
	w, _ := tests.MakeGET(BrokersPath, tests.GenerateToken("admin@exfin.org", 1))
	assert.Equal(t, 200, w.Code)
}

func TestGetBrokerDoesntExist(t *testing.T) {
	w, _ := tests.MakeGET(BrokersPath+"/66644", tests.GenerateToken("admin@exfin.org", 1))
	assert.Equal(t, 404, w.Code)
	w.Body.String()
}

func TestGetBrokerBadParams(t *testing.T) {
	w, _ := tests.MakeGET(BrokersPath+"/nostring66644", tests.GenerateToken("admin@exfin.org", 1))
	assert.Equal(t, 400, w.Code)
	w.Body.String()
}

func TestGetBrokerSuccessful(t *testing.T) {
	conn := db.Connection()
	broker := m.Broker{
		Name:        "Existing Name",
		Institution: "Existing Institution",
		AccountId:   1,
	}
	_ = conn.Insert(&broker)
	w, _ := tests.MakeGET(fmt.Sprintf(BrokersPath+"/%d", broker.Id), tests.GenerateToken("admin@exfin.org", 1))
	assert.Equal(t, 200, w.Code)
	var responseBroker m.Broker
	assert.NoError(t, json.Unmarshal([]byte(w.Body.String()), &responseBroker))
	assert.Equal(t, broker.Id, responseBroker.Id)
	t.Cleanup(func() {
		_ = conn.Delete(&broker)
	})
}

func TestCreateBrokerMissingRequiredParameters(t *testing.T) {
	req := CreateBrokerRequest()
	token := tests.GenerateToken("admin@exfin.org", 1)
	delete(req, "name")
	w, _ := tests.MakePOST(BrokersPath, req, token)
	assert.Equal(t, 400, w.Code)
}

func TestCreateBrokerSuccessful(t *testing.T) {
	conn := db.Connection()
	var account m.Broker
	req := CreateBrokerRequest()
	token := tests.GenerateToken("admin@exfin.org", 1)
	w, _ := tests.MakePOST(BrokersPath, req, token)
	assert.Equal(t, 201, w.Code)
	assert.NoError(t, json.Unmarshal([]byte(w.Body.String()), &account))
	assert.Equal(t, account.Name, req["name"])
	t.Cleanup(func() {
		_ = conn.Delete(&account)
	})
}

func TestUpdateUnknownBroker(t *testing.T) {
	req := CreateBrokerRequest()
	token := tests.GenerateToken("admin@exfin.org", 1)
	w, _ := tests.MakePATCH(BrokersPath+"/987987", req, token)
	assert.Equal(t, 404, w.Code)
}

func TestUpdateInvalidBrokerId(t *testing.T) {
	req := CreateBrokerRequest()
	token := tests.GenerateToken("admin@exfin.org", 1)
	w, _ := tests.MakePATCH(BrokersPath+"/fid", req, token)
	assert.Equal(t, 400, w.Code)
}

func TestUpdateBrokerSuccessful(t *testing.T) {
	conn := db.Connection()
	broker := m.Broker{
		Name:        "Existing Name",
		Institution: "Existing Institution",
		AccountId:   1,
	}
	_ = conn.Insert(&broker)
	req := CreateAccountRequest()
	delete(req, "institution")
	token := tests.GenerateToken("admin@exfin.org", 1)
	w, _ := tests.MakePATCH(fmt.Sprintf(BrokersPath+"/%d", broker.Id), req, token)
	assert.Equal(t, 200, w.Code)
	assert.NoError(t, json.Unmarshal([]byte(w.Body.String()), &broker))
	assert.Equal(t, broker.Name, req["name"])
	t.Cleanup(func() {
		_ = conn.Delete(&broker)
	})
}
