resource "googlecalendar_event" "peddy_kyle" {
  summary     = "Sit around and shoot shit"
  description = "See summary"
  location    = "Fort Asshole"

  start = "2019-07-18T13:00:00-05:00"
  end   = "2019-07-18T14:00:00-05:00"

  recurrence = [
    "RRULE:FREQ=WEEKLY;BYDAY=TH",
  ]

  attendee {
    email = "pedrampejman2010@gmail.com"
  }
  attendee {
    email = "kylevonbredow@gmail.com"
  }

  guests_can_modify = true
}
