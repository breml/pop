package pop

import (
	"reflect"
)

// AfterFindable callback will be called after a record, or records,
// has been retrieved from the database.
type AfterFindable interface {
	AfterFind(*Connection) error
}

// AfterEagerFindable callback will be called after a record, or records,
// has been retrieved from the database and their associations have been
// eagerly loaded.
type AfterEagerFindable interface {
	AfterEagerFind(*Connection) error
}

func (m *Model) afterFind(c *Connection, eager bool) error {
	if eager {
		if x, ok := m.Value.(AfterEagerFindable); ok {
			if err := x.AfterEagerFind(c); err != nil {
				return err
			}
		}
	} else {
		if x, ok := m.Value.(AfterFindable); ok {
			if err := x.AfterFind(c); err != nil {
				return err
			}
		}
	}

	// if the "model" is a slice/array we want
	// to loop through each of the elements in the collection
	// and call AfterFind on them if they exist.
	rv := reflect.Indirect(reflect.ValueOf(m.Value))
	kind := rv.Kind()
	if kind != reflect.Slice && kind != reflect.Array {
		return nil
	}

	for i := 0; i < rv.Len(); i++ {
		elem := rv.Index(i).Addr()

		if eager {
			x, ok := elem.Interface().(AfterEagerFindable)
			if !ok {
				continue
			}
			err := x.AfterEagerFind(c)
			if err != nil {
				return err
			}
		} else {
			x, ok := elem.Interface().(AfterFindable)
			if !ok {
				continue
			}
			err := x.AfterFind(c)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// BeforeSaveable callback will be called before a record is
// either created or updated in the database.
type BeforeSaveable interface {
	BeforeSave(*Connection) error
}

func (m *Model) beforeSave(c *Connection) error {
	if x, ok := m.Value.(BeforeSaveable); ok {
		return x.BeforeSave(c)
	}
	return nil
}

// BeforeCreateable callback will be called before a record is
// created in the database.
type BeforeCreateable interface {
	BeforeCreate(*Connection) error
}

func (m *Model) beforeCreate(c *Connection) error {
	if x, ok := m.Value.(BeforeCreateable); ok {
		return x.BeforeCreate(c)
	}
	return nil
}

// BeforeUpdateable callback will be called before a record is
// updated in the database.
type BeforeUpdateable interface {
	BeforeUpdate(*Connection) error
}

func (m *Model) beforeUpdate(c *Connection) error {
	if x, ok := m.Value.(BeforeUpdateable); ok {
		return x.BeforeUpdate(c)
	}
	return nil
}

// BeforeDestroyable callback will be called before a record is
// destroyed in the database.
type BeforeDestroyable interface {
	BeforeDestroy(*Connection) error
}

func (m *Model) beforeDestroy(c *Connection) error {
	if x, ok := m.Value.(BeforeDestroyable); ok {
		return x.BeforeDestroy(c)
	}
	return nil
}

// BeforeValidateable callback will be called before a record is
// validated during
// ValidateAndCreate, ValidateAndUpdate, or ValidateAndSave
type BeforeValidateable interface {
	BeforeValidate(*Connection) error
}

func (m *Model) beforeValidate(c *Connection) error {
	if x, ok := m.Value.(BeforeValidateable); ok {
		return x.BeforeValidate(c)
	}
	return nil
}

// AfterDestroyable callback will be called after a record is
// destroyed in the database.
type AfterDestroyable interface {
	AfterDestroy(*Connection) error
}

func (m *Model) afterDestroy(c *Connection) error {
	if x, ok := m.Value.(AfterDestroyable); ok {
		return x.AfterDestroy(c)
	}
	return nil
}

// AfterUpdateable callback will be called after a record is
// updated in the database.
type AfterUpdateable interface {
	AfterUpdate(*Connection) error
}

func (m *Model) afterUpdate(c *Connection) error {
	if x, ok := m.Value.(AfterUpdateable); ok {
		return x.AfterUpdate(c)
	}
	return nil
}

// AfterCreateable callback will be called after a record is
// created in the database.
type AfterCreateable interface {
	AfterCreate(*Connection) error
}

func (m *Model) afterCreate(c *Connection) error {
	if x, ok := m.Value.(AfterCreateable); ok {
		return x.AfterCreate(c)
	}
	return nil
}

// AfterSaveable callback will be called after a record is
// either created or updated in the database.
type AfterSaveable interface {
	AfterSave(*Connection) error
}

func (m *Model) afterSave(c *Connection) error {
	if x, ok := m.Value.(AfterSaveable); ok {
		return x.AfterSave(c)
	}
	return nil
}
