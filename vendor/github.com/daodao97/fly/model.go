package fly

import (
	"database/sql"
	"time"

	"github.com/pkg/errors"

	"github.com/daodao97/fly/interval/util"
)

type Model interface {
	PrimaryKey() string
	Select(opt ...Option) (rows *Rows)
	SelectOne(opt ...Option) *Row
	Count(opt ...Option) (count int64, err error)
	Insert(record interface{}) (lastId int64, err error)
	Update(record interface{}, opt ...Option) (ok bool, err error)
	Delete(opt ...Option) (ok bool, err error)
	Exec(query string, args ...interface{}) (sql.Result, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
}

type model struct {
	connection      string
	database        string
	table           string
	fakeDelKey      string
	primaryKey      string
	columnHook      map[string]HookData
	columnValidator []Validator
	hasOne          []HasOpts
	hasMany         []HasOpts
	options         *Options
	client          *sql.DB
	saveZero        bool
	err             error
}

func New(table string, baseOpt ...With) *model {
	m := &model{
		connection: "default",
		primaryKey: "id",
		table:      table,
	}

	if table == "" {
		m.err = errors.New("table name is empty")
		return m
	}

	for _, v := range baseOpt {
		v(m)
	}

	if m.client == nil {
		m.client, m.err = db(m.connection)
	}
	return m
}

func (m *model) PrimaryKey() string {
	return m.primaryKey
}

func (m *model) Select(opt ...Option) (rows *Rows) {
	var kv []interface{}
	var err error
	defer dbLog("Select", time.Now(), &err, &kv)

	if m.err != nil {
		err = m.err
		return &Rows{Err: m.err}
	}
	opts := new(Options)
	opt = append(opt, table(m.table), database(m.database))
	if m.fakeDelKey != "" {
		opt = append(opt, WhereEq(m.fakeDelKey, 0))
	}
	for _, o := range opt {
		o(opts)
	}

	_sql, args := SelectBuilder(opt...)
	res, err := query(m.client, _sql, args...)
	kv = append(kv, "sql:", _sql, "args:", args)
	if err != nil {
		return &Rows{Err: err}
	}

	for _, has := range m.hasOne {
		res, err = m.hasOneData(res, has)
		if err != nil {
			return &Rows{Err: err}
		}
	}

	for _, has := range m.hasMany {
		res, err = m.hasManyData(res, has)
		if err != nil {
			return &Rows{Err: err}
		}
	}

	for k, v := range m.columnHook {
		for i, r := range res {
			for field, val := range r.Data {
				if k == field {
					overVal, err1 := v.Output(res[i].Data, val)
					if err1 != nil {
						err = err1
						return &Rows{Err: err}
					}
					res[i].Data[field] = overVal
				}
			}
		}
	}

	return &Rows{List: res, Err: err}
}

func (m *model) SelectOne(opt ...Option) *Row {
	opt = append(opt, Limit(1))
	rows := m.Select(opt...)
	if rows.Err != nil {
		return &Row{Err: rows.Err}
	}
	if len(rows.List) == 0 {
		return &Row{}
	}
	return &rows.List[0]
}

func (m *model) Count(opt ...Option) (count int64, err error) {
	opt = append(opt, table(m.table), AggregateCount("*"))
	var result struct {
		Count int64
	}
	err = m.SelectOne(opt...).Binding(&result)
	if err != nil {
		return 0, err
	}

	return result.Count, nil
}

func (m *model) Insert(record interface{}) (lastId int64, err error) {
	if m.err != nil {
		return 0, m.err
	}

	var kv []interface{}
	defer dbLog("Insert", time.Now(), &err, &kv)

	_record, err := util.DecodeToMap(record, m.saveZero)
	if err != nil {
		return 0, err
	}

	_record, err = m.hookInput(_record)
	if err != nil {
		return 0, err
	}

	for _, v := range m.columnValidator {
		for _, h := range v.Handle {
			ok, err := h(m, _record, _record[v.Field])
			if err != nil {
				return 0, err
			}
			if !ok {
				return 0, errors.New("ValidateHandle err " + v.Msg)
			}
		}
	}

	delete(_record, m.primaryKey)
	if len(_record) == 0 {
		return 0, errors.New("empty record to insert")
	}

	ks, vs := m.recordToKV(_record)
	_sql, args := InsertBuilder(table(m.table), Field(ks...), Value(vs...))
	kv = append(kv, "sql:", _sql, "args:", vs)

	result, err := exec(m.client, _sql, args...)
	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}

func (m *model) Update(record interface{}, opt ...Option) (ok bool, err error) {
	if m.err != nil {
		return false, m.err
	}

	var kv []interface{}
	defer dbLog("Update", time.Now(), &err, &kv)

	_record, err := util.DecodeToMap(record, m.saveZero)
	if err != nil {
		return false, err
	}

	if id, ok := _record[m.primaryKey]; ok {
		kv = append(kv, "id:", id)
		opt = append(opt, WhereEq(m.primaryKey, id))
	}

	_record, err = m.hookInput(_record)
	if err != nil {
		return false, err
	}

	delete(_record, m.primaryKey)
	if len(_record) == 0 {
		return false, errors.New("empty record to update")
	}

	for _, v := range m.columnValidator {
		for _, h := range v.Handle {
			ok, err := h(m, _record, _record[v.Field])
			if err != nil {
				return false, err
			}
			if !ok {
				return false, errors.New("ValidateHandle err " + v.Msg)
			}
		}
	}

	ks, vs := m.recordToKV(_record)
	opt = append(opt, table(m.table), Field(ks...), Value(vs...))

	_sql, args := UpdateBuilder(opt...)
	kv = append(kv, "sql:", _sql, "args:", vs)

	result, err := exec(m.client, _sql, args...)
	if err != nil {
		return false, err
	}

	effect, err := result.RowsAffected()
	if err != nil {
		return false, err
	}

	return effect >= int64(0), nil
}

func (m *model) Delete(opt ...Option) (ok bool, err error) {
	if len(opt) == 0 {
		return false, errors.New("danger, delete query must with some condition")
	}

	if m.err != nil {
		return false, m.err
	}

	opt = append(opt, table(m.table))
	if m.fakeDelKey != "" {
		return m.Update(map[string]interface{}{m.fakeDelKey: 1}, opt...)
	}

	var kv []interface{}
	defer dbLog("Delete", time.Now(), &err, &kv)

	_sql, args := DeleteBuilder(opt...)
	kv = append(kv, "slq:", _sql, "args:", args)

	result, err := exec(m.client, _sql, args...)
	if err != nil {
		return false, err
	}
	effect, err := result.RowsAffected()
	if err != nil {
		return false, err
	}
	return effect > int64(0), nil
}

func (m *model) Exec(query string, args ...interface{}) (sql.Result, error) {
	return m.client.Exec(query, args...)
}

func (m *model) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return m.client.Query(query, args...)
}

func (m *model) hookInput(record map[string]interface{}) (map[string]interface{}, error) {
	for k, v := range m.columnHook {
		for field, val := range record {
			if k == field {
				overVal, err := v.Input(record, val)
				if err != nil {
					return nil, err
				}
				record[field] = overVal
			}
		}
	}
	return record, nil
}

func (m *model) recordToKV(record map[string]interface{}) (ks []string, vs []interface{}) {
	for k, v := range record {
		ks = append(ks, k)
		vs = append(vs, v)
	}

	return ks, vs
}