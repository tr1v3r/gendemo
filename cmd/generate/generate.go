package main

import (
	"gorm.io/gen"
	"gorm.io/gen/field"

	"gendemo/dal"
	"gendemo/dal/method"
	"gendemo/dal/model"
)

// unexported struct will be ignored.
type test struct {
	id  int
	Xxx string
	Ttt int
}

func main() {
	// dal.init()

	g := gen.NewGenerator(gen.Config{
		OutPath: "../../dal/query",
		Mode:    gen.WithDefaultQuery, /*WithQueryInterface, WithoutContext*/

		WithUnitTest: true,
	})
	g.UseDB(dal.DB)

	g.ApplyBasic(model.Customer{}, model.CreditCard{}, model.Bank{}) // Associations
	g.ApplyBasic(g.GenerateModel("people"), g.GenerateModelAs("user", "JustUser"))
	g.ApplyInterface(func(method.Method) {}, &model.User{}, &model.Company{}, test{}) // struct test will be ignored

	g.ApplyBasic(model.Company{}, model.TestField{})
	g.ApplyBasic(model.Passport{}, g.GenerateModel("people"),
		g.GenerateModelAs("address", "Addr",
			gen.FieldIgnore("deleted_at"),
			gen.FieldNewTag("id", field.Tag{"newTag": "tag info"}),
		),
	)
	g.ApplyInterface(func(method.Method) {}, model.TestField{}, &model.Company{}, model.User{}, test{})
	g.ApplyInterface(func(method method.UserMethod) {}, model.User{})
	g.Execute()
}
