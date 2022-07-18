package options

import "gorm.io/gorm"

type Where struct {
	Query interface{}
	Args  []interface{}
}

type Paginate struct {
	Page     int
	PageSize int
}

type Option struct {
	Where
	Paginate

	Unscoped bool
}

type Opt func(*Option)

func WithUnscoped(unscoped bool) Opt {
	return func(option *Option) {
		option.Unscoped = unscoped
	}
}

func WithWhere(query interface{}, args ...interface{}) Opt {
	return func(option *Option) {
		option.Where.Query = query
		option.Where.Args = args
	}
}

func WithQuery(query interface{}) Opt {
	return func(option *Option) {
		option.Where.Query = query
	}
}

func WithArgs(args ...interface{}) Opt {
	return func(option *Option) {
		option.Where.Args = args
	}
}

func WithPaginate(page, size int) Opt {
	return func(option *Option) {
		option.Paginate.Page = page
		option.Paginate.PageSize = size
	}
}

func ScopesPaginate(o *Option) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page := o.Page
		if page == 0 {
			page = 1
		}

		pageSize := o.PageSize
		switch {
		case pageSize > 200:
			pageSize = 200
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
