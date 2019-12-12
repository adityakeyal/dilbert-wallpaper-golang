package main

import (
	"fmt"
	"github.com/reujab/wallpaper"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func main() {

	ymd:=time.Now().Format("2006-01-02")
	fmt.Print(ymd)
	response, _ := http.Get("https://dilbert.com/strip/"+ymd)
	defer response.Body.Close()

	bodyArr, _ := ioutil.ReadAll(response.Body)
	body := string(bodyArr)

	//fmt.Print(body)
	img1 := strings.Index(body,"img")+3
	img2 := strings.Index(body[img1:],"<img")+img1+3
	img3 := strings.Index(body[img2:],"<img")+img2+3
	img4 := strings.Index(body[img3:],"<img")+img3+3
	imgsrc := strings.Index(body[img4:],"src=")+img4+3


	count :=0
	url := ""
	for _, char := range body[imgsrc:] {
		//fmt.Print(string(char))



		if char == '"'{
			count=count+1
		}
		if count == 1 {
			url=url+string(char)
		}
		if count == 2{
			break
		}
	}

	url="https:"+url[1:]
	//background, err := wallpaper.Get()

	//if err != nil {
	//	panic(err)
	//}

	//fmt.Println("Current wallpaper:", background)
	//wallpaper.SetFromFile("/usr/share/backgrounds/gnome/adwaita-day.jpg")
	//wallpaper.SetFromURL("https://i.imgur.com/pIwrYeM.jpg")
	wallpaper.SetFromURL(url)
	//bodyImg , _ := http.Get(url)
	//defer bodyImg.Body.Close()

	//b,_ :=ioutil.ReadAll(bodyImg.Body)
	//ioutil.WriteFile("wp.jpg" ,b,0777)



}
