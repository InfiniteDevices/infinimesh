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

package main

import (
	"context"

	jwt "github.com/dgrijalva/jwt-go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/golang/protobuf/ptypes/empty"

	"github.com/infinimesh/infinimesh/pkg/apiserver/apipb"
	"github.com/infinimesh/infinimesh/pkg/node/nodepb"
)

type accountAPI struct {
	signingSecret []byte
	client        nodepb.AccountServiceClient
}

func (a *accountAPI) SelfAccount(ctx context.Context, request *empty.Empty) (response *nodepb.Account, err error) {
	account, ok := ctx.Value("account_id").(string)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "Unauthenticated")
	}

	return a.client.GetAccount(ctx, &nodepb.GetAccountRequest{
		Id: account,
	})
}

func (a *accountAPI) GetAccount(ctx context.Context, request *nodepb.GetAccountRequest) (response *nodepb.Account, err error) {
	account, ok := ctx.Value("account_id").(string)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "Unauthenticated")
	}

	if res, err := a.client.IsRoot(ctx, &nodepb.IsRootRequest{
		Account: account,
	}); err == nil && res.GetIsRoot() {
		return a.client.GetAccount(ctx, request)
	}
	return &nodepb.Account{}, status.Error(codes.PermissionDenied, "Insufficient permissions")

}

func (a *accountAPI) Token(ctx context.Context, request *apipb.TokenRequest) (response *apipb.TokenResponse, err error) {
	resp, err := a.client.Authenticate(ctx, &nodepb.AuthenticateRequest{Username: request.GetUsername(), Password: request.GetPassword()})
	if err != nil {
		return nil, err
	}

	if resp.GetSuccess() {
		if resp.Account == nil {
			return nil, status.Error(codes.Internal, "Failed to check credentials")
		}

		claim := jwt.MapClaims{}
		claim[accountIDClaim] = resp.Account.Uid
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

		// Sign and get the complete encoded token as a string using the secret
		tokenString, err := token.SignedString(a.signingSecret)
		if err != nil {
			return nil, status.Error(codes.Internal, "Failed to sign token")
		}

		return &apipb.TokenResponse{Token: tokenString}, nil
	}

	return nil, status.Error(codes.Unauthenticated, "Invalid credentials")
}

func (a *accountAPI) UpdateAccount(ctx context.Context, request *nodepb.UpdateAccountRequest) (response *nodepb.Account, err error) {
	account, ok := ctx.Value("account_id").(string)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "Unauthenticated")
	}

	if res, err := a.client.IsRoot(ctx, &nodepb.IsRootRequest{
		Account: account,
	}); err == nil && res.GetIsRoot() {
		res, err := a.client.UpdateAccount(ctx, request)
		return res, err
	}
	return &nodepb.Account{}, status.Error(codes.PermissionDenied, "Insufficient permissions")
}

func (a *accountAPI) CreateUserAccount(ctx context.Context, request *nodepb.CreateUserAccountRequest) (response *nodepb.CreateUserAccountResponse, err error) {
	account, ok := ctx.Value("account_id").(string)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "Unauthenticated")
	}

	if res, err := a.client.IsRoot(ctx, &nodepb.IsRootRequest{
		Account: account,
	}); err == nil && res.GetIsRoot() {
		res, err := a.client.CreateUserAccount(ctx, request)
		return res, err
	}

	return &nodepb.CreateUserAccountResponse{}, status.Error(codes.PermissionDenied, "Insufficient permissions")
}

func (a *accountAPI) AuthFuncOverride(ctx context.Context, fullMethodName string) (context.Context, error) {
	if fullMethodName != "/infinimesh.api.Accounts/Token" {
		return jwtAuth(ctx)
	}
	return ctx, nil
}

func (a *accountAPI) ListAccounts(ctx context.Context, request *nodepb.ListAccountsRequest) (response *nodepb.ListAccountsResponse, err error) {
	account, ok := ctx.Value("account_id").(string)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "Unauthenticated")
	}

	if res, err := a.client.IsRoot(ctx, &nodepb.IsRootRequest{
		Account: account,
	}); err == nil && res.GetIsRoot() {
		res, err := a.client.ListAccounts(ctx, request)
		return res, err
	}

	return &nodepb.ListAccountsResponse{}, status.Error(codes.PermissionDenied, "Insufficient permissions")

}
