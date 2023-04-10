var calendarInstance1 = calendarJs( "calendar", {
    manualEditingEnabled: false,
    showDayNumberOrdinals: false,
    showHolidays: false,
    events: [{
        title: "test",
        from: new Date('2023-04-10T15:00:00'),
        to: new Date('2023-04-10T17:00:00'),
    }],
    // All your options can be set here
  } );