package auditContract

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/v2/contractapi"
	"github.com/ranger-hadoop-blockchain/hyper/chaincode/utils"
)

// AuditContract provides functions for managing an Asset
type AuditContract struct {
	contractapi.Contract
}

// Asset describes basic details of what makes up a simple asset
// Insert struct field in alphabetic order => to achieve determinism across languages
// golang keeps the order when marshal to json but doesn't order automatically
type AuditStruct struct {
	ID        string `json:"id"`
	Repo      string `json:"repo"`
	Result  int `json:"result"`
	Tags    []string `json:"tags"`
	PolicyVersion    int `json:"policyVersion"`
	Resource string `json:"resource"`
	Timestamp string `json:"timestamp"`
	CliIP string `json:"cliIP"`
	Policy int `json:"policy"`
	ReqUser string `json:"reqUser"`
	EvtTime string `json:"evtTime"`
	ZoneName string `json:"zoneName"`
	AgentHost string `json:"agentHost"`
	ResType string `json:"resType"`
	SeqNum int `json:"seq_num"`
	Cluster string `json:"cluster"`
	ReqData string `json:"reqData"`
	EventCount int `json:"event_count"`
	EventDurMs int `json:"event_dur_ms"`
	Action string `json:"action"`
	Reason string `json:"reason"`
	LogType string `json:"logType"`
	RepoType int `json:"repoType"`
	Sess string `json:"sess"`
	Agent string `json:"agent"`
	Access string `json:"access"`
	Enforcer string `json:"enforcer"`
}

// InitLedger adds a base set of assets to the ledger
func (s *AuditContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
	audits := []AuditStruct{
		// {ID: "1", User: "user123", Resource: "resourceX", Action: "read", Status: "success",Timestamp: "2024-12-21T10:00:00"},
	}

	for _, audit := range audits {
		assetJSON, err := json.Marshal(audit)
		if err != nil {
			println("error");
			return err
		}

		err = ctx.GetStub().PutState(audit.ID, assetJSON)
		if err != nil {
			return fmt.Errorf("failed to put to world state. %v", err)
		}
	}

	return nil
}

// CreateAudit issues a new audit to the world state with given details.
func (s *AuditContract) CreateAudit(ctx contractapi.TransactionContextInterface, id string, jsonData string) error {
	jsonValid := utils.PreprocessToJSON(jsonData)
	var audit AuditStruct
	err := json.Unmarshal([]byte(jsonValid), &audit)
	if err != nil {
		return fmt.Errorf("failed to parse JSON: %v", err)
	}

	auditJSON, err := json.Marshal(audit)
	if err != nil {
		return fmt.Errorf("failed to parse struct: %v", err)
	}

	return ctx.GetStub().PutState(audit.ID, auditJSON)
}

// ReadAsset returns the asset stored in the world state with given id.
// func (s *AuditContract) ReadAsset(ctx contractapi.TransactionContextInterface, id string) (*Asset, error) {
// 	assetJSON, err := ctx.GetStub().GetState(id)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to read from world state: %v", err)
// 	}
// 	if assetJSON == nil {
// 		return nil, fmt.Errorf("the asset %s does not exist", id)
// 	}

// 	var asset Asset
// 	err = json.Unmarshal(assetJSON, &asset)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &asset, nil
// }

// UpdateAsset updates an existing asset in the world state with provided parameters.
// func (s *AuditContract) UpdateAsset(ctx contractapi.TransactionContextInterface, id string, color string, size int, owner string, appraisedValue int) error {
// 	exists, err := s.AssetExists(ctx, id)
// 	if err != nil {
// 		return err
// 	}
// 	if !exists {
// 		return fmt.Errorf("the asset %s does not exist", id)
// 	}

// 	// overwriting original asset with new asset
// 	asset := Asset{
// 		ID:             id,
// 		Color:          color,
// 		Size:           size,
// 		Owner:          owner,
// 		AppraisedValue: appraisedValue,
// 	}
// 	assetJSON, err := json.Marshal(asset)
// 	if err != nil {
// 		return err
// 	}

// 	return ctx.GetStub().PutState(id, assetJSON)
// }

// DeleteAsset deletes an given asset from the world state.
// func (s *AuditContract) DeleteAsset(ctx contractapi.TransactionContextInterface, id string) error {
// 	exists, err := s.AssetExists(ctx, id)
// 	if err != nil {
// 		return err
// 	}
// 	if !exists {
// 		return fmt.Errorf("the asset %s does not exist", id)
// 	}

// 	return ctx.GetStub().DelState(id)
// }

// AuditExists returns true when asset with given ID exists in world state
func (s *AuditContract) AuditExists(ctx contractapi.TransactionContextInterface, id string) (bool, error) {
	auditJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return false, fmt.Errorf("failed to read from world state: %v", err)
	}

	return auditJSON != nil, nil
}

// TransferAsset updates the owner field of asset with given id in world state, and returns the old owner.
// func (s *AuditContract) TransferAsset(ctx contractapi.TransactionContextInterface, id string, newOwner string) (string, error) {
// 	asset, err := s.ReadAsset(ctx, id)
// 	if err != nil {
// 		return "", err
// 	}

// 	oldOwner := asset.Owner
// 	asset.Owner = newOwner

// 	assetJSON, err := json.Marshal(asset)
// 	if err != nil {
// 		return "", err
// 	}

// 	err = ctx.GetStub().PutState(id, assetJSON)
// 	if err != nil {
// 		return "", err
// 	}

// 	return oldOwner, nil
// }

// GetAllAudits returns all audit found in world state
func (s *AuditContract) GetAllAudits(ctx contractapi.TransactionContextInterface) ([]*AuditStruct, error) {
	// range query with empty string for startKey and endKey does an
	// open-ended query of all audit in the chaincode namespace.
	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var audits []*AuditStruct
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var audit AuditStruct
		err = json.Unmarshal(queryResponse.Value, &audit)
		if err != nil {
			return nil, err
		}
		audits = append(audits, &audit)
	}

	return audits, nil
}
