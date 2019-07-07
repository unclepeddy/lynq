resource "googlecalendar_event" "peddy_kyle" {
  summary     = "Sit around and shoot shit"
  description = "See summary"
  location    = "Fort Asshole"

  start = "2019-07-07T13:00:00-05:00"
  end   = "2019-07-07T14:00:00-05:00"

  attendee {
    email = "pedrampejman2010@gmail.com"
  }
  attendee {
    email = "kylevonbredow@gmail.com"
  }
  attendee {
    email = "peddy@google.com"
    optional = true
  }
}
