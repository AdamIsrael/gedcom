package relationship

import (
	"fmt"

	"github.com/adamisrael/gedcom/types"
	"github.com/dustin/go-humanize"
)

/*
Handle all calculations to determine the relationship between Individuals
in a tree.

TODO: Support half relations

*/

// Relationship represents the genealogical relationship
// between two individuals, used for calculations
type Relationship struct {
	Xref         string
	Person       types.Individual
	Generations  int
	Removed      int
	Relationship string
}

func CalculateRelationship(home types.Individual, target types.Individual) string {
	var relationship = ""
	// iterate through each generation to find a common ancestor
	a := findAncestors(&home, 0)
	b := findAncestors(&target, 0)

	// var seen = make(map[string]int)
	var found = 0

	for _, ancestorA := range a {
		if ancestorA.Xref == target.Xref {
			relationship = ancestorA.Relationship
			break
		}

		for _, ancestorB := range b {
			if ancestorA.Xref == ancestorB.Xref {
				// fmt.Printf("Found common ancestor at generation %d: %s\n", ancestorA.Generations, ancestorA.Person.Name[0].Name)
				// fmt.Printf("Generation %d vs %d\n", ancestorA.Generations, ancestorB.Generations)
				// fmt.Printf("Removed %d vs %d\n", ancestorA.Removed, ancestorB.Removed)

				if ancestorA.Generations >= ancestorB.Generations {
					removed := ancestorA.Generations - ancestorB.Generations
					relationship = getChildRelationship(ancestorB.Generations, removed, &target)
				} else {
					removed := ancestorB.Generations - ancestorA.Generations
					if ancestorA.Generations == 0 {
						relationship = getSiblingChildRelationship(ancestorB.Generations, removed, &target)
					} else {
						relationship = getChildRelationship(ancestorA.Generations, removed, &target)

					}

				}

				found = 1
				break
			}

			if found == 1 {
				break
			}
		}
	}

	return relationship
}

// findAncestors iterates recursively through an Individual's parents in order
// and returns a slice of their Relationships
func findAncestors(home *types.Individual, generation int) []Relationship {
	var relationships []Relationship

	for _, parent := range home.Parents {
		if parent.Family.Husband != nil {
			relation := Relationship{
				Xref:         parent.Family.Husband.Xref,
				Person:       *parent.Family.Husband,
				Generations:  generation,
				Removed:      0,
				Relationship: GetAncestorRelationship(generation+1, "M"),
			}
			relationships = append(relationships, relation)
			relationships = append(relationships, findAncestors(parent.Family.Husband, generation+1)...)
		}

		if parent.Family.Wife != nil {
			relation := Relationship{
				Xref:         parent.Family.Wife.Xref,
				Person:       *parent.Family.Wife,
				Generations:  generation,
				Removed:      0,
				Relationship: GetAncestorRelationship(generation+1, "F"),
			}
			relationships = append(relationships, relation)
			relationships = append(relationships, findAncestors(parent.Family.Wife, generation+1)...)
		}
	}
	return relationships
}

// findSiblings returns a slice of Relationship for a given Individual
func findSiblings(home *types.Individual) []Relationship {
	fmt.Printf("Finding siblings for %s\n", home.Name[0].Name)
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
func findChildren(home *types.Individual, generation int, removed ...int) []Relationship {
	fmt.Printf("Finding children for %s, generation %d, %d removed\n", home.Name[0].Name, generation, removed)
	var relationships []Relationship

	// Use a variadic param for generation so the initial call
	// doesn't need to specify the generation
	var _removed = 0
	if len(removed) == 0 {
		_removed = 0
	} else {
		_removed = removed[0] + 1
	}

	/* TODO: Figure out how to deal with step-children and half-siblings
	* Maybe merge the (n) families and keep a count of how many families the
	* child appears in. If occurrences < n then they are half-siblings to
	 */
	for _, family := range home.Family {
		for _, child := range family.Family.Child {
			if child != nil && home.Xref != child.Xref {
				relation := Relationship{
					Xref:         child.Xref,
					Person:       *child,
					Generations:  generation,
					Removed:      _removed,
					Relationship: getChildRelationship(generation, _removed, child),
				}
				fmt.Printf("Child: %s (%d/%d)\n", child.Name[0].Name, generation, _removed)
				relationships = append(relationships, relation)

				// for _, family := range child.Family {
				// 	for _, ch := range family.Family.Child {
				// 		relationships = append(relationships, findChildren(ch, generation, _removed)...)
				// 	}
				// }
			}

		}
	}

	return relationships
}

// findAncestralRelationship finds the relationship between an Individual and
// a common Ancestor
func findAncestralRelationship(home *types.Individual, ancestor *types.Individual, generation int) *Relationship {
	var relationship = Relationship{
		Xref:    home.Xref,
		Person:  *home,
		Removed: 0,
	}

	generation++
	for _, parent := range home.Parents {
		if parent.Family.Husband != nil {
			if parent.Family.Husband.Xref == ancestor.Xref {
				relationship.Generations = generation
				relationship.Relationship = GetAncestorRelationship(generation, "M")
				return &relationship
			}
		}

		if parent.Family.Wife != nil {
			if parent.Family.Wife.Xref == ancestor.Xref {
				relationship.Generations = generation
				relationship.Relationship = GetAncestorRelationship(generation, "F")
				return &relationship
			}
		}

		if parent.Family.Husband != nil {
			var r = findAncestralRelationship(home, parent.Family.Husband, generation)
			if r != nil {
				return r
			}
		}

		if parent.Family.Wife != nil {
			var r = findAncestralRelationship(home, parent.Family.Wife, generation)
			if r != nil {
				return r
			}
		}

	}

	return nil
}

func GetAncestorRelationship(generation int, gender string) string {
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
			description = fmt.Sprintf("%s Great-Grandfather", humanize.Ordinal(generation-2))
		} else if gender == "F" {
			description = fmt.Sprintf("%s Great-Grandmother", humanize.Ordinal(generation-2))
		} else {
			description = fmt.Sprintf("%s Great-Grandarent", humanize.Ordinal(generation-2))
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

//  getSiblingChildRelationship returns the relationship between self and a child of a siling
func getSiblingChildRelationship(generation int, removed int, child *types.Individual) string {
	var relation = ""

	if generation == 1 {
		if child.Sex == "M" {
			relation = "Nephew"
		} else if child.Sex == "F" {
			relation = "Niece"
		} else {
			// TODO: Find a gender-neutral term
			relation = "Nephew/Niece"
		}
	} else if generation == 2 {
		if child.Sex == "M" {
			relation = "Grandnephew"
		} else if child.Sex == "F" {
			relation = "Grandniece"
		} else {
			// TODO: Find a gender-neutral term
			relation = "Grandnephew/Grandniece"
		}
		// grand niece/nephew
	} else if generation == 3 {
		if child.Sex == "M" {
			relation = "Great-Grandnephew"
		} else if child.Sex == "F" {
			relation = "Great-Grandniece"
		} else {
			// TODO: Find a gender-neutral term
			relation = "Great-Grandnephew/Grandniece"
		}
		// great-grandniece/nephew
	} else {
		if child.Sex == "M" {
			relation = fmt.Sprintf("%s Great-Grandnephew", humanize.Ordinal(removed-2))
		} else if child.Sex == "F" {
			relation = fmt.Sprintf("%s Great-Grandniece", humanize.Ordinal(removed-2))
		} else {
			// TODO: Find a gender-neutral term
			relation = fmt.Sprintf("%s Great-Grandnephew/Grandniece", humanize.Ordinal(removed-2))
		}
		// nth great-grandniece/nephew
	}
	return relation
}

func getChildRelationship(generation int, removed int, child *types.Individual) string {
	var relation = ""

	if generation == 0 {
		if removed == 0 {
			if child.Sex == "M" {
				relation = "Brother"
			} else if child.Sex == "F" {
				relation = "Sister"
			}
		} else if removed == 1 {
			if child.Sex == "M" {
				relation = "Uncle"
			} else if child.Sex == "F" {
				relation = "Aunt"
			}
		} else if removed == 2 {
			if child.Sex == "M" {
				relation = "Great-Uncle"
			} else if child.Sex == "F" {
				relation = "Great-Aunt"
			}
		} else if removed == 3 {
			if child.Sex == "M" {
				relation = "Great-Granduncle"
			} else if child.Sex == "F" {
				relation = "Great-Grandaunt"
			}
		} else {
			if child.Sex == "M" {
				relation = fmt.Sprintf("%s Great-Granduncle", humanize.Ordinal(removed-2))
			} else if child.Sex == "F" {
				relation = fmt.Sprintf("%s Great-Grandaunt", humanize.Ordinal(removed-2))
			}
		}

	} else {
		if removed == 0 {
			relation = fmt.Sprintf("%s Cousin", humanize.Ordinal(generation))
		} else {
			relation = fmt.Sprintf("%s Cousin %dx Removed", humanize.Ordinal(generation), removed)
		}
	}

	return relation
}
