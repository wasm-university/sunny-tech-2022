package main

import (
	//"fmt"
	"math"
	"strconv"
	"syscall/js"
	"time"
)

var (
	// js.Value can be any JS object/type/constructor
	window, doc, body, canvas, context js.Value
	windowSize struct{ w, h float64 }
)

type Circle struct {
	x float64
	y float64
	radius float64
	color string
	borderWidth float64
	borderColor string
}

type Constraints struct {
	border      float64
	width       float64
	height      float64
	maxVelocity float64
}

type Cow struct {
	nickName string
	x float64
	y float64
	constraints Constraints
	xVelocity float64
	yVelocity float64
}

func (cow *Cow) initialize(nickName string, x float64, y float64, constraints Constraints) {
	cow.nickName = nickName
	cow.x = x
	cow.y = y
	cow.constraints = constraints
	cow.xVelocity = 1.0
	cow.yVelocity = -1.0
}

func (this *Cow) move() {
	this.x += this.xVelocity
	this.y += this.yVelocity
	if(this.x <= this.constraints.border || this.x >= this.constraints.width - this.constraints.border) {
		this.x -= this.xVelocity
		this.x = math.Max(this.x, this.constraints.border)
		this.x = math.Min(this.x, this.constraints.width - this.constraints.border)
		this.xVelocity = -this.xVelocity
		this.x += this.xVelocity
	}
	if(this.y <= this.constraints.border || this.y >= this.constraints.height - this.constraints.border) {
		this.y -= this.yVelocity
		this.y = math.Max(this.y, this.constraints.border)
		this.y = math.Min(this.y, this.constraints.height - this.constraints.border)
		this.yVelocity = -this.yVelocity
		this.y += this.yVelocity
	}

}

func (this *Cow) distance(boid Cow) float64 {
	var distX = this.x - boid.x
	var distY = this.y - boid.y
	return math.Sqrt(distX * distX + distY * distY)
}

func (this *Cow) moveAway (boids []*Cow, minDistance float64) {
	var distanceX = 0.0
	var distanceY = 0.0
	var numClose = 0.0

	for i := 0; i < len(boids); i++ {
		var boid = boids[i];

		if(boid.x == this.x && boid.y == this.y) {
			continue
		}

		var distance = this.distance(*boid)

		if(distance < minDistance) {
			numClose++
			var xdiff = (this.x - boid.x)
			var ydiff = (this.y - boid.y)

			if(xdiff >= 0) {
				xdiff = math.Sqrt(minDistance) - xdiff
			} else if(xdiff < 0) {
				xdiff = -math.Sqrt(minDistance) - xdiff
			}

			if(ydiff >= 0) {
				ydiff = math.Sqrt(minDistance) - ydiff
			} else if(ydiff < 0) {
				ydiff = -math.Sqrt(minDistance) - ydiff
			}
			distanceX += xdiff
			distanceY += ydiff
			//boid = nil;
		}
	}

	if(numClose == 0) {
		return
	}
	this.xVelocity -= distanceX / 5
	this.yVelocity -= distanceY / 5
}

func (this *Cow) moveCloser (boids []*Cow, distance float64) {
	if(len(boids) < 1) { return }

	var avgX = 0.0
	var avgY = 0.0

	for i := 0; i < len(boids); i++ {
		var boid = boids[i]
		if(boid.x == this.x && boid.y == this.y) {continue }
		if(this.distance(*boid) > distance) { continue }

		avgX += (this.x - boid.x)
		avgY += (this.y - boid.y)
		//boid = null
	}

	avgX /= float64(len(boids))
	avgY /= float64(len(boids))

	distance = math.Sqrt((avgX * avgX) + (avgY * avgY)) * -1.0

	if(distance == 0.0) {
		return
	}

	this.xVelocity= math.Min(this.xVelocity + (avgX / distance) * 0.15, this.constraints.maxVelocity)
	this.yVelocity = math.Min(this.yVelocity + (avgY / distance) * 0.15, this.constraints.maxVelocity)
}

func (this *Cow) moveWith (boids []*Cow, distance float64) {
	if(len(boids) < 1) { return }

	// calculate the average velocity of the other boids
	var avgX = 0.0
	var avgY = 0.0

	for i := 0; i < len(boids); i++ {
		var boid = boids[i]
		if(boid.x == this.x && boid.y == this.y) { continue }
		if(this.distance(*boid) > distance) { continue }

		avgX += boid.xVelocity
		avgY += boid.yVelocity
		//boid = null
	}
	avgX /= float64(len(boids))
	avgY /= float64(len(boids))

	distance = math.Sqrt((avgX * avgX) + (avgY * avgY)) * 1.0
	if(distance == 0.0) { return }

	this.xVelocity= math.Min(this.xVelocity + (avgX / distance) * 0.05, this.constraints.maxVelocity)
	this.yVelocity = math.Min(this.yVelocity + (avgY / distance) * 0.05, this.constraints.maxVelocity)
}

func (this *Cow) draw(context js.Value, formerX float64, formerY float64) {

	drawCircle :=
		func(circle Circle) {
			context.Call("beginPath")
			context.Call("arc", circle.x, circle.y, circle.radius, 0, 2 * math.Pi, false)
			context.Set("fillStyle", circle.color)
			context.Call("fill")
			context.Set("lineWidth", circle.borderWidth)
			context.Set("strokeStyle", circle.borderColor)
			context.Call("stroke")
		}

	previousCircle := Circle{
		x: formerX,
		y: formerY,
		radius: 4.0,
		color: "white",
		borderWidth: 2.0,
		borderColor: "white",
	}

	currentCircle := Circle{
		x: this.x,
		y: this.y,
		radius: 4.0,
		color: "blue",
		borderWidth: 1.0,
		borderColor: "blue",
	}

	drawCircle(previousCircle)
	drawCircle(currentCircle)

}


func main() {

	window = js.Global()
	doc = window.Get("document")
	body = doc.Get("body")
	canvas = doc.Call("getElementById", "canvas")

	context = canvas.Call("getContext", "2d")

	constraints := Constraints{
		border:      5.0,
		width:       800.0,
		height:      800.0,
		maxVelocity: 5.0,
	}


	var cowsList []*Cow

	for i := 0; i < 30; i++ {
		var cow = Cow{}
		cow.initialize("Bob"+strconv.Itoa(i), 10.0, 10.0, constraints)
		cowsList = append(cowsList, &cow)
	}

	for true {
		for _, cow := range cowsList {
			var formerX = cow.x
			var formerY = cow.y

			cow.moveWith(cowsList, 300.0)
			cow.moveCloser(cowsList, 300.0)
			cow.moveAway(cowsList, 15.0)
			cow.move()

			cow.draw(context, formerX, formerY)
		}
		time.Sleep(1 * time.Millisecond)
	}

	<-make(chan bool)


}

