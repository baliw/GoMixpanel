// This package implements the Mixpanel API in Go
//
package mixpanel

import (
	"encoding/json"
	"net/http"
	"encoding/base64"
	"io/ioutil"
)

type Mixpanel struct {
	apiToken  string
}

func NewMixpanel() *Mixpanel {
	return &Mixpanel{}
}

func (self *Mixpanel) ApiToken() string {
	return self.apiToken
}

func (self *Mixpanel) SetApiToken(apiToken string) *Mixpanel {
	self.apiToken = apiToken
	return self
}

func (self *Mixpanel) NewEvent() *MixpanelEvent {
	e := NewMixpanelEvent()
	e.SetMixpanel(self)
	return e
}

func (self *Mixpanel) SendEvent(name string, properties map[string]interface{}) (success bool, err error) {
	e := self.NewEvent()
	e.SetName(name)
	e.SetProperties(properties)
	return e.Send()
}

type MixpanelEvent struct {
	mixpanel	*Mixpanel
	name    	string
	properties 	map[string]interface{}
	apiToken  	string
}

func NewMixpanelEvent() *MixpanelEvent {
	m := &MixpanelEvent{}
	m.properties = map[string]interface{}{}
	return m
}

func (self *MixpanelEvent) Mixpanel() *Mixpanel {
	return self.mixpanel
}

func (self *MixpanelEvent) SetMixpanel(mixpanel *Mixpanel) *MixpanelEvent {
	self.mixpanel = mixpanel
	return self
}

func (self *MixpanelEvent) Name() string {
	return self.name
}

func (self *MixpanelEvent) SetName(name string) *MixpanelEvent {
	self.name = name
	return self
}

func (self *MixpanelEvent) Properties() map[string]interface{} {
	return self.properties
}

func (self *MixpanelEvent) SetProperties(properties map[string]interface{}) *MixpanelEvent {
	self.properties = properties
	return self
}

func (self *MixpanelEvent) SetProperty(name string, value interface{}) *MixpanelEvent {
	self.properties[name] = value
	return self
}

func (self *MixpanelEvent) Send() (success bool, err error) {
	m := map[string]interface{}{}
	m["event"] = self.name
	
	p := map[string]interface{}{}
	for k, v := range self.properties {
	    p[k] = v
	}
	p["token"] = self.mixpanel.ApiToken()
	m["properties"] = p
	
	//TODO handle errors
	data, err := json.Marshal(m)
	if err != nil {
		return false, err
	}
	
	encodedData := base64.URLEncoding.EncodeToString(data)
	resp, err := http.Get("http://api.mixpanel.com/track?data=" + encodedData)
	if err != nil {
		return false, err
	}
	
	body, err := ioutil.ReadAll(resp.Body)
	
	return string(body) == "1", err
}