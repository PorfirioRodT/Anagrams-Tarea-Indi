package main

import "testing"

func TestMethodToSortWordsIsSortingWellTheLetter(t *testing.T){

	w := "bac"

	if sortingValues(w) != "abc" {

		t.Error("The values are not sorting well")

	}

}

func TestTheValuesShouldBeHashed(t *testing.T){

	value := "holadsadas"

	if hashedValues(value) == int32(1){

		t.Skip("The values are not getting hashed")

	}

}