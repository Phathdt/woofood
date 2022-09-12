package pb

type Opt func(*composer)

type composer struct {
	evt *Event
}

func EventComposer(title string, data interface{}, opts ...Opt) *composer {
	comp := new(composer)
	comp.evt = &Event{Title: title, Data: data}

	for _, o := range opts {
		o(comp)
	}

	return comp
}

func (c *composer) Event() *Event {
	return c.evt
}

func WithSenderIdAndObject(senderId, obj string) Opt {
	return func(c *composer) {
		if c.evt.Author == nil {
			c.evt.Author = &EntityDetail{Id: senderId, Object: obj}
			return
		}
		c.evt.Author.Id = senderId
		c.evt.Author.Object = obj
	}
}

func WithReceiverIdAndObject(receiverId, obj string) Opt {
	return func(c *composer) {
		if c.evt.Receiver == nil {
			c.evt.Receiver = &EntityDetail{Id: receiverId, Object: obj}
			return
		}
		c.evt.Receiver.Id = receiverId
		c.evt.Receiver.Object = obj
	}
}
