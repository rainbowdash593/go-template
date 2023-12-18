package database

import "time"

type Option func(*Database)

func MaxIdleConns(size int) Option {
	return func(c *Database) {
		c.maxIdleConns = size
	}
}

func MaxOpenConns(size int) Option {
	return func(c *Database) {
		c.maxOpenConns = size
	}
}

func ConnMaxLifetime(duration time.Duration) Option {
	return func(c *Database) {
		c.connMaxLifetime = duration
	}
}
