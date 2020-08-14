//--------------------------------------------------------------------------
// Copyright 2018 Infinite Devices GmbH
// www.infinimesh.io
//
//   Licensed under the Apache License, Version 2.0 (the "License");
//   you may not use this file except in compliance with the License.
//   You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.
//--------------------------------------------------------------------------

package dgraph

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/dgraph-io/dgo/protos/api"

	"github.com/infinimesh/infinimesh/pkg/node/nodepb"
)

func (s *DGraphRepo) ListAccounts(ctx context.Context) (accounts []*nodepb.Account, err error) {
	txn := s.Dg.NewReadOnlyTxn()

	const q = `query accounts{
                     accounts(func: eq(type, "account")) {
                       uid
                       name
                       enabled
                       isRoot
                     }
                   }`

	res, err := txn.Query(ctx, q)
	if err != nil {
		return nil, err
	}

	var result struct {
		Accounts []*Account `json:"accounts"`
	}

	if err := json.Unmarshal(res.Json, &result); err != nil {
		return nil, err
	}

	for _, account := range result.Accounts {
		accounts = append(accounts, &nodepb.Account{
			Uid:     account.UID,
			Name:    account.Name,
			IsRoot:  account.IsRoot,
			Enabled: account.Enabled,
		})
	}

	return accounts, nil
}

func (s *DGraphRepo) UpdateAccount(ctx context.Context, account *nodepb.UpdateAccountRequest) (err error) {
	txn := s.Dg.NewTxn()

	q := `query userExists($id: string) {
                exists(func: uid($id)) @filter(eq(type, "account")) {
                  uid
                }
              }
             `

	var result struct {
		Exists []map[string]interface{} `json:"exists"`
	}

	resp, err := txn.QueryWithVars(ctx, q, map[string]string{"$id": account.Account.Uid})
	if err != nil {
		return err
	}
	err = json.Unmarshal(resp.Json, &result)
	if err != nil {
		return err
	}

	if len(result.Exists) == 0 {
		return errors.New("Account not found")
	}

	// TODO this may override fields with zero-values
	acc := &Account{
		Node: Node{
			Type: "account",
			UID:  account.Account.Uid,
		},
	}

	for _, field := range account.FieldMask.Paths {
		switch field {
		case "Name":
			acc.Name = account.Account.Name
		case "IsRoot":
			acc.IsRoot = account.Account.IsRoot
		case "Enabled":
			acc.Enabled = account.Account.Enabled
		}
	}

	js, err := json.Marshal(acc)
	if err != nil {
		return err
	}

	m := &api.Mutation{SetJson: js}
	_, err = txn.Mutate(ctx, m)
	if err != nil {
		return err
	}

	err = txn.Commit(ctx)
	if err != nil {
		return errors.New("Failed to commit")
	}
	return nil
}

func (s *DGraphRepo) CreateUserAccount(ctx context.Context, username, password string, isRoot, enabled bool) (uid string, err error) {
	// TODO move this to the controller

	txn := s.Dg.NewTxn()

	q := `query userExists($name: string) {
                exists(func: eq(name, $name)) @filter(eq(type, "account")) {
                  uid
                }
              }
             `

	var result struct {
		Exists []map[string]interface{} `json:"exists"`
	}

	resp, err := txn.QueryWithVars(ctx, q, map[string]string{"$name": username})
	if err != nil {
		return "", err
	}
	err = json.Unmarshal(resp.Json, &result)
	if err != nil {
		return "", err
	}

	//If the user doesnt exists then create user
	if len(result.Exists) == 0 {

		//Create Default namespace for the new user being created
		defaultNs, err := s.CreateNamespace(ctx, username)
		if err != nil {
			return "", err
		}

		//Build the json data structure to create the user for DGraph
		js, err := json.Marshal(&Account{
			Node: Node{
				Type: "account",
				UID:  "_:user",
			},
			Name:    username,
			IsRoot:  isRoot,
			Enabled: enabled,
			HasCredentials: []*UsernameCredential{
				{
					Node: Node{
						Type: "credentials",
					},
					Username: username,
					Password: password,
				},
			},
			AccessToNamespace: []*Namespace{
				&Namespace{
					Node: Node{
						UID: defaultNs,
					},
					AccessToPermission: nodepb.Action_WRITE.String(),
				},
			},
			DefaultNamespace: []*Namespace{
				&Namespace{
					Node: Node{
						UID: defaultNs,
					},
				},
			},
		})
		if err != nil {
			return "", err
		}

		//Create the new user in the Dgraph DB
		m := &api.Mutation{SetJson: js}
		a, err := txn.Mutate(ctx, m)
		if err != nil {
			return "", err
		}

		err = txn.Commit(ctx)
		if err != nil {
			return "", errors.New("Failed to commit")
		}
		userUID := a.GetUids()["user"]
		return userUID, nil

	}

	//If the user exists then return error
	return "", errors.New("User exists already")
}

func (s *DGraphRepo) GetAccount(ctx context.Context, name string) (account *nodepb.Account, err error) {
	txn := s.Dg.NewReadOnlyTxn()
	const q = `query accounts($account: string) {
                     accounts(func: uid($account)) @filter(eq(type, "account"))  {
                       uid
                       name
                       type
                       isRoot
                       enabled
                       default.namespace {
                         name
                         uid
                       }
                     }
                   }`

	response, err := txn.QueryWithVars(ctx, q, map[string]string{"$account": name})
	if err != nil {
		return nil, err
	}

	var result struct {
		Account []*Account `json:"accounts"`
	}

	err = json.Unmarshal(response.Json, &result)
	if err != nil {
		return nil, err
	}

	if len(result.Account) == 0 {
		return nil, errors.New("Account not found")
	}

	account = &nodepb.Account{
		Uid:     result.Account[0].UID,
		Name:    result.Account[0].Name,
		IsRoot:  result.Account[0].IsRoot,
		Enabled: result.Account[0].Enabled,
	}

	if len(result.Account[0].DefaultNamespace) == 1 {
		account.DefaultNamespace = &nodepb.Namespace{
			Name: result.Account[0].DefaultNamespace[0].Name,
			Id:   result.Account[0].DefaultNamespace[0].UID,
		}
	}

	return account, err
}
