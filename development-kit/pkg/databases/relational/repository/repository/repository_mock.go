// Copyright 2020 ZUP IT SERVICOS EM TECNOLOGIA E INOVACAO SA
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package repository

import (
	SQL "github.com/ZupIT/horusec/development-kit/pkg/databases/relational"
	accountEntities "github.com/ZupIT/horusec/development-kit/pkg/entities/account"
	"github.com/ZupIT/horusec/development-kit/pkg/entities/roles"
	mockUtils "github.com/ZupIT/horusec/development-kit/pkg/utils/mock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type Mock struct {
	mock.Mock
}

func (m *Mock) Create(_ *accountEntities.Repository, _ SQL.InterfaceWrite) error {
	args := m.MethodCalled("Create")
	return mockUtils.ReturnNilOrError(args, 0)
}

func (m *Mock) Update(_ uuid.UUID, _ *accountEntities.Repository) (*accountEntities.Repository, error) {
	args := m.MethodCalled("Update")
	return args.Get(0).(*accountEntities.Repository), mockUtils.ReturnNilOrError(args, 1)
}

func (m *Mock) Get(_ uuid.UUID) (*accountEntities.Repository, error) {
	args := m.MethodCalled("Get")
	return args.Get(0).(*accountEntities.Repository), mockUtils.ReturnNilOrError(args, 1)
}

func (m *Mock) List(_, _ uuid.UUID) (*[]accountEntities.RepositoryResponse, error) {
	args := m.MethodCalled("List")
	return args.Get(0).(*[]accountEntities.RepositoryResponse), mockUtils.ReturnNilOrError(args, 1)
}

func (m *Mock) Delete(_ uuid.UUID) error {
	args := m.MethodCalled("Delete")
	return mockUtils.ReturnNilOrError(args, 0)
}

func (m *Mock) GetAllAccountsInRepository(_ uuid.UUID) (*[]roles.AccountRole, error) {
	args := m.MethodCalled("GetAllAccountsInRepository")
	return args.Get(0).(*[]roles.AccountRole), mockUtils.ReturnNilOrError(args, 1)
}

func (m *Mock) GetByName(_ uuid.UUID, _ string) (*accountEntities.Repository, error) {
	args := m.MethodCalled("GetByName")
	return args.Get(0).(*accountEntities.Repository), mockUtils.ReturnNilOrError(args, 1)
}

func (m *Mock) GetAccountCompanyRole(_, _ uuid.UUID) (*roles.AccountCompany, error) {
	args := m.MethodCalled("GetAccountCompanyRole")
	return args.Get(0).(*roles.AccountCompany), mockUtils.ReturnNilOrError(args, 1)
}

func (m *Mock) ListByLdapPermissions(_ uuid.UUID, _ []string) (*[]accountEntities.RepositoryResponse, error) {
	args := m.MethodCalled("ListByLdapPermissions")
	return args.Get(0).(*[]accountEntities.RepositoryResponse), mockUtils.ReturnNilOrError(args, 1)
}
