package main

import (
	"fmt"
)

func main() {
	flags := ParseFlags()
	project := CreateProject(flags.Name)
	fmt.Println(*project.ServerPath)
	//switch *mode {
	//case "model":
	//	doActionModel(action, names)
	//case "api":
	//	doActionApi(action, names)
	//case "service":
	//	doActionService(action, names)
	//}
	//
	//db := getDb()
	//defer db.Close()
	//var migrator *migrate.Migrator
	//if *mode == "migration" {
	//	migrator = migrate.NewMigrator(db, migrations.Migrations)
	//} else {
	//	migrator = migrate.NewMigrator(db, seeding.Migrations)
	//}
	//doAction(migrator, action, name)
}

//
//func doAction(migrator *migrate.Migrator, action *string, name *string) {
//	switch *action {
//	case "init":
//		initMigration(migrator)
//	case "dropDatabase":
//		dropDatabase(migrator)
//	case "create":
//		createMigrationSql(migrator, name)
//	case "migrate":
//		runMigration(migrator)
//	case "status":
//		ms, err := migrator.MigrationsWithStatus(context.TODO())
//		if err != nil {
//			panic(err)
//		}
//		fmt.Printf("migrations: %s\n", ms)
//		fmt.Printf("unapplied migrations: %s\n", ms.Unapplied())
//		fmt.Printf("last migration group: %s\n", ms.LastGroup())
//	default:
//		log.Fatal("cannot parse action")
//	}
//}
//
//func createModel(names *nameBuilder.NameFormats) {
//	path := fmt.Sprintf("models/%s.go", *names.PascalCase)
//	writeFile(path, model.Fabric(names))
//}
//
//func writeFile(path string, data *string) {
//	file, err := os.Create(path)
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer file.Close()
//	_, err = file.WriteString(*data)
//	if err != nil {
//		log.Fatal(err)
//	}
//}
//
//func getDb() *bun.DB {
//	conf, err := config.LoadConfig()
//	if err != nil {
//		log.Fatal("cannot load config:", err)
//	}
//
//	return connect.InitDB(conf)
//}
