package db_lawsuit_status_types

import (
	model "gitlab.aescorp.ru/dsp_dev/claim/common/object_model"
	"gitlab.aescorp.ru/dsp_dev/claim/nikitin/config"
	"gitlab.aescorp.ru/dsp_dev/claim/nikitin/postgres_gorm"
	"testing"
)

const Postgres_ID_Test = 1

func TestRead(t *testing.T) {
	config.LoadEnv()

	postgres_gorm.Connect()
	defer postgres_gorm.CloseConnection()

	crud := Crud_DB{}
	Otvet := model.LawsuitStatusType{}
	Otvet.ID = Postgres_ID_Test
	err := crud.Read(&Otvet)
	if err != nil {
		t.Error("TestRead() error: ", err)
	}

	if Otvet.Name == "" {
		t.Error(TableName + "_test.TestRead() error name= '' ")
	} else {
		t.Log(TableName+"_test.TestRead() Otvet: ", Otvet.Name)
	}
}

func TestSave(t *testing.T) {
	config.LoadEnv()

	postgres_gorm.Connect()
	defer postgres_gorm.CloseConnection()

	crud := Crud_DB{}
	Otvet := model.LawsuitStatusType{}
	Otvet.ID = Postgres_ID_Test
	err := crud.Read(&Otvet)
	if err != nil {
		t.Error("TestSave() error: ", err)
	}

	if Otvet.Name == "" {
		t.Error(TableName + "_test.TestSave() error name= '' ")
	}

	err = crud.Save(&Otvet)
	if err != nil {
		t.Error("TestSave() error: ", err)
	}
	t.Log(TableName+"_test.TestSave() Otvet: ", Otvet.Name)

}

func TestDelete(t *testing.T) {
	config.LoadEnv()

	postgres_gorm.Connect()
	defer postgres_gorm.CloseConnection()

	crud := Crud_DB{}
	Otvet := model.LawsuitStatusType{}
	Otvet.ID = Postgres_ID_Test
	err := crud.Read(&Otvet)
	if err != nil {
		t.Error("TestDelete() error: ", err)
	}

	if Otvet.IsDeleted == false {
		err = crud.Delete(&Otvet)
		if err != nil {
			t.Error("TestDelete() error: ", err)
		}

		err = crud.Restore(&Otvet)
		if err != nil {
			t.Error("TestDelete() error: ", err)
		}
	} else {
		err = crud.Restore(&Otvet)
		if err != nil {
			t.Error("TestDelete() error: ", err)
		}

		err = crud.Delete(&Otvet)
		if err != nil {
			t.Error("TestDelete() error: ", err)
		}

	}

}

//func TestFind_ByExtID(t *testing.T) {
//	config.LoadEnv()
//	postgres_gorm.Connect()
//	defer postgres_gorm.CloseConnection()
//
//	Otvet, err := Find_ByExtID(1, constants.CONNECTION_ID_TEST)
//	if err != nil {
//		t.Error("TestFind_ByExtID() error: ", err)
//	}
//
//	if Otvet.ID == 0 {
//		t.Error("TestFind_ByExtID() error: ID =0")
//	}
//}
