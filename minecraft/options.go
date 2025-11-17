package minecraft

import (
	"fmt"
	"reflect"
	"strconv"
)


type (
	Options struct {
		MainHand string `toml:"main_hand" txt:"mainHand"`
		AO bool `toml:"smooth_lighting" txt:"ao"`
		Fov int `toml:"fov" txt:"fov"`
	}
)

func WriteOptions(options Options, path string) {
	plainText := "version:9999\n"

	reflectOptions := reflect.TypeOf(options)

	for i := 0; i < reflectOptions.NumField(); i++ {
		field := reflectOptions.Field(i)

		txt := field.Tag.Get("txt")
		if (txt == "") {
			txt = field.Name
		}
		value := reflect.ValueOf(options).Field(i)
		
		switch field.Type.Kind(){
		case reflect.String:
			plainText += txt + ":" + "\"" + value.String() + "\"\n"
		case reflect.Int:
			plainText += txt + ":" + strconv.Itoa(int(value.Int())) + "\n"
		case reflect.Bool:
			plainText += txt + ":" + strconv.FormatBool(value.Bool()) + "\n"
		}
	}
	fmt.Print(plainText)
}
