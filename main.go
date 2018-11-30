package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"runtime"

	"github.com/bglebrun/PCarsUART/pcars"
	yaml "gopkg.in/yaml.v2"
)

// Config is our baud port configuration
type Config struct {
	PortName        string `yaml:"name"`
	BaudRate        int32  `yaml:"baud"`
	ProjectCarsAddr string `yaml:"memaddr"`
}

var clear map[string]func() //create a map for storing clear funcs

func init() {
	clear = make(map[string]func()) //Initialize it
	clear["linux"] = func() {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func CallClear() {
	value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
	if ok {                          //if we defined a clear func for that platform:
		value() //we execute it
	} else { //unsupported platform
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}

func main() {
	conf := Config{}
	conf.getConfig()
	/*
		c := &serial.Config{Name: conf.PortName, Baud: conf.BaudRate}
		s, err := serial.OpenPort(c)
		if err != nil {
			log.Fatal(err)
		}
		// Project cars shit goes here
	*/

	for true {
		d := pcars.ExtractData()
		fmt.Println("RPM: ", d.mRpm, " / ", d.mMaxRPM)
		fmt.Println("GEAR: ", d.mGear, " / ", d.mNumGears)
		fmt.Println("SPEED: ", d.mSpeed)
		fmt.Println("SHIFT? ", d.mClutch)
		CallClear()
	}

}

func (c *Config) getConfig() *Config {
	yamlFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Printf("Failed to open config.yaml file #%v", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal failed: %v", err)
	}
	return c
}
