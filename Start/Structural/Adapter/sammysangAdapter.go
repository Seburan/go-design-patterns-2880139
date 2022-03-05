package main

type SammysangAdapter struct {
	// TODO: add a field for the SammysangTV reference
	sstv *SammysangTV
}

func (ss *SammysangAdapter) turnOn() {
	// TODO
	ss.sstv.setOnState(true);
}

func (ss *SammysangAdapter) turnOff() {
	// TODO
	ss.sstv.setOnState(false);
}

func (ss *SammysangAdapter) volumeUp() int {
	// TODO
	vol := ss.sstv.getVolume() + 1;
	ss.sstv.setVolume(vol);

	return vol;
}

func (ss *SammysangAdapter) volumeDown() int {
	// TODO
	vol := ss.sstv.getVolume() - 1;
	ss.sstv.setVolume(vol);

	return vol;
}

func (ss *SammysangAdapter) channelUp() int {
	// TODO
	channel := ss.sstv.getChannel() + 1;
	ss.sstv.setChannel(channel);

	return channel;
}

func (ss *SammysangAdapter) channelDown() int {
	// TODO
	channel := ss.sstv.getChannel() - 1;
	ss.sstv.setChannel(channel);

	return channel;
}

func (ss *SammysangAdapter) goToChannel(ch int) {
	// TODO
	ss.sstv.setChannel(ch);
}
