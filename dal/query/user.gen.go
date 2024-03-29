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

	"gendemo/dal/model"
)

func newJustUser(db *gorm.DB, opts ...gen.DOOption) justUser {
	_justUser := justUser{}

	_justUser.justUserDo.UseDB(db, opts...)
	_justUser.justUserDo.UseModel(&model.JustUser{})

	tableName := _justUser.justUserDo.TableName()
	_justUser.ALL = field.NewAsterisk(tableName)
	_justUser.ID = field.NewInt32(tableName, "id")
	_justUser.Name = field.NewString(tableName, "name")
	_justUser.Address = field.NewString(tableName, "address")
	_justUser.RegisterTime = field.NewTime(tableName, "register_time")
	_justUser.Alive = field.NewBool(tableName, "alive")
	_justUser.CreatedAt = field.NewTime(tableName, "created_at")
	_justUser.CompanyID = field.NewInt32(tableName, "company_id")
	_justUser.PrivateURL = field.NewString(tableName, "private_url")
	_justUser.XMLHTTPRequest = field.NewString(tableName, "xmlHTTPRequest")
	_justUser.JStr = field.NewString(tableName, "jStr")
	_justUser.Geo = field.NewString(tableName, "geo")
	_justUser.Mint = field.NewInt32(tableName, "mint")
	_justUser.Blank = field.NewString(tableName, "blank")

	_justUser.fillFieldMap()

	return _justUser
}

type justUser struct {
	justUserDo justUserDo

	ALL            field.Asterisk
	ID             field.Int32
	Name           field.String
	Address        field.String
	RegisterTime   field.Time
	Alive          field.Bool
	CreatedAt      field.Time
	CompanyID      field.Int32
	PrivateURL     field.String
	XMLHTTPRequest field.String
	JStr           field.String
	Geo            field.String
	Mint           field.Int32
	Blank          field.String

	fieldMap map[string]field.Expr
}

func (j justUser) Table(newTableName string) *justUser {
	j.justUserDo.UseTable(newTableName)
	return j.updateTableName(newTableName)
}

func (j justUser) As(alias string) *justUser {
	j.justUserDo.DO = *(j.justUserDo.As(alias).(*gen.DO))
	return j.updateTableName(alias)
}

func (j *justUser) updateTableName(table string) *justUser {
	j.ALL = field.NewAsterisk(table)
	j.ID = field.NewInt32(table, "id")
	j.Name = field.NewString(table, "name")
	j.Address = field.NewString(table, "address")
	j.RegisterTime = field.NewTime(table, "register_time")
	j.Alive = field.NewBool(table, "alive")
	j.CreatedAt = field.NewTime(table, "created_at")
	j.CompanyID = field.NewInt32(table, "company_id")
	j.PrivateURL = field.NewString(table, "private_url")
	j.XMLHTTPRequest = field.NewString(table, "xmlHTTPRequest")
	j.JStr = field.NewString(table, "jStr")
	j.Geo = field.NewString(table, "geo")
	j.Mint = field.NewInt32(table, "mint")
	j.Blank = field.NewString(table, "blank")

	j.fillFieldMap()

	return j
}

func (j *justUser) WithContext(ctx context.Context) *justUserDo { return j.justUserDo.WithContext(ctx) }

func (j justUser) TableName() string { return j.justUserDo.TableName() }

func (j justUser) Alias() string { return j.justUserDo.Alias() }

func (j justUser) Columns(cols ...field.Expr) gen.Columns { return j.justUserDo.Columns(cols...) }

func (j *justUser) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := j.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (j *justUser) fillFieldMap() {
	j.fieldMap = make(map[string]field.Expr, 13)
	j.fieldMap["id"] = j.ID
	j.fieldMap["name"] = j.Name
	j.fieldMap["address"] = j.Address
	j.fieldMap["register_time"] = j.RegisterTime
	j.fieldMap["alive"] = j.Alive
	j.fieldMap["created_at"] = j.CreatedAt
	j.fieldMap["company_id"] = j.CompanyID
	j.fieldMap["private_url"] = j.PrivateURL
	j.fieldMap["xmlHTTPRequest"] = j.XMLHTTPRequest
	j.fieldMap["jStr"] = j.JStr
	j.fieldMap["geo"] = j.Geo
	j.fieldMap["mint"] = j.Mint
	j.fieldMap["blank"] = j.Blank
}

func (j justUser) clone(db *gorm.DB) justUser {
	j.justUserDo.ReplaceConnPool(db.Statement.ConnPool)
	return j
}

func (j justUser) replaceDB(db *gorm.DB) justUser {
	j.justUserDo.ReplaceDB(db)
	return j
}

type justUserDo struct{ gen.DO }

func (j justUserDo) Debug() *justUserDo {
	return j.withDO(j.DO.Debug())
}

func (j justUserDo) WithContext(ctx context.Context) *justUserDo {
	return j.withDO(j.DO.WithContext(ctx))
}

func (j justUserDo) ReadDB() *justUserDo {
	return j.Clauses(dbresolver.Read)
}

func (j justUserDo) WriteDB() *justUserDo {
	return j.Clauses(dbresolver.Write)
}

func (j justUserDo) Session(config *gorm.Session) *justUserDo {
	return j.withDO(j.DO.Session(config))
}

func (j justUserDo) Clauses(conds ...clause.Expression) *justUserDo {
	return j.withDO(j.DO.Clauses(conds...))
}

func (j justUserDo) Returning(value interface{}, columns ...string) *justUserDo {
	return j.withDO(j.DO.Returning(value, columns...))
}

func (j justUserDo) Not(conds ...gen.Condition) *justUserDo {
	return j.withDO(j.DO.Not(conds...))
}

func (j justUserDo) Or(conds ...gen.Condition) *justUserDo {
	return j.withDO(j.DO.Or(conds...))
}

func (j justUserDo) Select(conds ...field.Expr) *justUserDo {
	return j.withDO(j.DO.Select(conds...))
}

func (j justUserDo) Where(conds ...gen.Condition) *justUserDo {
	return j.withDO(j.DO.Where(conds...))
}

func (j justUserDo) Order(conds ...field.Expr) *justUserDo {
	return j.withDO(j.DO.Order(conds...))
}

func (j justUserDo) Distinct(cols ...field.Expr) *justUserDo {
	return j.withDO(j.DO.Distinct(cols...))
}

func (j justUserDo) Omit(cols ...field.Expr) *justUserDo {
	return j.withDO(j.DO.Omit(cols...))
}

func (j justUserDo) Join(table schema.Tabler, on ...field.Expr) *justUserDo {
	return j.withDO(j.DO.Join(table, on...))
}

func (j justUserDo) LeftJoin(table schema.Tabler, on ...field.Expr) *justUserDo {
	return j.withDO(j.DO.LeftJoin(table, on...))
}

func (j justUserDo) RightJoin(table schema.Tabler, on ...field.Expr) *justUserDo {
	return j.withDO(j.DO.RightJoin(table, on...))
}

func (j justUserDo) Group(cols ...field.Expr) *justUserDo {
	return j.withDO(j.DO.Group(cols...))
}

func (j justUserDo) Having(conds ...gen.Condition) *justUserDo {
	return j.withDO(j.DO.Having(conds...))
}

func (j justUserDo) Limit(limit int) *justUserDo {
	return j.withDO(j.DO.Limit(limit))
}

func (j justUserDo) Offset(offset int) *justUserDo {
	return j.withDO(j.DO.Offset(offset))
}

func (j justUserDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *justUserDo {
	return j.withDO(j.DO.Scopes(funcs...))
}

func (j justUserDo) Unscoped() *justUserDo {
	return j.withDO(j.DO.Unscoped())
}

func (j justUserDo) Create(values ...*model.JustUser) error {
	if len(values) == 0 {
		return nil
	}
	return j.DO.Create(values)
}

func (j justUserDo) CreateInBatches(values []*model.JustUser, batchSize int) error {
	return j.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (j justUserDo) Save(values ...*model.JustUser) error {
	if len(values) == 0 {
		return nil
	}
	return j.DO.Save(values)
}

func (j justUserDo) First() (*model.JustUser, error) {
	if result, err := j.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.JustUser), nil
	}
}

func (j justUserDo) Take() (*model.JustUser, error) {
	if result, err := j.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.JustUser), nil
	}
}

func (j justUserDo) Last() (*model.JustUser, error) {
	if result, err := j.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.JustUser), nil
	}
}

func (j justUserDo) Find() ([]*model.JustUser, error) {
	result, err := j.DO.Find()
	return result.([]*model.JustUser), err
}

func (j justUserDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.JustUser, err error) {
	buf := make([]*model.JustUser, 0, batchSize)
	err = j.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (j justUserDo) FindInBatches(result *[]*model.JustUser, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return j.DO.FindInBatches(result, batchSize, fc)
}

func (j justUserDo) Attrs(attrs ...field.AssignExpr) *justUserDo {
	return j.withDO(j.DO.Attrs(attrs...))
}

func (j justUserDo) Assign(attrs ...field.AssignExpr) *justUserDo {
	return j.withDO(j.DO.Assign(attrs...))
}

func (j justUserDo) Joins(fields ...field.RelationField) *justUserDo {
	for _, _f := range fields {
		j = *j.withDO(j.DO.Joins(_f))
	}
	return &j
}

func (j justUserDo) Preload(fields ...field.RelationField) *justUserDo {
	for _, _f := range fields {
		j = *j.withDO(j.DO.Preload(_f))
	}
	return &j
}

func (j justUserDo) FirstOrInit() (*model.JustUser, error) {
	if result, err := j.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.JustUser), nil
	}
}

func (j justUserDo) FirstOrCreate() (*model.JustUser, error) {
	if result, err := j.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.JustUser), nil
	}
}

func (j justUserDo) FindByPage(offset int, limit int) (result []*model.JustUser, count int64, err error) {
	result, err = j.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = j.Offset(-1).Limit(-1).Count()
	return
}

func (j justUserDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = j.Count()
	if err != nil {
		return
	}

	err = j.Offset(offset).Limit(limit).Scan(result)
	return
}

func (j justUserDo) Scan(result interface{}) (err error) {
	return j.DO.Scan(result)
}

func (j justUserDo) Delete(models ...*model.JustUser) (result gen.ResultInfo, err error) {
	return j.DO.Delete(models)
}

func (j *justUserDo) withDO(do gen.Dao) *justUserDo {
	j.DO = *do.(*gen.DO)
	return j
}
