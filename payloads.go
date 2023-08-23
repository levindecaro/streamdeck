package streamdeck

import "encoding/json"

type LogMessagePayload struct {
	Message string `json:"message"`
}

type OpenURLPayload struct {
	URL string `json:"url"`
}

type SetTitlePayload struct {
	Title  string `json:"title"`
	Target Target `json:"target"`
}

type SetFeedbackPayloadMap struct {
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
}

type SetFeedbackTitlePayload struct {
	Title  string `json:"title"`
	Target Target `json:"target"`
}

type SetFeedbackIconPayload struct {
	Icon   string `json:"icon"`
	Target Target `json:"target"`
}

type SetFeedbackValuePayload struct {
	Value  string `json:"value"`
	Target Target `json:"target"`
}

type SetImagePayload struct {
	Base64Image string `json:"image"`
	Target      Target `json:"target"`
}

type SetStatePayload struct {
	State int `json:"state"`
}

type SwitchProfilePayload struct {
	Profile string `json:"profile"`
}

type DidReceiveSettingsPayload struct {
	Settings        json.RawMessage `json:"settings,omitempty"`
	Coordinates     Coordinates     `json:"coordinates,omitempty"`
	IsInMultiAction bool            `json:"isInMultiAction,omitempty"`
}

type SendToPluginPayload struct {
	Settings json.RawMessage `json:"settings,omitempty"`
}

type Coordinates struct {
	Column int `json:"column,omitempty"`
	Row    int `json:"row,omitempty"`
}

type DidReceiveGlobalSettingsPayload struct {
	Settings json.RawMessage `json:"settings,omitempty"`
}

type TouchTapPayload struct {
	Settings    json.RawMessage `json:"settings,omitempty"`
	Coordinates Coordinates     `json:"coordinates,omitempty"`
	TapPos      []int           `json:"tapPos,omitempty"`
	Hold        bool            `json:"hold,omitempty"`
}

type DialDownPayload struct {
	Settings    json.RawMessage `json:"settings,omitempty"`
	Coordinates Coordinates     `json:"coordinates,omitempty"`
	Pressed     bool            `json:"pressed,omitempty"`
}

type DialUpPayload struct {
	Settings    json.RawMessage `json:"settings,omitempty"`
	Coordinates Coordinates     `json:"coordinates,omitempty"`
	Pressed     bool            `json:"pressed,omitempty"`
}

type DialRotateEventPayload struct {
	Settings    json.RawMessage `json:"settings,omitempty"`
	Coordinates Coordinates     `json:"coordinates,omitempty"`
	Ticks       int             `json:"ticks,omitempty"`
	Pressed     bool            `json:"pressed,omitempty"`
}

type KeyDownPayload struct {
	Settings         json.RawMessage `json:"settings,omitempty"`
	Coordinates      Coordinates     `json:"coordinates,omitempty"`
	State            int             `json:"state,omitempty"`
	UserDesiredState int             `json:"userDesiredState,omitempty"`
	IsInMultiAction  bool            `json:"isInMultiAction,omitempty"`
}

type KeyUpPayload struct {
	Settings         json.RawMessage `json:"settings,omitempty"`
	Coordinates      Coordinates     `json:"coordinates,omitempty"`
	State            int             `json:"state,omitempty"`
	UserDesiredState int             `json:"userDesiredState,omitempty"`
	IsInMultiAction  bool            `json:"isInMultiAction,omitempty"`
}

type WillAppearPayload struct {
	Settings        json.RawMessage `json:"settings,omitempty"`
	Coordinates     Coordinates     `json:"coordinates,omitempty"`
	State           int             `json:"state,omitempty"`
	IsInMultiAction bool            `json:"isInMultiAction,omitempty"`
}

type WillDisappearPayload struct {
	Settings        json.RawMessage `json:"settings,omitempty"`
	Coordinates     Coordinates     `json:"coordinates,omitempty"`
	State           int             `json:"state,omitempty"`
	IsInMultiAction bool            `json:"isInMultiAction,omitempty"`
}

type TitleParametersDidChangePayload struct {
	Settings        json.RawMessage `json:"settings,omitempty"`
	Coordinates     Coordinates     `json:"coordinates,omitempty"`
	State           int             `json:"state,omitempty"`
	Title           string          `json:"title,omitempty"`
	TitleParameters TitleParameters `json:"titleParameters,omitempty"`
}

type TitleParameters struct {
	FontFamily     string `json:"fontFamily,omitempty"`
	FontSize       int    `json:"fontSize,omitempty"`
	FontStyle      string `json:"fontStyle,omitempty"`
	FontUnderline  bool   `json:"fontUnderline,omitempty"`
	ShowTitle      bool   `json:"showTitle,omitempty"`
	TitleAlignment string `json:"titleAlignment,omitempty"`
	TitleColor     string `json:"titleColor,omitempty"`
}

type ApplicationDidLaunchPayload struct {
	Application string `json:"application,omitempty"`
}

type ApplicationDidTerminatePayload struct {
	Application string `json:"application,omitempty"`
}
