package streamdeck

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
	sdcontext "github.com/samwho/streamdeck/context"
)

type EventHandler func(ctx context.Context, client *Client, event Event) error

type Client struct {
	ctx      context.Context
	params   RegistrationParams
	c        *websocket.Conn
	actions  map[string]*Action
	handlers map[string][]EventHandler
	done     chan struct{}
}

func NewClient(ctx context.Context, params RegistrationParams) *Client {
	return &Client{
		ctx:     ctx,
		params:  params,
		actions: make(map[string]*Action),
		done:    make(chan struct{}),
	}
}
func (client *Client) Action(uuid string) *Action {
	_, ok := client.actions[uuid]
	if !ok {
		client.actions[uuid] = &Action{
			uuid:     uuid,
			handlers: make(map[string][]EventHandler),
		}
	}
	return client.actions[uuid]
}

func (client *Client) Run() error {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "ws", Host: fmt.Sprintf("127.0.0.1:%d", client.params.Port)}
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		return err
	}

	client.c = c

	go func() {
		defer close(client.done)
		for {
			messageType, message, err := client.c.ReadMessage()
			if err != nil {
				log.Printf("read error: %v\n", err)
				return
			}

			if messageType == websocket.PingMessage {
				log.Printf("received ping message\n")
				if err := client.c.WriteMessage(websocket.PongMessage, []byte{}); err != nil {
					log.Printf("error while ponging: %v\n", err)
				}
				continue
			}

			event := Event{}
			if err := json.Unmarshal(message, &event); err != nil {
				log.Printf("failed to unmarshal received event: %s\n", string(message))
				continue
			}

			log.Println("recv: ", string(message))

			ctx := sdcontext.WithContext(client.ctx, event.Context)
			ctx = sdcontext.WithDevice(ctx, event.Device)
			ctx = sdcontext.WithAction(ctx, event.Action)

			if event.Action == "" {
				for _, f := range client.handlers[event.Event] {
					if err := f(ctx, client, event); err != nil {
						log.Printf("error in handler for event %v: %v\n", event.Event, err)
					}
				}
				continue
			}

			action, ok := client.actions[event.Action]
			if !ok {
				log.Printf("received event for nonexistent action: %v\n", event.Action)
				continue
			}

			for _, f := range action.handlers[event.Event] {
				if err := f(ctx, client, event); err != nil {
					log.Printf("error in handler for event %v: %v\n", event.Event, err)
				}
			}
		}
	}()

	if err := client.register(client.params); err != nil {
		return err
	}

	select {
	case <-client.done:
		return nil
	case <-interrupt:
		log.Printf("interrupted, closing...\n")
		return client.Close()
	}
}

func (client *Client) register(params RegistrationParams) error {
	if err := client.send(Event{UUID: params.PluginUUID, Event: params.RegisterEvent}); err != nil {
		client.Close()
		return err
	}
	return nil
}

func (client *Client) send(event Event) error {
	j, _ := json.Marshal(event)
	log.Printf("sending message: %v\n", string(j))
	return client.c.WriteJSON(event)
}

func (client *Client) SetSettings(ctx context.Context, settings interface{}) error {
	return client.send(NewEvent(ctx, SetSettings, settings))
}

func (client *Client) GetSettings(ctx context.Context) error {
	return client.send(NewEvent(ctx, GetSettings, nil))
}

func (client *Client) SetGlobalSettings(ctx context.Context, settings interface{}) error {
	return client.send(NewEvent(ctx, SetGlobalSettings, settings))
}

func (client *Client) GetGlobalSettings(ctx context.Context) error {
	return client.send(NewEvent(ctx, GetGlobalSettings, nil))
}

func (client *Client) OpenURL(ctx context.Context, u url.URL) error {
	return client.send(NewEvent(ctx, OpenURL, OpenURLPayload{URL: u.String()}))
}

func (client *Client) LogMessage(message string) error {
	return client.send(NewEvent(nil, LogMessage, LogMessagePayload{Message: message}))
}

func (client *Client) SetTitle(ctx context.Context, title string, target Target) error {
	return client.send(NewEvent(ctx, SetTitle, SetTitlePayload{Title: title, Target: target}))
}

func (client *Client) SetImage(ctx context.Context, base64image string, target Target) error {
	return client.send(NewEvent(ctx, SetImage, SetImagePayload{Base64Image: base64image, Target: target}))
}

func (client *Client) ShowAlert(ctx context.Context) error {
	return client.send(NewEvent(ctx, ShowAlert, nil))
}

func (client *Client) ShowOk(ctx context.Context) error {
	return client.send(NewEvent(ctx, ShowOk, nil))
}

func (client *Client) SetState(ctx context.Context, state int) error {
	return client.send(NewEvent(ctx, SetState, SetStatePayload{State: state}))
}

func (client *Client) SwitchToProfile(ctx context.Context, profile string) error {
	return client.send(NewEvent(ctx, SwitchToProfile, SwitchProfilePayload{Profile: profile}))
}

func (client *Client) SendToPropertyInspector(ctx context.Context, payload interface{}) error {
	return client.send(NewEvent(ctx, SendToPropertyInspector, payload))
}

func (client *Client) SendToPlugin(ctx context.Context, payload interface{}) error {
	return client.send(NewEvent(ctx, SendToPlugin, payload))
}

func (client *Client) Close() error {
	err := client.c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	if err != nil {
		return err
	}
	select {
	case <-client.done:
	case <-time.After(time.Second):
	}
	return client.c.Close()
}