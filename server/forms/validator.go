package forms

import (
	"reflect"
	"regexp"
	"sync"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type DefaultValidator struct {
	once     sync.Once
	validate *validator.Validate
}

var _ binding.StructValidator = &DefaultValidator{}

// Engine ...
func (v *DefaultValidator) Engine() interface{} {
	v.lazyinit()
	return v.validate
}

// ValidateStruct ...
func (v *DefaultValidator) ValidateStruct(obj interface{}) error {

	if kindOfData(obj) == reflect.Struct {

		v.lazyinit()

		if err := v.validate.Struct(obj); err != nil {
			return err
		}
	}

	return nil
}

func (v *DefaultValidator) lazyinit() {
	v.once.Do(func() {

		v.validate = validator.New()
		v.validate.SetTagName("binding")

		// add any custom validations etc. here

		//Custom rule for user name
		v.validate.RegisterValidation("UserName", ValidateUserName)
		v.validate.RegisterValidation("UserID", ValidateUserID)
	})
}

func kindOfData(data interface{}) reflect.Kind {

	value := reflect.ValueOf(data)
	valueType := value.Kind()

	if valueType == reflect.Ptr {
		valueType = value.Elem().Kind()
	}
	return valueType
}

func ValidateUserName(fl validator.FieldLevel) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9_]*$`)
	return re.MatchString(fl.Field().String())
}

func ValidateUserID(fl validator.FieldLevel) bool {
	return fl.Field().Int() > 0
}
