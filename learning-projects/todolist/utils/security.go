package utils

import (
	"errors"
	"regexp"
)

func ValidateSanitizedTodo(title string) (string, error){ 
  if len(title) == 0 || len(title) > 255 {
    return "", errors.New("title must be between 0 and 255 characters")
  }
  // if len(title) > 255 {
  //   return "", errors.New("title must be less than 255 characters")
  // }
  sanitizedTitle := regexp.MustCompile(`[^\w\s\-\'.,]`).ReplaceAllString(title, "")
  
  return sanitizedTitle, nil
}

func ValidateID(id int) error{
  if id < 0 || id > 1000 {
    return errors.New("Invalid ID: must between 1 and 1000")
  }
  return nil
}
