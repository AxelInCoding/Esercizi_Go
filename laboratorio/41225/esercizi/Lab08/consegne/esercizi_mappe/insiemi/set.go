package main
import (
  "fmt"
  "os"
  "bufio"
  "strings"
  "sort"
)

//tipo StringSet
type StringSet map[string]struct{}

//crea un nuovo setvuoto
func NewStringSet() StringSet {
  return make(StringSet)
}

//converte un set in slice ordinata
func ToSlice(s StringSet) []string {
	slice := make([]string, 0, len(s))
	for elem := range s {
		slice = append(slice, elem)
	}
	sort.Strings(slice)
	return slice
}

//unione di due set
func Union(set1, set2 StringSet) StringSet {
  result := NewStringSet()
  
  for elem := range set1 {
    result[elem] = struct{}{}
  }

  for elem := range set2 {
    result[elem] = struct{}{}
  }

  return result
} 

//intersezioni di due set
func Intersection(set1, set2 StringSet) StringSet {
  result := NewStringSet()
  for elem := range set1 {
    if _, ok := set2[elem]; ok {
      result[elem] = struct{}{}
    } 
  }

  return result
}

//differenza di due set
func Difference(set1, set2 StringSet) StringSet{
  result := NewStringSet()

  for elem := range set1 {
    if _, ok := set2[elem]; !ok {
      result[elem] = struct{}{}
    } 
  }

  return result
}

func main() {
  A := NewStringSet()
  B := NewStringSet()

  scanner := bufio.NewScanner(os.Stdin)
  for scanner.Scan() {
    line := scanner.Text()
    line = strings.TrimSpace(line)
    if len(line) < 3 || line [1] != ':' {
      continue
    }

    prefix := line[0]
    value := line[2:]

    if prefix == 'A' {
      A[value] = struct{}{}
    } else if prefix == 'B' {
      B[value] = struct{}{}
    }
  }

  fmt.Println("set A:", ToSlice(A))
	fmt.Println("set B:", ToSlice(B))
	fmt.Println("unione:", ToSlice(Union(A, B)))
	fmt.Println("intersezione:", ToSlice(Intersection(A, B)))
	fmt.Println("differenza:", ToSlice(Difference(A, B)))
  
}
