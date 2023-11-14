//File generated automatic with crud_generator app
//Do not change anything here.

package server_grpc

import (
	"context"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/network/grpc/grpc_proto"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/object_model/entities/bank_account_organizations"
)

// BankAccountOrganization_Read - читает и возвращает модель из БД
func (s *ServerGRPC) BankAccountOrganization_Read(ctx context.Context, Request *grpc_proto.RequestId) (*grpc_proto.Response, error) {
	var Otvet grpc_proto.Response
	var err error

	// проверим совпадения версии модели
	VersionServer := bank_account_organizations.BankAccountOrganization{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(bank_account_organizations.BankAccountOrganization{})
		return &Otvet, err
	}

	// запрос в БД
	Model := &bank_account_organizations.BankAccountOrganization{}
	Model.ID = Request.Id
	err = Model.Read()
	if err != nil {
		return &Otvet, err
	}

	// заполяем ответ
	ModelString, err := Model.GetJSON()
	if err != nil {
		return &Otvet, err
	}
	Otvet.ModelString = ModelString

	return &Otvet, err
}

// BankAccountOrganization_Delete - записывает в БД is_deleted = true и возвращает модель из БД
func (s *ServerGRPC) BankAccountOrganization_Delete(ctx context.Context, Request *grpc_proto.RequestId) (*grpc_proto.Response, error) {
	var Otvet grpc_proto.Response
	var err error

	// проверим совпадения версии модели
	VersionServer := bank_account_organizations.BankAccountOrganization{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(bank_account_organizations.BankAccountOrganization{})
		return &Otvet, err
	}

	// запрос в БД
	Model := &bank_account_organizations.BankAccountOrganization{}
	Model.ID = Request.Id
	err = Model.Delete()
	if err != nil {
		return &Otvet, err
	}

	// заполяем ответ
	ModelString, err := Model.GetJSON()
	if err != nil {
		return &Otvet, err
	}
	Otvet.ModelString = ModelString

	return &Otvet, err
}

// BankAccountOrganization_Restore - записывает в БД is_deleted = false и возвращает модель из БД
func (s *ServerGRPC) BankAccountOrganization_Restore(ctx context.Context, Request *grpc_proto.RequestId) (*grpc_proto.Response, error) {
	var Otvet grpc_proto.Response
	var err error

	// проверим совпадения версии модели
	VersionServer := bank_account_organizations.BankAccountOrganization{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(bank_account_organizations.BankAccountOrganization{})
		return &Otvet, err
	}

	// запрос в БД
	Model := &bank_account_organizations.BankAccountOrganization{}
	Model.ID = Request.Id
	err = Model.Restore()
	if err != nil {
		return &Otvet, err
	}

	// заполяем ответ
	ModelString, err := Model.GetJSON()
	if err != nil {
		return &Otvet, err
	}
	Otvet.ModelString = ModelString

	return &Otvet, err
}

// BankAccountOrganization_Create - создаёт новую запись в БД
func (s *ServerGRPC) BankAccountOrganization_Create(ctx context.Context, Request *grpc_proto.RequestModel) (*grpc_proto.Response, error) {
	var Otvet grpc_proto.Response
	var err error

	// проверим совпадения версии модели
	VersionServer := bank_account_organizations.BankAccountOrganization{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(bank_account_organizations.BankAccountOrganization{})
		return &Otvet, err
	}

	// получим модель из строки JSON
	Model := &bank_account_organizations.BankAccountOrganization{}
	err = Model.GetModelFromJSON(Request.ModelString)
	if err != nil {
		return &Otvet, err
	}

	// запрос в БД
	err = Model.Create()
	if err != nil {
		return &Otvet, err
	}

	// заполяем ответ
	ModelString, err := Model.GetJSON()
	if err != nil {
		return &Otvet, err
	}
	Otvet.ModelString = ModelString

	return &Otvet, err
}

// BankAccountOrganization_Update - обновляет новую запись в БД
func (s *ServerGRPC) BankAccountOrganization_Update(ctx context.Context, Request *grpc_proto.RequestModel) (*grpc_proto.Response, error) {
	var Otvet grpc_proto.Response
	var err error

	// проверим совпадения версии модели
	VersionServer := bank_account_organizations.BankAccountOrganization{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(bank_account_organizations.BankAccountOrganization{})
		return &Otvet, err
	}

	// получим модель из строки JSON
	Model := &bank_account_organizations.BankAccountOrganization{}
	err = Model.GetModelFromJSON(Request.ModelString)
	if err != nil {
		return &Otvet, err
	}

	// запрос в БД
	err = Model.Update()
	if err != nil {
		return &Otvet, err
	}

	// заполяем ответ
	ModelString, err := Model.GetJSON()
	if err != nil {
		return &Otvet, err
	}
	Otvet.ModelString = ModelString

	return &Otvet, err
}

// BankAccountOrganization_Save - записывает (создаёт или обновляет) запись в БД
func (s *ServerGRPC) BankAccountOrganization_Save(ctx context.Context, Request *grpc_proto.RequestModel) (*grpc_proto.Response, error) {
	var Otvet grpc_proto.Response
	var err error

	// проверим совпадения версии модели
	VersionServer := bank_account_organizations.BankAccountOrganization{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(bank_account_organizations.BankAccountOrganization{})
		return &Otvet, err
	}

	// получим модель из строки JSON
	Model := bank_account_organizations.BankAccountOrganization{}
	err = Model.GetModelFromJSON(Request.ModelString)
	if err != nil {
		return &Otvet, err
	}

	// запрос в БД
	err = Model.Save()
	if err != nil {
		return &Otvet, err
	}

	// заполяем ответ
	ModelString, err := Model.GetJSON()
	if err != nil {
		return &Otvet, err
	}
	Otvet.ModelString = ModelString

	return &Otvet, err
}

// BankAccountOrganization_FindByExtID - возвращает запись из БД по ext_id и connection_id
func (s *ServerGRPC) BankAccountOrganization_FindByExtID(ctx context.Context, Request *grpc_proto.RequestExtId) (*grpc_proto.Response, error) {
	var Otvet grpc_proto.Response
	var err error

	//проверим совпадения версии модели
	VersionServer := bank_account_organizations.BankAccountOrganization{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(bank_account_organizations.BankAccountOrganization{})
		return &Otvet, err
	}

	//запрос в БД
	Model := &bank_account_organizations.BankAccountOrganization{}
	Model.ExtID = Request.ExtId
	Model.ConnectionID = Request.ConnectionId
	err = Model.Find_ByExtID()
	if err != nil {
		return &Otvet, err
	}

	//заполяем ответ
	ModelString, err := Model.GetJSON()
	if err != nil {
		return &Otvet, err
	}
	Otvet.ModelString = ModelString

	return &Otvet, err
}
