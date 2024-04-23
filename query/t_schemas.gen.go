// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"hs_short_link/model"
)

func newTSchema(db *gorm.DB, opts ...gen.DOOption) tSchema {
	_tSchema := tSchema{}

	_tSchema.tSchemaDo.UseDB(db, opts...)
	_tSchema.tSchemaDo.UseModel(&model.TSchema{})

	tableName := _tSchema.tSchemaDo.TableName()
	_tSchema.ALL = field.NewAsterisk(tableName)
	_tSchema.ID = field.NewInt64(tableName, "id")
	_tSchema.No = field.NewString(tableName, "no")
	_tSchema.AllowRegion = field.NewString(tableName, "allow_region")
	_tSchema.BanProxy = field.NewBool(tableName, "ban_proxy")
	_tSchema.DestLink = field.NewString(tableName, "dest_link")
	_tSchema.BanLink = field.NewString(tableName, "ban_link")

	_tSchema.fillFieldMap()

	return _tSchema
}

type tSchema struct {
	tSchemaDo

	ALL         field.Asterisk
	ID          field.Int64
	No          field.String
	AllowRegion field.String
	BanProxy    field.Bool
	DestLink    field.String
	BanLink     field.String

	fieldMap map[string]field.Expr
}

func (t tSchema) Table(newTableName string) *tSchema {
	t.tSchemaDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t tSchema) As(alias string) *tSchema {
	t.tSchemaDo.DO = *(t.tSchemaDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *tSchema) updateTableName(table string) *tSchema {
	t.ALL = field.NewAsterisk(table)
	t.ID = field.NewInt64(table, "id")
	t.No = field.NewString(table, "no")
	t.AllowRegion = field.NewString(table, "allow_region")
	t.BanProxy = field.NewBool(table, "ban_proxy")
	t.DestLink = field.NewString(table, "dest_link")
	t.BanLink = field.NewString(table, "ban_link")

	t.fillFieldMap()

	return t
}

func (t *tSchema) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *tSchema) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 6)
	t.fieldMap["id"] = t.ID
	t.fieldMap["no"] = t.No
	t.fieldMap["allow_region"] = t.AllowRegion
	t.fieldMap["ban_proxy"] = t.BanProxy
	t.fieldMap["dest_link"] = t.DestLink
	t.fieldMap["ban_link"] = t.BanLink
}

func (t tSchema) clone(db *gorm.DB) tSchema {
	t.tSchemaDo.ReplaceConnPool(db.Statement.ConnPool)
	return t
}

func (t tSchema) replaceDB(db *gorm.DB) tSchema {
	t.tSchemaDo.ReplaceDB(db)
	return t
}

type tSchemaDo struct{ gen.DO }

type ITSchemaDo interface {
	gen.SubQuery
	Debug() ITSchemaDo
	WithContext(ctx context.Context) ITSchemaDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ITSchemaDo
	WriteDB() ITSchemaDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ITSchemaDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ITSchemaDo
	Not(conds ...gen.Condition) ITSchemaDo
	Or(conds ...gen.Condition) ITSchemaDo
	Select(conds ...field.Expr) ITSchemaDo
	Where(conds ...gen.Condition) ITSchemaDo
	Order(conds ...field.Expr) ITSchemaDo
	Distinct(cols ...field.Expr) ITSchemaDo
	Omit(cols ...field.Expr) ITSchemaDo
	Join(table schema.Tabler, on ...field.Expr) ITSchemaDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ITSchemaDo
	RightJoin(table schema.Tabler, on ...field.Expr) ITSchemaDo
	Group(cols ...field.Expr) ITSchemaDo
	Having(conds ...gen.Condition) ITSchemaDo
	Limit(limit int) ITSchemaDo
	Offset(offset int) ITSchemaDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ITSchemaDo
	Unscoped() ITSchemaDo
	Create(values ...*model.TSchema) error
	CreateInBatches(values []*model.TSchema, batchSize int) error
	Save(values ...*model.TSchema) error
	First() (*model.TSchema, error)
	Take() (*model.TSchema, error)
	Last() (*model.TSchema, error)
	Find() ([]*model.TSchema, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TSchema, err error)
	FindInBatches(result *[]*model.TSchema, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.TSchema) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ITSchemaDo
	Assign(attrs ...field.AssignExpr) ITSchemaDo
	Joins(fields ...field.RelationField) ITSchemaDo
	Preload(fields ...field.RelationField) ITSchemaDo
	FirstOrInit() (*model.TSchema, error)
	FirstOrCreate() (*model.TSchema, error)
	FindByPage(offset int, limit int) (result []*model.TSchema, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ITSchemaDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (t tSchemaDo) Debug() ITSchemaDo {
	return t.withDO(t.DO.Debug())
}

func (t tSchemaDo) WithContext(ctx context.Context) ITSchemaDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t tSchemaDo) ReadDB() ITSchemaDo {
	return t.Clauses(dbresolver.Read)
}

func (t tSchemaDo) WriteDB() ITSchemaDo {
	return t.Clauses(dbresolver.Write)
}

func (t tSchemaDo) Session(config *gorm.Session) ITSchemaDo {
	return t.withDO(t.DO.Session(config))
}

func (t tSchemaDo) Clauses(conds ...clause.Expression) ITSchemaDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t tSchemaDo) Returning(value interface{}, columns ...string) ITSchemaDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t tSchemaDo) Not(conds ...gen.Condition) ITSchemaDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t tSchemaDo) Or(conds ...gen.Condition) ITSchemaDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t tSchemaDo) Select(conds ...field.Expr) ITSchemaDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t tSchemaDo) Where(conds ...gen.Condition) ITSchemaDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t tSchemaDo) Order(conds ...field.Expr) ITSchemaDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t tSchemaDo) Distinct(cols ...field.Expr) ITSchemaDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t tSchemaDo) Omit(cols ...field.Expr) ITSchemaDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t tSchemaDo) Join(table schema.Tabler, on ...field.Expr) ITSchemaDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t tSchemaDo) LeftJoin(table schema.Tabler, on ...field.Expr) ITSchemaDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t tSchemaDo) RightJoin(table schema.Tabler, on ...field.Expr) ITSchemaDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t tSchemaDo) Group(cols ...field.Expr) ITSchemaDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t tSchemaDo) Having(conds ...gen.Condition) ITSchemaDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t tSchemaDo) Limit(limit int) ITSchemaDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t tSchemaDo) Offset(offset int) ITSchemaDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t tSchemaDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ITSchemaDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t tSchemaDo) Unscoped() ITSchemaDo {
	return t.withDO(t.DO.Unscoped())
}

func (t tSchemaDo) Create(values ...*model.TSchema) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t tSchemaDo) CreateInBatches(values []*model.TSchema, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t tSchemaDo) Save(values ...*model.TSchema) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t tSchemaDo) First() (*model.TSchema, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.TSchema), nil
	}
}

func (t tSchemaDo) Take() (*model.TSchema, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.TSchema), nil
	}
}

func (t tSchemaDo) Last() (*model.TSchema, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.TSchema), nil
	}
}

func (t tSchemaDo) Find() ([]*model.TSchema, error) {
	result, err := t.DO.Find()
	return result.([]*model.TSchema), err
}

func (t tSchemaDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TSchema, err error) {
	buf := make([]*model.TSchema, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t tSchemaDo) FindInBatches(result *[]*model.TSchema, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t tSchemaDo) Attrs(attrs ...field.AssignExpr) ITSchemaDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t tSchemaDo) Assign(attrs ...field.AssignExpr) ITSchemaDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t tSchemaDo) Joins(fields ...field.RelationField) ITSchemaDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Joins(_f))
	}
	return &t
}

func (t tSchemaDo) Preload(fields ...field.RelationField) ITSchemaDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Preload(_f))
	}
	return &t
}

func (t tSchemaDo) FirstOrInit() (*model.TSchema, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.TSchema), nil
	}
}

func (t tSchemaDo) FirstOrCreate() (*model.TSchema, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.TSchema), nil
	}
}

func (t tSchemaDo) FindByPage(offset int, limit int) (result []*model.TSchema, count int64, err error) {
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

func (t tSchemaDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t tSchemaDo) Scan(result interface{}) (err error) {
	return t.DO.Scan(result)
}

func (t tSchemaDo) Delete(models ...*model.TSchema) (result gen.ResultInfo, err error) {
	return t.DO.Delete(models)
}

func (t *tSchemaDo) withDO(do gen.Dao) *tSchemaDo {
	t.DO = *do.(*gen.DO)
	return t
}
