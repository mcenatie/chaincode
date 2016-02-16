/*
Contact Andreas Kind @IBM, 2016
*/

package main

import (
	"errors"
	"strings"
	"fmt"
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
	jsonState := `
{
    "mortgageDocs": {
        "myDocProperty": {
            "name": "myDocProperty",
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
        }
    },
    "mortgageProcesses": {
        "mortgageOrigination": {
            "nodes": [{
                "id": "node0",
                "name": "Borrower Information",
                "status": "open",
                "next": "node3",
                "gridx": 1,
                "gridy": 0,
                "docs": ["myDocBorrower"]
            }, {
                "id": "node1",
                "name": "Property Information",
                "status": "open",
                "next": "node2",
                "gridx": 0,
                "gridy": 1,
                "docs": ["myDocProperty"]
            }, {
                "id": "node2",
                "name": "Purchase Contract Information",
                "status": "closed",
                "next": "node3",
                "gridx": 1,
                "gridy": 1,
                "docs": []
            }, {
                "id": "node3",
                "name": "Mortgage Application",
                "status": "closed",
                "next": "node4",
                "gridx": 2,
                "gridy": 1,
                "docs": []
            }, {
                "id": "node4",
                "name": "Good Faith Estimate",
                "status": "closed",
                "next": "node5",
                "gridx": 3,
                "gridy": 1,
                "docs": []
            }, {
                "id": "node5",
                "name": "Truth in Lending Statement",
                "status": "closed",
                "next": "node6",
                "gridx": 4,
                "gridy": 1,
                "docs": []
            }, {
                "id": "node6",
                "name": "Borrower Verification",
                "status": "closed",
                "next": "node7",
                "gridx": 5,
                "gridy": 1,
                "docs": []
            }, {
                "id": "node7",
                "name": "Purchase Contract Verification",
                "status": "closed",
                "next": "node8",
                "gridx": 6,
                "gridy": 1,
                "docs": []
            }, {
                "id": "node8",
                "name": "Employment Verification",
                "status": "closed",
                "next": "node9",
                "gridx": 6,
                "gridy": 2,
                "docs": []
            }, {
                "id": "node9",
                "name": "Rent Verification",
                "status": "closed",
                "next": "node10",
                "gridx": 5,
                "gridy": 2,
                "docs": []
            }, {
                "id": "node10",
                "name": "Desktop Underwriter / Loan Prospector Program",
                "status": "closed",
                "next": "node11",
                "gridx": 4,
                "gridy": 2,
                "docs": []
            }, {
                "id": "node11",
                "name": "Title Order, Home Insurance",
                "status": "closed",
                "next": "node12",
                "gridx": 3,
                "gridy": 2,
                "docs": []
            }, {
                "id": "node12",
                "name": "Home Appraisal",
                "status": "closed",
                "next": "node13",
                "gridx": 2,
                "gridy": 2,
                "docs": []
            }, {
                "id": "node13",
                "name": "Mortgage Contract",
                "status": "closed",
                "next": "",
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
            ]
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
		_, _, rights, err := getMortgageDoc(stub, state, docId, role)
		if err != nil {
			return nil, err
		}

		op := "read"
		if rights[op] != "OK" {
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
			fmt.Printf("User of role '" + role + "' has no right to " + op + " document '" + docId + "'; document masked\n")
		}
	}
	bytes, _ := json.Marshal(state)

	fmt.Printf("Returning state...\n")
	//fmt.Printf("Returning state: " + string(bytes) + "\n")
	return bytes, nil
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
	formStr := args[3]
	signature := args[4]

	// Get user role
	role := token[:1]

	// Get submitted mortage fields
	var form map[string]interface{}
	err = json.Unmarshal([]byte(formStr), &form)
	if err != nil {
		return nil, myError("Cannot unmarshal form for " + op)
	}
	state, err := getMortgageStateFull(stub)
	if err != nil {
		return nil, err
	}
	doc, f, rights, err := getMortgageDoc(stub, state, docId, role)
	if err != nil {
		return nil, err
	}
	if rights[op] == "OK" {
		if op == "sign" {
			// Sign document
			s := doc["signatures"].(map[string]interface{})
			if s == nil {
				return nil, myError("No signature provided; " + op + " ignored")
			}
			s[role] = signature
			doc["status"] = "signed"
			fmt.Printf("Signed " + docId + "\n")

			// Get process nodes
			processes := state["mortgageProcesses"].(map[string]interface{})
			if processes == nil {
				return nil, myError("No mortgage processses in mortgage state")
			}
			process := processes["mortgageOrigination"].(map[string]interface{})
			if process == nil {
				return nil, myError("No mortgage origination processs in mortgage state")
			}
			nodes := process["nodes"].([]interface{})
			if nodes == nil {
				return nil, myError("No mortgage processs nodes in mortgage origination process")
			}

			// Walk through all nodes and check which have all required docs signed off
			for _, n := range nodes {
				var allDocsSigned = true
				var noDocs = true
				var nn map[string]interface{}
				nn = n.(map[string]interface{})
				dIds := nn["docs"].([]interface{})
				for _, dId := range dIds {
					noDocs = false
					d, _, _, err := getMortgageDoc(stub, state, dId.(string), role)
					if err != nil {
						return nil, err
					}
					if d["status"] != "signed" {
						allDocsSigned = false
					}
				}
				if allDocsSigned && !noDocs {
					// Node n has all docs signed off
					nn["status"] = "signed"
					nxt := nn["next"].(string)
					// Open next 
					for _, n2 := range nodes {
						var nn2 map[string]interface{}
						nn2 = n2.(map[string]interface{})
						if nxt == nn2["id"].(string) {
							nn2["status"] = "open"
							fmt.Printf("Opened next node " + nxt + "\n")
						}
					}
				}
			}	
		}
		// Update fields
		for k, v := range form {
			fmt.Printf("Update document fields %s -> %s\n", k, v)
			f[k] = v;
		}
		fmt.Printf("User of role '" + role + " updated document '" + docId + "'!\n")
	} else {
		return nil, myError("User of role '" + role + "' has no right to " + op + " document '" + docId + "'")
	}

	// Write updated state back to ledger
	bytes, _ := json.Marshal(state)
		
	//fmt.Printf("State: " + string(bytes) + "\n")
	fmt.Printf("Writing mortgage state to shared ledger...\n")
	key := "myMortgageState"
	err = stub.PutState(key, bytes)
	if err != nil {
		return nil, myError("Failed to put document '" + docId + "' into ledger")
	}
	
	return nil, nil
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

func getMortgageDoc(stub *shim.ChaincodeStub, state map[string]interface{}, docId string, role string) (map[string]interface{}, map[string]interface{}, map[string]interface{}, error) {
	// Walk through the form fields and update state
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
		return nil, nil, nil, myError("Bad rol '" + role)
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
