/*
Contact Andreas Kind @IBM, 2016
*/

package main

import (
	"errors"
	"strings"
	"fmt"
	"time"
	"encoding/json"
	//"strconv"
	"github.com/openblockchain/obc-peer/openchain/chaincode/shim"
)

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

func (t *SimpleChaincode) init(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	// Initialize mortgage state in ledger
	key := "myMortgageState"
	jsonState := `{
    "mortgageDocs": {
        "myDocProperty": {
            "name": "myDocProperty",
            "status": "open",
            "signatures": {},
            "form": {},
            "access": {
                "L": {
                    "read": "OK",
                    "save": "OK",
                    "sign": ""
                },
                "B": {
                    "read": "OK",
                    "save": "OK",
                    "sign": ""
                },
                "S": {
                    "read": "OK",
                    "save": "OK",
                    "sign": "OK"
                },
                "A": {
                    "read": "OK",
                    "save": "",
                    "sign": ""
                },
                "N": {
                    "read": "OK",
                    "save": "",
                    "sign": ""
                }
            }
        },
        "myDocBorrower": {
            "name": "myDocBorrower",
            "status": "open",
            "signatures": {},
            "form": {},
            "access": {
                "L": {
                    "read": "OK",
                    "save": "OK",
                    "sign": ""
                },
                "B": {
                    "read": "OK",
                    "save": "OK",
                    "sign": "OK"
                },
                "S": {
                    "read": "OK",
                    "save": "",
                    "sign": ""
                },
                "A": {
                    "read": "OK",
                    "save": "",
                    "sign": ""
                },
                "N": {
                    "read": "OK",
                    "save": "",
                    "sign": ""
                }
            }
        },
        "myDocEmployment": {
            "name": "myDocEmployment",
            "status": "open",
            "signatures": {},
            "form": {},
            "access": {
                "L": {
                    "read": "OK",
                    "save": "OK",
                    "sign": ""
                },
                "B": {
                    "read": "OK",
                    "save": "OK",
                    "sign": "OK"
                },
                "S": {
                    "read": "",
                    "save": "",
                    "sign": ""
                },
                "A": {
                    "read": "",
                    "save": "",
                    "sign": ""
                },
                "N": {
                    "read": "",
                    "save": "",
                    "sign": ""
                }
            }
        },
            "myDocMonthlyIncomeHousingExpense": {
            "name": "myDocMonthlyIncomeHousingExpense",
            "status": "open",
            "signatures": {},
            "form": {},
            "access": {
                "L": {
                    "read": "OK",
                    "save": "OK",
                    "sign": ""
                },
                "B": {
                    "read": "OK",
                    "save": "OK",
                    "sign": "OK"
                },
                "S": {
                    "read": "",
                    "save": "",
                    "sign": ""
                },
                "A": {
                    "read": "",
                    "save": "",
                    "sign": ""
                },
                "N": {
                    "read": "",
                    "save": "",
                    "sign": ""
                }
            }
        },
	    "myDocAssetsLiabilities": {
            "name": "myDocAssetsLiabilities",
            "status": "open",
            "signatures": {},
            "form": {},
            "access": {
                "L": {
                    "read": "OK",
                    "save": "OK",
                    "sign": ""
                },
                "B": {
                    "read": "OK",
                    "save": "OK",
                    "sign": "OK"
                },
                "S": {
                    "read": "",
                    "save": "",
                    "sign": ""
                },
                "A": {
                    "read": "",
                    "save": "",
                    "sign": ""
                },
                "N": {
                    "read": "OK",
                    "save": "",
                    "sign": ""
                }
            }
        },
        "myDocMortgageType": {
            "name": "myDocMortgageType",
            "status": "open",
            "signatures": {},
            "form": {},
            "access": {
                "L": {
                    "read": "OK",
                    "save": "OK",
                    "sign": ""
                },
                "B": {
                    "read": "OK",
                    "save": "OK",
                    "sign": "OK"
                },
                "S": {
                    "read": "",
                    "save": "",
                    "sign": ""
                },
                "A": {
                    "read": "",
                    "save": "",
                    "sign": ""
                },
                "N": {
                    "read": "OK",
                    "save": "",
                    "sign": ""
                }
            }
        },
        "myDocMortgagePurpose": {
            "name": "myDocMortgagePurpose",
	    "status": "open",
            "signatures": {},
            "form": {},
            "access": {
                "L": {
                    "read": "OK",
                    "save": "",
                    "sign": ""
                },
                "B": {
                    "read": "OK",
                    "save": "OK",
                    "sign": "OK"
                },
                "S": {
                    "read": "",
                    "save": "",
                    "sign": ""
                },
                "A": {
                    "read": "",
                    "save": "",
                    "sign": ""
                },
                "N": {
                    "read": "OK",
                    "save": "",
                    "sign": ""
                }
            }
        },
	    "myDocDetailsOfTransaction": {
            "name": "myDocDetailsOfTransaction",
            "status": "open",
            "signatures": {},
            "form": {},
            "access": {
                "L": {
                    "read": "OK",
                    "save": "OK",
                    "sign": ""
                },
                "B": {
                    "read": "OK",
                    "save": "OK",
                    "sign": "OK"
                },
                "S": {
                    "read": "",
                    "save": "",
                    "sign": ""
                },
                "A": {
                    "read": "",
                    "save": "",
                    "sign": ""
                },
                "N": {
                    "read": "OK",
                    "save": "",
                    "sign": ""
                }
            }
        },
	    "myDocForGovMonitoring": {
            "name": "myDocForGovMonitoring",
            "status": "open",
            "signatures": {},
            "form": {},
            "access": {
                "L": {
                    "read": "OK",
                    "save": "OK",
                    "sign": "OK"
                },
                "B": {
                    "read": "OK",
                    "save": "OK",
                    "sign": "OK"
                },
                "S": {
                    "read": "",
                    "save": "",
                    "sign": ""
                },
                "A": {
                    "read": "",
                    "save": "",
                    "sign": ""
                },
                "N": {
                    "read": "",
                    "save": "",
                    "sign": ""
                }
            }
        },
            "myDocAcknowledgementAgreement": {
            "name": "myDocAcknowledgementAgreement",
            "status": "open",
            "signatures": {},
            "form": {},
            "access": {
                "L": {
                    "read": "OK",
                    "save": "",
                    "sign": ""
                },
                "B": {
                    "read": "OK",
                    "save": "OK",
                    "sign": "OK"
                },
                "S": {
                    "read": "",
                    "save": "",
                    "sign": ""
                },
                "A": {
                    "read": "",
                    "save": "",
                    "sign": ""
                },
                "N": {
                    "read": "OK",
                    "save": "",
                    "sign": ""
                }
            }
        },
        "myDocPurchaseContract": {
            "name": "myDocPurchaseContract",
            "status": "open",
            "signatures": {},
            "form": {},
            "access": {
                "L": {
                    "read": "OK",
                    "save": "",
                    "sign": ""
                },
                "B": {
                    "read": "OK",
                    "save": "OK",
                    "sign": "OK"
                },
                "S": {
                    "read": "OK",
                    "save": "OK",
                    "sign": "OK"
                },
                "A": {
                    "read": "",
                    "save": "",
                    "sign": ""
                },
                "N": {
                    "read": "OK",
                    "save": "",
                    "sign": ""
                }
            }
        }
    },
    "mortgageProcesses": {
        "mortgageOrigination": {
            "nodes": [{
                "id": "node0",
                "name": "Borrower Information",
                "status": "open",
                "next": "node3",
                "prev": [],
                "gridx": 1,
                "gridy": 0,
                "docs": ["myDocBorrower","myDocEmployment","myDocMonthlyIncomeHousingExpense","myDocAssetsLiabilities"]
            }, {
                "id": "node1",
                "name": "Property Information",
                "status": "open",
                "next": "node2",
                "prev": [],
                "gridx": 0,
                "gridy": 1,
                "docs": ["myDocProperty"]
            }, {
                "id": "node2",
                "name": "Purchase Contract Information",
                "status": "closed",
                "next": "node3",
                "prev": ["node1"],
                "gridx": 1,
                "gridy": 1,
                "docs": ["myDocPurchaseContract"]
            }, {
                "id": "node3",
                "name": "Mortgage Application",
                "status": "closed",
                "next": "node4",
                "prev": ["node0", "node2"],
                "gridx": 2,
                "gridy": 1,
                "docs": ["myDocMortgageType","myDocMortgagePurpose","myDocDetailsOfTransaction","myDocForGovMonitoring","myDocAcknowledgementAgreement"]
            }, {
                "id": "node4",
                "name": "Good Faith Estimate",
                "status": "closed",
                "next": "node5",
                "prev": ["node3"],
                "gridx": 3,
                "gridy": 1,
                "docs": []
            }, {
                "id": "node5",
                "name": "Truth in Lending Statement",
                "status": "closed",
                "next": "node6",
                "prev": ["node4"],
                "gridx": 4,
                "gridy": 1,
                "docs": []
            }, {
                "id": "node6",
                "name": "Borrower Verification",
                "status": "closed",
                "next": "node7",
                "prev": ["node5"],
                "gridx": 5,
                "gridy": 1,
                "docs": []
            }, {
                "id": "node7",
                "name": "Purchase Contract Verification",
                "status": "closed",
                "next": "node8",
                "prev": ["node6"],
                "gridx": 6,
                "gridy": 1,
                "docs": []
            }, {
                "id": "node8",
                "name": "Employment Verification",
                "status": "closed",
                "next": "node9",
                "prev": ["node7"],
                "gridx": 6,
                "gridy": 2,
                "docs": []
            }, {
                "id": "node9",
                "name": "Rent Verification",
                "status": "closed",
                "next": "node10",
                "prev": ["node8"],
                "gridx": 5,
                "gridy": 2,
                "docs": []
            }, {
                "id": "node10",
                "name": "Desktop Underwriter / Loan Prospector Program",
                "status": "closed",
                "next": "node11",
                "prev": ["node9"],
                "gridx": 4,
                "gridy": 2,
                "docs": []
            }, {
                "id": "node11",
                "name": "Title Order, Home Insurance",
                "status": "closed",
                "next": "node12",
                "prev": ["node10"],
                "gridx": 3,
                "gridy": 2,
                "docs": []
            }, {
                "id": "node12",
                "name": "Home Appraisal",
                "status": "closed",
                "next": "node13",
                "prev": ["node11"],
                "gridx": 2,
                "gridy": 2,
                "docs": []
            }, {
                "id": "node13",
                "name": "Mortgage Contract",
                "status": "closed",
                "next": "",
                "prev": ["node12"],
                "gridx": 1,
                "gridy": 2,
                "docs": []
            }],
            "edges": [
                [{
                    "gridx": 1.0,
                    "gridy": 0.0
                }, {
                    "gridx": 1.5,
                    "gridy": 0.0
                }, {
                    "gridx": 2.0,
                    "gridy": 0.5
                }, {
                    "gridx": 2.0,
                    "gridy": 1.0
                }],
                [{
                    "gridx": 0.0,
                    "gridy": 1.0
                }, {
                    "gridx": 1.0,
                    "gridy": 1.0
                }, {
                    "gridx": 2.0,
                    "gridy": 1.0
                }, {
                    "gridx": 3.0,
                    "gridy": 1.0
                }, {
                    "gridx": 4.0,
                    "gridy": 1.0
                }, {
                    "gridx": 5.0,
                    "gridy": 1.0
                }, {
                    "gridx": 6.0,
                    "gridy": 1.0
                }],
                [{
                    "gridx": 6.0,
                    "gridy": 1.5
                }, {
                    "gridx": 6.0,
                    "gridy": 2.0
                }],
                [{
                    "gridx": 1.0,
                    "gridy": 2.0
                }, {
                    "gridx": 2.0,
                    "gridy": 2.0
                }, {
                    "gridx": 3.0,
                    "gridy": 2.0
                }, {
                    "gridx": 4.0,
                    "gridy": 2.0
                }, {
                    "gridx": 5.0,
                    "gridy": 2.0
                }, {
                    "gridx": 6.0,
                    "gridy": 2.0
                }]
            ],
            "auditing": {
                "auditLogs": ""
	    }
        }
    },
    "rootHash": ""
}`
	
	fmt.Printf("Initializing mortgage state in shared ledger...\n")
	err := stub.PutState(key, []byte(jsonState))
	if err != nil {
		return nil, err
	}
	
	return nil, nil
}

func (t *SimpleChaincode) getMortgageState(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	var err error
	
	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting token")
	}

	token := args[0]
	role := token[:1]

	state, err := getMortgageStateFull(stub)
	if err != nil {
		return nil, err
	}
	docs := state["mortgageDocs"].(map[string]interface{})
	if docs == nil {
		return nil, myError("Missing documents in mortage state")
	}
	for docId := range docs {
		_, _, rights, err := getMortgageDoc(state, docId, role)
		if err != nil {
			return nil, err
		}

		if rights["read"] != "OK" {
			// Mask doc
			var form map[string]interface{}
			err = json.Unmarshal([]byte(`{}`), &form)
			if err != nil {
				return nil, myError("Cannot unmarshal form '{}'")
			}
			
			doc := docs[docId].(map[string]interface{})
			if doc == nil {
				return nil, myError("Missing dcument '" + docId + "' in mortgage state")
			}
			doc["form"] = form
			msg := fmt.Sprintf("User of role '" + role + "' has no right to read document '" + docId + "'; document masked")
			//myAuditLog(state, "INFO", msg)
			fmt.Printf(msg + "\n")
		} else {
			msg := fmt.Sprintf("User of role '" + role + "' read document '" + docId + "'")
			//myAuditLog(state, "INFO", msg)
			fmt.Printf(msg + "\n")
		}
	}
	
	bytes, err := json.Marshal(state)
	if err != nil {
		//return bytes, err
		return bytes, myError("Cannot marshal mortgage state")
	}

	fmt.Printf("Returning state...\n")
	//fmt.Printf("Returning state: " + string(bytes) + "\n")
	return bytes, nil
}

func myAuditLog(state map[string]interface{}, prefix string, msg string)(error) {
	// Get the log
	processes := state["mortgageProcesses"].(map[string]interface{})
	if processes == nil {
		return myError("No mortgage processses in mortgage state")
	}
	process := processes["mortgageOrigination"].(map[string]interface{})
	if process == nil {
		return myError("No mortgage origination processs in mortgage state")
	}
	auditing := process["auditing"].(map[string]interface{})
	if auditing == nil {
		return myError("No mortgage processs auditing information in mortgage origination process")
	}
	logs := auditing["auditLogs"].(string)

	// Get time prefix
	t := time.Now()
	nowStr := t.Format("2006-01-02 15:04:05")

	// Write first line 
	if logs == "" {
		logs = "<pre><b style=\"color: white\">" + nowStr + " [INFO]:</b> Audit log started...</pre>"
	}

	// Get topic prefix
	color := "white"
	if prefix == "SIGN" {
		color = "green"
	} else if prefix == "REVOKE" {
		color = "red"
	} else if prefix == "SAVE" {
		color = "orange"
	} else {
		color = "white"
	}
	fmt.Printf(msg + "\n");
	auditing["auditLogs"] = fmt.Sprintf("%s<pre><b style=\"color: " + color + "\">" + nowStr + " [" + prefix + "]:</b> " + msg + "</pre>", logs)
	return nil
}

// Set mortgage document
func (t *SimpleChaincode) setMortgageDoc(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	var err error
	
	if len(args) != 5 {
		return nil, myError("Incorrect number of arguments. Expecting 5: token, docId, operation, docForm and signature")
	}
	
	token := args[0]
	docId := args[1]
	op := args[2]
	fieldsStr := args[3]
	signature := args[4]

	if token == "" {
		return nil, myError("Missing token")
	}
	if docId == "" {
		return nil, myError("Missing document identification")
	}
	if op == "" {
		return nil, myError("Missing operation")
	}
	if fieldsStr == "" {
		return nil, myError("Missing fields information (eg, '{}')")
	}
	if signature == "" {
		return nil, myError("Missing signature")
	}

	// Get user role
	role := token[:1]

	// Get submitted mortage fields
	var fields map[string]interface{}
	err = json.Unmarshal([]byte(fieldsStr), &fields)
	if err != nil {
		return nil, myError("Cannot unmarshal fields for " + op)
	}
	state, err := getMortgageStateFull(stub)
	if err != nil {
		return nil, err
	}
	
	// Get document, form and user access rights
	doc, form, rights, err := getMortgageDoc(state, docId, role)	
	if err != nil {
		return nil, err
	}
	
	// Get doc signatures
	sigs := doc["signatures"].(map[string]interface{})
	if sigs == nil {
		return nil, myError("No signature provided; " + op + " ignored")
	}	

	// Check if operation is allowed; who can sign, can also revoke
	if rights[op] != "OK" && (op == "revo" && rights["sign"] != "OK") {
		return nil, myError("User of role '" + role + "' has no right to " + op + " document '" + docId + "'")
	}

	if op == "save" {
		i, unit, err := updateDocFields(form, fields)
		if err != nil {
			return nil, err
		}
		msg := fmt.Sprintf("Role '" + role + "' updated document '" + docId + "' (%d " + unit + ")", i)
		myAuditLog(state, "SAVE", msg)
		
	} else if op == "sign" {
		// Already signed by user?
		if sigs[role] != nil && sigs[role] != "" {
			return nil, myError("User of role '" + role + "' signed document '" + docId + "' already; no action")
		}
			
		// Mark doc as signed by user
		sigs[role] = signature

		_, err = updateDocStatus(token, doc, op)
		if err != nil {
			return nil, err
		}

		i, unit, err := updateDocFields(form, fields)
		if err != nil {
			return nil, err
		}	
		msg := fmt.Sprintf("User of role '" + role + "' updated and signed document '" + docId + "' (%d " + unit + ")", i)
		myAuditLog(state, "SIGN", msg)
		
		err = updateNodeStatus(token, state, op)
		if err != nil {
			return nil, err
		}

	} else if (op == "revo") {
		// Cannot revoke if not yet signed by user
		if sigs[role] == nil || sigs[role] == "" {
			return nil, myError("User of role '" + role + "' has not signed document '" + docId + "' yet; no action")
		}
			
		// Unmark doc as signed by user
		sigs[role] = ""
		doc["status"] = "closed"
		msg := fmt.Sprintf("Role '" + role + "' revoked signature for document '" + docId + "'")
		myAuditLog(state, "REVOKE", msg)
		
		_, err = updateDocStatus(token, doc, op)
		if err != nil {
			return nil, err
		}
		
		err = updateNodeStatus(token, state, op)
		if err != nil {
			return nil, err
		}
		
	}
		
	_, err = writeMortgageState(stub, state)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func updateDocFields(form map[string]interface{}, fields map[string]interface{}) (int, string, error) {
	// Update fields
	i := 0
	unit := "fields"
	for k, v := range fields {
		fmt.Printf("Update document field: %s -> %s\n", k, v)
		form[k] = v
		i++
		if i == 1 {
			unit = "field"
		}
	}
	return i, unit, nil
}

func allNodeDocsSigned(state map[string]interface{}, node map[string]interface{}, role string) (bool, bool) {
	dIds := node["docs"].([]interface{})
	noDocs := true
	allSigned := true
	// Walk through all docs of node
	for _, dId := range dIds {
		d, _, _, err := getMortgageDoc(state, dId.(string), role)
		if err != nil {
			fmt.Printf("Document '" + dId.(string) + "' does not exist!\n")
			continue
		}
		noDocs = false
		
		if d["status"] != "signed" {
			allSigned = false
		}
	}
	return allSigned, noDocs
}

func allPrevNodesSigned(nodes []interface{}, node map[string]interface{}) (bool, bool) {
	nIds := node["prev"].([]interface{})
	noPrev := true
	allSigned := true
	// Walk through all prev nodes of node
	for _, nId := range nIds {	
		prev := findNode(nodes, nId.(string))
		if prev == nil {
			continue
		}
		noPrev = false
		
		if prev["status"] != "signed" {
			allSigned = false
		}
	}
	return allSigned, noPrev
}

func updateNodeStatus(token string, state map[string]interface{}, op string) (error) {
	// Get user role
	role := token[:1]
	
	// Get process nodes to see later if nodes should change status
	processes := state["mortgageProcesses"].(map[string]interface{})
	if processes == nil {
		return myError("No mortgage processses in mortgage state")
	}
	process := processes["mortgageOrigination"].(map[string]interface{})
	if process == nil {
		return myError("No mortgage origination processs in mortgage state")
	}
	nodes := process["nodes"].([]interface{})
	if nodes == nil {
		return myError("No mortgage processs nodes in mortgage origination process")
	}
	
	// Walk through all nodes and check which have all required docs signed off
	for _, n := range nodes {
		var nn map[string]interface{}
		nn = n.(map[string]interface{})
		allDocsSigned, _ := allNodeDocsSigned(state, nn, role)
		if allDocsSigned && nn["status"] == "open" {
			// Node nn has all docs signed off
			fmt.Printf("Node nn has all docs signed off:\n")
			nn["status"] = "signed"
			updateNextNode(nodes, nn, "sign")
		} else if !allDocsSigned && nn["status"] == "signed" {
			// Node nn has not all docs signed off
			//fmt.Printf("Node nn has not all docs signed off:\n")
			nn["status"] = "open"
			updateNextNode(nodes, nn, "revo")
		}
		fmt.Printf("Node " + nn["id"].(string) + " changed to status " + nn["status"].(string) + "\n")
	}
	return nil
}

func findNode(nodes []interface{}, id string)(map[string]interface{}) {
	// Find node with id
	// TODO: should change state from list to map data structure
	var nn map[string]interface{}
	for _, n := range nodes {
		nn = n.(map[string]interface{})
		if id == nn["id"].(string) {
			return nn
		}
	}
	return nil
}

func updateNextNode(nodes []interface{}, node map[string]interface{}, op string) (error) {
	// Get next node
	nxtId := node["next"].(string)
	nxt := findNode(nodes, nxtId)
	if nxt == nil {
		return nil;
	}
	
	if op == "sign" {
		// On signing: open next only if all prev nodes of next are signed
		allPrevSigned, _ := allPrevNodesSigned(nodes, nxt)
		if allPrevSigned {
			nxt["status"] = "open"
		}
	} else if op == "revo" {
		// On revoction: always close next
		nxt["status"] = "close"
		updateNextNode(nodes, nxt, op)
	}
	return nil
}

func updateDocStatus(token string, doc map[string]interface{}, op string) (string, error){
	// Get document name
	docId := doc["name"].(string)
	if docId == "" {
		return "", myError("Missing document name")
	}
	// Get document access rights
	access := doc["access"].(map[string]interface{})
	if access == nil {
		return "", myError("Missing access rights in document '" + docId)
	}
	// Get doc signatures
	sigs := doc["signatures"].(map[string]interface{})
	if sigs == nil {
		return "", myError("No signatures in document docId")
	}
	
	roles := "LBSAN"
	allSignatures := true	
	// Determine if everyone signed doc
	for i := 0; i < len(roles); i++ {
		ro := roles[i:i+1]
		rs := access[ro].(map[string]interface{})
		if rs == nil {
			return "", myError("Missing access rights for role '" + ro + "' in document '" + docId)
		}
		if rs["sign"] == "OK" && (sigs[ro] == nil || sigs[ro] == "") {
			// Role should sign, but has not yet signed
			allSignatures = false
			break
		}
	}
	status := "open"
	if allSignatures {
		status = "signed"
	}
	doc["status"] = status
	fmt.Printf("Changed status of document '" + docId + "' to " + status + "\n")
	return doc["status"].(string), nil
}

func writeMortgageState(stub *shim.ChaincodeStub, state map[string]interface{}) ([]byte, error) {
	// Write updated state back to ledger
	bytes, err := json.Marshal(state)
	if err != nil {
		return bytes, err
	}
		
	//fmt.Printf("State: " + string(bytes) + "\n")
	fmt.Printf("Writing mortgage state to shared ledger...\n")
	err = stub.PutState("myMortgageState", bytes)
	if err != nil {
		return bytes, myError("Failed to write mortgage state to ledger")
		//return bytes, err
	}
	return bytes, nil
}

func myError(msg string) (error) {
	fmt.Printf(msg + "!\n")
	jsonResp := "{\"Error\":\"" + msg + "\"}"
	return errors.New(jsonResp)
}

func getMortgageStateFull(stub *shim.ChaincodeStub) (map[string]interface{}, error) {
	// Get mortgage state from ledger
	fmt.Printf("Reading mortgage state from shared ledger...\n")
	var state map[string]interface{}
	key := "myMortgageState"
	bytes, err := stub.GetState(key)
	if err != nil || bytes == nil {
		return nil, myError("No mortage state in ledger")
	}
	err = json.Unmarshal(bytes, &state)
	if err != nil {
		return nil, myError("Cannot unmarshal mortage state")
	}
	return state, nil
}

func getMortgageDoc(state map[string]interface{}, docId string, role string) (map[string]interface{}, map[string]interface{}, map[string]interface{}, error) {
	// Retrieve doc, form and access rights from ledger
	docs := state["mortgageDocs"].(map[string]interface{})
	if docs == nil {
		return nil, nil, nil, myError("Missing documents in mortage state")
	}

	doc := docs[docId].(map[string]interface{})
	if doc == nil {
		return nil, nil, nil, myError("Missing dcument '" + docId + "' in mortgage state")
	}
	form := doc["form"].(map[string]interface{})
	if form == nil {
		return nil, nil, nil, myError("Missing form in document '" + docId)
	}
	access := doc["access"].(map[string]interface{})
	if access == nil {
		return nil, nil, nil, myError("Missing access rights in document '" + docId)
	}
	if strings.Index("LBSAN", role) < 0 {
		return nil, nil, nil, myError("Bad role '" + role + "'")
	}
	rights := access[role].(map[string]interface{})
	if rights == nil {
		return nil, nil, nil, myError("Missing access rights for role '" + role + "' in document '" + docId)
	}
	
	return doc, form, rights, nil
}

func (t *SimpleChaincode) Run(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
	// Handle different functions
	if function == "init" {
		return t.init(stub, args)
	} else if function == "invoke" {
		return t.invoke(stub, args)
	}

	return nil, errors.New("Received unknown function invocation")
}

func (t *SimpleChaincode) invoke(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	return t.setMortgageDoc(stub, args)
}

func (t *SimpleChaincode) Query(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
	return t.getMortgageState(stub, args)
}

func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}
