package data

import "log"

func CreateStepsTable() {
	createTableSQL := `CREATE TABLE IF NOT EXISTS steps (
		"id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"pillar_id" INTEGER NOT NULL,
		"name" TEXT NOT NULL,
		"description" TEXT,
		"created" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		"frequency" TEXT NOT NULL,
		FOREIGN KEY (pillar_id) REFERENCES pillars(id) ON DELETE CASCADE,
		CHECK (frequency IN ('daily', 'weekly', 'monthly', 'yearly'))
	);`

	statement, err := db.Prepare(createTableSQL)
	if err != nil {
		log.Fatal(err.Error())
	}

	statement.Exec()
	log.Println("Steps Table Created")
}
