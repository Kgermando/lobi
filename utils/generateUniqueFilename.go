package utils

import (
	"path/filepath"
	"strings"

	"github.com/google/uuid"
)


func GenerateUniqueFilename(original string) (string) {
	// Extract filename and extension
	ext := filepath.Ext(original)
	filename := strings.TrimSuffix(original, ext)
  
	// Generate a unique UUID
	uuid, err := uuid.NewRandom()
	if err != nil {
	  return ""
	}
  
	// Combine filename, UUID, and extension
	return filename + "_" + uuid.String() + ext
  }
  
  