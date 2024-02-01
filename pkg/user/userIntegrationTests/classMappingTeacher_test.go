package userIntegrationTests

import (
	"github.com/komalreddy3/Attendance-go/pkg/user/userRepository"
	"github.com/komalreddy3/Attendance-go/pkg/user/userRepository/userModels"
	"go.uber.org/zap"
	"testing"
)

func TestClassMappingTeacher(t *testing.T) {
	// Set up the test database
	testDB, err := setupTestDB()
	if err != nil {
		t.Fatal(err)
	}
	defer teardownTestDB(testDB)

	// Initialize the repository with the test database
	repo := userRepository.NewUserRepositoryImpl(testDB, zap.NewNop().Sugar())

	// Insert classes for testing
	classNames := []string{"class1", "class2", "class3"}
	for _, className := range classNames {
		_, err := repo.InsertClass(className)
		if err != nil {
			t.Fatalf("Error inserting class for testing: %v", err)
		}
	}

	// Test the ClassMappingTeacher function
	classIDMap, err := repo.ClassMappingTeacher(classNames)

	// Assert the result as per your expectations
	if err != nil {
		t.Errorf("ClassMappingTeacher failed: %v", err)
	}

	// Verify the mapping
	for _, className := range classNames {
		classID, exists := classIDMap[className]
		if !exists {
			t.Errorf("Expected class ID for class %s, but not found", className)
		} else {
			// Optionally, you can fetch the class from the database and verify its existence
			fetchedClass := &userModels.Class{ClassID: classID}
			err = testDB.Model(fetchedClass).WherePK().Select()
			if err != nil {
				t.Errorf("Error fetching class from the database: %v", err)
			}

			if fetchedClass.ClassName != className {
				t.Errorf("Expected class name %s, got %s", className, fetchedClass.ClassName)
			}
		}
	}
}
