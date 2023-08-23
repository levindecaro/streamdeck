package streamdeck

const (
	DidReceiveSettings            = "didReceiveSettings"
	DidReceiveGlobalSettings      = "didReceiveGlobalSettings"
	KeyDown                       = "keyDown"
	KeyUp                         = "keyUp"
	TouchTap                      = "touchTap"
	DialDown                      = "dialDown"
	DialUp                        = "dialUp"
	DialRotate                    = "dialRotate"
	WillAppear                    = "willAppear"
	WillDisappear                 = "willDisappear"
	TitleParametersDidChange      = "titleParametersDidChange"
	DeviceDidConnect              = "deviceDidConnect"
	DeviceDidDisconnect           = "deviceDidDisconnect"
	ApplicationDidLaunch          = "applicationDidLaunch"
	ApplicationDidTerminate       = "applicationDidTerminate"
	SystemDidWakeUp               = "systemDidWakeUp"
	PropertyInspectorDidAppear    = "propertyInspectorDidAppear"
	PropertyInspectorDidDisappear = "propertyInspectorDidDisappear"
	SendToPlugin                  = "sendToPlugin"
	SendToPropertyInspector       = "sendToPropertyInspector"

	SetSettings       = "setSettings"
	GetSettings       = "getSettings"
	SetGlobalSettings = "setGlobalSettings"
	GetGlobalSettings = "getGlobalSettings"
	OpenURL           = "openUrl"
	LogMessage        = "logMessage"
	SetTitle          = "setTitle"
	SetFeedback       = "setFeedback"
	SetFeedbackTitle  = "setFeedback"
	SetFeedbackValue  = "setFeedback"
	SetFeedbackIcon   = "setFeedback"
	SetImage          = "setImage"
	ShowAlert         = "showAlert"
	ShowOk            = "showOk"
	SetState          = "setState"
	SwitchToProfile   = "switchToProfile"
)

type Target int

const (
	HardwareAndSoftware Target = 0
	OnlyHardware        Target = 1
	OnlySoftware        Target = 2
)
