// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"database/sql"

	"gorm.io/gorm"

	"gorm.io/gen"

	"gorm.io/plugin/dbresolver"
)

var (
	Q               = new(Query)
	AdminPermission *adminPermission
	AdminRole       *adminRole
	AdminUser       *adminUser
)

func SetDefault(db *gorm.DB, opts ...gen.DOOption) {
	*Q = *Use(db, opts...)
	AdminPermission = &Q.AdminPermission
	AdminRole = &Q.AdminRole
	AdminUser = &Q.AdminUser
}

func Use(db *gorm.DB, opts ...gen.DOOption) *Query {
	return &Query{
		db:              db,
		AdminPermission: newAdminPermission(db, opts...),
		AdminRole:       newAdminRole(db, opts...),
		AdminUser:       newAdminUser(db, opts...),
	}
}

type Query struct {
	db *gorm.DB

	AdminPermission adminPermission
	AdminRole       adminRole
	AdminUser       adminUser
}

func (q *Query) Available() bool { return q.db != nil }

func (q *Query) clone(db *gorm.DB) *Query {
	return &Query{
		db:              db,
		AdminPermission: q.AdminPermission.clone(db),
		AdminRole:       q.AdminRole.clone(db),
		AdminUser:       q.AdminUser.clone(db),
	}
}

func (q *Query) ReadDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Read))
}

func (q *Query) WriteDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Write))
}

func (q *Query) ReplaceDB(db *gorm.DB) *Query {
	return &Query{
		db:              db,
		AdminPermission: q.AdminPermission.replaceDB(db),
		AdminRole:       q.AdminRole.replaceDB(db),
		AdminUser:       q.AdminUser.replaceDB(db),
	}
}

type queryCtx struct {
	AdminPermission IAdminPermissionDo
	AdminRole       IAdminRoleDo
	AdminUser       IAdminUserDo
}

func (q *Query) WithContext(ctx context.Context) *queryCtx {
	return &queryCtx{
		AdminPermission: q.AdminPermission.WithContext(ctx),
		AdminRole:       q.AdminRole.WithContext(ctx),
		AdminUser:       q.AdminUser.WithContext(ctx),
	}
}

func (q *Query) Transaction(fc func(tx *Query) error, opts ...*sql.TxOptions) error {
	return q.db.Transaction(func(tx *gorm.DB) error { return fc(q.clone(tx)) }, opts...)
}

func (q *Query) Begin(opts ...*sql.TxOptions) *QueryTx {
	tx := q.db.Begin(opts...)
	return &QueryTx{Query: q.clone(tx), Error: tx.Error}
}

type QueryTx struct {
	*Query
	Error error
}

func (q *QueryTx) Commit() error {
	return q.db.Commit().Error
}

func (q *QueryTx) Rollback() error {
	return q.db.Rollback().Error
}

func (q *QueryTx) SavePoint(name string) error {
	return q.db.SavePoint(name).Error
}

func (q *QueryTx) RollbackTo(name string) error {
	return q.db.RollbackTo(name).Error
}
