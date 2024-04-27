// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"hs-api/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"
)

func newTClientInfo(db *gorm.DB, opts ...gen.DOOption) tClientInfo {
	_tClientInfo := tClientInfo{}

	_tClientInfo.tClientInfoDo.UseDB(db, opts...)
	_tClientInfo.tClientInfoDo.UseModel(&model.TClientInfo{})

	tableName := _tClientInfo.tClientInfoDo.TableName()
	_tClientInfo.ALL = field.NewAsterisk(tableName)
	_tClientInfo.ID = field.NewInt64(tableName, "id")
	_tClientInfo.IP = field.NewString(tableName, "ip")
	_tClientInfo.No = field.NewString(tableName, "no")
	_tClientInfo.Ua = field.NewString(tableName, "ua")
	_tClientInfo.Lang = field.NewString(tableName, "lang")
	_tClientInfo.Referer = field.NewString(tableName, "referer")
	_tClientInfo.VisitTime = field.NewString(tableName, "visit_time")
	_tClientInfo.Status = field.NewInt64(tableName, "status")
	_tClientInfo.Isp = field.NewString(tableName, "isp")
	_tClientInfo.Organization = field.NewString(tableName, "organization")
	_tClientInfo.Country = field.NewString(tableName, "country")
	_tClientInfo.Longitude = field.NewFloat64(tableName, "longitude")
	_tClientInfo.Latitude = field.NewFloat64(tableName, "latitude")
	_tClientInfo.ContinentCode = field.NewString(tableName, "continent_code")
	_tClientInfo.CountryCode = field.NewString(tableName, "country_code")
	_tClientInfo.CloakResponse = field.NewString(tableName, "cloak_response")

	_tClientInfo.fillFieldMap()

	return _tClientInfo
}

type tClientInfo struct {
	tClientInfoDo

	ALL           field.Asterisk
	ID            field.Int64
	IP            field.String
	No            field.String
	Ua            field.String
	Lang          field.String
	Referer       field.String
	VisitTime     field.String
	Status        field.Int64
	Isp           field.String
	Organization  field.String
	Country       field.String
	Longitude     field.Float64
	Latitude      field.Float64
	ContinentCode field.String
	CountryCode   field.String
	CloakResponse field.String

	fieldMap map[string]field.Expr
}

func (t tClientInfo) Table(newTableName string) *tClientInfo {
	t.tClientInfoDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t tClientInfo) As(alias string) *tClientInfo {
	t.tClientInfoDo.DO = *(t.tClientInfoDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *tClientInfo) updateTableName(table string) *tClientInfo {
	t.ALL = field.NewAsterisk(table)
	t.ID = field.NewInt64(table, "id")
	t.IP = field.NewString(table, "ip")
	t.No = field.NewString(table, "no")
	t.Ua = field.NewString(table, "ua")
	t.Lang = field.NewString(table, "lang")
	t.Referer = field.NewString(table, "referer")
	t.VisitTime = field.NewString(table, "visit_time")
	t.Status = field.NewInt64(table, "status")
	t.Isp = field.NewString(table, "isp")
	t.Organization = field.NewString(table, "organization")
	t.Country = field.NewString(table, "country")
	t.Longitude = field.NewFloat64(table, "longitude")
	t.Latitude = field.NewFloat64(table, "latitude")
	t.ContinentCode = field.NewString(table, "continent_code")
	t.CountryCode = field.NewString(table, "country_code")
	t.CloakResponse = field.NewString(table, "cloak_response")

	t.fillFieldMap()

	return t
}

func (t *tClientInfo) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *tClientInfo) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 16)
	t.fieldMap["id"] = t.ID
	t.fieldMap["ip"] = t.IP
	t.fieldMap["no"] = t.No
	t.fieldMap["ua"] = t.Ua
	t.fieldMap["lang"] = t.Lang
	t.fieldMap["referer"] = t.Referer
	t.fieldMap["visit_time"] = t.VisitTime
	t.fieldMap["status"] = t.Status
	t.fieldMap["isp"] = t.Isp
	t.fieldMap["organization"] = t.Organization
	t.fieldMap["country"] = t.Country
	t.fieldMap["longitude"] = t.Longitude
	t.fieldMap["latitude"] = t.Latitude
	t.fieldMap["continent_code"] = t.ContinentCode
	t.fieldMap["country_code"] = t.CountryCode
	t.fieldMap["cloak_response"] = t.CloakResponse
}

func (t tClientInfo) clone(db *gorm.DB) tClientInfo {
	t.tClientInfoDo.ReplaceConnPool(db.Statement.ConnPool)
	return t
}

func (t tClientInfo) replaceDB(db *gorm.DB) tClientInfo {
	t.tClientInfoDo.ReplaceDB(db)
	return t
}

type tClientInfoDo struct{ gen.DO }

type ITClientInfoDo interface {
	gen.SubQuery
	Debug() ITClientInfoDo
	WithContext(ctx context.Context) ITClientInfoDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ITClientInfoDo
	WriteDB() ITClientInfoDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ITClientInfoDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ITClientInfoDo
	Not(conds ...gen.Condition) ITClientInfoDo
	Or(conds ...gen.Condition) ITClientInfoDo
	Select(conds ...field.Expr) ITClientInfoDo
	Where(conds ...gen.Condition) ITClientInfoDo
	Order(conds ...field.Expr) ITClientInfoDo
	Distinct(cols ...field.Expr) ITClientInfoDo
	Omit(cols ...field.Expr) ITClientInfoDo
	Join(table schema.Tabler, on ...field.Expr) ITClientInfoDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ITClientInfoDo
	RightJoin(table schema.Tabler, on ...field.Expr) ITClientInfoDo
	Group(cols ...field.Expr) ITClientInfoDo
	Having(conds ...gen.Condition) ITClientInfoDo
	Limit(limit int) ITClientInfoDo
	Offset(offset int) ITClientInfoDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ITClientInfoDo
	Unscoped() ITClientInfoDo
	Create(values ...*model.TClientInfo) error
	CreateInBatches(values []*model.TClientInfo, batchSize int) error
	Save(values ...*model.TClientInfo) error
	First() (*model.TClientInfo, error)
	Take() (*model.TClientInfo, error)
	Last() (*model.TClientInfo, error)
	Find() ([]*model.TClientInfo, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TClientInfo, err error)
	FindInBatches(result *[]*model.TClientInfo, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.TClientInfo) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ITClientInfoDo
	Assign(attrs ...field.AssignExpr) ITClientInfoDo
	Joins(fields ...field.RelationField) ITClientInfoDo
	Preload(fields ...field.RelationField) ITClientInfoDo
	FirstOrInit() (*model.TClientInfo, error)
	FirstOrCreate() (*model.TClientInfo, error)
	FindByPage(offset int, limit int) (result []*model.TClientInfo, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ITClientInfoDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (t tClientInfoDo) Debug() ITClientInfoDo {
	return t.withDO(t.DO.Debug())
}

func (t tClientInfoDo) WithContext(ctx context.Context) ITClientInfoDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t tClientInfoDo) ReadDB() ITClientInfoDo {
	return t.Clauses(dbresolver.Read)
}

func (t tClientInfoDo) WriteDB() ITClientInfoDo {
	return t.Clauses(dbresolver.Write)
}

func (t tClientInfoDo) Session(config *gorm.Session) ITClientInfoDo {
	return t.withDO(t.DO.Session(config))
}

func (t tClientInfoDo) Clauses(conds ...clause.Expression) ITClientInfoDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t tClientInfoDo) Returning(value interface{}, columns ...string) ITClientInfoDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t tClientInfoDo) Not(conds ...gen.Condition) ITClientInfoDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t tClientInfoDo) Or(conds ...gen.Condition) ITClientInfoDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t tClientInfoDo) Select(conds ...field.Expr) ITClientInfoDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t tClientInfoDo) Where(conds ...gen.Condition) ITClientInfoDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t tClientInfoDo) Order(conds ...field.Expr) ITClientInfoDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t tClientInfoDo) Distinct(cols ...field.Expr) ITClientInfoDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t tClientInfoDo) Omit(cols ...field.Expr) ITClientInfoDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t tClientInfoDo) Join(table schema.Tabler, on ...field.Expr) ITClientInfoDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t tClientInfoDo) LeftJoin(table schema.Tabler, on ...field.Expr) ITClientInfoDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t tClientInfoDo) RightJoin(table schema.Tabler, on ...field.Expr) ITClientInfoDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t tClientInfoDo) Group(cols ...field.Expr) ITClientInfoDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t tClientInfoDo) Having(conds ...gen.Condition) ITClientInfoDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t tClientInfoDo) Limit(limit int) ITClientInfoDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t tClientInfoDo) Offset(offset int) ITClientInfoDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t tClientInfoDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ITClientInfoDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t tClientInfoDo) Unscoped() ITClientInfoDo {
	return t.withDO(t.DO.Unscoped())
}

func (t tClientInfoDo) Create(values ...*model.TClientInfo) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t tClientInfoDo) CreateInBatches(values []*model.TClientInfo, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t tClientInfoDo) Save(values ...*model.TClientInfo) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t tClientInfoDo) First() (*model.TClientInfo, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.TClientInfo), nil
	}
}

func (t tClientInfoDo) Take() (*model.TClientInfo, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.TClientInfo), nil
	}
}

func (t tClientInfoDo) Last() (*model.TClientInfo, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.TClientInfo), nil
	}
}

func (t tClientInfoDo) Find() ([]*model.TClientInfo, error) {
	result, err := t.DO.Find()
	return result.([]*model.TClientInfo), err
}

func (t tClientInfoDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TClientInfo, err error) {
	buf := make([]*model.TClientInfo, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t tClientInfoDo) FindInBatches(result *[]*model.TClientInfo, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t tClientInfoDo) Attrs(attrs ...field.AssignExpr) ITClientInfoDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t tClientInfoDo) Assign(attrs ...field.AssignExpr) ITClientInfoDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t tClientInfoDo) Joins(fields ...field.RelationField) ITClientInfoDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Joins(_f))
	}
	return &t
}

func (t tClientInfoDo) Preload(fields ...field.RelationField) ITClientInfoDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Preload(_f))
	}
	return &t
}

func (t tClientInfoDo) FirstOrInit() (*model.TClientInfo, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.TClientInfo), nil
	}
}

func (t tClientInfoDo) FirstOrCreate() (*model.TClientInfo, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.TClientInfo), nil
	}
}

func (t tClientInfoDo) FindByPage(offset int, limit int) (result []*model.TClientInfo, count int64, err error) {
	result, err = t.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = t.Offset(-1).Limit(-1).Count()
	return
}

func (t tClientInfoDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t tClientInfoDo) Scan(result interface{}) (err error) {
	return t.DO.Scan(result)
}

func (t tClientInfoDo) Delete(models ...*model.TClientInfo) (result gen.ResultInfo, err error) {
	return t.DO.Delete(models)
}

func (t *tClientInfoDo) withDO(do gen.Dao) *tClientInfoDo {
	t.DO = *do.(*gen.DO)
	return t
}
