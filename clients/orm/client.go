package orm

import (
	"database/sql"

	"github.com/pubgo/funk/result"
	"github.com/pubgo/funk/vars"
	"gorm.io/gorm"
)

const Name = "orm"

type Client struct {
	*gorm.DB
}

func (c *Client) Ping() error {
	var _db, err = c.DB.DB()
	if err != nil {
		return err
	}
	return _db.Ping()
}

func (c *Client) Vars() vars.Value {
	return func() interface{} {
		_db, err := c.DB.DB()
		if err != nil {
			return err.Error()
		} else {
			return _db.Stats()
		}
	}
}

func (c *Client) Close() error {
	var db, err = c.DB.DB()
	if err != nil {
		return err
	}
	return db.Close()
}

func (c *Client) Stats() (r result.Result[sql.DBStats]) {
	var db, err = c.DB.DB()
	if err != nil {
		return r.WithErr(err)
	}
	return r.WithVal(db.Stats())
}
