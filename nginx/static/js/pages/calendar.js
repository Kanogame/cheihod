"use strict"
import PostConnection from "../utils/post.js";

var calendarInstance1 = calendarJs( "calendar", {
    manualEditingEnabled: false,
    showDayNumberOrdinals: false,
    showHolidays: false,
    events: [{
        title: "test1",
        from: new Date('2023-04-10T15:00:00'),
        to: new Date('2023-04-10T17:00:00'),
    }],
    // All your options can be set here
  } );

  const post = new PostConnection("http://127.0.0.1/places/get/30");
  const res = await post.SendDataJson();
  console.log(res);
