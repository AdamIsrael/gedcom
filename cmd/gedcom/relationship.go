package main

import (
	"fmt"
	"sort"

	"github.com/adamisrael/gedcom/types"
	"github.com/dustin/go-humanize"
)

// Relationship represents the genealogical relationship
// between two individuals, used for calculations
type Relationship struct {
	Xref         string
	Person       types.Individual
	Generations  int
	Removed      int
	Relationship string
}

func calculateRelationshipTest(home types.Individual, target types.Individual) string {
	// Experiment with Depth-first search (DSF)

	ancestors := findAncestors(&home)
	siblings := findSiblings(&home)

	fmt.Printf("%d sibling(s) found\n", len(siblings))
	for _, sibling := range siblings {
		fmt.Printf("%s - %s\n", sibling.Person.Name[0].Name, sibling.Relationship)
	}

	// Sort Ancestors by generation
	sort.Slice(ancestors, func(i, j int) bool {
		return ancestors[i].Generations < ancestors[j].Generations
	})

	fmt.Printf("%d generation(s) found\n", len(ancestors))
	for _, ancestor := range ancestors {
		fmt.Printf("(%d) %s - %s\n", ancestor.Generations, ancestor.Person.Name[0].Name, ancestor.Relationship)
	}

	return ""
}

// findAncestors iterates recursively through an Individual's parents in order
// and returns a slice of their Relationships
func findAncestors(home *types.Individual, gen ...int) []Relationship {
	var relationships []Relationship

	// Use a variadic param for generation so the initial call
	// doesn't need to specify the generation
	var generation = 0
	if len(gen) == 0 {
		generation = 1
	} else {
		generation = gen[0] + 1
	}

	for _, parent := range home.Parents {
		if parent.Family.Husband != nil {
			relation := Relationship{
				Xref:         parent.Family.Husband.Xref,
				Person:       *parent.Family.Husband,
				Generations:  generation,
				Removed:      0,
				Relationship: getAncestorRelationship(generation, "M"),
			}
			relationships = append(relationships, relation)
			relationships = append(relationships, findAncestors(parent.Family.Husband, generation)...)
		}

		if parent.Family.Wife != nil {
			relation := Relationship{
				Xref:         parent.Family.Wife.Xref,
				Person:       *parent.Family.Wife,
				Generations:  generation,
				Removed:      0,
				Relationship: getAncestorRelationship(generation, "F"),
			}
			relationships = append(relationships, relation)
			relationships = append(relationships, findAncestors(parent.Family.Wife, generation)...)
		}
	}
	return relationships
}

// findSiblings returns a slice of Relationship for a given Individual
func findSiblings(home *types.Individual) []Relationship {
	var relationships []Relationship

	for _, parent := range home.Parents {
		for _, child := range parent.Family.Child {
			if home.Xref != child.Xref {
				relation := Relationship{
					Xref:         child.Xref,
					Person:       *child,
					Generations:  0,
					Removed:      0,
					Relationship: getSiblingRelationship(child),
				}
				relationships = append(relationships, relation)
			}
		}

	}
	return relationships
}

// findChildren recursively returns a slice of Relationship representing
// an Individual's children (and their children) to calculate cousinship
func findChildren(home *types.Individual) []Relationship {
	var relationships []Relationship

	return relationships
}

// // calculateRelationship will compare two Individuals and determine
// // the relationship between them, if any is found.
// func calculateRelationship(home types.Individual, target types.Individual) string {

// 	if home.Xref == target.Xref {
// 		return "Self"
// 	}

// 	var generation = 0
// 	var removed = 0
// 	// var err error
// 	var people []types.Individual
// 	var siblings []types.Individual
// 	// var person = home

// 	// Seed the loop
// 	people = append(people, home)

// 	// Loop through each generation
// 	for {
// 		var relationships []Relationship // Relationships in this generation

// 		// Check the parents
// 		// for _, p := range person.Parents {
// 		for _, p1 := range people {
// 			for _, p := range p1.Parents {
// 				fmt.Printf("Scanning parents of %s\n", p1.Name[0].Name)
// 				fmt.Printf("Generation: %d\n", generation)
// 				/*
// 				 * To scan a single generation, we scan the parents
// 				 * and then scan their children.
// 				 * For each child, we then scan their children (incremending removed)
// 				 */

// 				if p.Family.Husband != nil {
// 					var relation = ""
// 					if generation == 0 {
// 						relation = "Father"
// 					} else if generation == 1 {
// 						relation = "Grandfather"
// 					} else if generation == 2 {
// 						relation = "Great-Grandfather"
// 					} else {
// 						relation = fmt.Sprintf("%s Great-Grandfather", humanize.Ordinal(generation-1))
// 					}

// 					r := Relationship{
// 						Xref:         p.Family.Husband.Xref,
// 						Person:       *p.Family.Husband,
// 						Generations:  generation,
// 						Relationship: relation,
// 					}
// 					relationships = append(relationships, r)

// 					people = append(people, *p.Family.Husband)
// 				}

// 				if p.Family.Wife != nil {
// 					r := Relationship{
// 						Xref:         p.Family.Wife.Xref,
// 						Person:       *p.Family.Wife,
// 						Generations:  generation,
// 						Relationship: getGeneration(generation, "F"),
// 					}
// 					relationships = append(relationships, r)

// 					people = append(people, *p.Family.Wife)
// 				}

// 				// Check the children
// 				if p.Family.Child != nil {
// 					fmt.Printf("Scanning children\n")
// 					for _, child := range p.Family.Child {
// 						if child != nil {
// 							var relation = ""
// 							if generation == 0 {
// 								relation = "Sibling"
// 								if child.Sex == "M" {
// 									relation = "Brother"
// 								} else if child.Sex == "F" {
// 									relation = "Sister"
// 								}
// 							} else if generation == 1 {
// 								if child.Sex == "M" {
// 									relation = "Uncle"
// 								} else if child.Sex == "F" {
// 									relation = "Aunt"
// 								}
// 							} else if generation == 2 {
// 								if child.Sex == "M" {
// 									relation = "Great-Uncle"
// 								} else if child.Sex == "F" {
// 									relation = "Great-Aunt"
// 								}
// 							} else if generation == 3 {
// 								if child.Sex == "M" {
// 									relation = "Great-Granduncle"
// 								} else if child.Sex == "F" {
// 									relation = "Great-Grandaunt"
// 								}
// 							} else {
// 								if child.Sex == "M" {
// 									relation = fmt.Sprintf("%s Great-Uncle", humanize.Ordinal(generation-1))
// 								} else if child.Sex == "F" {
// 									relation = fmt.Sprintf("%s Great-Aunt", humanize.Ordinal(generation-1))
// 								}
// 							}

// 							r := Relationship{
// 								Xref:         child.Xref,
// 								Person:       *child,
// 								Generations:  generation,
// 								Relationship: relation,
// 							}
// 							relationships = append(relationships, r)

// 							siblings = append(siblings, *child)
// 						}
// 					}
// 				}

// 				// Check the children's children
// 				fmt.Printf("Scanning for cousins\n")
// 				relationships = nil
// 				var cgeneration = 1
// 				var family = p.Family

// 				// var scanCousins func(types.Individual, int, int) (string, error)
// 				// scanCousins = func(person types.Individual, generation int, removed int) (string, error) {

// 				// }

// 				// // Get a list of this generation's children
// 				// cousins := make([]*types.Individual, len(p.Family.Child))
// 				// copy(cousins, p.Family.Child)

// 				// for {

// 				// 	for cousin := range cousins {
// 				// 		relationship, err := scanCousins(cousin, cgeneration, removed)
// 				// 	}

// 				// 	cgeneration++
// 				// }
// 				// Check for match

// 				// For each child, repeat the match check (resursive?)

// 				for {
// 					// for _, child := range p.Family.Child {
// 					for _, child := range family.Child {
// 						for _, f := range child.Family {
// 							if f.Family.Child != nil {
// 								for _, c := range f.Family.Child {
// 									var relationship = ""
// 									if removed == 0 {
// 										relationship = fmt.Sprintf("%s Cousin", humanize.Ordinal(cgeneration))
// 									} else {
// 										relationship = fmt.Sprintf("%s Cousin %dx Removed", humanize.Ordinal(cgeneration), removed)
// 									}
// 									r := Relationship{
// 										Xref:         c.Xref,
// 										Person:       *c,
// 										Generations:  generation,
// 										Removed:      removed,
// 										Relationship: relationship,
// 									}
// 									relationships = append(relationships, r)
// 								}
// 							}
// 						}
// 						removed++

// 					}

// 					fmt.Printf("%d relationships to check\n", len(relationships))
// 					for _, relationship := range relationships {
// 						if target.Xref == relationship.Xref {
// 							fmt.Printf("Found target person: %s - %s\n", relationship.Person.Name[0].Name, relationship.Relationship)
// 							return relationship.Relationship
// 						}
// 					}
// 					relationships = nil
// 					removed = 0
// 					if len(relationships) == 0 {
// 						fmt.Printf("No relationships found at depth %d.\n", removed)
// 						break
// 					}
// 					fmt.Printf("%d relationships to check\n", len(relationships))
// 					cgeneration++
// 				}

// 			}

// 			fmt.Printf("%d relationships to check\n", len(relationships))
// 			for _, relationship := range relationships {
// 				if target.Xref == relationship.Xref {
// 					fmt.Printf("Found target person: %s\n", relationship.Relationship)
// 					return relationship.Relationship
// 				}
// 			}

// 			// Clear the relationships we've already checked
// 			relationships = nil

// 		}

// 		// fmt.Printf("len=%d cap=%d %v\n", len(siblings), cap(siblings), siblings)
// 		if len(people) == 0 {
// 			fmt.Printf("No more people.\n")
// 			break
// 			// } else {
// 			// 	fmt.Printf("Before: %d people\n", len(people))
// 			// 	// Take a person from the array
// 			// 	person = people[0]
// 			// 	people = people[1:]
// 			// 	fmt.Printf("After: %d people\n", len(people))

// 		} else {
// 			fmt.Printf("%d people to check\n", len(people))
// 		}

// 		fmt.Printf("Incrementing generation %d => %d\n", generation, generation+1)
// 		generation++
// 	}

// 	// // Figure out how many generations there are between the two Individuals
// 	// var scanGeneration func(types.Individual, int, int) (string, error)
// 	// // var scanChildren func(types.Individual, int, int) (int, error)

// 	// scanGeneration = func(person types.Individual, generation int, removed int) (string, error) {
// 	// 	var relationships []Relationship

// 	// 	for _, p := range person.Parents {
// 	// 		/*
// 	// 		 * To scan a single generation, we scan the parents
// 	// 		 * and then scan their children.
// 	// 		 * For each child, we then scan their children (incremending removed)
// 	// 		 */

// 	// 		if p.Family.Husband != nil {
// 	// 			r := Relationship{
// 	// 				Xref:         p.Family.Husband.Xref,
// 	// 				Person:       *p.Family.Husband,
// 	// 				Generations:  generation,
// 	// 				Relationship: "Father",
// 	// 			}
// 	// 			relationships = append(relationships, r)

// 	// 			people = append(people, *p.Family.Husband)
// 	// 		}

// 	// 		if p.Family.Wife != nil {
// 	// 			r := Relationship{
// 	// 				Xref:         p.Family.Wife.Xref,
// 	// 				Person:       *p.Family.Wife,
// 	// 				Generations:  generation,
// 	// 				Relationship: "Mother",
// 	// 			}
// 	// 			relationships = append(relationships, r)

// 	// 			people = append(people, *p.Family.Wife)
// 	// 		}

// 	// 		if p.Family.Child != nil {
// 	// 			for _, child := range p.Family.Child {
// 	// 				if child != nil {
// 	// 					var relation = "Sibling"
// 	// 					if child.Sex == "M" {
// 	// 						relation = "Brother"
// 	// 					} else if child.Sex == "F" {
// 	// 						relation = "Sister"
// 	// 					}
// 	// 					r := Relationship{
// 	// 						Xref:         child.Xref,
// 	// 						Person:       *child,
// 	// 						Generations:  generation,
// 	// 						Relationship: relation,
// 	// 					}
// 	// 					relationships = append(relationships, r)

// 	// 					people = append(people, *child)
// 	// 				}
// 	// 			}
// 	// 		}
// 	// 		fmt.Printf("len=%d cap=%d %v\n", len(people), cap(people), people)
// 	// 		fmt.Printf("len=%d cap=%d %v\n", len(relationships), cap(relationships), relationships)

// 	// 		for _, relationship := range relationships {
// 	// 			if target.Xref == relationship.Xref {
// 	// 				fmt.Printf("Found target person: %s\n", relationship.Relationship)
// 	// 				return relationship.Relationship, nil
// 	// 			}
// 	// 		}

// 	// 		// If we hit this point, we didn't find the person in this generation. Scan down through the children first
// 	// 		// generation + removed before incrementing generation and starting over

// 	// 		relationships = nil
// 	// 		if p.Family.Child != nil {
// 	// 			// Loop until we run out of children
// 	// 			for {
// 	// 				for _, child := range p.Family.Child {
// 	// 					for _, f := range child.Family {
// 	// 						if f.Family.Child != nil {
// 	// 							for _, c := range f.Family.Child {
// 	// 								var relationship = ""
// 	// 								if removed == 0 {
// 	// 									relationship = fmt.Sprintf("%s Cousin", humanize.Ordinal(generation))
// 	// 								} else {
// 	// 									relationship = fmt.Sprintf("%s Cousin %dx Removed", humanize.Ordinal(generation), removed)
// 	// 								}
// 	// 								r := Relationship{
// 	// 									Xref:         c.Xref,
// 	// 									Person:       *c,
// 	// 									Generations:  generation,
// 	// 									Relationship: relationship,
// 	// 								}
// 	// 								relationships = append(relationships, r)
// 	// 							}
// 	// 						}
// 	// 					}
// 	// 				}

// 	// 				if len(relationships) == 0 {
// 	// 					fmt.Printf("No relationships found at depth %d.\n", removed)
// 	// 					break
// 	// 				}

// 	// 				removed++
// 	// 			}

// 	// 		} else {
// 	// 			fmt.Println("This generation's children have no children.")
// 	// 		}
// 	// 	}

// 	// 	return "", errors.New("individual not found")
// 	// }

// 	// var relationship = ""
// 	// relationship, err = scanGeneration(home, generation, removed)
// 	// if err != nil {
// 	// 	return ""
// 	// } else {
// 	// 	return relationship
// 	// 	// return "generation goes here"
// 	// 	// return getGeneration(generation, target.Sex)
// 	// }
// 	return ""
// }

func getAncestorRelationship(generation int, gender string) string {
	var description = ""

	if generation == 0 {
		description = "Self"
	} else if generation == 1 {
		if gender == "M" {
			description = "Father"
		} else if gender == "F" {
			description = "Mother"
		} else {
			description = "Parent"
		}
	} else if generation == 2 {
		if gender == "M" {
			description = "Grandfather"
		} else if gender == "F" {
			description = "Grandmother"
		} else {
			description = "Grandparent"
		}
	} else if generation == 3 {
		if gender == "M" {
			description = "Great-Grandfather"
		} else if gender == "F" {
			description = "Great-Grandmother"
		} else {
			description = "Great-Grandparent"
		}
	} else {
		// Calculate the nth great-grandparant
		if gender == "M" {
			description = fmt.Sprintf("%s Great-Grandfather", humanize.Ordinal(generation-1))
		} else if gender == "F" {
			description = fmt.Sprintf("%s Great-Grandmother", humanize.Ordinal(generation-1))
		} else {
			description = fmt.Sprintf("%s Great-Grandarent", humanize.Ordinal(generation-1))
		}

	}

	return description
}

func getSiblingRelationship(sibling *types.Individual) string {
	var relation = "Sibling"

	if sibling.Sex == "M" {
		relation = "Brother"
	} else if sibling.Sex == "F" {
		relation = "Sister"
	}
	return relation
}
