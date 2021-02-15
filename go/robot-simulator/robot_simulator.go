package robot

import "fmt"

const (
	N Dir = iota
	E
	S
	W
)

func (d Dir) String() string {
	return [...]string{"N", "E", "S", "W"}[d]
}

func Advance() {
	switch Step1Robot.Dir {
	case N:
		Step1Robot.Y++
	case E:
		Step1Robot.X++
	case S:
		Step1Robot.Y--
	case W:
		Step1Robot.X--
	}
}

func Right() {
	switch Step1Robot.Dir {
	case N:
		Step1Robot.Dir = E
	case E:
		Step1Robot.Dir = S
	case S:
		Step1Robot.Dir = W
	case W:
		Step1Robot.Dir = N
	}
}

func Left() {
	switch Step1Robot.Dir {
	case N:
		Step1Robot.Dir = W
	case E:
		Step1Robot.Dir = N
	case S:
		Step1Robot.Dir = E
	case W:
		Step1Robot.Dir = S
	}
}

type Action string

type Action3 struct {
	Action
	name string
}

func StartRobot(cmd chan Command, act chan Action) {
	defer close(act)

	for c := range cmd {
		act <- Action(c)
	}
}

func Room(room Rect, robot Step2Robot, action chan Action, report chan Step2Robot) {
	defer close(report)

	for a := range action {
		switch a {
		case "A":
			robot.advance(room)
		case "R":
			robot.right()
		case "L":
			robot.left()
		}
	}
	report <- robot
}

func (r *Step2Robot) advance(room Rect) {
	switch r.Dir {
	case N:
		if r.Pos.Northing+1 <= room.Max.Northing {
			r.Pos.Northing++
		}
	case E:
		if r.Pos.Easting+1 <= room.Max.Easting {
			r.Pos.Easting++
		}
	case S:
		if r.Pos.Northing-1 >= room.Min.Northing {
			r.Pos.Northing--
		}
	case W:
		if r.Pos.Easting-1 >= room.Min.Easting {
			r.Pos.Easting--
		}
	}
}

func (r *Step2Robot) right() {
	switch r.Dir {
	case N:
		r.Dir = E
	case E:
		r.Dir = S
	case S:
		r.Dir = W
	case W:
		r.Dir = N
	}
}

func (r *Step2Robot) left() {
	switch r.Dir {
	case N:
		r.Dir = W
	case E:
		r.Dir = N
	case S:
		r.Dir = E
	case W:
		r.Dir = S
	}
}

func StartRobot3(name string, commands string, action chan Action3, log chan string) {
	defer func() {
		action <- Action3{Action('X'), name}
	}()

	if len(name) == 0 {
		log <- "invalid name"
	}

	for _, c := range commands {
		// if c != 'A' && c != 'R' && c != 'L' {
		// 	log <- "invalid command"
		// 	break
		// }
		action <- Action3{Action(c), name}
	}

}

func isRobotInRoom(robot *Step3Robot, room Rect) bool {
	return robot.Pos.Northing > room.Max.Northing ||
		robot.Pos.Northing < room.Min.Northing ||
		robot.Pos.Easting > room.Max.Easting ||
		robot.Pos.Easting < room.Min.Easting
}

func canRobotOccupyPosition(robots map[string]*Step3Robot, name string, pos Pos) bool {
	for _, r := range robots {
		if r.Name != name && r.Pos.Easting == pos.Easting && r.Pos.Northing == pos.Northing {
			return false
		}
	}

	return true
}

func Room3(room Rect, robots []Step3Robot, action chan Action3, report chan []Step3Robot, log chan string) {
	defer func() {
		report <- robots
	}()

	nameToRobot := make(map[string]*Step3Robot)
	for i := 0; i < len(robots); i++ {
		robot := &robots[i]
		if _, exits := nameToRobot[robot.Name]; exits {
			log <- fmt.Sprintf("duplicate robot name %v", robot.Name)
			return
		} else if isRobotInRoom(robot, room) {
			log <- fmt.Sprintf("robot %v is outside the room", robot.Name)
			return
		} else if !canRobotOccupyPosition(nameToRobot, robot.Name, robot.Pos) {
			log <- fmt.Sprintf("robot %v landed on another robot", robot.Name)
			return
		}

		nameToRobot[robot.Name] = robot
	}

	completeCount := 0
	for a := range action {
		robot, ok := nameToRobot[a.name]
		if !ok {
			log <- fmt.Sprintf("unknown robot %v", a.name)
			return
		}

		if robot == nil {
			println("nil robot pointer")
		}
		switch a.Action {
		default:
			log <- fmt.Sprintf("undefined command %v", a.Action)
			return
		case "A":
			if !robot.advance(room, nameToRobot) {
				log <- fmt.Sprintf("robot %v failed to move", robot.Name)
			}
		case "R":
			robot.right()
		case "L":
			robot.left()
		case "X":
			completeCount++
			if len(robots) == completeCount {
				return
			}
		}
	}
}

func (r *Step3Robot) advance(room Rect, robots map[string]*Step3Robot) bool {
	moved := false
	switch r.Dir {
	case N:
		if r.Pos.Northing+1 <= room.Max.Northing && canRobotOccupyPosition(robots, r.Name, Pos{r.Pos.Easting, r.Pos.Northing + 1}) {
			r.Pos.Northing++
			moved = true
		}
	case E:
		if r.Pos.Easting+1 <= room.Max.Easting && canRobotOccupyPosition(robots, r.Name, Pos{r.Pos.Easting + 1, r.Pos.Northing}) {
			r.Pos.Easting++
			moved = true
		}
	case S:
		if r.Pos.Northing-1 >= room.Min.Northing && canRobotOccupyPosition(robots, r.Name, Pos{r.Pos.Easting, r.Pos.Northing - 1}) {
			r.Pos.Northing--
			moved = true
		}
	case W:
		if r.Pos.Easting-1 >= room.Min.Easting && canRobotOccupyPosition(robots, r.Name, Pos{r.Pos.Easting - 1, r.Pos.Northing}) {
			r.Pos.Easting--
			moved = true
		}
	}

	return moved
}

func (r *Step3Robot) right() {
	switch r.Dir {
	case N:
		r.Dir = E
	case E:
		r.Dir = S
	case S:
		r.Dir = W
	case W:
		r.Dir = N
	}
}

func (r *Step3Robot) left() {
	switch r.Dir {
	case N:
		r.Dir = W
	case E:
		r.Dir = N
	case S:
		r.Dir = E
	case W:
		r.Dir = S
	}
}
