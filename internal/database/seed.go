package database

import (
	"log"

	"glossika-assignment/internal/model"

	"github.com/go-faker/faker/v4"
	"gorm.io/gorm"
)

// SeedDatabase seeds the database with initial data
func SeedDatabase(db *gorm.DB) error {
	// Seed recommendations with default count of 100
	if err := SeedRecommendations(db, 100); err != nil {
		return err
	}

	log.Println("Database seeding completed successfully")
	return nil
}

// SeedRecommendations seeds the recommendations table with sample data
// count parameter determines how many recommendations to generate
func SeedRecommendations(db *gorm.DB, count int) error {
	log.Printf("Seeding %d recommendations...", count)

	// Check if recommendations already exist
	var existingCount int64
	if err := db.Model(&model.Recommendation{}).Count(&existingCount).Error; err != nil {
		return err
	}

	// Skip seeding if recommendations already exist
	if existingCount > 0 {
		log.Printf("Found %d existing recommendations, skipping...", existingCount)
		return nil
	}

	// Generate and insert recommendations
	recommendations := make([]model.Recommendation, count)
	for i := 0; i < count; i++ {
		recommendations[i] = model.Recommendation{
			Title:       faker.Sentence(),
			Description: faker.Paragraph(),
			Score:       float64(faker.RandomUnixTime()%40+10) / 10, // Random score between 1.0 and 5.0
		}
	}

	// Insert in batches
	batchSize := 10
	if count < 10 {
		batchSize = count
	}

	if err := db.CreateInBatches(recommendations, batchSize).Error; err != nil {
		return err
	}

	log.Printf("Successfully seeded %d recommendations", count)
	return nil
}
