package main

import "fmt"

func main() {
	// TODO: Create a NotificationBuilder and use it to set properties
	var bldr = newNotificationBuilder();

	// TODO: Use the builder to set some properties
	bldr.SetTitle("New notification");
	bldr.SetSubTitle("New subtitle");
	bldr.SetIcon("icon.png");
	bldr.SetImage("image.png");
	bldr.SetPriority(5);
	bldr.SetMessage("This is a basic notification with text");
	bldr.SetType("Alert Type");

	notif, err := bldr.Build();
	if err != nil {
		fmt.Println("Error creating the notification", err);
	} else {
		fmt.Printf("Notification :%+v\n", notif);
	}


	// TODO: Use the Build function to create a finished object

}
