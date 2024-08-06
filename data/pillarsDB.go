package data

import "log"

func CreatePillarTable() {
	createTableSQL := `CREATE TABLE IF NOT EXISTS pillars(
		"id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"name" TEXT NOT NULL,
		"description" TEXT
	);`

	statement, err := db.Prepare(createTableSQL)
	if err != nil {
		log.Fatal(err.Error())
	}

	statement.Exec()
	log.Println("Pillars Table Created")

	addPillars := `INSERT INTO pillars (name, description) VALUES
	("Physical Health", "Prioritize your body's well-being through exercise, nutrition, and proper rest."),
	("Mental Health", "Foster a positive mindset with mindfulness, stress management, and emotional resilience."),
	("Personal Growth", "Expand your horizons by learning new skills, pursuing hobbies, and engaging in continuous education."),
	("Relationships", "Strengthen your connections with family, friends, and colleagues through effective communication and meaningful interactions."),
	("Professional Development", "Advance your career and enhance productivity with goal setting, efficient time management, and work-life balance strategies.");`

	statement, err = db.Prepare(addPillars)
	if err != nil {
		log.Fatal(err.Error())
	}

	statement.Exec()
	log.Println("Pillars Added")
}

func GetPillars() []Pillar {
	rows, err := db.Query("SELECT name, description FROM pillars")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer rows.Close()

	pillars := []Pillar{}
	for rows.Next() {
		var name, description string
		err = rows.Scan(&name, &description)
		if err != nil {
			log.Fatal(err.Error())
		}
		pillar := Pillar{
			Name:        name,
			Description: description,
		}
		pillars = append(pillars, pillar)
	}

	return pillars
}
