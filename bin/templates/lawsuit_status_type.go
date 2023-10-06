package object_model

import (
	"encoding/json"
	"github.com/vmihailenco/msgpack/v5"
	"reflect"
)

// versionLawsuitStatusType - версия структуры модели, с учётом имен и типов полей
var versionLawsuitStatusType uint32

// crud_LawsuitStatusType - объект контроллер crud операций
var crud_LawsuitStatusType ICrud_LawsuitStatusType

// LawsuitStatusType Статусы дел (справочник).
type LawsuitStatusType struct {
	CommonStruct
	NameStruct
	Code string `json:"code" gorm:"column:code;default:0"`
}

type ICrud_LawsuitStatusType interface {
	Read(l *LawsuitStatusType) error
	Save(l *LawsuitStatusType) error
	Update(l *LawsuitStatusType) error
	Create(l *LawsuitStatusType) error
	Delete(l *LawsuitStatusType) error
	Restore(l *LawsuitStatusType) error
}

// TableName - возвращает имя таблицы в БД, нужен для gorm
func (l LawsuitStatusType) TableName() string {
	return "lawsuit_status_types"
}

// NewLawsuitStatusType - возвращает новый	объект
func NewLawsuitStatusType() LawsuitStatusType {
	return LawsuitStatusType{}
}

// AsLawsuitStatusType - создаёт объект из упакованного объекта в массиве байтов
func AsLawsuitStatusType(b []byte) (LawsuitStatusType, error) {
	c := NewLawsuitStatusType()
	err := msgpack.Unmarshal(b, &c)
	if err != nil {
		return NewLawsuitStatusType(), err
	}
	return c, nil
}

// LawsuitStatusTypeAsBytes - упаковывает объект в массив байтов
func LawsuitStatusTypeAsBytes(l *LawsuitStatusType) ([]byte, error) {
	b, err := msgpack.Marshal(l)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// GetStructVersion - возвращает версию модели
func (l LawsuitStatusType) GetStructVersion() uint32 {
	if versionLawsuitStatusType == 0 {
		versionLawsuitStatusType = CalcStructVersion(reflect.TypeOf(l))
	}

	return versionLawsuitStatusType
}

// GetModelFromJSON - создаёт модель из строки json
func (l *LawsuitStatusType) GetModelFromJSON(sModel string) error {
	var err error

	var bytes []byte
	bytes = []byte(sModel)

	err = json.Unmarshal(bytes, l)

	return err
}

// GetJSON - возвращает строку json из модели
func (l LawsuitStatusType) GetJSON() (string, error) {
	var ReturnVar string
	var err error

	bytes, err := json.Marshal(l)
	if err != nil {
		return ReturnVar, err
	}
	ReturnVar = string(bytes)
	return ReturnVar, err
}

//---------------------------- CRUD операции ------------------------------------------------------------

// Read - находит запись в БД по ID, и заполняет в объект
func (l *LawsuitStatusType) Read() error {
	err := crud_LawsuitStatusType.Read(l)

	return err
}

// Save - записывает объект в БД по ID
func (l *LawsuitStatusType) Save() error {
	err := crud_LawsuitStatusType.Save(l)

	return err
}

// Update - обновляет объект в БД по ID
func (l *LawsuitStatusType) Update() error {
	err := crud_LawsuitStatusType.Update(l)

	return err
}

// Create - создаёт объект в БД с новым ID
func (l *LawsuitStatusType) Create() error {
	err := crud_LawsuitStatusType.Create(l)

	return err
}

// Delete - устанавливает признак пометки удаления в БД
func (l *LawsuitStatusType) Delete() error {
	err := crud_LawsuitStatusType.Delete(l)

	return err
}

// Restore - снимает признак пометки удаления в БД
func (l *LawsuitStatusType) Restore() error {
	err := crud_LawsuitStatusType.Restore(l)

	return err
}

// SetCrudInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (c LawsuitStatusType) SetCrudInterface(crud ICrud_LawsuitStatusType) {
	crud_LawsuitStatusType = crud

	return
}

//---------------------------- конец CRUD операции ------------------------------------------------------------
