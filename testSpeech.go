package main

import (
	"bytes"
	"fmt"
	"github.com/IBM/go-sdk-core/tree/master/v5/core"
	"github.com/watson-developer-cloud/go-sdk/tree/master/v2/texttospeechv1"
	"os"
)

func ConvertTextToSpeech(audioFilePath string, audioFileName string, text string, apikey string, url string) {

//Instantiate the Watson Text to Speech service
authenticator := &core.IamAuthenticator{ApiKey: "{DzNMNRO0WhhDyBGrsTRJE1Txmvf06_BNhdVdqDJ7oyV-}",}

service, serviceErr := texttospeechv1.NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{Authenticator: authenticator,})

//Check successful instantiation
if serviceErr != nil { panic(serviceErr)}

// SYNTHESIZE
synthesizeOptions := service.NewSynthesizeOptions(text).SetAccept("audio/mp3").SetVoice("en-GB_KateVoice")

// Call the TexttoSpeech Synthesize method
synthesizeResult, _,responseErr := service.Synthesize(synthesizeOptions)

//Check successful call
if responseErr != nil { panic(responseErr)}

//Check succesfull call
if responseErr != nil { panic(responseErr)}

}

// Check successful casting
if synthesizeResult != nil {
	buff := new(bytes.Buffer)
	buff.ReadFrom(synthesizeResult)
	file,_ := os.Create(audioFileName + audioFilePath)
	file.Write(buff.Bytes())
	file.Close()
}
synthesizeResult.Close()
}

func main() {
	Speech()
}

func Speech string() {
	var audioFileName,audioFilePath,text,apikey,url string

	audioFilePath ="D:/Go/Arun/src/Watson/"
	audioFileName ="speech.wav"
	text = "Hello World"
	apikey = "<api key>"
	url = "https://api.us-south.text-to-speech.watson.cloud.ibm.com"

	ConvertTextToSpeech(audioFilePath, audioFileName, text, apikey,url)
	fmt.Println("Conversion from Text to Speech completed")
	return("Success")
}
