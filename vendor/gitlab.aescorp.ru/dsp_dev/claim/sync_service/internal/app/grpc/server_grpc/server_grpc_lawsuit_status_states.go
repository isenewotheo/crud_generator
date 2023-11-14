//File generated automatic with crud_generator app
//Do not change anything here.

package server_grpc

import (
	"context"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/network/grpc/grpc_proto"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/object_model/entities/lawsuit_status_states"
)

// LawsuitStatusState_Read - читает и возвращает модель из БД
func (s *ServerGRPC) LawsuitStatusState_Read(ctx context.Context, Request *grpc_proto.RequestId) (*grpc_proto.Response, error) {
	var Otvet grpc_proto.Response
	var err error

	// проверим совпадения версии модели
	VersionServer := lawsuit_status_states.LawsuitStatusState{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(lawsuit_status_states.LawsuitStatusState{})
		return &Otvet, err
	}

	// запрос в БД
	Model := &lawsuit_status_states.LawsuitStatusState{}
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

// LawsuitStatusState_Delete - записывает в БД is_deleted = true и возвращает модель из БД
func (s *ServerGRPC) LawsuitStatusState_Delete(ctx context.Context, Request *grpc_proto.RequestId) (*grpc_proto.Response, error) {
	var Otvet grpc_proto.Response
	var err error

	// проверим совпадения версии модели
	VersionServer := lawsuit_status_states.LawsuitStatusState{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(lawsuit_status_states.LawsuitStatusState{})
		return &Otvet, err
	}

	// запрос в БД
	Model := &lawsuit_status_states.LawsuitStatusState{}
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

// LawsuitStatusState_Restore - записывает в БД is_deleted = false и возвращает модель из БД
func (s *ServerGRPC) LawsuitStatusState_Restore(ctx context.Context, Request *grpc_proto.RequestId) (*grpc_proto.Response, error) {
	var Otvet grpc_proto.Response
	var err error

	// проверим совпадения версии модели
	VersionServer := lawsuit_status_states.LawsuitStatusState{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(lawsuit_status_states.LawsuitStatusState{})
		return &Otvet, err
	}

	// запрос в БД
	Model := &lawsuit_status_states.LawsuitStatusState{}
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

// LawsuitStatusState_Create - создаёт новую запись в БД
func (s *ServerGRPC) LawsuitStatusState_Create(ctx context.Context, Request *grpc_proto.RequestModel) (*grpc_proto.Response, error) {
	var Otvet grpc_proto.Response
	var err error

	// проверим совпадения версии модели
	VersionServer := lawsuit_status_states.LawsuitStatusState{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(lawsuit_status_states.LawsuitStatusState{})
		return &Otvet, err
	}

	// получим модель из строки JSON
	Model := &lawsuit_status_states.LawsuitStatusState{}
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

// LawsuitStatusState_Update - обновляет новую запись в БД
func (s *ServerGRPC) LawsuitStatusState_Update(ctx context.Context, Request *grpc_proto.RequestModel) (*grpc_proto.Response, error) {
	var Otvet grpc_proto.Response
	var err error

	// проверим совпадения версии модели
	VersionServer := lawsuit_status_states.LawsuitStatusState{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(lawsuit_status_states.LawsuitStatusState{})
		return &Otvet, err
	}

	// получим модель из строки JSON
	Model := &lawsuit_status_states.LawsuitStatusState{}
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

// LawsuitStatusState_Save - записывает (создаёт или обновляет) запись в БД
func (s *ServerGRPC) LawsuitStatusState_Save(ctx context.Context, Request *grpc_proto.RequestModel) (*grpc_proto.Response, error) {
	var Otvet grpc_proto.Response
	var err error

	// проверим совпадения версии модели
	VersionServer := lawsuit_status_states.LawsuitStatusState{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(lawsuit_status_states.LawsuitStatusState{})
		return &Otvet, err
	}

	// получим модель из строки JSON
	Model := lawsuit_status_states.LawsuitStatusState{}
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
