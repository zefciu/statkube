package main

import (
	"fmt"
	"strings"
	"testing"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"github.com/Mirantis/statkube/models"
)

const TEST_DATA = `{
	"users": [{
		"github_id": "user1",
		"emails": ["user1@example.com", "user1@example.pl"],
		"companies": [{
				"company_name": "Intel",
				"end_date": "2015-May-01"
			}, {
				"company_name": "Mirantis",
				"end_date": null
			}

		]
	}, {
		"launchpad_id": "user2",
		"emails": ["user2@example.com", "user2@example.pl"],
		"companies": [{
				"company_name": "Intel",
				"end_date": "2015-May-01"
			}, {
				"company_name": "Mirantis",
				"end_date": null
			}

		]
	}],
	"companies": [
	]
} 
`

const TEST_DATA_COMPLEMENTED = `{
	"users": [{
		"github_id": "user1",
		"emails": ["user1@example.com", "user1@example.pl"],
		"companies": [{
				"company_name": "Intel",
				"end_date": "2015-May-01"
			}, {
				"company_name": "Mirantis",
				"end_date": null
			}

		]
	}, {
		"launchpad_id": "user2",
		"github_id": "user2-gh",
		"emails": ["user2@example.com", "user2@example.pl"],
		"companies": [{
				"company_name": "Intel",
				"end_date": "2015-May-01"
			}, {
				"company_name": "Mirantis",
				"end_date": null
			}

		]
	}],
	"companies": [
	]
} 
`

func getDB() *gorm.DB {
	db, err := gorm.Open("sqlite3", ":memory:")
	if err != nil {
		panic(fmt.Sprintf("Failed to create in-memory db: %v", err.Error()))
	}
	models.Migrate(db)
	return db
}

func getTestData() *strings.Reader {
	return strings.NewReader(TEST_DATA)
}

func loadData() *gorm.DB {
	db := getDB()
	data := getTestData()
	loadAll(db, data)
	return db
}

// TestLoadNewDeveloper tests whether we correctly load a user with github_id
func TestLoadNewDeveloper(t *testing.T) {
	var developer models.Developer
	var emails []models.Email
	db := loadData()
	db.Where("github_id = ?", "user1").Find(&developer)
	if developer.GithubID != "user1" {
		t.Error("Didn't find user1")
	}
	db.Model(&developer).Related(&emails)
	if len(emails) != 2 {
		t.Error("Number of emails wrong")
	}
}

// TestLoadNewDeveloperByLaunchpad tests whether we correctly load a user without github_id
func TestLoadNewDeveloperByLaunchpad(t *testing.T) {
	var developer models.Developer
	var emails []models.Email
	db := loadData()
	db.Where("launchpad_id = ?", "user2").Find(&developer)
	if developer.LaunchpadID != "user2" {
		t.Error("Didn't find user2")
	}
	db.Model(&developer).Related(&emails)
	if len(emails) != 2 {
		t.Error("Number of emails wrong")
	}
}

func TestComplementGithubID(t *testing.T) {
	var developer models.Developer
	var count int
	db := loadData()
	loadAll(db, strings.NewReader(TEST_DATA_COMPLEMENTED))
	db.Table("developers").Where("launchpad_id = ?", "user2").Count(&count)
	if count != 1 {
		t.Error("User duplicated instead of complimenting her data")
	}
	db.Where("launchpad_id = ?", "user2").Find(&developer)
	if developer.GithubID != "user2-gh" {
		t.Error("Github id not supplemented")
	}
}
