package day11

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strings"
)

func nth(s string) int {
	switch s {
	case "first":
		return 1
	case "second":
		return 2
	case "third":
		return 3
	case "fourth":
		return 4
	}
	return 0
}

type object struct {
	kind      int
	generator bool
}

func newObject(kind int, generator bool) object {
	return object{
		kind:      kind,
		generator: generator,
	}
}

type objects []object

func (o objects) Len() int {
	return len(o)
}

func (o objects) Swap(i, j int) {
	o[i], o[j] = o[j], o[i]
}

func (o objects) Less(i, j int) bool {
	if o[i].kind == o[j].kind {
		if o[i].generator {
			return true
		} else {
			return false
		}
	}
	return o[i].kind < o[j].kind
}

type pair struct {
	first  object
	second object
}

type floor struct {
	objects []object
}

func newFloor() floor {
	return floor{
		objects: make([]object, 0),
	}
}

func (f *floor) newCopy() floor {
	result := *f
	result.objects = make([]object, len(f.objects))
	copy(result.objects, f.objects)
	return result
}

func (f floor) String() string {
	var buffer bytes.Buffer
	for _, o := range f.objects {
		generator := 1
		if !o.generator {
			generator = 0
		}
		buffer.WriteString(fmt.Sprintf("%d%d", o.kind, generator))
	}
	return buffer.String()
}

func (f *floor) addObject(obj object) {
	if obj.kind == 0 {
		return
	}
	f.objects = append(f.objects, obj)
	sort.Sort(objects(f.objects))
}

func (f *floor) removeObject(obj object) bool {
	if obj.kind == 0 {
		return true
	}
	for i, o := range f.objects {
		if o == obj {
			f.objects = append(f.objects[:i], f.objects[i+1:]...)
			return true
		}
	}
	return false
}

func (f *floor) hasGenerator(kind int) bool {
	for _, o := range f.objects {
		if o.generator && (kind == 0 || o.kind == kind) {
			return true
		}
	}
	return false
}

func (f floor) check() bool {
	if !f.hasGenerator(0) {
		return true
	}
	for _, o := range f.objects {
		if !o.generator && !f.hasGenerator(o.kind) {
			return false
		}
	}
	return true
}

type facility struct {
	floors   map[int]floor
	elevator int
	moves    int
	distance int
}

func newFacility() facility {
	floors := make(map[int]floor)
	for i := 1; i < 5; i++ {
		floors[i] = newFloor()
	}
	return facility{
		floors:   floors,
		elevator: 1,
		moves:    0,
		distance: 0,
	}
}

func (f *facility) newCopy() facility {
	result := *f
	result.floors = make(map[int]floor)
	for k, v := range f.floors {
		result.floors[k] = v.newCopy()
	}
	return result
}

func (f *facility) addObject(floor int, kind int, generator bool) {
	fl := f.floors[floor]
	fl.addObject(newObject(kind, generator))
	f.floors[floor] = fl
	f.distance += 4 - floor
}

func (f *facility) ready() bool {
	return f.distance == 0
	//return len(f.floors[1].objects) == 0 && len(f.floors[2].objects) == 0 && len(f.floors[3].objects) == 0
}

func (f *facility) move(up bool, obj1 object, obj2 object) bool {
	f.moves++
	if (up && f.elevator == 4) || (!up && f.elevator == 1) {
		return false
	}
	floor := f.floors[f.elevator]
	floor.removeObject(obj1)
	floor.removeObject(obj2)
	f.floors[f.elevator] = floor
	prev := f.elevator
	if up {
		f.elevator++
	} else {
		f.elevator--
	}
	floor = f.floors[f.elevator]
	floor.addObject(obj1)
	floor.addObject(obj2)
	f.floors[f.elevator] = floor

	f.distance -= 4 - prev
	if obj1.kind != 0 && obj2.kind != 0 {
		f.distance -= 4 - prev
	}
	f.distance += 4 - f.elevator
	if obj1.kind != 0 && obj2.kind != 0 {
		f.distance += 4 - f.elevator
	}

	return f.floors[f.elevator].check() && f.floors[prev].check()
}

func (f *facility) variants() []pair {
	result := make([]pair, 0)
	for _, obj := range f.floors[f.elevator].objects {
		result = append(result, pair{first: obj, second: object{kind: 0}})
		for _, second := range f.floors[f.elevator].objects {
			if obj != second {
				result = append(result, pair{first: obj, second: second})
			}
		}
	}
	return result
}

func (f *facility) String() string {
	return fmt.Sprintf("e%df1%sf2%sf3%sf4%s", f.elevator, f.floors[1], f.floors[2], f.floors[3], f.floors[4])
}

func pop(facilities []facility) (facility, []facility) {
	if len(facilities) == 1 {
		return facilities[0], make([]facility, 0)
	}
	return facilities[0], facilities[1:]
}

type facilities []facility

func (f facilities) Len() int {
	return len(f)
}

func (f facilities) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}

func (f facilities) Less(i, j int) bool {
	d1, d2 := f[i].distance, f[j].distance
	m1, m2 := f[i].moves, f[j].moves
	if m1 == m2 {
		return d1 < d2
	}
	return m1 < m2
}

func Run(part int) {
	if len(os.Args) != 4 {
		log.Fatalf("Expected input file parameter")
	}
	filename := os.Args[3]
	log.Printf("Loading file %s", filename)
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("Failed to read input file - %s", err)
	}

	lines := strings.Split(string(data), "\n")
	f := newFacility()
	names := make(map[string]int)
	for _, line := range lines {
		parts := strings.Split(line, " ")
		n := nth(parts[1])
		for i := 2; i < len(parts); i++ {
			switch strings.Trim(parts[i], ",.") {
			case "generator":
				name := parts[i-1]
				if _, exists := names[name]; !exists {
					names[name] = len(names) + 1
				}
				f.addObject(n, names[name], true)
			case "microchip":
				name := strings.TrimSuffix(parts[i-1], "-compatible")
				if _, exists := names[name]; !exists {
					names[name] = len(names) + 1
				}
				f.addObject(n, names[name], false)
			}
		}
	}

	switch part {
	case 1:
		fs := make([]facility, 1)
		fs[0] = f
		var current facility
		foundFacilities := make(map[string]bool)
		foundMoves := 0
		i := 0
		for len(fs) != 0 {
			i++
			if i%1000 == 0 {
				log.Printf("Queue length is %d (min moves in queue %d)", len(fs), fs[0].moves)
			}
			current, fs = pop(fs)
			if _, exists := foundFacilities[current.String()]; exists {
				continue
			}
			foundFacilities[current.String()] = true
			variants := current.variants()
			for _, p := range variants {
				newUp := current.newCopy()
				if newUp.move(true, p.first, p.second) {
					if _, exists := foundFacilities[newUp.String()]; !exists {
						if newUp.ready() {
							log.Printf("Found solution in %d moves, queue length is %d", newUp.moves, len(fs))
							return
						}
						if foundMoves == 0 || newUp.moves < foundMoves {
							fs = append(fs, newUp)
						}
					}
				}
				newDown := current.newCopy()
				if newDown.move(false, p.first, p.second) {
					if _, exists := foundFacilities[newDown.String()]; !exists {
						if newDown.ready() {
							log.Printf("Found solution in %d moves, queue length is %d", newDown.moves, len(fs))
							return
						}
						if foundMoves == 0 || newDown.moves < foundMoves {
							fs = append(fs, newDown)
						}
					}
				}
			}
			sort.Sort(facilities(fs))
		}
	case 2:
		fs := make([]facility, 1)
		fs[0] = f
		var current facility
		foundFacilities := make(map[string]bool)
		foundMoves := 0
		i := 0
		for len(fs) != 0 {
			i++
			if i%1000 == 0 {
				log.Printf("Queue length is %d (min moves in queue %d)", len(fs), fs[0].moves)
			}
			current, fs = pop(fs)
			if _, exists := foundFacilities[current.String()]; exists {
				continue
			}
			foundFacilities[current.String()] = true
			variants := current.variants()
			for _, p := range variants {
				newUp := current.newCopy()
				if newUp.move(true, p.first, p.second) {
					if _, exists := foundFacilities[newUp.String()]; !exists {
						if newUp.ready() {
							log.Printf("Found solution in %d moves, queue length is %d", newUp.moves+24, len(fs))
							return
						}
						if foundMoves == 0 || newUp.moves < foundMoves {
							fs = append(fs, newUp)
						}
					}
				}
				newDown := current.newCopy()
				if newDown.move(false, p.first, p.second) {
					if _, exists := foundFacilities[newDown.String()]; !exists {
						if newDown.ready() {
							log.Printf("Found solution in %d moves, queue length is %d", newDown.moves+24, len(fs))
							return
						}
						if foundMoves == 0 || newDown.moves < foundMoves {
							fs = append(fs, newDown)
						}
					}
				}
			}
			sort.Sort(facilities(fs))
		}
	}
}
